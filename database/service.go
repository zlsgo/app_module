package database

import (
	"errors"
	"sync"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/builder"
	"github.com/zlsgo/zdb/driver"
)

var builderDialectMu sync.Mutex

// initDB 根据配置创建数据库实例
func initDB(db Options) (*zdb.DB, error) {
	var dbConf driver.IfeConfig

	d, err := resolveDriver(db)
	if err != nil {
		return nil, err
	}
	if d == "" {
		return nil, errors.New("初始化数据库失败: 未配置数据库类型")
	}

	dri, ok := getDriver(d)
	if !ok {
		return nil, errors.New("初始化数据库失败: 未知数据库类型[" + d + "]")
	}

	dbConf, err = dri(db)
	if err != nil {
		return nil, zerror.With(err, "数据库配置失败")
	}

	dialect, ok := dbConf.(driver.Dialect)
	if !ok {
		return nil, errors.New("数据库配置失败: driver 未实现 Dialect")
	}
	if err := setBuilderDialect(dialect); err != nil {
		return nil, err
	}

	var e *zdb.DB
	if d == "sqlite" {
		if db := dbConf.DB(); db != nil {
			db.SetMaxOpenConns(1)
			db.SetMaxIdleConns(1)
		}
		e, err = zdb.New(dbConf)
	} else {
		e, err = zdb.New(dbConf)
	}

	if err != nil {
		return nil, zerror.With(err, "数据库连接失败")
	}

	return e, nil
}

func resolveDriver(db Options) (string, error) {
	d := db.Driver
	if d != "" {
		return d, nil
	}

	var candidates []string
	if db.MySQL != nil && db.MySQL.Host != "" {
		candidates = append(candidates, "mysql")
	}
	if db.Postgres != nil && db.Postgres.Host != "" {
		candidates = append(candidates, "postgres")
	}
	if db.Sqlite != nil && db.Sqlite.Path != "" {
		candidates = append(candidates, "sqlite")
	}

	if len(candidates) > 1 {
		return "", errors.New("初始化数据库失败: 多个数据库配置同时存在, 请显式指定 driver")
	}
	if len(candidates) == 1 {
		return candidates[0], nil
	}
	return "", nil
}

func setBuilderDialect(dialect driver.Dialect) error {
	if dialect == nil {
		return errors.New("数据库配置失败: driver Dialect 为空")
	}
	builderDialectMu.Lock()
	defer builderDialectMu.Unlock()
	builder.DefaultDriver = dialect
	return nil
}
