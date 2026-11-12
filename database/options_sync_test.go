package database

import (
	"path/filepath"
	"runtime"
	"sync"
	"testing"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model"
)

// TestCurrentOptionsSnapshotWhileReloading 锁住「读配置」与「热重载」不互相踩踏。
//
// 分区管理器在构造时读一次当前配置；同一进程里配置热重载会替换它。若读侧不加锁，
// 并发时既可能读到撕裂的驱动与连接参数组合，也会被 -race 判为 data race。
func TestCurrentOptionsSnapshotWhileReloading(t *testing.T) {
	original := CurrentOptions()
	t.Cleanup(func() { setOptions(original) })
	// Reload 会经 applyMode 改进程级的 model 模式选项，跑完要还原，
	// 否则同包后续测试会继承这次重载的删除列策略。
	originalSchemaOptions := model.DefaultSchemaOptions
	t.Cleanup(func() { model.DefaultSchemaOptions = originalSchemaOptions })

	path := filepath.Join(t.TempDir(), "reload.db")
	conf := service.NewConf()(zdi.New())
	conf.Set("database.driver", "sqlite")
	conf.Set("database.sqlite.path", path)

	p := &Plugin{}
	db, err := initDB(Options{Driver: "sqlite", Sqlite: &Sqlite{Path: path}})
	if err != nil {
		t.Fatal(err)
	}
	p.db = db
	t.Cleanup(func() { _ = db.Close() })

	stop := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stop:
				return
			default:
			}
			snapshot := CurrentOptions()
			if snapshot.Driver != "" && snapshot.Sqlite == nil {
				t.Errorf("torn snapshot: driver %q without sqlite config", snapshot.Driver)
				return
			}
		}
	}()

	// 次数够触发读写交错即可：每次 Reload 都会新建并排空一个连接池，
	// 循环太多会把整仓并行测试的机器负载抬高，反而让无关的偶发问题更容易出现。
	for i := 0; i < 5; i++ {
		if err := p.Reload(conf); err != nil {
			t.Fatalf("reload %d: %v", i, err)
		}
		runtime.Gosched()
	}
	close(stop)
	wg.Wait()
}
