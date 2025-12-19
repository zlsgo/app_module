package model

import (
	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/ztype"
)

func (m *Schema) ParseLables(items ztype.Maps) ztype.Maps {
	return zarray.Map(items, func(_ int, v ztype.Map) ztype.Map {
		for k := range v {
			s, ok := m.getField(k)
			if !ok {
				continue
			}

			if len(s.Options.Enum) > 0 {
				for i := range s.Options.Enum {
					if s.Options.Enum[i].Value == v.Get(k).String() {
						v[k+"_label"] = s.Options.Enum[i].Label
						break
					}
				}
			}
		}
		return v
	}, 10)
}

func (m *Schema) GetViewFields(view string) []string {
	v := m.define.Extend.Get("views").Get(view).Map()
	if v.Get("disabled").Bool() {
		return []string{}
	}

	fields := v.Get("fields").Slice().String()
	if len(fields) == 0 {
		switch view {
		case "info":
			fields = m.fullFields
		default:
			fields = m.GetFields()
		}
	}
	return zarray.Unique(append(fields, idKey))
}

func parseViewLists(m *Schema) ztype.Map {
	columns := make(map[string]ztype.Map, 0)
	data := m.define.Extend.Get("views").Get("lists").Map()

	if data.Get("disabled").Bool() {
		return ztype.Map{}
	}

	fields := append([]string{idKey}, data.Get("fields").Slice().String()...)
	if len(fields) == 1 {
		fields = append(fields, m.GetFields()...)
	}
	fields = zarray.Unique(fields)

	for i := range fields {
		name := fields[i]
		column, ok := m.getField(name)
		if !ok {
			continue
		}
		layout := data.Get("layouts").Get(name).Map()
		columns[name] = ztype.Map{
			"title":        column.Label,
			"type":         column.Type,
			"ModelOptions": column.Options,
			"layout":       layout,
		}
		if *m.define.Options.CryptID && name == idKey {
			columns[name]["type"] = "string"
		}
	}

	title := data.Get("title").String()
	if title == "" {
		title = m.GetName()
	}
	info := ztype.Map{
		"title":   title,
		"columns": columns,
		"fields":  fields,
	}

	return info
}

func parseViewInfo(m *Schema) ztype.Map {
	info := ztype.Map{}

	data := m.define.Extend.Get("views").Get("info").Map()

	if data.Get("disabled").Bool() {
		return ztype.Map{}
	}

	columns := make(map[string]ztype.Map, 0)

	fields := append([]string{idKey}, data.Get("fields").Slice().String()...)
	if len(fields) == 1 {
		fields = append(fields, m.fullFields...)
	}
	fields = zarray.Unique(fields)

	layouts := data.Get("layouts").Map()
	for i := range fields {
		name := fields[i]
		column, ok := m.getField(name)
		if !ok {
			continue
		}
		layout := layouts.Get(name).Map()
		columns[name] = ztype.Map{
			"label":        column.Label,
			"type":         column.Type,
			"readonly":     column.Options.ReadOnly,
			"size":         column.Size,
			"layout":       layout,
			"disabled":     m.isInlayField(name),
			"ModelOptions": column.Options,
		}

		if *m.define.Options.CryptID && name == idKey {
			columns[name]["type"] = "string"
		}
	}

	if *m.define.Options.SoftDeletes {
		delete(columns, DeletedAtKey)
		fields = zarray.Filter(fields, func(_ int, v string) bool {
			return v != DeletedAtKey
		})
	}

	info["columns"] = columns
	info["fields"] = fields
	if !layouts.IsEmpty() {
		info["layouts"] = layouts
	}
	return info
}

func parseViews(m *Schema) ztype.Map {
	views := ztype.Map{}

	views["lists"] = parseViewLists(m)

	views["info"] = parseViewInfo(m)
	return views
}

func (m *Schema) GetViews() ztype.Map {
	return m.views
}
