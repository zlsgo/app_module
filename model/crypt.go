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
func (m *Schema) DeCrypt(row ztype.Map) (success bool) {
	if m.define.Options.CryptID {
		if id, ok := row[idKey]; ok {
			var (
				err error
				raw int64
			)
			success = true
			switch i := id.(type) {
			case string:
				raw, err = hashid.DecryptID(m.Hashid, i)
				if err == nil {
					row[idKey] = raw
				} else {
					success = false
				}

			case []interface{}:
				for k, v := range i {
					raw, err = hashid.DecryptID(m.Hashid, ztype.ToString(v))
					if err == nil {
						i[k] = raw
					} else {
						success = false
					}
				}

			case []string:
				ids := make([]int64, len(i))
				for k, v := range i {
					raw, err = hashid.DecryptID(m.Hashid, v)
					if err == nil {
						ids[k] = raw
					} else {
						success = false
					}
				}
				row[idKey] = ids

			default:
				raw, err = hashid.DecryptID(m.Hashid, ztype.ToString(id))
				if err == nil {
					row[idKey] = raw
				} else {
					success = false
				}
			}
		}
	}

	return true
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
