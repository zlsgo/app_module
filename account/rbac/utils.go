package rbac

import (
	"github.com/sohaha/zlsgo/zstring"
)

type Matcher func(target string) (bool, error)

func globMatch(target string) Matcher {
	return func(t string) (bool, error) {
		return zstring.Match(t, target), nil
	}
}

func newGlobPermission(action, target string, deny ...bool) Permission {
	isDeny := false
	if len(deny) > 0 {
		isDeny = deny[0]
	}
	return newPermission(globMatch(action), globMatch(target), isDeny)
}

func newRegexPermission(action, target string, deny ...bool) Permission {
	isDeny := false
	if len(deny) > 0 {
		isDeny = deny[0]
	}
	return newPermission(globMatch(action), func(t string) (bool, error) {
		return zstring.RegexMatch(target, t), nil
	}, isDeny)
}
