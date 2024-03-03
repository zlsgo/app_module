package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zsync"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb/schema"
)

func TestDefine(t *testing.T) {
	tt := zlsgo.NewTest(t)
	m := Define{
		Name: "test",
		Fields: map[string]Field{
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
		Options: Options{DisabledMigrator: true},
	}
	s := New(nil)

	var wg zsync.WaitGroup
	for i := 0; i < 2; i++ {
		ii := i
		wg.Go(func() {
			mod, _ := s.Reg("test"+ztype.ToString(ii), m, false)
			tt.Log(mod.model.Fields)
		})

	}
	wg.Wait()
}
