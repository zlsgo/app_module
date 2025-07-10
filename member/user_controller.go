package member

import (
	"errors"
	"reflect"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zreflect"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zvalid"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/account/jwt"
	"github.com/zlsgo/app_module/account/limiter"
	"github.com/zlsgo/app_module/model"
	"github.com/zlsgo/app_module/restapi"
	"golang.org/x/crypto/bcrypt"
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

	// 无需权限校验
	{
		noPerm := r.Group("/", func(e *znet.Engine) {
			e.Use(limiter.IPMiddleware())
		})
		noPerm.POST("/register", h.register)
		noPerm.POST("/login", h.login)
	}

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
	return restapi.UpdateById(c, model, user.Id, func(_, data ztype.Map) (ztype.Map, error) {
		// 敏感字段不允许修改
		for _, k := range []string{"password", "salt", "account", "login_at", "status"} {
			delete(data, k)
		}

		return data, nil
	})
}

// register 注册
func (h *UserServer) register(c *znet.Context) (any, error) {
	if !h.module.Options.EnableRegister {
		return nil, zerror.InvalidInput.Text("注册功能已关闭")
	}

	j, err := c.GetJSONs()
	if err != nil {
		return nil, zerror.InvalidInput.Text("数据格式错误")
	}

	account := j.Get("account").String()
	if !zvalid.Text(account).Regex("^[a-zA-Z]").Ok() {
		return nil, zerror.InvalidInput.Text("账号必须以字母开头")
	}
	if account == "" {
		return nil, zerror.InvalidInput.Text("账号不能为空")
	}

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
		"avatar":   j.Get("avatar").String(),
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
func (h *UserServer) login(c *znet.Context) (data any, err error) {
	json, _ := c.GetJSONs()
	account := json.Get("account").String()
	password := json.Get("password").String()

	if account == "" {
		err = zerror.InvalidInput.Text("请输入账号")
		return
	}

	if password == "" {
		err = zerror.InvalidInput.Text("请输入密码")
		return
	}

	userModel, _ := h.module.UserModel()
	user, err := userModel.FindOne(model.Filter{
		"account": account,
	})
	if err != nil {
		return nil, zerror.InvalidInput.Text(err.Error())
	}

	if user.IsEmpty() {
		return nil, zerror.InvalidInput.Text("用户不存在")
	}

	err = bcrypt.CompareHashAndPassword(user.Get("password").Bytes(), zstring.String2Bytes(password))
	if err != nil {
		return nil, zerror.InvalidInput.Text("账号或密码错误")
	}

	status := user.Get("status").Int()
	if status != 1 {
		switch status {
		case 0:
			return nil, zerror.WrapTag(zerror.Unauthorized)(errors.New("账号待激活"))
		default:
			return nil, zerror.WrapTag(zerror.Unauthorized)(errors.New("账号已停用"))
		}
	}

	salt := user.Get("salt").String()

	if h.module.Options.Only || salt == "" {
		salt = zstring.Rand(saltLen)
	}

	uid := user.Get(model.IDKey()).String()

	_, err = userModel.UpdateByID(uid, ztype.Map{
		"login_at": ztime.Now(),
		"salt":     salt,
	})
	if err != nil {
		return nil, zerror.InvalidInput.Text(err.Error())
	}

	accessToken, refreshToken, err := jwt.GenToken(salt+uid, h.module.Options.key, int64(h.module.Options.Expire), int64(0))
	if err != nil {
		return nil, err
	}

	return ztype.Map{
		"uid":           uid,
		"token":         accessToken,
		"refresh_token": refreshToken,
	}, nil
}
