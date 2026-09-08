package member

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_module/model"
)

type Instance struct {
	module     *Module
	middleware func(c *znet.Context) error
}

func (ins *Instance) GetMiddleware() func(c *znet.Context) error {
	return ins.middleware
}

func (ins *Instance) GetMemberModel() *model.Store {
	return ins.module.schemas.MustGet(modelName).Model()
}

func (ins *Instance) GetModel(name string) (*model.Store, bool) {
	s, ok := ins.module.schemas.Get(name)
	if !ok {
		return nil, false
	}
	return s.Model(), true
}

func (m *Module) Instance() *Instance {
	if m.instance == nil {
		_ = initInstance(m)
	}
	return m.instance
}

func initInstance(m *Module) error {
	m.instance = &Instance{
		module: m,
		middleware: func(c *znet.Context) error {
			member := &User{}
			c.Injector().Map(member)

			authMember, err := m.memberFromAuthSession(c)
			if err != nil {
				return err
			}

			member.Id = authMember.Id
			member.RawId = authMember.RawId
			member.Info = authMember.Info
			_ = member.Info.Delete("auth_user_id")

			c.Next()
			return nil
		},
	}
	return nil
}
