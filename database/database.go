package database

import (
	"errors"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/zdb"
)

type InitFunc func(di zdi.Invoker) (*zdb.DB, error)

var options = Options{
	disableWrite: true,
	Driver:       "sqlite",
	Sqlite: &Sqlite{
		Path:       "db.db",
		Parameters: "_pragma=busy_timeout(5000)",
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

func New(o ...Options) *Plugin {
	if len(o) > 0 {
		options = o[0]
	}

	service.DefaultConf = append(service.DefaultConf, options)
	return &Plugin{}
}

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

	p.db = db
	options = nOptions
	return nil
}

func (p *Plugin) DB() (*zdb.DB, error) {
	if p.db == nil {
		return nil, errors.New("database not init")
	}

	return p.db, nil
}
