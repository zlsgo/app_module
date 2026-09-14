package html

import (
	"reflect"
	"sync"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/app_core/service"
)

type Module struct {
	service.ModuleLifeCycle
}

var moduleOptions sync.Map // map[*znet.Engine]Options

var (
	_ service.Module = &Module{}
	_                = reflect.TypeFor[*Module]()
)

func New(opt ...func(*Options)) (m *Module) {
	localOptions := zutil.Optional(Options{}, opt...)

	service.DefaultConf = append(service.DefaultConf, &localOptions)

	m = &Module{}
	m.ModuleLifeCycle = service.ModuleLifeCycle{
		OnStart: func(di zdi.Invoker) error {
			return di.InvokeWithErrorOnly(func(r *znet.Engine, conf *service.Conf) error {
				moduleOptions.Store(r, localOptions)
				return registerStatic(r, localOptions.StaticPrefix)
			})
		},
		OnDone: func(di zdi.Invoker) error {
			return di.InvokeWithErrorOnly(func(r *znet.Engine) {
				moduleOptions.Delete(r)
			})
		},
	}
	return m
}
