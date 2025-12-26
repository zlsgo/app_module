package restapi

import (
	"errors"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model"
)

type controller struct {
	service.App
	options *Options
	Path    string
}

const defaultMaxPageSize = 1000

func (h *controller) Init(r *znet.Engine) error {
	var models *model.Stores
	err := h.DI.Resolve(&models)
	if err != nil {
		return errors.New("functional model has not been registered")
	}

	if h.options == nil {
		return errors.New("restapi options is required")
	}

	h.options.AllowMethods = normalizeAllowMethods(h.options.AllowMethods)
	h.options.AllowQueryKeys = normalizeAllowQueryKeys(h.options.AllowQueryKeys)

	if !h.options.DisableErrorHandler {
		handler := h.options.ErrorHandler
		if handler == nil {
			handler = defaultErrorHandler
		}
		r.Use(znet.RewriteErrorHandler(handler))
	}

	if h.options.Middleware != nil {
		r.Use(h.options.Middleware)
	}

	r.Any("/*", func(c *znet.Context) (any, error) {
		modelName, id, err := parseRoute(c.GetParam("*"))
		if err != nil {
			return nil, err
		}

		mod, ok := models.Get(modelName)
		if !ok {
			return nil, zerror.NotFound.Text("model not found")
		}

		method := c.Request.Method

		if h.options.ResponseHook != nil && !h.options.ResponseHook(c, modelName, id, method) {
			r.HandleNotFound(c)
			return nil, nil
		}

		if !methodAllowed(h.options.AllowMethods, method) {
			handleMethodNotAllowed(c, h.options.AllowMethods)
			return nil, nil
		}

		maxPageSize := h.options.MaxPageSize
		if maxPageSize <= 0 {
			maxPageSize = defaultMaxPageSize
		}

		if err := validateQueryParams(c, h.options); err != nil {
			return nil, err
		}

		switch method {
		case "GET":
			filter, queryFn, err := buildQuery(c, mod, models, h.options)
			if err != nil {
				return nil, err
			}
			return find(c, mod, id, filter, queryFn, maxPageSize)
		case "POST":
			if id != "" {
				return nil, zerror.InvalidInput.Text("id not allowed for POST")
			}
			return Insert(c, mod, nil)
		case "PUT", "PATCH":
			if id == "" {
				return nil, zerror.InvalidInput.Text("id required")
			}
			return UpdateById(c, mod, id, nil)
		case "DELETE":
			if id == "" {
				return nil, zerror.InvalidInput.Text("id required")
			}
			return DeleteById(c, mod, id, nil)
		default:
			r.HandleNotFound(c)
			return nil, nil
		}
	})
	return nil
}

func find(
	c *znet.Context,
	mod *model.Store,
	id string,
	filter model.Filter,
	fn func(o *model.CondOptions),
	maxPageSize int,
) (any, error) {
	if filter == nil {
		filter = model.Filter{}
	}
	switch id {
	case "":
		page, pagesize, err := model.Common.VarPages(c)
		if err != nil {
			return nil, err
		}
		if maxPageSize > 0 && pagesize > maxPageSize {
			pagesize = maxPageSize
		}
		return mod.Pages(page, pagesize, filter, func(o *model.CondOptions) {
			o.OrderBy = []model.OrderByItem{{Field: model.IDKey(), Direction: "DESC"}}

			if fn != nil {
				fn(o)
			}
		})
	case "*":
		return nil, zerror.InvalidInput.Text("全量查询不允许，请使用分页")
	default:
		filter[model.IDKey()] = id
		row, err := mod.FindOne(filter, fn)
		if err != nil {
			if errors.Is(err, model.ErrNoRecord) {
				return nil, zerror.NotFound.Text("id not found")
			}
			return nil, err
		}

		return row, nil
	}
}

func parseRoute(path string) (modelName string, id string, err error) {
	trimmed := strings.Trim(path, "/")
	if trimmed == "" {
		return "", "", zerror.InvalidInput.Text("model required")
	}
	parts := strings.Split(trimmed, "/")
	if len(parts) > 2 {
		return "", "", zerror.InvalidInput.Text("invalid path")
	}
	modelName = parts[0]
	if modelName == "" {
		return "", "", zerror.InvalidInput.Text("model required")
	}
	if len(parts) == 2 {
		id = parts[1]
		if id == "" {
			return "", "", zerror.InvalidInput.Text("id required")
		}
	}
	return modelName, id, nil
}

func methodAllowed(allow map[string]bool, method string) bool {
	if len(allow) == 0 {
		return true
	}
	return allow[strings.ToUpper(method)]
}

func normalizeAllowMethods(allow map[string]bool) map[string]bool {
	if len(allow) == 0 {
		return allow
	}
	out := make(map[string]bool, len(allow))
	for method, ok := range allow {
		m := strings.ToUpper(strings.TrimSpace(method))
		if m == "" {
			continue
		}
		if ok {
			out[m] = true
			continue
		}
		if _, exists := out[m]; !exists {
			out[m] = false
		}
	}
	return out
}

func handleMethodNotAllowed(c *znet.Context, allow map[string]bool) {
	if c == nil {
		return
	}
	allowHeader := buildAllowHeader(allow)
	if allowHeader != "" {
		c.SetHeader("Allow", allowHeader, true)
	}
	c.JSON(int32(http.StatusMethodNotAllowed), znet.ApiData{
		Code: int32(http.StatusMethodNotAllowed),
		Msg:  "method not allowed",
	})
	c.Abort()
}

func buildAllowHeader(allow map[string]bool) string {
	if len(allow) == 0 {
		return ""
	}
	methods := make([]string, 0, len(allow))
	for method, ok := range allow {
		if !ok {
			continue
		}
		m := strings.ToUpper(strings.TrimSpace(method))
		if m == "" {
			continue
		}
		methods = append(methods, m)
	}
	if len(methods) == 0 {
		return ""
	}
	sort.Strings(methods)
	return strings.Join(methods, ", ")
}

func buildQuery(c *znet.Context, mod *model.Store, models *model.Stores, opts *Options) (model.Filter, func(o *model.CondOptions), error) {
	var schema *model.Schema
	if mod != nil {
		schema = mod.Schema()
	}
	filter, err := parseFilterParam(c, schema, opts)
	if err != nil {
		return nil, nil, err
	}
	if filter == nil {
		filter = model.Filter{}
	}

	queryFn, err := buildCondOptions(c, schema, models, opts)
	if err != nil {
		return nil, nil, err
	}

	return filter, queryFn, nil
}

func buildCondOptions(c *znet.Context, schema *model.Schema, models *model.Stores, opts *Options) (func(o *model.CondOptions), error) {
	if c == nil {
		return nil, nil
	}

	fields, err := parseFieldsParam(c, schema, opts)
	if err != nil {
		return nil, err
	}
	relations, err := parseRelationsParam(c, schema, models, opts)
	if err != nil {
		return nil, err
	}
	orderBy, err := parseOrderParam(c, schema, opts)
	if err != nil {
		return nil, err
	}

	if len(fields) == 0 && len(relations) == 0 && len(orderBy) == 0 {
		return nil, nil
	}

	return func(o *model.CondOptions) {
		if len(fields) > 0 {
			o.Fields = append([]string(nil), fields...)
		}
		if len(relations) > 0 {
			o.Relations = append([]string(nil), relations...)
		}
		if len(orderBy) > 0 {
			o.OrderBy = append([]model.OrderByItem(nil), orderBy...)
		}
	}, nil
}

func parseFilterParam(c *znet.Context, schema *model.Schema, opts *Options) (model.Filter, error) {
	raw, ok := getParamValue(c, "filter")
	raw = strings.TrimSpace(raw)
	if !ok {
		return nil, nil
	}
	if raw == "" {
		return nil, zerror.InvalidInput.Text("invalid filter")
	}
	if !zjson.Valid(raw) {
		return nil, zerror.InvalidInput.Text("invalid filter")
	}
	j := zjson.Parse(raw)
	if !j.IsObject() {
		return nil, zerror.InvalidInput.Text("invalid filter")
	}

	return normalizeFilterMap(j.Map(), schema, opts)
}

func normalizeFilterMap(input ztype.Map, schema *model.Schema, opts *Options) (model.Filter, error) {
	out := model.Filter{}
	for key, raw := range input {
		k := strings.TrimSpace(key)
		if k == "" {
			continue
		}

		logicKey, isLogic := normalizeLogicKey(k)
		if isLogic {
			normalized, err := normalizeLogicFilter(raw, schema, opts)
			if err != nil {
				return nil, err
			}
			out[logicKey] = normalized
			continue
		}

		if !isSafeName(k) {
			return nil, zerror.InvalidInput.Text("invalid filter field")
		}
		if !filterFieldAllowed(opts, k) {
			return nil, zerror.InvalidInput.Text("filter field not allowed")
		}
		if !schemaFieldExists(schema, k) {
			return nil, zerror.InvalidInput.Text("invalid filter field")
		}

		if opMap, ok := raw.(map[string]any); ok {
			if err := applyFieldOperators(out, k, ztype.Map(opMap)); err != nil {
				return nil, err
			}
			continue
		}
		if opMap, ok := raw.(ztype.Map); ok {
			if err := applyFieldOperators(out, k, opMap); err != nil {
				return nil, err
			}
			continue
		}

		out[k] = raw
	}
	return out, nil
}

func normalizeLogicKey(key string) (string, bool) {
	switch strings.ToUpper(key) {
	case "$OR":
		return "$OR", true
	case "$AND":
		return "$AND", true
	default:
		return "", false
	}
}

func normalizeLogicFilter(value any, schema *model.Schema, opts *Options) ([]ztype.Map, error) {
	items := ztype.New(value).SliceValue()
	if len(items) == 0 {
		return nil, zerror.InvalidInput.Text("invalid filter")
	}

	out := make([]ztype.Map, 0, len(items))
	for _, item := range items {
		m := ztype.ToMap(item)
		if len(m) == 0 {
			return nil, zerror.InvalidInput.Text("invalid filter")
		}
		normalized, err := normalizeFilterMap(ztype.Map(m), schema, opts)
		if err != nil {
			return nil, err
		}
		out = append(out, ztype.Map(normalized))
	}
	return out, nil
}

func applyFieldOperators(out model.Filter, field string, ops ztype.Map) error {
	if len(ops) == 0 {
		return zerror.InvalidInput.Text("invalid filter")
	}

	for key, value := range ops {
		op := strings.ToLower(strings.TrimSpace(key))
		if op == "" || !strings.HasPrefix(op, "$") {
			return zerror.InvalidInput.Text("invalid filter")
		}

		switch op {
		case "$eq":
			out[field] = value
		case "$ne":
			out[field+" !="] = value
		case "$gt":
			out[field+" >"] = value
		case "$gte":
			out[field+" >="] = value
		case "$lt":
			out[field+" <"] = value
		case "$lte":
			out[field+" <="] = value
		case "$like":
			out[field+" LIKE"] = value
		case "$in":
			values := ztype.New(value).SliceValue()
			if len(values) == 0 {
				return zerror.InvalidInput.Text("invalid filter")
			}
			out[field+" IN"] = values
		case "$nin":
			values := ztype.New(value).SliceValue()
			if len(values) == 0 {
				return zerror.InvalidInput.Text("invalid filter")
			}
			out[field+" NOT IN"] = values
		case "$between":
			values := ztype.New(value).SliceValue()
			if len(values) != 2 {
				return zerror.InvalidInput.Text("invalid filter")
			}
			out[field+" BETWEEN"] = values
		case "$null":
			if !truthy(value) {
				return zerror.InvalidInput.Text("invalid filter")
			}
			out[field+" IS NULL"] = nil
		case "$notnull":
			if !truthy(value) {
				return zerror.InvalidInput.Text("invalid filter")
			}
			out[field+" IS NOT NULL"] = nil
		default:
			return zerror.InvalidInput.Text("invalid filter")
		}
	}
	return nil
}

func truthy(value any) bool {
	if value == nil {
		return true
	}
	b, ok := value.(bool)
	return ok && b
}

func parseFieldsParam(c *znet.Context, schema *model.Schema, opts *Options) ([]string, error) {
	raw, ok := getParamValue(c, "fields")
	raw = strings.TrimSpace(raw)
	if !ok {
		if opts != nil && opts.RequireFields {
			return nil, zerror.InvalidInput.Text("fields required")
		}
		if opts != nil && len(opts.DefaultFields) > 0 {
			fields := normalizeList(opts.DefaultFields)
			if len(fields) == 0 {
				return nil, zerror.InvalidInput.Text("fields required")
			}
			return validateFields(fields, schema, opts)
		}
		return nil, nil
	}
	if raw == "" {
		return nil, zerror.InvalidInput.Text("invalid field")
	}
	fields := splitList(raw)
	if len(fields) == 0 {
		return nil, zerror.InvalidInput.Text("invalid field")
	}
	return validateFields(fields, schema, opts)
}

func parseRelationsParam(c *znet.Context, schema *model.Schema, models *model.Stores, opts *Options) ([]string, error) {
	withRaw, withOK := getParamValue(c, "with")
	relRaw, relOK := getParamValue(c, "relations")
	if withOK && relOK {
		return nil, zerror.InvalidInput.Text("invalid relation")
	}
	raw := ""
	if withOK {
		raw = withRaw
	} else if relOK {
		raw = relRaw
	} else {
		return nil, nil
	}
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, zerror.InvalidInput.Text("invalid relation")
	}
	relations := splitList(raw)
	if len(relations) == 0 {
		return nil, zerror.InvalidInput.Text("invalid relation")
	}
	for _, rel := range relations {
		if !isSafePath(rel) {
			return nil, zerror.InvalidInput.Text("invalid relation")
		}
		if !relationAllowed(opts, rel) {
			return nil, zerror.InvalidInput.Text("relation not allowed")
		}
		if err := validateRelationPath(schema, models, rel); err != nil {
			return nil, err
		}
	}
	return relations, nil
}

func parseOrderParam(c *znet.Context, schema *model.Schema, opts *Options) ([]model.OrderByItem, error) {
	raw, ok := getParamValue(c, "order")
	raw = strings.TrimSpace(raw)
	if !ok {
		if opts != nil && len(opts.DefaultOrder) > 0 {
			orderBy := append([]model.OrderByItem(nil), opts.DefaultOrder...)
			for _, item := range orderBy {
				if !isSafePath(item.Field) {
					return nil, zerror.InvalidInput.Text("invalid order field")
				}
				if !orderFieldAllowed(opts, item.Field) {
					return nil, zerror.InvalidInput.Text("order field not allowed")
				}
				if !schemaFieldExists(schema, item.Field) {
					return nil, zerror.InvalidInput.Text("invalid order field")
				}
			}
			return orderBy, nil
		}
		return nil, nil
	}
	if raw == "" {
		return nil, zerror.InvalidInput.Text("invalid order")
	}

	items := splitList(raw)
	if len(items) == 0 {
		return nil, zerror.InvalidInput.Text("invalid order")
	}

	orderBy := make([]model.OrderByItem, 0, len(items))
	for _, item := range items {
		field, dir, err := parseOrderItem(item)
		if err != nil {
			return nil, err
		}
		if !isSafePath(field) {
			return nil, zerror.InvalidInput.Text("invalid order field")
		}
		if !orderFieldAllowed(opts, field) {
			return nil, zerror.InvalidInput.Text("order field not allowed")
		}
		if !schemaFieldExists(schema, field) {
			return nil, zerror.InvalidInput.Text("invalid order field")
		}
		orderBy = append(orderBy, model.OrderByItem{
			Field:     field,
			Direction: dir,
		})
	}
	return orderBy, nil
}

func parseOrderItem(item string) (field string, direction string, err error) {
	part := strings.TrimSpace(item)
	if part == "" {
		return "", "", zerror.InvalidInput.Text("invalid order")
	}

	direction = "ASC"
	switch part[0] {
	case '-':
		direction = "DESC"
		part = strings.TrimSpace(part[1:])
	case '+':
		part = strings.TrimSpace(part[1:])
	}

	if strings.Contains(part, ":") {
		parts := strings.SplitN(part, ":", 2)
		part = strings.TrimSpace(parts[0])
		if part == "" {
			return "", "", zerror.InvalidInput.Text("invalid order")
		}
		dir := strings.ToUpper(strings.TrimSpace(parts[1]))
		if dir == "" {
			dir = "ASC"
		}
		if dir != "ASC" && dir != "DESC" {
			return "", "", zerror.InvalidInput.Text("invalid order")
		}
		direction = dir
	}

	if strings.ContainsAny(part, " \t") {
		return "", "", zerror.InvalidInput.Text("invalid order")
	}

	if part == "" {
		return "", "", zerror.InvalidInput.Text("invalid order")
	}

	return part, direction, nil
}

func getParamValue(c *znet.Context, key string) (string, bool) {
	if c == nil {
		return "", false
	}
	if value, ok := c.GetPostForm(key); ok {
		return value, true
	}
	if value, ok := c.GetQuery(key); ok {
		return value, true
	}
	return "", false
}

func normalizeAllowQueryKeys(allow map[string]bool) map[string]bool {
	if len(allow) == 0 {
		return allow
	}
	out := make(map[string]bool, len(allow))
	for key, ok := range allow {
		if !ok {
			continue
		}
		k := strings.TrimSpace(key)
		if k == "" {
			continue
		}
		out[k] = true
	}
	return out
}

func validateQueryParams(c *znet.Context, opts *Options) error {
	if c == nil || opts == nil || !opts.RejectUnknownQuery {
		return nil
	}
	values := c.GetAllQuery()
	if len(values) == 0 {
		return nil
	}
	allowed := defaultAllowedQueryKeys()
	if len(opts.AllowQueryKeys) > 0 {
		for key, ok := range opts.AllowQueryKeys {
			if !ok {
				continue
			}
			k := strings.TrimSpace(key)
			if k == "" {
				continue
			}
			allowed[k] = struct{}{}
		}
	}
	invalidErr := zerror.InvalidInput.Text("invalid query")
	for key, vals := range values {
		if key == "" {
			return invalidErr
		}
		if _, ok := allowed[key]; !ok {
			return invalidErr
		}
		if len(vals) != 1 {
			return invalidErr
		}
		value := strings.TrimSpace(vals[0])
		if value == "" {
			return invalidErr
		}
		if key == "page" || key == "pagesize" {
			num, err := strconv.Atoi(value)
			if err != nil || num < 1 {
				return invalidErr
			}
		}
	}
	return nil
}

func defaultAllowedQueryKeys() map[string]struct{} {
	return map[string]struct{}{
		"fields":    {},
		"with":      {},
		"relations": {},
		"order":     {},
		"filter":    {},
		"page":      {},
		"pagesize":  {},
	}
}

func splitList(value string) []string {
	if value == "" {
		return nil
	}
	parts := strings.Split(value, ",")
	if len(parts) == 0 {
		return nil
	}
	seen := make(map[string]struct{}, len(parts))
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		if _, ok := seen[part]; ok {
			continue
		}
		seen[part] = struct{}{}
		out = append(out, part)
	}
	return out
}

func isSafePath(value string) bool {
	if value == "" {
		return false
	}
	parts := strings.Split(value, ".")
	for _, part := range parts {
		if !isSafeName(part) {
			return false
		}
	}
	return true
}

func isSafeName(value string) bool {
	if value == "" {
		return false
	}
	for _, c := range value {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_' {
			continue
		}
		return false
	}
	return true
}

func normalizeList(items []string) []string {
	if len(items) == 0 {
		return nil
	}
	seen := make(map[string]struct{}, len(items))
	out := make([]string, 0, len(items))
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}
	return out
}

func validateFields(fields []string, schema *model.Schema, opts *Options) ([]string, error) {
	for _, field := range fields {
		if field == "*" {
			if !fieldAllowed(opts, field) {
				return nil, zerror.InvalidInput.Text("field not allowed")
			}
			continue
		}
		if !isSafePath(field) {
			return nil, zerror.InvalidInput.Text("invalid field")
		}
		if !fieldAllowed(opts, field) {
			return nil, zerror.InvalidInput.Text("field not allowed")
		}
		if !schemaFieldExists(schema, field) {
			return nil, zerror.InvalidInput.Text("invalid field")
		}
	}
	return fields, nil
}

func schemaFieldExists(schema *model.Schema, field string) bool {
	if schema == nil {
		return true
	}
	_, ok := schema.GetField(field)
	return ok
}

func validateRelationPath(schema *model.Schema, models *model.Stores, path string) error {
	if schema == nil {
		return nil
	}
	if path == "" {
		return zerror.InvalidInput.Text("invalid relation")
	}
	parts := strings.Split(path, ".")
	cur := schema
	for i := 0; i < len(parts); i++ {
		part := strings.TrimSpace(parts[i])
		if part == "" {
			return zerror.InvalidInput.Text("invalid relation")
		}
		rel, ok := cur.GetDefine().Relations[part]
		if ok {
			if i == len(parts)-1 {
				return nil
			}
			if models == nil {
				return zerror.InvalidInput.Text("invalid relation")
			}
			next, ok := models.Get(rel.Schema)
			if !ok || next == nil {
				return zerror.InvalidInput.Text("invalid relation")
			}
			cur = next.Schema()
			if cur == nil {
				return zerror.InvalidInput.Text("invalid relation")
			}
			continue
		}
		if i == 0 {
			return zerror.InvalidInput.Text("invalid relation")
		}
		if !schemaFieldExists(cur, part) {
			return zerror.InvalidInput.Text("invalid relation")
		}
		if i != len(parts)-1 {
			return zerror.InvalidInput.Text("invalid relation")
		}
		return nil
	}
	return zerror.InvalidInput.Text("invalid relation")
}

func fieldAllowed(opts *Options, field string) bool {
	if opts == nil || len(opts.AllowFields) == 0 {
		return true
	}
	return opts.AllowFields[field]
}

func filterFieldAllowed(opts *Options, field string) bool {
	if opts == nil {
		return true
	}
	if len(opts.AllowFilterFields) > 0 {
		return opts.AllowFilterFields[field]
	}
	if len(opts.AllowFields) > 0 {
		return opts.AllowFields[field]
	}
	return true
}

func orderFieldAllowed(opts *Options, field string) bool {
	if opts == nil {
		return true
	}
	if len(opts.AllowOrderFields) > 0 {
		return opts.AllowOrderFields[field]
	}
	if len(opts.AllowFields) > 0 {
		return opts.AllowFields[field]
	}
	return true
}

func relationAllowed(opts *Options, relation string) bool {
	if opts == nil || len(opts.AllowRelations) == 0 {
		return true
	}
	if opts.AllowRelations[relation] {
		return true
	}
	root := relation
	if idx := strings.Index(relation, "."); idx != -1 {
		root = relation[:idx]
	}
	return opts.AllowRelations[root]
}

func defaultErrorHandler(c *znet.Context, err error) {
	if err == nil {
		return
	}
	code := statusCodeFromError(err)
	c.JSON(int32(code), znet.ApiData{
		Code: int32(code),
		Msg:  err.Error(),
	})
}

func statusCodeFromError(err error) int {
	switch zerror.GetTag(err) {
	case zerror.InvalidInput:
		return http.StatusBadRequest
	case zerror.NotFound:
		return http.StatusNotFound
	case zerror.Unauthorized:
		return http.StatusUnauthorized
	case zerror.PermissionDenied:
		return http.StatusForbidden
	case zerror.Cancelled:
		return http.StatusRequestTimeout
	case zerror.Internal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
