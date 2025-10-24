package el

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
)

func TestRenderBytes(t *testing.T) {
	tt := zlsgo.NewTest(t)

	now := ztype.ToString(time.Now().Unix())
	node := HTML(HEAD(
		TITLE(Text("App")),
		SCRIPT(Raw("console.log('hello world');")),
	),
		BODY(
			DIV(Text("Hello, World!")),

			Class("container"),
			Attr("data-app", "zlsgo"),
			Style("color:red;"),

			Component(func(ctx context.Context) Node {
				return DIV(Text("Time:" + now))
			}),
		),
	)

	html, err := RenderBytes(context.Background(), node)
	tt.NoError(err)
	tt.Equal(`<html><head><title>App</title><script>console.log('hello world');</script></head><body class="container" data-app="zlsgo" style="color:red;"><div>Hello, World!</div><div>Time:`+now+`</div></body></html>`, string(html))
}

func TestRenderEdgeCases(t *testing.T) {
	t.Run("RenderNode", func(t *testing.T) {
		tt := zlsgo.NewTest(t)
		node := DIV(ID("main"), Text("Hello"))
		buf := new(bytes.Buffer)

		if err := RenderNode(context.Background(), buf, node); err != nil {
			t.Fatalf("RenderNode returned error: %v", err)
		}
		tt.Equal(`<div id="main">Hello</div>`, buf.String())
	})

	t.Run("RenderNode with nil node", func(t *testing.T) {
		tt := zlsgo.NewTest(t)
		buf := new(bytes.Buffer)
		err := RenderNode(context.Background(), buf, nil)
		tt.NoError(err)
		tt.Equal("", buf.String())
	})

	t.Run("RenderNode with nil writer", func(t *testing.T) {
		tt := zlsgo.NewTest(t)
		err := RenderNode(context.Background(), nil, DIV())
		tt.NoError(err)
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

	t.Run("Render with empty items", func(t *testing.T) {
		tt := zlsgo.NewTest(t)
		buf := new(bytes.Buffer)
		err := Render(context.Background(), buf)
		tt.NoError(err)
		tt.Equal("", buf.String())
	})

	t.Run("Render with nil writer", func(t *testing.T) {
		tt := zlsgo.NewTest(t)
		err := Render(context.Background(), nil, DIV())
		tt.NoError(err)
	})

	t.Run("Render with nil items", func(t *testing.T) {
		tt := zlsgo.NewTest(t)
		buf := new(bytes.Buffer)
		err := Render(context.Background(), buf, nil, nil)
		tt.NoError(err)
		tt.Equal("", buf.String())
	})

	t.Run("Render with invalid item", func(t *testing.T) {
		tt := zlsgo.NewTest(t)
		buf := new(bytes.Buffer)
		err := Render(context.Background(), buf, Attr("id", "test"))
		tt.Equal(true, err != nil)
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

	t.Run("MustRenderString success", func(t *testing.T) {
		tt := zlsgo.NewTest(t)
		out := MustRenderBytes(context.Background(), DIV(Text("test")))
		tt.Equal("<div>test</div>", string(out))
	})
}
