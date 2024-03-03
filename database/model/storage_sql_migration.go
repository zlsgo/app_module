package model

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/builder"
	"github.com/zlsgo/zdb/schema"
)

type Migration struct {
	Model *Model
	DB    *zdb.DB
}

func (m *Migration) Auto(oldColumn DealOldColumn) (err error) {
	if m.Model.TableName() == "" {
		return errors.New("表名不能为空")
	}

	exist := m.HasTable()

	defer func() {
		if err != nil {
			return
		}

		err = m.Indexs()
		if err == nil {
			err = m.InitValue(!exist)
		}

		if err == nil && m.Model.model.MigrationDone != nil {
			err = m.Model.Define().MigrationDone(m.DB, m.Model)
		}
	}()

	if !exist {
		err = m.CreateTable()
		return
	}

	err = m.UpdateTable(oldColumn)

	return
}

func (m *Migration) InitValue(first bool) error {
	if !first {
		row, err := FindOne(m.Model, ztype.Map{}, func(o *CondOptions) error {
			o.Fields = []string{"COUNT(*) AS count"}
			return nil
		})
		if err == nil {
			first = row.Get("count").Int() == 0
		}
	}
	for _, data := range m.Model.model.Values {
		// data, ok := v.(map[string]interface{})
		// if !ok {
		// 	return errors.New("初始化数据格式错误")
		// }
		if !first {
			if _, ok := data[IDKey]; ok {
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
	table := builder.NewTable(m.Model.TableName()).Create()
	table.SetDriver(m.DB.GetDriver())
	sql, values, process := table.Has()
	res, err := m.DB.QueryToMaps(sql, values...)

	if err != nil {
		return false
	}

	return process(res)
}

func (m *Migration) UpdateTable(oldColumn DealOldColumn) error {
	table := builder.NewTable(m.Model.TableName())
	table.SetDriver(m.DB.GetDriver())
	sql, values, process := table.GetColumn()
	res, err := m.DB.QueryToMaps(sql, values...)
	if err != nil {
		return err
	}

	modelFields := m.Model.GetModelFields()
	newColumns := zarray.Keys(modelFields)
	newColumns = append(newColumns, IDKey)

	currentColumns := process(res)
	oldColumns := zarray.Keys(currentColumns)

	{
		if m.Model.model.Options.SoftDeletes {
			newColumns = append(newColumns, DeletedAtKey)
		}

		// if m.Model.model.Options.CreatedBy {
		// 	newColumns = append(newColumns, CreatedByKey)
		// }

		if m.Model.model.Options.Timestamps {
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

	for _, v := range deleteColumns {
		if oldColumn == DealOldColumnNone || isDisableMigratioField(m.Model, v) {
			continue
		}
		if oldColumn == DealOldColumnDelete {
			sql, values = table.DropColumn(v)
		} else if oldColumn == DealOldColumnRename {
			sql, values = table.RenameColumn(v, deleteFieldPrefix+v)
		}

		_, err := m.DB.Exec(sql, values...)
		if err != nil {
			return err
		}
	}

	if m.Model.model.Options.SoftDeletes {
		if !zarray.Contains(oldColumns, DeletedAtKey) {
			sql, values := table.AddColumn(DeletedAtKey, "int", func(f *schema.Field) {
				f.Comment = "删除时间戳"
			})
			_, err := m.DB.Exec(sql, values...)
			if err != nil {
				return err
			}
		}
	}

	// if m.Model.model.Options.CreatedBy {
	// 	if !zarray.Contains(oldColumns, CreatedByKey) {
	// 		sql, values := table.AddColumn(CreatedByKey, "string", func(f *schema.Field) {
	// 			f.Comment = "创建人 ID"
	// 			f.NotNull = false
	// 			f.Size = 120
	// 		})
	// 		_, err := m.DB.Exec(sql, values...)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	if m.Model.model.Options.Timestamps {
		if !zarray.Contains(oldColumns, CreatedAtKey) {
			sql, values := table.AddColumn(CreatedAtKey, "time", func(f *schema.Field) {
				f.Comment = "更新时间"

			})
			_, err := m.DB.Exec(sql, values...)
			if err != nil {
				return err
			}
		}
		if !zarray.Contains(oldColumns, UpdatedAtKey) {
			sql, values := table.AddColumn(UpdatedAtKey, "time", func(f *schema.Field) {
				f.Comment = "更新时间"
			})
			_, err := m.DB.Exec(sql, values...)
			if err != nil {
				return err
			}
		}
	}

	deleteColumn := oldColumn == DealOldColumnDelete
	if len(addColumns) > 0 {
		if len(m.Model.model.Options.FieldsSort) > 0 {
			for _, n := range m.Model.model.Options.FieldsSort {
				for i, v := range addColumns {
					if v != n {
						continue
					}
					if err := m.execAddColumn(deleteColumn, modelFields, v, table, oldColumns); err != nil {
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

			if err := m.execAddColumn(deleteColumn, modelFields, v, table, oldColumns); err != nil {
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

func (m *Migration) execAddColumn(deleteColumn bool, modelFields Fields, v string, table *builder.TableBuilder, oldColumns []string) error {
	var (
		ok    bool
		field *Field
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
		recovery := deleteFieldPrefix + v
		_, ok := zarray.Find(oldColumns, func(i int, n string) bool {
			return n == recovery
		})
		if ok {
			sql, values = table.RenameColumn(recovery, v)
		}
	}

	_, err := m.DB.Exec(sql, values...)
	if err != nil {
		return err
	}
	return nil
}

func (m *Migration) fillField(fields []*schema.Field) []*schema.Field {
	if m.Model.model.Options.SoftDeletes {
		fields = append(fields, schema.NewField(DeletedAtKey, schema.Int, func(f *schema.Field) {
			f.Size = 9999999999
			f.NotNull = false
			f.Comment = "删除时间"
		}))
	}

	if m.Model.model.Options.Timestamps {
		fields = append(fields, schema.NewField(CreatedAtKey, schema.Time, func(f *schema.Field) {
			f.Comment = "创建时间"
		}))
		fields = append(fields, schema.NewField(UpdatedAtKey, schema.Time, func(f *schema.Field) {
			f.Comment = "更新时间"
		}))
	}

	// if m.Model.model.Options.CreatedBy {
	// 	fields = append(fields, schema.NewField(CreatedByKey, schema.String, func(f *schema.Field) {
	// 		f.Comment = "创建人 ID"
	// 		f.NotNull = false
	// 		f.Size = 120
	// 	}))
	// }

	return fields
}

func (m *Migration) CreateTable() error {
	table := builder.NewTable(m.Model.TableName()).Create()
	table.SetDriver(m.DB.GetDriver())
	modelFields := m.Model.GetModelFields()
	fields := make([]*schema.Field, 0, len(modelFields))
	fields = append(fields, m.getPrimaryKey())
	for name := range modelFields {
		if isDisableMigratioField(m.Model, name) {
			continue
		}

		field := modelFields[name]
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

	fields = m.fillField(fields)

	if len(m.Model.model.Options.FieldsSort) > 0 {
		for _, n := range m.Model.model.Options.FieldsSort {
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
	_, err = m.DB.Exec(sql, values...)
	// if err == nil && len(sideFields) > 0 {
	// 	err = m.createSideTable(sideFields)
	// }

	return err
}

func (m *Migration) getPrimaryKey() *schema.Field {
	return schema.NewField(IDKey, schema.Uint, func(f *schema.Field) {
		f.Comment = "ID"
		f.PrimaryKey = true
		f.AutoIncrement = true
	})
}

func (m *Migration) Indexs() error {
	table := builder.NewTable(m.Model.TableName()).Create()
	table.SetDriver(m.DB.GetDriver())
	modelFields := m.Model.GetModelFields()
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

	for name, v := range uniques {
		name = m.Model.TableName() + "__unique__" + name
		sql, values, process := table.HasIndex(name)
		res, err := m.DB.QueryToMaps(sql, values...)

		if err == nil && !process(res) {
			sql, values := table.CreateIndex(name, v, "UNIQUE")
			_, err = m.DB.Exec(sql, values...)
			if err != nil {
				return err
			}
		}
	}

	for name, v := range indexs {
		name = m.Model.TableName() + "__idx__" + name
		sql, values, process := table.HasIndex(name)
		res, err := m.DB.QueryToMaps(sql, values...)
		if err == nil && !process(res) {
			sql, values := table.CreateIndex(name, v, "")
			_, err = m.DB.Exec(sql, values...)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
