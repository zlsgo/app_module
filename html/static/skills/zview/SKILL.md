---
name: zview
description: 当用户要在浏览器页面里通过 CDN/UMD 使用、解释、排障或生成 `@sohaha/zview` 示例时使用。适用于静态 HTML、服务端渲染页面、零构建集成、`$z` 信号、指令写法、事件处理、`z-req` 请求流、swap 行为以及 HTML 属性语法场景。此 skill 统一按 CDN 接入和 `window.Zview` 组织回答。
---

# zview

这个 skill 只描述 zview 自身的公开行为，面向通过 CDN 使用 `@sohaha/zview` 的包使用者。

## 资料路由

- 要生成可直接运行的页面：先读 `references/cdn-quickstart.md`。
- 要核对属性语法、请求边界或 swap 排障：再读 `references/directive-gotchas.md`。
- 要核对完整 API、配置、响应头和行为细节：以 `docs/zview.md`、`packages/view/src/` 和测试为准。

## 默认工作流

1. 默认把目标环境视为浏览器 CDN/UMD。
2. 普通页面优先用 `data-init` 自动初始化。
3. 只有在需要显式初始化顺序、运行时注册或条件化启动时，才用手动 `Zview.init(...)`。
4. 除非用户明确要求程序化 API，否则默认按 HTML-first 的声明式方式回答。
5. 输出的 HTML 片段必须使用 zview 当前真实实现的属性语法，并给出完整可运行上下文。

## 硬规则

- CDN 使用 `https://unpkg.com/@sohaha/zview/dist/zview.umd.min.js`，全局入口是 `window.Zview`。
- `data-config` 只有在脚本标签同时带 `data-init` 时才生效。
- 使用 `z-on-click`、`z-bind-foo`、`z-class-active` 这类连字符属性，不输出冒号语法。
- `window.$z` 只有在 `Zview.init()` 之后存在；自定义指令也必须在初始化后注册。
- 请求默认使用 `z-req`，按需配合 `z-trigger`、`z-target`、`z-swap`、`z-config`、`z-data` 和 `z-confirm`。
- `z-preserve` 只依据响应片段中的 marker 判断，必须有唯一 `id`；旧节点整体复用，不参与 morph patch，也不重新执行其内部脚本。
- `morph` / `morph-all` 会尽量按 `id`、`z-key`、`data-key`、`key` 复用节点，并恢复被复用表单控件的 value、checked、焦点和选区；不要把它描述成通用虚拟 DOM diff。
- `z-html` 只有轻量净化；表达式基于 `new Function`，两者都不能被描述成安全隔离。

## 回答契约

- “怎么用”类问题：先给最小完整 HTML，再解释关键属性。
- “为什么不生效”类问题：依次检查 CDN、初始化、连字符属性、signal 名称、请求目标和 swap 后重新激活。
- “怎么发请求”类问题：优先给 `z-req` 声明式写法；只有用户明确要求时才改用命令式 API。
- 解释行为时区分“源码/测试已确认”和“建议用法”，不编造不存在的指令、配置或事件。
- 示例必须保持当前实现可运行，避免使用历史 README 中的冒号语法或旧属性名。

## 结束前检查

- 是否使用了正确 CDN/UMD 入口？
- 是否包含 `data-init` 或明确调用 `Zview.init()`？
- 属性是否全部使用当前连字符语法？
- 请求是否明确了 `z-target` / `z-swap`，响应内容是否包含需要的指令？
- 是否把安全边界、`z-preserve`、morph 语义说准确？
