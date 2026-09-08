package member

import (
	"errors"
	"strings"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model"
)

func (m *Module) memberFromAuthSession(c *znet.Context) (*User, error) {
	authUser, err := m.authModule.CurrentUser(c)
	if err != nil {
		return nil, err
	}

	return m.ensureMemberForAuthUser(authUser)
}

func (m *Module) ensureMemberForAuthUser(authUser ztype.Map) (*User, error) {
	userModel, ok := m.UserModel()
	if !ok {
		return nil, errors.New("member model not ready")
	}

	authUserID := strings.TrimSpace(authUser.Get(model.IDKey()).String())
	if authUserID == "" {
		return nil, errors.New("auth user id is empty")
	}

	existing, err := m.findExistingMemberByAuthUser(userModel, authUserID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return existing, nil
	}

	nickname := strings.TrimSpace(authUser.Get("nickname").String())
	if nickname == "" {
		nickname = strings.TrimSpace(authUser.Get("email").String())
	}
	if nickname == "" {
		nickname = authUserID
	}

	id, err := userModel.Insert(ztype.Map{
		"auth_user_id": authUserID,
		"nickname":     nickname,
		"avatar":       authUser.Get("avatar").String(),
		"status":       1,
		"extension":    ztype.Map{},
	})
	if err != nil {
		if !isUniqueConflict(err) {
			return nil, err
		}

		existing, findErr := m.findExistingMemberByAuthUser(userModel, authUserID)
		if findErr != nil {
			return nil, findErr
		}
		if existing != nil {
			return existing, nil
		}
		return nil, err
	}

	return m.UserById(id)
}

func (m *Module) findExistingMemberByAuthUser(userModel *model.Store, authUserID string) (*User, error) {
	existing, err := userModel.FindOne(model.Filter{"auth_user_id": authUserID})
	if err == nil {
		return m.UserById(existing.Get(model.IDKey()).String())
	}
	if err != nil && !errors.Is(err, model.ErrNoRecord) {
		return nil, err
	}

	return nil, nil
}

func isUniqueConflict(err error) bool {
	if err == nil {
		return false
	}

	msg := err.Error()
	return strings.Contains(msg, "Duplicate entry") || strings.Contains(msg, "UNIQUE")
}
