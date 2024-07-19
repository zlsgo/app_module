package account

import (
	"reflect"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model"
)

type User struct {
	service.App
	module *Module
	Path   string
}

var _ = reflect.TypeOf(&User{})

func (h *User) Init(r *znet.Engine) error {
	return PermisMiddleware(r)
}

// Get 用户列表
func (h *User) Get(c *znet.Context) (data *model.PageData, err error) {
	filter := ztype.Map{
		"inlay": false,
	}
	account, _ := c.GetQuery("account")
	if account != "" {
		filter["account"] = account + "%"
	}
	page, pagesize, _ := model.Common.VarPages(c)

	data, err = GetAccountModel().Pages(page, pagesize, filter, func(co *model.CondOptions) {
		co.OrderBy = map[string]string{
			model.IDKey(): "desc",
		}
		co.Fields = GetAccountModel().m.GetFields("password", "salt")
		return
	})
	data.Items.ForEach(func(i int, item ztype.Map) bool {
		id, _ := GetAccountModel().DeCryptID(item.Get(model.IDKey()).String())
		_ = item.Set("id", id)
		return true
	})
	return
}

// Post 新增用户
func (h *User) Post(c *znet.Context) (id interface{}, err error) {
	j, _ := c.GetJSONs()
	data := j.Map()
	return Inside.CreateUser(data)
}

// UIDPut 修改用户
func (h *User) UIDPut(c *znet.Context) (res interface{}, err error) {
	id := c.GetParam("uid")
	j, _ := c.GetJSONs()
	return Inside.UpdateUser(id, j.Map())
}

// UIDDELETE 删除用户
func (h *User) UIDDELETE(c *znet.Context) (res interface{}, err error) {
	id := c.GetParam("uid")
	user, err := GetAccountModel().FindOneByID(id)
	if err != nil {
		return nil, err
	}

	if user.IsEmpty() {
		return nil, zerror.InvalidInput.Text("用户不存在")
	}

	if user.Get("inlay").Bool() {
		return nil, zerror.InvalidInput.Text("不能删除内置用户")
	}

	_, err = GetAccountModel().DeleteByID(id)
	return nil, err
}
