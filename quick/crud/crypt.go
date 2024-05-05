package crud

import (
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/database/hashid"
)

type crypt struct {
}

var Crypt = &crypt{}

func (c *crypt) ID(q *Crud, nid string) (id string, err error) {
	if q.define.Options.CryptID && nid != "" {
		rid, err := hashid.DecryptID(q.process.Hashid, ztype.ToString(nid))
		if err != nil {
			return "", err
		}
		id = ztype.ToString(rid)
	} else {
		id = nid
	}

	return
}

func (c *crypt) CryptID(q *Crud, id string) (nid string, err error) {
	if q.define.Options.CryptID && id != "" {
		return q.process.EnCryptID(id)
	}
	id = nid

	return
}
