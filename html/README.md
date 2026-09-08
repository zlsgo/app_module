# HTML 渲染模块

HTML 模块提供基于 Go 语言的 HTML 组件化 DSL，与 `znet` 路由和 `service.Module` 生命周期集成，可在服务端生成结构化页面，满足纯后端渲染或半交互式场景需求。

## 功能特性

- **声明式组件 DSL**: 借助 `pkg/app_module/html/el/` 提供的大量元素与属性构造器，在 Go 代码中声明 DOM 结构。
- **增强的属性系统**: `Attr` 和 `Data` 函数支持多种类型（string、Map、bool），自动处理 JSON 序列化和转义。
- **模块化接入**: `Module` 实现 `service.ModuleLifeCycle`，注册后自动绑定渲染器与依赖注入。
- **多返回模式**: 支持直接返回 `*el.Element`、状态码与元素组合，以及注入 `html.ZViewJS` 的高级用法。
- **ZView.js 集成**: 前后端协同的增强交互能力，支持局部更新、重定向、历史记录管理等。
- **错误兜底**: 通过 `Options.ErrorPage` 指定备用页面，渲染失败时自动回退，提升稳健性。

## 模块结构

```
html/
├── el/            # HTML 元素 DSL 与渲染引擎
├── module.go      # 模块生命周期定义
├── options.go     # Options 配置
├── render.go      # 渲染器注册与实现
├── static/        # 静态资源源文件（zcss.js / zview.js）
├── gen.go         # 静态资源生成脚本（go generate）
├── static_data.go # 由脚本生成的静态资源 Go 文件（勿手改）
└── static.go      # 静态资源路由注册（znet 直接输出）
```

## 快速开始

### 注册模块并输出页面

```go
package main

import (
    "github.com/sohaha/zlsgo/znet"
    "github.com/zlsgo/app_core/service"
    "github.com/zlsgo/app_module/html"
    "github.com/zlsgo/app_module/html/el"
)

func main() {
    app := service.NewApp()(nil)

    // 注册 HTML 模块并配置错误页兜底
    htmlMod := html.New(func(o *html.Options) {
        o.ErrorPage = el.HTML(
            el.HEAD(el.TITLE(el.Text("Render Error"))),
            el.BODY(
                el.H1(el.Text("服务器渲染失败")),
                el.P(el.Text("请稍后重试")),
            ),
        )
    })

    if err := service.InitModule([]service.Module{htmlMod}, app); err != nil {
        panic(err)
    }

    _ = app.DI.InvokeWithErrorOnly(func(r *znet.Engine) {
        r.GET("/", func(c *znet.Context) *el.Element {
            return el.HTML(
                el.HEAD(el.TITLE(el.Text("首页"))),
                el.BODY(
                    el.H1(el.Text("Hello HTML")),
                    el.P(el.Text("该页面由服务器渲染")),
                ),
            )
        })
    })

    service.RunWeb(app)
}
```

## 配置说明

`pkg/app_module/html/options.go` 定义了可选项：

| 字段           | 类型          | 说明                                          | 默认值          |
| -------------- | ------------- | --------------------------------------------- | --------------- |
| `ErrorPage`    | `*el.Element` | 渲染失败时返回的兜底页面                      | `nil`           |
| `StaticPrefix` | `string`      | zcss/zview 静态资源的路由前缀（空则用默认值） | `/__static_html` |

> `Options.DisableWrite()` 返回 `true`，该模块不会将配置写入文件，可直接在模块注册时通过代码注入。

## 渲染模式

`pkg/app_module/html/render.go` 注册了常见的处理函数签名：

### 基础模式

- `func(c *znet.Context) *el.Element` — 直接返回元素，状态码为 200
- `func(c *znet.Context) (*el.Element, error)` — 返回元素或错误，错误时使用 ErrorPage
- `func(c *znet.Context) (int, *el.Element)` — 返回自定义状态码与元素

```go
// 基础模式 - 直接返回元素
r.GET("/dashboard", func(c *znet.Context) *el.Element {
    return el.HTML(
        el.HEAD(el.TITLE(el.Text("Dashboard"))),
        el.BODY(
            el.H1(el.Text("仪表盘")),
            el.DIV(el.Text("完整页面内容")),
        ),
    )
})

// 错误处理模式 - 可返回 error
r.GET("/user", func(c *znet.Context) (*el.Element, error) {
    user, err := getUser(c)
    if err != nil {
        return nil, err
    }
    return el.DIV(el.Text(user.Name)), nil
})

// 自定义状态码模式
r.GET("/error", func(c *znet.Context) (int, *el.Element) {
    return 404, el.HTML(
        el.BODY(el.H1(el.Text("页面未找到"))),
    )
})
```

## 元素 DSL 示例

`pkg/app_module/html/el` 提供大量元素、属性与工具函数，可组合成组件：

```go
snippet := el.DIV(
    el.Class("card"),
    el.H2(el.Text("标题")),
    el.P(el.Textf("当前时间: %s", time.Now().Format(time.RFC3339))),
)

htmlBytes, err := el.RenderBytes(context.Background(), snippet)
if err != nil {
    panic(err)
}
```

## 最佳实践

1. **封装常用结构**: 将重复出现的元素组合提炼为函数，保持代码可读性。
2. **规划错误回退**: 为 `ErrorPage` 提供明确提示，避免渲染异常导致空白页面。
3. **结合其他模块**: 与 `account`、`restapi` 等模块协同，实现纯后端页面与接口的统一管理。

## 增强的属性系统

### 泛型属性支持

`Attr` 和 `Data` 函数现在支持多种类型，自动处理类型转换和序列化：

```go
// 字符串属性（传统方式）
el.DIV(el.Attr("title", "提示文本"))

// Map 属性（自动 JSON 序列化，渲染时统一转义）
config := ztype.Map{
    "timeout": 3000,
    "retries": 3,
    "endpoint": "/api/data",
}
el.DIV(el.Data("config", config))
// 输出: <div data-config="{&#34;timeout&#34;:3000,&#34;retries&#34;:3,&#34;endpoint&#34;:&#34;/api/data&#34;}"></div>

// 字符串形式的 JSON 同样被正确转义，不会破坏属性边界
el.DIV(el.Attr("z-config", `{"format":"urlencoded"}`))
// 输出: <div z-config="{&#34;format&#34;:&#34;urlencoded&#34;}"></div>

// 布尔属性
el.INPUT(
    el.Attr("disabled", true),
    el.Attr("checked", false),
)
```

**类型约束**: `AttrValue` 接口限定为 `string | ztype.Map | bool`

**自动处理**:
- `string`: 直接使用原始值
- `bool`: 转换为字符串 `"true"` 或 `"false"`
- `ztype.Map`: JSON 序列化，失败时使用 `"{}"`

**统一转义**: 所有属性值（含 `SetAttribute`、`DeferredAttr`）都在渲染输出阶段统一执行 HTML 转义，`Attribute.Value` 与 `GetAttribute` 始终返回未经转义的原始值，因此值中包含双引号、尖括号等字符（例如 JSON 文本）时也能保证最终 HTML 合法。浏览器解析后会还原为原始值，不影响前端读取。

## ZView.js 增强交互

### 注入 ZViewJS 对象

在高级模式下，可注入 `*html.ZViewJS` 来实现前后端协同交互：

```go
r.GET("/partial", func(c *znet.Context, z *html.ZViewJS) *el.Element {
    // 检查是否为 ZView.js 请求
    if z.Is() {
        // 设置页面标题
        z.SetTitle("更新后的标题")

        // 返回局部内容
        return el.DIV(el.Text("局部更新内容"))
    }

    // 完整页面
    return el.HTML(
        el.HEAD(el.TITLE(el.Text("完整页面"))),
        el.BODY(el.DIV(el.Text("完整内容"))),
    )
})
```

### ZViewJS API 方法

| 方法 | 说明 | 示例 |
|------|------|------|
| `Is()` | 判断是否为 ZView.js 请求 | `if z.Is() { ... }` |
| `IsPartial()` | 判断是否为局部更新请求 | `if z.IsPartial() { ... }` |
| `GetUrl(fallback...)` | 读取 `Z-Url`，可选回退值 | `path := z.GetUrl("/")` |
| `GetOrigin()` | 读取 `Z-Origin` 请求来源 | `origin := z.GetOrigin()` |
| `GetTarget()` | 读取 `Z-Target` 目标选择器 | `target := z.GetTarget()` |
| `SetTitle(title)` | 设置页面标题（响应头 `Z-Title`） | `z.SetTitle("新标题")` |
| `SetRedirect(url)` | 发送 `Z-Redirect` 头；非 ZView 请求会执行标准重定向 | `z.SetRedirect("/login")` |
| `SetLocation(url)` | 发送 `Z-Location` 头进行客户端导航 | `z.SetLocation("/dashboard")` |
| `SetHistory(url)` | 发送 `Z-History` 头更新浏览器历史 | `z.SetHistory("/new-url")` |
| `SetHistoryReplace(url)` | 发送 `Z-History-Replace` 头替换当前历史 | `z.SetHistoryReplace("/current")` |
| `SetSwap(value)` | 控制内容替换策略 | `z.SetSwap("inner")` |
| `SetTarget(value)` | 指定目标选择器（响应头 `Z-Target`） | `z.SetTarget("#panel")` |

### 响应头规范

所有 ZView.js 相关的响应头统一使用大写格式：

- `Z-Title`: 页面标题
- `Z-Redirect`: 重定向 URL
- `Z-Location`: 客户端导航 URL
- `Z-Origin`: 请求来源
- `Z-Target`: 目标元素选择器
- `Z-History`: 历史记录 URL
- `Z-History-Replace`: 替换当前历史的 URL
- `Z-Swap`: 内容替换策略

### 实际应用示例

```go
// 表单提交后的处理
r.POST("/submit", func(c *znet.Context, z *html.ZViewJS) (int, *el.Element) {
    // 处理表单数据
    if err := processForm(c); err != nil {
        return 400, el.DIV(
            el.Class("error"),
            el.Text("提交失败: " + err.Error()),
        )
    }

    // ZView.js 请求 - 局部更新
    if z.Is() {
        z.SetTitle("提交成功")
        z.SetHistory("/success")
        return 200, el.DIV(
            el.Class("success"),
            el.Text("提交成功！"),
        )
    }

    // 标准请求 - 重定向
    z.SetRedirect("/success")
    return 302, nil
})
```

## 静态资源（zcss / zview）

HTML 模块的前端交互依赖 `zcss.js` 与 `zview.js` 两个静态脚本。
它们通过脚本直接转换成 Go 源文件引入（不使用 `//go:embed` 内嵌方式），并由 znet 直接输出，方便自定义路由前缀或按需裁剪。

### 脚本生成

源码存放于 `html/static/`（`zcss.js`、`zview.js`），生成器为 `html/gen.go`：

```bash
go generate ./html
```

运行后生成 `html/static_data.go`，内含两份资源的名称与内容字节，无需再读取磁盘或 embed 目录。

> 变更 `html/static/*.js` 后请重新执行 `go generate ./html` 同步产物。

### 自定义前缀路由

默认路由前缀为 `/__static_html`，对应地址：

- `/__static_html/zcss.js`
- `/__static_html/zview.js`

如需自定义前缀，通过 `Options.StaticPrefix` 配置：

```go
htmlMod := html.New(func(o *html.Options) {
    o.StaticPrefix = "/assets" // 前端引用地址需同步
})
```

### go tag 控制是否引入

默认构建会包含这两份静态资源。若无需内置（例如由外部 CDN 加载、或追求最小体积），使用构建标签 `nostatic` 即可完全剔除生成的数据文件与对应静态路由：

```bash
go build -tags nostatic ./...
```

> 使用 `-tags nostatic` 构建后，`zcss.js` / `zview.js` 的资源字节与静态路由均不会进入产物，仅保留 zview.Context 依赖注入等渲染能力。
