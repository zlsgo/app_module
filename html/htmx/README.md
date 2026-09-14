# HTML HTMX 集成

`html/htmx` 为 `github.com/zlsgo/app_module/html/el` 提供 HTMX 2.x 属性常量与便捷构造器。它只负责生成 HTML 属性，不会自动注入或打包 HTMX JavaScript；页面仍需自行引入 HTMX（例如官方 CDN）或通过项目静态资源提供。

HTMX 专用属性的语义只在本包中定义；底层 `html/el` 仅提供通用的
`Attr`/`BareAttr` 属性渲染能力，不维护 `hx-*`、SSE 或 WebSocket 属性名单。

例如，页面可以在 `head` 中引入 HTMX 2.x：

```go
el.SCRIPT(el.Src("https://cdn.jsdelivr.net/npm/htmx.org@2.0.10/dist/htmx.min.js"))
```

具体版本请按应用的依赖锁定策略调整。

## 使用

```go
import (
	"github.com/zlsgo/app_module/html/el"
	"github.com/zlsgo/app_module/html/htmx"
)

content := el.DIV(
	htmx.Get("/todos"),
	htmx.Target("#todo-list"),
	htmx.Swap("innerHTML"),
	htmx.Trigger("click"),
	el.Text("刷新"),
)
```

也可以直接使用常量与 `el.Attr`：

```go
el.BUTTON(
	el.Attr(htmx.HXPost, "/todos"),
	el.Attr(htmx.HXConfirm, "确定删除吗？"),
	el.Attr(htmx.HXTarget, "closest li"),
	el.Text("删除"),
)
```

## 支持的属性

包提供 HTMX 2.x 核心属性（`HXGet`、`HXPost`、`HXPushURL`、`HXSelect`、`HXSelectOOB`、`HXSwap`、`HXSwapOOB`、`HXTarget`、`HXTrigger`、`HXVals`）与附加属性（`HXBoost`、`HXConfirm`、`HXDelete`、`HXDisable`、`HXDisabledElt`、`HXDisinherit`、`HXEncoding`、`HXExt`、`HXHeaders`、`HXHistory`、`HXHistoryElt`、`HXInclude`、`HXIndicator`、`HXInherit`、`HXParams`、`HXPatch`、`HXPreserve`、`HXPrompt`、`HXPut`、`HXReplaceURL`、`HXRequest`、`HXSync`、`HXValidate`）。

SSE 和 WebSocket 扩展属性也有常量：`SSEConnect`、`SSESwap`、`SSEClose`、`WSConnect`、`WSSend`。这些扩展仍需额外加载对应的 HTMX 扩展脚本。

所有 HTMX 事件属性均以 `HXOn...` 常量提供，例如 `HXOnBeforeRequest`、`HXOnAfterSwap`。动态事件可使用：

```go
el.BUTTON(
	htmx.On("htmx:beforeRequest", "console.log(event)"),
	el.Text("请求"),
)
```

`On` 会把 `htmx:beforeRequest`、`beforeRequest`、`before-request`、`hx-on::before-request` 和 `hx-on-htmx-before-request` 等形式规范化为 HTMX 2.x 的 `hx-on--before-request`。普通 DOM 事件则使用单连字符形式，例如 `On("click", "...")` 输出 `hx-on-click`，`On("hx-on:click", "...")` 输出同样的结果。

## JSON 属性

`hx-vals`、`hx-headers`、`hx-request` 等支持 JSON 的属性可以传入 `ztype.Map`：

```go
import "github.com/sohaha/zlsgo/ztype"

el.DIV(
	htmx.Vals(ztype.Map{"page": 2, "filter": "open"}),
	htmx.Request(ztype.Map{"timeout": 1000}),
)
```

`el.Attr` 会在最终渲染阶段统一 HTML 转义属性值，因此不需要手动转义 JSON。若传入字符串形式的 JSON，保持 JSON 原文即可；渲染结果会安全地包含在属性引号内。

## HTMX 版本

本包按 HTMX 2.x 参考实现。`HXVars`、`HXOn`、`HXSSE`、`HXWS` 等旧版常量仍保留以便迁移，但已标记为过时，优先使用 `HXVals`、`HXOn...`、`SSE*` 和 `WS*`。

参考：

- [HTMX reference](https://htmx.org/reference/)
- [hx-on attribute](https://htmx.org/attributes/hx-on/)
- [elem-go htmx README](https://github.com/chasefleming/elem-go/blob/main/htmx/README.md)
