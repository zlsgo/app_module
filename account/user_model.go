package account

import (
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model"
	mSchema "github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb/schema"
)

type AccountModel struct {
	*model.Store
	mod *Module
	m   *model.Schema
}

var accountModel *AccountModel

func GetAccountModel() *AccountModel {
	if accountModel == nil {
		panic("account model not define")
	}

	return accountModel
}

const accountName = "account"

func accountModelDefine(p *Module) error {
	adminDefaultPassword := p.Options.AdminDefaultPassword
	if adminDefaultPassword == "" {
		// TODO: 默认密码，后续是不是要改成随机密码
		adminDefaultPassword = "qw123456."
	}
	inlayUser := append(ztype.Maps{{
		model.IDKey():   1,
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

	b := true
	mod, err := p.mods.Reg(accountName, mSchema.Schema{
		Name: accountName,
		Options: mSchema.Options{
			CryptID:    &b,
			Timestamps: &b,
		},
		Fields: map[string]mSchema.Field{
			"avatar": {
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
				Options: mSchema.FieldOption{
					Enum: []mSchema.FieldEnum{
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
				Options:  mSchema.FieldOption{},
				Label:    "登录时间",
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
			"administrator": {
				Type:     schema.Bool,
				Label:    "是否超级管理员",
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
			"role": {
				Type:     schema.JSON,
				Default:  "[]",
				Nullable: true,
				Label:    "绑定角色",
				Options: mSchema.FieldOption{
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
				Validations: []mSchema.Validations{
					{
						Method: "minLength",
						Args:   3,
					},
					{
						Method: "maxLength",
						Args:   20,
					},
				},
				Options: mSchema.FieldOption{
					ReadOnly: true,
				},
			},
			"password": {
				Label: "密码",
				Type:  schema.String,
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
			},
		},
		Values: inlayUser,
	}, false)

	if err == nil {
		accountModel = &AccountModel{Store: mod.Model(), mod: p, m: mod}
	}
	return err
}
