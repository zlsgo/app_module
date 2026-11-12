package database

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// TestNewPartitionsReusesModuleConnection 锁住零参数入口：默认连接与配置都
// 取自数据库模块，调用方只需要调分区自己的参数。
func TestNewPartitionsReusesModuleConnection(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "db.db")
	base, err := initDB(Options{Driver: "sqlite", Sqlite: &Sqlite{Path: path}})
	if err != nil {
		t.Fatal(err)
	}
	originalOptions, originalDB := CurrentOptions(), CurrentDB()
	setOptions(Options{Driver: "sqlite", Sqlite: &Sqlite{Path: path}})
	setCurrentDB(base)
	t.Cleanup(func() {
		setOptions(originalOptions)
		setCurrentDB(originalDB)
		_ = base.Close()
	})

	partitions, err := NewPartitions()
	if err != nil {
		t.Fatalf("new partitions without options: %v", err)
	}
	t.Cleanup(func() { _ = partitions.Close() })

	storage, err := partitions.Get("")
	if err != nil || storage.DB != base || storage.Prefix != "" {
		t.Fatalf("default storage = %+v, %v", storage, err)
	}
	if _, err := partitions.Get("west"); err != nil {
		t.Fatalf("open west: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "db_west.db")); err != nil {
		t.Fatalf("default path was not inherited: %v", err)
	}
}

// TestNewPartitionsNormalizesDriverName 锁住驱动名大小写不影响分区形态。
//
// 驱动注册与查询都走 normalizeName，配置里写 "SQLite" 也必须建独立文件；
// 否则同一个配置会因为大小写静默切换到共享连接 + 表名前缀。
func TestNewPartitionsNormalizesDriverName(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "db.db")
	base, err := initDB(Options{Driver: "sqlite", Sqlite: &Sqlite{Path: path}})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = base.Close() })

	partitions, err := NewPartitions(func(o *PartitionOptions) {
		o.DB = base
		o.Driver = "SQLite"
		o.Sqlite = &Sqlite{Path: path}
	})
	if err != nil {
		t.Fatalf("new partitions with mixed-case driver: %v", err)
	}
	t.Cleanup(func() { _ = partitions.Close() })

	storage, err := partitions.Get("west")
	if err != nil {
		t.Fatalf("open west: %v", err)
	}
	if storage.DB == base {
		t.Fatal("mixed-case driver fell back to the shared connection instead of a partition file")
	}
	if storage.Prefix != "" {
		t.Fatalf("partition prefix = %q, want the default prefix", storage.Prefix)
	}
	if _, err := os.Stat(filepath.Join(dir, "db_west.db")); err != nil {
		t.Fatalf("mixed-case driver did not create a partition file: %v", err)
	}
}

func TestNewPartitionsAppliesOptionsInOrder(t *testing.T) {
	_, _, base := newTestPartitions(t, 0)
	path := filepath.Join(t.TempDir(), "custom.db")
	var order []int
	partitions, err := NewPartitions(
		func(o *PartitionOptions) {
			order = append(order, 1)
			o.Prefix = "data_"
			o.Sqlite = &Sqlite{Path: path}
			o.Max = 1
		},
		func(o *PartitionOptions) {
			order = append(order, 2)
			o.Max = 2
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = partitions.Close() })
	if !reflect.DeepEqual(order, []int{1, 2}) {
		t.Fatalf("option order = %v", order)
	}
	for _, key := range []Key{"west", "east"} {
		if _, err := partitions.Get(key); err != nil {
			t.Fatalf("Get(%q): %v", key, err)
		}
		if _, err := os.Stat(filepath.Join(filepath.Dir(path), "custom_"+string(key)+".db")); err != nil {
			t.Fatalf("custom path for %q: %v", key, err)
		}
	}
	if _, err := partitions.Get("north"); !errors.Is(err, ErrPartitionLimit) {
		t.Fatalf("later option did not override the limit: %v", err)
	}
	if got, want := partitions.Default().DB, base.DB; got != want {
		t.Fatal("the default partition changed its connection")
	}
}

func TestNewPartitionsKeepsOptionsIndependent(t *testing.T) {
	originalOptions, originalDB := CurrentOptions(), CurrentDB()
	t.Cleanup(func() {
		setOptions(originalOptions)
		setCurrentDB(originalDB)
	})
	dir := t.TempDir()
	setOptions(Options{Driver: "sqlite", Sqlite: &Sqlite{Path: filepath.Join(dir, "default.db")}})
	base, err := initDB(CurrentOptions())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = base.Close() })
	setCurrentDB(base)

	var configured *PartitionOptions
	custom, err := NewPartitions(func(o *PartitionOptions) {
		o.DB = base
		o.Prefix = "data_"
		o.Sqlite.Path = filepath.Join(dir, "custom.db")
		o.Max = 2
		configured = o
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = custom.Close() })
	if CurrentOptions().Sqlite.Path != filepath.Join(dir, "default.db") {
		t.Fatal("the callback changed the database module's default options")
	}
	defaults, err := NewPartitions()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = defaults.Close() })

	// 构造完成后修改调用方持有的配置，不应重定向已经创建的管理器。
	configured.Sqlite.Path = filepath.Join(dir, "changed-custom.db")
	configured.Max = 99
	setOptions(Options{Driver: "sqlite", Sqlite: &Sqlite{Path: filepath.Join(dir, "changed-default.db")}})

	if got := custom.Default().Prefix; got != "data_" {
		t.Fatalf("custom default prefix = %q", got)
	}
	for _, key := range []Key{"west", "east"} {
		if _, err := custom.Get(key); err != nil {
			t.Fatalf("custom Get(%q): %v", key, err)
		}
		if _, err := os.Stat(filepath.Join(dir, "custom_"+string(key)+".db")); err != nil {
			t.Fatalf("custom partition changed its path: %v", err)
		}
	}
	if _, err := custom.Get("north"); !errors.Is(err, ErrPartitionLimit) {
		t.Fatalf("custom partition changed its quota: %v", err)
	}

	// defaults 在构造时复制了当时的默认配置，之后的 Reload 不会重定向它。
	if _, err := defaults.Get("west"); err != nil {
		t.Fatalf("default Get(west): %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "default_west.db")); err != nil {
		t.Fatalf("default partition changed its path: %v", err)
	}
}

func TestNewPartitionsRejectsMissingDB(t *testing.T) {
	original := CurrentDB()
	t.Cleanup(func() { setCurrentDB(original) })
	setCurrentDB(nil)

	if partitions, err := NewPartitions(); err == nil || partitions != nil {
		t.Fatalf("NewPartitions() = %v, %v, want an error when the module has no connection", partitions, err)
	}
}

func TestNewPartitionsRequiresDriver(t *testing.T) {
	_, _, base := newTestPartitions(t, 0)
	partitions, err := NewPartitions(func(o *PartitionOptions) {
		o.DB = base.DB
		o.Options = Options{}
	})
	if err == nil || partitions != nil {
		t.Fatalf("NewPartitions() = %v, %v, want an error without a driver", partitions, err)
	}
}
