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
	plugin *Module
	Path   string
}

var (
	_ = reflect.TypeOf(&User{})
)

func (h *User) Init(r *znet.Engine) error {
	return h.plugin.RegMiddleware(r)
}

// Get 用户列表
func (h *User) Get(c *znet.Context) (data *restapi.PageData, err error) {
	filter := ztype.Map{}
	key, _ := c.GetQuery("key")
	if key != "" {
		filter["nickname like"] = "%" + key + "%"
		filter["account like"] = "%" + key + "%"
	}

	// uid := Ctx.UID(c)
	page, pagesize, _ := common.VarPages(c)
	data, err = h.plugin.AccountModel().Pages(page, pagesize, filter, func(co *restapi.CondOptions) error {
		co.OrderBy = map[string]string{
			restapi.IDKey: "desc",
		}
		co.Fields = h.plugin.AccountModel().m.GetFields("password")
		return nil
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
	if exist, _ := h.plugin.AccountModel().Exists(ztype.Map{
		"account": account,
	}); exist {
		return 0, zerror.WrapTag(zerror.InvalidInput)(errors.New("账号已存在"))
	}

	return h.plugin.AccountModel().Insert(data)
}

// UIDPut 修改用户
func (h *User) UIDPut(c *znet.Context) (err error) {
	id := c.GetParam("uid")
	user, err := h.plugin.AccountModel().FindOneByID(id)
	if err != nil {
		return err
	}

	if user.IsEmpty() {
		return zerror.InvalidInput.Text("用户不存在")
	}

	// 禁止修改超级管理员
	if user.Get("administrator").Bool() {
		return zerror.InvalidInput.Text("不能修改超级管理员")
	}

	j, _ := c.GetJSONs()
	data := j.Map()
	if err = fixUserData(j, &data); err != nil {
		return zerror.InvalidInput.Text(err.Error())
	}

	_, err = h.plugin.AccountModel().UpdateByID(id, data)

	return err
}

// fixUserData 修复并兼容用户数据各种情况
func fixUserData(j *zjson.Res, data *ztype.Map) error {
	// 禁止添加超级管理员
	_ = data.Delete("administrator")

	// 禁止标记为内置用户
	_ = data.Delete("inlay")

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