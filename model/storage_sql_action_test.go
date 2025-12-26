package model

import (
	"testing"

	"github.com/sohaha/zlsgo/ztype"
)

func TestStorageFindUsesCondFields(t *testing.T) {
	_, m := newTestDB(t, "storage_find_fields")

	_, err := m.Model().Insert(ztype.Map{
		"name":   "Alice",
		"email":  "alice@example.com",
		"age":    30,
		"status": 1,
	})
	if err != nil {
		t.Fatalf("insert failed: %v", err)
	}

	rows, err := m.Storage.Find(m.GetTableName(), ztype.Map{"name": "Alice"}, func(co *CondOptions) {
		co.Fields = append(co.Fields[:0], "name")
	})
	if err != nil {
		t.Fatalf("find failed: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("unexpected rows: %d", len(rows))
	}
	if _, ok := rows[0]["email"]; ok {
		t.Fatalf("unexpected field: email")
	}
	if rows[0].Get("name").String() != "Alice" {
		t.Fatalf("unexpected name: %s", rows[0].Get("name").String())
	}
}

func TestStoragePagesUsesCondFields(t *testing.T) {
	_, m := newTestDB(t, "storage_pages_fields")

	_, err := m.Model().Insert(ztype.Map{
		"name":   "Bob",
		"email":  "bob@example.com",
		"age":    22,
		"status": 1,
	})
	if err != nil {
		t.Fatalf("insert failed: %v", err)
	}

	rows, _, err := m.Storage.Pages(m.GetTableName(), 1, 10, ztype.Map{}, func(co *CondOptions) {
		co.Fields = append(co.Fields[:0], "name")
	})
	if err != nil {
		t.Fatalf("pages failed: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("unexpected rows: %d", len(rows))
	}
	if _, ok := rows[0]["email"]; ok {
		t.Fatalf("unexpected field: email")
	}
	if rows[0].Get("name").String() != "Bob" {
		t.Fatalf("unexpected name: %s", rows[0].Get("name").String())
	}
}
