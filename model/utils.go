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
		if !strings.ContainsRune(f[i], '.') {
			f[i] = table + f[i]
		}
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

func initModels(m *Module, di zdi.Invoker) (err error) {
	var (
		db  *zdb.DB
		opt = &m.Options
	)
	if opt.SetDB != nil {
		db, err = opt.SetDB()
	} else {
		err = di.Resolve(&db)
	}

	if err != nil {
		return zerror.With(err, "init db error")
	}

	m.Schemas = NewSchemas(di.(zdi.Injector), NewSQL(db, func(o *SQLOptions) {
		o.Prefix = m.Options.Prefix
	}))

	mapper := di.(zdi.TypeMapper)
	m.Models = &Models{items: zarray.NewHashMap[string, *Model]()}

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
			return errors.New("model name can not be empty, schema path: " + d.SchemaPath)
		}

		s, err := m.Schemas.Reg(d.Name, d, false)
		if err != nil {
			return err
		}

		m.Models.items.Set(d.Name, s.Model())
	}

	if opt.SetAlternateModels != nil {
		m.Schemas.getWrapModels = zutil.Once(func() []*Model {
			lists, err := opt.SetAlternateModels()
			if err != nil {
				panic(err)
			}
			return lists
		})
	}

	_ = mapper.Maps(m.Schemas, m.Models)

	zlog.Debugf("Models %s\n", m.Schemas)

	return nil
}
