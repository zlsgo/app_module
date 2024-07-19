package auth

import (
	"github.com/sohaha/zlsgo/znet"
)

type Provider struct {
	Provider         string `json:"provider,omitempty"`
	ProviderID       string `json:"provider_id,omitempty"`
	ProviderUsername string `json:"provider_username,omitempty"`
}

type AuthProvider interface {
	Name() string
	Init() error
	Login(*znet.Context) (any, error)
	Callback(*znet.Context) (Provider, error)
}
