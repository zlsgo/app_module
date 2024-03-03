package rbac

import (
	"errors"
	"unsafe"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zcache"
)

type MatchMode uint

const (
	// MatchPriorityDeny 按优先匹配，只要有封禁就禁止通过
	MatchPriorityDeny MatchMode = iota
	// MatchSomeDeny 只要有封禁就禁止通过
	MatchSomeDeny
	// MatchPrioritySomeAllow 按优先级匹配，只要有权限就允许通过
	MatchPrioritySomeAllow
	// MatchSomeAllow 只要有权限就可以通过
	MatchSomeAllow
)

type Role struct {
	cache       *zcache.FastCache
	permissions permissions
	roles       []*Role
	matchMode   MatchMode
	maxStack    uint
}

func NewRole(matchMode ...MatchMode) *Role {
	m := MatchPriorityDeny
	if len(matchMode) > 0 {
		m = matchMode[0]
	}
	return &Role{
		matchMode:   m,
		maxStack:    10,
		permissions: permissions{permission: make(map[int][]Permission), priority: make([]int, 0)},
		cache:       zcache.NewFast(),
	}
}

func (r *Role) AddGlobPermission(priority int, action, target string, deny ...bool) {
	r.AddPermission(priority, newGlobPermission(action, target, deny...))
}

func (r *Role) AddRegexPermission(priority int, action, target string, deny ...bool) {
	r.AddPermission(priority, newRegexPermission(action, target, deny...))
}

func (r *Role) AddPermission(priority int, p Permission) *Role {
	r.permissions.add(priority, p)
	return r
}

func (r *Role) AddRole(role *Role) *Role {
	r.roles = append(r.roles, role)
	return r
}

func (r *Role) Can(action, target string) (ok bool, err error) {
	cacheKey := action + ":" + target

	if val, ok := r.cache.Get(cacheKey); ok {
		return val.(bool), nil
	}

	var rolesPointer []unsafe.Pointer
	ok, rolesPointer, err = r.can(action, target, []unsafe.Pointer{unsafe.Pointer(r)}, 0)
	if err != nil {
		return
	}

	if len(rolesPointer) == 0 {
		r.cache.Set(cacheKey, ok)
	}

	return ok, nil
}

func (r *Role) can(action, target string, rp []unsafe.Pointer, stack uint) (ok bool, rolesPointer []unsafe.Pointer, err error) {
	if stack >= r.maxStack {
		return false, nil, errors.New("max stack")
	}

	ok, err = r.permissions.exist(r.matchMode, action, target)
	if err != nil {
		return false, nil, err
	}

	if ok {
		return
	}

	rolesPointer = make([]unsafe.Pointer, 0, len(r.roles)+len(rp))

	if len(r.roles) > 0 {
		if rp != nil {
			rolesPointer = append(rolesPointer, rp...)
		}

		rolesPointer = zarray.Unique(rolesPointer)
		for i := range r.roles {
			var rp []unsafe.Pointer
			p := unsafe.Pointer(r.roles[i])
			if zarray.Contains(rolesPointer, p) {
				continue
			}

			ok, rp, err = r.roles[i].can(action, target, rolesPointer, stack+1)
			rolesPointer = append(rolesPointer, p)
			rolesPointer = append(rolesPointer, rp...)
			if err != nil {
				return false, rolesPointer, err
			}

			if ok {
				return true, rolesPointer, nil
			}
		}
	}

	return false, rolesPointer, nil
}
