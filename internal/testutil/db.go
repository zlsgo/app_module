// Package testutil 提供仅测试使用的辅助函数，不参与业务逻辑。
package testutil

import (
	"fmt"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

var dbSeq atomic.Uint64

// NewMemoryDB 创建一个独占的内存 SQLite 库。
//
// zdb 会把 File 拼成 "file:<名字>?cache=shared&mode=memory"，名字相同就是同一个
// 库。历史上测试统一用 File: ":memory:"，于是同一个测试进程里所有用例、以及
// `go test -count>1` 的每一轮都挤在同一个库上：数据会累加（计数类断言变成 2/4），
// 先后顺序不同的表结构也会互相影响。这里给每次调用一个唯一名字，让每个用例
// 拥有自己的库。
func NewMemoryDB(tb testing.TB) *zdb.DB {
	tb.Helper()

	db, err := zdb.New(&sqlite3.Config{
		File:       uniqueName(tb, "db"),
		Memory:     true,
		Parameters: "_pragma=busy_timeout(3000)",
	})
	if err != nil {
		tb.Fatalf("failed to create db: %v", err)
	}
	return db
}

// uniqueName 把测试名清洗成安全字符后拼上单调递增序号，保证进程内唯一。
func uniqueName(tb testing.TB, prefix string) string {
	tb.Helper()

	name := strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9':
			return r
		default:
			return '-'
		}
	}, tb.Name())

	return fmt.Sprintf("app-module-%s-%s-%d", prefix, name, dbSeq.Add(1))
}

// UniqueName 供测试注册全局名字（例如 znet 服务器名、hashid 注册名）时使用，
// 避免 `-count>1` 时重复注册直接失败或返回错误。
func UniqueName(tb testing.TB, prefix string) string {
	return uniqueName(tb, prefix)
}
