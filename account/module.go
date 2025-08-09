package account

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	zsession "github.com/sohaha/zlsgo/znet/session"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/account/rbac"
	"github.com/zlsgo/app_module/model"
	"github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb"
)

type Module struct {
	service.ModuleLifeCycle
	service.App
	db          *zdb.DB
	mods        *model.Schemas
	index       *Index
	Request     *requestWith
	Inside      *inside
	Controllers []service.Controller
	Options     Options
	permission  *rbac.RBAC
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
	key                  string                   `z:"key"`
	ApiPrefix            string                   `z:"prefix"`
	RBACFile             string                   `z:"rbac_file"`
	AdminDefaultPassword string                   `z:"admin_default_password"`
	ModelPrefix          string                   `z:"model_prefix"`
	InlayUser            ztype.Maps               `z:"inlay_user"`
	Models               []schema.Schema          `z:"-"`
	SSE                  znet.SSEOption           `z:"-"`
	Expire               int64                    `z:"expire"`
	RefreshExpire        int64                    `z:"refresh_expire"`
	Only                 bool                     `z:"only"`
	DisabledLogIP        bool                     `z:"disabled_ip"`
	EnableRegister       bool                     `z:"register"`
	Session              zsession.Store           `z:"-"`
}

func (o Options) ConfKey() string {
	return "account"
}

func (o Options) DisableWrite() bool {
	return true
}

func New(key string, opt ...func(o *Options)) *Module {
	m := &Module{
		Options: zutil.Optional(Options{key: key, ApiPrefix: "/manage"}, opt...),
		index:   &Index{},
		Request: &requestWith{},
	}
	m.Inside = &inside{m: m}
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
				// 删除一个月前的日志
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

func (m *Module) Load(zdi.Invoker) (any, error) {
	return nil, m.DI.InvokeWithErrorOnly(func(c *service.Conf) error {
		if m.Options.key == "" {
			return errors.New("not account key")
		}
		m.Options.key = zstring.Pad(m.Options.key, 32, "0", zstring.PadRight)
		m.Options.ApiPrefix = strings.TrimSuffix(m.Options.ApiPrefix, "/")

		m.index.Path = m.Options.ApiPrefix + "/base"

		m.index.module = m
		m.Controllers = []service.Controller{
			m.index,
			&Message{
				module: m,
				Path:   m.Options.ApiPrefix + "/message",
			},
			&User{
				module: m,
				Path:   m.Options.ApiPrefix + "/user",
			},
			&Role{
				module: m,
				Path:   m.Options.ApiPrefix + "/role",
			},
			&Permission{
				module: m,
				Path:   m.Options.ApiPrefix + "/permission",
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

	m.mods = model.NewSchemas(di.(zdi.Injector), model.NewSQL(m.db, m.Options.ModelPrefix), model.SchemaOptions{})

	if err = initModel(m); err != nil {
		return zerror.With(err, "init accoutModel error")
	}

	mod, ok := m.mods.Get(accountName)
	if !ok {
		return errors.New("account accoutModel not found")
	}

	m.index.accoutModel = mod
	m.index.permModel, _ = m.mods.Get(permName)
	m.index.roleModel, _ = m.mods.Get(roleName)

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

	m.permission = permission

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
