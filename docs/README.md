# App Module 文档

app_module 是一个基于 Go 语言和 zlsgo 框架的通用模块库，提供了完整的 Web 应用开发所需的核心功能模块。

## 模块概览

### 核心模块

| 模块 | 功能描述 | 主要用途 |
|------|----------|----------|
| [account](./account.md) | 账户管理系统 | 用户认证、权限管理、角色控制 |
| [database](./database.md) | 数据库管理 | 数据库连接、迁移、操作封装 |
| [member](./member.md) | 会员系统 | 用户注册、登录、第三方认证 |
| [model](./model.md) | 数据模型 | ORM 封装、数据验证、迁移管理 |
| [restapi](./restapi.md) | REST API | RESTful 接口封装、文件上传 |
| [html](./html.md) | HTML 渲染 | 模板渲染、静态资源管理 |
| [template](./template.md) | 模板引擎 | 模板管理、布局系统 |

### 设计理念

- **模块化设计**: 每个模块都是独立的功能单元，可按需引入
- **统一接口**: 所有模块都遵循 `service.Module` 接口规范
- **配置驱动**: 通过配置文件灵活控制模块行为
- **依赖注入**: 基于 zdi 容器实现依赖管理
- **中间件支持**: 完整的中间件生态

## 快速开始

### 安装

```bash
go get github.com/zlsgo/app_module
```

### 基本使用

```go
package main

import (
    "github.com/sohaha/zlsgo/zlog"
    "github.com/zlsgo/app_module/database"
    "github.com/zlsgo/app_module/account"
    "github.com/zlsgo/app_core/service"
    "github.com/zlsgo/zdb"
)

func main() {
    // 初始化应用
    app := service.NewApp()(nil)

    // 数据库模块
    dbMod := database.New()

    // 账户模块
    accMod := account.New("your-secret-key", func(o *account.Options) {
        o.ApiPrefix = "/api"
        o.EnableRegister = true
        o.AdminDefaultPassword = "admin123"
    })

    // 注册全部模块
    err := service.InitModule([]service.Module{dbMod, accMod}, app)
    if err != nil {
        panic(err)
    }

    // 使用数据库对象
    err = app.DI.InvokeWithErrorOnly(func(db *zdb.DB) {
        // 执行数据库操作
        rows, err := db.QueryToMaps("SELECT COUNT(*) as count FROM users")
        if err != nil {
            zlog.Error(err)
            return
        }
        zlog.Infof("User count: %v", rows)
    })
    if err != nil {
        panic(err)
    }

    // 注册路由
    _ = service.Global.DI.InvokeWithErrorOnly(func(r *znet.Engine) {
        r.GET("/", func(c *znet.Context) string {
            return "Hello, World!"
        })
    })

    // 启动服务
    service.RunWeb(app)
}
```

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
└── docs/         # 文档目录
```

## 配置说明

所有模块都支持通过配置文件进行配置，配置文件通常为 `config.yaml` 或 `config.json`。

### 示例配置

```yaml
# 数据库配置
database:
  driver: "mysql"
  dsn: "user:password@tcp(localhost:3306)/dbname"
  mode:
    log: true

# 账户配置
account:
  key: "your-secret-key"
  prefix: "/api"
  register: true
  rbac_file: "./rbac.yaml"

# 会员配置
member:
  key: "your-member-secret"
  prefix: "/member"
  enable_register: true
  expire: 3600
```

## 模块开发

### 创建自定义模块

```go
package mymodule

import (
    "github.com/sohaha/zlsgo/zdi"
    "github.com/zlsgo/app_core/service"
)

type Module struct {
    service.ModuleLifeCycle
}

func (m *Module) Name() string {
    return "MyModule"
}

func (m *Module) Load(di zdi.Invoker) (any, error) {
    // 模块初始化逻辑
    return nil, nil
}

func (m *Module) Start(di zdi.Invoker) error {
    // 模块启动逻辑
    return nil
}

func (m *Module) Controller() []service.Controller {
    // 返回模块控制器
    return []service.Controller{}
}

func New() *Module {
    return &Module{}
}
```

## 贡献指南

1. Fork 本项目
2. 创建功能分支
3. 提交代码变更
4. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证，详见 [LICENSE](../LICENSE) 文件。
