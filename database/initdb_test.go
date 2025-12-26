package database

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb/builder"
	"github.com/zlsgo/zdb/driver"
	"github.com/zlsgo/zdb/schema"
)

type dialectlessConfig struct{}

func (c *dialectlessConfig) GetDsn() string           { return "" }
func (c *dialectlessConfig) SetDsn(string)            {}
func (c *dialectlessConfig) GetDriver() string        { return "" }
func (c *dialectlessConfig) DB() *sql.DB              { return nil }
func (c *dialectlessConfig) MustDB() (*sql.DB, error) { return nil, nil }
func (c *dialectlessConfig) SetDB(*sql.DB)            {}

type dialectA struct{}

func (d dialectA) Value() driver.Typ { return driver.MySQL }
func (d dialectA) DataTypeOf(field *schema.Field, only ...bool) string {
	return ""
}
func (d dialectA) HasTable(table string) (string, []interface{}, func(ztype.Maps) bool) {
	return "", nil, func(ztype.Maps) bool { return false }
}
func (d dialectA) GetColumn(table string) (string, []interface{}, func(result ztype.Maps) ztype.Map) {
	return "", nil, func(ztype.Maps) ztype.Map { return ztype.Map{} }
}
func (d dialectA) RenameColumn(table, oldName, newName string) (string, []interface{}) {
	return "", nil
}
func (d dialectA) HasIndex(table, name string) (string, []interface{}, func(ztype.Maps) bool) {
	return "", nil, func(ztype.Maps) bool { return false }
}
func (d dialectA) RenameIndex(table, oldName, newName string) (string, []interface{}) {
	return "", nil
}
func (d dialectA) CreateIndex(table, name string, columns []string, indexType string) (string, []interface{}) {
	return "", nil
}

type dialectB struct{ dialectA }

func (d dialectB) Value() driver.Typ { return driver.SQLite }

func TestInitDBNoDriver(t *testing.T) {
	tt := zlsgo.NewTest(t)

	_, err := initDB(Options{})
	tt.Equal(true, err != nil)
	if err != nil {
		tt.Equal(false, strings.Contains(err.Error(), "unknown"))
	}
}

func TestInitDBUnknownDriver(t *testing.T) {
	tt := zlsgo.NewTest(t)

	_, err := initDB(Options{Driver: "unknown"})
	tt.Equal(true, err != nil)
	if err != nil {
		tt.Equal(true, strings.Contains(err.Error(), "[unknown]"))
	}
}

func TestInitDBSqliteMissingPath(t *testing.T) {
	if _, ok := drivers["sqlite"]; !ok {
		t.Skip("sqlite driver not registered")
	}

	tt := zlsgo.NewTest(t)

	_, err := initDB(Options{Driver: "sqlite", Sqlite: &Sqlite{Path: ""}})
	tt.Equal(true, err != nil)
	if err != nil {
		var zerr *zerror.Error
		tt.Equal(true, errors.As(err, &zerr))
	}
}

func TestInitDBSqliteSuccess(t *testing.T) {
	if _, ok := drivers["sqlite"]; !ok {
		t.Skip("sqlite driver not registered")
	}

	tt := zlsgo.NewTest(t)

	path := filepath.Join(t.TempDir(), "test.db")
	db, err := initDB(Options{Driver: "sqlite", Sqlite: &Sqlite{Path: path}})
	tt.NoError(err)
	tt.Equal(true, db != nil)
	if err == nil {
		_, statErr := os.Stat(path)
		tt.NoError(statErr)
		tt.NoError(db.Close())
	}
}

func TestInitDBAutoSelectSqlite(t *testing.T) {
	if _, ok := drivers["sqlite"]; !ok {
		t.Skip("sqlite driver not registered")
	}

	tt := zlsgo.NewTest(t)

	path := filepath.Join(t.TempDir(), "auto.db")
	db, err := initDB(Options{Sqlite: &Sqlite{Path: path}})
	tt.NoError(err)
	tt.Equal(true, db != nil)
	if err == nil {
		_, statErr := os.Stat(path)
		tt.NoError(statErr)
		tt.NoError(db.Close())
	}
}

func TestInitDBMultipleConfigs(t *testing.T) {
	tt := zlsgo.NewTest(t)

	_, err := initDB(Options{
		MySQL:  &Mysql{Host: "127.0.0.1"},
		Sqlite: &Sqlite{Path: "data/test.db"},
	})
	tt.Equal(true, err != nil)
	if err != nil {
		tt.Equal(true, strings.Contains(err.Error(), "多个数据库配置"))
	}
}

func TestInitDBDialectMissing(t *testing.T) {
	tt := zlsgo.NewTest(t)

	orig := copyDrivers()
	defer func() {
		driversMu.Lock()
		drivers = orig
		driversMu.Unlock()
	}()

	name := strings.ToLower("dialectless_" + t.Name())
	tt.NoError(Register(name, func(Options) (driver.IfeConfig, error) {
		return &dialectlessConfig{}, nil
	}))

	_, err := initDB(Options{Driver: name})
	tt.Equal(true, err != nil)
	if err != nil {
		tt.Equal(true, strings.Contains(err.Error(), "Dialect"))
	}
}

func TestInitDBSqliteCreateDir(t *testing.T) {
	if _, ok := drivers["sqlite"]; !ok {
		t.Skip("sqlite driver not registered")
	}

	tt := zlsgo.NewTest(t)

	path := filepath.Join(t.TempDir(), "nested", "test.db")
	db, err := initDB(Options{Driver: "sqlite", Sqlite: &Sqlite{Path: path}})
	tt.NoError(err)
	tt.Equal(true, db != nil)
	if err == nil {
		_, statErr := os.Stat(path)
		tt.NoError(statErr)
		tt.NoError(db.Close())
	}
}

func TestSetBuilderDialectMultiple(t *testing.T) {
	tt := zlsgo.NewTest(t)

	tt.NoError(setBuilderDialect(dialectA{}))
	tt.NoError(setBuilderDialect(dialectB{}))
	tt.Equal(driver.SQLite, builder.DefaultDriver.Value())
}
