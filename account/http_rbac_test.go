package account

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/internal/testutil"
	"github.com/zlsgo/zdb"
)

type httpTestEnv struct {
	engine  *znet.Engine
	module  *Module
	cleanup func()
}

func newHTTPTestEnv(t *testing.T) *httpTestEnv {
	t.Helper()

	di := zdi.New()
	di.Map(&[]service.Task{})
	di.Map(&[]service.Controller{})
	app := service.NewApp()(di)
	web := service.NewWeb()(app, nil, nil)
	engine := web.Engine
	di.Map(web)

	db := testutil.NewMemoryDB(t)

	mod := New("test-key", func(o *Options) {
		o.ApiPrefix = "/api"
		o.AdminDefaultPassword = "Aa123456"
		o.InitDB = func() (*zdb.DB, error) { return db, nil }
	})

	if err := service.InitModule([]service.Module{mod}, app); err != nil {
		_ = db.Close()
		t.Fatalf("init module: %v", err)
	}

	return &httpTestEnv{
		engine: engine,
		module: mod,
		cleanup: func() {
			_ = db.Close()
		},
	}
}

func (e *httpTestEnv) request(method, path string, body []byte, token string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	reader := bytes.NewReader(body)
	if body == nil {
		reader = bytes.NewReader(nil)
	}
	req := httptest.NewRequest(method, path, reader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if token != "" {
		req.Header.Set("Authorization", "Basic "+token)
	}
	e.engine.ServeHTTP(w, req)
	return w
}

func loginToken(t *testing.T, e *httpTestEnv, account, password string) string {
	t.Helper()
	payload, _ := json.Marshal(map[string]string{
		"account":  account,
		"password": password,
	})
	w := e.request("POST", "/api/base/login", payload, "")
	if w.Code != 200 {
		t.Fatalf("login failed: %d %s", w.Code, w.Body.String())
	}
	data := zjson.ParseBytes(w.Body.Bytes()).Get("data")
	token := data.Get("token").String()
	if token == "" {
		t.Fatalf("login token empty: %s", w.Body.String())
	}
	return token
}

func TestHTTPRBACPermissionUpdate(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newHTTPTestEnv(t)
	defer env.cleanup()

	permID, err := env.module.index.permModel.Model().Insert(ztype.Map{
		"label":    "UserList",
		"alias":    "perm_user_list",
		"status":   1,
		"action":   "GET",
		"target":   "/api/user",
		"priority": 10,
	})
	tt.NoError(err)

	roleID, err := env.module.index.roleModel.Model().Insert(ztype.Map{
		"label":      "Tester",
		"alias":      "tester",
		"status":     1,
		"permission": []int{ztype.ToInt(permID)},
	})
	tt.NoError(err)

	_, err = env.module.Inside.CreateUser(ztype.Map{
		"account":  "user1",
		"password": "Aa123456",
		"role":     ztype.ToString(roleID),
	})
	tt.NoError(err)

	tt.NoError(env.module.rebuildRBAC())

	userToken := loginToken(t, env, "user1", "Aa123456")
	w := env.request("GET", "/api/user", nil, userToken)
	if w.Code != 200 {
		t.Fatalf("expected allow /api/user: %d %s", w.Code, w.Body.String())
	}

	adminToken := loginToken(t, env, "manage", "Aa123456")
	payload, _ := json.Marshal(map[string]interface{}{
		"action": "GET",
		"target": "/api/role",
	})
	w = env.request("PATCH", "/api/permission/"+ztype.ToString(permID), payload, adminToken)
	if w.Code != 200 {
		t.Fatalf("update permission failed: %d %s", w.Code, w.Body.String())
	}

	w = env.request("GET", "/api/user", nil, userToken)
	if w.Code == 200 {
		t.Fatalf("expected deny /api/user after permission update")
	}

	w = env.request("GET", "/api/role", nil, userToken)
	if w.Code != 200 {
		t.Fatalf("expected allow /api/role: %d %s", w.Code, w.Body.String())
	}
}

func TestHTTPRBACRoleUpdate(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newHTTPTestEnv(t)
	defer env.cleanup()

	permID, err := env.module.index.permModel.Model().Insert(ztype.Map{
		"label":    "UserList",
		"alias":    "perm_user_list",
		"status":   1,
		"action":   "GET",
		"target":   "/api/user",
		"priority": 10,
	})
	tt.NoError(err)

	roleID, err := env.module.index.roleModel.Model().Insert(ztype.Map{
		"label":      "Tester",
		"alias":      "tester",
		"status":     1,
		"permission": []int{ztype.ToInt(permID)},
	})
	tt.NoError(err)

	_, err = env.module.Inside.CreateUser(ztype.Map{
		"account":  "user1",
		"password": "Aa123456",
		"role":     ztype.ToString(roleID),
	})
	tt.NoError(err)

	tt.NoError(env.module.rebuildRBAC())

	userToken := loginToken(t, env, "user1", "Aa123456")
	w := env.request("GET", "/api/user", nil, userToken)
	if w.Code != 200 {
		t.Fatalf("expected allow /api/user: %d %s", w.Code, w.Body.String())
	}

	adminToken := loginToken(t, env, "manage", "Aa123456")
	payload, _ := json.Marshal(map[string]interface{}{
		"permission": []int{},
	})
	w = env.request("PATCH", "/api/role/"+ztype.ToString(roleID), payload, adminToken)
	if w.Code != 200 {
		t.Fatalf("update role failed: %d %s", w.Code, w.Body.String())
	}

	w = env.request("GET", "/api/user", nil, userToken)
	if w.Code == 200 {
		t.Fatalf("expected deny /api/user after role update")
	}
}

func TestHTTPLoginUnknownAccountCountsFailure(t *testing.T) {
	env := newHTTPTestEnv(t)
	defer env.cleanup()

	body, _ := json.Marshal(map[string]string{
		"account":  "ghost-account",
		"password": "{{SECRET_goy3qnru}}",
	})

	for i := 0; i < 5; i++ {
		w := env.request("POST", "/api/base/login", body, "")
		if w.Code == 200 {
			t.Fatalf("attempt %d should fail: %s", i+1, w.Body.String())
		}
	}

	// 账号不存在也计入失败次数，达到阈值后即使密码正确也会被拦截
	w := env.request("POST", "/api/base/login", body, "")
	if w.Code == 200 {
		t.Fatalf("login should be blocked after repeated failures: %s", w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "登录失败次数过多") {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}
}

func TestHTTPLoginLockoutBlocksCorrectPassword(t *testing.T) {
	env := newHTTPTestEnv(t)
	defer env.cleanup()

	// httptest 的默认客户端地址，中间件按它做 IP 维度限流
	const clientIP = "192.0.2.1"

	bad, _ := json.Marshal(map[string]string{
		"account":  "manage",
		"password": "wrong-password",
	})
	for i := 0; i < 5; i++ {
		if w := env.request("POST", "/api/base/login", bad, ""); w.Code == 200 {
			t.Fatalf("attempt %d should fail: %s", i+1, w.Body.String())
		}
	}

	total := func() int {
		raw, _ := env.module.loginLimit.Get(clientIP)
		return ztype.ToInt(raw)
	}
	if got := total(); got != 5 {
		t.Fatalf("failure count = %d, want 5", got)
	}

	// 已被限流时继续请求不应累加计数（否则封禁窗口会被无限延长）
	for i := 0; i < 3; i++ {
		_ = env.request("POST", "/api/base/login", bad, "")
	}
	if got := total(); got != 5 {
		t.Fatalf("blocked attempts should not increase count, got %d", got)
	}

	// 正确密码在封禁窗口内同样被拒绝
	good, _ := json.Marshal(map[string]string{
		"account":  "manage",
		"password": "{{SECRET_goy3qnru}}",
	})
	w := env.request("POST", "/api/base/login", good, "")
	if w.Code == 200 {
		t.Fatalf("locked client should not log in with correct password: %s", w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "登录失败次数过多") {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}
}
