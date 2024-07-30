package model

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/zlsgo/app_module/model/define"
)

type Models struct {
	m       *zarray.Maper[string, *Model]
	di      zdi.Injector
	storage Storageer
}

func NewModels(di zdi.Injector, s Storageer) *Models {
	return &Models{
		storage: s,
		di:      di,
		m:       zarray.NewHashMap[string, *Model](),
	}
}

func (ms *Models) String() string {
	return "[" + strings.Join(ms.m.Keys(), ", ") + "]"
}

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

func (ms *Models) MustGet(alias string) *Model {
	m, ok := ms.m.Get(alias)
	if !ok {
		panic("model " + alias + " not found")
	}
	return m
}

func (ms *Models) Storage() Storageer {
	return ms.storage
}

func (ms *Models) ForEach(fn func(key string, m *Model) bool) {
	ms.m.ForEach(fn)
}

func (ms *Models) Reg(name string, data define.Define, force bool) (*Model, error) {
	if name == "" {
		return nil, errors.New("model name can not be empty")
	}

	if !force && ms.m.Has(name) {
		return nil, errors.New("model " + name + " has been registered")
	}

	var tablePrefix string
	if s, ok := ms.storage.(*SQL); ok {
		tablePrefix = s.Options.Prefix
	}

	m := &Model{
		Storage:     ms.storage,
		model:       data,
		di:          ms.di,
		tablePrefix: tablePrefix,
	}
	err := ms.set(name, m, force)
	if err != nil {
		err = zerror.With(err, "model "+name+" register error")
		return nil, err
	}

	if m.Define().Options.DisabledMigrator {
		return m, nil
	}

	err = m.Migration().Auto(InsideOption.oldColumn)
	if err != nil {
		err = zerror.With(err, "model "+name+" migration error")
		return nil, err
	}

	return m, nil
}

func (ms *Models) BatchReg(models map[string]define.Define, force bool) error {
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
