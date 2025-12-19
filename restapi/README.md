# REST API 模块

REST API 模块提供了完整的 RESTful API 功能，包括自动路由、数据操作、文件上传和关联关系处理。

## 功能特性

- 🚀 通配路由自动分发（按 `{prefix}/{model}/...`）
- 📝 标准 CRUD 操作（GET/POST/PUT/PATCH/DELETE）
- 📁 文件上传封装（`HanderUpload`）
- 🛡️ 中间件支持（`Options.Middleware`）
- 🔄 响应钩子（`Options.ResponseHook`）

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
| GET       | `{prefix}/{model}`      | 分页查询（读取 `page`/`pagesize`，最大 1000），默认按 `id desc`   |
| GET       | `{prefix}/{model}/{id}` | 查询单条                                                          |
| POST      | `{prefix}/{model}`      | 插入（返回 `{ "id": ... }`）                                      |
| PUT/PATCH | `{prefix}/{model}/{id}` | 更新（返回 `{ "total": ... }`）                                   |
| DELETE    | `{prefix}/{model}/{id}` | 删除（返回 `{ "total": ... }`）                                   |

> **注意**：出于安全考虑，已禁用全量查询（`{prefix}/{model}/*`），请使用分页接口。

其中 `{model}` 来自 `model` 模块已注册的 `Stores`（`model.Store`）。

### Options

`Options` 与实现保持一致：

- `Prefix string`
- `Middleware znet.Handler`：若不为 `nil`，会在通配路由前 `r.Use(Middleware)`
- `ResponseHook func(c *znet.Context, model, args, method string) bool`：返回 `false` 时会走 `404`（中断请求）

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