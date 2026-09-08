# Account 账户管理模块

Account 模块提供了完整的用户认证、授权和权限管理功能，支持基于角色的访问控制（RBAC）。

## 功能特性

- 🔐 用户认证（登录/登出）
- 👥 用户管理（增删改查）
- 🔑 权限管理（RBAC）
- 🛡️ 安全中间件
- 📝 操作日志
- 🚦 访问限制
- 🔄 会话管理
- 📡 服务器推送事件（SSE）

## 模块结构

```
account/
├── auth/              # 认证相关
├── jwt/               # JWT 处理
├── limiter/           # 访问限制
├── rbac/              # 角色权限控制
├── controller.go      # 控制器
├── middleware.go      # 中间件
├── model.go           # 数据模型
├── module.go          # 模块定义
└── ...
```

## 快速开始

### 基本配置

```go
package main

import (
    "github.com/zlsgo/app_module/account"
    "github.com/zlsgo/app_module/database"
    "github.com/zlsgo/app_core/service"
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

    // 账户模块核心功能演示
    // ApiPrefix 默认为 /manage，本示例设置为 /api
    // 初始化后可访问：
    // - POST /api/base/login           - 用户登录
    // - POST /api/base/logout          - 用户退出
    // - ANY  /api/base/refresh-token   - 刷新令牌
    // - GET  /api/base/info            - 获取当前用户信息
    // - GET  /api/base/site            - 获取系统信息
    // - POST /api/base/register        - 用户注册（需 EnableRegister）
    // - POST /api/base/password        - 修改当前用户密码
    // - PATCH /api/base/me             - 更新当前用户信息
    // - POST /api/base/avatar          - 上传头像
    // - GET  /api/user                 - 用户列表
    // - POST /api/user                 - 创建用户
    // - PATCH /api/user/{uid}          - 更新用户
    // - DELETE /api/user/{uid}         - 删除用户
    // - ANY  /api/message/realtime     - 建立 SSE 实时推送
}
```

### 配置文件

```yaml
account:
  key: "your-secret-key"              # 加密密钥（模块内部会自动补齐到32位）
  prefix: "/api"                      # API 前缀
  admin_default_password: "admin123"  # 默认管理员密码
  model_prefix: "sys_"                # 数据表前缀
  register: true                      # 允许注册
  only: false                         # 仅模式
  disabled_ip: false                  # 禁用IP记录
  expire: 7200                        # token 过期时间（秒）
  refresh_expire: 604800              # 刷新token过期时间（秒）
  rbac_file: "./rbac.yaml"            # RBAC 配置文件
```

## API 接口

### 基础接口（`{prefix}` 为 `Options.ApiPrefix`）

| 方法  | 路径                          | 描述                   | 权限 |
| ----- | ----------------------------- | ---------------------- | ---- |
| POST  | `{prefix}/base/login`         | 用户登录               | 公开 |
| POST  | `{prefix}/base/logout`        | 用户退出登录           | 认证 |
| ANY   | `{prefix}/base/refresh-token` | 刷新访问令牌           | 认证 |
| GET   | `{prefix}/base/info`          | 获取当前用户信息       | 认证 |
| GET   | `{prefix}/base/site`          | 获取系统信息           | 公开 |
| POST  | `{prefix}/base/register`      | 用户注册（需启用注册） | 公开 |
| POST  | `{prefix}/base/password`      | 修改当前用户密码       | 认证 |
| PATCH | `{prefix}/base/me`            | 更新当前用户资料       | 认证 |
| POST  | `{prefix}/base/avatar`        | 上传当前用户头像       | 认证 |

> `register` 接口仅当 `Options.EnableRegister` 为 `true` 时可访问。
>
> `GET {prefix}/base/site`（公开）返回系统信息，字段包括：`name`（模块名）、`prefix`（API 前缀）、
> `register`（是否开放注册）、`time`（服务器当前时间）。

### 消息通知

| 方法 | 路径                        | 描述                  | 权限 |
| ---- | --------------------------- | --------------------- | ---- |
| GET  | `{prefix}/message`          | 获取未读通知统计      | 认证 |
| ANY  | `{prefix}/message/realtime` | 建立 SSE 实时推送连接 | 认证 |

### 用户管理

| 方法   | 路径                  | 描述     | 权限     |
| ------ | --------------------- | -------- | -------- |
| GET    | `{prefix}/user`       | 用户列表 | 用户管理 |
| POST   | `{prefix}/user`       | 创建用户 | 用户管理 |
| PATCH  | `{prefix}/user/{uid}` | 更新用户 | 用户管理 |
| DELETE | `{prefix}/user/{uid}` | 删除用户 | 用户管理 |

### 角色管理

| 方法   | 路径                  | 描述     | 权限     |
| ------ | --------------------- | -------- | -------- |
| GET    | `{prefix}/role`       | 角色列表 | 角色管理 |
| POST   | `{prefix}/role`       | 创建角色 | 角色管理 |
| PATCH  | `{prefix}/role/{rid}` | 更新角色 | 角色管理 |
| DELETE | `{prefix}/role/{rid}` | 删除角色 | 角色管理 |

### 权限管理

| 方法   | 路径                        | 描述     | 权限     |
| ------ | --------------------------- | -------- | -------- |
| GET    | `{prefix}/permission`       | 权限列表 | 权限管理 |
| POST   | `{prefix}/permission`       | 创建权限 | 权限管理 |
| PATCH  | `{prefix}/permission/{pid}` | 更新权限 | 权限管理 |
| DELETE | `{prefix}/permission/{pid}` | 删除权限 | 权限管理 |

> `uid`、`rid`、`pid` 为对应资源的加密主键，框架会自动解析。

## 使用示例

### 中间件与上下文使用

```go
package main

import (
    "github.com/sohaha/zlsgo/znet"
    "github.com/sohaha/zlsgo/ztype"
    "github.com/zlsgo/app_core/service"
    "github.com/zlsgo/app_module/account"
)

func main() {
    app := service.NewApp()(nil)

    accMod := account.New("your-secret-key", func(o *account.Options) {
        o.ApiPrefix = "/api"
        o.EnableRegister = true
    })

    if err := service.InitModule([]service.Module{accMod}, app); err != nil {
        panic(err)
    }

    _ = service.Global.DI.InvokeWithErrorOnly(func(r *znet.Engine) {
        if err := accMod.UsePermisMiddleware(r, nil,
            accMod.Options.ApiPrefix+"/base/login",
            accMod.Options.ApiPrefix+"/base/register",
        ); err != nil {
            panic(err)
        }

        r.GET("/profile", func(c *znet.Context) {
            uid := accMod.Request.UID(c)
            if uid == "" {
                c.Fail(401, "请先登录")
                return
            }

            realUID, err := accMod.Request.RealUID(c)
            if err != nil {
                c.Fail(400, err.Error())
                return
            }

            user := accMod.Request.User(c)
            c.JSON(200, ztype.Map{
                "uid":      uid,
                "real_uid": realUID,
                "info":     user,
            })
        })
    })
}
```

### 自定义权限处理

```go
func requireRole(acc *account.Module, role string) znet.Handler {
    return func(c *znet.Context) error {
        roles := acc.Request.Roles(c)
        if !zarray.Contains(roles, role) {
            return zerror.PermissionDenied.Text("权限不足")
        }
        c.Next()
        return nil
    }
}
```

通过 `acc.Request.IgnorePerm(c)` 可以在特定处理函数中临时放行权限校验，适用于公开接口或自定义鉴权流程。
## RBAC 配置

创建 `rbac.yaml` 文件：

```yaml
# 角色定义
roles:
  admin:
    name: "管理员"
    description: "系统管理员"
    permissions:
      - "*"
  user:
    name: "普通用户"
    description: "普通用户"
    permissions:
      - "user:read"
      - "user:update:self"
  guest:
    name: "访客"
    description: "访客用户"
    permissions:
      - "public:read"

# 权限定义
permissions:
  - name: "user:read"
    description: "查看用户"
  - name: "user:update"
    description: "更新用户"
  - name: "user:delete"
    description: "删除用户"
  - name: "public:read"
    description: "公开内容访问"
```

## 安全特性

- **密码加密**: 使用 bcrypt 对密码进行哈希存储
- **JWT 认证**: 支持访问令牌与刷新令牌
- **登录保护**: 登录失败次数限制并在成功登录后刷新盐值
- **访问限制**: 集成 `limiter` 中间件进行 IP/频率限制
- **操作日志**: 自动记录敏感操作请求
- **SSE 会话**: `session.go` 管理实时推送连接并支持断线恢复

## 最佳实践

1. **密钥管理**: 使用强密钥并定期更换
2. **权限最小化**: 只授予必要的权限
3. **定期审计**: 定期检查用户权限和操作日志
4. **安全配置**: 在生产环境中禁用调试模式
