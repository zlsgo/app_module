package schema

import (
	"reflect"
	"testing"
	"time"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
)

type SimpleUser struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" field:"size:100,label:用户名"`
	Email string `json:"email" field:"size:200,nullable,unique"`
	Age   int    `json:"age" field:"default:0"`
}

func TestFieldsFromStruct_Simple(t *testing.T) {
	tt := zlsgo.NewTest(t)

	fields := FieldsFromStruct[SimpleUser]()
	tt.Equal(4, len(fields))

	tt.Equal(Uint, fields["id"].Type)

	tt.Equal(String, fields["name"].Type)
	tt.Equal(uint64(100), fields["name"].Size)
	tt.Equal("用户名", fields["name"].Label)

	tt.Equal(String, fields["email"].Type)
	tt.Equal(uint64(200), fields["email"].Size)
	tt.Equal(true, fields["email"].Nullable)
	tt.Equal(true, fields["email"].Unique)

	tt.Equal(Int, fields["age"].Type)
	tt.Equal("0", fields["age"].Default)
}

type FullUser struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name" field:"size:50,label:姓名"`
	Profile   *string    `json:"profile" field:"size:500"`
	Age       *int       `json:"age"`
	Score     float64    `json:"score"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Data      ztype.Map  `json:"data"`
	Tags      []string   `json:"tags" field:"array"`
	Password  string     `json:"password" field:"crypt:password,readonly"`
	Ignored   string     `json:"-"`
	Skip      string     `field:"-"`
}

func TestFieldsFromStruct_FullTypes(t *testing.T) {
	tt := zlsgo.NewTest(t)

	fields := FieldsFromStruct[FullUser]()

	tt.Equal(Uint, fields["id"].Type)
	tt.Equal(false, fields["id"].Nullable)

	tt.Equal(String, fields["name"].Type)
	tt.Equal(uint64(50), fields["name"].Size)
	tt.Equal("姓名", fields["name"].Label)

	tt.Equal(String, fields["profile"].Type)
	tt.Equal(true, fields["profile"].Nullable)
	tt.Equal(uint64(500), fields["profile"].Size)

	tt.Equal(Int, fields["age"].Type)
	tt.Equal(true, fields["age"].Nullable)

	tt.Equal(Float, fields["score"].Type)

	tt.Equal(Bool, fields["is_active"].Type)

	tt.Equal(Time, fields["created_at"].Type)
	tt.Equal(false, fields["created_at"].Nullable)

	tt.Equal(Time, fields["updated_at"].Type)
	tt.Equal(true, fields["updated_at"].Nullable)

	tt.Equal(JSON, fields["data"].Type)

	tt.Equal(JSON, fields["tags"].Type)
	tt.Equal(true, fields["tags"].Options.IsArray)

	tt.Equal(String, fields["password"].Type)
	tt.Equal("password", fields["password"].Options.Crypt)
	tt.Equal(true, fields["password"].Options.ReadOnly)

	_, hasIgnored := fields["Ignored"]
	tt.Equal(false, hasIgnored)

	_, hasSkip := fields["skip"]
	tt.Equal(false, hasSkip)
}

type IntTypes struct {
	A int    `json:"a"`
	B int8   `json:"b"`
	C int16  `json:"c"`
	D int32  `json:"d"`
	E int64  `json:"e"`
	F uint   `json:"f"`
	G uint8  `json:"g"`
	H uint16 `json:"h"`
	I uint32 `json:"i"`
	J uint64 `json:"j"`
}

func TestFieldsFromStruct_IntTypes(t *testing.T) {
	tt := zlsgo.NewTest(t)

	fields := FieldsFromStruct[IntTypes]()

	tt.Equal(Int, fields["a"].Type)
	tt.Equal(Int8, fields["b"].Type)
	tt.Equal(Int16, fields["c"].Type)
	tt.Equal(Int32, fields["d"].Type)
	tt.Equal(Int64, fields["e"].Type)
	tt.Equal(Uint, fields["f"].Type)
	tt.Equal(Uint8, fields["g"].Type)
	tt.Equal(Uint16, fields["h"].Type)
	tt.Equal(Uint32, fields["i"].Type)
	tt.Equal(Uint64, fields["j"].Type)
}

type BaseModel struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type ExtendedUser struct {
	BaseModel
	Name string `json:"name" field:"size:100"`
}

func TestFieldsFromStruct_Embedded(t *testing.T) {
	tt := zlsgo.NewTest(t)

	fields := FieldsFromStruct[ExtendedUser]()

	tt.Equal(3, len(fields))
	tt.Equal(Uint, fields["id"].Type)
	tt.Equal(Time, fields["created_at"].Type)
	tt.Equal(String, fields["name"].Type)
	tt.Equal(uint64(100), fields["name"].Size)
}

type NoJsonTag struct {
	ID       uint
	UserName string `field:"size:50"`
}

type MetaUser struct {
	Meta   `name:"meta_user" table:"meta_users" comment:"Meta Users" options:"timestamps,soft_deletes,crypt_id" crypt_len:"16" low_fields:"secret|password" fields_sort:"id|name"`
	ID     uint   `json:"id"`
	Status int    `json:"status" field:"enum:1=active|2=disabled,unique:status_u,index:status_i"`
	Score  int    `json:"score" field:"valid:min=1|max=10"`
	Secret string `json:"secret" field:"disable_migration"`
}

type RelProfile struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type RelRole struct {
	ID uint `json:"id"`
}

type RelUser struct {
	Meta      `name:"rel_user"`
	ID        uint        `json:"id"`
	ProfileID uint        `json:"profile_id"`
	Profile   *RelProfile `relation:"type:single,schema:profiles,foreign:profile_id,fields:id|name"`
	Roles     []RelRole   `relation:"type:many_to_many,schema:roles,pivot_foreign:user_id,pivot_related:role_id"`
}

func TestFieldsFromStruct_NoJsonTag(t *testing.T) {
	tt := zlsgo.NewTest(t)

	fields := FieldsFromStruct[NoJsonTag]()

	tt.Equal(2, len(fields))
	tt.Equal(Uint, fields["id"].Type)
	tt.Equal(String, fields["user_name"].Type)
	tt.Equal(uint64(50), fields["user_name"].Size)
}

func TestFieldsFromStruct_Pointer(t *testing.T) {
	tt := zlsgo.NewTest(t)

	fields := FieldsFromStruct[*SimpleUser]()
	tt.Equal(4, len(fields))
}

func TestFieldsFromStruct_NonStruct(t *testing.T) {
	fields := FieldsFromStruct[string]()
	if fields != nil {
		t.Error("expected nil for non-struct type")
	}
}

func TestNewFromStruct(t *testing.T) {
	tt := zlsgo.NewTest(t)

	s := NewFromStruct[SimpleUser]("user", "users")
	tt.Equal("user", s.Name)
	tt.Equal("users", s.Table.Name)
	tt.Equal(4, len(s.Fields))
	tt.Equal(String, s.Fields["name"].Type)
}

func TestNewFromStructValue(t *testing.T) {
	tt := zlsgo.NewTest(t)

	s := NewFromStructValue("user", SimpleUser{}, "users")
	tt.Equal("user", s.Name)
	tt.Equal("users", s.Table.Name)
	tt.Equal(4, len(s.Fields))
	tt.Equal(String, s.Fields["name"].Type)
}

func TestNewFromStructType(t *testing.T) {
	tt := zlsgo.NewTest(t)

	s := NewFromStructType("user", reflect.TypeOf(SimpleUser{}), "users")
	tt.Equal("user", s.Name)
	tt.Equal("users", s.Table.Name)
	tt.Equal(4, len(s.Fields))
	tt.Equal(String, s.Fields["name"].Type)
}

func TestNewFromStruct_MetaAndTags(t *testing.T) {
	tt := zlsgo.NewTest(t)

	s := NewFromStruct[MetaUser]("", "")
	tt.Equal("meta_user", s.Name)
	tt.Equal("meta_users", s.Table.Name)
	tt.Equal("Meta Users", s.Table.Comment)

	tt.Equal(true, s.Options.Timestamps != nil && *s.Options.Timestamps)
	tt.Equal(true, s.Options.SoftDeletes != nil && *s.Options.SoftDeletes)
	tt.Equal(true, s.Options.CryptID != nil && *s.Options.CryptID)
	tt.Equal(16, s.Options.CryptLen)
	tt.Equal([]string{"secret", "password"}, s.Options.LowFields)
	tt.Equal([]string{"id", "name"}, s.Options.FieldsSort)

	status := s.Fields["status"]
	tt.Equal("status_u", status.Unique)
	tt.Equal("status_i", status.Index)
	tt.Equal(2, len(status.Options.Enum))
	tt.Equal("1", status.Options.Enum[0].Value)
	tt.Equal("active", status.Options.Enum[0].Label)

	score := s.Fields["score"]
	tt.Equal(2, len(score.Validations))
	tt.Equal("min", score.Validations[0].Method)
	tt.Equal("1", score.Validations[0].Args)

	secret := s.Fields["secret"]
	tt.Equal(true, secret.Options.DisableMigration)
}

func TestNewFromStruct_RelationTags(t *testing.T) {
	tt := zlsgo.NewTest(t)

	s := NewFromStruct[RelUser]("", "")

	profile, ok := s.Relations["profile"]
	tt.Equal(true, ok)
	tt.Equal(RelationSingle, profile.Type)
	tt.Equal("profiles", profile.Schema)
	tt.Equal([]string{"profile_id"}, profile.ForeignKey)
	tt.Equal([]string{"id"}, profile.SchemaKey)
	tt.Equal([]string{"id", "name"}, profile.Fields)

	_, hasProfileField := s.Fields["profile"]
	tt.Equal(false, hasProfileField)

	roles, ok := s.Relations["roles"]
	tt.Equal(true, ok)
	tt.Equal(RelationManyToMany, roles.Type)
	tt.Equal([]string{"user_id"}, roles.PivotKeys.Foreign)
	tt.Equal([]string{"role_id"}, roles.PivotKeys.Related)
}
