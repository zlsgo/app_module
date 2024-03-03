package account

import (
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/database/model"
	"github.com/zlsgo/zdb/schema"
)

func initModel(p *Module) error {
	adminDefaultPassword := p.Options.AdminDefaultPassword
	if adminDefaultPassword == "" {
		// TODO: 默认密码，后续是不是要改成随机密码
		adminDefaultPassword = "qw123456."
	}
	inlayUser := append(ztype.Maps{{
		model.IDKey:     1,
		"inlay":         true,
		"administrator": true,
		"avatar":        "data:image/svg+xml,%3Csvg viewBox='0 0 36 36' fill='none' role='img' xmlns='http://www.w3.org/2000/svg' width='128' height='128'%3E%3Ctitle%3EMary Roebling%3C/title%3E%3Cmask id='mask__beam' maskUnits='userSpaceOnUse' x='0' y='0' width='36' height='36'%3E%3Crect width='36' height='36' fill='%23FFFFFF'%3E%3C/rect%3E%3C/mask%3E%3Cg mask='url(%23mask__beam)'%3E%3Crect width='36' height='36' fill='%23f0f0d8'%3E%3C/rect%3E%3Crect x='0' y='0' width='36' height='36' transform='translate(5 -1) rotate(155 18 18) scale(1.2)' fill='%23000000' rx='6'%3E%3C/rect%3E%3Cg transform='translate(3 -4) rotate(-5 18 18)'%3E%3Cpath d='M15 21c2 1 4 1 6 0' stroke='%23FFFFFF' fill='none' stroke-linecap='round'%3E%3C/path%3E%3Crect x='14' y='14' width='1.5' height='2' rx='1' stroke='none' fill='%23FFFFFF'%3E%3C/rect%3E%3Crect x='20' y='14' width='1.5' height='2' rx='1' stroke='none' fill='%23FFFFFF'%3E%3C/rect%3E%3C/g%3E%3C/g%3E%3C/svg%3E",
		// "avatar":        "https://avatars.githubusercontent.com/u/18098083?s=220&v=4",
		"nickname": "admin",
		"status":   "1",
		"account":  "manage",
		"password": adminDefaultPassword,
		"role":     []string{"admin"},
	}}, p.Options.InlayUser...)

	for _, err := range []error{
		accountModel(p.ms, inlayUser),
		roleModel(p.ms),
		permModel(p.ms),
		logModel(p.ms),
	} {
		if err != nil {
			return err
		}
	}
	return nil
}

const accountName = "account"

func accountModel(ms *model.Models, inlayUser ztype.Maps) error {
	_, err := ms.Reg(accountName, model.Define{
		Name: accountName,
		Options: model.Options{
			CryptID:    true,
			Timestamps: true,
		},
		Fields: map[string]model.Field{
			"avatar": {
				Label:    "头像",
				Nullable: true,
				Default:  "",
				Type:     schema.String,
				Size:     1024 * 2,
				Validations: []model.Validations{
					{
						Method: "regex",
						Args:   "^(data:image/|http://|https://|/)",
					},
				},
			},
			"nickname": {
				Type:  schema.String,
				Size:  20,
				Label: "昵称",
			},
			"status": {
				Type:  schema.Int8,
				Size:  9,
				Label: "状态",
				Options: model.FieldOption{
					Enum: []model.FieldEnum{
						{Value: "0", Label: "待激活"},
						{Value: "1", Label: "正常"},
						{Value: "2", Label: "禁用"},
					},
				},
			},
			"salt": {
				Type:     schema.String,
				Size:     4,
				Nullable: true,
				Label:    "盐",
			},
			"login_at": {
				Type:     schema.Time,
				Nullable: true,
				Options:  model.FieldOption{},
				Label:    "登录时间",
			},
			"inlay": {
				Type:     schema.Bool,
				Label:    "是否内置数据",
				Default:  false,
				Nullable: true,
			},
			"administrator": {
				Type:     schema.Bool,
				Label:    "是否超级管理员",
				Default:  false,
				Nullable: true,
			},
			"remark": {
				Type:     schema.String,
				Size:     100,
				Default:  "",
				Nullable: true,
				Label:    "备注",
			},
			"role": {
				Type:     schema.JSON,
				Default:  "[]",
				Nullable: true,
				Label:    "绑定角色",
				Options: model.FieldOption{
					IsArray: true,
				},
			},
			"extension": {
				Type:     schema.JSON,
				Default:  "{}",
				Nullable: true,
				Label:    "扩展信息",
			},
			"account": {
				Label:  "账号",
				Type:   schema.String,
				Unique: true,
				Validations: []model.Validations{
					{
						Method: "minLength",
						Args:   3,
					},
					{
						Method: "maxLength",
						Args:   10,
					},
				},
			},
			"password": {
				Label: "密码",
				Type:  schema.String,
				Options: model.FieldOption{
					Crypt: "PASSWORD",
				},
				Validations: []model.Validations{
					{
						Method: "minLength",
						Args:   3,
					},
					{
						Method: "maxLength",
						Args:   250,
					},
				},
			},
		},
		Values: inlayUser,
	}, false)
	return err
}

const roleName = "role"

func roleModel(ms *model.Models) error {
	_, err := ms.Reg(roleName, model.Define{
		Name: roleName,
		Options: model.Options{
			CryptID:    true,
			Timestamps: true,
		},
		Fields: map[string]model.Field{
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
				Validations: []model.Validations{
					{
						Method:  "regex",
						Args:    "^[a-zA-Z0-9_]+$",
						Message: "角色名别名不能包含特殊符号",
					},
				},
				Label: "角色名别名",
			},
			"status": {
				Type:  schema.Uint8,
				Size:  9,
				Label: "状态",
				Options: model.FieldOption{
					Enum: []model.FieldEnum{
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
				Options: model.FieldOption{
					IsArray: true,
				},
			},
			"permission": {
				Type:     schema.JSON,
				Default:  "[]",
				Nullable: true,
				Label:    "包含权限",
				Options: model.FieldOption{
					IsArray: true,
				},
			},
		},
		Values: []ztype.Map{
			{
				model.IDKey:  1,
				"label":      "管理员",
				"status":     "1",
				"alias":      "admin",
				"inlay":      true,
				"permission": []uint{1},
			},
		},
	}, false)
	return err
}

const permName = "permission"

func permModel(ms *model.Models) error {
	_, err := ms.Reg(permName, model.Define{
		Name: permName,
		Options: model.Options{
			Timestamps: true,
		},
		Fields: map[string]model.Field{
			"label": {
				Type:  schema.String,
				Size:  20,
				Label: "规则名称",
			},
			"alias": {
				Type:     schema.String,
				Size:     20,
				Comment:  "如果不为空，必须唯一",
				Nullable: true,
				Unique:   true,
				Validations: []model.Validations{
					{
						Method:  "regex",
						Args:    "^[a-zA-Z0-9_]+$",
						Message: "规则别名不能包含特殊符号",
					},
				},
				Label: "规则别名",
			},
			"status": {
				Type:  schema.Uint8,
				Size:  9,
				Label: "状态",
				Options: model.FieldOption{
					Enum: []model.FieldEnum{
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
			"action": {
				Type:     schema.String,
				Size:     100,
				Default:  "",
				Nullable: true,
				Label:    "操作",
			},
			"target": {
				Type:     schema.String,
				Size:     225,
				Default:  "",
				Nullable: true,
				Label:    "目标",
			},
			"priority": {
				Type:    schema.Uint32,
				Default: 0,
				Label:   "优先级",
			},
		},
		Values: []ztype.Map{
			{
				model.IDKey: 1,
				"label":     "全局访问",
				"remark":    "可访问全部接口",
				"status":    "1",
				"alias":     "global_allow",
				"inlay":     true,
				"target":    "*",
				"action":    "/*",
			},
		},
	}, false)
	return err
}

const logsName = "logs"

func logModel(ms *model.Models) error {
	_, err := ms.Reg(logsName, model.Define{
		Name: logsName,
		Options: model.Options{
			CryptID: true,
		},
		Fields: map[string]model.Field{
			"account": {
				Type:  schema.String,
				Size:  120,
				Label: "用户账号",
				Index: true,
			},
			"ip": {
				Type:     schema.String,
				Size:     20,
				Label:    "IP",
				Nullable: true,
			},
			"method": {
				Type:     schema.String,
				Size:     20,
				Label:    "请求方法",
				Nullable: true,
			},
			"path": {
				Type:     schema.String,
				Size:     100,
				Label:    "请求路径",
				Nullable: true,
			},
			"params": {
				Type:     schema.JSON,
				Default:  "{}",
				Nullable: true,
				Label:    "请求参数",
			},
			"status": {
				Type:     schema.Uint16,
				Size:     999,
				Label:    "响应状态",
				Nullable: true,
			},
			"message": {
				Type:     schema.String,
				Size:     100,
				Label:    "消息",
				Nullable: true,
			},
			"remark": {
				Type:     schema.Text,
				Default:  "",
				Nullable: true,
				Comment:  "",
				Label:    "请求数据",
			},
			"record_at": {
				Type:     schema.Time,
				Default:  "",
				Nullable: true,
				Comment:  "",
				Label:    "记录时间",
			},
		},
	}, false)
	return err
}
