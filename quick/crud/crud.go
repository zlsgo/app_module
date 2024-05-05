package crud

import (
	"errors"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/zlsgo/app_module/quick/define"
	"github.com/zlsgo/app_module/quick/process"
	"github.com/zlsgo/app_module/quick/sqlstorage"
	"github.com/zlsgo/app_module/quick/storage"
)

type Crud struct {
	define       *define.Define
	storage      storage.Storage
	options      Options
	tableName    string
	process      *process.Process
	inlayFields  []string
	Fields       []string `json:"-"`
	JSON         []byte
	fullFields   []string // 所有字段
	readOnlyKeys []string // 只读字段
}

func New(s storage.Storage, d define.Define, o ...func(Options) Options) (q *Crud, err error) {
	q = &Crud{
		define:  &d,
		storage: s,
		process: &process.Process{},
	}
	if len(o) > 0 {
		for i := range o {
			q.options = o[i](q.options)
		}
	}

	if err = setup(q); err != nil {
		return nil, err
	}

	if !q.define.Options.DisabledMigrator {
		if s.Type() == "SQL" {
			sqlStorage, ok := s.(*sqlstorage.SQL)
			if !ok {
				return nil, errors.New("sql storage not support")
			}
			err = sqlStorage.Migration(q.tableName, q.define, q.process)
			if err != nil {
				return nil, err
			}
		}
	}

	return
}

func (crud *Crud) GetFields(exclude ...string) []string {
	f := crud.fullFields
	if len(exclude) == 0 {
		exclude = crud.define.Options.ShowFields
		if len(exclude) == 0 {
			return f
		}
	}

	return zarray.Filter(f, func(_ int, v string) bool {
		return !zarray.Contains(exclude, v)
	})
}

func (crud *Crud) Define() define.Define {
	return *crud.define
}
