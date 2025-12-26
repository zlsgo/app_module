package model

import (
	"errors"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

type TestUser struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Status int8   `json:"status"`
}

type TestUserCreate struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Status int8   `json:"status"`
}

type TestUserPatch struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type TestUserFilter struct {
	Name   string `json:"name,omitempty"`
	Status int8   `json:"status,omitempty"`
}

func newTestDB(t *testing.T, tableName string) (*zdb.DB, *Schema) {
	db, err := zdb.New(&sqlite3.Config{
		File:       ":memory:",
		Memory:     true,
		Parameters: "_pragma=busy_timeout(3000)",
	})
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}

	b := true
	userSchema := schema.Schema{
		Name: tableName,
		Table: schema.Table{
			Name:    tableName,
			Comment: "Test Users",
		},
		Options: schema.Options{
			Timestamps: &b,
		},
		Fields: map[string]schema.Field{
			"name": {
				Type:  "string",
				Label: "Name",
				Size:  100,
			},
			"email": {
				Type:  "string",
				Label: "Email",
				Size:  200,
			},
			"age": {
				Type:    "int",
				Label:   "Age",
				Default: "0",
			},
			"status": {
				Type:    "int8",
				Label:   "Status",
				Default: "1",
			},
		},
	}

	schemas := NewSchemas(nil, NewSQL(db, ""), SchemaOptions{})
	m, err := schemas.Reg(tableName, userSchema, false)
	if err != nil {
		t.Fatalf("failed to register schema: %v", err)
	}

	err = m.Migration().Auto(DealOldColumnNone)
	if err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	return db, m
}

func TestRepositoryMapCRUD(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "crud_map")

	repo := m.Model().Repository()

	id, err := repo.Insert(ztype.Map{
		"name":   "Alice",
		"email":  "alice@example.com",
		"age":    25,
		"status": 1,
	})
	tt.NoError(err)
	tt.Log("Insert ID:", id)

	user, err := repo.FindByID(id)
	tt.NoError(err)
	tt.Equal("Alice", user.Get("name").String())
	tt.Equal("alice@example.com", user.Get("email").String())
	tt.Equal(25, user.Get("age").Int())

	count, err := repo.Update(ID(id), ztype.Map{"age": 26})
	tt.NoError(err)
	tt.Equal(int64(1), count)

	user, err = repo.FindByID(id)
	tt.NoError(err)
	tt.Equal(26, user.Get("age").Int())

	count, err = repo.DeleteByID(id)
	tt.NoError(err)
	tt.Equal(int64(1), count)

	_, err = repo.FindByID(id)
	tt.Equal(true, errors.Is(err, ErrNoRecord))
}

func TestRepositoryStructCRUD(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "crud_struct")

	repo := Repo[TestUser, TestUserFilter, TestUserCreate, TestUserPatch](m.Model())

	id, err := repo.Insert(TestUserCreate{
		Name:   "Bob",
		Email:  "bob@example.com",
		Age:    30,
		Status: 1,
	})
	tt.NoError(err)
	tt.Log("Insert ID:", id)

	user, err := repo.FindByID(id)
	tt.NoError(err)
	tt.Equal("Bob", user.Name)
	tt.Equal("bob@example.com", user.Email)
	tt.Equal(30, user.Age)

	count, err := repo.UpdateByID(id, TestUserPatch{Name: "Bobby", Age: 31})
	tt.NoError(err)
	tt.Equal(int64(1), count)

	user, err = repo.FindByID(id)
	tt.NoError(err)
	tt.Equal("Bobby", user.Name)
	tt.Equal(31, user.Age)
}

func TestRepositoryQuery(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "query_basic")

	repo := m.Model().Repository()

	testData := []ztype.Map{
		{"name": "User1", "email": "user1@test.com", "age": 20, "status": 1},
		{"name": "User2", "email": "user2@test.com", "age": 25, "status": 1},
		{"name": "User3", "email": "user3@test.com", "age": 30, "status": 2},
		{"name": "User4", "email": "user4@test.com", "age": 35, "status": 1},
		{"name": "User5", "email": "user5@test.com", "age": 40, "status": 2},
	}
	for _, data := range testData {
		_, err := repo.Insert(data)
		tt.NoError(err)
	}

	users, err := repo.Query().
		Where("status", 1).
		Find()
	tt.NoError(err)
	tt.Equal(3, len(users))

	users, err = repo.Query().
		WhereGt("age", 25).
		Find()
	tt.NoError(err)
	tt.Equal(3, len(users))

	users, err = repo.Query().
		Where("status", 1).
		OrderByDesc("age").
		Limit(2).
		Find()
	tt.NoError(err)
	tt.Equal(2, len(users))
	tt.Equal("User4", users[0].Get("name").String())

	count, err := repo.Query().
		Where("status", 1).
		Count()
	tt.NoError(err)
	tt.Equal(uint64(3), count)

	exists, err := repo.Query().
		Where("name", "User1").
		Exists()
	tt.NoError(err)
	tt.Equal(true, exists)

	exists, err = repo.Query().
		Where("name", "NotExist").
		Exists()
	tt.NoError(err)
	tt.Equal(false, exists)
}

func TestRepositoryQueryOr(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "query_or")

	repo := m.Model().Repository()

	testData := []ztype.Map{
		{"name": "Admin", "email": "admin@test.com", "age": 30, "status": 1},
		{"name": "Editor", "email": "editor@test.com", "age": 25, "status": 1},
		{"name": "Viewer", "email": "viewer@test.com", "age": 20, "status": 2},
	}
	for _, data := range testData {
		_, err := repo.Insert(data)
		tt.NoError(err)
	}

	users, err := repo.Find(Or(Eq("name", "Admin"), Eq("name", "Editor")))
	tt.NoError(err)
	tt.Equal(2, len(users))

	users, err = repo.Query().
		OrWhere(Eq("name", "Admin"), Eq("status", 2)).
		Find()
	tt.NoError(err)
	tt.Equal(2, len(users))
}

func TestRepositoryQueryAdvanced(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "query_advanced")

	repo := m.Model().Repository()

	testData := []ztype.Map{
		{"name": "Test1", "email": "test1@test.com", "age": 18, "status": 1},
		{"name": "Test2", "email": "test2@test.com", "age": 25, "status": 1},
		{"name": "Test3", "email": "test3@test.com", "age": 35, "status": 2},
		{"name": "Test4", "email": "test4@test.com", "age": 45, "status": 1},
	}
	for _, data := range testData {
		_, err := repo.Insert(data)
		tt.NoError(err)
	}

	users, err := repo.Query().
		WhereBetween("age", 20, 40).
		Find()
	tt.NoError(err)
	tt.Equal(2, len(users))

	users, err = repo.Query().
		WhereIn("name", []string{"Test1", "Test3"}).
		Find()
	tt.NoError(err)
	tt.Equal(2, len(users))

	users, err = repo.Query().
		WhereLike("email", "%test.com").
		Find()
	tt.NoError(err)
	tt.Equal(4, len(users))

	users, err = repo.Query().
		Select("name", "age").
		Where("status", 1).
		Find()
	tt.NoError(err)
	tt.Equal(3, len(users))
}

func TestRepositoryPages(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "query_pages")

	repo := m.Model().Repository()

	for i := 1; i <= 25; i++ {
		_, err := repo.Insert(ztype.Map{
			"name":   "PageUser" + ztype.ToString(i),
			"email":  "page" + ztype.ToString(i) + "@test.com",
			"age":    20 + i,
			"status": 1,
		})
		tt.NoError(err)
	}

	page1, err := repo.Query().
		OrderBy("age").
		Pages(1, 10)
	tt.NoError(err)
	tt.Equal(10, len(page1.Items))
	tt.Equal(uint(25), page1.Page.Total)

	page2, err := repo.Query().
		OrderBy("age").
		Pages(2, 10)
	tt.NoError(err)
	tt.Equal(10, len(page2.Items))

	page3, err := repo.Query().
		OrderBy("age").
		Pages(3, 10)
	tt.NoError(err)
	tt.Equal(5, len(page3.Items))
}

func TestRepositoryStructQuery(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "query_struct")

	repo := Repo[TestUser, TestUserFilter, TestUserCreate, TestUserPatch](m.Model())

	testData := []TestUserCreate{
		{Name: "StructUser1", Email: "struct1@test.com", Age: 22, Status: 1},
		{Name: "StructUser2", Email: "struct2@test.com", Age: 28, Status: 2},
		{Name: "StructUser3", Email: "struct3@test.com", Age: 33, Status: 1},
	}
	for _, data := range testData {
		_, err := repo.Insert(data)
		tt.NoError(err)
	}

	users, err := repo.Query().
		Where("status", 1).
		OrderBy("age").
		Find()
	tt.NoError(err)
	tt.Equal(2, len(users))
	tt.Equal("StructUser1", users[0].Name)
	tt.Equal(22, users[0].Age)
	tt.Equal("StructUser3", users[1].Name)

	user, err := repo.Query().
		Where("name", "StructUser2").
		FindOne()
	tt.NoError(err)
	tt.Equal("StructUser2", user.Name)
	tt.Equal(28, user.Age)
	tt.Equal(int8(2), user.Status)
}

func TestRepositoryStructFilter(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "struct_filter")

	repo := Repo[TestUser, TestUserFilter, TestUserCreate, TestUserPatch](m.Model())

	_, err := repo.Insert(TestUserCreate{
		Name:   "FilterUser1",
		Email:  "f1@test.com",
		Age:    20,
		Status: 1,
	})
	tt.NoError(err)

	_, err = repo.Insert(TestUserCreate{
		Name:   "FilterUser2",
		Email:  "f2@test.com",
		Age:    25,
		Status: 2,
	})
	tt.NoError(err)

	users, err := repo.Find(TestUserFilter{Status: 1})
	tt.NoError(err)
	tt.Equal(1, len(users))

	users, err = repo.Query().WhereFilter(TestUserFilter{Name: "FilterUser2"}).Find()
	tt.NoError(err)
	tt.Equal(1, len(users))
	tt.Equal("FilterUser2", users[0].Name)
}

func TestRepositoryGeneric(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "typed_repo")

	repo := Repo[TestUser, TestUserFilter, TestUserCreate, TestUserPatch](m.Model())

	id, err := repo.Insert(TestUserCreate{
		Name:   "TypedUser",
		Email:  "typed@test.com",
		Age:    26,
		Status: 1,
	})
	tt.NoError(err)
	tt.Equal(true, id != nil)

	users, err := repo.Find(TestUserFilter{Status: 1})
	tt.NoError(err)
	tt.Equal(1, len(users))

	count, err := repo.Update(TestUserFilter{Name: "TypedUser"}, TestUserPatch{Age: 27})
	tt.NoError(err)
	tt.Equal(int64(1), count)

	pageData, err := repo.Pages(1, 10, TestUserFilter{Status: 1})
	tt.NoError(err)
	tt.Equal(1, len(pageData.Items))

	users, err = repo.Query().WhereFilter(TestUserFilter{Status: 1}).Find()
	tt.NoError(err)
	tt.Equal(1, len(users))
}

func TestRepositoryTransaction(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "tx_test")

	repo := m.Model().Repository()

	err := repo.Tx(func(txRepo *Repository[ztype.Map, QueryFilter, ztype.Map, ztype.Map]) error {
		_, err := txRepo.Insert(ztype.Map{
			"name":   "TxUser1",
			"email":  "tx1@test.com",
			"age":    25,
			"status": 1,
		})
		if err != nil {
			return err
		}

		_, err = txRepo.Insert(ztype.Map{
			"name":   "TxUser2",
			"email":  "tx2@test.com",
			"age":    30,
			"status": 1,
		})
		return err
	})
	tt.NoError(err)

	count, err := repo.Query().Count()
	tt.NoError(err)
	tt.Equal(uint64(2), count)
}

func TestRepositoryTransactionRollback(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "tx_rollback")

	repo := m.Model().Repository()

	_, err := repo.Insert(ztype.Map{
		"name":   "ExistingUser",
		"email":  "existing@test.com",
		"age":    20,
		"status": 1,
	})
	tt.NoError(err)

	err = repo.Tx(func(txRepo *Repository[ztype.Map, QueryFilter, ztype.Map, ztype.Map]) error {
		_, err := txRepo.Insert(ztype.Map{
			"name":   "TxUser",
			"email":  "tx@test.com",
			"age":    25,
			"status": 1,
		})
		if err != nil {
			return err
		}
		return errors.New("rollback test")
	})
	tt.Equal(true, err != nil)

	count, err := repo.Query().Count()
	tt.NoError(err)
	tt.Equal(uint64(1), count)
}

func TestRepositoryBatchInsert(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "batch_insert")

	repo := m.Model().Repository()

	data := make(ztype.Maps, 50)
	for i := 0; i < 50; i++ {
		data[i] = ztype.Map{
			"name":   "BatchUser" + ztype.ToString(i),
			"email":  "batch" + ztype.ToString(i) + "@test.com",
			"age":    20 + i%30,
			"status": 1,
		}
	}

	ids, err := repo.BatchInsert(data, BatchSize(10))
	tt.NoError(err)
	tt.Equal(50, len(ids))

	count, err := repo.Query().Count()
	tt.NoError(err)
	tt.Equal(uint64(50), count)
}

func TestRepositoryBatchInsertTx(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "batch_insert_tx")

	repo := m.Model().Repository()

	data := make(ztype.Maps, 30)
	for i := 0; i < 30; i++ {
		data[i] = ztype.Map{
			"name":   "BatchTxUser" + ztype.ToString(i),
			"email":  "batchtx" + ztype.ToString(i) + "@test.com",
			"age":    20 + i%30,
			"status": 1,
		}
	}

	ids, err := repo.BatchInsertTx(data, BatchSize(10))
	tt.NoError(err)
	tt.Equal(30, len(ids))

	count, err := repo.Query().Count()
	tt.NoError(err)
	tt.Equal(uint64(30), count)
}

func TestRepositoryRepoFunc(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "repo_func")

	repo := Repo[TestUser, TestUserFilter, TestUserCreate, TestUserPatch](m.Model())

	_, err := repo.Insert(TestUserCreate{
		Name:   "RepoUser",
		Email:  "repo@test.com",
		Age:    25,
		Status: 1,
	})
	tt.NoError(err)

	user, err := repo.Query().Where("name", "RepoUser").FindOne()
	tt.NoError(err)
	tt.Equal("RepoUser", user.Name)
	tt.Equal(25, user.Age)
}

func TestRepositoryFindByIDs(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "find_by_ids")

	repo := m.Model().Repository()

	ids := make([]any, 3)
	for i := 0; i < 3; i++ {
		id, err := repo.Insert(ztype.Map{
			"name":   "User" + ztype.ToString(i+1),
			"email":  "user" + ztype.ToString(i+1) + "@test.com",
			"age":    20 + i,
			"status": 1,
		})
		tt.NoError(err)
		ids[i] = id
	}

	users, err := repo.FindByIDs(ids)
	tt.NoError(err)
	tt.Equal(3, len(users))
}

func TestRepositoryAll(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "find_all")

	repo := m.Model().Repository()

	for i := 0; i < 5; i++ {
		_, err := repo.Insert(ztype.Map{
			"name":   "AllUser" + ztype.ToString(i+1),
			"email":  "all" + ztype.ToString(i+1) + "@test.com",
			"age":    20 + i,
			"status": 1,
		})
		tt.NoError(err)
	}

	users, err := repo.All()
	tt.NoError(err)
	tt.Equal(5, len(users))
}

func TestRepositoryFirst(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "first_method")

	repo := m.Model().Repository()

	_, err := repo.Insert(ztype.Map{
		"name":   "FirstUser",
		"email":  "first@test.com",
		"age":    30,
		"status": 1,
	})
	tt.NoError(err)

	user, err := repo.First(Eq("name", "FirstUser"))
	tt.NoError(err)
	tt.Equal("FirstUser", user.Get("name").String())
}

func TestRepositoryUpdateByIDs(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "update_by_ids")

	repo := m.Model().Repository()

	ids := make([]any, 3)
	for i := 0; i < 3; i++ {
		id, err := repo.Insert(ztype.Map{
			"name":   "UpdateUser" + ztype.ToString(i+1),
			"email":  "update" + ztype.ToString(i+1) + "@test.com",
			"age":    20,
			"status": 1,
		})
		tt.NoError(err)
		ids[i] = id
	}

	count, err := repo.UpdateByIDs(ids, ztype.Map{"age": 25})
	tt.NoError(err)
	tt.Equal(int64(3), count)

	users, err := repo.FindByIDs(ids)
	tt.NoError(err)
	for _, user := range users {
		tt.Equal(25, user.Get("age").Int())
	}
}

func TestRepositoryDeleteByIDs(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "delete_by_ids")

	repo := m.Model().Repository()

	ids := make([]any, 3)
	for i := 0; i < 3; i++ {
		id, err := repo.Insert(ztype.Map{
			"name":   "DeleteUser" + ztype.ToString(i+1),
			"email":  "delete" + ztype.ToString(i+1) + "@test.com",
			"age":    20,
			"status": 1,
		})
		tt.NoError(err)
		ids[i] = id
	}

	count, err := repo.DeleteByIDs(ids)
	tt.NoError(err)
	tt.Equal(int64(3), count)

	users, err := repo.FindByIDs(ids)
	tt.NoError(err)
	tt.Equal(0, len(users))
}
