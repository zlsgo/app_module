package restapi

import (
	"github.com/zlsgo/zdb"
)

type SQL struct {
	db *zdb.DB
}

func NewSQL(db *zdb.DB) Storageer {
	return &SQL{
		db: db,
	}
}

func (s *SQL) GetStorageType() StorageType {
	return SQLStorage
}

func (s *SQL) Migration(model *Model) Migrationer {
	return &Migration{
		Model: model,
		DB:    s.db,
	}
}

func sqlOrderBy(orderBy [][]string) (o []string) {
	l := len(orderBy)
	if l == 0 {
		return nil
	}

	o = make([]string, 0, l)
	for i := range orderBy {
		if len(orderBy[i]) < 2 {
			o = append(o, orderBy[i][0])
			continue
		}
		switch orderBy[i][1] {
		case "-1":
			o = append(o, orderBy[i][0]+" DESC")
		case "1", "0":
			o = append(o, orderBy[i][0]+" ASC")
		default:
			o = append(o, orderBy[i][0]+" "+orderBy[i][1])
		}
	}
	return
}
