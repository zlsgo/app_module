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
├── action.go            # 操作定义
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
    "github.com/zlsgo/app_module/member/auth"
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

```yaml
member:
  key: "your-secret-key"              # 加密密钥（必须32位）
  prefix: "/member"                   # API 前缀
  model_prefix: "mem_"                # 数据表前缀
  enable_register: true               # 允许注册
  only: false                         # 仅模式
  expire: 7200                        # token 过期时间（秒）

  # 第三方登录提供商
  providers:
    - provider: "weapp"
      app_id: "wx55a57ece33099d66"
      app_secret: "your-app-secret"

  # 启用的提供商
  enabled_providers:
    - "weapp"
```

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

### 用户注册

```go
package main

import (
    "github.com/sohaha/zlsgo/znet"
    "github.com/sohaha/zlsgo/ztype"
)

func registerHandler(c *znet.Context) {
    // 获取注册数据
    data, err := c.GetJSONs()
    if err != nil {
        c.Fail(400, "数据格式错误")
        return
    }

    // 验证必填字段
    account := data.Get("account").String()
    password := data.Get("password").String()
    nickname := data.Get("nickname").String()

    if account == "" || password == "" {
        c.Fail(400, "账号和密码不能为空")
        return
    }

    if nickname == "" {
        nickname = account
    }

    // 构造注册数据
    userData := ztype.Map{
        "account":  account,
        "password": password,
        "nickname": nickname,
        "avatar":   data.Get("avatar").String(),
    }

    // 调用会员模块注册
    result, err := member.Register(userData)
    if err != nil {
        c.Fail(400, err.Error())
        return
    }

    c.Okay(result)
}
```

### 用户登录

```go
func loginHandler(c *znet.Context) {
    // 获取登录数据
    data, err := c.GetJSONs()
    if err != nil {
        c.Fail(400, "数据格式错误")
        return
    }

    account := data.Get("account").String()
    password := data.Get("password").String()

    if account == "" || password == "" {
        c.Fail(400, "请输入账号和密码")
        return
    }

    // 执行登录
    result, err := member.Login(account, password)
    if err != nil {
        c.Fail(401, err.Error())
        return
    }

    c.Okay(result)
}
```

### 获取用户信息

```go
func getUserInfoHandler(c *znet.Context) {
    // 从请求头获取token
    token := c.GetHeader("Authorization")
    if token == "" {
        c.Fail(401, "未提供认证token")
        return
    }

    // 解析token获取用户信息
    user, err := member.ParseToken(token)
    if err != nil {
        c.Fail(401, "无效的token")
        return
    }

    c.Okay(user)
}
```

### 更新用户信息

```go
func updateProfileHandler(c *znet.Context) {
    // 获取当前用户
    user := member.GetCurrentUser(c)
    if user == nil {
        c.Fail(401, "用户未登录")
        return
    }

    // 获取更新数据
    data, err := c.GetJSONs()
    if err != nil {
        c.Fail(400, "数据格式错误")
        return
    }

    // 过滤敏感字段
    allowedFields := []string{"nickname", "avatar", "email", "phone"}
    updateData := ztype.Map{}

    for _, field := range allowedFields {
        if value := data.Get(field); !value.IsEmpty() {
            updateData[field] = value.String()
        }
    }

    // 更新用户信息
    err = member.UpdateUser(user.ID, updateData)
    if err != nil {
        c.Fail(500, "更新失败")
        return
    }

    c.Okay("更新成功")
}
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

// 微信小程序登录处理
func weappLoginHandler(c *znet.Context) {
    // 获取小程序登录code
    type WeappLoginReq struct {
        Code string `json:"code"`
        UserInfo struct {
            NickName string `json:"nickName"`
            Avatar   string `json:"avatarUrl"`
        } `json:"userInfo"`
    }

    var req WeappLoginReq
    if err := c.Parse(&req); err != nil {
        c.Fail(400, "数据格式错误")
        return
    }

    // 调用第三方登录
    result, err := member.AuthWithProvider("weapp", auth.ProviderInfo{
        Code:      req.Code,
        Username:  req.UserInfo.NickName,
        Avatar:    req.UserInfo.Avatar,
    })
    if err != nil {
        c.Fail(400, err.Error())
        return
    }

    c.Okay(result)
}
```

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

```go
// 自定义权限检查
func requirePermission(permission string) znet.HandlerFunc {
    return func(c *znet.Context) {
        user := member.GetCurrentUser(c)
        if user == nil {
            c.Fail(401, "用户未登录")
            c.Abort()
            return
        }

        // 检查用户权限
        if !member.HasPermission(user.ID, permission) {
            c.Fail(403, "权限不足")
            c.Abort()
            return
        }

        c.Next()
    }
}

// 使用权限中间件
r.GET("/admin/users", requirePermission("user:read"), adminUsersHandler)
```
