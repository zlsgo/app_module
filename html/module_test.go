package html

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/html/el"
	"github.com/zlsgo/app_module/internal/testutil"
)

func TestOptionsModuleContract(t *testing.T) {
	tt := zlsgo.NewTest(t)

	o := Options{}
	tt.Equal("html", o.ConfKey())
	// 配置只读：回写会把默认值落盘，覆盖用户配置
	tt.Equal(true, o.DisableWrite())
}

// TestModuleLifecycle 覆盖 New 里挂的 OnStart / OnDone。
//
// OnStart 把该引擎的选项存进包级 sync.Map 并注册静态路由；OnDone 负责清理。
// 这两个钩子平时只在框架启动/退出时走，缺测试时清理遗漏会变成跨用例串味。
func TestModuleLifecycle(t *testing.T) {
	tt := zlsgo.NewTest(t)

	defaultConf := service.DefaultConf
	t.Cleanup(func() { service.DefaultConf = defaultConf })

	e := znet.New(testutil.UniqueName(t, "html-module"))
	t.Cleanup(func() { moduleOptions.Delete(e) })

	di := zdi.New()
	di.Map(e)
	conf := service.NewConf()(di)
	di.Map(conf)

	mod := New(func(o *Options) {
		o.StaticPrefix = "/assets"
		o.ErrorPage = el.DIV(el.Text("module-error"))
	})
	tt.Equal(true, mod.ModuleLifeCycle.OnStart != nil)
	tt.Equal(true, mod.ModuleLifeCycle.OnDone != nil)

	tt.NoError(mod.ModuleLifeCycle.OnStart(di))

	stored, ok := moduleOptions.Load(e)
	tt.Equal(true, ok)
	opts, ok := stored.(Options)
	tt.Equal(true, ok)
	tt.Equal("/assets", opts.StaticPrefix)
	tt.Equal(true, opts.ErrorPage != nil)

	// 静态资源与 zview 注入都应已生效
	w := httptest.NewRecorder()
	e.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/assets/zview.js", nil))
	tt.Equal(http.StatusOK, w.Code)
	tt.Equal(true, len(w.Body.String()) > 0)

	tt.NoError(mod.ModuleLifeCycle.OnDone(di))
	_, ok = moduleOptions.Load(e)
	tt.Equal(false, ok)
}

// TestRendererInvokers 直接调用各个 PreInvoker 实现。
//
// 这些类型由 znet.RegisterRender 按函数签名挑选；HTTP 测试只覆盖了其中
// 几条分支，错误页回退与状态码透传需要逐个点一遍。
func TestRendererInvokers(t *testing.T) {
	tt := zlsgo.NewTest(t)

	errStub := errors.New("渲染失败")

	e := znet.New(testutil.UniqueName(t, "html-invokers"))
	t.Cleanup(func() { moduleOptions.Delete(e) })
	moduleOptions.Store(e, Options{ErrorPage: el.DIV(el.Text("fallback"))})

	e.GET("/element", func(c *znet.Context) {
		_, err := invoker(func(*znet.Context) *el.Element {
			return el.MAIN(el.Text("element"))
		}).Invoke([]any{c})
		if err != nil {
			c.String(500, err.Error())
		}
	})
	e.GET("/code", func(c *znet.Context) {
		_, err := invokerCode(func(*znet.Context) (*el.Element, error) {
			return el.DIV(el.Text("created")), nil
		}).Invoke([]any{c})
		if err != nil {
			c.String(500, err.Error())
		}
	})
	e.GET("/code-error", func(c *znet.Context) {
		_, err := invokerCode(func(*znet.Context) (*el.Element, error) {
			return nil, errStub
		}).Invoke([]any{c})
		if err != nil {
			c.String(500, err.Error())
		}
	})
	e.GET("/status", func(c *znet.Context) {
		_, err := invokerError(func(*znet.Context) (int, *el.Element) {
			return http.StatusTeapot, el.DIV(el.Text("teapot"))
		}).Invoke([]any{c})
		if err != nil {
			c.String(500, err.Error())
		}
	})

	for _, c := range []struct {
		path string
		code int
		body string
	}{
		{"/element", http.StatusOK, "<main>element</main>"},
		{"/code", http.StatusOK, "<div>created</div>"},
		// 没有 ErrorPage 的引擎返回原始错误，有则渲染 500 错误页
		{"/code-error", http.StatusInternalServerError, "<div>fallback</div>"},
		{"/status", http.StatusTeapot, "<div>teapot</div>"},
	} {
		w := httptest.NewRecorder()
		e.ServeHTTP(w, httptest.NewRequest(http.MethodGet, c.path, nil))
		tt.Equal(c.code, w.Code)
		tt.Equal(c.body, w.Body.String())
	}
}

// TestWritePageWithoutEngine 确认拿不到引擎选项时不会 panic，
// 且无 ErrorPage 时把原始错误透传出去。
func TestWritePageWithoutEngine(t *testing.T) {
	tt := zlsgo.NewTest(t)

	tt.Equal(Options{}, optionsFor(nil))
	tt.NoError(writePage(nil, http.StatusOK, nil, nil))

	boom := errors.New("boom")
	tt.Equal(true, errors.Is(writePage(nil, http.StatusOK, nil, boom), boom))
}
