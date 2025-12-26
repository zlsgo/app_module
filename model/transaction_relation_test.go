package model

import (
	"errors"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

func TestTxRelationUsesSameStorage(t *testing.T) {
	tt := zlsgo.NewTest(t)

	parents := schema.Schema{
		Name: "parents_tx",
		Table: schema.Table{
			Name: "parents_tx",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
		Relations: map[string]schema.Relation{
			"children": {
				Label:      "Children",
				Type:       schema.RelationMany,
				Schema:     "children_tx",
				ForeignKey: []string{IDKey()},
				SchemaKey:  []string{"parent_id"},
				Fields:     []string{"value"},
				Nullable:   false,
			},
		},
	}

	children := schema.Schema{
		Name: "children_tx",
		Table: schema.Table{
			Name: "children_tx",
		},
		Fields: map[string]schema.Field{
			"parent_id": {Type: "int"},
			"value":     {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, parents, children)
	parentsRepo := schemas.MustGet("parents_tx").Model().Repository()
	childSchema := schemas.MustGet("children_tx")

	err := parentsRepo.Tx(func(txRepo *Repository[ztype.Map, QueryFilter, ztype.Map, ztype.Map]) error {
		pid, err := txRepo.Insert(ztype.Map{"name": "p1"})
		if err != nil {
			return err
		}

		_, err = txRepo.Schema().Storage.Insert(
			childSchema.GetTableName(),
			ztype.Map{"parent_id": pid, "value": "v1"},
		)
		if err != nil {
			return err
		}

		rows, err := txRepo.Query().
			Select(IDKey()).
			WithRelation("children.value").
			Find()
		if err != nil {
			return err
		}
		if len(rows) != 1 {
			return errors.New("unexpected parent rows")
		}

		childrenAny := rows[0].Get("children").Value()
		childRows, ok := childrenAny.(ztype.Maps)
		if !ok || len(childRows) != 1 {
			return errors.New("relation not loaded in tx")
		}
		return nil
	})
	tt.NoError(err)
}
