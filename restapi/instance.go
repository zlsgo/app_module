package restapi

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zerror"
)

type Models struct {
	m *zarray.Maper[string, *Model]
	// db      *zdb.DB
	storage       Storageer
	ModelsOptions ModelsOptions
}

type ModelsOptions struct {
	// 前缀
	Prefix string
}

func NewModels(s Storageer, opt ...func(*ModelsOptions)) *Models {
	o := ModelsOptions{}
	for _, v := range opt {
		v(&o)
	}

	return &Models{
		storage:       s,
		m:             zarray.NewHashMap[string, *Model](),
		ModelsOptions: o,
	}
}

// func (ms *Models) DB() *zdb.DB {
// 	return ms.db
// }

func (ms *Models) set(alias string, m *Model, force ...bool) (err error) {
	if m.model.Table.Name == "" {
		tableName := strings.Replace(alias, "-", "_", -1)
		tableName = strings.Replace(tableName, "::", "__", -1)
		m.model.Table.Name = tableName
	}

	err = perfect(alias, m)

	ms.m.Set(alias, m)
	return
}

func (ms *Models) Get(alias string) (*Model, bool) {
	return ms.m.Get(alias)
}

func (ms *Models) Storage() Storageer {
	return ms.storage
}

func (ms *Models) ForEach(fn func(key string, m *Model) bool) {
	ms.m.ForEach(fn)
}

func (ms *Models) Reg(name string, data Define, force bool) (*Model, error) {
	if !force && ms.m.Has(name) {
		return nil, errors.New("model " + name + " has been registered")
	}

	m := &Model{
		Storage:     ms.storage,
		tablePrefix: ms.ModelsOptions.Prefix,
		model:       data,
	}
	err := ms.set(name, m, force)
	if err != nil {
		err = zerror.With(err, "model "+name+" register error")
		return nil, err
	}

	if m.Define().Options.DisabledMigrator {
		return m, nil
	}

	err = m.Migration().Auto(Inside.oldColumn)
	if err != nil {
		err = zerror.With(err, "model "+name+" migration error")
		return nil, err
	}

	return m, nil
}

func (ms *Models) BatchReg(models map[string]Define, force bool) error {
	for name, data := range models {
		err := zerror.TryCatch(func() error {
			_, err := ms.Reg(name, data, force)
			return err
		})
		if err != nil {
			return zerror.With(err, "register "+name+" model error")
		}
	}
	return nil
}