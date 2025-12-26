package account

import (
	"errors"

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
	if g.m == nil || g.m.accountModel == nil {
		return nil, zerror.InvalidInput.Text("账号模型未初始化")
	}

	if exist, _ := g.m.accountModel.Exists(model.Filter{
		"account": account,
	}); exist {
		return nil, zerror.InvalidInput.Text("账号已存在")
	}

	id, err := g.m.accountModel.Insert(data)
	if err != nil {
		return nil, err
	}
	return ztype.Map{"id": id}, nil
}

func (g *inside) UpdateUser(uid interface{}, data ztype.Map) (resp ztype.Map, err error) {
	if g.m == nil || g.m.accountModel == nil {
		return nil, zerror.InvalidInput.Text("账号模型未初始化")
	}

	user, err := g.m.accountModel.FindOneByID(uid)
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			return nil, zerror.InvalidInput.Text("用户不存在")
		}
		return nil, err
	}

	// 禁止修改超级管理员
	if user.Get("administrator").Bool() {
		return nil, zerror.InvalidInput.Text("不能修改超级管理员")
	}

	if err = g.fixUserData(&data); err != nil {
		return nil, zerror.InvalidInput.Text(err.Error())
	}

	total, err := g.m.accountModel.UpdateByID(uid, data)
	if err != nil {
		return nil, err
	}

	g.m.clearCache("", ztype.ToString(uid))
	return ztype.Map{"total": total}, nil
}

func (g *inside) DeleteUser(uid interface{}) (resp ztype.Map, err error) {
	if g.m == nil || g.m.accountModel == nil {
		return nil, zerror.InvalidInput.Text("账号模型未初始化")
	}

	user, err := g.m.accountModel.FindOneByID(uid)
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			return nil, zerror.InvalidInput.Text("用户不存在")
		}
		return nil, err
	}

	if user.Get("inlay").Bool() {
		return nil, zerror.InvalidInput.Text("不能删除内置用户")
	}

	_, err = g.m.accountModel.DeleteByID(uid)
	if err != nil {
		return nil, err
	}

	g.m.clearCache("", ztype.ToString(uid))

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
		if g.m != nil && g.m.index != nil && g.m.index.roleModel != nil {
			roleModel := g.m.index.roleModel
			inputs := make([]string, 0, len(r))
			for _, item := range r {
				if item != "" {
					inputs = append(inputs, item)
				}
			}

			roleSet := make(map[string]struct{}, len(inputs))
			roles := make([]string, 0, len(inputs))

			if len(inputs) > 0 {
				aliasRoles, _ := roleModel.Model().Find(model.In("alias", inputs), func(co *model.CondOptions) {
					co.Fields = []string{"alias"}
				})
				for i := range aliasRoles {
					alias := aliasRoles[i].Get("alias").String()
					if alias == "" {
						continue
					}
					if _, ok := roleSet[alias]; ok {
						continue
					}
					roleSet[alias] = struct{}{}
					roles = append(roles, alias)
				}
			}

			idList := make([]string, 0, len(inputs))
			for _, item := range inputs {
				if _, ok := roleSet[item]; ok {
					continue
				}
				if _, err := roleModel.DeCryptID(item); err == nil {
					idList = append(idList, item)
				}
			}

			if len(idList) > 0 {
				idRoles, _ := roleModel.Model().Find(model.In(model.IDKey(), idList), func(co *model.CondOptions) {
					co.Fields = []string{"alias"}
				})
				for i := range idRoles {
					alias := idRoles[i].Get("alias").String()
					if alias == "" {
						continue
					}
					if _, ok := roleSet[alias]; ok {
						continue
					}
					roleSet[alias] = struct{}{}
					roles = append(roles, alias)
				}
			}

			data.Set("role", roles)
		}
	}

	password := data.Get("password")
	if password.Exists() && password.String() == "" {
		return errors.New("密码不能为空")
	}
	if password.Exists() {
		pwd := password.String()
		ok, msg := ValidatePassword(pwd, DefaultPasswordConfig)
		if !ok {
			return errors.New(msg)
		}
		data.Set("password", pwd)
	}

	return nil
}
