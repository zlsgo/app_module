//go:build postgres
// +build postgres

package database

import (
	"errors"

	"github.com/zlsgo/zdb/driver"
	"github.com/zlsgo/zdb/driver/postgres"
)

func init() {
	options.Postgres = &Postgres{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "root",
		Password: "",
		DBName:   "zls",
	}

	drivers["postgres"] = func(db Options) (dbConf driver.IfeConfig, err error) {
		if db.Postgres == nil {
			return nil, errors.New("初始化数据库失败: postgres 未配置")
		}
		dbConf = &postgres.Config{
			Host:     db.Postgres.Host,
			Port:     db.Postgres.Port,
			User:     db.Postgres.User,
			Password: db.Postgres.Password,
			DBName:   db.Postgres.DBName,
			SSLMode:  db.Postgres.SSLMode,
		}
		return
	}
}
