package restapi

import (
	"github.com/sohaha/zlsgo/ztype"
)

func (m *Model) Operation() *Operation {
	return &Operation{
		model: m,
	}
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

// FindCols 查询指定字段
func (o *Operation) FindCols(field string, filter ztype.Map) (ztype.SliceType, error) {
	return FindCols(o.model, field, filter, func(co *CondOptions) error {
		co.Fields = []string{field}
		return nil
	})
}

func (o *Operation) Find(filter ztype.Map, fn ...func(*CondOptions) error) (ztype.Maps, error) {
	return Find(o.model, filter, fn...)
}

func (o *Operation) FindOne(filter ztype.Map, fn ...func(*CondOptions) error) (ztype.Map, error) {
	return FindOne(o.model, filter, fn...)
}
