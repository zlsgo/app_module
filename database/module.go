package database

import (
	"reflect"

	"github.com/zlsgo/app_module/model"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/zdb"
)

// Plugin 为数据库主模块
type Plugin struct {
	service.App
	db *zdb.DB
}

var (
	_ service.Module = &Plugin{}
	_                = reflect.TypeOf(&Plugin{})
)

// Name 返回模块名
func (p *Plugin) Name() string {
	return "Database"
}

// Tasks 返回模块任务
func (p *Plugin) Tasks() []service.Task {
	return []service.Task{}
}

// Load 初始化模块
func (p *Plugin) Load(zdi.Invoker) (any, error) {
	return nil, p.DI.InvokeWithErrorOnly(func(d zdi.TypeMapper, c *service.Conf) (err error) {
		err = c.Unmarshal((Options{}).ConfKey(), &options)
		if err != nil {
			return
		}

		options.Mode = normalizeMode(options.Mode)
		applyMode(options.Mode)

		p.db, err = initDB(options)
		if err != nil {
			return err
		}

		d.Map(p.db)

		return
	})
}

// Start 启动模块
func (p *Plugin) Start(zdi.Invoker) error {
	return nil
}

// Done 停止模块
func (p *Plugin) Done(zdi.Invoker) error {
	return nil
}

// Controller 返回模块控制器
func (p *Plugin) Controller() []service.Controller {
	return []service.Controller{}
}

func normalizeMode(mode *Mode) *Mode {
	if mode == nil {
		return &Mode{}
	}
	return mode
}

func applyMode(mode *Mode) {
	if mode == nil {
		mode = &Mode{}
	}
	if mode.DelteColumn {
		model.SetDefaultSchemaOptions(func(o *model.SchemaOptions) {
			o.OldColumn = model.DealOldColumnDelete
		})
	} else {
		model.SetDefaultSchemaOptions(func(o *model.SchemaOptions) {
			o.OldColumn = model.DealOldColumnNone
		})
	}
}
