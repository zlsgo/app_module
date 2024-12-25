package model

import (
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/builder"
)

type StorageType uint8

const (
	SQLStorage StorageType = iota + 1
	NoSQLStorage
)

func (s StorageType) String() string {
	switch s {
	case SQLStorage:
		return "sql"
	case NoSQLStorage:
		return "nosql"
	default:
		return "unknown"
	}
}

type StorageJoin struct {
	Table        string
	As           string
	Expr         string
	ModelOptions builder.JoinOption
}

// type StorageWhere struct {
// 	Expr string
// 	// Cond  string
// 	Field string
// 	Value interface{}
// }

type StorageModelOptions struct{}

type CondOptions struct {
	// 查询字段 默认查询所有字段，如果字段包含空格那么会跳过追加表名前缀
	Fields  []string
	GroupBy []string
	OrderBy map[string]string
	Join    []StorageJoin
	Limit   int
	Offset  int
	// Relations []string
}

type InsertOptions struct {
	Options string
}

type Storageer interface {
	GetStorageType() StorageType
	GetOptions() ztype.Map
	Transaction(run func(s *SQL) error) (err error)
	Find(table string, fields []string, filter ztype.Map, fn ...func(*CondOptions)) (ztype.Maps, error)
	First(table string, fields []string, filter ztype.Map, fn ...func(*CondOptions)) (ztype.Map, error)
	Pages(table string, fields []string, page, pagesize int, filter ztype.Map, fn ...func(*CondOptions)) (ztype.Maps, PageInfo, error)
	Migration(model *Schema) Migrationer
	Insert(table string, fields []string, data ztype.Map, fn ...func(*InsertOptions)) (lastId interface{}, err error)
	InsertMany(table string, fields []string, data ztype.Maps, fn ...func(*InsertOptions)) (lastIds []interface{}, err error)
	Delete(table string, fields []string, filter ztype.Map, fn ...func(*CondOptions)) (int64, error)
	Update(table string, fields []string, data ztype.Map, filter ztype.Map, fn ...func(*CondOptions)) (int64, error)
}

type PageInfo struct {
	zdb.Pages
}

type Migrationer interface {
	Auto(deleteColumn ...DealOldColumn) (err error)
	HasTable() bool
	GetFields() (ztype.Map, error)
}
