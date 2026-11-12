package database

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/zlsgo/zdb/driver"
)

func TestPartitionsFailedOpenCanBeRetried(t *testing.T) {
	dir := t.TempDir()
	blocker := filepath.Join(dir, "blocked")
	if err := os.WriteFile(blocker, []byte("temporary obstruction"), 0o644); err != nil {
		t.Fatal(err)
	}
	_, _, base := newTestPartitions(t, 0)
	partitions, err := NewPartitions(func(o *PartitionOptions) {
		o.DB = base.DB
		o.Prefix = base.Prefix
		o.Sqlite = &Sqlite{Path: filepath.Join(blocker, "db.db")}
		o.Max = 1
	})
	if err != nil {
		t.Fatalf("new partitions: %v", err)
	}
	t.Cleanup(func() { _ = partitions.Close() })

	if _, err := partitions.Get("west"); err == nil {
		t.Fatal("opening under a regular file must fail")
	}
	// 仅修复本测试创建的临时阻挡文件，不改管理器的配置。
	if err := os.Remove(blocker); err != nil {
		t.Fatal(err)
	}
	storage, err := partitions.Get("west")
	if err != nil {
		t.Fatalf("retry after repairing the filesystem: %v", err)
	}
	if _, err := storage.DB.QueryToMaps("SELECT 1"); err != nil {
		t.Fatalf("retried partition is unusable: %v", err)
	}
	if _, err := os.Stat(filepath.Join(blocker, "db_west.db")); err != nil {
		t.Fatalf("retried file: %v", err)
	}
	if _, err := partitions.Get("east"); !errors.Is(err, ErrPartitionLimit) {
		t.Fatalf("successful retry did not consume exactly one slot: %v", err)
	}
}

func TestPartitionsOpenPanicCanBeRetried(t *testing.T) {
	dir := t.TempDir()
	_, _, base := newTestPartitions(t, 0)
	partitions, err := NewPartitions(func(o *PartitionOptions) {
		o.DB = base.DB
		o.Prefix = base.Prefix
		o.Sqlite = &Sqlite{Path: filepath.Join(dir, "db.db")}
		o.Max = 1
	})
	if err != nil {
		t.Fatalf("new partitions: %v", err)
	}
	t.Cleanup(func() { _ = partitions.Close() })

	original, ok := getDriver("sqlite")
	if !ok {
		t.Fatal("sqlite is not registered")
	}
	driversMu.Lock()
	drivers["sqlite"] = func(opts Options) (driver.IfeConfig, error) {
		if opts.Sqlite != nil && opts.Sqlite.Path == filepath.Join(dir, "db_boom.db") {
			panic("driver boom")
		}
		return original(opts)
	}
	driversMu.Unlock()

	restored := false
	restore := func() {
		if restored {
			return
		}
		restored = true
		driversMu.Lock()
		drivers["sqlite"] = original
		driversMu.Unlock()
	}
	t.Cleanup(restore)

	if _, err := partitions.Get("boom"); err == nil || !strings.Contains(err.Error(), "panic: driver boom") {
		t.Fatalf("driver panic must surface as an error, got %v", err)
	}
	// 恢复驱动后重试同一个键：panic 不应让 entry 停在半初始化状态，
	// 也不应该占掉 Max 配额。
	restore()

	storage, err := partitions.Get("boom")
	if err != nil {
		t.Fatalf("retry after panic: %v", err)
	}
	if _, err := storage.DB.QueryToMaps("SELECT 1"); err != nil {
		t.Fatalf("retried partition is unusable: %v", err)
	}
	if _, err := partitions.Get("other"); !errors.Is(err, ErrPartitionLimit) {
		t.Fatalf("retry did not consume exactly one slot: %v", err)
	}
}

func TestPartitionsCloseWaitsForOpeningConnections(t *testing.T) {
	_, _, base := newTestPartitions(t, 0)
	path := filepath.Join(t.TempDir(), "db.db")
	partitions, err := NewPartitions(func(o *PartitionOptions) {
		o.DB = base.DB
		o.Prefix = base.Prefix
		o.Sqlite = &Sqlite{Path: path}
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = partitions.Close() })

	// 在真实 SQLite 工厂外设置调度门，不替换数据库连接或 SQL 行为。
	// 测试不并行，且只阻塞它自己的临时路径。
	original, ok := getDriver("sqlite")
	if !ok {
		t.Fatal("sqlite is not registered")
	}
	started := make(chan driver.IfeConfig, 1)
	release := make(chan struct{})
	var releaseOnce sync.Once
	resume := func() { releaseOnce.Do(func() { close(release) }) }
	driversMu.Lock()
	drivers["sqlite"] = func(opts Options) (driver.IfeConfig, error) {
		config, err := original(opts)
		if err == nil && opts.Sqlite.Path == filepath.Join(filepath.Dir(path), "db_slow.db") {
			started <- config
			<-release
		}
		return config, err
	}
	driversMu.Unlock()
	t.Cleanup(func() {
		resume()
		driversMu.Lock()
		drivers["sqlite"] = original
		driversMu.Unlock()
	})

	type result struct {
		storage Storage
		err     error
	}
	opened := make(chan result, 1)
	go func() {
		storage, err := partitions.Get("slow")
		opened <- result{storage: storage, err: err}
	}()
	var config driver.IfeConfig
	select {
	case config = <-started:
	case <-time.After(2 * time.Second):
		t.Fatal("the opening connection did not reach the scheduling gate")
	}
	t.Cleanup(func() { _ = config.DB().Close() })

	closed := make(chan error, 2)
	for i := 0; i < 2; i++ {
		go func() { closed <- partitions.Close() }()
	}
	// 先观察公共接口上的关闭边界，避免仅凭 sleep 推断 Close 已经开始。
	deadline := time.Now().Add(2 * time.Second)
	for {
		if _, err := partitions.Get(""); errors.Is(err, ErrPartitionsClosed) {
			break
		}
		if time.Now().After(deadline) {
			t.Fatal("Close did not reject new storage requests")
		}
		runtime.Gosched()
	}

	earlyReplies := 0
	select {
	case <-closed:
		earlyReplies++
	case <-time.After(25 * time.Millisecond):
	}
	resume()
	var got result
	select {
	case got = <-opened:
	case <-time.After(2 * time.Second):
		t.Fatal("Storage did not finish after releasing the gate")
	}
	for replies := earlyReplies; replies < 2; replies++ {
		select {
		case err := <-closed:
			if err != nil {
				t.Errorf("Close: %v", err)
			}
		case <-time.After(2 * time.Second):
			t.Fatal("concurrent Close did not finish")
		}
	}
	if earlyReplies != 0 {
		t.Error("Close returned before the in-flight connection could be owned and closed")
	}
	if !errors.Is(got.err, ErrPartitionsClosed) || got.storage.DB != nil {
		t.Errorf("in-flight Storage crossed the close boundary: %v, %v", got.storage, got.err)
	}
	if err := config.DB().Ping(); err == nil {
		t.Error("Close leaked the connection created while it was waiting")
	}
	if _, err := base.DB.QueryToMaps("SELECT 1"); err != nil {
		t.Errorf("Close touched the caller-owned connection: %v", err)
	}
}
