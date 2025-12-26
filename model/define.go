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
	// Schema 模型结构定义
	Schema struct {
		define         schema.Schema
		Storage        Storageer
		di             zdi.Injector
		model          *Store
		cryptKeys      map[string]CryptProcess
		Hashid         *hashid.HashID `json:"-"`
		idCrypter      IDCrypter
		afterProcess   map[string][]afterProcess
		beforeProcess  map[string][]beforeProcess
		views          ztype.Map
		getSchema      func(alias string) (*Schema, bool)
		JSONPath       string
		alias          string
		tablePrefix    string
		fullFields     []string
		fullFieldsMap  map[string]struct{}
		fieldsSet      map[string]struct{}
		lowFields      []string
		readOnlyKeys   []string
		inlayFields    []string
		inlayFieldsMap map[string]struct{}
		JSON           []byte
		fields         []string `json:"-"`
		StorageType    StorageType
	}

	// ColumnEnum 列举值结构
	ColumnEnum struct {
		Value string `json:"value"`
		Label string `json:"label"`
	}

	// Store 模型存储实例
	Store struct {
		schema *Schema
	}
)

const (
	// CreatedAtKey 创建时间字段名
	CreatedAtKey = "created_at"
	// CreatedByKey 创建人字段名（已注释）
	// CreatedByKey = "created_by"
	// UpdatedAtKey 更新时间字段名
	UpdatedAtKey = "updated_at"
	// DeletedAtKey 删除时间字段名
	DeletedAtKey = "deleted_at"
)

// idKey 主键字段名变量
var idKey = builder.IDKey

// IDKey 返回数据库表的主键字段名
func IDKey() string {
	return idKey
}

// deleteFieldPrefix 删除字段前缀
const deleteFieldPrefix = "__del__"

// DataTime 自定义时间类型
type DataTime struct {
	time.Time
}

// UnmarshalJSON 将 JSON 数据解析为 DataTime 类型
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

// MarshalJSON 将 DataTime 序列化为 JSON 格式
func (t DataTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return zstring.String2Bytes(ztime.FormatTime(t.Time, "\"Y-m-d H:i:s\"")), nil
}

// Value 返回 driver.Value，实现 driver.Valuer 接口
func (t DataTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.IsZero() || t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// String 返回格式化的时间字符串
func (t *DataTime) String() string {
	if t.Time.IsZero() {
		return "0000-00-00 00:00:00"
	}
	return ztime.FormatTime(t.Time)
}

// Scan 将数据库值扫描到 DataTime，实现 sql.Scanner 接口
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
