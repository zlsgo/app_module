package model

import (
	"errors"
	"strings"
	"time"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/sohaha/zlsgo/zvalid"
	"github.com/zlsgo/app_module/model/define"
	"github.com/zlsgo/zdb/schema"
)

func (m *Schema) filterFields(fields []string) []string {
	return zarray.Filter(fields, func(_ int, f string) bool {
		f = zstring.TrimSpace(f)
		if strings.ContainsRune(f, '(') || strings.ContainsRune(f, ' ') {
			return true
		}
		return zarray.Contains(m.fullFields, f)
	})
}

func (m *Schema) GetField(name string) (define.Field, bool) {
	f, ok := m.getField(name)
	if !ok {
		return define.Field{}, false
	}
	return *f, true
}

func (m *Schema) getField(name string) (*define.Field, bool) {
	for fname := range m.define.Fields {
		if name == fname {
			field := m.define.Fields[fname]
			return &field, true
		}
	}

	if name == idKey {
		return &define.Field{
			Type:     schema.Int,
			Nullable: false,
			Label:    "ID",
			Options: define.FieldOption{
				ReadOnly: true,
			},
		}, true
	}
	if m.define.Options.Timestamps {
		switch name {
		case CreatedAtKey:
			return &define.Field{
				Type:     schema.Time,
				Nullable: true,
				Label:    "创建时间",
				Options: define.FieldOption{
					ReadOnly: true,
				},
			}, true
		case UpdatedAtKey:
			return &define.Field{
				Type:     schema.Time,
				Nullable: true,
				Label:    "更新时间",
			}, true
		}
	}

	if m.define.Options.SoftDeletes {
		if name == DeletedAtKey {
			return &define.Field{
				Type:     schema.Int,
				Size:     11,
				Nullable: true,
				Default:  0,
				Label:    "删除时间戳",
			}, true
		}
	}

	// if m.model.Options.CreatedBy {
	// 	if name == CreatedByKey {
	// 		return &Field{
	// 			Type:     schema.String,
	// 			Nullable: true,
	// 			Default:  "",
	// 			Size:     120,
	// 			ModelOptions: FieldOption{
	// 				ReadOnly: true,
	// 			},
	// 			Label: "创建人 ID"}, true
	// 	}
	// }

	return nil, false
}

func (m *Schema) GetModelFields() define.Fields {
	return m.define.Fields
}

func (m *Schema) isInlayField(field string) bool {
	inlayFields := []string{idKey}
	if m.define.Options.Timestamps {
		inlayFields = append(inlayFields, CreatedAtKey, UpdatedAtKey)
	}
	// if m.model.Options.CreatedBy {
	// 	inlayFields = append(inlayFields, CreatedByKey)
	// }
	if m.define.Options.SoftDeletes {
		inlayFields = append(inlayFields, DeletedAtKey)
	}
	return zarray.Contains(inlayFields, field)
}

func perfectField(m *Schema) ([]string, error) {
	fields := make([]string, 0, len(m.define.Fields))
	if len(m.JSON) > 0 {
		j := zjson.ParseBytes(m.JSON).Get("fields")
		j.ForEach(func(key, _ *zjson.Res) bool {
			fields = append(fields, key.String())
			return true
		})
	}

	nFields := make(define.Fields, len(m.define.Fields))
	for name := range m.define.Fields {
		field := m.define.Fields[name]
		if err := parseField(m, name, &field); err != nil {
			return nil, err
		}

		fields = append(fields, name)
		nFields[name] = field
	}
	m.define.Fields = nFields

	return fields, nil
}

func parseField(m *Schema, name string, f *define.Field) error {
	if f == nil {
		return nil
	}

	if f.Default != nil {
		f.Nullable = true
	}

	if f.Label == "" {
		f.Label = name
	}

	if f.Options.ReadOnly {
		m.readOnlyKeys = append(m.readOnlyKeys, name)
	}

	switch f.Type {
	case schema.Bool:
		f.Before = append(f.Before, "bool")
		f.After = append(f.After, "bool")
	case schema.JSON:
		jsonProcess := zutil.IfVal(f.Options.IsArray, "jsons", "json").(string)
		f.Before = append(f.Before, jsonProcess)
		f.After = append(f.After, jsonProcess)
	case schema.Time:
		format := f.Options.FormatTime
		if format == "" {
			format = "date|Y-m-d H:i:s"
		} else {
			format = "date|" + format
		}

		f.Before = append(f.Before, format)
		f.After = append(f.After, format)
	}

	if f.Options.Crypt != "" {
		p, err := m.GetCryptProcess(f.Options.Crypt)
		if err != nil {
			return err
		}
		m.cryptKeys[name] = p
	}

	if len(f.Before) > 0 {
		ps, err := m.GetBeforeProcess(f.Before)
		if err != nil {
			return err
		}
		m.beforeProcess[name] = ps
	}

	if len(f.After) > 0 {
		ps, err := m.GetAfterProcess(f.Before)
		if err != nil {
			return err
		}
		m.afterProcess[name] = ps
	}

	parseFieldValidRule(name, f)
	parseFieldModelOptions(name, f)
	return nil
}

func parseFieldModelOptions(_ string, c *define.Field) {
	if len(c.Options.Enum) > 0 {
		c.Options.Enum = zarray.Map(c.Options.Enum, func(_ int, v define.FieldEnum) define.FieldEnum {
			if v.Label == "" {
				v.Label = v.Value
			}
			return v
		})

		c.ValidRules = c.ValidRules.EnumString(zarray.Map(c.Options.Enum, func(_ int, v define.FieldEnum) string {
			return v.Value
		}))
	}
}

func parseFieldValidRule(name string, c *define.Field) {
	label := c.Label
	rule := zvalid.New().SetAlias(label)
	if c.Type == schema.JSON {
		rule = rule.Required().IsJSON(name + "必须是JSON格式")
	}

	if c.Size > 0 {
		switch c.Type {
		case schema.JSON:
		case schema.String:
			rule = rule.MaxUTF8Length(int(c.Size))
		case schema.Int, schema.Int8, schema.Int16, schema.Int32, schema.Int64, schema.Uint, schema.Uint8, schema.Uint16, schema.Uint32, schema.Uint64:
			rule = rule.MaxInt(int(c.Size))
		case schema.Float:
			rule = rule.MaxFloat(float64(c.Size))
		case schema.Time:
			rule.Customize(func(rawValue string, err error) (newValue string, newErr error) {
				if err != nil {
					return "", err
				}
				if ztime.Unix(int64(c.Size)).After(time.Now()) {
					return rawValue, errors.New(label + "时间不能大于指定时间")
				}
				return
			})
		}
	}

	for _, valid := range c.Validations {
		switch valid.Method {
		case "regex":
			rule = rule.Regex(ztype.ToString(valid.Args), valid.Message)
		case "json":
			rule = rule.IsJSON(valid.Message)
		case "enum":
			switch val := valid.Args.(type) {
			case []float64:
				rule = rule.EnumFloat64(val)
			case []string:
				rule = rule.EnumString(val)
			case []int:
				rule = rule.EnumInt(val)
			default:
				rule = rule.Customize(func(rawValue string, err error) (string, error) {
					ok := zarray.Contains(ztype.ToSlice(val).String(), rawValue)
					if !ok {
						return "", errors.New(label + "枚举值不在合法范围")
					}
					return rawValue, nil
				})
			}
		case "mobile":
			rule = rule.IsMobile(valid.Message)
		case "mail":
			rule = rule.IsMail(valid.Message)
		case "url":
			rule = rule.IsURL(valid.Message)
		case "ip":
			rule = rule.IsIP(valid.Message)
		case "minLength":
			rule = rule.MinUTF8Length(ztype.ToInt(valid.Args), valid.Message)
		case "maxLength":
			rule = rule.MaxUTF8Length(ztype.ToInt(valid.Args), valid.Message)
		case "min":
			rule = rule.MinFloat(ztype.ToFloat64(valid.Args), valid.Message)
		case "max":
			rule = rule.MaxFloat(ztype.ToFloat64(valid.Args), valid.Message)
		}
	}

	c.ValidRules = rule
}

func isDisableMigratioField(m *Schema, name string) bool {
	for n := range m.define.Fields {
		if name != n {
			continue
		}
		if m.define.Fields[n].Options.DisableMigration {
			return true
		}
	}
	return false
}
