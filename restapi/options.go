package restapi

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_module/model/define"
)

type Options struct {
	Prefix       string
	DisableAuth  bool
	ModelsDefine []define.Define
	ResponseHook func(c *znet.Context, model, args, method string) (next bool)
}
