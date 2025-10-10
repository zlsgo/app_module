package account

import (
	"errors"
	"time"

	"github.com/zlsgo/app_module/account/jwt"
	"github.com/zlsgo/app_module/model"

	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/ztime"
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
	}, time.Hour)

	info, ok := user.(ztype.Map)
	if !ok || info.IsEmpty() {
		return nil, zerror.WrapTag(zerror.NotFound)(errors.New("用户不存在"))
	}

	return info, nil
}

func deleteUserForCache(uid string) {
	userCache.Delete(uid)
}

var (
	jwtCache        = zcache.NewFast()
	errUnauthorized = zerror.WrapTag(zerror.Unauthorized)(errors.New("登录状态过期，请重新登录"))
)

func getJWTForCache(m *model.Schema, token, jwtKey string) (string, error) {
	resp, _ := jwtCache.ProvideGet(token, func() (interface{}, bool) {
		info, err := jwt.Parse(token, jwtKey)
		if err != nil {
			return [2]interface{}{}, false
		}

		ttl := time.Until(time.Unix(info.ExpiresAt, 0))
		if ttl <= 0 {
			return [2]interface{}{}, false
		}

		salt := info.Info[:saltLen]
		uid := info.Info[saltLen:]
		f, err := model.FindCols(m, "salt", uid)
		if err != nil || f.Index(0).String() != salt {
			return [2]interface{}{}, false
		}

		return [2]interface{}{uid, info.ExpiresAt}, true
	}, time.Hour) // 使用固定1小时，避免重复解析

	v, ok := resp.([2]interface{})
	if !ok {
		return "", errUnauthorized
	}

	uid, uidOK := v[0].(string)
	expiresAt, expiresAtOK := v[1].(int64)
	if !uidOK || uid == "" || !expiresAtOK || expiresAt == 0 {
		return "", errUnauthorized
	}

	if ztime.Clock()/1000000 > expiresAt {
		deleteJWTForCache(token)
		return "", errUnauthorized
	}

	return uid, nil
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

var (
	roleCache = zcache.NewFast(func(o *zcache.Options) {
		o.AutoCleaner = true
		o.Expiration = time.Minute * 5
	}) // 角色缓存
	permissionCache = zcache.NewFast(func(o *zcache.Options) {
		o.AutoCleaner = true
		o.Expiration = time.Minute * 5
	}) // 权限缓存
)

// getRolesForCache 获取缓存的角色列表
func getRolesForCache(roleModel *model.Schema) ([]ztype.Map, error) {
	result, _ := roleCache.ProvideGet("all_active_roles", func() (interface{}, bool) {
		roles, err := model.Find(roleModel, ztype.Map{
			"status": 1,
		})
		if err != nil {
			return []ztype.Map{}, false
		}
		return roles, true
	})

	if roles, ok := result.([]ztype.Map); ok {
		return roles, nil
	}

	// 如果缓存获取失败，直接查询数据库
	return model.Find(roleModel, ztype.Map{
		"status": 1,
	})
}

// invalidateRoleCache 使角色缓存失效
func invalidateRoleCache() {
	roleCache.Delete("all_active_roles")
}

// getUserPermissionForCache 获取用户权限缓存
func getUserPermissionForCache(userID string, getter func() (interface{}, error)) (interface{}, error) {
	cacheKey := "user_perm_" + userID

	result, _ := permissionCache.ProvideGet(cacheKey, func() (interface{}, bool) {
		perm, err := getter()
		if err != nil {
			return nil, false
		}
		return perm, true
	})

	if result == nil {
		return getter()
	}

	return result, nil
}

// invalidateUserPermCache 使用户权限缓存失效
func invalidateUserPermCache(userID string) {
	cacheKey := "user_perm_" + userID
	permissionCache.Delete(cacheKey)
}
