package model

import (
	"sync"
)

const (
	defaultFieldsCap  = 8
	defaultOrderByCap = 4
	defaultGroupByCap = 4
)

var condOptionsPool = sync.Pool{
	New: func() any {
		return &CondOptions{
			Fields:    make([]string, 0, defaultFieldsCap),
			Relations: make([]string, 0, defaultFieldsCap),
			OrderBy:   make([]OrderByItem, 0, defaultOrderByCap),
			GroupBy:   make([]string, 0, defaultGroupByCap),
		}
	},
}

func acquireCondOptions() *CondOptions {
	return condOptionsPool.Get().(*CondOptions)
}

func releaseCondOptions(opts *CondOptions) {
	opts.Fields = opts.Fields[:0]
	opts.Relations = opts.Relations[:0]
	opts.OrderBy = opts.OrderBy[:0]
	opts.GroupBy = opts.GroupBy[:0]
	opts.Join = nil
	opts.Limit = 0
	opts.Offset = 0
	condOptionsPool.Put(opts)
}
