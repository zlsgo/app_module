package model

import (
	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model/define"
	"github.com/zlsgo/zdb"
)

type (
	Options struct {
		Prefix           string
		ModelsDefine     define.Defines
		DisabledMigrator bool
		GetDB            func() (*zdb.DB, error)
	}
)

func New(o ...func(*Options)) (m *Module) {
	opt := zutil.Optional(Options{Prefix: "model_", ModelsDefine: make([]define.Define, 0)}, o...)

	return &Module{
		Options: opt,
		ModuleLifeCycle: service.ModuleLifeCycle{
			OnDone: func(di zdi.Invoker) error {
				return di.InvokeWithErrorOnly(func() (err error) {
					var db *zdb.DB
					if opt.GetDB == nil {
						err = di.Resolve(&db)
					} else {
						db, err = opt.GetDB()
					}

					if err != nil {
						return zerror.With(err, "get db error")
					}

					mod := NewModels(di.(zdi.Injector), NewSQL(db, func(o *SQLOptions) {
						o.Prefix = m.Options.Prefix
					}))

					mapper := di.(zdi.TypeMapper)
					opers := &Operations{m: zarray.NewHashMap[string, *Operation]()}
					for _, d := range opt.ModelsDefine {
						if opt.DisabledMigrator {
							d.Options.DisabledMigrator = true
						}
						m, err := mod.Reg(d.Name, d, false)
						if err != nil {
							return err
						}
						opers.m.Set(d.Name, m.Operation())
						// mapper.Map(d)
					}

					_ = mapper.Maps(mod, opers)

					// m.Models = mod
					// m.Operations = mops
					return nil
				})
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
