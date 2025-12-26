package account

import (
	"errors"
	"time"

	"github.com/zlsgo/app_module/account/jwt"
	"github.com/zlsgo/app_module/model"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
)

func (m *Module) getUserForCache(schema *model.Schema, uid string) (ztype.Map, error) {
	user, _ := m.userCache.ProvideGet(uid, func() (interface{}, bool) {
		f, err := model.FindOne[ztype.Map](schema.Model(), model.ID(uid), func(o *model.CondOptions) {
			o.Fields = schema.GetFields("password", "salt")
		})
		if err != nil {
			return ztype.Map{}, false
		}
		if schema.GetDefine().Options.CryptID != nil && *schema.GetDefine().Options.CryptID {
			id, _ := schema.DeCryptID(uid)
			_ = f.Set("raw_id", id)
		}
		_ = f.Delete("password")
		_ = f.Delete("salt")
		return f, true
	}, time.Hour)

	info, ok := user.(ztype.Map)
	if !ok || info.IsEmpty() {
		return nil, zerror.WrapTag(zerror.NotFound)(errors.New("用户不存在"))
	}

	_ = info.Delete("password")
	_ = info.Delete("salt")
	return info, nil
}

func (m *Module) deleteUserForCache(uid string) {
	m.userCache.Delete(uid)
}

func (m *Module) getJWTForCache(schema *model.Schema, token, jwtKey string) (string, error) {
	errUnauthorized := zerror.WrapTag(zerror.Unauthorized)(errors.New("登录状态过期，请重新登录"))
	resp, _ := m.jwtCache.ProvideGet(token, func() (interface{}, bool) {
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
		f, err := model.FindCols[string](schema.Model(), "salt", model.ID(uid))
		if err != nil || len(f) == 0 || f[0] != salt {
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
		m.deleteJWTForCache(token)
		return "", errUnauthorized
	}

	return uid, nil
}

func (m *Module) deleteJWTForCache(token string) {
	m.jwtCache.Delete(token)
}

func (m *Module) clearCache(token, uid string) {
	if token != "" {
		m.deleteJWTForCache(token)
	}
	if uid != "" {
		m.deleteUserForCache(uid)
	}
}

// getRolesForCache 获取缓存的角色列表
func (m *Module) getRolesForCache(roleModel *model.Schema) ([]ztype.Map, error) {
	result, _ := m.roleCache.ProvideGet("all_active_roles", func() (interface{}, bool) {
		roles, err := model.FindMaps(roleModel.Model(), model.Filter{
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
	return model.FindMaps(roleModel.Model(), model.Filter{
		"status": 1,
	})
}

func (m *Module) loadActiveRoles(roleModel *model.Schema) ([]ztype.Map, error) {
	roles, err := model.FindMaps(roleModel.Model(), model.Filter{
		"status": 1,
	})
	if err != nil {
		return nil, err
	}
	m.roleCache.Set("all_active_roles", roles)
	return roles, nil
}

// invalidateRoleCache 使角色缓存失效
func (m *Module) invalidateRoleCache() {
	m.roleCache.Delete("all_active_roles")
}

// getUserPermissionForCache 获取用户权限缓存
func (m *Module) getUserPermissionForCache(userID string, getter func() (interface{}, error)) (interface{}, error) {
	cacheKey := "user_perm_" + userID

	result, _ := m.permissionCache.ProvideGet(cacheKey, func() (interface{}, bool) {
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
func (m *Module) invalidateUserPermCache(userID string) {
	cacheKey := "user_perm_" + userID
	m.permissionCache.Delete(cacheKey)
}
