package el_test

import (
	"io"
	"testing"

	"github.com/zlsgo/app_module/html/el"
)

// BenchmarkElementRender 简单元素渲染
func BenchmarkElementRender(b *testing.B) {
	element := el.DIV(el.Text("Hello World"))

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		element.Render(io.Discard)
	}
}

// BenchmarkElementWithAttributes 带属性的元素
func BenchmarkElementWithAttributes(b *testing.B) {
	element := el.DIV(el.Text("Hello World")).
		Attr("id", "main").
		Attr("data-test", "value").
		Class("container", "active").
		Style("padding", "20px").
		Style("margin", "10px")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		element.Render(io.Discard)
	}
}

// BenchmarkElementNested 嵌套元素
func BenchmarkElementNested(b *testing.B) {
	element := el.DIV(
		el.H1(el.Text("Title")),
		el.P(el.Text("Paragraph 1")),
		el.P(el.Text("Paragraph 2")),
		el.DIV(
			el.SPAN(el.Text("Nested content")),
		),
	)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		element.Render(io.Discard)
	}
}

// BenchmarkElementComplex 复杂HTML结构
func BenchmarkElementComplex(b *testing.B) {
	element := el.HTML(
		el.HEAD(
			el.TITLE(el.Text("Test Page")),
			el.META().Attr("charset", "utf-8"),
		),
		el.BODY(
			el.DIV(
				el.H1(el.Text("Title")).Class("title"),
				el.P(el.Text("Content")).Class("content"),
				el.FORM(
					el.INPUT().Attr("type", "text").Attr("name", "username"),
					el.INPUT().Attr("type", "password").Attr("name", "password"),
					el.BUTTON(el.Text("Submit")).Attr("type", "submit"),
				).Attr("action", "/login").Attr("method", "post"),
			).Class("container").Style("padding", "20px"),
		),
	).Attr("lang", "en")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		element.Render(io.Discard)
	}
}

// BenchmarkClassManipulation CSS类操作
func BenchmarkClassManipulation(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		element := el.DIV().
			Class("class1").
			Class("class2", "class3").
			RemoveClass("class2").
			Class("class4")
		element.Render(io.Discard)
	}
}

// BenchmarkStyleManipulation CSS样式操作
func BenchmarkStyleManipulation(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		element := el.DIV().
			Style("color", "red").
			Style("font-size", "16px").
			Style("padding", "10px").
			Style("margin", "5px")
		element.Render(io.Discard)
	}
}

// BenchmarkDeepNesting 深度嵌套性能测试（10层）
func BenchmarkDeepNesting(b *testing.B) {
	// 构建10层嵌套的DIV结构
	var buildNested func(depth int) *el.Element
	buildNested = func(depth int) *el.Element {
		if depth == 0 {
			return el.DIV(el.Text("Deep content"))
		}
		return el.DIV(buildNested(depth - 1))
	}

	element := buildNested(10)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		element.Render(io.Discard)
	}
}

// BenchmarkMassiveAttributes 大量属性测试
func BenchmarkMassiveAttributes(b *testing.B) {
	element := el.DIV(el.Text("Content"))

	// 添加20个不同属性
	for i := 0; i < 10; i++ {
		element.Attr("data-attr-"+string(rune('a'+i)), "value"+string(rune('0'+i)))
		element.Attr("custom-"+string(rune('a'+i)), "custom-value"+string(rune('0'+i)))
	}

	// 添加常用属性
	element.Attr("id", "massive-test").
		Class("class1", "class2", "class3", "class4", "class5").
		Style("color", "red").
		Style("background", "blue").
		Style("padding", "10px").
		Style("margin", "5px").
		Style("width", "100px").
		Style("height", "50px")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		element.Render(io.Discard)
	}
}

// BenchmarkConcurrentRender 并发渲染测试
func BenchmarkConcurrentRender(b *testing.B) {
	element := el.HTML(
		el.HEAD(el.TITLE(el.Text("Concurrent Test"))),
		el.BODY(
			el.DIV(
				el.H1(el.Text("Title")).Class("header"),
				el.P(el.Text("Content paragraph")).Attr("id", "content"),
			).Class("container").Style("padding", "20px"),
		),
	)

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			element.Render(io.Discard)
		}
	})
}

// BenchmarkMemoryPressure 内存压力测试
func BenchmarkMemoryPressure(b *testing.B) {
	// 创建大量不同的元素，测试内存池效果
	elements := make([]*el.Element, 100)
	for i := 0; i < 100; i++ {
		elements[i] = el.DIV(
			el.H2(el.Text("Header "+string(rune('0'+i%10)))),
			el.P(el.Text("Paragraph content "+string(rune('0'+i%10)))),
			el.SPAN(el.Text("Span "+string(rune('0'+i%10)))),
		).Attr("id", "elem-"+string(rune('0'+i%10))).
			Class("item", "test").
			Style("margin", "10px")
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		elem := elements[i%100]
		elem.Render(io.Discard)
	}
}

// BenchmarkLargeDocument 大文档渲染测试
func BenchmarkLargeDocument(b *testing.B) {
	// 构建一个大型HTML文档（模拟真实网页）
	var rows []el.ElementRenderer
	for i := 0; i < 50; i++ { // 50行表格
		row := el.TR(
			el.TD(el.Text("Cell 1-"+string(rune('0'+i%10)))),
			el.TD(el.Text("Cell 2-"+string(rune('0'+i%10)))),
			el.TD(el.Text("Cell 3-"+string(rune('0'+i%10)))),
			el.TD(el.Text("Cell 4-"+string(rune('0'+i%10)))),
		)
		rows = append(rows, row)
	}

	document := el.HTML(
		el.HEAD(
			el.TITLE(el.Text("Large Document Test")),
			el.META().Attr("charset", "utf-8"),
			el.META().Attr("viewport", "width=device-width, initial-scale=1"),
		),
		el.BODY(
			el.DIV(
				el.H1(el.Text("Large Document")).Class("title"),
				el.P(el.Text("This is a test document with many elements")).Class("description"),
				el.TABLE(rows...).Class("data-table").Style("width", "100%"),
			).Class("container").Style("padding", "20px"),
		),
	).Attr("lang", "en")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		document.Render(io.Discard)
	}
}

// BenchmarkSimpleElementOptimization 简单元素零分配优化测试
func BenchmarkSimpleElementOptimization(b *testing.B) {
	element := el.SPAN() // 无属性、无子元素的简单元素

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		element.Render(io.Discard)
	}
}
