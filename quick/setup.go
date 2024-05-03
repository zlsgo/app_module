package quick

import (
	"errors"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/zlsgo/app_module/database/hashid"
	"github.com/zlsgo/app_module/quick/define"
	"github.com/zlsgo/app_module/quick/process"
)

func setup(q *Quick) error {
	if q.define.Name == "" {
		return errors.New("model name is empty")
	}

	q.define.Name = zstring.CamelCaseToSnakeCase(q.define.Name)
	q.tableName = q.options.TablePrefix + q.define.Name
	// q.storage.setup(q)
	return perfect(q)
}

func perfect(q *Quick) (err error) {
	if q.define.Options.CryptID {
		salt := q.define.Options.Salt
		// if salt == "" {
		// 	return errors.New("salt is empty")
		// }
		cryptLen := q.define.Options.CryptLen
		if cryptLen <= 0 {
			cryptLen = 12
		}
		q.process.Hashid = hashid.New(salt, cryptLen)
	}

	q.readOnlyKeys = make([]string, 0)
	q.process.Define = q.define
	q.process.CryptKeys = make(map[string]process.CryptProcess)
	q.process.AfterProcess = make(map[string][]process.AfterProcess)
	q.process.BeforeProcess = make(map[string][]process.BeforeProcess)

	q.Fields, err = perfectField(q)
	if err != nil {
		return
	}

	q.inlayFields = []string{define.Inside.IDKey()}
	if q.define.Options.Timestamps {
		if zarray.Contains(q.Fields, define.Inside.CreatedAtKey()) {
			err = errors.New(define.Inside.CreatedAtKey() + " is a reserved field")
			return
		}
		if zarray.Contains(q.Fields, define.Inside.UpdatedAtKey()) {
			err = errors.New(define.Inside.UpdatedAtKey() + " is a reserved field")
			return
		}
		var after []process.AfterProcess
		after, err = process.GetAfterProcess([]string{"date|Y-m-d H:i:s"})
		if err != nil {
			return
		}
		q.process.AfterProcess[define.Inside.CreatedAtKey()] = after
		q.process.AfterProcess[define.Inside.UpdatedAtKey()] = after
		q.inlayFields = append(q.inlayFields, define.Inside.CreatedAtKey(), define.Inside.UpdatedAtKey())
	}

	if q.define.Options.SoftDeletes {
		if zarray.Contains(q.Fields, define.Inside.DeletedAtKey()) {
			err = errors.New(define.Inside.DeletedAtKey() + " is a reserved field")
			return
		}
		q.inlayFields = append(q.inlayFields, define.Inside.DeletedAtKey())
	}

	q.fullFields = append([]string{define.Inside.IDKey()}, q.Fields...)
	q.fullFields = zarray.Unique(append(q.fullFields, q.inlayFields...))

	if q.define.Options.SoftDeletes {
		flen := len(q.fullFields)
		for i := 0; i < flen; i++ {
			f := q.fullFields[i]
			if f == define.Inside.DeletedAtKey() {
				q.fullFields = append(q.fullFields[0:i], q.fullFields[i+1:]...)
				break
			}
		}
	}

	if len(q.define.Relations) > 0 {
		for k := range q.define.Relations {
			v := q.define.Relations[k]
			if v.Foreign == "" {
				q.define.Relations[k].Foreign = define.Inside.IDKey()
			}
		}

		newRelations := make(map[string]*define.ModelRelation, len(q.define.Relations))
		for k := range q.define.Relations {
			v := q.define.Relations[k]
			newRelations[zstring.CamelCaseToSnakeCase(k)] = v
		}
		q.define.Relations = newRelations
	} else {
		q.define.Relations = make(map[string]*define.ModelRelation)
	}

	return
}
