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
	return h.module.UsePermisMiddleware(r, nil)
}

// Get 角色列表
func (h *Role) Get(c *znet.Context) (data *model.PageData, err error) {
	filter := model.Filter{}
	key, _ := c.GetQuery("key")
	if key != "" {
		filter["key"] = key + "%"
	}

	return restapi.Page(c, h.module.index.roleModel.Model(), filter, func(o *model.CondOptions) {
		o.OrderBy = []model.OrderByItem{{Field: model.IDKey(), Direction: "DESC"}}
	})
}

// Post 新增角色
func (h *Role) Post(c *znet.Context) (interface{}, error) {
	var (
		alias string
	)

	resp, err := restapi.Insert(c, h.module.index.roleModel.Model(), func(data ztype.Map) (ztype.Map, error) {
		_ = data.Delete("inlay")

		alias = data.Get("alias").String()
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
	if err != nil {
		return nil, err
	}

	h.module.invalidateRoleCache()
	if err = h.module.rebuildRBAC(); err != nil {
		return nil, err
	}

	return resp, nil
}

// RIDPatch 修改角色
func (h *Role) RIDPatch(c *znet.Context) (resp interface{}, err error) {
	id := c.GetParam("rid")
	resp, err = restapi.UpdateById(c, h.module.index.roleModel.Model(), id, func(old ztype.Map, data ztype.Map) (ztype.Map, error) {
		_ = data.Delete("inlay")
		_ = data.Delete("alias")
		return data, nil
	})
	if err != nil {
		return nil, err
	}

	h.module.invalidateRoleCache()
	if err = h.module.rebuildRBAC(); err != nil {
		return nil, err
	}

	return resp, err
}

// RIDDELETE 删除角色
func (h *Role) RIDDELETE(c *znet.Context) (interface{}, error) {
	id := c.GetParam("rid")
	resp, err := restapi.DeleteById(c, h.module.index.roleModel.Model(), id, func(old ztype.Map) error {
		if old.Get("inlay").Bool() {
			return zerror.InvalidInput.Text("不能删除内置角色")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.module.invalidateRoleCache()
	if err = h.module.rebuildRBAC(); err != nil {
		return nil, err
	}

	return resp, err
}
