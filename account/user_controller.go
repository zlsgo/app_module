package account

import (
	"reflect"

	"github.com/sohaha/zlsgo/znet"
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
	return UsePermisMiddleware(r, nil)
}

// Get 用户列表
func (h *User) Get(c *znet.Context) (data *model.PageData, err error) {
	filter := model.Filter{
		"inlay": false,
	}
	account, _ := c.GetQuery("key")
	if account != "" {
		filter["account"] = account + "%"
	}
	page, pagesize, _ := model.Common.VarPages(c)

	data, err = GetAccountModel().Pages(page, pagesize, filter, func(co *model.CondOptions) {
		co.OrderBy = []model.OrderByItem{{Field: model.IDKey(), Direction: "DESC"}}
		co.Fields = GetAccountModel().m.GetFields("password", "salt")
	})
	// data.Items.ForEach(func(i int, item ztype.Map) bool {
	// 	id, _ := GetAccountModel().Schema().DeCryptID(item.Get(model.IDKey()).String())
	// 	_ = item.Set("uid", id)
	// 	return true
	// })
	return
}

// Post 新增用户
func (h *User) Post(c *znet.Context) (id interface{}, err error) {
	j, _ := c.GetJSONs()
	data := j.Map()
	return h.module.Inside.CreateUser(data)
}

// UIDPATCH 修改用户
func (h *User) UIDPATCH(c *znet.Context) (resp interface{}, err error) {
	id := c.GetParam("uid")
	j, _ := c.GetJSONs()

	return h.module.Inside.UpdateUser(id, j.Map())
}

// UIDDELETE 删除用户
func (h *User) UIDDELETE(c *znet.Context) (resp interface{}, err error) {
	return h.module.Inside.DeleteUser(c.GetParam("uid"))
}
