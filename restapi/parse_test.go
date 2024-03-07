package restapi

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

func TestSet(t *testing.T) {

	tt := zlsgo.NewTest(t)
	data := Define{
		Name: "日志模型",
		Table: Table{
			Name:    "lowcode_logs",
			Comment: "日志表",
		},
		ModelOptions: ModelOptions{
			Timestamps: true,
			CryptID:    true,
		},
	}

	data.Fields = map[string]Field{
		"action": {
			Type:    "string",
			Label:   "操作",
			Default: "",
			Validations: []Validations{
				{
					Method: "minLength",
					Args:   1,
				},
				{
					Method: "maxLength",
					Args:   60,
				},
			},
		},
		"ip": {
			Type:    "string",
			Label:   "请求 IP",
			Size:    100,
			Default: "",
			Validations: []Validations{
				{Method: "ip"},
			},
		},
		"status": {
			Type:    "int8",
			Label:   "状态",
			Size:    9,
			Default: "1",
			ModelOptions: FieldOption{
				// Quote: true,
				Enum: []FieldEnum{
					{Value: "1", Label: "未读"},
					{Value: "2", Label: "已读"},
				},
			},
		},
	}

	db, err := zdb.New(&sqlite3.Config{
		File:       "data/db.db",
		Memory:     true,
		Parameters: "_pragma=busy_timeout(3000)",
	})
	model := New(NewSQL(db), func(o *ModelOptions) {
		o.Prefix = "model_"
	})

	m, err := model.Reg("test_model", data, false)
	tt.NoError(err)

	zdb.Debug.Store(true)

	err = m.Migration().Auto(DealOldColumnNone)
	tt.NoError(err)

	tt.Equal(m.Name(), "日志模型")

	id, err := Insert(m, map[string]interface{}{"action": "test", "ip": "127.0.0.1", "status": 1})
	tt.NoError(err)
	tt.Log(id)

	id, err = Insert(m, map[string]interface{}{"action": "demo", "ip": "127.0.0.2", "status": "2"})
	tt.NoError(err)
	tt.Log(id)

	_, _ = Insert(m, map[string]interface{}{"action": "demo", "ip": "127.0.0.3", "status": "1"})

	row, err := FindOne(m, ztype.Map{}, func(ModelOptions *CondOptions) error {
		ModelOptions.OrderBy = [][]string{{IDKey, "DESC"}}
		ModelOptions.Fields = []string{IDKey, "status"}
		return nil
	})
	tt.NoError(err)
	tt.Log(row)

	total, err := Update(m, row.Get(IDKey).String(), ztype.Map{"ip": "192.168.0.1", "status": 1})
	tt.NoError(err)
	tt.Log(total)

	row, err = FindOne(m, ztype.Map{}, func(ModelOptions *CondOptions) error {
		ModelOptions.OrderBy = [][]string{{IDKey, "DESC"}}
		return nil
	})
	tt.NoError(err)
	tt.Log(row)
}
