package html

import (
	"reflect"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/app_core/service"
)

type Module struct {
	service.ModuleLifeCycle
}

var (
	_ service.Module = &Module{}
	_                = reflect.TypeOf(&Module{})
)

func New(opt ...func(*Options)) (m *Module) {
	options = zutil.Optional(Options{}, opt...)

	service.DefaultConf = append(service.DefaultConf, &options)

	return &Module{
		ModuleLifeCycle: service.ModuleLifeCycle{
			OnStart: func(di zdi.Invoker) error {
				return di.InvokeWithErrorOnly(func(r *znet.Engine, conf *service.Conf) error {
					return registerStatic(r)
				})
			},
			OnDone: func(di zdi.Invoker) error {
				return di.InvokeWithErrorOnly(func(r *znet.Engine) {
				})
			},
		},
	}
}
