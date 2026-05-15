package member

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
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
	invoker func(c *znet.Context, member *User, o *model.Stores) (any, error)
)

var invokerValue zdi.PreInvoker = (invoker)(nil)

func (h invoker) Invoke(v []interface{}) ([]reflect.Value, error) {
	c, member, o := v[0].(*znet.Context), v[1].(*User), v[2].(*model.Stores)
	resp, err := h(c, member, o)
	if err != nil {
		return []reflect.Value{zreflect.ValueOf(err)}, nil
	}

	return []reflect.Value{zreflect.ValueOf(resp)}, nil
}

// Init 初始化路由
func (h *UserServer) Init(r *znet.Engine) error {
	_ = h.DI.Resolve(&h.module)
	r.Use(znet.RewriteErrorHandler(memberErrorHandler))
	r.Use(h.module.authModule.SessionMiddleware())

	znet.RegisterRender(invokerValue)

	r.Use(h.module.instance.GetMiddleware())

	return nil
}

// GETInfo 获取用户信息
func (h *UserServer) GETInfo(c *znet.Context, user *User) (any, error) {
	return user, nil
}

// PATCHMe 修改用户
func (h *UserServer) PATCHMe(c *znet.Context, user *User) (any, error) {
	model, _ := h.module.UserModel()
	resp, err := restapi.UpdateById(c, model, user.Id, func(_, data ztype.Map) (ztype.Map, error) {
		// 敏感字段不允许修改
		for _, k := range []string{"status", "auth_user_id"} {
			delete(data, k)
		}
		if extension, ok := data["extension"].(map[string]interface{}); ok {
			delete(extension, "auth")
			data["extension"] = extension
		}
		if extension, ok := data["extension"].(ztype.Map); ok {
			delete(extension, "auth")
			data["extension"] = extension
		}

		return data, nil
	})
	if err != nil {
		return nil, err
	}
	userCache.Delete(user.Id)
	return resp, nil
}

func memberErrorHandler(c *znet.Context, err error) {
	if err == nil {
		return
	}

	code := http.StatusInternalServerError
	switch zerror.GetTag(err) {
	case zerror.InvalidInput:
		code = http.StatusBadRequest
	case zerror.NotFound:
		code = http.StatusNotFound
	case zerror.Unauthorized:
		code = http.StatusUnauthorized
	case zerror.PermissionDenied:
		code = http.StatusForbidden
	case zerror.Cancelled:
		code = http.StatusRequestTimeout
	case zerror.Internal:
		code = http.StatusInternalServerError
	default:
		if parsed, convErr := strconv.Atoi(string(zerror.GetTag(err))); convErr == nil && parsed > 0 {
			code = parsed
		}
	}

	c.JSON(int32(code), znet.ApiData{
		Code: int32(code),
		Msg:  err.Error(),
	})
}
