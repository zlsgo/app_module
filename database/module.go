package database

import (
	"reflect"

	"github.com/zlsgo/app_module/restapi"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/zlsgo/app_core/common"
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
	return nil, p.DI.InvokeWithErrorOnly(func(d zdi.TypeMapper, c *service.Conf) (err error) {
		err = c.Unmarshal((Options{}).ConfKey(), &options)
		if err != nil {
			return
		}

		dealOldColumn := restapi.DealOldColumnNone
		if options.Mode == nil {
			options.Mode = &Mode{}
		}
		if options.Mode.DelteColumn {
			dealOldColumn = restapi.DealOldColumnDelete
		}
		restapi.Inside.DeleteOldColumn(dealOldColumn)

		p.db, err = initDB(options)
		common.Fatal(err)

		d.Map(p.db)

		return
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
