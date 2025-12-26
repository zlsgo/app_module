package account

import (
	"testing"

	"github.com/sohaha/zlsgo"
)

func TestGetUserForCacheSanitized(t *testing.T) {
	tt := zlsgo.NewTest(t)
	mod, cleanup := newTestModule(t)
	defer cleanup()

	uid, err := mod.accountModel.Schema().EnCryptID("1")
	tt.NoError(err)

	info, err := mod.getUserForCache(mod.index.accoutModel, uid)
	tt.NoError(err)
	tt.Equal(false, info.Has("password"))
	tt.Equal(false, info.Has("salt"))
}
