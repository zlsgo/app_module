package member

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/account/auth"
	"github.com/zlsgo/app_module/model"
)

type Auth struct {
	service.App
	module        *Module
	memberModel   *model.Store
	providerModel *model.Store
	Path          string
}

var _ = reflect.TypeOf(&Auth{})

func (h *Auth) initModel() (err error) {
	err = h.DI.Resolve(&h.module)
	if err != nil {
		return err
	}

	var ok bool
	h.memberModel, ok = h.module.models.Get(modelName)
	if !ok {
		return fmt.Errorf("model %s not found", modelName)
	}
	h.providerModel, ok = h.module.models.Get(modelProviderName)
	if !ok {
		return fmt.Errorf("model %s not found", modelProviderName)
	}
	return nil
}

func (h *Auth) Init(r *znet.Engine) (err error) {
	if err = h.initModel(); err != nil {
		return err
	}

	storageType := h.module.schemas.Storage().GetStorageType()
	if storageType != model.SQLStorage {
		return fmt.Errorf("storage type %s not supported", storageType)
	}

	regController := auth.NewRouter(
		h.module.Options.key,
		h.module.Options.Expire,
		func(ctx *znet.Context, p auth.Provider) (mid string, err error) {
			if p.Provider == "" || p.ProviderID == "" {
				return "", zerror.InvalidInput.Text("提供商不能为空")
			}

			provider, err := h.providerModel.FindOne(model.Filter{
				"provider":    p.Provider,
				"provider_id": p.ProviderID,
			}, func(co *model.CondOptions) {
				co.Fields = []string{"member_id"}
			})
			if err != nil {
				return "", err
			}

			if provider.IsEmpty() {
				// 没有找到提供者，则需要先创建会员
				account := "__" + p.Provider + "_" + p.ProviderID
				password := zstring.Rand(16)
				err = h.module.schemas.Storage().Transaction(func(s *model.SQL) (err error) {
					memberSchema := h.memberModel.Schema(s)
					_, exists, _ := model.FindCol(memberSchema, "account", ztype.Map{"account": account})
					if exists {
						return errors.New("账号已存在，无法自动注册")
					}

					nickname := p.ProviderUsername
					if nickname == "" {
						nickname = account
					}

					id, err := model.Insert(memberSchema, ztype.Map{
						"account":  account,
						"password": password,
						"nickname": nickname,
						"avatar":   p.ProviderAvatar,
						"login_at": ztime.Now(),
					})
					if err != nil {
						return zerror.With(err, "用户注册失败")
					}

					mid = ztype.ToString(id)
					account, err := memberSchema.DeCryptID(mid)
					if err != nil {
						return zerror.With(err, "用户信息解析失败")
					}

					data := ztype.Map{
						"provider":          p.Provider,
						"member_id":         account,
						"provider_id":       p.ProviderID,
						"provider_username": p.ProviderUsername,
						"provider_avatar":   p.ProviderAvatar,
					}
					if p.ProviderExtension != nil {
						data["provider_extension"] = p.ProviderExtension
					}
					_, err = model.Insert(h.providerModel.Schema(s), data)
					if err != nil {
						return zerror.With(err, "用户信息保存失败")
					}
					return nil
				})

				return
			}

			return provider.Get("member_id").String(), nil
		},
	)

	h.module.jwtParse, err = regController(r, h.module.Options.Providers, h.module.Options.EnabledProviders)
	if err != nil {
		return err
	}

	return err
}
