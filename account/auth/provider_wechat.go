package auth

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/wechat"
)

type Weapp struct {
	wx        *wechat.Engine
	AppId     string
	AppSecret string
}

var _ AuthProvider = (*Weapp)(nil)

func (w *Weapp) Name() string {
	return "weapp"
}

func (w *Weapp) Init(r *znet.Engine) error {
	w.wx = wechat.New(&wechat.Weapp{
		AppID:     w.AppId,
		AppSecret: w.AppSecret,
	})
	return nil
}

func (w *Weapp) Login(c *znet.Context) error {
	return nil
}

func (w *Weapp) Callback(c *znet.Context) (Provider, error) {
	code := c.DefaultQuery("code", "")

	info, err := w.wx.GetAuthInfo(code)
	if err != nil {
		return Provider{}, err
	}

	openid := info.Get("openid").String()

	return Provider{
		Provider:         w.Name(),
		ProviderID:       openid,
		ProviderUsername: info.Get("nickname").String(),
	}, nil
}
