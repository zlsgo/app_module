package model

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb"
	zdbschema "github.com/zlsgo/zdb/schema"
	"golang.org/x/crypto/bcrypt"
)

type testCrypter struct{}

func (testCrypter) Encrypt(id int64) (string, error) {
	return "x", nil
}

func (testCrypter) Decrypt(encrypted string) (int64, error) {
	return 42, nil
}

func TestPageDataStringAndMap(t *testing.T) {
	tt := zlsgo.NewTest(t)

	data := &PageData{
		Items:    ztype.Maps{{"name": "a"}, {"name": "b"}},
		Page:     PageInfo{},
		pagesize: 2,
	}

	data.Map(func(index int, item ztype.Map) ztype.Map {
		item["idx"] = index
		return item
	}, 1)

	tt.Equal(2, len(data.Items))
	tt.Equal(0, data.Items[0].Get("idx").Int())
	tt.Equal(1, data.Items[1].Get("idx").Int())
	tt.Equal(true, strings.Contains(data.String(), "\"items\""))
}

func TestFindMapsAndFindCol(t *testing.T) {
	tt := zlsgo.NewTest(t)

	people := schema.Schema{
		Name: "people_cover",
		Table: schema.Table{
			Name: "people_cover",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
			"age":  {Type: "int"},
		},
	}

	_, schemas := newTestSchemas(t, people)
	repo := schemas.MustGet("people_cover").Model().Repository()

	id1, err := repo.Insert(ztype.Map{"name": "alice", "age": 20})
	tt.NoError(err)
	_, err = repo.Insert(ztype.Map{"name": "bob", "age": 30})
	tt.NoError(err)

	rows, err := FindMaps(schemas.MustGet("people_cover").Model(), Filter{})
	tt.NoError(err)
	tt.Equal(2, len(rows))

	col, ok, err := FindCol[string](schemas.MustGet("people_cover").Model(), "name", ID(id1))
	tt.NoError(err)
	tt.Equal(true, ok)
	tt.Equal("alice", col)
}

func TestBatchUpdateDelete(t *testing.T) {
	tt := zlsgo.NewTest(t)

	items := schema.Schema{
		Name: "batch_cover",
		Table: schema.Table{
			Name: "batch_cover",
		},
		Fields: map[string]schema.Field{
			"name":   {Type: "string", Size: 80},
			"status": {Type: "int"},
		},
	}

	_, schemas := newTestSchemas(t, items)
	repo := schemas.MustGet("batch_cover").Model().Repository()

	for i := 0; i < 5; i++ {
		_, err := repo.Insert(ztype.Map{"name": fmt.Sprintf("n%d", i), "status": 0})
		tt.NoError(err)
	}

	updated, err := repo.BatchUpdate(Q(ztype.Map{"status": 0}), ztype.Map{"status": 1}, BatchSize(2))
	tt.NoError(err)
	tt.Equal(int64(5), updated)

	rows, err := repo.Find(Q(ztype.Map{}))
	tt.NoError(err)
	tt.Equal(5, len(rows))
	for i := range rows {
		tt.Equal(1, rows[i].Get("status").Int())
	}

	deleted, err := repo.BatchDelete(Q(ztype.Map{"status": 1}), BatchSize(2))
	tt.NoError(err)
	tt.Equal(int64(5), deleted)

	rows, err = repo.Find(Q(ztype.Map{}))
	tt.NoError(err)
	tt.Equal(0, len(rows))
}

func TestCryptHelpers(t *testing.T) {
	tt := zlsgo.NewTest(t)
	b := true

	users := schema.Schema{
		Name: "crypt_cover",
		Table: schema.Table{
			Name: "crypt_cover",
		},
		Options: schema.Options{
			CryptID: &b,
			Salt:    "salt",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	db, schemas := newTestSchemas(t, users)
	m := schemas.MustGet("crypt_cover")
	store := m.Model()

	fn, err := m.GetCryptProcess("md5")
	tt.NoError(err)
	hash, err := fn("abc")
	tt.NoError(err)
	tt.Equal(32, len(hash))
	tt.Equal(true, hash != "abc")

	fn, err = m.GetCryptProcess("password")
	tt.NoError(err)
	hash, err = fn("pass")
	tt.NoError(err)
	tt.NoError(bcrypt.CompareHashAndPassword([]byte(hash), []byte("pass")))

	_, err = m.GetCryptProcess("missing")
	tt.Equal(true, err != nil)

	encrypted, err := m.EnCryptID("1")
	tt.NoError(err)
	decrypted, err := m.DeCryptID(encrypted)
	tt.NoError(err)
	tt.Equal("1", decrypted)

	m.SetIDCrypter(testCrypter{})
	encrypted, err = m.EnCryptID("2")
	tt.NoError(err)
	tt.Equal("x", encrypted)
	decrypted, err = m.DeCryptID("x")
	tt.NoError(err)
	tt.Equal("42", decrypted)

	engine := GetEngine[*zdb.DB](store)
	tt.Equal(true, engine != nil)
	tt.Equal(db, engine)
}

func TestDataTimeHelpers(t *testing.T) {
	tt := zlsgo.NewTest(t)

	var dt DataTime
	err := dt.UnmarshalJSON([]byte("\"\""))
	tt.NoError(err)
	tt.Equal("0000-00-00 00:00:00", dt.String())

	encoded, err := dt.MarshalJSON()
	tt.NoError(err)
	tt.Equal("null", string(encoded))

	now := time.Date(2025, 1, 2, 3, 4, 5, 0, time.UTC)
	dt = DataTime{Time: now}
	encoded, err = dt.MarshalJSON()
	tt.NoError(err)
	tt.Equal(true, string(encoded) != "null")

	val, err := dt.Value()
	tt.NoError(err)
	valTime, ok := val.(time.Time)
	tt.Equal(true, ok)
	tt.Equal(now, valTime)

	var dt2 DataTime
	err = dt2.Scan(now)
	tt.NoError(err)
	err = dt2.Scan([]byte("2025-01-02 03:04:05"))
	tt.NoError(err)
	err = dt2.Scan(123)
	tt.Equal(true, err != nil)
}

func TestModelErrors(t *testing.T) {
	tt := zlsgo.NewTest(t)

	root := errors.New("root")
	me := NewModelError("insert", "t", root, "msg")
	tt.Equal(true, strings.Contains(me.Error(), "model: insert"))
	tt.Equal(root, errors.Unwrap(me))

	qe := NewQueryError("sql", []any{1}, root, "bad")
	tt.Equal(true, strings.Contains(qe.Error(), "query error"))
	tt.Equal(root, errors.Unwrap(qe))

	he := NewHookError("before", root)
	tt.Equal(true, strings.Contains(he.Error(), "hook before"))
	tt.Equal(root, errors.Unwrap(he))
}

func TestDefaultSchemaOptions(t *testing.T) {
	tt := zlsgo.NewTest(t)

	old := DefaultSchemaOptions
	SetDefaultSchemaOptions(func(o *SchemaOptions) {
		o.OldColumn = DealOldColumnDelete
		o.SoftDeleteIsTime = true
	})
	tt.Equal(DealOldColumnDelete, DefaultSchemaOptions.OldColumn)
	tt.Equal(true, DefaultSchemaOptions.SoftDeleteIsTime)
	DefaultSchemaOptions = old
}

func TestGetFieldHelper(t *testing.T) {
	tt := zlsgo.NewTest(t)
	b := true
	softTime := false

	s := schema.Schema{
		Name: "field_cover",
		Table: schema.Table{
			Name: "field_cover",
		},
		Options: schema.Options{
			Timestamps:       &b,
			SoftDeletes:      &b,
			SoftDeleteIsTime: &softTime,
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, s)
	m := schemas.MustGet("field_cover")

	f, ok := m.GetField("name")
	tt.Equal(true, ok)
	tt.Equal(zdbschema.String, f.Type)

	f, ok = m.GetField(IDKey())
	tt.Equal(true, ok)
	tt.Equal(true, f.Options.ReadOnly)

	f, ok = m.GetField(CreatedAtKey)
	tt.Equal(true, ok)
	tt.Equal(zdbschema.Time, f.Type)

	f, ok = m.GetField(UpdatedAtKey)
	tt.Equal(true, ok)
	tt.Equal(zdbschema.Time, f.Type)

	f, ok = m.GetField(DeletedAtKey)
	tt.Equal(true, ok)
	tt.Equal(zdbschema.Uint, f.Type)
}

func TestGetBeforeProcess(t *testing.T) {
	tt := zlsgo.NewTest(t)

	s := &Schema{}
	fns, err := s.GetBeforeProcess([]string{"bool"})
	tt.NoError(err)
	val, err := fns[0](true)
	tt.NoError(err)
	tt.Equal("1", val)

	fns, err = s.GetBeforeProcess([]string{"json"})
	tt.NoError(err)
	val, err = fns[0](ztype.Map{"a": 1})
	tt.NoError(err)
	tt.Equal(true, strings.HasPrefix(val, "{"))

	_, err = s.GetBeforeProcess([]string{"missing"})
	tt.Equal(true, err != nil)
}
