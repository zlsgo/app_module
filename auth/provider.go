package auth

import (
	accountauth "github.com/zlsgo/app_module/account/auth"
)

type Provider = accountauth.Provider

type AuthProvider = accountauth.AuthProvider

type ProviderAdapter struct {
	AuthProvider
}
