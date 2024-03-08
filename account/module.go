package account

import (
	"errors"
	"reflect"
	"time"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/account/rbac"
	"github.com/zlsgo/app_module/restapi"
	"github.com/zlsgo/zdb"
)

type Module struct {
	service.App
	service.ModuleLifeCycle
	db          *zdb.DB
	mods        *restapi.Models
	Options     Options
	controllers []service.Controller
}

var (
	_ service.Module = &Module{}
	_                = reflect.TypeOf(&Module{})
)

func (m *Module) Name() string {
	return "Account"
}

type Options struct {
	InitDB               func() (*zdb.DB, error) `json:"-"`
	key                  string
	InlayRBAC            *rbac.RBAC       `json:"-"`
	RBACFile             string           `json:"rbac_file"`
	ApiPrefix            string           `json:"prefix"`
	InlayUser            ztype.Maps       `json:"inlay_user"`
	AdminDefaultPassword string           `json:"admin_default_password"`
	Expire               int              `json:"expire"`
	Only                 bool             `json:"only"`
	DisabledLogIP        bool             `json:"disabled_ip"`
	Models               []restapi.Define `json:"-"`
}

func (o Options) ConfKey() string {
	return "account"
}

var options = Options{}

func New(key string, opt ...func(o *Options)) *Module {
	options.key = key
	options.ApiPrefix = "/manage"

	for _, f := range opt {
		f(&options)
	}

	service.DefaultConf = append(service.DefaultConf, &options)

	p := &Module{}

	return p
}

func (m *Module) Tasks() []service.Task {
	return []service.Task{
		{
			Run: func() {
				lm, ok := m.mods.Get(logsName)
				if !ok {
					return
				}

				t := time.Now().AddDate(0, -1, 0)
				_, err := restapi.DeleteMany(lm, ztype.Map{
					"record_at <": ztime.FormatTime(t),
				}, func(so *restapi.CondOptions) error {
					return nil
				})
				if err != nil {
					return
				}
			},
			Name: "clean_logs",
			Cron: "0 2 * * *",
		},
	}
}

var index = &Index{
	// Path: "/manage/base",
}

func (m *Module) Load(zdi.Invoker) (any, error) {
	return nil, m.DI.InvokeWithErrorOnly(func(c *service.Conf) error {
		m.Options = options
		if m.Options.key == "" {
			return errors.New("not account key")
		}
		index.Path = m.Options.ApiPrefix + "/base"
		m.Options.key = zstring.Pad(m.Options.key, 32, "0", zstring.PadLeft)

		index.plugin = m
		m.controllers = []service.Controller{
			index,
		}
		return nil
	})
}

func (m *Module) Start(zdi.Invoker) (err error) {
	if m.Options.InitDB != nil {
		m.db, err = m.Options.InitDB()
	} else {
		err = m.DI.Resolve(&m.db)
	}
	if err != nil || m.db == nil {
		return zerror.With(err, "init db error")
	}
	m.mods = restapi.NewModels(restapi.NewSQL(m.db, func(o *restapi.SQLOptions) {
		var restapiModule *restapi.Module
		if err := m.DI.Resolve(&restapiModule); err == nil {
			o.Prefix = restapiModule.Options.Prefix
		}
	}))

	if err = initModel(m); err != nil {
		return zerror.With(err, "init accoutModel error")
	}

	mod, ok := m.mods.Get(accountName)
	if !ok {
		return errors.New("account accoutModel not found")
	}

	index.accoutModel = mod
	index.permModel, _ = m.mods.Get(permName)
	index.roleModel, _ = m.mods.Get(roleName)

	permission := m.Options.InlayRBAC
	if permission == nil {
		permission = rbac.New()
	}

	if m.Options.RBACFile != "" {
		fPermission, err := rbac.ParseFile(m.Options.RBACFile)
		if err != nil {
			return zerror.With(err, "parse rbac file error")
		}

		err = permission.Merge(fPermission)
		if err != nil {
			return zerror.With(err, "merge rbac file error")
		}
	}

	if err = m.initMiddleware(permission); err != nil {
		return err
	}

	noLogIP = m.Options.DisabledLogIP
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
