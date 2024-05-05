package crud

import (
	"errors"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/zlsgo/app_module/quick/define"
	"github.com/zlsgo/app_module/quick/storage"
	"github.com/zlsgo/app_module/quick/utils"
)

type (
	Models struct {
		models  *zarray.Maper[string, *Crud]
		storage storage.Storage
		options define.ModelsOptions
	}
)

func NewModels(s storage.Storage, fn ...func(define.ModelsOptions) define.ModelsOptions) *Models {
	return &Models{
		storage: s,
		options: utils.Optional(define.ModelsOptions{Prefix: "model_"}, fn...),
		models:  zarray.NewHashMap[string, *Crud](),
	}
}

func (m *Models) Reg(d define.Define, force bool) (q *Crud, err error) {
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

func (m *Models) Get(name string) (*Crud, bool) {
	return m.models.Get(name)
}
