package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zsync"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/define"
	"github.com/zlsgo/zdb/schema"
)

func TestDefine(t *testing.T) {
	tt := zlsgo.NewTest(t)
	m := define.Define{
		Name: "test",
		Fields: map[string]define.Field{
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
		Options: define.ModelOptions{DisabledMigrator: true},
	}
	s := NewModels(nil, nil)

	var wg zsync.WaitGroup
	for i := 0; i < 2; i++ {
		ii := i
		wg.Go(func() {
			mod, _ := s.Reg("test"+ztype.ToString(ii), m, false)
			tt.Log(mod.model.Fields)
		})

	}
	_ = wg.Wait()
}
