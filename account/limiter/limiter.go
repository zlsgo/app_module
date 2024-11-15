package limiter

import (
	"errors"
	"net/http"
	"time"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/znet/limiter"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zutil"
)

var MaxIPRequestPerSecond = 10

var IPMiddleware = zutil.Once(func() func(c *znet.Context) error {
	limit := limiter.NewRule()
	if MaxIPRequestPerSecond == 0 {
		MaxIPRequestPerSecond = 10
	}
	limit.AddRule(time.Second, MaxIPRequestPerSecond)
	tooManyRequestsTag := zerror.WrapTag(zerror.TagKind(ztype.ToString(http.StatusTooManyRequests)))(errors.New("Too many requests"))
	return func(c *znet.Context) error {
		if !limit.AllowVisitByIP(c.GetClientIP()) {
			return tooManyRequestsTag
		}
		c.Next()
		return nil
	}
})
