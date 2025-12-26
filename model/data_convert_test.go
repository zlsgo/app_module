package model

import (
	"reflect"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

type convertUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestDataToMap(t *testing.T) {
	tt := zlsgo.NewTest(t)

	m, err := dataToMap(convertUser{Name: "n1", Age: 18})
	tt.NoError(err)
	tt.Equal("n1", m["name"])
	tt.Equal(18, m["age"])

	m, err = dataToMap(&convertUser{Name: "n2", Age: 20})
	tt.NoError(err)
	tt.Equal("n2", m["name"])

	m, err = dataToMap(ztype.Map{"name": "n3"})
	tt.NoError(err)
	tt.Equal("n3", m["name"])

	_, err = dataToMap(nil)
	tt.Equal(true, err == ErrInvalidData)

	_, err = dataToMap(123)
	tt.Equal(true, err == ErrInvalidData)
}

func TestDataToMaps(t *testing.T) {
	tt := zlsgo.NewTest(t)

	maps, err := dataToMaps([]convertUser{{Name: "a"}, {Name: "b"}})
	tt.NoError(err)
	tt.Equal(2, len(maps))
	tt.Equal("a", maps[0]["name"])

	maps, err = dataToMaps([]map[string]interface{}{{"name": "c"}})
	tt.NoError(err)
	tt.Equal(1, len(maps))
	tt.Equal("c", maps[0]["name"])

	maps, err = dataToMaps(nil)
	tt.NoError(err)
	tt.Equal(0, len(maps))

	_, err = dataToMaps(convertUser{Name: "x"})
	tt.Equal(true, err == ErrInvalidData)
}

func TestGetFilter(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "filter_convert")

	f := getFilter(m, ID(123))
	tt.Equal(123, f[idKey])

	f = getFilter(m, Q(TestUserFilter{Status: 1}))
	tt.Equal(int8(1), f["status"])
	_, hasName := f["name"]
	tt.Equal(false, hasName)

	f = getFilter(m, Filter{"status": 1, "unknown": 2})
	tt.Equal(1, f["status"])
	_, ok := f["unknown"]
	tt.Equal(false, ok)
}

func TestGetFilterSoftDeletes(t *testing.T) {
	tt := zlsgo.NewTest(t)

	b := true
	soft := schema.Schema{
		Name: "soft_users",
		Table: schema.Table{
			Name: "soft_users",
		},
		Options: schema.Options{
			SoftDeletes: &b,
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Label: "Name"},
		},
	}

	softTime := true

	soft.Options.SoftDeleteIsTime = &softTime
	_, schemas := newTestSchemas(t, soft)
	m := schemas.MustGet("soft_users")
	f := getFilter(m, Filter{})
	_, ok := f[DeletedAtKey]
	tt.Equal(true, ok)
}

func TestGetFilterSoftDeletesOverride(t *testing.T) {
	tt := zlsgo.NewTest(t)

	b := true
	soft := schema.Schema{
		Name: "soft_users_override",
		Table: schema.Table{
			Name: "soft_users_override",
		},
		Options: schema.Options{
			SoftDeletes: &b,
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Label: "Name"},
		},
	}

	softTime := true
	soft.Options.SoftDeleteIsTime = &softTime
	_, schemas := newTestSchemas(t, soft)
	m := schemas.MustGet("soft_users_override")

	f := getFilter(m, Filter{DeletedAtKey + " IS NOT NULL": true})
	_, hasDefault := f[DeletedAtKey]
	tt.Equal(false, hasDefault)
	_, hasExplicit := f[DeletedAtKey+" IS NOT NULL"]
	tt.Equal(true, hasExplicit)
}

func TestGetFilterSoftDeletesQualifiedDoesNotOverride(t *testing.T) {
	tt := zlsgo.NewTest(t)

	b := true
	soft := schema.Schema{
		Name: "soft_users_qualified",
		Table: schema.Table{
			Name: "soft_users_qualified",
		},
		Options: schema.Options{
			SoftDeletes: &b,
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Label: "Name"},
		},
	}

	softTime := true
	soft.Options.SoftDeleteIsTime = &softTime
	_, schemas := newTestSchemas(t, soft)
	m := schemas.MustGet("soft_users_qualified")

	f := getFilter(m, Filter{"other.deleted_at IS NOT NULL": true})
	_, hasDefault := f[DeletedAtKey]
	tt.Equal(true, hasDefault)
	_, hasQualified := f["other.deleted_at IS NOT NULL"]
	tt.Equal(true, hasQualified)
}

func TestFieldsFromType(t *testing.T) {
	tt := zlsgo.NewTest(t)

	fields := schema.FieldsFromType(reflect.TypeOf(convertUser{}))
	tt.Equal(2, len(fields))
	tt.Equal(schema.String, fields["name"].Type)
}
