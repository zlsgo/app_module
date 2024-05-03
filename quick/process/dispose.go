package process

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/ztype"
)

type BeforeProcess func(interface{}) (string, error)

func GetBeforeProcess(p []string) (fn []BeforeProcess, err error) {
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

func (p *Process) ValuesBeforeProcess(data ztype.Map) (ztype.Map, error) {
	var err error
	for k := range p.CryptKeys {
		if _, ok := data[k]; ok {
			data[k], err = p.CryptKeys[k](data.Get(k).String())
			if err != nil {
				return nil, err
			}
		}
	}

	for name, fns := range p.BeforeProcess {
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

type AfterProcess func(string) (interface{}, error)

func GetAfterProcess(p []string) (fn []AfterProcess, err error) {
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
