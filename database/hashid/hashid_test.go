package hashid

import (
	"testing"

	"github.com/sohaha/zlsgo"
)

func TestNew(t *testing.T) {
	tt := zlsgo.NewTest(t)

	h := New("", 6)
	str, err := EncryptID(h, 1)
	tt.NoError(err)
	tt.Equal("lejRej", str)
	i, err := DecryptID(h, str)
	tt.NoError(err)
	tt.Equal(int64(1), i)

	h = New("1", 6)
	str, err = EncryptID(h, 1)
	tt.NoError(err)
	tt.Equal("4dkLrv", str)
	i, err = DecryptID(h, str)
	tt.NoError(err)
	tt.Equal(int64(1), i)
}

func TestSetGet(t *testing.T) {
	tt := zlsgo.NewTest(t)

	tt.Equal(false, h.Has("test"))

	tt.NoError(Set("test", "1"))
	tt.EqualTrue(Set("test", "1") != nil)

	h, ok := Get("test")
	tt.Equal(true, ok)
	tt.EqualTrue(h != nil)
}
