package auth

import (
	"errors"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	zsession "github.com/sohaha/zlsgo/znet/session"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/account"
	"github.com/zlsgo/app_module/model"
	"golang.org/x/crypto/bcrypt"
)

// 限流配置常量
const (
	// 登录限流
	loginMaxAttemptsPerEmail = 5
	loginMaxAttemptsPerIP    = 10
	loginBaseDelay           = 30 * time.Second
	loginMaxDelay            = time.Hour

	// 找回密码限流
	forgotMaxAttemptsPerEmail = 3
	forgotMaxAttemptsPerIP    = 10
	forgotBaseDelay           = 30 * time.Second
	forgotMaxDelay            = time.Hour

	// OAuth flow
	oauthFlowTTL = 10 * time.Minute
)

// 邮箱验证正则
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type UserController struct {
	service.App
	module *Module
	Path   string
}

var _ = reflect.TypeOf(&UserController{})

func (h *UserController) Init(r *znet.Engine) error {
	if err := h.DI.Resolve(&h.module); err != nil {
		return err
	}

	r.Use(znet.RewriteErrorHandler(authErrorHandler))
	r.Use(zsession.New(h.module.Options.Session, func(c *zsession.Config) {
		c.CookieName = h.module.Options.CookieName
		c.ExpiresAt = time.Duration(h.module.Options.SessionTTL) * time.Second
		c.AutoRenew = true
	}))

	r.POST("/create", h.create)
	r.POST("/auth", h.auth)
	r.GET("/get", h.get)
	r.GET("/signout", h.signout)
	r.POST("/update", h.update)
	r.POST("/forgotpassword", h.forgotPassword)
	r.POST("/resetpassword", h.resetPassword)

	return nil
}

func (h *UserController) create(c *znet.Context) (any, error) {
	input, err := getPayload(c)
	if err != nil {
		return nil, toHTTPError(err)
	}

	email := normalizeEmail(input.Get("email").String())
	password := input.Get("password").String()
	nickname := input.Get("nickname").String()
	if nickname == "" {
		nickname = email
	}

	if !isValidEmail(email) {
		return nil, toHTTPError(ErrInvalidEmail)
	}
	if ok, msg := account.ValidatePassword(password, account.DefaultPasswordConfig); !ok {
		return nil, respondBadRequest(msg)
	}

	users, ok := h.module.UserModel()
	if !ok {
		return nil, toHTTPError(ErrModelNotReady)
	}

	// 预检查邮箱是否已存在（快速失败，但不依赖此检查保证唯一性）
	if exist, _ := users.Exists(model.Filter{"email": email}); exist {
		return nil, toHTTPError(ErrEmailExists)
	}

	id, err := users.Insert(ztype.Map{
		"email":           email,
		"password":        password,
		"nickname":        nickname,
		"avatar":          input.Get("avatar").String(),
		"status":          1,
		"settings":        "{}",
		"login_at":        ztime.Now(),
		"session_version": 1,
	})
	if err != nil {
		// 处理并发导致的唯一约束冲突
		if strings.Contains(err.Error(), "Duplicate entry") || strings.Contains(err.Error(), "UNIQUE") {
			return nil, toHTTPError(ErrEmailExists)
		}
		return nil, toHTTPError(err)
	}

	user, err := users.FindOneByID(id)
	if err != nil {
		return nil, toHTTPError(err)
	}

	session, err := zsession.Get(c)
	if err != nil {
		return nil, toHTTPError(err)
	}

	if err = h.module.attachSession(user, session); err != nil {
		return nil, toHTTPError(err)
	}

	return sanitizeUser(user), nil
}

func (h *UserController) auth(c *znet.Context) (any, error) {
	input, err := getPayload(c)
	if err != nil {
		return nil, toHTTPError(err)
	}

	email := normalizeEmail(input.Get("email").String())
	password := input.Get("password").String()
	if !isValidEmail(email) {
		return nil, toHTTPError(ErrInvalidEmail)
	}
	if password == "" {
		return nil, toHTTPError(ErrInvalidPassword)
	}
	loginKey := "login:" + email
	loginIPKey := "login-ip:" + c.GetClientIP()
	if h.module.tooManyRequests(loginKey, loginMaxAttemptsPerEmail, "login") || h.module.tooManyRequests(loginIPKey, loginMaxAttemptsPerIP, "login") {
		return nil, toHTTPError(ErrTooManyRequests)
	}

	user, err := h.module.userByEmail(email)
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			h.module.markFailed(loginKey, "login")
			h.module.markFailed(loginIPKey, "login")
			return nil, toHTTPError(ErrInvalidCredentials)
		}
		return nil, toHTTPError(err)
	}

	if user.Get("status").Int() != 1 {
		return nil, toHTTPError(ErrUserDisabled)
	}

	if err = bcrypt.CompareHashAndPassword(user.Get("password").Bytes(), []byte(password)); err != nil {
		h.module.markFailed(loginKey, "login")
		h.module.markFailed(loginIPKey, "login")
		return nil, toHTTPError(ErrInvalidCredentials)
	}

	users, _ := h.module.UserModel()
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

	session, err := zsession.Get(c)
	if err != nil {
		return nil, toHTTPError(err)
	}

	if err = h.module.attachSession(user, session); err != nil {
		return nil, toHTTPError(err)
	}

	h.module.clearFailed(loginKey, "login")
	h.module.clearFailed(loginIPKey, "login")

	return sanitizeUser(user), nil
}

func (h *UserController) get(c *znet.Context) (any, error) {
	user, _, err := h.module.currentUser(c)
	if err != nil {
		return nil, toHTTPError(err)
	}
	return sanitizeUser(user), nil
}

func (h *UserController) signout(c *znet.Context) (any, error) {
	_, session, err := h.module.currentUser(c)
	if err != nil {
		return nil, toHTTPError(err)
	}

	if err = h.module.invalidateSession(session.ID()); err != nil {
		return nil, toHTTPError(err)
	}
	if err = session.Destroy(); err != nil {
		return nil, toHTTPError(err)
	}

	return ztype.Map{"signed_out": true}, nil
}

func (h *UserController) update(c *znet.Context) (any, error) {
	user, session, err := h.module.currentUser(c)
	if err != nil {
		return nil, toHTTPError(err)
	}

	input, err := getPayload(c)
	if err != nil {
		return nil, toHTTPError(err)
	}

	email := normalizeEmail(input.Get("email").String())
	password := input.Get("password").String()
	currentPassword := input.Get("current_password").String()
	if email == "" && password == "" {
		return nil, toHTTPError(ErrMissingUpdateField)
	}
	hasLocalPassword := userHasLocalPassword(user)
	allowPasswordBootstrap := !hasLocalPassword && email == "" && password != ""
	if !allowPasswordBootstrap && currentPassword == "" {
		return nil, toHTTPError(ErrMissingCurrentPassword)
	}
	if !allowPasswordBootstrap {
		err = bcrypt.CompareHashAndPassword(user.Get("password").Bytes(), []byte(currentPassword))
		if err != nil {
			return nil, toHTTPError(ErrCurrentPasswordWrong)
		}
	}

	update := ztype.Map{}
	if email != "" {
		if !isValidEmail(email) {
			return nil, toHTTPError(ErrInvalidEmail)
		}
		if exists, err := h.module.emailInUse(email, user.Get(model.IDKey()).String()); err != nil {
			return nil, toHTTPError(err)
		} else if exists {
			return nil, toHTTPError(ErrEmailExists)
		}
		update["email"] = email
	}

	passwordChanged := false
	if password != "" {
		if ok, msg := account.ValidatePassword(password, account.DefaultPasswordConfig); !ok {
			return nil, respondBadRequest(msg)
		}
		update["password"] = password
		update["session_version"] = user.Get("session_version").Int() + 1
		passwordChanged = true
	}

	users, _ := h.module.UserModel()
	_, err = users.UpdateByID(user.Get(model.IDKey()).String(), update)
	if err != nil {
		return nil, toHTTPError(err)
	}

	if passwordChanged {
		if err = h.module.invalidateUserSessions(user.Get(model.IDKey()).String()); err != nil {
			return nil, toHTTPError(err)
		}
	}

	user, err = users.FindOneByID(user.Get(model.IDKey()).String())
	if err != nil {
		return nil, toHTTPError(err)
	}

	if passwordChanged {
		if err = h.module.attachSession(user, session); err != nil {
			return nil, toHTTPError(err)
		}
	}

	return sanitizeUser(user), nil
}

func getPayload(c *znet.Context) (ztype.Map, error) {
	data := ztype.Map{}
	if json, err := c.GetJSONs(); err == nil && json != nil {
		if mapped := json.Map(); len(mapped) > 0 {
			data = mapped
		}
	}

	for _, key := range []string{
		"email",
		"password",
		"nickname",
		"avatar",
		"current_password",
		"token",
		"provider",
		"provider_id",
		"provider_email",
		"provider_username",
		"provider_avatar",
	} {
		if value := strings.TrimSpace(c.DefaultFormOrQuery(key, "")); value != "" {
			data[key] = value
		}
	}

	return data, nil
}

func isValidEmail(email string) bool {
	return email != "" && emailRegex.MatchString(email)
}

func normalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}

func sanitizeUser(user ztype.Map) ztype.Map {
	out := ztype.Map{}
	for k, v := range user {
		out[k] = v
	}
	delete(out, "password")
	return out
}

func userHasLocalPassword(user ztype.Map) bool {
	return strings.TrimSpace(user.Get("password").String()) != ""
}

func authErrorHandler(c *znet.Context, err error) {
	if err == nil {
		return
	}

	// 将错误转换为 HTTP 错误并提取状态码
	httpErr := toHTTPError(err)
	if he, ok := httpErr.(*HTTPError); ok {
		c.String(int32(he.Code()), he.Error())
		return
	}

	c.String(http.StatusInternalServerError, err.Error())
}

var zerrorTagTooManyRequests = zerror.TagKind(strconv.Itoa(http.StatusTooManyRequests))
