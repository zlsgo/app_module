package crud

import (
	"strings"

	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/quick/define"
	"github.com/zlsgo/app_module/quick/storage"
	"github.com/zlsgo/app_module/quick/utils"
	"github.com/zlsgo/zdb"
	"golang.org/x/exp/constraints"
)

type Filter interface {
	ztype.Map | constraints.Integer | string
}

func getFilter[T Filter](m *Crud, filter T) (filterMap ztype.Map) {
	var ok bool
	f := (interface{})(filter)
	filterMap, ok = f.(ztype.Map)
	if !ok {
		idVal := f
		if m.define.Options.CryptID {
			if id, err := m.process.DeCryptID(ztype.ToString(filter)); err == nil {
				idVal = id
			}
		}

		filterMap = ztype.Map{
			define.Inside.IDKey(): idVal,
		}

	} else if filterMap == nil {
		filterMap = ztype.Map{}
	}

	if m.define.Options.SoftDeletes {
		filterMap[define.Inside.DeletedAtKey()] = 0
	}

	return
}

func find(m *Crud, filter ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (ztype.Maps, error) {
	_ = m.process.DeCrypt(filter)
	rows, err := m.storage.Find(m.tableName, filter, func(so storage.CondOptions) storage.CondOptions {
		so = utils.Optional(so, fn...)
		if len(so.Fields) > 0 && len(so.Join) == 0 {
			so.Fields = m.filterFields(so.Fields)
		} else if so.Limit != 1 && len(so.Fields) == 0 {
			so.Fields = m.GetFields()
		}
		return so
	})
	if err != nil {
		return rows, err
	}

	if len(m.process.AfterProcess) > 0 {
		for i := range rows {
			row := &rows[i]
			for k, v := range m.process.AfterProcess {
				if _, ok := (*row)[k]; ok {
					(*row)[k], err = v[0](row.Get(k).String())
					if err != nil {
						return nil, err
					}
				}
			}
			_ = m.process.EnCrypt(row)
		}
	}
	return rows, nil
}

func (crud *Crud) Find(filter ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (ztype.Maps, error) {
	return find(crud, getFilter(crud, filter), fn...)
}

func (crud *Crud) FindOne(filter ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (ztype.Map, error) {
	rows, err := find(crud, getFilter(crud, filter), func(so storage.CondOptions) storage.CondOptions {
		so = utils.Optional(so, fn...)
		so.Limit = 1
		return so
	})

	if err != nil {
		return ztype.Map{}, err
	}

	return rows.Index(0), nil
}

func (crud *Crud) FindCols(field string, filter ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (ztype.SliceType, error) {
	rows, err := find(crud, getFilter(crud, filter), func(so storage.CondOptions) storage.CondOptions {
		so.Fields = []string{field}
		return utils.Optional(so, fn...)
	})
	if err != nil {
		return ztype.SliceType{}, err
	}

	data := make(ztype.SliceType, rows.Len())
	f := strings.Split(field, " ")
	field = f[len(f)-1]
	for i := range rows {
		data[i] = rows[i].Get(field)
	}
	return data, nil
}

func (crud *Crud) FindOneByID(id any, fn ...func(storage.CondOptions) storage.CondOptions) (ztype.Map, error) {
	return crud.FindOne(ztype.Map{define.Inside.IDKey(): id}, fn...)
}

func (crud *Crud) Count(filter ztype.Map) (int64, error) {
	row, err := crud.FindOne(filter, func(options storage.CondOptions) storage.CondOptions {
		options.Fields = []string{"count(*) as count"}
		return options
	})
	if err != nil {
		return 0, err
	}
	return row.Get("count").Int64(), nil
}

func (crud *Crud) Exists(filter ztype.Map) (bool, error) {
	tatol, err := crud.Count(filter)
	if err != nil {
		return false, err
	}
	return tatol > 0, nil
}

type PageData struct {
	Items ztype.Maps `json:"items"`
	Page  zdb.Pages  `json:"page"`
}

func (crud *Crud) Pages(page, pagesize int, filter ztype.Map, fn ...func(storage.CondOptions) storage.CondOptions) (*PageData, error) {
	f := getFilter(crud, filter)
	_ = crud.process.DeCrypt(f)

	rows, pages, err := crud.storage.Pages(crud.tableName, page, pagesize, f, func(so storage.CondOptions) storage.CondOptions {
		so = utils.Optional(so, fn...)
		if len(so.Fields) > 0 && len(so.Join) == 0 {
			so.Fields = crud.filterFields(so.Fields)
		}
		return so
	})

	data := &PageData{Items: rows, Page: pages}
	if err != nil {
		return data, err
	}

	afterProcess := crud.process.AfterProcess
	if len(afterProcess) == 0 {
		return data, nil
	}

	for i := range data.Items {
		row := &data.Items[i]
		for k, v := range afterProcess {
			if _, ok := (*row)[k]; ok {
				(*row)[k], err = v[0](row.Get(k).String())
				if err != nil {
					return data, err
				}
			}
		}
		if crud.define.Options.CryptID {
			err = crud.process.EnCrypt(row)
			if err != nil {
				return data, err
			}
		}
	}

	return data, nil
}
