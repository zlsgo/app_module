package model

import (
	"errors"
	"testing"

	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/hook"
	"github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

func newHookTestDB(t *testing.T, tableName string, hookFn func(hook.Event, ...any) error) (*zdb.DB, *Schema) {
	db, err := zdb.New(&sqlite3.Config{
		File:   ":memory:",
		Memory: true,
	})
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}

	b := true
	testSchema := schema.Schema{
		Name: tableName,
		Table: schema.Table{
			Name: tableName,
		},
		Options: schema.Options{
			Timestamps: &b,
			Hook:       hookFn,
		},
		Fields: map[string]schema.Field{
			"name": {
				Type:  "string",
				Label: "Name",
			},
			"age": {
				Type:  "int",
				Label: "Age",
			},
			"status": {
				Type:  "int",
				Label: "Status",
			},
		},
	}

	schemas := NewSchemas(nil, NewSQL(db, ""), SchemaOptions{})
	s, err := schemas.Reg(tableName, testSchema, false)
	if err != nil {
		t.Fatalf("failed to register schema: %v", err)
	}

	if err := s.Migration().Auto(); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	return db, s
}

// TestHookInsert 测试 Insert Hook
func TestHookInsert(t *testing.T) {
	var (
		beforeInsertCalled = false
		afterInsertCalled  = false
		insertedData       ztype.Map
		insertedID         interface{}
	)

	hookFn := func(e hook.Event, data ...any) error {
		switch e {
		case hook.EventBeforeInsert:
			beforeInsertCalled = true
			if len(data) > 0 {
				if m, ok := data[0].(ztype.Map); ok {
					insertedData = m
				}
			}
		case hook.EventAfterInsert:
			afterInsertCalled = true
			if len(data) > 0 {
				insertedID = data[0]
			}
		}
		return nil
	}

	_, s := newHookTestDB(t, "hook_insert", hookFn)
	m := s.Model()

	// 插入数据
	id, err := m.Insert(ztype.Map{
		"name":   "Alice",
		"age":    25,
		"status": 1,
	})

	if err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	if !beforeInsertCalled {
		t.Error("BeforeInsert hook not called")
	}

	if !afterInsertCalled {
		t.Error("AfterInsert hook not called")
	}

	if insertedData == nil {
		t.Error("insertedData should not be nil")
	} else {
		if insertedData.Get("name").String() != "Alice" {
			t.Errorf("Expected name Alice, got %v", insertedData.Get("name"))
		}
	}

	if insertedID == nil {
		t.Error("insertedID should not be nil")
	}

	if id == nil {
		t.Error("Returned ID should not be nil")
	}

	t.Logf("Insert successful, ID: %v", id)
}

// TestHookInsertReject 测试 BeforeInsert Hook 拒绝插入
func TestHookInsertReject(t *testing.T) {
	hookFn := func(e hook.Event, data ...any) error {
		if e == hook.EventBeforeInsert {
			return errors.New("insert rejected by hook")
		}
		return nil
	}

	_, s := newHookTestDB(t, "hook_insert_reject", hookFn)
	m := s.Model()

	// 尝试插入数据
	_, err := m.Insert(ztype.Map{"name": "Bob", "age": 30, "status": 1})

	if err == nil {
		t.Error("Expected error, got nil")
	}

	if err.Error() != "insert rejected by hook" {
		t.Errorf("Expected 'insert rejected by hook', got '%v'", err)
	}

	t.Logf("Insert correctly rejected: %v", err)
}

// TestHookUpdate 测试 Update Hook
func TestHookUpdate(t *testing.T) {
	var (
		beforeUpdateCalled = false
		afterUpdateCalled  = false
		updateFilter       ztype.Map
		updateData         ztype.Map
		affectedRows       int64
	)

	hookFn := func(e hook.Event, data ...any) error {
		switch e {
		case hook.EventBeforeUpdate:
			beforeUpdateCalled = true
			if len(data) >= 2 {
				if f, ok := data[0].(ztype.Map); ok {
					updateFilter = f
				}
				if d, ok := data[1].(ztype.Map); ok {
					updateData = d
				}
			}
		case hook.EventAfterUpdate:
			afterUpdateCalled = true
			if len(data) >= 3 {
				if rows, ok := data[2].(int64); ok {
					affectedRows = rows
				}
			}
		}
		return nil
	}

	_, s := newHookTestDB(t, "hook_update", hookFn)
	m := s.Model()

	// 先插入数据
	_, _ = m.Insert(ztype.Map{"name": "Charlie", "age": 30, "status": 1})

	// 使用 name 字段更新数据（避免 ID 加密导致的 filter 不匹配问题）
	affected, err := m.Update(Filter{"name": "Charlie"}, ztype.Map{"status": 2})

	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	if !beforeUpdateCalled {
		t.Error("BeforeUpdate hook not called")
	}

	if !afterUpdateCalled {
		t.Error("AfterUpdate hook not called")
	}

	if updateFilter == nil {
		t.Error("updateFilter should not be nil")
	}

	if updateData == nil {
		t.Error("updateData should not be nil")
	} else {
		if updateData.Get("status").Int() != 2 {
			t.Errorf("Expected status 2, got %v", updateData.Get("status"))
		}
	}

	if affectedRows != 1 {
		t.Errorf("Hook should receive 1 affected row, got %d", affectedRows)
	}

	if affected != 1 {
		t.Errorf("Expected 1 affected row, got %d", affected)
	}

	t.Logf("Update successful, affected rows: %d", affected)
}

// TestHookDelete 测试 Delete Hook
func TestHookDelete(t *testing.T) {
	var (
		beforeDeleteCalled = false
		afterDeleteCalled  = false
		deleteFilter       ztype.Map
		deletedRows        int64
	)

	hookFn := func(e hook.Event, data ...any) error {
		switch e {
		case hook.EventBeforeDelete:
			beforeDeleteCalled = true
			if len(data) > 0 {
				if f, ok := data[0].(ztype.Map); ok {
					deleteFilter = f
				}
			}
		case hook.EventAfterDelete:
			afterDeleteCalled = true
			if len(data) >= 2 {
				if rows, ok := data[1].(int64); ok {
					deletedRows = rows
				}
			}
		}
		return nil
	}

	_, s := newHookTestDB(t, "hook_delete", hookFn)
	m := s.Model()

	// 先插入数据
	id, _ := m.Insert(ztype.Map{"name": "David", "age": 40, "status": 1})

	// 删除数据
	affected, err := m.Delete(Filter{idKey: id})

	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	if !beforeDeleteCalled {
		t.Error("BeforeDelete hook not called")
	}

	if !afterDeleteCalled {
		t.Error("AfterDelete hook not called")
	}

	if deleteFilter == nil {
		t.Error("deleteFilter should not be nil")
	}

	if deletedRows != 1 {
		t.Errorf("Expected 1 deleted row, got %d", deletedRows)
	}

	if affected != 1 {
		t.Errorf("Expected 1 affected row, got %d", affected)
	}

	t.Logf("Delete successful, affected rows: %d", affected)
}

// TestHookInsertMany 测试批量插入 Hook
func TestHookInsertMany(t *testing.T) {
	var (
		beforeInsertCalled = false
		afterInsertCalled  = false
		batchSize          int
	)

	hookFn := func(e hook.Event, data ...any) error {
		switch e {
		case hook.EventBeforeInsert:
			beforeInsertCalled = true
			if len(data) > 0 {
				if maps, ok := data[0].(ztype.Maps); ok {
					batchSize = len(maps)
				}
			}
		case hook.EventAfterInsert:
			afterInsertCalled = true
		}
		return nil
	}

	_, s := newHookTestDB(t, "hook_insert_many", hookFn)
	m := s.Model()

	// 批量插入
	ids, err := m.InsertMany(ztype.Maps{
		{"name": "User1", "age": 21, "status": 1},
		{"name": "User2", "age": 22, "status": 1},
		{"name": "User3", "age": 23, "status": 1},
	})

	if err != nil {
		t.Fatalf("InsertMany failed: %v", err)
	}

	if !beforeInsertCalled {
		t.Error("BeforeInsert hook not called")
	}

	if !afterInsertCalled {
		t.Error("AfterInsert hook not called")
	}

	if batchSize != 3 {
		t.Errorf("Expected batch size 3, got %d", batchSize)
	}

	idsSlice, ok := ids.([]interface{})
	if !ok {
		t.Fatalf("Expected IDs to be []interface{}, got %T", ids)
	}

	if len(idsSlice) != 3 {
		t.Errorf("Expected 3 IDs, got %d", len(idsSlice))
	}

	t.Logf("InsertMany successful, IDs: %v", idsSlice)
}
