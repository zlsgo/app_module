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
	"github.com/zlsgo/app_module/model/define"
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

func parseSchema(dir string) ([]define.Schema, error) {
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

	schemaModelsDefine := zarray.Map(files, func(_ int, v string) (d define.Schema) {
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
	if opt.GetDB == nil {
		err = di.Resolve(&db)
	} else {
		db, err = opt.GetDB()
	}

	if err != nil {
		return zerror.With(err, "init db error")
	}

	mod := NewSchemas(di.(zdi.Injector), NewSQL(db, func(o *SQLOptions) {
		o.Prefix = m.Options.Prefix
	}))

	mapper := di.(zdi.TypeMapper)
	opers := &Models{items: zarray.NewHashMap[string, *Model]()}

	if opt.SchemaDir != "" {
		schemaModelsDefine, err := parseSchema(opt.SchemaDir)
		if err != nil {
			return err
		}

		opt.Schemas = append(opt.Schemas, schemaModelsDefine...)
	}

	for i := range opt.Schemas {
		d := opt.Schemas[i]
		if opt.DisabledMigrator {
			d.Options.DisabledMigrator = true
		}

		if d.Name == "" && d.SchemaPath != "" {
			return errors.New("model name can not be empty, schema path: " + d.SchemaPath)
		}

		m, err := mod.Reg(d.Name, d, false)
		if err != nil {
			return err
		}

		opers.items.Set(d.Name, m.Operation())
	}

	_ = mapper.Maps(mod, opers)

	m.Models = mod
	m.Operations = opers

	zlog.Debugf("Models %s\n", mod)

	return nil
}
