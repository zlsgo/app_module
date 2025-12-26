# Model 模块

Model 是 app_module 中提供的数据建模与访问层。它围绕 Schema 驱动的模型定义，统一处理字段校验、自动迁移、关联装载、查询过滤与数据前后处理，内置 SQL 存储实现，亦支持通过接口接入自定义存储。

## 核心能力

- **Schema 驱动**：通过结构化 Schema 描述字段、关系、初始数据及视图元数据。
- **自动迁移**：启动阶段自动建表/增量同步、索引维护和初始数据写入。
- **统一 CRUD**：`Store` 封装插入、查询、更新、删除、分页等常见操作，并内建软删除/时间戳逻辑。
- **Repository 模式**：泛型 `Repository[T, F, C, U]` 提供类型安全的数据访问层，过滤/创建/更新类型显式化。
- **链式查询**：`Query[T, F, C, U]` 提供流式 API 构建复杂查询，支持条件组合、排序、分页等。
- **类型安全过滤**：`QueryFilter` 构建函数（`Eq`、`In`、`Like` 等）与结构体过滤输入，统一到 Map 条件。
- **字段管线**：支持 JSON/布尔/时间格式的 Before/After 处理、字段加密（密码/MD5）以及枚举标签映射。
- **关联装载**：基于定义的 Relation 自动选择单条/合并/多条关联数据。
- **CRUD 钩子**：支持 Insert/Update/Delete 前后的钩子事件。
- **Schema API**：可选对外暴露模型元信息，便于管理端动态渲染。

## 目录结构

```
model/
├── action_*.go            // Store CRUD 操作（按功能拆分）
├── define.go              // Schema 运行时结构
├── field.go               // 字段解析与校验
├── hook/define.go         // 钩子事件定义（迁移与 CRUD）
├── instance.go            // Schemas/Stores 容器
├── mapper.go              // Mapper 接口与实现
├── module.go / model.go   // 模块初始化入口
├── operation.go           // Filter/Store API 封装
├── parse.go               // Schema 解析与完善
├── pool.go                // CondOptions 对象池
├── query_filter.go        // QueryFilter 接口与构建函数
├── query.go               // 链式查询 Query Builder
├── repository.go          // 泛型 Repository 层
├── storage.go / storage_sql*.go // 存储接口与 SQL 实现
├── utils.go               // 初始化辅助
├── valid.go               // 写入数据校验
├── views.go               // 视图元信息解析
└── var.go                 // 常用工具 (分页)
```

## 快速集成

```go
import (
    "github.com/zlsgo/app_module/model"
    "github.com/zlsgo/app_module/model/schema"
)

mod := model.New(func(o *model.Options) {
    // 选择存储（默认会从 DI 中解析 *zdb.DB 并构造 SQL 存储）
    // o.SetStorageer = func() (model.Storageer, error) { ... }

    // 表名前缀（SQL 存储会用于生成表名）
    o.Prefix = "model_"

    // 自动注册 Schema API（可选）
    o.SchemaApi = "/api/v1/schema"
    o.SchemaMiddleware = func() []znet.Handler { return []znet.Handler{middleware.Auth()} }

    // 全局默认 Schema 行为
    o.SchemaOptions = model.SchemaOptions{
        DisabledMigrator: false,
        SoftDeletes:      true,
        Timestamps:       true,
        CryptID:          false,
        SoftDeleteIsTime: false,
        OldColumn:        model.DealOldColumnNone,
    }

    // 直接在代码中定义 Schema
    user := schema.New("user", "users")
    _ = user.AddField("username", schema.Field{Label: "用户名", Type: schema.String, Size: 50, Unique: true})
    _ = user.AddField("email",    schema.Field{Label: "邮箱",   Type: schema.String, Size: 100, Unique: true})
    _ = user.AddField("password", schema.Field{Label: "密码",   Type: schema.String, Size: 255, Options: schema.FieldOption{Crypt: "password"}})
    o.Schemas.Append(user)

    // 或者指定目录，自动解析 JSON Schema 文件
    // o.SchemaDir = "./schema"

    // 动态模型（延迟加载）
    // o.SetAlternateModels = func() ([]*model.Store, error) { ... }
})

// 通过 service.InitModule 将模块注册进应用生命周期，自动执行迁移并注入依赖
```

> **依赖注入**：模块会从 DI 中解析 `*zdb.DB`，若不存在需通过 `SetStorageer` 手动提供 `Storageer` 实现。

## 定义 Schema

### Schema 结构

`schema.Schema` 由以下部分组成：

- **基本信息**：`Name`（别名）、`Table{Name, Comment}`（表定义）。
- **Fields**：`map[string]schema.Field`，使用 `AddField` 添加，键为字段名。
- **Relations**：`map[string]schema.Relation`，描述模型间关联。
- **Extend**：自定义扩展信息（`ztype.Map`），常用于视图描述。
- **Values**：`[]ztype.Map`，表首次初始化时插入的默认数据。
- **Options**：`schema.Options`，控制模型级行为。

### 通过结构体定义

```go
type User struct {
    schema.Meta `name:"user" table:"users" comment:"用户表" options:"timestamps,soft_deletes,crypt_id" low_fields:"password" fields_sort:"id|name"`
    ID       uint   `json:"id"`
    Name     string `json:"name" field:"size:100,label:用户名"`
    Email    string `json:"email" field:"size:200,unique"`
    Status   int8   `json:"status" field:"default:1,enum:1=active|2=disabled"`
    Password string `json:"password" field:"crypt:password,readonly"`
    Profile  *Profile `relation:"type:single,schema:profiles,foreign:profile_id,fields:id|name"`
}

user := schema.NewFromStruct[User]("", "")
o.Schemas.Append(user)
```

运行时结构体：

```go
user := schema.NewFromStructValue("", User{})
o.Schemas.Append(user)
```

字段 tag 规则：

- 使用 `json` 定义字段名，缺省为字段名 snake_case。
- 使用 `field` 定义字段参数：`size`/`default`/`label`/`nullable`/`unique`/`index`/`crypt`/`format`/`enum`/`valid`/`disable_migration`。
- `enum`/`valid` 列表使用 `|` 分隔，`valid` 采用 `method=arg@message` 格式。
- 使用匿名嵌入 `schema.Meta` 定义表元信息：`name`/`table`/`comment`/`options`/`low_fields`/`fields_sort`/`crypt_salt`/`crypt_len`。
- 使用 `relation` 定义关联字段，列表参数用 `|` 分隔：`type`/`schema`/`foreign`/`schema_key`/`fields`/`nullable`/`pivot_*`/`cascade`。

### 字段 Field

`schema.Field` 关键属性：

| 属性               | 说明                                                                                    |
| ------------------ | --------------------------------------------------------------------------------------- |
| `Label`            | 字段显示名称，默认使用键名                                                              |
| `Type`             | 字段类型，使用 `schema.Bool / Int / Uint / Float / String / Text / JSON / Time / Bytes` |
| `Size`             | 对字符串/数字长度的约束，自动用于校验                                                   |
| `Nullable`         | 是否允许为 `NULL`                                                                       |
| `Default`          | 默认值，存在时写入前会将字段设为可空                                                    |
| `Unique` / `Index` | 支持布尔或详细配置，自动生成索引                                                        |
| `Validations`      | 额外校验规则，`[]schema.Validations{Method, Args, Message}`                             |
| `Options`          | `schema.FieldOption`，见下                                                              |

`FieldOption` 支持：

- `Crypt`：对写入值执行加密处理，内置 `"md5"` 与 `"password"`（bcrypt，大小写不敏感）。
- `Enum`：`[]FieldEnum{Value, Label}`，生成下拉枚举并在查询结果中附加 `<field>_label`。
- `FormatTime`：时间格式化模板（`date|Y-m-d H:i:s` 等）。
- `IsArray`：针对 JSON 字段控制数组/对象期望格式。
- `ReadOnly`：更新操作会自动过滤此字段。
- `DisableMigration`：字段不会参与自动迁移。

字段在解析时会为 JSON、布尔、时间类型自动挂载 Before/After 处理器，实现写入前转换与读取后反序列化。

### 校验规则

`field.Validations` 通过 `Method` 指定规则：

- 字符串：`minLength`、`maxLength`、`regex`、`mail`、`mobile`、`url`、`ip`
- 数值：`min`、`max`
- 枚举：`enum`（支持字符串/浮点/整型 slice）
- JSON：`json`

内置校验引擎会根据字段类型自动追加必填/长度限制等规则，验证失败返回 `zerror.InvalidInput`。

### 模型 Options

`schema.Options` 覆盖模型级行为：

- `DisabledMigrator`：禁用自动迁移。
- `SoftDeletes`：启用软删除字段（默认 `deleted_at`）。
- `Timestamps`：自动维护 `created_at` / `updated_at`。
- `CryptID`：对主键 ID 进行 Hash 加密（使用 HashID）。
- `Hook`：模型生命周期钩子，签名 `func(event hook.Event, data ...any) error`，事件枚举：
  - 迁移事件：
    - `hook.EventMigrationStart`
    - `hook.EventMigrationIndexDone`
    - `hook.EventMigrationDone`
  - CRUD 事件：
    - `hook.EventBeforeInsert` / `hook.EventAfterInsert`
    - `hook.EventBeforeUpdate` / `hook.EventAfterUpdate`
    - `hook.EventBeforeDelete` / `hook.EventAfterDelete`
- `Salt` / `CryptLen`：ID 加密参数。
- `LowFields`：在 Schema API 响应中隐藏的字段列表。
- `FieldsSort`：字段排序优先级。

> 模块级 `Options.SchemaOptions` 作为默认值，单个 Schema 可以在 Options 中覆盖。

### 关系 Relation

`schema.Relation` 属性：

- `Type`：`single`（一对一）、`single_merge`（一对一并合并字段）、`many`（一对多）。
- `Schema`：目标 Schema 别名。
- `ForeignKey`：本表外键字段列表。
- `SchemaKey`：目标表匹配字段列表（需与 ForeignKey 数量一致）。
- `Fields`：关联查询时要加载的字段，留空默认全部或 `*`。
- `Filter`：附加筛选条件。
- `Nullable`：没有匹配数据时返回空对象/数组。

解析阶段会自动将关系键名转换为 snake_case，供查询时匹配。

### Extend 视图

`Schema.Extend["views"]` 可定义 `lists`、`info` 等视图：

- `fields`：视图字段列表。
- `layouts`：字段布局描述（保留结构不做解析）。
- `disabled`：禁用视图。
- `title`：视图标题。

运行时可通过 `Schema.GetViews()`、`Schema.GetViewFields(view)` 获取解析结果。结果会自动补齐主键及必要字段，并针对加密 ID 调整字段类型。

### 初始数据 Values

`Schema.Values` 用于初始化表数据：

- 表首次迁移成功后或表为空时写入。
- 支持在数据中指定主键（若启用 CryptID 会在写入前解密）。

## 存储与自动迁移

- **Storageer 接口**：定义了 `Find/Pages/Insert/Update/Delete/Migration` 等操作，默认实现为 SQL 存储（`model.NewSQL`）。
- **依赖注入**：模块会从 DI 获取 `*zdb.DB` 并生成 SQL Storage；也可通过 `Options.SetStorageer` 注入自定义实现（如 NoSQL）。
- **SchemaDir**：指定目录时会递归读取 JSON 文件并反序列化为 Schema。
- **自动迁移**：`Module.Done` 时调用 `initModels` → `Migration.Auto`。
  - 新表：直接创建并执行初始值写入。
  - 老表：根据字段差集添加列、可选删除/重命名旧列。
  - 索引：迁移完成后统一创建。
  - 初始值：若表为空或首次创建，写入 `Schema.Values`。

### 旧字段策略 & 软删除

通过 `SchemaOptions` / `schema.Options` 控制：

- `SchemaOptions.OldColumn`：默认迁移策略（`DealOldColumnNone`/`DealOldColumnDelete`/`DealOldColumnRename`）。
- `SchemaOptions.SoftDeleteIsTime`：默认软删除字段使用时间类型。
- `schema.Options.SoftDeleteIsTime`：单个 Schema 覆盖软删除字段类型。

## Store API

`Module.MustGetStore(name)` / `Module.GetStore(name)` 返回 `*model.Store`，提供：

- 写入：`Insert`、`InsertMany`
- 查询：`Find`、`FindOne`、`FindCols`、`FindCol`、`Pages`
- 统计：`Count`、`Exists`
- 更新：`Update`、`UpdateMany`、`UpdateByID`
- 删除：`Delete`、`DeleteMany`、`DeleteByID`

Store 的写入与过滤参数支持 `ztype.Map`/`map[string]any`/结构体输入，结构体需提供 `z` 或 `json` tag；更新建议使用 `omitempty` 或指针字段避免覆盖零值。

Store 会自动：

- 在写入前执行字段 Before 管线与校验（`VerifiData`）。
- 自动填充 `created_at` / `updated_at` / `deleted_at` 等内建字段。
- 根据配置对 ID/字段执行加密或解密。

## Repository 模式

`Repository[T, F, C, U]` 是基于 `Store` 的泛型封装层，提供类型安全的数据访问，其中 F 为过滤类型、C 为创建类型、U 为更新类型。

### 创建 Repository

```go
store := mod.MustGetStore("user")

// 方式一：使用工厂函数
mapRepo := model.NewMapRepository(store)
structRepo := model.NewStructRepository[User, UserFilter, UserCreate, UserPatch](store)

// 方式二：从 Store 直接获取（推荐）
mapRepo := store.Repository() // 返回 *Repository[ztype.Map, QueryFilter, ztype.Map, ztype.Map]
structRepo := model.Repo[User, UserFilter, UserCreate, UserPatch](store) // 返回 *Repository[User, UserFilter, UserCreate, UserPatch]

// 方式三：自定义 Mapper
customRepo := model.NewRepository[User, UserFilter, UserCreate, UserPatch](store, customMapper)
```

结构体映射示例：

```go
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}

type UserCreate struct {
    Username string `json:"username"`
    Email    string `json:"email"`
}

type UserPatch struct {
    Username string `json:"username,omitempty"`
    Email    string `json:"email,omitempty"`
    Status   int    `json:"status,omitempty"`
}

type UserFilter struct {
    Username string `json:"username,omitempty"`
    Status   int    `json:"status,omitempty"`
}
```

### Repository API

> Repository 的 filter 参数由 F 决定：结构体过滤直接传 F；若 F 为 QueryFilter，可用 Q(...) / Eq / ID / Filter。

```go
// 查询
users, err := repo.Find(UserFilter{Status: 1})
user, err := repo.FindOne(UserFilter{Username: "john"})
user, err := repo.First(UserFilter{Status: 1})  // FindOne 别名
user, err := repo.FindByID(1)
users, err := repo.FindByIDs([]any{1, 2, 3})
all, err := repo.All()

// 分页
pageData, err := repo.Pages(1, 20, UserFilter{})

// 统计
count, err := repo.Count(UserFilter{Status: 1})
exists, err := repo.Exists(UserFilter{Username: "john"})

// 写入
id, err := repo.Insert(UserCreate{Username: "john", Email: "john@example.com"})
id, err := repo.InsertMany([]UserCreate{...})

// 更新
affected, err := repo.Update(UserFilter{Username: "john"}, UserPatch{Username: "john"})
affected, err := repo.UpdateByID(1, UserPatch{Email: "john@example.com"})
affected, err := repo.UpdateMany(UserFilter{Status: 1}, UserPatch{Status: 2})
affected, err := repo.UpdateByIDs([]any{1, 2}, UserPatch{Status: 2})

// 删除
affected, err := repo.Delete(UserFilter{Username: "john"})
affected, err := repo.DeleteByID(1)
affected, err := repo.DeleteMany(UserFilter{Status: 2})
affected, err := repo.DeleteByIDs([]any{1, 2})

// 辅助方法
store := repo.Store()    // 获取底层 *Store
schema := repo.Schema()  // 获取 *Schema
```

### 泛型用法

```go
type UserFilter struct {
    Status int `json:"status,omitempty"`
}

type UserCreate struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Status   int    `json:"status"`
}

type UserPatch struct {
    Email  string `json:"email,omitempty"`
    Status int    `json:"status,omitempty"`
}

repo := model.Repo[User, UserFilter, UserCreate, UserPatch](store)
users, err := repo.Find(UserFilter{Status: 1})
id, err := repo.Insert(UserCreate{Username: "john", Email: "john@example.com", Status: 1})
_, err = repo.Update(UserFilter{Status: 1}, UserPatch{Status: 2})
pageData, err := repo.Query().Where("status", 1).Pages(1, 20)
```

需要使用 `Eq/Or/ID` 等 QueryFilter 构建时，可将 F 指定为 `model.QueryFilter`，或直接使用 `store.Repository()` 返回的 Map Repository。

### 泛型查询函数

```go
store := mod.MustGetStore("user")
schema := store.Schema()

user, err := model.FindOne[User](store, model.Q(UserFilter{Status: 1}))
page, err := model.Pages[User](schema, 1, 20, model.Q(UserFilter{}))
ages, err := model.FindCols[int](store, "age", model.Q(UserFilter{Status: 1}))
age, ok, err := model.FindCol[int](store, "age", model.Q(UserFilter{Status: 1}))
```

### QueryFilter 构建函数

提供类型安全的条件构建（支持 Map/Struct 输入）：

| 函数        | 说明         | 示例                                    |
| ----------- | ------------ | --------------------------------------- |
| `Q(any)`    | 原始条件     | `Q(UserFilter{Status: 1})`              |
| `ID(v)`     | 主键匹配     | `ID(1)`                                 |
| `Eq(f, v)`  | 等于         | `Eq("name", "john")`                    |
| `Ne(f, v)`  | 不等于       | `Ne("status", 0)`                       |
| `Gt(f, v)`  | 大于         | `Gt("age", 18)`                         |
| `Ge(f, v)`  | 大于等于     | `Ge("age", 18)`                         |
| `Lt(f, v)`  | 小于         | `Lt("age", 60)`                         |
| `Le(f, v)`  | 小于等于     | `Le("age", 60)`                         |
| `In(f, v)`  | IN 集合      | `In("id", []int{1,2,3})`                |
| `NotIn`     | NOT IN       | `NotIn("status", []int{0, -1})`         |
| `Like`      | 模糊匹配     | `Like("name", "%john%")`                |
| `Between`   | 区间         | `Between("age", 18, 60)`                |
| `IsNull`    | 为空         | `IsNull("deleted_at")`                  |
| `IsNotNull` | 不为空       | `IsNotNull("email")`                    |
| `And(...)`  | 条件组合 AND | `And(Q(UserFilter{Status: 1}), Gt("b", 2))` |
| `Or(...)`   | 条件组合 OR  | `Or(Eq("status", 1), Eq("status", 2))`  |

## Query Builder（链式查询）

`Query[T, F, C, U]` 提供流畅的链式 API 构建复杂查询：

```go
repo := model.NewMapRepository(store)

// 链式查询
users, err := repo.Query().
    Where("status", 1).
    WhereLike("name", "%john%").
    WhereGt("age", 18).
    Select("id", "username", "email").
    OrderBy("created_at", "DESC").
    Limit(10).
    Find()

// 分页查询
pageData, err := repo.Query().
    Where("status", 1).
    OrderByDesc("id").
    Pages(1, 20)

// 单条查询
user, err := repo.Query().
    WhereID(1).
    FindOne()

// 统计
count, err := repo.Query().
    Where("status", 1).
    Count()

// 更新
affected, err := repo.Query().
    Where("status", 0).
    Update(ztype.Map{"status": 1})

// 删除
affected, err := repo.Query().
    Where("status", -1).
    Delete()
```

### Query 方法

| 方法                        | 说明                       |
| --------------------------- | -------------------------- |
| `Where(field, value)`       | 等值条件                   |
| `WhereFilter(filter)`       | 自定义过滤结构 F           |
| `WhereID(id)`               | 主键条件                   |
| `WhereIn(field, values)`    | IN 条件                    |
| `WhereNot(field, value)`    | 不等于                     |
| `WhereGt/Ge/Lt/Le`          | 比较条件                   |
| `WhereLike(field, pattern)` | 模糊匹配                   |
| `WhereBetween(field, a, b)` | 区间条件                   |
| `WhereNull/WhereNotNull`    | 空值判断                   |
| `OrWhere(filters...)`       | OR 条件组（F）             |
| `Select(fields...)`         | 指定返回字段               |
| `OrderBy(field, dir)`       | 排序（默认 ASC）           |
| `OrderByDesc(field)`        | 降序排序                   |
| `GroupBy(fields...)`        | 分组                       |
| `Limit(n)` / `Offset(n)`    | 限制与偏移                 |
| `WithRelation(names...)`    | 加载关联                   |
| `Find()` / `FindOne()`      | 执行查询                   |
| `First()`                   | 等同 FindOne               |
| `Pages(page, pagesize)`     | 分页查询                   |
| `Count()` / `Exists()`      | 统计                       |
| `Update(data)`              | 执行更新                   |
| `Delete()`                  | 执行删除                   |

## 批量操作

`Repository[T, F, C, U]` 提供批量操作方法，适用于大数据量场景：

```go
repo := model.NewMapRepository(store)

// 批量插入（自动分批，默认每批 1000 条）
ids, err := repo.BatchInsert(ztype.Maps{
    {"username": "user1"},
    {"username": "user2"},
    // ... 大量数据
})

// 自定义批量大小
ids, err := repo.BatchInsert(data, model.BatchSize(500))

// 事务内批量插入（全部成功或全部回滚）
ids, err := repo.BatchInsertTx(data)

// 批量更新
affected, err := repo.BatchUpdate(model.Eq("status", 0), ztype.Map{"status": 1})

// 批量删除
affected, err := repo.BatchDelete(model.Lt("created_at", expireTime))
```

| 方法                         | 说明                              |
| ---------------------------- | --------------------------------- |
| `BatchInsert(data []C, opts...)`   | 分批插入，返回所有插入的 ID       |
| `BatchInsertTx(data []C, opts...)` | 事务内分批插入，失败时全部回滚    |
| `BatchUpdate(filter, data, opts...)` | 分批更新匹配的记录          |
| `BatchDelete(filter, opts...)` | 分批删除匹配的记录                |

> 默认批量大小 `DefaultBatchSize = 1000`，可通过 `BatchSize(n)` 调整。

## 查询与过滤器

当 F 为 `QueryFilter` 时可直接使用 QueryFilter 构建条件，结构体/Map 需通过 `Q(...)` 或 `Filter` 转换：

| 语法         | 示例                                                                                                 | 说明               |
| ------------ | ---------------------------------------------------------------------------------------------------- | ------------------ |
| 等值         | `Filter{"status": 1}`                                                                                | 字段等于某值       |
| 范围         | `Filter{"age >": 18, "age <=": 60}`                                                                  | 自动生成比较表达式 |
| IN/NOT IN    | `Filter{"id": []int{1,2}}`、`Filter{"id IN": []any{1,2}}`、`Filter{"id NOT IN": []int{3}}`           | 集合匹配           |
| LIKE         | `Filter{"name LIKE": "%john%"}`                                                                      | 模糊查询           |
| NULL         | `Filter{"deleted_at": nil}`, `Filter{"deleted_at IS NOT NULL": true}`                                | 空值判断           |
| BETWEEN      | `Filter{"created_at BETWEEN": []string{"2024-01-01", "2024-12-31"}}`                                 | 区间               |
| 逻辑组合     | `Filter{"$OR": ztype.Map{"status":1, "name LIKE":"%a%"}}`、`Filter{"$AND": ...}`                     | 嵌套 AND / OR      |
| 自定义表达式 | `Filter{}.Cond(func(c *builder.BuildCond) string { return c.Expr("JSON_CONTAINS(tags, '"foo"')") })` | 直接注入 SQL 片段  |

其他特性：

- key 中包含空格会被视为 `字段 + 操作符`，其余字符串作为参数。
- 递归深度限制为 50，避免过深嵌套。
- SoftDeletes 启用时，查询会自动追加 `deleted_at` 过滤。
- CryptID 启用时，传入/返回的 `id` 会在查询前后自动解密/加密。
- `Filter.Set()`/`Filter.Get()` 辅助构建条件。

## CondOptions 与关联装载

查询方法可接受 `func(*model.CondOptions)` 定制：

- `Fields`：返回字段列表。包含 `关系` 或 `关系.字段` 时会触发关联查询。
- `OrderBy`：`[]OrderByItem`，每项包含 `Field` 和 `Direction`（`ASC` / `DESC`）。
- `GroupBy`：字段分组。
- `Join`：手动追加 `StorageJoin`（表名、别名、表达式）。
- `Limit` / `Offset`：限制条数与偏移量。

关联装载实现：

- `single`：在结果中附加对象。
- `single_merge`：把子对象字段合并到父记录。
- `many`：附加数组结果。
- 若视图字段中包含关联字段，会自动补齐外键字段并在结果返回后移除。

## 分页返回结构

`store.Pages` 返回 `*model.PageData`：

- `Items`：`ztype.Maps` 列表。
- `Page`：`model.PageInfo`，继承 `zdb.Pages`（包含 Page、PageSize、Total 等）。
- `Map` 方法支持对结果逐条加工，默认并发度与分页大小一致。

## 字段处理与写入校验

- Before/After 管线通过 `Field.Before` / `Field.After` 触发：
  - `bool`：0/1 与布尔互转。
  - `json`/`jsons`：对象/数组 JSON 编解码。
  - `date|<format>`：日期字符串与时间戳转换。
- 字段加密：`FieldOption.Crypt` 对写入值执行 MD5 或 bcrypt。
- 自动字段：
  - `Timestamps`：写入/更新自动填充 `created_at` / `updated_at`。
  - `SoftDeletes`：删除时更新 `deleted_at`（时间或时间戳）。
  - `CryptID`：写入/返回时自动处理主键。
- 校验流程：`VerifiData` 根据字段定义及 Validations 动态校验，失败返回错误并中断写入。

## 事务

存储接口提供 `Transaction`，推荐使用 `Repository.Tx` 进行事务操作：

```go
repo := model.NewMapRepository(store)

err := repo.Tx(func(txRepo *model.Repository[ztype.Map, model.QueryFilter, ztype.Map, ztype.Map]) error {
    id, err := txRepo.Insert(ztype.Map{"username": "john"})
    if err != nil {
        return err
    }
    // 同一事务内的其他操作
    return nil
})
```

也可以直接使用底层存储的 `Transaction` 方法：

```go
err := store.Schema().Storage.Transaction(func(txStorage model.Storageer) error {
    // 使用 txStorage 进行事务操作
    return nil
})
```

## Schema API 与视图元数据

设置 `Options.SchemaApi` 后会注册 `schemaController`：

- `GET /api/.../schema`：返回所有 Schema 的名称、注释、字段列表（过滤 `LowFields`）。
- `GET /api/.../schema/{name}`（内部使用）：访问单个 Schema。

视图元数据会在响应中返回，便于前端/管理端动态生成列表与详情页面。

## 辅助方法

- `model.Common.VarPages(c *znet.Context)`：从请求参数解析 `page` / `pagesize`。
- `model.GetEngine[*zdb.DB](store)`：获取底层 `*zdb.DB` 引擎。
- `store.Schema(Storageer)`：临时替换存储实现（例如事务内使用新的 SQL 实例）。

## 最佳实践

1. **Schema 规划**：避免在 Schema 中直接使用保留字段名（`id`, `created_at`, `updated_at`, `deleted_at`）。
2. **索引管理**：在字段定义阶段设置 `Unique` / `Index`，让迁移统一维护索引。
3. **验证与脱敏**：使用 `Validations` 与 `FieldOption.Crypt`，通过 `LowFields` 和 `Enum` Label 控制输出。
4. **关联查询**：只在 `Fields` 中声明必要字段，避免加载全部关联造成性能压力。
5. **迁移策略**：在生产环境慎用 `OldColumnDelete`，推荐先 `OldColumnRename` 观察无误再删除。
6. **自定义存储**：实现 `Storageer` 接口即可扩展至其他数据库/服务，同时复用 Schema 与 Store 能力。

通过以上机制，Model 模块提供了统一、可配置且易扩展的数据访问能力，适合构建后台管理、业务中台等需要快速迭代的数据模型场景。
