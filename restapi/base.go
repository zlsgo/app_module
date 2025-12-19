package restapi

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model"
)

type controller struct {
	service.App
	options *Options
	Path    string
}

func (h *controller) Init(r *znet.Engine) error {
	var models *model.Stores
	err := h.DI.Resolve(&models)
	if err != nil {
		return errors.New("functional model has not been registered")
	}

	if h.options == nil {
		return errors.New("restapi options is required")
	}

	if h.options.Middleware != nil {
		r.Use(h.options.Middleware)
	}

	r.Any("/*", func(c *znet.Context) (any, error) {
		path := strings.SplitN(c.GetParam("*"), "/", 2)

		mod, ok := models.Get(path[0])
		if !ok {
			return nil, zerror.InvalidInput.Text("model not found")
		}

		method := c.Request.Method
		args := ""
		if len(path) > 1 {
			args = path[1]
		}

		if h.options.ResponseHook == nil || !h.options.ResponseHook(c, path[0], args, method) {
			r.HandleNotFound(c)
			return nil, nil
		}

		switch method {
		case "GET":
			return find(c, mod, args, model.Filter{}, nil)
		case "POST":
			return Insert(c, mod, nil)
		case "PUT":
			return UpdateById(c, mod, args, nil)
		case "PATCH":
			return UpdateById(c, mod, args, nil)
		case "DELETE":
			return DeleteById(c, mod, args, nil)
		default:
			r.HandleNotFound(c)
			return nil, nil
		}
	})
	return nil
}

func find(
	c *znet.Context,
	mod *model.Store,
	id string,
	filter model.Filter,
	fn func(o *model.CondOptions),
) (any, error) {
	switch id {
	case "":
		page, pagesize, err := model.Common.VarPages(c)
		if err != nil {
			return nil, err
		}
		const maxPageSize = 1000
		if pagesize > maxPageSize {
			pagesize = maxPageSize
		}
		return mod.Pages(page, pagesize, filter, func(o *model.CondOptions) {
			o.OrderBy = []model.OrderByItem{{Field: model.IDKey(), Direction: "DESC"}}

			if fn != nil {
				fn(o)
			}
		})
	case "*":
		return nil, zerror.InvalidInput.Text("全量查询不允许，请使用分页")
	default:
		filter[model.IDKey()] = id
		row, err := mod.FindOne(filter)
		if err != nil {
			return nil, err
		}

		if row.IsEmpty() {
			return nil, zerror.InvalidInput.Text("id not found")
		}

		return row, nil
	}
}
