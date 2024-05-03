package process

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/database/hashid"
	"github.com/zlsgo/app_module/quick/define"
	"golang.org/x/crypto/bcrypt"
)

type CryptProcess func(string) (string, error)

func GetCryptProcess(cryptName string) (fn CryptProcess, err error) {
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
func (p *Process) DeCrypt(row ztype.Map) (err error) {
	if p.Define.Options.CryptID {
		if id, ok := row[define.Inside.IDKey()]; ok {
			switch i := id.(type) {
			case string:
				row[define.Inside.IDKey()], err = hashid.DecryptID(p.Hashid, i)
			case []interface{}:
				for k, v := range i {
					i[k], err = hashid.DecryptID(p.Hashid, ztype.ToString(v))
				}
			case []string:
				ids := make([]int64, len(i))
				for k, v := range i {
					ids[k], err = hashid.DecryptID(p.Hashid, v)
				}
				row[define.Inside.IDKey()] = ids
			default:
				row[define.Inside.IDKey()], err = hashid.DecryptID(p.Hashid, ztype.ToString(id))
			}
		}
	}
	return
}

// EnCrypt  加密 ID
func (p *Process) EnCrypt(row *ztype.Map) (err error) {
	if p.Define.Options.CryptID {
		if _, ok := (*row)[define.Inside.IDKey()]; ok {
			(*row)[define.Inside.IDKey()], err = hashid.EncryptID(p.Hashid, (*row).Get(define.Inside.IDKey()).Int64())
		}
	}

	return
}

// EnCryptID  加密 ID
func (p *Process) EnCryptID(id string) (nid string, err error) {
	i := ztype.ToInt64(id)
	if p.Define.Options.CryptID && id != "" {
		if i == 0 {
			return "", errors.New("id cannot be empty")
		}
		nid, err = hashid.EncryptID(p.Hashid, i)
	} else {
		return id, nil
	}

	return
}

// DeCryptID 解密 ID
func (p *Process) DeCryptID(nid string) (id string, err error) {
	if p.Define.Options.CryptID && nid != "" {
		rid, err := hashid.DecryptID(p.Hashid, ztype.ToString(nid))
		if err != nil {
			return "", err
		}
		id = ztype.ToString(rid)
	} else {
		id = nid
	}

	return
}
