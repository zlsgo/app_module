package member

import (
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/app_module/model"
	mSchema "github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb/schema"
)

type Model struct {
	model.Model
}

const modelName = "member"

var modelDefine = zutil.Once(func() mSchema.Schema {
	s := mSchema.New(modelName)
	s.SetOptions().SetCryptID(true).SetTimestamps(true).SetSoftDeletes(true)
	s.AddField("avatar", mSchema.Field{
		Label:    "头像",
		Nullable: true,
		Default:  "",
		Type:     schema.String,
		Size:     1024 * 2,
		Validations: []mSchema.Validations{
			{
				Method: "regex",
				Args:   "^(data:image/|http://|https://|/)",
			},
		},
	})

	s.AddField("nickname", mSchema.Field{
		Type:     schema.String,
		Size:     20,
		Label:    "昵称",
		Nullable: true,
		Default:  "",
	})

	s.AddField("status", mSchema.Field{
		Type:    schema.Int8,
		Size:    9,
		Label:   "状态",
		Default: 1,
		Options: mSchema.FieldOption{
			Enum: []mSchema.FieldEnum{
				{Value: "1", Label: "正常"},
				{Value: "0", Label: "禁用"},
			},
		},
	})

	s.AddField("salt", mSchema.Field{
		Type:     schema.String,
		Size:     4,
		Nullable: true,
		Label:    "盐",
	})

	s.AddField("login_at", mSchema.Field{
		Type:     schema.Time,
		Nullable: true,
		Options:  mSchema.FieldOption{},
		Label:    "登录时间",
	})

	s.AddField("remark", mSchema.Field{
		Type:     schema.String,
		Size:     100,
		Default:  "",
		Nullable: true,
		Label:    "备注",
	})

	s.AddField("extension", mSchema.Field{
		Type:     schema.JSON,
		Default:  "{}",
		Nullable: true,
		Label:    "扩展信息",
	})

	s.AddField("provider", mSchema.Field{
		Type:     schema.String,
		Default:  "",
		Nullable: true,
		Label:    "第三方登录",
	})

	s.AddField("provider_id", mSchema.Field{
		Type:     schema.String,
		Default:  "",
		Nullable: true,
		Label:    "第三方ID",
	})

	s.AddField("provider_username", mSchema.Field{
		Type:     schema.String,
		Default:  "",
		Nullable: true,
		Label:    "第三方用户名",
	})

	s.AddField("account", mSchema.Field{
		Label:  "账号",
		Type:   schema.String,
		Unique: true,
		Validations: []mSchema.Validations{
			{
				Method: "minLength",
				Args:   3,
			},
			{
				Method: "maxLength",
				Args:   120,
			},
		},
		Options: mSchema.FieldOption{
			ReadOnly: true,
		},
	})

	s.AddField("password", mSchema.Field{
		Label:    "密码",
		Type:     schema.String,
		Nullable: true,
		Options: mSchema.FieldOption{
			Crypt: "PASSWORD",
		},
		Validations: []mSchema.Validations{
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

	return s
})
