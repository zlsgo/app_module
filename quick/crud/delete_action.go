package crud

import (
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/quick/define"
	"github.com/zlsgo/app_module/quick/storage"
	"github.com/zlsgo/app_module/quick/utils"
)

func (crud *Crud) Delete(filter ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (int64, error) {
	return crud.DeleteMany(filter, func(so storage.CondOptions) storage.CondOptions {
		so = utils.Optional(so, fn...)
		so.Limit = 1
		return so
	})
}

func (crud *Crud) DeleteByID(id any, fn ...func(storage.CondOptions) storage.CondOptions) (int64, error) {
	return crud.DeleteMany(ztype.Map{define.Inside.IDKey(): id}, func(so storage.CondOptions) storage.CondOptions {
		so = utils.Optional(so, fn...)
		return so
	})
}

func (crud *Crud) DeleteMany(filter ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (int64, error) {
	f := getFilter(crud, filter)
	_ = crud.process.DeCrypt(f)
	if crud.define.Options.SoftDeletes {
		return crud.storage.Update(crud.tableName, ztype.Map{
			define.Inside.DeletedAtKey(): ztime.Time().Unix(),
		}, f)
	}

	return crud.storage.Delete(crud.tableName, f, fn...)
}
