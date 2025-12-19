package html

import (
	"net/http"
	"reflect"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_module/html/el"
)

type (
	invoker func(c *znet.Context) *el.Element
)

var invokerValue zdi.PreInvoker = (invoker)(nil)

func (h invoker) Invoke(v []interface{}) ([]reflect.Value, error) {
	c := v[0].(*znet.Context)
	resp := h(c)
	if resp == nil {
		return []reflect.Value{}, nil
	}

	html, err := el.RenderBytes(c.Request.Context(), resp)
	statusCode := http.StatusOK
	if err != nil {
		if options.ErrorPage == nil {
			return nil, err
		}
		html, err = el.RenderBytes(c.Request.Context(), options.ErrorPage)
		if err != nil {
			return nil, err
		}
		statusCode = http.StatusInternalServerError
	}

	c.SetContentType(znet.ContentTypeHTML)
	c.Byte(int32(statusCode), html)
	return []reflect.Value{}, nil
}

type (
	invokerCode func(c *znet.Context) (*el.Element, error)
)

var invokerCodeValue zdi.PreInvoker = (invokerCode)(nil)

func (h invokerCode) Invoke(v []interface{}) ([]reflect.Value, error) {
	c := v[0].(*znet.Context)
	resp, err := h(c)
	if err != nil {
		return []reflect.Value{}, err
	}
	if resp == nil {
		return []reflect.Value{}, nil
	}

	html, err := el.RenderBytes(c.Request.Context(), resp)
	statusCode := http.StatusOK
	if err != nil {
		if options.ErrorPage == nil {
			return nil, err
		}
		html, err = el.RenderBytes(c.Request.Context(), options.ErrorPage)
		if err != nil {
			return nil, err
		}
		statusCode = http.StatusInternalServerError
	}

	c.SetContentType(znet.ContentTypeHTML)
	c.Byte(int32(statusCode), html)
	return []reflect.Value{}, nil
}

type (
	invokerError func(c *znet.Context) (int, *el.Element)
)

var invokerErrorValue zdi.PreInvoker = (invokerError)(nil)

func (h invokerError) Invoke(v []interface{}) ([]reflect.Value, error) {
	c := v[0].(*znet.Context)
	code, resp := h(c)
	if resp == nil {
		return []reflect.Value{}, nil
	}

	html, err := el.RenderBytes(c.Request.Context(), resp)
	if err != nil {
		if options.ErrorPage == nil {
			return nil, err
		}
		html, err = el.RenderBytes(c.Request.Context(), options.ErrorPage)
		if err != nil {
			return nil, err
		}
		code = http.StatusInternalServerError
	}

	c.SetContentType(znet.ContentTypeHTML)
	c.Byte(int32(code), html)
	return []reflect.Value{}, nil
}

type (
	invokerZ func(c *znet.Context, x *ZHTML) *el.Element
)

var invokerZValue zdi.PreInvoker = (invokerZ)(nil)

func (h invokerZ) Invoke(v []interface{}) ([]reflect.Value, error) {
	c := v[0].(*znet.Context)
	z := v[1].(*ZHTML)
	resp := h(c, z)
	if resp == nil {
		return []reflect.Value{}, nil
	}

	html, err := el.RenderBytes(c.Request.Context(), resp)
	statusCode := http.StatusOK
	if err != nil {
		if options.ErrorPage == nil {
			return nil, err
		}
		html, err = el.RenderBytes(c.Request.Context(), options.ErrorPage)
		if err != nil {
			return nil, err
		}
		statusCode = http.StatusInternalServerError
	}

	c.SetContentType(znet.ContentTypeHTML)
	c.Byte(int32(statusCode), html)
	return []reflect.Value{}, nil
}

type (
	invokerCodeZ func(c *znet.Context, x *ZHTML) (int, *el.Element)
)

var invokerCodeZValue zdi.PreInvoker = (invokerCodeZ)(nil)

func (h invokerCodeZ) Invoke(v []interface{}) ([]reflect.Value, error) {
	c := v[0].(*znet.Context)
	z := v[1].(*ZHTML)
	code, resp := h(c, z)
	if resp == nil {
		return []reflect.Value{}, nil
	}

	html, err := el.RenderBytes(c.Request.Context(), resp)
	if err != nil {
		if options.ErrorPage == nil {
			return nil, err
		}
		html, err = el.RenderBytes(c.Request.Context(), options.ErrorPage)
		if err != nil {
			return nil, err
		}
		code = http.StatusInternalServerError
	}

	c.SetContentType(znet.ContentTypeHTML)
	c.Byte(int32(code), html)

	return []reflect.Value{}, nil
}

func init() {
	znet.RegisterRender(invokerValue)
	znet.RegisterRender(invokerCodeValue)
	znet.RegisterRender(invokerErrorValue)
	znet.RegisterRender(invokerZValue)
	znet.RegisterRender(invokerCodeZValue)
}
