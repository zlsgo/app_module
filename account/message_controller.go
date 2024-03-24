package account

import (
	"reflect"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
)

type Message struct {
	service.App
	module *Module
	Path   string
}

var (
	_ = reflect.TypeOf(&Message{})
)

func (h *Message) Init(r *znet.Engine) error {
	return PermisMiddleware(r)
}

// Get 站内通知列表
func (h *Message) Get(c *znet.Context) (data ztype.Map, err error) {
	uid := Request.UID(c)
	m, _ := GetMessageModel()
	unread, _ := m.Unread(uid)

	return ztype.Map{
		"unread": unread,
	}, err
}

// AnyRealtime 实时通知
func (h *Message) AnyRealtime(c *znet.Context) error {
	if !c.IsSSE() {
		return zerror.InvalidInput.Text("不支持的请求类型")
	}

	sse, remove, err := h.module.newSession(c)
	if err != nil {
		return err
	}
	defer remove()

	sse.Push()
	return nil
}
