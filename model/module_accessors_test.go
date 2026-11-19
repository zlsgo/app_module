package model

import (
	"io"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zlog"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/internal/testutil"
	"github.com/zlsgo/app_module/model/schema"
)

// TestModuleAccessors 覆盖 Module 上的各访问器与 Schemas 的遍历/批量注册。
//
// 这些是模块对外的取用口（业务侧靠 GetStore / MustGetSchema 拿模型），
// 少一个 nil 判断就是线上 panic，所以都过一遍。
func TestModuleAccessors(t *testing.T) {
	tt := zlsgo.NewTest(t)

	db := testutil.NewMemoryDB(t)
	defer func() { _ = db.Close() }()

	tt.Equal("[]", (&Module{}).String())

	users := schema.Schema{
		Name:  "acc_users",
		Table: schema.Table{Name: "acc_users"},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 64},
		},
	}

	mod := New(func(o *Options) {
		o.Prefix = "acc_"
		o.Schemas = schema.Schemas{users}
	})
	tt.Equal("Model", mod.Name())
	tt.Equal("[]", mod.String())

	di := zdi.New()
	di.Map(db)
	app := service.NewApp()(di)

	tt.NoError(service.InitModule([]service.Module{mod}, app))

	tt.Equal(true, mod.Schemas() != nil)
	tt.Equal(true, mod.Stores() != nil)
	tt.Equal("[acc_users]", mod.String())

	_, ok := mod.GetSchema("acc_users")
	tt.Equal(true, ok)
	tt.Equal(true, mod.MustGetSchema("acc_users") != nil)

	store, ok := mod.GetStore("acc_users")
	tt.Equal(true, ok)
	tt.Equal(store, mod.MustGetStore("acc_users"))

	// Schemas 自身的遍历与存储访问
	tt.Equal(true, mod.Schemas().Storage() != nil)
	keys := make([]string, 0, 1)
	mod.Schemas().ForEach(func(key string, _ *Schema) bool {
		keys = append(keys, key)
		return true
	})
	tt.Equal([]string{"acc_users"}, keys)

	tt.NoError(mod.Schemas().BatchReg(map[string]schema.Schema{
		"acc_roles": {
			Name:  "acc_roles",
			Table: schema.Table{Name: "acc_roles"},
			Fields: map[string]schema.Field{
				"title": {Type: "string", Size: 32},
			},
		},
	}, false))
	_, ok = mod.GetSchema("acc_roles")
	tt.Equal(true, ok)
}

// TestDefaultLoggerBranches 走一遍默认 logger 的各级别方法。
//
// 默认实现只是 zlog 的转发，但接口方法与 zlog 的方法名并不完全一致
// （Debug/Debugf 等），写错一个就是编译期通过、运行时静默不输出。
func TestDefaultLoggerBranches(t *testing.T) {
	tt := zlsgo.NewTest(t)

	original := GetLogger()
	t.Cleanup(func() { SetLogger(original) })

	l := &defaultLogger{log: zlog.New()}
	l.log.SetOut(io.Discard)

	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")
	l.Debugf("debug %d", 1)
	l.Infof("info %d", 1)
	l.Warnf("warn %d", 1)
	l.Errorf("error %d", 1)

	// 默认 logger 生效时 SetLogger(nil) 不应把它换掉
	SetLogger(l)
	tt.Equal(l, GetLogger())
	SetLogger(nil)
	tt.Equal(l, GetLogger())
}
