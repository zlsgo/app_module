package database

import (
	"errors"
	"strings"

	"github.com/zlsgo/app_core/common"
	"github.com/zlsgo/app_core/service"

	"github.com/sohaha/zlsgo/zfile"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/builder"
	"github.com/zlsgo/zdb/driver"
	"github.com/zlsgo/zdb/driver/mysql"
	"github.com/zlsgo/zdb/driver/postgres"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

func initDB(c *service.Conf) (*zdb.DB, error) {
	var (
		dbConf driver.IfeConfig
		db     Options
	)
	_ = c.Unmarshal((Options{}).ConfKey(), &db)

	d := strings.ToLower(db.Driver)
	if d == "" {
		if db.MySQL != nil && db.MySQL.Host != "" {
			d = "mysql"
		} else if db.Postgres != nil && db.Postgres.Host != "" {
			d = "postgres"
		} else if db.Sqlite != nil && db.Sqlite.Path != "" {
			d = "sqlite"
		}
	}

	switch d {
	case "mysql":
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
	case "postgres":
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
	case "sqlite":
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
	}

	if dbConf == nil {
		tip := "未知数据库类型[" + d + "]"
		if d == "" {
			tip = "未配置数据库类型"
		}
		common.Fatal(errors.New("初始化数据库失败: " + tip))
	}

	builder.DefaultDriver = dbConf.(driver.Dialect)

	return zdb.New(dbConf)
}
