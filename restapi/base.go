package restapi

import (
	"strings"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
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
			return handerGet(c, mod, args, model.Filter{}, nil)
		case "POST":
			return HanderPost(c, mod, nil)
		case "PUT":
			return HanderPut(c, mod, args)
		case "PATCH":
			return HanderPATCH(c, mod, args, nil)
		case "DELETE":
			return HanderDelete(c, mod, args, nil)
		default:
			r.HandleNotFound(c)
			return nil, nil
		}
	})
	return nil
}

func HanderGet(
	c *znet.Context,
	mod *model.Model,
	id string,
	fn func(o *model.CondOptions),
) (ztype.Map, error) {
	res, err := handerGet(c, mod, id, model.Filter{}, fn)
	if err != nil {
		return nil, err
	}

	return res.(ztype.Map), nil
}

func HanderPage(
	c *znet.Context,
	mod *model.Model,
	filter model.Filter,
	fn func(o *model.CondOptions),
) (*model.PageData, error) {
	res, err := handerGet(c, mod, "", filter, fn)
	if err != nil {
		return nil, err
	}

	return res.(*model.PageData), nil
}

func HanderGets(
	c *znet.Context,
	mod *model.Model,
	filter model.Filter,
	fn func(o *model.CondOptions),
) (ztype.Maps, error) {
	res, err := handerGet(c, mod, "*", filter, fn)
	if err != nil {
		return nil, err
	}

	return res.(ztype.Maps), nil
}

func handerGet(
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

func HanderPost(
	c *znet.Context,
	mod *model.Model,
	fn func(data ztype.Map) (ztype.Map, error),
) (any, error) {
	j, _ := c.GetJSONs()
	data := j.Map()

	if fn != nil {
		var err error
		data, err = fn(data)
		if err != nil {
			return nil, err
		}
	}

	id, err := mod.Insert(data)
	if err != nil {
		return nil, err
	}

	return ztype.Map{"id": id}, nil
}

func HanderDelete(
	c *znet.Context,
	mod *model.Model,
	id string,
	handler func(old ztype.Map) error,
) (any, error) {
	if id == "" {
		return nil, zerror.InvalidInput.Text("id cannot empty")
	}
	if handler != nil {
		info, err := mod.FindOneByID(id)
		if err != nil {
			return nil, err
		}
		if info.IsEmpty() {
			return nil, zerror.InvalidInput.Text("id not found")
		}
		err = handler(info)
		if err != nil {
			return nil, err
		}
	}

	total, err := mod.DeleteByID(id)
	return ztype.Map{"total": total}, err
}

func HanderPut(c *znet.Context, mod *model.Model, id string) (any, error) {
	if id == "" {
		return nil, zerror.InvalidInput.Text("id cannot empty")
	}

	j, err := c.GetJSONs()
	if err != nil {
		return nil, err
	}

	total, err := mod.UpdateByID(id, j.Map())
	return ztype.Map{"total": total}, err
}

func HanderPATCH(
	c *znet.Context,
	mod *model.Model,
	id string,
	handler func(old ztype.Map, data ztype.Map) (ztype.Map, error),
) (any, error) {
	if id == "" {
		return nil, zerror.InvalidInput.Text("id cannot empty")
	}

	j, err := c.GetJSONs()
	if err != nil {
		return nil, err
	}

	data := j.Map()

	if handler != nil {
		info, err := mod.FindOneByID(id)
		if err != nil {
			return nil, err
		}

		if info.IsEmpty() {
			return nil, zerror.InvalidInput.Text("id not found")
		}

		data, err = handler(info, data)
		if err != nil {
			return nil, err
		}
	}

	total, err := mod.UpdateByID(id, data)
	return ztype.Map{"total": total}, err
}
