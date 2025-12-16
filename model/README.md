# Model 模块

Model 是 app_module 中提供的数据建模与访问层。它围绕 Schema 驱动的模型定义，统一处理字段校验、自动迁移、关联装载、查询过滤与数据前后处理，内置 SQL 存储实现，亦支持通过接口接入自定义存储。

## 核心能力

- **Schema 驱动**：通过结构化 Schema 描述字段、关系、初始数据及视图元数据。
- **自动迁移**：启动阶段自动建表/增量同步、索引维护和初始数据写入。
- **统一 CRUD**：`Store` 封装插入、查询、更新、删除、分页等常见操作，并内建软删除/时间戳逻辑。
- **字段管线**：支持 JSON/布尔/时间格式的 Before/After 处理、字段加密（密码/MD5）以及枚举标签映射。
- **关联装载**：基于定义的 Relation 自动选择单条/合并/多条关联数据。
- **Schema API**：可选对外暴露模型元信息，便于管理端动态渲染。

## 目录结构

```
model/
├── action.go              // Store CRUD 入口
├── define.go              // Schema 运行时结构
├── field.go               // 字段解析与校验
├── hook/define.go         // 迁移阶段钩子事件
├── instance.go            // Schemas/Stores 容器
├── module.go / model.go   // 模块初始化入口
├── operation.go           // Filter/Store API 封装
├── parse.go               // Schema 解析与完善
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
- `Hook`：迁移阶段钩子，签名 `func(event hook.Event, data ...any) error`，事件枚举：
  - `hook.EventMigrationStart`
  - `hook.EventMigrationIndexDone`
  - `hook.EventMigrationDone`
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

解析阶段会自动将关系键名转换为 snake_case，并缓存关系键，供查询时匹配。

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

通过包级 `model.InsideOption` 控制：

- `InsideOption.OldColumnDelete()`：迁移时删除旧字段。
- `InsideOption.OldColumnRename()`：将旧字段重命名为 `__del__<name>`。
- `InsideOption.OldColumnIgnore()`：保留旧字段（默认）。
- `InsideOption.SoftDeleteIsTime(true)`：软删除字段使用时间类型，否则使用整型时间戳。

## Store API

`Module.MustGetStore(name)` / `Module.GetStore(name)` 返回 `*model.Store`，提供：

- 写入：`Insert`、`InsertMany`
- 查询：`Find`、`FindOne`、`FindCols`、`FindCol`、`Pages`
- 统计：`Count`、`Exists`
- 更新：`Update`、`UpdateMany`、`UpdateByID`
- 删除：`Delete`、`DeleteMany`、`DeleteByID`

Store 会自动：

- 在写入前执行字段 Before 管线与校验（`VerifiData`）。
- 自动填充 `created_at` / `updated_at` / `deleted_at` 等内建字段。
- 根据配置对 ID/字段执行加密或解密。

## 查询与过滤器

过滤条件使用 `model.Filter`（`type Filter ztype.Map`）或任意 `ztype.Map`：

| 语法         | 示例                                                                                                 | 说明               |
| ------------ | ---------------------------------------------------------------------------------------------------- | ------------------ |
| 等值         | `{"status": 1}`                                                                                      | 字段等于某值       |
| 范围         | `{"age >": 18, "age <=": 60}`                                                                        | 自动生成比较表达式 |
| IN/NOT IN    | `{"id": []int{1,2}}`、`{"id IN": []any{1,2}}`、`{"id NOT IN": []int{3}}`                             | 集合匹配           |
| LIKE         | `{"name LIKE": "%john%"}`                                                                            | 模糊查询           |
| NULL         | `{"deleted_at": nil}`, `{"deleted_at IS NOT NULL": true}`                                            | 空值判断           |
| BETWEEN      | `{"created_at BETWEEN": []string{"2024-01-01", "2024-12-31"}}`                                       | 区间               |
| 逻辑组合     | `{"$OR": ztype.Map{"status":1, "name LIKE":"%a%"}}`、`{"$AND": ...}`                                 | 嵌套 AND / OR      |
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
- `OrderBy`：`map[string]string`，支持 `ASC` / `DESC` / `1` / `-1`。
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

SQL 存储提供 `Transaction`：

```go
store := mod.MustGetStore("user")

err := store.Schema().Storage.(*model.SQL).Transaction(func(tx *model.SQL) error {
    id, err := tx.Insert("users", nil, ztype.Map{"username": "john"})
    if err != nil {
        return err
    }

    _, err = tx.Insert("user_logs", nil, ztype.Map{"user_id": id, "action": "register"})
    return err
})
```

也可以直接使用 `module.schemas.Storage().Transaction(...)` 对底层存储开启事务。

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