package model

import (
	"strings"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
	"golang.org/x/exp/constraints"
)

// filter 定义可接受的过滤器类型
type filter interface {
	ztype.Map | constraints.Integer | string | Filter
}

// getFilter 将各种类型的过滤器转换为统一的 ztype.Map 格式
// 并自动处理软删除字段过滤
func getFilter[T filter | any](m *Schema, filter T) (filterMap ztype.Map) {
	f := (interface{})(filter)

	filterData, ok := f.(Filter)
	if ok {
		filterMap = ztype.Map(filterData)
	} else {
		filterMap, ok = f.(ztype.Map)
	}

	if !ok {
		mapData := ztype.ToMap(f)
		if len(mapData) > 0 {
			if len(mapData) == 1 {
				if mapData.Get("0").Value() != f {
					filterMap = mapData
					ok = true
				}
			} else {
				filterMap = mapData
				ok = true
			}
		}
	}

	if !ok {
		idVal := f
		filterMap = ztype.Map{
			idKey: idVal,
		}
	} else if filterMap == nil {
		filterMap = ztype.Map{}
	}

	// 过滤无效字段：排除不在模型定义中的字段
	for key := range filterMap {
		k := zstring.TrimSpace(key)
		if k == "" || strings.Contains(k, placeHolder) {
			continue
		}
		// 删除包含点号或空格的非法 key（防止 SQL 注入）
		if strings.Contains(k, ".") || strings.Contains(k, " ") {
			delete(filterMap, key)
			continue
		}
		if !zarray.Contains(m.GetFields(), k) {
			delete(filterMap, key)
		}
	}

	// 自动添加软删除过滤条件
	if *m.define.Options.SoftDeletes {
		if InsideOption.softDeleteIsTime {
			filterMap[DeletedAtKey] = nil
		} else {
			filterMap[DeletedAtKey] = 0
		}
	}

	return
}
