package define

import (
	"github.com/sohaha/zlsgo/zvalid"
	"github.com/zlsgo/zdb/schema"
)

type (
	Fields map[string]Field
	Field  struct {
		Default     interface{}     `json:"default"`
		Unique      interface{}     `json:"unique"`
		Index       interface{}     `json:"index"`
		Comment     string          `json:"comment"`
		Label       string          `json:"label"`
		Type        schema.DataType `json:"type"`
		Validations []Validations   `json:"validations"`
		Options     FieldOption     `json:"ModelOptions"`
		Before      []string        `json:"before"`
		After       []string        `json:"after"`
		ValidRules  zvalid.Engine
		Size        uint64 `json:"size"`
		Nullable    bool   `json:"nullable"`
		// quoteName   string `json:"-"`
	}
	FieldEnum struct {
		Value string `json:"value"`
		Label string `json:"label"`
	}
	FieldOption struct {
		FormatTime       string      `json:"format_time"`
		Crypt            string      `json:"crypt"`
		Enum             []FieldEnum `json:"enum"`
		IsArray          bool        `json:"is_array"`
		ReadOnly         bool        `json:"readonly"`
		DisableMigration bool        `json:"disable_migration"`
		// Quote      bool        `json:"quote"`
	}
)

func (f *Field) GetValidations() *zvalid.Engine {
	return &f.ValidRules
}
