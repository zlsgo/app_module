package member

import (
	"github.com/zlsgo/app_module/model/define"
	"github.com/zlsgo/zdb/schema"
)

const modelName = "member"

var modelDefine = define.Define{
	Name: modelName,
	Options: define.ModelOptions{
		CryptID:    true,
		Timestamps: true,
	},
	Fields: map[string]define.Field{
		"avatar": {
			Label:    "头像",
			Nullable: true,
			Default:  "",
			Type:     schema.String,
			Size:     1024 * 2,
			Validations: []define.Validations{
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
			Options: define.FieldOption{
				Enum: []define.FieldEnum{
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
			Options:  define.FieldOption{},
			Label:    "登录时间",
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
		"auth": {
			Type:     schema.JSON,
			Default:  "[]",
			Nullable: true,
			Label:    "第三方平台授权",
			Options: define.FieldOption{
				IsArray: true,
			},
		},
		"account": {
			Label:  "账号",
			Type:   schema.String,
			Unique: true,
			Validations: []define.Validations{
				{
					Method: "minLength",
					Args:   3,
				},
				{
					Method: "maxLength",
					Args:   20,
				},
			},
			Options: define.FieldOption{
				ReadOnly: true,
			},
		},
		"password": {
			Label:    "密码",
			Type:     schema.String,
			Nullable: true,
			Options: define.FieldOption{
				Crypt: "PASSWORD",
			},
			Validations: []define.Validations{
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
}
