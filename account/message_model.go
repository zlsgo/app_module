package account

import (
	"errors"

	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model"
	"github.com/zlsgo/app_module/model/define"
	"github.com/zlsgo/zdb/schema"
)

type MessageModel struct {
	*model.Operation
	model  *model.Model
	module *Module
}

var messageModel *MessageModel

func GetMessageModel() (*MessageModel, error) {
	if messageModel == nil {
		return nil, errors.New("message model not define")
	}
	return messageModel, nil
}

func messageModelDefine(m *Module) error {
	const messageName = "message"
	mod, err := m.mods.Reg(messageName, define.Define{
		Name: messageName,
		Options: define.ModelOptions{
			CryptID:    true,
			Timestamps: true,
		},
		Fields: map[string]define.Field{
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
			"mtype": {
				Type:     schema.String,
				Size:     100,
				Label:    "消息类型",
				Nullable: true,
			},
		},
	}, false)

	if err == nil {
		messageModel = &MessageModel{model: mod, module: m, Operation: mod.Operation()}
	}
	return err
}

func (m *MessageModel) Unread(uid string) (ztype.Map, error) {
	id, err := GetAccountModel().DeCryptID(uid)
	if err != nil {
		return nil, errors.New("用户 ID 错误")
	}

	resp, err := m.Find(ztype.Map{"to": id, "status": 0}, func(co *model.CondOptions) error {
		co.Fields = []string{model.IDKey, model.CreatedAtKey, "mtype"}
		co.OrderBy = map[string]string{model.IDKey: "desc"}
		return nil
	})
	if err != nil {
		return nil, err
	}

	last, unread, mtype := int64(0), len(resp), ""
	if !resp.IsEmpty() {
		first := resp.First()
		mtype = first.Get("mtype").String()
		t, _ := first.Get(model.CreatedAtKey).Time()
		last = t.Unix()
	}

	return ztype.Map{
		"unread":    unread,
		"last_time": last,
		"mtype":     mtype,
	}, nil
}

const maxCutMessageTitle = 10

func (m *MessageModel) SendMessage(from, to, title, message string, mtype ...string) (err error) {
	if message == "" {
		return errors.New("消息内容不能为空")
	}

	if from == "" || to == "" {
		return errors.New("接收者/发送者 ID 不能为空")
	}

	to, err = GetAccountModel().DeCryptID(to)
	if err != nil {
		return errors.New("接收者 ID 错误")
	}

	from, err = GetAccountModel().DeCryptID(from)
	if err != nil {
		return errors.New("发送者 ID 错误")
	}

	data := ztype.Map{
		"from":    from,
		"to":      to,
		"message": message,
		"title":   title,
		"mtype":   "",
	}

	if len(title) == 0 {
		t := zstring.Substr(message, 0, maxCutMessageTitle)
		if len(message) > maxCutMessageTitle {
			t += "..."
		}
		data["title"] = t
	}

	if len(mtype) > 0 {
		data["mtype"] = mtype[0]
	}

	_, err = m.Insert(data)
	if err != nil {
		err = zerror.With(err, "发送消息失败")
	}

	return
}
