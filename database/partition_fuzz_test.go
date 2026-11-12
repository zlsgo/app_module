package database

import (
	"path/filepath"
	"strings"
	"testing"
)

// FuzzKeyRoundTripAndPathSafety 校验两条分区不变量：
//   - KeyOf 与 Key.Int64 对任意 int64 互逆；
//   - 合法键必须能安全地作为文件名片段，派生路径不会逃出默认库所在目录。
func FuzzKeyRoundTripAndPathSafety(f *testing.F) {
	f.Add("", int64(0))
	f.Add("tenant-west", int64(-1))
	f.Add("../etc", int64(1001))
	f.Add("a/b", int64(1<<62))
	f.Add("a\\b", int64(-1<<63))

	f.Fuzz(func(t *testing.T, raw string, value int64) {
		key := KeyOf(value)
		got, err := key.Int64()
		if err != nil || got != value {
			t.Fatalf("KeyOf(%d).Int64() = %d, %v", value, got, err)
		}

		if !validPartitionKey(Key(raw)) || raw == "" {
			return
		}
		if strings.ContainsAny(raw, `/\`) || filepath.Base(raw) != raw {
			t.Fatalf("valid key is not a safe path component: %q", raw)
		}

		path, err := partitionSQLitePath(filepath.Join("data", "db.db"), Key(raw))
		if err != nil {
			return
		}
		if filepath.Dir(path) != "data" {
			t.Fatalf("partition path escaped the database directory: %q -> %q", raw, path)
		}
	})
}
