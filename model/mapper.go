package model

import (
	"github.com/sohaha/zlsgo/ztype"
)

type Mapper[T any] interface {
	MapOne(row ztype.Map) (T, error)
	MapMany(rows ztype.Maps) ([]T, error)
}

type MapMapper struct{}

func (m MapMapper) MapOne(row ztype.Map) (ztype.Map, error) {
	return row, nil
}

func (m MapMapper) MapMany(rows ztype.Maps) ([]ztype.Map, error) {
	return rows, nil
}

type StructMapper[T any] struct{}

func (m StructMapper[T]) MapOne(row ztype.Map) (T, error) {
	var result T
	err := ztype.To(row, &result)
	return result, err
}

func (m StructMapper[T]) MapMany(rows ztype.Maps) ([]T, error) {
	var result []T
	err := ztype.To(rows, &result)
	return result, err
}
