package model

import (
	"errors"
	"github.com/zlsgo/app_module/database/hashid"
	"strings"

	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
	"golang.org/x/crypto/bcrypt"
)

type CryptProcess func(string) (string, error)

func (m *Model) GetCryptProcess(cryptName string) (fn CryptProcess, err error) {
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
func (m *Model) DeCrypt(row ztype.Map) (err error) {
	if m.model.Options.CryptID {
		if id, ok := row[IDKey]; ok {
			switch i := id.(type) {
			case string:
				row[IDKey], err = hashid.DecryptID(m.Hashid, i)
			case []interface{}:
				for k, v := range i {
					i[k], err = hashid.DecryptID(m.Hashid, ztype.ToString(v))
				}
			case []string:
				ids := make([]int64, len(i))
				for k, v := range i {
					ids[k], err = hashid.DecryptID(m.Hashid, v)
				}
				row[IDKey] = ids
			default:
				row[IDKey], err = hashid.DecryptID(m.Hashid, ztype.ToString(id))
			}
		}
	}

	return

}

// EnCrypt  加密 ID
func (m *Model) EnCrypt(row *ztype.Map) (err error) {
	if m.model.Options.CryptID {
		if _, ok := (*row)[IDKey]; ok {
			(*row)[IDKey], err = hashid.EncryptID(m.Hashid, (*row).Get(IDKey).Int64())
		}
	}

	return
}

// EnCryptID  加密 ID
func (m *Model) EnCryptID(id string) (nid string, err error) {
	i := ztype.ToInt64(id)
	if m.model.Options.CryptID && id != "" {
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
func (m *Model) DeCryptID(nid string) (id string, err error) {
	if m.model.Options.CryptID && nid != "" {
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
