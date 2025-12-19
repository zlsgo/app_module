package model

import (
	"strings"

	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb"
)

type SQL struct {
	db      *zdb.DB
	Options SQLOptions
}

type SQLOptions struct {
	prefix string
}

func NewSQL(db *zdb.DB, tablePrefix string, o ...func(*SQLOptions)) Storageer {
	opt := SQLOptions{
		prefix: tablePrefix,
	}
	for _, f := range o {
		f(&opt)
	}
	return &SQL{
		db:      db,
		Options: opt,
	}
}

func (s *SQL) GetOptions() ztype.Map {
	return ztype.Map{
		"prefix": s.Options.prefix,
	}
}

func (s *SQL) GetStorageType() StorageType {
	return SQLStorage
}

func (s *SQL) GetDB() *zdb.DB {
	return s.db
}

func (s *SQL) Migration(model *Schema) Migrationer {
	return &Migration{
		Model: model,
		DB:    s.db,
	}
}

func (s *SQL) Transaction(run func(s Storageer) error) (err error) {
	return s.db.Transaction(func(db *zdb.DB) (err error) {
		return run(&SQL{
			db:      db,
			Options: s.Options,
		})
	})
}

func sqlOrderBy(orderBy []OrderByItem, fieldPrefix string) (o []string) {
	l := len(orderBy)
	if l == 0 {
		return nil
	}

	o = make([]string, 0, l)
	hasPrefix := fieldPrefix != ""
	for _, item := range orderBy {
		field := item.Field
		if !isValidFieldName(field) {
			continue
		}
		if hasPrefix && !strings.ContainsRune(field, '.') {
			field = fieldPrefix + field
		}
		dir := strings.ToUpper(item.Direction)
		if dir != "ASC" && dir != "DESC" {
			dir = "ASC"
		}
		o = append(o, field+" "+dir)
	}
	return
}

func isValidFieldName(field string) bool {
	if field == "" {
		return false
	}
	for _, c := range field {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_' || c == '.') {
			return false
		}
	}
	return true
}
