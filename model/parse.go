package model

import (
	"errors"
	"strings"

	"github.com/zlsgo/app_module/database/hashid"
	"github.com/zlsgo/app_module/model/schema"

	"github.com/sohaha/zlsgo/zarray"
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
		s.refreshFieldsSet()
	} else {
		b := true
		s.define.Options.DisabledMigrator = &b
	}

	if len(s.define.Relations) > 0 {
		storageType := StorageType(0)
		if s.Storage != nil {
			storageType = s.Storage.GetStorageType()
		}

		newRelations := make(map[string]schema.Relation, len(s.define.Relations))
		for k := range s.define.Relations {
			v := s.define.Relations[k]
			if v.Type == schema.RelationManyToMany {
				if s.Storage != nil && storageType != SQLStorage {
					return errors.New("many_to_many relation requires sql storage")
				}
				parentKeys := v.ForeignKey
				if len(parentKeys) == 0 {
					parentKeys = []string{idKey}
				}
				relatedKeys := v.SchemaKey
				if len(relatedKeys) == 0 {
					relatedKeys = []string{idKey}
				}
				if len(v.PivotKeys.Foreign) != len(parentKeys) {
					return errors.New("Pivot foreign key length mismatch")
				}
				if len(v.PivotKeys.Related) != len(relatedKeys) {
					return errors.New("Pivot related key length mismatch")
				}
			} else {
				if len(v.ForeignKey) == 0 {
					return errors.New("relation foreign_key required")
				}
				if len(v.SchemaKey) == 0 {
					if len(v.ForeignKey) == 1 {
						v.SchemaKey = []string{idKey}
					} else {
						return errors.New("relation schema_key required for composite foreign_key")
					}
				}
				if len(v.ForeignKey) != len(v.SchemaKey) {
					return errors.New("ForeignKey and SchemaKey must be the same length")
				}
			}

			normalized := camelToSnake(k)
			if _, exists := newRelations[normalized]; exists {
				return errors.New("relation " + normalized + " already exists after normalization")
			}
			newRelations[normalized] = v
		}
		s.define.Relations = newRelations
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

func camelToSnake(name string) string {
	if name == "" {
		return ""
	}
	var b strings.Builder
	b.Grow(len(name) + len(name)/2)
	for i := 0; i < len(name); i++ {
		c := name[i]
		if c == '_' {
			b.WriteByte(c)
			continue
		}
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				prev := name[i-1]
				nextLower := i+1 < len(name) && name[i+1] >= 'a' && name[i+1] <= 'z'
				prevLower := prev >= 'a' && prev <= 'z'
				prevDigit := prev >= '0' && prev <= '9'
				prevUpper := prev >= 'A' && prev <= 'Z'
				if prevLower || prevDigit || (prevUpper && nextLower) {
					b.WriteByte('_')
				}
			}
			b.WriteByte(c + 'a' - 'A')
			continue
		}
		b.WriteByte(c)
	}
	return b.String()
}
