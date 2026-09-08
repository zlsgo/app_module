package el

import (
	"context"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
)

// TestAttrEscapingAtRender 回归验证：属性值一律在渲染序列化时才统一转义。
// 旧实现只在 Attr 的 Map 分支预先转义，字符串/SetAttribute/DeferredAttr 中
// 包含双引号（例如 JSON 文本）时会破坏属性边界，产出形如
// z-config="{"format":"urlencoded"}" 的畸形 HTML。
func TestAttrEscapingAtRender(t *testing.T) {
	tt := zlsgo.NewTest(t)

	cases := []ChunkTest{
		{
			// 字符串形式的 JSON：双引号必须转义
			Node:     DIV(Attr("z-config", `{"format":"urlencoded"}`)),
			Rendered: `<div z-config="{&#34;format&#34;:&#34;urlencoded&#34;}"></div>`,
		},
		{
			// Map 与等效字符串渲染结果一致（渲染时统一转义）
			Node:     DIV(Data("config", ztype.Map{"format": "urlencoded"})),
			Rendered: `<div data-config="{&#34;format&#34;:&#34;urlencoded&#34;}"></div>`,
		},
		{
			// 普通特殊字符
			Node:     DIV(Attr("title", `say "hi" & bye`)),
			Rendered: `<div title="say &#34;hi&#34; &amp; bye"></div>`,
		},
		{
			// 通过 SetAttribute 写入的原始值同样在渲染时转义
			Node: func() Node {
				e := DIV()
				e.SetAttribute("data-x", `{"a":"b"}`)
				return e
			}(),
			Rendered: `<div data-x="{&#34;a&#34;:&#34;b&#34;}"></div>`,
		},
		{
			// 延迟属性同样在渲染时转义
			Node:     BODY(DeferredAttr("data-y", func(ctx context.Context) string { return `{"a":"b"}` })),
			Rendered: `<body data-y="{&#34;a&#34;:&#34;b&#34;}"></body>`,
			Impure:   true,
		},
		{
			// 空值属性仍然输出为裸属性名
			Node:     INPUT(Attr("checked", "")),
			Rendered: "<input checked>",
		},
	}

	for _, c := range cases {
		c.Assert(tt)
	}
}

// TestAttributeStoresRawValue 明确存储语义：Attribute.Value 与 GetAttribute
// 返回原始值，转义只发生在渲染输出阶段。
func TestAttributeStoresRawValue(t *testing.T) {
	tt := zlsgo.NewTest(t)

	a := Attr("z-config", `{"format":"urlencoded"}`)
	tt.Equal(`{"format":"urlencoded"}`, a.Value)

	div := DIV(a)
	tt.Equal(`{"format":"urlencoded"}`, div.GetAttribute("z-config"))

	attrs := div.GetAttributes()
	tt.Equal(`{"format":"urlencoded"}`, attrs["z-config"])
}
