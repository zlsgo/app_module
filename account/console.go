package account

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model"
)

// 控制台只读视图（工单 06）。
//
// 控制台是服务端渲染的页面，需要在渲染前拿到「控制台用户 + 角色」的结构化
// 数据；复用 HTTP 接口会让页面渲染绕一圈回自己的服务。这里只开两个读口子，
// 写入仍然走 account 的 JSON 接口（Inside.CreateUser / UpdateUser / DeleteUser），
// 不在模块外重写一套用户写逻辑。

// errConsoleModelMissing 表示模块尚未完成 Load（模型未装配）。
var errConsoleModelMissing = errors.New("account: model not initialized")

// isNoRecord 把「空结果」与真错误区分开：空表应该渲染成空列表，不是 500。
func isNoRecord(err error) bool {
	return errors.Is(err, model.ErrNoRecord)
}

// ConsoleUser 控制台用户列表的一行。
//
// UID 是加密后的对外 ID（与 JSON 接口、绑定接口里用的是同一个值），
// 控制台拿它拼表单地址；密码与盐不出现在视图里。
type ConsoleUser struct {
	UID           string
	Account       string
	Nickname      string
	Status        int
	Roles         []string
	Administrator bool
	Inlay         bool
}

// ConsoleRole 角色下拉项。
type ConsoleRole struct {
	Alias  string
	Label  string
	Status int
	Inlay  bool
}

// ConsoleUsers 按账号/昵称关键字分页读控制台用户。
//
// keyword 为空时取全部；匹配方式与 account 自己的列表接口一致（前缀匹配），
// 避免两处「搜同一个账号却得到不同结果」。
func (m *Module) ConsoleUsers(keyword string, page, pageSize int) ([]ConsoleUser, int, error) {
	if m == nil || m.accountModel == nil {
		return nil, 0, errConsoleModelMissing
	}
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 200 {
		pageSize = 50
	}

	filter := model.Filter{}
	if kw := strings.TrimSpace(keyword); kw != "" {
		// account 的列表接口是 key + 前缀匹配；这里同时匹配昵称，
		// 因为控制台最常见的找法是「我记得这个人叫什么」。
		filter["account"] = kw + "%"
	}

	data, err := m.accountModel.Pages(page, pageSize, filter, func(co *model.CondOptions) {
		co.OrderBy = []model.OrderByItem{{Field: model.IDKey(), Direction: "ASC"}}
		co.Fields = m.accountModel.m.GetFields("password", "salt")
	})
	if err != nil {
		if isNoRecord(err) {
			return []ConsoleUser{}, 0, nil
		}
		return nil, 0, err
	}

	items := make([]ConsoleUser, 0, len(data.Items))
	for _, row := range data.Items {
		items = append(items, consoleUserFromRow(row))
	}
	return items, int(data.Page.Total), nil
}

// ConsoleUser 按 uid 读单个控制台用户。
//
// uid 是加密后的对外 ID（与 JSON 接口一致）。不存在时返回 model.ErrNoRecord，
// 由调用方决定渲染 404 还是 500；不把「找不到」与「读失败」揉成一个字串。
func (m *Module) ConsoleUser(uid string) (ConsoleUser, error) {
	if m == nil || m.accountModel == nil {
		return ConsoleUser{}, errConsoleModelMissing
	}
	row, err := m.accountModel.FindOneByID(uid, func(co *model.CondOptions) {
		co.Fields = m.accountModel.m.GetFields("password", "salt")
	})
	if err != nil {
		return ConsoleUser{}, err
	}
	return consoleUserFromRow(row), nil
}

// ConsoleRoles 返回角色下拉需要的全部角色（管理员与 SaaS 用户都在其中）。
func (m *Module) ConsoleRoles() ([]ConsoleRole, error) {
	if m == nil || m.mods == nil {
		return nil, errConsoleModelMissing
	}
	schema, ok := m.mods.Get(roleName)
	if !ok || schema == nil {
		return nil, errConsoleModelMissing
	}

	data, err := schema.Model().Pages(1, 200, model.Filter{}, func(co *model.CondOptions) {
		co.OrderBy = []model.OrderByItem{{Field: model.IDKey(), Direction: "ASC"}}
	})
	if err != nil {
		if isNoRecord(err) {
			return []ConsoleRole{}, nil
		}
		return nil, err
	}

	roles := make([]ConsoleRole, 0, len(data.Items))
	for _, row := range data.Items {
		roles = append(roles, ConsoleRole{
			Alias:  row.Get("alias").String(),
			Label:  row.Get("label").String(),
			Status: row.Get("status").Int(),
			Inlay:  row.Get("inlay").Bool(),
		})
	}
	return roles, nil
}

// consoleUserFromRow 把模型行转成视图；零值字段保持原样，不在这里编造。
func consoleUserFromRow(row ztype.Map) ConsoleUser {
	return ConsoleUser{
		UID:           row.Get(model.IDKey()).String(),
		Account:       row.Get("account").String(),
		Nickname:      row.Get("nickname").String(),
		Status:        row.Get("status").Int(),
		Roles:         row.Get("role").SliceString(),
		Administrator: row.Get("administrator").Bool(),
		Inlay:         row.Get("inlay").Bool(),
	}
}
