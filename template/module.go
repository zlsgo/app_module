package template

import (
	"reflect"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/jet"
)

type Module struct {
	service.ModuleLifeCycle
	engine *jet.Engine
}

var (
	_ service.Module = &Module{}
	_                = reflect.TypeOf(&Module{})
)

func New(opt ...func(*Options)) (m *Module) {
	for _, f := range opt {
		f(&options)
	}

	service.DefaultConf = append(service.DefaultConf, &options)

	return &Module{
		ModuleLifeCycle: service.ModuleLifeCycle{
			OnStart: func(di zdi.Invoker) error {
				return di.InvokeWithErrorOnly(func(r *znet.Engine, conf *service.Conf) error {
					if options.Static != "" && options.StaticDir != "" {
						r.Static(options.Static, zfile.RealPath(options.StaticDir))
					}

					dir := zfile.RealPath(options.Dir)

					j := jet.New(r, dir, func(o *jet.Options) {
						o.DelimLeft = "{{:"
						o.DelimRight = "}}"
						o.Reload = options.Reload || conf.Base.Debug
					})

					if options.Funcs != nil {
						for k := range options.Funcs {
							j.AddFunc(k, options.Funcs[k])
						}
					}

					r.SetTemplate(j)
					if err := j.Load(); err != nil {
						if !zfile.DirExist(dir) {
							return nil
						}

						return err
					}

					m.engine = j
					return nil
				})
			},
			OnDone: func(di zdi.Invoker) error {
				return di.InvokeWithErrorOnly(func(r *znet.Engine) {
				})
			},
		},
	}
}
