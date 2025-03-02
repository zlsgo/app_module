package jwt

import (
	"testing"

	"github.com/sohaha/zlsgo"
)

func TestToken(t *testing.T) {
	tt := zlsgo.NewTest(t)

	value := "AO632VolejRejNmG"
	key := "wV94b4wBV8jGOJTVcM4ZUsCx9kJF30aj"

	token, refreshToken, err := GenToken(value, key, 0, 0)
	tt.NoError(err)
	tt.Log(token)

	result, err := Parse(token, key)
	tt.Equal(false, result.IsRefresh)
	tt.NoError(err)

	refreshResult, err := Parse(refreshToken, key)
	tt.Equal(true, refreshResult.IsRefresh)
	tt.NoError(err)

	tt.Equal(value, result.Info)
	tt.Log(result)
	tt.Log(refreshResult)

	tt.Log(Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJbmZvIjoiQU82MzJWb2xlalJlak5tRyIsIklzUmVmcmVzaCI6ZmFsc2UsImV4cCI6MTcwOTI5MDI0M30.i2ukxHRXoW350p_zRutjonAJvGRxPaHGrHIx-D6Ui14", key))
}
