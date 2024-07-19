package member

import (
	"reflect"

	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model"
)

type UserServer struct {
	service.App

	module *Module
	Model  func() (*model.Model, bool)
	Path   string
}

var _ = reflect.TypeOf(&UserServer{})

func (h *UserServer) Init(r *znet.Engine) error {
	r.Use(h.module.Middleware(true))
	return nil
}

func (h *UserServer) GETMe(c *znet.Context, user *User) (any, error) {
	return user, nil
}
