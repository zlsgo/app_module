package account

import (
	"errors"

	"github.com/zlsgo/app_module/account/jwt"
	"github.com/zlsgo/app_module/account/rbac"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/restapi"
)

var (
	verifyPermissions func(c *znet.Context) error
)

func (m *Module) RegMiddleware(r *znet.Engine, ignore ...string) error {
	if verifyPermissions == nil {
		return errors.New("middleware not initialized, please call Init first")
	}

	if len(ignore) > 0 {
		r.Use(func(c *znet.Context) error {
			for _, v := range ignore {
				if zstring.Match(c.Request.URL.Path, v) {
					c.Next()
					return nil
				}
			}
			return verifyPermissions(c)
		})
		return nil
	}

	r.Use(verifyPermissions)
	return nil
}

func (m *Module) initMiddleware(permission *rbac.RBAC) error {
	permissionDenied := zerror.WrapTag(zerror.PermissionDenied)

	userModel, ok := m.mods.Get(accountName)
	if !ok {
		return errors.New(accountName + " accoutModel not found")
	}

	logModel, ok := m.mods.Get(logsName)
	if !ok {
		return errors.New(logsName + " logsName not found")
	}

	roleModel, ok := m.mods.Get(roleName)
	if !ok {
		return errors.New(roleName + " roleName not found")
	}

	permModel, ok := m.mods.Get(permName)
	if !ok {
		return errors.New(permName + " permName not found")
	}

	// TODO: 可能需要独立出来方便做缓存
	roles, err := restapi.Find(roleModel, ztype.Map{
		"status": 1,
	})
	if err != nil {
		return err
	}

	// 添加权限规则
	for _, r := range roles {
		role := rbac.NewRole(rbac.MatchPriorityDeny)
		perms, err := restapi.Find(permModel, ztype.Map{
			restapi.IDKey: r.Get("permission").SliceInt(),
			"status":      1,
		}, func(o *restapi.CondOptions) error {
			o.Fields = []string{"action", "alias", "target", "priority"}
			return nil
		})
		if err != nil {
			return err
		}

		for _, perm := range perms {
			role.AddGlobPermission(perm.Get("priority").Int(), perm.Get("action").String(), perm.Get("target").String())
		}
		err = permission.MergerRole(r.Get("name").String(), role)
		if err != nil {
			return err
		}
	}

	verifyPermissions = func(c *znet.Context) error {
		token := jwt.GetToken(c)
		if token == "" {
			return zerror.WrapTag(zerror.Unauthorized)(errors.New("无法访问，请先登录"))
		}

		if userModel == nil {
			return errors.New(accountName + " accoutModel not found")
		}

		uid, err := getJWTForCache(userModel, token, m.Options.key)
		if err != nil {
			return err
		}

		c.WithValue(ctxWithUID, uid)

		u, err := getUserForCache(userModel, uid)
		if err != nil {
			return err
		}

		if u.Get("status").Int() != 1 {
			return permissionDenied(errors.New("用户已被禁用"))
		}

		logRequest(c, logModel, u)

		isInlayAdmin := u.Get("administrator").Bool()
		c.WithValue(ctxWithIsInlayAdmin, isInlayAdmin)
		if isInlayAdmin {
			return nil
		}

		roles := u.Get("role").SliceString()
		c.WithValue(ctxWithRole, roles)
		for _, r := range roles {
			isAllow, _ := permission.Can(r, c.Request.Method, c.Request.URL.Path)
			if isAllow {
				return nil
			}
		}

		// 是否忽略权限限制
		if b, ok := c.Value(ctxWithIgnorePerm); ok && b.(bool) {
			return nil
		}

		return permissionDenied(errors.New("无权访问"))
	}
	return nil
}
