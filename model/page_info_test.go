package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/internal/testutil"
	"github.com/zlsgo/app_module/model/schema"
)

// newPagesStore 造一个装了两行数据的模型存储，供分页断言使用。
func newPagesStore(t *testing.T) *Store {
	t.Helper()

	db := testutil.NewMemoryDB(t)
	t.Cleanup(func() { _ = db.Close() })

	mod := New(func(o *Options) {
		o.Prefix = "page_"
		o.Schemas = schema.Schemas{
			{
				Name:  "page_users",
				Table: schema.Table{Name: "page_users"},
				Fields: map[string]schema.Field{
					"name": {Type: "string", Size: 32},
				},
			},
		}
	})

	di := zdi.New()
	di.Map(db)
	app := service.NewApp()(di)
	if err := service.InitModule([]service.Module{mod}, app); err != nil {
		t.Fatalf("init module: %v", err)
	}

	store := mod.MustGetStore("page_users")
	for _, name := range []string{"a", "b"} {
		if _, err := store.Insert(ztype.Map{"name": name}); err != nil {
			t.Fatalf("insert %s: %v", name, err)
		}
	}

	return store
}

// TestPagesEmptyPageKeepsTotal 锁住越界页（空页）的分页信息。
//
// zdb 在越界页返回 ErrNotFound 时会丢掉已经统计好的 total，调用方只能看到
// total=0 / curpage=0。前端分页器拿到这份数据会认定「没有数据」并把页码重置
// 回第一页，用户翻到最后一页再点下一页就跳回开头，所以空页也必须报告真实的
// 总数与当前页号。
func TestPagesEmptyPageKeepsTotal(t *testing.T) {
	tt := zlsgo.NewTest(t)

	store := newPagesStore(t)

	// 正常页：total/count/curpage 都要对得上
	p1, err := store.Pages(1, 1, nil)
	tt.NoError(err)
	tt.Equal(1, len(p1.Items))
	tt.Equal(uint(2), p1.Page.Total)
	tt.Equal(uint(2), p1.Page.Count)
	tt.Equal(uint(1), p1.Page.Curpage)

	// 越界页：没有数据，但总数与页号仍然是真实的
	p3, err := store.Pages(3, 1, nil)
	tt.NoError(err)
	tt.Equal(0, len(p3.Items))
	tt.Equal(uint(2), p3.Page.Total)
	tt.Equal(uint(2), p3.Page.Count)
	tt.Equal(uint(3), p3.Page.Curpage)

	// 页数按调用方给的 pagesize 重算，而不是按补查时的 1 条一页
	p2, err := store.Pages(2, 10, nil)
	tt.NoError(err)
	tt.Equal(0, len(p2.Items))
	tt.Equal(uint(2), p2.Page.Total)
	tt.Equal(uint(1), p2.Page.Count)
	tt.Equal(uint(2), p2.Page.Curpage)

	// 查询条件命中 0 行：空页维持 total=0，且不因为补查而报错
	pf, err := store.Pages(1, 10, Filter{"name": "nobody"})
	tt.NoError(err)
	tt.Equal(0, len(pf.Items))
	tt.Equal(uint(0), pf.Page.Total)
}
