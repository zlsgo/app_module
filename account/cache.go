package account

import (
	"errors"

	"github.com/zlsgo/app_module/account/jwt"
	"github.com/zlsgo/app_module/model"

	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/ztype"
)

var userCache = zcache.NewFast()

func getUserForCache(m *model.Schema, uid string) (ztype.Map, error) {
	user, _ := userCache.ProvideGet(uid, func() (interface{}, bool) {
		f, err := model.FindOne(m, uid)
		if err != nil || f.IsEmpty() {
			return ztype.Map{}, false
		}
		if *m.GetDefine().Options.CryptID {
			id, _ := m.DeCryptID(uid)
			_ = f.Set("raw_id", id)
		}
		return f, true
	})

	info, ok := user.(ztype.Map)
	if !ok || info.IsEmpty() {
		return nil, zerror.WrapTag(zerror.NotFound)(errors.New("用户不存在"))
	}

	return info, nil
}

func deleteUserForCache(uid string) {
	userCache.Delete(uid)
}

var jwtCache = zcache.NewFast()

func getJWTForCache(m *model.Schema, token, jwtKey string) (string, error) {
	uid, _ := jwtCache.ProvideGet(token, func() (interface{}, bool) {
		info, err := jwt.Parse(token, jwtKey)
		if err != nil {
			return "", false
		}

		salt := info.Info[:saltLen]
		uid := info.Info[saltLen:]
		f, err := model.FindCols(m, "salt", uid)
		if err != nil || f.Index(0).String() != salt {
			return "", false
		}

		return uid, true
	})

	id, ok := uid.(string)
	if !ok || id == "" {
		return "", zerror.WrapTag(zerror.Unauthorized)(errors.New("登录状态过期，请重新登录"))
	}

	return id, nil
}

func deleteJWTForCache(token string) {
	jwtCache.Delete(token)
}

func clearCache(token, uid string) {
	if token != "" {
		deleteJWTForCache(token)
	}
	if uid != "" {
		deleteUserForCache(uid)
	}
}
