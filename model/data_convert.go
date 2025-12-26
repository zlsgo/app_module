package model

import (
	"reflect"

	"github.com/sohaha/zlsgo/ztype"
)

func dataToMap(value any) (ztype.Map, error) {
	if value == nil {
		return nil, ErrInvalidData
	}

	switch v := value.(type) {
	case ztype.Map:
		if v == nil {
			return ztype.Map{}, nil
		}
		return v, nil
	case map[string]interface{}:
		if v == nil {
			return ztype.Map{}, nil
		}
		return ztype.Map(v), nil
	case Filter:
		if v == nil {
			return ztype.Map{}, nil
		}
		return ztype.Map(v), nil
	}

	rv := reflect.ValueOf(value)
	if !rv.IsValid() {
		return nil, ErrInvalidData
	}
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil, ErrInvalidData
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Map && rv.Kind() != reflect.Struct {
		return nil, ErrInvalidData
	}

	m := ztype.ToMap(rv.Interface())
	if m == nil {
		m = ztype.Map{}
	}
	return m, nil
}

func dataToMaps(value any) (ztype.Maps, error) {
	if value == nil {
		return ztype.Maps{}, nil
	}

	switch v := value.(type) {
	case ztype.Maps:
		if v == nil {
			return ztype.Maps{}, nil
		}
		return v, nil
	case []ztype.Map:
		if v == nil {
			return ztype.Maps{}, nil
		}
		return ztype.Maps(v), nil
	case []map[string]interface{}:
		if v == nil {
			return ztype.Maps{}, nil
		}
		result := make(ztype.Maps, len(v))
		for i := range v {
			result[i] = ztype.Map(v[i])
		}
		return result, nil
	}

	rv := reflect.ValueOf(value)
	if !rv.IsValid() {
		return nil, ErrInvalidData
	}
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return ztype.Maps{}, nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Slice && rv.IsNil() {
		return ztype.Maps{}, nil
	}
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return nil, ErrInvalidData
	}

	l := rv.Len()
	result := make(ztype.Maps, 0, l)
	for i := 0; i < l; i++ {
		item := rv.Index(i).Interface()
		itemMap, err := dataToMap(item)
		if err != nil {
			return nil, err
		}
		if !itemMap.IsEmpty() {
			result = append(result, itemMap)
		}
	}

	return result, nil
}
