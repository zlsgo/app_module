package database

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/builder"
	"github.com/zlsgo/zdb/driver"
)

func initDB(db Options) (*zdb.DB, error) {
	var (
		dbConf driver.IfeConfig
	)

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

	dri, ok := drivers[d]
	if !ok {
		tip := "未知数据库类型[" + d + "]"
		if d == "" {
			tip = "未配置数据库类型"
		}
		return nil, errors.New("初始化数据库失败: " + tip)
	}

	dbConf, err := dri(db)
	if err != nil {
		return nil, zerror.With(err, "数据库配置失败")
	}

	builder.DefaultDriver = dbConf.(driver.Dialect)

	e, err := zdb.New(dbConf)
	if err != nil {
		return nil, zerror.With(err, "数据库连接失败")
	}

	return e, nil
}
