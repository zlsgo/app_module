---
name: zcss
description: 当用户要在浏览器页面里通过 CDN/UMD 使用、解释、排障或生成 `@sohaha/zcss` 示例时使用。适用于静态 HTML、服务端渲染页面、CMS 片段、零构建集成、运行时 class 编译、`zcss.tw()`、`zcss.observe()`、`zcss.setup()` 以及 `css()` / `keyframes()` 的 CSS-in-JS 场景。此 skill 统一按 CDN 接入和 `window.zcss` 组织回答。
---

# zcss

这个 skill 面向通过 CDN 使用 `@sohaha/zcss` 的包使用者。

## 适用范围

- 默认把目标环境视为浏览器 CDN/UMD。
- 优先给生产可用的 `https://unpkg.com/@sohaha/zcss/dist/zcss.umd.min.js` 示例。
- 使用真实全局对象名：`window.zcss`。

## 默认工作流

1. 先从 `references/cdn-quickstart.md` 的 CDN 快速开始写起。
2. 页面只需要 HTML 声明式初始化时，优先用自动初始化。
3. 需要显式顺序、动态配置或直接调 API 时，再用手动初始化。
4. 需求涉及 theme、config 或运行时行为时，再读 `references/runtime-gotchas.md`。
5. 生成的示例默认应能直接粘进纯 HTML 文件并在浏览器中运行。

## 输出目标

- 零构建接入。
- 正确使用 `data-init` / `data-config`。
- 正确调用运行时 API：`setup`、`observe`、`scan`、`disconnect`、`tw`、`css`、`keyframes`。
- 代码片段尽量短，但必须保留真实执行顺序。

## 硬规则

- `data-config` 只有在脚本标签同时带 `data-init` 时才生效。
- 尽量在第一次 `tw()`、`observe()` 或初始扫描前先调用 `zcss.setup(...)`。
- `theme` 必须使用扁平键，例如 `colors: { 'brand-500': '#3b82f6' }`。
- 跳过扫描用 `[z-ignore]`，不要编造选择器式 ignore 行为。
- 只要涉及深浅色示例，就显式写出 `darkMode` 模式。
- 给 CDN 示例时，保持代码自包含、可直接运行。

## 回答模式

- 遇到“怎么用 zcss”，先给完整 HTML 片段。
- 遇到“为什么这个类不生效”，优先检查初始化顺序、构建 profile 假设、`darkMode` 模式、theme 键形状、以及元素是否被扫描。
- 遇到“怎么做动态样式”，优先用 `zcss.tw()` 处理工具类，用 `zcss.css()` / `zcss.keyframes()` 处理生成类名。

## 参考资料

- 快速开始与 API 片段：`references/cdn-quickstart.md`
- 运行时行为与常见误区：`references/runtime-gotchas.md`
