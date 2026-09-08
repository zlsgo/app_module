package auth

import (
	"errors"
	"time"

	"github.com/sohaha/zlsgo/znet"
	zsession "github.com/sohaha/zlsgo/znet/session"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model"
)

// parseDBTime 解析数据库时间字符串，统一使用 UTC
func parseDBTimeInSession(value string) (time.Time, error) {
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

func (m *Module) currentUser(c *znet.Context) (ztype.Map, zsession.Session, error) {
	session, err := zsession.Get(c)
	if err != nil {
		return nil, nil, err
	}

	userID := session.Get("user_id").String()
	if userID == "" {
		return nil, nil, ErrNotAuthenticated
	}

	user, err := m.userByID(userID)
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			return nil, nil, ErrUserNotFound
		}
		return nil, nil, err
	}

	if user.Get("status").Int() != 1 {
		return nil, nil, ErrUserDisabled
	}

	if err := m.validateSession(session, user); err != nil {
		return nil, nil, err
	}

	m.updateSessionActivity(session.ID())

	return user, session, nil
}

// validateSession 验证会话有效性
func (m *Module) validateSession(session zsession.Session, user ztype.Map) error {
	sessionRow, err := m.sessionByKey(session.ID())
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			return ErrSessionInvalid
		}
		return err
	}

	if sessionRow.Get("status").Int() != 1 {
		return ErrSessionInvalid
	}

	if session.Get("session_version").Int() != user.Get("session_version").Int() {
		return ErrSessionInvalid
	}

	if expireAt := sessionRow.Get("expire_at").String(); expireAt != "" {
		if exp, err := parseDBTimeInSession(expireAt); err == nil && exp.Before(time.Now()) {
			return ErrSessionExpired
		}
	}

	return nil
}

// updateSessionActivity 更新会话活跃时间
func (m *Module) updateSessionActivity(sessionKey string) {
	sessions, ok := m.SessionModel()
	if !ok {
		return
	}
	_, _ = sessions.UpdateMany(model.Filter{
		"session_key": sessionKey,
	}, ztype.Map{
		"last_seen_at": ztime.Now(),
	})
}

func (m *Module) attachSession(user ztype.Map, session zsession.Session) error {
	if session == nil {
		return errors.New("session not initialized")
	}

	session.Set("user_id", user.Get(model.IDKey()).String())
	session.Set("session_version", user.Get("session_version").Int())
	if err := session.Save(); err != nil {
		return err
	}

	rawUserID, err := m.rawUserID(user.Get(model.IDKey()).String())
	if err != nil {
		return err
	}

	sessions, ok := m.SessionModel()
	if !ok {
		return errors.New("session model not ready")
	}

	expireAt := time.Now().Add(time.Duration(m.Options.SessionTTL) * time.Second)
	row, err := sessions.FindOne(model.Filter{"session_key": session.ID()})
	if err != nil && !errors.Is(err, model.ErrNoRecord) {
		return err
	}

	data := ztype.Map{
		"user_id":      rawUserID,
		"session_key":  session.ID(),
		"status":       1,
		"expire_at":    ztime.FormatTime(expireAt),
		"last_seen_at": ztime.Now(),
	}

	if errors.Is(err, model.ErrNoRecord) {
		_, err = sessions.Insert(data)
		return err
	}

	_, err = sessions.UpdateByID(row.Get(model.IDKey()).String(), data)
	return err
}

func (m *Module) invalidateSession(sessionKey string) error {
	sessions, ok := m.SessionModel()
	if !ok {
		return errors.New("session model not ready")
	}

	_, err := sessions.UpdateMany(model.Filter{"session_key": sessionKey}, ztype.Map{
		"status": 0,
	})
	return err
}

func (m *Module) invalidateUserSessions(userID string) error {
	sessions, ok := m.SessionModel()
	if !ok {
		return errors.New("session model not ready")
	}

	rawUserID, err := m.rawUserID(userID)
	if err != nil {
		return err
	}

	_, err = sessions.UpdateMany(model.Filter{"user_id": rawUserID}, ztype.Map{
		"status": 0,
	})
	return err
}

func (m *Module) sessionByKey(sessionKey string) (ztype.Map, error) {
	sessions, ok := m.SessionModel()
	if !ok {
		return nil, errors.New("session model not ready")
	}
	return sessions.FindOne(model.Filter{"session_key": sessionKey})
}

func (m *Module) userByEmail(email string) (ztype.Map, error) {
	users, ok := m.UserModel()
	if !ok {
		return nil, errors.New("user model not ready")
	}
	return users.FindOne(model.Filter{"email": email})
}

func (m *Module) userByID(id string) (ztype.Map, error) {
	users, ok := m.UserModel()
	if !ok {
		return nil, errors.New("user model not ready")
	}
	return users.FindOneByID(id)
}

func (m *Module) rawUserID(id string) (string, error) {
	schema, ok := m.schemas.Get(userModelName)
	if !ok {
		return "", errors.New("user schema not ready")
	}
	return schema.DeCryptID(id)
}

func (m *Module) emailInUse(email, currentUserID string) (bool, error) {
	user, err := m.userByEmail(email)
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			return false, nil
		}
		return false, err
	}
	return user.Get(model.IDKey()).String() != currentUserID, nil
}
