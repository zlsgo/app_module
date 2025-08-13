package html

import (
	"reflect"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zreflect"
	"github.com/zlsgo/app_module/html/elements"
)

type (
	invoker func(c *znet.Context) elements.ElementRenderer
)

var invokerValue zdi.PreInvoker = (invoker)(nil)

func (h invoker) Invoke(v []interface{}) ([]reflect.Value, error) {
	c := v[0].(*znet.Context)
	resp := h(c)
	if resp == nil {
		return []reflect.Value{zreflect.ValueOf("")}, nil
	}

	w := c.GetWriter(200)
	if err := resp.Render(w); err != nil {
		return nil, err
	}

	return []reflect.Value{}, nil
}

type (
	invokerCode func(c *znet.Context) (int, elements.ElementRenderer)
)

var invokerCodeValue zdi.PreInvoker = (invokerCode)(nil)

func (h invokerCode) Invoke(v []interface{}) ([]reflect.Value, error) {
	c := v[0].(*znet.Context)
	code, resp := h(c)
	if resp == nil {
		return []reflect.Value{zreflect.ValueOf("")}, nil
	}

	w := c.GetWriter(int32(code))
	if err := resp.Render(w); err != nil {
		return nil, err
	}

	return []reflect.Value{}, nil
}

type (
	invokerZ func(c *znet.Context, x *ZHTML) elements.ElementRenderer
)

var invokerZValue zdi.PreInvoker = (invokerZ)(nil)

func (h invokerZ) Invoke(v []interface{}) ([]reflect.Value, error) {
	c := v[0].(*znet.Context)
	z := v[1].(*ZHTML)
	resp := h(c, z)
	if resp == nil {
		return []reflect.Value{zreflect.ValueOf("")}, nil
	}

	w := c.GetWriter(200)
	if err := resp.Render(w); err != nil {
		return nil, err
	}

	return []reflect.Value{}, nil
}

type (
	invokerCodeZ func(c *znet.Context, x *ZHTML) (int, elements.ElementRenderer)
)

var invokerCodeZValue zdi.PreInvoker = (invokerCodeZ)(nil)

func (h invokerCodeZ) Invoke(v []interface{}) ([]reflect.Value, error) {
	c := v[0].(*znet.Context)
	z := v[1].(*ZHTML)
	code, resp := h(c, z)
	if resp == nil {
		return []reflect.Value{zreflect.ValueOf("")}, nil
	}

	w := c.GetWriter(int32(code))
	if err := resp.Render(w); err != nil {
		return nil, err
	}

	return []reflect.Value{}, nil
}

func init() {
	znet.RegisterRender(invokerValue)
	znet.RegisterRender(invokerCodeValue)
	znet.RegisterRender(invokerZValue)
	znet.RegisterRender(invokerCodeZValue)
}
