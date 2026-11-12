package account

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

func TestStatusFromError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want int
	}{
		{"invalid input", zerror.InvalidInput.Text("bad"), http.StatusBadRequest},
		{"not found", zerror.NotFound.Text("missing"), http.StatusNotFound},
		{"unauthorized", zerror.Unauthorized.Text("please login"), http.StatusUnauthorized},
		{"permission denied", zerror.PermissionDenied.Text("denied"), http.StatusForbidden},
		{"cancelled", zerror.Cancelled.Text("timeout"), http.StatusRequestTimeout},
		{"internal", zerror.Internal.Text("boom"), http.StatusInternalServerError},
		{"numeric tag", zerror.TagKind("429").Text("too many"), http.StatusTooManyRequests},
		{"untagged", errors.New("plain"), http.StatusInternalServerError},
		{"out of range numeric tag", zerror.TagKind("999").Text("weird"), http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := statusFromError(tt.err); got != tt.want {
				t.Fatalf("statusFromError = %d, want %d", got, tt.want)
			}
		})
	}
}

// 身份缺失与登录失败都是客户端错误，不能一律回 500。
func TestHTTPErrorStatus(t *testing.T) {
	env := newHTTPTestEnv(t)
	defer env.cleanup()

	w := env.request("GET", "/api/user", nil, "")
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("no token status = %d, want 401: %s", w.Code, w.Body.String())
	}
	if msg := zjson.ParseBytes(w.Body.Bytes()).Get("msg").String(); msg == "" {
		t.Fatalf("empty msg: %s", w.Body.String())
	}

	w = env.request("POST", "/api/base/login", []byte(`{"account":"nobody","password":"bad"}`), "")
	if w.Code != http.StatusBadRequest {
		t.Fatalf("bad login status = %d, want 400: %s", w.Code, w.Body.String())
	}
}

// 缺失的权限在 RBAC 拒绝时回 403，而不是 500。
func TestHTTPPermissionDeniedStatus(t *testing.T) {
	env := newHTTPTestEnv(t)
	defer env.cleanup()

	if _, err := env.module.Inside.CreateUser(map[string]interface{}{
		"account":  "noperm",
		"password": "Aa123456",
	}); err != nil {
		t.Fatalf("create user: %v", err)
	}
	if err := env.module.rebuildRBAC(); err != nil {
		t.Fatalf("rebuild rbac: %v", err)
	}

	token := loginToken(t, env, "noperm", "Aa123456")
	w := env.request("GET", "/api/user", nil, token)
	if w.Code != http.StatusForbidden {
		t.Fatalf("denied status = %d, want 403: %s", w.Code, w.Body.String())
	}
}

// envRequest 在指定引擎上发一次无请求体的请求。
func envRequest(engine *znet.Engine, method, path string, body []byte) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(method, path, bytes.NewReader(body))
	engine.ServeHTTP(w, req)
	return w
}

// 应用自己通过 DI 提供错误处理器时，模块不应覆盖它。
func TestCustomErrorHandlerIsPreserved(t *testing.T) {
	di := zdi.New()
	di.Map(&[]service.Task{})
	di.Map(&[]service.Controller{})
	di.Map(&service.Conf{Base: service.BaseConf{}})
	called := false
	di.Map(znet.ErrHandlerFunc(func(c *znet.Context, err error) {
		called = true
		c.String(418, "custom")
	}))

	app := service.NewApp()(di)
	web := service.NewWeb()(app, nil, nil)
	di.Map(web)

	db, err := zdb.New(&sqlite3.Config{File: ":memory:", Memory: true, Parameters: "_pragma=busy_timeout(3000)"})
	if err != nil {
		t.Fatalf("create db: %v", err)
	}
	defer db.Close()

	mod := New("test-key", func(o *Options) {
		o.ApiPrefix = "/api"
		o.InitDB = func() (*zdb.DB, error) { return db, nil }
	})
	if err := service.InitModule([]service.Module{mod}, app); err != nil {
		t.Fatalf("init module: %v", err)
	}

	w := envRequest(web.Engine, "GET", "/api/user", nil)
	if !called {
		t.Fatal("custom error handler was not used")
	}
	if w.Code != 418 {
		t.Fatalf("status = %d, want 418", w.Code)
	}
}
