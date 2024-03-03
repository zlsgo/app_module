package database

import (
	"errors"
	"reflect"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/zdb"
)

type Single struct {
	service.App
	db     *zdb.DB
	initDB InitFunc
}

var (
	_ service.Module = &Single{}
	_                = reflect.TypeOf(&Single{})
)

func NewSingle(initDB InitFunc) *Single {
	return &Single{
		initDB: initDB,
	}
}

func (p *Single) DB() (*zdb.DB, error) {
	if p.db == nil {
		return nil, errors.New("database not init")
	}
	return p.db, nil
}

func (p *Single) Name() string {
	return "SingleDatabase"
}

func (p *Single) Tasks() []service.Task {
	return []service.Task{}
}

func (p *Single) Load(di zdi.Invoker) (any, error) {
	err := di.InvokeWithErrorOnly(func(c *service.Conf) (err error) {
		p.db, err = p.initDB(p.DI)
		return
	})

	return nil, err
}

func (p *Single) Start(zdi.Invoker) error {
	return nil
}

func (p *Single) Done(zdi.Invoker) error {
	return nil
}

func (p *Single) Controller() []service.Controller {
	return []service.Controller{}
}
