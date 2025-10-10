package account

import (
	"errors"
	"time"

	"github.com/zlsgo/app_module/account/jwt"
	"github.com/zlsgo/app_module/account/rbac"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zcli"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	zsession "github.com/sohaha/zlsgo/znet/session"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
)

var (
	verifyPermissions = []znet.Handler{}
	ignoreRoutes      = []string{}
)

func UsePermisMiddleware(r *znet.Engine, authErrHandler func(c *znet.Context, err error) error, ignore ...string) error {
	if verifyPermissions == nil {
		return errors.New("middleware not initialized, please call Init first")
	}

	if len(ignore) > 0 {
		r.Use(func(c *znet.Context) error {
			c.WithValue(ctxWithPermCheck, authErrHandler)
			for _, v := range ignore {
				if zstring.Match(c.Request.URL.Path, v) {
					c.WithValue(ctxWithIgnorePerm, true)
					break
				}
			}
			c.Next()
			return nil
		})
	}

	r.Use(verifyPermissions...)

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

	// 使用缓存获取角色权限
	roles, err := getRolesForCache(roleModel)
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

	verifyPermissions = []znet.Handler{}

	if m.Options.Session != nil {
		s := zsession.New(m.Options.Session, func(c *zsession.Config) {
			c.ExpiresAt = time.Duration(m.Options.Expire) * time.Second
		})

		go func() {
			timer := time.NewTicker(time.Duration(m.Options.Expire) * time.Second)
			defer timer.Stop()
			for {
				select {
				case <-timer.C:
					m.Options.Session.Collect()
				case <-zcli.SingleKillSignal():
					return
				}
			}
		}()

		verifyPermissions = append(verifyPermissions, s)
	}

	verifyPermissions = append(verifyPermissions, func(c *znet.Context) error {
		var ignorePerm bool
		if b, ok := c.Value(ctxWithIgnorePerm); ok && b.(bool) {
			ignorePerm = b.(bool)
		}

		token := jwt.GetToken(c)
		if token == "" {
			if m.Options.Session != nil {
				s, err := zsession.Get(c)
				if err == nil {
					token = s.Get("token").String()
				}
			}
			if token == "" {
				if ignorePerm {
					return nil
				}

				err = zerror.WrapTag(zerror.Unauthorized)(errors.New("无法访问，请先登录"))
				if b, ok := c.Value(ctxWithPermCheck); ok && b != nil {
					err = b.(func(c *znet.Context, err error) error)(c, err)
				}

				return err
			}
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

		// 管理员直接通过，避免权限检查
		if isInlayAdmin {
			c.Next()
			return nil
		}

		// 管理员直接通过
		if isInlayAdmin {
			c.Next()
			return nil
		}

		// 检查公共路由
		if zarray.Contains(publicRoutes, c.Request.URL.Path) {
			c.Next()
			return nil
		}

		// 检查是否忽略权限限制
		if ignorePerm {
			c.Next()
			return nil
		}

		// 检查权限
		for _, r := range roles {
			isAllow, _ := permission.Can(r, c.Request.Method, c.Request.URL.Path)
			if isAllow {
				c.Next()
				return nil
			}
		}

		err = permissionDenied(errors.New("没有访问权限"))
		if b, ok := c.Value(ctxWithPermCheck); ok && b != nil {
			err = b.(func(c *znet.Context, err error) error)(c, err)
		}

		return err
	},
		func(c *znet.Context) error {
			c.Next()

			u, ok := c.Value(ctxWithUser)
			if !ok {
				return nil
			}

			logRequest(c, logModel, u.(ztype.Map))
			return nil
		})

	return nil
}
