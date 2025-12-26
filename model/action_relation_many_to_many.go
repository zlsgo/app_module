package model

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/builder"
	zdbschema "github.com/zlsgo/zdb/schema"
)

type ManyToManyRelation struct {
	schema   *Schema
	related  *Schema
	relation schema.Relation
	pivot    *PivotManager
}

func NewManyToManyRelation(parent *Schema, related *Schema, relation schema.Relation) *ManyToManyRelation {
	return &ManyToManyRelation{
		schema:   parent,
		related:  related,
		relation: relation,
		pivot:    NewPivotManager(parent),
	}
}

func (mtr *ManyToManyRelation) Load(rows ztype.Maps, key string, fields []string) (ztype.Maps, error) {
	parentKeys := mtr.parentKeyFields()
	relatedKeys := mtr.relatedKeyFields()
	pivotForeignKeys := mtr.relation.PivotKeys.Foreign
	pivotRelatedKeys := mtr.relation.PivotKeys.Related

	if len(pivotForeignKeys) == 0 || len(pivotRelatedKeys) == 0 {
		return nil, errRelationMismatch(errors.New("pivot keys not configured"))
	}
	if len(parentKeys) != len(pivotForeignKeys) {
		return nil, errRelationMismatch(errors.New("parent key and pivot foreign key length mismatch"))
	}
	if len(relatedKeys) != len(pivotRelatedKeys) {
		return nil, errRelationMismatch(errors.New("related key and pivot related key length mismatch"))
	}

	if err := mtr.pivot.SyncPivotSchema(&mtr.relation); err != nil {
		return nil, err
	}

	pivotTable, err := mtr.pivot.GetPivotTableName(&mtr.relation)
	if err != nil {
		return nil, err
	}

	parentTuples := collectKeyTuples(rows, parentKeys)
	pivotFilter := buildCompositeFilter(pivotForeignKeys, parentTuples)

	if len(mtr.relation.PivotFilter) > 0 {
		if pivotFilter == nil {
			pivotFilter = ztype.Map{}
		}
		for k := range mtr.relation.PivotFilter {
			pivotFilter[k] = mtr.relation.PivotFilter[k]
		}
	}

	if len(pivotFilter) == 0 {
		if mtr.relation.Nullable {
			return relationsonValue(key, schema.RelationManyToMany, fields, rows), nil
		}
		return zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
			row.Set(key, ztype.Maps{})
			return row
		}, 10), nil
	}

	pivotSelect := zarray.Unique(append(append([]string{}, pivotForeignKeys...), pivotRelatedKeys...))
	if len(mtr.relation.PivotFields) > 0 {
		pivotSelect = zarray.Unique(append(pivotSelect, mtr.relation.PivotFields...))
	}

	pivotRows, err := mtr.schema.Storage.Find(pivotTable, pivotFilter, func(co *CondOptions) {
		co.Fields = append(co.Fields[:0], pivotSelect...)
	})
	if err != nil {
		return nil, err
	}
	if len(pivotRows) == 0 {
		if mtr.relation.Nullable {
			return relationsonValue(key, schema.RelationManyToMany, fields, rows), nil
		}
		return zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
			row.Set(key, ztype.Maps{})
			return row
		}, 10), nil
	}

	type pivotRecord struct {
		relatedLookup string
		extra         ztype.Map
	}

	pivotMap := make(map[string][]pivotRecord)

	for _, pr := range pivotRows {
		parentLookup, ok := buildLookupKeyStrict(pr, pivotForeignKeys)
		if !ok {
			continue
		}
		relatedLookup, ok := buildLookupKeyStrict(pr, pivotRelatedKeys)
		if !ok {
			continue
		}

		extra := make(ztype.Map, len(mtr.relation.PivotFields))
		for _, f := range mtr.relation.PivotFields {
			extra["pivot_"+f] = pr.Get(f).Value()
		}

		pivotMap[parentLookup] = append(pivotMap[parentLookup], pivotRecord{
			relatedLookup: relatedLookup,
			extra:         extra,
		})

	}

	relatedTuples := collectKeyTuples(pivotRows, pivotRelatedKeys)
	relatedFilter := buildCompositeFilter(relatedKeys, relatedTuples)
	if len(mtr.relation.Filter) > 0 {
		if relatedFilter == nil {
			relatedFilter = ztype.Map{}
		}
		for k := range mtr.relation.Filter {
			relatedFilter[k] = mtr.relation.Filter[k]
		}
	}
	if len(relatedFilter) == 0 {
		if mtr.relation.Nullable {
			return relationsonValue(key, schema.RelationManyToMany, fields, rows), nil
		}
		return zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
			row.Set(key, ztype.Maps{})
			return row
		}, 10), nil
	}

	tmpKeys := make([]string, 0, len(relatedKeys))
	items, err := findMaps(mtr.related.Model(), getFilter(mtr.related, Filter(relatedFilter)), false, func(co *CondOptions) {
		co.Fields = fields
		if len(co.Fields) == 0 {
			co.Fields = allFields
		} else {
			for i := range relatedKeys {
				if !zarray.Contains(co.Fields, relatedKeys[i]) {
					tmpKeys = append(tmpKeys, relatedKeys[i])
					co.Fields = append(co.Fields, relatedKeys[i])
				}
			}
			co.Fields = zarray.Unique(co.Fields)
		}
	})
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		if mtr.relation.Nullable {
			return relationsonValue(key, schema.RelationManyToMany, fields, rows), nil
		}
		return zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
			row.Set(key, ztype.Maps{})
			return row
		}, 10), nil
	}

	tmpKeysMap := make(map[string]struct{}, len(tmpKeys))
	for _, k := range tmpKeys {
		tmpKeysMap[k] = struct{}{}
	}

	itemsMap := buildRelationMapSingle(items, relatedKeys)
	rows = zarray.Map(rows, func(_ int, row ztype.Map) ztype.Map {
		parentLookup, ok := buildLookupKeyStrict(row, parentKeys)
		if !ok {
			if mtr.relation.Nullable {
				row.Set(key, nil)
			} else {
				row.Set(key, ztype.Maps{})
			}
			return row
		}

		records := pivotMap[parentLookup]
		if len(records) == 0 {
			if mtr.relation.Nullable {
				row.Set(key, nil)
			} else {
				row.Set(key, ztype.Maps{})
			}
			return row
		}

		values := make(ztype.Maps, 0, len(records))
		for _, r := range records {
			idx, ok := itemsMap[r.relatedLookup]
			if !ok || idx < 0 || idx >= len(items) {
				continue
			}

			value := make(ztype.Map, len(items[idx])+len(r.extra))
			for k := range items[idx] {
				if _, skip := tmpKeysMap[k]; skip {
					continue
				}
				value[k] = items[idx][k]
			}
			for k := range r.extra {
				value[k] = r.extra[k]
			}
			values = append(values, value)
		}

		if len(values) == 0 {
			if mtr.relation.Nullable {
				row.Set(key, nil)
			} else {
				row.Set(key, ztype.Maps{})
			}
			return row
		}
		row.Set(key, values)
		return row
	}, 10)

	return rows, nil
}

func (mtr *ManyToManyRelation) parentKeyFields() []string {
	if len(mtr.relation.ForeignKey) > 0 {
		return mtr.relation.ForeignKey
	}
	return []string{idKey}
}

func (mtr *ManyToManyRelation) relatedKeyFields() []string {
	if len(mtr.relation.SchemaKey) > 0 {
		return mtr.relation.SchemaKey
	}
	return []string{idKey}
}

func buildLookupKeyStrict(row ztype.Map, keys []string) (string, bool) {
	if len(keys) == 0 {
		return "", false
	}

	keyParts := make([]string, len(keys))
	for i, fk := range keys {
		val := row.Get(fk)
		if val.Value() == nil {
			return "", false
		}
		keyParts[i] = val.String()
	}
	return strings.Join(keyParts, relationKeySeparator), true
}

type PivotManager struct {
	schema *Schema
}

func NewPivotManager(schema *Schema) *PivotManager {
	return &PivotManager{schema: schema}
}

func (pm *PivotManager) GetPivotTableName(relation *schema.Relation) (string, error) {
	if relation.PivotTable != "" {
		return pm.schema.tablePrefix + relation.PivotTable, nil
	}
	relatedSchema, ok := pm.schema.getSchema(relation.Schema)
	if !ok {
		return "", errRelationMismatch(errors.New("related schema not found"))
	}

	tables := []string{pm.schema.define.Table.Name, relatedSchema.define.Table.Name}
	if strings.Compare(tables[0], tables[1]) > 0 {
		tables[0], tables[1] = tables[1], tables[0]
	}
	return pm.schema.tablePrefix + strings.Join(tables, "_"), nil
}

func (pm *PivotManager) CreatePivotTable(relation *schema.Relation) error {
	db, err := pm.getDB()
	if err != nil {
		return err
	}
	pivotTable, err := pm.GetPivotTableName(relation)
	if err != nil {
		return err
	}

	foreign := relation.PivotKeys.Foreign
	related := relation.PivotKeys.Related
	if len(foreign) == 0 || len(related) == 0 {
		return errRelationMismatch(errors.New("pivot keys not configured"))
	}

	parentKeys := relation.ForeignKey
	if len(parentKeys) == 0 {
		parentKeys = []string{idKey}
	}
	relatedKeys := relation.SchemaKey
	if len(relatedKeys) == 0 {
		relatedKeys = []string{idKey}
	}

	create := builder.NewTable(pivotTable).Create().IfNotExists()
	create.SetDriver(db.GetDriver())

	fields := make([]*zdbschema.Field, 0, len(foreign)+len(related)+len(relation.PivotFields))
	for i, name := range foreign {
		fields = append(fields, zdbschema.NewField(name, pm.pivotKeyType(relation, parentKeys, relatedKeys, true, i)))
	}
	for i, name := range related {
		fields = append(fields, zdbschema.NewField(name, pm.pivotKeyType(relation, parentKeys, relatedKeys, false, i)))
	}
	for _, name := range relation.PivotFields {
		fields = append(fields, zdbschema.NewField(name, pm.pivotFieldType(name), func(f *zdbschema.Field) {
			f.NotNull = false
		}))
	}

	fields = pm.uniquePivotFields(fields)
	create.Column(fields...)

	sql, values, err := create.Build()
	if err != nil {
		return err
	}
	_, err = db.Exec(sql, values...)
	return err
}

func (pm *PivotManager) SyncPivotSchema(relation *schema.Relation) error {
	db, err := pm.getDB()
	if err != nil {
		return err
	}

	if err := pm.CreatePivotTable(relation); err != nil {
		return err
	}

	pivotTable, err := pm.GetPivotTableName(relation)
	if err != nil {
		return err
	}

	table := builder.NewTable(pivotTable)
	table.SetDriver(db.GetDriver())
	sql, values, process := table.GetColumn()
	res, err := db.QueryToMaps(sql, values...)
	if err != nil {
		return err
	}
	columns := process(res)

	required := make([]string, 0, len(relation.PivotKeys.Foreign)+len(relation.PivotKeys.Related)+len(relation.PivotFields))
	required = append(required, relation.PivotKeys.Foreign...)
	required = append(required, relation.PivotKeys.Related...)
	required = append(required, relation.PivotFields...)
	required = zarray.Unique(required)

	parentKeys := relation.ForeignKey
	if len(parentKeys) == 0 {
		parentKeys = []string{idKey}
	}
	relatedKeys := relation.SchemaKey
	if len(relatedKeys) == 0 {
		relatedKeys = []string{idKey}
	}

	for _, col := range required {
		if _, ok := columns[col]; ok {
			continue
		}
		sql, values = table.AddColumn(col, pm.pivotColumnType(relation, parentKeys, relatedKeys, col), func(f *zdbschema.Field) {
			if zarray.Contains(relation.PivotKeys.Foreign, col) || zarray.Contains(relation.PivotKeys.Related, col) {
				f.NotNull = true
				return
			}
			f.NotNull = false
		})
		_, err = db.Exec(sql, values...)
		if err != nil {
			return err
		}
	}

	return pm.ensurePivotIndexes(db, pivotTable, relation)
}

func (pm *PivotManager) ensurePivotIndexes(db *zdb.DB, pivotTable string, relation *schema.Relation) error {
	table := builder.NewTable(pivotTable).Create()
	table.SetDriver(db.GetDriver())

	foreign := relation.PivotKeys.Foreign
	related := relation.PivotKeys.Related

	uniqueIndex := pivotTable + "__u__" + strings.Join(append(append([]string{}, foreign...), related...), "_")
	sql, values, process := table.HasIndex(uniqueIndex)
	res, err := db.QueryToMaps(sql, values...)
	if err != nil {
		return err
	}
	if !process(res) {
		sql, values = table.CreateIndex(uniqueIndex, append(append([]string{}, foreign...), related...), "UNIQUE")
		_, err = db.Exec(sql, values...)
		if err != nil {
			return err
		}
	}

	foreignIndex := pivotTable + "__i__" + strings.Join(foreign, "_")
	sql, values, process = table.HasIndex(foreignIndex)
	res, err = db.QueryToMaps(sql, values...)
	if err != nil {
		return err
	}
	if !process(res) {
		sql, values = table.CreateIndex(foreignIndex, foreign, "")
		_, err = db.Exec(sql, values...)
		if err != nil {
			return err
		}
	}

	relatedIndex := pivotTable + "__i__" + strings.Join(related, "_")
	sql, values, process = table.HasIndex(relatedIndex)
	res, err = db.QueryToMaps(sql, values...)
	if err != nil {
		return err
	}
	if !process(res) {
		sql, values = table.CreateIndex(relatedIndex, related, "")
		_, err = db.Exec(sql, values...)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pm *PivotManager) pivotFieldType(name string) zdbschema.DataType {
	if name == CreatedAtKey || name == UpdatedAtKey || strings.HasSuffix(name, "_at") {
		return zdbschema.Time
	}
	if strings.HasSuffix(name, "_id") {
		return zdbschema.Uint
	}
	return zdbschema.Text
}

func (pm *PivotManager) pivotKeyType(relation *schema.Relation, parentKeys, relatedKeys []string, isForeign bool, idx int) zdbschema.DataType {
	if relation == nil {
		return zdbschema.Uint
	}

	if isForeign {
		if idx >= 0 && idx < len(parentKeys) {
			if f, ok := pm.schema.getField(parentKeys[idx]); ok {
				if f.Type != "" {
					return zdbschema.DataType(f.Type)
				}
			}
		}
		return zdbschema.Uint
	}

	if idx >= 0 && idx < len(relatedKeys) {
		if relatedSchema, ok := pm.schema.getSchema(relation.Schema); ok && relatedSchema != nil {
			if f, ok := relatedSchema.getField(relatedKeys[idx]); ok {
				if f.Type != "" {
					return zdbschema.DataType(f.Type)
				}
			}
		}
	}

	return zdbschema.Uint
}

func (pm *PivotManager) pivotColumnType(relation *schema.Relation, parentKeys, relatedKeys []string, col string) zdbschema.DataType {
	for i, name := range relation.PivotKeys.Foreign {
		if name == col {
			return pm.pivotKeyType(relation, parentKeys, relatedKeys, true, i)
		}
	}
	for i, name := range relation.PivotKeys.Related {
		if name == col {
			return pm.pivotKeyType(relation, parentKeys, relatedKeys, false, i)
		}
	}
	return pm.pivotFieldType(col)
}

func (pm *PivotManager) uniquePivotFields(fields []*zdbschema.Field) []*zdbschema.Field {
	seen := make(map[string]struct{}, len(fields))
	out := make([]*zdbschema.Field, 0, len(fields))
	for _, f := range fields {
		if f == nil || f.Name == "" {
			continue
		}
		if _, ok := seen[f.Name]; ok {
			continue
		}
		seen[f.Name] = struct{}{}
		out = append(out, f)
	}
	return out
}

func (pm *PivotManager) getDB() (*zdb.DB, error) {
	sqlStorage, ok := pm.schema.Storage.(*SQL)
	if !ok || sqlStorage == nil {
		return nil, errors.New("pivot table only supported for sql storage")
	}
	return sqlStorage.GetDB(), nil
}
