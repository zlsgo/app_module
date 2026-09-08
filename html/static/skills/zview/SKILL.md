---
name: zview
description: 当用户要在浏览器页面里通过 CDN/UMD 使用、解释、排障或生成 `@sohaha/zview` 示例时使用。适用于静态 HTML、服务端渲染页面、零构建集成、`$z` 信号、指令写法、事件处理、`z-req` 请求流、swap 行为以及 HTML 属性语法场景。也覆盖从 `htmx`/`datastar` 迁移的概念对照与能力边界。此 skill 统一按 CDN 接入和 `window.Zview` 组织回答。
---

# zview

这个 skill 面向通过 CDN 使用 `@sohaha/zview` 的包使用者。

## 适用范围

- 默认把目标环境视为浏览器 CDN/UMD。
- 优先给生产可用的 `https://unpkg.com/@sohaha/zview/dist/zview.umd.min.js` 示例。
- 使用真实全局入口：`window.Zview`。
- 除非用户明确要求程序化 API，否则默认按 HTML-first 的声明式使用方式回答。

## 默认工作流

1. 先从 `references/cdn-quickstart.md` 开始。
2. 普通页面优先用 `data-init` 自动初始化。
3. 只有在需要显式初始化顺序、运行时注册或条件化启动时，才用手动 `Zview.init(...)`。
4. 如果需求涉及指令、请求、swap 或语法混淆，再读 `references/directive-gotchas.md`。
5. 输出的 HTML 片段必须使用 zview 当前真实实现的属性语法。
6. 若用户从 `htmx`/`datastar` 迁移过来、或想复刻它们的交互，先读 `references/htmx-datastar-mapping.md` 再给答案。

## 输出目标

- 正确的 CDN 初始化。
- 正确的连字符指令写法，例如 `z-on-click`。
- 正确使用 `$z` 和 `$signal`。
- 正确解释请求与 swap 流程，不编造不存在的 API。

## 硬规则

- `data-config` 只有在脚本标签同时带 `data-init` 时才生效。
- 使用 `z-on-click`、`z-bind-foo`、`z-class-active` 这类属性名。
- `window.$z` 只有在 `Zview.init()` 之后才存在。
- 需要自定义指令时，先 `Zview.init()`，再 `Zview.register(...)`。
- `z-html` 的净化只是轻量规则，不要把它描述成完整 sanitizer。
- 表达式执行基于 `new Function`，不要把它说成沙箱。

## 回答模式

- 遇到“怎么用 zview”，先给完整 HTML 片段。
- 遇到“为什么这个指令不运行”，优先检查初始化时机、属性命名、信号命名，以及 swap 后内容是否被 zview 重新激活。
- 遇到“怎么发请求”，默认用 `z-req` 配合 `z-trigger`、`z-target`、`z-swap` 和相关配置，而不是手写 fetch，除非用户明确要求别的方案。

## 参考资料

- 快速开始与可直接粘贴的示例：`references/cdn-quickstart.md`
- 指令语法与排障清单：`references/directive-gotchas.md`
- `htmx` / `datastar` → zview 概念对照与能力边界：`references/htmx-datastar-mapping.md`
