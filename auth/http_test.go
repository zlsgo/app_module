package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"
	zsession "github.com/sohaha/zlsgo/znet/session"
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

func newHTTPTestEnv(t *testing.T, opt ...func(*Options)) *httpTestEnv {
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

	mod := New(func(o *Options) {
		o.InitDB = func() (*zdb.DB, error) { return db, nil }
		o.Session = zsession.NewMemoryStore()
	})
	for i := range opt {
		opt[i](&mod.Options)
	}

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

func (e *httpTestEnv) request(method, path string, body any, cookies ...*http.Cookie) *httptest.ResponseRecorder {
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

func sessionCookie(resp *httptest.ResponseRecorder, name string) *http.Cookie {
	for _, cookie := range resp.Result().Cookies() {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}

func responseState(resp *httptest.ResponseRecorder) string {
	body := zjson.ParseBytes(resp.Body.Bytes())
	state := body.Get("data.state").String()
	if state != "" {
		return state
	}
	return body.Get("state").String()
}

type resetCapture struct {
	job   ResetPasswordJob
	calls int
}

func (c *resetCapture) Send(job ResetPasswordJob) error {
	c.job = job
	c.calls++
	return nil
}

type oauthStubProvider struct {
	name string
}

func (s *oauthStubProvider) Name() string {
	return s.name
}

func (s *oauthStubProvider) Init(*znet.Engine) error {
	return nil
}

func (s *oauthStubProvider) Login(c *znet.Context) error {
	c.JSON(200, ztype.Map{
		"provider": s.name,
		"state":    c.Request.URL.Query().Get("state"),
	})
	return nil
}

func (s *oauthStubProvider) Callback(c *znet.Context) (Provider, error) {
	query := c.Request.URL.Query()
	return Provider{
		Provider:              s.name,
		ProviderID:            query.Get("id"),
		ProviderEmail:         query.Get("email"),
		ProviderEmailVerified: query.Get("verified") == "true",
		ProviderUsername:      query.Get("name"),
		ProviderAvatar:        query.Get("avatar"),
	}, nil
}

func TestHTTPAuthRegisterAndSignout(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newHTTPTestEnv(t)
	defer env.cleanup()

	w := env.request("POST", "/auth/user/create", map[string]string{
		"email":    "demo@example.com",
		"password": "Aa123456",
		"nickname": "demo",
	})
	tt.Equal(200, w.Code)
	cookie := sessionCookie(w, env.module.Options.CookieName)
	tt.Equal(true, cookie != nil)
	tt.Equal("demo@example.com", zjson.ParseBytes(w.Body.Bytes()).Get("data.email").String())

	w = env.request("GET", "/auth/user/get", nil, cookie)
	tt.Equal(200, w.Code)
	tt.Equal("demo@example.com", zjson.ParseBytes(w.Body.Bytes()).Get("data.email").String())

	w = env.request("GET", "/auth/user/signout", nil, cookie)
	tt.Equal(200, w.Code)

	w = env.request("GET", "/auth/user/get", nil, cookie)
	if w.Code != 401 {
		t.Fatalf("expected 401 after signout, got %d: %s", w.Code, w.Body.String())
	}
}

func TestHTTPAuthPasswordChangeInvalidatesOldSession(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newHTTPTestEnv(t)
	defer env.cleanup()

	w := env.request("POST", "/auth/user/create", map[string]string{
		"email":    "demo@example.com",
		"password": "Aa123456",
	})
	tt.Equal(200, w.Code)

	w = env.request("POST", "/auth/user/auth", map[string]string{
		"email":    "demo@example.com",
		"password": "Aa123456",
	})
	tt.Equal(200, w.Code)
	cookie1 := sessionCookie(w, env.module.Options.CookieName)
	tt.Equal(true, cookie1 != nil)

	w = env.request("POST", "/auth/user/auth", map[string]string{
		"email":    "demo@example.com",
		"password": "Aa123456",
	})
	tt.Equal(200, w.Code)
	cookie2 := sessionCookie(w, env.module.Options.CookieName)
	tt.Equal(true, cookie2 != nil)

	w = env.request("POST", "/auth/user/update", map[string]string{
		"password":         "Bb123456",
		"current_password": "Aa123456",
	}, cookie1)
	tt.Equal(200, w.Code)

	w = env.request("GET", "/auth/user/get", nil, cookie1)
	if w.Code != 200 {
		t.Fatalf("current session should remain valid, got %d: %s", w.Code, w.Body.String())
	}

	w = env.request("GET", "/auth/user/get", nil, cookie2)
	if w.Code != 401 {
		t.Fatalf("old session should be invalidated, got %d: %s", w.Code, w.Body.String())
	}

	w = env.request("POST", "/auth/user/auth", map[string]string{
		"email":    "demo@example.com",
		"password": "Bb123456",
	})
	tt.Equal(200, w.Code)
}

func TestHTTPForgotAndResetPassword(t *testing.T) {
	tt := zlsgo.NewTest(t)
	capture := &resetCapture{}
	env := newHTTPTestEnv(t, func(o *Options) {
		o.BaseURL = "https://auth.example.com"
		o.SendResetPassword = capture.Send
	})
	defer env.cleanup()

	w := env.request("POST", "/auth/user/create", map[string]string{
		"email":    "reset@example.com",
		"password": "Aa123456",
	})
	tt.Equal(200, w.Code)
	cookie := sessionCookie(w, env.module.Options.CookieName)
	tt.Equal(true, cookie != nil)

	w = env.request("POST", "/auth/user/forgotpassword", map[string]string{
		"email": "reset@example.com",
	})
	tt.Equal(200, w.Code)
	tt.Equal(1, capture.calls)
	tt.Equal("reset@example.com", capture.job.User.Email)
	tt.Equal(true, capture.job.Token != "")
	tt.Equal("https://auth.example.com/reset-password?token="+url.QueryEscape(capture.job.Token), capture.job.ResetURL)

	w = env.request("POST", "/auth/user/resetpassword", map[string]string{
		"token":    capture.job.Token,
		"password": "Bb123456",
	})
	tt.Equal(200, w.Code)

	w = env.request("GET", "/auth/user/get", nil, cookie)
	tt.Equal(401, w.Code)

	w = env.request("POST", "/auth/user/auth", map[string]string{
		"email":    "reset@example.com",
		"password": "Aa123456",
	})
	tt.Equal(401, w.Code)

	w = env.request("POST", "/auth/user/auth", map[string]string{
		"email":    "reset@example.com",
		"password": "Bb123456",
	})
	tt.Equal(200, w.Code)

	w = env.request("POST", "/auth/user/resetpassword", map[string]string{
		"token":    capture.job.Token,
		"password": "Cc123456",
	})
	tt.Equal(401, w.Code)
}

func TestHTTPOAuthCallbackAutoCreateUser(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newHTTPTestEnv(t, func(o *Options) {
		o.Providers = []AuthProvider{&oauthStubProvider{name: "demo"}}
		o.EnabledProviders = []string{"demo"}
	})
	defer env.cleanup()

	w := env.request("GET", "/auth/user/oauth/login/demo", nil)
	tt.Equal(200, w.Code)
	state := responseState(w)
	if state == "" {
		t.Fatalf("expected state in login response, got body: %s", w.Body.String())
	}
	loginCookie := sessionCookie(w, env.module.Options.CookieName)
	tt.Equal(true, loginCookie != nil)

	w = env.request("GET", "/auth/user/oauth/callback/demo?id=oauth-1&email=oauth@example.com&verified=true&name=OAuthDemo&state="+url.QueryEscape(state), nil, loginCookie)
	tt.Equal(200, w.Code)
	cookie := sessionCookie(w, env.module.Options.CookieName)
	if cookie == nil {
		cookie = loginCookie
	}
	tt.Equal(true, cookie != nil)
	tt.Equal("oauth@example.com", zjson.ParseBytes(w.Body.Bytes()).Get("data.email").String())

	w = env.request("GET", "/auth/user/get", nil, cookie)
	tt.Equal(200, w.Code)
	tt.Equal("oauth@example.com", zjson.ParseBytes(w.Body.Bytes()).Get("data.email").String())
}

func TestHTTPOAuthAddAndRemoveBinding(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newHTTPTestEnv(t, func(o *Options) {
		o.Providers = []AuthProvider{&oauthStubProvider{name: "demo"}}
		o.EnabledProviders = []string{"demo"}
	})
	defer env.cleanup()

	w := env.request("POST", "/auth/user/create", map[string]string{
		"email":    "bind@example.com",
		"password": "Aa123456",
	})
	tt.Equal(200, w.Code)
	cookie := sessionCookie(w, env.module.Options.CookieName)
	tt.Equal(true, cookie != nil)

	w = env.request("POST", "/auth/user/oauth/add", map[string]string{
		"provider": "demo",
	}, cookie)
	tt.Equal(200, w.Code)
	state := responseState(w)
	if state == "" {
		t.Fatalf("expected state in add response, got body: %s", w.Body.String())
	}

	w = env.request("GET", "/auth/user/oauth/callback/demo?id=bind-1&name=BoundDemo&state="+url.QueryEscape(state), nil, cookie)
	tt.Equal(200, w.Code)
	tt.Equal("bind-1", zjson.ParseBytes(w.Body.Bytes()).Get("data.provider.provider_id").String())

	w = env.request("POST", "/auth/user/oauth/remove", map[string]string{
		"provider":    "demo",
		"provider_id": "bind-1",
	}, cookie)
	tt.Equal(200, w.Code)

	w = env.request("GET", "/auth/user/oauth/login/demo", nil)
	tt.Equal(200, w.Code)
	state = responseState(w)
	if state == "" {
		t.Fatalf("expected state in login response, got body: %s", w.Body.String())
	}
	loginCookie := sessionCookie(w, env.module.Options.CookieName)
	tt.Equal(true, loginCookie != nil)

	w = env.request("GET", "/auth/user/oauth/callback/demo?id=bind-1&name=OtherDemo&state="+url.QueryEscape(state), nil, loginCookie)
	tt.Equal(200, w.Code)
	if zjson.ParseBytes(w.Body.Bytes()).Get("data.email").String() == "bind@example.com" {
		t.Fatalf("expected oauth callback after unlink to avoid old account, got %s", w.Body.String())
	}
}

func TestHTTPOAuthCallbackRequiresState(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newHTTPTestEnv(t, func(o *Options) {
		o.Providers = []AuthProvider{&oauthStubProvider{name: "demo"}}
		o.EnabledProviders = []string{"demo"}
	})
	defer env.cleanup()

	w := env.request("GET", "/auth/user/oauth/callback/demo?id=oauth-1&email=oauth@example.com&verified=true&name=OAuthDemo", nil)
	tt.Equal(400, w.Code)
}

func TestHTTPOAuthEmailMergeRequiresVerifiedEmail(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newHTTPTestEnv(t, func(o *Options) {
		o.Providers = []AuthProvider{&oauthStubProvider{name: "demo"}}
		o.EnabledProviders = []string{"demo"}
	})
	defer env.cleanup()

	w := env.request("POST", "/auth/user/create", map[string]string{
		"email":    "merge@example.com",
		"password": "Aa123456",
	})
	tt.Equal(200, w.Code)

	w = env.request("GET", "/auth/user/oauth/login/demo", nil)
	tt.Equal(200, w.Code)
	state := responseState(w)
	if state == "" {
		t.Fatalf("expected state in login response, got body: %s", w.Body.String())
	}
	loginCookie := sessionCookie(w, env.module.Options.CookieName)
	tt.Equal(true, loginCookie != nil)

	w = env.request("GET", "/auth/user/oauth/callback/demo?id=oauth-merge&email=merge@example.com&verified=false&name=MergeDemo&state="+url.QueryEscape(state), nil, loginCookie)
	tt.Equal(200, w.Code)
	email := zjson.ParseBytes(w.Body.Bytes()).Get("data.email").String()
	if email == "merge@example.com" {
		t.Fatalf("expected unverified provider email to avoid merging local account, got %s", w.Body.String())
	}
}

func TestHTTPOAuthOnlyUserCanSetPasswordAndThenUnlink(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newHTTPTestEnv(t, func(o *Options) {
		o.Providers = []AuthProvider{&oauthStubProvider{name: "demo"}}
		o.EnabledProviders = []string{"demo"}
	})
	defer env.cleanup()

	w := env.request("GET", "/auth/user/oauth/login/demo", nil)
	tt.Equal(200, w.Code)
	state := responseState(w)
	if state == "" {
		t.Fatalf("expected state in login response, got body: %s", w.Body.String())
	}
	loginCookie := sessionCookie(w, env.module.Options.CookieName)
	tt.Equal(true, loginCookie != nil)

	w = env.request("GET", "/auth/user/oauth/callback/demo?id=oauth-pass&name=OAuthPass&state="+url.QueryEscape(state), nil, loginCookie)
	tt.Equal(200, w.Code)
	cookie := sessionCookie(w, env.module.Options.CookieName)
	if cookie == nil {
		cookie = loginCookie
	}
	tt.Equal(true, cookie != nil)

	w = env.request("POST", "/auth/user/oauth/remove", map[string]string{
		"provider":    "demo",
		"provider_id": "oauth-pass",
	}, cookie)
	tt.Equal(400, w.Code)

	w = env.request("POST", "/auth/user/update", map[string]string{
		"password": "Cc123456",
	}, cookie)
	tt.Equal(200, w.Code)

	w = env.request("POST", "/auth/user/oauth/remove", map[string]string{
		"provider":    "demo",
		"provider_id": "oauth-pass",
	}, cookie)
	tt.Equal(200, w.Code)

	email := zjson.ParseBytes(env.request("GET", "/auth/user/get", nil, cookie).Body.Bytes()).Get("data.email").String()
	w = env.request("GET", "/auth/user/signout", nil, cookie)
	tt.Equal(200, w.Code)

	w = env.request("POST", "/auth/user/auth", map[string]string{
		"email":    email,
		"password": "Cc123456",
	})
	tt.Equal(200, w.Code)
}
