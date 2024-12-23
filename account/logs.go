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
	m, _ := h.module.mods.Get(logsName)
	user := h.module.Request.User(c)
	account := user.Get("account").String()
	if account == "" {
		return nil, errors.New("用户不存在")
	}

	page, pagesize, _ := common.VarPages(c)
	return model.Pages(m, page, pagesize, ztype.Map{
		"account": account,
	}, func(co *model.CondOptions) {
		co.OrderBy = map[string]string{model.IDKey(): "desc"}
	})
}

// 记录日志
func logRequest(c *znet.Context, m *model.Schema, u ztype.Map) {
	msg, ok := c.Value(ctxWithLog)
	if !ok {
		return
	}

	var remark string

	if r, ok := c.Value(ctxWithLogRemark); ok {
		remark = ztype.ToString(r)
	}

	_, _ = insertLog(c, m, u.Get("account").String(), c.PrevContent().Code.Load(), msg.(string), remark)
}

func insertLog(c *znet.Context, m *model.Schema, account string, status int32, msg string, remark ...string) (interface{}, error) {
	var r string
	if len(remark) > 0 {
		r = remark[0]
	}

	ip := ""
	if !noLogIP {
		ip = c.GetClientIP()
	}

	return model.Insert(m, ztype.Map{
		"account":   account,
		"ip":        ip,
		"method":    c.Request.Method,
		"path":      c.Request.URL.String(),
		"status":    status,
		"message":   msg,
		"remark":    r,
		"params":    c.Request.URL.Query().Encode(),
		"record_at": ztime.Now(),
	})
}
