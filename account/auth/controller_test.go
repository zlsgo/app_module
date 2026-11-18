package auth

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_module/account/jwt"
)

const testJWTKey = "controller-test-key"

type stubProvider struct {
	name        string
	initCalls   int
	callback    Provider
	callbackErr error
}

func (s *stubProvider) Name() string { return s.name }

func (s *stubProvider) Init(_ *znet.Engine) error {
	s.initCalls++
	return nil
}

func (s *stubProvider) Login(c *znet.Context) error {
	c.JSON(200, map[string]string{"provider": s.name})
	return nil
}

func (s *stubProvider) Callback(_ *znet.Context) (Provider, error) {
	return s.callback, s.callbackErr
}

func serve(engine *znet.Engine, method, path, token string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(method, path, bytes.NewReader(nil))
	if token != "" {
		req.Header.Set(jwt.AuthorizationKey, "Bearer "+token)
	}
	engine.ServeHTTP(w, req)
	return w
}

func TestRegControllerSkipsDisabledProvider(t *testing.T) {
	tt := zlsgo.NewTest(t)

	engine := znet.New()
	provider := &stubProvider{name: "demo"}

	jwtParse, err := regController(engine, []AuthProvider{provider}, nil, testJWTKey, 60, func(*znet.Context, Provider) (string, error) {
		return "user-1", nil
	})
	tt.NoError(err)
	tt.Equal(0, provider.initCalls)

	engine.Any("/parse", func(c *znet.Context) (string, error) {
		return jwtParse(c)
	})

	w := serve(engine, "GET", "/parse", "")
	if w.Code == 200 {
		t.Fatalf("missing token should fail, got %d: %s", w.Code, w.Body.String())
	}
	if !bytes.Contains(w.Body.Bytes(), []byte("token not found")) {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}

	// 未启用的 provider 不应注册路由
	if w = serve(engine, "GET", "/demo/login", ""); w.Code == 200 {
		t.Fatalf("disabled provider route should not exist: %s", w.Body.String())
	}
}

func TestRegControllerCallbackIssuesToken(t *testing.T) {
	tt := zlsgo.NewTest(t)

	engine := znet.New()
	provider := &stubProvider{
		name:     "demo",
		callback: Provider{Provider: "demo", ProviderID: "provider-1"},
	}

	_, err := regController(engine, []AuthProvider{provider}, []string{"demo"}, testJWTKey, 3600, func(_ *znet.Context, info Provider) (string, error) {
		if info.ProviderID != "provider-1" {
			return "", errors.New("unexpected provider id: " + info.ProviderID)
		}
		return "user-1", nil
	})
	tt.NoError(err)
	tt.Equal(1, provider.initCalls)

	w := serve(engine, "GET", "/demo/callback", "")
	tt.Equal(200, w.Code)

	token := zjson.ParseBytes(w.Body.Bytes()).Get("data.token").String()
	if token == "" {
		t.Fatalf("token missing: %s", w.Body.String())
	}
	if refresh := zjson.ParseBytes(w.Body.Bytes()).Get("data.refresh_token").String(); refresh == "" {
		t.Fatalf("refresh token missing: %s", w.Body.String())
	}

	info, err := jwt.Parse(token, testJWTKey)
	tt.NoError(err)
	tt.Equal("user-1", info.Info)

	// jwtParse 能从请求头还原用户标识
	engine.Any("/parse", func(c *znet.Context) (string, error) {
		return jwt.GetToken(c), nil
	})
	w = serve(engine, "GET", "/parse", token)
	tt.Equal(200, w.Code)
	if !bytes.Contains(w.Body.Bytes(), []byte(token)) {
		t.Fatalf("token should be parsed from header: %s", w.Body.String())
	}
}

func TestRegControllerCallbackErrors(t *testing.T) {
	tt := zlsgo.NewTest(t)

	engine := znet.New()
	emptyID := &stubProvider{name: "empty"}
	_, err := regController(engine, []AuthProvider{emptyID}, []string{"empty"}, testJWTKey, 60, func(*znet.Context, Provider) (string, error) {
		return "", nil
	})
	tt.NoError(err)

	w := serve(engine, "GET", "/empty/callback", "")
	if w.Code == 200 {
		t.Fatalf("empty user id should fail: %s", w.Body.String())
	}
	if !bytes.Contains(w.Body.Bytes(), []byte("user id not found")) {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}

	failing := &stubProvider{name: "failing", callbackErr: errors.New("provider down")}
	engine2 := znet.New()
	_, err = regController(engine2, []AuthProvider{failing}, []string{"failing"}, testJWTKey, 60, func(*znet.Context, Provider) (string, error) {
		return "user-1", nil
	})
	tt.NoError(err)

	w = serve(engine2, "GET", "/failing/callback", "")
	if w.Code == 200 {
		t.Fatalf("provider error should fail: %s", w.Body.String())
	}
	if !bytes.Contains(w.Body.Bytes(), []byte("provider down")) {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}
}

func TestWeappProviderLoginReturnsState(t *testing.T) {
	tt := zlsgo.NewTest(t)

	provider := &Weapp{AppId: "app-id", AppSecret: "app-secret"}
	tt.Equal("weapp", provider.Name())

	engine := znet.New()
	tt.NoError(provider.Init(engine))
	engine.Any("/login", provider.Login)

	w := serve(engine, "GET", "/login?state=abc", "")
	tt.Equal(200, w.Code)

	data := zjson.ParseBytes(w.Body.Bytes())
	tt.Equal("weapp", data.Get("provider").String())
	tt.Equal("abc", data.Get("state").String())
}
