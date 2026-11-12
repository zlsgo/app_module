# REST API 模块

`restapi` 模块基于已注册的 `model.Store` 提供通配路由分发与标准 RESTful CRUD 能力，适合快速为数据模型暴露 HTTP 接口。

## 功能特性

- 🚀 通配路由自动分发（按 `{prefix}/{model}` / `{prefix}/{model}/{id}`）
- 📝 标准 CRUD 操作（GET/POST/PUT/PATCH/DELETE）
- 🔎 查询支持：字段选择、关联装载、排序与过滤
- 📄 分页查询与 `MaxPageSize` 限制
- 📁 文件上传封装（`HanderUpload`）
- 🛡️ 中间件支持（`Options.Middleware`）
- 🔄 响应钩子（`Options.ResponseHook`）
- ⚙️ 可配置方法/字段/排序/关联白名单与严格查询模式

## 模块结构

```
restapi/
├── base.go           # 基础控制器与通用处理
├── methods.go        # CRUD 方法实现
├── module.go         # 模块定义与路由注册
├── options.go        # 配置选项
├── relation.go       # 关联分页 helper（Deprecated）
├── restapi.go        # REST API 接口与 helper
└── upload.go         # 文件上传功能
```

## 快速开始

```go
package main

import (
    "github.com/sohaha/zlsgo/znet"
    "github.com/zlsgo/app_module/restapi"
    "github.com/zlsgo/app_module/model"
    "github.com/zlsgo/app_module/database"
    "github.com/zlsgo/app_core/service"
)

func main() {
    app := service.NewApp()(nil)

    dbMod := database.New()
    modelMod := model.New()
    restApiMod := restapi.New(func(o *restapi.Options) {
        o.Prefix = "/api/v1"
        o.Middleware = func(c *znet.Context) error {
            c.Next()
            return nil
        }
        o.ResponseHook = func(c *znet.Context, model, args, method string) bool {
            return true
        }
    })

    err := service.InitModule([]service.Module{dbMod, modelMod, restApiMod}, app)
    if err != nil {
        panic(err)
    }

    service.RunWeb(app)
}
```

> `restapi` 依赖 `model` 模块中已注册的 `Store`。请确保相关 Schema 已在 `model` 模块中完成注册与迁移。

## 自动路由规则

模块启动后会在 `Options.Prefix` 下注册通配路由 `/*`，并根据 HTTP Method 转发到对应的 `model.Store`：

| 方法      | 路径                    | 行为                                                              |
| --------- | ----------------------- | ----------------------------------------------------------------- |
| GET       | `{prefix}/{model}`      | 分页查询（读取 `page`/`pagesize`，受 `MaxPageSize` 限制），默认按 `id desc` |
| GET       | `{prefix}/{model}/{id}` | 查询单条                                                          |
| POST      | `{prefix}/{model}`      | 插入（返回 `{ "id": ... }`）                                      |
| PUT/PATCH | `{prefix}/{model}/{id}` | 更新（返回 `{ "total": ... }`）                                   |
| DELETE    | `{prefix}/{model}/{id}` | 删除（返回 `{ "total": ... }`）                                   |

说明：

- `{model}` 来自 `model` 模块已注册的 `Stores`。
- 仅允许上述两级路径，超出将返回错误。
- 出于安全考虑，已禁用全量查询（`{prefix}/{model}/*`），请使用分页接口。
- 模型或 id 不存在时，默认返回 404。
- 方法被禁止时，返回 405（包含 `Allow` 头）。

## 查询参数

| 参数    | 说明 |
| ------- | ---- |
| `page`  | 页码，严格模式下必须为正整数 |
| `pagesize` | 每页条数，严格模式下必须为正整数，且受 `MaxPageSize` 限制 |
| `fields` | 逗号分隔字段列表（如 `id,name` 或 `*`） |
| `with`  | 逗号分隔关联路径（如 `profile` / `profile.nickname`） |
| `order` | 逗号分隔排序字段（如 `name:asc,-id`） |
| `filter` | JSON 过滤对象（需 URL 编码） |

校验与约束：

- `fields` / `with` / `order` / `filter` 中的字段和关系会与 `Store.Schema()` 做严格校验，不存在即返回 4xx。
- 空值参数（如 `fields=` / `order=` / `with=` / `filter=`）会被视为无效并返回 400。
- `with` 与 `relations` 互斥，不能同时传。
- 开启 `RejectUnknownQuery` 时，未知 query 参数会被拒绝。
- 严格模式下按 key 原样匹配（大小写敏感），可通过 `AllowQueryKeys` 追加允许的自定义 key。
- 严格模式会对所有 HTTP 方法的 URL query 生效，并要求 query key 单值。

### `filter` 语法

`filter` 为 JSON 对象，支持以下操作符：

- 比较：`$eq` `$ne` `$gt` `$gte` `$lt` `$lte`
- 集合与匹配：`$in` `$nin` `$like` `$between`
- 空值：`$null` `$notnull`
- 逻辑组合：`$and` `$or`（数组）

示例：

```text
filter={"name":{"$like":"%foo%"},"age":{"$gte":18},"$or":[{"status":"active"},{"status":"pending"}]}
```

## 配置选项（Options）

| 字段 | 类型 | 说明 |
| ---- | ---- | ---- |
| `Prefix` | `string` | API 路由前缀 |
| `Middleware` | `znet.Handler` | 非 `nil` 时在通配路由前执行 |
| `ResponseHook` | `func(c *znet.Context, model, args, method string) bool` | 返回 `false` 时走 404（中断请求）；为 `nil` 时默认放行 |
| `MaxPageSize` | `int` | 分页最大值（默认 1000） |
| `AllowMethods` | `map[string]bool` | 仅允许显式为 `true` 的 HTTP 方法（大小写不敏感） |
| `AllowFields` | `map[string]bool` | 允许查询字段白名单 |
| `DefaultFields` | `[]string` | 未传 `fields` 时使用的默认字段 |
| `RequireFields` | `bool` | 未传 `fields` 时直接返回 400 |
| `AllowFilterFields` | `map[string]bool` | 允许过滤字段白名单（为空则回退到 `AllowFields`） |
| `AllowOrderFields` | `map[string]bool` | 允许排序字段白名单（为空则回退到 `AllowFields`） |
| `AllowRelations` | `map[string]bool` | 允许关联白名单（支持根关联名或完整路径） |
| `DefaultOrder` | `[]model.OrderByItem` | 默认排序（未传 `order` 时生效） |
| `ErrorHandler` | `znet.ErrHandlerFunc` | 自定义错误处理器（输出 ApiData） |
| `DisableErrorHandler` | `bool` | 禁用内置错误处理器 |
| `RejectUnknownQuery` | `bool` | 拒绝未知 query 参数（对所有方法的 URL query 生效） |
| `AllowQueryKeys` | `map[string]bool` | 严格模式下允许的额外 query key（区分大小写） |

## 错误响应格式

默认错误响应为 `ApiData`：

```json
{
  "data": null,
  "msg": "invalid filter",
  "code": 400
}
```

可通过 `Options.ErrorHandler` 自定义错误输出格式。

## 辅助函数

模块提供若干基于当前请求上下文 `*znet.Context` 的 helper，便于在自定义路由中复用：

- `Page(c, store, filter, fn)`
- `Find(c, store, filter, fn)`
- `FindById(c, store, id, fn)`
- `Insert(c, store, handler)` / `InsertMany(c, store, handler)`
- `UpdateById(c, store, id, handler)`
- `DeleteById(c, store, id, handler)`

## 文件上传

`HanderUpload(c, subDirName, ...)` 是对 `common.Upload` 的封装：

```go
file, err := restapi.HanderUpload(c, "images", func(o *common.UploadOption) {
    // 配置上传选项
})
```

## 关联关系

### 通过查询参数装载（推荐）

GET 请求可通过 `with`（或 `relations`）携带逗号分隔的关联路径（如 `profile` / `profile.nickname`），`restapi` 会将其写入 `CondOptions.Relations` 并在 `model` 层完成关系装载。

```text
GET /api/v1/user?with=profile,orders
```

关联路径需与 `Options.AllowRelations` 白名单一致（支持根关联名或完整路径），且 `with` 与 `relations` 互斥，不能同时传。

### 辅助函数（Deprecated）

`HanderPageRelation(c, store, filter, relations)` 是基于当前上下文的 legacy 分页 helper，仅包装 `Page(...)` 返回分页结果，不做额外关系装载。新代码请直接使用 `with`/`relations` 查询参数以启用关联装载。

## 最佳实践

- 通过 `AllowMethods`、`AllowFields`、`AllowFilterFields`、`AllowOrderFields`、`AllowRelations` 明确暴露范围，避免过度开放。
- 在生产环境启用 `RejectUnknownQuery` 并配置 `AllowQueryKeys`，减少非预期参数进入查询链路。
- 合理设置 `MaxPageSize`，避免大分页请求造成性能压力。
- 对外接口建议配合鉴权中间件与 `ResponseHook` 做统一访问控制。
- 关联装载优先使用 `with` 查询参数，避免使用已废弃的 `HanderPageRelation`。