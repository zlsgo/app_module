package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zsync"
	"github.com/sohaha/zlsgo/ztype"
	mSchema "github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb/schema"
)

func TestDefine(t *testing.T) {
	tt := zlsgo.NewTest(t)

	b := true
	m := mSchema.Schema{
		Name: "test",
		Fields: map[string]mSchema.Field{
			"test11111111": {
				Type: schema.String,
				Size: 255,
			},
			// "测试": {
			// 	Type:  schema.String,
			// 	Size:  255,
			// 	Label: "新的",
			// },
		},
		Options: mSchema.Options{DisabledMigrator: &b},
	}
	s := NewSchemas(nil, nil, SchemaOptions{})

	var wg zsync.WaitGroup
	for i := 0; i < 2; i++ {
		ii := i
		wg.Go(func() {
			mod, _ := s.Reg("test"+ztype.ToString(ii), m, false)
			tt.Log(mod.define.Fields)
		})

	}
	_ = wg.Wait()
}
