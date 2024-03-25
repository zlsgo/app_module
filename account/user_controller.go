package account

import (
	"reflect"

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

	data, err = GetAccountModel().Pages(page, pagesize, filter, func(co *restapi.CondOptions) error {
		co.OrderBy = map[string]string{
			restapi.IDKey: "desc",
		}
		co.Fields = GetAccountModel().m.GetFields("password", "salt")
		return nil
	})
	data.Items.ForEach(func(i int, item ztype.Map) bool {
		id, _ := GetAccountModel().DeCryptID(item.Get(restapi.IDKey).String())
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
