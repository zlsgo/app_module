package model

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/internal/testutil"
	"github.com/zlsgo/app_module/model/schema"
)

// newControllerModule 组装一个带 schema api 的模型模块，并完成 InitModule。
//
// schemaController 只在 OnDone（initModels）之后才有非空 schemas，所以这里
// 走真实的模块初始化，而不是手工塞字段——否则测的是假对象。
func newControllerModule(t *testing.T, middleware func() []znet.Handler) (*Module, *znet.Engine) {
	t.Helper()

	db := testutil.NewMemoryDB(t)
	t.Cleanup(func() { _ = db.Close() })

	mod := New(func(o *Options) {
		o.Prefix = "ctl_"
		o.SchemaApi = "/schemas"
		o.SchemaMiddleware = middleware
		o.Schemas = schema.Schemas{
			{
				Name:  "ctl_users",
				Table: schema.Table{Name: "ctl_users", Comment: "users table"},
				Fields: map[string]schema.Field{
					"name": {Type: "string", Size: 64},
				},
			},
			{
				Name:  "ctl_roles",
				Table: schema.Table{Name: "ctl_roles"},
				Fields: map[string]schema.Field{
					"title": {Type: "string", Size: 32},
				},
			},
		}
	})

	di := zdi.New()
	di.Map(db)
	app := service.NewApp()(di)
	if err := service.InitModule([]service.Module{mod}, app); err != nil {
		t.Fatalf("init module: %v", err)
	}

	return mod, znet.New(testutil.UniqueName(t, "ctl"))
}

// TestSchemaControllerRegistration 确认 SchemaApi 为空时不注册控制器：
// 空路径会在 znet 里注册出无意义的路由，属于配置错误而非功能。
func TestSchemaControllerRegistration(t *testing.T) {
	tt := zlsgo.NewTest(t)

	mod := New(func(o *Options) { o.SchemaApi = "/schemas" })
	tt.Equal(1, len(mod.Service.Controllers))

	sc, ok := mod.Service.Controllers[0].(*schemaController)
	tt.Equal(true, ok)
	tt.Equal("/schemas", sc.Path)

	plain := New()
	tt.Equal(0, len(plain.Service.Controllers))
}

// TestSchemaControllerServe 走完整 HTTP 链路：Init 挂中间件、GET 输出
// schema 清单。业务侧就是靠这个接口排查模型字段，字段名/注释错一位都
// 会误导排查。
func TestSchemaControllerServe(t *testing.T) {
	tt := zlsgo.NewTest(t)

	var hits int
	mod, e := newControllerModule(t, func() []znet.Handler {
		return []znet.Handler{func(c *znet.Context) {
			hits++
			c.Next()
		}}
	})

	sc, ok := mod.Service.Controllers[0].(*schemaController)
	tt.Equal(true, ok)

	tt.NoError(sc.Init(e))
	// Init 只注册不执行，请求到达才计数
	tt.Equal(0, hits)

	e.GET(sc.Path, func(c *znet.Context) {
		data, err := sc.GET(c)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, data)
	})

	w := httptest.NewRecorder()
	e.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/schemas", nil))

	tt.Equal(http.StatusOK, w.Code)
	tt.Equal(1, hits)

	var body map[string]struct {
		Name    string   `json:"name"`
		Comment string   `json:"comment"`
		Fields  []string `json:"fields"`
	}
	tt.NoError(json.Unmarshal(w.Body.Bytes(), &body))
	tt.Equal(2, len(body))

	users, ok := body["ctl_users"]
	tt.Equal(true, ok)
	tt.Equal("ctl_users", users.Name)
	tt.Equal("users table", users.Comment)
	// id 是模型隐式补的主键，业务自定义字段只有一个 name
	tt.Equal(2, len(users.Fields))
	tt.Equal(true, slices.Contains(users.Fields, "name"))

	// 没有 Table.Comment 时注释回落到模型名
	roles, ok := body["ctl_roles"]
	tt.Equal(true, ok)
	tt.Equal("ctl_roles", roles.Comment)
}

// TestSchemaControllerInitWithoutMiddleware 覆盖 SchemaMiddleware 未配置的
// 分支：Init 必须直接返回 nil，而不是因为调用 nil 中间件 panic。
func TestSchemaControllerInitWithoutMiddleware(t *testing.T) {
	tt := zlsgo.NewTest(t)

	mod, e := newControllerModule(t, nil)
	sc, ok := mod.Service.Controllers[0].(*schemaController)
	tt.Equal(true, ok)
	tt.NoError(sc.Init(e))

	// MarshalJSON / DI 是业务侧另外两个取用口，顺手确认它们可用
	s := mod.MustGetSchema("ctl_users")
	raw, err := json.Marshal(s)
	tt.NoError(err)
	tt.Equal(true, strings.Contains(string(raw), "ctl_users"))
	// 注入器由 initModels 传入，未接 DI 时为 nil；这里只确认调用安全
	_ = s.DI()
}
