package database

import (
	"reflect"

	"github.com/zlsgo/app_module/database/model"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/zdb"
)

type Plugin struct {
	service.App
	db *zdb.DB
}

var (
	_ service.Module = &Plugin{}
	_                = reflect.TypeOf(&Plugin{})
)

func (p *Plugin) Name() string {
	return "Database"
}

func (p *Plugin) Tasks() []service.Task {
	return []service.Task{}
}

func (p *Plugin) Load(zdi.Invoker) (any, error) {
	return nil, p.DI.InvokeWithErrorOnly(func(d zdi.TypeMapper, c *service.Conf) {
		model.Inside.DeleteOldColumn(model.DealOldColumn(c.Get((Options{}).ConfKey() + ".Model.OldColumn").Uint8()))
		p.db = initDB(c)
		d.Map(p.db)
	})
}

func (p *Plugin) Start(zdi.Invoker) error {
	return nil
}

func (p *Plugin) Done(zdi.Invoker) error {
	return nil
}

func (p *Plugin) Controller() []service.Controller {
	return []service.Controller{}
}
