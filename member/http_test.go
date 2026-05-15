package member

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	authmodule "github.com/zlsgo/app_module/auth"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

type memberHTTPTestEnv struct {
	engine     http.Handler
	authModule *authmodule.Module
	member     *Module
	cleanup    func()
}

func newMemberHTTPTestEnv(t *testing.T, authOpt func(*authmodule.Options), memberOpt func(*Options)) *memberHTTPTestEnv {
	t.Helper()
	userCache = zcache.NewFast()

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

	authMod := authmodule.New(func(o *authmodule.Options) {
		o.InitDB = func() (*zdb.DB, error) { return db, nil }
		if authOpt != nil {
			authOpt(o)
		}
	})
	memberMod := New(func(o *Options) {
		o.InitDB = func() (*zdb.DB, error) { return db, nil }
		if memberOpt != nil {
			memberOpt(o)
		}
	})

	if err := service.InitModule([]service.Module{authMod, memberMod}, app); err != nil {
		_ = db.Close()
		t.Fatalf("init module: %v", err)
	}

	return &memberHTTPTestEnv{
		engine:     engine,
		authModule: authMod,
		member:     memberMod,
		cleanup: func() {
			_ = db.Close()
		},
	}
}

func (e *memberHTTPTestEnv) request(method, path string, body any, cookies ...*http.Cookie) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	var payload []byte
	switch v := body.(type) {
	case nil:
	case []byte:
		payload = v
	default:
		payload, _ = json.Marshal(v)
	}

	req := httptest.NewRequest(method, path, bytes.NewReader(payload))
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	for _, cookie := range cookies {
		if cookie != nil {
			req.AddCookie(cookie)
		}
	}

	e.engine.ServeHTTP(w, req)
	return w
}

func memberSessionCookie(resp *httptest.ResponseRecorder, name string) *http.Cookie {
	for _, cookie := range resp.Result().Cookies() {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}

func TestMemberRequiresAuthModule(t *testing.T) {
	userCache = zcache.NewFast()

	di := zdi.New()
	di.Map(&[]service.Task{})
	di.Map(&[]service.Controller{})
	conf := &service.Conf{Base: service.BaseConf{}}
	app := service.NewApp()(conf, di)
	web, _ := service.NewWeb()(app, nil, nil)
	di.Map(web)

	db, err := zdb.New(&sqlite3.Config{File: ":memory:", Memory: true, Parameters: "_pragma=busy_timeout(3000)"})
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}
	defer func() { _ = db.Close() }()

	memberMod := New(func(o *Options) {
		o.InitDB = func() (*zdb.DB, error) { return db, nil }
	})

	err = service.InitModule([]service.Module{memberMod}, app)
	if err == nil {
		t.Fatal("expected member module to require auth module")
	}
}

func TestMemberInfoAcceptsAuthSession(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newMemberHTTPTestEnv(t, nil, nil)
	defer env.cleanup()

	w := env.request("POST", "/auth/user/create", map[string]string{
		"email":    "bridge@example.com",
		"password": "Aa123456",
		"nickname": "bridge-user",
	})
	tt.Equal(200, w.Code)
	cookie := memberSessionCookie(w, env.authModule.Options.CookieName)
	if cookie == nil {
		t.Fatalf("expected auth session cookie, got %s", w.Body.String())
	}

	w = env.request("GET", "/member/info", nil, cookie)
	tt.Equal(200, w.Code)
	data := zjson.ParseBytes(w.Body.Bytes()).Get("data")
	tt.Equal("bridge-user", data.Get("info.nickname").String())
	tt.Equal("", data.Get("info.auth_user_id").String())

	memberID := data.Get("id").String()
	if memberID == "" {
		t.Fatalf("expected member id, got %s", w.Body.String())
	}

	userModel, _ := env.member.UserModel()
	memberRow, err := userModel.FindOneByID(memberID)
	if err != nil {
		t.Fatalf("find member row: %v", err)
	}
	if strings.TrimSpace(memberRow.Get("auth_user_id").String()) == "" {
		t.Fatalf("expected auth_user_id, got %#v", memberRow)
	}
	tt.Equal("", memberRow.Get("account").String())
	tt.Equal("", memberRow.Get("password").String())
	tt.Equal("", memberRow.Get("extension.auth.user_id").String())
}

func TestMemberInfoRejectsMissingAuthSession(t *testing.T) {
	env := newMemberHTTPTestEnv(t, nil, nil)
	defer env.cleanup()

	w := env.request("GET", "/member/info", nil)
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d: %s", w.Code, w.Body.String())
	}
}

func TestMemberPatchWithAuthSession(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newMemberHTTPTestEnv(t, nil, nil)
	defer env.cleanup()

	w := env.request("POST", "/auth/user/create", map[string]string{
		"email":    "patch@example.com",
		"password": "Aa123456",
		"nickname": "before",
	})
	tt.Equal(200, w.Code)
	cookie := memberSessionCookie(w, env.authModule.Options.CookieName)
	if cookie == nil {
		t.Fatalf("expected auth session cookie, got %s", w.Body.String())
	}

	w = env.request("PATCH", "/member/me", map[string]string{
		"nickname": "after",
	}, cookie)
	tt.Equal(200, w.Code)

	w = env.request("GET", "/member/info", nil, cookie)
	tt.Equal(200, w.Code)
	tt.Equal("after", zjson.ParseBytes(w.Body.Bytes()).Get("data.info.nickname").String())
}

func TestMemberPatchIgnoresAuthBindingFields(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newMemberHTTPTestEnv(t, nil, nil)
	defer env.cleanup()

	w := env.request("POST", "/auth/user/create", map[string]string{
		"email":    "patch-bind@example.com",
		"password": "Aa123456",
		"nickname": "before",
	})
	tt.Equal(200, w.Code)
	cookie := memberSessionCookie(w, env.authModule.Options.CookieName)
	if cookie == nil {
		t.Fatalf("expected auth session cookie, got %s", w.Body.String())
	}

	w = env.request("GET", "/member/info", nil, cookie)
	tt.Equal(200, w.Code)
	memberID := zjson.ParseBytes(w.Body.Bytes()).Get("data.id").String()
	if memberID == "" {
		t.Fatalf("expected member id, got %s", w.Body.String())
	}

	w = env.request("PATCH", "/member/me", ztype.Map{
		"nickname":     "after",
		"auth_user_id": "forged-auth-user",
		"extension": ztype.Map{
			"auth": ztype.Map{
				"user_id": "forged-auth-user",
			},
			"custom": "ok",
		},
	}, cookie)
	tt.Equal(200, w.Code)

	userModel, _ := env.member.UserModel()
	memberRow, err := userModel.FindOneByID(memberID)
	if err != nil {
		t.Fatalf("find member row: %v", err)
	}
	tt.Equal("after", memberRow.Get("nickname").String())
	tt.Equal("", memberRow.Get("extension.auth.user_id").String())
	tt.Equal("ok", memberRow.Get("extension.custom").String())
}

func TestMemberDoesNotReuseLegacyBridgeRows(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newMemberHTTPTestEnv(t, nil, nil)
	defer env.cleanup()

	w := env.request("POST", "/auth/user/create", map[string]string{
		"email":    "legacy@example.com",
		"password": "Aa123456",
		"nickname": "legacy-user",
	})
	tt.Equal(200, w.Code)
	cookie := memberSessionCookie(w, env.authModule.Options.CookieName)
	if cookie == nil {
		t.Fatalf("expected auth session cookie, got %s", w.Body.String())
	}

	authResp := zjson.ParseBytes(w.Body.Bytes()).Get("data")
	authUserID := authResp.Get("id").String()
	if authUserID == "" {
		t.Fatalf("expected auth user id, got %s", w.Body.String())
	}

	userModel, _ := env.member.UserModel()
	_, err := userModel.Insert(ztype.Map{
		"account":  legacyBridgeAccount(authUserID),
		"nickname": "legacy-profile",
		"status":   1,
		"extension": ztype.Map{
			"auth": ztype.Map{
				"user_id": authUserID,
			},
		},
	})
	if err == nil {
		t.Fatal("expected legacy row without auth_user_id to be rejected")
	}
	if !strings.Contains(err.Error(), "Auth用户ID不能为空") {
		t.Fatalf("unexpected legacy insert error: %v", err)
	}

	w = env.request("GET", "/member/info", nil, cookie)
	tt.Equal(200, w.Code)
	memberID := zjson.ParseBytes(w.Body.Bytes()).Get("data.id").String()
	memberRow, err := userModel.FindOneByID(memberID)
	if err != nil {
		t.Fatalf("find member row: %v", err)
	}
	tt.Equal(authUserID, memberRow.Get("auth_user_id").String())
	if memberRow.Get("account").String() != "" {
		t.Fatalf("expected new profile row without legacy account, got %#v", memberRow)
	}
}

func TestMemberLegacyRoutesRemoved(t *testing.T) {
	env := newMemberHTTPTestEnv(t, nil, nil)
	defer env.cleanup()

	for _, path := range []string{
		"/member/register",
		"/member/login",
		"/member/auth/demo/login",
		"/member/auth/demo/callback",
	} {
		w := env.request("POST", path, map[string]string{})
		if w.Code != http.StatusNotFound {
			t.Fatalf("expected %s to be removed, got %d: %s", path, w.Code, w.Body.String())
		}
	}
}

func legacyBridgeAccount(authUserID string) string {
	return "__auth_legacy_" + strings.TrimSpace(authUserID)
}
