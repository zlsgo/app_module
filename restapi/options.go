package restapi

import (
	"github.com/sohaha/zlsgo/znet"
)

type Options struct {
	Middleware   znet.Handler
	ResponseHook func(c *znet.Context, model, args, method string) (next bool)
	Prefix       string
}
