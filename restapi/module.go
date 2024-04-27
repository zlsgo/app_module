package restapi

import (
	"reflect"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/zlsgo/app_core/service"
)

type Module struct {
	options    Options
	controller []service.Controller
}

var (
	_ service.Module = &Module{}
	_                = reflect.TypeOf(&Module{})
)

func (m Module) Name() string {
	return "Restapi"
}

func (m Module) Tasks() []service.Task {
	return nil
}

func (m Module) Controller() []service.Controller {
	return []service.Controller{
		&controller{
			Path:    m.options.Prefix,
			options: &m.options,
		},
	}
}

func (m Module) Load(invoker zdi.Invoker) (any, error) {
	return nil, nil
}

func (m Module) Start(invoker zdi.Invoker) error {
	return nil
}

func (m Module) Done(invoker zdi.Invoker) error {
	return nil
}
