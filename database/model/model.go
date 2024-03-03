package model

import "github.com/sohaha/zlsgo/zarray"

func (m *Model) Name() string {
	return m.model.Name
}

func (m *Model) Alias() string {
	return m.alias
}

func (m *Model) Define() Define {
	return m.model
}

func (m *Model) TableName() string {
	return m.tablePrefix + m.model.Table.Name
}

func (m *Model) Migration() Migrationer {
	return m.Storage.Migration(m)
}

func (m *Model) GetFields(exclude ...string) []string {
	f := m.fullFields
	if len(exclude) == 0 {
		exclude = m.Define().Options.LowFields
		if len(exclude) == 0 {
			return f
		}
	}

	return zarray.Filter(f, func(_ int, v string) bool {
		return !zarray.Contains(exclude, v)
	})
}
