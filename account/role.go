package account

import (
	"errors"

	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/account/rbac"
	"github.com/zlsgo/app_module/model"
	mSchema "github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb/schema"
)

const roleName = "role"

func roleModel(ms *model.Schemas) error {
	b := true
	_, err := ms.Reg(roleName, mSchema.Schema{
		Name: roleName,
		Options: mSchema.Options{
			CryptID:    &b,
			Timestamps: &b,
		},
		Fields: map[string]mSchema.Field{
			"label": {
				Type:  schema.String,
				Size:  20,
				Label: "角色名称",
			},
			"alias": {
				Type:    schema.String,
				Size:    20,
				Comment: "必须唯一",
				Unique:  true,
				Default: "",
				Validations: []mSchema.Validations{
					{
						Method:  "regex",
						Args:    "^[a-zA-Z0-9_]+$",
						Message: "角色名别名不能包含特殊符号",
					},
				},
				Label: "角色名别名",
				Options: mSchema.FieldOption{
					ReadOnly: true,
				},
			},
			"status": {
				Type:  schema.Uint8,
				Size:  9,
				Label: "状态",
				Options: mSchema.FieldOption{
					Enum: []mSchema.FieldEnum{
						{Value: "0", Label: "待激活"},
						{Value: "1", Label: "正常"},
						{Value: "2", Label: "禁用"},
					},
				},
			},
			"inlay": {
				Type:     schema.Bool,
				Label:    "是否内置数据",
				Default:  false,
				Nullable: true,
				Options: mSchema.FieldOption{
					ReadOnly: true,
				},
			},
			"remark": {
				Type:     schema.String,
				Size:     100,
				Default:  "",
				Nullable: true,
				Label:    "备注",
			},
			"extension": {
				Type:     schema.JSON,
				Default:  "{}",
				Nullable: true,
				Label:    "扩展信息",
			},
			"children": {
				Type:     schema.JSON,
				Default:  "[]",
				Nullable: true,
				Label:    "包含角色",
				Options: mSchema.FieldOption{
					IsArray: true,
				},
			},
			"permission": {
				Type:     schema.JSON,
				Default:  "[]",
				Nullable: true,
				Label:    "包含权限",
				Options: mSchema.FieldOption{
					IsArray: true,
				},
			},
		},
		Values: []ztype.Map{
			{
				model.IDKey(): 1,
				"label":       "管理员",
				"status":      "1",
				"alias":       "admin",
				"inlay":       true,
				"permission":  []uint{1},
			},
		},
	}, false)

	return err
}

func (m *Module) setPermission(permission *rbac.RBAC, roleInfo ztype.Map) error {
	permModel, ok := m.mods.Get(permName)
	if !ok {
		return errors.New(permName + " permName not found")
	}

	role := rbac.NewRole(rbac.MatchPriorityDeny)
	permissionIds := roleInfo.Get("permission").SliceInt()
	if len(permissionIds) == 0 {
		return permission.RemoveRole(roleInfo.Get("alias").String())
	}

	perms, err := model.Find(permModel, ztype.Map{
		model.IDKey(): permissionIds,
		"status":      1,
	}, func(o *model.CondOptions) {
		o.Fields = []string{"action", "alias", "target", "priority"}
	})
	if err != nil {
		return err
	}

	for _, perm := range perms {
		role.AddGlobPermission(perm.Get("priority").Int(), perm.Get("action").String(), perm.Get("target").String())
	}
	return permission.MergerRole(roleInfo.Get("alias").String(), role)
}
