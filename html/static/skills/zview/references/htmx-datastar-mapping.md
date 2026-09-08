# htmx / datastar → zview 对照

当用户是从 `htmx` 或 `datastar` 迁移过来、或想拿 zview 复刻它们的交互时，读这个文件。
所有映射都锚定 zview 真实实现（`packages/view/src/` 与 `docs/zview.md`），不编造 API。

## htmx → zview（超媒体 / AJAX 局部更新）

| htmx | zview 等价物 | 备注 |
|---|---|---|
| `hx-get`/`hx-post`/... | `z-req`（`GET /url`、`POST:url`、`POST /url`） | 默认触发器：`form→submit`、`input/select/textarea→change`、其余→click |
| `hx-trigger` | `z-trigger` | 除 DOM 事件外还内置特殊触发器（分类见下方「特殊触发器」小节）；`every(ms)` 是轮询间隔解析、配合 `load` 等使用，不是独立事件触发器 |
| `hx-target` | `z-target` | 支持 `this/inherit/parent/child/sibling` 与链式 `a\|b\|c` |
| `hx-swap` | `z-swap` | `inner/replace/append/prepend/beforebegin/afterend/morph/morph-all/skip` |
| `hx-indicator` | `z-indicator` | 请求进行中的指示器 |
| `hx-disabled-elt` | `z-disabled-elt` | 请求期间禁用触发元素自身或匹配 selector 的元素；值可为逗号分隔的 CSS 选择器，空值禁用自身 |
| `hx-confirm` | `z-confirm` | 请求发出前弹窗确认（对标 htmx `hx-confirm`）；在带 `z-req` 的元素上加 `z-confirm="提示文案"`，用户取消则不发出请求 |
| `hx-select` | `z-config.selector` | 从 HTML 响应提取片段 |
| `hx-vals` / `hx-include` | `z-data` / 表单自动收集 | 见请求体默认格式 |
| `hx-boost` | 用 `z-req` + `z-history` 手写 | 无整站自动 boost，需在链接上加 `z-req` |
| `hx-push-url`/`hx-replace-url` | `z-history` / `z-history-replace` | 也可由服务端响应头 `Z-History` 触发 |
| 服务端 `HX-Redirect`/`HX-*` 响应头 | 服务端 `Z-Redirect`/`Z-Location`/`Z-Title` 响应头 | `Z-Redirect` 全页跳转，`Z-Location` 客户端导航（SPA 式） |
| 表单提交 AJAX | 表单元素上的 `z-req` 默认接管 `submit` | `z-before` 可 `preventDefault`/返回 `false` 取消请求 |
| 生命周期 | `z-before` / `z-after` / `z-error` / `z-finally` | 请求在触发元素 dispatch 的自定义事件 |

### z-trigger 特殊触发器分类

`z-trigger` 除常规 DOM 事件（如 `click`/`submit`/`change`）外，还内置以下特殊触发器，按触发来源分类：

- **生命周期**：`load`（页面加载完成后执行一次）
- **滚动**：`scroll`（监听窗口/页面滚动）、`windowScroll`（监听触发元素自身的滚动）
- **外部点击**：`outside`（点击 document 中触发元素之外的位置）
- **可见性（IntersectionObserver）**：`visible`（进入视口）/ `invisible`（完全离开）/ `closeby`（进入半屏邻近区）/ `away`（离开半屏邻近区）
- **DOM 变化（MutationObserver）**：`remove`（元素被移除）/ `childrenChange`（子节点变动）/ `empty`（子节点清空）/ `notempty`（出现子节点）

> `every(ms)` 是轮询**间隔解析**，与上述某个触发器（常用 `load`）配合使用，表示“每隔 ms 触发一次”；它不是独立事件触发器。

zview 没有的 htmx 概念（暂不要假装支持）：

- `hx-swap-oob`（同一响应里多处 out-of-band 交换）
- swap 的 `settle`/CSS 过渡类（`htmx-swapping` 等）；zview 目前只 dispatch `z-swap-before`/`z-swap-after`

## datastar → zview（信号 / 后端驱动）

| datastar | zview 等价物 | 备注 |
|---|---|---|
| `data-signals` | `z-signals` | JSON 初始化信号 |
| `@click` / 事件绑定 | `z-on-click` | 含修饰符：`.prevent/.stop/.self/.outside/.window`、`.enter/.right`、`.debounce.300/.delay.300/.throttle.500` |
| `data-bind` | `z-bind` | 表单双向绑定 |
| 计算值 | `z-computed` / `z-computed-{name}` | `z-computed-total="$price * $qty"` |
| 文本/条件渲染 | `z-text` / `z-show` | |
| 客户端从后端拉数据写信号 | `z-req` + `z-json` / `z-json-*` | 设置 `Accept: application/json`，把响应写入信号并跳过 swap |
| 后端驱动流式更新（SSE） | 无 | zview 基于 XHR 请求/响应，不是 SSE 长连接 |

## 迁移时应给用户的关键提醒

1. 别写冒号写法 `z-on:click` / `z-bind:value`；zview 是连字符 `z-on-click`。
2. 事件表达式里 `event`、`target`、`currentTarget` 都可用（也有 `$event`/`$target`/`$currentTarget`）。
3. `z-show` 控制的是 display 显隐，没有 `z-if`。
4. 想让某个交互区不换页而是局部刷新：`z-req` + `z-target` + `z-swap`，不要手写 fetch。
5. swap 进来的新内容会被 zview 重新激活；若没生效先查返回 markup 的指令名与连字符写法。
