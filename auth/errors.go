package auth

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/sohaha/zlsgo/zerror"
)

// 预定义错误类型
var (
	ErrUserNotFound         = errors.New("用户不存在")
	ErrUserDisabled         = errors.New("账号已禁用")
	ErrInvalidCredentials   = errors.New("邮箱或密码错误")
	ErrInvalidEmail         = errors.New("请输入正确的邮箱")
	ErrInvalidPassword      = errors.New("密码格式不正确")
	ErrEmailExists          = errors.New("邮箱已存在")
	ErrNotAuthenticated     = errors.New("未登录")
	ErrSessionInvalid       = errors.New("登录态已失效")
	ErrSessionExpired       = errors.New("登录态已过期")
	ErrInvalidToken         = errors.New("重置令牌无效")
	ErrTokenUsed            = errors.New("重置令牌已使用")
	ErrTokenExpired         = errors.New("重置令牌已过期")
	ErrProviderNotEnabled   = errors.New("未启用该 provider")
	ErrProviderInfoIncomplete = errors.New("provider 信息不完整")
	ErrProviderAlreadyBound = errors.New("该第三方账号已绑定其他用户")
	ErrOAuthStateInvalid    = errors.New("OAuth 状态不存在或已失效")
	ErrOAuthStateExpired    = errors.New("OAuth 状态已过期")
	ErrOAuthStateMismatch   = errors.New("OAuth 状态校验失败")
	ErrProviderMismatch     = errors.New("OAuth provider 不匹配")
	ErrOAuthUserMismatch    = errors.New("OAuth 绑定用户不匹配")
	ErrLastProviderCannotUnlink = errors.New("请先设置本地密码，再解绑最后一个第三方登录")
	ErrBindingNotFound      = errors.New("未找到绑定关系")
	ErrMissingCurrentPassword = errors.New("请输入当前密码")
	ErrCurrentPasswordWrong = errors.New("当前密码错误")
	ErrMissingUpdateField   = errors.New("请至少提供一个更新字段")
	ErrTooManyRequests      = errors.New("请求过多，请稍后再试")
	ErrModelNotReady        = errors.New("model not ready")
	ErrSessionNotInit       = errors.New("session not initialized")
	ErrResetSenderNotConfigured = errors.New("reset password sender not configured")
	ErrMissingResetToken    = errors.New("缺少重置令牌")
)

// HTTP 错误包装器
type HTTPError struct {
	code int
	err  error
	tag  zerror.TagKind
}

func (e *HTTPError) Error() string {
	return e.err.Error()
}

func (e *HTTPError) Unwrap() error {
	return e.err
}

func (e *HTTPError) Code() int {
	return e.code
}

func (e *HTTPError) Tag() zerror.TagKind {
	return e.tag
}

// 错误响应构造器
func respondUnauthorized(text string) error {
	return wrapHTTPError(http.StatusUnauthorized, zerror.Unauthorized, text)
}

func respondBadRequest(text string) error {
	return wrapHTTPError(http.StatusBadRequest, zerror.InvalidInput, text)
}

func respondTooManyRequests(text string) error {
	return wrapHTTPError(http.StatusTooManyRequests, zerrorTagTooManyRequests, text)
}

func respondInternalError(text string) error {
	return wrapHTTPError(http.StatusInternalServerError, zerror.TagKind(strconv.Itoa(http.StatusInternalServerError)), text)
}

func wrapHTTPError(code int, tag zerror.TagKind, text string) error {
	return &HTTPError{
		code: code,
		err:  errors.New(text),
		tag:  tag,
	}
}

// 将预定义错误转换为 HTTP 错误
func toHTTPError(err error) error {
	if err == nil {
		return nil
	}

	// 如果已经是 HTTPError，直接返回
	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		return err
	}

	// 根据错误类型映射到 HTTP 错误
	switch {
	case errors.Is(err, ErrUserNotFound):
		return respondUnauthorized(err.Error())
	case errors.Is(err, ErrUserDisabled):
		return respondUnauthorized(err.Error())
	case errors.Is(err, ErrInvalidCredentials):
		return respondUnauthorized(err.Error())
	case errors.Is(err, ErrNotAuthenticated):
		return respondUnauthorized(err.Error())
	case errors.Is(err, ErrSessionInvalid):
		return respondUnauthorized(err.Error())
	case errors.Is(err, ErrSessionExpired):
		return respondUnauthorized(err.Error())
	case errors.Is(err, ErrInvalidToken):
		return respondUnauthorized(err.Error())
	case errors.Is(err, ErrTokenUsed):
		return respondUnauthorized(err.Error())
	case errors.Is(err, ErrTokenExpired):
		return respondUnauthorized(err.Error())
	case errors.Is(err, ErrOAuthStateInvalid):
		return respondBadRequest(err.Error())
	case errors.Is(err, ErrOAuthStateExpired):
		return respondBadRequest(err.Error())
	case errors.Is(err, ErrOAuthStateMismatch):
		return respondBadRequest(err.Error())
	case errors.Is(err, ErrProviderMismatch):
		return respondBadRequest(err.Error())
	case errors.Is(err, ErrOAuthUserMismatch):
		return respondUnauthorized(err.Error())
	case errors.Is(err, ErrCurrentPasswordWrong):
		return respondUnauthorized(err.Error())
	case errors.Is(err, ErrTooManyRequests):
		return respondTooManyRequests(err.Error())
	case errors.Is(err, ErrLastProviderCannotUnlink):
		return respondBadRequest(err.Error())
	default:
		// 未知错误，返回内部服务器错误
		return wrapHTTPError(http.StatusInternalServerError, zerror.TagKind(strconv.Itoa(http.StatusInternalServerError)), "服务器内部错误")
	}
}
