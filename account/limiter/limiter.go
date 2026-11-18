package limiter

import (
	"net/http"
	"time"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/znet/limiter"
	"github.com/sohaha/zlsgo/zutil"
)

var MaxIPRequestPerSecond = 10

var IPMiddleware = zutil.Once(func() func(c *znet.Context) error {
	limit := limiter.NewRule()
	if MaxIPRequestPerSecond == 0 {
		MaxIPRequestPerSecond = 10
	}
	limit.AddRule(time.Second, MaxIPRequestPerSecond)
	return func(c *znet.Context) error {
		if !limit.AllowVisitByIP(c.GetClientIP()) {
			c.String(http.StatusTooManyRequests, "Too many requests")
			c.Abort()
			return nil
		}
		c.Next()
		return nil
	}
})
