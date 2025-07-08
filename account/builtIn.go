package account

import (
	"errors"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model"
)

type inside struct {
	m *Module
}

func (g *inside) CreateUser(data ztype.Map) (resp ztype.Map, err error) {
	if err = g.fixUserData(&data); err != nil {
		return nil, zerror.InvalidInput.Text(err.Error())
	}

	// DEV: 需要校验角色是否存在
	roles := data.Get("role").SliceString()
	_ = roles

	account := data.Get("account").String()
	if account == "" {
		return nil, zerror.InvalidInput.Text("账号不能为空")
	}

	if data.Get("nickname").String() == "" {
		data.Set("nickname", account)
	}

	if !data.Get("status").Exists() {
		data.Set("status", 1)
	}

	// 检查账号是否存在
	if exist, _ := GetAccountModel().Exists(model.Filter{
		"account": account,
	}); exist {
		return nil, zerror.InvalidInput.Text("账号已存在")
	}

	id, err := GetAccountModel().Insert(data)
	if err != nil {
		return nil, err
	}
	return ztype.Map{"id": id}, nil
}

func (g *inside) UpdateUser(uid interface{}, data ztype.Map) (resp ztype.Map, err error) {
	user, err := GetAccountModel().FindOneByID(uid)
	if err != nil {
		return nil, err
	}

	if user.IsEmpty() {
		return nil, zerror.InvalidInput.Text("用户不存在")
	}

	// 禁止修改超级管理员
	if user.Get("administrator").Bool() {
		return nil, zerror.InvalidInput.Text("不能修改超级管理员")
	}

	if err = g.fixUserData(&data); err != nil {
		return nil, zerror.InvalidInput.Text(err.Error())
	}

	total, err := GetAccountModel().UpdateByID(uid, data)
	if err != nil {
		return nil, err
	}

	clearCache("", ztype.ToString(uid))
	return ztype.Map{"total": total}, nil
}

func (g *inside) DeleteUser(uid interface{}) (resp ztype.Map, err error) {
	user, err := GetAccountModel().FindOneByID(uid)
	if err != nil {
		return nil, err
	}

	if user.IsEmpty() {
		return nil, zerror.InvalidInput.Text("用户不存在")
	}

	if user.Get("inlay").Bool() {
		return nil, zerror.InvalidInput.Text("不能删除内置用户")
	}

	_, err = GetAccountModel().DeleteByID(uid)
	if err != nil {
		return nil, err
	}

	clearCache("", ztype.ToString(uid))

	return nil, err
}

// fixUserData 修复并兼容用户数据各种情况
func (g *inside) fixUserData(data *ztype.Map) error {
	// 禁止添加超级管理员
	_ = data.Delete("administrator")
	// 禁止标记为内置用户
	_ = data.Delete("inlay")
	// 验证盐不应该可以人为修改
	_ = data.Delete("salt")
	// 登录时间不应该可以人为修改
	_ = data.Delete("login_at")

	role := data.Get("role")
	if role.Exists() {
		r := data.Get("role").SliceString(true)
		if len(r) == 0 {
			role := data.Get("role").String()
			if role != "" {
				r = []string{role}
			}
		}

		// 检查角色是否存在
		if g.m != nil {
			roles, _ := g.m.index.roleModel.Model().Find(model.Filter{"id": r}, func(co *model.CondOptions) {
				co.Fields = []string{"alias"}
			})
			data.Set("role", zarray.Map(roles, func(i int, v ztype.Map) string {
				return v.Get("alias").String()
			}))
		}
	}

	password := data.Get("password")
	if password.Exists() && password.String() == "" {
		return errors.New("密码不能为空")
	}

	return nil
}
