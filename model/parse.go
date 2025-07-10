package model

import (
	"errors"

	"github.com/zlsgo/app_module/database/hashid"
	"github.com/zlsgo/app_module/model/schema"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zstring"
)

// perfect is perfect
func perfect(alias string, s *Schema, o *SchemaOptions) (err error) {
	s.alias = alias

	cryptLen := s.define.Options.CryptLen
	if cryptLen <= 0 {
		cryptLen = 12
	}
	s.Hashid = hashid.New(s.define.Options.Salt, cryptLen)

	s.readOnlyKeys = make([]string, 0, 4)
	s.relationsKeys = make([]string, 0, len(s.define.Relations))
	s.cryptKeys = make(map[string]CryptProcess, 2)
	s.afterProcess = make(map[string][]afterProcess, 4)
	s.beforeProcess = make(map[string][]beforeProcess, 4)

	isNotFields := len(s.define.Fields) == 0
	s.fields, err = perfectField(s)
	if err != nil {
		return
	}

	_ = perfectOptions(s, o)

	if !isNotFields {
		s.inlayFields = []string{idKey}

		if *s.define.Options.Timestamps {
			if zarray.Contains(s.fields, CreatedAtKey) {
				return errors.New(CreatedAtKey + " is a reserved field")
			}
			if zarray.Contains(s.fields, UpdatedAtKey) {
				return errors.New(UpdatedAtKey + " is a reserved field")
			}

			after, err := s.GetAfterProcess([]string{"date|Y-m-d H:i:s"})
			if err != nil {
				return err
			}
			s.afterProcess[CreatedAtKey] = after
			s.afterProcess[UpdatedAtKey] = after
			s.inlayFields = append(s.inlayFields, CreatedAtKey, UpdatedAtKey)
		}

		// if m.models.Options.CreatedBy {
		// 	if zarray.Contains(m.Fields, CreatedByKey) {
		// 		err = errors.New(CreatedByKey + " is a reserved field")
		// 		return
		// 	}
		// 	m.inlayFields = append(m.inlayFields, CreatedByKey)
		// }

		if *s.define.Options.SoftDeletes {
			if zarray.Contains(s.fields, DeletedAtKey) {
				return errors.New(DeletedAtKey + " is a reserved field")
			}
			s.inlayFields = append(s.inlayFields, DeletedAtKey)
		}

		capacity := 1 + len(s.fields) + len(s.inlayFields)
		s.fullFields = make([]string, 0, capacity)
		s.fullFields = append(s.fullFields, idKey)
		s.fullFields = append(s.fullFields, s.fields...)
		s.fullFields = zarray.Unique(append(s.fullFields, s.inlayFields...))

		if *s.define.Options.SoftDeletes {
			for i := range s.fullFields {
				if s.fullFields[i] == DeletedAtKey {
					s.fullFields = append(s.fullFields[0:i], s.fullFields[i+1:]...)
					break
				}
			}
		}

		s.lowFields = s.define.Options.LowFields
	} else {
		b := true
		s.define.Options.DisabledMigrator = &b
	}

	if len(s.define.Relations) > 0 {
		for k := range s.define.Relations {
			v := s.define.Relations[k]
			if len(v.ForeignKey) != len(v.SchemaKey) {
				return errors.New("ForeignKey and SchemaKey must be the same length")
			}

		}

		newRelations := make(map[string]schema.Relation, len(s.define.Relations))
		for k := range s.define.Relations {
			v := s.define.Relations[k]
			newRelations[zstring.CamelCaseToSnakeCase(k)] = v
		}
		s.define.Relations = newRelations
		s.relationsKeys = zarray.Keys(s.define.Relations)
	} else {
		s.define.Relations = make(map[string]schema.Relation)
	}

	// if m.models.Options.CreatedBy {
	// 	c := &ModelRelation{
	// 		Key:     CreatedByKey,
	// 		Model:   define.UserModel,
	// 		Foreign: "_id",
	// 		Fields: []string{
	// 			"account",
	// 			"nickname",
	// 		},
	// 	}
	// 	m.models.Relations[zstring.SnakeCaseToCamelCase(CreatedByKey, true)] = c
	// }

	s.views = parseViews(s)
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
