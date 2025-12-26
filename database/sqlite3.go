//go:build !nosqlite
// +build !nosqlite

package database

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/sohaha/zlsgo/zfile"
	"github.com/zlsgo/zdb/driver"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

// init 注册 sqlite 驱动工厂
func init() {
	if err := Register("sqlite", func(db Options) (dbConf driver.IfeConfig, err error) {
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
		dir := filepath.Dir(db.Sqlite.Path)
		if dir != "." && dir != "" {
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return nil, err
			}
		}
		if !zfile.FileExist(db.Sqlite.Path) {
			err := zfile.WriteFile(db.Sqlite.Path, []byte(""))
			if err != nil {
				return nil, err
			}
		}
		return
	}); err != nil {
		panic(err)
	}
}
