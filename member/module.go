package member

import (
	"errors"
	"reflect"
	"strings"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/account/auth"
	"github.com/zlsgo/app_module/model"
	mSchema "github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb"
)

type Module struct {
	service.ModuleLifeCycle
	db          *zdb.DB
	schemas     *model.Schemas
	models      *model.Models
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
	key              string                  `z:"-"`
	Providers        []auth.AuthProvider     `z:"-"`
	EnabledProviders []string                `z:"enabled_providers"`
	Expire           int                     `z:"expire"`
	ModelPrefix      string                  `z:"model_prefix"`
	EnableRegister   bool                    `z:"enable_register"`
	Only             bool                    `z:"only"`
}

func (o Options) ConfKey() string {
	return "member"
}

func (o Options) DisableWrite() bool {
	return true
}

func New(key string, opt ...func(o *Options)) *Module {
	m := &Module{
		Options: zutil.Optional(Options{key: key, ApiPrefix: "/member"}, opt...),
	}

	service.DefaultConf = append(service.DefaultConf, &m.Options)

	return m
}

func (m *Module) Tasks() []service.Task {
	return []service.Task{}
}

func (m *Module) Load(di zdi.Invoker) (any, error) {
	return nil, di.InvokeWithErrorOnly(func(c *service.Conf) error {
		if m.Options.key == "" {
			return errors.New("not set key")
		}
		m.Options.ApiPrefix = strings.TrimSuffix(m.Options.ApiPrefix, "/")

		m.Options.key = zstring.Pad(m.Options.key, 32, "0", zstring.PadRight)

		injector := di.(zdi.Injector)

		if err := initInstance(m); err != nil {
			return err
		}

		_ = injector.Map(m.instance)

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

	m.schemas = model.NewSchemas(di.(zdi.Injector), model.NewSQL(m.db, func(s *model.SQLOptions) {
		if m.Options.ModelPrefix != "" {
			s.Prefix = m.Options.ModelPrefix
		}
	}), model.SchemaOptions{})

	for modelName, modelDefine := range map[string]func() mSchema.Schema{
		modelName:         modelDefine,
		modelProviderName: modelProviderDefine,
	} {
		_, err := m.schemas.Reg(modelName, modelDefine(), false)
		if err != nil {
			return err
		}

		// di.(zdi.Injector).Map(&Model{
		// 	Model: *schema.Model(),
		// })
	}

	m.models = m.schemas.Models()
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
