package member

import (
	"errors"
	"reflect"
	"strings"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/app_core/service"
	authmodule "github.com/zlsgo/app_module/auth"
	"github.com/zlsgo/app_module/model"
	mSchema "github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb"
)

type Module struct {
	service.ModuleLifeCycle
	db          *zdb.DB
	schemas     *model.Schemas
	models      *model.Stores
	instance    *Instance
	authModule  *authmodule.Module
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
	InitDB      func() (*zdb.DB, error) `z:"-"`
	ApiPrefix   string                  `z:"prefix"`
	ModelPrefix string                  `z:"model_prefix"`
}

func (o Options) ConfKey() string {
	return "member"
}

func (o Options) DisableWrite() bool {
	return true
}

func New(opt ...func(o *Options)) *Module {
	m := &Module{
		Options: zutil.Optional(Options{ApiPrefix: "/member"}, opt...),
	}

	service.DefaultConf = append(service.DefaultConf, &m.Options)

	return m
}

func (m *Module) Tasks() []service.Task {
	return []service.Task{}
}

func (m *Module) Load(di zdi.Invoker) (any, error) {
	return nil, di.InvokeWithErrorOnly(func(c *service.Conf) error {
		m.Options.ApiPrefix = strings.TrimSuffix(m.Options.ApiPrefix, "/")

		injector := di.(zdi.Injector)

		if err := initInstance(m); err != nil {
			return err
		}
		if err := di.Resolve(&m.authModule); err != nil {
			return errors.New("member module requires auth module")
		}

		_ = injector.Map(m.instance)

		m.controllers = []service.Controller{
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

	m.schemas = model.NewSchemas(di.(zdi.Injector), model.NewSQL(m.db, m.Options.ModelPrefix, func(s *model.SQLOptions) {
	}), model.SchemaOptions{})

	for modelName, modelDefine := range map[string]func() mSchema.Schema{
		modelName: modelDefine,
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
