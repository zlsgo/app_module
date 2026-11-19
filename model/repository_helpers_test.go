package model

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/znet"
)

// TestRepositoryAccessors 确认仓储能正确暴露底层 store 与 schema。
func TestRepositoryAccessors(t *testing.T) {
	tt := zlsgo.NewTest(t)
	repo := newQueryRepo(t, "repo_batch")

	tt.Equal(true, repo.Store() != nil)
	tt.Equal(true, repo.Schema() != nil)
	tt.Equal("repo_batch", repo.Schema().define.Table.Name)
	tt.Equal(repo.Store().schema, repo.Schema())
}

// TestRepositoryCRUDHelpers 用空表跑一遍仓储的增删改查便捷方法。
func TestRepositoryCRUDHelpers(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "repo_crud")
	repo := Repo[TestUser, TestUserFilter, TestUserCreate, TestUserPatch](m.Model())

	many, err := repo.InsertMany([]TestUserCreate{
		{Name: "u1", Email: "{{EMAIL_nmv7cgof}}", Age: 10, Status: 1},
		{Name: "u2", Email: "{{EMAIL_ytr3hldb}}", Age: 20, Status: 1},
		{Name: "u3", Email: "{{EMAIL_nghnrbyl}}", Age: 30, Status: 2},
	})
	tt.NoError(err)
	tt.Equal(true, many != nil)

	count, err := repo.Count(TestUserFilter{Status: 1})
	tt.NoError(err)
	tt.Equal(uint64(2), count)

	exists, err := repo.Exists(TestUserFilter{Name: "u3"})
	tt.NoError(err)
	tt.Equal(true, exists)

	updated, err := repo.Update(TestUserFilter{Status: 1}, TestUserPatch{Age: 99})
	tt.NoError(err)
	// Update 内部固定 Limit=1，语义是「只改第一条」，所以命中 2 条也只返回 1
	tt.Equal(int64(1), updated)

	updated, err = repo.UpdateMany(TestUserFilter{Status: 2}, TestUserPatch{Name: "u3-renamed"})
	tt.NoError(err)
	tt.Equal(int64(1), updated)

	got, err := repo.FindOne(TestUserFilter{Name: "u3-renamed"})
	tt.NoError(err)
	tt.Equal("u3-renamed", got.Name)

	deleted, err := repo.Delete(TestUserFilter{Status: 2})
	tt.NoError(err)
	tt.Equal(int64(1), deleted)

	deleted, err = repo.DeleteMany(TestUserFilter{Status: 1})
	tt.NoError(err)
	tt.Equal(int64(2), deleted)

	count, err = repo.Count(TestUserFilter{})
	tt.NoError(err)
	tt.Equal(uint64(0), count)
}

// TestQueryNotIn 覆盖 NotIn 过滤条件（在 store 层用，Query 只接受领域 filter 类型）。
func TestQueryNotIn(t *testing.T) {
	tt := zlsgo.NewTest(t)
	repo := newQueryRepo(t, "query_not_in")

	rows, err := repo.Store().Find(NotIn("name", []string{"alice"}))
	tt.NoError(err)
	tt.Equal(2, len(rows))

	rows, err = repo.Store().Find(NotIn("name", []string{"alice", "bob"}))
	tt.NoError(err)
	tt.Equal(1, len(rows))
	tt.Equal("carol", rows[0].Get("name").String())
}

// TestNewRepositoryConstructors 覆盖三个仓储构造函数与 StorageType 文本。
func TestNewRepositoryConstructors(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "repo_ctor")

	store := m.Model()

	r1 := NewRepository[TestUser, TestUserFilter, TestUserCreate, TestUserPatch](store, StructMapper[TestUser]{})
	tt.Equal(store, r1.Store())

	r2 := NewStructRepository[TestUser, TestUserFilter, TestUserCreate, TestUserPatch](store)
	tt.Equal(store, r2.Store())

	r3 := NewMapRepository(store)
	tt.Equal(store, r3.Store())

	tt.Equal("sql", SQLStorage.String())
	tt.Equal("nosql", NoSQLStorage.String())
	tt.Equal("unknown", StorageType(0).String())
}

// TestVarPages 校验分页参数解析的兜底逻辑。
func TestVarPages(t *testing.T) {
	tt := zlsgo.NewTest(t)

	run := func(query string) (int, int) {
		// 每个用例独立引擎：同一引擎重复注册同一路由会被忽略，导致 handler 不生效
		e := znet.New()
		w := httptest.NewRecorder()
		var page, pagesize int
		e.GET("/", func(c *znet.Context) {
			page, pagesize, _ = Common.VarPages(c)
			c.String(http.StatusOK, "ok")
		})
		e.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/"+query, nil))
		return page, pagesize
	}

	page, pagesize := run("?page=3&pagesize=20")
	tt.Equal(3, page)
	tt.Equal(20, pagesize)

	// 缺省与非法值都回落到 1 / 10
	page, pagesize = run("")
	tt.Equal(1, page)
	tt.Equal(10, pagesize)

	page, pagesize = run("?page=0&pagesize=-5")
	tt.Equal(1, page)
	tt.Equal(10, pagesize)

	page, pagesize = run("?page=abc&pagesize=")
	tt.Equal(1, page)
	tt.Equal(10, pagesize)
}
