package account

import (
	"reflect"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
)

type Message struct {
	service.App
	plugin *Module
	Path   string
}

var (
	_ = reflect.TypeOf(&Message{})
)

func (h *Message) Init(r *znet.Engine) error {
	return h.plugin.RegMiddleware(r)
}

// Get 站内通知列表
func (h *Message) Get(c *znet.Context) (data ztype.Map, err error) {
	uid := Ctx.UID(c)
	unread, _ := h.plugin.messageModel.Unread(uid)

	return ztype.Map{
		"unread": unread,
	}, err
}
