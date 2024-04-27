package database

import (
	"errors"

	"github.com/sohaha/zlsgo/zfile"
	"github.com/zlsgo/zdb/driver"
	"github.com/zlsgo/zdb/driver/mysql"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

var (
	drivers = map[string]func(Options) (driver.IfeConfig, error){
		"mysql": func(db Options) (dbConf driver.IfeConfig, err error) {
			if db.MySQL == nil {
				return nil, errors.New("初始化数据库失败: mysql 未配置")
			}
			dbConf = &mysql.Config{
				Host:       db.MySQL.Host,
				Port:       db.MySQL.Port,
				User:       db.MySQL.User,
				Password:   db.MySQL.Password,
				DBName:     db.MySQL.DBName,
				Parameters: db.MySQL.Parameters,
				Charset:    db.MySQL.Charset,
				// Zone:       db.MySQL.Zone,
			}
			return
		},
		"sqlite": func(db Options) (dbConf driver.IfeConfig, err error) {
			if db.Sqlite == nil {
				return nil, errors.New("初始化数据库失败: sqlite 未配置")
			}

			if db.Sqlite.Path == "" {
				return nil, errors.New("初始化数据库失败: sqlite path 未配置")
			}

			dbConf = &sqlite3.Config{
				File:       db.Sqlite.Path,
				Parameters: db.Sqlite.Parameters,
			}
			if !zfile.FileExist(db.Sqlite.Path) {
				err := zfile.WriteFile(db.Sqlite.Path, []byte(""))
				if err != nil {
					return nil, err
				}
			}
			return
		},
	}
)
