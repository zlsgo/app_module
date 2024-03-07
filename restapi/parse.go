package restapi

import (
	"errors"

	"github.com/zlsgo/app_module/database/hashid"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zstring"
)

// perfect 完善模型
func perfect(alias string, m *Model) (err error) {
	m.alias = alias

	salt := ""
	salt = m.model.Options.Salt
	cryptLen := m.model.Options.CryptLen
	if cryptLen <= 0 {
		cryptLen = 12
	}
	m.Hashid = hashid.New(salt, cryptLen)

	m.readOnlyKeys = make([]string, 0)
	m.cryptKeys = make(map[string]CryptProcess, 0)
	m.afterProcess = make(map[string][]afterProcess, 0)
	m.beforeProcess = make(map[string][]beforeProcess, 0)

	m.Fields, err = perfectField(m)
	if err != nil {
		return
	}

	m.inlayFields = []string{IDKey}
	if m.model.Options.Timestamps {
		if zarray.Contains(m.Fields, CreatedAtKey) {
			err = errors.New(CreatedAtKey + " is a reserved field")
			return
		}
		if zarray.Contains(m.Fields, UpdatedAtKey) {
			err = errors.New(UpdatedAtKey + " is a reserved field")
			return
		}
		var after []afterProcess
		after, err = m.GetAfterProcess([]string{"date|Y-m-d H:i:s"})
		if err != nil {
			return
		}
		m.afterProcess[CreatedAtKey] = after
		m.afterProcess[UpdatedAtKey] = after
		m.inlayFields = append(m.inlayFields, CreatedAtKey, UpdatedAtKey)
	}

	// if m.model.Options.CreatedBy {
	// 	if zarray.Contains(m.Fields, CreatedByKey) {
	// 		err = errors.New(CreatedByKey + " is a reserved field")
	// 		return
	// 	}
	// 	m.inlayFields = append(m.inlayFields, CreatedByKey)
	// }

	if m.model.Options.SoftDeletes {
		if zarray.Contains(m.Fields, DeletedAtKey) {
			err = errors.New(DeletedAtKey + " is a reserved field")
			return
		}
		m.inlayFields = append(m.inlayFields, DeletedAtKey)
	}

	m.fullFields = append([]string{IDKey}, m.Fields...)
	m.fullFields = zarray.Unique(append(m.fullFields, m.inlayFields...))

	if m.model.Options.SoftDeletes {
		flen := len(m.fullFields)
		for i := 0; i < flen; i++ {
			f := m.fullFields[i]
			if f == DeletedAtKey {
				m.fullFields = append(m.fullFields[0:i], m.fullFields[i+1:]...)
				break
			}
		}
	}

	m.lowFields = m.model.Options.LowFields

	if len(m.model.Relations) > 0 {
		for k := range m.model.Relations {
			v := m.model.Relations[k]
			if v.Foreign == "" {
				m.model.Relations[k].Foreign = IDKey
			}
		}

		newRelations := make(map[string]*ModelRelation, len(m.model.Relations))
		for k := range m.model.Relations {
			v := m.model.Relations[k]
			newRelations[zstring.CamelCaseToSnakeCase(k)] = v
		}
		m.model.Relations = newRelations
	} else {
		m.model.Relations = make(map[string]*ModelRelation)
	}

	// if m.model.Options.CreatedBy {
	// 	c := &ModelRelation{
	// 		Key:     CreatedByKey,
	// 		Model:   define.UserModel,
	// 		Foreign: "_id",
	// 		Fields: []string{
	// 			"account",
	// 			"nickname",
	// 		},
	// 	}
	// 	m.model.Relations[zstring.SnakeCaseToCamelCase(CreatedByKey, true)] = c
	// }

	m.views = parseViews(m)
	return
}

// func parseColumn(m *Model, c *Column) {
// 	if c.Default != nil {
// 		c.Nullable = true
// 	}

// 	if c.ReadOnly {
// 		m.readOnlyKeys = append(m.readOnlyKeys, c.Name)
// 	}

// 	if c.Type == schema.JSON {
// 		if len(c.Before) == 0 {
// 			c.Before = []string{"json"}
// 		}
// 		if len(c.After) == 0 {
// 			c.After = []string{"json"}
// 		}
// 	}

// 	if c.Crypt != "" {
// 		p, err := m.GetCryptProcess(c.Crypt)
// 		if err == nil {
// 			m.cryptKeys[c.Name] = p
// 		}
// 	}

// 	if len(c.Before) > 0 {
// 		ps, err := m.GetBeforeProcess(c.Before)
// 		if err == nil {
// 			m.beforeProcess[c.Name] = ps
// 		}
// 	}

// 	if len(c.After) > 0 {
// 		ps, err := m.GetAfterProcess(c.Before)
// 		if err == nil {
// 			m.afterProcess[c.Name] = ps
// 		}
// 	}
// }
