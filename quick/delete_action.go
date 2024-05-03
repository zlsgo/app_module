package quick

import (
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/quick/define"
	"github.com/zlsgo/app_module/quick/storage"
	"github.com/zlsgo/app_module/quick/utils"
)

func (m *Quick) Delete(filter ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (int64, error) {
	return m.DeleteMany(filter, func(so storage.CondOptions) storage.CondOptions {
		so = utils.Optional(so, fn...)
		so.Limit = 1
		return so
	})
}

func (m *Quick) DeleteMany(filter ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (int64, error) {
	f := getFilter(m, filter)
	_ = m.process.DeCrypt(f)
	if m.define.Options.SoftDeletes {
		return m.storage.Update(m.tableName, ztype.Map{
			define.Inside.DeletedAtKey(): ztime.Time().Unix(),
		}, f)
	}

	return m.storage.Delete(m.tableName, f, fn...)
}
