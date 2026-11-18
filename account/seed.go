package account

import (
	"errors"

	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model"
)

// PermissionRule 宿主侧的种子权限规则（glob 形式，供 EnsurePermissionRules）。
//
// target/action 里 `*` 匹配任意字符序列（含 /），与 RBAC 的 glob 行为一致；
// alias 必须满足 ^[a-zA-Z0-9_]+$ 且 ≤20 字符（permission 表的字段约束）。
type PermissionRule struct {
	Alias    string
	Label    string
	Action   string
	Target   string
	Remark   string
	Priority uint32
}

// EnsurePermissionRules 幂等写入权限规则：按 alias 查重，已存在的行不改动
// （部署方在管理界面人工调整过的规则不被种子覆盖）。返回与 rules 一一对应的
// permission 行 ID。
//
// 必须在模块 Start 之后调用（此时模型才装配完成）。写入后会重建 RBAC，
// 中间件立即生效。
func (m *Module) EnsurePermissionRules(rules []PermissionRule) ([]int, error) {
	if m == nil || m.index == nil || m.index.permModel == nil {
		return nil, errors.New("account: permission model is not initialized (call after Start)")
	}
	ids := make([]int, 0, len(rules))
	for _, r := range rules {
		if r.Alias == "" || r.Action == "" || r.Target == "" {
			return nil, errors.New("account: permission rule requires alias/action/target")
		}
		rows, err := m.index.permModel.Model().Find(model.Filter{"alias": r.Alias})
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			ids = append(ids, rows[0].Get(model.IDKey()).Int())
			continue
		}
		id, err := m.index.permModel.Model().Insert(ztype.Map{
			"label":    r.Label,
			"alias":    r.Alias,
			"status":   1,
			"inlay":    true,
			"remark":   r.Remark,
			"action":   r.Action,
			"target":   r.Target,
			"priority": r.Priority,
		})
		if err != nil {
			return nil, err
		}
		ids = append(ids, ztype.ToInt(id))
	}
	return ids, nil
}

// EnsureRole 幂等创建角色：alias 已存在时原样返回，不改动它的权限清单
// （人工调整不被覆盖）。permissionIDs 是 permission 表的行 ID。
// 写入后重建 RBAC 与角色缓存。
func (m *Module) EnsureRole(alias, label string, permissionIDs []int) error {
	if m == nil || m.index == nil || m.index.roleModel == nil {
		return errors.New("account: role model is not initialized (call after Start)")
	}
	if alias == "" {
		return errors.New("account: role alias is required")
	}
	rows, err := m.index.roleModel.Model().Find(model.Filter{"alias": alias})
	if err != nil {
		return err
	}
	if len(rows) > 0 {
		return nil
	}
	if _, err := m.index.roleModel.Model().Insert(ztype.Map{
		"label":      label,
		"alias":      alias,
		"status":     1,
		"permission": permissionIDs,
	}); err != nil {
		return err
	}
	return m.rebuildRBAC()
}
