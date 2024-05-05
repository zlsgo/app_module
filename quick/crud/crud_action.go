package crud

import (
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/quick/define"
	"github.com/zlsgo/app_module/quick/storage"
)

func (crud *Crud) InsertMany(data ztype.Maps) (lastId interface{}, err error) {
	return InsertMany(crud, data)
}

func (crud *Crud) Insert(data ztype.Map) (lastId interface{}, err error) {
	return Insert(crud, data)
}

func (crud *Crud) Update(filter ztype.Map, data ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (total int64, err error) {
	return Update(crud, filter, data, fn...)
}

func (crud *Crud) UpdateByID(id any, data ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (total int64, err error) {
	return Update(crud, ztype.Map{define.Inside.IDKey(): id}, data, fn...)
}

func (crud *Crud) UpdateMany(filter ztype.Map, data ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (total int64, err error) {
	return UpdateMany(crud, filter, data, fn...)
}
