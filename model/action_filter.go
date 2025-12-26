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

	filterMap = cloneFilterMap(filterMap)

	// 过滤无效字段：排除不在模型定义中的字段
	for key := range filterMap {
		k := zstring.TrimSpace(key)
		if k == "" || strings.Contains(k, placeHolder) {
			continue
		}
		if strings.Contains(k, ".") {
			continue
		}
		fieldName := k
		if spaceIdx := strings.Index(k, " "); spaceIdx > 0 {
			fieldName = k[:spaceIdx]
		}
		if fieldName == DeletedAtKey {
			if !*m.define.Options.SoftDeletes {
				delete(filterMap, key)
			}
			continue
		}
		if m.fullFieldsMap != nil {
			if _, ok := m.fullFieldsMap[fieldName]; !ok {
				delete(filterMap, key)
			}
			continue
		}
		if len(m.fullFields) > 0 {
			if !zarray.Contains(m.fullFields, fieldName) {
				delete(filterMap, key)
			}
			continue
		}
		if !zarray.Contains(m.GetFields(), fieldName) {
			delete(filterMap, key)
		}
	}

	if *m.define.Options.SoftDeletes {
		if !hasFieldInFilter(filterMap, DeletedAtKey) {
			if *m.define.Options.SoftDeleteIsTime {
				filterMap[DeletedAtKey] = nil
			} else {
				filterMap[DeletedAtKey] = 0
			}
		}
	}

	return
}

func cloneFilterMap(src ztype.Map) ztype.Map {
	if src == nil {
		return ztype.Map{}
	}
	dst := make(ztype.Map, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
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
	if len(filter) == 0 {
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
			if hasFieldInNestedFilter(v, field) {
				return true
			}
			continue
		}
		if strings.Contains(trimmedKey, placeHolder) {
			continue
		}
		fieldName := trimmedKey
		if spaceIdx := strings.IndexAny(fieldName, " \t"); spaceIdx > 0 {
			fieldName = fieldName[:spaceIdx]
		}
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
	switch v := value.(type) {
	case ztype.Map:
		return hasFieldInFilter(v, field)
	case map[string]interface{}:
		return hasFieldInFilter(ztype.Map(v), field)
	case ztype.Maps:
		for i := range v {
			if hasFieldInFilter(v[i], field) {
				return true
			}
		}
		return false
	case []ztype.Map:
		for i := range v {
			if hasFieldInFilter(v[i], field) {
				return true
			}
		}
		return false
	case []map[string]interface{}:
		for i := range v {
			if hasFieldInFilter(ztype.Map(v[i]), field) {
				return true
			}
		}
		return false
	case []interface{}:
		for i := range v {
			switch vv := v[i].(type) {
			case ztype.Map:
				if hasFieldInFilter(vv, field) {
					return true
				}
			case map[string]interface{}:
				if hasFieldInFilter(ztype.Map(vv), field) {
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
		return hasFieldInFilter(m, field)
	}
}
