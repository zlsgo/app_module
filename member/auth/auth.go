package auth

import (
	"fmt"

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

var providers = make(map[string]AuthProvider)

func Register(provider AuthProvider) error {
	if _, ok := providers[provider.Name()]; ok {
		return fmt.Errorf("auth provider %s already exists", provider.Name())
	}
	providers[provider.Name()] = provider
	return nil
}

func GetProvider(name string) (AuthProvider, error) {
	provider, ok := providers[name]
	if !ok {
		return nil, fmt.Errorf("auth provider %s not found", name)
	}
	return provider, nil
}

func GetProviders() map[string]AuthProvider {
	return providers
}
