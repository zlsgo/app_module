package html

import (
	"net/http"
	"reflect"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_module/html/el"
	"github.com/zlsgo/app_module/html/zview"
)

type (
	invoker func(c *znet.Context) *el.Element
)

var invokerValue zdi.PreInvoker = (invoker)(nil)

func (h invoker) Invoke(v []any) ([]reflect.Value, error) {
	c := v[0].(*znet.Context)
	return []reflect.Value{}, writePage(c, http.StatusOK, h(c), nil)
}

type (
	invokerCode func(c *znet.Context) (*el.Element, error)
)

var invokerCodeValue zdi.PreInvoker = (invokerCode)(nil)

func (h invokerCode) Invoke(v []any) ([]reflect.Value, error) {
	c := v[0].(*znet.Context)
	resp, err := h(c)
	if err != nil {
		return []reflect.Value{}, writePage(c, http.StatusInternalServerError, nil, err)
	}
	return []reflect.Value{}, writePage(c, http.StatusOK, resp, nil)
}

type (
	invokerError func(c *znet.Context) (int, *el.Element)
)

var invokerErrorValue zdi.PreInvoker = (invokerError)(nil)

func (h invokerError) Invoke(v []any) ([]reflect.Value, error) {
	c := v[0].(*znet.Context)
	code, resp := h(c)
	return []reflect.Value{}, writePage(c, code, resp, nil)
}

type (
	invokerZ func(c *znet.Context, x *zview.Context) *el.Element
)

var invokerZValue zdi.PreInvoker = (invokerZ)(nil)

func (h invokerZ) Invoke(v []any) ([]reflect.Value, error) {
	c := v[0].(*znet.Context)
	z := v[1].(*zview.Context)
	return []reflect.Value{}, writePage(c, http.StatusOK, h(c, z), nil)
}

type (
	invokerCodeZ func(c *znet.Context, x *zview.Context) (int, *el.Element)
)

var invokerCodeZValue zdi.PreInvoker = (invokerCodeZ)(nil)

func (h invokerCodeZ) Invoke(v []any) ([]reflect.Value, error) {
	c := v[0].(*znet.Context)
	z := v[1].(*zview.Context)
	code, resp := h(c, z)
	return []reflect.Value{}, writePage(c, code, resp, nil)
}

func optionsFor(c *znet.Context) Options {
	if c == nil || c.Engine == nil {
		return Options{}
	}
	if value, ok := moduleOptions.Load(c.Engine); ok {
		return value.(Options)
	}
	return Options{}
}

func writePage(c *znet.Context, code int, resp *el.Element, cause error) error {
	if resp == nil && cause == nil {
		return nil
	}

	opts := optionsFor(c)
	page := resp
	statusCode := code
	if cause != nil {
		if opts.ErrorPage == nil {
			return cause
		}
		page = opts.ErrorPage
		statusCode = http.StatusInternalServerError
	}

	html, err := el.RenderBytes(c.Request.Context(), page)
	if err != nil {
		if opts.ErrorPage == nil || page == opts.ErrorPage {
			return err
		}
		html, err = el.RenderBytes(c.Request.Context(), opts.ErrorPage)
		if err != nil {
			return err
		}
		statusCode = http.StatusInternalServerError
	}

	c.SetContentType(znet.ContentTypeHTML)
	c.Byte(int32(statusCode), html)
	return nil
}

func init() {
	znet.RegisterRender(invokerValue)
	znet.RegisterRender(invokerCodeValue)
	znet.RegisterRender(invokerErrorValue)
	znet.RegisterRender(invokerZValue)
	znet.RegisterRender(invokerCodeZValue)
}
