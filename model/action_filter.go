package model

import (
	"reflect"
	"strings"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
)

// getFilter 将各种类型的过滤器转换为统一的 ztype.Map 格式
// 并自动处理软删除字段过滤
func getFilter(m *Schema, filter QueryFilter) (filterMap ztype.Map) {
	if filter == nil {
		filterMap = ztype.Map{}
	} else {
		filterMap = filter.ToMap()
	}

	// 深拷贝：后续的白名单剔除与 ID 解密等均为原地修改，
	// 若只做浅拷贝，嵌套分组（$OR/$AND）仍会与调用方共享并污染其原始过滤器。
	filterMap = cloneFilterMap(filterMap)

	sanitizeFilterFields(m, filterMap, 0)

	if *m.define.Options.SoftDeletes {
		if *m.define.Options.SoftDeleteIsTime {
			// 显式 `deleted_at = nil`（与 insertData 时间型软删的写入写法一致）会被解析为
			// `deleted_at = NULL`（恒不成立），统一规范为 IS NULL。
			normalizeDeletedAtNull(filterMap, 0)
		}
		if !hasFieldInFilter(filterMap, DeletedAtKey) {
			if *m.define.Options.SoftDeleteIsTime {
				// 时间型软删除：未删除行的 deleted_at 为 NULL。
				// 注意不能注入 filterMap[DeletedAtKey] = nil，那会被解析为
				// `deleted_at = NULL`（恒不成立，查询/更新/软删全部匹配 0 行）。
				filterMap[DeletedAtKey+" IS NULL"] = true
			} else {
				filterMap[DeletedAtKey] = 0
			}
		}
	}

	return
}

// sanitizeFilterFields 过滤无效字段：排除不在模型定义中的字段（含嵌套 $OR/$AND 分组）。
func sanitizeFilterFields(m *Schema, filterMap ztype.Map, depth int) {
	if len(filterMap) == 0 || depth >= maxParseDepth {
		return
	}
	for key := range filterMap {
		k := zstring.TrimSpace(key)
		if k == "" {
			continue
		}
		upperKey := strings.ToUpper(k)
		isGroup := upperKey == placeHolderOR || upperKey == placeHolderAND
		if strings.Contains(k, placeHolder) && !isGroup {
			// 运行时占位条件（如 Filter.Cond 注入的 $N 回调），不做字段校验
			continue
		}
		if isGroup {
			if nv, keep := sanitizeFilterGroupValue(m, filterMap[key], depth+1); keep {
				filterMap[key] = nv
			} else {
				delete(filterMap, key)
			}
			continue
		}
		if strings.Contains(k, ".") {
			// 关系/已限定表名前缀的列，交由下游处理
			continue
		}
		if !isAllowedFilterField(m, splitFilterFieldName(k)) {
			delete(filterMap, key)
		}
	}
}

// isAllowedFilterField 判断字段是否允许出现在当前模型的过滤条件中
func isAllowedFilterField(m *Schema, fieldName string) bool {
	if fieldName == DeletedAtKey {
		if *m.define.Options.SoftDeletes {
			// 框架软删除：deleted_at 为内置列（被移出 fullFields），显式条件原样保留。
			return true
		}
		// 未启用框架软删除时 deleted_at 只是普通列（例如业务自行以 ""/时间戳
		// 手写软删标记的表）：交给下方字段白名单校验，schema 声明了该列才保留
		// （查询 `deleted_at = ''` 等条件才会真正生效）。
	}
	if m.fullFieldsMap != nil {
		_, ok := m.fullFieldsMap[fieldName]
		return ok
	}
	if len(m.fullFields) > 0 {
		return zarray.Contains(m.fullFields, fieldName)
	}
	return zarray.Contains(m.GetFields(), fieldName)
}

// splitFilterKey 从「字段 [操作符]」形式的过滤键中提取字段名与操作符，
// 分隔符兼容空格 / tab / 换行（与 storage 侧 parseExprs 的切分保持一致）。
func splitFilterKey(key string) (field, operator string) {
	field = strings.TrimSpace(key)
	if spaceIdx := strings.IndexAny(field, " \t\n\r"); spaceIdx > 0 {
		return field[:spaceIdx], strings.TrimSpace(field[spaceIdx:])
	}
	return field, ""
}

// splitFilterFieldName 从「字段 [操作符]」形式的过滤键中提取字段名。
func splitFilterFieldName(key string) string {
	field, _ := splitFilterKey(key)
	return field
}

// sanitizeFilterGroupValue 递归校验 $OR/$AND 分组内的字段白名单，
// 返回净化后的值。keep=false 表示分组原本为空或元素全部为无条件空条件（按旧行为整体移除）；
// 被净化清空的非空分组一律以恒假占位代替（见 emptyFilterGroup），
// 避免“看似有约束实则无约束”的查询意外扩大为全表扫描。
func sanitizeFilterGroupValue(m *Schema, value any, depth int) (any, bool) {
	if depth >= maxParseDepth {
		return value, true
	}
	switch v := value.(type) {
	case ztype.Map:
		if len(v) == 0 {
			return v, false
		}
		sanitizeFilterFields(m, v, depth)
		if len(v) == 0 {
			return emptyFilterGroup(), true
		}
		return v, true
	case map[string]interface{}:
		if len(v) == 0 {
			return v, false
		}
		sanitizeFilterFields(m, ztype.Map(v), depth)
		if len(v) == 0 {
			return map[string]interface{}(emptyFilterGroup()), true
		}
		return v, true
	case ztype.Maps:
		return sanitizeFilterGroupMaps(m, v, depth)
	case []ztype.Map:
		return sanitizeFilterGroupMaps(m, v, depth)
	case []map[string]interface{}:
		if len(v) == 0 {
			return v, false
		}
		out := make([]map[string]interface{}, 0, len(v))
		for i := range v {
			// 原本就为空 map 的元素 = 无条件（恒真），忽略之（与顶层 Filter{} 语义一致）；
			// 仅“原本有字段但被白名单净化清空”的元素需要恒假占位 fail-closed。
			if len(v[i]) == 0 {
				continue
			}
			mv := ztype.Map(v[i])
			sanitizeFilterFields(m, mv, depth)
			if len(mv) > 0 {
				out = append(out, map[string]interface{}(mv))
			} else {
				out = append(out, map[string]interface{}(emptyFilterGroup()))
			}
		}
		if len(out) == 0 {
			return v, false
		}
		return out, true
	case []interface{}:
		if len(v) == 0 {
			return v, false
		}
		out := make([]interface{}, 0, len(v))
		for i := range v {
			switch vv := v[i].(type) {
			case ztype.Map:
				// 同 []map：空 map 元素无条件，忽略；净化后为空才恒假占位。
				if len(vv) == 0 {
					continue
				}
				sanitizeFilterFields(m, vv, depth)
				if len(vv) > 0 {
					out = append(out, vv)
				} else {
					out = append(out, emptyFilterGroup())
				}
			case map[string]interface{}:
				if len(vv) == 0 {
					continue
				}
				mv := ztype.Map(vv)
				sanitizeFilterFields(m, mv, depth)
				if len(mv) > 0 {
					out = append(out, mv)
				} else {
					out = append(out, emptyFilterGroup())
				}
			default:
				out = append(out, v[i])
			}
		}
		if len(out) == 0 {
			return v, false
		}
		return out, true
	default:
		return value, true
	}
}

// sanitizeFilterGroupMaps 净化分组内的 map 列表。
// 原本就为空的元素视为无条件而忽略；净化后被清空的子条件保留为恒假占位；
// 元素全部被忽略（或整组原本为空）时返回 keep=false。
func sanitizeFilterGroupMaps(m *Schema, v []ztype.Map, depth int) (any, bool) {
	if len(v) == 0 {
		return v, false
	}
	out := make([]ztype.Map, 0, len(v))
	for i := range v {
		if len(v[i]) == 0 {
			continue
		}
		sanitizeFilterFields(m, v[i], depth)
		if len(v[i]) > 0 {
			out = append(out, v[i])
		} else {
			out = append(out, emptyFilterGroup())
		}
	}
	if len(out) == 0 {
		return v, false
	}
	return out, true
}

// emptyFilterGroup 恒假条件占位：净化后被清空的子条件/分组以此代替，
// 在 SQL 中被解析为 `1 = 0`（zdb builder 对空 IN 的原生表示），恒不成立。
func emptyFilterGroup() ztype.Map {
	return ztype.Map{idKey + " IN": []interface{}{}}
}

// normalizeDeletedAtNull 将时间型软删除下显式的 `deleted_at = nil` / `deleted_at =` 条件规范为
// `deleted_at IS NULL`，避免生成恒不成立的 `deleted_at = NULL`。
// 与 hasFieldInFilter 一样递归处理嵌套 $OR/$AND 分组（depth 用于防自引用栈溢出）。
func normalizeDeletedAtNull(filter ztype.Map, depth int) {
	if depth >= maxParseDepth {
		return
	}
	for key, value := range filter {
		k := strings.TrimSpace(key)
		if k == "" {
			continue
		}
		upperKey := strings.ToUpper(k)
		if upperKey == placeHolderOR || upperKey == placeHolderAND {
			normalizeDeletedAtNested(value, depth+1)
			continue
		}
		if strings.Contains(k, placeHolder) || strings.Contains(k, ".") {
			continue
		}
		if value != nil {
			continue
		}
		fieldName, operator := splitFilterKey(k)
		if fieldName != DeletedAtKey {
			continue
		}
		// 仅无操作符或显式 `=` 的 nil 需要规范化；其余操作符与 nil 无意义，交 parse 层。
		if operator != "" && strings.ToUpper(operator) != "=" {
			continue
		}
		delete(filter, key)
		filter[DeletedAtKey+" IS NULL"] = true
	}
}

func normalizeDeletedAtNested(value any, depth int) {
	if depth >= maxParseDepth {
		return
	}
	switch v := value.(type) {
	case ztype.Map:
		normalizeDeletedAtNull(v, depth)
	case map[string]interface{}:
		normalizeDeletedAtNull(ztype.Map(v), depth)
	case ztype.Maps:
		for i := range v {
			normalizeDeletedAtNull(v[i], depth)
		}
	case []ztype.Map:
		for i := range v {
			normalizeDeletedAtNull(v[i], depth)
		}
	case []map[string]interface{}:
		for i := range v {
			normalizeDeletedAtNull(ztype.Map(v[i]), depth)
		}
	case []interface{}:
		for i := range v {
			switch vv := v[i].(type) {
			case ztype.Map:
				normalizeDeletedAtNull(vv, depth)
			case map[string]interface{}:
				normalizeDeletedAtNull(ztype.Map(vv), depth)
			}
		}
	}
}

// cloneFilterMap 深拷贝过滤 map，保护调用方持有的原始过滤器。
func cloneFilterMap(src ztype.Map) ztype.Map {
	return cloneFilterMapDepth(src, 0)
}

// cloneFilterMapDepth 带深度限制的深拷贝实现。
// depth 超过 maxParseDepth（仅程序化构造的自引用/超深嵌套才会触达）时
// 返回当前层浅拷贝并终止递归：键集合复制、元素共享引用，
// 避免无限递归导致栈溢出，同时该层与调用方隔离。
func cloneFilterMapDepth(src ztype.Map, depth int) ztype.Map {
	if src == nil {
		return ztype.Map{}
	}
	if depth >= maxParseDepth {
		cp := make(ztype.Map, len(src))
		for k, v := range src {
			cp[k] = v
		}
		return cp
	}
	dst := make(ztype.Map, len(src))
	for k, v := range src {
		dst[k] = cloneFilterValueDepth(v, depth+1)
	}
	return dst
}

// cloneFilterValue 深拷贝过滤值中的嵌套容器（map / 切片）。
// 其余类型（标量、func、时间等）保持原引用。
func cloneFilterValue(v any) any {
	return cloneFilterValueDepth(v, 0)
}

func cloneFilterValueDepth(v any, depth int) any {
	switch val := v.(type) {
	case ztype.Map:
		return cloneFilterMapDepth(val, depth)
	case map[string]interface{}:
		// 保持原始动态类型不变量（map[string]interface{} 仍为 map[string]interface{}），
		// 避免调用方对过滤值做精确类型断言时出现类型漂移。
		cp := make(map[string]interface{}, len(val))
		for k, v := range val {
			cp[k] = cloneFilterValueDepth(v, depth+1)
		}
		return cp
	case ztype.Maps:
		cp := make(ztype.Maps, len(val))
		for i := range val {
			cp[i] = cloneFilterMapDepth(val[i], depth+1)
		}
		return cp
	case []ztype.Map:
		cp := make([]ztype.Map, len(val))
		for i := range val {
			cp[i] = cloneFilterMapDepth(val[i], depth+1)
		}
		return cp
	case []map[string]interface{}:
		cp := make([]map[string]interface{}, len(val))
		for i := range val {
			cp[i] = map[string]interface{}(cloneFilterMapDepth(ztype.Map(val[i]), depth+1))
		}
		return cp
	case []interface{}:
		cp := make([]interface{}, len(val))
		for i := range val {
			cp[i] = cloneFilterValueDepth(val[i], depth+1)
		}
		return cp
	default:
		return v
	}
}

func filterToMap(value any) (ztype.Map, bool) {
	if value == nil {
		return nil, false
	}

	switch v := value.(type) {
	case QueryFilter:
		return v.ToMap(), true
	case Filter:
		return ztype.Map(v), true
	case ztype.Map:
		return v, true
	}

	rv := reflect.ValueOf(value)
	if !rv.IsValid() {
		return nil, false
	}
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil, false
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Map && rv.Kind() != reflect.Struct {
		return nil, false
	}

	return ztype.ToMap(rv.Interface()), true
}

func hasFieldInFilter(filter ztype.Map, field string) bool {
	return hasFieldInFilterDepth(filter, field, 0)
}

func hasFieldInFilterDepth(filter ztype.Map, field string, depth int) bool {
	if len(filter) == 0 || depth >= maxParseDepth {
		return false
	}
	for k, v := range filter {
		if k == "" {
			continue
		}
		trimmedKey := strings.TrimSpace(k)
		if trimmedKey == "" {
			continue
		}
		upperKey := strings.ToUpper(trimmedKey)
		if upperKey == placeHolderOR || upperKey == placeHolderAND {
			if hasFieldInNestedFilterDepth(v, field, depth+1) {
				return true
			}
			continue
		}
		if strings.Contains(trimmedKey, placeHolder) {
			continue
		}
		fieldName, _ := splitFilterKey(trimmedKey)
		if strings.Contains(fieldName, ".") {
			continue
		}
		if fieldName == field {
			return true
		}
	}
	return false
}

func hasFieldInNestedFilter(value any, field string) bool {
	return hasFieldInNestedFilterDepth(value, field, 0)
}

func hasFieldInNestedFilterDepth(value any, field string, depth int) bool {
	if depth >= maxParseDepth {
		return false
	}
	switch v := value.(type) {
	case ztype.Map:
		return hasFieldInFilterDepth(v, field, depth+1)
	case map[string]interface{}:
		return hasFieldInFilterDepth(ztype.Map(v), field, depth+1)
	case ztype.Maps:
		for i := range v {
			if hasFieldInFilterDepth(v[i], field, depth+1) {
				return true
			}
		}
		return false
	case []ztype.Map:
		for i := range v {
			if hasFieldInFilterDepth(v[i], field, depth+1) {
				return true
			}
		}
		return false
	case []map[string]interface{}:
		for i := range v {
			if hasFieldInFilterDepth(ztype.Map(v[i]), field, depth+1) {
				return true
			}
		}
		return false
	case []interface{}:
		for i := range v {
			switch vv := v[i].(type) {
			case ztype.Map:
				if hasFieldInFilterDepth(vv, field, depth+1) {
					return true
				}
			case map[string]interface{}:
				if hasFieldInFilterDepth(ztype.Map(vv), field, depth+1) {
					return true
				}
			}
		}
		return false
	default:
		m := ztype.New(value).Map()
		if len(m) == 0 {
			return false
		}
		return hasFieldInFilterDepth(m, field, depth+1)
	}
}
