package account

import (
	"errors"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/common"
	"github.com/zlsgo/app_module/model"
)

var noLogIP = false

// GetLogs 操作日志
func (h *Index) GetLogs(c *znet.Context) (data any, err error) {
	m, ok := h.module.mods.Get(logsName)
	if !ok || m == nil {
		return nil, errors.New("日志模型未找到")
	}

	user := h.module.Request.User(c)
	account := user.Get("account").String()
	if account == "" {
		return nil, errors.New("用户不存在")
	}

	page, pagesize, err := common.VarPages(c)
	if err != nil {
		return nil, err
	}
	return model.Pages(m, page, pagesize, ztype.Map{
		"account": account,
	}, func(co *model.CondOptions) {
		co.OrderBy = []model.OrderByItem{{Field: model.IDKey(), Direction: "DESC"}}
	})
}

// 记录日志
func logRequest(c *znet.Context, m *model.Schema, u ztype.Map) {
	msg, ok := c.Value(ctxWithLog)
	if !ok {
		return
	}

	account := u.Get("account").String()
	status := c.PrevContent().Code.Load()
	msgStr, _ := msg.(string)

	var remark string
	if r, ok := c.Value(ctxWithLogRemark); ok {
		remark = ztype.ToString(r)
	}

	ip := ""
	if !noLogIP {
		ip = c.GetClientIP()
	}
	method := c.Request.Method
	path := c.Request.URL.String()
	params := c.Request.URL.Query().Encode()

	go func() {
		_, _ = insertLog(m, account, ip, method, path, status, msgStr, params, remark)
	}()
}

func insertLog(m *model.Schema, account, ip, method, path string, status int32, msg, params, remark string) (interface{}, error) {
	return model.Insert(m, ztype.Map{
		"account":   account,
		"ip":        ip,
		"method":    method,
		"path":      path,
		"status":    status,
		"message":   msg,
		"remark":    remark,
		"params":    params,
		"record_at": ztime.Now(),
	})
}
