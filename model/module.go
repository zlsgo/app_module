package model

import (
	"reflect"

	"github.com/zlsgo/app_core/service"
)

type (
	Module struct {
		service.ModuleLifeCycle
		schemas *Schemas
		models  *Models
		Options Options
	}
)

var (
	_                = reflect.TypeOf(&Module{})
	_ service.Module = &Module{}
)

func (m *Module) Name() string {
	return "Model"
}

func (m *Module) String() string {
	if m != nil && m.schemas != nil {
		return m.schemas.String()
	}

	return "[]"
}

func (m *Module) Models() *Models {
	return m.models
}

func (m *Module) GetModel(name string) (*Store, bool) {
	return m.models.Get(name)
}

func (m *Module) MustGetModel(name string) *Store {
	return m.models.MustGet(name)
}

func (m *Module) Schemas() *Schemas {
	return m.schemas
}

func (m *Module) GetSchema(name string) (*Schema, bool) {
	return m.schemas.Get(name)
}

func (m *Module) MustGetSchema(name string) *Schema {
	return m.schemas.MustGet(name)
}
