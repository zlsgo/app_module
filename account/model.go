package account

import (
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/restapi"
	"github.com/zlsgo/zdb/schema"
)

func initModel(p *Module) error {
	for _, err := range []error{
		accountModelDefine(p),
		roleModel(p.mods),
		permModel(p.mods),
		logModel(p.mods),
		messageModelDefine(p),
	} {
		if err != nil {
			return err
		}
	}
	return nil
}

const roleName = "role"

func roleModel(ms *restapi.Models) error {
	_, err := ms.Reg(roleName, restapi.Define{
		Name: roleName,
		Options: restapi.ModelOptions{
			CryptID:    true,
			Timestamps: true,
		},
		Fields: map[string]restapi.Field{
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
				Validations: []restapi.Validations{
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
				Options: restapi.FieldOption{
					Enum: []restapi.FieldEnum{
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
				Options: restapi.FieldOption{
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
				Options: restapi.FieldOption{
					IsArray: true,
				},
			},
			"permission": {
				Type:     schema.JSON,
				Default:  "[]",
				Nullable: true,
				Label:    "包含权限",
				Options: restapi.FieldOption{
					IsArray: true,
				},
			},
		},
		Values: []ztype.Map{
			{
				restapi.IDKey: 1,
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

const permName = "permission"

func permModel(ms *restapi.Models) error {
	_, err := ms.Reg(permName, restapi.Define{
		Name: permName,
		Options: restapi.ModelOptions{
			Timestamps: true,
		},
		Fields: map[string]restapi.Field{
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
				Validations: []restapi.Validations{
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
				Options: restapi.FieldOption{
					Enum: []restapi.FieldEnum{
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
				restapi.IDKey: 1,
				"label":       "全局访问",
				"remark":      "可访问全部接口",
				"status":      "1",
				"alias":       "global_allow",
				"inlay":       true,
				"target":      "*",
				"action":      "/*",
			},
		},
	}, false)
	return err
}

const logsName = "logs"

func logModel(ms *restapi.Models) error {
	_, err := ms.Reg(logsName, restapi.Define{
		Name: logsName,
		Options: restapi.ModelOptions{
			CryptID: true,
		},
		Fields: map[string]restapi.Field{
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
