package model

import "testing"

// 调用方会把包级切片直接赋给 CondOptions 的字段（例如
// `o.OrderBy = appPrimaryOrder`、`o.Fields = appQueryFields`）。
// releaseCondOptions 若只把长度切成 0、保留底层数组，池化对象下一次
// 被复用时的 append 就会写穿到那些包级切片上，导致跨查询的字段/排序串写
// （典型症状：对没有该列的表生成 ORDER BY，数据库报 no such column）。
func TestReleaseCondOptionsDetachesCallerBackingArray(t *testing.T) {
	t.Run("OrderBy", func(t *testing.T) {
		shared := []OrderByItem{{Field: "app_id", Direction: "ASC"}}
		opts := &CondOptions{OrderBy: shared}

		releaseCondOptions(opts)

		// 模拟另一个调用方复用同一个池对象并 append。
		opts.OrderBy = append(opts.OrderBy, OrderByItem{Field: "update_sequence", Direction: "DESC"})
		if shared[0].Field != "app_id" {
			t.Fatalf("池化 append 污染了调用方切片: shared[0]=%q", shared[0].Field)
		}
	})

	t.Run("Fields", func(t *testing.T) {
		shared := []string{"app_id"}
		opts := &CondOptions{Fields: shared}

		releaseCondOptions(opts)

		opts.Fields = append(opts.Fields, "update_sequence")
		if shared[0] != "app_id" {
			t.Fatalf("池化 append 污染了调用方切片: shared[0]=%q", shared[0])
		}
	})
}
