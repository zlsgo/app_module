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
		"password":      123,
	}
	g := &inside{}
	err := g.fixUserData(&data)
	tt.NoError(err)
	tt.Equal(false, data.Get("administrator").Exists())
	tt.Equal("123", data.Get("password").String())
	tt.Equal([]string{"admin"}, data.Get("role").SliceString())
	tt.Log(data)

	data = ztype.Map{
		"administrator": true,
		"inlay":         true,
		"role":          []string{"admin"},
		"password":      123,
	}
	err = g.fixUserData(&data)
	tt.NoError(err)
	tt.Equal(false, data.Get("administrator").Exists())
	tt.Equal("123", data.Get("password").String())
	tt.Equal([]string{"admin"}, data.Get("role").SliceString())
	tt.Log(data)
}
