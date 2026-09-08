package auth

import (
	"net/http"

	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zutil"
)

func (m *Module) tooManyRequests(cacheKey string, limit int, bucket string) bool {
	var cache = m.loginLimit

	if bucket == "forgot" {
		cache = m.forgotLimit
	}

	total, ok := cache.Get(cacheKey)
	if !ok {
		return false
	}
	return ztype.ToInt(total) >= limit
}

func (m *Module) markFailed(cacheKey string, bucket string) {
	var cache = m.loginLimit
	var baseDelay = loginBaseDelay
	var maxDelay = loginMaxDelay

	if bucket == "forgot" {
		cache = m.forgotLimit
		baseDelay = forgotBaseDelay
		maxDelay = forgotMaxDelay
	}

	total, _ := cache.Get(cacheKey)
	count := ztype.ToInt(total) + 1
	cache.Set(cacheKey, count, zutil.BackOffDelay(ztype.ToInt(total), baseDelay, maxDelay))
}

func (m *Module) clearFailed(cacheKey string, bucket string) {
	var cache = m.loginLimit
	if bucket == "forgot" {
		cache = m.forgotLimit
	}

	cache.Delete(cacheKey)
}

func tooManyRequestsError(text string) error {
	return wrapHTTPError(http.StatusTooManyRequests, zerrorTagTooManyRequests, text)
}
