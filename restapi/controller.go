package restapi

import (
	"strings"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model"
)

type controller struct {
	Path    string
	options *Options
	service.App
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

	r.Any("/*", func(c *znet.Context) (interface{}, error) {
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
			return h.handerGet(oper, c, args)
		case "POST":
			return h.handerPost(oper, c, args)
		case "PUT":
			return h.handerPut(oper, c, args)
		case "DELETE":
			return h.handerDelete(oper, c, args)
		default:
			r.HandleNotFound(c)
			return nil, nil
		}
	})
	return nil
}

func (h *controller) handerGet(oper *model.Operation, c *znet.Context, args string) (interface{}, error) {
	switch args {
	case "":
		page, pagesize, _ := model.Common.VarPages(c)
		return oper.Pages(page, pagesize, nil, func(o *model.CondOptions) {
			o.OrderBy = map[string]string{model.IDKey(): "desc"}
		})
	default:
		row, err := oper.FindOneByID(args)
		if err != nil {
			return nil, err
		}

		if row.IsEmpty() {
			return nil, zerror.InvalidInput.Text("id not found")
		}

		return row, nil
	}
}

func (h *controller) handerPost(oper *model.Operation, c *znet.Context, _ string) (interface{}, error) {
	j, _ := c.GetJSONs()
	data := j.Map()

	id, err := oper.Insert(data)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (h *controller) handerDelete(oper *model.Operation, _ *znet.Context, args string) (interface{}, error) {
	if args == "" {
		return nil, zerror.InvalidInput.Text("id not found")
	}

	return oper.DeleteByID(args)
}

func (h *controller) handerPut(oper *model.Operation, c *znet.Context, args string) (interface{}, error) {
	if args == "" {
		return nil, zerror.InvalidInput.Text("id not found")
	}

	j, _ := c.GetJSONs()
	data := j.Map()

	return oper.UpdateByID(args, data)
}
