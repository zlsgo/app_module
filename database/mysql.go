//go:build !nomysql
// +build !nomysql

package database

import (
	"errors"

	"github.com/zlsgo/zdb/driver"
	"github.com/zlsgo/zdb/driver/mysql"
)

func init() {
	drivers["mysql"] = func(db Options) (dbConf driver.IfeConfig, err error) {
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
	}
}
