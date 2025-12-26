package model

import (
	"errors"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/builder"
)

func fillFilterTablePrefix(f ztype.Map, table string) ztype.Map {
	if table == "" || len(f) == 0 {
		return f
	}

	result := make(ztype.Map, len(f))
	for k, v := range f {
		if k == "" {
			result[k] = v
			continue
		}

		if !strings.ContainsRune(k, '.') {
			result[table+k] = v
		} else {
			result[k] = v
		}
	}

	return result
}

func fillFieldsTablePrefix(f []string, table string) []string {
	if table == "" || len(f) == 0 {
		return f
	}

	result := make([]string, len(f))
	for i, field := range f {
		if strings.ContainsRune(field, '.') || strings.ContainsRune(field, ' ') {
			result[i] = field
		} else {
			result[i] = table + field
		}
	}

	return result
}

func parseSchema(dir string) ([]schema.Schema, error) {
	var files []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			if strings.ToLower(filepath.Ext(path)) != ".json" {
				return nil
			}
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	schemaModelsDefine := make([]schema.Schema, 0, len(files))
	for _, filePath := range files {
		text, err := zfile.ReadFile(filePath)
		if err != nil {
			return nil, zerror.With(err, "read schema file: "+filePath)
		}

		var d schema.Schema
		if err := zjson.Unmarshal(text, &d); err != nil {
			return nil, zerror.With(err, "invalid schema file: "+filePath)
		}

		d.SchemaPath = filePath
		schemaModelsDefine = append(schemaModelsDefine, d)
	}

	return schemaModelsDefine, nil
}

func parseExprsBuildCond(d *builder.BuildCond, value interface{}, exprs []string) ([]string, error) {
	var expr string
	switch val := value.(type) {
	case func(*builder.BuildCond) string:
		expr = val(d)
	case func() string:
		expr = val()
	default:
		return nil, errors.New("unknown type")
	}

	if expr != "" {
		exprs = append(exprs, expr)
	}
	return exprs, nil
}

func initModels(m *Module, di zdi.Invoker) (err error) {
	opt := &m.Options

	var storageer Storageer
	if opt.SetStorageer != nil {
		if storageer, err = opt.SetStorageer(); err != nil {
			return zerror.With(err, "init storageer error")
		}
	} else {
		var db *zdb.DB
		if err = di.Resolve(&db); err != nil {
			return zerror.With(err, "please set db")
		}
		storageer = NewSQL(db, opt.Prefix)
	}

	injector, ok := di.(zdi.Injector)
	if !ok {
		return errors.New("di injector not supported")
	}
	mapper, ok := di.(zdi.TypeMapper)
	if !ok {
		return errors.New("di type mapper not supported")
	}

	m.schemas = NewSchemas(injector, storageer, opt.SchemaOptions)
	m.stores = &Stores{items: zarray.NewHashMap[string, *Store]()}

	if opt.SchemaDir != "" {
		schemaModelsDefine, err := parseSchema(opt.SchemaDir)
		if err != nil {
			return err
		}

		opt.Schemas = append(opt.Schemas, schemaModelsDefine...)
	}

	for i := range opt.Schemas {
		d := opt.Schemas[i]

		if d.Name == "" && d.SchemaPath != "" {
			return errors.New("models name can not be empty, schema path: " + d.SchemaPath)
		}

		s, err := m.schemas.Reg(d.Name, d, false)
		if err != nil {
			return err
		}

		m.stores.items.Set(d.Name, s.Model())
	}

	if opt.SetAlternateModels != nil {
		m.schemas.getWrapModels = zutil.Once(func() []*Store {
			lists, err := opt.SetAlternateModels()
			if err != nil {
				panic(err)
			}
			return lists
		})
	}

	_ = mapper.Maps(m.schemas, m.stores)

	zlog.Debugf("Models %s\n", m.schemas)

	return nil
}
