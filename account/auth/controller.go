package auth

import (
	"errors"
	"fmt"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/account/jwt"
)

func NewRouter(key string, expire int, login func(*znet.Context, any) error, callback func(*znet.Context, Provider) (id string, err error)) func(r *znet.Engine, providers []AuthProvider, enabledProviders []string) (jwtParse func(c *znet.Context) (string, error), err error) {
	return func(r *znet.Engine, providers []AuthProvider, enabledProviders []string) (jwtParse func(c *znet.Context) (string, error), err error) {
		return regController(r, providers, enabledProviders, key, expire, login, callback)
	}
}

func regController(r *znet.Engine, providers []AuthProvider, enabledProviders []string, key string, expire int, login func(*znet.Context, any) error, callback func(*znet.Context, Provider) (string, error)) (jwtParse func(c *znet.Context) (string, error), err error) {
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

	for _, provider := range providers {
		name := provider.Name()
		if !zarray.Contains(enabledProviders, name) {
			continue
		}
		if err = provider.Init(); err != nil {
			return
		}

		r.GET(fmt.Sprintf("%s/callback", name), func(c *znet.Context) (ztype.Map, error) {
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
		r.GET(fmt.Sprintf("%s/login", name), func(c *znet.Context) error {
			value, err := provider.Login(c)
			if err != nil {
				return err
			}

			return login(c, value)
		})
	}

	return
}
