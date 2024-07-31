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
	"github.com/zlsgo/app_module/model"
	"github.com/zlsgo/app_module/model/define"
	"github.com/zlsgo/zdb"
)

type Module struct {
	service.ModuleLifeCycle
	service.App
	db           *zdb.DB
	mods         *model.Schemas
	accountModel *AccountModel
	Controllers  []service.Controller
	Options      Options
}

var (
	_ service.Module = &Module{}
	_                = reflect.TypeOf(&Module{})
)

func (m *Module) Name() string {
	return "Account"
}

type Options struct {
	InitDB               func() (*zdb.DB, error)  `z:"-"`
	SSEReconnect         func(uid, lastID string) `z:"-"`
	InlayRBAC            *rbac.RBAC               `z:"-"`
	AdminDefaultPassword string                   `z:"admin_default_password"`
	ApiPrefix            string                   `z:"prefix"`
	RBACFile             string                   `z:"rbac_file"`
	key                  string
	InlayUser            ztype.Maps      `z:"inlay_user"`
	Models               []define.Schema `z:"-"`
	SSE                  znet.SSEOption  `z:"-"`
	Expire               int             `z:"expire"`
	Only                 bool            `z:"only"`
	DisabledLogIP        bool            `z:"disabled_ip"`
}

func (o Options) ConfKey() string {
	return "account"
}

func (o Options) DisableWrite() bool {
	return true
}

func New(key string, opt ...func(o *Options)) *Module {
	m := &Module{
		Options: Options{key: key, ApiPrefix: "/manage"},
	}

	for _, f := range opt {
		f(&m.Options)
	}

	return m
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
				_, err := model.DeleteMany(lm, ztype.Map{
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
		if m.Options.key == "" {
			return errors.New("not account key")
		}
		m.Options.key = zstring.Pad(m.Options.key, 32, "0", zstring.PadRight)

		index.Path = m.Options.ApiPrefix + "/base"

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

	m.mods = model.NewSchemas(di.(zdi.Injector), model.NewSQL(m.db, func(o *model.SQLOptions) {
		var restapiModule *model.Module
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
	return m.Controllers
}

func (m *Module) Stop() error {
	return nil
}
