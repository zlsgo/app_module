package schema

import (
	"errors"

	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zutil"
)

type (
	Schemas []Schema
	Schema  struct {
		Options    Options    `json:"options,omitempty"`
		Fields     Fields     `json:"fields"`
		Extend     ztype.Map  `json:"extend,omitempty"`
		Relations  Relations  `json:"relations,omitempty"`
		Table      Table      `json:"table,omitempty"`
		Name       string     `json:"name"`
		SchemaPath string     `json:"-"`
		Values     ztype.Maps `json:"values,omitempty"`
	}

	Table struct {
		Name    string `json:"name,omitempty"`
		Comment string `json:"comment,omitempty"`
	}

	Validations struct {
		Args    interface{} `json:"args"`
		Method  string      `json:"method"`
		Message string      `json:"message,omitempty"`
		// Trigger ValidTriggerType `json:"trigger"`
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
		Relations: map[string]Relation{},
		Options:   Options{},
	}
}

func (d *Schema) SetComment(comment string) {
	d.Table.Comment = comment
}

func (d *Schema) AddField(name string, field Field) error {
	if err := d.exists(name); err != nil {
		return err
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

func (d *Schema) SetOptions(fn ...func(opt *Options)) *Options {
	d.Options = zutil.Optional(d.Options, fn...)
	return &d.Options
}

func (d *Schema) GetOptions() Options {
	return d.Options
}

func (d *Schema) exists(name string) error {
	if _, ok := d.GetField(name); ok {
		return errors.New("field " + name + " already exists")
	}
	if _, ok := d.Relations[name]; ok {
		return errors.New("relation " + name + " already exists")
	}
	return nil
}

func (d *Schema) AddRelation(name string, relation Relation) error {
	if err := d.exists(name); err != nil {
		return err
	}
	d.Relations[name] = relation
	return nil
}

func (d *Schemas) Append(define ...Schema) {
	*d = append(*d, define...)
}
