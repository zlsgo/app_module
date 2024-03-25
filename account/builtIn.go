package account

import (
	"errors"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/ztype"
)

type inside struct {
}

var Inside = &inside{}

func (g *inside) CreateUser(data ztype.Map) (resp ztype.Map, err error) {
	if err := fixUserData(&data); err != nil {
		return nil, zerror.WrapTag(zerror.InvalidInput)(err)
	}

	// DEV: 需要校验角色是否存在
	roles := data.Get("role").SliceString()
	_ = roles

	account := data.Get("account").String()
	if !data.Get("nickname").Exists() {
		data.Set("nickname", account)
	}

	if !data.Get("status").Exists() {
		data.Set("status", 1)
	}

	// 检查账号是否存在
	if exist, _ := GetAccountModel().Exists(ztype.Map{
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

func (g *inside) UpdateUser(id interface{}, data ztype.Map) (resp ztype.Map, err error) {
	user, err := GetAccountModel().FindOneByID(id)
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

	if err = fixUserData(&data); err != nil {
		return nil, zerror.InvalidInput.Text(err.Error())
	}

	total, err := GetAccountModel().UpdateByID(id, data)
	if err != nil {
		return nil, err
	}
	return ztype.Map{"total": total}, nil
}

// fixUserData 修复并兼容用户数据各种情况
func fixUserData(data *ztype.Map) error {
	// 禁止添加超级管理员
	_ = data.Delete("administrator")
	// 禁止标记为内置用户
	_ = data.Delete("inlay")
	// 验证盐不应该可以人为修改
	_ = data.Delete("salt")
	// 登录时间不应该可以人为修改
	_ = data.Delete("login_at")

	r := data.Get("role").Slice(true)
	if r.Len() == 0 {
		role := data.Get("role").String()
		if role != "" {
			data.Set("role", []interface{}{role})
		}
	}

	if data.Has("password") {
		if data.Get("password").String() == "" {
			return errors.New("密码不能为空")
		}
	}

	return nil
}
