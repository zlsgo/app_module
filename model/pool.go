package model

import (
	"sync"
)

var condOptionsPool = sync.Pool{
	New: func() any {
		// releaseCondOptions 会把切片置空，所以这里不需要预分配容量；
		// 首次 append 由调用方自然分配，且不会与任何共享数组别名。
		return &CondOptions{}
	},
}

func acquireCondOptions() *CondOptions {
	return condOptionsPool.Get().(*CondOptions)
}

func releaseCondOptions(opts *CondOptions) {
	// 这里必须把切片置空而不是 opts.Fields[:0]：调用方会把包级切片直接
	// 赋给这些字段（如 o.OrderBy = appPrimaryOrder、o.Fields = appQueryFields），
	// 两者共享底层数组。若只把长度切成 0，池中对象下一次被复用时的 append
	// 会写穿到调用方的包级切片，造成跨查询的字段/排序串写。
	opts.Fields = nil
	opts.Relations = nil
	opts.OrderBy = nil
	opts.GroupBy = nil
	opts.Join = nil
	opts.Limit = 0
	opts.Offset = 0
	condOptionsPool.Put(opts)
}
