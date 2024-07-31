package member

import (
	"reflect"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zreflect"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model"
	"github.com/zlsgo/app_module/restapi"
)

type UserServer struct {
	service.App

	module *Module
	Path   string
}

var _ = reflect.TypeOf(&UserServer{})

type (
	invoker func(c *znet.Context, member *User, o *model.Models) (any, error)
)

var invokerValue zdi.PreInvoker = (invoker)(nil)

func (h invoker) Invoke(v []interface{}) ([]reflect.Value, error) {
	c, member, o := v[0].(*znet.Context), v[1].(*User), v[2].(*model.Models)
	resp, err := h(c, member, o)
	if err != nil {
		return []reflect.Value{zreflect.ValueOf(err)}, nil
	}

	return []reflect.Value{zreflect.ValueOf(resp)}, nil
}

// Init 初始化路由
func (h *UserServer) Init(r *znet.Engine) error {
	_ = h.DI.Resolve(&h.module)

	znet.RegisterPreInvoker(invokerValue)

	r.Use(h.module.instance.GetMiddleware())

	return nil
}

// GETMe 获取用户
func (h *UserServer) GETMe(c *znet.Context, user *User, opers *model.Models) (any, error) {
	return user, nil
}

// PATCHMe 修改用户
func (h *UserServer) PATCHMe(c *znet.Context, user *User, opers *model.Models) (any, error) {
	oper := opers.MustGet(modelName)
	return restapi.HanderPATCH(c, oper, user.Id, func(_, data ztype.Map) (ztype.Map, error) {
		// 敏感字段不允许修改
		for _, k := range []string{"password", "salt", "account", "login_at", "provider", "provider_id", "provider_username", "status"} {
			delete(data, k)
		}

		return data, nil
	})
}
