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

var _ = reflect.TypeOf(&Message{})

func (h *Message) Init(r *znet.Engine) error {
	return h.module.UsePermisMiddleware(r, nil)
}

// Get 站内通知列表
func (h *Message) Get(c *znet.Context) (data ztype.Map, err error) {
	uid := h.module.Request.UID(c)
	if h.module.messageModel == nil {
		return nil, zerror.InvalidInput.Text("消息模型未初始化")
	}
	unread, _ := h.module.messageModel.Unread(uid)

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
