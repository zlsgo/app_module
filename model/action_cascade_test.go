package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

func TestCascadeRestrictMany(t *testing.T) {
	tt := zlsgo.NewTest(t)

	parents := schema.Schema{
		Name: "parents_restrict",
		Table: schema.Table{
			Name: "parents_restrict",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
		Relations: map[string]schema.Relation{
			"children": {
				Label:       "Children",
				Type:        schema.RelationMany,
				Schema:      "children_restrict",
				ForeignKey:  []string{IDKey()},
				SchemaKey:   []string{"parent_id"},
				Fields:      []string{"value"},
				CascadeType: schema.CascadeTypeRestrict,
			},
		},
	}

	children := schema.Schema{
		Name: "children_restrict",
		Table: schema.Table{
			Name: "children_restrict",
		},
		Fields: map[string]schema.Field{
			"parent_id": {Type: "int"},
			"value":     {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, parents, children)
	parentsRepo := schemas.MustGet("parents_restrict").Model().Repository()
	childrenRepo := schemas.MustGet("children_restrict").Model().Repository()

	pid, err := parentsRepo.Insert(ztype.Map{"name": "p1"})
	tt.NoError(err)
	_, err = childrenRepo.Insert(ztype.Map{"parent_id": pid, "value": "v1"})
	tt.NoError(err)

	_, err = parentsRepo.DeleteByID(pid)
	tt.Equal(true, err != nil)

	rows, err := parentsRepo.Find(Q(ztype.Map{}))
	tt.NoError(err)
	tt.Equal(1, len(rows))
}

func TestCascadeSetNullMany(t *testing.T) {
	tt := zlsgo.NewTest(t)

	parents := schema.Schema{
		Name: "parents_setnull",
		Table: schema.Table{
			Name: "parents_setnull",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
		Relations: map[string]schema.Relation{
			"children": {
				Label:       "Children",
				Type:        schema.RelationMany,
				Schema:      "children_setnull",
				ForeignKey:  []string{IDKey()},
				SchemaKey:   []string{"parent_id"},
				Fields:      []string{"value"},
				CascadeType: schema.CascadeTypeSetNull,
			},
		},
	}

	children := schema.Schema{
		Name: "children_setnull",
		Table: schema.Table{
			Name: "children_setnull",
		},
		Fields: map[string]schema.Field{
			"parent_id": {Type: "int", Nullable: true},
			"value":     {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, parents, children)
	parentsRepo := schemas.MustGet("parents_setnull").Model().Repository()
	childrenRepo := schemas.MustGet("children_setnull").Model().Repository()

	pid, err := parentsRepo.Insert(ztype.Map{"name": "p1"})
	tt.NoError(err)
	_, err = childrenRepo.Insert(ztype.Map{"parent_id": pid, "value": "v1"})
	tt.NoError(err)

	_, err = parentsRepo.DeleteByID(pid)
	tt.NoError(err)

	rows, err := childrenRepo.Find(Q(ztype.Map{}))
	tt.NoError(err)
	tt.Equal(1, len(rows))
	tt.Equal(nil, rows[0].Get("parent_id").Value())
}

func TestCascadeDeleteMany(t *testing.T) {
	tt := zlsgo.NewTest(t)

	parents := schema.Schema{
		Name: "parents_cascade",
		Table: schema.Table{
			Name: "parents_cascade",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
		Relations: map[string]schema.Relation{
			"children": {
				Label:       "Children",
				Type:        schema.RelationMany,
				Schema:      "children_cascade",
				ForeignKey:  []string{IDKey()},
				SchemaKey:   []string{"parent_id"},
				Fields:      []string{"value"},
				CascadeType: schema.CascadeTypeCascade,
			},
		},
	}

	children := schema.Schema{
		Name: "children_cascade",
		Table: schema.Table{
			Name: "children_cascade",
		},
		Fields: map[string]schema.Field{
			"parent_id": {Type: "int"},
			"value":     {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, parents, children)
	parentsRepo := schemas.MustGet("parents_cascade").Model().Repository()
	childrenRepo := schemas.MustGet("children_cascade").Model().Repository()

	pid, err := parentsRepo.Insert(ztype.Map{"name": "p1"})
	tt.NoError(err)
	_, err = childrenRepo.Insert(ztype.Map{"parent_id": pid, "value": "v1"})
	tt.NoError(err)

	_, err = parentsRepo.DeleteByID(pid)
	tt.NoError(err)

	rows, err := childrenRepo.Find(Q(ztype.Map{}))
	tt.NoError(err)
	tt.Equal(0, len(rows))
}

func TestCascadeManyToManyPivotDelete(t *testing.T) {
	tt := zlsgo.NewTest(t)

	users := schema.Schema{
		Name: "users_pivot",
		Table: schema.Table{
			Name: "users_pivot",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	roles := schema.Schema{
		Name: "roles_pivot",
		Table: schema.Table{
			Name: "roles_pivot",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	relation := schema.Relation{
		Label:       "Roles",
		Type:        schema.RelationManyToMany,
		Schema:      "roles_pivot",
		ForeignKey:  []string{IDKey()},
		SchemaKey:   []string{IDKey()},
		PivotKeys:   schema.PivotKeys{Foreign: []string{"user_id"}, Related: []string{"role_id"}},
		CascadeType: schema.CascadeTypeCascade,
		Nullable:    true,
	}
	users.Relations = map[string]schema.Relation{
		"roles": relation,
	}

	db, schemas := newTestSchemas(t, users, roles)
	usersRepo := schemas.MustGet("users_pivot").Model().Repository()
	rolesRepo := schemas.MustGet("roles_pivot").Model().Repository()

	uid, err := usersRepo.Insert(ztype.Map{"name": "u1"})
	tt.NoError(err)
	r1, err := rolesRepo.Insert(ztype.Map{"name": "r1"})
	tt.NoError(err)
	r2, err := rolesRepo.Insert(ztype.Map{"name": "r2"})
	tt.NoError(err)

	usersSchema := schemas.MustGet("users_pivot")
	pm := NewPivotManager(usersSchema)
	err = pm.SyncPivotSchema(&relation)
	tt.NoError(err)

	pivotTable, err := pm.GetPivotTableName(&relation)
	tt.NoError(err)

	_, err = db.Exec("INSERT INTO "+pivotTable+" (user_id, role_id) VALUES (?, ?)", uid, r1)
	tt.NoError(err)
	_, err = db.Exec("INSERT INTO "+pivotTable+" (user_id, role_id) VALUES (?, ?)", uid, r2)
	tt.NoError(err)

	_, err = usersRepo.DeleteByID(uid)
	tt.NoError(err)

	pivotRows, err := db.QueryToMaps("SELECT * FROM " + pivotTable)
	tt.NoError(err)
	tt.Equal(0, len(pivotRows))

	roleRows, err := rolesRepo.Find(Q(ztype.Map{}))
	tt.NoError(err)
	tt.Equal(2, len(roleRows))
}

func TestCascadeSoftDelete(t *testing.T) {
	tt := zlsgo.NewTest(t)

	b := true
	softTime := false
	parents := schema.Schema{
		Name: "parents_soft",
		Table: schema.Table{
			Name: "parents_soft",
		},
		Options: schema.Options{
			SoftDeletes:      &b,
			SoftDeleteIsTime: &softTime,
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
		Relations: map[string]schema.Relation{
			"children": {
				Label:       "Children",
				Type:        schema.RelationMany,
				Schema:      "children_soft",
				ForeignKey:  []string{IDKey()},
				SchemaKey:   []string{"parent_id"},
				Fields:      []string{"value"},
				CascadeType: schema.CascadeTypeCascade,
			},
		},
	}

	children := schema.Schema{
		Name: "children_soft",
		Table: schema.Table{
			Name: "children_soft",
		},
		Options: schema.Options{
			SoftDeletes:      &b,
			SoftDeleteIsTime: &softTime,
		},
		Fields: map[string]schema.Field{
			"parent_id": {Type: "int"},
			"value":     {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, parents, children)
	parentsRepo := schemas.MustGet("parents_soft").Model().Repository()
	childrenSchema := schemas.MustGet("children_soft")
	childrenRepo := childrenSchema.Model().Repository()

	pid, err := parentsRepo.Insert(ztype.Map{"name": "p1"})
	tt.NoError(err)
	_, err = childrenRepo.Insert(ztype.Map{"parent_id": pid, "value": "v1"})
	tt.NoError(err)

	_, err = parentsRepo.DeleteByID(pid)
	tt.NoError(err)

	rows, err := childrenSchema.Storage.Find(
		childrenSchema.GetTableName(),
		ztype.Map{"parent_id": pid},
		func(co *CondOptions) {
			co.Fields = append(co.Fields[:0], IDKey(), "parent_id", DeletedAtKey)
		},
	)
	tt.NoError(err)
	tt.Equal(1, len(rows))
	tt.Equal(true, rows[0].Get(DeletedAtKey).Int64() > 0)
}

func TestCascadeManyToManyRestrict(t *testing.T) {
	tt := zlsgo.NewTest(t)

	users := schema.Schema{
		Name: "users_pivot_restrict",
		Table: schema.Table{
			Name: "users_pivot_restrict",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	roles := schema.Schema{
		Name: "roles_pivot_restrict",
		Table: schema.Table{
			Name: "roles_pivot_restrict",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	relation := schema.Relation{
		Label:       "Roles",
		Type:        schema.RelationManyToMany,
		Schema:      "roles_pivot_restrict",
		ForeignKey:  []string{IDKey()},
		SchemaKey:   []string{IDKey()},
		PivotKeys:   schema.PivotKeys{Foreign: []string{"user_id"}, Related: []string{"role_id"}},
		CascadeType: schema.CascadeTypeRestrict,
		Nullable:    true,
	}
	users.Relations = map[string]schema.Relation{
		"roles": relation,
	}

	db, schemas := newTestSchemas(t, users, roles)
	usersRepo := schemas.MustGet("users_pivot_restrict").Model().Repository()
	rolesRepo := schemas.MustGet("roles_pivot_restrict").Model().Repository()

	uid, err := usersRepo.Insert(ztype.Map{"name": "u1"})
	tt.NoError(err)
	r1, err := rolesRepo.Insert(ztype.Map{"name": "r1"})
	tt.NoError(err)

	usersSchema := schemas.MustGet("users_pivot_restrict")
	pm := NewPivotManager(usersSchema)
	err = pm.SyncPivotSchema(&relation)
	tt.NoError(err)

	pivotTable, err := pm.GetPivotTableName(&relation)
	tt.NoError(err)

	_, err = db.Exec("INSERT INTO "+pivotTable+" (user_id, role_id) VALUES (?, ?)", uid, r1)
	tt.NoError(err)

	_, err = usersRepo.DeleteByID(uid)
	tt.Equal(true, err != nil)
}
