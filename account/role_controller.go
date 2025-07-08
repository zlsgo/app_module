package account

import (
	"reflect"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model"
	"github.com/zlsgo/app_module/restapi"
)

type Role struct {
	service.App
	module *Module
	Path   string
}

var _ = reflect.TypeOf(&Role{})

func (h *Role) Init(r *znet.Engine) error {
	return PermisMiddleware(r)
}

// Get 角色列表
func (h *Role) Get(c *znet.Context) (data *model.PageData, err error) {
	filter := model.Filter{}
	key, _ := c.GetQuery("key")
	if key != "" {
		filter["key"] = key + "%"
	}

	return restapi.Page(c, h.module.index.roleModel.Model(), filter, func(o *model.CondOptions) {
		o.OrderBy = map[string]string{
			model.IDKey(): "desc",
		}
	})
}

// Post 新增角色
func (h *Role) Post(c *znet.Context) (id interface{}, err error) {
	return restapi.Insert(c, h.module.index.roleModel.Model(), func(data ztype.Map) (ztype.Map, error) {
		_ = data.Delete("inlay")

		alias := data.Get("alias").String()
		if alias == "" {
			return nil, zerror.InvalidInput.Text("别名不能为空")
		}

		id, err := h.module.index.roleModel.Model().FindCols("id", model.Filter{"alias": alias})
		if err != nil {
			return nil, err
		}
		if len(id) > 0 {
			return nil, zerror.InvalidInput.Text("别名已存在")
		}

		return data, nil
	})
}

// UIDPut 修改角色
func (h *Role) UIDPut(c *znet.Context) (res interface{}, err error) {
	id := c.GetParam("rid")
	j, _ := c.GetJSONs()
	return h.module.Inside.UpdateUser(id, j.Map())
}

// UIDDELETE 删除角色
func (h *Role) UIDDELETE(c *znet.Context) (res interface{}, err error) {
	id := c.GetParam("rid")
	user, err := GetAccountModel().FindOneByID(id)
	if err != nil {
		return nil, err
	}

	if user.IsEmpty() {
		return nil, zerror.InvalidInput.Text("角色不存在")
	}

	if user.Get("inlay").Bool() {
		return nil, zerror.InvalidInput.Text("不能删除内置角色")
	}

	_, err = GetAccountModel().DeleteByID(id)
	return nil, err
}
