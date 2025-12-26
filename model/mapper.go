package model

import (
	"github.com/sohaha/zlsgo/ztype"
)

// Mapper 数据映射接口
type Mapper[T any] interface {
	MapOne(row ztype.Map) (T, error)
	MapMany(rows ztype.Maps) ([]T, error)
}

// MapMapper Map 映射器
type MapMapper struct{}

// MapOne 映射单条记录为 Map
func (m MapMapper) MapOne(row ztype.Map) (ztype.Map, error) {
	return row, nil
}

// MapMany 映射多条记录为 Map 数组
func (m MapMapper) MapMany(rows ztype.Maps) ([]ztype.Map, error) {
	return rows, nil
}

// StructMapper 结构体映射器
type StructMapper[T any] struct{}

// MapOne 映射单条记录为结构体
func (m StructMapper[T]) MapOne(row ztype.Map) (T, error) {
	var result T
	err := ztype.To(row, &result)
	return result, err
}

// MapMany 映射多条记录为结构体数组
func (m StructMapper[T]) MapMany(rows ztype.Maps) ([]T, error) {
	var result []T
	err := ztype.To(rows, &result)
	return result, err
}
