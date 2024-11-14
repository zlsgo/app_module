package schema

import (
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb/builder"
)

type RelationType string

const (
	// RelationSingle 单对单
	RelationSingle RelationType = "single"
	// RelationSingleMerge 单对单，结果合并
	RelationSingleMerge RelationType = "single_merge"
	// RelationMany 单对多
	RelationMany RelationType = "many"
)

type (
	Relations map[string]Relation
	Relation  struct {
		Filter     ztype.Map          `json:"filter,omitempty"`
		Comment    string             `json:"comment,omitempty"`
		Label      string             `json:"label"`
		Type       RelationType       `json:"type"`
		Join       builder.JoinOption `json:"-"`
		Schema     string             `json:"schema"`
		ForeignKey []string           `json:"foreign_key"`
		SchemaKey  []string           `json:"schema_key"`
		Fields     []string           `json:"fields,omitempty"`
		Nullable   bool               `json:"nullable,omitempty"`
	}
)
