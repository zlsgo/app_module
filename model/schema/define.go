package schema

import (
	"errors"

	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/zdb/builder"
)

type (
	// Schemas 模型集合类型
	Schemas []Schema
	// Schema 模型结构定义
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

	// Table 表结构定义
	Table struct {
		Name    string `json:"name,omitempty"`
		Comment string `json:"comment,omitempty"`
	}

	// Validations 验证规则定义
	Validations struct {
		Args    interface{} `json:"args"`
		Method  string      `json:"method"`
		Message string      `json:"message,omitempty"`
		// Trigger ValidTriggerType `json:"trigger"`
	}
)

// RelationType 关系类型
type RelationType string

const (
	// RelationSingle 单一关系
	RelationSingle RelationType = "single"
	// RelationSingleMerge 单一合并关系
	RelationSingleMerge RelationType = "single_merge"
	// RelationMany 多重关系
	RelationMany RelationType = "many"
	// RelationManyToMany 多对多关系
	RelationManyToMany RelationType = "many_to_many"
)

// CascadeType 级联操作类型
type CascadeType string

const (
	// CascadeTypeCascade 级联删除
	CascadeTypeCascade CascadeType = "CASCADE"
	// CascadeTypeSetNull 设置为空
	CascadeTypeSetNull CascadeType = "SET_NULL"
	// CascadeTypeRestrict 限制删除
	CascadeTypeRestrict CascadeType = "RESTRICT"
)

// PivotKeys 中间表键定义
type PivotKeys struct {
	Foreign []string `json:"foreign"`
	Related []string `json:"related"`
}

type (
	// Relations 关系集合类型
	Relations map[string]Relation
	// Relation 关系定义
	Relation  struct {
		Filter      ztype.Map          `json:"filter,omitempty"`
		Comment     string             `json:"comment,omitempty"`
		Label       string             `json:"label"`
		Type        RelationType       `json:"type"`
		Join        builder.JoinOption `json:"-"`
		Schema      string             `json:"schema"`
		ForeignKey  []string           `json:"foreign_key"`
		SchemaKey   []string           `json:"schema_key"`
		Keys        []string           `json:"keys,omitempty"`
		ForeignKeys []string           `json:"foreign_keys,omitempty"`
		Fields      []string           `json:"fields,omitempty"`
		Nullable    bool               `json:"nullable,omitempty"`

		PivotTable  string    `json:"pivot_table,omitempty"`
		PivotKeys   PivotKeys `json:"pivot_keys,omitempty"`
		PivotFields []string  `json:"pivot_fields,omitempty"`
		PivotFilter ztype.Map `json:"pivot_filter,omitempty"`

		Cascade     string      `json:"cascade,omitempty"`
		CascadeType CascadeType `json:"cascade_type,omitempty"`

		Options ztype.Map `json:"options,omitempty"`
		Inverse string    `json:"inverse,omitempty"`
	}
)

// New 创建新的模型结构定义
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

// SetComment 设置表注释
func (d *Schema) SetComment(comment string) {
	d.Table.Comment = comment
}

// AddField 添加字段到模型定义
func (d *Schema) AddField(name string, field Field) error {
	if err := d.exists(name); err != nil {
		return err
	}

	d.Fields[name] = field
	return nil
}

// GetField 获取指定名称的字段
func (d *Schema) GetField(name string) (Field, bool) {
	f, ok := d.Fields[name]
	if !ok {
		return Field{}, false
	}
	return f, true
}

// SetOptions 设置模型选项
func (d *Schema) SetOptions(fn ...func(opt *Options)) *Options {
	d.Options = zutil.Optional(d.Options, fn...)
	return &d.Options
}

// GetOptions 获取模型选项
func (d *Schema) GetOptions() Options {
	return d.Options
}

// exists 检查字段或关系是否已存在
func (d *Schema) exists(name string) error {
	if _, ok := d.GetField(name); ok {
		return errors.New("field " + name + " already exists")
	}
	if _, ok := d.Relations[name]; ok {
		return errors.New("relation " + name + " already exists")
	}
	return nil
}

// AddRelation 添加关系到模型定义
func (d *Schema) AddRelation(name string, relation Relation) error {
	if err := d.exists(name); err != nil {
		return err
	}
	d.Relations[name] = relation
	return nil
}

// Append 追加模型定义到集合
func (d *Schemas) Append(define ...Schema) {
	*d = append(*d, define...)
}
