package account

import (
	"errors"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
)

type requestWith struct {
	module *Module
}

const (
	ctxWithUID  = "m::account::uid"
	ctxWithUser = "m::account::user"
	// ctxWithRawUID       = "m::account::rawUID"
	ctxWithRole         = "m::account::role"
	ctxWithIsInlayAdmin = "m::account::administrator"
	// ctxWithDisabledLog  = "m::account::disabledLog"
	ctxWithLog        = "m::account::log"
	ctxWithLogRemark  = "m::account::logRemark"
	ctxWithIgnorePerm = "m::account::IgnorePerm"
	ctxWithPermCheck  = "m::account::PermCheck"
)

func (requestWith) UID(c *znet.Context) string {
	uid, ok := c.Value(ctxWithUID)
	if !ok {
		return ""
	}
	return uid.(string)
}

func (requestWith) User(c *znet.Context) ztype.Map {
	uid, ok := c.Value(ctxWithUser)
	if !ok {
		return ztype.Map{}
	}
	return uid.(ztype.Map)
}

func (r requestWith) RealUID(c *znet.Context) (string, error) {
	uid := r.UID(c)
	if uid == "" {
		return "", errors.New("uid is empty")
	}
	if r.module == nil || r.module.accountModel == nil {
		return "", errors.New("account model not initialized")
	}
	nid, err := r.module.accountModel.Schema().DeCryptID(uid)
	if err != nil || nid == "" {
		if err == nil {
			err = errors.New("uid decrypt failed")
		}
		return "", err
	}
	return nid, nil
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
