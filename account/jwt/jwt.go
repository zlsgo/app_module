package jwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
)

type Info struct {
	Info      string
	IsRefresh bool
	jwt.StandardClaims
}

func GenToken(info string, key string, expire int) (accessToken, refreshToken string, err error) {
	if expire == 0 {
		expire = 3600 * 24
	}
	expiresAt := time.Now().Add(time.Duration(expire) * time.Second).Unix()
	claims := Info{
		Info: info,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	if accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(zstring.String2Bytes(key)); err != nil {
		return "", "", fmt.Errorf("生成签名失败: %v", err)
	}

	claims.IsRefresh = true
	claims.StandardClaims.ExpiresAt = expiresAt + 3600*24*30
	if refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(zstring.String2Bytes(key)); err != nil {
		return "", "", fmt.Errorf("生成续期签名失败: %v", err)
	}

	return
}

func Parse(token string, tokenKey string) (*Info, error) {
	t, err := jwt.ParseWithClaims(token, &Info{}, func(token *jwt.Token) (i interface{}, err error) {
		return zstring.String2Bytes(tokenKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*Info); ok && t.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GetToken(c *znet.Context) string {
	authorization := c.GetHeader("Authorization")
	slen := len("Basic ")
	if len(authorization) > slen {
		authorization = zstring.TrimSpace(authorization[slen:])
		split := strings.Split(authorization, ".")
		if len(split) == 3 {
			return authorization
		}
		v, err := zstring.Base64Decode(zstring.String2Bytes(authorization))
		if err != nil {
			return ""
		}
		return strings.Split(zstring.Bytes2String(v), ":")[0]
	}
	return c.DefaultFormOrQuery("Authorization", "")
}
