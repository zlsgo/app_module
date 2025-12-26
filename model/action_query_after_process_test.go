package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

func TestFindAfterProcessChain(t *testing.T) {
	tt := zlsgo.NewTest(t)

	events := schema.Schema{
		Name: "events",
		Table: schema.Table{
			Name: "events",
		},
		Fields: map[string]schema.Field{
			"event_time": {
				Type:  "string",
				After: []string{"date|Y-m-d H:i:s", "date|Y-m-d"},
			},
		},
	}

	_, schemas := newTestSchemas(t, events)
	repo := schemas.MustGet("events").Model().Repository()

	_, err := repo.Insert(ztype.Map{"event_time": "2025-12-22 10:11:12"})
	tt.NoError(err)

	row, err := repo.FindOne(Filter{})
	tt.NoError(err)
	tt.Equal("2025-12-22", row.Get("event_time").String())
}
