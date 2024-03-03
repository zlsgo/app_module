package hashid

import (
	"github.com/sohaha/zlsgo/zarray"

	"errors"
	"github.com/speps/go-hashids/v2"
)

type HashID struct {
	hash *hashids.HashID
}

var h = zarray.NewHashMap[string, *HashID]()

func Set(name, salt string, minLength ...int) error {
	if h.Has(name) {
		return errors.New("repeat the initialization of the hashid")
	}

	l := 6
	if len(minLength) > 0 {
		l = minLength[0]
	}

	h.Set(name, New(salt, l))

	return nil
}

func Get(name string) (*HashID, bool) {
	hashid, ok := h.Get(name)
	if !ok {
		return nil, false
	}
	return hashid, true
}

func New(salt string, MinLength int) *HashID {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = MinLength
	hash, _ := hashids.NewWithData(hd)
	return &HashID{hash: hash}
}

func EncryptID(hashI *HashID, id int64) (string, error) {
	return hashI.hash.EncodeInt64([]int64{id})
}

func DecryptID(hashI *HashID, hid string) (int64, error) {
	ids, err := hashI.hash.DecodeInt64WithError(hid)
	if len(ids) == 0 || err != nil {
		return 0, err
	}
	return ids[0], nil
}
