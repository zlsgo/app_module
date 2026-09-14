# HTML Styles

`html/styles` 为 `html/el` 提供独立的 CSS 辅助能力。

## 行内样式

```go
style := styles.Props{
    styles.Display:    "flex",
    styles.Gap:        styles.Pixels(8),
    styles.FontWeight: "600",
}

button := el.BUTTON(style.Attr(), el.Text("确定"))
```

`Props.ToInline()` 会按 CSS 属性名排序，生成确定性的 `key: value;` 声明。`Merge`
可用于组合基础样式和覆盖样式，后者优先。

## 可复用样式表

```go
manager := styles.NewStyleManager()
className := manager.AddCompositeStyle(styles.CompositeStyle{
    Default: styles.Props{styles.Color: "#111827"},
    PseudoClasses: map[string]styles.Props{
        styles.PseudoHover: {styles.Color: "#2563eb"},
    },
})

page := el.HTML(
    el.HEAD(manager.StyleTag()),
    el.BODY(el.DIV(el.Class(className), el.Text("内容"))),
)
```

如果只需要将样式注册后直接应用到元素，也可以使用 `manager.Class(props)` 或
`manager.CompositeClass(style)`，它们直接返回 `*el.Attribute`：

```go
button := el.BUTTON(
    manager.Class(styles.Props{styles.Padding: styles.Pixels(8)}),
    el.Text("确定"),
)
```

`StyleManager` 支持：

- `AddStyle`：普通 class
- `AddCompositeStyle`：伪类、伪元素和媒体查询
- `AddAnimation`：keyframes
- `GenerateCSS` / `StyleTag`：生成 CSS 或 `<style>` 元素
- 相同样式去重，并生成稳定 class 名

CSS 选择器和值按稳定顺序输出，便于缓存和快照测试。CSS 原文不会被转义或清洗，
请不要直接把不可信用户输入拼入 CSS。

## CSS 值辅助函数

包提供常用的 `Pixels`、`Percent`、`Em`、`Rem`、`ViewportWidth`、`RGB`、`RGBA`、
`HSL`、`URL`、`Var` 等函数，避免在业务代码中重复拼接单位。
