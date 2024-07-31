package member

import (
	"errors"
	"reflect"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/account/auth"
	"github.com/zlsgo/app_module/model"
	"github.com/zlsgo/zdb"
)

type Module struct {
	service.ModuleLifeCycle
	db          *zdb.DB
	mods        *model.Schemas
	jwtParse    func(c *znet.Context) (string, error)
	instance    *Instance
	controllers []service.Controller
	Options     Options
}

var (
	_ service.Module = &Module{}
	_                = reflect.TypeOf(&Module{})
)

func (m *Module) Name() string {
	return "Member"
}

type Options struct {
	InitDB           func() (*zdb.DB, error) `z:"-"`
	ApiPrefix        string                  `z:"prefix"`
	Key              string                  `z:"key"`
	Providers        []auth.AuthProvider     `z:"-"`
	EnabledProviders []string                `z:"enabled_providers"`
	Expire           int                     `z:"expire"`
}

func (o Options) ConfKey() string {
	return "member"
}

func (o Options) DisableWrite() bool {
	return true
}

func New(key string, opt ...func(o *Options)) *Module {
	m := &Module{
		Options: Options{Key: key, ApiPrefix: "/member"},
	}

	for _, f := range opt {
		f(&m.Options)
	}

	service.DefaultConf = append(service.DefaultConf, &m.Options)

	return m
}

func (m *Module) Tasks() []service.Task {
	return []service.Task{}
}

func (m *Module) Load(di zdi.Invoker) (any, error) {
	return nil, di.InvokeWithErrorOnly(func(c *service.Conf) error {
		if m.Options.Key == "" {
			return errors.New("not set key")
		}

		m.Options.Key = zstring.Pad(m.Options.Key, 32, "0", zstring.PadRight)

		injector := di.(zdi.Injector)

		_ = initInstance(m)

		injector.Map(m.instance)

		m.controllers = []service.Controller{
			&Auth{
				Path: m.Options.ApiPrefix + "/auth",
			},
			&UserServer{
				Path: m.Options.ApiPrefix,
			},
		}
		return nil
	})
}

func (m *Module) Start(di zdi.Invoker) (err error) {
	if m.Options.InitDB != nil {
		m.db, err = m.Options.InitDB()
	} else {
		err = di.Resolve(&m.db)
	}
	if err != nil || m.db == nil {
		return zerror.With(err, "init db error")
	}

	m.mods = model.NewSchemas(di.(zdi.Injector), model.NewSQL(m.db))

	mod, err := m.mods.Reg(modelName, modelDefine(), false)
	if err != nil {
		return err
	}
	di.(zdi.Injector).Map(&Operation{
		Model: *mod.Operation(),
	})
	return
}

func (m *Module) Done(zdi.Invoker) (err error) {
	return nil
}

func (m *Module) Controller() []service.Controller {
	return m.controllers
}

func (m *Module) Stop() error {
	return nil
}
