package model

import (
	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/ztype"
)

func (m *Model) Operation() *Operation {
	if m.operation == nil {
		m.operation = &Operation{
			model: m,
		}
	}
	return m.operation
}

type Operations struct {
	items *zarray.Maper[string, *Operation]
}

// Get 获取操作对象
func (m *Operations) Get(name string) (*Operation, bool) {
	return m.items.Get(name)
}

// MustGet 获取操作对象
func (m *Operations) MustGet(name string) *Operation {
	o, _ := m.items.Get(name)
	if o == nil {
		panic("operation " + name + " not found")
	}
	return o
}

// Model 获取模型
func (o *Operation) Model() *Model {
	return o.model
}

// EnCryptID 加密 ID
func (o *Operation) EnCryptID(id string) (nid string, err error) {
	return o.model.EnCryptID(id)
}

// DeCryptID 解密 ID
func (o *Operation) DeCryptID(nid string) (id string, err error) {
	return o.model.DeCryptID(nid)
}

// Insert 插入数据
func (o *Operation) Insert(data ztype.Map) (lastId interface{}, err error) {
	return Insert(o.model, data)
}

// InsertMany 批量插入数据
func (o *Operation) InsertMany(data ztype.Maps) (lastId interface{}, err error) {
	return InsertMany(o.model, data)
}

// Count 统计数量
func (o *Operation) Count(filter ztype.Map) (int64, error) {
	resp, err := FindCols(o.model, "count", filter, func(co *CondOptions) {
		co.Fields = []string{"count(*) as count"}
	})
	if err != nil {
		return 0, err
	}

	return resp.First().Int64(), nil
}

// Exists 数据是否存在
func (o *Operation) Exists(filter ztype.Map) (bool, error) {
	total, err := o.Count(filter)
	return total > 0, err
}

// FindCols 查询指定字段
func (o *Operation) FindCols(field string, filter ztype.Map) (ztype.SliceType, error) {
	return FindCols(o.model, field, filter, func(co *CondOptions) {
		co.Fields = []string{field}
	})
}

// Find 查询数据
func (o *Operation) Find(filter ztype.Map, fn ...func(*CondOptions)) (ztype.Maps, error) {
	return Find(o.model, filter, fn...)
}

// FindOne 查询一条数据
func (o *Operation) FindOne(filter ztype.Map, fn ...func(*CondOptions)) (ztype.Map, error) {
	return FindOne(o.model, filter, fn...)
}

// FindOneByID 通过ID查询
func (o *Operation) FindOneByID(id any, fn ...func(*CondOptions)) (ztype.Map, error) {
	return FindOne(o.model, ztype.Map{idKey: id}, fn...)
}

// Pages 分页查询
func (o *Operation) Pages(page, pagesize int, filter ztype.Map, fn ...func(*CondOptions)) (*PageData, error) {
	return Pages(o.model, page, pagesize, filter, fn...)
}

// Update 更新数据
func (o *Operation) Update(filter ztype.Map, data ztype.Map, fn ...func(*CondOptions)) (total int64, err error) {
	return Update(o.model, filter, data, fn...)
}

// UpdateMany 更新多条数据
func (o *Operation) UpdateMany(filter ztype.Map, data ztype.Map, fn ...func(*CondOptions)) (total int64, err error) {
	return UpdateMany(o.model, filter, data, fn...)
}

// UpdateByID 通过ID更新
func (o *Operation) UpdateByID(id any, data ztype.Map, fn ...func(*CondOptions)) (total int64, err error) {
	filter := ztype.Map{idKey: id}
	return Update(o.model, filter, data, fn...)
}

// Delete 删除数据
func (o *Operation) Delete(id any, filter ztype.Map, fn ...func(*CondOptions)) (total int64, err error) {
	return Delete(o.model, filter, fn...)
}

// DeleteMany 删除多条数据
func (o *Operation) DeleteMany(id any, filter ztype.Map, fn ...func(*CondOptions)) (total int64, err error) {
	return DeleteMany(o.model, filter, fn...)
}

// DeleteByID 通过ID删除数据
func (o *Operation) DeleteByID(id any, fn ...func(*CondOptions)) (total int64, err error) {
	filter := ztype.Map{idKey: id}
	return Delete(o.model, filter, fn...)
}
