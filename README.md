# app_module

Common modules for zlsgo framework

## 安装

```bash
go get github.com/zlsgo/app_module
```

## 模块列表

### [Account](./account/README.md) - 账户管理模块
提供完整的用户认证、授权和权限管理功能，支持基于角色的访问控制（RBAC）。

- 🔐 用户认证（登录/登出）
- 👥 用户管理（增删改查）
- 🔑 权限管理（RBAC）
- 🛡️ 安全中间件
- 📝 操作日志
- 🚦 访问限制
- 🔄 会话管理
- 📡 服务器推送事件（SSE）

### [Database](./database/README.md) - 数据库模块
提供统一的数据库连接和管理功能，支持多种数据库驱动和高级特性。

- 🔌 多数据库支持（MySQL、PostgreSQL（构建标签启用）、SQLite）
- 🔄 数据库驱动管理
- 📊 基础连接池功能
- 🗃️ 数据库配置管理
- 🛠️ 模块化设计

### [HTML](./html/README.md) - HTML 渲染模块
提供基于 Go 语言的 HTML 组件化 DSL，与 `znet` 路由和 `service.Module` 生命周期集成。

- 声明式组件 DSL
- 模块化接入
- 多返回模式
- 错误兜底

### [Member](./member/README.md) - 会员模块
提供完整的会员注册、登录和第三方认证功能，支持多种认证提供商和灵活的用户管理。

- 📝 用户注册和登录
- 🔑 JWT 认证机制
- 🌐 第三方登录支持（微信小程序等）
- 🛡️ 安全中间件
- 📊 用户信息管理
- 🔒 访问限制
- 🔄 刷新Token机制

### [Model](./model/README.md) - 模型模块
提供数据建模与访问层，围绕 Schema 驱动的模型定义，统一处理字段校验、自动迁移、关联装载、查询过滤与数据前后处理。

- Schema 驱动
- 自动迁移
- 统一 CRUD
- 字段管线
- 关联装载
- Schema API

### [REST API](./restapi/README.md) - REST API 模块
提供基于已注册 `model.Store` 的通配路由转发与基础 CRUD helper。

- 🚀 通配路由自动分发（按 `{prefix}/{model}/...`）
- 📝 标准 CRUD 操作（GET/POST/PUT/PATCH/DELETE）
- 📁 文件上传封装（`HanderUpload`）
- 🛡️ 中间件支持（`Options.Middleware`）
- 🔄 响应钩子（`Options.ResponseHook`）

## 目录结构

```
app_module/
├── account/      # 账户管理模块
├── database/     # 数据库模块
├── member/       # 会员模块
├── model/        # 数据模型模块
├── restapi/      # REST API 模块
├── html/         # HTML 渲染模块
├── template/     # 模板引擎模块
├── go.mod        # 模块定义
└── ...
```

## 快速开始

```go
package main

import (
    "github.com/zlsgo/app_module/account"
    "github.com/zlsgo/app_module/database"
    "github.com/zlsgo/app_module/model"
    "github.com/zlsgo/app_module/restapi"
    "github.com/zlsgo/app_core/service"
)

func main() {
    app := service.NewApp()(nil)

    // 注册模块
    dbMod := database.New()
    modelMod := model.New()
    accMod := account.New("your-secret-key", func(o *account.Options) {
        o.ApiPrefix = "/api"
        o.EnableRegister = true
    })
    restApiMod := restapi.New(func(o *restapi.Options) {
        o.Prefix = "/api/v1"
    })

    // 初始化所有模块
    err := service.InitModule([]service.Module{dbMod, modelMod, accMod, restApiMod}, app)
    if err != nil {
        panic(err)
    }

    service.RunWeb(app)
}
```

[完整文档 https://docs.73zls.com/zls-go/#/](https://docs.73zls.com/zlsgo/#/13e0bca6-003e-8046-9330-ffe1e3cf20c8)
