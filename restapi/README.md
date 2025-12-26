# REST API 模块

REST API 模块提供了完整的 RESTful API 功能，包括自动路由、数据操作、文件上传和关联关系处理。

## 功能特性

- 🚀 通配路由自动分发（按 `{prefix}/{model}` / `{prefix}/{model}/{id}`）
- 📝 标准 CRUD 操作（GET/POST/PUT/PATCH/DELETE）
- 📁 文件上传封装（`HanderUpload`）
- 🛡️ 中间件支持（`Options.Middleware`）
- 🔄 响应钩子（`Options.ResponseHook`）
- 🔎 查询字段/关联/排序（`fields` / `with` / `order`）

## 模块结构

```
restapi/
├── base.go           # 基础控制器
├── methods.go        # CRUD 方法
├── module.go         # 模块定义
├── options.go        # 配置选项
├── relation.go       # 关联关系处理
├── restapi.go        # REST API 接口
└── upload.go         # 文件上传功能
```

## 快速开始

### 基本使用

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
    // 初始化应用
    app := service.NewApp()(nil)

    // 数据库模块
    dbMod := database.New()

    // 模型模块
    modelMod := model.New()

    // REST API 模块
    restApiMod := restapi.New(func(o *restapi.Options) {
        o.Prefix = "/api/v1"
        o.Middleware = func(c *znet.Context) error {
            c.Next()
            return nil
        }
        o.ResponseHook = func(c *znet.Context, model, args, method string) bool { return true }
    })

    // 注册全部模块
    err := service.InitModule([]service.Module{dbMod, modelMod, restApiMod}, app)
    if err != nil {
        panic(err)
    }

  

    // 启动服务
    service.RunWeb(app)
}
```

## API 接口

### 自动路由规则

模块启动后会在 `Options.Prefix` 下注册一个通配路由 `/*`，并根据 HTTP Method 转发到 `model.Store`：

| 方法      | 路径                    | 行为                                                              |
| --------- | ----------------------- | ----------------------------------------------------------------- |
| GET       | `{prefix}/{model}`      | 分页查询（读取 `page`/`pagesize`，最大 `MaxPageSize`），默认按 `id desc` |
| GET       | `{prefix}/{model}/{id}` | 查询单条                                                          |
| POST      | `{prefix}/{model}`      | 插入（返回 `{ "id": ... }`）                                      |
| PUT/PATCH | `{prefix}/{model}/{id}` | 更新（返回 `{ "total": ... }`）                                   |
| DELETE    | `{prefix}/{model}/{id}` | 删除（返回 `{ "total": ... }`）                                   |

> **注意**：仅允许上述两级路径，超出将返回错误。出于安全考虑，已禁用全量查询（`{prefix}/{model}/*`），请使用分页接口。

其中 `{model}` 来自 `model` 模块已注册的 `Stores`（`model.Store`）。
当模型或 id 不存在时，默认返回 404。
当方法被禁止时，返回 405（包含 `Allow` 头）。

### 查询参数

- `fields`: 逗号分隔字段列表（如 `id,name` 或 `*`）
- `with`: 逗号分隔关联路径（如 `profile` / `profile.nickname`）
- `order`: 逗号分隔排序字段（如 `name:asc,-id`）
- `filter`: JSON 过滤对象（需 URL 编码）

`fields` / `with` / `order` / `filter` 中的字段和关系会与 `Store.Schema()` 做严格校验，不存在即返回 4xx。
空值参数（如 `fields=` / `order=` / `with=` / `filter=`）会被视为无效并返回 400。
`with` 与 `relations` 互斥，不能同时传。
开启 `RejectUnknownQuery` 时，未知 query 参数会被拒绝。
严格模式下按 key 原样匹配（大小写敏感），可通过 `AllowQueryKeys` 追加允许的自定义 key。
严格模式会对所有 HTTP 方法的 URL query 生效。
严格模式下要求 query key 单值，且 `page`/`pagesize` 必须为正整数。

`filter` 支持操作符：

- `$eq` `$ne` `$gt` `$gte` `$lt` `$lte`
- `$in` `$nin` `$like` `$between`
- `$null` `$notnull`
- `$and` `$or`（数组）

示例：

```text
filter={"name":{"$like":"%foo%"},"age":{"$gte":18},"$or":[{"status":"active"},{"status":"pending"}]}
```

### Options

`Options` 与实现保持一致：

- `Prefix string`
- `Middleware znet.Handler`：若不为 `nil`，会在通配路由前 `r.Use(Middleware)`
- `ResponseHook func(c *znet.Context, model, args, method string) bool`：返回 `false` 时会走 `404`（中断请求）；为 `nil` 时默认放行
- `MaxPageSize int`：分页最大值（默认 1000）
- `AllowMethods map[string]bool`：仅允许显式为 `true` 的 HTTP 方法（如 `GET`/`POST`，大小写不敏感）
- `AllowFields map[string]bool`：允许查询字段白名单
- `DefaultFields []string`：未传 `fields` 时使用的默认字段
- `RequireFields bool`：未传 `fields` 时直接返回 400
- `AllowFilterFields map[string]bool`：允许过滤字段白名单（为空则回退到 `AllowFields`）
- `AllowOrderFields map[string]bool`：允许排序字段白名单（为空则回退到 `AllowFields`）
- `AllowRelations map[string]bool`：允许关联白名单（支持根关联名或完整路径）
- `DefaultOrder []model.OrderByItem`：默认排序（未传 `order` 时生效）
- `ErrorHandler znet.ErrHandlerFunc`：自定义错误处理器（输出 ApiData）
- `DisableErrorHandler bool`：禁用内置错误处理器
- `RejectUnknownQuery bool`：拒绝未知 query 参数（对所有方法的 URL query 生效）
- `AllowQueryKeys map[string]bool`：严格模式下允许的额外 query key（区分大小写）

### 错误响应格式

默认错误响应为 `ApiData`：

```json
{
  "data": null,
  "msg": "invalid filter",
  "code": 400
}
```

### 辅助函数

该模块还提供若干直接调用的 helper（基于当前请求上下文 `*znet.Context`）：

- `Page(c, store, filter, fn)`
- `Find(c, store, filter, fn)`
- `FindById(c, store, id, fn)`
- `Insert(c, store, handler)` / `InsertMany(c, store, handler)`
- `UpdateById(c, store, id, handler)`
- `DeleteById(c, store, id, handler)`

### 文件上传

`HanderUpload(c, subDirName, ...)` 是对 `common.Upload` 的封装：

- `restapi.HanderUpload(c, "images", func(o *common.UploadOption){ ... })`

### 关联关系

当前 `HanderPageRelation(...)` 仅包装了 `Page(...)` 并原样返回分页结果，关系装载逻辑尚未在该模块内实现。
