package model

import (
	"github.com/sohaha/zlsgo/ztype"
)

type Repository[T any] struct {
	store  *Store
	mapper Mapper[T]
}

func NewRepository[T any](store *Store, mapper Mapper[T]) *Repository[T] {
	return &Repository[T]{
		store:  store,
		mapper: mapper,
	}
}

func NewMapRepository(store *Store) *Repository[ztype.Map] {
	return &Repository[ztype.Map]{
		store:  store,
		mapper: MapMapper{},
	}
}

func NewStructRepository[T any](store *Store) *Repository[T] {
	return &Repository[T]{
		store:  store,
		mapper: StructMapper[T]{},
	}
}

func (r *Repository[T]) Store() *Store {
	return r.store
}

func (r *Repository[T]) Schema() *Schema {
	return r.store.schema
}

func (r *Repository[T]) Find(filter QueryFilter, fn ...func(*CondOptions)) ([]T, error) {
	rows, err := r.store.Find(Filter(filter.ToMap()), fn...)
	if err != nil {
		return nil, err
	}
	return r.mapper.MapMany(rows)
}

func (r *Repository[T]) FindOne(filter QueryFilter, fn ...func(*CondOptions)) (T, error) {
	row, err := r.store.FindOne(Filter(filter.ToMap()), fn...)
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

func (r *Repository[T]) FindByID(id any, fn ...func(*CondOptions)) (T, error) {
	return r.FindOne(ID(id), fn...)
}

func (r *Repository[T]) First(filter QueryFilter, fn ...func(*CondOptions)) (T, error) {
	return r.FindOne(filter, fn...)
}

func (r *Repository[T]) FindByIDs(ids []any, fn ...func(*CondOptions)) ([]T, error) {
	return r.Find(In(idKey, ids), fn...)
}

func (r *Repository[T]) All(fn ...func(*CondOptions)) ([]T, error) {
	return r.Find(Q(ztype.Map{}), fn...)
}

func (r *Repository[T]) Pages(page, pagesize int, filter QueryFilter, fn ...func(*CondOptions)) (*RepositoryPageData[T], error) {
	pageData, err := r.store.Pages(page, pagesize, Filter(filter.ToMap()), fn...)
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

func (r *Repository[T]) Count(filter QueryFilter, fn ...func(*CondOptions)) (uint64, error) {
	return r.store.Count(Filter(filter.ToMap()), fn...)
}

func (r *Repository[T]) Exists(filter QueryFilter, fn ...func(*CondOptions)) (bool, error) {
	return r.store.Exists(Filter(filter.ToMap()), fn...)
}

func (r *Repository[T]) Insert(data ztype.Map, fn ...func(*InsertOptions)) (any, error) {
	return r.store.Insert(data, fn...)
}

func (r *Repository[T]) InsertMany(data ztype.Maps, fn ...func(*InsertOptions)) (any, error) {
	return r.store.InsertMany(data, fn...)
}

func (r *Repository[T]) Update(filter QueryFilter, data ztype.Map, fn ...func(*CondOptions)) (int64, error) {
	return r.store.Update(Filter(filter.ToMap()), data, fn...)
}

func (r *Repository[T]) UpdateMany(filter QueryFilter, data ztype.Map, fn ...func(*CondOptions)) (int64, error) {
	return r.store.UpdateMany(Filter(filter.ToMap()), data, fn...)
}

func (r *Repository[T]) UpdateByID(id any, data ztype.Map, fn ...func(*CondOptions)) (int64, error) {
	return r.Update(ID(id), data, fn...)
}

func (r *Repository[T]) UpdateByIDs(ids []any, data ztype.Map, fn ...func(*CondOptions)) (int64, error) {
	return r.UpdateMany(In(idKey, ids), data, fn...)
}

func (r *Repository[T]) Delete(filter QueryFilter, fn ...func(*CondOptions)) (int64, error) {
	return r.store.Delete(Filter(filter.ToMap()), fn...)
}

func (r *Repository[T]) DeleteMany(filter QueryFilter, fn ...func(*CondOptions)) (int64, error) {
	return r.store.DeleteMany(Filter(filter.ToMap()), fn...)
}

func (r *Repository[T]) DeleteByID(id any, fn ...func(*CondOptions)) (int64, error) {
	return r.Delete(ID(id), fn...)
}

func (r *Repository[T]) DeleteByIDs(ids []any, fn ...func(*CondOptions)) (int64, error) {
	return r.DeleteMany(In(idKey, ids), fn...)
}

func (r *Repository[T]) Tx(fn func(txRepo *Repository[T]) error) error {
	storage := r.store.schema.Storage

	return storage.Transaction(func(txStorage Storageer) error {
		txSchema := *r.store.schema
		txSchema.Storage = txStorage
		txStore := &Store{schema: &txSchema}
		txRepo := &Repository[T]{
			store:  txStore,
			mapper: r.mapper,
		}
		return fn(txRepo)
	})
}

type RepositoryPageData[T any] struct {
	Items    []T      `json:"items"`
	Page     PageInfo `json:"page"`
	pagesize uint     `json:"-"`
}
