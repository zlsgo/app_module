package model

import (
	"errors"
	"strings"
	"sync"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/zlsgo/app_module/model/schema"
)

type Schemas struct {
	data          *zarray.Maper[string, *Schema]
	di            zdi.Injector
	storage       Storageer
	getWrapModels func() []*Store
	models        *Stores
	SchemaOption  SchemaOptions
	cacheGet      map[string]*Schema
	mu            sync.RWMutex
}

// NewSchemas 创建新的模型集合实例
func NewSchemas(di zdi.Injector, s Storageer, o SchemaOptions) *Schemas {
	return &Schemas{
		storage:      s,
		di:           di,
		data:         zarray.NewHashMap[string, *Schema](),
		SchemaOption: o,
		cacheGet:     make(map[string]*Schema),
	}
}

// String 返回模型集合的字符串表示
func (ss *Schemas) String() string {
	return "[" + strings.Join(ss.data.Keys(), ", ") + "]"
}

// set 设置模型到集合中
func (ss *Schemas) set(alias string, s *Schema, force ...bool) (err error) {
	if s.define.Table.Name == "" {
		tableName := strings.Replace(alias, "-", "_", -1)
		tableName = strings.Replace(tableName, "::", "__", -1)
		s.define.Table.Name = tableName
	}

	err = perfect(alias, s, &ss.SchemaOption)

	ss.mu.Lock()
	ss.data.Set(alias, s)
	ss.cacheGet[alias] = s
	if ss.models != nil {
		ss.models.items.Set(alias, s.Model())
	}
	ss.mu.Unlock()
	return
}

// Get 获取指定别名的模型
func (ss *Schemas) Get(alias string) (*Schema, bool) {
	ss.mu.RLock()
	if cached, exists := ss.cacheGet[alias]; exists {
		ss.mu.RUnlock()
		return cached, true
	}

	s, ok := ss.data.Get(alias)
	ss.mu.RUnlock()
	if !ok && ss.getWrapModels != nil {
		for _, m := range ss.getWrapModels() {
			if alias == m.schema.GetAlias() {
				ss.mu.Lock()
				ss.data.Set(alias, m.schema)
				ss.cacheGet[alias] = m.schema
				if ss.models != nil {
					ss.models.items.Set(alias, m.schema.Model())
				}
				ss.mu.Unlock()
				return m.schema, true
			}
		}
	}

	if ok {
		ss.mu.Lock()
		ss.cacheGet[alias] = s
		ss.mu.Unlock()
	}

	return s, ok
}

// MustGet 获取指定别名的模型，不存在则 panic
func (ss *Schemas) MustGet(alias string) *Schema {
	m, ok := ss.Get(alias)
	if !ok {
		panic("models " + alias + " not found")
	}
	return m
}

// Models 返回模型存储集合
func (ss *Schemas) Models() *Stores {
	ss.mu.Lock()
	if ss.models == nil {
		ss.models = &Stores{items: zarray.NewHashMap[string, *Store]()}
		ss.data.ForEach(func(key string, m *Schema) bool {
			ss.models.items.Set(key, m.Model())
			return true
		})
	}
	models := ss.models
	ss.mu.Unlock()
	return models
}

// Storage 返回存储实例
func (ss *Schemas) Storage() Storageer {
	return ss.storage
}

// ForEach 遍历所有模型
func (ss *Schemas) ForEach(fn func(key string, m *Schema) bool) {
	ss.mu.RLock()
	items := make([]struct {
		key string
		m   *Schema
	}, 0, ss.data.Len())
	ss.data.ForEach(func(key string, m *Schema) bool {
		items = append(items, struct {
			key string
			m   *Schema
		}{key: key, m: m})
		return true
	})
	ss.mu.RUnlock()

	for _, item := range items {
		if !fn(item.key, item.m) {
			return
		}
	}
}

// Reg 注册模型到集合中
func (ss *Schemas) Reg(name string, data schema.Schema, force bool) (*Schema, error) {
	if name == "" {
		return nil, errors.New("models name can not be empty")
	}

	if !force && ss.data.Has(name) {
		return nil, errors.New("models " + name + " has been registered")
	}

	var tablePrefix string
	if ss.storage != nil {
		if opts := ss.storage.GetOptions(); opts != nil {
			prefixVal := opts.Get("prefix")
			if prefixVal.Exists() {
				tablePrefix = prefixVal.String()
			}
		}
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
		if ss.storage != nil {
			migration := m.Migration()
			if migration.HasTable() {
				if mFields, err := migration.GetFields(); err == nil {
					inlayFields := zarray.Keys(mFields)
					m.inlayFields = zarray.Unique(append(m.inlayFields, inlayFields...))
					m.fullFields = zarray.Unique(append(m.fullFields, inlayFields...))
				}
			}
		}
		m.refreshFieldsSet()
		return m, nil
	}

	if ss.storage != nil {
		err = m.Migration().Auto(ss.SchemaOption.OldColumn)
		if err != nil {
			err = zerror.With(err, "models "+name+" migration error")
			return nil, err
		}
	}

	return m, nil
}

// BatchReg 批量注册模型
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
