package model

import (
	"strconv"
	"sync/atomic"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb/builder"
)

// condCounter 条件计数器
var condCounter uint64

// Filter 过滤器类型
type Filter ztype.Map

// NewFilter 创建新的过滤器
func NewFilter() Filter {
	return Filter{}
}

// ToMap 转换为 Map
func (f Filter) ToMap() ztype.Map {
	return ztype.Map(f)
}

// Cond 添加条件表达式
func (f Filter) Cond(fn func(*builder.BuildCond) (exprs string)) Filter {
	id := atomic.AddUint64(&condCounter, 1)
	f[placeHolder+strconv.FormatUint(id, 10)] = fn
	return f
}

// Set 设置字段条件
func (f Filter) Set(field string, cond any) Filter {
	f[field] = cond
	return f
}

// Get 获取字段值
func (f Filter) Get(field string) ztype.Type {
	return ztype.New(f[field])
}

// Model 创建模型存储实例
func (m *Schema) Model() *Store {
	return &Store{schema: m}
}

// Stores 存储集合
type Stores struct {
	items *zarray.Maper[string, *Store]
}

// Get 获取指定名称的存储实例
func (m *Stores) Get(name string) (*Store, bool) {
	return m.items.Get(name)
}

// MustGet 获取存储实例，不存在则 panic
func (m *Stores) MustGet(name string) *Store {
	o, _ := m.items.Get(name)
	if o == nil {
		panic("operation " + name + " not found")
	}
	return o
}

// All 获取所有存储实例
func (m *Stores) All() (models []*Store) {
	models = make([]*Store, 0, m.items.Len())

	m.items.ForEach(func(key string, value *Store) bool {
		models = append(models, value)
		return true
	})

	return
}

// Schema 获取模型结构定义
func (o *Store) Schema(Storage ...Storageer) *Schema {
	if len(Storage) > 0 {
		return cloneSchemaWithStorage(o.schema, Storage[0])
	}
	return o.schema
}

// Insert 插入单条数据
func (o *Store) Insert(data any, fn ...func(*InsertOptions)) (lastId interface{}, err error) {
	return Insert(o.schema, data, fn...)
}

// InsertMany 批量插入数据
func (o *Store) InsertMany(data any, fn ...func(*InsertOptions)) (lastId interface{}, err error) {
	return InsertMany(o.schema, data, fn...)
}

// Count 统计记录数量
func (o *Store) Count(filter QueryFilter, fn ...func(*CondOptions)) (uint64, error) {
	return Count(o, filter, fn...)
}

// Exists 检查记录是否存在
func (o *Store) Exists(filter QueryFilter, fn ...func(*CondOptions)) (bool, error) {
	total, err := Count(o, filter, fn...)
	return total > 0, err
}

// FindCols 查询指定字段值
func (o *Store) FindCols(field string, filter QueryFilter) (ztype.SliceType, error) {
	return findColsRaw(o, field, filter)
}

// Find 查询多条记录
func (o *Store) Find(filter QueryFilter, fn ...func(*CondOptions)) (ztype.Maps, error) {
	return Find[ztype.Map](o, filter, fn...)
}

// FindOne 查询单条记录
func (o *Store) FindOne(filter QueryFilter, fn ...func(*CondOptions)) (ztype.Map, error) {
	return FindOne[ztype.Map](o, filter, fn...)
}

// FindOneByID 根据 ID 查询单条记录
func (o *Store) FindOneByID(id any, fn ...func(*CondOptions)) (ztype.Map, error) {
	return FindOne[ztype.Map](o, ID(id), fn...)
}

// Pages 分页查询记录
func (o *Store) Pages(page, pagesize int, filter QueryFilter, fn ...func(*CondOptions)) (*PageData, error) {
	return pages(o.schema, page, pagesize, getFilter(o.schema, filter), true, fn...)
}

// Update 更新符合条件的记录
func (o *Store) Update(filter QueryFilter, data any, fn ...func(*CondOptions)) (total int64, err error) {
	return Update(o.schema, filter, data, fn...)
}

// UpdateMany 批量更新记录
func (o *Store) UpdateMany(filter QueryFilter, data any, fn ...func(*CondOptions)) (total int64, err error) {
	return UpdateMany(o.schema, filter, data, fn...)
}

// UpdateByID 根据 ID 更新记录
func (o *Store) UpdateByID(id any, data any, fn ...func(*CondOptions)) (total int64, err error) {
	return Update(o.schema, ID(id), data, fn...)
}

// Delete 删除符合条件的记录
func (o *Store) Delete(filter QueryFilter, fn ...func(*CondOptions)) (total int64, err error) {
	return Delete(o.schema, filter, fn...)
}

// DeleteMany 批量删除记录
func (o *Store) DeleteMany(filter QueryFilter, fn ...func(*CondOptions)) (total int64, err error) {
	return DeleteMany(o.schema, filter, fn...)
}

// DeleteByID 根据 ID 删除记录
func (o *Store) DeleteByID(id any, fn ...func(*CondOptions)) (total int64, err error) {
	return Delete(o.schema, ID(id), fn...)
}

// Repository 创建 Map 类型仓储
func (o *Store) Repository() *Repository[ztype.Map, QueryFilter, ztype.Map, ztype.Map] {
	return NewMapRepository(o)
}

// Repo 创建泛型仓储
func Repo[T any, F any, C any, U any](o *Store) *Repository[T, F, C, U] {
	return NewStructRepository[T, F, C, U](o)
}
