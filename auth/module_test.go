package auth

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/internal/testutil"
	"github.com/zlsgo/zdb"
)

func TestModuleInit(t *testing.T) {
	tt := zlsgo.NewTest(t)

	di := zdi.New()
	di.Map(&[]service.Task{})
	di.Map(&[]service.Controller{})
	app := service.NewApp()(di)

	db := testutil.NewMemoryDB(t)
	defer func() {
		_ = db.Close()
	}()

	mod := New(func(o *Options) {
		o.InitDB = func() (*zdb.DB, error) { return db, nil }
	})

	err := service.InitModule([]service.Module{mod}, app)
	tt.NoError(err)

	_, ok := mod.UserModel()
	tt.Equal(true, ok)
	_, ok = mod.ProviderModel()
	tt.Equal(true, ok)
	_, ok = mod.PasswordResetTokenModel()
	tt.Equal(true, ok)
	_, ok = mod.SessionModel()
	tt.Equal(true, ok)
}

func TestEnabledProviders(t *testing.T) {
	tt := zlsgo.NewTest(t)

	mod := New(func(o *Options) {
		o.Providers = []AuthProvider{
			&stubProvider{name: "one"},
			&stubProvider{name: "two"},
		}
		o.EnabledProviders = []string{"two"}
	})

	got := mod.EnabledProviders()
	tt.Equal(1, len(got))
	tt.Equal("two", got[0].Name())
}

type stubProvider struct {
	name string
}

func (s *stubProvider) Name() string {
	return s.name
}

func (s *stubProvider) Init(*znet.Engine) error {
	return nil
}

func (s *stubProvider) Login(*znet.Context) error {
	return nil
}

func (s *stubProvider) Callback(*znet.Context) (Provider, error) {
	return Provider{}, nil
}
