# Member 会员模块

Member 模块提供了完整的会员注册、登录和第三方认证功能，支持多种认证提供商和灵活的用户管理。

## 功能特性

- 📝 用户注册和登录
- 🔑 JWT 认证机制
- 🌐 第三方登录支持（微信小程序等）
- 🛡️ 安全中间件
- 📊 用户信息管理
- 🔒 访问限制
- 🔄 刷新Token机制

## 模块结构

```
member/
├── auth_controller.go    # 认证控制器
├── user_controller.go    # 用户控制器
├── model.go             # 数据模型
├── instance.go          # 实例管理
├── user.go              # 用户实体
├── module.go            # 模块定义
└── README.md            # 基础说明
```

## 快速开始

### 基本使用

```go
package main

import (
    "github.com/sohaha/zlsgo/zlog"
    "github.com/zlsgo/app_module/member"
    "github.com/zlsgo/app_module/account/auth"
    "github.com/zlsgo/app_module/database"
    "github.com/zlsgo/app_core/service"
    "github.com/zlsgo/zdb"
)

func main() {
    // 初始化应用
    app := service.NewApp()(nil)

    // 数据库模块
    dbMod := database.New()

    // 会员模块
    memberMod := member.New("your-secret-key", func(o *member.Options) {
        o.ApiPrefix = "/member"
        o.EnableRegister = true
        o.Expire = 7200
        o.Providers = []auth.AuthProvider{
            &auth.Weapp{
                AppId:     "wx55a57ece33099d66",
                AppSecret: "your-app-secret",
            },
        }
        o.EnabledProviders = []string{"weapp"}
    })

    // 注册全部模块
    err := service.InitModule([]service.Module{dbMod, memberMod}, app)
    if err != nil {
        panic(err)
    }

    // 启动服务
    service.RunWeb(app)
}
```

### 配置文件

以下配置项可通过配置文件注入（字段 tag 对应的 `z` key 如下）：

```yaml
member:
  prefix: "/member"       # API 前缀
  model_prefix: "mem_"    # 数据表前缀（可留空）
  enable_register: true   # 允许注册
  only: false             # 仅模式
  expire: 7200            # token 过期时间（秒）
  enabled_providers:
    - "weapp"             # 启用的提供商（默认全部启用）
```

> ⚠️ 注意：`key` 与 `Providers` 字段的配置 tag 为 `z:"-"`，**无法通过配置文件注入**，
> 必须在代码中调用 `member.New("your-secret-key", func(o *member.Options){ o.Providers = ... })` 传入
> （见上文「快速开始」）。若未在代码中配置提供商，则第三方登录不可用。

## API 接口

### 用户认证

| 方法  | 路径               | 描述         | 权限 |
| ----- | ------------------ | ------------ | ---- |
| POST  | `/member/register` | 用户注册     | 公开 |
| POST  | `/member/login`    | 用户登录     | 公开 |
| GET   | `/member/info`     | 获取用户信息 | 认证 |
| PATCH | `/member/me`       | 更新用户信息 | 认证 |

### 第三方认证

| 方法 | 路径                               | 描述       | 权限 |
| ---- | ---------------------------------- | ---------- | ---- |
| POST | `/member/auth/{provider}`          | 第三方登录 | 公开 |
| GET  | `/member/auth/{provider}/callback` | 登录回调   | 公开 |

## 使用示例

本模块主要通过 HTTP API 使用（路由由模块内部 controller 注册），不提供 `member.Register/Login/ParseToken/...` 这类包级函数。

### 用户注册

```bash
curl -X POST http://127.0.0.1:8080/member/register \
  -H 'Content-Type: application/json' \
  -d '{"account":"demo","password":"demo1234","nickname":"demo"}'
```

> 密码强度要求：长度至少 8 位，且必须同时包含字母和数字。

### 用户登录

```bash
curl -X POST http://127.0.0.1:8080/member/login \
  -H 'Content-Type: application/json' \
  -d '{"account":"demo","password":"demo1234"}'
```

登录成功返回：

- **`token`**: access token
- **`refresh_token`**: refresh token

后续请求使用 `Authorization` 头（模块内部通过 `account/jwt.GetToken` 解析）。

### 获取当前用户信息

```bash
curl http://127.0.0.1:8080/member/info \
  -H 'Authorization: Basic <token>'
```

### 更新当前用户信息

```bash
curl -X PATCH http://127.0.0.1:8080/member/me \
  -H 'Authorization: Basic <token>' \
  -H 'Content-Type: application/json' \
  -d '{"nickname":"new_name","avatar":"/avatar.png"}'
```

## 第三方登录

### 微信小程序登录

```go
// 配置微信小程序登录
memberModule := member.New("your-secret-key", func(o *member.Options) {
    o.Providers = []auth.AuthProvider{
        &auth.Weapp{
            AppId:     "wx55a57ece33099d66",
            AppSecret: "your-app-secret",
        },
    }
    o.EnabledProviders = []string{"weapp"}
})
```

第三方登录同样通过模块内部路由使用：

```bash
curl -X POST http://127.0.0.1:8080/member/auth/weapp \
  -H 'Content-Type: application/json' \
  -d '{"code":"<wx-code>"}'
```

`/member/auth/{provider}` 由 `Auth` controller 初始化时通过 `account/auth.NewRouter(...)` 挂载，登录/回调流程由各 `auth.AuthProvider` 实现。

### 自定义认证提供商

```go
// 实现自定义认证提供商
type CustomProvider struct {
    auth.AuthProvider
    clientID     string
    clientSecret string
}

func (p *CustomProvider) Authenticate(code string) (*auth.UserInfo, error) {
    // 实现自定义认证逻辑
    // 调用第三方API获取用户信息
    userInfo, err := p.getUserInfoFromThirdParty(code)
    if err != nil {
        return nil, err
    }

    return &auth.UserInfo{
        Provider:      "custom",
        ProviderID:    userInfo.ID,
        Username:      userInfo.Name,
        Avatar:        userInfo.Avatar,
        Email:         userInfo.Email,
        ExtensionData: userInfo.RawData,
    }, nil
}

// 注册自定义提供商
memberModule := member.New("your-secret-key", func(o *member.Options) {
    o.Providers = []auth.AuthProvider{
        &CustomProvider{
            clientID:     "your-client-id",
            clientSecret: "your-client-secret",
        },
    }
    o.EnabledProviders = []string{"custom"}
})
```
## 中间件

### 权限中间件

模块在 `UserServer.Init` 中会为除 `POST /register`、`POST /login` 外的路由启用鉴权中间件：

- 未携带 token 时返回 `401 Unauthorized`
- token 校验逻辑位于本模块 `instance.go`（内部调用 `account/jwt.Parse`）