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

type Permission struct {
	service.App
	module *Module
	Path   string
}

var _ = reflect.TypeOf(&Permission{})

func (h *Permission) Init(r *znet.Engine) error {
	return UsePermisMiddleware(r, nil)
}

// Get 规则列表
func (h *Permission) Get(c *znet.Context) (data *model.PageData, err error) {
	filter := model.Filter{}
	key, _ := c.GetQuery("key")
	if key != "" {
		filter["key"] = key + "%"
	}

	return restapi.Page(c, h.module.index.permModel.Model(), filter, func(o *model.CondOptions) {
		o.OrderBy = map[string]string{
			model.IDKey(): "desc",
		}
	})
}

// Post 新增规则
func (h *Permission) Post(c *znet.Context) (interface{}, error) {
	var (
		alias      string
		permission []int
	)

	resp, err := restapi.Insert(c, h.module.index.permModel.Model(), func(data ztype.Map) (ztype.Map, error) {
		_ = data.Delete("inlay")

		alias = data.Get("alias").String()
		if alias == "" {
			return nil, zerror.InvalidInput.Text("别名不能为空")
		}

		id, err := h.module.index.permModel.Model().FindCols("id", model.Filter{"alias": alias})
		if err != nil {
			return nil, err
		}
		if len(id) > 0 {
			return nil, zerror.InvalidInput.Text("别名已存在")
		}

		permission = data.Get("permission").SliceInt()
		return data, nil
	})
	if err != nil {
		return nil, err
	}

	if alias != "" {
		h.module.setPermission(h.module.permission, ztype.Map{
			"alias":      alias,
			"permission": permission,
		})
	}

	return resp, nil
}

// PIDPATCH 修改规则
func (h *Permission) PIDPATCH(c *znet.Context) (resp interface{}, err error) {
	id := c.GetParam("pid")
	resp, err = restapi.UpdateById(c, h.module.index.permModel.Model(), id, func(old ztype.Map, data ztype.Map) (ztype.Map, error) {
		_ = data.Delete("inlay")
		_ = data.Delete("alias")
		return data, nil
	})
	if err != nil {
		return nil, err
	}

	return resp, err
}

// PIDDELETE 删除规则
func (h *Permission) PIDDELETE(c *znet.Context) (interface{}, error) {
	id := c.GetParam("pid")
	resp, err := restapi.DeleteById(c, h.module.index.permModel.Model(), id, func(old ztype.Map) error {
		if old.Get("inlay").Bool() {
			return zerror.InvalidInput.Text("不能删除内置规则")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return resp, err
}
