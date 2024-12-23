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
	if table == "" {
		return f
	}

	for k := range f {
		if k == "" {
			continue
		}
		if !strings.ContainsRune(k, '.') {
			f[table+k] = f[k]
			delete(f, k)
		}
	}

	return f
}

func fillFieldsTablePrefix(f []string, table string) []string {
	if table == "" {
		return f
	}

	for i := range f {
		if strings.ContainsRune(f[i], '.') || strings.ContainsRune(f[i], ' ') {
			continue
		}

		f[i] = table + f[i]
	}

	return f
}

func parseSchema(dir string) ([]schema.Schema, error) {
	files := make([]string, 0)
	_ = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})

	schemaModelsDefine := zarray.Map(files, func(_ int, v string) (d schema.Schema) {
		text, err := zfile.ReadFile(v)
		if err != nil {
			return
		}
		zjson.Unmarshal(text, &d)
		d.SchemaPath = v
		return
	}, 10)

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

	m.schemas = NewSchemas(di.(zdi.Injector), storageer, opt.SchemaOptions)

	mapper := di.(zdi.TypeMapper)
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
