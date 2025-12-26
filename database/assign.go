package database

import (
	"errors"
	"reflect"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/zdb"
)

// Single 为单库模块
type Single struct {
	service.App
	db     *zdb.DB
	initDB InitFunc
}

var (
	_ service.Module = &Single{}
	_                = reflect.TypeOf(&Single{})
)

// NewSingle 创建单库模块
func NewSingle(initDB InitFunc) *Single {
	return &Single{
		initDB: initDB,
	}
}

// DB 返回当前数据库实例
func (p *Single) DB() (*zdb.DB, error) {
	if p.db == nil {
		return nil, errors.New("database not init")
	}
	return p.db, nil
}

// Name 返回模块名
func (p *Single) Name() string {
	return "SingleDatabase"
}

// Tasks 返回模块任务
func (p *Single) Tasks() []service.Task {
	return []service.Task{}
}

// Load 初始化模块
func (p *Single) Load(di zdi.Invoker) (any, error) {
	err := di.InvokeWithErrorOnly(func(c *service.Conf) (err error) {
		p.db, err = p.initDB(p.DI)
		return
	})

	return nil, err
}

// Start 启动模块
func (p *Single) Start(zdi.Invoker) error {
	return nil
}

// Done 停止模块
func (p *Single) Done(zdi.Invoker) error {
	return nil
}

// Controller 返回模块控制器
func (p *Single) Controller() []service.Controller {
	return []service.Controller{}
}
