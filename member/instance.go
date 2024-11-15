package member

import (
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/zlsgo/app_module/account/jwt"
	"github.com/zlsgo/app_module/model"
)

type Instance struct {
	module     *Module
	middleware func(optionalRoute ...string) (middleware func(c *znet.Context) error)
}

func (ins *Instance) GetMiddleware(optionalRoute ...string) (middleware func(c *znet.Context) error) {
	return ins.middleware(optionalRoute...)
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
		middleware: func(optionalRoute ...string) func(c *znet.Context) error {
			return func(c *znet.Context) error {
				member := &User{}
				c.Injector().Map(member)

				isOptionalRoute := false
				for i := range optionalRoute {
					if zstring.Match(c.Request.URL.Path, optionalRoute[i]) {
						isOptionalRoute = true
						break
					}
				}

				token := jwt.GetToken(c)
				if token == "" && !isOptionalRoute {
					return zerror.Unauthorized.Text("please login first")
				}

				if token == "" {
					return nil
				}

				info, err := jwt.Parse(token, m.Options.key)
				if err != nil {
					return zerror.Unauthorized.Text(err.Error())
				}

				// salt := info.Info[:saltLen]
				uid := info.Info[saltLen:]

				user, err := m.UserById(uid)
				if err != nil {
					return zerror.Unauthorized.Text(err.Error())
				}

				member.Id = user.Id
				member.RawId = user.RawId
				member.Info = user.Info

				// 删除敏感信息
				_ = member.Info.Delete("password")
				_ = member.Info.Delete("salt")

				c.Next()

				return nil
			}
		},
	}
	return nil
}
