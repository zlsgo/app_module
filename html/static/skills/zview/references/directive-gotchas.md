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
- 需要在请求发出前弹窗确认时，在带 `z-req` 的元素上加 `z-confirm="提示文案"`；用户取消则该次请求不会发出。
- 请求期间需要禁用触发元素自身或匹配元素时，用 `z-disabled-elt`；值可为逗号分隔的 CSS 选择器，空值禁用自身。
- swap 进来的内容会被 zview 重新激活；如果还不生效，优先检查返回的 markup 和指令名称。
- 替换/插入等 swap 会先对旧子树统一退场：旧 `$refs`、监听器、观察器与轮询会随 `Zview.deactivate` 语义被清理，不会因 DOM 替换残留。
- 请求被 `z-before` 取消时，`z-finally` 仍会触发，`aria-busy`/`z-active` 状态也会复位；不要为了“取消就跳过收尾”而依赖异常路径。

## z-preserve（保留节点）

- 语义：**响应片段里的 preserve 标记**才是本次保留请求；旧节点是否标记不影响。服务端要在响应中返回带 marker + 同 `id` 的占位节点。
- 必须带唯一 `id`：缺失 id 或重复 id 的 preserve 节点不会被保留（会告警并跳过）。
- 旧 target 的后代中必须存在同 id 节点；target 根节点本身不能作为 preserve 对象。
- 保留的旧节点整体复用：不参与 morph patch，其内部 DOM、事件、`$refs` 与运行时状态原样保留；它内部或占位节点内的 `<script>` 不会被重新执行/替换。
- 响应移除 marker 后，旧节点可以被正常更新或删除（不会因以前标记过而永久保留）。
- 支持 `inner` / `replace` / `morph` / `morph-all`；`append` / `prepend` / `beforebegin` / `afterend` 是插入类，不处理 preserve。

## morph / morph-all 注意点

- 递归合并同类型节点，节点身份按 `id` / `z-key` / `data-key` / `key` 复用（支持顺序变化），不是 `innerHTML` 重写。
- 被复用的表单控件保留 value / checked / 选中项 / 焦点 / 选区，不会被响应里的 `value="..."` 属性覆盖。
- 只有新增、删除、tag 变化或指令属性变化的节点会退场后重新激活；未变化节点保留激活状态。
- morph 响应为空/纯文本时会把目标清空。
- preserve 标记只接受 `z-preserve` 这一种属性名。

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
