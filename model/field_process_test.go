package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	mSchema "github.com/zlsgo/app_module/model/schema"
	zdbschema "github.com/zlsgo/zdb/schema"
)

func TestFieldAfterProcessUsesAfterList(t *testing.T) {
	tt := zlsgo.NewTest(t)

	m := &Schema{
		define:        mSchema.Schema{Fields: map[string]mSchema.Field{}},
		cryptKeys:     map[string]CryptProcess{},
		beforeProcess: map[string][]beforeProcess{},
		afterProcess:  map[string][]afterProcess{},
	}

	f := mSchema.Field{
		Type:   zdbschema.String,
		Before: []string{"bool"},
		After:  []string{"json"},
	}

	tt.NoError(parseField(m, "payload", &f))
	fns := m.afterProcess["payload"]
	tt.Equal(true, len(fns) > 0)

	out, err := fns[0]("{\"a\":1}")
	tt.NoError(err)

	mval, ok := out.(ztype.Map)
	tt.Equal(true, ok)
	tt.Equal("1", mval.Get("a").String())
}

func TestFieldAfterProcessChain(t *testing.T) {
	tt := zlsgo.NewTest(t)

	m := &Schema{
		define:        mSchema.Schema{Fields: map[string]mSchema.Field{}},
		cryptKeys:     map[string]CryptProcess{},
		beforeProcess: map[string][]beforeProcess{},
		afterProcess:  map[string][]afterProcess{},
	}

	f := mSchema.Field{
		Type:  zdbschema.String,
		After: []string{"date|Y-m-d H:i:s", "date|Y-m-d"},
	}

	tt.NoError(parseField(m, "event_time", &f))
	fns := m.afterProcess["event_time"]
	tt.Equal(true, len(fns) == 2)

	val := interface{}("2025-12-22 10:11:12")
	var err error
	for i := range fns {
		val, err = fns[i](val)
		tt.NoError(err)
	}

	tt.Equal("2025-12-22", val)
}
