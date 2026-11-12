package zview

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sohaha/zlsgo/znet"
)

// newRequestContext 通过真实路由拿到一个 znet.Context，避免手搓内部状态。
// setup 里的 Set* 写进 znet 的 header 缓存，处理器末尾写一次响应才会落到 Writer。
func newRequestContext(t *testing.T, setup func(view *Context)) (*Context, *httptest.ResponseRecorder) {
	t.Helper()

	r := znet.New()
	var view *Context
	r.GET("/view", func(c *znet.Context) error {
		view = New(c)
		if setup != nil {
			setup(view)
		}
		c.String(200, "ok")
		return nil
	})

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, httptest.NewRequest("GET", "/view", nil))
	if view == nil {
		t.Fatal("handler did not run")
	}
	return view, resp
}

func TestNewReadsHeaders(t *testing.T) {
	view, _ := newRequestContext(t, func(view *Context) {
		view.c.Request.Header.Set("Accept", "text/html+partial")
		view.c.Request.Header.Set("Z-Url", "/users")
		view.c.Request.Header.Set("Z-Origin", "https://example.com")
		view.c.Request.Header.Set("Z-Target", "#main")
		// New 在 setup 之前已经读过 header，这里重新构造一次验证读取结果。
		*view = *New(view.c)
	})

	if got := view.GetUrl(); got != "/users" {
		t.Fatalf("GetUrl = %q, want /users", got)
	}
	if got := view.GetOrigin(); got != "https://example.com" {
		t.Fatalf("GetOrigin = %q, want https://example.com", got)
	}
	if got := view.GetTarget(); got != "#main" {
		t.Fatalf("GetTarget = %q, want #main", got)
	}
	if !view.Is() {
		t.Fatal("Is = false, want true when Z-Url is set")
	}
	if !view.IsPartial() {
		t.Fatal("IsPartial = false, want true for text/html+partial")
	}
	if view.String() == "" {
		t.Fatal("String() should not be empty")
	}
}

func TestGetUrlFallback(t *testing.T) {
	view, _ := newRequestContext(t, nil)

	if got := view.GetUrl(); got != "" {
		t.Fatalf("GetUrl = %q, want empty", got)
	}
	if got := view.GetUrl("/fallback"); got != "/fallback" {
		t.Fatalf("GetUrl fallback = %q, want /fallback", got)
	}
	if got := view.GetUrl(); got != "" {
		t.Fatalf("GetUrl after fallback = %q, want empty", got)
	}
	if view.Is() {
		t.Fatal("Is = true, want false without Z-Url")
	}
	if view.IsPartial() {
		t.Fatal("IsPartial = true, want false without partial Accept")
	}
}

func TestSettersWriteHeaders(t *testing.T) {
	_, resp := newRequestContext(t, func(view *Context) {
		view.SetTitle("Home")
		view.SetLocation("/users/1")
		view.SetHistory("/users")
		view.SetHistoryReplace("/users/2")
		view.SetSwap(SwapMorph)
		view.SetTarget("#list")
	})

	want := map[string]string{
		"Z-Title":           "Home",
		"Z-Location":        "/users/1",
		"Z-History":         "/users",
		"Z-History-Replace": "/users/2",
		"Z-Swap":            "morph",
		"Z-Target":          "#list",
	}
	for key, value := range want {
		if got := resp.Header().Get(key); got != value {
			t.Fatalf("header %s = %q, want %q", key, got, value)
		}
	}
}

func TestSetRedirectPartialOnlySetsHeader(t *testing.T) {
	_, resp := newRequestContext(t, func(view *Context) {
		view.c.Request.Header.Set("Z-Url", "/dashboard")
		*view = *New(view.c)
		view.SetRedirect("/login")
	})

	if got := resp.Header().Get("Z-Redirect"); got != "/login" {
		t.Fatalf("Z-Redirect = %q, want /login", got)
	}
	if got := resp.Header().Get("Location"); got != "" {
		t.Fatalf("Location = %q, want empty for partial request", got)
	}
}

func TestSetRedirectFullRequestRedirects(t *testing.T) {
	_, resp := newRequestContext(t, func(view *Context) {
		view.SetRedirect("/login")
	})

	// Redirect 走的是 Writer 上的 Location，不依赖 znet 的 header 缓存。
	// CompletionLink 会补全成绝对地址，这里只校验路径部分。
	if got := resp.Header().Get("Location"); !strings.HasSuffix(got, "/login") {
		t.Fatalf("Location = %q, want path /login", got)
	}
}
