package account

import (
	"errors"
	"time"

	"github.com/zlsgo/app_module/account/jwt"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zcli"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	zsession "github.com/sohaha/zlsgo/znet/session"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
)

func (m *Module) UsePermisMiddleware(r *znet.Engine, authErrHandler func(c *znet.Context, err error) error, ignore ...string) error {
	if m.verifyPermissions == nil {
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

	r.Use(m.verifyPermissions...)

	return nil
}

func (m *Module) initMiddleware() error {
	permissionDenied := zerror.WrapTag(zerror.PermissionDenied)

	userModel, ok := m.mods.Get(accountName)
	if !ok {
		return errors.New(accountName + " accoutModel not found")
	}

	logModel, ok := m.mods.Get(logsName)
	if !ok {
		return errors.New(logsName + " logsName not found")
	}

	// 无需角色权限校验的接口
	publicRoutes := []string{
		m.Options.ApiPrefix + "/base/info",
		m.Options.ApiPrefix + "/base/logs",
	}

	m.verifyPermissions = []znet.Handler{}

	if m.Options.Session != nil {
		expireDuration := time.Duration(m.Options.Expire) * time.Second
		if expireDuration <= 0 {
			expireDuration = time.Hour * 24
		}

		s := zsession.New(m.Options.Session, func(c *zsession.Config) {
			c.ExpiresAt = expireDuration
		})

		go func() {
			timer := time.NewTicker(expireDuration)
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

		m.verifyPermissions = append(m.verifyPermissions, s)
	}

	unauthorized := zerror.WrapTag(zerror.Unauthorized)(errors.New("无法访问，请先登录"))
	m.verifyPermissions = append(m.verifyPermissions, func(c *znet.Context) error {
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

				authErr := unauthorized
				if b, ok := c.Value(ctxWithPermCheck); ok && b != nil {
					r, ok := b.(func(c *znet.Context, err error) error)
					if ok && r != nil {
						authErr = zerror.TryCatch(func() error {
							return r(c, authErr)
						})
					}
				}

				return authErr
			}
		}

		if userModel == nil {
			return errors.New(accountName + " accoutModel not found")
		}

		uid, err := m.getJWTForCache(userModel, token, m.Options.key)
		if err != nil {
			return err
		}

		c.WithValue(ctxWithUID, uid)

		u, err := m.getUserForCache(userModel, uid)
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
		permission := m.permission.Load()
		if permission == nil {
			return errors.New("rbac not initialized")
		}
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

			m.logRequest(c, logModel, u.(ztype.Map))
			return nil
		})

	return nil
}
