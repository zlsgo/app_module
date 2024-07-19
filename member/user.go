package member

import (
	"errors"

	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/ztype"
)

type User struct {
	Id   string    `json:"id"`
	Info ztype.Map `json:"info"`
}

var userCache = zcache.NewFast()

func (m *Module) UserById(id any) (ztype.Map, error) {
	mod, ok := m.mods.Get(modelName)
	if !ok {
		return nil, errors.New("not found model")
	}

	info, ok := userCache.ProvideGet(ztype.ToString(id), func() (interface{}, bool) {
		user, err := mod.Operation().FindOneByID(id)
		if err != nil {
			return nil, false
		}

		if user.IsEmpty() {
			return nil, false
		}

		return user, true
	})

	if !ok {
		return nil, errors.New("not found user")
	}

	return ztype.ToMap(info), nil
}
