package account

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/zlsgo/app_module/model"
)

func TestEnsurePermissionRulesIdempotent(t *testing.T) {
	tt := zlsgo.NewTest(t)
	mod, cleanup := newTestModule(t)
	defer cleanup()

	rules := []PermissionRule{
		{Alias: "seed_read", Label: "读种子数据", Action: "GET", Target: "/api/seed/*"},
		{Alias: "seed_write", Label: "写种子数据", Action: "POST", Target: "/api/seed/*"},
	}

	ids, err := mod.EnsurePermissionRules(rules)
	tt.NoError(err)
	tt.Equal(2, len(ids))
	tt.Equal(true, ids[0] > 0)
	tt.Equal(true, ids[1] > 0)

	// 重复调用不应新增行，返回同样的 ID
	again, err := mod.EnsurePermissionRules(rules)
	tt.NoError(err)
	tt.Equal(ids, again)

	rows, err := mod.index.permModel.Model().Find(model.Filter{"alias": "seed_read"})
	tt.NoError(err)
	tt.Equal(1, len(rows))
	tt.Equal("/api/seed/*", rows[0].Get("target").String())
	tt.Equal(uint64(1), rows[0].Get("status").Uint64())

	// 缺字段直接报错，不会写入半成品
	_, err = mod.EnsurePermissionRules([]PermissionRule{{Alias: "", Action: "GET", Target: "/api/seed/*"}})
	tt.Equal(true, err != nil)

	// 空规则集：不报错，也不触发重建
	empty, err := mod.EnsurePermissionRules(nil)
	tt.NoError(err)
	tt.Equal(0, len(empty))
}

func TestEnsurePermissionRulesRequiresStartedModule(t *testing.T) {
	tt := zlsgo.NewTest(t)

	var mod *Module
	_, err := mod.EnsurePermissionRules([]PermissionRule{{Alias: "a", Action: "GET", Target: "/a"}})
	tt.Equal(true, err != nil)

	mod = New("test-key")
	_, err = mod.EnsurePermissionRules([]PermissionRule{{Alias: "a", Action: "GET", Target: "/a"}})
	tt.Equal(true, err != nil)
	tt.Equal(true, mod.EnsureRole("a", "b", nil) != nil)
}

func TestEnsureRoleSeedsEffectivePermissions(t *testing.T) {
	tt := zlsgo.NewTest(t)
	mod, cleanup := newTestModule(t)
	defer cleanup()

	ids, err := mod.EnsurePermissionRules([]PermissionRule{
		{Alias: "seed_role_read", Label: "读", Action: "GET", Target: "/api/seed/*"},
	})
	tt.NoError(err)
	tt.Equal(1, len(ids))

	tt.NoError(mod.EnsureRole("seed_role", "种子角色", ids))

	// 种子写入后 RBAC 立即生效，无需宿主再手动重建
	perm := mod.permission.Load()
	tt.NotNil(perm)
	ok, err := perm.Can("seed_role", "GET", "/api/seed/42")
	tt.NoError(err)
	tt.Equal(true, ok)

	ok, err = perm.Can("seed_role", "DELETE", "/api/seed/42")
	tt.NoError(err)
	tt.Equal(false, ok)

	// 已存在的角色保持原样，不被再次调用覆盖
	tt.NoError(mod.EnsureRole("seed_role", "改过的标签", nil))
	rows, err := mod.index.roleModel.Model().Find(model.Filter{"alias": "seed_role"})
	tt.NoError(err)
	tt.Equal(1, len(rows))
	tt.Equal("种子角色", rows[0].Get("label").String())

	// 空别名直接报错
	tt.Equal(true, mod.EnsureRole("", "x", nil) != nil)
}
