package restapi

import (
	"reflect"

	"github.com/zlsgo/app_core/service"
)

type (
	Module struct {
		service.ModuleLifeCycle
		Operations *Operations
		Models     *Models
		Options    Options
	}
)

var (
	_                = reflect.TypeOf(&Module{})
	_ service.Module = &Module{}
)

func (p *Module) Name() string {
	return "RestAPI"
}
