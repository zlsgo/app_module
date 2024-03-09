package account

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
)

// getSite 系统信息
func (h *Index) getSite(c *znet.Context) (data ztype.Map, err error) {
	return ztype.Map{}, err
}

// GetMessage 站内消息
func (h *Index) GetMessage(c *znet.Context) (data ztype.Map, err error) {
	return ztype.Map{
		"unread": 110,
	}, err
}
