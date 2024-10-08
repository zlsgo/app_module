package model

import (
	"strconv"
	"unsafe"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb/builder"
)

type Filter ztype.Map

func NewFilter() Filter {
	return Filter{}
}

func (f Filter) Cond(fn func(*builder.BuildCond) string) Filter {
	f["$"+strconv.FormatInt(int64(uintptr(unsafe.Pointer(&fn))), 10)] = fn
	return f
}

func (f Filter) Set(field string, cond any) Filter {
	f[field] = cond
	return f
}

func (f Filter) Get(field string) ztype.Type {
	return ztype.New(f[field])
}

func (m *Schema) Model() *Model {
	if m.model == nil {
		m.model = &Model{
			schema: m,
		}
	}
	return m.model
}

// Models 快捷操作
type Models struct {
	items *zarray.Maper[string, *Model]
}

// Get 获取操作对象
func (m *Models) Get(name string) (*Model, bool) {
	return m.items.Get(name)
}

// MustGet 获取操作对象
func (m *Models) MustGet(name string) *Model {
	o, _ := m.items.Get(name)
	if o == nil {
		panic("operation " + name + " not found")
	}
	return o
}

// All 全部模型
func (m *Models) All() (models []*Model) {
	models = make([]*Model, 0, m.items.Len())

	m.items.ForEach(func(key string, value *Model) bool {
		models = append(models, value)
		return true
	})

	return
}

// Schema 获取模型
func (o *Model) Schema() *Schema {
	return o.schema
}

// Insert 插入数据
func (o *Model) Insert(data ztype.Map, fn ...func(*InsertOptions)) (lastId interface{}, err error) {
	return Insert(o.schema, data, fn...)
}

// InsertMany 批量插入数据
func (o *Model) InsertMany(data ztype.Maps, fn ...func(*InsertOptions)) (lastId interface{}, err error) {
	return InsertMany(o.schema, data, fn...)
}

// Count 统计数量
func (o *Model) Count(filter Filter, fn ...func(*CondOptions)) (uint64, error) {
	return Count(o.schema, filter, fn...)
}

// Exists 数据是否存在
func (o *Model) Exists(filter Filter, fn ...func(*CondOptions)) (bool, error) {
	total, err := Count(o.schema, filter, fn...)
	return total > 0, err
}

// FindCols 查询指定字段
func (o *Model) FindCols(field string, filter Filter) (ztype.SliceType, error) {
	return FindCols(o.schema, field, filter, func(co *CondOptions) {
		co.Fields = []string{field}
	})
}

// Find 查询数据
func (o *Model) Find(filter Filter, fn ...func(*CondOptions)) (ztype.Maps, error) {
	return Find(o.schema, filter, fn...)
}

// FindOne 查询一条数据
func (o *Model) FindOne(filter Filter, fn ...func(*CondOptions)) (ztype.Map, error) {
	return FindOne(o.schema, filter, fn...)
}

// FindOneByID 通过ID查询
func (o *Model) FindOneByID(id any, fn ...func(*CondOptions)) (ztype.Map, error) {
	return FindOne(o.schema, ztype.Map{idKey: id}, fn...)
}

// Pages 分页查询
func (o *Model) Pages(page, pagesize int, filter Filter, fn ...func(*CondOptions)) (*PageData, error) {
	return Pages(o.schema, page, pagesize, filter, fn...)
}

// Update 更新数据
func (o *Model) Update(filter Filter, data ztype.Map, fn ...func(*CondOptions)) (total int64, err error) {
	return Update(o.schema, filter, data, fn...)
}

// UpdateMany 更新多条数据
func (o *Model) UpdateMany(filter Filter, data ztype.Map, fn ...func(*CondOptions)) (total int64, err error) {
	return UpdateMany(o.schema, filter, data, fn...)
}

// UpdateByID 通过ID更新
func (o *Model) UpdateByID(id any, data ztype.Map, fn ...func(*CondOptions)) (total int64, err error) {
	filter := ztype.Map{idKey: id}
	return Update(o.schema, filter, data, fn...)
}

// Delete 删除数据
func (o *Model) Delete(filter Filter, fn ...func(*CondOptions)) (total int64, err error) {
	return Delete(o.schema, filter, fn...)
}

// DeleteMany 删除多条数据
func (o *Model) DeleteMany(filter Filter, fn ...func(*CondOptions)) (total int64, err error) {
	return DeleteMany(o.schema, filter, fn...)
}

// DeleteByID 通过ID删除数据
func (o *Model) DeleteByID(id any, fn ...func(*CondOptions)) (total int64, err error) {
	filter := ztype.Map{idKey: id}
	return Delete(o.schema, filter, fn...)
}
