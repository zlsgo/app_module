package model

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/app_module/model/hook"
	mSchema "github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/builder"
	"github.com/zlsgo/zdb/schema"
)

type Migration struct {
	Model *Schema
	DB    *zdb.DB
}

func (m *Migration) Auto(oldColumn ...DealOldColumn) (err error) {
	if m.Model.GetTableName() == "" {
		return errors.New("表名不能为空")
	}

	if err = m.Model.hook(hook.EventMigrationStart, m); err != nil {
		return err
	}

	var exist bool

	err = m.DB.Transaction(func(db *zdb.DB) (err error) {
		oldDB := m.DB
		m.DB = db
		defer func() { m.DB = oldDB }()

		if exist = m.HasTable(); !exist {
			err = m.CreateTable(db)
		} else {
			err = m.UpdateTable(db, oldColumn...)
		}
		if err != nil {
			return
		}

		return m.Indexs(db)
	})

	if err == nil {
		err = m.InitValue(!exist)
	}

	if err == nil {
		err = m.Model.hook(hook.EventMigrationDone, m)
	}
	return
}

func (m *Migration) InitValue(first bool) error {
	if !first {
		row, err := FindOne(m.Model, ztype.Map{}, func(o *CondOptions) {
			o.Fields = []string{"COUNT(*) AS count"}
		})
		if err == nil {
			first = row.Get("count").Int() == 0
		}
	}

	for _, data := range m.Model.define.Values {
		if !first {
			if _, ok := data[idKey]; ok {
				continue
			}
		}
		_, err := Insert(m.Model, data)
		if err != nil {
			return zerror.With(err, "初始化数据失败")
		}
	}

	return nil
}

func (m *Migration) HasTable() bool {
	table := builder.NewTable(m.Model.GetTableName()).Create()
	table.SetDriver(m.DB.GetDriver())
	sql, values, process := table.Has()
	res, err := m.DB.QueryToMaps(sql, values...)
	if err != nil {
		return false
	}

	return process(res)
}

func (m *Migration) GetFields() (ztype.Map, error) {
	table := builder.NewTable(m.Model.GetTableName())
	table.SetDriver(m.DB.GetDriver())
	sql, values, process := table.GetColumn()
	res, err := m.DB.QueryToMaps(sql, values...)
	if err != nil {
		return ztype.Map{}, err
	}

	return process(res), nil
}

func (m *Migration) UpdateTable(db *zdb.DB, oldColumn ...DealOldColumn) error {
	modelFields := m.Model.GetDefineFields()
	newColumns := zarray.Keys(modelFields)
	newColumns = append(newColumns, idKey)

	currentColumns, err := m.GetFields()
	if err != nil {
		return err
	}

	oldColumns := zarray.Keys(currentColumns)

	{
		if *m.Model.define.Options.SoftDeletes {
			newColumns = append(newColumns, DeletedAtKey)
		}

		// if m.Model.models.Options.CreatedBy {
		// 	newColumns = append(newColumns, CreatedByKey)
		// }

		if *m.Model.define.Options.Timestamps {
			if zarray.Contains(oldColumns, CreatedAtKey) {
				newColumns = append(newColumns, CreatedAtKey)
			}
			if zarray.Contains(oldColumns, UpdatedAtKey) {
				newColumns = append(newColumns, UpdatedAtKey)
			}
		}
	}

	// updateColumns := zarray.Var(zarray.Filter(m.Model.Columns, func(_ int, n *Column) bool {
	// 	c := currentColumns.Get(n.Name)
	// 	if !c.Exists() {
	// 		return false
	// 	}
	// 	nf := schema.NewField(n.Name, schema.DataType(n.Type), func(f *schema.Field) {
	// 		f.Size = n.Size
	// 	})
	// 	t := d.DataTypeOf(nf, true)
	// 	return !strings.EqualFold(t, c.Get("type").String())
	// }), func(i int, v *Column) string { return v.Name })

	addColumns := zarray.Filter(newColumns, func(_ int, n string) bool {
		return !zarray.Contains(oldColumns, n)
	})

	deleteColumns := zarray.Filter(oldColumns, func(_ int, n string) bool {
		return !zarray.Contains(newColumns, n) && !strings.HasPrefix(n, deleteFieldPrefix)
	})

	var dealOldColumn DealOldColumn
	if len(oldColumn) > 0 {
		dealOldColumn = oldColumn[0]
	}
	var (
		sql    string
		values []interface{}
		table  = builder.NewTable(m.Model.GetTableName())
	)

	table.SetDriver(db.GetDriver())

	for _, v := range deleteColumns {
		if dealOldColumn == dealOldColumnNone || isDisableMigratioField(m.Model, v) {
			continue
		}
		if dealOldColumn == dealOldColumnDelete {
			sql, values = table.DropColumn(v)
		} else if dealOldColumn == dealOldColumnRename {
			sql, values = table.RenameColumn(v, deleteFieldPrefix+v)
		}

		_, err := db.Exec(sql, values...)
		if err != nil {
			return err
		}
	}

	if *m.Model.define.Options.SoftDeletes {
		if !zarray.Contains(oldColumns, DeletedAtKey) {
			var (
				sql    string
				values []interface{}
			)
			if InsideOption.softDeleteIsTime {
				sql, values = table.AddColumn(DeletedAtKey, schema.Time, func(f *schema.Field) {
					f.Comment = "删除时间"
					f.NotNull = false
				})
			} else {
				sql, values = table.AddColumn(DeletedAtKey, schema.Uint, func(f *schema.Field) {
					f.Comment = "删除时间戳"
					f.NotNull = false
				})
			}
			_, err := db.Exec(sql, values...)
			if err != nil {
				return err
			}
		}
	}

	// if m.Model.models.Options.CreatedBy {
	// 	if !zarray.Contains(oldColumns, CreatedByKey) {
	// 		sql, values := table.AddColumn(CreatedByKey, "string", func(f *schema.Field) {
	// 			f.Comment = "创建人 ID"
	// 			f.NotNull = false
	// 			f.Size = 120
	// 		})
	// 		_, err := db.Exec(sql, values...)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	if *m.Model.define.Options.Timestamps {
		if !zarray.Contains(oldColumns, CreatedAtKey) {
			sql, values := table.AddColumn(CreatedAtKey, schema.Time, func(f *schema.Field) {
				f.Comment = "更新时间"
			})
			_, err := db.Exec(sql, values...)
			if err != nil {
				return err
			}
		}
		if !zarray.Contains(oldColumns, UpdatedAtKey) {
			sql, values := table.AddColumn(UpdatedAtKey, schema.Time, func(f *schema.Field) {
				f.Comment = "更新时间"
			})
			_, err := db.Exec(sql, values...)
			if err != nil {
				return err
			}
		}
	}

	deleteColumn := dealOldColumn == dealOldColumnDelete
	if len(addColumns) > 0 {
		if len(m.Model.define.Options.FieldsSort) > 0 {
			for _, n := range m.Model.define.Options.FieldsSort {
				for i, v := range addColumns {
					if v != n {
						continue
					}
					if err := m.execAddColumn(db, deleteColumn, modelFields, v, table, oldColumns); err != nil {
						return err
					}
					addColumns = append(addColumns[:i], addColumns[i+1:]...)
					break
				}
			}
		}
		for _, v := range addColumns {
			if isDisableMigratioField(m.Model, v) {
				continue
			}

			if err := m.execAddColumn(db, deleteColumn, modelFields, v, table, oldColumns); err != nil {
				return err
			}
		}
	}

	// TODO 是否需要支持修改字段类型
	// if len(updateColumns) > 0 {
	// 	zlog.Warn("暂不支持修改字段类型：", updateColumns)
	// }

	return nil
}

func (m *Migration) execAddColumn(
	db *zdb.DB,
	deleteColumn bool,
	modelFields mSchema.Fields,
	v string,
	table *builder.TableBuilder,
	oldColumns []string,
) error {
	var (
		ok    bool
		field *mSchema.Field
	)

	for name := range modelFields {
		if name == v {
			ok = true
			f := modelFields[name]
			field = &f
			break
		}
	}

	if !ok {
		return nil
	}

	sql, values := table.AddColumn(v, field.Type, func(f *schema.Field) {
		f.Comment = ztype.ToString(zutil.IfVal(field.Comment != "", field.Comment, field.Label))
		f.NotNull = !field.Nullable
		f.Size = field.Size
	})

	if !deleteColumn {
		recovery := deleteFieldPrefix + v
		_, ok := zarray.Find(oldColumns, func(i int, n string) bool {
			return n == recovery
		})
		if ok {
			sql, values = table.RenameColumn(recovery, v)
		}
	}

	_, err := db.Exec(sql, values...)
	if err != nil {
		return err
	}
	return nil
}

func (m *Migration) fillField(fields []*schema.Field) []*schema.Field {
	if *m.Model.define.Options.SoftDeletes {
		if InsideOption.softDeleteIsTime {
			fields = append(fields, schema.NewField(DeletedAtKey, schema.Time, func(f *schema.Field) {
				f.Size = 9999999999
				f.NotNull = false
				f.Comment = "删除时间"
			}))
		} else {
			fields = append(fields, schema.NewField(DeletedAtKey, schema.Int, func(f *schema.Field) {
				f.Size = 9999999999
				f.NotNull = false
				f.Comment = "删除时间"
			}))
		}
	}

	if *m.Model.define.Options.Timestamps {
		fields = append(fields, schema.NewField(CreatedAtKey, schema.Time, func(f *schema.Field) {
			f.Comment = "创建时间"
		}))
		fields = append(fields, schema.NewField(UpdatedAtKey, schema.Time, func(f *schema.Field) {
			f.Comment = "更新时间"
		}))
	}

	// if m.Model.models.Options.CreatedBy {
	// 	fields = append(fields, schema.NewField(CreatedByKey, schema.String, func(f *schema.Field) {
	// 		f.Comment = "创建人 ID"
	// 		f.NotNull = false
	// 		f.Size = 120
	// 	}))
	// }

	return fields
}

func (m *Migration) CreateTable(db *zdb.DB) error {
	table := builder.NewTable(m.Model.GetTableName()).Create()
	table.SetDriver(db.GetDriver())
	modelFields := m.Model.GetDefineFields()
	fields := make([]*schema.Field, 0, len(modelFields))
	fields = append(fields, m.getPrimaryKey())
	for name := range modelFields {
		if isDisableMigratioField(m.Model, name) {
			continue
		}

		field := modelFields[name]
		f := schema.NewField(name, field.Type, func(f *schema.Field) {
			f.Comment = ztype.ToString(zutil.IfVal(field.Comment != "", field.Comment, field.Label))
			f.NotNull = !field.Nullable
			f.Size = field.Size
		})
		fields = append(fields, f)
	}

	if len(fields) == 0 {
		return errors.New("表字段不能为空")
	}

	fields = m.fillField(fields)

	if len(m.Model.define.Options.FieldsSort) > 0 {
		for _, n := range m.Model.define.Options.FieldsSort {
			for i := range fields {
				f := &fields[i]
				if (*f).Name == n {
					table.Column(*f)
					fields = append(fields[:i], fields[i+1:]...)
					break
				}
			}
		}
	}

	table.Column(fields...)

	sql, values, err := table.Build()
	if err != nil {
		return err
	}

	_, err = db.Exec(sql, values...)

	return err
}

func (m *Migration) getPrimaryKey() *schema.Field {
	return schema.NewField(idKey, schema.Uint, func(f *schema.Field) {
		f.Comment = "ID"
		f.PrimaryKey = true
		f.AutoIncrement = true
	})
}

func (m *Migration) Indexs(db *zdb.DB) error {
	table := builder.NewTable(m.Model.GetTableName()).Create()
	table.SetDriver(db.GetDriver())

	modelFields := m.Model.GetDefineFields()
	uniques := make(map[string][]string, 0)
	indexs := make(map[string][]string, 0)
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

	if *m.Model.define.Options.SoftDeletes {
		indexs[DeletedAtKey] = []string{DeletedAtKey}
	}

	for name, v := range uniques {
		name = m.Model.GetTableName() + "__u__" + name
		sql, values, process := table.HasIndex(name)
		res, err := db.QueryToMaps(sql, values...)

		if err == nil && !process(res) {
			sql, values := table.CreateIndex(name, v, "UNIQUE")
			_, err = db.Exec(sql, values...)
			if err != nil {
				return err
			}
		}
	}

	for name, v := range indexs {
		name = m.Model.GetTableName() + "__i__" + name
		sql, values, process := table.HasIndex(name)
		res, err := db.QueryToMaps(sql, values...)
		if err == nil && !process(res) {
			sql, values := table.CreateIndex(name, v, "")
			_, err = db.Exec(sql, values...)
			if err != nil {
				return err
			}
		}
	}

	if err := m.Model.hook(hook.EventMigrationIndexDone, m, db, table); err != nil {
		return err
	}

	return nil
}
