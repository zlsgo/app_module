package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/account"
	"github.com/zlsgo/app_module/model"
)

// hashToken 生成 token 的哈希值用于数据库存储
func hashToken(token string) string {
	h := sha256.New()
	h.Write([]byte(token))
	return hex.EncodeToString(h.Sum(nil))
}

func (h *UserController) forgotPassword(c *znet.Context) (any, error) {
	if h.module.Options.SendResetPassword == nil {
		return nil, toHTTPError(ErrResetSenderNotConfigured)
	}

	input, err := getPayload(c)
	if err != nil {
		return nil, toHTTPError(err)
	}

	email := normalizeEmail(input.Get("email").String())
	if !isValidEmail(email) {
		return nil, toHTTPError(ErrInvalidEmail)
	}

	forgotKey := "forgot:" + email
	forgotIPKey := "forgot-ip:" + c.GetClientIP()
	if h.module.tooManyRequests(forgotKey, forgotMaxAttemptsPerEmail, "forgot") || h.module.tooManyRequests(forgotIPKey, forgotMaxAttemptsPerIP, "forgot") {
		return nil, toHTTPError(ErrTooManyRequests)
	}
	h.module.markFailed(forgotKey, "forgot")
	h.module.markFailed(forgotIPKey, "forgot")

	user, err := h.module.userByEmail(email)
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			return ztype.Map{"sent": true}, nil
		}
		return nil, toHTTPError(err)
	}

	job, err := h.module.createResetPasswordJob(c, user)
	if err != nil {
		return nil, toHTTPError(err)
	}
	if err = h.module.Options.SendResetPassword(job); err != nil {
		return nil, toHTTPError(err)
	}

	return ztype.Map{"sent": true}, nil
}

func (h *UserController) resetPassword(c *znet.Context) (any, error) {
	input, err := getPayload(c)
	if err != nil {
		return nil, toHTTPError(err)
	}

	token := strings.TrimSpace(input.Get("token").String())
	password := input.Get("password").String()
	if token == "" {
		return nil, toHTTPError(ErrMissingResetToken)
	}
	if ok, msg := account.ValidatePassword(password, account.DefaultPasswordConfig); !ok {
		return nil, respondBadRequest(msg)
	}

	tokenRow, user, err := h.module.validateResetToken(token)
	if err != nil {
		return nil, toHTTPError(err)
	}

	users, _ := h.module.UserModel()
	tokens, _ := h.module.PasswordResetTokenModel()
	sessions, _ := h.module.SessionModel()
	rawUserID, err := h.module.rawUserID(user.Get(model.IDKey()).String())
	if err != nil {
		return nil, toHTTPError(err)
	}

	err = h.module.schemas.Storage().Transaction(func(s model.Storageer) error {
		txUsers := users.Schema(s).Model()
		txTokens := tokens.Schema(s).Model()
		txSessions := sessions.Schema(s).Model()

		_, err := txUsers.UpdateByID(user.Get(model.IDKey()).String(), ztype.Map{
			"password":        password,
			"session_version": user.Get("session_version").Int() + 1,
		})
		if err != nil {
			return err
		}

		_, err = txTokens.UpdateByID(tokenRow.Get(model.IDKey()).String(), ztype.Map{
			"used_at": ztime.Now(),
		})
		if err != nil {
			return err
		}

		_, err = txSessions.UpdateMany(model.Filter{"user_id": rawUserID}, ztype.Map{
			"status": 0,
		})
		return err
	})
	if err != nil {
		return nil, toHTTPError(err)
	}

	return ztype.Map{"reset": true}, nil
}

func (m *Module) createResetPasswordJob(c *znet.Context, user ztype.Map) (ResetPasswordJob, error) {
	tokens, ok := m.PasswordResetTokenModel()
	if !ok {
		return ResetPasswordJob{}, errors.New("password reset token model not ready")
	}

	rawUserID, err := m.rawUserID(user.Get(model.IDKey()).String())
	if err != nil {
		return ResetPasswordJob{}, err
	}

	token := zstring.UUID()
	tokenHash := hashToken(token)
	expireAt := time.Now().Add(time.Duration(m.Options.PasswordResetTokenTTL) * time.Second)

	err = m.schemas.Storage().Transaction(func(s model.Storageer) error {
		txTokens := tokens.Schema(s).Model()
		_, err := txTokens.DeleteMany(model.Filter{"user_id": rawUserID})
		if err != nil {
			return err
		}

		_, err = txTokens.Insert(ztype.Map{
			"user_id":    rawUserID,
			"token":      token,
			"token_hash": tokenHash,
			"expire_at":  ztime.FormatTime(expireAt),
		})
		return err
	})
	if err != nil {
		return ResetPasswordJob{}, err
	}

	return ResetPasswordJob{
		User: ResetPasswordUser{
			ID:       user.Get(model.IDKey()).String(),
			Email:    user.Get("email").String(),
			Nickname: user.Get("nickname").String(),
		},
		Token:    token,
		ResetURL: m.resetPasswordURL(c, token),
	}, nil
}

func (m *Module) validateResetToken(token string) (ztype.Map, ztype.Map, error) {
	tokens, ok := m.PasswordResetTokenModel()
	if !ok {
		return nil, nil, ErrModelNotReady
	}

	tokenHash := hashToken(token)
	row, err := tokens.FindOne(model.Filter{"token_hash": tokenHash})
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			return nil, nil, ErrInvalidToken
		}
		return nil, nil, err
	}

	// 二次验证 token 匹配
	if row.Get("token").String() != token {
		return nil, nil, ErrInvalidToken
	}

	if row.Get("used_at").String() != "" {
		return nil, nil, ErrTokenUsed
	}

	expireAt, err := parseDBTime(row.Get("expire_at").String())
	if err != nil {
		return nil, nil, err
	}
	if expireAt.Before(time.Now()) {
		return nil, nil, ErrTokenExpired
	}

	user, err := m.userByRawID(row.Get("user_id").String())
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			return nil, nil, ErrUserNotFound
		}
		return nil, nil, err
	}

	return row, user, nil
}

func (m *Module) resetPasswordURL(c *znet.Context, token string) string {
	path := m.Options.ResetPasswordPath + "?token=" + url.QueryEscape(token)
	base := strings.TrimSuffix(m.Options.BaseURL, "/")
	if base == "" && c != nil && c.Request != nil {
		scheme := "http"
		if forwarded := strings.TrimSpace(strings.Split(c.Request.Header.Get("X-Forwarded-Proto"), ",")[0]); forwarded != "" {
			scheme = forwarded
		} else if c.Request.TLS != nil {
			scheme = "https"
		}

		host := strings.TrimSpace(strings.Split(c.Request.Header.Get("X-Forwarded-Host"), ",")[0])
		if host == "" {
			host = c.Request.Host
		}
		if host != "" {
			base = scheme + "://" + host
		}
	}
	if base == "" {
		return path
	}

	return base + path
}

func (m *Module) userByRawID(rawID string) (ztype.Map, error) {
	schema, ok := m.schemas.Get(userModelName)
	if !ok {
		return nil, errors.New("user schema not ready")
	}

	id, err := schema.EnCryptID(rawID)
	if err != nil {
		return nil, err
	}

	return m.userByID(id)
}

func parseDBTime(value string) (time.Time, error) {
	// 尝试多种时间格式，优先使用 UTC
	formats := []string{
		time.RFC3339,
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05Z",
		"2006-01-02 15:04:05.999999999",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, value); err == nil {
			// 如果格式不包含时区信息，假设为 UTC
			if format == "2006-01-02 15:04:05" || format == "2006-01-02 15:04:05.999999999" {
				return t.UTC(), nil
			}
			return t, nil
		}
	}

	return time.Time{}, errors.New("unsupported time format: " + value)
}
