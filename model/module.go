package model

import (
	"reflect"

	"github.com/zlsgo/app_core/service"
)

type (
	Module struct {
		service.ModuleLifeCycle
		schemas *Schemas
		stores  *Stores
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

func (m *Module) Stores() *Stores {
	return m.stores
}

func (m *Module) GetStore(name string) (*Store, bool) {
	return m.stores.Get(name)
}

func (m *Module) MustGetStore(name string) *Store {
	return m.stores.MustGet(name)
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
