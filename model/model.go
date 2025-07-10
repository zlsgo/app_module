package model

import (
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model/schema"
)

type (
	SchemaOptions struct {
		// DisabledMigrator 禁用自动迁移
		DisabledMigrator bool `z:"disabled_migrator,omitempty"`
		// SoftDeletes 开启软删除
		SoftDeletes bool `z:"soft_deletes,omitempty"`
		// Timestamps 注入创建/更新时间
		Timestamps bool `z:"timestamps,omitempty"`
		// CryptID 加密 ID
		CryptID bool `z:"crypt_id,omitempty"`
	}
	Options struct {
		// SetStorageer 手动设置数据库
		SetStorageer func() (Storageer, error)
		// SetAlternateModels 动态设置关联表
		SetAlternateModels func() ([]*Store, error)
		// SchemaMiddleware 自定义 schema 中间件
		SchemaMiddleware func() []znet.Handler
		// Prefix 模型前缀
		Prefix string
		// SchemaDir 模型定义目录
		SchemaDir string
		// SchemaApi schema api 路径
		SchemaApi string
		// Schemas 定义模型
		Schemas schema.Schemas
		// SchemaOptions 模型选项
		SchemaOptions
	}
)

func New(o ...func(*Options)) (m *Module) {
	opt := zutil.Optional(Options{Prefix: "model_", SchemaDir: "", Schemas: make([]schema.Schema, 0)}, o...)

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
