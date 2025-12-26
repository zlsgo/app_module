package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

func TestRelationFilterMany(t *testing.T) {
	tt := zlsgo.NewTest(t)

	parents := schema.Schema{
		Name: "parents_filter",
		Table: schema.Table{
			Name: "parents_filter",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
		Relations: map[string]schema.Relation{
			"children": {
				Label:      "Children",
				Type:       schema.RelationMany,
				Schema:     "children_filter",
				ForeignKey: []string{IDKey()},
				SchemaKey:  []string{"parent_id"},
				Fields:     []string{"value"},
				Filter:     ztype.Map{"status": 1},
				Nullable:   false,
			},
		},
	}

	children := schema.Schema{
		Name: "children_filter",
		Table: schema.Table{
			Name: "children_filter",
		},
		Fields: map[string]schema.Field{
			"parent_id": {Type: "int"},
			"status":    {Type: "int"},
			"value":     {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, parents, children)
	parentsRepo := schemas.MustGet("parents_filter").Model().Repository()
	childrenRepo := schemas.MustGet("children_filter").Model().Repository()

	pid, err := parentsRepo.Insert(ztype.Map{"name": "p1"})
	tt.NoError(err)
	_, err = childrenRepo.Insert(ztype.Map{"parent_id": pid, "status": 1, "value": "ok"})
	tt.NoError(err)
	_, err = childrenRepo.Insert(ztype.Map{"parent_id": pid, "status": 0, "value": "skip"})
	tt.NoError(err)

	rows, err := parentsRepo.Query().
		Select(IDKey()).
		WithRelation("children.value").
		Find()
	tt.NoError(err)
	tt.Equal(1, len(rows))

	childrenAny := rows[0].Get("children").Value()
	childRows, ok := childrenAny.(ztype.Maps)
	tt.Equal(true, ok)
	tt.Equal(1, len(childRows))
	tt.Equal("ok", childRows[0].Get("value").String())
}
