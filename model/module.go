package model

import (
	"reflect"

	"github.com/zlsgo/app_core/service"
)

type (
	Module struct {
		service.ModuleLifeCycle
		Options Options

		Models     *Models
		Operations *Operations
	}
)

var (
	_                = reflect.TypeOf(&Module{})
	_ service.Module = &Module{}
)

func (p *Module) Name() string {
	return "Model"
}

func (p *Module) String() string {
	if p != nil && p.Models != nil {
		return p.Models.String()
	}
	return "[]"
}
