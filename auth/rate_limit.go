package auth

import (
	"net/http"
	"time"

	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zutil"
)

// limitBucket 返回限流桶对应的缓存与退避参数
func (m *Module) limitBucket(bucket string) (*zcache.FastCache, time.Duration, time.Duration) {
	switch bucket {
	case "forgot":
		return m.forgotLimit, forgotBaseDelay, forgotMaxDelay
	case "reset":
		return m.resetLimit, resetBaseDelay, resetMaxDelay
	default:
		return m.loginLimit, loginBaseDelay, loginMaxDelay
	}
}

func (m *Module) tooManyRequests(cacheKey string, limit int, bucket string) bool {
	cache, _, _ := m.limitBucket(bucket)

	total, ok := cache.Get(cacheKey)
	if !ok {
		return false
	}
	return ztype.ToInt(total) >= limit
}

func (m *Module) markFailed(cacheKey string, bucket string) {
	cache, baseDelay, maxDelay := m.limitBucket(bucket)

	total, _ := cache.Get(cacheKey)
	count := ztype.ToInt(total) + 1
	cache.Set(cacheKey, count, zutil.BackOffDelay(ztype.ToInt(total), baseDelay, maxDelay))
}

func (m *Module) clearFailed(cacheKey string, bucket string) {
	cache, _, _ := m.limitBucket(bucket)

	cache.Delete(cacheKey)
}

func tooManyRequestsError(text string) error {
	return wrapHTTPError(http.StatusTooManyRequests, zerrorTagTooManyRequests, text)
}
