package template

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/zutil"
)

func (m *Module) Render(template string, data interface{}, layout ...string) ([]byte, error) {
	b := zutil.GetBuff()
	defer zutil.PutBuff(b)

	if err := m.engine.Render(b, template, data, layout...); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (m *Module) Template(c *znet.Context, template string, data interface{}, layout ...string) (err error) {
	r, err := m.Render(template, data, layout...)
	if err != nil {
		return err
	}
	c.HTML(200, zstring.Bytes2String(r))
	return
}
