package database

import (
	"github.com/zlsgo/zdb/driver"
)

var drivers = map[string]func(Options) (driver.IfeConfig, error){}
