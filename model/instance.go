package model

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/zlsgo/app_module/model/schema"
)

type Schemas struct {
	data          *zarray.Maper[string, *Schema]
	di            zdi.Injector
	storage       Storageer
	getWrapModels func() []*Model
	models        *Models
}

func NewSchemas(di zdi.Injector, s Storageer) *Schemas {
	return &Schemas{
		storage: s,
		di:      di,
		data:    zarray.NewHashMap[string, *Schema](),
	}
}

func (ms *Schemas) String() string {
	return "[" + strings.Join(ms.data.Keys(), ", ") + "]"
}

func (ms *Schemas) StorageType() string {
	switch ms.storage.(type) {
	case *SQL:
		return "sql"
	default:
		return "unknown"
	}
}

func (ms *Schemas) set(alias string, m *Schema, force ...bool) (err error) {
	if m.define.Table.Name == "" {
		tableName := strings.Replace(alias, "-", "_", -1)
		tableName = strings.Replace(tableName, "::", "__", -1)
		m.define.Table.Name = tableName
	}

	err = perfect(alias, m)

	ms.data.Set(alias, m)
	return
}

func (ms *Schemas) Get(alias string) (*Schema, bool) {
	s, ok := ms.data.Get(alias)
	if !ok && ms.getWrapModels != nil {
		for _, m := range ms.getWrapModels() {
			if alias == m.schema.GetAlias() {
				ms.data.Set(alias, m.schema)
				return m.schema, true
			}
		}
	}

	return s, ok
}

func (ms *Schemas) MustGet(alias string) *Schema {
	m, ok := ms.Get(alias)
	if !ok {
		panic("model " + alias + " not found")
	}
	return m
}

func (ms *Schemas) Models() *Models {
	if ms.models == nil {
		ms.models = &Models{items: zarray.NewHashMap[string, *Model]()}
		ms.ForEach(func(key string, m *Schema) bool {
			ms.models.items.Set(key, m.Model())
			return true
		})
	}

	return ms.models
}

func (ms *Schemas) Storage() Storageer {
	return ms.storage
}

func (ms *Schemas) ForEach(fn func(key string, m *Schema) bool) {
	ms.data.ForEach(fn)
}

func (ms *Schemas) Reg(name string, data schema.Schema, force bool) (*Schema, error) {
	if name == "" {
		return nil, errors.New("model name can not be empty")
	}

	if !force && ms.data.Has(name) {
		return nil, errors.New("model " + name + " has been registered")
	}

	var tablePrefix string
	if s, ok := ms.storage.(*SQL); ok {
		tablePrefix = s.Options.Prefix
	}

	m := &Schema{
		Storage:     ms.storage,
		define:      data,
		di:          ms.di,
		getSchema:   ms.Get,
		tablePrefix: tablePrefix,
	}

	err := ms.set(name, m, force)
	if err != nil {
		err = zerror.With(err, "model "+name+" register error")
		return nil, err
	}

	if m.GetDefine().Options.DisabledMigrator {
		migration := m.Migration()
		if migration.HasTable() {
			go func() {
				if mFields, err := migration.GetFields(); err == nil {
					inlayFields := zarray.Keys(mFields)
					m.inlayFields = zarray.Unique(append(m.inlayFields, inlayFields...))
					m.fullFields = zarray.Unique(append(m.fullFields, inlayFields...))
				}
			}()
		}
		return m, nil
	}

	err = m.Migration().Auto(InsideOption.oldColumn)
	if err != nil {
		err = zerror.With(err, "model "+name+" migration error")
		return nil, err
	}

	return m, nil
}

func (ms *Schemas) BatchReg(models map[string]schema.Schema, force bool) error {
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
