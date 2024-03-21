package restapi

import (
	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/zdb"
)

type (
	Options struct {
		ModelsDefine []Define
		Prefix       string
	}
)

func New(o ...func(*Options)) (m *Module) {
	opt := Options{
		Prefix: "model_",
	}
	for _, f := range o {
		f(&opt)
	}

	return &Module{
		Options: opt,
		ModuleLifeCycle: service.ModuleLifeCycle{
			OnDone: func(di zdi.Invoker) error {
				return di.InvokeWithErrorOnly(func(db *zdb.DB) error {
					mod := NewModels(di.(zdi.Injector), NewSQL(db, func(o *SQLOptions) {
						o.Prefix = m.Options.Prefix
					}))

					mapper := di.(zdi.TypeMapper)
					mops := &Operations{m: zarray.NewHashMap[string, *Operation]()}
					for _, d := range opt.ModelsDefine {
						m, err := mod.Reg(d.Name, d, false)
						if err != nil {
							return err
						}
						mops.m.Set(d.Name, m.Operation())
						mapper.Map(d)
					}

					_ = di.(zdi.TypeMapper).Map(mod)
					_ = di.(zdi.TypeMapper).Map(mops)

					m.Models = mod
					m.Operations = mops
					return nil
				})
			},
		},
	}
}
