package quick

import (
	"github.com/zlsgo/app_module/quick/crud"
)

func (m *Module) Models() *crud.Models {
	return m.models
}

func (m *Module) GetModel(name string) (*crud.Crud, bool) {
	return m.models.Get(name)
}
