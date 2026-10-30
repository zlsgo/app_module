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

// Done 在启动阶段（Load/Start 之后）被框架调用，不是退出钩子，
// 这里不能释放连接。真正的关闭由 Stop 负责，而 Stop 由业务方
// （如 app_core 的 RunWeb 退出路径）显式调用。
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
