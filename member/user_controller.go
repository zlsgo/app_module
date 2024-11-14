package member

import (
	"errors"
	"reflect"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zreflect"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/account/limiter"
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

	// 无需权限校验
	{
		noPerm := r.Group("/", func(e *znet.Engine) {
			e.Use(limiter.IPMiddleware())
		})
		noPerm.POST("/register", h.register)
		noPerm.POST("/login", h.login)
	}

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
	return restapi.UpdateById(c, oper, user.Id, func(_, data ztype.Map) (ztype.Map, error) {
		// 敏感字段不允许修改
		for _, k := range []string{"password", "salt", "account", "login_at", "provider", "provider_id", "provider_username", "status"} {
			delete(data, k)
		}

		return data, nil
	})
}

// register 注册
func (h *UserServer) register(c *znet.Context) (any, error) {
	if !h.module.Options.EnableRegister {
		return nil, zerror.WrapTag(zerror.InvalidInput)(errors.New("系统未开启注册"))
	}

	j, err := c.GetJSONs()
	if err != nil {
		return nil, zerror.InvalidInput.Text("数据格式错误")
	}

	account := j.Get("account").String()
	nickname := j.Get("nickname").String()
	if nickname == "" {
		nickname = account
	}

	password := j.Get("password").String()
	if password == "" {
		return nil, zerror.InvalidInput.Text("密码不能为空")
	}

	data := ztype.Map{
		"nickname": nickname,
		"account":  account,
		"password": password,
	}

	if account == "" {
		return nil, zerror.InvalidInput.Text("账号不能为空")
	}

	userModel, _ := h.module.UserModel()
	if exist, _ := userModel.Exists(model.Filter{
		"account": account,
	}); exist {
		return nil, zerror.InvalidInput.Text("账号已存在")
	}

	return userModel.Insert(data)
}

// login 登录
func (h *UserServer) login(c *znet.Context) (any, error) {
	return nil, nil
}
