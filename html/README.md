# HTML 渲染模块

`html` 模块提供基于 Go 的服务端 HTML 组件化 DSL，可与 `znet` 路由和 `service.Module` 生命周期集成，用于构建结构化页面，适用于纯后端渲染或半交互式场景。

## 功能特性

- **声明式组件 DSL**：通过 `pkg/app_module/html/el/` 提供的元素与属性构造器在 Go 代码中声明 DOM 结构
- **HTMX 2.x 集成**：`pkg/app_module/html/htmx/` 提供 HTMX 属性常量、事件属性与 JSON 属性构造器
- **CSS 样式支持**：`pkg/app_module/html/styles/` 提供行内样式、CSS 值辅助函数与样式表管理器
- **增强的属性系统**：`Attr`/`Data` 支持字符串、数字、布尔值和 Map，自动处理 JSON 序列化与 HTML 转义
- **模块化接入**：实现 `service.ModuleLifeCycle`，注册后自动绑定渲染器与依赖注入
- **多返回模式**：支持直接返回 `*el.Element`、状态码+元素、带 `error` 的返回及 `ZViewJS` 增强模式
- **ZView.js 集成**：支持局部更新、重定向、历史记录管理等前后端协同交互
- **错误兜底**：通过 `Options.ErrorPage` 指定备用页面，渲染失败时自动回退

## 模块结构

```
html/
├── el/                 # HTML 元素 DSL 与渲染引擎
├── htmx/               # HTMX 2.x 属性常量与构造器
├── styles/             # CSS 声明、值辅助函数与 StyleManager
├── zview/              # ZView 请求上下文与交换策略
├── static/             # 静态资源源文件（zcss.js / zview.js）
├── module.go           # 模块生命周期定义
├── options.go          # Options 配置
├── render.go           # 渲染器注册与实现
├── gen.go              # 静态资源生成脚本（go generate）
├── static_data.go      # 由脚本生成的静态资源 Go 文件（勿手改）
├── static.go           # 静态资源路由注册
└── static_disabled.go  # nostatic 构建标签下的空实现
```

## 快速开始

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

| 字段           | 类型          | 说明                                          | 默认值           |
| -------------- | ------------- | --------------------------------------------- | ---------------- |
| `ErrorPage`    | `*el.Element` | 渲染失败时返回的兜底页面                      | `nil`            |
| `StaticPrefix` | `string`      | zcss/zview 静态资源的路由前缀（空则用默认值） | `/__static_html` |

> `Options.DisableWrite()` 返回 `true`，该模块不会将配置写入文件，可直接在模块注册时通过代码注入。

## 渲染模式

渲染器会设置 `Content-Type: text/html`。当处理函数返回 `error` 或元素渲染失败时：若配置了 `Options.ErrorPage`，会以 500 状态输出兜底页面；否则将错误交给 znet 的错误处理流程。

### 基础模式

- `func(c *znet.Context) *el.Element` — 直接返回元素，状态码为 200
- `func(c *znet.Context) (*el.Element, error)` — 返回元素或错误，错误时使用 ErrorPage
- `func(c *znet.Context) (int, *el.Element)` — 返回自定义状态码与元素

```go
// 直接返回元素
r.GET("/dashboard", func(c *znet.Context) *el.Element {
    return el.HTML(
        el.HEAD(el.TITLE(el.Text("Dashboard"))),
        el.BODY(el.H1(el.Text("仪表盘"))),
    )
})

// 可返回 error
r.GET("/user", func(c *znet.Context) (*el.Element, error) {
    user, err := getUser(c)
    if err != nil {
        return nil, err
    }
    return el.DIV(el.Text(user.Name)), nil
})

// 自定义状态码
r.GET("/error", func(c *znet.Context) (int, *el.Element) {
    return 404, el.HTML(el.BODY(el.H1(el.Text("页面未找到"))))
})
```

## 元素 DSL 示例

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

## HTMX 局部更新

HTMX 属性可通过 `html/htmx` 子包类型安全地加入元素。该子包只负责生成属性，不会自动加载 HTMX 脚本；请在页面中按需引入 HTMX 2.x。

```go
import (
    "github.com/zlsgo/app_module/html/el"
    "github.com/zlsgo/app_module/html/htmx"
)

page := el.DIV(
    el.BUTTON(
        htmx.Get("/users"),
        htmx.Target("#users"),
        htmx.Swap("innerHTML"),
        el.Text("加载用户"),
    ),
    el.DIV(el.ID("users"), el.Text("尚未加载")),
)
```

更多属性与示例见 [`html/htmx/README.md`](./htmx/README.md)。

## CSS 样式

`html/styles` 是独立的服务端 CSS 辅助包，不依赖 HTMX，也不替代 `html/static` 中的前端 `zcss.js`。

```go
import "github.com/zlsgo/app_module/html/styles"

buttonStyle := styles.Props{
    styles.BackgroundColor: "#2563eb",
    styles.Color:           "white",
    styles.Padding:         styles.Pixels(12) + " " + styles.Pixels(20),
}

button := el.BUTTON(buttonStyle.Attr(), el.Text("保存"))
```

需要复用 class、伪类、媒体查询或 keyframes 时使用 `StyleManager`：

```go
manager := styles.NewStyleManager()
buttonClass := manager.AddCompositeStyle(styles.CompositeStyle{
    Default: styles.Props{
        styles.BackgroundColor: "#2563eb",
        styles.Color:           "white",
    },
    PseudoClasses: map[string]styles.Props{
        styles.PseudoHover: {styles.BackgroundColor: "#1d4ed8"},
    },
})

page := el.HTML(
    el.HEAD(manager.StyleTag()),
    el.BODY(el.BUTTON(el.Class(buttonClass), el.Text("保存"))),
)
```

更多说明见 [`html/styles/README.md`](./styles/README.md)。

## 增强的属性系统

`Attr` 和 `Data` 支持字符串、数字、布尔值和 `ztype.Map`（自动 JSON 序列化）。所有属性值在渲染阶段统一执行 HTML 转义，`Attribute.Value` 与 `GetAttribute` 返回未经转义的原始值。

```go
config := ztype.Map{"timeout": 3000, "endpoint": "/api/data"}
el.DIV(el.Data("config", config))

el.INPUT(el.Attr("disabled", true), el.Attr("checked", false))
```

常用通用工具：`Attrs`（稳定排序）、`On`、`SetAttribute`/`RemoveAttribute`/`HasAttribute`、`BareAttr`。渲染不会自动释放元素树，可重复渲染；不再使用时可调用顶层节点 `Release()`。

## ZView.js 增强交互

在高级模式下可注入 `*html.ZViewJS`（`*html/zview.Context` 的兼容别名）实现前后端协同交互：

```go
r.GET("/partial", func(c *znet.Context, z *html.ZViewJS) *el.Element {
    if z.Is() {
        z.SetTitle("更新后的标题")
        return el.DIV(el.Text("局部更新内容"))
    }

    return el.HTML(
        el.HEAD(el.TITLE(el.Text("完整页面"))),
        el.BODY(el.DIV(el.Text("完整内容"))),
    )
})
```

### 常用 API

| 方法 | 说明 |
| ---- | ---- |
| `Is()` / `IsPartial()` | 判断是否为 ZView 请求 / 局部更新请求 |
| `GetUrl(fallback...)` / `GetOrigin()` / `GetTarget()` | 读取请求头信息 |
| `SetTitle(title)` | 设置页面标题（`Z-Title`） |
| `SetRedirect(url)` / `SetLocation(url)` | 重定向或客户端导航 |
| `SetHistory(url)` / `SetHistoryReplace(url)` | 更新浏览器历史 |
| `SetSwap(value)` / `SetTarget(value)` | 控制替换策略与目标选择器 |

## 静态资源（zcss / zview）

前端交互依赖 `zcss.js` 与 `zview.js`。源码位于 `html/static/`，通过 `go generate ./html` 生成 `html/static_data.go`，由 znet 直接输出。

- 默认路由前缀：`/__static_html`（可通过 `Options.StaticPrefix` 自定义）
  - `/__static_html/zcss.js`
  - `/__static_html/zview.js`
- 使用 `-tags nostatic` 可剔除内置静态资源与路由

更多说明见 [`html/static/README.md`](./static/README.md)。

## 最佳实践

1. **封装常用结构**：将重复的元素组合提炼为函数，保持代码可读性。
2. **规划错误回退**：为 `ErrorPage` 提供明确提示，避免渲染异常导致空白页面。
3. **按需引入前端能力**：HTMX/ZView 静态资源按需加载，生产环境可结合 `nostatic` 与外部资源策略。
4. **结合其他模块**：与 `account`、`restapi` 等模块协同，实现纯后端页面与接口的统一管理。