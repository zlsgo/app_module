package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

func TestManyToManyNonNullableEmpty(t *testing.T) {
	tt := zlsgo.NewTest(t)

	users := schema.Schema{
		Name: "users_m2m_nonnull",
		Table: schema.Table{
			Name: "users_m2m_nonnull",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	roles := schema.Schema{
		Name: "roles_m2m_nonnull",
		Table: schema.Table{
			Name: "roles_m2m_nonnull",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	relation := schema.Relation{
		Label:      "Roles",
		Type:       schema.RelationManyToMany,
		Schema:     "roles_m2m_nonnull",
		ForeignKey: []string{IDKey()},
		SchemaKey:  []string{IDKey()},
		PivotKeys:  schema.PivotKeys{Foreign: []string{"user_id"}, Related: []string{"role_id"}},
		Nullable:   false,
	}
	users.Relations = map[string]schema.Relation{
		"roles": relation,
	}

	db, schemas := newTestSchemas(t, users, roles)
	usersRepo := schemas.MustGet("users_m2m_nonnull").Model().Repository()
	rolesRepo := schemas.MustGet("roles_m2m_nonnull").Model().Repository()

	u1, err := usersRepo.Insert(ztype.Map{"name": "u1"})
	tt.NoError(err)
	u2, err := usersRepo.Insert(ztype.Map{"name": "u2"})
	tt.NoError(err)
	r1, err := rolesRepo.Insert(ztype.Map{"name": "r1"})
	tt.NoError(err)

	usersSchema := schemas.MustGet("users_m2m_nonnull")
	pm := NewPivotManager(usersSchema)
	err = pm.SyncPivotSchema(&relation)
	tt.NoError(err)

	pivotTable, err := pm.GetPivotTableName(&relation)
	tt.NoError(err)

	_, err = db.Exec("INSERT INTO "+pivotTable+" (user_id, role_id) VALUES (?, ?)", u1, r1)
	tt.NoError(err)

	rows, err := usersRepo.Query().
		OrderBy(IDKey()).
		WithRelation("roles.name").
		Find()
	tt.NoError(err)
	tt.Equal(2, len(rows))

	rolesAny := rows[1].Get("roles").Value()
	rolesRows, ok := rolesAny.(ztype.Maps)
	tt.Equal(true, ok)
	tt.Equal(0, len(rolesRows))

	tt.Equal(u2, rows[1].Get(IDKey()).Value())
}
