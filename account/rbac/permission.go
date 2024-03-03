package rbac

import (
	"sort"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/zsync"
	"github.com/sohaha/zlsgo/zutil"
)

type Permission func(action string, target string) (ok, isDeny bool, err error)

type permissions struct {
	permission map[int][]Permission
	priority   []int
}

func (p *permissions) add(priority int, permission Permission) {
	permissions, ok := p.permission[priority]
	if !ok {
		permissions = make([]Permission, 0)
		p.priority = append(p.priority, priority)
		sort.Sort(sort.Reverse(sort.IntSlice(p.priority)))
	}
	p.permission[priority] = append(permissions, permission)
}

func (p *permissions) exist(matchMode MatchMode, action string, target string) (ok bool, err error) {
	deny, allow := int32(0), int32(0)

	allowSum := zutil.NewInt32(0)
	denySum := zutil.NewInt32(0)

	isPriority := matchMode == MatchPriorityDeny || matchMode == MatchPrioritySomeAllow
	for _, k := range p.priority {
		permissions := p.permission[k]
		if isPriority {
			var isAllow bool
			for i := range permissions {
				exist, isDeny, err := permissions[i](action, target)
				if err != nil {
					return false, nil
				}

				if !exist {
					continue
				}

				if isDeny && matchMode == MatchPriorityDeny {
					return false, nil
				}

				isAllow = !isDeny
				if isAllow && matchMode == MatchPrioritySomeAllow {
					return true, err
				}
			}

			if isAllow {
				return true, nil
			}
		}

		var wg zsync.WaitGroup
		for i := range permissions {
			i := i
			wg.GoTry(func() {
				ok, isDeny, err := permissions[i](action, target)
				zerror.Panic(err)
				if ok {
					if isDeny {
						denySum.Add(1)
					} else {
						allowSum.Add(1)
					}
				}
			})
		}
		err = wg.Wait()

		if err != nil {
			return false, err
		}

		callow := allowSum.Load()
		cdeny := denySum.Load()

		allow += callow
		deny += cdeny

	}

	if deny > 0 && matchMode == MatchSomeDeny {
		return false, err
	}

	if allow > 0 && matchMode == MatchSomeAllow {
		return true, err
	}
	return allow > 0, nil
}

func newPermission(actionMatcher, targetMatcher Matcher, deny bool) Permission {
	return func(action string, target string) (bool, bool, error) {
		actionMatch, err := actionMatcher(action)
		if err != nil {
			return false, false, err
		}

		if !actionMatch {
			return false, false, nil
		}

		ok, err := targetMatcher(target)
		return ok, deny, err
	}
}

func AllowAll(action, target string) (bool, error) {
	return true, nil
}
