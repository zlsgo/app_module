package el

import (
	"bytes"
	"context"
	"iter"
	"testing"

	"github.com/sohaha/zlsgo"
)

func TestBase(t *testing.T) {
	tt := zlsgo.NewTest(t)

	div := El("div",
		Attr("id", "main"),
		Text("Hello, world!"),
		El("span", Text("Nested span")),
	)

	tt.Equal("div", div.Tag())
	tt.Equal("main", div.GetAttribute("id"))

	nodes := div.GetNodes()
	tt.Equal(2, len(nodes))

	textNode, ok := nodes[0].(Text)
	tt.EqualTrue(ok)
	tt.Equal("Hello, world!", string(textNode))

	spanNode, ok := nodes[1].(*Element)
	tt.EqualTrue(ok)
	tt.Equal("span", spanNode.Tag())

	spanChildren := spanNode.GetNodes()
	tt.Equal(1, len(spanChildren))
	innerText, ok := spanChildren[0].(Text)
	tt.EqualTrue(ok)
	tt.Equal("Nested span", string(innerText))

	ChunkTest{
		Node:     div,
		Rendered: `<div id="main">Hello, world!<span>Nested span</span></div>`,
	}.Assert(tt)
}

type ChunkTest struct {
	Node     Node
	N        int
	Impure   bool
	Rendered string
}

func (c ChunkTest) Assert(tt *zlsgo.TestUtil) {
	cw := NewChunkWriter()
	c.Node.Render(cw)
	chunks := cw.Chunks()

	if c.N > 0 {
		tt.Equal(c.N, len(chunks))
	}

	buf := new(bytes.Buffer)

	for _, chunk := range chunks {
		if err := RenderChunk(chunk, context.Background(), buf); err != nil {
			tt.Fatal(err.Error())
		}
	}

	tt.Equal(c.Rendered, buf.String())
}

func TestRenderHelpers(t *testing.T) {
	t.Run("RenderNode", func(t *testing.T) {
		tt := zlsgo.NewTest(t)
		node := DIV(ID("main"), Text("Hello"))
		buf := new(bytes.Buffer)

		if err := RenderNode(context.Background(), buf, node); err != nil {
			t.Fatalf("RenderNode returned error: %v", err)
		}
		tt.Equal(`<div id="main">Hello</div>`, buf.String())
	})

	t.Run("RenderMultiple", func(t *testing.T) {
		tt := zlsgo.NewTest(t)
		buf := new(bytes.Buffer)
		err := Render(context.Background(), buf,
			DIV(Text("A")),
			SPAN(Text("B")),
		)
		if err != nil {
			t.Fatalf("Render returned error: %v", err)
		}
		tt.Equal("<div>A</div><span>B</span>", buf.String())
	})

	t.Run("RenderString", func(t *testing.T) {
		tt := zlsgo.NewTest(t)
		out, err := RenderBytes(context.Background(), P(Text("ok")))
		if err != nil {
			t.Fatalf("RenderString returned error: %v", err)
		}
		tt.Equal("<p>ok</p>", string(out))
	})

	t.Run("MustRenderString panics on invalid input", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("expected panic from MustRenderString")
			}
		}()
		_ = MustRenderBytes(context.Background(), Attr("id", "oops"))
	})
}

func TestNodeChunks(t *testing.T) {
	tests := []ChunkTest{
		{
			Node:     Text("Hello, World!"),
			Rendered: "Hello, World!",
			N:        1,
		},
		{
			Node:     DIV(),
			Rendered: "<div></div>",
		},
		{
			Node:     Doctype,
			Rendered: "<!DOCTYPE html>",
		},
		{
			Node:     INPUT(P()),
			Rendered: "<input>",
		},
		{
			Node:     DIV(ID("app"), Data("app", "zlsgo")),
			Rendered: `<div id="app" data-app="zlsgo"></div>`,
		},
		{
			Node:     DIV(P(Text("Hello, World!"))),
			Rendered: "<div><p>Hello, World!</p></div>",
		},
		{
			Node:     Fragment(HEAD(), BODY()),
			Rendered: "<head></head><body></body>",
		},
		{
			Node: Component(func(ctx context.Context) Node {
				return DIV(ID("app"))
			}),
			Rendered: `<div id="app"></div>`,
			Impure:   true,
		},
		{
			Node:     BODY(DeferredAttr("class", func(ctx context.Context) string { return "dark" })),
			Rendered: `<body class="dark"></body>`,
			Impure:   true,
		},
		{
			Node:     BODY(ID("app"), DeferredAttr("class", func(ctx context.Context) string { return "dark" })),
			Rendered: `<body class="dark" id="app"></body>`,
			Impure:   true,
		},
	}

	tt := zlsgo.NewTest(t)
	for _, test := range tests {
		test.Assert(tt)
	}
}

func TestMap(t *testing.T) {
	tt := zlsgo.NewTest(t)

	items := []string{"a", "b", "c"}
	node := Map(items, func(s string) Node {
		return Text(s)
	})

	ChunkTest{
		Node:     node,
		Rendered: "abc",
	}.Assert(tt)
}

func TestMapIdx(t *testing.T) {
	tt := zlsgo.NewTest(t)

	items := []string{"a", "b", "c"}
	node := MapIdx(items, func(s string, i int) Node {
		return Textf("%d:%s", i, s)
	})

	ChunkTest{
		Node:     node,
		Rendered: "0:a1:b2:c",
	}.Assert(tt)
}

func TestIter(t *testing.T) {
	tt := zlsgo.NewTest(t)

	seq := func(yield func(string) bool) {
		for _, v := range []string{"x", "y", "z"} {
			if !yield(v) {
				return
			}
		}
	}

	node := Iter(seq, func(s string) Node {
		return Text(s)
	})

	ChunkTest{
		Node:     node,
		Rendered: "xyz",
	}.Assert(tt)
}

func TestIter2(t *testing.T) {
	tt := zlsgo.NewTest(t)

	seq := func(yield func(int, string) bool) {
		for i, v := range []string{"a", "b"} {
			if !yield(i, v) {
				return
			}
		}
	}

	node := Iter2(iter.Seq2[int, string](seq), func(k int, v string) Node {
		return Textf("%d:%s", k, v)
	})

	ChunkTest{
		Node:     node,
		Rendered: "0:a1:b",
	}.Assert(tt)
}

func TestIf(t *testing.T) {
	tt := zlsgo.NewTest(t)

	t.Run("true condition", func(t *testing.T) {
		result := If(true, Text("yes"))
		node, ok := result.(Node)
		tt.EqualTrue(ok)
		ChunkTest{
			Node:     node,
			Rendered: "yes",
		}.Assert(tt)
	})

	t.Run("false condition", func(t *testing.T) {
		result := If(false, Text("no"))
		node, ok := result.(Node)
		tt.EqualTrue(ok)
		ChunkTest{
			Node:     node,
			Rendered: "",
		}.Assert(tt)
	})
}

func TestIfFn(t *testing.T) {
	tt := zlsgo.NewTest(t)

	t.Run("true condition", func(t *testing.T) {
		result := IfFn(true, func() Item {
			return Text("computed")
		})
		node, ok := result.(Node)
		tt.EqualTrue(ok)
		ChunkTest{
			Node:     node,
			Rendered: "computed",
		}.Assert(tt)
	})

	t.Run("false condition", func(t *testing.T) {
		called := false
		result := IfFn(false, func() Item {
			called = true
			return Text("should not run")
		})
		node, ok := result.(Node)
		tt.EqualTrue(ok)
		tt.EqualFalse(called)
		ChunkTest{
			Node:     node,
			Rendered: "",
		}.Assert(tt)
	})
}

func TestSwitch(t *testing.T) {
	tt := zlsgo.NewTest(t)

	t.Run("matching case", func(t *testing.T) {
		result := Switch("b",
			Case("a", Text("first")),
			Case("b", Text("second")),
			Case("c", Text("third")),
		)
		node, ok := result.(Node)
		tt.EqualTrue(ok)
		ChunkTest{
			Node:     node,
			Rendered: "second",
		}.Assert(tt)
	})

	t.Run("no match", func(t *testing.T) {
		result := Switch("x",
			Case("a", Text("first")),
			Case("b", Text("second")),
		)
		node, ok := result.(Node)
		tt.EqualTrue(ok)
		ChunkTest{
			Node:     node,
			Rendered: "",
		}.Assert(tt)
	})

	t.Run("with default", func(t *testing.T) {
		result := Switch("x",
			Case("a", Text("first")),
			Case("b", Text("second")),
			DefaultEl[string](Text("default")),
		)
		node, ok := result.(Node)
		tt.EqualTrue(ok)
		ChunkTest{
			Node:     node,
			Rendered: "default",
		}.Assert(tt)
	})
}

func TestCaseFn(t *testing.T) {
	tt := zlsgo.NewTest(t)

	result := Switch(2,
		CaseFn(1, func() Item { return Text("one") }),
		CaseFn(2, func() Item { return Text("two") }),
	)
	node, ok := result.(Node)
	tt.EqualTrue(ok)
	ChunkTest{
		Node:     node,
		Rendered: "two",
	}.Assert(tt)
}

func TestDefaultFn(t *testing.T) {
	tt := zlsgo.NewTest(t)

	result := Switch(99,
		Case(1, Text("one")),
		DefaultElFn[int](func() Item { return Text("fallback") }),
	)
	node, ok := result.(Node)
	tt.EqualTrue(ok)
	ChunkTest{
		Node:     node,
		Rendered: "fallback",
	}.Assert(tt)
}

func TestJSON(t *testing.T) {
	tt := zlsgo.NewTest(t)

	t.Run("basic json", func(t *testing.T) {
		data := map[string]any{"key": "value"}
		node := JSON(data)

		ChunkTest{
			Node:     node,
			Rendered: `{"key":"value"}` + "\n",
		}.Assert(tt)
	})

	t.Run("with indent", func(t *testing.T) {
		data := map[string]any{"a": 1, "b": 2}
		node := JSON(data).WithIndent("  ")

		ChunkTest{
			Node:     node,
			Rendered: "{\n  \"a\": 1,\n  \"b\": 2\n}\n",
		}.Assert(tt)
	})
}

func TestRaw(t *testing.T) {
	tt := zlsgo.NewTest(t)

	node := DIV(Raw("<b>bold</b>"))
	ChunkTest{
		Node:     node,
		Rendered: "<div><b>bold</b></div>",
	}.Assert(tt)
}

func TestTextf(t *testing.T) {
	tt := zlsgo.NewTest(t)

	node := Textf("Hello %s, you are %d", "Alice", 25)
	ChunkTest{
		Node:     node,
		Rendered: "Hello Alice, you are 25",
	}.Assert(tt)
}

func TestComponentNode(t *testing.T) {
	tt := zlsgo.NewTest(t)

	comp := Component(func(ctx context.Context) Node {
		return DIV(Text("from component"))
	})

	ChunkTest{
		Node:     comp,
		Rendered: "<div>from component</div>",
		Impure:   true,
	}.Assert(tt)
}
