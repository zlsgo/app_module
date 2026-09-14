package html

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_module/html/el"
)

func TestHTMLRenderersAndErrorPage(t *testing.T) {
	e := znet.New("html-render-test")
	defer moduleOptions.Delete(e)
	moduleOptions.Store(e, Options{ErrorPage: el.DIV(el.ID("error"), el.Text("fallback"))})

	e.GET("/page", func(c *znet.Context) *el.Element {
		return el.DIV(el.Text("ok"))
	})
	e.GET("/status", func(c *znet.Context) (int, *el.Element) {
		return http.StatusCreated, el.DIV(el.Text("created"))
	})
	e.GET("/error", func(c *znet.Context) (*el.Element, error) {
		return nil, errors.New("expected")
	})

	tests := []struct {
		path string
		code int
		body string
	}{
		{path: "/page", code: http.StatusOK, body: "<div>ok</div>"},
		{path: "/status", code: http.StatusCreated, body: "<div>created</div>"},
		{path: "/error", code: http.StatusInternalServerError, body: `<div id="error">fallback</div>`},
	}
	for _, test := range tests {
		w := httptest.NewRecorder()
		e.ServeHTTP(w, httptest.NewRequest(http.MethodGet, test.path, nil))
		if w.Code != test.code || w.Body.String() != test.body {
			t.Errorf("GET %s = (%d, %q), want (%d, %q)", test.path, w.Code, w.Body.String(), test.code, test.body)
		}
		if got := w.Header().Get("Content-Type"); got == "" {
			t.Errorf("GET %s did not set Content-Type", test.path)
		}
	}
}

func TestHTMLRenderersKeepElementsReusable(t *testing.T) {
	e := znet.New("html-render-reuse-test")
	defer moduleOptions.Delete(e)
	e.GET("/reuse", func(c *znet.Context) *el.Element {
		return reusablePage
	})

	for i := range 2 {
		w := httptest.NewRecorder()
		e.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/reuse", nil))
		if w.Code != http.StatusOK || w.Body.String() != `<main id="reusable">content</main>` {
			t.Fatalf("request %d = (%d, %q)", i, w.Code, w.Body.String())
		}
	}
}

func TestHTMLRenderersUsePerEngineOptions(t *testing.T) {
	first := znet.New("html-options-first")
	second := znet.New("html-options-second")
	defer moduleOptions.Delete(first)
	defer moduleOptions.Delete(second)
	moduleOptions.Store(first, Options{ErrorPage: el.DIV(el.Text("first"))})
	moduleOptions.Store(second, Options{ErrorPage: el.DIV(el.Text("second"))})
	handler := func(c *znet.Context) (*el.Element, error) { return nil, errors.New("failed") }
	first.GET("/error", handler)
	second.GET("/error", handler)

	for _, test := range []struct {
		engine *znet.Engine
		want   string
	}{
		{engine: first, want: "<div>first</div>"},
		{engine: second, want: "<div>second</div>"},
	} {
		w := httptest.NewRecorder()
		test.engine.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/error", nil))
		if w.Code != http.StatusInternalServerError || w.Body.String() != test.want {
			t.Errorf("engine response = (%d, %q), want 500/%q", w.Code, w.Body.String(), test.want)
		}
	}
}

var reusablePage = el.MAIN(el.ID("reusable"), el.Text("content"))
