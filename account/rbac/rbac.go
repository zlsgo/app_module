package rbac

import (
	"errors"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/conf"
)

type RBAC struct {
	roles *zarray.Maper[string, *Role]
}

func New() *RBAC {
	return &RBAC{
		roles: zarray.NewHashMap[string, *Role](),
	}
}

func ParseConfig(path string) (*RBAC, error) {
	c := conf.New(zfile.RealPath(path))
	err := c.Read()
	if err != nil {
		return nil, err
	}

	r := &RBAC{
		roles: zarray.NewHashMap[string, *Role](),
	}

	ztype.ToMap(c.GetAll()).ForEach(func(key string, value ztype.Type) bool {
		m := value.Get("mode").Uint()
		if m > uint(MatchSomeAllow) {
			m = uint(MatchPriorityDeny)
		}

		role := NewRole(MatchMode(m))

		value.Get("permission").Slice().Maps().ForEach(func(_ int, value ztype.Map) bool {
			target := value.Get("target").String()
			if target == "" {
				err = errors.New("target is empty")
				return false
			}

			role.AddGlobPermission(
				value.Get("priority").Int(0),
				value.Get("action").String("*"),
				target,
				value.Get("deny").Bool(),
			)
			return true
		})

		if err != nil {
			return false
		}

		err = r.AddRole(key, role)
		return err == nil
	})

	return r, err
}

func (r *RBAC) AddRole(roleName string, role *Role) error {
	if r.roles.Has(roleName) {
		return errors.New("role already exists")
	}

	r.roles.Set(roleName, role)
	return nil
}

func (r *RBAC) RemoveRole(roleName string) error {
	r.roles.Delete(roleName)
	return nil
}

func (r *RBAC) Can(roleName string, action, target string) (ok bool, err error) {
	role, ok := r.roles.Get(roleName)
	if !ok {
		return false, nil
	}

	return role.Can(action, target)
}
