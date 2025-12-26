package model

import (
	"testing"
	"time"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	mSchema "github.com/zlsgo/app_module/model/schema"
	zdbschema "github.com/zlsgo/zdb/schema"
)

func TestVerifiDataBytesInput(t *testing.T) {
	tt := zlsgo.NewTest(t)

	m := &Schema{
		define:        mSchema.Schema{Fields: map[string]mSchema.Field{}},
		cryptKeys:     map[string]CryptProcess{},
		beforeProcess: map[string][]beforeProcess{},
		afterProcess:  map[string][]afterProcess{},
	}

	f := mSchema.Field{
		Type: zdbschema.Bytes,
		Size: 4,
	}
	tt.NoError(parseField(m, "blob", &f))

	columns := mSchema.Fields{"blob": f}
	data, err := VerifiData(ztype.Map{"blob": "abcd"}, columns, activeCreate)
	tt.NoError(err)
	_, ok := data["blob"].([]byte)
	tt.Equal(true, ok)

	_, err = VerifiData(ztype.Map{"blob": "abcde"}, columns, activeCreate)
	tt.Equal(true, err != nil)
}

func TestVerifiDataTimeNumericAndMax(t *testing.T) {
	tt := zlsgo.NewTest(t)

	max := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC).Unix()

	m := &Schema{
		define:        mSchema.Schema{Fields: map[string]mSchema.Field{}},
		cryptKeys:     map[string]CryptProcess{},
		beforeProcess: map[string][]beforeProcess{},
		afterProcess:  map[string][]afterProcess{},
	}

	f := mSchema.Field{
		Type: zdbschema.Time,
		Size: uint64(max),
	}
	tt.NoError(parseField(m, "event_time", &f))

	columns := mSchema.Fields{"event_time": f}
	data, err := VerifiData(ztype.Map{"event_time": max}, columns, activeCreate)
	tt.NoError(err)
	_, ok := data["event_time"].(DataTime)
	if !ok {
		t.Fatalf("expected DataTime, got %T", data["event_time"])
	}

	_, err = VerifiData(ztype.Map{"event_time": max + 3600}, columns, activeCreate)
	tt.Equal(true, err != nil)
}
