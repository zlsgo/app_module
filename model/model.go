package model

import (
	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model/define"
	"github.com/zlsgo/zdb"
)

type (
	Options struct {
		GetDB            func() (*zdb.DB, error)
		Prefix           string
		ModelsDefine     define.Defines
		DisabledMigrator bool
		SchemaDir        string
	}
)

func New(o ...func(*Options)) (m *Module) {
	opt := zutil.Optional(Options{Prefix: "model_", SchemaDir: "data/schemas", ModelsDefine: make([]define.Define, 0)}, o...)

	return &Module{
		Options: opt,
		ModuleLifeCycle: service.ModuleLifeCycle{
			OnDone: func(di zdi.Invoker) error {
				return initModels(m, di)
			},
		},
	}
}

func (m *Model) Name() string {
	return m.model.Name
}

func (m *Model) Alias() string {
	return m.alias
}

func (m *Model) Define() define.Define {
	return m.model
}

func (m *Model) TableName() string {
	return m.tablePrefix + m.model.Table.Name
}

func (m *Model) Migration() Migrationer {
	return m.Storage.Migration(m)
}

func (m *Model) GetFields(exclude ...string) []string {
	f := m.fullFields
	if len(exclude) == 0 {
		exclude = m.Define().Options.LowFields
		if len(exclude) == 0 {
			return f
		}
	}

	return zarray.Filter(f, func(_ int, v string) bool {
		return !zarray.Contains(exclude, v)
	})
}

func (m *Model) MarshalJSON() ([]byte, error) {
	json, err := zjson.Marshal(m.Define())
	return json, err
}

func (m *Model) DI() zdi.Injector {
	return m.di
}

func (m *Model) hook(name string) error {
	// TODO: 钩子
	// if m.model.Hook == nil {
	// 	return nil
	// }
	// return m.model.Hook(name, m)
	return nil
}
