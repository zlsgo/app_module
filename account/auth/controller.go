package auth

import (
	"errors"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/account/jwt"
)

func NewRouter(key string, expire int, callback func(*znet.Context, Provider) (id string, err error)) func(r *znet.Engine, providers []AuthProvider, enabledProviders []string) (jwtParse func(c *znet.Context) (string, error), err error) {
	return func(r *znet.Engine, providers []AuthProvider, enabledProviders []string) (jwtParse func(c *znet.Context) (string, error), err error) {
		return regController(r, providers, enabledProviders, key, expire, callback)
	}
}

func regController(r *znet.Engine, providers []AuthProvider, enabledProviders []string, key string, expire int, callback func(*znet.Context, Provider) (string, error)) (jwtParse func(c *znet.Context) (string, error), err error) {
	jwtParse = func(c *znet.Context) (string, error) {
		token := jwt.GetToken(c)
		if token == "" {
			return "", errors.New("token not found")
		}

		info, err := jwt.Parse(token, key)
		if err != nil {
			return "", err
		}

		return info.Info, nil
	}

	for i := range providers {
		provider := providers[i]
		name := provider.Name()
		if !zarray.Contains(enabledProviders, name) {
			continue
		}

		e := r.Group(name)
		if err = provider.Init(e); err != nil {
			return
		}

		e.GET("/callback", func(c *znet.Context) (ztype.Map, error) {
			info, err := provider.Callback(c)
			if err != nil {
				return nil, err
			}

			id, err := callback(c, info)
			if err != nil {
				return nil, err
			}

			if id == "" {
				return nil, errors.New("user id not found")
			}

			accessToken, refreshToken, err := jwt.GenToken(id, key, expire)
			if err != nil {
				return nil, err
			}
			return ztype.Map{
				"token":         accessToken,
				"refresh_token": refreshToken,
			}, nil
		})

		e.GET("/login", provider.Login)
	}

	return
}
