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
		if *m.define.Options.SoftDeleteIsTime {
			data[DeletedAtKey] = nil
		} else {
			data[DeletedAtKey] = 0
		}
	}
	data, err = m.valuesCryptProcess(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Insert 插入单条记录
func Insert[D any](m *Schema, data D, fn ...func(*InsertOptions)) (lastId interface{}, err error) {
	dataMap, err := dataToMap(data)
	if err != nil {
		return 0, err
	}
	dataMap, err = insertData(m, dataMap)
	if err != nil {
		return 0, err
	}

	// BeforeInsert hook
	if err = m.hook(hook.EventBeforeInsert, dataMap); err != nil {
		return 0, err
	}

	id, err := m.Storage.Insert(m.GetTableName(), dataMap, fn...)
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
	_ = m.hook(hook.EventAfterInsert, id, dataMap)

	return id, nil
}

// InsertMany 批量插入记录
func InsertMany[D any](m *Schema, datas D, fn ...func(*InsertOptions)) (lastIds []interface{}, err error) {
	dataMaps, err := dataToMaps(datas)
	if err != nil {
		return []interface{}{}, err
	}

	d := make(ztype.Maps, 0, len(dataMaps))
	for i := range dataMaps {
		data, err := insertData(m, dataMaps[i])
		if err != nil {
			return []interface{}{}, err
		}
		if !data.IsEmpty() {
			d = append(d, data)
		}
	}
	if len(d) == 0 {
		return []interface{}{}, nil
	}

	// BeforeInsert hook (批量数据)
	if err = m.hook(hook.EventBeforeInsert, d); err != nil {
		return []interface{}{}, err
	}

	lastIds, err = m.Storage.InsertMany(m.GetTableName(), d, fn...)
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
func Delete(m *Schema, filter QueryFilter, fn ...func(*CondOptions)) (int64, error) {
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
func DeleteMany(m *Schema, filter QueryFilter, fn ...func(*CondOptions)) (int64, error) {
	f := getFilter(m, filter)
	m.DeCrypt(f)

	// BeforeDelete hook
	if err := m.hook(hook.EventBeforeDelete, f); err != nil {
		return 0, err
	}

	if fields := cascadeFields(m); len(fields) > 0 {
		rows, err := m.Storage.Find(m.GetTableName(), f, func(so *CondOptions) {
			for i := range fn {
				if fn[i] != nil {
					fn[i](so)
				}
			}
			so.Fields = append(so.Fields[:0], fields...)
		})
		if err != nil {
			return 0, err
		}
		if err := cascadeDelete(m, rows); err != nil {
			return 0, err
		}
	}

	var total int64
	var err error

	if *m.define.Options.SoftDeletes {
		data := make(ztype.Map, 1)
		now := ztime.Time()
		if *m.define.Options.SoftDeleteIsTime {
			data[DeletedAtKey] = now
		} else {
			data[DeletedAtKey] = now.Unix()
		}
		total, err = m.Storage.Update(m.GetTableName(), data, f, fn...)
	} else {
		total, err = m.Storage.Delete(m.GetTableName(), f, fn...)
	}

	if err != nil {
		return 0, err
	}

	// AfterDelete hook (不返回错误，避免数据不一致)
	_ = m.hook(hook.EventAfterDelete, f, total)

	return total, nil
}

// Update 更新单条记录
func Update[D any](
	m *Schema,
	filter QueryFilter,
	data D,
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
func UpdateMany[D any](
	m *Schema,
	filter QueryFilter,
	data D,
	fn ...func(*CondOptions),
) (total int64, err error) {
	dataMap, err := dataToMap(data)
	if err != nil {
		return 0, err
	}
	dataMap = filterDate(dataMap, m.readOnlyKeys)
	dataMap, err = m.valuesBeforeProcess(dataMap)
	if err != nil {
		return 0, err
	}

	if len(m.GetDefineFields()) > 0 {
		dataMap, err = VerifiData(dataMap, m.GetDefineFields(), activeUpdate)
		if err != nil {
			return 0, errDataValidation(err)
		}
	}
	if *m.define.Options.Timestamps {
		dataMap[UpdatedAtKey] = ztime.Now()
	}
	dataMap, err = m.valuesCryptProcess(dataMap)
	if err != nil {
		return 0, err
	}

	f := getFilter(m, filter)

	if ok := m.DeCrypt(f); !ok {
		return 0, errDecryptionFailed(errors.New("data decryption failed"))
	}

	// BeforeUpdate hook
	if err = m.hook(hook.EventBeforeUpdate, f, dataMap); err != nil {
		return 0, err
	}

	total, err = m.Storage.Update(m.GetTableName(), dataMap, f, fn...)
	if err != nil {
		return 0, err
	}

	// AfterUpdate hook (不返回错误，避免数据不一致)
	_ = m.hook(hook.EventAfterUpdate, f, dataMap, total)

	return total, nil
}
