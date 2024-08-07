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
	Options struct {
		SetDB              func() (*zdb.DB, error)
		SetAlternateModels func() ([]*Model, error)
		SchemaMiddleware   func() []znet.Handler
		Prefix             string
		SchemaDir          string
		SchemaApi          string
		Schemas            schema.Schemas
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
