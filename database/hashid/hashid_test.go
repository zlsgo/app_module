package hashid

import (
	"errors"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/zlsgo/app_module/internal/testutil"
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

	// 注册表是进程级全局的，用唯一名称保证 `-count=N` 重复运行时不报重复注册
	name := testutil.UniqueName(t, "hashid")

	tt.Equal(false, h.Has(name))

	tt.NoError(Set(name, "1"))
	tt.EqualTrue(Set(name, "1") != nil)

	h, ok := Get(name)
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
