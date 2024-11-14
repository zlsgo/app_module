package model

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/zlsgo/app_module/database/hashid"
	"github.com/zlsgo/app_module/model/schema"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb/builder"
)

type (
	Schema struct {
		define        schema.Schema
		Storage       Storageer
		di            zdi.Injector
		model         *Model
		cryptKeys     map[string]CryptProcess
		Hashid        *hashid.HashID `json:"-"`
		afterProcess  map[string][]afterProcess
		beforeProcess map[string][]beforeProcess
		views         ztype.Map
		getSchema     func(alias string) (*Schema, bool)
		JSONPath      string
		alias         string
		tablePrefix   string
		fullFields    []string
		lowFields     []string
		readOnlyKeys  []string
		relationsKeys []string
		inlayFields   []string
		JSON          []byte
		fields        []string `json:"-"`
		StorageType   StorageType
	}

	ColumnEnum struct {
		Value string `json:"value"`
		Label string `json:"label"`
	}

	Model struct {
		schema *Schema
	}
)

const (
	CreatedAtKey = "created_at"
	// CreatedByKey = "created_by"
	UpdatedAtKey = "updated_at"
	DeletedAtKey = "deleted_at"
)

var idKey = builder.IDKey

func IDKey() string {
	return idKey
}

const deleteFieldPrefix = "__del__"

type DataTime struct {
	time.Time
}

func (t *DataTime) UnmarshalJSON(data []byte) error {
	if len(data) == 2 {
		*t = DataTime{Time: time.Time{}}
		return nil
	}
	now, err := ztime.Parse(zstring.Bytes2String(data))
	if err != nil {
		return err
	}
	*t = DataTime{Time: now}
	return nil
}

func (t DataTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return zstring.String2Bytes(ztime.FormatTime(t.Time, "\"Y-m-d H:i:s\"")), nil
}

func (t DataTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.IsZero() || t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *DataTime) String() string {
	if t.Time.IsZero() {
		return "0000-00-00 00:00:00"
	}
	return ztime.FormatTime(t.Time)
}

func (t *DataTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = DataTime{Time: value}
		return nil
	}
	if b, ok := v.([]byte); ok {
		parse, err := ztime.Parse(zstring.Bytes2String(b))
		if err != nil {
			return err
		}
		*t = DataTime{Time: parse}
		return nil
	}

	return fmt.Errorf("can not convert %v to timestamp", v)
}
