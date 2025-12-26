package model

import "errors"

// ErrNoRecord 记录未找到错误
var ErrNoRecord = errors.New("record not found")

// OrderByItem 排序项
type OrderByItem struct {
	Field     string
	Direction string
}

// Query 查询构建器
type Query[T any, F any, C any, U any] struct {
	repo      *Repository[T, F, C, U]
	filter    QueryFilter
	fields    []string
	orderBy   []OrderByItem
	groupBy   []string
	limit     int
	offset    int
	relations []string
}

const (
	// defaultQueryFieldsCap 默认字段容量
	defaultQueryFieldsCap = 8
	// defaultQueryOrderByCap 默认排序容量
	defaultQueryOrderByCap = 4
	// defaultQueryRelationsCap 默认关联容量
	defaultQueryRelationsCap = 2
)

// Query 创建新的查询构建器
func (r *Repository[T, F, C, U]) Query() *Query[T, F, C, U] {
	return &Query[T, F, C, U]{
		repo:      r,
		filter:    Filter{},
		fields:    make([]string, 0, defaultQueryFieldsCap),
		orderBy:   make([]OrderByItem, 0, defaultQueryOrderByCap),
		relations: make([]string, 0, defaultQueryRelationsCap),
	}
}

// Where 添加 WHERE 条件
func (q *Query[T, F, C, U]) Where(field string, value any) *Query[T, F, C, U] {
	q.filter = And(q.filter, Eq(field, value))
	return q
}

// WhereFilter 添加过滤器条件
func (q *Query[T, F, C, U]) WhereFilter(filter F) *Query[T, F, C, U] {
	q.filter = And(q.filter, Q(filter))
	return q
}

// WhereID 添加 ID 条件
func (q *Query[T, F, C, U]) WhereID(id any) *Query[T, F, C, U] {
	q.filter = And(q.filter, ID(id))
	return q
}

// WhereIn 添加 IN 条件
func (q *Query[T, F, C, U]) WhereIn(field string, values any) *Query[T, F, C, U] {
	q.filter = And(q.filter, In(field, values))
	return q
}

// WhereNot 添加 NOT 条件
func (q *Query[T, F, C, U]) WhereNot(field string, value any) *Query[T, F, C, U] {
	q.filter = And(q.filter, Ne(field, value))
	return q
}

// WhereGt 添加大于条件
func (q *Query[T, F, C, U]) WhereGt(field string, value any) *Query[T, F, C, U] {
	q.filter = And(q.filter, Gt(field, value))
	return q
}

// WhereGe 添加大于等于条件
func (q *Query[T, F, C, U]) WhereGe(field string, value any) *Query[T, F, C, U] {
	q.filter = And(q.filter, Ge(field, value))
	return q
}

// WhereLt 添加小于条件
func (q *Query[T, F, C, U]) WhereLt(field string, value any) *Query[T, F, C, U] {
	q.filter = And(q.filter, Lt(field, value))
	return q
}

// WhereLe 添加小于等于条件
func (q *Query[T, F, C, U]) WhereLe(field string, value any) *Query[T, F, C, U] {
	q.filter = And(q.filter, Le(field, value))
	return q
}

// WhereLike 添加 LIKE 条件
func (q *Query[T, F, C, U]) WhereLike(field string, pattern string) *Query[T, F, C, U] {
	q.filter = And(q.filter, Like(field, pattern))
	return q
}

// WhereBetween 添加 BETWEEN 条件
func (q *Query[T, F, C, U]) WhereBetween(field string, start, end any) *Query[T, F, C, U] {
	q.filter = And(q.filter, Between(field, start, end))
	return q
}

// WhereNull 添加 IS NULL 条件
func (q *Query[T, F, C, U]) WhereNull(field string) *Query[T, F, C, U] {
	q.filter = And(q.filter, IsNull(field))
	return q
}

// WhereNotNull 添加 IS NOT NULL 条件
func (q *Query[T, F, C, U]) WhereNotNull(field string) *Query[T, F, C, U] {
	q.filter = And(q.filter, IsNotNull(field))
	return q
}

// OrWhere adds an OR condition that groups the provided filters.
// Note: This creates "AND (filter1 OR filter2 OR ...)" pattern.
// For a pure OR without preceding AND, use repo.Find(Or(...)) when F is QueryFilter.
func (q *Query[T, F, C, U]) OrWhere(filters ...F) *Query[T, F, C, U] {
	if len(filters) == 0 {
		return q
	}
	orFilters := make([]QueryFilter, 0, len(filters))
	for _, filter := range filters {
		orFilters = append(orFilters, Q(filter))
	}
	q.filter = And(q.filter, Or(orFilters...))
	return q
}

// Select 设置查询字段
func (q *Query[T, F, C, U]) Select(fields ...string) *Query[T, F, C, U] {
	q.fields = fields
	return q
}

// OrderBy 添加排序条件
func (q *Query[T, F, C, U]) OrderBy(field string, direction ...string) *Query[T, F, C, U] {
	dir := "ASC"
	if len(direction) > 0 {
		dir = direction[0]
	}
	q.orderBy = append(q.orderBy, OrderByItem{Field: field, Direction: dir})
	return q
}

// OrderByDesc 添加降序排序
func (q *Query[T, F, C, U]) OrderByDesc(field string) *Query[T, F, C, U] {
	q.orderBy = append(q.orderBy, OrderByItem{Field: field, Direction: "DESC"})
	return q
}

// GroupBy 添加分组条件
func (q *Query[T, F, C, U]) GroupBy(fields ...string) *Query[T, F, C, U] {
	q.groupBy = fields
	return q
}

// Limit 设置查询限制数量
func (q *Query[T, F, C, U]) Limit(limit int) *Query[T, F, C, U] {
	q.limit = limit
	return q
}

// Offset 设置查询偏移量
func (q *Query[T, F, C, U]) Offset(offset int) *Query[T, F, C, U] {
	q.offset = offset
	return q
}

// WithRelation 添加关联查询
func (q *Query[T, F, C, U]) WithRelation(names ...string) *Query[T, F, C, U] {
	q.relations = append(q.relations, names...)
	return q
}

// buildCondOptions 构建查询条件选项
func (q *Query[T, F, C, U]) buildCondOptions() func(*CondOptions) {
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
			opts.Relations = append(opts.Relations[:0], q.relations...)
		}
	}
}

// Find 执行查询并返回多条记录
func (q *Query[T, F, C, U]) Find() ([]T, error) {
	return q.repo.find(q.filter, q.buildCondOptions())
}

// FindOne 执行查询并返回单条记录
func (q *Query[T, F, C, U]) FindOne() (T, error) {
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

// First 返回第一条记录
func (q *Query[T, F, C, U]) First() (T, error) {
	return q.FindOne()
}

// Count 统计符合条件的记录数量
func (q *Query[T, F, C, U]) Count() (uint64, error) {
	return q.repo.store.Count(q.filter, q.buildCondOptions())
}

// Exists 检查符合条件的记录是否存在
func (q *Query[T, F, C, U]) Exists() (bool, error) {
	return q.repo.store.Exists(q.filter, q.buildCondOptions())
}

// Pages 分页查询记录
func (q *Query[T, F, C, U]) Pages(page, pagesize int) (*RepositoryPageData[T], error) {
	return q.repo.pages(page, pagesize, q.filter, q.buildCondOptions())
}

// Update 更新符合条件的记录
func (q *Query[T, F, C, U]) Update(data U) (int64, error) {
	return q.repo.store.Update(q.filter, data, q.buildCondOptions())
}

// Delete 删除符合条件的记录
func (q *Query[T, F, C, U]) Delete() (int64, error) {
	return q.repo.store.Delete(q.filter, q.buildCondOptions())
}
