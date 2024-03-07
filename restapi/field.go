package restapi

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
	"github.com/zlsgo/zdb/schema"
)

type (
	Fields map[string]Field
	Field  struct {
		Default     interface{}     `json:"default"`
		Unique      interface{}     `json:"unique"`
		Index       interface{}     `json:"index"`
		Comment     string          `json:"comment"`
		Label       string          `json:"label"`
		Type        schema.DataType `json:"type"`
		Validations []Validations   `json:"validations"`
		Options     FieldOption     `json:"ModelOptions"`
		Before      []string        `json:"before"`
		After       []string        `json:"after"`
		validRules  zvalid.Engine
		Size        uint64 `json:"size"`
		Nullable    bool   `json:"nullable"`
		// quoteName   string `json:"-"`
	}
	FieldEnum struct {
		Value string `json:"value"`
		Label string `json:"label"`
	}
	FieldOption struct {
		FormatTime       string      `json:"format_time"`
		Crypt            string      `json:"crypt"`
		Enum             []FieldEnum `json:"enum"`
		IsArray          bool        `json:"is_array"`
		ReadOnly         bool        `json:"readonly"`
		DisableMigration bool        `json:"disable_migration"`
		// Quote      bool        `json:"quote"`
	}
)

func (f *Field) GetValidations() *zvalid.Engine {
	return &f.validRules
}

func (m *Model) filterFields(fields []string) []string {
	return zarray.Filter(fields, func(_ int, f string) bool {
		f = zstring.TrimSpace(f)
		if strings.ContainsRune(f, '(') || strings.ContainsRune(f, ' ') {
			return true
		}
		return zarray.Contains(m.fullFields, f)
	})
}

func (m *Model) GetField(name string) (Field, bool) {
	f, ok := m.getField(name)
	if !ok {
		return Field{}, false
	}
	return *f, true
}

func (m *Model) getField(name string) (*Field, bool) {
	for fname := range m.model.Fields {
		if name == fname {
			field := m.model.Fields[fname]
			return &field, true
		}
	}

	if name == IDKey {
		return &Field{
			Type:     schema.Int,
			Nullable: false,
			Label:    "ID",
			Options: FieldOption{
				ReadOnly: true,
			},
		}, true

	}
	if m.model.Options.Timestamps {
		switch name {
		case CreatedAtKey:
			return &Field{
				Type:     schema.Time,
				Nullable: true,
				Label:    "创建时间"}, true
		case UpdatedAtKey:
			return &Field{
				Type:     schema.Time,
				Nullable: true,
				Label:    "更新时间"}, true
		}
	}

	if m.model.Options.SoftDeletes {
		if name == DeletedAtKey {
			return &Field{
				Type:    schema.Int,
				Size:    11,
				Default: 0,
				Label:   "删除时间戳"}, true
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

func (m *Model) GetModelFields() Fields {
	return m.model.Fields
}

func (m *Model) isInlayField(field string) bool {
	inlayFields := []string{IDKey}
	if m.model.Options.Timestamps {
		inlayFields = append(inlayFields, CreatedAtKey, UpdatedAtKey)
	}
	// if m.model.Options.CreatedBy {
	// 	inlayFields = append(inlayFields, CreatedByKey)
	// }
	if m.model.Options.SoftDeletes {
		inlayFields = append(inlayFields, DeletedAtKey)
	}
	return zarray.Contains(inlayFields, field)
}

func perfectField(m *Model) ([]string, error) {
	fields := make([]string, 0, len(m.model.Fields))
	if len(m.JSON) > 0 {
		j := zjson.ParseBytes(m.JSON).Get("fields")
		j.ForEach(func(key, _ *zjson.Res) bool {
			fields = append(fields, key.String())
			return true
		})
	}

	nFields := make(Fields, len(m.model.Fields))
	for name := range m.model.Fields {
		field := m.model.Fields[name]
		if err := parseField(m, name, &field); err != nil {
			return nil, err
		}

		fields = append(fields, name)
		nFields[name] = field
	}
	m.model.Fields = nFields

	return fields, nil
}

func parseField(m *Model, name string, f *Field) error {
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

func parseFieldModelOptions(name string, c *Field) {
	if len(c.Options.Enum) > 0 {
		c.Options.Enum = zarray.Map(c.Options.Enum, func(_ int, v FieldEnum) FieldEnum {
			if v.Label == "" {
				v.Label = v.Value
			}
			return v
		})

		c.validRules = c.validRules.EnumString(zarray.Map(c.Options.Enum, func(_ int, v FieldEnum) string {
			return v.Value
		}))
	}
}

func parseFieldValidRule(name string, c *Field) {
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

	c.validRules = rule
}

func isDisableMigratioField(m *Model, name string) bool {
	for n := range m.model.Fields {
		if name != n {
			continue
		}
		if m.model.Fields[n].Options.DisableMigration {
			return true
		}
	}
	return false
}
