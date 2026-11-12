package html

import (
	"testing"

	"github.com/sohaha/zlsgo/ztype"
)

// TestOptionsConfigKeys 保证结构体字段名与配置文件里的键完全一致。
//
// 缺 z 标签时 ztype 生成的键是 staticprefix / errorpage，user 配置里的
// static_prefix / error_page 在归一化比较（忽略大小写和 _ -）下同样命中字段，
// 两份键同时存在时由 map 迭代顺序决定谁生效，配置文件会随机失效。
func TestOptionsConfigKeys(t *testing.T) {
	m := ztype.ToMap(Options{})
	for _, key := range []string{"static_prefix", "error_page"} {
		if !m.Has(key) {
			t.Fatalf("默认配置缺少键 %q，实际为 %v", key, m)
		}
	}
}
