package member

import (
	"fmt"
	"reflect"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/member/auth"
)

type Auth struct {
	service.App
	module *Module
	Path   string
}

var _ = reflect.TypeOf(&Auth{})

func (h *Auth) Init(r *znet.Engine) error {
	mod, _ := h.module.mods.Get(modelName)
	providers := auth.GetProviders()
	for _, provider := range providers {
		name := provider.Name()
		if err := provider.Init(); err != nil {
			return err
		}
		r.GET(fmt.Sprintf("%s/callback", name), func(c *znet.Context) (ztype.Map, error) {
			info, err := provider.Callback(c)
			if err != nil {
				return nil, err
			}
			_ = mod
			return ztype.Map{
				"provider":          info.Provider,
				"provider_id":       info.ProviderID,
				"provider_username": info.ProviderUsername,
			}, nil
		})
		r.GET(fmt.Sprintf("%s/login", name), provider.Login)
	}
	return nil
}
