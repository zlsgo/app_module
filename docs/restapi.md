# REST API 模块

REST API 模块提供了完整的 RESTful API 功能，包括自动路由、数据操作、文件上传和关联关系处理。

## 功能特性

- 🚀 自动 RESTful 路由生成
- 📝 标准 CRUD 操作
- 📁 文件上传支持
- 🔗 关联关系处理
- 🛡️ 中间件支持
- 📊 分页查询
- 🔄 响应钩子
- ⚡ 高性能数据处理

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
    "github.com/sohaha/zlsgo/zlog"
    "github.com/zlsgo/app_module/restapi"
    "github.com/zlsgo/app_module/model"
    "github.com/zlsgo/app_module/database"
    "github.com/zlsgo/app_core/service"
    "github.com/zlsgo/zdb"
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
        o.Middleware = []znet.HandlerFunc{
            middleware.Auth(),
            middleware.RateLimit(),
        }
        o.ResponseHook = func(c *znet.Context, model, id, method string) bool {
            // 响应钩子，返回 true 继续处理，false 中断请求
            return true
        }
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

### 配置文件

```yaml
restapi:
  prefix: "/api/v1"           # API 路径前缀
  middleware: []              # 中间件列表
  response_hook: true         # 启用响应钩子

  # 上传配置
  upload:
    max_size: 10485760        # 最大文件大小（10MB）
    allowed_types:            # 允许的文件类型
      - "image/jpeg"
      - "image/png"
      - "image/gif"
    storage_path: "./uploads" # 存储路径
    url_prefix: "/uploads"    # URL 前缀
```

## API 接口

### 自动生成的端点

基于注册的模型，REST API 模块会自动生成以下端点：

| 方法 | 路径 | 描述 | 示例 |
|------|------|------|------|
| GET | `/api/v1/{model}` | 获取模型列表 | `/api/v1/users` |
| GET | `/api/v1/{model}/{id}` | 获取单条记录 | `/api/v1/users/1` |
| POST | `/api/v1/{model}` | 创建记录 | `/api/v1/users` |
| PUT | `/api/v1/{model}/{id}` | 更新记录 | `/api/v1/users/1` |
| DELETE | `/api/v1/{model}/{id}` | 删除记录 | `/api/v1/users/1` |

### 查询参数

#### 列表查询

```bash
GET /api/v1/users?page=1&pagesize=20&status=1&order=created_at:desc
```

参数说明：
- `page`: 页码（默认：1）
- `pagesize`: 每页数量（默认：20）
- `order`: 排序字段和方向（如：`created_at:desc`）
- 其他字段作为过滤条件

#### 字段过滤

```bash
GET /api/v1/users?fields=id,username,email
```

只返回指定字段，逗号分隔。

## 使用示例

### 基本 CRUD 操作

#### 获取数据

```go
// 获取用户列表
GET /api/v1/users

// 获取分页数据
GET /api/v1/users?page=2&pagesize=10

// 获取过滤数据
GET /api/v1/users?status=1&role=admin

// 获取单条记录
GET /api/v1/users/123

// 获取指定字段
GET /api/v1/users/123?fields=id,username,email
```

#### 创建数据

```bash
POST /api/v1/users
Content-Type: application/json

{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "secure_password",
    "status": 1
}
```

响应：
```json
{
    "id": 123
}
```

#### 更新数据

```bash
PUT /api/v1/users/123
Content-Type: application/json

{
    "email": "newemail@example.com",
    "status": 1
}
```

响应：
```json
{
    "total": 1
}
```

#### 删除数据

```bash
DELETE /api/v1/users/123
```

响应：
```json
{
    "total": 1
}
```

### 高级查询

```go
// 使用 REST API 方法
func handleUsers(c *znet.Context) {
    store := model.GetStore("user")

    // 分页查询
    pageData, err := restapi.Page(c, store, model.Filter{
        "status": 1,
    }, func(options *model.CondOptions) {
        options.OrderBy = map[string]string{
            "created_at": "desc",
        }
        options.Fields = []string{"id", "username", "email"}
    })

    if err != nil {
        c.Fail(500, err.Error())
        return
    }

    c.Okay(pageData)
}
```

### 文件上传

```go
// 上传文件处理
func uploadHandler(c *znet.Context) {
    results, err := restapi.HanderUpload(c, "avatars", func(o *common.UploadOption) {
        o.MaxSize = 5 * 1024 * 1024 // 5MB
        o.AllowedTypes = []string{
            "image/jpeg",
            "image/png",
            "image/gif",
        }
        o.Rename = true
    })

    if err != nil {
        c.Fail(400, err.Error())
        return
    }

    c.Okay(results)
}

// 前端使用
const formData = new FormData();
formData.append('file', fileInput.files[0]);

fetch('/api/v1/upload', {
    method: 'POST',
    body: formData
})
.then(response => response.json())
.then(data => console.log(data));
```

### 关联关系处理

```go
// 关联查询示例
func getUsersWithProfile(c *znet.Context) {
    store := model.GetStore("user")

    // 定义关联关系
    relations := map[string]restapi.Relation{
        "profile": {
            Operation: model.GetStore("user_profile"),
        },
        "roles": {
            Operation: model.GetStore("role"),
        },
    }

    // 分页关联查询
    pageData, err := restapi.HanderPageRelation(c, store, model.Filter{
        "status": 1,
    }, relations)

    if err != nil {
        c.Fail(500, err.Error())
        return
    }

    c.Okay(pageData)
}
```

### 自定义处理逻辑

```go
// 自定义插入处理
func customInsert(c *znet.Context) {
    store := model.GetStore("user")

    result, err := restapi.Insert(c, store, func(data ztype.Map) (ztype.Map, error) {
        // 数据预处理
        data["created_at"] = time.Now()
        data["status"] = 1

        // 密码加密
        if password := data.Get("password").String(); password != "" {
            hashedPassword, err := bcrypt.GenerateFromPassword(
                []byte(password), bcrypt.DefaultCost)
            if err != nil {
                return nil, err
            }
            data["password"] = string(hashedPassword)
        }

        return data, nil
    })

    if err != nil {
        c.Fail(400, err.Error())
        return
    }

    c.Okay(result)
}

// 自定义更新处理
func customUpdate(c *znet.Context) {
    store := model.GetStore("user")
    id := c.Param("id")

    result, err := restapi.UpdateById(c, store, id, func(oldData, newData ztype.Map) (ztype.Map, error) {
        // 权限检查
        currentUser := getCurrentUser(c)
        if currentUser.ID != oldData.Get("id").Int() && !currentUser.IsAdmin() {
            return nil, errors.New("权限不足")
        }

        // 数据验证
        if email := newData.Get("email").String(); email != "" {
            if !isValidEmail(email) {
                return nil, errors.New("邮箱格式不正确")
            }
        }

        // 敏感字段过滤
        delete(newData, "password")
        delete(newData, "role")

        newData["updated_at"] = time.Now()

        return newData, nil
    })

    if err != nil {
        c.Fail(400, err.Error())
        return
    }

    c.Okay(result)
}
```

## 中间件

### 认证中间件

```go
func AuthMiddleware() znet.HandlerFunc {
    return func(c *znet.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.Fail(401, "未提供认证token")
            c.Abort()
            return
        }

        // 验证 token
        user, err := validateToken(token)
        if err != nil {
            c.Fail(401, "无效的token")
            c.Abort()
            return
        }

        // 将用户信息存储到上下文
        c.Set("user", user)
        c.Next()
    }
}
```

### 权限中间件

```go
func PermissionMiddleware(requiredPermission string) znet.HandlerFunc {
    return func(c *znet.Context) {
        user := c.Get("user").(*User)

        if !user.HasPermission(requiredPermission) {
            c.Fail(403, "权限不足")
            c.Abort()
            return
        }

        c.Next()
    }
}
```

### 限流中间件

```go
func RateLimitMiddleware() znet.HandlerFunc {
    limiter := rate.NewLimiter(rate.Every(time.Second), 10) // 每秒10次

    return func(c *znet.Context) {
        if !limiter.Allow() {
            c.Fail(429, "请求过于频繁")
            c.Abort()
            return
        }

        c.Next()
    }
}
```

## 响应格式

### 成功响应

```json
{
    "code": 200,
    "msg": "success",
    "data": {
        "id": 123
    }
}
```

### 分页响应

```json
{
    "code": 200,
    "msg": "success",
    "data": {
        "items": [
            {
                "id": 1,
                "username": "john_doe",
                "email": "john@example.com"
            }
        ],
        "total": 100,
        "page": 1,
        "pagesize": 20,
        "pages": 5
    }
}
```

### 错误响应

```json
{
    "code": 400,
    "msg": "参数错误",
    "data": null
}
```

## 响应钩子

```go
// 全局响应钩子
restApiModule := restapi.New(func(o *restapi.Options) {
    o.ResponseHook = func(c *znet.Context, modelName, id, method string) bool {
        // 记录访问日志
        log.Printf("%s %s %s %s", method, modelName, id, c.ClientIP())

        // 权限检查
        if modelName == "admin" && !isAdmin(c) {
            c.Fail(403, "权限不足")
            return false
        }

        return true
    }
})
```

## 文件上传

### 配置选项

```go
type UploadOptions struct {
    MaxSize     int64             // 最大文件大小
    AllowedTypes []string         // 允许的文件类型
    StoragePath string            // 存储路径
    URLPrefix   string            // URL 前缀
    Rename      bool              // 是否重命名文件
    Compression bool              // 是否压缩图片
    Thumbnail   bool              // 是否生成缩略图
}

// 使用上传功能
func uploadHandler(c *znet.Context) {
    results, err := restapi.HanderUpload(c, "images", func(o *common.UploadOption) {
        o.MaxSize = 10 * 1024 * 1024 // 10MB
        o.AllowedTypes = []string{
            "image/jpeg",
            "image/png",
            "image/gif",
            "image/webp",
        }
        o.StoragePath = "./uploads/images"
        o.URLPrefix = "/uploads"
        o.Rename = true
        o.Compression = true
        o.Thumbnail = true
    })

    if err != nil {
        c.Fail(400, err.Error())
        return
    }

    c.Okay(results)
}
```

### 上传响应格式

```json
{
    "code": 200,
    "msg": "success",
    "data": [
        {
            "name": "image_1234567890.jpg",
            "path": "uploads/images/2024/01/01/image_1234567890.jpg",
            "url": "/uploads/images/2024/01/01/image_1234567890.jpg",
            "size": 1024000,
            "type": "image/jpeg",
            "thumbnail": "/uploads/images/2024/01/01/thumb_image_1234567890.jpg"
        }
    ]
}
```

## 关联关系

### 一对一关系

```go
func getUserWithProfile(c *znet.Context) {
    store := model.GetStore("user")

    relations := map[string]restapi.Relation{
        "profile": {
            Operation: model.GetStore("user_profile"),
        },
    }

    pageData, err := restapi.HanderPageRelation(c, store, model.Filter{}, relations)
    if err != nil {
        c.Fail(500, err.Error())
        return
    }

    c.Okay(pageData)
}
```

### 一对多关系

```go
func getUserWithArticles(c *znet.Context) {
    store := model.GetStore("user")

    relations := map[string]restapi.Relation{
        "articles": {
            Operation: model.GetStore("article"),
        },
    }

    pageData, err := restapi.HanderPageRelation(c, store, model.Filter{
        "id": c.Param("id"),
    }, relations)

    if err != nil {
        c.Fail(500, err.Error())
        return
    }

    c.Okay(pageData)
}
```

## 性能优化

### 查询优化

```go
// 使用字段过滤
func optimizedQuery(c *znet.Context) {
    store := model.GetStore("user")

    pageData, err := restapi.Page(c, store, model.Filter{
        "status": 1,
    }, func(options *model.CondOptions) {
        // 只查询需要的字段
        options.Fields = []string{"id", "username", "avatar"}

        // 添加索引排序
        options.OrderBy = map[string]string{
            "id": "desc",
        }

        // 限制查询数量
        options.Limit = 50
    })

    if err != nil {
        c.Fail(500, err.Error())
        return
    }

    c.Okay(pageData)
}
```

### 缓存策略

```go
// 带缓存的查询
func cachedQuery(c *znet.Context) {
    cacheKey := fmt.Sprintf("users:page:%d", c.QueryInt("page", 1))

    // 尝试从缓存获取
    if cached := cache.Get(cacheKey); cached != nil {
        c.Okay(cached)
        return
    }

    // 查询数据库
    store := model.GetStore("user")
    pageData, err := restapi.Page(c, store, model.Filter{})
    if err != nil {
        c.Fail(500, err.Error())
        return
    }

    // 存储到缓存
    cache.Set(cacheKey, pageData, 5*time.Minute)

    c.Okay(pageData)
}
```

## 最佳实践

### 1. API 设计

- 使用清晰的资源名称
- 保持 URL 简洁和语义化
- 使用合适的 HTTP 方法
- 提供统一的响应格式

### 2. 错误处理

- 提供清晰的错误信息
- 使用合适的 HTTP 状态码
- 记录详细的错误日志
- 避免暴露敏感信息

### 3. 安全考虑

- 实施适当的认证和授权
- 验证和过滤输入数据
- 使用 HTTPS 传输
- 防止常见攻击（XSS、CSRF 等）

### 4. 性能优化

- 使用分页减少数据传输
- 实施查询缓存
- 优化数据库查询
- 使用 CDN 处理静态资源

## 故障排除

### 常见问题

1. **路由不匹配**
   - 检查 API 前缀配置
   - 确认模型名称正确
   - 验证 HTTP 方法

2. **数据验证失败**
   - 检查数据格式
   - 确认必填字段
   - 验证数据类型

3. **文件上传失败**
   - 检查文件大小限制
   - 确认文件类型允许
   - 验证存储路径权限

## 更新日志

### v1.0.0
- 初始版本发布
- 基本 CRUD 操作
- 文件上传功能
- 关联关系处理
- 中间件支持
- 响应钩子功能
