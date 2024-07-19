package member

import (
	"errors"
	"reflect"
	"time"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/account/auth"
	"github.com/zlsgo/app_module/model"
)

type Auth struct {
	service.App
	module    *Module
	Path      string
	userModel func() (*model.Model, bool)
}

var _ = reflect.TypeOf(&Auth{})

func (h *Auth) Init(r *znet.Engine) (err error) {
	if h.userModel == nil {
		return errors.New("user model not found")
	}
	if _, ok := h.userModel(); !ok {
		return errors.New("user model not found")
	}

	regController := auth.NewRouter(
		h.module.Options.Key,
		h.module.Options.Expire,
		func(ctx *znet.Context, _ any) error {
			return nil
		},
		func(ctx *znet.Context, p auth.Provider) (string, error) {
			mod, _ := h.userModel()
			data := ztype.Map{
				"provider":    p.Provider,
				"provider_id": p.ProviderID,
			}

			user, _ := mod.Operation().FindOne(data)
			if !user.IsEmpty() {
				id := user.Get(model.IDKey()).String()
				mod.Operation().UpdateByID(id, ztype.Map{
					"login_at": time.Now(),
				})
				return id, nil
			}

			data["account"] = p.Provider + "_" + p.ProviderID
			data["provider_username"] = p.ProviderUsername

			id, err := mod.Operation().Insert(data)
			return ztype.ToString(id), err
		},
	)

	h.module.jwtParse, err = regController(r, h.module.Options.Providers, h.module.Options.EnabledProviders)
	if err != nil {
		return err
	}

	return err
}
