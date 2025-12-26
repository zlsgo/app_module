package account

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/driver/sqlite3"
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
	conf := &service.Conf{Base: service.BaseConf{}}
	app := service.NewApp()(conf, di)
	web, engine := service.NewWeb()(app, nil, nil)
	di.Map(web)

	db, err := zdb.New(&sqlite3.Config{File: ":memory:", Memory: true, Parameters: "_pragma=busy_timeout(3000)"})
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}

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
