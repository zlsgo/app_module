package schema

import (
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

func TestFieldsFromStruct_NoJsonTag(t *testing.T) {
	tt := zlsgo.NewTest(t)

	fields := FieldsFromStruct[NoJsonTag]()

	tt.Equal(2, len(fields))
	tt.Equal(Uint, fields["id"].Type)
	tt.Equal(String, fields["username"].Type)
	tt.Equal(uint64(50), fields["username"].Size)
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
