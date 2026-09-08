package auth

import (
	"errors"
	"reflect"
	"strings"

	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	zsession "github.com/sohaha/zlsgo/znet/session"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model"
	mSchema "github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb"
)

type Module struct {
	service.ModuleLifeCycle
	service.App
	db          *zdb.DB
	schemas     *model.Schemas
	models      *model.Stores
	controllers []service.Controller
	loginLimit  *zcache.FastCache
	forgotLimit *zcache.FastCache
	Options     Options
}

var (
	_ service.Module = (*Module)(nil)
	_                = reflect.TypeOf(&Module{})
)

type Options struct {
	InitDB                func() (*zdb.DB, error) `z:"-"`
	ApiPrefix             string                  `z:"prefix"`
	ModelPrefix           string                  `z:"model_prefix"`
	BaseURL               string                  `z:"base_url"`
	ResetPasswordPath     string                  `z:"reset_password_path"`
	CookieName            string                  `z:"cookie_name"`
	SessionTTL            int64                   `z:"session_ttl"`
	PasswordResetTokenTTL int64                   `z:"password_reset_token_ttl"`
	Session               zsession.Store          `z:"-"`
	Providers             []AuthProvider          `z:"-"`
	EnabledProviders      []string                `z:"enabled_providers"`
	SendResetPassword     ResetPasswordSender     `z:"-"`
}

type ResetPasswordSender func(job ResetPasswordJob) error

type ResetPasswordJob struct {
	User     ResetPasswordUser
	Token    string
	ResetURL string
}

type ResetPasswordUser struct {
	ID       string
	Email    string
	Nickname string
}

func (o Options) ConfKey() string {
	return "auth"
}

func (o Options) DisableWrite() bool {
	return true
}

func New(opt ...func(o *Options)) *Module {
	m := &Module{
		Options: zutil.Optional(Options{
			ApiPrefix:             "/auth",
			ModelPrefix:           "auth_",
			ResetPasswordPath:     "/reset-password",
			CookieName:            "auth_session",
			SessionTTL:            30 * 24 * 60 * 60,
			PasswordResetTokenTTL: 60 * 60,
		}, opt...),
		loginLimit:  zcache.NewFast(),
		forgotLimit: zcache.NewFast(),
	}

	service.DefaultConf = append(service.DefaultConf, &m.Options)
	return m
}

func (m *Module) Name() string {
	return "Auth"
}

func (m *Module) Tasks() []service.Task {
	return nil
}

func (m *Module) Load(di zdi.Invoker) (any, error) {
	return nil, di.InvokeWithErrorOnly(func(*service.Conf) error {
		m.Options.ApiPrefix = strings.TrimSuffix(m.Options.ApiPrefix, "/")
		if m.Options.ApiPrefix == "" {
			return errors.New("auth prefix not set")
		}
		if m.Options.ResetPasswordPath == "" {
			m.Options.ResetPasswordPath = "/reset-password"
		}
		if !strings.HasPrefix(m.Options.ResetPasswordPath, "/") {
			m.Options.ResetPasswordPath = "/" + m.Options.ResetPasswordPath
		}
		if m.Options.Session == nil {
			m.Options.Session = zsession.NewMemoryStore()
		}

		m.controllers = []service.Controller{
			&UserController{
				Path: m.Options.ApiPrefix + "/user",
			},
			&OAuthController{
				Path: m.Options.ApiPrefix + "/user/oauth",
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

	m.schemas = model.NewSchemas(di.(zdi.Injector), model.NewSQL(m.db, m.Options.ModelPrefix), model.SchemaOptions{})
	for name, define := range map[string]func() mSchema.Schema{
		userModelName:               userModelDefine,
		providerModelName:           providerModelDefine,
		passwordResetTokenModelName: passwordResetTokenModelDefine,
		sessionModelName:            sessionModelDefine,
	} {
		_, err = m.schemas.Reg(name, define(), false)
		if err != nil {
			return err
		}
	}
	m.models = m.schemas.Models()

	return nil
}

func (m *Module) Done(zdi.Invoker) error {
	return nil
}

func (m *Module) Controller() []service.Controller {
	return m.controllers
}

func (m *Module) Stop() error {
	return nil
}

func (m *Module) UserModel() (*model.Store, bool) {
	if m.models == nil {
		return nil, false
	}
	return m.models.Get(userModelName)
}

func (m *Module) ProviderModel() (*model.Store, bool) {
	if m.models == nil {
		return nil, false
	}
	return m.models.Get(providerModelName)
}

func (m *Module) PasswordResetTokenModel() (*model.Store, bool) {
	if m.models == nil {
		return nil, false
	}
	return m.models.Get(passwordResetTokenModelName)
}

func (m *Module) SessionModel() (*model.Store, bool) {
	if m.models == nil {
		return nil, false
	}
	return m.models.Get(sessionModelName)
}

func (m *Module) EnabledProviders() []AuthProvider {
	if len(m.Options.EnabledProviders) == 0 {
		return nil
	}

	enabled := make([]AuthProvider, 0, len(m.Options.EnabledProviders))
	for _, name := range m.Options.EnabledProviders {
		for i := range m.Options.Providers {
			if m.Options.Providers[i].Name() == name {
				enabled = append(enabled, m.Options.Providers[i])
				break
			}
		}
	}

	return enabled
}
