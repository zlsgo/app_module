package quick

import (
	"reflect"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/quick/crud"
	"github.com/zlsgo/app_module/quick/define"
	"github.com/zlsgo/app_module/quick/sqlstorage"
	"github.com/zlsgo/app_module/quick/utils"
	"github.com/zlsgo/zdb"
)

type (
	Module struct {
		service.ModuleLifeCycle

		models  *crud.Models
		Options Options
	}
)

var (
	_                = reflect.TypeOf(&Module{})
	_ service.Module = &Module{}
)

func (p *Module) Name() string {
	return "QuickCrud"
}

type (
	Options struct {
		Prefix           string
		ModelsDefine     []define.Define
		DisabledMigrator bool
		GetDB            func() (*zdb.DB, error)
	}
)

func New(o ...func(Options) Options) (m *Module) {
	opt := utils.Optional(Options{Prefix: "model_", ModelsDefine: make([]define.Define, 0)}, o...)

	return &Module{
		Options: opt,
		ModuleLifeCycle: service.ModuleLifeCycle{
			OnDone: func(di zdi.Invoker) error {
				return di.InvokeWithErrorOnly(func() (err error) {
					var db *zdb.DB
					if opt.GetDB != nil {
						db, err = opt.GetDB()
					} else {
						err = di.Resolve(&db)
					}

					if err != nil {
						return zerror.With(err, "db not found")
					}

					m.models = crud.NewModels(sqlstorage.NewSQL(db), func(o define.ModelsOptions) define.ModelsOptions {
						o.Prefix = opt.Prefix
						return o
					})

					for _, d := range opt.ModelsDefine {
						if opt.DisabledMigrator {
							d.Options.DisabledMigrator = true
						}
						_, err := m.models.Reg(d, false)
						if err != nil {
							return err
						}
					}

					return nil
				})
			},
		},
	}
}
