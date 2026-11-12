package database

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"testing"

	"github.com/zlsgo/zdb"
)

// newTestPartitions 用临时文件创建默认库，再按零参数入口构造分区管理器。
func newTestPartitions(t *testing.T, max int) (*Partitions, string, Storage) {
	t.Helper()
	return newTestPartitionsWithPath(t, max, "db.db")
}

func newTestPartitionsWithPath(t *testing.T, max int, name string) (*Partitions, string, Storage) {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, name)
	base, err := initDB(Options{Driver: "sqlite", Sqlite: &Sqlite{Path: path}})
	if err != nil {
		t.Fatalf("init default db: %v", err)
	}
	t.Cleanup(func() { _ = base.Close() })

	originalOptions, originalDB := CurrentOptions(), CurrentDB()
	setOptions(Options{Driver: "sqlite", Sqlite: &Sqlite{Path: path}})
	setCurrentDB(base)
	t.Cleanup(func() {
		setOptions(originalOptions)
		setCurrentDB(originalDB)
	})

	partitions, err := NewPartitions(func(o *PartitionOptions) {
		o.Prefix = "data_"
		o.Max = max
	})
	if err != nil {
		t.Fatalf("new partitions: %v", err)
	}
	t.Cleanup(func() { _ = partitions.Close() })
	return partitions, dir, Storage{DB: base, Prefix: "data_"}
}

func TestPartitionsSQLiteUsesSeparateFile(t *testing.T) {
	partitions, dir, base := newTestPartitions(t, 0)
	got, err := partitions.Get("1001")
	if err != nil {
		t.Fatalf("storage: %v", err)
	}
	if got.DB == base.DB {
		t.Fatal("partition reused the default connection")
	}
	if got.Prefix != "data_" {
		t.Fatalf("prefix = %q, want data_", got.Prefix)
	}
	if _, err := os.Stat(filepath.Join(dir, "db_1001.db")); err != nil {
		t.Fatalf("partition file: %v", err)
	}
}

func TestPartitionsSameKeyIsStable(t *testing.T) {
	partitions, _, _ := newTestPartitions(t, 0)
	first, err := partitions.Get("1001")
	if err != nil {
		t.Fatalf("first: %v", err)
	}
	second, err := partitions.Get("1001")
	if err != nil {
		t.Fatalf("second: %v", err)
	}
	if first != second {
		t.Fatal("the same key returned a different storage")
	}
}

func TestPartitionsConcurrentFirstUseCreatesOnce(t *testing.T) {
	partitions, dir, _ := newTestPartitions(t, 0)
	const workers = 16
	var wg sync.WaitGroup
	seen := make([]*zdb.DB, workers)
	errs := make([]error, workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			got, err := partitions.Get("worker-pool")
			seen[idx], errs[idx] = got.DB, err
		}(i)
	}
	wg.Wait()
	for i, err := range errs {
		if err != nil {
			t.Fatalf("worker %d: %v", i, err)
		}
		if seen[i] != seen[0] {
			t.Fatalf("worker %d received a different connection", i)
		}
	}
	if _, err := os.Stat(filepath.Join(dir, "db_worker-pool.db")); err != nil {
		t.Fatalf("partition file: %v", err)
	}
	if got := partitions.Keys(); !reflect.DeepEqual(got, []Key{"", "worker-pool"}) {
		t.Fatalf("Keys() = %v", got)
	}
}

func TestPartitionsDefaultIsEmptyKey(t *testing.T) {
	partitions, dir, base := newTestPartitions(t, 0)
	got, err := partitions.Get("")
	if err != nil {
		t.Fatalf("default storage: %v", err)
	}
	if got != base || partitions.Default() != base {
		t.Fatal("empty key did not return the default partition")
	}
	zero, err := partitions.Get(KeyOf(0))
	if err != nil {
		t.Fatalf("zero key storage: %v", err)
	}
	if zero.DB == base.DB {
		t.Fatal("the numeric key 0 was interpreted as the default partition")
	}
	if _, err := os.Stat(filepath.Join(dir, "db_0.db")); err != nil {
		t.Fatalf("zero partition file: %v", err)
	}
}

func TestPartitionsRejectsInvalidKeys(t *testing.T) {
	partitions, _, _ := newTestPartitions(t, 0)
	for _, key := range []Key{"../outside", "..", ".", "a/b", "a\\b", "a.b", "a b", " a", "a ", "a\n", "a\x00", "a'", "a\"", "a;", "a%2f", "租户", "é", "a:1"} {
		t.Run(string(key), func(t *testing.T) {
			if storage, err := partitions.Get(key); !errors.Is(err, ErrInvalidKey) || storage.DB != nil {
				t.Fatalf("Get(%q) = %v, %v, want ErrInvalidKey and no connection", key, storage, err)
			}
		})
	}
	if got := partitions.Keys(); !reflect.DeepEqual(got, []Key{""}) {
		t.Fatalf("invalid keys created partitions: %v", got)
	}
}

func TestPartitionsLimit(t *testing.T) {
	partitions, _, _ := newTestPartitions(t, 2)
	if _, err := partitions.Get(""); err != nil {
		t.Fatalf("default: %v", err)
	}
	for _, key := range []Key{"first", "second"} {
		if _, err := partitions.Get(key); err != nil {
			t.Fatalf("storage(%q): %v", key, err)
		}
	}
	if _, err := partitions.Get("third"); !errors.Is(err, ErrPartitionLimit) {
		t.Fatalf("over limit: %v", err)
	}
	for _, key := range []Key{"", "first"} {
		if _, err := partitions.Get(key); err != nil {
			t.Fatalf("existing partition at limit: %v", err)
		}
	}
}

func TestPartitionsNegativeLimitIsUnlimited(t *testing.T) {
	partitions, _, _ := newTestPartitions(t, -1)
	for _, key := range []Key{"first", "second", "third"} {
		if _, err := partitions.Get(key); err != nil {
			t.Fatalf("storage(%q): %v", key, err)
		}
	}
}

func TestPartitionsKeysIncludesDefaultAndSorted(t *testing.T) {
	partitions, _, _ := newTestPartitions(t, 0)
	for _, key := range []Key{"tenant-z", "2", "10", "Tenant-a"} {
		if _, err := partitions.Get(key); err != nil {
			t.Fatalf("storage(%q): %v", key, err)
		}
	}
	want := []Key{"", "10", "2", "Tenant-a", "tenant-z"}
	if got := partitions.Keys(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Keys() = %v, want %v", got, want)
	}
	// 返回值是调用方自己的快照，修改它不会改变管理器。
	keys := partitions.Keys()
	keys[0] = "changed"
	if got := partitions.Keys(); !reflect.DeepEqual(got, want) {
		t.Fatalf("caller changed Keys(): %v", got)
	}
}

func TestPartitionsClosedRejectsGet(t *testing.T) {
	partitions, _, base := newTestPartitions(t, 0)
	if _, err := partitions.Get("first"); err != nil {
		t.Fatalf("first: %v", err)
	}
	if err := partitions.Close(); err != nil {
		t.Fatalf("close: %v", err)
	}
	if err := partitions.Close(); err != nil {
		t.Fatalf("second close: %v", err)
	}
	for _, key := range []Key{"", "first", "second"} {
		if _, err := partitions.Get(key); !errors.Is(err, ErrPartitionsClosed) {
			t.Fatalf("closed storage(%q): %v", key, err)
		}
	}
	if _, err := base.DB.QueryToMaps("SELECT 1"); err != nil {
		t.Fatalf("Close closed the caller-owned default connection: %v", err)
	}
	if got := partitions.Keys(); !reflect.DeepEqual(got, []Key{"", "first"}) {
		t.Fatalf("closed Keys() lost its snapshot: %v", got)
	}
}

func TestPartitionsRejectsMemorySQLite(t *testing.T) {
	_, _, base := newTestPartitions(t, 0)
	for _, path := range []string{":memory:", "file::memory:?cache=shared"} {
		partitions, err := NewPartitions(func(o *PartitionOptions) {
			o.DB = base.DB
			o.Prefix = base.Prefix
			o.Sqlite = &Sqlite{Path: path}
		})
		if err != nil {
			t.Fatalf("new partitions for %q: %v", path, err)
		}
		t.Cleanup(func() { _ = partitions.Close() })
		if _, err := partitions.Get("first"); err == nil {
			t.Fatalf("file partitioning accepted the in-memory database %q", path)
		}
	}
}

func TestPartitionsSQLitePathWithoutExtension(t *testing.T) {
	partitions, dir, _ := newTestPartitionsWithPath(t, 0, "store")
	if _, err := partitions.Get("west"); err != nil {
		t.Fatalf("storage: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "store_west")); err != nil {
		t.Fatalf("extensionless partition file: %v", err)
	}
}

func TestPartitionsSharedConnectionNamespaces(t *testing.T) {
	_, _, base := newTestPartitions(t, 0)
	for _, driver := range []string{"mysql", "postgres"} {
		t.Run(driver, func(t *testing.T) {
			// 只验证路由接缝，不拨号到外部服务：非 SQLite 策略必须借用 base。
			partitions, err := NewPartitions(func(o *PartitionOptions) {
				o.DB = base.DB
				o.Prefix = base.Prefix
				o.Driver = driver
			})
			if err != nil {
				t.Fatalf("new partitions: %v", err)
			}
			t.Cleanup(func() { _ = partitions.Close() })
			for _, key := range []Key{"west-2", "east_1"} {
				storage, err := partitions.Get(key)
				if err != nil {
					t.Fatalf("storage: %v", err)
				}
				want := "data_west-2_"
				if key == "east_1" {
					want = "data_east_1_"
				}
				if storage.DB != base.DB || storage.Prefix != want {
					t.Fatalf("storage(%q) = %v, want shared connection and prefix %q", key, storage, want)
				}
			}
		})
	}
}

func TestPartitionsSharedConnectionWithoutPrefix(t *testing.T) {
	_, _, base := newTestPartitions(t, 0)
	partitions, err := NewPartitions(func(o *PartitionOptions) {
		o.DB = base.DB
		o.Prefix = ""
		o.Driver = "mysql"
	})
	if err != nil {
		t.Fatalf("new partitions: %v", err)
	}
	t.Cleanup(func() { _ = partitions.Close() })

	storage, err := partitions.Get("west")
	if err != nil {
		t.Fatalf("storage: %v", err)
	}
	if storage.DB != base.DB || storage.Prefix != "west_" {
		t.Fatalf("storage = %+v, want base connection and prefix west_", storage)
	}
}

func TestPartitionsFailedOpenStaysOutOfKeys(t *testing.T) {
	dir := t.TempDir()
	blocker := filepath.Join(dir, "blocker")
	if err := os.WriteFile(blocker, []byte("not a directory"), 0o644); err != nil {
		t.Fatalf("write blocker: %v", err)
	}
	_, _, base := newTestPartitions(t, 0)
	partitions, err := NewPartitions(func(o *PartitionOptions) {
		o.DB = base.DB
		o.Prefix = base.Prefix
		o.Sqlite = &Sqlite{Path: filepath.Join(blocker, "db.db")}
	})
	if err != nil {
		t.Fatalf("new partitions: %v", err)
	}
	t.Cleanup(func() { _ = partitions.Close() })
	if _, err := partitions.Get("failed"); err == nil {
		t.Fatal("opening inside a regular file must fail")
	}
	if got := partitions.Keys(); !reflect.DeepEqual(got, []Key{""}) {
		t.Fatalf("failed partition appeared in Keys(): %v", got)
	}
}

func TestPartitionsLimitHoldsUnderConcurrentFirstVisit(t *testing.T) {
	partitions, _, _ := newTestPartitions(t, 2)
	var wg sync.WaitGroup
	succeeded := make([]bool, 8)
	for i := range succeeded {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			_, err := partitions.Get(KeyOf(int64(idx)))
			succeeded[idx] = err == nil
		}(i)
	}
	wg.Wait()
	var count int
	for _, ok := range succeeded {
		if ok {
			count++
		}
	}
	if count != 2 {
		t.Fatalf("concurrent first visits opened %d partitions, want 2", count)
	}
}
