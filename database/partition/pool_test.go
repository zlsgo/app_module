package partition

import (
	"errors"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type testResource struct {
	key string
}

func TestNewPoolRequiresOpen(t *testing.T) {
	if pool, err := NewPool(Config[int]{}); err == nil || pool != nil {
		t.Fatalf("NewPool() = %v, %v, want an error without an open callback", pool, err)
	}
}

func TestPoolBaseIsReturnedByEmptyKey(t *testing.T) {
	base := &testResource{key: "base"}
	pool, err := NewPool(Config[*testResource]{
		Base: base,
		Open: func(key string) (*testResource, error) {
			return &testResource{key: key}, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if got := pool.Base(); got != base {
		t.Fatalf("Base() = %#v, want the configured base", got)
	}
	if got, err := pool.Get(""); err != nil || got != base {
		t.Fatalf("Get(\"\") = %#v, %v, want the base resource", got, err)
	}
}

// TestPoolOpensKeyOnce 锁住并发首访只调用一次 Open。
func TestPoolOpensKeyOnce(t *testing.T) {
	var opens atomic.Int64
	pool, err := NewPool(Config[*testResource]{
		Open: func(key string) (*testResource, error) {
			opens.Add(1)
			// 放大首次打开的窗口，让并发调用都能进入 once 的等待路径。
			time.Sleep(2 * time.Millisecond)
			return &testResource{key: key}, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	const workers = 16
	var wg sync.WaitGroup
	values := make([]*testResource, workers)
	errs := make([]error, workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			values[index], errs[index] = pool.Get("west")
		}(i)
	}
	wg.Wait()

	if got := opens.Load(); got != 1 {
		t.Fatalf("Open called %d times, want 1", got)
	}
	for i := range values {
		if errs[i] != nil {
			t.Fatalf("worker %d: %v", i, errs[i])
		}
		if values[i] != values[0] {
			t.Fatalf("worker %d received a different resource", i)
		}
	}
	if got, want := pool.Keys(), []string{"", "west"}; !equalKeys(got, want) {
		t.Fatalf("Keys() = %v, want %v", got, want)
	}
}

// TestPoolCountsInflightAgainstLimit 锁住“在途打开也占配额”：
// 否则两个首次访问同时通过上限检查会超发。
func TestPoolCountsInflightAgainstLimit(t *testing.T) {
	started := make(chan struct{})
	release := make(chan struct{})
	var startOnce sync.Once
	pool, err := NewPool(Config[*testResource]{
		Max: 1,
		Open: func(key string) (*testResource, error) {
			if key == "slow" {
				startOnce.Do(func() { close(started) })
				<-release
			}
			return &testResource{key: key}, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	opened := make(chan result, 1)
	go func() {
		value, err := pool.Get("slow")
		opened <- result{value: value, err: err}
	}()
	select {
	case <-started:
	case <-time.After(2 * time.Second):
		t.Fatal("slow key did not start opening")
	}

	if _, err := pool.Get("west"); !errors.Is(err, ErrLimit) {
		t.Fatalf("in-flight open did not consume quota: %v", err)
	}
	if got, want := pool.Keys(), []string{""}; !equalKeys(got, want) {
		t.Fatalf("Keys() during open = %v, want %v", got, want)
	}

	close(release)
	got := <-opened
	if got.err != nil {
		t.Fatalf("slow key: %v", got.err)
	}
	if _, err := pool.Get("slow"); err != nil {
		t.Fatalf("reopen slow key: %v", err)
	}
	if _, err := pool.Get("west"); !errors.Is(err, ErrLimit) {
		t.Fatalf("quota was not consumed after open: %v", err)
	}
	if got, want := pool.Keys(), []string{"", "slow"}; !equalKeys(got, want) {
		t.Fatalf("Keys() = %v, want %v", got, want)
	}
}

// TestPoolFailureDoesNotConsumeQuota 锁住“打开失败不占配额且可以重试”。
func TestPoolFailureDoesNotConsumeQuota(t *testing.T) {
	var broken atomic.Bool
	broken.Store(true)
	pool, err := NewPool(Config[*testResource]{
		Max: 1,
		Open: func(key string) (*testResource, error) {
			if key == "broken" && broken.Load() {
				return nil, errors.New("boom")
			}
			return &testResource{key: key}, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := pool.Get("broken"); err == nil {
		t.Fatal("broken key must fail")
	}
	if got, want := pool.Keys(), []string{""}; !equalKeys(got, want) {
		t.Fatalf("failed key leaked into Keys(): %v", got)
	}
	if _, err := pool.Get("good"); err != nil {
		t.Fatalf("failure consumed quota: %v", err)
	}

	broken.Store(false)
	if _, err := pool.Get("broken"); !errors.Is(err, ErrLimit) {
		t.Fatalf("retry bypassed the limit: %v", err)
	}
}

func TestPoolPanicBecomesRetryableError(t *testing.T) {
	var attempts atomic.Int64
	pool, err := NewPool(Config[*testResource]{
		Max: 1,
		Open: func(key string) (*testResource, error) {
			if attempts.Add(1) == 1 {
				panic("driver boom")
			}
			return &testResource{key: key}, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := pool.Get("boom"); err == nil || !strings.Contains(err.Error(), "panic: driver boom") {
		t.Fatalf("panic was not converted to an error: %v", err)
	}
	if got, want := pool.Keys(), []string{""}; !equalKeys(got, want) {
		t.Fatalf("panicking key leaked into Keys(): %v", got)
	}

	value, err := pool.Get("boom")
	if err != nil {
		t.Fatalf("retry after panic: %v", err)
	}
	if value == nil || value.key != "boom" {
		t.Fatalf("retry returned %#v", value)
	}
	if _, err := pool.Get("other"); !errors.Is(err, ErrLimit) {
		t.Fatalf("retry did not consume exactly one slot: %v", err)
	}
}

// TestPoolDoesNotOwnResourcesWithoutClose 锁住共享连接模式：
// 没有 Close 回调时，池不能把调用方的资源记进待关闭列表。
func TestPoolDoesNotOwnResourcesWithoutClose(t *testing.T) {
	pool, err := NewPool(Config[*testResource]{
		Open: func(key string) (*testResource, error) {
			return &testResource{key: key}, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Get("west"); err != nil {
		t.Fatal(err)
	}
	pool.mu.RLock()
	owned := len(pool.own)
	pool.mu.RUnlock()
	if owned != 0 {
		t.Fatalf("pool owns %d resources without a Close callback", owned)
	}
	if err := pool.Close(); err != nil {
		t.Fatalf("Close without callback: %v", err)
	}
	if _, err := pool.Get("west"); !errors.Is(err, ErrClosed) {
		t.Fatalf("Get after Close = %v, want ErrClosed", err)
	}
}

// TestPoolCloseWithConcurrentGetAndKeys 锁住关闭与读操作的并发边界：
// 关闭期间 Get/Keys 可以任意交错，但不能 panic 或改变关闭后的拒绝语义。
func TestPoolCloseWithConcurrentGetAndKeys(t *testing.T) {
	pool, err := NewPool(Config[*testResource]{
		Open: func(key string) (*testResource, error) {
			return &testResource{key: key}, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Get("west"); err != nil {
		t.Fatal(err)
	}

	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			<-start
			if index%2 == 0 {
				if _, err := pool.Get("east"); err != nil && !errors.Is(err, ErrClosed) {
					t.Errorf("concurrent Get: %v", err)
				}
				return
			}
			_ = pool.Keys()
		}(i)
	}

	closed := make(chan error, 1)
	go func() {
		<-start
		closed <- pool.Close()
	}()
	close(start)
	wg.Wait()
	if err := <-closed; err != nil {
		t.Fatalf("Close: %v", err)
	}

	for _, key := range []string{"", "west", "east"} {
		if _, err := pool.Get(key); !errors.Is(err, ErrClosed) {
			t.Fatalf("Get(%q) after Close = %v, want ErrClosed", key, err)
		}
	}
}

// TestPoolCloseWaitsAndClosesOwned 锁住关闭边界：拒绝新访问、等待在途打开、
// 只关闭池自己打开的资源，并把关闭错误返回给所有并发调用。
func TestPoolCloseWaitsAndClosesOwned(t *testing.T) {
	var closes atomic.Int64
	closeErr := errors.New("close failed")
	base := &testResource{key: ""}
	started := make(chan struct{})
	release := make(chan struct{})
	var startOnce sync.Once

	pool, err := NewPool(Config[*testResource]{
		Base: base,
		Open: func(key string) (*testResource, error) {
			startOnce.Do(func() { close(started) })
			<-release
			return &testResource{key: key}, nil
		},
		Close: func(*testResource) error {
			closes.Add(1)
			return closeErr
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	opened := make(chan result, 1)
	go func() {
		value, err := pool.Get("slow")
		opened <- result{value: value, err: err}
	}()
	select {
	case <-started:
	case <-time.After(2 * time.Second):
		t.Fatal("slow key did not start opening")
	}

	closed := make(chan error, 2)
	for i := 0; i < 2; i++ {
		go func() { closed <- pool.Close() }()
	}

	deadline := time.Now().Add(2 * time.Second)
	for {
		if _, err := pool.Get(""); errors.Is(err, ErrClosed) {
			break
		}
		if time.Now().After(deadline) {
			t.Fatal("Close did not reject new requests")
		}
		time.Sleep(time.Millisecond)
	}

	select {
	case <-closed:
		t.Fatal("Close returned before the in-flight open finished")
	case <-time.After(25 * time.Millisecond):
	}

	close(release)
	got := <-opened
	if !errors.Is(got.err, ErrClosed) {
		t.Fatalf("in-flight open crossed the close boundary: %v", got.err)
	}

	for i := 0; i < 2; i++ {
		select {
		case err := <-closed:
			if !errors.Is(err, closeErr) {
				t.Fatalf("Close() = %v, want %v", err, closeErr)
			}
		case <-time.After(2 * time.Second):
			t.Fatal("concurrent Close did not finish")
		}
	}
	if got := closes.Load(); got != 1 {
		t.Fatalf("Close called the close callback %d times, want 1", got)
	}
	if base.key != "" {
		t.Fatal("the base resource was modified")
	}
}

type result struct {
	value *testResource
	err   error
}

func equalKeys(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
