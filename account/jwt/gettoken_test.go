package jwt

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/znet"
)

func getTokenFromRequest(t *testing.T, mutate func(*znet.Context)) string {
	t.Helper()

	engine := znet.New()
	engine.Any("/token", func(c *znet.Context) string {
		if mutate != nil {
			mutate(c)
		}
		return GetToken(c)
	})

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, httptest.NewRequest("GET", "/token", nil))

	body := strings.TrimSpace(w.Body.String())
	return strings.Trim(body, `"`)
}

func TestGetTokenFromHeaderAndCookie(t *testing.T) {
	tt := zlsgo.NewTest(t)

	tt.Equal("", getTokenFromRequest(t, nil))

	bearer := getTokenFromRequest(t, func(c *znet.Context) {
		c.Request.Header.Set(AuthorizationKey, "Bearer abc.def.ghi")
	})
	tt.Equal("abc.def.ghi", bearer)

	cookie := getTokenFromRequest(t, func(c *znet.Context) {
		c.Request.AddCookie(&http.Cookie{Name: AuthorizationKey, Value: "cookie.token.value"})
	})
	tt.Equal("cookie.token.value", cookie)
}

func TestGetTokenDecodesBasicAuthorization(t *testing.T) {
	tt := zlsgo.NewTest(t)

	raw := base64.StdEncoding.EncodeToString([]byte("admin:secret"))
	token := getTokenFromRequest(t, func(c *znet.Context) {
		c.Request.Header.Set(AuthorizationKey, "Basic "+raw)
	})
	tt.Equal("admin", token)

	// 非法 base64 不应 panic，返回空串
	tt.Equal("", getTokenFromRequest(t, func(c *znet.Context) {
		c.Request.Header.Set(AuthorizationKey, "Basic @@@not-base64@@@")
	}))
}

func TestGetTokenFromQuery(t *testing.T) {
	tt := zlsgo.NewTest(t)

	engine := znet.New()
	engine.Any("/token", func(c *znet.Context) string {
		return GetToken(c)
	})

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, httptest.NewRequest("GET", "/token?Authorization=query-token", nil))
	tt.Equal("query-token", strings.Trim(strings.TrimSpace(w.Body.String()), `"`))
}
