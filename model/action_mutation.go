package model

import (
	"errors"

	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/hook"
)

// insertData 插入数据预处理
// 执行字段验证、默认值填充、时间戳设置等
func insertData(m *Schema, data ztype.Map) (ztype.Map, error) {
	data, err := m.valuesBeforeProcess(data)
	if err != nil {
		return nil, err
	}

	if len(m.GetDefineFields()) > 0 {
		data, err = VerifiData(data, m.GetDefineFields(), activeCreate)
		if err != nil {
			return nil, err
		}
	}

	if *m.define.Options.Timestamps {
		data[CreatedAtKey] = ztime.Now()
		data[UpdatedAtKey] = ztime.Now()
	}

	if *m.define.Options.SoftDeletes {
		if InsideOption.softDeleteIsTime {
			data[DeletedAtKey] = nil
		} else {
			data[DeletedAtKey] = 0
		}
	}
	return data, nil
}

// Insert 插入单条记录
func Insert(m *Schema, data ztype.Map, fn ...func(*InsertOptions)) (lastId interface{}, err error) {
	data, err = insertData(m, data)
	if err != nil {
		return 0, err
	}

	// BeforeInsert hook
	if err = m.hook(hook.EventBeforeInsert, data); err != nil {
		return 0, err
	}

	id, err := m.Storage.Insert(m.GetTableName(), m.GetFields(), data, fn...)
	if err != nil {
		return 0, err
	}

	if *m.define.Options.CryptID {
		id, err = m.EnCryptID(ztype.ToString(id))
		if err != nil {
			return 0, err
		}
	}

	// AfterInsert hook (不返回错误，避免数据不一致)
	_ = m.hook(hook.EventAfterInsert, id, data)

	return id, nil
}

// InsertMany 批量插入记录
func InsertMany(m *Schema, datas ztype.Maps, fn ...func(*InsertOptions)) (lastIds []interface{}, err error) {
	d := make(ztype.Maps, 0, len(datas))
	for i := range datas {
		data, err := insertData(m, datas[i])
		if err != nil {
			return []interface{}{}, err
		}
		if !data.IsEmpty() {
			d = append(d, data)
		}
	}

	// BeforeInsert hook (批量数据)
	if err = m.hook(hook.EventBeforeInsert, d); err != nil {
		return []interface{}{}, err
	}

	lastIds, err = m.Storage.InsertMany(m.GetTableName(), m.GetFields(), d, fn...)
	if err != nil {
		return []interface{}{}, err
	}

	if *m.define.Options.CryptID {
		for i := range lastIds {
			lastIds[i], err = m.EnCryptID(ztype.ToString(lastIds[i]))
			if err != nil {
				return []interface{}{}, err
			}
		}
	}

	// AfterInsert hook (批量数据，不返回错误，避免数据不一致)
	_ = m.hook(hook.EventAfterInsert, lastIds, d)

	return lastIds, nil
}

// Delete 删除单条记录
func Delete[T filter](m *Schema, filter T, fn ...func(*CondOptions)) (int64, error) {
	return DeleteMany(m, filter, func(so *CondOptions) {
		for i := range fn {
			if fn[i] != nil {
				fn[i](so)
			}
		}
		so.Limit = 1
	})
}

// DeleteMany 删除多条记录（支持软删除）
func DeleteMany[T filter](m *Schema, filter T, fn ...func(*CondOptions)) (int64, error) {
	f := getFilter(m, filter)
	m.DeCrypt(f)

	// BeforeDelete hook
	if err := m.hook(hook.EventBeforeDelete, f); err != nil {
		return 0, err
	}

	var total int64
	var err error

	if *m.define.Options.SoftDeletes {
		data := make(ztype.Map, 1)
		now := ztime.Time()
		if InsideOption.softDeleteIsTime {
			data[DeletedAtKey] = now
		} else {
			data[DeletedAtKey] = now.Unix()
		}
		total, err = m.Storage.Update(m.GetTableName(), m.GetFields(), data, f, fn...)
	} else {
		total, err = m.Storage.Delete(m.GetTableName(), m.GetFields(), f, fn...)
	}

	if err != nil {
		return 0, err
	}

	// AfterDelete hook (不返回错误，避免数据不一致)
	_ = m.hook(hook.EventAfterDelete, f, total)

	return total, nil
}

// Update 更新单条记录
func Update[T filter](
	m *Schema,
	filter T,
	data ztype.Map,
	fn ...func(*CondOptions),
) (total int64, err error) {
	return UpdateMany(m, filter, data, func(so *CondOptions) {
		for i := range fn {
			if fn[i] != nil {
				fn[i](so)
			}
		}
		so.Limit = 1
	})
}

// UpdateMany 更新多条记录
func UpdateMany[T filter](
	m *Schema,
	filter T,
	data ztype.Map,
	fn ...func(*CondOptions),
) (total int64, err error) {
	data = filterDate(data, m.readOnlyKeys)
	data, err = m.valuesBeforeProcess(data)
	if err != nil {
		return 0, err
	}

	if len(m.GetDefineFields()) > 0 {
		data, err = VerifiData(data, m.GetDefineFields(), activeUpdate)
		if err != nil {
			return 0, errDataValidation(err)
		}
	}
	if *m.define.Options.Timestamps {
		data[UpdatedAtKey] = ztime.Now()
	}

	f := getFilter(m, filter)

	if ok := m.DeCrypt(f); !ok {
		return 0, errDecryptionFailed(errors.New("data decryption failed"))
	}

	// BeforeUpdate hook
	if err = m.hook(hook.EventBeforeUpdate, f, data); err != nil {
		return 0, err
	}

	total, err = m.Storage.Update(m.GetTableName(), m.GetFields(), data, f, fn...)
	if err != nil {
		return 0, err
	}

	// AfterUpdate hook (不返回错误，避免数据不一致)
	_ = m.hook(hook.EventAfterUpdate, f, data, total)

	return total, nil
}
