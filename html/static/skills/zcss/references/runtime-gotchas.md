# zcss 运行时注意点

当用户在排查行为问题，而不是单纯要一个示例时，读这个文件。

## 初始化顺序

- 尽量在第一次编译前调用 `zcss.setup(...)`。
- 如果类名是在页面加载后才加入，而当前又没有观察器，就需要显式执行 `zcss.scan(node)` 或 `zcss.observe()`。

## 配置更新语义

- 已注入的工具类 token 会按 token 去重。
- 重复使用完全相同的 token，不会强制重新注入。
- 如果在旧 token 已经编译之后再改配置，通常需要重新扫描或重新触发对应 DOM 来源。

## Theme 结构

- `theme` 的各个 map 都是扁平记录，不是嵌套 Tailwind 对象。
- 正确：`colors: { 'brand-500': '#3b82f6' }`
- 错误：`colors: { brand: { 500: '#3b82f6' } }`

## 深浅色模式

- `darkMode: 'class'` 依赖 DOM 中存在 `.dark`。
- `darkMode: 'media'` 或未配置时，走 `prefers-color-scheme`。
- 如果 `dark:` 示例表现不对，先确认页面实际配置的是哪一种模式。

## preset 与变量

- `preset: true` 会注入 base rules 和一批 CSS 变量。
- 如果用户要覆盖默认颜色或 spacing 变量，CSS 自定义属性往往比 theme fallback map 更关键。

## 观察器规则

- ignore 是基于属性的：用 `[z-ignore]`。
- 不要承诺存在选择器过滤式 ignore。
- `observe()` 同时只保留一个活跃观察器；再次调用会替换之前的观察器。

## CSS-in-JS 限制

- 数字值会被序列化为带 `px` 的长度；像 `opacity`、`zIndex`、无单位 flex 值这类属性应显式写成字符串。
- `zcss.css()` 里的 `$dark` / `$light` 会跟随当前 zcss 的 `darkMode` 行为。

## 构建 profile 假设

- CDN 使用者默认应走完整 UMD 构建，除非用户明确要求更小的 profile。
- 不要默认 tiny、lite、runtime 构建包含完整工具类和所有 helper。
