package restapi

import (
	"github.com/sohaha/zlsgo/zdi"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/zdb"
)

type (
	Options struct {
		ModelsDefine []Define
	}
)

func New(o ...func(*Options)) (p *Module) {
	opt := &Options{}
	for _, f := range o {
		f(opt)
	}

	return &Module{
		ModuleLifeCycle: service.ModuleLifeCycle{
			OnDone: func(di zdi.Invoker) error {
				return di.InvokeWithErrorOnly(func(db *zdb.DB) error {

					mod := NewModels(NewSQL(db), func(o *ModelsOptions) {
						o.Prefix = "model_"
					})
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
