package member

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
	"github.com/sohaha/zlsgo/zstring"
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

	mod := New("member-test-key", func(o *Options) {
		o.ApiPrefix = "/member"
		o.EnableRegister = true
		o.Expire = 7200
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

func memberRegister(t *testing.T, e *httpTestEnv, account, password string) *httptest.ResponseRecorder {
	t.Helper()
	payload, _ := json.Marshal(map[string]string{
		"account":  account,
		"password": password,
		"nickname": "测试会员",
	})
	return e.request("POST", "/member/register", payload, "")
}

func TestMemberRegisterAndLogin(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newHTTPTestEnv(t)
	defer env.cleanup()

	// 正常注册成功（HTTP 200，业务 code=0）
	w := memberRegister(t, env, "demo001", "pass1234")
	tt.Equal(200, w.Code)
	tt.Equal(0, zjson.ParseBytes(w.Body.Bytes()).Get("code").Int())

	// 重复注册同一账号应失败（业务错误，返回 500 并含提示）
	w = memberRegister(t, env, "demo001", "pass1234")
	tt.NoEqual(200, w.Code)
	if !strings.Contains(w.Body.String(), "账号已存在") {
		tt.Fatal("duplicate register should report account exists, got: " + w.Body.String())
	}
}

func TestMemberRegisterPasswordPolicy(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newHTTPTestEnv(t)
	defer env.cleanup()

	cases := []struct{ name, password string }{
		{"too short", "abc123"},   // 7 位，长度不足 8
		{"no number", "abcdefgh"}, // 仅字母
		{"no letter", "12345678"}, // 仅数字
	}
	for _, c := range cases {
		w := memberRegister(t, env, "demo"+zstring.Rand(6), c.password)
		tt.NoEqual(200, w.Code)
		if strings.Contains(w.Body.String(), "code\":0") {
			tt.T().Fatalf("[%s] weak password should be rejected, got success: %s", c.name, w.Body.String())
		}
		tt.Logf("[%s] rejected as expected: %s", c.name, w.Body.String())
	}

	// 合规密码应可注册成功
	w := memberRegister(t, env, "demo999", "valid1234")
	tt.Equal(200, w.Code)
	tt.Equal(0, zjson.ParseBytes(w.Body.Bytes()).Get("code").Int())
}

func TestMemberLoginAndInfo(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newHTTPTestEnv(t)
	defer env.cleanup()

	account := "demo002"
	password := "pass1234"
	w := memberRegister(t, env, account, password)
	tt.Equal(200, w.Code)
	tt.Equal(0, zjson.ParseBytes(w.Body.Bytes()).Get("code").Int())

	// 登录成功返回 token
	payload, _ := json.Marshal(map[string]string{"account": account, "password": password})
	w = env.request("POST", "/member/login", payload, "")
	tt.Equal(200, w.Code)
	tt.Logf("login resp: %s", w.Body.String())
	data := zjson.ParseBytes(w.Body.Bytes()).Get("data")
	token := data.Get("token").String()
	if token == "" {
		tt.Fatal("login token empty: " + w.Body.String())
	}

	// 错误密码登录应失败
	badPayload, _ := json.Marshal(map[string]string{"account": account, "password": "wrongpw1"})
	w = env.request("POST", "/member/login", badPayload, "")
	tt.NoEqual(200, w.Code)

	// 携带 token 获取当前用户信息
	w = env.request("GET", "/member/info", nil, token)
	tt.Equal(200, w.Code)
	tt.Logf("info resp: %s", w.Body.String())
	info := zjson.ParseBytes(w.Body.Bytes()).Get("data")
	tt.Equal(account, info.Get("info.account").String())

	// 未携带 token 访问受保护接口应 401/403
	w = env.request("GET", "/member/info", nil, "")
	tt.NoEqual(200, w.Code)
}
