package account

import (
	"errors"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/restapi"
	"github.com/zlsgo/zdb/schema"
)

type MessageModel struct {
	*restapi.Operation
	mod *Module
}

func (m *Module) MessageModel() *MessageModel {
	return m.messageModel
}

func messageModelDefine(m *Module) error {
	const messageName = "message"
	mod, err := m.mods.Reg(messageName, restapi.Define{
		Name: messageName,
		Options: restapi.ModelOptions{
			CryptID:    true,
			Timestamps: true,
		},
		Fields: map[string]restapi.Field{
			"from": {
				Type:  schema.Int64,
				Label: "发送者",
			},
			"to": {
				Type:  schema.Int64,
				Label: "接收者",
				Index: true,
			},
			"status": {
				Type:    schema.Uint16,
				Size:    999,
				Label:   "消息状态",
				Default: 0,
				Comment: "0 未读 1 已读",
			},
			"title": {
				Type:     schema.String,
				Size:     100,
				Label:    "标题",
				Nullable: true,
			},
			"message": {
				Type:    schema.Text,
				Default: "",
				Comment: "",
				Label:   "消息",
			},
		},
	}, false)

	if err == nil {
		m.messageModel = &MessageModel{Operation: mod.Operation(), mod: m}
	}
	return err
}

func (m *MessageModel) CountUnread(uid string) (int64, error) {
	id, err := m.mod.accountModel.DeCryptID(uid)
	if err != nil {
		return 0, errors.New("用户 ID 错误")
	}
	return m.Count(ztype.Map{"to": id, "status": 0})
}

const maxCutMessageTitle = 10

func (m *MessageModel) SendMessage(from, to, message string, title ...string) (err error) {
	if message == "" {
		return errors.New("消息内容不能为空")
	}

	if from == "" || to == "" {
		return errors.New("接收者/发送者 ID 不能为空")
	}

	to, err = m.mod.accountModel.DeCryptID(to)
	if err != nil {
		return errors.New("接收者 ID 错误")
	}

	from, err = m.mod.accountModel.DeCryptID(from)
	if err != nil {
		return errors.New("发送者 ID 错误")
	}

	data := ztype.Map{
		"from":    from,
		"to":      to,
		"message": message,
	}

	if len(title) > 0 {
		data["title"] = title[0]
	} else {
		title := zstring.Substr(message, 0, maxCutMessageTitle)
		if len(message) > maxCutMessageTitle {
			title += "..."
		}
		data["title"] = title
	}

	_, err = m.Insert(data)
	if err != nil {
		err = zerror.With(err, "发送消息失败")
	}

	return
}
