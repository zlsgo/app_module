package account

import "github.com/sohaha/zlsgo/znet"

func ParseUID(c *znet.Context) string {
	uid, ok := c.Value(contextWithUID)
	if !ok {
		return ""
	}
	return uid.(string)
}

func ParseRoles(c *znet.Context) []string {
	roles, ok := c.Value(contextWithRole)
	if !ok {
		return []string{}
	}
	return roles.([]string)
}

func IsInlayAdmin(c *znet.Context) bool {
	b, ok := c.Value(contextWithIsInlayAdmin)
	if !ok {
		return false
	}
	return b.(bool)
}
