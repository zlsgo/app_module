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

type nestedRelationMap map[string]*nestedRelationNode

type nestedRelationNode struct {
	fields      []string
	children    nestedRelationMap
	loadDefault bool
}

func parseNestedRelationPath(path string) []string {
	path = strings.TrimSpace(path)
	if path == "" {
		return nil
	}

	raw := strings.Split(path, ".")
	parts := make([]string, 0, len(raw))
	for i := range raw {
		p := strings.TrimSpace(raw[i])
		if p == "" {
			continue
		}
		parts = append(parts, p)
	}

	return parts
}

func buildNestedRelationMap(m *Schema, relationFields []string) nestedRelationMap {
	relations := make(nestedRelationMap)
	if m == nil || len(relationFields) == 0 {
		return relations
	}

	for _, f := range relationFields {
		parts := parseNestedRelationPath(f)
		if len(parts) == 0 {
			continue
		}

		relName := parts[0]
		d, ok := m.define.Relations[relName]
		if !ok {
			continue
		}

		rootNode, ok := relations[relName]
		if !ok {
			rootNode = &nestedRelationNode{}
			relations[relName] = rootNode
		}

		if len(parts) == 1 {
			rootNode.loadDefault = true
			continue
		}

		childSchema, ok := m.getSchema(d.Schema)
		if !ok {
			field := strings.Join(parts[1:], ".")
			if field != "" {
				rootNode.fields = zarray.Unique(append(rootNode.fields, field))
			}
			continue
		}

		curSchema := childSchema
		curNode := rootNode

		for i := 1; i < len(parts); i++ {
			part := parts[i]

			nextRel, isRel := curSchema.define.Relations[part]
			if !isRel {
				field := strings.Join(parts[i:], ".")
				if field != "" {
					curNode.fields = zarray.Unique(append(curNode.fields, field))
				}
				break
			}

			if curNode.children == nil {
				curNode.children = make(nestedRelationMap)
			}

			nextNode, ok := curNode.children[part]
			if !ok {
				nextNode = &nestedRelationNode{}
				curNode.children[part] = nextNode
			}

			if i == len(parts)-1 {
				nextNode.loadDefault = true
				break
			}

			nextSchema, ok := curSchema.getSchema(nextRel.Schema)
			if !ok {
				field := strings.Join(parts[i+1:], ".")
				if field != "" {
					nextNode.fields = zarray.Unique(append(nextNode.fields, field))
				}
				break
			}

			curSchema = nextSchema
			curNode = nextNode
		}
	}

	return relations
}

func buildManyToManyFields(relatedSchema *Schema, d schema.Relation, node *nestedRelationNode) (fields []string, queryFields []string, tmpKeys []string) {
	fields = mergeRelationFields(d.Fields, node)
	queryFields = append([]string(nil), fields...)
	if len(queryFields) == 0 {
		queryFields = allFields
		return fields, queryFields, nil
	}

	fieldSet := make(map[string]struct{}, len(queryFields))
	for _, field := range queryFields {
		fieldSet[field] = struct{}{}
	}
	tmpSet := make(map[string]struct{})
	addKey := func(key string) {
		if key == "" {
			return
		}
		if _, ok := fieldSet[key]; ok {
			return
		}
		fieldSet[key] = struct{}{}
		queryFields = append(queryFields, key)
		if _, ok := tmpSet[key]; ok {
			return
		}
		tmpSet[key] = struct{}{}
		tmpKeys = append(tmpKeys, key)
	}

	relatedKeys := d.SchemaKey
	if len(relatedKeys) == 0 {
		relatedKeys = []string{idKey}
	}
	for _, k := range relatedKeys {
		addKey(k)
	}

	if node != nil && len(node.children) > 0 && relatedSchema != nil {
		for relName := range node.children {
			rel, ok := relatedSchema.define.Relations[relName]
			if !ok {
				continue
			}
			for _, fk := range rel.ForeignKey {
				addKey(fk)
			}
		}
	}

	return fields, queryFields, tmpKeys
}

func buildRelationFields(childSchema *Schema, d schema.Relation, node *nestedRelationNode) (fields []string, queryFields []string, tmpKeys []string) {
	fields = mergeRelationFields(d.Fields, node)
	queryFields = append([]string(nil), fields...)
	if len(queryFields) == 0 {
		queryFields = allFields
		return fields, queryFields, nil
	}

	fieldSet := make(map[string]struct{}, len(queryFields))
	for _, field := range queryFields {
		fieldSet[field] = struct{}{}
	}
	tmpSet := make(map[string]struct{})
	addKey := func(key string) {
		if key == "" {
			return
		}
		if _, ok := fieldSet[key]; ok {
			return
		}
		fieldSet[key] = struct{}{}
		queryFields = append(queryFields, key)
		if _, ok := tmpSet[key]; ok {
			return
		}
		tmpSet[key] = struct{}{}
		tmpKeys = append(tmpKeys, key)
	}

	for i := range d.SchemaKey {
		k := d.SchemaKey[i]
		addKey(k)
	}

	if node != nil && len(node.children) > 0 && childSchema != nil {
		for relName := range node.children {
			rel, ok := childSchema.define.Relations[relName]
			if !ok {
				continue
			}
			for _, fk := range rel.ForeignKey {
				addKey(fk)
			}
		}
	}

	return fields, queryFields, tmpKeys
}

func mergeRelationFields(base []string, node *nestedRelationNode) []string {
	if node == nil {
		return uniqueStrings(base)
	}
	if node.loadDefault || len(node.fields) == 0 {
		if len(node.fields) == 0 {
			return uniqueStrings(base)
		}
		merged := append(append([]string(nil), base...), node.fields...)
		return uniqueStrings(merged)
	}
	return uniqueStrings(node.fields)
}

func uniqueStrings(items []string) []string {
	if len(items) == 0 {
		return nil
	}
	set := make(map[string]struct{}, len(items))
	out := make([]string, 0, len(items))
	for _, item := range items {
		if item == "" {
			continue
		}
		if _, ok := set[item]; ok {
			continue
		}
		set[item] = struct{}{}
		out = append(out, item)
	}
	return out
}

// relationson 解析查询选项中的关联字段
// 返回需要加载的子关联和外键字段列表
func relationson(
	m *Schema,
	so *CondOptions,
) (childRelationson nestedRelationMap, foreignKeys []string) {
	childRelationson = make(nestedRelationMap)
	if m == nil || so == nil || len(so.Relations) == 0 {
		return childRelationson, nil
	}

	originFields := append([]string(nil), so.Fields...)
	explicitFields := len(originFields) > 0
	relationFields := make([]string, 0, len(so.Relations))
	originFieldSet := make(map[string]struct{}, len(originFields))
	if explicitFields {
		for _, field := range originFields {
			originFieldSet[field] = struct{}{}
		}
	}
	foreignKeySet := make(map[string]struct{})

	for _, f := range so.Relations {
		f = strings.TrimSpace(f)
		if f == "" {
			continue
		}

		field := strings.SplitN(f, ".", 2)
		d, ok := m.define.Relations[field[0]]
		if !ok {
			continue
		}

		relationFields = append(relationFields, f)
		if !explicitFields {
			continue
		}

		if d.Type == schema.RelationManyToMany {
			fk := d.ForeignKey
			if len(fk) == 0 {
				fk = []string{idKey}
			}
			for i := range fk {
				if _, ok := originFieldSet[fk[i]]; ok {
					continue
				}
				if _, ok := foreignKeySet[fk[i]]; ok {
					continue
				}
				foreignKeySet[fk[i]] = struct{}{}
				foreignKeys = append(foreignKeys, fk[i])
			}
			continue
		}

		for i := range d.ForeignKey {
			fk := d.ForeignKey[i]
			if _, ok := originFieldSet[fk]; ok {
				continue
			}
			if _, ok := foreignKeySet[fk]; ok {
				continue
			}
			foreignKeySet[fk] = struct{}{}
			foreignKeys = append(foreignKeys, fk)
		}
	}

	if len(foreignKeys) > 0 {
		so.Fields = uniqueStrings(append(so.Fields, foreignKeys...))
	}

	if len(relationFields) > 0 {
		childRelationson = buildNestedRelationMap(m, relationFields)
	}

	return childRelationson, foreignKeys
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

	var value any
	switch typ {
	case schema.RelationSingle, schema.RelationSingleMerge:
		if typ == schema.RelationSingleMerge {
			return rows
		}
		value = nil
	case schema.RelationMany, schema.RelationManyToMany:
		value = nil
	default:
		return rows
	}

	return zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
		row[key] = value
		return row
	}, 10)
}

func relationEmptyValue(typ schema.RelationType) (any, bool) {
	switch typ {
	case schema.RelationSingle:
		return ztype.Map{}, true
	case schema.RelationMany, schema.RelationManyToMany:
		return ztype.Maps{}, true
	case schema.RelationSingleMerge:
		return nil, false
	default:
		return nil, false
	}
}

// handlerRelationson 处理关联数据装载（核心函数）
// 根据配置的关联关系，查询关联表并将数据合并到主结果中
func handlerRelationson(
	m *Schema,
	rows ztype.Maps,
	childRelationson nestedRelationMap,
	foreignKeys []string,
) (ztype.Maps, error) {
	if m == nil || len(rows) == 0 {
		return rows, nil
	}

	for key, node := range childRelationson {
		d, ok := m.define.Relations[key]
		if !ok {
			continue
		}

		childSchema, ok := m.getSchema(d.Schema)
		if !ok {
			continue
		}

		if d.Type == schema.RelationManyToMany {
			fields, queryFields, tmpKeys := buildManyToManyFields(childSchema, d, node)
			m2m := NewManyToManyRelation(m, childSchema, d)
			loaded, err := m2m.Load(rows, key, queryFields)
			if err != nil {
				return nil, err
			}
			rows = loaded

			relatedRows := make(ztype.Maps, 0)
			for i := range rows {
				v := rows[i].Get(key).Value()
				switch vv := v.(type) {
				case ztype.Maps:
					relatedRows = append(relatedRows, vv...)
				case []ztype.Map:
					relatedRows = append(relatedRows, vv...)
				}
			}

			if node != nil && len(node.children) > 0 && len(relatedRows) > 0 {
				_, err := handlerRelationson(childSchema, relatedRows, node.children, tmpKeys)
				if err != nil {
					return nil, err
				}
			} else if len(tmpKeys) > 0 && len(relatedRows) > 0 {
				for i := range relatedRows {
					for j := range tmpKeys {
						delete(relatedRows[i], tmpKeys[j])
					}
				}
			}

			if d.Nullable && len(relatedRows) == 0 {
				rows = relationsonValue(key, d.Type, fields, rows)
			}

			continue
		}

		schemaKeyLen := len(d.SchemaKey)

		if len(d.ForeignKey) != schemaKeyLen {
			return nil, errRelationMismatch(errors.New("schema key and foreign key length mismatch"))
		}

		fields, queryFields, tmpKeys := buildRelationFields(childSchema, d, node)

		tuples := collectKeyTuples(rows, d.ForeignKey)
		filter := buildCompositeFilter(d.SchemaKey, tuples)

		if len(d.Filter) > 0 {
			if filter == nil {
				filter = ztype.Map{}
			}
			for k := range d.Filter {
				filter[k] = d.Filter[k]
			}
		}

		if len(filter) == 0 {
			if d.Nullable {
				rows = relationsonValue(key, d.Type, fields, rows)
			} else if emptyVal, ok := relationEmptyValue(d.Type); ok {
				rows = zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
					row[key] = emptyVal
					return row
				}, 10)
			}
			continue
		}

		items, err := findMaps(childSchema.Model(), getFilter(childSchema, Filter(filter)), false, func(co *CondOptions) {
			co.Fields = queryFields
		})
		if err != nil {
			return nil, err
		}

		if len(items) == 0 {
			if d.Nullable {
				rows = relationsonValue(key, d.Type, fields, rows)
			} else if emptyVal, ok := relationEmptyValue(d.Type); ok {
				rows = zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
					row[key] = emptyVal
					return row
				}, 10)
			}
			continue
		}

		tmpKeysMap := make(map[string]struct{}, len(tmpKeys))
		for _, k := range tmpKeys {
			tmpKeysMap[k] = struct{}{}
		}

		var itemsMapSingle map[string]int
		var itemsMapMany map[string][]int
		switch d.Type {
		case schema.RelationSingle, schema.RelationSingleMerge:
			itemsMapSingle = buildRelationMapSingle(items, d.SchemaKey)
		case schema.RelationMany:
			itemsMapMany = buildRelationMapMany(items, d.SchemaKey)
		}

		if node != nil && len(node.children) > 0 {
			items, err = handlerRelationson(childSchema, items, node.children, tmpKeys)
			if err != nil {
				return nil, err
			}
		}

		switch d.Type {
		case schema.RelationSingle:
			rows = zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
				mapKey := buildLookupKey(row, d.ForeignKey)
				if idx, ok := itemsMapSingle[mapKey]; ok && idx >= 0 && idx < len(items) {
					value := make(ztype.Map, len(items[idx]))
					for k := range items[idx] {
						if _, skip := tmpKeysMap[k]; skip {
							continue
						}
						value[k] = items[idx][k]
					}
					row.Set(key, value)
				} else {
					if d.Nullable {
						row.Set(key, nil)
					} else {
						row.Set(key, ztype.Map{})
					}
				}
				return row
			}, 10)

		case schema.RelationSingleMerge:
			rows = zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
				mapKey := buildLookupKey(row, d.ForeignKey)
				if idx, ok := itemsMapSingle[mapKey]; ok && idx >= 0 && idx < len(items) {
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
			rows = zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
				mapKey := buildLookupKey(row, d.ForeignKey)
				if indices, ok := itemsMapMany[mapKey]; ok {
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
					if d.Nullable {
						row.Set(key, nil)
					} else {
						row.Set(key, ztype.Maps{})
					}
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
