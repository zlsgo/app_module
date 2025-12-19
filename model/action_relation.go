package model

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

// allFields 表示查询所有字段的标记
var allFields = []string{"*"}

// relationKeySeparator 关联键分隔符（使用不可见字符避免冲突）
const (
	relationKeySeparator = "\x1F"
)

// relationson 解析查询选项中的关联字段
// 返回需要加载的子关联和外键字段列表
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

// buildRelationMapSingle 构建关联关系的HashMap，用于O(1)查找
func buildRelationMapSingle(items ztype.Maps, schemaKeys []string) map[string]int {
	if len(items) == 0 || len(schemaKeys) == 0 {
		return make(map[string]int)
	}

	schemaKeyLen := len(schemaKeys)
	itemsMap := make(map[string]int, len(items))

	keyParts := make([]string, schemaKeyLen)

	for i := range items {
		for si := 0; si < schemaKeyLen; si++ {
			val := items[i].Get(schemaKeys[si])
			if val.Value() == nil {
				keyParts[si] = "<NULL>"
			} else {
				keyParts[si] = val.String()
			}
		}
		mapKey := strings.Join(keyParts, relationKeySeparator)
		itemsMap[mapKey] = i
	}

	return itemsMap
}

// buildRelationMapMany 构建一对多关联的HashMap
func buildRelationMapMany(items ztype.Maps, schemaKeys []string) map[string][]int {
	if len(items) == 0 || len(schemaKeys) == 0 {
		return make(map[string][]int)
	}

	schemaKeyLen := len(schemaKeys)
	itemsMap := make(map[string][]int, len(items))

	keyParts := make([]string, schemaKeyLen)

	for i := range items {
		for si := 0; si < schemaKeyLen; si++ {
			val := items[i].Get(schemaKeys[si])
			// 使用 Value() 检查 nil，nil 值使用特殊标记
			if val.Value() == nil {
				keyParts[si] = "<NULL>"
			} else {
				keyParts[si] = val.String()
			}
		}
		mapKey := strings.Join(keyParts, relationKeySeparator)
		itemsMap[mapKey] = append(itemsMap[mapKey], i)
	}

	return itemsMap
}

// buildLookupKey 根据外键构建查找key
func buildLookupKey(row ztype.Map, foreignKeys []string) string {
	if len(foreignKeys) == 0 {
		return ""
	}

	keyParts := make([]string, len(foreignKeys))
	for i, fk := range foreignKeys {
		val := row.Get(fk)
		// 使用 Value() 检查 nil，nil 值使用特殊标记
		if val.Value() == nil {
			keyParts[i] = "<NULL>"
		} else {
			keyParts[i] = val.String()
		}
	}
	return strings.Join(keyParts, relationKeySeparator)
}

// relationsonValue 为没有关联数据的记录生成默认值
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

// handlerRelationson 处理关联数据装载（核心函数）
// 根据配置的关联关系，查询关联表并将数据合并到主结果中
func handlerRelationson(
	m *Schema,
	rows ztype.Maps,
	childRelationson map[string][]string,
	foreignKeys []string,
) (ztype.Maps, error) {
	for key := range childRelationson {
		d := m.define.Relations[key]
		childSchema, ok := m.getSchema(d.Schema)
		if !ok {
			continue
		}

		schemaKeyLen, fields := len(d.SchemaKey), childRelationson[key]

		// 边界检查: 确保 SchemaKey 和 ForeignKey 长度一致
		if len(d.ForeignKey) != schemaKeyLen {
			return nil, errRelationMismatch(errors.New("schema key and foreign key length mismatch"))
		}

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
		items, err := findMaps(childSchema.Model(), getFilter(childSchema, filter), false, func(co *CondOptions) {
			co.Fields = fields

			if len(co.Fields) == 0 {
				co.Fields = allFields
			} else {
				for i := 0; i < schemaKeyLen; i++ {
					if !zarray.Contains(co.Fields, d.SchemaKey[i]) {
						tmpKeys = append(tmpKeys, d.SchemaKey[i])
						co.Fields = append(co.Fields, d.SchemaKey[i])
					}
				}

				co.Fields = zarray.Unique(co.Fields)
			}
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

		tmpKeysMap := make(map[string]struct{}, len(tmpKeys))
		for _, k := range tmpKeys {
			tmpKeysMap[k] = struct{}{}
		}

		switch d.Type {
		case schema.RelationSingle:
			itemsMap := buildRelationMapSingle(items, d.SchemaKey)
			rows = zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
				mapKey := buildLookupKey(row, d.ForeignKey)
				if idx, ok := itemsMap[mapKey]; ok && idx >= 0 && idx < len(items) {
					value := make(ztype.Map, len(items[idx]))
					for k := range items[idx] {
						if _, skip := tmpKeysMap[k]; skip {
							continue
						}
						value[k] = items[idx][k]
					}
					row.Set(key, value)
				} else {
					row.Set(key, ztype.Map{})
				}
				return row
			}, 10)

		case schema.RelationSingleMerge:
			itemsMap := buildRelationMapSingle(items, d.SchemaKey)
			rows = zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
				mapKey := buildLookupKey(row, d.ForeignKey)
				if idx, ok := itemsMap[mapKey]; ok && idx >= 0 && idx < len(items) {
					for k := range items[idx] {
						if _, skip := tmpKeysMap[k]; skip {
							continue
						}
						row.Set(k, items[idx][k])
					}
				}
				return row
			}, 10)

		case schema.RelationMany:
			itemsMap := buildRelationMapMany(items, d.SchemaKey)
			rows = zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
				mapKey := buildLookupKey(row, d.ForeignKey)
				if indices, ok := itemsMap[mapKey]; ok {
					values := make(ztype.Maps, 0, len(indices))
					for _, idx := range indices {
						if idx >= 0 && idx < len(items) {
							value := make(ztype.Map, len(items[idx]))
							for k := range items[idx] {
								if _, skip := tmpKeysMap[k]; skip {
									continue
								}
								value[k] = items[idx][k]
							}
							values = append(values, value)
						}
					}
					row.Set(key, values)
				} else {
					row.Set(key, ztype.Maps{})
				}
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
