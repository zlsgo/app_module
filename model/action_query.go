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
		childRelationson nestedRelationMap
		foreignKeys      []string
	)

	resp, err = m.schema.Storage.Find(m.schema.GetTableName(), filter, func(so *CondOptions) {
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
				val, ok := (*row)[k]
				if !ok {
					continue
				}
				for j := range v {
					val, err = v[j](val)
					if err != nil {
						return
					}
				}
				(*row)[k] = val
			}
			if cryptId && *m.schema.define.Options.CryptID {
				m.schema.EnCrypt(row)
			}
		}
	}

	return resp, nil
}

// find 泛型查询函数（支持结构体映射）
func find[R any](m *Store, filter ztype.Map, cryptId bool, fn ...func(*CondOptions)) (rows []R, err error) {
	resp, err := findMaps(m, filter, cryptId, fn...)
	if err != nil {
		return nil, err
	}

	return mapRows[R](resp)
}

// Find 查询多条记录（公开 API）
func Find[R any](m *Store, filter QueryFilter, fn ...func(*CondOptions)) ([]R, error) {
	return find[R](m, getFilter(m.schema, filter), true, fn...)
}

// FindMaps 查询多条记录（返回 ztype.Maps）
func FindMaps(m *Store, filter QueryFilter, fn ...func(*CondOptions)) (ztype.Maps, error) {
	return Find[ztype.Map](m, filter, fn...)
}

// FindOne 查询单条记录
func FindOne[R any](m *Store, filter QueryFilter, fn ...func(*CondOptions)) (R, error) {
	var zero R
	rows, err := find[R](m, getFilter(m.schema, filter), true, func(so *CondOptions) {
		for i := range fn {
			if fn[i] == nil {
				continue
			}
			fn[i](so)
		}
		so.Limit = 1
	})
	if err != nil {
		return zero, err
	}
	if len(rows) == 0 {
		return zero, ErrNoRecord
	}
	return rows[0], nil
}

func findColsRaw(
	m *Store,
	field string,
	filter QueryFilter,
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

// FindCols 查询指定列的值（返回数组）
func FindCols[T any](m *Store, field string, filter QueryFilter, fn ...func(*CondOptions)) ([]T, error) {
	values, err := findColsRaw(m, field, filter, fn...)
	if err != nil {
		return nil, err
	}
	return mapValues[T](values)
}

// FindCol 查询单个字段的单个值
func FindCol[T any](
	m *Store,
	field string,
	filter QueryFilter,
	fn ...func(*CondOptions),
) (T, bool, error) {
	var zero T
	values, err := findColsRaw(m, field, filter, fn...)
	if err != nil || values.Len() == 0 {
		return zero, false, err
	}
	value, err := mapValue[T](values.First())
	if err != nil {
		return zero, true, err
	}
	return value, true, nil
}

// Count 统计记录数量
func Count(m *Store, filter QueryFilter, fn ...func(*CondOptions)) (uint64, error) {
	data, err := FindCols[uint64](m, "count(*) as count", filter, func(co *CondOptions) {
		for i := range fn {
			fn[i](co)
		}
		co.OrderBy = nil
	})
	if err != nil {
		return 0, err
	}
	if len(data) == 0 {
		return 0, nil
	}
	return data[0], nil
}

func mapRows[R any](resp ztype.Maps) ([]R, error) {
	var rows []R
	var r R
	if _, ok := any(r).(ztype.Map); ok {
		rows = make([]R, len(resp))
		for i := range resp {
			rows[i] = any(resp[i]).(R)
		}
		return rows, nil
	}

	err := ztype.To(resp, &rows)
	return rows, err
}

func mapValues[T any](values ztype.SliceType) ([]T, error) {
	if values.Len() == 0 {
		return []T{}, nil
	}

	var zero T
	switch any(zero).(type) {
	case ztype.Type:
		result := make([]T, values.Len())
		for i := range values {
			result[i] = any(values[i]).(T)
		}
		return result, nil
	}

	raw := values.Value()
	if len(raw) == 0 {
		return []T{}, nil
	}

	var result []T
	err := ztype.To(raw, &result)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return []T{}, nil
	}
	return result, nil
}

func mapValue[T any](value ztype.Type) (T, error) {
	var out T
	switch any(out).(type) {
	case ztype.Type:
		return any(value).(T), nil
	}
	err := ztype.To(value.Value(), &out)
	return out, err
}
