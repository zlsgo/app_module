package limiter

import (
	"net/http/httptest"
	"testing"

	"github.com/sohaha/zlsgo/znet"
)

// doRequest 以指定客户端地址请求一次 /ping，返回状态码。
func doRequest(t *testing.T, r *znet.Engine, remoteAddr string) int {
	t.Helper()
	req := httptest.NewRequest("GET", "/ping", nil)
	if remoteAddr != "" {
		req.RemoteAddr = remoteAddr
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	return resp.Code
}

func TestIPMiddlewareLimitsPerIP(t *testing.T) {
	limit := MaxIPRequestPerSecond
	if limit <= 0 {
		t.Fatalf("default limit = %d, want positive", limit)
	}

	r := znet.New()
	r.Use(IPMiddleware())
	r.GET("/ping", func(c *znet.Context) error {
		c.String(200, "pong")
		return nil
	})

	for i := 0; i < limit; i++ {
		if code := doRequest(t, r, "10.1.2.3:1234"); code != 200 {
			t.Fatalf("request %d status = %d, want 200", i+1, code)
		}
	}

	if code := doRequest(t, r, "10.1.2.3:1234"); code != 429 {
		t.Fatalf("request %d status = %d, want 429", limit+1, code)
	}

	// 另一个 IP 不受影响：限流按 IP 隔离。
	if code := doRequest(t, r, "10.9.9.9:1234"); code != 200 {
		t.Fatalf("other ip status = %d, want 200", code)
	}
}

func TestNewIPMiddlewareCustomRate(t *testing.T) {
	r := znet.New()
	r.Use(NewIPMiddleware(2))
	r.GET("/ping", func(c *znet.Context) error {
		c.String(200, "pong")
		return nil
	})

	for i := 0; i < 2; i++ {
		if code := doRequest(t, r, "10.3.3.3:1234"); code != 200 {
			t.Fatalf("request %d status = %d, want 200", i+1, code)
		}
	}

	if code := doRequest(t, r, "10.3.3.3:1234"); code != 429 {
		t.Fatalf("request over custom limit status = %d, want 429", code)
	}
}

func TestIPMiddlewareInstancesAreIsolated(t *testing.T) {
	newEngine := func() *znet.Engine {
		r := znet.New()
		r.Use(NewIPMiddleware(1))
		r.GET("/ping", func(c *znet.Context) error {
			c.String(200, "pong")
			return nil
		})
		return r
	}

	r1 := newEngine()
	r2 := newEngine()

	if code := doRequest(t, r1, "10.4.4.4:1234"); code != 200 {
		t.Fatalf("first engine first request status = %d, want 200", code)
	}
	if code := doRequest(t, r1, "10.4.4.4:1234"); code != 429 {
		t.Fatalf("first engine second request status = %d, want 429", code)
	}

	// 第二个引擎使用独立限流器，不应受第一个引擎计数影响
	if code := doRequest(t, r2, "10.4.4.4:1234"); code != 200 {
		t.Fatalf("second engine first request status = %d, want 200", code)
	}
}

func TestIPMiddlewarePassesThrough(t *testing.T) {
	r := znet.New()
	r.Use(IPMiddleware())
	r.GET("/ping", func(c *znet.Context) error {
		c.String(200, "ok")
		return nil
	})

	if code := doRequest(t, r, "10.2.2.2:1234"); code != 200 {
		t.Fatalf("status = %d, want 200", code)
	}
}
