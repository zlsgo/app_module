package account

import (
	"errors"
	"reflect"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/common"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/restapi"
)

type User struct {
	service.App
	module *Module
	Path   string
}

var (
	_ = reflect.TypeOf(&User{})
)

func (h *User) Init(r *znet.Engine) error {
	return PermisMiddleware(r)
}

// Get 用户列表
func (h *User) Get(c *znet.Context) (data *restapi.PageData, err error) {
	filter := ztype.Map{}
	account, _ := c.GetQuery("account")
	if account != "" {
		filter["account"] = account + "%"
	}
	page, pagesize, _ := common.VarPages(c)

	data, err = h.module.AccountModel().Pages(page, pagesize, filter, func(co *restapi.CondOptions) error {
		co.OrderBy = map[string]string{
			restapi.IDKey: "desc",
		}
		co.Fields = h.module.AccountModel().m.GetFields("password", "salt")
		return nil
	})
	data.Items.ForEach(func(i int, item ztype.Map) bool {
		id, _ := h.module.AccountModel().DeCryptID(item.Get(restapi.IDKey).String())
		_ = item.Set("id", id)
		return true
	})
	return
}

// Post 新增用户
func (h *User) Post(c *znet.Context) (id interface{}, err error) {
	j, _ := c.GetJSONs()
	data := j.Map()

	if err := fixUserData(j, &data); err != nil {
		return nil, zerror.WrapTag(zerror.InvalidInput)(err)
	}

	// DEV: 需要校验角色是否存在
	roles := data.Get("role").SliceString()
	_ = roles

	account := data.Get("account").String()
	if !data.Get("nickname").Exists() {
		data.Set("nickname", account)
	}

	if !data.Get("status").Exists() {
		data.Set("status", 1)
	}

	// 检查账号是否存在
	if exist, _ := h.module.AccountModel().Exists(ztype.Map{
		"account": account,
	}); exist {
		return 0, zerror.WrapTag(zerror.InvalidInput)(errors.New("账号已存在"))
	}

	return h.module.AccountModel().Insert(data)
}

// UIDPut 修改用户
func (h *User) UIDPut(c *znet.Context) (res interface{}, err error) {
	id := c.GetParam("uid")
	user, err := h.module.AccountModel().FindOneByID(id)
	if err != nil {
		return nil, err
	}

	if user.IsEmpty() {
		return nil, zerror.InvalidInput.Text("用户不存在")
	}

	// 禁止修改超级管理员
	if user.Get("administrator").Bool() {
		return nil, zerror.InvalidInput.Text("不能修改超级管理员")
	}

	j, _ := c.GetJSONs()
	data := j.Map()
	if err = fixUserData(j, &data); err != nil {
		return nil, zerror.InvalidInput.Text(err.Error())
	}
	_, err = h.module.AccountModel().UpdateByID(id, data)

	return nil, err
}

// fixUserData 修复并兼容用户数据各种情况
func fixUserData(j *zjson.Res, data *ztype.Map) error {
	// 禁止添加超级管理员
	_ = data.Delete("administrator")
	// 禁止标记为内置用户
	_ = data.Delete("inlay")
	// 验证盐不应该可以人为修改
	_ = data.Delete("salt")
	// 登录时间不应该可以人为修改
	_ = data.Delete("login_at")

	if !j.Get("role").IsArray() {
		role := data.Get("role").String()
		if role != "" {
			data.Set("role", []interface{}{role})
		}
	}

	if data.Has("password") {
		if data.Get("password").String() == "" {
			return errors.New("密码不能为空")
		}
	}

	return nil
}
