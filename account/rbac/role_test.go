package rbac

import (
	"testing"

	"github.com/sohaha/zlsgo"
)

func testRole(matchMode MatchMode) *Role {
	role := NewRole(matchMode)
	role.AddGlobPermission(1, "GET", "/api/test/1/get")
	role.AddGlobPermission(2, "GET", "/api/test/2/get")
	role.AddGlobPermission(2, "GET", "/api/test/2/get/1")
	role.AddGlobPermission(2, "GET", "/api/test/3/2", true)
	role.AddGlobPermission(3, "GET", "/api/test/2/*", true)
	role.AddGlobPermission(3, "GET", "/api/test/2/2")
	role.AddGlobPermission(3, "GET", "/api/test/3/*")
	role.AddGlobPermission(3, "GET", "/api/test")
	role.AddGlobPermission(3, "GET", "/api/test/3/3", true)
	role.AddGlobPermission(3, "GET", "/api/test/3/3", true)
	role.AddRegexPermission(3, "*", "/api/1[\\d]{10}$")
	return role
}

func TestRolePriorityDeny(t *testing.T) {
	tt := zlsgo.NewTest(t)
	role := testRole(MatchPriorityDeny)

	for b, strings := range map[bool][]string{
		true:  {"/api/test/1/get", "/api/test/3/get", "/api/test/3/2", "/api/11234567890"},
		false: {"/api/test/2/get", "/api/test/2/get/1", "/api/test/2/2", "/api/test/3/3", "/api/test/4/get"},
	} {
		for _, s := range strings {
			tt.Run(":"+s, func(tt *zlsgo.TestUtil) {
				ok, err := role.Can("GET", s)
				tt.NoError(err)
				tt.Equal(b, ok)
			})
		}
	}

	ok, err := role.Can("POST", "/api/test/1/get")
	tt.NoError(err)
	tt.Equal(false, ok)

	ok, err = role.Can("POST", "/api/10123456789")
	tt.NoError(err)
	tt.Equal(true, ok)

	ok, err = role.Can("GET", "/api/test/1/get")
	tt.NoError(err)
	tt.Equal(true, ok)
}

func TestRoleMatchSomeDeny(t *testing.T) {
	tt := zlsgo.NewTest(t)
	role := testRole(MatchSomeDeny)

	for b, strings := range map[bool][]string{
		true:  {"/api/test/1/get", "/api/test/3/get"},
		false: {"/api/test/2/get", "/api/test/2/get/1", "/api/test/2/2", "/api/test/3/3", "/api/test/3/2", "/api/test/4/get"},
	} {
		for _, s := range strings {
			tt.Run(":"+s, func(tt *zlsgo.TestUtil) {
				ok, err := role.Can("GET", s)
				tt.NoError(err)
				tt.Equal(b, ok)
			})
		}
	}
}

func TestRoleMatchPrioritySomeAllow(t *testing.T) {
	tt := zlsgo.NewTest(t)
	role := testRole(MatchPrioritySomeAllow)

	for b, strings := range map[bool][]string{
		true:  {"/api/test/1/get", "/api/test/3/2", "/api/test/2/2", "/api/test/3/3", "/api/test/2/get/1", "/api/test/2/get", "/api/test/3/get"},
		false: {"/api/test/4/get"},
	} {
		for _, s := range strings {
			tt.Run(":"+s, func(tt *zlsgo.TestUtil) {
				ok, err := role.Can("GET", s)
				tt.NoError(err)
				tt.Equal(b, ok)
			})
		}
	}
}

func TestRoleMatchSomeAllow(t *testing.T) {
	tt := zlsgo.NewTest(t)
	role := testRole(MatchSomeAllow)

	for b, strings := range map[bool][]string{
		true:  {"/api/test/1/get", "/api/test/2/2", "/api/test/3/2", "/api/test/3/3", "/api/test/2/get/1", "/api/test/2/get", "/api/test/3/get"},
		false: {"/api/test/4/get"},
	} {
		for _, s := range strings {
			tt.Run(":"+s, func(tt *zlsgo.TestUtil) {
				ok, err := role.Can("GET", s)
				tt.NoError(err)
				tt.Equal(b, ok)
			})
		}
	}
}

func TestSubRole(t *testing.T) {
	tt := zlsgo.NewTest(t)

	subRole := testRole(MatchSomeAllow)

	role := NewRole()
	role.AddRole(subRole)

	// circular reference
	subRole.AddRole(role)

	for b, strings := range map[bool][]string{
		true:  {"/api/test/1/get", "/api/test/2/2", "/api/test/3/2", "/api/test/3/3", "/api/test/2/get/1", "/api/test/2/get", "/api/test/3/get"},
		false: {"/api/test/4/get"},
	} {
		for _, s := range strings {
			tt.Run(":"+s, func(tt *zlsgo.TestUtil) {
				ok, err := role.Can("GET", s)
				tt.NoError(err)
				tt.Equal(b, ok)
			})
		}
	}
}

func BenchmarkRole(b *testing.B) {
	tt := zlsgo.NewTest(b)
	r := testRole(MatchPriorityDeny)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ok, err := r.Can("GET", "/api/test/3/2")
			tt.NoError(err)
			tt.EqualTrue(ok)

			ok, err = r.Can("GET", "/api/test/3/3")
			tt.NoError(err)
			tt.EqualTrue(!ok)
		}
	})
}
