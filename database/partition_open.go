package database

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

// sqlitePartitionOpener 返回按文件隔离的打开与关闭回调。
//
// SQLite 每个键一个副库文件，连接由本管理器持有，Close 时统一关闭。
func sqlitePartitionOpener(conf Options, base Storage) (func(string) (Storage, error), func(Storage) error) {
	open := func(key string) (Storage, error) {
		if conf.Sqlite == nil {
			return Storage{}, errors.New("分区存储不可用: sqlite 未配置")
		}

		path, err := partitionSQLitePath(conf.Sqlite.Path, Key(key))
		if err != nil {
			return Storage{}, err
		}

		opts := conf
		sqlite := *conf.Sqlite
		sqlite.Path = path
		opts.Sqlite = &sqlite

		db, err := initDB(opts)
		if err != nil {
			return Storage{}, fmt.Errorf("打开分区[%s]失败: %w", key, err)
		}

		// 文件已经隔离，表名不需要再带键，调用方无需改写 schema 的表名。
		return Storage{DB: db, Prefix: base.Prefix}, nil
	}
	close := func(storage Storage) error {
		return storage.DB.Close()
	}
	return open, close
}

// sharedPartitionOpener 共用同一个数据库实例，只把表名空间分开。
func sharedPartitionOpener(base Storage) func(string) (Storage, error) {
	return func(key string) (Storage, error) {
		return Storage{DB: base.DB, Prefix: partitionPrefix(base.Prefix, Key(key))}, nil
	}
}

// partitionSQLitePath 把默认库路径改写成该键专用的文件。
//
// data/db.db + tenant-west -> data/db_tenant-west.db
//
// 本适配器按文件隔离，不接受内存库（":memory:" 一类）。
func partitionSQLitePath(base string, key Key) (string, error) {
	trimmed := strings.TrimSpace(base)
	if trimmed == "" {
		return "", errors.New("分区存储不可用: sqlite path 未配置")
	}
	// 只对校验用 trim 后的串，派生仍用原串：默认库按原串打开，
	// 这里再裁剪一次会让分区文件落到与默认库不同的目录。
	// ":memory:" 与 "file::memory:?cache=shared" 都是内存库，后者没有前导冒号。
	if strings.HasPrefix(trimmed, ":") || strings.Contains(trimmed, ":memory:") {
		return "", errors.New("分区存储不可用: 内存库不支持文件分区")
	}

	ext := filepath.Ext(base)
	stem := strings.TrimSuffix(base, ext)
	return stem + "_" + string(key) + ext, nil
}

// partitionPrefix 把默认表前缀改写成该键专用的表名空间。
//
// data_ + tenant-west -> data_tenant-west_。
func partitionPrefix(base string, key Key) string {
	return base + string(key) + "_"
}
