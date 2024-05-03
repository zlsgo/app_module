package quick

import (
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/quick/define"
	"github.com/zlsgo/app_module/quick/storage"
)

func (q *Quick) InsertMany(data ztype.Maps) (lastId interface{}, err error) {
	return InsertMany(q, data)
}

func (q *Quick) Insert(data ztype.Map) (lastId interface{}, err error) {
	return Insert(q, data)
}

func (q *Quick) Update(filter ztype.Map, data ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (total int64, err error) {
	return Update(q, filter, data, fn...)
}

func (q *Quick) UpdateByID(id any, data ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (total int64, err error) {
	return Update(q, ztype.Map{define.Inside.IDKey(): id}, data, fn...)
}

func (q *Quick) UpdateMany(filter ztype.Map, data ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (total int64, err error) {
	return UpdateMany(q, filter, data, fn...)
}
