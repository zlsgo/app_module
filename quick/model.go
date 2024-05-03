package quick

import (
	"errors"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/zlsgo/app_module/quick/define"
	"github.com/zlsgo/app_module/quick/storage"
	"github.com/zlsgo/app_module/quick/utils"
)

type (
	Models struct {
		models  *zarray.Maper[string, *Quick]
		storage storage.Storage
		options ModelsOptions
	}

	ModelsOptions struct {
		Prefix string
	}
)

func NewModels(s storage.Storage, fn ...func(ModelsOptions) ModelsOptions) *Models {
	options := utils.Optional(ModelsOptions{Prefix: "model_"}, fn...)

	return &Models{
		storage: s,
		options: options,
		models:  zarray.NewHashMap[string, *Quick](),
	}
}

func (m *Models) Reg(d define.Define, force bool) (q *Quick, err error) {
	q, err = New(m.storage, d, func(options Options) Options {
		options.TablePrefix = m.options.Prefix
		return options
	})
	if err != nil {
		return nil, err
	}

	if !force && m.models.Has(d.Name) {
		return nil, errors.New("model " + d.Name + " has been registered")
	}

	m.models.Set(d.Name, q)
	return
}

func (m *Models) Get(name string) (*Quick, bool) {
	return m.models.Get(name)
}
