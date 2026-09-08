# Template 模块

Template 模块基于 [CloudyKit/jet](https://github.com/CloudyKit/jet) 为 `znet` 应用提供服务端模板渲染能力，与 `service.Module` 生命周期集成，支持模板目录、静态资源、自定义函数与热重载。

## 功能特性

- 🚀 基于 `jet` 引擎的模板渲染（`{{: ... }}` 定界符）
- 📁 模板目录加载（默认 `./views`）
- 📦 静态资源托管（`Static` / `StaticDir`）
- 🛠️ 自定义模板函数（`Options.Funcs`）
- 🔄 调试/开发模式热重载（`Reload`）
- 📄 页面渲染与字节返回两种形态

## 快速开始

```go
package main

import (
    "github.com/sohaha/zlsgo/znet"
    "github.com/zlsgo/app_core/service"
    "github.com/zlsgo/app_module/template"
)

func main() {
    app := service.NewApp()(nil)

    tmplMod := template.New(func(o *template.Options) {
        o.Dir = "./views"      // 模板目录
        o.Static = "/static"   // 静态资源路由前缀
        o.StaticDir = "./public"
        o.Funcs = map[string]interface{}{
            "upper": func(s string) string { return strings.ToUpper(s) },
        }
    })

    err := service.InitModule([]service.Module{tmplMod}, app)
    if err != nil {
        panic(err)
    }

    service.RunWeb(app)
}
```

在控制器中渲染：

```go
func index(c *znet.Context) error {
    return tmplMod.Template(c, "index", ztype.Map{"title": "Home"})
}
```

## 配置项（Options）

| 字段       | 类型                   | 说明                                   |
| ---------- | ---------------------- | -------------------------------------- |
| `Funcs`    | `map[string]interface{}` | 自定义模板函数，注册到 jet 引擎          |
| `Dir`      | `string`               | 模板目录（默认 `./views`）              |
| `Static`   | `string`               | 静态资源路由前缀（配合 `StaticDir` 生效）|
| `StaticDir`| `string`               | 静态资源本地目录                        |
| `Reload`   | `bool`                 | 是否热重载模板（调试模式下自动开启）      |

配置项 `ConfKey` 为 `templates`，可通过配置文件覆盖，且不参与写回（`DisableWrite`）。

## 渲染 API

- `Render(template string, data interface{}, layout ...string) ([]byte, error)`：渲染并返回字节
- `Template(c *znet.Context, template string, data interface{}, layout ...string) error`：渲染并直接以 HTML 输出

## 定界符

引擎使用非默认定界符 `{{: ... }}` 与 `}}`，避免与前端框架（如 Vue、Angular 的 `{{ }}`）冲突，适合在 HTML 中混用前端模板。

## 使用约定

- 若 `Dir` 对应目录不存在，启动时静默跳过加载（不报错）。
- 模板加载失败（目录存在但解析出错）会在启动阶段返回错误。
- 静态托管仅在同时设置 `Static` 与 `StaticDir` 时启用。
