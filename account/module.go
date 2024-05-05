package account

import (
	"errors"
	"reflect"
	"time"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/account/rbac"
	"github.com/zlsgo/app_module/quick/crud"
	"github.com/zlsgo/app_module/quick/define"
	"github.com/zlsgo/app_module/quick/sqlstorage"
	"github.com/zlsgo/zdb"
)

type Module struct {
	service.ModuleLifeCycle
	service.App
	db          *zdb.DB
	Options     *Options
	Controllers []service.Controller
	quick       *crud.Models
}

var (
	_ service.Module = &Module{}
	_                = reflect.TypeOf(&Module{})
)

func (m *Module) Name() string {
	return "Account"
}

type Options struct {
	InitDB               func() (*zdb.DB, error)  `json:"-"`
	SSEReconnect         func(uid, lastID string) `json:"-"`
	InlayRBAC            *rbac.RBAC               `json:"-"`
	AdminDefaultPassword string                   `json:"admin_default_password"`
	ApiPrefix            string                   `json:"prefix"`
	RBACFile             string                   `json:"rbac_file"`
	key                  string
	InlayUser            ztype.Maps      `json:"inlay_user"`
	Models               []define.Define `json:"-"`
	SSE                  znet.SSEOption  `json:"-"`
	Expire               int             `json:"expire"`
	Only                 bool            `json:"only"`
	DisabledLogIP        bool            `json:"disabled_ip"`
}

func (o Options) ConfKey() string {
	return "account"
}

func (o Options) DisableWrite() bool {
	return true
}

var options = Options{}

func New(key string, opt ...func(o Options) Options) *Module {
	options.key = key
	options.ApiPrefix = "/manage"

	for _, f := range opt {
		options = f(options)
	}

	service.DefaultConf = append(service.DefaultConf, &options)

	return &Module{}
}

func (m *Module) Tasks() []service.Task {
	return []service.Task{
		{
			Run: func() {
				lm, ok := m.quick.Get(logsName)
				if !ok {
					return
				}

				t := time.Now().AddDate(0, -1, 0)
				_, err := lm.DeleteMany(ztype.Map{
					"record_at <": ztime.FormatTime(t),
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
		m.Options = &options
		if m.Options.key == "" {
			return errors.New("not account key")
		}
		index.Path = m.Options.ApiPrefix + "/base"
		m.Options.key = zstring.Pad(m.Options.key, 32, "0", zstring.PadLeft)

		index.module = m
		m.Controllers = []service.Controller{
			index,
			&Message{
				module: m,
				Path:   m.Options.ApiPrefix + "/message",
			},
			&User{
				module: m,
				Path:   m.Options.ApiPrefix + "/user",
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

	// var (
	// 	restapiModule *restapi.Module
	// 	prefix        string
	// )
	// if err := m.DI.Resolve(&restapiModule); err == nil {
	// 	prefix = restapiModule.Options.Prefix
	// }

	sqlStorage := sqlstorage.NewSQL(m.db)
	m.quick = crud.NewModels(sqlStorage, func(o define.ModelsOptions) define.ModelsOptions {
		// if prefix != "" {
		// 	o.Prefix = prefix
		// }
		return o
	})

	if err = initModel(m); err != nil {
		return zerror.With(err, "init accout model error")
	}

	mod, ok := m.quick.Get(accountName)
	if !ok {
		return errors.New("accout model not found")
	}

	index.accoutModel = mod
	index.permModel, _ = m.quick.Get(permName)
	index.roleModel, _ = m.quick.Get(roleName)

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
	return m.Controllers
}

func (m *Module) Stop() error {
	return nil
}
