package database

import (
	"errors"
	"sync"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/zdb"
)

// InitFunc 定义数据库初始化函数
type InitFunc func(di zdi.Invoker) (*zdb.DB, error)

// options 保存默认配置，读写都经过 optionsMu。
//
// 配置热重载在 HTTP 请求里发生，分区管理器构造可能同时在另一处读这份配置，
// 所以不能直接暴露结构体：撕裂的快照会拿到上一代的 driver 配这一代的连接参数。
var (
	optionsMu sync.RWMutex
	dbMu      sync.RWMutex
	options   = Options{
		DisableAutoBringConfig: true,
		Driver:                 "sqlite",
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
	currentDB *zdb.DB
)

// New 返回数据库插件
func New(o ...Options) *Plugin {
	if len(o) > 0 {
		setOptions(o[0])
	}

	registerDefaultConf(CurrentOptions())
	return &Plugin{}
}

// CurrentOptions 返回当前生效的数据库配置。
//
// 分区管理器需要知道 driver 与连接参数才能建出同类型的副库，而这两项只有
// database 模块自己知道。注意这是浅拷贝：Sqlite / MySQL 等指针字段仍指向
// 同一份配置，调用方不应修改它们。NewPartitions 会自动复制当前配置，
// 自定义分区配置时直接使用它的可选回调即可，无需先调用 CurrentOptions。
func CurrentOptions() Options {
	optionsMu.RLock()
	defer optionsMu.RUnlock()
	return options
}

// CurrentDB 返回数据库模块当前连接，尚未初始化时返回 nil。
//
// NewPartitions 默认使用它作为默认分区的连接；多模块各自建库时，
// 最后 Load 的那个生效，需要固定连接时通过 PartitionOptions.DB 显式指定。
func CurrentDB() *zdb.DB {
	dbMu.RLock()
	defer dbMu.RUnlock()
	return currentDB
}

func setCurrentDB(db *zdb.DB) {
	dbMu.Lock()
	currentDB = db
	dbMu.Unlock()
}

// setOptions 替换当前模块配置。
func setOptions(o Options) {
	optionsMu.Lock()
	defer optionsMu.Unlock()
	options = o
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
	confKey := (Options{}).ConfKey()
	if conf.Get(confKey).Value() == nil {
		return nil
	}

	var nOptions Options

	err := conf.Unmarshal(confKey, &nOptions)
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

	setOptions(nOptions)
	setCurrentDB(p.db)
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
