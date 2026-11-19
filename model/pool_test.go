package model

import "testing"

// 调用方会把包级切片直接赋给 CondOptions 的字段（例如
// `o.OrderBy = appPrimaryOrder`、`o.Fields = appQueryFields`）。
// releaseCondOptions 若只把长度切成 0、保留底层数组，池化对象下一次
// 被复用时的 append 就会写穿到那些包级切片上，导致跨查询的字段/排序串写
// （典型症状：对没有该列的表生成 ORDER BY，数据库报 no such column）。
//
// 注意：对象一旦归还给 sync.Pool，测试就不能再往它上面写数据，否则会给
// 池里留下带排序列/字段的脏对象，同样造成上面那种跨查询串写。这里统一
// 通过 acquireCondOptions 取对象，用完立刻 releaseCondOptions 归还。
func TestReleaseCondOptionsDetachesCallerBackingArray(t *testing.T) {
	t.Run("OrderBy", func(t *testing.T) {
		shared := []OrderByItem{{Field: "app_id", Direction: "ASC"}}

		opts := acquireCondOptions()
		opts.OrderBy = shared
		releaseCondOptions(opts)

		// 归还的对象必须与调用方切片断开，而不是留着一个容量仍在的切片头。
		if c := cap(opts.OrderBy); c != 0 {
			t.Fatalf("归还后仍持有底层数组: len=%d cap=%d", len(opts.OrderBy), c)
		}

		// 模拟池对象被下一次查询复用并 append。
		reused := acquireCondOptions()
		reused.OrderBy = append(reused.OrderBy, OrderByItem{Field: "update_sequence", Direction: "DESC"})
		if shared[0].Field != "app_id" {
			t.Fatalf("池化 append 污染了调用方切片: shared[0]=%q", shared[0].Field)
		}
		releaseCondOptions(reused)
	})

	t.Run("Fields", func(t *testing.T) {
		shared := []string{"app_id"}

		opts := acquireCondOptions()
		opts.Fields = shared
		releaseCondOptions(opts)

		if c := cap(opts.Fields); c != 0 {
			t.Fatalf("归还后仍持有底层数组: len=%d cap=%d", len(opts.Fields), c)
		}

		reused := acquireCondOptions()
		reused.Fields = append(reused.Fields, "update_sequence")
		if shared[0] != "app_id" {
			t.Fatalf("池化 append 污染了调用方切片: shared[0]=%q", shared[0])
		}
		releaseCondOptions(reused)
	})
}
