package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/model/schema"
)

// 回归：时间型软删默认过滤必须注入 IS NULL 而非 nil（nil 会生成 `= NULL` 查不到行）。
func TestGetFilterTimeSoftDeleteEndToEnd(t *testing.T) {
	tt := zlsgo.NewTest(t)

	b := true
	softTime := true
	s := schema.Schema{
		Name: "audit_fix_soft_time",
		Table: schema.Table{
			Name: "audit_fix_soft_time",
		},
		Options: schema.Options{
			SoftDeletes:      &b,
			SoftDeleteIsTime: &softTime,
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, s)
	m := schemas.MustGet("audit_fix_soft_time")
	repo := m.Model().Repository()

	id1, err := repo.Insert(ztype.Map{"name": "a"})
	tt.NoError(err)
	_, err = repo.Insert(ztype.Map{"name": "b"})
	tt.NoError(err)

	// 默认过滤应能查到全部未删除行
	rows, err := repo.Find(Filter{})
	tt.NoError(err)
	tt.Equal(2, len(rows))

	// 软删后不再返回
	_, err = repo.DeleteByID(id1)
	tt.NoError(err)
	rows, err = repo.Find(Filter{})
	tt.NoError(err)
	tt.Equal(1, len(rows))
}

// 回归：嵌套 $OR/$AND 分组的字段同样需经过白名单校验。
func TestGetFilterNestedGroupWhitelist(t *testing.T) {
	tt := zlsgo.NewTest(t)

	s := schema.Schema{
		Name: "audit_fix_nested",
		Table: schema.Table{
			Name: "audit_fix_nested",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	db, schemas := newTestSchemas(t, s)
	_ = db
	m := schemas.MustGet("audit_fix_nested")
	repo := m.Model().Repository()
	_, err := repo.Insert(ztype.Map{"name": "a"})
	tt.NoError(err)
	_, err = repo.Insert(ztype.Map{"name": "b"})
	tt.NoError(err)

	// 顶层非法字段仍被剔除
	top := getFilter(m, Filter{"bogus_col": 1, "name": "b"})
	_, hasBogusTop := top["bogus_col"]
	tt.Equal(false, hasBogusTop)

	// 嵌套分组内非法字段被剔除；净化后为空的子条件保留为恒假占位，合法条件不受影响
	f := getFilter(m, Filter{"$OR": []ztype.Map{{"bogus_col": 1}, {"name": "b"}}})
	group, ok := f["$OR"].([]ztype.Map)
	tt.Equal(true, ok)
	tt.Equal(2, len(group))
	_, hasBogus := group[0]["bogus_col"]
	tt.Equal(false, hasBogus)
	_, isFalse := group[0][idKey+" IN"]
	tt.Equal(true, isFalse)
	tt.Equal("b", group[1]["name"])

	// 净化后的过滤器应能正常执行 SQL 且只返回合法条件命中的行
	rows, err := m.Storage.Find(m.GetTableName(), f)
	tt.NoError(err)
	tt.Equal(1, len(rows))
	tt.Equal("b", rows[0]["name"])
}

// 回归：$AND 数组分组必须按 AND 组合（旧实现会把整组拍平成 OR，导致结果集扩大）。
func TestGetFilterAndArraySemantics(t *testing.T) {
	tt := zlsgo.NewTest(t)

	s := schema.Schema{
		Name: "audit_fix_and_arr",
		Table: schema.Table{
			Name: "audit_fix_and_arr",
		},
		Fields: map[string]schema.Field{
			"name":   {Type: "string", Size: 80},
			"status": {Type: "int"},
		},
	}

	_, schemas := newTestSchemas(t, s)
	m := schemas.MustGet("audit_fix_and_arr")
	repo := m.Model().Repository()
	_, err := repo.Insert(ztype.Map{"name": "Admin", "status": 1})
	tt.NoError(err)
	_, err = repo.Insert(ztype.Map{"name": "Editor", "status": 2})
	tt.NoError(err)

	// $AND：互斥条件 → 0 行（若被误解析为 OR 会返回 2 行）
	andF := getFilter(m, Filter{"$AND": []ztype.Map{{"name": "Admin"}, {"status": 2}}})
	rows, err := m.Storage.Find(m.GetTableName(), andF)
	tt.NoError(err)
	tt.Equal(0, len(rows))

	// $OR：任一条件命中 → 2 行
	orF := getFilter(m, Filter{"$OR": []ztype.Map{{"name": "Admin"}, {"status": 2}}})
	rows, err = m.Storage.Find(m.GetTableName(), orF)
	tt.NoError(err)
	tt.Equal(2, len(rows))
}

// 回归：$OR 元素内的多个字段必须保持组内 AND，
// 不得被拍平成跨元素的扁平 OR（否则 (A AND B) OR C 会变成 A OR B OR C）。
func TestGetFilterOrElementKeepsConjunction(t *testing.T) {
	tt := zlsgo.NewTest(t)

	s := schema.Schema{
		Name: "audit_fix_or_el",
		Table: schema.Table{
			Name: "audit_fix_or_el",
		},
		Fields: map[string]schema.Field{
			"name":   {Type: "string", Size: 80},
			"status": {Type: "int"},
		},
	}

	_, schemas := newTestSchemas(t, s)
	m := schemas.MustGet("audit_fix_or_el")
	repo := m.Model().Repository()
	_, err := repo.Insert(ztype.Map{"name": "Admin", "status": 1})
	tt.NoError(err)
	_, err = repo.Insert(ztype.Map{"name": "Editor", "status": 2})
	tt.NoError(err)

	// (name=Admin AND status=2) OR name=Editor → 只有 Editor 命中
	f := getFilter(m, Filter{"$OR": []ztype.Map{{"name": "Admin", "status": 2}, {"name": "Editor"}}})
	rows, err := m.Storage.Find(m.GetTableName(), f)
	tt.NoError(err)
	tt.Equal(1, len(rows))
	tt.Equal("Editor", rows[0]["name"])
}

// 回归：净化后为空的分组保留恒假条件，查询不得扩大为全表扫描。
func TestGetFilterEmptyGroupFailClosed(t *testing.T) {
	tt := zlsgo.NewTest(t)

	s := schema.Schema{
		Name: "audit_fix_empty_group",
		Table: schema.Table{
			Name: "audit_fix_empty_group",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, s)
	m := schemas.MustGet("audit_fix_empty_group")
	repo := m.Model().Repository()
	_, err := repo.Insert(ztype.Map{"name": "a"})
	tt.NoError(err)
	_, err = repo.Insert(ztype.Map{"name": "b"})
	tt.NoError(err)

	// $OR 内全部为非法字段 → 整组变成恒假，返回 0 行而不是全表
	f := getFilter(m, Filter{"$OR": []ztype.Map{{"bogus_col": 1}}})
	rows, err := m.Storage.Find(m.GetTableName(), f)
	tt.NoError(err)
	tt.Equal(0, len(rows))

	// $AND 数组混入非法子条件 → 整组恒假，而不是退化为剩余条件
	f = getFilter(m, Filter{"$AND": []ztype.Map{{"bogus_col": 1}, {"name": "a"}}})
	rows, err = m.Storage.Find(m.GetTableName(), f)
	tt.NoError(err)
	tt.Equal(0, len(rows))
}

// 回归：时间型软删除下显式的 `deleted_at = nil` 应规范为 IS NULL，
// 不能生成恒不成立的 `deleted_at = NULL`。
func TestGetFilterDeletedAtNilNormalized(t *testing.T) {
	tt := zlsgo.NewTest(t)

	b := true
	softTime := true
	s := schema.Schema{
		Name: "audit_fix_del_nil",
		Table: schema.Table{
			Name: "audit_fix_del_nil",
		},
		Options: schema.Options{
			SoftDeletes:      &b,
			SoftDeleteIsTime: &softTime,
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, s)
	m := schemas.MustGet("audit_fix_del_nil")
	repo := m.Model().Repository()
	id1, err := repo.Insert(ztype.Map{"name": "a"})
	tt.NoError(err)
	_, err = repo.Insert(ztype.Map{"name": "b"})
	tt.NoError(err)
	_, err = repo.DeleteByID(id1)
	tt.NoError(err)

	// 顶层显式 nil 被规范为 IS NULL
	f := getFilter(m, Filter{DeletedAtKey: nil})
	_, hasIsNull := f[DeletedAtKey+" IS NULL"]
	tt.Equal(true, hasIsNull)
	_, hasRaw := f[DeletedAtKey]
	tt.Equal(false, hasRaw)

	// 嵌套分组内的显式 nil 同样被规范
	f = getFilter(m, Filter{"$OR": []ztype.Map{{DeletedAtKey: nil}}})
	group, ok := f["$OR"].([]ztype.Map)
	tt.Equal(true, ok)
	tt.Equal(1, len(group))
	_, hasIsNull = group[0][DeletedAtKey+" IS NULL"]
	tt.Equal(true, hasIsNull)

	// 生效：只返回未删除行
	rows, err := repo.Find(Filter{DeletedAtKey: nil})
	tt.NoError(err)
	tt.Equal(1, len(rows))
	tt.Equal("b", rows[0]["name"])
}

// 回归：getFilter 的深拷贝应避免解密等原地修改污染调用方原始过滤器。
func TestGetFilterNestedCloneProtectsCaller(t *testing.T) {
	tt := zlsgo.NewTest(t)

	enc := true
	s := schema.Schema{
		Name: "audit_fix_crypt",
		Table: schema.Table{
			Name: "audit_fix_crypt",
		},
		Options: schema.Options{
			CryptID:  &enc,
			Salt:     "test-salt",
			CryptLen: 8,
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, s)
	m := schemas.MustGet("audit_fix_crypt")

	encID, err := m.EnCryptID("42")
	tt.NoError(err)

	userFilter := ztype.Map{"$OR": []ztype.Map{{idKey: encID}}}
	f := getFilter(m, Filter(userFilter))

	// 解密只应发生在副本上
	ok := m.DeCrypt(f)
	tt.Equal(true, ok)
	got := userFilter["$OR"].([]ztype.Map)[0][idKey]
	tt.Equal(encID, got)

	decrypted := f["$OR"].([]ztype.Map)[0][idKey]
	tt.Equal(int64(42), decrypted)
}

// 回归：tab/多空格分隔的「字段 操作符」键不应被白名单静默丢弃，且能正确执行。
func TestGetFilterTabOperatorKey(t *testing.T) {
	tt := zlsgo.NewTest(t)

	s := schema.Schema{
		Name: "audit_fix_tab",
		Table: schema.Table{
			Name: "audit_fix_tab",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, s)
	m := schemas.MustGet("audit_fix_tab")
	repo := m.Model().Repository()
	_, err := repo.Insert(ztype.Map{"name": "a"})
	tt.NoError(err)
	_, err = repo.Insert(ztype.Map{"name": "b"})
	tt.NoError(err)

	for _, key := range []string{"name\t=", "name  ="} {
		f := getFilter(m, Filter{key: "b"})
		_, ok := f[key]
		tt.Equal(true, ok)

		rows, err := m.Storage.Find(m.GetTableName(), f)
		tt.NoError(err)
		tt.Equal(1, len(rows))
		tt.Equal("b", rows[0]["name"])
	}
}

// 回归：nil 值条件应生成 IS NULL 而非 `= NULL`（SQL 恒不成立）。
// 覆盖默认无操作符、显式 `=`、`!=`/`<>` 三种形态。
func TestGetFilterNilValueBecomesIsNull(t *testing.T) {
	tt := zlsgo.NewTest(t)

	b := true
	s := schema.Schema{
		Name: "audit_fix_nil_eq",
		Table: schema.Table{
			Name: "audit_fix_nil_eq",
		},
		Options: schema.Options{
			SoftDeletes:      &b,
			SoftDeleteIsTime: &b,
		},
		Fields: map[string]schema.Field{
			"name":  {Type: "string", Size: 80},
			"score": {Type: "int", Nullable: true},
		},
	}

	_, schemas := newTestSchemas(t, s)
	m := schemas.MustGet("audit_fix_nil_eq")
	repo := m.Model().Repository()

	id1, err := repo.Insert(ztype.Map{"name": "a"})
	tt.NoError(err)
	_, err = repo.Insert(ztype.Map{"name": "b", "score": 5})
	tt.NoError(err)

	// score IS NULL → 命中 a
	rows, err := repo.Find(Filter{"score": nil})
	tt.NoError(err)
	tt.Equal(1, len(rows))
	tt.Equal("a", rows[0]["name"])

	// score = nil（显式操作符）→ IS NULL
	f := getFilter(m, Filter{"score =": nil})
	rows, err = m.Storage.Find(m.GetTableName(), f)
	tt.NoError(err)
	tt.Equal(1, len(rows))
	tt.Equal("a", rows[0]["name"])

	// score != nil → IS NOT NULL → 命中 b
	f = getFilter(m, Filter{"score !=": nil})
	rows, err = m.Storage.Find(m.GetTableName(), f)
	tt.NoError(err)
	tt.Equal(1, len(rows))
	tt.Equal("b", rows[0]["name"])

	// score <> nil 同样语义
	f = getFilter(m, Filter{"score <>": nil})
	rows, err = m.Storage.Find(m.GetTableName(), f)
	tt.NoError(err)
	tt.Equal(1, len(rows))

	// 保留引用完整性：显式 nil 不应污染调用方 filter
	raw := Filter{"score": nil}
	f = getFilter(m, raw)
	rows, err = m.Storage.Find(m.GetTableName(), f)
	tt.NoError(err)
	tt.Equal(1, len(rows))

	_ = id1
}

// 回归：净化后为空的 $OR/$AND 元素保持恒假（fail-closed），
// 而原本就为空的 {} 子条件表示无条件，不产生恒假条件。
func TestGetFilterEmptyGroupElementSemantics(t *testing.T) {
	tt := zlsgo.NewTest(t)

	s := schema.Schema{
		Name: "audit_fix_empty_el",
		Table: schema.Table{
			Name: "audit_fix_empty_el",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, s)
	m := schemas.MustGet("audit_fix_empty_el")
	repo := m.Model().Repository()
	_, err := repo.Insert(ztype.Map{"name": "a"})
	tt.NoError(err)
	_, err = repo.Insert(ztype.Map{"name": "b"})
	tt.NoError(err)

	// $OR 内原本空 {} 元素 = 无条件，应忽略而非恒假 → 命中 name=b
	f := getFilter(m, Filter{"$OR": []ztype.Map{{}, {"name": "b"}}})
	group, ok := f["$OR"].([]ztype.Map)
	tt.Equal(true, ok)
	tt.Equal(1, len(group))
	rows, err := m.Storage.Find(m.GetTableName(), f)
	tt.NoError(err)
	tt.Equal(1, len(rows))
	tt.Equal("b", rows[0]["name"])

	// $AND 内原本空 {} 元素同样忽略
	f = getFilter(m, Filter{"$AND": []ztype.Map{{}, {"name": "b"}}})
	rows, err = m.Storage.Find(m.GetTableName(), f)
	tt.NoError(err)
	tt.Equal(1, len(rows))
	tt.Equal("b", rows[0]["name"])

	// $OR 全为空元素 → 无条件 → 命中全部（a、b）
	f = getFilter(m, Filter{"$OR": []ztype.Map{{}, {}}})
	rows, err = m.Storage.Find(m.GetTableName(), f)
	tt.NoError(err)
	tt.Equal(2, len(rows))
}

// 回归：自引用/环状 filter 不得导致 clone / sanitize / hasField 无限递归或栈溢出。
func TestGetFilterSelfReferenceNoPanic(t *testing.T) {
	tt := zlsgo.NewTest(t)

	s := schema.Schema{
		Name: "audit_fix_selfref",
		Table: schema.Table{
			Name: "audit_fix_selfref",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, s)
	m := schemas.MustGet("audit_fix_selfref")

	// 程序化构造的环：$OR 子条件引用自身 map
	selfRef := ztype.Map{"name": "a"}
	_ = selfRef
	cyclic := ztype.Map{"$OR": []interface{}{selfRef}}
	selfRef["$OR"] = cyclic["$OR"] // 形成引用环（值级）

	// clone 与净化都应在有限步内结束
	got := cloneFilterMap(cyclic)
	tt.Equal(true, len(got) > 0)

	// 深度受限的深层嵌套不报错也不崩溃
	deep := ztype.Map{"name": "a"}
	for i := 0; i < 200; i++ {
		deep = ztype.Map{"$AND": []ztype.Map{deep}}
	}
	f := getFilter(m, Filter(deep))
	tt.Equal(true, len(f) > 0)
}

// 回归：带空白前缀/后缀的 $OR/$AND 占位键在净化后仍能被正确解析。
func TestGetFilterPlaceholderWhitespaceKey(t *testing.T) {
	tt := zlsgo.NewTest(t)

	s := schema.Schema{
		Name: "audit_fix_ph_ws",
		Table: schema.Table{
			Name: "audit_fix_ph_ws",
		},
		Fields: map[string]schema.Field{
			"name": {Type: "string", Size: 80},
		},
	}

	_, schemas := newTestSchemas(t, s)
	m := schemas.MustGet("audit_fix_ph_ws")
	repo := m.Model().Repository()
	_, err := repo.Insert(ztype.Map{"name": "a"})
	tt.NoError(err)
	_, err = repo.Insert(ztype.Map{"name": "b"})
	tt.NoError(err)

	f := getFilter(m, Filter{"  $OR ": []ztype.Map{{"name": "b"}}})
	rows, err := m.Storage.Find(m.GetTableName(), f)
	tt.NoError(err)
	tt.Equal(1, len(rows))
	tt.Equal("b", rows[0]["name"])
}
