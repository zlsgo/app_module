package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
)

type Info struct {
	jwt.RegisteredClaims
	Info      string
	IsRefresh bool
}

func GenToken(info string, key string, expire, refreshExpire int64) (accessToken, refreshToken string, err error) {
	if expire == 0 {
		expire = 86400 // 默认过期时间为 24 小时
	}
	expiresAt := time.Now().Add(time.Duration(expire) * time.Second)
	claims := Info{
		Info: info,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	if accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(zstring.String2Bytes(key)); err != nil {
		return "", "", zerror.With(err, "failed to generate signature")
	}

	expiresAtUnix := expiresAt.Unix()
	if refreshExpire == 0 || refreshExpire < expiresAtUnix {
		refreshExpire = 604800 + expiresAtUnix // refresh token 的默认过期时间为 7 天后
	}

	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Unix(refreshExpire, 0))
	claims.IsRefresh = true
	claims.ID = zstring.Rand(16)
	if refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(zstring.String2Bytes(key)); err != nil {
		return "", "", zerror.With(err, "failed to generate refresh token signature")
	}

	return
}

func Parse(token string, tokenKey string) (*Info, error) {
	t, err := jwt.ParseWithClaims(token, &Info{}, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return zstring.String2Bytes(tokenKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token expired")
		}
		return nil, errors.New("invalid token")
	}

	if claims, ok := t.Claims.(*Info); ok && t.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

const AuthorizationKey = "Authorization"

func GetToken(c *znet.Context) string {
	authorization := c.GetHeader(AuthorizationKey)
	if authorization == "" {
		authorization = c.GetCookie(AuthorizationKey)
	}

	if authorization != "" {
		if i := strings.IndexByte(authorization, ' '); i > 0 {
			authorization = zstring.TrimSpace(authorization[i+1:])
		}

		if len(strings.Split(authorization, ".")) == 3 {
			return authorization
		}

		// 兼容 Basic base64(user:pass)
		v, err := zstring.Base64Decode(zstring.String2Bytes(authorization))
		if err != nil {
			return ""
		}
		return strings.Split(zstring.Bytes2String(v), ":")[0]
	}

	return c.DefaultFormOrQuery(AuthorizationKey, "")
}
