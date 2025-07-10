package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
)

type Info struct {
	jwt.StandardClaims
	Info      string
	IsRefresh bool
	JTI       string
}

func GenToken(info string, key string, expire, refreshExpire int64) (accessToken, refreshToken string, err error) {
	if expire == 0 {
		expire = 86400 // 默认过期时间为 24 小时
	}
	expiresAt := time.Now().Add(time.Duration(expire) * time.Second).Unix()
	claims := Info{
		Info: info,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	if accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(zstring.String2Bytes(key)); err != nil {
		return "", "", zerror.With(err, "failed to generate signature")
	}

	if refreshExpire == 0 || refreshExpire < expiresAt {
		refreshExpire = 604800 + expiresAt // refresh token 的默认过期时间为 7 天后
	}

	claims.StandardClaims.ExpiresAt = refreshExpire
	claims.IsRefresh = true
	claims.JTI = zstring.Rand(16)
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
		return nil, errors.New("invalid token")
	}

	if claims, ok := t.Claims.(*Info); ok && t.Valid {
		now := time.Now().Unix()
		if claims.ExpiresAt < now {
			return nil, errors.New("token expired")
		}
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

	return c.DefaultFormOrQuery(AuthorizationKey, "")
}
