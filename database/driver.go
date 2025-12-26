package database

import (
	"errors"
	"sync"

	"github.com/zlsgo/zdb/driver"
)

// drivers 保存已注册驱动工厂
var (
	drivers   = map[string]func(Options) (driver.IfeConfig, error){}
	driversMu sync.RWMutex
)

// Register 注册驱动工厂
func Register(name string, driver func(Options) (driver.IfeConfig, error)) error {
	if name == "" {
		return errors.New("数据库驱动名为空")
	}
	driversMu.Lock()
	defer driversMu.Unlock()
	if _, ok := drivers[name]; ok {
		return errors.New("数据库驱动[" + name + "]已注册")
	}
	drivers[name] = driver
	return nil
}

func getDriver(name string) (func(Options) (driver.IfeConfig, error), bool) {
	if name == "" {
		return nil, false
	}
	driversMu.RLock()
	defer driversMu.RUnlock()
	dri, ok := drivers[name]
	return dri, ok
}
