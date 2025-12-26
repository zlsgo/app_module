package database

import (
	"strings"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/zlsgo/zdb/driver"
)

func copyDrivers() map[string]func(Options) (driver.IfeConfig, error) {
	driversMu.RLock()
	defer driversMu.RUnlock()
	copied := make(map[string]func(Options) (driver.IfeConfig, error), len(drivers))
	for k, v := range drivers {
		copied[k] = v
	}
	return copied
}

func TestRegister(t *testing.T) {
	tt := zlsgo.NewTest(t)

	orig := copyDrivers()
	defer func() {
		driversMu.Lock()
		drivers = orig
		driversMu.Unlock()
	}()

	rawName := "  TeSt_Register_" + t.Name() + "  "
	name := strings.ToLower(strings.TrimSpace(rawName))
	called := 0

	err := Register(rawName, func(Options) (driver.IfeConfig, error) {
		called = 1
		return nil, nil
	})
	tt.NoError(err)

	err = Register(strings.ToUpper(name), func(Options) (driver.IfeConfig, error) {
		called = 2
		return nil, nil
	})
	tt.Equal(true, err != nil)

	dri, ok := getDriver(name)
	tt.Equal(true, ok)
	_, _ = dri(Options{})
	tt.Equal(1, called)
}
