package model

import (
	"errors"

	"github.com/zlsgo/app_module/database/hashid"
	"github.com/zlsgo/app_module/model/schema"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zstring"
)

// perfect is perfect
func perfect(alias string, m *Schema) (err error) {
	m.alias = alias

	salt := ""
	salt = m.define.Options.Salt
	cryptLen := m.define.Options.CryptLen
	if cryptLen <= 0 {
		cryptLen = 12
	}
	m.Hashid = hashid.New(salt, cryptLen)

	m.readOnlyKeys = make([]string, 0)
	m.relationsKeys = make([]string, 0)
	m.cryptKeys = make(map[string]CryptProcess, 0)
	m.afterProcess = make(map[string][]afterProcess, 0)
	m.beforeProcess = make(map[string][]beforeProcess, 0)

	isNotFields := len(m.define.Fields) == 0
	m.fields, err = perfectField(m)
	if err != nil {
		return
	}

	if !isNotFields {
		m.inlayFields = []string{idKey}
		if m.define.Options.Timestamps {
			if zarray.Contains(m.fields, CreatedAtKey) {
				err = errors.New(CreatedAtKey + " is a reserved field")
				return
			}
			if zarray.Contains(m.fields, UpdatedAtKey) {
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

		if m.define.Options.SoftDeletes {
			if zarray.Contains(m.fields, DeletedAtKey) {
				err = errors.New(DeletedAtKey + " is a reserved field")
				return
			}
			m.inlayFields = append(m.inlayFields, DeletedAtKey)
		}

		m.fullFields = append([]string{idKey}, m.fields...)
		m.fullFields = zarray.Unique(append(m.fullFields, m.inlayFields...))

		if m.define.Options.SoftDeletes {
			flen := len(m.fullFields)
			for i := 0; i < flen; i++ {
				f := m.fullFields[i]
				if f == DeletedAtKey {
					m.fullFields = append(m.fullFields[0:i], m.fullFields[i+1:]...)
					break
				}
			}
		}

		m.lowFields = m.define.Options.LowFields
	} else {
		m.define.Options.DisabledMigrator = true
	}

	if len(m.define.Relations) > 0 {
		for k := range m.define.Relations {
			v := m.define.Relations[k]
			if len(v.ForeignKey) != len(v.SchemaKey) {
				return errors.New("ForeignKey and SchemaKey must be the same length")
			}

		}

		newRelations := make(map[string]schema.Relation, len(m.define.Relations))
		for k := range m.define.Relations {
			v := m.define.Relations[k]
			newRelations[zstring.CamelCaseToSnakeCase(k)] = v
		}
		m.define.Relations = newRelations
		m.relationsKeys = zarray.Keys(m.define.Relations)
	} else {
		m.define.Relations = make(map[string]schema.Relation)
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
