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

// Reload 重新加载配置并初始化数据库
func (p *Plugin) Reload(conf *service.Conf) error {
	if conf.Get(options.ConfKey()).Value() == nil {
		return nil
	}

	var nOptions Options

	err := conf.Unmarshal(nOptions.ConfKey(), &nOptions)
	if err != nil {
		return err
	}

	db, err := initDB(nOptions)
	if err != nil {
		return zerror.With(err, "新配置初始化数据库失败")
	}

	nOptions.Mode = normalizeMode(nOptions.Mode)
	applyMode(nOptions.Mode)

	if p.db != nil {
		_ = p.db.Close()
	}

	p.db = db
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
