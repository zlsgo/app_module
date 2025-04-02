//go:build !nosqlite
// +build !nosqlite

package database

import (
	"errors"

	"github.com/sohaha/zlsgo/zfile"
	"github.com/zlsgo/zdb/driver"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

func init() {
	drivers["sqlite"] = func(db Options) (dbConf driver.IfeConfig, err error) {
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
	}
}
