package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

func newViewsSchema(name string, ext ztype.Map) schema.Schema {
	s := schema.Schema{
		Name:  name,
		Table: schema.Table{Name: name},
		Fields: map[string]schema.Field{
			"name":  {Type: "string", Label: "名称", Size: 80},
			"email": {Type: "string", Label: "邮箱", Size: 120},
			"age":   {Type: "int", Label: "年龄"},
			"status": {
				Type:  "int",
				Label: "状态",
				Options: schema.FieldOption{
					Enum: []schema.FieldEnum{{Value: "1", Label: "启用"}, {Value: "0", Label: "停用"}},
				},
			},
		},
	}
	if ext != nil {
		s.Extend = ext
	}
	return s
}

func TestParseViewLists(t *testing.T) {
	tt := zlsgo.NewTest(t)

	users := newViewsSchema("view_lists", ztype.Map{
		"views": ztype.Map{
			"lists": ztype.Map{
				"title":    "用户列表",
				"fields":   []string{"name", "status"},
				"disabled": false,
			},
		},
	})

	_, schemas := newTestSchemas(t, users)
	m := schemas.MustGet("view_lists")

	lists := m.GetViews().Get("lists")
	fields := lists.Get("fields").Slice().String()
	// 自动补齐主键
	tt.Equal(true, zarray.Contains(fields, idKey))
	// 声明的字段都在列
	tt.Equal(true, zarray.Contains(fields, "name"))
	tt.Equal(true, zarray.Contains(fields, "status"))
	tt.Equal("用户列表", lists.Get("title").String())
	// 未声明字段不进入（除主键外）
	tt.Equal(false, zarray.Contains(fields, "email"))

	// 运行时 GetViewFields 与解析结果一致
	tt.Equal(true, zarray.Contains(m.GetViewFields("lists"), idKey))
	tt.Equal(true, zarray.Contains(m.GetViewFields("lists"), "name"))
}

func TestParseViewInfoAndGetViewFields(t *testing.T) {
	tt := zlsgo.NewTest(t)

	users := newViewsSchema("view_info", ztype.Map{
		"views": ztype.Map{
			"info": ztype.Map{
				"fields": []string{"name", "email"},
				"layouts": ztype.Map{
					"name": ztype.Map{"x": 1},
				},
			},
		},
	})

	_, schemas := newTestSchemas(t, users)
	m := schemas.MustGet("view_info")

	info := m.GetViews().Get("info")
	fields := info.Get("fields").Slice().String()
	tt.Equal(true, zarray.Contains(fields, idKey))
	tt.Equal(true, zarray.Contains(fields, "name"))
	tt.Equal(true, zarray.Contains(fields, "email"))
	// layouts 保留
	tt.Equal(float64(1), info.Get("layouts").Get("name").Get("x").Float64())

	// 未匹配到明确视图时回退到全字段
	def := m.GetViewFields("not_exist")
	tt.Equal(true, zarray.Contains(def, "name"))
	tt.Equal(true, zarray.Contains(def, "age"))

	// info 视图 (未声明 disabled) 未把 email 排除在字段里
	tt.Equal(true, zarray.Contains(info.Get("fields").Slice().String(), "email"))
}

func TestParseViewsDisabled(t *testing.T) {
	tt := zlsgo.NewTest(t)

	users := newViewsSchema("view_disabled", ztype.Map{
		"views": ztype.Map{
			"lists": ztype.Map{
				"fields":   []string{"name"},
				"disabled": true,
			},
			"info": ztype.Map{
				"fields":   []string{"name"},
				"disabled": true,
			},
		},
	})

	_, schemas := newTestSchemas(t, users)
	m := schemas.MustGet("view_disabled")

	// disabled 视图返回空字段
	tt.Equal(true, len(m.GetViewFields("lists")) == 0)
	tt.Equal(true, len(m.GetViewFields("info")) == 0)
}

func TestParseLables(t *testing.T) {
	tt := zlsgo.NewTest(t)

	users := newViewsSchema("view_labels", nil)
	_, schemas := newTestSchemas(t, users)
	m := schemas.MustGet("view_labels")

	_, err := m.Model().Repository().Insert(ztype.Map{"name": "a", "email": "a@b.com", "age": 18, "status": 1})
	tt.NoError(err)
	rows, err := m.Model().Repository().Find(Eq("name", "a"))
	tt.NoError(err)
	tt.Equal(1, len(rows))

	out := m.ParseLables(rows)
	// 命中枚举生成 *_label 字段
	tt.Equal("启用", out[0].Get("status_label").String())
	// 无枚举字段不生成 label
	tt.Equal("", out[0].Get("name_label").String())
	// 不存在的字段被安全跳过
	_ = m.ParseLables(ztype.Maps{{"not_exists": "1"}})
}

func TestViewsCryptID(t *testing.T) {
	tt := zlsgo.NewTest(t)

	b := true
	s := newViewsSchema("view_crypt", ztype.Map{
		"views": ztype.Map{
			"lists": ztype.Map{
				"fields": []string{"name"},
			},
			"info": ztype.Map{
				"fields": []string{"name"},
			},
		},
	})
	s.Options = schema.Options{CryptID: &b, Salt: "v-salt", CryptLen: 8}

	_, schemas := newTestSchemas(t, s)
	m := schemas.MustGet("view_crypt")

	lists := m.GetViews().Get("lists")
	// 加密 ID 时 id 列类型调整为 string
	tt.Equal("string", lists.Get("columns").Get(idKey).Get("type").String())
	info := m.GetViews().Get("info")
	tt.Equal("string", info.Get("columns").Get(idKey).Get("type").String())
}
