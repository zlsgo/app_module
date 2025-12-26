package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

func TestSchemasModelsRefresh(t *testing.T) {
	tt := zlsgo.NewTest(t)

	db, err := zdb.New(&sqlite3.Config{
		File:       ":memory:",
		Memory:     true,
		Parameters: "_pragma=busy_timeout(3000)",
	})
	tt.NoError(err)

	ss := NewSchemas(nil, NewSQL(db, ""), SchemaOptions{})

	first := schema.Schema{
		Name: "reg_a",
		Table: schema.Table{
			Name: "reg_a",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	second := schema.Schema{
		Name: "reg_b",
		Table: schema.Table{
			Name: "reg_b",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	_, err = ss.Reg("reg_a", first, false)
	tt.NoError(err)

	models := ss.Models()
	_, ok := models.Get("reg_a")
	tt.Equal(true, ok)

	_, err = ss.Reg("reg_b", second, false)
	tt.NoError(err)

	_, ok = models.Get("reg_b")
	tt.Equal(true, ok)
}
