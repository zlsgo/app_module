# zview 指令注意点

当用户在排查语法、请求行为或 swap 后内容时，读这个文件。

## 属性语法

- 使用连字符属性，例如 `z-on-click`、`z-bind-email`、`z-class-active`。
- 不要输出 `z-on:click`、`z-bind:value` 这类写法。

## Signal 访问

- 使用 `$count`、`$user.name`、`$z.$refs.form`、`$z.$all()`。
- `z-for` 注入的循环变量是 `item`、`index` 这类普通变量名，不是 `$item`。

## 事件处理

- 处理器上下文里有 `event`、`target`、`currentTarget`，以及 `$event`、`$target`、`$currentTarget`。
- 如果处理器返回 `false` 且事件可取消，则会阻止默认行为。

## 请求与 swap

- HTML-first 的请求流优先使用 `z-req`。
- 配合 `z-trigger`、`z-target`、`z-swap`、`z-config` 一起表达，不要凭空再包一层命令式 fetch 胶水代码。
- 需要在请求发出前弹窗确认（对标 htmx `hx-confirm`）时，在带 `z-req` 的元素上加 `z-confirm="提示文案"`；用户取消则该次请求不会发出。
- 请求期间需要禁用触发元素自身或匹配元素（对标 htmx `hx-disabled-elt`）时，用 `z-disabled-elt`；值可为逗号分隔的 CSS 选择器，空值禁用自身。
- swap 进来的内容会被 zview 重新激活；如果还不生效，优先检查返回的 markup 和指令名称。
- 替换/插入等 swap 会先对旧子树统一退场：旧 `$refs`、监听器、观察器与轮询会随 `Zview.deactivate` 语义被清理，不会因 DOM 替换残留。
- 请求被 `z-before` 取消时，`z-finally` 仍会触发，`aria-busy`/`z-active` 状态也会复位；不要为了“取消就跳过收尾”而依赖异常路径。

## 初始化顺序

- `window.$z` 在 `Zview.init()` 之前不存在。
- 注册自定义指令或读取 `$z` 之前，先调用 `Zview.init()`。

## HTML 渲染与安全边界

- `z-html` 只用了轻量净化规则。
- 表达式求值基于 `new Function`；不要把它描述成隔离执行，也不要说它适合不可信输入。

## 常见排障清单

- CDN 脚本是否真的从 `@sohaha/zview` 加载？
- 页面依赖 `data-config` 时，是否同时带了 `data-init`？
- 指令是否使用了连字符语法？
- 写入和读取的 signal 名称是否一致？
- 如果是手动插入的 DOM，是否在需要时调用了 `Zview.activate(node)`？
