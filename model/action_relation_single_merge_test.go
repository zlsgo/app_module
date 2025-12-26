package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

func TestRelationSingleMerge(t *testing.T) {
	tt := zlsgo.NewTest(t)

	profiles := schema.Schema{
		Name: "profiles_merge",
		Table: schema.Table{
			Name: "profiles_merge",
		},
		Fields: map[string]schema.Field{
			"nickname": {Type: "string", Size: 80},
			"email":    {Type: "string", Size: 120},
		},
	}

	parents := schema.Schema{
		Name: "parents_merge",
		Table: schema.Table{
			Name: "parents_merge",
		},
		Fields: map[string]schema.Field{
			"name":       {Type: "string", Size: 80},
			"profile_id": {Type: "int", Nullable: true},
		},
		Relations: map[string]schema.Relation{
			"profile": {
				Label:      "Profile",
				Type:       schema.RelationSingleMerge,
				Schema:     "profiles_merge",
				ForeignKey: []string{"profile_id"},
				SchemaKey:  []string{IDKey()},
				Fields:     []string{"nickname"},
				Nullable:   true,
			},
		},
	}

	_, schemas := newTestSchemas(t, profiles, parents)
	profilesRepo := schemas.MustGet("profiles_merge").Model().Repository()
	parentsRepo := schemas.MustGet("parents_merge").Model().Repository()

	profileID, err := profilesRepo.Insert(ztype.Map{"nickname": "n1", "email": "e1"})
	tt.NoError(err)

	_, err = parentsRepo.Insert(ztype.Map{"name": "p1", "profile_id": profileID})
	tt.NoError(err)
	_, err = parentsRepo.Insert(ztype.Map{"name": "p2", "profile_id": nil})
	tt.NoError(err)

	rows, err := parentsRepo.Query().
		Select(IDKey(), "name").
		OrderBy(IDKey()).
		WithRelation("profile.nickname").
		Find()
	tt.NoError(err)
	tt.Equal(2, len(rows))

	first := rows[0]
	tt.Equal("p1", first.Get("name").String())
	tt.Equal("n1", first.Get("nickname").String())
	_, hasProfile := first["profile"]
	tt.Equal(false, hasProfile)
	_, hasProfileID := first["profile_id"]
	tt.Equal(false, hasProfileID)

	second := rows[1]
	tt.Equal("p2", second.Get("name").String())
	_, hasNickname := second["nickname"]
	tt.Equal(false, hasNickname)
}
