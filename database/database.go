package database

import (
	"errors"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/zdb"
)

type InitFunc func(di zdi.Invoker) (*zdb.DB, error)

var options Options

func New(o ...Options) *Plugin {
	if len(o) > 0 {
		options = o[0]
	} else {
		options = Options{
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
			Postgres: &Postgres{
				Host:     "127.0.0.1",
				Port:     5432,
				User:     "root",
				Password: "",
				DBName:   "zls",
			},
		}
	}

	service.DefaultConf = append(service.DefaultConf, options)
	return &Plugin{}
}

func (p *Plugin) DB() (*zdb.DB, error) {
	if p.db == nil {
		return nil, errors.New("database not init")
	}

	return p.db, nil
}
