package template

import (
	"errors"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/zutil"
)

// ErrNotReady 表示模板模块尚未启动完成，引擎不可用。
var ErrNotReady = errors.New("template 模块未启动，无法渲染")

func (m *Module) Render(template string, data interface{}, layout ...string) ([]byte, error) {
	if m == nil || m.engine == nil {
		return nil, ErrNotReady
	}

	b := zutil.GetBuff()
	defer zutil.PutBuff(b)

	if err := m.engine.Render(b, template, data, layout...); err != nil {
		return nil, err
	}

	// PutBuff 会 Reset 缓冲并放回池，返回的切片必须脱离池化数组，
	// 否则下一次 GetBuff 写入会改写已返回的内容。
	return append([]byte(nil), b.Bytes()...), nil
}

func (m *Module) Template(c *znet.Context, template string, data interface{}, layout ...string) (err error) {
	r, err := m.Render(template, data, layout...)
	if err != nil {
		return err
	}
	c.HTML(200, zstring.Bytes2String(r))
	return
}
