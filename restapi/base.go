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
	var operations *model.Operations
	err := h.DI.Resolve(&operations)
	if err != nil {
		return err
	}

	if h.options.Middleware != nil {
		r.Use(h.options.Middleware)
	}

	r.Any("/*", func(c *znet.Context) (any, error) {
		path := strings.SplitN(c.GetParam("*"), "/", 2)

		oper, ok := operations.Get(path[0])
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
			return handerGet(c, oper, args, ztype.Map{}, nil)
		case "POST":
			return HanderPost(c, oper, nil)
		case "PUT":
			return HanderPut(oper, c, args)
		case "PATCH":
			return HanderPATCH(c, oper, args, nil)
		case "DELETE":
			return HanderDelete(c, oper, args, nil)
		default:
			r.HandleNotFound(c)
			return nil, nil
		}
	})
	return nil
}

func HanderGet(c *znet.Context, oper *model.Operation, id string, fn func(o *model.CondOptions)) (ztype.Map, error) {
	res, err := handerGet(c, oper, id, ztype.Map{}, fn)
	if err != nil {
		return nil, err
	}

	return res.(ztype.Map), nil
}

func HanderPage(c *znet.Context, oper *model.Operation, filter ztype.Map, fn func(o *model.CondOptions)) (*model.PageData, error) {
	res, err := handerGet(c, oper, "", filter, fn)
	if err != nil {
		return nil, err
	}

	return res.(*model.PageData), nil
}

func handerGet(c *znet.Context, oper *model.Operation, id string, filter ztype.Map, fn func(o *model.CondOptions)) (any, error) {
	switch id {
	case "":
		page, pagesize, _ := model.Common.VarPages(c)
		return oper.Pages(page, pagesize, filter, func(o *model.CondOptions) {
			o.OrderBy = map[string]string{model.IDKey(): "desc"}

			if fn != nil {
				fn(o)
			}
		})
	default:
		filter[model.IDKey()] = id
		row, err := oper.FindOne(filter)
		if err != nil {
			return nil, err
		}

		if row.IsEmpty() {
			return nil, zerror.InvalidInput.Text("id not found")
		}

		return row, nil
	}
}

func HanderPost(c *znet.Context, oper *model.Operation, fn func(data ztype.Map) (ztype.Map, error)) (any, error) {
	j, _ := c.GetJSONs()
	data := j.Map()

	if fn != nil {
		var err error
		data, err = fn(data)
		if err != nil {
			return nil, err
		}
	}

	id, err := oper.Insert(data)
	if err != nil {
		return nil, err
	}

	return ztype.Map{"id": id}, nil
}

func HanderDelete(c *znet.Context, oper *model.Operation, id string, handler func(old ztype.Map) error) (any, error) {
	if id == "" {
		return nil, zerror.InvalidInput.Text("id cannot empty")
	}
	if handler != nil {
		info, err := oper.FindOneByID(id)
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

	total, err := oper.DeleteByID(id)
	return ztype.Map{"total": total}, err
}

func HanderPut(oper *model.Operation, c *znet.Context, id string) (any, error) {
	if id == "" {
		return nil, zerror.InvalidInput.Text("id cannot empty")
	}

	j, _ := c.GetJSONs()
	data := j.Map()

	total, err := oper.UpdateByID(id, data)
	return ztype.Map{"total": total}, err
}

func HanderPATCH(c *znet.Context, oper *model.Operation, id string, handler func(old ztype.Map, data ztype.Map) (ztype.Map, error)) (any, error) {
	if id == "" {
		return nil, zerror.InvalidInput.Text("id cannot empty")
	}

	j, _ := c.GetJSONs()
	data := j.Map()

	if handler != nil {
		info, err := oper.FindOneByID(id)
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

	total, err := oper.UpdateByID(id, data)
	return ztype.Map{"total": total}, err
}
