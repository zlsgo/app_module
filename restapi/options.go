package restapi

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_module/model"
)

type Options struct {
	Middleware          znet.Handler
	ResponseHook        func(c *znet.Context, model, args, method string) (next bool)
	Prefix              string
	MaxPageSize         int
	AllowMethods        map[string]bool
	AllowFields         map[string]bool
	DefaultFields       []string
	RequireFields       bool
	AllowFilterFields   map[string]bool
	AllowOrderFields    map[string]bool
	AllowRelations      map[string]bool
	DefaultOrder        []model.OrderByItem
	ErrorHandler        znet.ErrHandlerFunc
	DisableErrorHandler bool
	RejectUnknownQuery  bool
	AllowQueryKeys      map[string]bool
}
