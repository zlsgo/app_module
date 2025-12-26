package account

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model"
)

func TestFixUserDataRoleAliasAndID(t *testing.T) {
	tt := zlsgo.NewTest(t)
	mod, cleanup := newTestModule(t)
	defer cleanup()

	roleID, err := mod.index.roleModel.Model().Insert(ztype.Map{
		"label":  "Tester",
		"alias":  "tester",
		"status": 1,
	})
	tt.NoError(err)

	encID := ztype.ToString(roleID)

	_, err = mod.Inside.CreateUser(ztype.Map{
		"account":  "roleuser",
		"password": "Aa123456",
		"role":     []string{"tester", encID},
	})
	tt.NoError(err)

	user, err := mod.accountModel.FindOne(model.Filter{"account": "roleuser"})
	tt.NoError(err)
	roles := user.Get("role").SliceString()
	tt.Equal(1, len(roles))
	tt.Equal("tester", roles[0])
}

func TestFixUserDataRoleAliasPriority(t *testing.T) {
	tt := zlsgo.NewTest(t)
	mod, cleanup := newTestModule(t)
	defer cleanup()

	alias, err := mod.index.roleModel.EnCryptID("1")
	tt.NoError(err)

	_, err = mod.index.roleModel.Model().Insert(ztype.Map{
		"label":  "Shadow",
		"alias":  alias,
		"status": 1,
	})
	tt.NoError(err)

	_, err = mod.Inside.CreateUser(ztype.Map{
		"account":  "aliasfirst",
		"password": "Aa123456",
		"role":     []string{alias},
	})
	tt.NoError(err)

	user, err := mod.accountModel.FindOne(model.Filter{"account": "aliasfirst"})
	tt.NoError(err)
	roles := user.Get("role").SliceString()
	tt.Equal(1, len(roles))
	tt.Equal(alias, roles[0])
}
