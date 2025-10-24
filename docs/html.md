# HTML 渲染模块

HTML 模块提供基于 Go 语言的 HTML 组件化 DSL，与 `znet` 路由和 `service.Module` 生命周期集成，可在服务端生成结构化页面，满足纯后端渲染或半交互式场景需求。

## 功能特性

- **声明式组件 DSL**: 借助 `pkg/app_module/html/el/` 提供的大量元素与属性构造器，在 Go 代码中声明 DOM 结构。
- **模块化接入**: `Module` 实现 `service.ModuleLifeCycle`，注册后自动绑定渲染器与依赖注入。
- **多返回模式**: 支持直接返回 `*el.Element`、状态码与元素组合，以及注入 `html.ZHTML` 的高级用法。
- **错误兜底**: 通过 `Options.ErrorPage` 指定备用页面，渲染失败时自动回退，提升稳健性。

## 模块结构

```
html/
├── el/          # HTML 元素 DSL 与渲染引擎
├── module.go    # 模块生命周期定义
├── options.go   # Options 配置
├── render.go    # 渲染器注册与实现
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

| 字段        | 类型          | 说明                     | 默认值 |
| ----------- | ------------- | ------------------------ | ------ |
| `ErrorPage` | `*el.Element` | 渲染失败时返回的兜底页面 | `nil`  |

> `Options.DisableWrite()` 返回 `true`，该模块不会将配置写入文件，可直接在模块注册时通过代码注入。

## 渲染模式

`pkg/app_module/html/render.go` 注册了常见的处理函数签名：

- `func(c *znet.Context) *el.Element`
- `func(c *znet.Context) (int, *el.Element)`


```go
r.GET("/dashboard", func(c *znet.Context) *el.Element {
    return el.HTML(
        el.HEAD(el.TITLE(el.Text("Dashboard"))),
        el.BODY(
            el.H1(el.Text("仪表盘")),
            el.DIV(el.Text("完整页面内容")),
        ),
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
