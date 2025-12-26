package model

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"io"
	"strings"

	"github.com/zlsgo/app_module/database/hashid"
	"github.com/zlsgo/zdb"

	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
	"golang.org/x/crypto/bcrypt"
)

type CryptProcess func(string) (string, error)

type IDCrypter interface {
	Encrypt(id int64) (string, error)
	Decrypt(encrypted string) (int64, error)
}

type HashIDCrypter struct {
	hashid *hashid.HashID
}

func NewHashIDCrypter(salt string, minLength int) *HashIDCrypter {
	return &HashIDCrypter{
		hashid: hashid.New(salt, minLength),
	}
}

func (c *HashIDCrypter) Encrypt(id int64) (string, error) {
	return hashid.EncryptID(c.hashid, id)
}

func (c *HashIDCrypter) Decrypt(encrypted string) (int64, error) {
	return hashid.DecryptID(c.hashid, encrypted)
}

type AESCrypter struct {
	block cipher.Block
}

func NewAESCrypter(key []byte) (*AESCrypter, error) {
	keyLen := len(key)
	if keyLen != 16 && keyLen != 24 && keyLen != 32 {
		return nil, ErrInvalidKeyLength
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return &AESCrypter{block: block}, nil
}

func (c *AESCrypter) Encrypt(id int64) (string, error) {
	plaintext := make([]byte, 8)
	binary.BigEndian.PutUint64(plaintext, uint64(id))

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(c.block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.RawURLEncoding.EncodeToString(ciphertext), nil
}

func (c *AESCrypter) Decrypt(encrypted string) (int64, error) {
	ciphertext, err := base64.RawURLEncoding.DecodeString(encrypted)
	if err != nil {
		return 0, ErrInvalidCryptedID
	}

	if len(ciphertext) < aes.BlockSize+8 {
		return 0, ErrInvalidCryptedID
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(c.block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return int64(binary.BigEndian.Uint64(ciphertext)), nil
}

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

func (m *Schema) getIDCrypter() IDCrypter {
	if m.idCrypter != nil {
		return m.idCrypter
	}
	if m.Hashid != nil {
		return &HashIDCrypter{hashid: m.Hashid}
	}
	return nil
}

func (m *Schema) SetIDCrypter(crypter IDCrypter) {
	m.idCrypter = crypter
}

func (m *Schema) DeCrypt(row ztype.Map) (success bool) {
	success = true
	if !*m.define.Options.CryptID {
		return
	}

	crypter := m.getIDCrypter()
	if crypter == nil {
		return false
	}

	return decryptFilterMap(row, crypter)
}

func (m *Schema) EnCrypt(row *ztype.Map) (err error) {
	if !*m.define.Options.CryptID {
		return nil
	}

	crypter := m.getIDCrypter()
	if crypter == nil {
		return ErrCrypterNotSet
	}

	if _, ok := (*row)[idKey]; ok {
		(*row)[idKey], err = crypter.Encrypt((*row).Get(idKey).Int64())
	}

	return
}

func (m *Schema) EnCryptID(id string) (nid string, err error) {
	i := ztype.ToInt64(id)
	if !*m.define.Options.CryptID || id == "" {
		return id, nil
	}

	if i == 0 {
		return "", errors.New("id cannot be empty")
	}

	crypter := m.getIDCrypter()
	if crypter == nil {
		return "", ErrCrypterNotSet
	}

	return crypter.Encrypt(i)
}

func (m *Schema) DeCryptID(nid string) (id string, err error) {
	if !*m.define.Options.CryptID || nid == "" {
		return nid, nil
	}

	crypter := m.getIDCrypter()
	if crypter == nil {
		return "", ErrCrypterNotSet
	}

	rid, err := crypter.Decrypt(nid)
	if err != nil {
		return "", err
	}
	return ztype.ToString(rid), nil
}

func decryptFilterMap(row ztype.Map, crypter IDCrypter) bool {
	if len(row) == 0 {
		return true
	}

	success := true
	for k, v := range row {
		if k == "" {
			continue
		}
		trimmedKey := strings.TrimSpace(k)
		if trimmedKey == "" {
			continue
		}

		upperKey := strings.ToUpper(trimmedKey)
		if upperKey == placeHolderOR || upperKey == placeHolderAND {
			if !decryptNestedFilters(v, crypter) {
				success = false
			}
			continue
		}
		if strings.Contains(trimmedKey, placeHolder) {
			continue
		}

		fieldName := trimmedKey
		op := ""
		if spaceIdx := strings.IndexAny(fieldName, " \t"); spaceIdx > 0 {
			op = strings.TrimSpace(fieldName[spaceIdx+1:])
			fieldName = fieldName[:spaceIdx]
		}
		if strings.Contains(fieldName, ".") {
			continue
		}
		if fieldName != idKey {
			continue
		}
		if strings.ToUpper(op) == "IS NULL" || strings.ToUpper(op) == "IS NOT NULL" || v == nil {
			continue
		}

		nv, ok := decryptIDValue(v, crypter)
		if !ok {
			success = false
			continue
		}
		row[k] = nv
	}

	return success
}

func decryptNestedFilters(value any, crypter IDCrypter) bool {
	switch v := value.(type) {
	case ztype.Map:
		return decryptFilterMap(v, crypter)
	case map[string]interface{}:
		return decryptFilterMap(ztype.Map(v), crypter)
	case ztype.Maps:
		ok := true
		for i := range v {
			if !decryptFilterMap(v[i], crypter) {
				ok = false
			}
		}
		return ok
	case []ztype.Map:
		ok := true
		for i := range v {
			if !decryptFilterMap(v[i], crypter) {
				ok = false
			}
		}
		return ok
	case []map[string]interface{}:
		ok := true
		for i := range v {
			if !decryptFilterMap(ztype.Map(v[i]), crypter) {
				ok = false
			}
		}
		return ok
	case []interface{}:
		ok := true
		for i := range v {
			switch vv := v[i].(type) {
			case ztype.Map:
				if !decryptFilterMap(vv, crypter) {
					ok = false
				}
			case map[string]interface{}:
				if !decryptFilterMap(ztype.Map(vv), crypter) {
					ok = false
				}
			}
		}
		return ok
	default:
		m := ztype.New(value).Map()
		if len(m) == 0 {
			return true
		}
		return decryptFilterMap(m, crypter)
	}
}

func decryptIDValue(value any, crypter IDCrypter) (any, bool) {
	switch v := value.(type) {
	case string:
		raw, err := crypter.Decrypt(v)
		if err != nil {
			return value, false
		}
		return raw, true
	case []interface{}:
		ok := true
		for i := range v {
			raw, err := crypter.Decrypt(ztype.ToString(v[i]))
			if err != nil {
				ok = false
				continue
			}
			v[i] = raw
		}
		return v, ok
	case []string:
		ok := true
		ids := make([]int64, len(v))
		for i := range v {
			raw, err := crypter.Decrypt(v[i])
			if err != nil {
				ok = false
				continue
			}
			ids[i] = raw
		}
		return ids, ok
	default:
		raw, err := crypter.Decrypt(ztype.ToString(value))
		if err != nil {
			return value, false
		}
		return raw, true
	}
}

func GetEngine[T *zdb.DB](m *Store) T {
	return m.schema.Storage.(*SQL).GetDB()
}
