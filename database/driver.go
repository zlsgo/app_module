package database

import (
	"errors"

	"github.com/zlsgo/zdb/driver"
)

var drivers = map[string]func(Options) (driver.IfeConfig, error){}

func Register(name string, driver func(Options) (driver.IfeConfig, error)) (err error) {
	if _, ok := drivers[name]; ok {
		err = errors.New("数据库驱动[" + name + "]已注册")
	}
	drivers[name] = driver
	return
}
