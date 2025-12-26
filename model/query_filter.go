package model

import (
	"github.com/sohaha/zlsgo/ztype"
)

type QueryFilter interface {
	ToMap() ztype.Map
}

type IDQueryFilter struct {
	ID any
}

func (f IDQueryFilter) ToMap() ztype.Map {
	return ztype.Map{idKey: f.ID}
}

type conditionFilter struct {
	field string
	op    string
	value any
}

func (f conditionFilter) ToMap() ztype.Map {
	key := f.field
	if f.op != "" {
		key = f.field + " " + f.op
	}
	return ztype.Map{key: f.value}
}

type andFilter struct {
	filters []QueryFilter
}

func (f andFilter) ToMap() ztype.Map {
	result := make(ztype.Map)
	for _, filter := range f.filters {
		for k, v := range filter.ToMap() {
			result[k] = v
		}
	}
	return result
}

type orFilter struct {
	filters []QueryFilter
}

func (f orFilter) ToMap() ztype.Map {
	subFilters := make([]ztype.Map, len(f.filters))
	for i, filter := range f.filters {
		subFilters[i] = filter.ToMap()
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
	return andFilter{filters: compactFilters(filters...)}
}

func Or(filters ...QueryFilter) QueryFilter {
	return orFilter{filters: compactFilters(filters...)}
}

func compactFilters(filters ...QueryFilter) []QueryFilter {
	if len(filters) == 0 {
		return nil
	}
	result := make([]QueryFilter, 0, len(filters))
	for _, filter := range filters {
		if filter == nil {
			continue
		}
		result = append(result, filter)
	}
	return result
}
