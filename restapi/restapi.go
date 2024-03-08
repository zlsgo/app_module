package restapi

import (
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
					mod := NewModels(NewSQL(db, func(o *SQLOptions) {
						o.Prefix = m.Options.Prefix
					}))

					for _, d := range opt.ModelsDefine {
						_, err := mod.Reg(d.Name, d, false)
						if err != nil {
							return err
						}
					}

					_ = di.(zdi.TypeMapper).Map(mod)

					return nil
				})
			},
		},
	}
}
