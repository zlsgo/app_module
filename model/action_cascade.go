package model

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

func cascadeFields(m *Schema) []string {
	if m == nil {
		return nil
	}

	fields := make([]string, 0)
	for _, rel := range m.define.Relations {
		if cascadeType(rel) == "" {
			continue
		}

		switch rel.Type {
		case schema.RelationManyToMany:
			keys := rel.ForeignKey
			if len(keys) == 0 {
				keys = []string{idKey}
			}
			fields = append(fields, keys...)
		default:
			if len(rel.ForeignKey) > 0 {
				fields = append(fields, rel.ForeignKey...)
			}
		}
	}

	return zarray.Unique(fields)
}

func cascadeDelete(m *Schema, rows ztype.Maps) error {
	if m == nil || len(rows) == 0 {
		return nil
	}

	for name, rel := range m.define.Relations {
		cType := cascadeType(rel)
		if cType == "" {
			continue
		}

		var err error
		switch rel.Type {
		case schema.RelationManyToMany:
			err = cascadeDeletePivot(m, rel, cType, rows)
		case schema.RelationSingle, schema.RelationSingleMerge, schema.RelationMany:
			err = cascadeDeleteRelation(m, rel, cType, rows)
		}
		if err != nil {
			return errors.New(name + ": " + err.Error())
		}
	}

	return nil
}

func cascadeType(rel schema.Relation) schema.CascadeType {
	if rel.CascadeType != "" {
		return rel.CascadeType
	}
	if rel.Cascade != "" {
		return schema.CascadeType(strings.ToUpper(rel.Cascade))
	}
	return ""
}

func cascadeDeleteRelation(m *Schema, rel schema.Relation, cType schema.CascadeType, rows ztype.Maps) error {
	childSchema, ok := m.getSchema(rel.Schema)
	if !ok {
		return nil
	}
	if len(rel.ForeignKey) == 0 || len(rel.SchemaKey) == 0 {
		return nil
	}

	parentTuples := collectKeyTuples(rows, rel.ForeignKey)
	filter := buildCompositeFilter(rel.SchemaKey, parentTuples)
	if len(filter) == 0 {
		return nil
	}

	switch cType {
	case schema.CascadeTypeRestrict:
		exists, err := hasRows(childSchema, filter)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("restrict")
		}
	case schema.CascadeTypeSetNull:
		if rel.Type != schema.RelationMany {
			return nil
		}
		data := make(ztype.Map, len(rel.SchemaKey))
		for _, k := range rel.SchemaKey {
			data[k] = nil
		}
		_, err := childSchema.Storage.Update(childSchema.GetTableName(), data, filter)
		return err
	case schema.CascadeTypeCascade:
		return deleteByFilter(childSchema, filter)
	}

	return nil
}

func cascadeDeletePivot(m *Schema, rel schema.Relation, cType schema.CascadeType, rows ztype.Maps) error {
	pm := NewPivotManager(m)
	if err := pm.SyncPivotSchema(&rel); err != nil {
		return err
	}
	pivotTable, err := pm.GetPivotTableName(&rel)
	if err != nil {
		return err
	}

	parentKeys := rel.ForeignKey
	if len(parentKeys) == 0 {
		parentKeys = []string{idKey}
	}
	if len(rel.PivotKeys.Foreign) == 0 {
		return nil
	}

	parentTuples := collectKeyTuples(rows, parentKeys)
	filter := buildCompositeFilter(rel.PivotKeys.Foreign, parentTuples)
	if len(filter) == 0 {
		return nil
	}

	switch cType {
	case schema.CascadeTypeRestrict:
		exists, err := hasRowsTable(m.Storage, pivotTable, filter)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("restrict")
		}
	case schema.CascadeTypeSetNull, schema.CascadeTypeCascade:
		_, err := m.Storage.Delete(pivotTable, filter)
		return err
	}

	return nil
}

func deleteByFilter(m *Schema, filter ztype.Map) error {
	if *m.define.Options.SoftDeletes {
		data := make(ztype.Map, 1)
		now := ztime.Time()
		if *m.define.Options.SoftDeleteIsTime {
			data[DeletedAtKey] = now
		} else {
			data[DeletedAtKey] = now.Unix()
		}
		_, err := m.Storage.Update(m.GetTableName(), data, filter)
		return err
	}
	_, err := m.Storage.Delete(m.GetTableName(), filter)
	return err
}

func hasRows(m *Schema, filter ztype.Map) (bool, error) {
	rows, err := m.Storage.Find(m.GetTableName(), filter, func(co *CondOptions) {
		co.Fields = append(co.Fields[:0], idKey)
		co.Limit = 1
	})
	if err != nil {
		return false, err
	}
	return len(rows) > 0, nil
}

func hasRowsTable(s Storageer, table string, filter ztype.Map) (bool, error) {
	rows, err := s.Find(table, filter, func(co *CondOptions) {
		co.Limit = 1
	})
	if err != nil {
		return false, err
	}
	return len(rows) > 0, nil
}
