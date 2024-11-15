package model

import (
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb"
)

type (
	SchemaOptions struct {
		DisabledMigrator bool `json:"disabled_migrator,omitempty"`
		SoftDeletes      bool `json:"soft_deletes,omitempty"`
		Timestamps       bool `json:"timestamps,omitempty"`
		CryptID          bool `json:"crypt_id,omitempty"`
	}
	Options struct {
		SetDB              func() (*zdb.DB, error)
		SetAlternateModels func() ([]*Store, error)
		SchemaMiddleware   func() []znet.Handler
		Prefix             string
		SchemaDir          string
		SchemaApi          string
		Schemas            schema.Schemas
		SchemaOptions
	}
)

func New(o ...func(*Options)) (m *Module) {
	opt := zutil.Optional(Options{Prefix: "model_", SchemaDir: "data/schemas", Schemas: make([]schema.Schema, 0)}, o...)

	m = &Module{
		Options: opt,
		ModuleLifeCycle: service.ModuleLifeCycle{
			Service: &service.ModuleService{},
			OnDone: func(di zdi.Invoker) error {
				return initModels(m, di)
			},
		},
	}

	if opt.SchemaApi != "" {
		m.Service.Controllers = append(m.Service.Controllers, &schemaController{
			Path:       opt.SchemaApi,
			module:     m,
			middleware: opt.SchemaMiddleware,
		})
	}

	return m
}
