package rbac

import (
	"path/filepath"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zsync"
)

func TestNew(t *testing.T) {
	tt := zlsgo.NewTest(t)
	admin := testRole(MatchPriorityDeny)

	r := New()
	err := r.AddRole("admin", admin)
	tt.NoError(err)

	newRole := NewRole(MatchSomeDeny)
	newRole.AddGlobPermission(1, "GET", "/web/user")
	r.MergerRole("admin", newRole)

	ok, err := r.Can("admin", "GET", "/web/user")
	tt.NoError(err)
	tt.EqualTrue(ok)

	ok, err = r.Can("admin", "POST", "/web/user")
	tt.NoError(err)
	tt.EqualTrue(!ok)

	ok, err = r.Can("admin", "GET", "/api/test")
	tt.NoError(err)
	tt.EqualTrue(ok)

	ok, err = r.Can("admin", "GET", "/api/test1")
	tt.NoError(err)
	tt.EqualTrue(!ok)

	_ = r.RemoveRole("admin")

	ok, err = r.Can("admin", "GET", "/api/test")
	tt.NoError(err)
	tt.Equal(false, ok)

	ok, err = r.Can("admin2", "GET", "/api/test")
	tt.NoError(err)
	tt.Equal(false, ok)
}

func TestMerge(t *testing.T) {
	tt := zlsgo.NewTest(t)

	admin := testRole(MatchPriorityDeny)
	e := New()
	e.AddRole("admin", admin)

	r := New()
	newRole := NewRole(MatchSomeDeny)
	newRole.AddGlobPermission(1, "GET", "/web/user")
	err := r.AddRole("admin2", newRole)
	tt.NoError(err)

	e.Merge(r)

	ok, err := e.Can("admin2", "GET", "/web/user")
	tt.NoError(err)
	tt.EqualTrue(ok)

	ok, err = e.Can("admin", "GET", "/api/test")
	tt.NoError(err)
	tt.EqualTrue(ok)
}

func TestNewGo(t *testing.T) {
	tt := zlsgo.NewTest(t)
	admin := testRole(MatchPriorityDeny)

	r := New()
	err := r.AddRole("admin", admin)
	tt.NoError(err)

	var wg zsync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Go(func() {
			for i := 0; i < 100; i++ {
				ok, err := r.Can("admin", "GET", "/api/test/3/2")
				tt.NoError(err)
				tt.EqualTrue(ok)

				ok, err = r.Can("admin", "GET", "/api/test/3/3")
				tt.NoError(err)
				tt.EqualTrue(!ok)

				ok, err = r.Can("admin2", "GET", "/api/test")
				tt.NoError(err)
				tt.Equal(false, ok)
			}
		})
	}

	tt.NoError(wg.Wait())
}

func TestParseConfig(t *testing.T) {
	tt := zlsgo.NewTest(t)

	path := filepath.Join(zfile.RootPath(), "testdata", "rbac.toml")
	r, err := ParseFile(path)
	tt.Log(path)
	tt.NoError(err, true)

	ok, err := r.Can("admin", "GET", "/api/test")
	tt.NoError(err)
	tt.EqualTrue(ok)

	ok, err = r.Can("admin", "POST", "/api/test1")
	tt.NoError(err)
	tt.EqualTrue(ok)

	ok, err = r.Can("admin", "GET", "/api/test1")
	tt.NoError(err)
	tt.EqualTrue(!ok)

	ok, err = r.Can("admin2", "GET", "/api/test")
	tt.NoError(err)
	tt.Equal(false, ok)
}
