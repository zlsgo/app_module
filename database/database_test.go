package database

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model"
	"github.com/zlsgo/zdb"
)

func TestOptionsHelpers(t *testing.T) {
	tt := zlsgo.NewTest(t)

	o := Options{disableWrite: true}
	tt.Equal("database", o.ConfKey())
	tt.Equal(true, o.DisableWrite())
	o.disableWrite = false
	tt.Equal(false, o.DisableWrite())
}

func TestDBGetters(t *testing.T) {
	tt := zlsgo.NewTest(t)

	p := &Plugin{}
	_, err := p.DB()
	tt.Equal(true, err != nil)

	db := &zdb.DB{}
	p.db = db
	got, err := p.DB()
	tt.NoError(err)
	tt.Equal(db, got)
	tt.Equal("Database", p.Name())

	s := &Single{}
	_, err = s.DB()
	tt.Equal(true, err != nil)

	s.db = db
	got, err = s.DB()
	tt.NoError(err)
	tt.Equal(db, got)
	tt.Equal("SingleDatabase", s.Name())
}

func TestApplyMode(t *testing.T) {
	tt := zlsgo.NewTest(t)

	old := model.DefaultSchemaOptions
	defer func() {
		model.DefaultSchemaOptions = old
	}()

	applyMode(&Mode{DelteColumn: true})
	tt.Equal(model.DealOldColumnDelete, model.DefaultSchemaOptions.OldColumn)

	applyMode(&Mode{DelteColumn: false})
	tt.Equal(model.DealOldColumnNone, model.DefaultSchemaOptions.OldColumn)

	applyMode(nil)
	tt.Equal(model.DealOldColumnNone, model.DefaultSchemaOptions.OldColumn)
}

func TestRegisterDefaultConfReplace(t *testing.T) {
	tt := zlsgo.NewTest(t)

	orig := service.DefaultConf
	defer func() {
		service.DefaultConf = orig
	}()

	service.DefaultConf = nil

	o1 := Options{Driver: "sqlite"}
	registerDefaultConf(o1)
	tt.Equal(1, len(service.DefaultConf))

	o2 := Options{Driver: "mysql"}
	registerDefaultConf(o2)
	tt.Equal(1, len(service.DefaultConf))

	got, ok := service.DefaultConf[0].(Options)
	tt.Equal(true, ok)
	tt.Equal("mysql", got.Driver)
}
