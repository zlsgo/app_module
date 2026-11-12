# app_module

Common modules for zlsgo framework.

## 安装

```bash
go get github.com/zlsgo/app_module
```

## 模块列表

### [Auth](./auth/README.md) - 通用认证模块

提供面向业务站点的认证基础设施，使用服务端 session + cookie 管理登录态，并支持密码找回与 OAuth 绑定。在当前仓库中作为**唯一身份中心（canonical identity）**。

- 🍪 服务端会话与显式失效
- 📝 邮箱注册、登录、登出、当前用户
- 🔐 修改密码、找回密码、reset token
- 🌐 OAuth 登录、自动建号、绑定解绑
- 🚦 登录与找回密码限流

### [Account](./account/README.md) - 账户管理模块

提供后台账号认证、授权与权限管理能力，支持基于角色的访问控制（RBAC），适用于管理端场景。

- 🔐 用户认证（登录/登出）
- 👥 用户管理（增删改查）
- 🔑 权限管理（RBAC）
- 🛡️ 安全中间件
- 📝 操作日志
- 🚦 访问限制
- 🔄 会话管理
- 📡 服务器推送事件（SSE）

### [Database](./database/README.md) - 数据库模块

提供统一的数据库连接与管理能力，支持多种数据库驱动、命名键分区与固定连接事务。

- 🔌 多数据库支持（MySQL、PostgreSQL（构建标签启用）、SQLite）
- 🔄 数据库驱动管理
- 📊 基础连接池功能
- 🗃️ 数据库配置管理
- 🧩 命名键分区与固定连接事务

### [HTML](./html/README.md) - HTML 渲染模块

提供基于 Go 的服务端 HTML 组件化 DSL，可与 `znet` 路由和 `service.Module` 生命周期集成。

- 声明式组件 DSL
- HTMX 2.x 集成
- CSS 样式辅助（`styles`）
- ZView.js 增强交互
- 多返回模式与错误兜底

### [Member](./member/README.md) - 会员资料模块

提供依赖 `auth` 登录态的会员资料接口；定位为业务 profile 层，不承担认证职责。

- 📊 会员资料管理
- 🍪 依赖 `auth` 会话态
- 🔗 `auth user -> member profile` 显式映射
- 🛡️ 认证中间件与访问限制

### [Model](./model/README.md) - 模型模块

提供 Schema 驱动的数据建模与访问层，统一处理字段校验、自动迁移、关联装载、查询过滤与数据前后处理。

- Schema 驱动与自动迁移
- 统一 CRUD / Repository / Query Builder
- 字段管线与校验
- 关联装载
- Schema API 与视图元数据

### [REST API](./restapi/README.md) - REST API 模块

基于已注册 `model.Store` 提供通配路由转发与标准 RESTful CRUD 能力。

- 🚀 通配路由自动分发（按 `{prefix}/{model}/...`）
- 📝 标准 CRUD 操作（GET/POST/PUT/PATCH/DELETE）
- 🔎 字段/关联/排序/过滤查询
- 📁 文件上传封装（`HanderUpload`）
- 🛡️ 中间件与响应钩子支持

### [Template](./template/README.md) - 模板引擎模块

基于 jet 的服务端模板渲染能力，与 `znet` 路由和 `service.Module` 生命周期集成。

- 🚀 jet 模板渲染（`{{: ... }}` 定界符）
- 📁 模板目录加载（默认 `./views`）
- 📦 静态资源托管
- 🛠️ 自定义模板函数
- 🔄 调试/开发热重载

## 目录结构

```
app_module/
├── auth/         # 通用认证模块
├── account/      # 账户管理模块
├── database/     # 数据库模块
├── html/         # HTML 渲染模块
├── member/       # 会员资料模块
├── model/        # 数据模型模块
├── restapi/      # REST API 模块
├── template/     # 模板引擎模块
├── go.mod        # 模块定义
└── ...
```

## 快速开始

```go
package main

import (
    authmodule "github.com/zlsgo/app_module/auth"
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
    authMod := authmodule.New()
    accMod := account.New("your-secret-key", func(o *account.Options) {
        o.ApiPrefix = "/api"
        o.EnableRegister = true
    })
    restApiMod := restapi.New(func(o *restapi.Options) {
        o.Prefix = "/api/v1"
    })

    // 初始化所有模块
    err := service.InitModule([]service.Module{dbMod, modelMod, authMod, accMod, restApiMod}, app)
    if err != nil {
        panic(err)
    }

    service.RunWeb(app)
}
```

> 完整文档：[https://docs.73zls.com/zlsgo/#/13e0bca6-003e-8046-9330-ffe1e3cf20c8](https://docs.73zls.com/zlsgo/#/13e0bca6-003e-8046-9330-ffe1e3cf20c8)