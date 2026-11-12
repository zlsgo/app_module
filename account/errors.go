package account

import (
	"net/http"
	"strconv"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
)

// statusFromError 把带 tag 的错误映射成 HTTP 状态码。
//
// 模块内部用 zerror.InvalidInput / Unauthorized / PermissionDenied 这类语义 tag，
// 限流中间件则直接传数字 tag（"429"）。znet 默认错误处理器不认 tag，会把所有
// 错误一律渲染成 500，所以这里显式映射。
func statusFromError(err error) int {
	switch zerror.GetTag(err) {
	case zerror.InvalidInput:
		return http.StatusBadRequest
	case zerror.NotFound:
		return http.StatusNotFound
	case zerror.Unauthorized:
		return http.StatusUnauthorized
	case zerror.PermissionDenied:
		return http.StatusForbidden
	case zerror.Cancelled:
		return http.StatusRequestTimeout
	case zerror.Internal:
		return http.StatusInternalServerError
	}

	if code, convErr := strconv.Atoi(string(zerror.GetTag(err))); convErr == nil && code >= 100 && code <= 599 {
		return code
	}

	return http.StatusInternalServerError
}

// defaultErrorHandler 把处理器返回的错误按 tag 渲染成带状态码的 JSON。
//
// 注册它会替换 znet 的默认处理器（500 纯文本），因此在应用自己通过 DI 提供
// znet.ErrHandlerFunc 时不安装，保留应用的自定义格式。
func defaultErrorHandler(c *znet.Context, err error) {
	if err == nil {
		return
	}

	code := statusFromError(err)
	c.JSON(int32(code), znet.ApiData{
		Code: int32(code),
		Msg:  err.Error(),
	})
}

// useDefaultErrorHandler 始终安装默认错误处理器，确保账户模块有效错误渲染。
func (m *Module) useDefaultErrorHandler(r *znet.Engine) {
	r.Use(znet.RewriteErrorHandler(defaultErrorHandler))
}
