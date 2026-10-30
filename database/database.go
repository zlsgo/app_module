package database

import (
	"errors"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/zdb"
)

// InitFunc 定义数据库初始化函数
type InitFunc func(di zdi.Invoker) (*zdb.DB, error)

// options 保存默认配置
var options = Options{
	disableWrite: true,
	Driver:       "sqlite",
	Sqlite: &Sqlite{
		Path:       "data/db.db",
		Parameters: "",
	},
	MySQL: &Mysql{
		Host:       "127.0.0.1",
		Port:       3306,
		User:       "root",
		Password:   "",
		DBName:     "zls",
		Parameters: "charset=utf8mb4&parseTime=True&loc=Local",
	},
}

// New 返回数据库插件
func New(o ...Options) *Plugin {
	if len(o) > 0 {
		options = o[0]
	}

	registerDefaultConf(options)
	return &Plugin{}
}

// Reload 重新加载配置并刷新数据库连接。
//
// 关键约束：DI 里 Map 的 p.db 指针要在整个进程生命期内保持不变。
// 所有下游模块（app_module/model 的仓储、业务 store）在各自 Load 时
// 把 *zdb.DB 存进结构体，如果这里换成新对象，它们手里的旧指针就
// 永远指向一个已 Close 的库，之后每次查询都报 `sql: database is closed`。
//
// 因此 Reload 只做“同一个句柄换内部连接池”：
//  1. 用新配置建一个临时 *zdb.DB（会 Ping，失败直接返回，旧池不动）；
//  2. 把它的池摘下来换进 p.db，原子生效；
//  3. 旧池交给 ClosePools：sql.DB.Close 会等在途查询结束，形成排空屏障。
func (p *Plugin) Reload(conf *service.Conf) error {
	if conf.Get(options.ConfKey()).Value() == nil {
		return nil
	}

	var nOptions Options

	err := conf.Unmarshal(nOptions.ConfKey(), &nOptions)
	if err != nil {
		return err
	}

	nOptions.Mode = normalizeMode(nOptions.Mode)

	if p.db == nil {
		return errors.New("database not init")
	}

	fresh, err := initDB(nOptions)
	if err != nil {
		return zerror.With(err, "新配置初始化数据库失败")
	}

	applyMode(nOptions.Mode)

	// 旧池必须先排空再关闭，且关闭不能阻塞 Reload 的调用方（配置热重载
	// 走的是 HTTP 请求），所以放在独立 goroutine 里等。
	old := p.db.SwapPools(fresh.TakePools())
	if len(old) > 0 {
		go func() { _ = zdb.ClosePools(old) }()
	}

	options = nOptions
	return nil
}

// DB 返回当前数据库实例
func (p *Plugin) DB() (*zdb.DB, error) {
	if p.db == nil {
		return nil, errors.New("database not init")
	}

	return p.db, nil
}

func registerDefaultConf(conf Options) {
	for i := range service.DefaultConf {
		if v, ok := service.DefaultConf[i].(service.DefaultConfValue); ok && v.ConfKey() == conf.ConfKey() {
			service.DefaultConf[i] = conf
			return
		}
	}
	service.DefaultConf = append(service.DefaultConf, conf)
}
