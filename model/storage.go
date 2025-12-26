package model

import (
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/builder"
)

// StorageType 存储类型
type StorageType uint8

const (
	// SQLStorage SQL 存储
	SQLStorage StorageType = iota + 1
	// NoSQLStorage NoSQL 存储
	NoSQLStorage
)

// String 返回存储类型的字符串表示
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

// StorageJoin 连接查询配置
type StorageJoin struct {
	Table        string
	As           string
	Expr         string
	ModelOptions builder.JoinOption
}

// StorageModelOptions 存储模型选项
type StorageModelOptions struct{}

// CondOptions 查询条件选项
type CondOptions struct {
	// 查询字段 默认查询所有字段，如果字段包含空格那么会跳过追加表名前缀
	Fields    []string
	Relations []string
	GroupBy   []string
	OrderBy   []OrderByItem
	Join      []StorageJoin
	Limit     int
	Offset    int
}

// InsertOptions 插入选项
type InsertOptions struct {
	Options string
}

type Storageer interface {
	GetStorageType() StorageType
	GetOptions() ztype.Map
	Transaction(run func(s Storageer) error) (err error)
	Find(table string, filter ztype.Map, fn ...func(*CondOptions)) (ztype.Maps, error)
	First(table string, filter ztype.Map, fn ...func(*CondOptions)) (ztype.Map, error)
	Pages(table string, page, pagesize int, filter ztype.Map, fn ...func(*CondOptions)) (ztype.Maps, PageInfo, error)
	Migration(model *Schema) Migrationer
	Insert(table string, data ztype.Map, fn ...func(*InsertOptions)) (lastId interface{}, err error)
	InsertMany(table string, data ztype.Maps, fn ...func(*InsertOptions)) (lastIds []interface{}, err error)
	Delete(table string, filter ztype.Map, fn ...func(*CondOptions)) (int64, error)
	Update(table string, data ztype.Map, filter ztype.Map, fn ...func(*CondOptions)) (int64, error)
}

// PageInfo 分页信息
type PageInfo struct {
	zdb.Pages
}

type Migrationer interface {
	Auto(deleteColumn ...DealOldColumn) (err error)
	HasTable() bool
	GetFields() (ztype.Map, error)
}
