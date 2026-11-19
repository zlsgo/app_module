package database

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/internal/testutil"
	"github.com/zlsgo/zdb"
)

// TestSingleModuleLifecycle 覆盖单库模块的模块协议方法。
//
// 这些方法平时由 app_core 的启动流程调用，单测里没人碰，一旦签名或语义
// 变动（例如 Load 不再 setCurrentDB）不会有人发现。
func TestSingleModuleLifecycle(t *testing.T) {
	tt := zlsgo.NewTest(t)

	defer func(prev *zdb.DB) { setCurrentDB(prev) }(CurrentDB())

	db := testutil.NewMemoryDB(t)
	defer func() { _ = db.Close() }()

	mod := NewSingle(func(zdi.Invoker) (*zdb.DB, error) { return db, nil })
	tt.Equal("SingleDatabase", mod.Name())
	tt.Equal(0, len(mod.Tasks()))
	tt.Equal(0, len(mod.Controller()))

	di := zdi.New()
	di.Map(&[]service.Task{})
	di.Map(&[]service.Controller{})
	app := service.NewApp()(di)

	tt.NoError(service.InitModule([]service.Module{mod}, app))

	got, err := mod.DB()
	tt.NoError(err)
	tt.Equal(db, got)
	tt.Equal(db, CurrentDB())

	tt.NoError(mod.Start(di))
	tt.NoError(mod.Done(di))
}

// TestSingleModuleLoadFailure 确认 initDB 失败时错误会传到框架，而不是被吞掉。
func TestSingleModuleLoadFailure(t *testing.T) {
	tt := zlsgo.NewTest(t)

	expected := errors.New("连接失败")
	mod := NewSingle(func(zdi.Invoker) (*zdb.DB, error) { return nil, expected })

	di := zdi.New()
	app := service.NewApp()(di)

	err := service.InitModule([]service.Module{mod}, app)
	tt.Equal(true, err != nil)
	tt.Equal(true, errors.Is(err, expected))
}

// TestPluginLoadFromConf 用真实配置文件走一遍 Plugin.Load，
// 顺带回填 options / currentDB 这两个包级状态。
func TestPluginLoadFromConf(t *testing.T) {
	tt := zlsgo.NewTest(t)

	defer func(prev Options, prevDB *zdb.DB) {
		setOptions(prev)
		setCurrentDB(prevDB)
	}(CurrentOptions(), CurrentDB())

	path := filepath.Join(t.TempDir(), "plugin.db")

	di := zdi.New()
	conf := service.NewConf()(di)
	conf.Set("database.driver", "sqlite")
	conf.Set("database.sqlite.path", path)
	di.Map(conf)

	p := New()
	tt.Equal("Database", p.Name())
	tt.Equal(0, len(p.Tasks()))
	tt.Equal(0, len(p.Controller()))

	app := service.NewApp()(di)
	tt.NoError(service.InitModule([]service.Module{p}, app))
	defer func() { _ = p.db.Close() }()

	db, err := p.DB()
	tt.NoError(err)
	tt.Equal("sqlite", CurrentOptions().Driver)
	tt.Equal(true, db != nil)

	// 驱动实例应已注册进 DI，后续模块可以直接注入 *zdb.DB
	if _, err := db.Exec("CREATE TABLE plugin_load (id INTEGER PRIMARY KEY)"); err != nil {
		t.Fatalf("exec: %v", err)
	}

	tt.NoError(p.Start(di))
	tt.NoError(p.Done(di))
}
