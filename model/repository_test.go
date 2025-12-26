package model

import (
	"testing"

	"github.com/sohaha/zlsgo/ztype"
)

func TestMapMapper(t *testing.T) {
	mapper := MapMapper{}

	row := ztype.Map{"id": 1, "name": "test"}
	result, err := mapper.MapOne(row)
	if err != nil {
		t.Errorf("MapOne error: %v", err)
	}
	if result["id"] != 1 || result["name"] != "test" {
		t.Errorf("MapOne result mismatch: %v", result)
	}

	rows := ztype.Maps{
		{"id": 1, "name": "test1"},
		{"id": 2, "name": "test2"},
	}
	results, err := mapper.MapMany(rows)
	if err != nil {
		t.Errorf("MapMany error: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("MapMany length mismatch: %d", len(results))
	}
}

func TestStructMapper(t *testing.T) {
	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	mapper := StructMapper[User]{}

	row := ztype.Map{"id": 1, "name": "test"}
	result, err := mapper.MapOne(row)
	if err != nil {
		t.Errorf("MapOne error: %v", err)
	}
	if result.ID != 1 || result.Name != "test" {
		t.Errorf("MapOne result mismatch: %+v", result)
	}

	rows := ztype.Maps{
		{"id": 1, "name": "test1"},
		{"id": 2, "name": "test2"},
	}
	results, err := mapper.MapMany(rows)
	if err != nil {
		t.Errorf("MapMany error: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("MapMany length mismatch: %d", len(results))
	}
	if results[0].ID != 1 || results[1].Name != "test2" {
		t.Errorf("MapMany result mismatch: %+v", results)
	}
}

func TestQueryFilter(t *testing.T) {
	f := Eq("status", "active")
	m := f.ToMap()
	if m["status"] != "active" {
		t.Errorf("Eq filter mismatch: %v", m)
	}

	f = Gt("age", 18)
	m = f.ToMap()
	if m["age >"] != 18 {
		t.Errorf("Gt filter mismatch: %v", m)
	}

	f = In("role", []string{"admin", "user"})
	m = f.ToMap()
	if m["role IN"] == nil {
		t.Errorf("In filter mismatch: %v", m)
	}

	f = And(Eq("status", "active"), Gt("age", 18))
	m = f.ToMap()
	if m["status"] != "active" || m["age >"] != 18 {
		t.Errorf("And filter mismatch: %v", m)
	}

	f = ID(123)
	m = f.ToMap()
	if m[idKey] != 123 {
		t.Errorf("ID filter mismatch: %v", m)
	}
}

func TestQueryFilterOr(t *testing.T) {
	f := Or(Eq("status", "active"), Eq("role", "admin"))
	m := f.ToMap()
	orFilters, ok := m[placeHolderOR].([]ztype.Map)
	if !ok {
		t.Errorf("Or filter type mismatch: %v", m)
	}
	if len(orFilters) != 2 {
		t.Errorf("Or filter length mismatch: %d", len(orFilters))
	}
}

func TestQueryFilterComplex(t *testing.T) {
	f := And(
		Eq("status", "active"),
		Or(
			Eq("role", "admin"),
			Eq("role", "superuser"),
		),
		Ge("created_at", "2024-01-01"),
	)
	m := f.ToMap()
	if m["status"] != "active" {
		t.Errorf("Complex filter status mismatch: %v", m)
	}
	if m["created_at >="] != "2024-01-01" {
		t.Errorf("Complex filter created_at mismatch: %v", m)
	}
}

func TestMapFilter(t *testing.T) {
	f := Q(ztype.Map{"status": "active", "age >": 18})
	m := f.ToMap()
	if m["status"] != "active" || m["age >"] != 18 {
		t.Errorf("MapFilter mismatch: %v", m)
	}
}

func TestStructFilter(t *testing.T) {
	type StatusFilter struct {
		Status int `json:"status"`
	}

	f := Q(StatusFilter{Status: 1})
	m := f.ToMap()
	if m["status"] != 1 {
		t.Errorf("Struct filter mismatch: %v", m)
	}
}

func TestBetweenFilter(t *testing.T) {
	f := Between("age", 18, 65)
	m := f.ToMap()
	vals, ok := m["age BETWEEN"].([]any)
	if !ok || len(vals) != 2 {
		t.Errorf("Between filter mismatch: %v", m)
	}
	if vals[0] != 18 || vals[1] != 65 {
		t.Errorf("Between filter values mismatch: %v", vals)
	}
}

func TestNullFilters(t *testing.T) {
	f := IsNull("deleted_at")
	m := f.ToMap()
	if _, ok := m["deleted_at IS NULL"]; !ok {
		t.Errorf("IsNull filter mismatch: %v", m)
	}

	f = IsNotNull("email")
	m = f.ToMap()
	if _, ok := m["email IS NOT NULL"]; !ok {
		t.Errorf("IsNotNull filter mismatch: %v", m)
	}
}
