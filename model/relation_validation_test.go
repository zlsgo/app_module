package model

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

type noopStorage struct{}

type noopMigration struct{}

func (noopStorage) GetStorageType() StorageType { return NoSQLStorage }

func (noopStorage) GetOptions() ztype.Map { return nil }

func (noopStorage) Transaction(run func(s Storageer) error) error { return run(noopStorage{}) }

func (noopStorage) Find(table string, filter ztype.Map, fn ...func(*CondOptions)) (ztype.Maps, error) {
	return ztype.Maps{}, nil
}

func (noopStorage) First(table string, filter ztype.Map, fn ...func(*CondOptions)) (ztype.Map, error) {
	return ztype.Map{}, nil
}

func (noopStorage) Pages(table string, page, pagesize int, filter ztype.Map, fn ...func(*CondOptions)) (ztype.Maps, PageInfo, error) {
	return ztype.Maps{}, PageInfo{}, nil
}

func (noopStorage) Migration(model *Schema) Migrationer { return noopMigration{} }

func (noopStorage) Insert(table string, data ztype.Map, fn ...func(*InsertOptions)) (interface{}, error) {
	return nil, nil
}

func (noopStorage) InsertMany(table string, data ztype.Maps, fn ...func(*InsertOptions)) ([]interface{}, error) {
	return []interface{}{}, nil
}

func (noopStorage) Delete(table string, filter ztype.Map, fn ...func(*CondOptions)) (int64, error) {
	return 0, nil
}

func (noopStorage) Update(table string, data ztype.Map, filter ztype.Map, fn ...func(*CondOptions)) (int64, error) {
	return 0, nil
}

func (noopMigration) Auto(deleteColumn ...DealOldColumn) error { return nil }

func (noopMigration) HasTable() bool { return false }

func (noopMigration) GetFields() (ztype.Map, error) { return ztype.Map{}, nil }

func TestRelationNameNormalizeCollision(t *testing.T) {
	schemas := NewSchemas(nil, nil, SchemaOptions{})
	userSchema := schema.Schema{
		Name: "user",
		Fields: map[string]schema.Field{
			"name": {Type: "string"},
		},
		Relations: map[string]schema.Relation{
			"FooBar": {
				Type:       schema.RelationSingle,
				Schema:     "profile",
				ForeignKey: []string{"profile_id"},
				SchemaKey:  []string{"id"},
			},
			"foo_bar": {
				Type:       schema.RelationSingle,
				Schema:     "profile",
				ForeignKey: []string{"profile_id"},
				SchemaKey:  []string{"id"},
			},
		},
	}

	_, err := schemas.Reg("user", userSchema, false)
	if err == nil || !strings.Contains(fmt.Sprintf("%v", err), "normalization") {
		t.Fatalf("expected normalization conflict error, got %v", err)
	}
}

func TestRelationMissingForeignKey(t *testing.T) {
	schemas := NewSchemas(nil, nil, SchemaOptions{})
	userSchema := schema.Schema{
		Name: "user",
		Fields: map[string]schema.Field{
			"name": {Type: "string"},
		},
		Relations: map[string]schema.Relation{
			"profile": {
				Type:      schema.RelationSingle,
				Schema:    "profile",
				SchemaKey: []string{"id"},
			},
		},
	}

	_, err := schemas.Reg("user", userSchema, false)
	if err == nil || !strings.Contains(fmt.Sprintf("%v", err), "foreign_key") {
		t.Fatalf("expected foreign_key error, got %v", err)
	}
}

func TestRelationSchemaKeyDefault(t *testing.T) {
	schemas := NewSchemas(nil, nil, SchemaOptions{})
	userSchema := schema.Schema{
		Name: "user",
		Fields: map[string]schema.Field{
			"name": {Type: "string"},
		},
		Relations: map[string]schema.Relation{
			"profile": {
				Type:       schema.RelationSingle,
				Schema:     "profile",
				ForeignKey: []string{"profile_id"},
			},
		},
	}

	m, err := schemas.Reg("user", userSchema, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	rel, ok := m.GetDefine().Relations["profile"]
	if !ok {
		t.Fatalf("relation not found")
	}

	if len(rel.SchemaKey) != 1 || rel.SchemaKey[0] != idKey {
		t.Fatalf("schema_key default mismatch: %v", rel.SchemaKey)
	}
}

func TestRelationManyToManyRequiresSQLStorage(t *testing.T) {
	schemas := NewSchemas(nil, noopStorage{}, SchemaOptions{})
	userSchema := schema.Schema{
		Name: "user",
		Fields: map[string]schema.Field{
			"name": {Type: "string"},
		},
		Relations: map[string]schema.Relation{
			"roles": {
				Type:      schema.RelationManyToMany,
				Schema:    "role",
				PivotKeys: schema.PivotKeys{Foreign: []string{"user_id"}, Related: []string{"role_id"}},
			},
		},
	}

	_, err := schemas.Reg("user", userSchema, false)
	if err == nil || !strings.Contains(fmt.Sprintf("%v", err), "many_to_many") {
		t.Fatalf("expected many_to_many storage error, got %v", err)
	}
}
