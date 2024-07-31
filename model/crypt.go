package model

import (
	"errors"
	"strings"

	"github.com/zlsgo/app_module/database/hashid"

	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
	"golang.org/x/crypto/bcrypt"
)

type CryptProcess func(string) (string, error)

func (m *Schema) GetCryptProcess(cryptName string) (fn CryptProcess, err error) {
	switch strings.ToLower(cryptName) {
	default:
		return nil, errors.New("crypt name not found")
	case "md5":
		fn = func(s string) (string, error) {
			return zstring.Md5(s), nil
		}
	case "password":
		fn = func(s string) (string, error) {
			bcost := bcrypt.DefaultCost
			bytes, err := bcrypt.GenerateFromPassword(zstring.String2Bytes(s), bcost)
			if err != nil {
				return "", err
			}
			return zstring.Bytes2String(bytes), nil
		}
	}
	return
}

// DeCrypt 解密 ID
func (m *Schema) DeCrypt(row ztype.Map) (err error) {
	if m.define.Options.CryptID {
		if id, ok := row[idKey]; ok {
			switch i := id.(type) {
			case string:
				row[idKey], err = hashid.DecryptID(m.Hashid, i)
			case []interface{}:
				for k, v := range i {
					i[k], err = hashid.DecryptID(m.Hashid, ztype.ToString(v))
				}
			case []string:
				ids := make([]int64, len(i))
				for k, v := range i {
					ids[k], err = hashid.DecryptID(m.Hashid, v)
				}
				row[idKey] = ids
			default:
				row[idKey], err = hashid.DecryptID(m.Hashid, ztype.ToString(id))
			}
		}
	}
	return
}

// EnCrypt  加密 ID
func (m *Schema) EnCrypt(row *ztype.Map) (err error) {
	if m.define.Options.CryptID {
		if _, ok := (*row)[idKey]; ok {
			(*row)[idKey], err = hashid.EncryptID(m.Hashid, (*row).Get(idKey).Int64())
		}
	}

	return
}

// EnCryptID  加密 ID
func (m *Schema) EnCryptID(id string) (nid string, err error) {
	i := ztype.ToInt64(id)
	if m.define.Options.CryptID && id != "" {
		if i == 0 {
			return "", errors.New("id cannot be empty")
		}
		nid, err = hashid.EncryptID(m.Hashid, i)
	} else {
		return id, nil
	}

	return
}

// DeCryptID 解密 ID
func (m *Schema) DeCryptID(nid string) (id string, err error) {
	if m.define.Options.CryptID && nid != "" {
		rid, err := hashid.DecryptID(m.Hashid, ztype.ToString(nid))
		if err != nil {
			return "", err
		}
		id = ztype.ToString(rid)
	} else {
		id = nid
	}

	return
}
