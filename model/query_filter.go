package model

import (
	"github.com/sohaha/zlsgo/ztype"
)

type QueryFilter interface {
	ToMap() ztype.Map
}

type MapFilter ztype.Map

func (f MapFilter) ToMap() ztype.Map {
	return ztype.Map(f)
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

func Q(m ztype.Map) QueryFilter {
	return MapFilter(m)
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
	return andFilter{filters: filters}
}

func Or(filters ...QueryFilter) QueryFilter {
	return orFilter{filters: filters}
}
