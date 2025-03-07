package model

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
	"golang.org/x/exp/constraints"
)

type filter interface {
	ztype.Map | constraints.Integer | string | Filter
}

func getFilter[T filter](m *Schema, filter T) (filterMap ztype.Map) {
	f := (interface{})(filter)

	filterData, ok := f.(Filter)
	if ok {
		filterMap = ztype.Map(filterData)
	} else {
		filterMap, ok = f.(ztype.Map)
	}

	if !ok {
		idVal := f
		filterMap = ztype.Map{
			idKey: idVal,
		}
	} else if filterMap == nil {
		filterMap = ztype.Map{}
	}

	for k := range filterMap {
		k = zstring.TrimSpace(k)
		if k == "" || strings.Contains(k, placeHolder) || strings.Contains(k, ".") || strings.Contains(k, " ") {
			continue
		}
		if !zarray.Contains(m.GetFields(), k) {
			delete(filterMap, k)
		}
	}

	if *m.define.Options.SoftDeletes {
		if InsideOption.softDeleteIsTime {
			filterMap[DeletedAtKey] = nil
		} else {
			filterMap[DeletedAtKey] = 0
		}
	}

	return
}

type PageData struct {
	Items    ztype.Maps `json:"items"`
	Page     PageInfo   `json:"page"`
	pagesize uint       `json:"-"`
}

func (p *PageData) String() string {
	json, err := zjson.Marshal(p)
	if err != nil {
		return ""
	}
	return zstring.Bytes2String(zjson.Format(json))
}

func (p *PageData) Map(fn func(index int, item ztype.Map) ztype.Map, parallel ...uint) *PageData {
	if len(parallel) == 0 {
		parallel = []uint{p.pagesize}
	}
	p.Items = zarray.Map(p.Items, fn, parallel[0])

	return p
}

func Pages[T filter](
	m *Schema,
	page, pagesize int,
	filter T,
	fn ...func(*CondOptions),
) (*PageData, error) {
	return pages(m, page, pagesize, getFilter(m, filter), true, fn...)
}

func pages(
	m *Schema,
	page, pagesize int,
	filter ztype.Map,
	cryptId bool,
	fn ...func(*CondOptions),
) (pagedata *PageData, err error) {
	if cryptId {
		_ = m.DeCrypt(filter)
	}

	var (
		childRelationson map[string][]string
		foreignKeys      []string
		data             = &PageData{pagesize: uint(pagesize)}
	)

	rows, pages, err := m.Storage.Pages(
		m.GetTableName(),
		m.GetFields(),
		page,
		pagesize,
		filter,
		func(so *CondOptions) {
			if len(fn) > 0 {
				fn[0](so)
			}

			childRelationson, foreignKeys = relationson(m, so)
			if len(so.Fields) > 0 && len(so.Join) == 0 {
				so.Fields = m.filterFields(so.Fields)
			} else if len(so.Fields) == 0 {
				so.Fields = m.GetFields()
			} else if len(so.Fields) == 0 {
				so.Fields = allFields
			}
		},
	)

	data.Items = rows
	data.Page = pages

	if err != nil {
		return data, err
	}

	data.Items, err = handlerRelationson(m, data.Items, childRelationson, foreignKeys)
	if err != nil {
		return data, err
	}

	afterProcess := m.afterProcess
	if len(afterProcess) == 0 {
		return data, nil
	}

	for i := range data.Items {
		row := &data.Items[i]
		for k, v := range afterProcess {
			if _, ok := (*row)[k]; ok {
				(*row)[k], err = v[0](row.Get(k).String())
				if err != nil {
					return data, err
				}
			}
		}

		if cryptId && *m.define.Options.CryptID {
			_ = m.EnCrypt(row)
		}
	}

	return data, nil
}

var allFields = []string{"*"}

func relationson(
	m *Schema,
	so *CondOptions,
) (childRelationson map[string][]string, foreignKeys []string) {
	childRelationson = make(map[string][]string)
	includeAllFields := zarray.Contains(so.Fields, allFields[0])
	joinAs := zarray.Map(so.Join, func(_ int, v StorageJoin) string {
		return v.As
	})
	so.Fields = zarray.Filter(so.Fields, func(_ int, f string) bool {
		if f == allFields[0] {
			return true
		}
		field := strings.SplitN(f, ".", 2)
		// 如果字段是join的as，则直接返回true
		if zarray.Contains(joinAs, field[0]) {
			return true
		}
		isRelation := zarray.Contains(m.relationsKeys, field[0])
		if isRelation {
			if len(field) == 1 {
				childRelationson[field[0]] = m.define.Relations[field[0]].Fields
			} else {
				childRelationson[field[0]] = append(childRelationson[field[0]], field[1])
			}

			if !includeAllFields {
				for fki := range m.define.Relations[field[0]].SchemaKey {
					if !zarray.Contains(so.Fields, m.define.Relations[field[0]].ForeignKey[fki]) {
						foreignKeys = append(
							foreignKeys,
							m.define.Relations[field[0]].ForeignKey[fki],
						)
					}
				}
			}
		}

		return !isRelation
	})

	if len(foreignKeys) > 0 {
		so.Fields = append(so.Fields, foreignKeys...)
	}
	return
}

func relationsonValue(
	key string,
	typ schema.RelationType,
	fields []string,
	rows ztype.Maps,
) ztype.Maps {
	if len(rows) == 0 {
		return rows
	}

	m := ztype.Map{}
	for _, field := range fields {
		m[field] = nil
	}

	var value any
	switch typ {
	case schema.RelationSingle:
		value = m
	case schema.RelationMany:
		value = ztype.Maps{m}
	case schema.RelationSingleMerge:
		value = m
	}

	return zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
		row[key] = value
		return row
	}, 10)
}

func handlerRelationson(
	m *Schema,
	rows ztype.Maps,
	childRelationson map[string][]string,
	foreignKeys []string,
) (ztype.Maps, error) {
	for key := range childRelationson {
		d := m.define.Relations[key]
		m, ok := m.getSchema(d.Schema)
		if !ok {
			continue
		}

		ok = true
		schemaKeyLen, fields := len(d.SchemaKey), childRelationson[key]
		filter := make(ztype.Map, schemaKeyLen)
		for i := 0; i < schemaKeyLen; i++ {
			value := make([]any, 0, len(rows))
			repeat := make(map[any]struct{}, len(rows))

			for ir := range rows {
				v := rows[ir].Get(d.ForeignKey[i]).Value()
				if _, ok := repeat[v]; ok {
					continue
				}
				repeat[v] = struct{}{}
				value = append(value, v)
			}
			if len(value) > 0 {
				filter[d.SchemaKey[i]] = value
			}
		}

		if len(d.Filter) > 0 {
			for k := range d.Filter {
				filter[k] = d.Filter[k]
			}
		}

		if len(filter) == 0 {
			if d.Nullable {
				rows = relationsonValue(key, d.Type, fields, rows)
			}
			continue
		}

		tmpKeys := make([]string, 0, schemaKeyLen)
		items, err := find(m, getFilter(m, filter), false, func(co *CondOptions) {
			co.Fields = fields

			if len(co.Fields) == 0 {
				co.Fields = allFields
			} else {
				for i := 0; i < schemaKeyLen; i++ {
					ok = zarray.Contains(co.Fields, d.SchemaKey[i])
					if !ok {
						tmpKeys = append(tmpKeys, d.SchemaKey[i])
						co.Fields = append(co.Fields, d.SchemaKey[i])
					}
				}

				co.Fields = zarray.Unique(co.Fields)
			}

			// if d.Type != schema.RelationOneMerge {
			// 	co.Limit = 1
			// }
		})
		if err != nil {
			return nil, err
		}

		if len(items) == 0 {
			if d.Nullable {
				rows = relationsonValue(key, d.Type, fields, rows)
			}
			continue
		}

		switch d.Type {
		case schema.RelationSingle:
			rows = zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
				eq := false
				for i := range items {
					eqSum := 0
					for si := 0; si < schemaKeyLen; si++ {
						if items[i].Get(d.SchemaKey[si]).
							String() ==
							row.Get(d.ForeignKey[si]).
								String() {
							eqSum++
						}
					}
					if eqSum == schemaKeyLen {
						eq = true
						value := make(ztype.Map, len(items[i]))
						for k := range items[i] {
							if zarray.Contains(tmpKeys, k) {
								continue
							}
							value[k] = items[i][k]
						}
						row.Set(key, value)
						break
					}
				}
				if !eq {
					row.Set(key, ztype.Map{})
				}
				return row
			}, 10)

		case schema.RelationSingleMerge:
			rows = zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
				for i := range items {
					eqSum := 0
					for si := 0; si < schemaKeyLen; si++ {
						if items[i].Get(d.SchemaKey[si]).
							String() ==
							row.Get(d.ForeignKey[si]).
								String() {
							eqSum++
						}
					}
					if eqSum == schemaKeyLen {
						for k := range items[i] {
							if zarray.Contains(tmpKeys, k) {
								continue
							}
							row.Set(k, items[i][k])
						}
						break
					}
				}
				return row
			}, 10)
		case schema.RelationMany:
			rows = zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
				values := make(ztype.Maps, 0)
				for i := range items {
					eqSum := 0
					for si := 0; si < schemaKeyLen; si++ {
						if items[i].Get(d.SchemaKey[si]).
							String() ==
							row.Get(d.ForeignKey[si]).
								String() {
							eqSum++
						}
					}
					if eqSum == schemaKeyLen {
						value := make(ztype.Map, len(items[i]))
						for k := range items[i] {
							if zarray.Contains(tmpKeys, k) {
								continue
							}
							value[k] = items[i][k]
						}
						values = append(values, value)
					}
				}

				row.Set(key, values)
				return row
			}, 10)
		}
	}

	if len(foreignKeys) > 0 {
		rows = zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
			for i := range foreignKeys {
				delete(row, foreignKeys[i])
			}
			return row
		}, 10)
	}

	return rows, nil
}

func find(m *Schema, filter ztype.Map, cryptId bool, fn ...func(*CondOptions)) (rows ztype.Maps, err error) {
	if cryptId {
		_ = m.DeCrypt(filter)
	}

	var (
		childRelationson map[string][]string
		foreignKeys      []string
	)

	rows, err = m.Storage.Find(m.GetTableName(), m.GetFields(), filter, func(so *CondOptions) {
		for i := range fn {
			if fn[i] == nil {
				continue
			}
			fn[i](so)
		}

		childRelationson, foreignKeys = relationson(m, so)
		if len(so.Fields) > 0 && len(so.Join) == 0 {
			so.Fields = m.filterFields(so.Fields)
		} else if len(so.Fields) == 0 {
			so.Fields = m.GetFields()
		} else if len(so.Fields) == 0 {
			so.Fields = allFields
		}
	})
	if err != nil {
		return rows, err
	}

	rows, err = handlerRelationson(m, rows, childRelationson, foreignKeys)
	if err != nil {
		return rows, err
	}

	if len(m.afterProcess) > 0 {
		for i := range rows {
			row := &rows[i]
			for k, v := range m.afterProcess {
				if _, ok := (*row)[k]; ok {
					(*row)[k], err = v[0](row.Get(k).String())
					if err != nil {
						return nil, err
					}
				}
			}
			if cryptId && *m.define.Options.CryptID {
				m.EnCrypt(row)
			}
		}
	}

	return rows, nil
}

func Find[T filter](m *Schema, filter T, fn ...func(*CondOptions)) (ztype.Maps, error) {
	return find(m, getFilter(m, filter), true, fn...)
}

func FindOne[T filter](m *Schema, filter T, fn ...func(*CondOptions)) (ztype.Map, error) {
	rows, err := find(m, getFilter(m, filter), true, func(so *CondOptions) {
		for i := range fn {
			if fn[i] == nil {
				continue
			}
			fn[i](so)
		}

		// for _, v := range so.Relations {
		// 	r, ok := m.define.Relations[v]
		// 	if !ok {
		// 		continue
		// 	}

		// 	s, ok := m.getSchema(r.Model)
		// 	if !ok {
		// 		continue
		// 	}
		// 	switch r.Type {
		// 	case schema.RelationO2O:
		// 		so.Join = append(so.Join, StorageJoin{
		// 			Table: s.GetTableName(),
		// 			As:    s.GetName(),
		// 			Expr:  m.GetTableName() + "." + r.Foreign + "=" + s.GetName() + "." + r.Key,
		// 		})
		// 		if len(so.Fields) == 0 {
		// 			so.Fields = m.GetFields()
		// 		}
		// 		so.Fields = append(so.Fields, zarray.Map(r.Fields, func(_ int, v string) string {
		// 			return s.GetName() + "." + v
		// 		})...)
		// 	}
		// }

		so.Limit = 1
	})
	if err != nil {
		return ztype.Map{}, err
	}

	return rows.Index(0), nil
}

func FindCols[T filter](
	m *Schema,
	field string,
	filter T,
	fn ...func(*CondOptions),
) (ztype.SliceType, error) {
	rows, err := find(m, getFilter(m, filter), true, func(so *CondOptions) {
		if fn != nil {
			for i := range fn {
				fn[i](so)
			}
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

func FindCol[T filter](
	m *Schema,
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

func Count[T Filter](m *Schema, filter T, fn ...func(*CondOptions)) (uint64, error) {
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
		data[CreatedAtKey] = ztime.Time()
		data[UpdatedAtKey] = ztime.Time()
	}

	// if m.models.Options.CreatedBy {
	// 	data[CreatedByKey] = createdBy
	// }

	if *m.define.Options.SoftDeletes {
		if InsideOption.softDeleteIsTime {
			data[DeletedAtKey] = nil
		} else {
			data[DeletedAtKey] = 0
		}
	}
	return data, nil
}

func Insert(m *Schema, data ztype.Map, fn ...func(*InsertOptions)) (lastId interface{}, err error) {
	data, err = insertData(m, data)
	if err != nil {
		return 0, err
	}

	id, err := m.Storage.Insert(m.GetTableName(), m.GetFields(), data, fn...)
	if err == nil && *m.define.Options.CryptID {
		id, err = m.EnCryptID(ztype.ToString(id))
	}
	return id, err
}

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
	lastIds, err = m.Storage.InsertMany(m.GetTableName(), m.GetFields(), d, fn...)
	if err == nil && *m.define.Options.CryptID {
		for i := range lastIds {
			lastIds[i], err = m.EnCryptID(ztype.ToString(lastIds[i]))
		}
	}
	return
}

func Delete[T filter](m *Schema, filter T, fn ...func(*CondOptions)) (int64, error) {
	return DeleteMany(m, filter, func(so *CondOptions) {
		if fn != nil {
			for i := range fn {
				fn[i](so)
			}
		}
		so.Limit = 1
	})
}

func DeleteMany[T filter](m *Schema, filter T, fn ...func(*CondOptions)) (int64, error) {
	f := getFilter(m, filter)
	m.DeCrypt(f)
	if *m.define.Options.SoftDeletes {
		data := make(ztype.Map, 1)
		now := ztime.Time()
		if InsideOption.softDeleteIsTime {
			data[DeletedAtKey] = now
		} else {
			data[DeletedAtKey] = now.Unix()
		}
		return m.Storage.Update(m.GetTableName(), m.GetFields(), data, f)
	}

	return m.Storage.Delete(m.GetTableName(), m.GetFields(), f, fn...)
}

func Update[T filter](
	m *Schema,
	filter T,
	data ztype.Map,
	fn ...func(*CondOptions),
) (total int64, err error) {
	return UpdateMany(m, filter, data, func(so *CondOptions) {
		if fn != nil {
			for i := range fn {
				fn[i](so)
			}
		}
		so.Limit = 1
	})
}

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
			return 0, zerror.InvalidInput.Text(err.Error())
		}
	}
	if *m.define.Options.Timestamps {
		data[UpdatedAtKey] = ztime.Time()
	}

	f := getFilter(m, filter)

	if ok := m.DeCrypt(f); !ok {
		return 0, errors.New("data decryption failed")
	}

	return m.Storage.Update(m.GetTableName(), m.GetFields(), data, f, fn...)
}

//
// func Replica[T Filter](m *Model, filter T, data ztype.Var, fn ...func(*CondOptions)error) (total int64, err error) {
// 	return ReplicaMany(m, filter, data, func(so *CondOptions) error {
// 		if fn != nil {
// 			for i := range fn {
// 				if err := fn[i](so); err != nil {
// 					return err
// 				}
// 			}
// 		}
// 		so.Limit = 1
// 		return nil
// 	})
// }
//
// func ReplicaMany[T Filter](m *Model, filter T, data ztype.Var, fn ...func(*CondOptions)error) (total int64, err error) {
// 	data = filterDate(data, m.readOnlyKeys)
// 	data, err = m.valuesBeforeProcess(data)
// 	if err != nil {
// 		return 0, err
// 	}
// 	data, err = VerifiData(data, m.GetModelFields(), activeUpdate)
// 	if err != nil {
// 		return 0, err
// 	}
//
// 	if m.models.Options.Timestamps {
// 		data[UpdatedAtKey] = ztime.Time()
// 	}
//
// 	f := getFilter(m, filter)
// 	err = m.DeCrypt(f)
// 	if err != nil {
// 		return 0, err
// 	}
//
// 	return m.Storage.Replica(data, f, fn...)
// }
