package member

import (
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/app_module/model/define"
	"github.com/zlsgo/zdb/schema"
)

const modelName = "member"

var modelDefine = zutil.Once(func() define.Define {
	modelDefine := define.New(modelName)
	modelDefine.SetOptions(define.ModelOptions{
		CryptID:    true,
		Timestamps: true,
	})

	modelDefine.AddField("avatar", define.Field{
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
	})

	modelDefine.AddField("nickname", define.Field{
		Type:     schema.String,
		Size:     20,
		Label:    "昵称",
		Nullable: true,
		Default:  "",
	})

	modelDefine.AddField("status", define.Field{
		Type:    schema.Int8,
		Size:    9,
		Label:   "状态",
		Default: 1,
		Options: define.FieldOption{
			Enum: []define.FieldEnum{
				{Value: "1", Label: "正常"},
				{Value: "0", Label: "禁用"},
			},
		},
	})

	modelDefine.AddField("salt", define.Field{
		Type:     schema.String,
		Size:     4,
		Nullable: true,
		Label:    "盐",
	})

	modelDefine.AddField("login_at", define.Field{
		Type:     schema.Time,
		Nullable: true,
		Options:  define.FieldOption{},
		Label:    "登录时间",
	})

	modelDefine.AddField("remark", define.Field{
		Type:     schema.String,
		Size:     100,
		Default:  "",
		Nullable: true,
		Label:    "备注",
	})

	modelDefine.AddField("extension", define.Field{
		Type:     schema.JSON,
		Default:  "{}",
		Nullable: true,
		Label:    "扩展信息",
	})

	modelDefine.AddField("provider", define.Field{
		Type:     schema.String,
		Default:  "",
		Nullable: true,
		Label:    "第三方登录",
	})

	modelDefine.AddField("provider_id", define.Field{
		Type:     schema.String,
		Default:  "",
		Nullable: true,
		Label:    "第三方ID",
	})

	modelDefine.AddField("provider_username", define.Field{
		Type:     schema.String,
		Default:  "",
		Nullable: true,
		Label:    "第三方用户名",
	})

	modelDefine.AddField("account", define.Field{
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
				Args:   120,
			},
		},
		Options: define.FieldOption{
			ReadOnly: true,
		},
	})

	modelDefine.AddField("password", define.Field{
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
	})

	return modelDefine
})
