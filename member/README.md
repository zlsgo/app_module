# Member 会员模块

`member` 模块现在是纯资料层。

它不再提供 JWT、本地账号密码或第三方登录；这些能力都统一归 `auth`。`member` 只负责：

- member profile / 业务资料
- `auth user -> member profile` 的显式映射
- 资料接口，例如 `/member/info`、`/member/me`

## 功能特性

- 📊 用户信息与资料接口
- 🍪 依赖 `auth` 会话态
- 🔗 `auth_user_id` 显式关联
- 🛡️ 安全中间件
- 🔒 访问限制

## 模块结构

```
member/
├── auth_bridge.go        # auth user -> member profile 映射
├── user_controller.go    # 用户控制器
├── model.go             # 数据模型
├── instance.go          # 实例管理
├── user.go              # 用户实体
├── action.go            # 操作定义
├── module.go            # 模块定义
└── README.md            # 基础说明
```

## 快速开始

### 基础接入

```go
package main

import (
    authmodule "github.com/zlsgo/app_module/auth"
    "github.com/zlsgo/app_module/database"
    "github.com/zlsgo/app_module/member"
    "github.com/zlsgo/app_core/service"
)

func main() {
    app := service.NewApp()(nil)

    dbMod := database.New()
    authMod := authmodule.New(func(o *authmodule.Options) {
        o.ApiPrefix = "/auth"
    })
    memberMod := member.New(func(o *member.Options) {
        o.ApiPrefix = "/member"
    })

    if err := service.InitModule([]service.Module{dbMod, authMod, memberMod}, app); err != nil {
        panic(err)
    }

    service.RunWeb(app)
}
```

开启后，`/member/info`、`/member/me` 会直接依赖 `auth` 模块的登录 cookie。首次访问资料接口时，如果当前 `auth user` 还没有对应的 member profile，模块会自动补建一条记录并写入 `auth_user_id`。

### 配置文件

```yaml
member:
  prefix: "/member"                   # API 前缀
  model_prefix: "mem_"                # 数据表前缀
```

## API 接口

### 资料接口

| 方法  | 路径               | 描述         | 权限 |
| ----- | ------------------ | ------------ | ---- |
| GET   | `/member/info`     | 获取用户信息 | 认证 |
| PATCH | `/member/me`       | 更新用户信息 | 认证 |

这两条接口都要求当前请求已经带有有效的 `auth` session。

## 使用示例

### 获取当前用户信息

```bash
curl http://127.0.0.1:8080/member/info \
  -H 'Cookie: auth_session=<session-cookie>'
```

### 更新当前用户信息

```bash
curl -X PATCH http://127.0.0.1:8080/member/me \
  -H 'Cookie: auth_session=<session-cookie>' \
  -H 'Content-Type: application/json' \
  -d '{"nickname":"new_name","avatar":"/avatar.png"}'
```

## 与 `auth` 模块的关系

- `auth` 是唯一身份中心，负责用户主体、密码、session、OAuth 与 password reset。
- `member` 负责业务资料和会员接口，不再作为认证系统演进。
- `member` 不再提供自己的 JWT、本地账号密码或第三方登录入口。
- `member` 读取当前用户时只信任 `auth` session。
- `member` 与 `auth.user` 的关系只认 `auth_user_id`，不再回填 legacy 绑定。
## 中间件

### 权限中间件

模块在 `UserServer.Init` 中会为资料路由启用鉴权中间件：

- 未携带有效 `auth_session` 时返回 `401 Unauthorized`
- 当前请求通过 `auth.SessionMiddleware()` 解析 session
- 命中有效 `auth user` 后，`member` 会按 `auth_user_id` 读取或补建对应 profile
