package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

func newTestSchemas(t *testing.T, schemas ...schema.Schema) (*zdb.DB, *Schemas) {
	return newTestSchemasWithOptions(t, SchemaOptions{}, schemas...)
}

func newTestSchemasWithOptions(t *testing.T, opts SchemaOptions, schemas ...schema.Schema) (*zdb.DB, *Schemas) {
	db, err := zdb.New(&sqlite3.Config{
		File:       ":memory:",
		Memory:     true,
		Parameters: "_pragma=busy_timeout(3000)",
	})
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}

	ss := NewSchemas(nil, NewSQL(db, ""), opts)
	for _, s := range schemas {
		m, err := ss.Reg(s.Name, s, false)
		if err != nil {
			t.Fatalf("failed to register schema %s: %v", s.Name, err)
		}
		err = m.Migration().Auto(DealOldColumnNone)
		if err != nil {
			t.Fatalf("failed to migrate schema %s: %v", s.Name, err)
		}
	}

	return db, ss
}

func TestRelationNullableSingle(t *testing.T) {
	tt := zlsgo.NewTest(t)

	profiles := schema.Schema{
		Name: "profiles",
		Table: schema.Table{
			Name: "profiles",
		},
		Fields: map[string]schema.Field{
			"nickname": {Type: "string", Label: "Nickname", Size: 80},
		},
	}

	parents := schema.Schema{
		Name: "parents",
		Table: schema.Table{
			Name: "parents",
		},
		Fields: map[string]schema.Field{
			"name":       {Type: "string", Size: 80},
			"profile_id": {Type: "int", Nullable: true},
		},
		Relations: map[string]schema.Relation{
			"profile": {
				Label:      "Profile",
				Type:       schema.RelationSingle,
				Schema:     "profiles",
				ForeignKey: []string{"profile_id"},
				SchemaKey:  []string{IDKey()},
				Fields:     []string{"nickname"},
				Nullable:   true,
			},
		},
	}

	_, schemas := newTestSchemas(t, profiles, parents)
	profilesRepo := schemas.MustGet("profiles").Model().Repository()
	parentsRepo := schemas.MustGet("parents").Model().Repository()

	profileID, err := profilesRepo.Insert(ztype.Map{"nickname": "p1"})
	tt.NoError(err)

	_, err = parentsRepo.Insert(ztype.Map{"name": "p1", "profile_id": profileID})
	tt.NoError(err)
	_, err = parentsRepo.Insert(ztype.Map{"name": "p2", "profile_id": nil})
	tt.NoError(err)

	rows, err := parentsRepo.Query().
		Select(IDKey()).
		OrderBy(IDKey()).
		WithRelation("profile.nickname").
		Find()
	tt.NoError(err)
	tt.Equal(2, len(rows))

	profileAny := rows[0].Get("profile").Value()
	profile, ok := profileAny.(ztype.Map)
	tt.Equal(true, ok)
	tt.Equal("p1", profile.Get("nickname").String())

	missing := rows[1].Get("profile").Value()
	tt.Equal(nil, missing)
}

func TestRelationNonNullableSingle(t *testing.T) {
	tt := zlsgo.NewTest(t)

	profiles := schema.Schema{
		Name: "profiles_nonnull",
		Table: schema.Table{
			Name: "profiles_nonnull",
		},
		Fields: map[string]schema.Field{
			"nickname": {Type: "string", Label: "Nickname", Size: 80},
		},
	}

	parents := schema.Schema{
		Name: "parents_nonnull",
		Table: schema.Table{
			Name: "parents_nonnull",
		},
		Fields: map[string]schema.Field{
			"name":       {Type: "string", Size: 80},
			"profile_id": {Type: "int", Nullable: true},
		},
		Relations: map[string]schema.Relation{
			"profile": {
				Label:      "Profile",
				Type:       schema.RelationSingle,
				Schema:     "profiles_nonnull",
				ForeignKey: []string{"profile_id"},
				SchemaKey:  []string{IDKey()},
				Fields:     []string{"nickname"},
				Nullable:   false,
			},
		},
	}

	_, schemas := newTestSchemas(t, profiles, parents)
	parentsRepo := schemas.MustGet("parents_nonnull").Model().Repository()

	_, err := parentsRepo.Insert(ztype.Map{"name": "p2", "profile_id": nil})
	tt.NoError(err)

	rows, err := parentsRepo.Query().
		Select(IDKey()).
		OrderBy(IDKey()).
		WithRelation("profile.nickname").
		Find()
	tt.NoError(err)
	tt.Equal(1, len(rows))

	missing := rows[0].Get("profile").Value()
	_, ok := missing.(ztype.Map)
	tt.Equal(true, ok)
}

func TestRelationNullableMany(t *testing.T) {
	tt := zlsgo.NewTest(t)

	parents := schema.Schema{
		Name: "parents_many",
		Table: schema.Table{
			Name: "parents_many",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
		Relations: map[string]schema.Relation{
			"children": {
				Label:      "Children",
				Type:       schema.RelationMany,
				Schema:     "children_many",
				ForeignKey: []string{IDKey()},
				SchemaKey:  []string{"parent_id"},
				Fields:     []string{"value"},
				Nullable:   true,
			},
		},
	}

	children := schema.Schema{
		Name: "children_many",
		Table: schema.Table{
			Name: "children_many",
		},
		Fields: map[string]schema.Field{
			"parent_id": {Type: "int"},
			"value":     {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, parents, children)
	parentsRepo := schemas.MustGet("parents_many").Model().Repository()

	_, err := parentsRepo.Insert(ztype.Map{"name": "p1"})
	tt.NoError(err)

	rows, err := parentsRepo.Query().
		Select(IDKey()).
		OrderBy(IDKey()).
		WithRelation("children.value").
		Find()
	tt.NoError(err)
	tt.Equal(1, len(rows))

	missing := rows[0].Get("children").Value()
	tt.Equal(nil, missing)
}

func TestRelationNonNullableMany(t *testing.T) {
	tt := zlsgo.NewTest(t)

	parents := schema.Schema{
		Name: "parents_many_nonnull",
		Table: schema.Table{
			Name: "parents_many_nonnull",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
		Relations: map[string]schema.Relation{
			"children": {
				Label:      "Children",
				Type:       schema.RelationMany,
				Schema:     "children_many_nonnull",
				ForeignKey: []string{IDKey()},
				SchemaKey:  []string{"parent_id"},
				Fields:     []string{"value"},
				Nullable:   false,
			},
		},
	}

	children := schema.Schema{
		Name: "children_many_nonnull",
		Table: schema.Table{
			Name: "children_many_nonnull",
		},
		Fields: map[string]schema.Field{
			"parent_id": {Type: "int"},
			"value":     {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, parents, children)
	parentsRepo := schemas.MustGet("parents_many_nonnull").Model().Repository()

	_, err := parentsRepo.Insert(ztype.Map{"name": "p1"})
	tt.NoError(err)

	rows, err := parentsRepo.Query().
		Select(IDKey()).
		OrderBy(IDKey()).
		WithRelation("children.value").
		Find()
	tt.NoError(err)
	tt.Equal(1, len(rows))

	missing := rows[0].Get("children").Value()
	_, ok := missing.(ztype.Maps)
	tt.Equal(true, ok)
}

func TestRelationCompositeKeys(t *testing.T) {
	tt := zlsgo.NewTest(t)

	parents := schema.Schema{
		Name: "parents_comp",
		Table: schema.Table{
			Name: "parents_comp",
		},
		Fields: map[string]schema.Field{
			"tenant_id": {Type: "int"},
			"code":      {Type: "string", Size: 40},
		},
		Relations: map[string]schema.Relation{
			"children": {
				Label:      "Children",
				Type:       schema.RelationMany,
				Schema:     "children_comp",
				ForeignKey: []string{"tenant_id", "code"},
				SchemaKey:  []string{"tenant_id", "code"},
				Fields:     []string{"value"},
				Nullable:   false,
			},
		},
	}

	children := schema.Schema{
		Name: "children_comp",
		Table: schema.Table{
			Name: "children_comp",
		},
		Fields: map[string]schema.Field{
			"tenant_id": {Type: "int"},
			"code":      {Type: "string", Size: 40},
			"value":     {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, parents, children)
	parentsRepo := schemas.MustGet("parents_comp").Model().Repository()
	childrenRepo := schemas.MustGet("children_comp").Model().Repository()

	_, err := parentsRepo.Insert(ztype.Map{"tenant_id": 1, "code": "a"})
	tt.NoError(err)
	_, err = parentsRepo.Insert(ztype.Map{"tenant_id": 1, "code": "b"})
	tt.NoError(err)

	_, err = childrenRepo.Insert(ztype.Map{"tenant_id": 1, "code": "a", "value": "va"})
	tt.NoError(err)
	_, err = childrenRepo.Insert(ztype.Map{"tenant_id": 1, "code": "b", "value": "vb"})
	tt.NoError(err)
	_, err = childrenRepo.Insert(ztype.Map{"tenant_id": 1, "code": "c", "value": "vc"})
	tt.NoError(err)

	rows, err := parentsRepo.Query().
		Select(IDKey()).
		OrderBy(IDKey()).
		WithRelation("children.value").
		Find()
	tt.NoError(err)
	tt.Equal(2, len(rows))

	first := rows[0].Get("children").Value().(ztype.Maps)
	second := rows[1].Get("children").Value().(ztype.Maps)
	tt.Equal(1, len(first))
	tt.Equal(1, len(second))
	tt.Equal("va", first[0].Get("value").String())
	tt.Equal("vb", second[0].Get("value").String())
}
