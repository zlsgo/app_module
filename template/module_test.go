package template

import (
	"errors"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
)

// writeViews 在临时目录下写入模板文件，返回目录路径。
func writeViews(t *testing.T, files map[string]string) string {
	t.Helper()
	dir := t.TempDir()
	for name, content := range files {
		path := filepath.Join(dir, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", name, err)
		}
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	return dir
}

// newTestModule 搭一个最小的 app/web 环境并启动模板模块。
func newTestModule(t *testing.T, opt func(*Options)) (*Module, *znet.Engine) {
	t.Helper()
	di := zdi.New()
	di.Map(&[]service.Task{})
	di.Map(&[]service.Controller{})
	di.Map(&service.Conf{Base: service.BaseConf{}})
	app := service.NewApp()(di)
	web := service.NewWeb()(app, nil, nil)
	di.Map(web)
	di.Map(web.Engine)

	m := New(opt)
	if err := service.InitModule([]service.Module{m}, app); err != nil {
		t.Fatalf("init module: %v", err)
	}
	return m, web.Engine
}

func TestNewKeepsDefaultDir(t *testing.T) {
	// 回归：New 曾从零值 Options 起算，把默认 Dir 覆盖成空串。
	got := defaultOptions()
	if got.Dir != "./views" {
		t.Fatalf("default Dir = %q, want ./views", got.Dir)
	}

	m := New()
	if m == nil {
		t.Fatal("New returned nil module")
	}
}

func TestRender(t *testing.T) {
	dir := writeViews(t, map[string]string{
		"greeting.html": `hello {{: name }}`,
	})
	m, _ := newTestModule(t, func(o *Options) {
		o.Dir = dir
	})

	got, err := m.Render("greeting", ztype.Map{"name": "zlsgo"})
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if string(got) != "hello zlsgo" {
		t.Fatalf("rendered = %q, want %q", got, "hello zlsgo")
	}
}

func TestRenderLayout(t *testing.T) {
	dir := writeViews(t, map[string]string{
		"page.html":         `<h1>{{: title }}</h1>`,
		"layouts/main.html": `<html>{{: slot() }}</html>`,
	})
	m, _ := newTestModule(t, func(o *Options) {
		o.Dir = dir
	})

	got, err := m.Render("page", ztype.Map{"title": "T"}, "layouts/main")
	if err != nil {
		t.Fatalf("render with layout: %v", err)
	}
	want := `<html><h1>T</h1></html>`
	if string(got) != want {
		t.Fatalf("rendered = %q, want %q", got, want)
	}
}

// TestRenderResultSurvivesPoolReuse 覆盖池化缓冲复用：
// Render 返回的字节不能被后续 GetBuff 写入覆盖。
func TestRenderResultSurvivesPoolReuse(t *testing.T) {
	// 两次渲染输出等长，落在同一个容量档的池缓冲上：
	// 没有脱离池化数组时，第二次渲染会覆盖第一次的结果。
	dir := writeViews(t, map[string]string{
		"a.html": `first-result`,
		"b.html": `other-result`,
	})
	m, _ := newTestModule(t, func(o *Options) {
		o.Dir = dir
	})

	first, err := m.Render("a", nil)
	if err != nil {
		t.Fatalf("render a: %v", err)
	}
	if _, err := m.Render("b", nil); err != nil {
		t.Fatalf("render b: %v", err)
	}

	if string(first) != "first-result" {
		t.Fatalf("first result was mutated to %q", string(first))
	}
}

func TestRenderMissingTemplate(t *testing.T) {
	dir := writeViews(t, map[string]string{
		"ok.html": `ok`,
	})
	m, _ := newTestModule(t, func(o *Options) {
		o.Dir = dir
	})

	if _, err := m.Render("does-not-exist", nil); err == nil {
		t.Fatal("expected error for missing template")
	}
}

// TestRenderWithoutViewsDir 目录不存在时启动不报错，渲染返回错误而不是 panic。
func TestRenderWithoutViewsDir(t *testing.T) {
	m, _ := newTestModule(t, func(o *Options) {
		o.Dir = filepath.Join(t.TempDir(), "missing")
	})

	if _, err := m.Render("anything", nil); err == nil {
		t.Fatal("expected error when views dir is missing")
	}
}

func TestRenderBeforeStart(t *testing.T) {
	m := New()
	got, err := m.Render("index", nil)
	if !errors.Is(err, ErrNotReady) {
		t.Fatalf("err = %v, want ErrNotReady", err)
	}
	if got != nil {
		t.Fatalf("result = %q, want nil", got)
	}

	var nilModule *Module
	if _, err := nilModule.Render("index", nil); !errors.Is(err, ErrNotReady) {
		t.Fatalf("nil module err = %v, want ErrNotReady", err)
	}
}

func TestTemplateWritesHTML(t *testing.T) {
	dir := writeViews(t, map[string]string{
		"index.html": `hi {{: who }}`,
	})
	m, engine := newTestModule(t, func(o *Options) {
		o.Dir = dir
	})

	engine.GET("/page", func(c *znet.Context) error {
		return m.Template(c, "index", ztype.Map{"who": "world"})
	})

	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/page", nil)
	engine.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Fatalf("status = %d, want 200", resp.Code)
	}
	if got := resp.Body.String(); got != "hi world" {
		t.Fatalf("body = %q, want %q", got, "hi world")
	}
}
