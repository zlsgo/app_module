package schema

import (
	"github.com/sohaha/zlsgo/zvalid"
	"github.com/zlsgo/zdb/schema"
)

type (
	Fields map[string]Field

	Field struct {
		Default     interface{}     `json:"default,omitempty"`
		Unique      interface{}     `json:"unique,omitempty"`
		Index       interface{}     `json:"index,omitempty"`
		Comment     string          `json:"comment,omitempty"`
		Label       string          `json:"label"`
		Type        schema.DataType `json:"type"`
		Validations []Validations   `json:"validations,omitempty"`
		Options     FieldOption     `json:"options,omitempty"`
		Before      []string        `json:"-"`
		After       []string        `json:"-"`
		ValidRules  zvalid.Engine   `json:"-"`
		// 如果是数字类型则为长度，如果是字符串类型则为最大长度
		Size     uint64 `json:"size,omitempty"`
		Nullable bool   `json:"nullable,omitempty"`
		// quoteName   string `json:"-"`
	}

	FieldEnum struct {
		Value string `json:"value"`
		Label string `json:"label"`
	}

	FieldOption struct {
		FormatTime       string      `json:"format_time,omitempty"`
		Crypt            string      `json:"crypt,omitempty"`
		Enum             []FieldEnum `json:"enum,omitempty"`
		IsArray          bool        `json:"is_array,omitempty"`
		ReadOnly         bool        `json:"readonly,omitempty"`
		DisableMigration bool        `json:"disable_migration,omitempty"`
		// Quote      bool        `json:"quote"`
	}
)

func (f *Field) GetValidations() *zvalid.Engine {
	return &f.ValidRules
}
