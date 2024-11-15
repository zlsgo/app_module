package auth

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
)

type Provider struct {
	Provider          string    `json:"provider,omitempty"`
	ProviderID        string    `json:"provider_id,omitempty"`
	ProviderUsername  string    `json:"provider_username,omitempty"`
	ProviderAvatar    string    `json:"provider_avatar,omitempty"`
	ProviderExtension ztype.Map `json:"provider_extension,omitempty"`
}

type AuthProvider interface {
	Name() string
	Init(r *znet.Engine) error
	Login(c *znet.Context) error
	Callback(c *znet.Context) (Provider, error)
}
