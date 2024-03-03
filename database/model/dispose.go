package model

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/ztype"
)

type beforeProcess func(interface{}) (string, error)

func (m *Model) GetBeforeProcess(p []string) (fn []beforeProcess, err error) {
	for _, v := range p {
		switch strings.ToLower(v) {
		default:
			v := strings.SplitN(v, "|", 2)
			if len(v) == 2 {
				switch v[0] {
				case "date":
					if v[1] != "-" {
						fn = append(fn, dateMarshalProcess(v[1]))
					}
					continue
				}
			}
			return nil, errors.New("before name not found")
		case "bool":
			fn = append(fn, boolMarshalProcess)
		case "json":
			fn = append(fn, jsonMarshalProcess(false))
		case "jsons":
			fn = append(fn, jsonMarshalProcess(true))
		}
	}
	return
}

func (m *Model) valuesBeforeProcess(data ztype.Map) (ztype.Map, error) {
	var err error
	for k := range m.cryptKeys {
		if _, ok := data[k]; ok {
			data[k], err = m.cryptKeys[k](data.Get(k).String())
			if err != nil {
				return nil, err
			}
		}
	}

	for name, fns := range m.beforeProcess {
		val := data.Get(name)
		if !val.Exists() {
			continue
		}
		v := val.Value()
		for _, fn := range fns {
			v, err = fn(v)
			if err != nil {
				return data, err
			}

		}
		_ = data.Set(name, v)
	}

	return data, nil
}

type afterProcess func(string) (interface{}, error)

func (m *Model) GetAfterProcess(p []string) (fn []afterProcess, err error) {
	for _, v := range p {
		switch strings.ToLower(v) {
		default:
			v := strings.SplitN(v, "|", 2)
			if len(v) == 2 {
				switch v[0] {
				case "date":
					if v[1] != "-" {
						fn = append(fn, dateUnmarshalProcess(v[1]))
					}
					continue
				}
			}
			return nil, errors.New("after name not found")
		case "json":
			fn = append(fn, jsonUnmarshalProcess(false))
		case "bool":
			fn = append(fn, boolUnmarshalProcess)
		case "jsons":
			fn = append(fn, jsonUnmarshalProcess(true))
		}
	}
	return
}
