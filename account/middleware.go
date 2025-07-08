package account

import (
	"errors"

	"github.com/zlsgo/app_module/account/jwt"
	"github.com/zlsgo/app_module/account/rbac"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model"
)

var verifyPermissions []func(c *znet.Context) error

func PermisMiddleware(r *znet.Engine, ignore ...string) error {
	if verifyPermissions == nil {
		return errors.New("middleware not initialized, please call Init first")
	}

	if len(ignore) > 0 {
		permissions := verifyPermissions[0]
		verifyPermissions[0] = func(c *znet.Context) error {
			for _, v := range ignore {
				if zstring.Match(c.Request.URL.Path, v) {
					c.Next()
					return nil
				}
			}
			return permissions(c)
		}
	}

	for i := range verifyPermissions {
		r.Use(verifyPermissions[i])
	}

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

	// TODO: 可能需要独立出来方便做缓存
	roles, err := model.Find(roleModel, ztype.Map{
		"status": 1,
	})
	if err != nil {
		return err
	}

	// 添加权限规则
	for _, r := range roles {
		if err := m.setPermission(permission, r); err != nil {
			return err
		}
	}

	// 全部角色通用权限
	permission.ForEachRole(func(key string, value *rbac.Role) bool {
		value.AddGlobPermission(1, "*", m.Options.ApiPrefix+"/message/realtime")
		return true
	})

	// 无需角色权限校验的接口
	publicRoutes := []string{
		m.Options.ApiPrefix + "/base/info",
		m.Options.ApiPrefix + "/base/logs",
	}

	verifyPermissions = []func(c *znet.Context) error{
		func(c *znet.Context) error {
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

			c.WithValue(ctxWithUser, u)

			if u.Get("status").Int() != 1 {
				return permissionDenied(errors.New("用户已被禁用"))
			}

			isInlayAdmin := u.Get("administrator").Bool()
			c.WithValue(ctxWithIsInlayAdmin, isInlayAdmin)

			roles := u.Get("role").SliceString()
			c.WithValue(ctxWithRole, roles)

			if isInlayAdmin {
				return nil
			}

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

			if zarray.Contains(publicRoutes, c.Request.URL.Path) {
				return nil
			}

			return permissionDenied(errors.New("没有访问权限"))
		},
		func(c *znet.Context) error {
			c.Next()

			u, ok := c.Value(ctxWithUser)
			if !ok {
				return nil
			}
			logRequest(c, logModel, u.(ztype.Map))
			return nil
		},
	}
	return nil
}
