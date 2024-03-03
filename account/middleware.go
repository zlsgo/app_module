package account

import (
	"errors"

	"github.com/zlsgo/app_module/account/jwt"
	"github.com/zlsgo/app_module/account/rbac"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/database/model"
)

var (
	verifyPermissions func(c *znet.Context) error
)

func (p *Module) RegMiddleware(r *znet.Engine) error {
	if verifyPermissions == nil {
		return errors.New("jwt key is empty")
	}

	r.Use(verifyPermissions)
	return nil
}

func (p *Module) initMiddleware() error {
	permissionDenied := zerror.WrapTag(zerror.PermissionDenied)

	userModel, ok := p.ms.Get(accountName)
	if !ok {
		return errors.New(accountName + " model not found")
	}

	logModel, ok := p.ms.Get(logsName)
	if !ok {
		return errors.New(logsName + " model not found")
	}

	roleModel, ok := p.ms.Get(roleName)
	if !ok {
		return errors.New(roleName + " model not found")
	}

	permModel, ok := p.ms.Get(permName)
	if !ok {
		return errors.New(permName + " model not found")
	}

	// TODO: 可能需要独立出来方便做缓存
	roles, err := model.Find(roleModel, ztype.Map{
		"status": 1,
	})
	if err != nil {
		return err
	}

	permission := rbac.New()
	// 添加权限规则
	for _, r := range roles {
		role := rbac.NewRole(rbac.MatchPriorityDeny)
		perms, err := model.Find(permModel, ztype.Map{
			model.IDKey: r.Get("permission").SliceInt(),
			"status":    1,
		}, func(o *model.CondOptions) error {
			o.Fields = []string{"action", "alias", "target", "priority"}
			return nil
		})
		if err != nil {
			return err
		}

		for _, perm := range perms {
			role.AddGlobPermission(perm.Get("priority").Int(), perm.Get("action").String(), perm.Get("target").String())
		}
		err = permission.AddRole(r.Get("name").String(), role)
		if err != nil {
			return err
		}
	}

	verifyPermissions = func(c *znet.Context) error {
		token := jwt.GetToken(c)
		if token == "" {
			return zerror.WrapTag(zerror.Unauthorized)(errors.New("无法访问，请先登录"))
		}

		if userModel == nil {
			return errors.New(accountName + " model not found")
		}

		uid, err := getJWTForCache(userModel, token, p.Options.key)
		if err != nil {
			return err
		}

		c.WithValue("uid", uid)

		u, err := getUserForCache(userModel, uid)
		if err != nil {
			return err
		}

		if u.Get("status").Int() != 1 {
			return permissionDenied(errors.New("用户已被禁用"))
		}

		logRequest(c, logModel, u)

		if u.Get("administrator").Bool() {
			return nil
		}

		for _, r := range u.Get("role").SliceString() {
			isAllow, _ := permission.Can(r, c.Request.Method, c.Request.URL.Path)
			if isAllow {
				return nil
			}
		}

		return permissionDenied(errors.New("无权访问"))
	}
	return nil
}
