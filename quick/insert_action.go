package quick

import (
	"github.com/sohaha/zlsgo/ztype"
)

func Insert(m *Quick, data ztype.Map) (lastId interface{}, err error) {
	data, err = m.process.InsertData(m.define, data)
	if err != nil {
		return 0, err
	}

	id, err := m.storage.Insert(m.tableName, data)
	if err == nil && m.define.Options.CryptID {
		id, err = m.process.EnCryptID(ztype.ToString(id))
	}
	return id, err
}

func InsertMany(m *Quick, datas ztype.Maps) (lastIds []interface{}, err error) {
	for i := range datas {
		datas[i], err = m.process.InsertData(m.define, datas[i])
		if err != nil {
			return []interface{}{}, err
		}
	}

	lastIds, err = m.storage.InsertMany(m.tableName, datas)
	if err == nil && m.define.Options.CryptID {
		for i := range lastIds {
			lastIds[i], err = m.process.EnCryptID(ztype.ToString(lastIds[i]))
		}
	}
	return
}
