package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

func TestNestedRelationLoad(t *testing.T) {
	tt := zlsgo.NewTest(t)

	parents := schema.Schema{
		Name: "parents_nested",
		Table: schema.Table{
			Name: "parents_nested",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
		Relations: map[string]schema.Relation{
			"children": {
				Label:      "Children",
				Type:       schema.RelationMany,
				Schema:     "children_nested",
				ForeignKey: []string{IDKey()},
				SchemaKey:  []string{"parent_id"},
				Fields:     []string{"name"},
				Nullable:   false,
			},
		},
	}

	children := schema.Schema{
		Name: "children_nested",
		Table: schema.Table{
			Name: "children_nested",
		},
		Fields: map[string]schema.Field{
			"parent_id": {Type: "int"},
			"name":      {Type: "string", Size: 80},
		},
		Relations: map[string]schema.Relation{
			"grandchildren": {
				Label:      "Grandchildren",
				Type:       schema.RelationMany,
				Schema:     "grandchildren_nested",
				ForeignKey: []string{IDKey()},
				SchemaKey:  []string{"child_id"},
				Fields:     []string{"value"},
				Nullable:   false,
			},
		},
	}

	grandchildren := schema.Schema{
		Name: "grandchildren_nested",
		Table: schema.Table{
			Name: "grandchildren_nested",
		},
		Fields: map[string]schema.Field{
			"child_id": {Type: "int"},
			"value":    {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, parents, children, grandchildren)
	parentsRepo := schemas.MustGet("parents_nested").Model().Repository()
	childrenRepo := schemas.MustGet("children_nested").Model().Repository()
	grandchildrenRepo := schemas.MustGet("grandchildren_nested").Model().Repository()

	p1, err := parentsRepo.Insert(ztype.Map{"name": "p1"})
	tt.NoError(err)
	_, err = parentsRepo.Insert(ztype.Map{"name": "p2"})
	tt.NoError(err)

	c1, err := childrenRepo.Insert(ztype.Map{"parent_id": p1, "name": "c1"})
	tt.NoError(err)
	c2, err := childrenRepo.Insert(ztype.Map{"parent_id": p1, "name": "c2"})
	tt.NoError(err)

	_, err = grandchildrenRepo.Insert(ztype.Map{"child_id": c1, "value": "g1"})
	tt.NoError(err)
	_, err = grandchildrenRepo.Insert(ztype.Map{"child_id": c2, "value": "g2"})
	tt.NoError(err)

	rows, err := parentsRepo.Query().
		Select(IDKey()).
		OrderBy(IDKey()).
		WithRelation("children.grandchildren.value").
		Find()
	tt.NoError(err)
	tt.Equal(2, len(rows))

	childrenAny := rows[0].Get("children").Value()
	childRows, ok := childrenAny.(ztype.Maps)
	tt.Equal(true, ok)
	tt.Equal(2, len(childRows))

	values := make(map[string]string)
	for i := range childRows {
		child := childRows[i]
		name := child.Get("name").String()
		_, hasID := child[IDKey()]
		tt.Equal(false, hasID)
		_, hasParent := child["parent_id"]
		tt.Equal(false, hasParent)
		gAny := child.Get("grandchildren").Value()
		gRows := gAny.(ztype.Maps)
		tt.Equal(1, len(gRows))
		values[name] = gRows[0].Get("value").String()
	}

	tt.Equal("g1", values["c1"])
	tt.Equal("g2", values["c2"])

	emptyAny := rows[1].Get("children").Value()
	emptyRows, ok := emptyAny.(ztype.Maps)
	tt.Equal(true, ok)
	tt.Equal(0, len(emptyRows))
}
