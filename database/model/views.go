package model

import (
	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/ztype"
)

// type View struct {
// 	Disabled bool          `json:"disabled"`
// 	Title    string        `json:"title"`
// 	Fields   []string      `json:"Fields"`
// 	Layouts  ztype.Var     `json:"layouts"`
// 	Filters  []interface{} `json:"filters"`
// }

func (m *Model) GetViewFields(view string) []string {
	v := m.model.Extend.Get("views").Get(view).Map()
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
	return zarray.Unique(append(fields, IDKey))
}

func parseViewLists(m *Model) ztype.Map {
	columns := make(map[string]ztype.Map, 0)
	data := m.model.Extend.Get("views").Get("lists").Map()

	if data.Get("disabled").Bool() {
		return ztype.Map{}
	}

	fields := append([]string{IDKey}, data.Get("fields").Slice().String()...)
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
			"title":   column.Label,
			"type":    column.Type,
			"options": column.Options,
			"layout":  layout,
		}
		if m.model.Options.CryptID && name == IDKey {
			columns[name]["type"] = "string"
		}
	}

	title := data.Get("title").String()
	if title == "" {
		title = m.Name()
	}
	info := ztype.Map{
		"title":   title,
		"columns": columns,
		"fields":  fields,
	}

	return info
}

func parseViewInfo(m *Model) ztype.Map {
	info := ztype.Map{}

	data := m.model.Extend.Get("views").Get("info").Map()

	if data.Get("disabled").Bool() {
		return ztype.Map{}
	}

	columns := make(map[string]ztype.Map, 0)

	fields := append([]string{IDKey}, data.Get("fields").Slice().String()...)
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
			"label":    column.Label,
			"type":     column.Type,
			"readonly": column.Options.ReadOnly,
			"size":     column.Size,
			"layout":   layout,
			"disabled": m.isInlayField(name),
			"options":  column.Options,
		}

		if m.model.Options.CryptID && name == IDKey {
			columns[name]["type"] = "string"
		}
	}

	if m.model.Options.SoftDeletes {
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

func parseViews(m *Model) ztype.Map {
	views := ztype.Map{}

	views["lists"] = parseViewLists(m)

	views["info"] = parseViewInfo(m)
	return views
}

func (m *Model) GetViews() ztype.Map {
	return m.views
}
