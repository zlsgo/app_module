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
	"github.com/zlsgo/app_module/model/schema"
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
	index        *Index
	Request      *requestWith
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
	AdminDefaultPassword string                   `z:"admin_default_password"` // 管理员默认密码
	ApiPrefix            string                   `z:"prefix"`                 // 接口前缀
	RBACFile             string                   `z:"rbac_file"`              // rbac 文件
	key                  string                   `z:"key"`                    // 密钥
	InlayUser            ztype.Maps               `z:"inlay_user"`             // 默认用户
	Models               []schema.Schema          `z:"-"`
	ModelPrefix          string                   `z:"model_prefix"` // 模型前缀
	SSE                  znet.SSEOption           `z:"-"`
	Expire               int                      `z:"expire"`      // access token 过期时间
	Only                 bool                     `z:"only"`        // 是否只允许一处登录
	DisabledLogIP        bool                     `z:"disabled_ip"` // 是否禁用日志 IP 记录
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

	m.index = &Index{}
	m.Request = &requestWith{}

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
		if m.Options.ModelPrefix != "" {
			o.Prefix = m.Options.ModelPrefix
		}
	}), model.SchemaOptions{})

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
