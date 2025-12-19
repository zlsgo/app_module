package model

import (
	"strings"

	"github.com/sohaha/zlsgo/ztype"
)

// findMaps 内部查询函数（返回 ztype.Maps）
// cryptId: 是否需要加密/解密 ID
func findMaps(m *Store, filter ztype.Map, cryptId bool, fn ...func(*CondOptions)) (resp ztype.Maps, err error) {
	if cryptId {
		_ = m.schema.DeCrypt(filter)
	}

	var (
		childRelationson map[string][]string
		foreignKeys      []string
	)

	resp, err = m.schema.Storage.Find(m.schema.GetTableName(), m.schema.GetFields(), filter, func(so *CondOptions) {
		for i := range fn {
			if fn[i] == nil {
				continue
			}
			fn[i](so)
		}

		childRelationson, foreignKeys = relationson(m.schema, so)
		if len(so.Fields) > 0 && len(so.Join) == 0 {
			so.Fields = m.schema.filterFields(so.Fields)
		} else if len(so.Fields) == 0 {
			so.Fields = m.schema.GetFields()
		}
	})
	if err != nil {
		return
	}

	resp, err = handlerRelationson(m.schema, resp, childRelationson, foreignKeys)
	if err != nil {
		return
	}

	if len(m.schema.afterProcess) > 0 {
		for i := range resp {
			row := &resp[i]
			for k, v := range m.schema.afterProcess {
				if _, ok := (*row)[k]; ok {
					(*row)[k], err = v[0](row.Get(k).String())
					if err != nil {
						return
					}
				}
			}
			if cryptId && *m.schema.define.Options.CryptID {
				m.schema.EnCrypt(row)
			}
		}
	}

	return resp, nil
}

// find 泛型查询函数（支持结构体映射）
func find[R any | ztype.Map](m *Store, filter ztype.Map, cryptId bool, fn ...func(*CondOptions)) (rows []R, err error) {
	resp, err := findMaps(m, filter, cryptId, fn...)
	if err != nil {
		return nil, err
	}

	var r R
	if _, ok := any(r).(ztype.Map); ok {
		rows = make([]R, len(resp))
		for i := range resp {
			rows[i] = any(resp[i]).(R)
		}
		return rows, nil
	}

	err = ztype.To(resp, &rows)
	return rows, err
}

// Find 查询多条记录（公开 API）
func Find[R any | ztype.Map, T filter | any](m *Store, filter T, fn ...func(*CondOptions)) ([]R, error) {
	return find[R](m, getFilter(m.schema, filter), true, fn...)
}

// FindMaps 查询多条记录（返回 ztype.Maps）
func FindMaps[T filter | any](m *Store, filter T, fn ...func(*CondOptions)) (ztype.Maps, error) {
	return Find[ztype.Map](m, filter, fn...)
}

// FindOne 查询单条记录
func FindOne[T filter | any](m *Store, filter T, fn ...func(*CondOptions)) (ztype.Map, error) {
	rows, err := findMaps(m, getFilter(m.schema, filter), true, func(so *CondOptions) {
		for i := range fn {
			if fn[i] == nil {
				continue
			}
			fn[i](so)
		}

		so.Limit = 1
	})
	if err != nil {
		return ztype.Map{}, err
	}

	return rows.Index(0), nil
}

// FindCols 查询指定列的值（返回数组）
func FindCols[T filter](
	m *Store,
	field string,
	filter T,
	fn ...func(*CondOptions),
) (ztype.SliceType, error) {
	rows, err := findMaps(m, getFilter(m.schema, filter), true, func(so *CondOptions) {
		for i := range fn {
			fn[i](so)
		}
		so.Fields = []string{field}
	})
	if err != nil {
		return ztype.SliceType{}, err
	}

	data := make(ztype.SliceType, rows.Len())
	f := strings.Split(field, " ")
	field = f[len(f)-1]
	for i := range rows {
		data[i] = rows[i].Get(field)
	}
	return data, nil
}

// FindCol 查询单个字段的单个值
func FindCol[T filter](
	m *Store,
	field string,
	filter T,
	fn ...func(*CondOptions),
) (ztype.Type, bool, error) {
	values, err := FindCols(m, field, filter, fn...)
	if err != nil || values.Len() == 0 {
		return ztype.Type{}, false, err
	}
	return values.First(), true, nil
}

// Count 统计记录数量
func Count[T Filter](m *Store, filter T, fn ...func(*CondOptions)) (uint64, error) {
	data, err := FindCols(m, "count(*) as count", filter, func(co *CondOptions) {
		for i := range fn {
			fn[i](co)
		}
		co.OrderBy = nil
	})
	if err != nil {
		return 0, err
	}
	return data.First().Uint64(), nil
}
