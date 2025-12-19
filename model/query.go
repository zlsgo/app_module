package model

import (
	"errors"

	"github.com/sohaha/zlsgo/ztype"
)

var ErrNoRecord = errors.New("record not found")

type OrderByItem struct {
	Field     string
	Direction string
}

type Query[T any] struct {
	repo      *Repository[T]
	filter    QueryFilter
	fields    []string
	orderBy   []OrderByItem
	groupBy   []string
	limit     int
	offset    int
	relations []string
}

const (
	defaultQueryFieldsCap    = 8
	defaultQueryOrderByCap   = 4
	defaultQueryRelationsCap = 2
)

func (r *Repository[T]) Query() *Query[T] {
	return &Query[T]{
		repo:    r,
		filter:  MapFilter{},
		fields:  make([]string, 0, defaultQueryFieldsCap),
		orderBy: make([]OrderByItem, 0, defaultQueryOrderByCap),
	}
}

func (q *Query[T]) Where(field string, value any) *Query[T] {
	q.filter = And(q.filter, Eq(field, value))
	return q
}

func (q *Query[T]) WhereFilter(filter QueryFilter) *Query[T] {
	q.filter = And(q.filter, filter)
	return q
}

func (q *Query[T]) WhereID(id any) *Query[T] {
	q.filter = And(q.filter, ID(id))
	return q
}

func (q *Query[T]) WhereIn(field string, values any) *Query[T] {
	q.filter = And(q.filter, In(field, values))
	return q
}

func (q *Query[T]) WhereNot(field string, value any) *Query[T] {
	q.filter = And(q.filter, Ne(field, value))
	return q
}

func (q *Query[T]) WhereGt(field string, value any) *Query[T] {
	q.filter = And(q.filter, Gt(field, value))
	return q
}

func (q *Query[T]) WhereGe(field string, value any) *Query[T] {
	q.filter = And(q.filter, Ge(field, value))
	return q
}

func (q *Query[T]) WhereLt(field string, value any) *Query[T] {
	q.filter = And(q.filter, Lt(field, value))
	return q
}

func (q *Query[T]) WhereLe(field string, value any) *Query[T] {
	q.filter = And(q.filter, Le(field, value))
	return q
}

func (q *Query[T]) WhereLike(field string, pattern string) *Query[T] {
	q.filter = And(q.filter, Like(field, pattern))
	return q
}

func (q *Query[T]) WhereBetween(field string, start, end any) *Query[T] {
	q.filter = And(q.filter, Between(field, start, end))
	return q
}

func (q *Query[T]) WhereNull(field string) *Query[T] {
	q.filter = And(q.filter, IsNull(field))
	return q
}

func (q *Query[T]) WhereNotNull(field string) *Query[T] {
	q.filter = And(q.filter, IsNotNull(field))
	return q
}

// OrWhere adds an OR condition that groups the provided filters.
// Note: This creates "AND (filter1 OR filter2 OR ...)" pattern.
// For a pure OR without preceding AND, use repo.Find(Or(filters...)).
func (q *Query[T]) OrWhere(filters ...QueryFilter) *Query[T] {
	q.filter = And(q.filter, Or(filters...))
	return q
}

func (q *Query[T]) Select(fields ...string) *Query[T] {
	q.fields = fields
	return q
}

func (q *Query[T]) OrderBy(field string, direction ...string) *Query[T] {
	dir := "ASC"
	if len(direction) > 0 {
		dir = direction[0]
	}
	q.orderBy = append(q.orderBy, OrderByItem{Field: field, Direction: dir})
	return q
}

func (q *Query[T]) OrderByDesc(field string) *Query[T] {
	q.orderBy = append(q.orderBy, OrderByItem{Field: field, Direction: "DESC"})
	return q
}

func (q *Query[T]) GroupBy(fields ...string) *Query[T] {
	q.groupBy = fields
	return q
}

func (q *Query[T]) Limit(limit int) *Query[T] {
	q.limit = limit
	return q
}

func (q *Query[T]) Offset(offset int) *Query[T] {
	q.offset = offset
	return q
}

func (q *Query[T]) WithRelation(names ...string) *Query[T] {
	q.relations = append(q.relations, names...)
	return q
}

func (q *Query[T]) buildCondOptions() func(*CondOptions) {
	return func(opts *CondOptions) {
		if len(q.fields) > 0 {
			opts.Fields = append(opts.Fields[:0], q.fields...)
		}
		if len(q.orderBy) > 0 {
			opts.OrderBy = append(opts.OrderBy[:0], q.orderBy...)
		}
		if len(q.groupBy) > 0 {
			opts.GroupBy = append(opts.GroupBy[:0], q.groupBy...)
		}
		if q.limit > 0 {
			opts.Limit = q.limit
		}
		if q.offset > 0 {
			opts.Offset = q.offset
		}
		if len(q.relations) > 0 {
			if len(opts.Fields) == 0 {
				opts.Fields = append(opts.Fields, "*")
			}
			opts.Fields = append(opts.Fields, q.relations...)
		}
	}
}

func (q *Query[T]) Find() ([]T, error) {
	return q.repo.Find(q.filter, q.buildCondOptions())
}

func (q *Query[T]) FindOne() (T, error) {
	q.limit = 1
	results, err := q.Find()
	if err != nil {
		var zero T
		return zero, err
	}
	if len(results) == 0 {
		var zero T
		return zero, ErrNoRecord
	}
	return results[0], nil
}

func (q *Query[T]) First() (T, error) {
	return q.FindOne()
}

func (q *Query[T]) Count() (uint64, error) {
	return q.repo.Count(q.filter, q.buildCondOptions())
}

func (q *Query[T]) Exists() (bool, error) {
	return q.repo.Exists(q.filter, q.buildCondOptions())
}

func (q *Query[T]) Pages(page, pagesize int) (*RepositoryPageData[T], error) {
	return q.repo.Pages(page, pagesize, q.filter, q.buildCondOptions())
}

func (q *Query[T]) Update(data ztype.Map) (int64, error) {
	return q.repo.Update(q.filter, data, q.buildCondOptions())
}

func (q *Query[T]) Delete() (int64, error) {
	return q.repo.Delete(q.filter, q.buildCondOptions())
}
