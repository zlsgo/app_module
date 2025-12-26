package hashid

import (
	"errors"

	"github.com/sohaha/zlsgo/zarray"

	"github.com/speps/go-hashids/v2"
)

// HashID 封装 hashids.HashID
type HashID struct {
	hash *hashids.HashID
	err  error
}

// h 全局实例
var h = zarray.NewHashMap[string, *HashID]()

// Set 注册命名实例
func Set(name, salt string, minLength ...int) error {
	if h.Has(name) {
		return errors.New("repeat the initialization of the hashid")
	}

	l := 6
	if len(minLength) > 0 {
		l = minLength[0]
	}

	hash := New(salt, l)
	if hash == nil || hash.hash == nil {
		if hash != nil && hash.err != nil {
			return hash.err
		}
		return errors.New("hashid init failed")
	}
	h.Set(name, hash)

	return nil
}

// Get 获取命名实例
func Get(name string) (*HashID, bool) {
	hashid, ok := h.Get(name)
	if !ok {
		return nil, false
	}
	return hashid, true
}

// New 创建实例
func New(salt string, MinLength int) *HashID {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = MinLength
	hash, err := hashids.NewWithData(hd)
	if err != nil {
		return &HashID{err: err}
	}
	return &HashID{hash: hash}
}

// EncryptID 编码 ID
func EncryptID(hashI *HashID, id int64) (string, error) {
	if hashI == nil || hashI.hash == nil {
		if hashI != nil && hashI.err != nil {
			return "", hashI.err
		}
		return "", errors.New("hashid is nil")
	}
	return hashI.hash.EncodeInt64([]int64{id})
}

// DecryptID 解码 ID
func DecryptID(hashI *HashID, hid string) (int64, error) {
	if hashI == nil || hashI.hash == nil {
		if hashI != nil && hashI.err != nil {
			return 0, hashI.err
		}
		return 0, errors.New("hashid is nil")
	}
	ids, err := hashI.hash.DecodeInt64WithError(hid)
	if len(ids) == 0 || err != nil {
		return 0, err
	}
	return ids[0], nil
}
