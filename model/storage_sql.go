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

func (s *SQL) Transaction(run func(s *SQL) error) (err error) {
	return s.db.Transaction(func(db *zdb.DB) (err error) {
		return run(&SQL{
			db:      db,
			Options: s.Options,
		})
	})
}

func sqlOrderBy(orderBy map[string]string, fieldPrefix string) (o []string) {
	l := len(orderBy)
	if l == 0 {
		return nil
	}

	o = make([]string, 0, l)
	for n := range orderBy {
		v := orderBy[n]
		if fieldPrefix != "" && !strings.ContainsRune(n, '.') {
			n = fieldPrefix + n
		}
		switch orderBy[n] {
		case "-1":
			o = append(o, n+" DESC")
		case "1", "0":
			o = append(o, n+" ASC")
		default:
			o = append(o, n+" "+v)
		}
	}
	return
}
