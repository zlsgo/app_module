package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

// newOperationSchema 构造通用操作模型
func newOperationSchema(name string) schema.Schema {
	return schema.Schema{
		Name:  name,
		Table: schema.Table{Name: name},
		Fields: map[string]schema.Field{
			"name":  {Type: "string", Label: "名称", Size: 80},
			"email": {Type: "string", Label: "邮箱", Size: 120},
			"age":   {Type: "int", Label: "年龄", Default: 18},
		},
	}
}

func TestStoreFacadeOperations(t *testing.T) {
	tt := zlsgo.NewTest(t)

	users := newOperationSchema("op_users")
	_, schemas := newTestSchemas(t, users)
	m := schemas.MustGet("op_users")
	store := m.Model()

	// Schema 返回模型定义
	tt.Equal("op_users", store.Schema().GetName())

	// Insert + FindOneByID
	id, err := store.Insert(ztype.Map{"name": "张三", "email": "z@b.com"})
	tt.NoError(err)
	row, err := store.FindOneByID(id)
	tt.NoError(err)
	tt.Equal("张三", row.Get("name").String())

	// Count / Exists
	n, err := store.Count(ID(id))
	tt.NoError(err)
	tt.Equal(uint64(1), n)
	ok, err := store.Exists(ID(id))
	tt.NoError(err)
	tt.Equal(true, ok)

	// FindCols
	names, err := store.FindCols("name", ID(id))
	tt.NoError(err)
	tt.Equal(1, len(names))
	tt.Equal("张三", names[0].String())

	// UpdateByID
	total, err := store.UpdateByID(id, ztype.Map{"age": 30})
	tt.NoError(err)
	tt.Equal(int64(1), total)
	row2, _ := store.FindOneByID(id)
	tt.Equal(int64(30), row2.Get("age").Int64())

	// DeleteByID
	total, err = store.DeleteByID(id)
	tt.NoError(err)
	tt.Equal(int64(1), total)
	n2, _ := store.Count(ID(id))
	tt.Equal(uint64(0), n2)
}

func TestFilterBuilder(t *testing.T) {
	tt := zlsgo.NewTest(t)

	f := NewFilter().
		Set("name", "张三").
		Set("age", 30)

	tt.Equal("张三", f.Get("name").String())
	tt.Equal(int64(30), f.Get("age").Int64())
	tt.Equal(true, f.ToMap().Has("name"))
	tt.Equal(2, len(f.ToMap()))

	// 空条件查询不被误判为缺失
	empty := NewFilter()
	tt.Equal(true, empty.ToMap() != nil)
}

func TestStoresContainer(t *testing.T) {
	tt := zlsgo.NewTest(t)

	a := newOperationSchema("store_a")
	b := newOperationSchema("store_b")
	_, schemas := newTestSchemas(t, a, b)

	stores := schemas.Models()

	// Get
	s, ok := stores.Get("store_a")
	tt.Equal(true, ok)
	tt.Equal("store_a", s.Schema().GetName())

	// MustGet
	tt.Equal("store_b", stores.MustGet("store_b").Schema().GetName())

	// All
	all := stores.All()
	tt.Equal(2, len(all))

}
