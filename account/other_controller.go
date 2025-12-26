package account

import (
	"errors"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
)

// getSite 系统信息
func (h *Index) getSite(c *znet.Context) (data ztype.Map, err error) {
	return ztype.Map{}, err
}

// GetMessage 站内消息
func (h *Index) GetMessage(c *znet.Context) (data ztype.Map, err error) {
	uid := h.module.Request.UID(c)
	if h.module.messageModel == nil {
		return nil, errors.New("message model not define")
	}
	unread, _ := h.module.messageModel.Unread(uid)

	return ztype.Map{
		"unread": unread,
	}, err
}
