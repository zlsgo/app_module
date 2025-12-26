package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

func TestCryptIDFilterOperators(t *testing.T) {
	tt := zlsgo.NewTest(t)

	b := true
	users := schema.Schema{
		Name: "crypt_users",
		Table: schema.Table{
			Name: "crypt_users",
		},
		Options: schema.Options{
			CryptID: &b,
			Salt:    "test-salt",
			CryptLen: 8,
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string"},
		},
	}

	_, schemas := newTestSchemas(t, users)
	repo := schemas.MustGet("crypt_users").Model().Repository()

	id1, err := repo.Insert(ztype.Map{"name": "a"})
	tt.NoError(err)
	id2, err := repo.Insert(ztype.Map{"name": "b"})
	tt.NoError(err)
	_, err = repo.Insert(ztype.Map{"name": "c"})
	tt.NoError(err)

	rows, err := repo.Find(In(IDKey(), []any{id1, id2}))
	tt.NoError(err)
	tt.Equal(2, len(rows))

	rows, err = repo.Find(Ne(IDKey(), id1))
	tt.NoError(err)
	tt.Equal(2, len(rows))

	rows, err = repo.Find(Or(Eq(IDKey(), id1), Eq(IDKey(), id2)))
	tt.NoError(err)
	tt.Equal(2, len(rows))
}

func TestCryptIDFilterQualifiedIDNoDecrypt(t *testing.T) {
	tt := zlsgo.NewTest(t)

	b := true
	users := schema.Schema{
		Name: "crypt_users_join",
		Table: schema.Table{
			Name: "crypt_users_join",
		},
		Options: schema.Options{
			CryptID:  &b,
			Salt:     "test-salt",
			CryptLen: 8,
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string"},
		},
	}

	_, schemas := newTestSchemas(t, users)
	m := schemas.MustGet("crypt_users_join")

	ok := m.DeCrypt(ztype.Map{"other.id": "1"})
	tt.Equal(true, ok)
}
