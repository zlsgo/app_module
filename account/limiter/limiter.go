package limiter

import (
	"net/http"
	"time"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/znet/limiter"
)

// MaxIPRequestPerSecond 默认的每 IP 每秒请求上限，供 IPMiddleware 使用
var MaxIPRequestPerSecond = 10

// NewIPMiddleware 创建一份独立的 IP 限流中间件，每秒最多允许 rps 次请求。
// rps <= 0 时回退到 MaxIPRequestPerSecond 的当前值。
// 每次调用都会新建限流器，因此不同引擎/路由组之间不会共享计数。
func NewIPMiddleware(rps int) func(c *znet.Context) error {
	if rps <= 0 {
		rps = MaxIPRequestPerSecond
	}
	if rps <= 0 {
		rps = 10
	}

	limit := limiter.NewRule()
	limit.AddRule(time.Second, rps)

	return func(c *znet.Context) error {
		if !limit.AllowVisitByIP(c.GetClientIP()) {
			c.String(http.StatusTooManyRequests, "Too many requests")
			c.Abort()
			return nil
		}
		c.Next()
		return nil
	}
}

// IPMiddleware 使用 MaxIPRequestPerSecond 创建 IP 限流中间件
func IPMiddleware() func(c *znet.Context) error {
	return NewIPMiddleware(MaxIPRequestPerSecond)
}
