package restapi

import (
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
	var models *model.Models
	err := h.DI.Resolve(&models)
	if err != nil {
		return err
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

		if h.options.ResponseHook != nil && !h.options.ResponseHook(c, path[0], args, method) {
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
	mod *model.Model,
	id string,
	filter model.Filter,
	fn func(o *model.CondOptions),
) (any, error) {
	switch id {
	case "":
		page, pagesize, _ := model.Common.VarPages(c)
		return mod.Pages(page, pagesize, filter, func(o *model.CondOptions) {
			o.OrderBy = map[string]string{model.IDKey(): "desc"}

			if fn != nil {
				fn(o)
			}
		})
	case "*":
		return mod.Find(filter, func(o *model.CondOptions) {
			o.OrderBy = map[string]string{model.IDKey(): "desc"}
			if fn != nil {
				fn(o)
			}
		})
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
