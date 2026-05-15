package auth

import (
	"errors"
	"time"

	"github.com/sohaha/zlsgo/znet"
	zsession "github.com/sohaha/zlsgo/znet/session"
	"github.com/sohaha/zlsgo/zstring"
)

const (
	oauthFlowActionLogin = "login"
	oauthFlowActionBind  = "bind"

	oauthFlowStateKey     = "auth_oauth_state"
	oauthFlowProviderKey  = "auth_oauth_provider"
	oauthFlowActionKey    = "auth_oauth_action"
	oauthFlowUserIDKey    = "auth_oauth_user_id"
	oauthFlowExpiresAtKey = "auth_oauth_expires_at"
)

type oauthFlow struct {
	State     string
	Provider  string
	Action    string
	UserID    string
	ExpiresAt time.Time
}

func (m *Module) beginOAuthFlow(c *znet.Context, action, provider, userID string) (string, error) {
	session, err := zsession.Get(c)
	if err != nil {
		return "", err
	}

	state := zstring.UUID()
	expiresAt := time.Now().Add(oauthFlowTTL)
	session.Set(oauthFlowStateKey, state)
	session.Set(oauthFlowProviderKey, provider)
	session.Set(oauthFlowActionKey, action)
	session.Set(oauthFlowUserIDKey, userID)
	session.Set(oauthFlowExpiresAtKey, expiresAt.Unix())
	if err = session.Save(); err != nil {
		return "", err
	}

	injectOAuthState(c, state)
	return state, nil
}

func (m *Module) currentOAuthFlow(c *znet.Context) (oauthFlow, zsession.Session, error) {
	session, err := zsession.Get(c)
	if err != nil {
		return oauthFlow{}, nil, err
	}

	flow := oauthFlow{
		State:    session.Get(oauthFlowStateKey).String(),
		Provider: session.Get(oauthFlowProviderKey).String(),
		Action:   session.Get(oauthFlowActionKey).String(),
		UserID:   session.Get(oauthFlowUserIDKey).String(),
	}
	expiresAt := session.Get(oauthFlowExpiresAtKey).Int64()
	if expiresAt > 0 {
		flow.ExpiresAt = time.Unix(expiresAt, 0)
	}

	if flow.State == "" || flow.Provider == "" || flow.Action == "" || flow.ExpiresAt.IsZero() {
		return oauthFlow{}, session, ErrOAuthStateInvalid
	}
	if flow.ExpiresAt.Before(time.Now()) {
		_ = m.clearOAuthFlow(session)
		return oauthFlow{}, session, ErrOAuthStateExpired
	}

	return flow, session, nil
}

func (m *Module) validateOAuthFlow(c *znet.Context, provider string) (oauthFlow, zsession.Session, error) {
	flow, session, err := m.currentOAuthFlow(c)
	if err != nil {
		return oauthFlow{}, session, err
	}
	if flow.Provider != provider {
		_ = m.clearOAuthFlow(session)
		return oauthFlow{}, session, ErrProviderMismatch
	}

	state := c.DefaultFormOrQuery("state", "")
	if state == "" || state != flow.State {
		_ = m.clearOAuthFlow(session)
		return oauthFlow{}, session, ErrOAuthStateMismatch
	}

	return flow, session, nil
}

func (m *Module) clearOAuthFlow(session zsession.Session) error {
	if session == nil {
		return errors.New("session not initialized")
	}
	for _, key := range []string{
		oauthFlowStateKey,
		oauthFlowProviderKey,
		oauthFlowActionKey,
		oauthFlowUserIDKey,
		oauthFlowExpiresAtKey,
	} {
		_ = session.Delete(key)
	}
	return session.Save()
}

func injectOAuthState(c *znet.Context, state string) {
	query := c.Request.URL.Query()
	query.Set("state", state)
	c.Request.URL.RawQuery = query.Encode()
}
