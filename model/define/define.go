package define

import (
	"errors"

	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb/builder"
)

type (
	Defines []Define
	Define  struct {
		Fields    Fields                    `json:"fields"`
		Extend    ztype.Map                 `json:"extend,omitempty"`
		Relations map[string]*ModelRelation `json:"relations,omitempty"`
		// Hook      func(name string, m *Model) error `json:"-"`
		Table      Table        `json:"table,omitempty"`
		Name       string       `json:"name"`
		Values     ztype.Maps   `json:"values,omitempty"`
		Options    ModelOptions `json:"options,omitempty"`
		SchemaPath string       `json:"-"`
	}

	Table struct {
		Name    string `json:"name,omitempty"`
		Comment string `json:"comment,omitempty"`
	}

	ModelOptions struct {
		Salt             string   `json:"crypt_salt,omitempty"`
		LowFields        []string `json:"low_fields,omitempty"`
		FieldsSort       []string `json:"fields_sort,omitempty"`
		CryptLen         int      `json:"crypt_len,omitempty"`
		DisabledMigrator bool     `json:"disabled_migrator,omitempty"`
		SoftDeletes      bool     `json:"soft_deletes,omitempty"`
		Timestamps       bool     `json:"timestamps,omitempty"`
		CryptID          bool     `json:"crypt_id,omitempty"`
	}

	Validations struct {
		Args    interface{} `json:"args"`
		Method  string      `json:"method"`
		Message string      `json:"message,omitempty"`
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
		Limit   int                `json:"limit,omitempty"`
	}
)

func New(name string) Define {
	return Define{
		Name: name,
		Table: Table{
			Name: name,
		},
		Fields:    Fields{},
		Extend:    ztype.Map{},
		Values:    ztype.Maps{},
		Relations: map[string]*ModelRelation{},
		Options:   ModelOptions{},
	}
}

func (d *Define) AddField(name string, field Field) error {
	if _, ok := d.Fields[name]; ok {
		return errors.New("field " + name + " already exists")
	}

	d.Fields[name] = field
	return nil
}

func (d *Define) GetField(name string) (Field, bool) {
	f, ok := d.Fields[name]
	if !ok {
		return Field{}, false
	}
	return f, true
}

func (d *Define) SetOptions(opt ModelOptions) {
	d.Options = opt
}

func (d *Define) GetOptions() ModelOptions {
	return d.Options
}

func (d *Defines) Append(define ...Define) {
	*d = append(*d, define...)
}
