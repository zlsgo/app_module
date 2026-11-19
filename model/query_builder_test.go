package model

import (
	"errors"
	"testing"

	"github.com/sohaha/zlsgo"
)

// newQueryRepo 建一个带三条数据的泛型仓储，供查询构建器用例复用。
func newQueryRepo(t *testing.T, table string) *Repository[TestUser, TestUserFilter, TestUserCreate, TestUserPatch] {
	t.Helper()

	_, m := newTestDB(t, table)
	repo := Repo[TestUser, TestUserFilter, TestUserCreate, TestUserPatch](m.Model())

	for _, c := range []TestUserCreate{
		{Name: "alice", Email: "alice@test.com", Age: 20, Status: 1},
		{Name: "bob", Email: "bob@test.com", Age: 30, Status: 1},
		{Name: "carol", Email: "carol@test.com", Age: 40, Status: 2},
	} {
		if _, err := repo.Insert(c); err != nil {
			t.Fatalf("insert %s: %v", c.Name, err)
		}
	}
	return repo
}

// TestQueryWhereVariants 逐个点过 Query 上的 WhereXXX 封装。
//
// 这些方法只是 appendFilter 的一行包装，但条件写错（比如 Ne 写成 Eq）不会报错，
// 只会静默返回错数据，所以每个都断言结果集。
func TestQueryWhereVariants(t *testing.T) {
	repo := newQueryRepo(t, "query_where")

	for _, c := range []struct {
		name  string
		query func() ([]TestUser, error)
		want  []string
	}{
		{"Where", func() ([]TestUser, error) {
			return repo.Query().Where("name", "alice").Find()
		}, []string{"alice"}},
		{"WhereNot", func() ([]TestUser, error) {
			return repo.Query().WhereNot("name", "alice").Find()
		}, []string{"bob", "carol"}},
		{"WhereGt", func() ([]TestUser, error) {
			return repo.Query().WhereGt("age", 20).Find()
		}, []string{"bob", "carol"}},
		{"WhereGe", func() ([]TestUser, error) {
			return repo.Query().WhereGe("age", 30).Find()
		}, []string{"bob", "carol"}},
		{"WhereLt", func() ([]TestUser, error) {
			return repo.Query().WhereLt("age", 30).Find()
		}, []string{"alice"}},
		{"WhereLe", func() ([]TestUser, error) {
			return repo.Query().WhereLe("age", 30).Find()
		}, []string{"alice", "bob"}},
		{"WhereLike", func() ([]TestUser, error) {
			return repo.Query().WhereLike("email", "%bob%").Find()
		}, []string{"bob"}},
		{"WhereBetween", func() ([]TestUser, error) {
			return repo.Query().WhereBetween("age", 20, 30).Find()
		}, []string{"alice", "bob"}},
		{"WhereIn", func() ([]TestUser, error) {
			return repo.Query().WhereIn("name", []string{"alice", "carol"}).Find()
		}, []string{"alice", "carol"}},
		{"WhereNull", func() ([]TestUser, error) {
			// name 是非空列，IS NULL 恒为假，验证 SQL 合法且不误命中
			return repo.Query().WhereNull("name").Find()
		}, []string{}},
		{"WhereNotNull", func() ([]TestUser, error) {
			return repo.Query().WhereNotNull("name").Find()
		}, []string{"alice", "bob", "carol"}},
		{"WhereFilter", func() ([]TestUser, error) {
			return repo.Query().WhereFilter(TestUserFilter{Status: 2}).Find()
		}, []string{"carol"}},
		{"OrWhere", func() ([]TestUser, error) {
			// OrWhere 生成的是 "AND (f1 OR f2 ...)"，用来把一组并列条件括起来
			return repo.Query().Where("status", 1).
				OrWhere(TestUserFilter{Name: "alice"}, TestUserFilter{Name: "bob"}).
				OrderBy("age").Find()
		}, []string{"alice", "bob"}},
	} {
		t.Run(c.name, func(t *testing.T) {
			tt := zlsgo.NewTest(t)
			users, err := c.query()
			tt.NoError(err)
			tt.Equal(c.want, queryNames(users))
		})
	}
}

func TestQueryWhereIDAndOrdering(t *testing.T) {
	tt := zlsgo.NewTest(t)
	repo := newQueryRepo(t, "query_order")

	all, err := repo.Query().OrderBy("age").Find()
	tt.NoError(err)
	tt.Equal([]string{"alice", "bob", "carol"}, queryNames(all))

	desc, err := repo.Query().OrderByDesc("age").Find()
	tt.NoError(err)
	tt.Equal([]string{"carol", "bob", "alice"}, queryNames(desc))

	// 显式方向参数与 OrderByDesc 等价
	explicit, err := repo.Query().OrderBy("age", "DESC").Find()
	tt.NoError(err)
	tt.Equal([]string{"carol", "bob", "alice"}, queryNames(explicit))

	limited, err := repo.Query().OrderBy("age").Limit(2).Find()
	tt.NoError(err)
	tt.Equal([]string{"alice", "bob"}, queryNames(limited))

	offset, err := repo.Query().OrderBy("age").Limit(2).Offset(1).Find()
	tt.NoError(err)
	tt.Equal([]string{"bob", "carol"}, queryNames(offset))

	first, err := repo.Query().OrderBy("age").First()
	tt.NoError(err)
	tt.Equal("alice", first.Name)

	id := all[0].ID
	byID, err := repo.Query().WhereID(id).FindOne()
	tt.NoError(err)
	tt.Equal("alice", byID.Name)
}

func TestQueryAggregates(t *testing.T) {
	tt := zlsgo.NewTest(t)
	repo := newQueryRepo(t, "query_aggregate")

	count, err := repo.Query().Where("status", 1).Count()
	tt.NoError(err)
	tt.Equal(uint64(2), count)

	exists, err := repo.Query().Where("name", "carol").Exists()
	tt.NoError(err)
	tt.Equal(true, exists)

	exists, err = repo.Query().Where("name", "nobody").Exists()
	tt.NoError(err)
	tt.Equal(false, exists)

	page, err := repo.Query().OrderBy("age").Pages(2, 2)
	tt.NoError(err)
	tt.Equal(1, len(page.Items))
	tt.Equal(uint(3), page.Page.Total)

	_, err = repo.Query().Where("name", "nobody").FindOne()
	tt.Equal(true, errors.Is(err, ErrNoRecord))
}

func TestQuerySelectGroupBy(t *testing.T) {
	tt := zlsgo.NewTest(t)
	repo := newQueryRepo(t, "query_group")

	// 只取 name 列，其余字段应保持零值
	rows, err := repo.Query().Select("name").Where("name", "alice").Find()
	tt.NoError(err)
	tt.Equal(1, len(rows))
	tt.Equal("alice", rows[0].Name)
	tt.Equal("", rows[0].Email)

	groups, err := repo.Query().Select("status").GroupBy("status").OrderBy("status").Find()
	tt.NoError(err)
	tt.Equal(2, len(groups))
}

func TestQueryUpdateDelete(t *testing.T) {
	tt := zlsgo.NewTest(t)
	repo := newQueryRepo(t, "query_write")

	updated, err := repo.Query().Where("name", "carol").Update(TestUserPatch{Name: "carol2"})
	tt.NoError(err)
	tt.Equal(int64(1), updated)

	after, err := repo.Query().Where("name", "carol2").FindOne()
	tt.NoError(err)
	tt.Equal("carol2", after.Name)

	deleted, err := repo.Query().Where("status", 2).Delete()
	tt.NoError(err)
	tt.Equal(int64(1), deleted)

	count, err := repo.Query().Count()
	tt.NoError(err)
	tt.Equal(uint64(2), count)
}

// TestQueryOrWhereEmpty 确认没有参数的 OrWhere 是空操作，不会把条件清空。
func TestQueryOrWhereEmpty(t *testing.T) {
	tt := zlsgo.NewTest(t)
	repo := newQueryRepo(t, "query_or_empty")

	users, err := repo.Query().Where("name", "bob").OrWhere().Find()
	tt.NoError(err)
	tt.Equal([]string{"bob"}, queryNames(users))
}

func queryNames(users []TestUser) []string {
	names := make([]string, 0, len(users))
	for _, u := range users {
		names = append(names, u.Name)
	}
	return names
}
