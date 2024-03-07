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
	ms          *restapi.Models
	Options     Options
	controllers []service.Controller
}

var (
	_ service.Module = &Module{}
	_                = reflect.TypeOf(&Module{})
)

func (p *Module) Name() string {
	return "Account"
}

type Options struct {
	InitDB               func() (*zdb.DB, error) `json:"-"`
	key                  string
	InlayRBAC            *rbac.RBAC       `json:"-"`
	RBACFile             string           `json:"rbac_file"`
	Prefix               string           `json:"prefix"`
	InlayUser            ztype.Maps       `json:"inlay_user"`
	AdminDefaultPassword string           `json:"admin_default_password"`
	Expire               int              `json:"expire"`
	Only                 bool             `json:"only"`
	NoLogIP              bool             `json:"no_ip"`
	Models               []restapi.Define `json:"-"`
}

func (o Options) ConfKey() string {
	return "account"
}

var options = Options{}

func New(key string, opt ...func(o *Options)) *Module {
	options.key = key
	options.Prefix = "/manage"

	for _, f := range opt {
		f(&options)
	}

	service.DefaultConf = append(service.DefaultConf, &options)

	p := &Module{}

	return p
}

func (p *Module) Tasks() []service.Task {
	return []service.Task{
		{
			Run: func() {
				lm, ok := p.ms.Get(logsName)
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

func (p *Module) Load(zdi.Invoker) (any, error) {
	return nil, p.DI.InvokeWithErrorOnly(func(c *service.Conf) error {
		p.Options = options
		if p.Options.key == "" {
			return errors.New("not account key")
		}
		index.Path = p.Options.Prefix + "/base"
		p.Options.key = zstring.Pad(p.Options.key, 32, "0", zstring.PadLeft)

		index.plugin = p
		p.controllers = []service.Controller{
			index,
		}
		return nil
	})
}

func (p *Module) Start(zdi.Invoker) (err error) {
	if p.Options.InitDB != nil {
		p.db, err = p.Options.InitDB()
	} else {
		err = p.DI.Resolve(&p.db)
	}
	if err != nil || p.db == nil {
		return zerror.With(err, "init db error")
	}
	p.ms = restapi.NewModels(restapi.NewSQL(p.db), func(o *restapi.ModelsOptions) {
		o.Prefix = "model_"
	})

	if err = initModel(p); err != nil {
		return zerror.With(err, "init model error")
	}

	m, ok := p.ms.Get(accountName)
	if !ok {
		return errors.New("account model not found")
	}

	index.model = m
	index.permModel, _ = p.ms.Get(permName)
	index.roleModel, _ = p.ms.Get(roleName)

	permission := p.Options.InlayRBAC
	if permission == nil {
		permission = rbac.New()
	}

	if p.Options.RBACFile != "" {
		fPermission, err := rbac.ParseFile(p.Options.RBACFile)
		if err != nil {
			return zerror.With(err, "parse rbac file error")
		}

		err = permission.Merge(fPermission)
		if err != nil {
			return zerror.With(err, "merge rbac file error")
		}
	}

	if err = p.initMiddleware(permission); err != nil {
		return err
	}

	noLogIP = p.Options.NoLogIP
	return
}

func (p *Module) Done(zdi.Invoker) (err error) {
	return nil
}

func (p *Module) Controller() []service.Controller {
	return p.controllers
}

func (p *Module) Stop() error {
	return nil
}
