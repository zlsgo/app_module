package model

import (
	"errors"
	"io/fs"
	"path/filepath"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/zlog"
	"github.com/zlsgo/app_module/model/define"
	"github.com/zlsgo/zdb"
)

func parseSchema(dir string) ([]define.Define, error) {
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

	schemaModelsDefine := zarray.Map(files, func(_ int, v string) (d define.Define) {
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

	mod := NewModels(di.(zdi.Injector), NewSQL(db, func(o *SQLOptions) {
		o.Prefix = m.Options.Prefix
	}))

	mapper := di.(zdi.TypeMapper)
	opers := &Operations{items: zarray.NewHashMap[string, *Operation]()}

	if opt.SchemaDir != "" {
		schemaModelsDefine, err := parseSchema(opt.SchemaDir)
		if err != nil {
			return err
		}

		opt.ModelsDefine = append(opt.ModelsDefine, schemaModelsDefine...)
	}

	for i := range opt.ModelsDefine {
		d := opt.ModelsDefine[i]
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
