package model

import (
	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

func (m *Schema) GetName() string {
	return m.define.Name
}

func (m *Schema) GetComment() string {
	if m.define.Table.Comment != "" {
		return m.define.Table.Comment
	}
	return m.GetName()
}

func (m *Schema) GetAlias() string {
	return m.alias
}

func (m *Schema) GetDefine() schema.Schema {
	return m.define
}

func (m *Schema) GetExtend() ztype.Map {
	return m.define.Extend
}

func (m *Schema) GetTableName() string {
	return m.tablePrefix + m.define.Table.Name
}

func (m *Schema) Migration() Migrationer {
	return m.Storage.Migration(m)
}

func (m *Schema) GetFields(exclude ...string) []string {
	f := m.fullFields
	if len(exclude) == 0 {
		exclude = m.GetDefine().Options.LowFields
		if len(exclude) == 0 {
			return f
		}
	}

	return zarray.Filter(f, func(_ int, v string) bool {
		return !zarray.Contains(exclude, v)
	})
}

func (m *Schema) MarshalJSON() ([]byte, error) {
	json, err := zjson.Marshal(m.GetDefine())
	return json, err
}

func (m *Schema) DI() zdi.Injector {
	return m.di
}

func (m *Schema) hook(name string) error {
	// TODO: 钩子
	// if m.model.Hook == nil {
	// 	return nil
	// }
	// return m.model.Hook(name, m)
	return nil
}

type schemaController struct {
	module     *Module
	middleware func() []znet.Handler
	Path       string
}

func (h *schemaController) Init(r *znet.Engine) error {
	if h.middleware != nil {
		r.Use(h.middleware()...)
	}
	return nil
}

func (h *schemaController) GET(c *znet.Context) (any, error) {
	schemas := ztype.Map{}

	h.module.Schemas.ForEach(func(key string, m *Schema) bool {
		schemas[key] = ztype.Map{
			"name":    m.GetName(),
			"comment": m.GetComment(),
			"fields":  m.GetFields(),
			"extend":  m.GetExtend(),
		}
		return true
	})

	return schemas, nil
}
