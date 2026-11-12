# Template 模块

`template` 模块基于 [CloudyKit/jet](https://github.com/CloudyKit/jet) 为 `znet` 应用提供服务端模板渲染能力，并与 `service.Module` 生命周期集成，支持模板目录加载、静态资源托管、自定义函数与热重载。

## 功能特性

- 🚀 基于 `jet` 引擎的模板渲染（使用 `{{: ... }}` 定界符）
- 📁 模板目录加载（默认 `./views`）
- 📦 静态资源托管（`Static` / `StaticDir`）
- 🛠️ 自定义模板函数（`Options.Funcs`）
- 🔄 调试/开发模式热重载（`Reload`）
- 📄 支持页面渲染与字节返回两种形态

## 快速开始

```go
package main

import (
    "strings"

    "github.com/sohaha/zlsgo/znet"
    "github.com/sohaha/zlsgo/ztype"
    "github.com/zlsgo/app_core/service"
    "github.com/zlsgo/app_module/template"
)

func main() {
    app := service.NewApp()(nil)

    tmplMod := template.New(func(o *template.Options) {
        o.Dir = "./views"       // 模板目录
        o.Static = "/static"    // 静态资源路由前缀
        o.StaticDir = "./public"
        o.Funcs = map[string]interface{}{
            "upper": func(s string) string { return strings.ToUpper(s) },
        }
    })

    if err := service.InitModule([]service.Module{tmplMod}, app); err != nil {
        panic(err)
    }

    _ = service.Global.DI.InvokeWithErrorOnly(func(r *znet.Engine) {
        r.GET("/", func(c *znet.Context) error {
            return tmplMod.Template(c, "index", ztype.Map{"title": "Home"})
        })
    })

    service.RunWeb(app)
}
```

## 配置项（Options）

| 字段        | 类型                     | 说明                                      |
| ----------- | ------------------------ | ----------------------------------------- |
| `Funcs`     | `map[string]interface{}` | 自定义模板函数，注册到 jet 引擎           |
| `Dir`       | `string`                 | 模板目录（默认 `./views`）                |
| `Static`    | `string`                 | 静态资源路由前缀（配合 `StaticDir` 生效） |
| `StaticDir` | `string`                 | 静态资源本地目录                          |
| `Reload`    | `bool`                   | 是否热重载模板（调试模式下自动开启）      |

> 配置项 `ConfKey` 为 `templates`，可通过配置文件覆盖，且不参与写回（`DisableWrite`）。

## 渲染 API

- `Render(template string, data interface{}, layout ...string) ([]byte, error)`：渲染并返回字节
- `Template(c *znet.Context, template string, data interface{}, layout ...string) error`：渲染并直接以 HTML 输出

示例：

```go
// 返回字节，适合自定义响应处理
html, err := tmplMod.Render("index", ztype.Map{"title": "Home"})
if err != nil {
    return err
}

// 直接输出 HTML 响应
return tmplMod.Template(c, "index", ztype.Map{"title": "Home"})
```

## 定界符

引擎使用非默认定界符 `{{: ... }}` 与 `}}`，避免与前端框架（如 Vue、Angular 的 `{{ }}`）冲突，适合在 HTML 中混用前端模板语法。

## 使用约定

- 若 `Dir` 对应目录不存在，启动时静默跳过加载（不报错）。
- 模板加载失败（目录存在但解析出错）会在启动阶段返回错误。
- 静态托管仅在同时设置 `Static` 与 `StaticDir` 时启用。

## 最佳实践

- 生产环境建议关闭热重载（`Reload=false`）以提升性能与稳定性。
- 将公共布局与片段模板拆分管理，保持页面模板简洁。
- 自定义函数尽量保持纯函数语义，避免在模板中引入副作用。
- 静态资源目录与模板目录分离，便于缓存策略与部署管理。