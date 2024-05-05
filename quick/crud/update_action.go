package crud

import (
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/quick/define"
	"github.com/zlsgo/app_module/quick/process"
	"github.com/zlsgo/app_module/quick/storage"
	"github.com/zlsgo/app_module/quick/utils"
)

func Update[T Filter](m *Crud, filter T, data ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (total int64, err error) {
	return UpdateMany(m, filter, data, func(so storage.CondOptions) storage.CondOptions {
		so = utils.Optional(so, fn...)
		so.Limit = 1
		return so
	})
}

func UpdateMany[T Filter](m *Crud, filter T, data ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (total int64, err error) {
	data = m.process.FilterDate(data, m.readOnlyKeys)
	data, err = m.process.ValuesBeforeProcess(data)
	if err != nil {
		return 0, zerror.With(err, "data preprocessing failed")
	}

	data, err = m.process.VerifiData(data, process.ActiveUpdate)
	if err != nil {
		return 0, zerror.With(err, "data verification failed")
	}

	if m.define.Options.Timestamps {
		data[define.Inside.UpdatedAtKey()] = ztime.Time()
	}

	f := getFilter(m, filter)
	err = m.process.DeCrypt(f)
	if err != nil {
		return 0, zerror.With(err, "data decryption failed")
	}

	return m.storage.Update(m.tableName, data, f, fn...)
}
