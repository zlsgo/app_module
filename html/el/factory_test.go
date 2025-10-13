package el_test

import (
	"strings"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/zlsgo/app_module/html/el"
)

func TestHTMLElements(t *testing.T) {
	tt := zlsgo.NewTest(t)

	// 测试各种HTML元素
	tests := []struct {
		name    string
		element el.ElementRenderer
		tag     string
	}{
		{"SPAN", el.SPAN(el.Text("span")), "span"},
		{"P", el.P(el.Text("paragraph")), "p"},
		{"H1", el.H1(el.Text("h1")), "h1"},
		{"H2", el.H2(el.Text("h2")), "h2"},
		{"H3", el.H3(el.Text("h3")), "h3"},
		{"BUTTON", el.BUTTON(el.Text("button")), "button"},
		{"TEXTAREA", el.TEXTAREA(el.Text("text")), "textarea"},
		{"SELECT", el.SELECT(el.OPTION(el.Text("opt"))), "select"},
		{"OPTION", el.OPTION(el.Text("option")), "option"},
		{"TABLE", el.TABLE(el.TR(el.TD(el.Text("cell")))), "table"},
		{"TR", el.TR(el.TD(el.Text("cell"))), "tr"},
		{"TD", el.TD(el.Text("cell")), "td"},
		{"TH", el.TH(el.Text("header")), "th"},
		{"UL", el.UL(el.LI(el.Text("item"))), "ul"},
		{"OL", el.OL(el.LI(el.Text("item"))), "ol"},
		{"LI", el.LI(el.Text("item")), "li"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b, err := el.RenderToBytes(tc.element)
			tt.NoError(err)
			result := string(b)
			tt.EqualTrue(strings.Contains(result, "<"+tc.tag))
			tt.EqualTrue(strings.Contains(result, "</"+tc.tag+">"))
		})
	}
}

func TestSelfClosingElements(t *testing.T) {
	tt := zlsgo.NewTest(t)

	// 测试自闭合元素
	tests := []struct {
		name    string
		element el.ElementRenderer
		tag     string
	}{
		{"IMG", el.IMG(), "img"},
		{"BR", el.BR(), "br"},
		{"HR", el.HR(), "hr"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b, err := el.RenderToBytes(tc.element)
			tt.NoError(err)
			result := string(b)
			tt.EqualTrue(strings.Contains(result, "<"+tc.tag))
			tt.EqualTrue(strings.Contains(result, "/>"))
			// 自闭合元素不应该有结束标签
			tt.EqualTrue(!strings.Contains(result, "</"+tc.tag+">"))
		})
	}
}

func TestElementFactory(t *testing.T) {
	tt := zlsgo.NewTest(t)

	// 测试已知元素
	div := el.Create("div", el.Text("test"))
	b, _ := el.RenderToBytes(div)
	tt.EqualTrue(strings.Contains(string(b), "<div>"))

	// 测试未知元素（应该创建自定义元素）
	custom := el.Create("custom-element", el.Text("custom"))
	b2, _ := el.RenderToBytes(custom)
	result := string(b2)
	tt.EqualTrue(strings.Contains(result, "<custom-element>"))
	tt.EqualTrue(strings.Contains(result, "</custom-element>"))
}

func TestRegisterElement(t *testing.T) {
	tt := zlsgo.NewTest(t)

	// 注册自定义元素
	config := &el.ElementConfig{
		Tag:           "my-element",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Custom element for testing",
	}
	el.RegisterElement("my-element", config)

	// 验证可以获取配置
	retrieved, ok := el.GetElementConfig("my-element")
	tt.Equal(true, ok)
	tt.Equal("my-element", retrieved.Tag)
	tt.Equal("Custom element for testing", retrieved.Description)
}

func TestGetElementConfig(t *testing.T) {
	tt := zlsgo.NewTest(t)

	// 测试获取已知元素配置
	config, ok := el.GetElementConfig("div")
	tt.Equal(true, ok)
	tt.Equal("div", config.Tag)

	// 测试获取不存在的元素配置
	_, ok2 := el.GetElementConfig("nonexistent-element")
	tt.Equal(false, ok2)
}

func TestListSupportedElements(t *testing.T) {
	tt := zlsgo.NewTest(t)

	elements := el.ListSupportedElements()
	tt.EqualTrue(len(elements) > 0)

	// 验证包含常见元素
	hasDiv := false
	hasSpan := false
	for _, el := range elements {
		if el == "div" {
			hasDiv = true
		}
		if el == "span" {
			hasSpan = true
		}
	}
	tt.Equal(true, hasDiv)
	tt.Equal(true, hasSpan)
}

func TestComplexForm(t *testing.T) {
	tt := zlsgo.NewTest(t)

	form := el.FORM(
		el.DIV(
			el.LABEL(el.Text("Username:")).Attr("for", "username"),
			el.INPUT().Attr("type", "text").Attr("id", "username").Attr("name", "username"),
		),
		el.DIV(
			el.LABEL(el.Text("Password:")).Attr("for", "password"),
			el.INPUT().Attr("type", "password").Attr("id", "password").Attr("name", "password"),
		),
		el.DIV(
			el.BUTTON(el.Text("Submit")).Attr("type", "submit"),
		),
	).Attr("action", "/login").Attr("method", "post")

	b, err := el.RenderToBytes(form)
	tt.NoError(err)
	result := string(b)

	tt.EqualTrue(strings.Contains(result, "<form"))
	tt.EqualTrue(strings.Contains(result, `action="/login"`))
	tt.EqualTrue(strings.Contains(result, `method="post"`))
	tt.EqualTrue(strings.Contains(result, "<label"))
	tt.EqualTrue(strings.Contains(result, "<input"))
	tt.EqualTrue(strings.Contains(result, "<button"))
}

func TestComplexTable(t *testing.T) {
	tt := zlsgo.NewTest(t)

	table := el.TABLE(
		el.TR(
			el.TH(el.Text("Name")),
			el.TH(el.Text("Age")),
		),
		el.TR(
			el.TD(el.Text("Alice")),
			el.TD(el.Text("30")),
		),
		el.TR(
			el.TD(el.Text("Bob")),
			el.TD(el.Text("25")),
		),
	)

	b, err := el.RenderToBytes(table)
	tt.NoError(err)
	result := string(b)

	tt.EqualTrue(strings.Contains(result, "<table>"))
	tt.EqualTrue(strings.Contains(result, "<tr>"))
	tt.EqualTrue(strings.Contains(result, "<th>"))
	tt.EqualTrue(strings.Contains(result, "<td>"))
	tt.EqualTrue(strings.Contains(result, "Alice"))
	tt.EqualTrue(strings.Contains(result, "Bob"))
}

func TestLists(t *testing.T) {
	tt := zlsgo.NewTest(t)

	// 无序列表
	ul := el.UL(
		el.LI(el.Text("Item 1")),
		el.LI(el.Text("Item 2")),
		el.LI(el.Text("Item 3")),
	)

	b1, _ := el.RenderToBytes(ul)
	result1 := string(b1)
	tt.EqualTrue(strings.Contains(result1, "<ul>"))
	tt.EqualTrue(strings.Contains(result1, "<li>"))
	tt.EqualTrue(strings.Contains(result1, "Item 1"))

	// 有序列表
	ol := el.OL(
		el.LI(el.Text("First")),
		el.LI(el.Text("Second")),
	)

	b2, _ := el.RenderToBytes(ol)
	result2 := string(b2)
	tt.EqualTrue(strings.Contains(result2, "<ol>"))
	tt.EqualTrue(strings.Contains(result2, "First"))
}
