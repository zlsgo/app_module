package model

import (
	"github.com/sohaha/zlsgo/ztype"
)

type QueryFilter interface {
	ToMap() ztype.Map
}

type filterMapAppender interface {
	appendToMap(dst ztype.Map) ztype.Map
}

type IDQueryFilter struct {
	ID any
}

func (f IDQueryFilter) ToMap() ztype.Map {
	return f.appendToMap(make(ztype.Map, 1))
}

func (f IDQueryFilter) appendToMap(dst ztype.Map) ztype.Map {
	if dst == nil {
		dst = make(ztype.Map, 1)
	}
	dst[idKey] = f.ID
	return dst
}

type conditionFilter struct {
	field string
	op    string
	value any
}

func (f conditionFilter) ToMap() ztype.Map {
	return f.appendToMap(make(ztype.Map, 1))
}

func (f conditionFilter) appendToMap(dst ztype.Map) ztype.Map {
	key := f.field
	if f.op != "" {
		key = f.field + " " + f.op
	}
	if dst == nil {
		dst = make(ztype.Map, 1)
	}
	dst[key] = f.value
	return dst
}

type andFilter struct {
	filters []QueryFilter
}

func (f andFilter) ToMap() ztype.Map {
	return f.appendToMap(make(ztype.Map, estimateFilterMapSize(f.filters)))
}

func (f andFilter) appendToMap(dst ztype.Map) ztype.Map {
	for _, filter := range f.filters {
		dst = appendFilterMap(dst, filter)
	}
	return dst
}

type orFilter struct {
	filters []QueryFilter
}

func (f orFilter) ToMap() ztype.Map {
	subFilters := make([]ztype.Map, 0, len(f.filters))
	for _, filter := range f.filters {
		subFilters = append(subFilters, filterToMapEntry(filter))
	}
	return ztype.Map{placeHolderOR: subFilters}
}

func Q[F any](m F) QueryFilter {
	if filter, ok := any(m).(QueryFilter); ok {
		return filter
	}
	if mapData, ok := filterToMap(m); ok {
		return Filter(mapData)
	}
	return Filter(ztype.Map{})
}

func ID(id any) QueryFilter {
	return IDQueryFilter{ID: id}
}

func Eq(field string, value any) QueryFilter {
	return conditionFilter{field: field, value: value}
}

func Ne(field string, value any) QueryFilter {
	return conditionFilter{field: field, op: "!=", value: value}
}

func Gt(field string, value any) QueryFilter {
	return conditionFilter{field: field, op: ">", value: value}
}

func Ge(field string, value any) QueryFilter {
	return conditionFilter{field: field, op: ">=", value: value}
}

func Lt(field string, value any) QueryFilter {
	return conditionFilter{field: field, op: "<", value: value}
}

func Le(field string, value any) QueryFilter {
	return conditionFilter{field: field, op: "<=", value: value}
}

func In(field string, values any) QueryFilter {
	return conditionFilter{field: field, op: "IN", value: values}
}

func NotIn(field string, values any) QueryFilter {
	return conditionFilter{field: field, op: "NOT IN", value: values}
}

func Like(field string, pattern string) QueryFilter {
	return conditionFilter{field: field, op: "LIKE", value: pattern}
}

func Between(field string, start, end any) QueryFilter {
	return conditionFilter{field: field, op: "BETWEEN", value: []any{start, end}}
}

func IsNull(field string) QueryFilter {
	return conditionFilter{field: field, op: "IS NULL", value: nil}
}

func IsNotNull(field string) QueryFilter {
	return conditionFilter{field: field, op: "IS NOT NULL", value: nil}
}

func And(filters ...QueryFilter) QueryFilter {
	compact := compactAndFilters(filters...)
	switch len(compact) {
	case 0:
		return Filter{}
	case 1:
		return compact[0]
	default:
		return andFilter{filters: compact}
	}
}

func Or(filters ...QueryFilter) QueryFilter {
	compact := compactOrFilters(filters...)
	switch len(compact) {
	case 0:
		return Filter{}
	case 1:
		return compact[0]
	default:
		return orFilter{filters: compact}
	}
}

func compactAndFilters(filters ...QueryFilter) []QueryFilter {
	if len(filters) == 0 {
		return nil
	}
	result := make([]QueryFilter, 0, len(filters))
	for _, filter := range filters {
		if isEmptyQueryFilter(filter) {
			continue
		}
		if nested, ok := filter.(andFilter); ok {
			result = append(result, nested.filters...)
			continue
		}
		result = append(result, filter)
	}
	return result
}

func compactOrFilters(filters ...QueryFilter) []QueryFilter {
	if len(filters) == 0 {
		return nil
	}
	result := make([]QueryFilter, 0, len(filters))
	for _, filter := range filters {
		if isEmptyQueryFilter(filter) {
			continue
		}
		if nested, ok := filter.(orFilter); ok {
			result = append(result, nested.filters...)
			continue
		}
		result = append(result, filter)
	}
	return result
}

func appendFilterMap(dst ztype.Map, filter QueryFilter) ztype.Map {
	if isEmptyQueryFilter(filter) {
		return dst
	}
	if dst == nil {
		dst = make(ztype.Map)
	}
	if appender, ok := filter.(filterMapAppender); ok {
		return appender.appendToMap(dst)
	}
	for k, v := range filter.ToMap() {
		dst[k] = v
	}
	return dst
}

func filterToMapEntry(filter QueryFilter) ztype.Map {
	if isEmptyQueryFilter(filter) {
		return ztype.Map{}
	}
	if appender, ok := filter.(filterMapAppender); ok {
		return appender.appendToMap(make(ztype.Map, estimateFilterMapSize([]QueryFilter{filter})))
	}
	return filter.ToMap()
}

func estimateFilterMapSize(filters []QueryFilter) int {
	size := 0
	for _, filter := range filters {
		switch v := filter.(type) {
		case nil:
			continue
		case andFilter:
			size += estimateFilterMapSize(v.filters)
		case Filter:
			size += len(v)
		default:
			size++
		}
	}
	if size == 0 {
		return 1
	}
	return size
}

func isEmptyQueryFilter(filter QueryFilter) bool {
	if filter == nil {
		return true
	}
	switch v := filter.(type) {
	case Filter:
		return len(v) == 0
	case andFilter:
		return len(v.filters) == 0
	case orFilter:
		return len(v.filters) == 0
	default:
		return false
	}
}
