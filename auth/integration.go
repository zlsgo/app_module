package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	zsession "github.com/sohaha/zlsgo/znet/session"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model"
)

// CurrentUser 暴露当前 auth 登录用户，供其他模块在同一请求上下文里复用 auth 会话态。
func (m *Module) CurrentUser(c *znet.Context) (ztype.Map, error) {
	user, _, err := m.currentUser(c)
	if err != nil {
		if zerror.GetTag(err) == zerror.Unauthorized || looksLikeAuthUnauthorized(err) {
			return nil, zerror.Unauthorized.Text(err.Error())
		}
		return nil, err
	}
	return user, nil
}

// UserByID 暴露 auth 用户读取，供兼容模块做显式身份校验。
func (m *Module) UserByID(id string) (ztype.Map, error) {
	user, err := m.userByID(id)
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			return nil, zerror.NotFound.Text("auth user not found")
		}
		return nil, err
	}
	return user, nil
}

// ValidateSessionBinding 验证 session 是否仍归属于目标用户且处于有效状态。
func (m *Module) ValidateSessionBinding(userID, sessionKey string, sessionVersion int) error {
	if strings.TrimSpace(userID) == "" || strings.TrimSpace(sessionKey) == "" {
		return zerror.Unauthorized.Text("auth session binding invalid")
	}

	user, err := m.userByID(userID)
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			return zerror.Unauthorized.Text("auth user not found")
		}
		return err
	}
	if user.Get("status").Int() != 1 {
		return zerror.Unauthorized.Text("auth user disabled")
	}
	if user.Get("session_version").Int() != sessionVersion {
		return zerror.Unauthorized.Text("auth session version changed")
	}

	sessionRow, err := m.sessionByKey(sessionKey)
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			return zerror.Unauthorized.Text("auth session invalid")
		}
		return err
	}
	if sessionRow.Get("status").Int() != 1 {
		return zerror.Unauthorized.Text("auth session invalid")
	}

	rawUserID, err := m.rawUserID(userID)
	if err != nil {
		return err
	}
	if sessionRow.Get("user_id").String() != rawUserID {
		return zerror.Unauthorized.Text("auth session user mismatch")
	}

	if expireAt := sessionRow.Get("expire_at").String(); expireAt != "" {
		if exp, parseErr := parseDBTimeInSession(expireAt); parseErr == nil && exp.Before(time.Now()) {
			return zerror.Unauthorized.Text("auth session expired")
		}
	}

	return nil
}

// SessionMiddleware 返回与 auth 控制器一致的 session 中间件，供其他模块挂载同一套会话解析逻辑。
func (m *Module) SessionMiddleware() znet.Handler {
	return zsession.New(m.Options.Session, func(c *zsession.Config) {
		c.CookieName = m.Options.CookieName
		c.ExpiresAt = time.Duration(m.Options.SessionTTL) * time.Second
		c.AutoRenew = true
	})
}

func looksLikeAuthUnauthorized(err error) bool {
	if err == nil {
		return false
	}

	msg := err.Error()
	for _, pattern := range []string{
		"未登录",
		"登录态已失效",
		"登录态已过期",
		"用户不存在",
		"账号已禁用",
	} {
		if strings.Contains(msg, pattern) {
			return true
		}
	}

	return false
}
