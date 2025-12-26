package model

import (
	"strings"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

func newTestDBWithManyToMany(t *testing.T, tablePrefix string) (*zdb.DB, *Schemas, schema.Relation) {
	db, err := zdb.New(&sqlite3.Config{
		File:       ":memory:",
		Memory:     true,
		Parameters: "_pragma=busy_timeout(3000)",
	})
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}

	schemas := NewSchemas(nil, NewSQL(db, ""), SchemaOptions{})

	users := schema.Schema{
		Name: "users",
		Table: schema.Table{
			Name:    tablePrefix + "_users",
			Comment: "Test Users",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Label: "Name", Size: 100},
		},
	}

	roles := schema.Schema{
		Name: "roles",
		Table: schema.Table{
			Name:    tablePrefix + "_roles",
			Comment: "Test Roles",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Label: "Name", Size: 100},
		},
	}

	relation := schema.Relation{
		Label:      "Roles",
		Type:       schema.RelationManyToMany,
		Schema:     "roles",
		ForeignKey: []string{IDKey()},
		SchemaKey:  []string{IDKey()},
		PivotKeys: schema.PivotKeys{
			Foreign: []string{"user_id"},
			Related: []string{"role_id"},
		},
		PivotFields: []string{"assigned_at"},
		Nullable:    true,
	}
	users.Relations = map[string]schema.Relation{
		"roles": relation,
	}

	for _, s := range []schema.Schema{users, roles} {
		m, err := schemas.Reg(s.Name, s, false)
		if err != nil {
			t.Fatalf("failed to register schema %s: %v", s.Name, err)
		}
		err = m.Migration().Auto(DealOldColumnNone)
		if err != nil {
			t.Fatalf("failed to migrate schema %s: %v", s.Name, err)
		}
	}

	return db, schemas, relation
}

func TestManyToManyRelationLoad(t *testing.T) {
	tt := zlsgo.NewTest(t)
	db, schemas, relation := newTestDBWithManyToMany(t, "m2m")

	usersRepo := schemas.MustGet("users").Model().Repository()
	rolesRepo := schemas.MustGet("roles").Model().Repository()

	u1, err := usersRepo.Insert(ztype.Map{"name": "u1"})
	tt.NoError(err)
	u2, err := usersRepo.Insert(ztype.Map{"name": "u2"})
	tt.NoError(err)

	r1, err := rolesRepo.Insert(ztype.Map{"name": "admin"})
	tt.NoError(err)
	r2, err := rolesRepo.Insert(ztype.Map{"name": "reader"})
	tt.NoError(err)

	usersSchema := schemas.MustGet("users")
	pm := NewPivotManager(usersSchema)
	err = pm.SyncPivotSchema(&relation)
	tt.NoError(err)

	pivotTable, err := pm.GetPivotTableName(&relation)
	tt.NoError(err)

	_, err = db.Exec(
		"INSERT INTO "+pivotTable+" (user_id, role_id, assigned_at) VALUES (?, ?, ?)",
		u1, r1, "2025-01-01",
	)
	tt.NoError(err)
	_, err = db.Exec(
		"INSERT INTO "+pivotTable+" (user_id, role_id, assigned_at) VALUES (?, ?, ?)",
		u1, r2, "2025-01-02",
	)
	tt.NoError(err)

	rows, err := usersRepo.Query().
		OrderBy(IDKey()).
		WithRelation("roles.name").
		Find()
	tt.NoError(err)
	tt.Equal(2, len(rows))
	tt.Equal(u1, rows[0].Get(IDKey()).Value())
	tt.Equal(u2, rows[1].Get(IDKey()).Value())

	rolesAny := rows[0].Get("roles").Value()
	roles, ok := rolesAny.(ztype.Maps)
	tt.Equal(true, ok)
	tt.Equal(2, len(roles))

	_, hasID := roles[0][IDKey()]
	tt.Equal(false, hasID)
	_, hasPivot := roles[0]["pivot_assigned_at"]
	tt.Equal(true, hasPivot)

	noRoles := rows[1].Get("roles").Value()
	tt.Equal(nil, noRoles)
}

func TestPivotKeyTypeInference(t *testing.T) {
	tt := zlsgo.NewTest(t)

	db, err := zdb.New(&sqlite3.Config{
		File:       ":memory:",
		Memory:     true,
		Parameters: "_pragma=busy_timeout(3000)",
	})
	tt.NoError(err)

	schemas := NewSchemas(nil, NewSQL(db, ""), SchemaOptions{})

	users := schema.Schema{
		Name: "users_text",
		Table: schema.Table{
			Name: "users_text",
		},
		Fields: map[string]schema.Field{
			"uuid": {Type: "string", Size: 64},
			"name": {Type: "string", Size: 100},
		},
	}

	roles := schema.Schema{
		Name: "roles_text",
		Table: schema.Table{
			Name: "roles_text",
		},
		Fields: map[string]schema.Field{
			"code":  {Type: "string", Size: 64},
			"label": {Type: "string", Size: 100},
		},
	}

	relation := schema.Relation{
		Label:      "Roles",
		Type:       schema.RelationManyToMany,
		Schema:     "roles_text",
		ForeignKey: []string{"uuid"},
		SchemaKey:  []string{"code"},
		PivotKeys: schema.PivotKeys{
			Foreign: []string{"user_uuid"},
			Related: []string{"role_code"},
		},
	}
	users.Relations = map[string]schema.Relation{
		"roles": relation,
	}

	for _, s := range []schema.Schema{users, roles} {
		m, err := schemas.Reg(s.Name, s, false)
		tt.NoError(err)
		err = m.Migration().Auto(DealOldColumnNone)
		tt.NoError(err)
	}

	usersSchema := schemas.MustGet("users_text")
	pm := NewPivotManager(usersSchema)
	err = pm.SyncPivotSchema(&relation)
	tt.NoError(err)

	pivotTable, err := pm.GetPivotTableName(&relation)
	tt.NoError(err)

	rows, err := db.QueryToMaps("PRAGMA table_info('" + pivotTable + "')")
	tt.NoError(err)
	tt.Equal(true, len(rows) > 0)

	types := map[string]string{}
	for _, row := range rows {
		name := row.Get("name").String()
		typ := strings.ToLower(row.Get("type").String())
		types[name] = typ
	}

	userType := types["user_uuid"]
	roleType := types["role_code"]
	tt.Equal(true, userType != "" && (strings.Contains(userType, "text") || strings.Contains(userType, "char")))
	tt.Equal(true, roleType != "" && (strings.Contains(roleType, "text") || strings.Contains(roleType, "char")))
}
