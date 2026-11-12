package database_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sohaha/zlsgo/zcache"
	"github.com/zlsgo/app_module/database"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

// ExamplePartitions 演示按应用键把数据落到独立的 SQLite 文件。
//
// 每个分区有自己的连接和表结构；这里为了聚焦分区行为，只建一张 records 表。
// 生产代码通常把 storage.DB 与 storage.Prefix 交给 model.NewSQL 或业务 schema。
func ExamplePartitions() {
	dir, err := os.MkdirTemp("", "partition-example")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "app.db")
	cfg := &sqlite3.Config{File: path}
	cfg.DB().SetMaxOpenConns(1)
	cfg.DB().SetMaxIdleConns(1)

	base, err := zdb.New(cfg)
	if err != nil {
		panic(err)
	}
	defer base.Close()

	// 默认分区沿用 base 与原表前缀 app_；其余分区按需打开 app_1001.db。
	partitions, err := database.NewPartitions(func(o *database.PartitionOptions) {
		o.DB = base
		o.Prefix = "app_"
		o.Options = database.Options{
			Driver: "sqlite",
			Sqlite: &database.Sqlite{Path: path},
		}
	})
	if err != nil {
		panic(err)
	}
	defer partitions.Close()

	for _, key := range []database.Key{"", database.KeyOf(1001), database.KeyOf(1002)} {
		storage, err := partitions.Get(key)
		if err != nil {
			panic(err)
		}
		table := storage.Prefix + "records"
		if _, err := storage.DB.Exec("CREATE TABLE IF NOT EXISTS " + table + " (marker TEXT)"); err != nil {
			panic(err)
		}
		if _, err := storage.DB.Exec("INSERT INTO "+table+" (marker) VALUES (?)", key.String()); err != nil {
			panic(err)
		}
	}

	fmt.Println("default:", partitions.Default().Prefix)
	fmt.Printf("keys: %q\n", partitions.Keys())

	key1001, err := partitions.Get(database.KeyOf(1001))
	if err != nil {
		panic(err)
	}
	rows, err := key1001.DB.QueryToMaps("SELECT marker FROM " + key1001.Prefix + "records")
	if err != nil {
		panic(err)
	}
	fmt.Println("partition 1001:", rows[0]["marker"])

	// Output:
	// default: app_
	// keys: ["" "1001" "1002"]
	// partition 1001: 1001
}

// tenantDB 是一个分区里的连接与表前缀。
type tenantDB struct {
	db     *zdb.DB
	prefix string
}

// tenantResult 把初始化结果和错误一起交给缓存，失败时 ok=false 不写缓存。
type tenantResult struct {
	tenant *tenantDB
	err    error
}

// tenantStore 按租户自动路由：业务代码只传 tenantID，不接触连接与表前缀。
//
// zcache.FastCache 负责并发读写与单飞：同一个租户的初始化只会跑一次，
// 失败不写缓存，下一次访问会重试。
type tenantStore struct {
	partitions *database.Partitions
	tenants    *zcache.FastCache
}

func newTenantStore(partitions *database.Partitions) *tenantStore {
	// FastCache 按 LRU 淘汰；条目数超过容量会重新初始化，建表是
	// IF NOT EXISTS，所以安全。分区数量大时按上限调大 Cap/Bucket。
	return &tenantStore{partitions: partitions, tenants: zcache.NewFast()}
}

func (s *tenantStore) dbFor(tenantID int64) (*tenantDB, error) {
	key := database.KeyOf(tenantID)
	v, _ := s.tenants.ProvideGet(key.String(), func() (interface{}, bool) {
		storage, err := s.partitions.Get(key)
		if err != nil {
			return tenantResult{err: err}, false
		}
		tenant := &tenantDB{db: storage.DB, prefix: storage.Prefix}
		if _, err := tenant.db.Exec("CREATE TABLE IF NOT EXISTS " + tenant.prefix + "notes (content TEXT NOT NULL)"); err != nil {
			return tenantResult{err: err}, false
		}
		return tenantResult{tenant: tenant}, true
	})
	res, ok := v.(tenantResult)
	if !ok {
		return nil, errors.New("分区初始化失败")
	}
	return res.tenant, res.err
}

func (s *tenantStore) Save(tenantID int64, content string) error {
	t, err := s.dbFor(tenantID)
	if err != nil {
		return err
	}
	_, err = t.db.Exec("INSERT INTO "+t.prefix+"notes (content) VALUES (?)", content)
	return err
}

func (s *tenantStore) List(tenantID int64) ([]string, error) {
	t, err := s.dbFor(tenantID)
	if err != nil {
		return nil, err
	}
	rows, err := t.db.QueryToMaps("SELECT content FROM " + t.prefix + "notes ORDER BY content")
	if err != nil {
		return nil, err
	}
	contents := make([]string, 0, len(rows))
	for _, row := range rows {
		contents = append(contents, fmt.Sprint(row["content"]))
	}
	return contents, nil
}

// ExamplePartitions_multiTenant 演示多租户写入时按 tenantID 自动落到对应分区。
//
// 租户 1001 与 1002 各有一个 SQLite 文件；业务侧只调用 Save/List。
func ExamplePartitions_multiTenant() {
	dir, err := os.MkdirTemp("", "partition-tenant")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "app.db")
	cfg := &sqlite3.Config{File: path}
	cfg.DB().SetMaxOpenConns(1)
	cfg.DB().SetMaxIdleConns(1)
	base, err := zdb.New(cfg)
	if err != nil {
		panic(err)
	}
	defer base.Close()

	partitions, err := database.NewPartitions(func(o *database.PartitionOptions) {
		o.DB = base
		o.Prefix = "t_"
		o.Options = database.Options{
			Driver: "sqlite",
			Sqlite: &database.Sqlite{Path: path},
		}
	})
	if err != nil {
		panic(err)
	}
	defer partitions.Close()

	store := newTenantStore(partitions)
	for _, note := range []struct {
		tenantID int64
		content  string
	}{
		{1001, "alice"},
		{1002, "bob"},
		{1001, "alice-2"},
	} {
		if err := store.Save(note.tenantID, note.content); err != nil {
			panic(err)
		}
	}

	for _, tenantID := range []int64{1001, 1002} {
		notes, err := store.List(tenantID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d %v\n", tenantID, notes)
	}

	// Output:
	// 1001 [alice alice-2]
	// 1002 [bob]
}
