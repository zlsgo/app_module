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
	"github.com/zlsgo/app_module/account/jwt"
	"github.com/zlsgo/app_module/model"
	"github.com/zlsgo/zdb"
)

type Module struct {
	service.ModuleLifeCycle
	service.App
	db          *zdb.DB
	mods        *model.Models
	Options     Options
	controllers []service.Controller
	jwtParse    func(c *znet.Context) (string, error)
	Middleware  func(must bool) func(c *znet.Context) error
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
	Expire           int                     `z:"expire"`
	Providers        []auth.AuthProvider     `z:"-"`
	EnabledProviders []string                `z:"enabled_providers"`
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
		Middleware: func(must bool) func(c *znet.Context) error {
			return func(c *znet.Context) error {
				return nil
			}
		},
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
	return nil, m.DI.InvokeWithErrorOnly(func(c *service.Conf) error {
		if m.Options.Key == "" {
			return errors.New("not set key")
		}

		m.Options.Key = zstring.Pad(m.Options.Key, 32, "0", zstring.PadRight)

		m.Middleware = func(must bool) func(c *znet.Context) error {
			return func(c *znet.Context) error {
				member := &User{}
				c.Injector().Map(member)

				token := jwt.GetToken(c)
				if !must && token == "" {
					return nil
				}

				if token == "" {
					return errors.New("token not found")
				}

				info, err := jwt.Parse(token, m.Options.Key)
				if err != nil {
					return err
				}

				member.Id = info.Info
				member.Info, err = m.UserById(info.Info)
				if err != nil {
					return err
				}

				// 删除敏感信息
				_ = member.Info.Delete("password")
				_ = member.Info.Delete("salt")

				c.Next()

				return nil
			}
		}

		m.controllers = []service.Controller{
			&Auth{
				module: m,
				userModel: func() (*model.Model, bool) {
					return m.mods.Get(modelName)
				},
				Path: m.Options.ApiPrefix + "/auth",
			},
			&UserServer{
				module: m,
				Model: func() (*model.Model, bool) {
					return m.mods.Get(modelName)
				},
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
		err = m.DI.Resolve(&m.db)
	}
	if err != nil || m.db == nil {
		return zerror.With(err, "init db error")
	}

	m.mods = model.NewModels(di.(zdi.Injector), model.NewSQL(m.db))

	_, err = m.mods.Reg(modelName, modelDefine(), false)
	if err != nil {
		return err
	}
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
