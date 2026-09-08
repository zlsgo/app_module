package auth

import (
	"errors"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/sohaha/zlsgo/znet"
	zsession "github.com/sohaha/zlsgo/znet/session"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model"
)

type OAuthController struct {
	service.App
	module *Module
	Path   string
}

var _ = reflect.TypeOf(&OAuthController{})

func (h *OAuthController) Init(r *znet.Engine) error {
	if err := h.DI.Resolve(&h.module); err != nil {
		return err
	}

	r.Use(znet.RewriteErrorHandler(authErrorHandler))
	r.Use(zsession.New(h.module.Options.Session, func(c *zsession.Config) {
		c.CookieName = h.module.Options.CookieName
		c.ExpiresAt = time.Duration(h.module.Options.SessionTTL) * time.Second
		c.AutoRenew = true
	}))
	for _, provider := range h.module.EnabledProviders() {
		if err := provider.Init(r.Group(provider.Name())); err != nil {
			return err
		}
	}

	r.POST("/add", h.add)
	r.POST("/remove", h.remove)
	r.Any("/login/:provider", h.login)
	r.Any("/callback/:provider", h.callback)

	return nil
}

func (h *OAuthController) add(c *znet.Context) (any, error) {
	user, _, err := h.module.currentUser(c)
	if err != nil {
		return nil, toHTTPError(err)
	}

	input, err := getPayload(c)
	if err != nil {
		return nil, toHTTPError(err)
	}

	provider, err := h.module.findProvider(input.Get("provider").String())
	if err != nil {
		return nil, toHTTPError(err)
	}

	if _, err = h.module.beginOAuthFlow(c, oauthFlowActionBind, provider.Name(), user.Get(model.IDKey()).String()); err != nil {
		return nil, toHTTPError(err)
	}
	if err = provider.Login(c); err != nil {
		return nil, toHTTPError(err)
	}
	if handledOAuthLogin(c) {
		return nil, nil
	}

	return ztype.Map{
		"provider": provider.Name(),
		"state":    c.Request.URL.Query().Get("state"),
	}, nil
}

func (h *OAuthController) remove(c *znet.Context) (any, error) {
	user, _, err := h.module.currentUser(c)
	if err != nil {
		return nil, toHTTPError(err)
	}

	input, err := getPayload(c)
	if err != nil {
		return nil, toHTTPError(err)
	}

	provider := strings.TrimSpace(input.Get("provider").String())
	providerID := strings.TrimSpace(input.Get("provider_id").String())
	if provider == "" || providerID == "" {
		return nil, toHTTPError(ErrProviderInfoIncomplete)
	}

	removed, err := h.module.unlinkProvider(user.Get(model.IDKey()).String(), provider, providerID)
	if err != nil {
		return nil, toHTTPError(err)
	}
	if !removed {
		return nil, toHTTPError(ErrBindingNotFound)
	}

	return ztype.Map{"removed": true}, nil
}

func (h *OAuthController) login(c *znet.Context) (any, error) {
	return h.module.OAuthLogin(c, c.GetParam("provider"))
}

func (h *OAuthController) callback(c *znet.Context) (any, error) {
	return h.module.OAuthCallback(c, c.GetParam("provider"))
}

func (m *Module) OAuthLogin(c *znet.Context, providerName string) (any, error) {
	provider, err := m.findProvider(providerName)
	if err != nil {
		return nil, toHTTPError(err)
	}

	if _, err = m.beginOAuthFlow(c, oauthFlowActionLogin, provider.Name(), ""); err != nil {
		return nil, toHTTPError(err)
	}
	if err = provider.Login(c); err != nil {
		return nil, toHTTPError(err)
	}
	if handledOAuthLogin(c) {
		return nil, nil
	}

	return ztype.Map{
		"provider": provider.Name(),
		"state":    c.Request.URL.Query().Get("state"),
	}, nil
}

func (m *Module) OAuthCallback(c *znet.Context, providerName string) (ztype.Map, error) {
	provider, err := m.findProvider(providerName)
	if err != nil {
		return nil, toHTTPError(err)
	}

	flow, flowSession, err := m.validateOAuthFlow(c, provider.Name())
	if err != nil {
		return nil, toHTTPError(err)
	}

	info, err := provider.Callback(c)
	if err != nil {
		return nil, toHTTPError(err)
	}
	if info.Provider == "" {
		info.Provider = provider.Name()
	}
	if info.Provider == "" || info.ProviderID == "" {
		return nil, toHTTPError(ErrProviderInfoIncomplete)
	}
	defer func() {
		_ = m.clearOAuthFlow(flowSession)
	}()

	if flow.Action == oauthFlowActionBind {
		sessionUser, _, err := m.currentUser(c)
		if err != nil {
			return nil, toHTTPError(err)
		}
		if sessionUser.Get(model.IDKey()).String() != flow.UserID {
			return nil, toHTTPError(ErrOAuthUserMismatch)
		}

		row, err := m.linkProviderToUser(sessionUser.Get(model.IDKey()).String(), info)
		if err != nil {
			return nil, toHTTPError(err)
		}
		return ztype.Map{
			"linked":   true,
			"provider": row,
			"user":     sanitizeUser(sessionUser),
		}, nil
	}

	user, err := m.resolveOAuthUser(info)
	if err != nil {
		return nil, toHTTPError(err)
	}

	session, err := zsession.Get(c)
	if err != nil {
		return nil, toHTTPError(err)
	}

	if err = m.attachSession(user, session); err != nil {
		return nil, toHTTPError(err)
	}

	users, _ := m.UserModel()
	_, err = users.UpdateByID(user.Get(model.IDKey()).String(), ztype.Map{
		"login_at": ztime.Now(),
	})
	if err != nil {
		return nil, toHTTPError(err)
	}

	user, err = users.FindOneByID(user.Get(model.IDKey()).String())
	if err != nil {
		return nil, toHTTPError(err)
	}

	return sanitizeUser(user), nil
}

func (m *Module) findProvider(name string) (AuthProvider, error) {
	name = strings.TrimSpace(name)
	for _, provider := range m.EnabledProviders() {
		if provider.Name() == name {
			return provider, nil
		}
	}

	return nil, ErrProviderNotEnabled
}

func (m *Module) resolveOAuthUser(info Provider) (ztype.Map, error) {
	if row, err := m.providerByKey(info.Provider, info.ProviderID); err == nil {
		return m.userByRawID(row.Get("user_id").String())
	} else if !errors.Is(err, model.ErrNoRecord) {
		return nil, err
	}

	if info.ProviderEmailVerified && isValidEmail(info.ProviderEmail) {
		user, err := m.userByEmail(normalizeEmail(info.ProviderEmail))
		if err == nil {
			_, err = m.linkProviderToUser(user.Get(model.IDKey()).String(), info)
			if err != nil {
				return nil, err
			}
			return user, nil
		}
		if !errors.Is(err, model.ErrNoRecord) {
			return nil, err
		}
	}

	return m.createOAuthUser(info)
}

func (m *Module) createOAuthUser(info Provider) (ztype.Map, error) {
	users, ok := m.UserModel()
	if !ok {
		return nil, errors.New("user model not ready")
	}

	email := normalizeEmail(info.ProviderEmail)
	if !info.ProviderEmailVerified || !isValidEmail(email) {
		email = oauthPlaceholderEmail(info.Provider, info.ProviderID)
	}

	nickname := strings.TrimSpace(info.ProviderUsername)
	if nickname == "" {
		nickname = info.Provider + "_" + info.ProviderID
	}

	id, err := users.Insert(ztype.Map{
		"email":           email,
		"password":        nil,
		"nickname":        nickname,
		"avatar":          info.ProviderAvatar,
		"status":          1,
		"settings":        "{}",
		"login_at":        ztime.Now(),
		"session_version": 1,
	})
	if err != nil {
		return nil, err
	}

	user, err := users.FindOneByID(id)
	if err != nil {
		return nil, err
	}

	_, err = m.linkProviderToUser(user.Get(model.IDKey()).String(), info)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Module) linkProviderToUser(userID string, info Provider) (ztype.Map, error) {
	if info.Provider == "" || info.ProviderID == "" {
		return nil, ErrProviderInfoIncomplete
	}

	providers, ok := m.ProviderModel()
	if !ok {
		return nil, ErrModelNotReady
	}

	rawUserID, err := m.rawUserID(userID)
	if err != nil {
		return nil, err
	}

	key := providerKey(info.Provider, info.ProviderID)
	data := ztype.Map{
		"user_id":            rawUserID,
		"provider":           info.Provider,
		"provider_id":        info.ProviderID,
		"provider_key":       key,
		"provider_email":     normalizeEmail(info.ProviderEmail),
		"provider_username":  info.ProviderUsername,
		"provider_avatar":    info.ProviderAvatar,
		"provider_extension": info.ProviderExtension,
	}
	if len(info.ProviderExtension) == 0 {
		data["provider_extension"] = "{}"
	}

	// 先检查是否已绑定到其他用户（快速失败，但不依赖此检查保证唯一性）
	row, err := providers.FindOne(model.Filter{"provider_key": key})
	if err != nil && !errors.Is(err, model.ErrNoRecord) {
		return nil, err
	}

	// 如果记录已存在
	if row != nil && row.Get("provider_key").String() != "" {
		if row.Get("user_id").String() != rawUserID {
			return nil, ErrProviderAlreadyBound
		}

		// 更新现有绑定
		_, err = providers.UpdateByID(row.Get(model.IDKey()).String(), data)
		if err != nil {
			return nil, err
		}
		return providers.FindOneByID(row.Get(model.IDKey()).String())
	}

	// 尝试插入新绑定
	id, err := providers.Insert(data)
	if err != nil {
		// 处理并发导致的唯一约束冲突
		if strings.Contains(err.Error(), "Duplicate entry") || strings.Contains(err.Error(), "UNIQUE") {
			// 重新查询以获取最新的绑定记录
			row, err := providers.FindOne(model.Filter{"provider_key": key})
			if err != nil {
				return nil, err
			}
			if row.Get("user_id").String() != rawUserID {
				return nil, ErrProviderAlreadyBound
			}
			// 如果是当前用户的绑定，返回更新后的记录
			return row, nil
		}
		return nil, err
	}

	return providers.FindOneByID(id)
}

func (m *Module) unlinkProvider(userID, provider, providerID string) (bool, error) {
	providers, ok := m.ProviderModel()
	if !ok {
		return false, ErrModelNotReady
	}
	user, err := m.userByID(userID)
	if err != nil {
		return false, err
	}

	rawUserID, err := m.rawUserID(userID)
	if err != nil {
		return false, err
	}

	if !userHasLocalPassword(user) {
		count, err := providers.Count(model.Filter{"user_id": rawUserID})
		if err != nil {
			return false, err
		}
		if count <= 1 {
			return false, ErrLastProviderCannotUnlink
		}
	}

	total, err := providers.DeleteMany(model.Filter{
		"user_id":      rawUserID,
		"provider_key": providerKey(provider, providerID),
	})
	return total > 0, err
}

func (m *Module) providerByKey(provider, providerID string) (ztype.Map, error) {
	providers, ok := m.ProviderModel()
	if !ok {
		return nil, ErrModelNotReady
	}

	return providers.FindOne(model.Filter{"provider_key": providerKey(provider, providerID)})
}

func providerKey(provider, providerID string) string {
	return strings.TrimSpace(provider) + ":" + strings.TrimSpace(providerID)
}

func oauthPlaceholderEmail(provider, providerID string) string {
	safeID := strings.NewReplacer("@", "_", "/", "_", "\\", "_", ":", "_", " ", "_").Replace(providerID)
	return strings.TrimSpace(provider) + "+" + safeID + "@oauth.local"
}

func handledOAuthLogin(c *znet.Context) bool {
	prev := c.PrevContent()
	if c.IsAbort() {
		return true
	}

	if prev.Code.Load() != 0 || len(prev.Content) > 0 {
		if prev.Code.Load() != 0 {
			c.Abort(prev.Code.Load())
		} else {
			c.Abort(http.StatusOK)
		}
		return true
	}

	return false
}
