package database

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestKeyOfPreservesIntegerIdentity(t *testing.T) {
	for _, test := range []struct {
		value int64
		want  Key
	}{
		{0, "0"},
		{1001, "1001"},
		{-1, "-1"},
		{-9223372036854775808, "-9223372036854775808"},
		{9223372036854775807, "9223372036854775807"},
	} {
		if got := KeyOf(test.value); got != test.want {
			t.Errorf("KeyOf(%d) = %q, want %q", test.value, got, test.want)
		}
		got, err := test.want.Int64()
		if err != nil || got != test.value {
			t.Errorf("Key(%q).Int64() = %d, %v, want %d", test.want, got, err, test.value)
		}
	}
	if _, err := Key("").Int64(); !errors.Is(err, ErrInvalidKey) {
		t.Fatalf("empty key Int64() error = %v", err)
	}
	if _, err := Key("tenant").Int64(); err == nil {
		t.Fatal("a non-decimal key must not parse as int64")
	}
}

func TestPartitionsAcceptsOpaqueStringKeys(t *testing.T) {
	partitions, dir, base := newTestPartitions(t, 0)

	first, err := partitions.Get(Key("tenant_West-2"))
	if err != nil {
		t.Fatalf("open string partition: %v", err)
	}
	second, err := partitions.Get(Key("tenant_West-2"))
	if err != nil {
		t.Fatalf("reopen string partition: %v", err)
	}
	if first.DB == base.DB || first.DB != second.DB || first.Prefix != base.Prefix {
		t.Fatal("a string key must have a stable, isolated connection with the original table prefix")
	}
	if _, err := os.Stat(filepath.Join(dir, "db_tenant_West-2.db")); err != nil {
		t.Fatalf("string partition file: %v", err)
	}
}

func TestPartitionsSQLiteKeepsDataSeparate(t *testing.T) {
	partitions, _, _ := newTestPartitions(t, 0)
	keys := []Key{"", "tenant_West-2", "tenant_East-9", KeyOf(-1)}
	for _, key := range keys {
		storage, err := partitions.Get(key)
		if err != nil {
			t.Fatalf("Get(%q): %v", key, err)
		}
		if _, err := storage.DB.Exec("CREATE TABLE IF NOT EXISTS data_records (marker TEXT NOT NULL)"); err != nil {
			t.Fatalf("create records for %q: %v", key, err)
		}
		if _, err := storage.DB.Exec("INSERT INTO data_records (marker) VALUES (?)", string(key)); err != nil {
			t.Fatalf("write records for %q: %v", key, err)
		}
	}
	for _, key := range keys {
		storage, err := partitions.Get(key)
		if err != nil {
			t.Fatalf("reopen %q: %v", key, err)
		}
		rows, err := storage.DB.QueryToMaps("SELECT marker FROM data_records")
		if err != nil {
			t.Fatalf("read records for %q: %v", key, err)
		}
		if len(rows) != 1 || rows[0]["marker"] != string(key) {
			t.Fatalf("partition %q leaked records: %#v", key, rows)
		}
	}
}
