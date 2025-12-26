package restapi

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_core/service"
	"github.com/zlsgo/app_module/model"
	mschema "github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb"
	"github.com/zlsgo/zdb/driver/sqlite3"
)

type testEnv struct {
	engine   *znet.Engine
	users    *model.Store
	profiles *model.Store
	cleanup  func()
}

func newTestEnv(t *testing.T, opts *Options) *testEnv {
	if opts == nil {
		opts = &Options{}
	}
	if opts.Prefix == "" {
		opts.Prefix = "/api"
	}
	if opts.MaxPageSize == 0 {
		opts.MaxPageSize = defaultMaxPageSize
	}

	db, err := zdb.New(&sqlite3.Config{File: ":memory:", Memory: true, Parameters: "_pragma=busy_timeout(3000)"})
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}

	schemas := model.NewSchemas(nil, model.NewSQL(db, ""), model.SchemaOptions{})

	profiles := mschema.Schema{
		Name:  "profiles",
		Table: mschema.Table{Name: "profiles"},
		Fields: map[string]mschema.Field{
			"nickname": {Type: mschema.String, Size: 80},
		},
	}

	users := mschema.Schema{
		Name:  "users",
		Table: mschema.Table{Name: "users"},
		Fields: map[string]mschema.Field{
			"name":       {Type: mschema.String, Size: 80},
			"profile_id": {Type: mschema.Int, Nullable: true},
		},
		Relations: map[string]mschema.Relation{
			"profile": {
				Label:      "Profile",
				Type:       mschema.RelationSingle,
				Schema:     "profiles",
				ForeignKey: []string{"profile_id"},
				SchemaKey:  []string{model.IDKey()},
				Fields:     []string{"nickname"},
				Nullable:   true,
			},
		},
	}

	for _, s := range []mschema.Schema{profiles, users} {
		m, err := schemas.Reg(s.Name, s, false)
		if err != nil {
			_ = db.Close()
			t.Fatalf("failed to register schema %s: %v", s.Name, err)
		}
		if err := m.Migration().Auto(model.DealOldColumnNone); err != nil {
			_ = db.Close()
			t.Fatalf("failed to migrate schema %s: %v", s.Name, err)
		}
	}

	stores := schemas.Models()
	userStore, ok := stores.Get("users")
	if !ok {
		_ = db.Close()
		t.Fatalf("users store not found")
	}
	profileStore, ok := stores.Get("profiles")
	if !ok {
		_ = db.Close()
		t.Fatalf("profiles store not found")
	}

	di := zdi.New()
	di.Maps(stores)

	app := service.App{DI: di}
	ctrl := &controller{App: app, options: opts}

	r := znet.New()
	r.SetMode(znet.ProdMode)
	if err := r.BindStruct(opts.Prefix, ctrl); err != nil {
		_ = db.Close()
		t.Fatalf("bind controller: %v", err)
	}

	return &testEnv{
		engine:   r,
		users:    userStore,
		profiles: profileStore,
		cleanup: func() {
			_ = db.Close()
		},
	}
}

func (e *testEnv) request(method, path string, body []byte) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	var reader *bytes.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	} else {
		reader = bytes.NewReader(nil)
	}
	req := httptest.NewRequest(method, path, reader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	e.engine.ServeHTTP(w, req)
	return w
}

func parseData(w *httptest.ResponseRecorder) *zjson.Res {
	return zjson.ParseBytes(w.Body.Bytes()).Get("data")
}

func parseCode(w *httptest.ResponseRecorder) int {
	return zjson.ParseBytes(w.Body.Bytes()).Get("code").Int()
}

func TestRestAPIListFieldsAndOrder(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api", AllowFields: map[string]bool{"id": true, "name": true}})
	defer env.cleanup()

	_, err := env.users.Insert(ztype.Map{"name": "b"})
	tt.NoError(err)
	_, err = env.users.Insert(ztype.Map{"name": "a"})
	tt.NoError(err)

	w := env.request("GET", "/api/users?fields=id,name&order=name:asc", nil)
	tt.Equal(200, w.Code)

	data := parseData(w)
	items := data.Get("items").Maps()
	tt.Equal(2, len(items))
	tt.Equal("a", items[0].Get("name").String())
	_, hasProfileID := items[0]["profile_id"]
	tt.Equal(false, hasProfileID)
}

func TestRestAPIRelationLoad(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{
		Prefix:         "/api",
		AllowFields:    map[string]bool{"id": true, "name": true, "profile_id": true},
		AllowRelations: map[string]bool{"profile": true},
	})
	defer env.cleanup()

	profileID, err := env.profiles.Insert(ztype.Map{"nickname": "p1"})
	tt.NoError(err)
	_, err = env.users.Insert(ztype.Map{"name": "u1", "profile_id": profileID})
	tt.NoError(err)

	w := env.request("GET", "/api/users?fields=id,name,profile_id&with=profile.nickname", nil)
	tt.Equal(200, w.Code)

	data := parseData(w)
	tt.Equal("p1", data.Get("items.0.profile.nickname").String())
}

func TestRestAPIMaxPageSize(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api", MaxPageSize: 2})
	defer env.cleanup()

	for i := 0; i < 5; i++ {
		_, err := env.users.Insert(ztype.Map{"name": fmt.Sprintf("u%d", i)})
		tt.NoError(err)
	}

	w := env.request("GET", "/api/users?pagesize=10", nil)
	tt.Equal(200, w.Code)

	items := parseData(w).Get("items").Maps()
	tt.Equal(2, len(items))
}

func TestRestAPIMethodAllowlist(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api", AllowMethods: map[string]bool{"GET": true}})
	defer env.cleanup()

	w := env.request("POST", "/api/users", []byte(`{"name":"x"}`))
	tt.Equal(405, w.Code)
	tt.Equal(405, parseCode(w))
	tt.Equal(true, strings.Contains(w.Header().Get("Allow"), "GET"))
}

func TestRestAPIMethodAllowlistNormalize(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api", AllowMethods: map[string]bool{"get": true}})
	defer env.cleanup()

	w := env.request("GET", "/api/users", nil)
	tt.Equal(200, w.Code)
}

func TestRestAPIInvalidJSONBody(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api"})
	defer env.cleanup()

	w := env.request("POST", "/api/users", []byte(`{"name":`))
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIInvalidJSONBodyOnUpdate(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api"})
	defer env.cleanup()

	id, err := env.users.Insert(ztype.Map{"name": "u1"})
	tt.NoError(err)

	w := env.request("PUT", fmt.Sprintf("/api/users/%v", id), []byte(`{"name":`))
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIFilter(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api", AllowFilterFields: map[string]bool{"name": true}})
	defer env.cleanup()

	_, err := env.users.Insert(ztype.Map{"name": "b"})
	tt.NoError(err)
	_, err = env.users.Insert(ztype.Map{"name": "a"})
	tt.NoError(err)

	filter := url.QueryEscape(`{"name":{"$eq":"a"}}`)
	w := env.request("GET", "/api/users?filter="+filter, nil)
	tt.Equal(200, w.Code)

	items := parseData(w).Get("items").Maps()
	tt.Equal(1, len(items))
	tt.Equal("a", items[0].Get("name").String())
}

func TestRestAPIInvalidFilterJSON(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api"})
	defer env.cleanup()

	w := env.request("GET", "/api/users?filter=invalid", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIEmptyFilterParam(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api"})
	defer env.cleanup()

	w := env.request("GET", "/api/users?filter=", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIFilterEmptyLogic(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api"})
	defer env.cleanup()

	filter := url.QueryEscape(`{"$and":[]}`)
	w := env.request("GET", "/api/users?filter="+filter, nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIFilterFieldNotAllowed(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api", AllowFilterFields: map[string]bool{"name": true}})
	defer env.cleanup()

	filter := url.QueryEscape(`{"profile_id":{"$eq":1}}`)
	w := env.request("GET", "/api/users?filter="+filter, nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIOrderAllowlistOverride(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{
		Prefix:           "/api",
		AllowFields:      map[string]bool{"id": true},
		AllowOrderFields: map[string]bool{"name": true},
	})
	defer env.cleanup()

	id1, err := env.users.Insert(ztype.Map{"name": "b"})
	tt.NoError(err)
	id2, err := env.users.Insert(ztype.Map{"name": "a"})
	tt.NoError(err)

	w := env.request("GET", "/api/users?fields=id&order=name:asc", nil)
	tt.Equal(200, w.Code)

	items := parseData(w).Get("items").Maps()
	tt.Equal(2, len(items))
	tt.Equal(fmt.Sprint(id2), items[0].Get("id").String())
	tt.Equal(fmt.Sprint(id1), items[1].Get("id").String())
}

func TestRestAPIDefaultFields(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{
		Prefix:        "/api",
		DefaultFields: []string{"id", "name"},
	})
	defer env.cleanup()

	profileID, err := env.profiles.Insert(ztype.Map{"nickname": "p1"})
	tt.NoError(err)
	_, err = env.users.Insert(ztype.Map{"name": "u1", "profile_id": profileID})
	tt.NoError(err)

	w := env.request("GET", "/api/users", nil)
	tt.Equal(200, w.Code)

	items := parseData(w).Get("items").Maps()
	tt.Equal(1, len(items))
	_, hasProfileID := items[0]["profile_id"]
	tt.Equal(false, hasProfileID)
	tt.Equal("u1", items[0].Get("name").String())
}

func TestRestAPIRequireFields(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{
		Prefix:        "/api",
		RequireFields: true,
	})
	defer env.cleanup()

	w := env.request("GET", "/api/users", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIInvalidRelationPath(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{
		Prefix:         "/api",
		AllowRelations: map[string]bool{"profile": true},
	})
	defer env.cleanup()

	w := env.request("GET", "/api/users?with=profile.unknown", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIInvalidOrderField(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api"})
	defer env.cleanup()

	w := env.request("GET", "/api/users?order=unknown:asc", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPINotFoundModel(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api"})
	defer env.cleanup()

	w := env.request("GET", "/api/unknown", nil)
	tt.Equal(404, w.Code)
	tt.Equal(404, parseCode(w))
}

func TestRestAPINotFoundID(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api"})
	defer env.cleanup()

	w := env.request("GET", "/api/users/999", nil)
	tt.Equal(404, w.Code)
	tt.Equal(404, parseCode(w))
}

func TestRestAPINotFoundUpdate(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api"})
	defer env.cleanup()

	w := env.request("PUT", "/api/users/999", []byte(`{"name":"x"}`))
	tt.Equal(404, w.Code)
	tt.Equal(404, parseCode(w))
}

func TestRestAPINotFoundDelete(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api"})
	defer env.cleanup()

	w := env.request("DELETE", "/api/users/999", nil)
	tt.Equal(404, w.Code)
	tt.Equal(404, parseCode(w))
}

func TestRestAPIEmptyFieldsParam(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api"})
	defer env.cleanup()

	w := env.request("GET", "/api/users?fields=", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIEmptyOrderParam(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api"})
	defer env.cleanup()

	w := env.request("GET", "/api/users?order=", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIEmptyRelationParam(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api"})
	defer env.cleanup()

	w := env.request("GET", "/api/users?with=", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))

	w = env.request("GET", "/api/users?relations=", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIRelationParamConflict(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{
		Prefix:         "/api",
		AllowRelations: map[string]bool{"profile": true},
	})
	defer env.cleanup()

	w := env.request("GET", "/api/users?with=profile&relations=profile", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIErrorHandlerOverride(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{
		Prefix: "/api",
		ErrorHandler: func(c *znet.Context, err error) {
			c.JSON(418, znet.ApiData{Code: 418, Msg: "custom"})
		},
	})
	defer env.cleanup()

	w := env.request("GET", "/api/users?order=unknown:asc", nil)
	tt.Equal(418, w.Code)
	tt.Equal(418, parseCode(w))
	tt.Equal("custom", zjson.ParseBytes(w.Body.Bytes()).Get("msg").String())
}

func TestRestAPIDisableErrorHandler(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{
		Prefix:              "/api",
		DisableErrorHandler: true,
	})
	defer env.cleanup()

	w := env.request("GET", "/api/users?order=unknown:asc", nil)
	tt.Equal(500, w.Code)
}

func TestRestAPIStrictQueryRejectUnknown(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{
		Prefix:             "/api",
		RejectUnknownQuery: true,
	})
	defer env.cleanup()

	w := env.request("GET", "/api/users?foo=bar", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIStrictQueryRejectUnknownPost(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{
		Prefix:             "/api",
		RejectUnknownQuery: true,
	})
	defer env.cleanup()

	w := env.request("POST", "/api/users?foo=bar", []byte(`{"name":"x"}`))
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIStrictQueryDisabled(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{Prefix: "/api"})
	defer env.cleanup()

	w := env.request("GET", "/api/users?foo=bar", nil)
	tt.Equal(200, w.Code)
}

func TestRestAPIStrictQueryAllowExtraKey(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{
		Prefix:             "/api",
		RejectUnknownQuery: true,
		AllowQueryKeys:     map[string]bool{"foo": true},
	})
	defer env.cleanup()

	w := env.request("GET", "/api/users?foo=bar", nil)
	tt.Equal(200, w.Code)
}

func TestRestAPIStrictQueryCaseSensitive(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{
		Prefix:             "/api",
		RejectUnknownQuery: true,
	})
	defer env.cleanup()

	w := env.request("GET", "/api/users?Fields=id", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIStrictQueryInvalidPageValues(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{
		Prefix:             "/api",
		RejectUnknownQuery: true,
	})
	defer env.cleanup()

	w := env.request("GET", "/api/users?page=abc", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))

	w = env.request("GET", "/api/users?pagesize=0", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))

	w = env.request("GET", "/api/users?page=", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}

func TestRestAPIStrictQueryMultiValue(t *testing.T) {
	tt := zlsgo.NewTest(t)
	env := newTestEnv(t, &Options{
		Prefix:             "/api",
		RejectUnknownQuery: true,
	})
	defer env.cleanup()

	w := env.request("GET", "/api/users?fields=id&fields=name", nil)
	tt.Equal(400, w.Code)
	tt.Equal(400, parseCode(w))
}
