package model

import (
	"github.com/sohaha/zlsgo/ztype"
)

// Repository 泛型仓储
type Repository[T any, F any, C any, U any] struct {
	store  *Store
	mapper Mapper[T]
}

// NewRepository 创建新的泛型仓储实例
func NewRepository[T any, F any, C any, U any](store *Store, mapper Mapper[T]) *Repository[T, F, C, U] {
	return &Repository[T, F, C, U]{
		store:  store,
		mapper: mapper,
	}
}

// NewMapRepository 创建基于 Map 的仓储实例
func NewMapRepository(store *Store) *Repository[ztype.Map, QueryFilter, ztype.Map, ztype.Map] {
	return &Repository[ztype.Map, QueryFilter, ztype.Map, ztype.Map]{
		store:  store,
		mapper: MapMapper{},
	}
}

// NewStructRepository 创建基于结构体的仓储实例
func NewStructRepository[T any, F any, C any, U any](store *Store) *Repository[T, F, C, U] {
	return &Repository[T, F, C, U]{
		store:  store,
		mapper: StructMapper[T]{},
	}
}

// Store 返回存储实例
func (r *Repository[T, F, C, U]) Store() *Store {
	return r.store
}

// Schema 返回模型结构定义
func (r *Repository[T, F, C, U]) Schema() *Schema {
	return r.store.schema
}

// find 查找多条记录
func (r *Repository[T, F, C, U]) find(filter QueryFilter, fn ...func(*CondOptions)) ([]T, error) {
	rows, err := r.store.Find(filter, fn...)
	if err != nil {
		return nil, err
	}
	return r.mapper.MapMany(rows)
}

// findOne 查找单条记录
func (r *Repository[T, F, C, U]) findOne(filter QueryFilter, fn ...func(*CondOptions)) (T, error) {
	row, err := r.store.FindOne(filter, fn...)
	if err != nil {
		var zero T
		return zero, err
	}
	if len(row) == 0 {
		var zero T
		return zero, ErrNoRecord
	}
	return r.mapper.MapOne(row)
}

// pages 分页查询记录
func (r *Repository[T, F, C, U]) pages(page, pagesize int, filter QueryFilter, fn ...func(*CondOptions)) (*RepositoryPageData[T], error) {
	pageData, err := r.store.Pages(page, pagesize, filter, fn...)
	if err != nil {
		return nil, err
	}

	items, err := r.mapper.MapMany(pageData.Items)
	if err != nil {
		return nil, err
	}

	return &RepositoryPageData[T]{
		Items:    items,
		Page:     pageData.Page,
		pagesize: pageData.pagesize,
	}, nil
}

// Find 根据过滤器查找多条记录
func (r *Repository[T, F, C, U]) Find(filter F, fn ...func(*CondOptions)) ([]T, error) {
	return r.find(Q(filter), fn...)
}

// FindOne 根据过滤器查找单条记录
func (r *Repository[T, F, C, U]) FindOne(filter F, fn ...func(*CondOptions)) (T, error) {
	return r.findOne(Q(filter), fn...)
}

// FindByID 根据 ID 查找记录
func (r *Repository[T, F, C, U]) FindByID(id any, fn ...func(*CondOptions)) (T, error) {
	return r.findOne(ID(id), fn...)
}

// First 查找第一条记录
func (r *Repository[T, F, C, U]) First(filter F, fn ...func(*CondOptions)) (T, error) {
	return r.FindOne(filter, fn...)
}

// FindByIDs 根据 ID 列表查找多条记录
func (r *Repository[T, F, C, U]) FindByIDs(ids []any, fn ...func(*CondOptions)) ([]T, error) {
	return r.find(In(idKey, ids), fn...)
}

// All 查找所有记录
func (r *Repository[T, F, C, U]) All(fn ...func(*CondOptions)) ([]T, error) {
	return r.find(Filter{}, fn...)
}

// Pages 分页查询记录
func (r *Repository[T, F, C, U]) Pages(page, pagesize int, filter F, fn ...func(*CondOptions)) (*RepositoryPageData[T], error) {
	return r.pages(page, pagesize, Q(filter), fn...)
}

// Count 统计记录数量
func (r *Repository[T, F, C, U]) Count(filter F, fn ...func(*CondOptions)) (uint64, error) {
	return r.store.Count(Q(filter), fn...)
}

// Exists 检查记录是否存在
func (r *Repository[T, F, C, U]) Exists(filter F, fn ...func(*CondOptions)) (bool, error) {
	return r.store.Exists(Q(filter), fn...)
}

// Insert 插入单条记录
func (r *Repository[T, F, C, U]) Insert(data C, fn ...func(*InsertOptions)) (any, error) {
	return r.store.Insert(data, fn...)
}

// InsertMany 批量插入记录
func (r *Repository[T, F, C, U]) InsertMany(data []C, fn ...func(*InsertOptions)) (any, error) {
	return r.store.InsertMany(data, fn...)
}

// Update 更新符合条件的记录
func (r *Repository[T, F, C, U]) Update(filter F, data U, fn ...func(*CondOptions)) (int64, error) {
	return r.store.Update(Q(filter), data, fn...)
}

// UpdateMany 批量更新记录
func (r *Repository[T, F, C, U]) UpdateMany(filter F, data U, fn ...func(*CondOptions)) (int64, error) {
	return r.store.UpdateMany(Q(filter), data, fn...)
}

// UpdateByID 根据 ID 更新记录
func (r *Repository[T, F, C, U]) UpdateByID(id any, data U, fn ...func(*CondOptions)) (int64, error) {
	return r.store.Update(ID(id), data, fn...)
}

// UpdateByIDs 根据 ID 列表批量更新记录
func (r *Repository[T, F, C, U]) UpdateByIDs(ids []any, data U, fn ...func(*CondOptions)) (int64, error) {
	return r.store.UpdateMany(In(idKey, ids), data, fn...)
}

// Delete 删除符合条件的记录
func (r *Repository[T, F, C, U]) Delete(filter F, fn ...func(*CondOptions)) (int64, error) {
	return r.store.Delete(Q(filter), fn...)
}

// DeleteMany 批量删除记录
func (r *Repository[T, F, C, U]) DeleteMany(filter F, fn ...func(*CondOptions)) (int64, error) {
	return r.store.DeleteMany(Q(filter), fn...)
}

// DeleteByID 根据 ID 删除记录
func (r *Repository[T, F, C, U]) DeleteByID(id any, fn ...func(*CondOptions)) (int64, error) {
	return r.store.Delete(ID(id), fn...)
}

// DeleteByIDs 根据 ID 列表批量删除记录
func (r *Repository[T, F, C, U]) DeleteByIDs(ids []any, fn ...func(*CondOptions)) (int64, error) {
	return r.store.DeleteMany(In(idKey, ids), fn...)
}

// Tx 在事务中执行操作
func (r *Repository[T, F, C, U]) Tx(fn func(txRepo *Repository[T, F, C, U]) error) error {
	storage := r.store.schema.Storage

	return storage.Transaction(func(txStorage Storageer) error {
		txSchema := cloneSchemaWithStorage(r.store.schema, txStorage)
		txStore := &Store{schema: txSchema}
		txRepo := &Repository[T, F, C, U]{
			store:  txStore,
			mapper: r.mapper,
		}
		return fn(txRepo)
	})
}

// RepositoryPageData 仓储分页数据
type RepositoryPageData[T any] struct {
	Items    []T      `json:"items"`
	Page     PageInfo `json:"page"`
	pagesize uint     `json:"-"`
}
