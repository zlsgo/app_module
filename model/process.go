package model

import (
	"errors"
	"strconv"
	"time"

	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
)

func jsonMarshalProcess(isArray bool) func(s interface{}) (string, error) {
	return func(s interface{}) (j string, err error) {
		defer func() {
			if j == "" {
				if isArray {
					j = "[]"
				} else {
					j = "{}"
				}
			} else {
				if isArray {
					if j[0] != '[' {
						err = errors.New("json must be array")
					}
				} else {
					if j[0] != '{' {
						err = errors.New("json must be object")
					}
				}
			}
		}()

		switch v := s.(type) {
		case string:
			if zjson.Valid(v) {
				return v, nil
			}
			return
		case ztype.Type:
			s = v.Value()
		}

		var jb []byte
		jb, err = zjson.Marshal(s)
		if err != nil {
			return
		}
		j = zstring.Bytes2String(jb)
		return
	}
}

func jsonUnmarshalProcess(isArray bool) func(s string) (interface{}, error) {
	return func(s string) (interface{}, error) {
		j := zjson.Parse(s)
		if s == "" {
			if isArray {
				return ztype.Maps{}, nil
			}
			return ztype.Map{}, nil
		}
		if !j.Exists() {
			return nil, errors.New("json parse error")
		}
		if j.IsArray() {
			return j.Slice().Value(), nil
		}
		return j.Map(), nil
	}
}

func boolMarshalProcess(s interface{}) (string, error) {
	return strconv.Itoa(ztype.ToInt(s)), nil
}

func boolUnmarshalProcess(s string) (interface{}, error) {
	return ztype.ToBool(s), nil
}

func dateMarshalProcess(format string) func(v interface{}) (string, error) {
	return func(v interface{}) (string, error) {
		if t, ok := v.(time.Time); ok {
			return ztime.FormatTime(t, format), nil
		}

		s := ztype.ToString(v)
		if s == "" {
			return "", nil
		}

		t, err := ztime.Parse(s, format)
		if err != nil {
			timestamp, err := strconv.Atoi(s)
			if err == nil {
				return ztime.FormatTimestamp(int64(timestamp)), nil
			}
			return "", errors.New("date parse error")
		}
		return ztime.FormatTime(t), nil
	}
}

func dateUnmarshalProcess(format string) func(v string) (interface{}, error) {
	return func(v string) (interface{}, error) {
		if v == "" {
			return "", nil
		}
		timestamp, err := strconv.Atoi(v)

		if err == nil {
			return ztime.FormatTimestamp(int64(timestamp), format), nil
		}

		t, err := ztime.Parse(v)
		if err != nil {
			return "", err
		}
		return ztime.FormatTime(t, format), nil
	}
}
