package auth

import (
	"github.com/sohaha/zlsgo/zutil"
	mSchema "github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb/schema"
)

const (
	userModelName               = "user"
	providerModelName           = "oauth_provider"
	passwordResetTokenModelName = "password_reset_token"
	sessionModelName            = "session"
)

var userModelDefine = zutil.Once(func() mSchema.Schema {
	s := mSchema.New(userModelName)
	s.SetOptions().SetCryptID(true).SetTimestamps(true).SetSoftDeletes(true)

	_ = s.AddField("email", mSchema.Field{
		Label:   "邮箱",
		Type:    schema.String,
		Size:    200,
		Unique:  true,
		Index:   true,
		Default: "",
		Validations: []mSchema.Validations{
			{Method: "minLength", Args: 3},
			{Method: "maxLength", Args: 200},
		},
	})
	_ = s.AddField("password", mSchema.Field{
		Label:    "密码",
		Type:     schema.String,
		Nullable: true,
		Options: mSchema.FieldOption{
			Crypt: "PASSWORD",
		},
	})
	_ = s.AddField("nickname", mSchema.Field{
		Label:    "昵称",
		Type:     schema.String,
		Size:     50,
		Nullable: true,
		Default:  "",
	})
	_ = s.AddField("avatar", mSchema.Field{
		Label:    "头像",
		Type:     schema.String,
		Size:     2048,
		Nullable: true,
		Default:  "",
	})
	_ = s.AddField("status", mSchema.Field{
		Label:   "状态",
		Type:    schema.Int8,
		Size:    9,
		Default: 1,
		Options: mSchema.FieldOption{
			Enum: []mSchema.FieldEnum{
				{Value: "1", Label: "正常"},
				{Value: "0", Label: "禁用"},
			},
		},
	})
	_ = s.AddField("settings", mSchema.Field{
		Label:    "设置",
		Type:     schema.JSON,
		Default:  "{}",
		Nullable: true,
	})
	_ = s.AddField("login_at", mSchema.Field{
		Label:    "登录时间",
		Type:     schema.Time,
		Nullable: true,
	})
	_ = s.AddField("session_version", mSchema.Field{
		Label:   "会话版本",
		Type:    schema.Int,
		Default: 1,
	})

	return s
})

var providerModelDefine = zutil.Once(func() mSchema.Schema {
	s := mSchema.New(providerModelName)
	s.SetOptions().SetTimestamps(true)

	_ = s.AddField("user_id", mSchema.Field{
		Label: "用户ID",
		Type:  schema.Int,
		Index: true,
	})
	_ = s.AddField("provider", mSchema.Field{
		Label:   "提供商",
		Type:    schema.String,
		Size:    50,
		Default: "",
		Index:   true,
	})
	_ = s.AddField("provider_id", mSchema.Field{
		Label:   "提供商用户ID",
		Type:    schema.String,
		Size:    120,
		Default: "",
		Index:   true,
	})
	_ = s.AddField("provider_key", mSchema.Field{
		Label:   "提供商唯一键",
		Type:    schema.String,
		Size:    180,
		Default: "",
		Unique:  true,
	})
	_ = s.AddField("provider_email", mSchema.Field{
		Label:    "提供商邮箱",
		Type:     schema.String,
		Size:     200,
		Nullable: true,
		Default:  "",
	})
	_ = s.AddField("provider_username", mSchema.Field{
		Label:    "提供商用户名",
		Type:     schema.String,
		Size:     120,
		Nullable: true,
		Default:  "",
	})
	_ = s.AddField("provider_avatar", mSchema.Field{
		Label:    "提供商头像",
		Type:     schema.String,
		Size:     2048,
		Nullable: true,
		Default:  "",
	})
	_ = s.AddField("provider_extension", mSchema.Field{
		Label:    "提供商扩展信息",
		Type:     schema.JSON,
		Default:  "{}",
		Nullable: true,
	})

	return s
})

var passwordResetTokenModelDefine = zutil.Once(func() mSchema.Schema {
	s := mSchema.New(passwordResetTokenModelName)
	s.SetOptions().SetTimestamps(true)

	_ = s.AddField("user_id", mSchema.Field{
		Label: "用户ID",
		Type:  schema.Int,
		Index: true,
	})
	_ = s.AddField("token", mSchema.Field{
		Label:   "重置令牌",
		Type:    schema.String,
		Size:    120,
		Unique:  true,
		Default: "",
	})
	_ = s.AddField("token_hash", mSchema.Field{
		Label:   "令牌哈希",
		Type:    schema.String,
		Size:    128,
		Index:   true,
		Default: "",
	})
	_ = s.AddField("expire_at", mSchema.Field{
		Label: "过期时间",
		Type:  schema.Time,
	})
	_ = s.AddField("used_at", mSchema.Field{
		Label:    "使用时间",
		Type:     schema.Time,
		Nullable: true,
	})

	return s
})

var sessionModelDefine = zutil.Once(func() mSchema.Schema {
	s := mSchema.New(sessionModelName)
	s.SetOptions().SetTimestamps(true)

	_ = s.AddField("user_id", mSchema.Field{
		Label: "用户ID",
		Type:  schema.Int,
		Index: true,
	})
	_ = s.AddField("session_key", mSchema.Field{
		Label:   "会话键",
		Type:    schema.String,
		Size:    120,
		Unique:  true,
		Default: "",
	})
	_ = s.AddField("status", mSchema.Field{
		Label:   "状态",
		Type:    schema.Int8,
		Size:    9,
		Default: 1,
		Options: mSchema.FieldOption{
			Enum: []mSchema.FieldEnum{
				{Value: "1", Label: "正常"},
				{Value: "0", Label: "失效"},
			},
		},
	})
	_ = s.AddField("expire_at", mSchema.Field{
		Label:    "过期时间",
		Type:     schema.Time,
		Nullable: true,
	})
	_ = s.AddField("last_seen_at", mSchema.Field{
		Label:    "最后活跃时间",
		Type:     schema.Time,
		Nullable: true,
	})

	return s
})
