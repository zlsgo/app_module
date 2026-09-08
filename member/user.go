package member

import (
	"errors"

	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model"
)

type User struct {
	Info  ztype.Map `json:"info"`
	Id    string    `json:"id"`
	RawId string    `json:"-"`
}

const saltLen = 4

func (m *Module) UserModel() (*model.Store, bool) {
	mod, ok := m.models.Get(modelName)
	if !ok {
		return nil, false
	}
	return mod, true
}

func (m *Module) UserById(id any) (u *User, err error) {
	mod, ok := m.UserModel()
	if !ok {
		return nil, errors.New("not found model")
	}

	if m.userCache == nil {
		m.userCache = zcache.NewFast()
	}
	info, ok := m.userCache.ProvideGet(ztype.ToString(id), func() (interface{}, bool) {
		info, err := mod.FindOneByID(id)
		if err != nil {
			return nil, false
		}

		if info.IsEmpty() {
			return nil, false
		}

		uid := ztype.ToString(id)
		rawId, _ := mod.Schema().DeCryptID(uid)
		return &User{Id: uid, Info: info, RawId: rawId}, true
	})

	if !ok || info == nil {
		return nil, errors.New("not found user")
	}

	return info.(*User), nil
}
