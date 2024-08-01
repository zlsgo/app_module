package member

import (
	"errors"

	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/ztype"
)

type User struct {
	Info  ztype.Map `json:"info"`
	Id    string    `json:"id"`
	RawId string    `json:"-"`
}

var userCache = zcache.NewFast()

func (m *Module) UserById(id any) (u *User, err error) {
	mod, ok := m.mods.Get(modelName)
	if !ok {
		return nil, errors.New("not found model")
	}

	info, ok := userCache.ProvideGet(ztype.ToString(id), func() (interface{}, bool) {
		info, err := mod.Model().FindOneByID(id)
		if err != nil {
			return nil, false
		}

		if info.IsEmpty() {
			return nil, false
		}

		uid := ztype.ToString(id)
		rawId, _ := mod.DeCryptID(uid)
		return &User{Id: uid, Info: info, RawId: rawId}, true
	})

	if !ok || info == nil {
		return nil, errors.New("not found user")
	}

	return info.(*User), nil
}
