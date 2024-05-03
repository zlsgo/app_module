package sqlstorage

import (
	"strings"

	"github.com/zlsgo/app_module/quick/storage"
	"github.com/zlsgo/zdb"
)

type SQL struct {
	db *zdb.DB
}

var _ storage.Storage = (*SQL)(nil)

func NewSQL(db *zdb.DB) storage.Storage {
	return &SQL{
		db: db,
	}
}

func (s *SQL) Type() storage.StorageType {
	return "SQL"
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
