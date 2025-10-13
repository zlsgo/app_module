package el_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/zlsgo/app_module/html/el"
)

func TestBasicElement(t *testing.T) {
	tt := zlsgo.NewTest(t)

	head := el.HEAD(
		el.TITLE(el.Text("测试页面标题")),
		el.META().Attrs("name", "description", "content", "测试页面描述"),
	)

	css := el.LINK().Attrs("rel", "stylesheet").AttrsMap(map[string]string{"href": "/css/style.css"})
	head.Children(css)

	html := el.HTML(head).Attr("lang", "en")

	form := el.FORM().
		AttrsMap(map[string]string{"method": "post", "action": "/submit", "enctype": "multipart/form-data"})

	body := el.BODY(form).Attrs("data-test")
	html.Children(body)

	a := el.A().Attr("href", "/").Text("首页")
	div := el.DIV().Children(a)
	group := el.Group(el.DIV().Text("1"), el.DIV().Text("2"))
	body.Children(div, group)

	script := el.SCRIPT().Text("console.log('hello world');")
	body.Children(script)

	b, err := el.RenderToBytes(html)
	tt.NoError(err)

	result := string(b)
	expects := []string{
		`<html lang="en">`,
		`<title>测试页面标题</title>`,
		`<meta name="description" content="测试页面描述" />`,
		`<link rel="stylesheet" href="/css/style.css" />`,
		`<body data-test>`,
		`<form`,
		`method="post"`,
		`action="/submit"`,
		`enctype="multipart/form-data"`,
		`<a href="/">首页</a>`,
		`<script>console.log('hello world');</script>`,
	}
	for _, expect := range expects {
		if !strings.Contains(result, expect) {
			tt.Log("result: ", result)
			tt.Fatal("expect: ", expect)
		}
	}
}

func TestElementAttributes(t *testing.T) {
	tt := zlsgo.NewTest(t)

	// 测试 Attr 方法
	div := el.DIV().Attr("id", "test").Attr("data-value", "123")
	b, _ := el.RenderToBytes(div)
	result := string(b)
	tt.EqualTrue(strings.Contains(result, `id="test"`))
	tt.EqualTrue(strings.Contains(result, `data-value="123"`))

	// 测试空属性
	div2 := el.DIV().Attr("disabled")
	b2, _ := el.RenderToBytes(div2)
	result2 := string(b2)
	tt.EqualTrue(strings.Contains(result2, "disabled"))
}

func TestElementAttrs(t *testing.T) {
	tt := zlsgo.NewTest(t)

	div := el.DIV().Attrs("id", "main", "class", "container")
	b, _ := el.RenderToBytes(div)
	result := string(b)
	tt.EqualTrue(strings.Contains(result, `id="main"`))
	tt.EqualTrue(strings.Contains(result, `class="container"`))
}

func TestElementAttrsMap(t *testing.T) {
	tt := zlsgo.NewTest(t)

	attrs := map[string]string{
		"id":    "test",
		"class": "box",
		"style": "color: red",
	}
	div := el.DIV().AttrsMap(attrs)
	b, _ := el.RenderToBytes(div)
	result := string(b)
	tt.EqualTrue(strings.Contains(result, `id="test"`))
	tt.EqualTrue(strings.Contains(result, `class="box"`))
}

func TestElementBoolAttr(t *testing.T) {
	tt := zlsgo.NewTest(t)

	input := el.INPUT().
		Attr("type", "checkbox").
		BoolAttr("checked", true).
		BoolAttr("disabled", false)

	b, _ := el.RenderToBytes(input)
	result := string(b)
	tt.EqualTrue(strings.Contains(result, "checked"))
	tt.EqualTrue(!strings.Contains(result, "disabled"))

	// 测试 BoolValue String 方法
	bv := el.BoolValue(true)
	tt.Equal("true", bv.String())
	bv2 := el.BoolValue(false)
	tt.Equal("false", bv2.String())
}

func TestElementClass(t *testing.T) {
	tt := zlsgo.NewTest(t)

	div := el.DIV().
		Class("container").
		Class("active", "highlight")

	b, _ := el.RenderToBytes(div)
	result := string(b)
	tt.EqualTrue(strings.Contains(result, "class"))
	tt.EqualTrue(strings.Contains(result, "container"))
	tt.EqualTrue(strings.Contains(result, "active"))
	tt.EqualTrue(strings.Contains(result, "highlight"))
}

func TestElementRemoveClass(t *testing.T) {
	tt := zlsgo.NewTest(t)

	div := el.DIV().
		Class("container", "active", "highlight").
		RemoveClass("active")

	b, _ := el.RenderToBytes(div)
	result := string(b)
	tt.EqualTrue(strings.Contains(result, "container"))
	tt.EqualTrue(!strings.Contains(result, "active"))
	tt.EqualTrue(strings.Contains(result, "highlight"))
}

func TestElementStyle(t *testing.T) {
	tt := zlsgo.NewTest(t)

	div := el.DIV().
		Style("color", "red").
		Style("font-size", "16px").
		Style("margin", "10px")

	b, _ := el.RenderToBytes(div)
	result := string(b)

	tt.Log(result)
	tt.EqualTrue(strings.Contains(result, "style"))
	tt.EqualTrue(strings.Contains(result, "color"))
	tt.EqualTrue(strings.Contains(result, "red"))
	tt.EqualTrue(strings.Contains(result, "font-size"))
	tt.EqualTrue(strings.Contains(result, "16px"))
}

func TestElementText(t *testing.T) {
	tt := zlsgo.NewTest(t)

	div := el.DIV().Text("Hello").Text(" World")
	b, _ := el.RenderToBytes(div)
	result := string(b)
	tt.EqualTrue(strings.Contains(result, "Hello"))
	tt.EqualTrue(strings.Contains(result, " World"))
}

func TestElementTextF(t *testing.T) {
	tt := zlsgo.NewTest(t)

	div := el.DIV().TextF("Count: %d", 42)
	b, _ := el.RenderToBytes(div)
	result := string(b)
	tt.EqualTrue(strings.Contains(result, "Count: 42"))
}

func TestElementEscaped(t *testing.T) {
	tt := zlsgo.NewTest(t)

	div := el.DIV().Escaped("<script>alert('xss')</script>")
	b, _ := el.RenderToBytes(div)
	result := string(b)
	tt.EqualTrue(!strings.Contains(result, "<script>"))
	tt.EqualTrue(strings.Contains(result, "&lt;script&gt;"))
}

func TestElementEscapedF(t *testing.T) {
	tt := zlsgo.NewTest(t)

	div := el.DIV().EscapedF("<b>%s</b>", "Bold")
	b, _ := el.RenderToBytes(div)
	result := string(b)
	tt.EqualTrue(strings.Contains(result, "&lt;b&gt;"))
	tt.EqualTrue(strings.Contains(result, "Bold"))
}

func TestElementIfText(t *testing.T) {
	tt := zlsgo.NewTest(t)

	div := el.DIV().
		IfText(true, "Visible").
		IfText(false, "Hidden")

	b, _ := el.RenderToBytes(div)
	result := string(b)
	tt.EqualTrue(strings.Contains(result, "Visible"))
	tt.EqualTrue(!strings.Contains(result, "Hidden"))
}

func TestElementIfEscaped(t *testing.T) {
	tt := zlsgo.NewTest(t)

	div1 := el.DIV().IfEscaped(true, "<b>Bold</b>")
	b1, _ := el.RenderToBytes(div1)
	result1 := string(b1)
	tt.EqualTrue(strings.Contains(result1, "&lt;b&gt;"))

	div2 := el.DIV().IfEscaped(false, "<b>Bold</b>")
	b2, _ := el.RenderToBytes(div2)
	result2 := string(b2)
	tt.EqualTrue(!strings.Contains(result2, "Bold"))
}

func TestElementChaining(t *testing.T) {
	tt := zlsgo.NewTest(t)

	div := el.DIV().
		Attr("id", "main").
		Class("container").
		Style("padding", "20px").
		Text("Hello World")

	b, _ := el.RenderToBytes(div)
	result := string(b)
	tt.EqualTrue(strings.Contains(result, `id="main"`))
	tt.EqualTrue(strings.Contains(result, "class"))
	tt.EqualTrue(strings.Contains(result, "container"))
	tt.EqualTrue(strings.Contains(result, "style"))
	tt.EqualTrue(strings.Contains(result, "padding"))
	tt.EqualTrue(strings.Contains(result, "Hello World"))
}

func TestIf(t *testing.T) {
	tt := zlsgo.NewTest(t)

	// 条件为 true
	result1 := el.If(true, el.DIV().Text("Visible"))
	b1, _ := el.RenderToBytes(result1)
	tt.EqualTrue(strings.Contains(string(b1), "Visible"))

	// 条件为 false
	result2 := el.If(false, el.DIV().Text("Hidden"))
	tt.Equal(true, result2 == nil)
}

func TestTern(t *testing.T) {
	tt := zlsgo.NewTest(t)

	// 条件为 true
	result1 := el.Tern(true, el.DIV().Text("True"), el.DIV().Text("False"))
	b1, _ := el.RenderToBytes(result1)
	tt.EqualTrue(strings.Contains(string(b1), "True"))

	// 条件为 false
	result2 := el.Tern(false, el.DIV().Text("True"), el.DIV().Text("False"))
	b2, _ := el.RenderToBytes(result2)
	tt.EqualTrue(strings.Contains(string(b2), "False"))
}

func TestRange(t *testing.T) {
	tt := zlsgo.NewTest(t)

	items := []string{"A", "B", "C"}
	result := el.Range(items, func(item string) el.ElementRenderer {
		return el.DIV().Text(item)
	})

	b, _ := el.RenderToBytes(result)
	str := string(b)
	tt.EqualTrue(strings.Contains(str, "A"))
	tt.EqualTrue(strings.Contains(str, "B"))
	tt.EqualTrue(strings.Contains(str, "C"))
}

func TestRangeI(t *testing.T) {
	tt := zlsgo.NewTest(t)

	items := []string{"A", "B", "C"}
	result := el.RangeI(items, func(i int, item string) el.ElementRenderer {
		return el.DIV().TextF("%d: %s", i, item)
	})

	b, _ := el.RenderToBytes(result)
	str := string(b)
	tt.EqualTrue(strings.Contains(str, "0: A"))
	tt.EqualTrue(strings.Contains(str, "1: B"))
	tt.EqualTrue(strings.Contains(str, "2: C"))
}

func TestDynGroup(t *testing.T) {
	tt := zlsgo.NewTest(t)

	result := el.DynGroup(
		func() el.ElementRenderer { return el.DIV().Text("1") },
		func() el.ElementRenderer { return el.DIV().Text("2") },
		func() el.ElementRenderer { return nil }, // 测试 nil 过滤
	)

	b, _ := el.RenderToBytes(result)
	str := string(b)
	tt.EqualTrue(strings.Contains(str, "1"))
	tt.EqualTrue(strings.Contains(str, "2"))
}

func TestDynIf(t *testing.T) {
	tt := zlsgo.NewTest(t)

	// 条件为 true
	result1 := el.DynIf(true,
		func() el.ElementRenderer { return el.DIV().Text("A") },
		func() el.ElementRenderer { return el.DIV().Text("B") },
	)
	b1, _ := el.RenderToBytes(result1)
	str1 := string(b1)
	tt.EqualTrue(strings.Contains(str1, "A"))
	tt.EqualTrue(strings.Contains(str1, "B"))

	// 条件为 false
	result2 := el.DynIf(false,
		func() el.ElementRenderer { return el.DIV().Text("C") },
	)
	tt.Equal(true, result2 == nil)
}

func TestDynTern(t *testing.T) {
	tt := zlsgo.NewTest(t)

	// 条件为 true
	result1 := el.DynTern(true,
		func() el.ElementRenderer { return el.DIV().Text("True") },
		func() el.ElementRenderer { return el.DIV().Text("False") },
	)
	b1, _ := el.RenderToBytes(result1)
	tt.EqualTrue(strings.Contains(string(b1), "True"))

	// 条件为 false
	result2 := el.DynTern(false,
		func() el.ElementRenderer { return el.DIV().Text("True") },
		func() el.ElementRenderer { return el.DIV().Text("False") },
	)
	b2, _ := el.RenderToBytes(result2)
	tt.EqualTrue(strings.Contains(string(b2), "False"))
}

func TestError(t *testing.T) {
	tt := zlsgo.NewTest(t)

	err := errors.New("test error")
	result := el.Error(err)
	b, _ := el.RenderToBytes(result)
	tt.Equal("test error", string(b))
}

func TestTextContent(t *testing.T) {
	tt := zlsgo.NewTest(t)

	text := el.Text("Hello World")
	b, _ := el.RenderToBytes(text)
	tt.Equal("Hello World", string(b))
}

func TestTextF(t *testing.T) {
	tt := zlsgo.NewTest(t)

	text := el.TextF("Count: %d, Name: %s", 42, "Test")
	b, _ := el.RenderToBytes(text)
	tt.Equal("Count: 42, Name: Test", string(b))
}

func TestEscapedContent(t *testing.T) {
	tt := zlsgo.NewTest(t)

	escaped := el.Escaped("<script>alert('xss')</script>")
	b, _ := el.RenderToBytes(escaped)
	tt.EqualTrue(strings.Contains(string(b), "&lt;script&gt;"))
}

func TestEscapedF(t *testing.T) {
	tt := zlsgo.NewTest(t)

	escaped := el.EscapedF("<b>%s</b>", "Bold")
	b, _ := el.RenderToBytes(escaped)
	result := string(b)
	tt.EqualTrue(strings.Contains(result, "&lt;b&gt;"))
	tt.EqualTrue(strings.Contains(result, "Bold"))
}

func TestGroup(t *testing.T) {
	tt := zlsgo.NewTest(t)

	group := el.Group(
		el.DIV().Text("1"),
		el.DIV().Text("2"),
		el.DIV().Text("3"),
	)

	b, err := el.RenderToBytes(group)
	tt.NoError(err)
	str := string(b)
	tt.EqualTrue(strings.Contains(str, "1"))
	tt.EqualTrue(strings.Contains(str, "2"))
	tt.EqualTrue(strings.Contains(str, "3"))
}
