package database

import (
	"path/filepath"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model"
)

// TestReloadKeepsHandleUsable 锁住热重载不换句柄的契约。
//
// 下游模块（app_module/model 仓储、业务 store）在各自 Load 时把 *zdb.DB
// 存进结构体，之后一直用这个指针。旧实现 Reload 会 Close 旧库并把插件字段
// 指向新对象，于是所有下游拿到的是一个已关闭的 *sql.DB，此后每次查询都报
// `sql: database is closed`。这里断言：Reload 之后同一个指针依然可查询。
func TestReloadKeepsHandleUsable(t *testing.T) {
	tt := zlsgo.NewTest(t)

	oldMode := model.DefaultSchemaOptions
	defer func() { model.DefaultSchemaOptions = oldMode }()
	oldOptions := options
	defer func() { options = oldOptions }()

	path := filepath.Join(t.TempDir(), "reload.db")

	conf := service.NewConf()(zdi.New())
	conf.Set("database.driver", "sqlite")
	conf.Set("database.sqlite.path", path)

	p := New(Options{Driver: "sqlite", Sqlite: &Sqlite{Path: path}})

	db, err := initDB(options)
	tt.NoError(err)
	p.db = db

	// 下游模块在此刻拿到的指针。
	handle := p.db

	tt.NoError(p.Reload(conf))
	tt.Equal(true, p.db == handle)

	// 旧指针必须仍然能查询（池已被原子换成新池）。
	rows, err := handle.Query("SELECT 1")
	tt.NoError(err)
	tt.NoError(rows.Close())

	tt.NoError(handle.Close())
}
