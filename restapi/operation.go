package restapi

import (
	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/ztype"
)

func (m *Model) Operation() *Operation {
	return &Operation{
		model: m,
	}
}

type Operations struct {
	m *zarray.Maper[string, *Operation]
}

// Get 获取操作对象
func (m *Operations) Get(name string) (*Operation, bool) {
	return m.m.Get(name)
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

// Count 统计数量
func (o *Operation) Count(filter ztype.Map) (int64, error) {
	resp, err := FindCols(o.model, "count", filter, func(co *CondOptions) error {
		co.Fields = []string{"count(*) as count"}
		return nil
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
	return FindCols(o.model, field, filter, func(co *CondOptions) error {
		co.Fields = []string{field}
		return nil
	})
}

// Find 查询数据
func (o *Operation) Find(filter ztype.Map, fn ...func(*CondOptions) error) (ztype.Maps, error) {
	return Find(o.model, filter, fn...)
}

// FindOne 查询一条数据
func (o *Operation) FindOne(filter ztype.Map, fn ...func(*CondOptions) error) (ztype.Map, error) {
	return FindOne(o.model, filter, fn...)
}

// FindOneByID 通过ID查询
func (o *Operation) FindOneByID(id any, fn ...func(*CondOptions) error) (ztype.Map, error) {
	return FindOne(o.model, ztype.Map{IDKey: id}, fn...)
}

// Pages 分页查询
func (o *Operation) Pages(page, pagesize int, filter ztype.Map, fn ...func(*CondOptions) error) (*PageData, error) {
	return Pages(o.model, page, pagesize, filter, fn...)
}

// Update 更新数据
func (o *Operation) Update(filter ztype.Map, data ztype.Map, fn ...func(*CondOptions) error) (total int64, err error) {
	return Update(o.model, filter, data, fn...)
}

// UpdateMany 更新多条数据
func (o *Operation) UpdateMany(filter ztype.Map, data ztype.Map, fn ...func(*CondOptions) error) (total int64, err error) {
	return UpdateMany(o.model, filter, data, fn...)
}

// UpdateByID 更新数据
func (o *Operation) UpdateByID(id any, data ztype.Map, fn ...func(*CondOptions) error) (total int64, err error) {
	filter := ztype.Map{IDKey: id}
	return Update(o.model, filter, data, fn...)
}
