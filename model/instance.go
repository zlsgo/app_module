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
	SchemaOption  SchemaOptions
}

func NewSchemas(di zdi.Injector, s Storageer, o SchemaOptions) *Schemas {
	return &Schemas{
		storage:      s,
		di:           di,
		data:         zarray.NewHashMap[string, *Schema](),
		SchemaOption: o,
	}
}

func (ss *Schemas) String() string {
	return "[" + strings.Join(ss.data.Keys(), ", ") + "]"
}

func (ss *Schemas) set(alias string, s *Schema, force ...bool) (err error) {
	if s.define.Table.Name == "" {
		tableName := strings.Replace(alias, "-", "_", -1)
		tableName = strings.Replace(tableName, "::", "__", -1)
		s.define.Table.Name = tableName
	}

	err = perfect(alias, s, &ss.SchemaOption)

	ss.data.Set(alias, s)
	return
}

func (ss *Schemas) Get(alias string) (*Schema, bool) {
	s, ok := ss.data.Get(alias)
	if !ok && ss.getWrapModels != nil {
		for _, m := range ss.getWrapModels() {
			if alias == m.schema.GetAlias() {
				ss.data.Set(alias, m.schema)
				return m.schema, true
			}
		}
	}

	return s, ok
}

func (ss *Schemas) MustGet(alias string) *Schema {
	m, ok := ss.Get(alias)
	if !ok {
		panic("models " + alias + " not found")
	}
	return m
}

func (ss *Schemas) Models() *Models {
	if ss.models == nil {
		ss.models = &Models{items: zarray.NewHashMap[string, *Model]()}
		ss.ForEach(func(key string, m *Schema) bool {
			ss.models.items.Set(key, m.Model())
			return true
		})
	}

	return ss.models
}

func (ss *Schemas) Storage() Storageer {
	return ss.storage
}

func (ss *Schemas) ForEach(fn func(key string, m *Schema) bool) {
	ss.data.ForEach(fn)
}

func (ss *Schemas) Reg(name string, data schema.Schema, force bool) (*Schema, error) {
	if name == "" {
		return nil, errors.New("models name can not be empty")
	}

	if !force && ss.data.Has(name) {
		return nil, errors.New("models " + name + " has been registered")
	}

	var tablePrefix string
	if s, ok := ss.storage.(*SQL); ok {
		tablePrefix = s.Options.Prefix
	}

	m := &Schema{
		Storage:     ss.storage,
		define:      data,
		di:          ss.di,
		getSchema:   ss.Get,
		tablePrefix: tablePrefix,
	}

	err := ss.set(name, m, force)
	if err != nil {
		err = zerror.With(err, "models "+name+" register error")
		return nil, err
	}

	if *m.GetDefine().Options.DisabledMigrator {
		migration := m.Migration()
		if migration.HasTable() {
			if mFields, err := migration.GetFields(); err == nil {
				inlayFields := zarray.Keys(mFields)
				m.inlayFields = zarray.Unique(append(m.inlayFields, inlayFields...))
				m.fullFields = zarray.Unique(append(m.fullFields, inlayFields...))
			}
		}
		return m, nil
	}

	err = m.Migration().Auto(InsideOption.oldColumn)
	if err != nil {
		err = zerror.With(err, "models "+name+" migration error")
		return nil, err
	}

	return m, nil
}

func (ss *Schemas) BatchReg(models map[string]schema.Schema, force bool) error {
	for name, data := range models {
		err := zerror.TryCatch(func() error {
			_, err := ss.Reg(name, data, force)
			return err
		})
		if err != nil {
			return zerror.With(err, "register "+name+" models error")
		}
	}
	return nil
}
