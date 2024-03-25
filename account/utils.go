package account

import (
	"github.com/sohaha/zlsgo/znet"
)

type requestWith struct {
}

const (
	ctxWithUID = "m::account::uid"
	// ctxWithRawUID       = "m::account::rawUID"
	ctxWithRole         = "m::account::role"
	ctxWithIsInlayAdmin = "m::account::administrator"
	// ctxWithDisabledLog  = "m::account::disabledLog"
	ctxWithLog        = "m::account::log"
	ctxWithLogRemark  = "m::account::logRemark"
	ctxWithIgnorePerm = "m::account::IgnorePerm"
)

var Request = &requestWith{}

func (requestWith) UID(c *znet.Context) string {
	uid, ok := c.Value(ctxWithUID)
	if !ok {
		return ""
	}
	return uid.(string)
}

func (r requestWith) RealUID(c *znet.Context) string {
	uid := r.UID(c)
	nid, _ := GetAccountModel().DeCryptID(uid)
	return nid
}

func (requestWith) Roles(c *znet.Context) []string {
	roles, ok := c.Value(ctxWithRole)
	if !ok {
		return []string{}
	}
	return roles.([]string)
}

func (requestWith) IsSuperAdmin(c *znet.Context) bool {
	b, ok := c.Value(ctxWithIsInlayAdmin)
	if !ok {
		return false
	}
	return b.(bool)
}

func (requestWith) IgnorePerm(c *znet.Context) *znet.Context {
	return c.WithValue(ctxWithIgnorePerm, true)
}

func (requestWith) WithLog(c *znet.Context, message string, remark ...string) *znet.Context {
	lastMsg := c.MustValue(ctxWithLog, "").(string)
	if lastMsg != "" {
		message = lastMsg + ": " + message
	}
	c.WithValue(ctxWithLog, message)
	if len(remark) > 0 {
		c.WithValue(ctxWithLogRemark, remark[0])
	}
	return c
}
