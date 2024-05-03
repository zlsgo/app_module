package sqlstorage

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/app_module/quick/define"
	"github.com/zlsgo/app_module/quick/process"
	"github.com/zlsgo/app_module/quick/storage"
	"github.com/zlsgo/app_module/quick/utils"
	"github.com/zlsgo/zdb/builder"
	"github.com/zlsgo/zdb/schema"
)

func (s *SQL) Migration(table string, d *define.Define, p *process.Process) (err error) {
	if table == "" {
		return errors.New("表名不能为空")
	}

	//
	// if err = m.Model.hook("migrationStart"); err != nil {
	// 	return err
	// }

	exist := s.HasTable(table)

	defer func() {
		if err != nil {
			return
		}

		err = s.Indexs(table, d)
		if err == nil {
			err = s.InitValue(table, d, p, !exist)
		}
		//
		// if err == nil {
		// 	err = s.Model.hook("migrationDone")
		// }
	}()

	if !exist {
		err = s.CreateTable(table, d)
		return
	}

	err = s.UpdateTable(table, d, d.Options.MigrationOldColumn)

	return
}

func (s *SQL) HasTable(table string) bool {
	t := builder.NewTable(table).Create()
	t.SetDriver(s.db.GetDriver())

	sql, values, p := t.Has()
	res, err := s.db.QueryToMaps(sql, values...)
	if err != nil {
		return false
	}

	return p(res)
}

func (s *SQL) Indexs(table string, d *define.Define) error {
	t := builder.NewTable(table).Create()
	t.SetDriver(s.db.GetDriver())

	modelFields := d.Fields
	uniques := make(map[string][]string)
	indexs := make(map[string][]string)
	for name := range modelFields {
		field := modelFields[name]
		unique := ztype.ToString(field.Unique)
		if unique != "" {
			if unique == "true" {
				unique = name
			}
			uniques[unique] = append(uniques[unique], name)
		}

		index := ztype.ToString(field.Index)
		if index != "" {
			if index == "true" {
				index = name
			}
			indexs[index] = append(indexs[index], name)
		}
	}

	for name, v := range uniques {
		name = table + "__unique__" + name
		sql, values, p := t.HasIndex(name)
		res, err := s.db.QueryToMaps(sql, values...)

		if err == nil && !p(res) {
			sql, values := t.CreateIndex(name, v, "UNIQUE")
			_, err = s.db.Exec(sql, values...)
			if err != nil {
				return err
			}
		}
	}

	for name, v := range indexs {
		name = table + "__idx__" + name
		sql, values, p := t.HasIndex(name)
		res, err := s.db.QueryToMaps(sql, values...)
		if err == nil && !p(res) {
			sql, values := t.CreateIndex(name, v, "")
			_, err = s.db.Exec(sql, values...)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getPrimaryKey() *schema.Field {
	return schema.NewField(define.Inside.IDKey(), schema.Uint, func(f *schema.Field) {
		f.Comment = "ID"
		f.PrimaryKey = true
		f.AutoIncrement = true
	})
}

func fillField(d *define.Define, fields []*schema.Field) []*schema.Field {
	if d.Options.SoftDeletes {
		fields = append(fields, schema.NewField(define.Inside.DeletedAtKey(), schema.Int, func(f *schema.Field) {
			f.Size = 9999999999
			f.NotNull = false
			f.Comment = "删除时间"
		}))
	}

	if d.Options.Timestamps {
		fields = append(fields, schema.NewField(define.Inside.CreatedAtKey(), schema.Time, func(f *schema.Field) {
			f.Comment = "创建时间"
		}))
		fields = append(fields, schema.NewField(define.Inside.UpdatedAtKey(), schema.Time, func(f *schema.Field) {
			f.Comment = "更新时间"
		}))
	}
	return fields
}

func (s *SQL) CreateTable(table string, d *define.Define) error {
	t := builder.NewTable(table).Create()
	t.SetDriver(s.db.GetDriver())
	fields := make([]*schema.Field, 0, len(d.Fields))
	fields = append(fields, getPrimaryKey())
	for name := range d.Fields {
		if utils.IsDisableMigratioField(d, name) {
			continue
		}

		field := d.Fields[name]
		f := schema.NewField(name, field.Type, func(f *schema.Field) {
			f.Comment = zutil.IfVal(field.Comment != "", field.Comment, field.Label).(string)
			f.NotNull = !field.Nullable
			f.Size = field.Size
		})
		fields = append(fields, f)
	}

	if len(fields) == 0 {
		return errors.New("表字段不能为空")
	}

	fields = fillField(d, fields)

	if len(d.Options.FieldsSort) > 0 {
		for _, n := range d.Options.FieldsSort {
			for i := range fields {
				f := &fields[i]
				if (*f).Name == n {
					t.Column(*f)
					fields = append(fields[:i], fields[i+1:]...)
					break
				}
			}
		}
	}

	t.Column(fields...)

	sql, values, err := t.Build()

	if err != nil {
		return err
	}

	_, err = s.db.Exec(sql, values...)

	return err
}

func (s *SQL) execAddColumn(deleteColumn bool, modelFields define.Fields, v string, table *builder.TableBuilder, oldColumns []string) error {
	var (
		ok    bool
		field *define.Field
	)

	for name := range modelFields {
		if name == v {
			ok = true
			var f = modelFields[name]
			field = &f
			break
		}
	}

	if !ok {
		return nil
	}

	sql, values := table.AddColumn(v, field.Type, func(f *schema.Field) {
		f.Comment = zutil.IfVal(field.Comment != "", field.Comment, field.Label).(string)
		f.NotNull = !field.Nullable
		f.Size = field.Size
	})

	if !deleteColumn {
		recovery := define.Inside.DeleteFieldPrefix() + v
		_, ok := zarray.Find(oldColumns, func(i int, n string) bool {
			return n == recovery
		})
		if ok {
			sql, values = table.RenameColumn(recovery, v)
		}
	}

	_, err := s.db.Exec(sql, values...)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQL) UpdateTable(table string, d *define.Define, oldColumn define.MigrationOldColumn) error {
	t := builder.NewTable(table)
	t.SetDriver(s.db.GetDriver())
	sql, values, process := t.GetColumn()
	res, err := s.db.QueryToMaps(sql, values...)
	if err != nil {
		return err
	}

	newColumns := zarray.Keys(d.Fields)
	newColumns = append(newColumns, define.Inside.IDKey())

	currentColumns := process(res)
	oldColumns := zarray.Keys(currentColumns)

	{
		if d.Options.SoftDeletes {
			newColumns = append(newColumns, define.Inside.DeletedAtKey())
		}

		if d.Options.Timestamps {
			if zarray.Contains(oldColumns, define.Inside.CreatedAtKey()) {
				newColumns = append(newColumns, define.Inside.CreatedAtKey())
			}
			if zarray.Contains(oldColumns, define.Inside.UpdatedAtKey()) {
				newColumns = append(newColumns, define.Inside.UpdatedAtKey())
			}
		}
	}

	addColumns := zarray.Filter(newColumns, func(_ int, n string) bool {
		return !zarray.Contains(oldColumns, n)
	})

	deleteColumns := zarray.Filter(oldColumns, func(_ int, n string) bool {
		return !zarray.Contains(newColumns, n) && !strings.HasPrefix(n, define.Inside.DeleteFieldPrefix())
	})

	for _, v := range deleteColumns {
		if oldColumn == define.DealOldColumnNone || utils.IsDisableMigratioField(d, v) {
			continue
		}
		if oldColumn == define.DealOldColumnDelete {
			sql, values = t.DropColumn(v)
		} else if oldColumn == define.DealOldColumnRename {
			sql, values = t.RenameColumn(v, define.Inside.DeleteFieldPrefix()+v)
		}

		_, err := s.db.Exec(sql, values...)
		if err != nil {
			return err
		}
	}

	if d.Options.SoftDeletes {
		if !zarray.Contains(oldColumns, define.Inside.DeletedAtKey()) {
			sql, values := t.AddColumn(define.Inside.DeletedAtKey(), "int", func(f *schema.Field) {
				f.Comment = "删除时间戳"
				f.NotNull = false
			})
			_, err := s.db.Exec(sql, values...)
			if err != nil {
				return err
			}
		}
	}

	if d.Options.Timestamps {
		if !zarray.Contains(oldColumns, define.Inside.CreatedAtKey()) {
			sql, values := t.AddColumn(define.Inside.CreatedAtKey(), "time", func(f *schema.Field) {
				f.Comment = "更新时间"

			})
			_, err := s.db.Exec(sql, values...)
			if err != nil {
				return err
			}
		}
		if !zarray.Contains(oldColumns, define.Inside.UpdatedAtKey()) {
			sql, values := t.AddColumn(define.Inside.UpdatedAtKey(), "time", func(f *schema.Field) {
				f.Comment = "更新时间"
			})
			_, err := s.db.Exec(sql, values...)
			if err != nil {
				return err
			}
		}
	}

	deleteColumn := oldColumn == define.DealOldColumnDelete
	if len(addColumns) > 0 {
		if len(d.Options.FieldsSort) > 0 {
			for _, n := range d.Options.FieldsSort {
				for i, v := range addColumns {
					if v != n {
						continue
					}
					if err := s.execAddColumn(deleteColumn, d.Fields, v, t, oldColumns); err != nil {
						return err
					}
					addColumns = append(addColumns[:i], addColumns[i+1:]...)
					break
				}
			}
		}
		for _, v := range addColumns {
			if utils.IsDisableMigratioField(d, v) {
				continue
			}

			if err := s.execAddColumn(deleteColumn, d.Fields, v, t, oldColumns); err != nil {
				return err
			}
		}
	}

	// TODO: 是否需要支持修改字段类型
	// if len(updateColumns) > 0 {
	// 	zlog.Warn("暂不支持修改字段类型：", updateColumns)
	// }

	return nil
}

func (s *SQL) InitValue(table string, d *define.Define, p *process.Process, first bool) error {
	if !first {
		rows, err := s.Find(table, ztype.Map{}, func(o storage.CondOptions) storage.CondOptions {
			o.Fields = []string{"COUNT(*) AS count"}
			return o
		})

		if err == nil {
			first = rows.Index(0).Get("count").Int() == 0
		}
	}

	var err error
	for _, data := range d.Values {
		if !first {
			if _, ok := data[define.Inside.IDKey()]; ok {
				continue
			}
		}

		data, err = p.InsertData(d, data)
		if err != nil {
			return zerror.With(err, "initial data value error")
		}

		if _, err = s.Insert(table, data); err != nil {
			return zerror.With(err, "initial data error")
		}
	}

	return nil
}
