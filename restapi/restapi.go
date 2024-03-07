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
					mods, err := NewModels(opt.ModelsDefine, db, true)
					if err != nil {
						return err
					}
					_ = mods
					// di.(zdi.TypeMapper).Map(mods)
					return nil
				})
			},
		},
	}
}
