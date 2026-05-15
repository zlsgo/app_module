package member

import (
	"github.com/sohaha/zlsgo/zutil"
	mSchema "github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb/schema"
)

// type Model struct {
// 	model.Model
// }

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

	s.AddField("auth_user_id", mSchema.Field{
		Label:    "Auth用户ID",
		Type:     schema.String,
		Size:     64,
		Unique:   true,
		Nullable: false,
		Comment:  "auth 模块用户主键，作为 member profile 的显式身份关联",
	})

	return s
})
