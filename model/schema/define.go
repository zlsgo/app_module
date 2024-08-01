package schema

import (
	"errors"

	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb/builder"
)

type (
	Schemas []Schema
	Schema  struct {
		Fields     Fields                   `json:"fields"`
		Extend     ztype.Map                `json:"extend,omitempty"`
		Relations  map[string]ModelRelation `json:"relations,omitempty"`
		Table      Table                    `json:"table,omitempty"`
		Name       string                   `json:"name"`
		SchemaPath string                   `json:"-"`
		Values     ztype.Maps               `json:"values,omitempty"`
		Options    ModelOptions             `json:"options,omitempty"`
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
		Type       RelationType       `json:"type"`
		Join       builder.JoinOption `json:"-"`
		Schema     string             `json:"schema"`
		ForeignKey string             `json:"foreign_key"`
		SchemaKey  string             `json:"schema_key"`
		Fields     []string           `json:"fields,omitempty"`
		Relations  []string           `json:"relations,omitempty"`
		// Limit   int                `json:"limit,omitempty"`
	}
)

func New(name string, tableName ...string) Schema {
	table := name
	if len(tableName) > 0 {
		table = tableName[0]
	}

	return Schema{
		Name:      name,
		Table:     Table{Name: table},
		Fields:    Fields{},
		Extend:    ztype.Map{},
		Values:    ztype.Maps{},
		Relations: map[string]ModelRelation{},
		Options:   ModelOptions{},
	}
}

func (d *Schema) SetComment(comment string) {
	d.Table.Comment = comment
}

func (d *Schema) AddField(name string, field Field) error {
	if _, ok := d.Fields[name]; ok {
		return errors.New("field " + name + " already exists")
	}

	d.Fields[name] = field
	return nil
}

func (d *Schema) GetField(name string) (Field, bool) {
	f, ok := d.Fields[name]
	if !ok {
		return Field{}, false
	}
	return f, true
}

func (d *Schema) SetOptions(opt ModelOptions) {
	d.Options = opt
}

func (d *Schema) GetOptions() ModelOptions {
	return d.Options
}

func (d *Schemas) Append(define ...Schema) {
	*d = append(*d, define...)
}
