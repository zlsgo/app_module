package model

import (
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/builder"
)

const (
	SQLStorage StorageType = iota + 1
	NoSQLStorage
)

type (
	StorageType uint8
	StorageJoin struct {
		Table        string
		As           string
		Expr         string
		ModelOptions builder.JoinOption
	}
)

// type StorageWhere struct {
// 	Expr string
// 	// Cond  string
// 	Field string
// 	Value interface{}
// }

type StorageModelOptions struct{}

type CondOptions struct {
	Fields  []string
	GroupBy []string
	OrderBy map[string]string
	Join    []StorageJoin
	Limit   int
	// Relations []string
}

type InsertOptions struct {
	Options string
}

type Storageer interface {
	GetStorageType() StorageType
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
