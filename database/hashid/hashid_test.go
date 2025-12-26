package hashid

import (
	"errors"
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

func TestNilHashID(t *testing.T) {
	tt := zlsgo.NewTest(t)

	_, err := EncryptID(nil, 1)
	tt.Equal(true, err != nil)

	_, err = DecryptID(nil, "x")
	tt.Equal(true, err != nil)
}

func TestHashIDErrorPropagation(t *testing.T) {
	tt := zlsgo.NewTest(t)

	expected := errors.New("init failed")
	h := &HashID{err: expected}

	_, err := EncryptID(h, 1)
	tt.Equal(expected, err)

	_, err = DecryptID(h, "x")
	tt.Equal(expected, err)
}
