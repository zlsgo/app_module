package model

import (
	"strings"

	"github.com/zlsgo/zdb"
)

type SQL struct {
	db      *zdb.DB
	Options SQLOptions
}
type SQLOptions struct {
	Prefix string
}

func NewSQL(db *zdb.DB, o ...func(*SQLOptions)) Storageer {
	opt := SQLOptions{
		Prefix: "model_",
	}
	for _, f := range o {
		f(&opt)
	}
	return &SQL{
		db:      db,
		Options: opt,
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
