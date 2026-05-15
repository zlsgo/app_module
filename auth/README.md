# Auth 认证模块

`auth` 模块提供面向业务站点的完整认证能力，使用 `znet/session + cookie` 维护登录态，并补齐密码找回、OAuth 登录、绑定解绑等能力。

在当前仓库的统一重构方向里，`auth` 是唯一身份中心（canonical identity）。

它与现有的 `account`、`member` 的分工是：

- `account` 侧重后台账号、权限与 RBAC
- `member` 侧重业务会员资料，不再承担任何登录能力
- `auth` 负责通用站点认证基础设施、服务端会话与第三方登录真源

## 功能特性

- 邮箱注册、登录、登出、获取当前用户
- 修改邮箱、修改密码
- 服务端会话索引，支持显式失效
- 密码找回与重置 token 闭环
- OAuth 登录、自动建号、登录后绑定、解绑
- 登录与找回密码限流
- 可向其他模块导出当前登录用户与 session 中间件

## 模块结构

```text
auth/
├── controller.go        # 本地账号接口
├── oauth_controller.go  # OAuth 登录/绑定/解绑
├── password_reset.go    # 找回密码与 reset token
├── session.go           # 会话索引与当前用户解析
├── rate_limit.go        # 限流
├── model.go             # 数据模型定义
├── module.go            # 模块入口与配置
└── README.md            # 使用说明
```

## 快速开始

```go
package main

import (
    authmodule "github.com/zlsgo/app_module/auth"
    accountauth "github.com/zlsgo/app_module/account/auth"
    "github.com/zlsgo/app_module/database"
    "github.com/zlsgo/app_core/service"
    "github.com/zlsgo/zdb"
)

func main() {
    app := service.NewApp()(nil)

    dbMod := database.New()
    authMod := authmodule.New(func(o *authmodule.Options) {
        o.ApiPrefix = "/auth"
        o.BaseURL = "https://example.com"
        o.SendResetPassword = func(job authmodule.ResetPasswordJob) error {
            // 在这里接入你自己的邮件发送器
            // job.ResetURL 是可直接发给用户的重置链接
            return nil
        }
        o.Providers = []accountauth.AuthProvider{
            &accountauth.Weapp{
                AppId:     "your-app-id",
                AppSecret: "your-app-secret",
            },
        }
        o.EnabledProviders = []string{"weapp"}
    })

    if err := service.InitModule([]service.Module{dbMod, authMod}, app); err != nil {
        panic(err)
    }

    service.RunWeb(app)
}
```

## 配置项

```yaml
auth:
  prefix: "/auth"                     # 路由前缀
  model_prefix: "auth_"               # 数据表前缀
  base_url: "https://example.com"     # reset 页面站点基准地址
  reset_password_path: "/reset-password" # reset 页面路径
  cookie_name: "auth_session"         # session cookie 名称
  session_ttl: 2592000                # 会话有效期（秒）
  password_reset_token_ttl: 3600      # reset token 有效期（秒）
  enabled_providers:
    - "weapp"
```

代码注入配置：

- `Options.InitDB`：自定义数据库初始化
- `Options.Session`：自定义 `znet/session` store
- `Options.Providers`：第三方登录 provider 列表
- `Options.SendResetPassword`：发送 reset 邮件的回调
- `Options.ResetPasswordPath`：邮件里生成的前端 reset 页面路径

## 对外复用

当其他模块希望复用 `auth` 的登录态时，可以直接依赖当前模块暴露的两个集成入口：

- `Module.SessionMiddleware()`：挂载与 `auth` 控制器一致的 session 解析逻辑
- `Module.CurrentUser(c)`：在当前请求上下文中读取已登录的 auth 用户

`member` 模块当前直接依赖这两个入口来读取当前登录用户，并把 `auth user` 映射成自己的 profile。

注意：

- `member` 不再保留自己的登录入口、JWT 或 provider 路径。
- 业务站点如果需要“当前会员资料”，应先拿到有效的 `auth` session，再访问 `member` 资料接口。

## API 接口

### 本地账号

| 方法 | 路径 | 描述 | 权限 |
| --- | --- | --- | --- |
| POST | `/auth/user/create` | 注册并创建会话 | 公开 |
| POST | `/auth/user/auth` | 邮箱密码登录 | 公开 |
| GET | `/auth/user/get` | 获取当前用户 | 认证 |
| GET | `/auth/user/signout` | 登出并销毁当前会话 | 认证 |
| POST | `/auth/user/update` | 更新邮箱或密码 | 认证 |
| POST | `/auth/user/forgotpassword` | 发送重置密码请求 | 公开 |
| POST | `/auth/user/resetpassword` | 使用 token 重置密码 | 公开 |

### OAuth

| 方法 | 路径 | 描述 | 权限 |
| --- | --- | --- | --- |
| ANY | `/auth/user/oauth/login/{provider}` | 进入第三方登录流程，生成一次性 `state` | 公开 |
| ANY | `/auth/user/oauth/callback/{provider}` | 第三方回调；未登录时自动建号/登录，已登录时执行绑定 | 公开/认证 |
| POST | `/auth/user/oauth/add` | 发起已登录用户的 provider 绑定流程 | 认证 |
| POST | `/auth/user/oauth/remove` | 解绑 provider 标识 | 认证 |

## 请求示例

### 注册

```bash
curl -X POST http://127.0.0.1:8080/auth/user/create \
  -H 'Content-Type: application/json' \
  -d '{"email":"demo@example.com","password":"Aa123456","nickname":"demo"}'
```

### 登录

```bash
curl -X POST http://127.0.0.1:8080/auth/user/auth \
  -H 'Content-Type: application/json' \
  -d '{"email":"demo@example.com","password":"Aa123456"}'
```

### 修改密码

```bash
curl -X POST http://127.0.0.1:8080/auth/user/update \
  -H 'Content-Type: application/json' \
  -H 'Cookie: auth_session=<session-cookie>' \
  -d '{"current_password":"Aa123456","password":"Bb123456"}'
```

### 发送密码重置邮件

```bash
curl -X POST http://127.0.0.1:8080/auth/user/forgotpassword \
  -H 'Content-Type: application/json' \
  -d '{"email":"demo@example.com"}'
```

### 使用 reset token 重置密码

```bash
curl -X POST http://127.0.0.1:8080/auth/user/resetpassword \
  -H 'Content-Type: application/json' \
  -d '{"token":"<reset-token>","password":"Cc123456"}'
```

### 发起 provider 绑定

```bash
curl -X POST http://127.0.0.1:8080/auth/user/oauth/add \
  -H 'Content-Type: application/json' \
  -H 'Cookie: auth_session=<session-cookie>' \
  -d '{"provider":"weapp"}'
```

## 行为说明

- 登录成功、注册成功、OAuth callback 成功后都会写入 session cookie。
- 修改密码与 reset password 都会提升 `session_version`，并让历史会话失效。
- `forgotpassword` 默认对外返回统一成功结构；是否真正发送邮件由 `SendResetPassword` 回调决定，`ResetURL` 指向前端 reset 页面而不是 POST API。
- `/oauth/login/{provider}` 与 `/oauth/add` 都会写入一次性 `state`；callback 必须带回匹配的 `state` 才会继续。
- `/oauth/add` 只负责发起第三方绑定流程，真正的绑定发生在 provider callback 且当前用户已登录，且 flow 意图必须是 `bind`。
- OAuth callback 未命中已有绑定时：
  - provider 返回已验证邮箱，且本地已有同邮箱用户：会按邮箱合并并创建绑定
  - 否则自动建号，并生成占位邮箱
- 没有本地密码的 OAuth-only 用户不能解绑最后一个第三方登录；需要先在已登录会话下设置本地密码。

## 测试

```bash
go test ./auth/...
```
