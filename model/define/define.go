package define

import (
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb/builder"
)

type (
	Define struct {
		Fields    Fields                    `json:"fields"`
		Extend    ztype.Map                 `json:"extend"`
		Relations map[string]*ModelRelation `json:"relations"`
		// Hook      func(name string, m *Model) error `json:"-"`
		Table   Table        `json:"table"`
		Name    string       `json:"name"`
		Values  ztype.Maps   `json:"values"`
		Options ModelOptions `json:"options"`
	}

	Table struct {
		Name    string `json:"name"`
		Comment string `json:"comment"`
	}

	ModelOptions struct {
		Salt             string   `json:"crypt_salt"`
		LowFields        []string `json:"low_fields"`
		FieldsSort       []string `json:"fields_sort"`
		CryptLen         int      `json:"crypt_len"`
		DisabledMigrator bool     `json:"disabled_migrator"`
		SoftDeletes      bool     `json:"soft_deletes"`
		Timestamps       bool     `json:"timestamps"`
		CryptID          bool     `json:"crypt_id"`
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
