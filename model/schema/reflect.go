package schema

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/zreflect"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb/builder"
	"github.com/zlsgo/zdb/schema"
)

type Meta struct{}

var (
	metaType    = reflect.TypeOf(Meta{})
	fieldsCache = zcache.NewFast()
	schemaCache = zcache.NewFast()
)

func FieldsFromStruct[T any]() Fields {
	var t T
	return fieldsFromType(zreflect.TypeOf(t))
}

func FieldsFromType(t reflect.Type) Fields {
	if t == nil {
		return nil
	}
	return fieldsFromType(zreflect.TypeOf(t))
}

func NewFromStruct[T any](name string, tableName ...string) Schema {
	var t T
	return NewFromStructType(name, zreflect.TypeOf(t), tableName...)
}

func NewFromStructType(name string, t reflect.Type, tableName ...string) Schema {
	if t != nil {
		t = zreflect.TypeOf(t)
	}
	if t == nil {
		return buildSchemaFromType(name, t, tableName...)
	}

	cacheKey := schemaCacheKey(t, name, tableName)
	cached, _ := schemaCache.ProvideGet(cacheKey, func() (interface{}, bool) {
		s := buildSchemaFromType(name, t, tableName...)
		return s, true
	})
	if cached == nil {
		return Schema{}
	}
	if s, ok := cached.(Schema); ok {
		return cloneSchema(s)
	}
	return Schema{}
}

func NewFromStructValue(name string, v any, tableName ...string) Schema {
	if v == nil {
		return NewFromStructType(name, nil, tableName...)
	}
	return NewFromStructType(name, zreflect.TypeOf(v), tableName...)
}

func buildSchemaFromType(name string, t reflect.Type, tableName ...string) Schema {
	hasName := name != ""
	hasTable := false

	var s Schema
	if len(tableName) > 0 {
		hasTable = tableName[0] != ""
		s = New(name, tableName...)
	} else {
		s = New(name)
	}
	s.Fields = Fields{}
	s.Relations = Relations{}

	tt := unwrapStructType(t)
	if tt == nil || tt.Kind() != reflect.Struct {
		if s.Fields == nil {
			s.Fields = Fields{}
		}
		return s
	}

	rootName := schemaNameFromType(tt)
	parseSchemaFromType(tt, &s, rootName, !hasName, !hasTable, true)

	if s.Name == "" {
		s.Name = rootName
	}
	if s.Table.Name == "" {
		s.Table.Name = s.Name
	}

	if s.Fields == nil {
		s.Fields = Fields{}
	}
	if s.Relations == nil {
		s.Relations = Relations{}
	}
	return s
}

func fieldsFromType(t reflect.Type) Fields {
	t = unwrapStructType(t)
	if t == nil || t.Kind() != reflect.Struct {
		return nil
	}

	cacheKey := fieldsCacheKey(t)
	cached, _ := fieldsCache.ProvideGet(cacheKey, func() (interface{}, bool) {
		return buildFieldsFromType(t), true
	})
	if cached == nil {
		return nil
	}
	if fields, ok := cached.(Fields); ok {
		return cloneFields(fields)
	}
	return nil
}

func buildFieldsFromType(t reflect.Type) Fields {
	fields := make(Fields)
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if isMetaField(sf) || relationTag(sf) != "" {
			continue
		}

		if !sf.IsExported() {
			continue
		}

		if sf.Anonymous {
			embeddedType := unwrapStructType(sf.Type)
			if embeddedType == nil || embeddedType == t {
				continue
			}
			embedded := fieldsFromType(embeddedType)
			for k, v := range embedded {
				if _, exists := fields[k]; !exists {
					fields[k] = v
				}
			}
			continue
		}

		fieldTag := sf.Tag.Get("field")
		if fieldTag == "-" {
			continue
		}

		name := getFieldName(sf)
		if name == "-" || name == "" {
			continue
		}

		ft := sf.Type
		nullable := false
		if ft.Kind() == reflect.Ptr {
			ft = ft.Elem()
			nullable = true
		}

		field := Field{
			Type:     goTypeToSchemaType(ft),
			Label:    name,
			Nullable: nullable,
		}

		parseFieldTag(fieldTag, &field)
		fields[name] = field
	}
	return fields
}

func parseSchemaFromType(t reflect.Type, s *Schema, rootName string, allowName, allowTable, allowMeta bool) {
	t = unwrapStructType(t)
	if t == nil || t.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)

		if allowMeta && isMetaField(sf) {
			applyMetaTags(s, sf, allowName, allowTable)
			continue
		}

		relTag := relationTag(sf)
		if relTag != "" {
			relName, rel := parseRelationTag(relTag, sf, rootName)
			if relName == "" || relName == "-" {
				continue
			}
			if rel.Label == "" {
				rel.Label = relName
			}
			if s.Relations == nil {
				s.Relations = Relations{}
			}
			if _, exists := s.Relations[relName]; !exists {
				s.Relations[relName] = rel
			}
			continue
		}

		if !sf.IsExported() {
			continue
		}

		if sf.Anonymous {
			parseSchemaFromType(sf.Type, s, rootName, false, false, false)
			continue
		}

		fieldTag := sf.Tag.Get("field")
		if fieldTag == "-" {
			continue
		}

		name := getFieldName(sf)
		if name == "-" || name == "" {
			continue
		}

		ft := sf.Type
		nullable := false
		if ft.Kind() == reflect.Ptr {
			ft = ft.Elem()
			nullable = true
		}

		field := Field{
			Type:     goTypeToSchemaType(ft),
			Label:    name,
			Nullable: nullable,
		}

		parseFieldTag(fieldTag, &field)
		if s.Fields == nil {
			s.Fields = Fields{}
		}
		if _, exists := s.Fields[name]; !exists {
			s.Fields[name] = field
		}
	}
}

func getFieldName(sf reflect.StructField) string {
	jsonTag := sf.Tag.Get("json")
	if jsonTag != "" {
		parts := strings.Split(jsonTag, ",")
		if parts[0] != "" {
			return parts[0]
		}
	}
	return camelToSnake(sf.Name)
}

func camelToSnake(name string) string {
	if name == "" {
		return ""
	}
	var b strings.Builder
	b.Grow(len(name) + len(name)/2)
	for i := 0; i < len(name); i++ {
		c := name[i]
		if c == '_' {
			b.WriteByte(c)
			continue
		}
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				prev := name[i-1]
				nextLower := i+1 < len(name) && name[i+1] >= 'a' && name[i+1] <= 'z'
				prevLower := prev >= 'a' && prev <= 'z'
				prevDigit := prev >= '0' && prev <= '9'
				prevUpper := prev >= 'A' && prev <= 'Z'
				if prevLower || prevDigit || (prevUpper && nextLower) {
					b.WriteByte('_')
				}
			}
			b.WriteByte(c + 'a' - 'A')
			continue
		}
		b.WriteByte(c)
	}
	return b.String()
}

func goTypeToSchemaType(t reflect.Type) schema.DataType {
	if t == reflect.TypeOf(time.Time{}) {
		return Time
	}

	switch t.Kind() {
	case reflect.Bool:
		return Bool
	case reflect.Int:
		return Int
	case reflect.Int8:
		return Int8
	case reflect.Int16:
		return Int16
	case reflect.Int32:
		return Int32
	case reflect.Int64:
		return Int64
	case reflect.Uint:
		return Uint
	case reflect.Uint8:
		return Uint8
	case reflect.Uint16:
		return Uint16
	case reflect.Uint32:
		return Uint32
	case reflect.Uint64:
		return Uint64
	case reflect.Float32, reflect.Float64:
		return Float
	case reflect.String:
		return String
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			return Bytes
		}
		return JSON
	case reflect.Map, reflect.Struct, reflect.Interface:
		return JSON
	default:
		return String
	}
}

func parseFieldTag(tag string, f *Field) {
	if tag == "" {
		return
	}

	for _, part := range strings.Split(tag, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		kv := strings.SplitN(part, ":", 2)
		key := strings.TrimSpace(kv[0])
		val := ""
		if len(kv) > 1 {
			val = strings.TrimSpace(kv[1])
		}

		switch key {
		case "type":
			f.Type = schema.DataType(val)
		case "size":
			f.Size, _ = strconv.ParseUint(val, 10, 64)
		case "label":
			f.Label = val
		case "default":
			f.Default = val
		case "comment":
			f.Comment = val
		case "nullable":
			f.Nullable = parseBoolDefaultTrue(val)
		case "unique":
			f.Unique = parseBoolOrString(val)
		case "index":
			f.Index = parseBoolOrString(val)
		case "readonly":
			f.Options.ReadOnly = parseBoolDefaultTrue(val)
		case "crypt":
			f.Options.Crypt = val
		case "array":
			f.Options.IsArray = parseBoolDefaultTrue(val)
		case "format":
			f.Options.FormatTime = val
		case "enum":
			f.Options.Enum = parseEnumList(val)
		case "valid", "validate":
			f.Validations = append(f.Validations, parseValidationList(val)...)
		case "disable_migration":
			f.Options.DisableMigration = parseBoolDefaultTrue(val)
		}
	}
}

func unwrapStructType(t reflect.Type) reflect.Type {
	if t == nil {
		return nil
	}
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func relationTag(sf reflect.StructField) string {
	tag := sf.Tag.Get("relation")
	if tag == "" {
		tag = sf.Tag.Get("rel")
	}
	if tag == "-" {
		return ""
	}
	return tag
}

func isMetaField(sf reflect.StructField) bool {
	if sf.Name == "_" {
		t := sf.Type
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		return t.Kind() == reflect.Struct
	}
	if !sf.Anonymous {
		return false
	}
	t := sf.Type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t == metaType
}

type tagKV struct {
	key string
	val string
}

func parseTagPairs(tag string) []tagKV {
	parts := strings.Split(tag, ",")
	pairs := make([]tagKV, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		kv := strings.SplitN(part, ":", 2)
		key := strings.TrimSpace(kv[0])
		val := ""
		if len(kv) > 1 {
			val = strings.TrimSpace(kv[1])
		}
		pairs = append(pairs, tagKV{key: key, val: val})
	}
	return pairs
}

func applyMetaTags(s *Schema, sf reflect.StructField, allowName, allowTable bool) {
	tag := sf.Tag
	if schemaTag := tag.Get("schema"); schemaTag != "" {
		for _, kv := range parseTagPairs(schemaTag) {
			applyMetaKey(s, kv.key, kv.val, allowName, allowTable)
		}
	}

	if nameTag := tag.Get("name"); nameTag != "" && allowName {
		s.Name = nameTag
	}
	if tableTag := tag.Get("table"); tableTag != "" && allowTable {
		s.Table.Name = tableTag
	}
	if commentTag := tag.Get("comment"); commentTag != "" {
		s.Table.Comment = commentTag
	}
	if optionsTag := tag.Get("options"); optionsTag != "" {
		applyOptionsList(s, optionsTag)
	}

	for _, key := range []string{
		"timestamps",
		"soft_deletes",
		"soft_delete_is_time",
		"crypt_id",
		"disabled_migrator",
		"low_fields",
		"fields_sort",
		"crypt_salt",
		"crypt_len",
	} {
		if val, ok := tag.Lookup(key); ok {
			applyOptionKey(s, key, val)
		}
	}
}

func applyMetaKey(s *Schema, key, val string, allowName, allowTable bool) {
	key = normalizeTagKey(key)
	switch key {
	case "name":
		if allowName {
			s.Name = val
		}
	case "table":
		if allowTable {
			s.Table.Name = val
		}
	case "comment":
		s.Table.Comment = val
	case "options":
		applyOptionsList(s, val)
	default:
		applyOptionKey(s, key, val)
	}
}

func applyOptionsList(s *Schema, val string) {
	items := splitOptionList(val)
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		kv := strings.SplitN(item, "=", 2)
		key := strings.TrimSpace(kv[0])
		value := ""
		if len(kv) > 1 {
			value = strings.TrimSpace(kv[1])
		}
		applyOptionKey(s, key, value)
	}
}

func applyOptionKey(s *Schema, key, val string) {
	key = normalizeTagKey(key)
	switch key {
	case "timestamps":
		setOptionBool(&s.Options.Timestamps, val)
	case "soft_deletes":
		setOptionBool(&s.Options.SoftDeletes, val)
	case "soft_delete_is_time":
		setOptionBool(&s.Options.SoftDeleteIsTime, val)
	case "crypt_id":
		setOptionBool(&s.Options.CryptID, val)
	case "disabled_migrator":
		setOptionBool(&s.Options.DisabledMigrator, val)
	case "crypt_salt":
		s.Options.Salt = val
	case "crypt_len":
		if val == "" {
			return
		}
		if i, err := strconv.Atoi(val); err == nil {
			s.Options.CryptLen = i
		}
	case "low_fields":
		s.Options.LowFields = splitOptionList(val)
	case "fields_sort":
		s.Options.FieldsSort = splitOptionList(val)
	}
}

func parseRelationTag(tag string, sf reflect.StructField, rootName string) (string, Relation) {
	rel := Relation{}
	relName := getFieldName(sf)
	relatedName := schemaNameFromType(sf.Type)

	for _, kv := range parseTagPairs(tag) {
		key := normalizeTagKey(kv.key)
		val := kv.val
		switch key {
		case "name":
			if val != "" {
				relName = val
			}
		case "type":
			rel.Type = RelationType(val)
		case "schema":
			rel.Schema = val
		case "foreign", "foreign_key":
			rel.ForeignKey = splitList(val)
		case "schema_key":
			rel.SchemaKey = splitList(val)
		case "fields":
			rel.Fields = splitList(val)
		case "nullable":
			rel.Nullable = parseBoolDefaultTrue(val)
		case "pivot_table":
			rel.PivotTable = val
		case "pivot_foreign":
			rel.PivotKeys.Foreign = splitList(val)
		case "pivot_related":
			rel.PivotKeys.Related = splitList(val)
		case "pivot_fields":
			rel.PivotFields = splitList(val)
		case "cascade":
			rel.Cascade = val
		case "cascade_type":
			rel.CascadeType = CascadeType(strings.ToUpper(val))
		case "inverse":
			rel.Inverse = val
		case "comment":
			rel.Comment = val
		case "label":
			rel.Label = val
		}
	}

	if rel.Schema == "" {
		rel.Schema = relatedName
	}
	if rel.Type == "" {
		rel.Type = inferRelationType(sf.Type)
	}
	applyRelationDefaults(&rel, relName, rootName, relatedName)

	return relName, rel
}

func inferRelationType(t reflect.Type) RelationType {
	if t == nil {
		return RelationSingle
	}
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		return RelationMany
	}
	return RelationSingle
}

func applyRelationDefaults(rel *Relation, relName, rootName, relatedName string) {
	idKey := builder.IDKey
	switch rel.Type {
	case RelationSingle, RelationSingleMerge:
		if len(rel.ForeignKey) == 0 && relName != "" {
			rel.ForeignKey = []string{relName + "_" + idKey}
		}
		if len(rel.SchemaKey) == 0 {
			rel.SchemaKey = []string{idKey}
		}
	case RelationMany:
		if len(rel.ForeignKey) == 0 {
			rel.ForeignKey = []string{idKey}
		}
		if len(rel.SchemaKey) == 0 && rootName != "" {
			rel.SchemaKey = []string{rootName + "_" + idKey}
		}
	case RelationManyToMany:
		if len(rel.ForeignKey) == 0 {
			rel.ForeignKey = []string{idKey}
		}
		if len(rel.SchemaKey) == 0 {
			rel.SchemaKey = []string{idKey}
		}
		if len(rel.PivotKeys.Foreign) == 0 && rootName != "" {
			rel.PivotKeys.Foreign = []string{rootName + "_" + idKey}
		}
		if len(rel.PivotKeys.Related) == 0 && relatedName != "" {
			rel.PivotKeys.Related = []string{relatedName + "_" + idKey}
		}
	}
}

func schemaNameFromType(t reflect.Type) string {
	for t != nil {
		switch t.Kind() {
		case reflect.Ptr, reflect.Slice, reflect.Array:
			t = t.Elem()
			continue
		case reflect.Struct:
			name := t.Name()
			if name == "" {
				return ""
			}
			return camelToSnake(name)
		default:
			return ""
		}
	}
	return ""
}

func normalizeTagKey(key string) string {
	key = strings.TrimSpace(strings.ToLower(key))
	return strings.ReplaceAll(key, "-", "_")
}

func splitList(val string) []string {
	return splitByDelims(val, []rune{'|', ';'})
}

func splitOptionList(val string) []string {
	return splitByDelims(val, []rune{'|', ';', ','})
}

func splitByDelims(val string, delims []rune) []string {
	if val == "" {
		return nil
	}
	items := strings.FieldsFunc(val, func(r rune) bool {
		for _, d := range delims {
			if r == d {
				return true
			}
		}
		return false
	})
	out := make([]string, 0, len(items))
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		out = append(out, item)
	}
	return out
}

func parseBoolValue(val string) (bool, bool) {
	if val == "" {
		return true, true
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		return false, false
	}
	return b, true
}

func parseBoolDefaultTrue(val string) bool {
	b, ok := parseBoolValue(val)
	if !ok {
		return true
	}
	return b
}

func parseBoolOrString(val string) interface{} {
	if val == "" {
		return true
	}
	if b, err := strconv.ParseBool(val); err == nil {
		return b
	}
	return val
}

func setOptionBool(dest **bool, val string) {
	b, ok := parseBoolValue(val)
	if !ok {
		return
	}
	*dest = &b
}

func parseEnumList(val string) []FieldEnum {
	items := splitList(val)
	out := make([]FieldEnum, 0, len(items))
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		keyVal := strings.SplitN(item, "=", 2)
		value := strings.TrimSpace(keyVal[0])
		label := value
		if len(keyVal) > 1 {
			label = strings.TrimSpace(keyVal[1])
		} else if strings.Contains(item, ":") {
			parts := strings.SplitN(item, ":", 2)
			value = strings.TrimSpace(parts[0])
			label = strings.TrimSpace(parts[1])
		}
		if value == "" {
			continue
		}
		out = append(out, FieldEnum{Value: value, Label: label})
	}
	return out
}

func parseValidationList(val string) []Validations {
	items := splitList(val)
	out := make([]Validations, 0, len(items))
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}

		message := ""
		methodPart := item
		if strings.Contains(item, "@") {
			parts := strings.SplitN(item, "@", 2)
			methodPart = strings.TrimSpace(parts[0])
			message = strings.TrimSpace(parts[1])
		}

		method := methodPart
		arg := ""
		if strings.Contains(methodPart, "=") {
			parts := strings.SplitN(methodPart, "=", 2)
			method = strings.TrimSpace(parts[0])
			arg = strings.TrimSpace(parts[1])
		}

		if method == "" {
			continue
		}

		v := Validations{Method: method}
		if arg != "" {
			v.Args = arg
		}
		if message != "" {
			v.Message = message
		}
		out = append(out, v)
	}
	return out
}

func fieldsCacheKey(t reflect.Type) string {
	key := typeCacheKey(t)
	if key == "" {
		return ""
	}
	return "fields::" + key
}

func schemaCacheKey(t reflect.Type, name string, tableName []string) string {
	key := typeCacheKey(t)
	if key == "" {
		return ""
	}
	table := ""
	if len(tableName) > 0 {
		table = tableName[0]
	}
	return "schema::" + key + "::name=" + name + "::table=" + table + "::tableSet=" + strconv.Itoa(len(tableName))
}

func typeCacheKey(t reflect.Type) string {
	t = unwrapStructType(t)
	if t == nil {
		return ""
	}
	pkg := t.PkgPath()
	if pkg == "" {
		return t.String()
	}
	return pkg + "::" + t.String()
}

func cloneSchema(s Schema) Schema {
	out := s
	out.Options = cloneOptions(s.Options)
	out.Fields = cloneFields(s.Fields)
	out.Relations = cloneRelations(s.Relations)
	out.Extend = cloneMap(s.Extend)
	out.Values = cloneMaps(s.Values)
	return out
}

func cloneOptions(o Options) Options {
	out := o
	if o.DisabledMigrator != nil {
		v := *o.DisabledMigrator
		out.DisabledMigrator = &v
	}
	if o.SoftDeletes != nil {
		v := *o.SoftDeletes
		out.SoftDeletes = &v
	}
	if o.SoftDeleteIsTime != nil {
		v := *o.SoftDeleteIsTime
		out.SoftDeleteIsTime = &v
	}
	if o.Timestamps != nil {
		v := *o.Timestamps
		out.Timestamps = &v
	}
	if o.CryptID != nil {
		v := *o.CryptID
		out.CryptID = &v
	}
	if o.LowFields != nil {
		out.LowFields = append([]string(nil), o.LowFields...)
	}
	if o.FieldsSort != nil {
		out.FieldsSort = append([]string(nil), o.FieldsSort...)
	}
	return out
}

func cloneFields(fields Fields) Fields {
	if fields == nil {
		return nil
	}
	out := make(Fields, len(fields))
	for k, v := range fields {
		out[k] = cloneField(v)
	}
	return out
}

func cloneField(f Field) Field {
	out := f
	if f.Validations != nil {
		out.Validations = append([]Validations(nil), f.Validations...)
	}
	if f.Before != nil {
		out.Before = append([]string(nil), f.Before...)
	}
	if f.After != nil {
		out.After = append([]string(nil), f.After...)
	}
	if f.Options.Enum != nil {
		out.Options.Enum = append([]FieldEnum(nil), f.Options.Enum...)
	}
	return out
}

func cloneRelations(relations Relations) Relations {
	if relations == nil {
		return nil
	}
	out := make(Relations, len(relations))
	for k, v := range relations {
		out[k] = cloneRelation(v)
	}
	return out
}

func cloneRelation(r Relation) Relation {
	out := r
	out.Filter = cloneMap(r.Filter)
	out.SchemaKey = append([]string(nil), r.SchemaKey...)
	out.ForeignKey = append([]string(nil), r.ForeignKey...)
	out.Keys = append([]string(nil), r.Keys...)
	out.ForeignKeys = append([]string(nil), r.ForeignKeys...)
	out.Fields = append([]string(nil), r.Fields...)
	out.PivotKeys = PivotKeys{
		Foreign: append([]string(nil), r.PivotKeys.Foreign...),
		Related: append([]string(nil), r.PivotKeys.Related...),
	}
	out.PivotFields = append([]string(nil), r.PivotFields...)
	out.PivotFilter = cloneMap(r.PivotFilter)
	out.Options = cloneMap(r.Options)
	return out
}

func cloneMap(in ztype.Map) ztype.Map {
	if in == nil {
		return nil
	}
	out := make(ztype.Map, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

func cloneMaps(in ztype.Maps) ztype.Maps {
	if in == nil {
		return nil
	}
	out := make(ztype.Maps, len(in))
	for i, m := range in {
		out[i] = cloneMap(m)
	}
	return out
}
