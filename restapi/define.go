package restapi

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/zlsgo/app_module/database/hashid"

	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/builder"
)

type (
	Define struct {
		Fields        Fields                           `json:"fields"`
		Extend        ztype.Map                        `json:"extend"`
		Relations     map[string]*ModelRelation        `json:"relations"`
		MigrationDone func(db *zdb.DB, m *Model) error `json:"-"`
		Table         Table                            `json:"table"`
		Name          string                           `json:"name"`
		Values        ztype.Maps                       `json:"values"`
		Options       ModelOptions                     `json:"options"`
	}

	Model struct {
		model         Define
		Storage       Storageer
		cryptKeys     map[string]CryptProcess
		Hashid        *hashid.HashID `json:"-"`
		afterProcess  map[string][]afterProcess
		beforeProcess map[string][]beforeProcess
		views         ztype.Map
		tablePrefix   string
		alias         string
		JSONPath      string
		fullFields    []string
		lowFields     []string
		readOnlyKeys  []string
		inlayFields   []string
		JSON          []byte
		Fields        []string `json:"-"`
		StorageType   StorageType
	}

	Table struct {
		Name    string `json:"name"`
		Comment string `json:"comment"`
	}

	ModelOptions struct {
		Salt             string   `json:"crypt_salt"`
		LowFields        []string `json:"low_fields"`
		FieldsSort       []string `json:"fields_sort"`
		CryptLen         int      `json:"crypt_len"`
		DisabledMigrator bool     `json:"disabled_migrator"`
		SoftDeletes      bool     `json:"soft_deletes"`
		Timestamps       bool     `json:"timestamps"`
		CryptID          bool     `json:"crypt_id"`
	}

	Validations struct {
		Args    interface{} `json:"args"`
		Method  string      `json:"method"`
		Message string      `json:"message"`
	}

	ColumnEnum struct {
		Value string `json:"value"`
		Label string `json:"label"`
	}
)

const (
	IDKey        = "_id"
	CreatedAtKey = "created_at"
	// CreatedByKey = "created_by"
	UpdatedAtKey = "updated_at"
	DeletedAtKey = "deleted_at"
)

func init() {
	builder.IDKey = IDKey
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
