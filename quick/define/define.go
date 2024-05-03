package define

import (
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb/builder"
)

type (
	Define struct {
		Name      string                    `json:"name"`
		Comment   string                    `json:"comment"`
		Fields    Fields                    `json:"fields"`
		Extend    ztype.Map                 `json:"extend"`
		Relations map[string]*ModelRelation `json:"relations"`
		// Hook      func(hook string, m *Model) error `json:"-"`
		Values  ztype.Maps `json:"values"`
		Options Options    `json:"options"`
	}

	Options struct {
		Salt               string             `json:"crypt_salt"`           // 加密盐
		ShowFields         []string           `json:"low_fields"`           // 显示字段
		FieldsSort         []string           `json:"fields_sort"`          // 字段排序
		CryptLen           int                `json:"crypt_len"`            // 加密长度
		DisabledMigrator   bool               `json:"disabled_migrator"`    // 禁用迁移
		MigrationOldColumn MigrationOldColumn `json:"migration_old_column"` // 旧字段处理方式
		SoftDeletes        bool               `json:"soft_deletes"`         // 软删除
		Timestamps         bool               `json:"timestamps"`           // 默认添加时间字段
		CryptID            bool               `json:"crypt_id"`             // 是否加密ID
	}

	Validations struct {
		Args    interface{} `json:"args"`
		Method  string      `json:"method"`
		Message string      `json:"message"`
		// Trigger ValidTriggerType `json:"trigger"`
	}

	ModelRelation struct {
		Name    string             `json:"name"`
		Type    string             `json:"type"`
		Join    builder.JoinOption `json:"-"`
		Model   string             `json:"model"`
		Foreign string             `json:"foreign"`
		Key     string             `json:"key"`
		Fields  []string           `json:"Fields"`
		Limit   int                `json:"limit"`
	}
)
