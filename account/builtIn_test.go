package account

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
)

func Test_fixUserData(t *testing.T) {
	tt := zlsgo.NewTest(t)

	data := ztype.Map{
		"administrator": true,
		"inlay":         true,
		"role":          "admin",
		"password":      "Aa123456",
	}
	g := &inside{}
	err := g.fixUserData(&data)
	tt.NoError(err)
	tt.Equal(false, data.Get("administrator").Exists())
	tt.Equal("Aa123456", data.Get("password").String())
	tt.Equal([]string{"admin"}, data.Get("role").SliceString())
	tt.Log(data)

	data = ztype.Map{
		"administrator": true,
		"inlay":         true,
		"role":          []string{"admin"},
		"password":      "Aa123456",
	}
	err = g.fixUserData(&data)
	tt.NoError(err)
	tt.Equal(false, data.Get("administrator").Exists())
	tt.Equal("Aa123456", data.Get("password").String())
	tt.Equal([]string{"admin"}, data.Get("role").SliceString())
	tt.Log(data)
}
