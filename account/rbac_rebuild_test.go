package account

import (
	"path/filepath"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

func newTestModule(t *testing.T, opt ...func(o *Options)) (*Module, func()) {
	t.Helper()

	db, err := zdb.New(&sqlite3.Config{File: ":memory:", Memory: true, Parameters: "_pragma=busy_timeout(3000)"})
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}

	mod := New("test-key")
	mod.Options.AdminDefaultPassword = "Aa123456"
	mod.Options.ApiPrefix = "/api"
	for _, fn := range opt {
		fn(&mod.Options)
	}
	mod.Options.InitDB = func() (*zdb.DB, error) {
		return db, nil
	}

	if err := mod.Start(zdi.New()); err != nil {
		_ = db.Close()
		t.Fatalf("start module: %v", err)
	}

	return mod, func() {
		_ = db.Close()
	}
}

func TestRebuildRBACPreservesFileRules(t *testing.T) {
	tt := zlsgo.NewTest(t)
	mod, cleanup := newTestModule(t, func(o *Options) {
		o.RBACFile = filepath.Join("rbac", "testdata", "rbac.toml")
	})
	defer cleanup()

	perm := mod.permission.Load()
	tt.NotNil(perm)
	ok, err := perm.Can("admin", "GET", "/api/test1")
	tt.NoError(err)
	tt.Equal(false, ok)

	tt.NoError(mod.rebuildRBAC())

	perm = mod.permission.Load()
	tt.NotNil(perm)
	ok, err = perm.Can("admin", "GET", "/api/test1")
	tt.NoError(err)
	tt.Equal(false, ok)
}

func TestRebuildRBACLoadsDBRoles(t *testing.T) {
	tt := zlsgo.NewTest(t)
	mod, cleanup := newTestModule(t)
	defer cleanup()

	permID, err := mod.index.permModel.Model().Insert(ztype.Map{
		"label":    "Test Perm",
		"alias":    "perm_test",
		"status":   1,
		"action":   "GET",
		"target":   "/api/custom",
		"priority": 10,
	})
	tt.NoError(err)

	_, err = mod.index.roleModel.Model().Insert(ztype.Map{
		"label":      "Tester",
		"alias":      "tester",
		"status":     1,
		"permission": []int{ztype.ToInt(permID)},
	})
	tt.NoError(err)

	tt.NoError(mod.rebuildRBAC())

	perm := mod.permission.Load()
	tt.NotNil(perm)
	ok, err := perm.Can("tester", "GET", "/api/custom")
	tt.NoError(err)
	tt.Equal(true, ok)
}
