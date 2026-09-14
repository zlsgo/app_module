package htmx

import (
	"context"
	"testing"

	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/html/el"
)

func TestAttributesRender(t *testing.T) {
	node := el.BUTTON(
		Get("/todos"),
		Target("#list"),
		Swap("innerHTML"),
		Trigger("click once"),
		Vals(ztype.Map{"page": 2, "filter": "open"}),
		el.Text("Refresh"),
	)

	got, err := el.RenderBytes(context.Background(), node)
	if err != nil {
		t.Fatal(err)
	}
	want := `<button hx-get="/todos" hx-target="#list" hx-swap="innerHTML" hx-trigger="click once" hx-vals="{&#34;filter&#34;:&#34;open&#34;,&#34;page&#34;:2}">Refresh</button>`
	if string(got) != want {
		t.Fatalf("rendered HTML = %q, want %q", got, want)
	}
}

func TestJSONAttributesEscapeAtRender(t *testing.T) {
	attr := Headers(ztype.Map{"X-Requested-With": "XMLHttpRequest"})
	if attr.Key != HXHeaders {
		t.Fatalf("attribute key = %q, want %q", attr.Key, HXHeaders)
	}
	if attr.Value != `{"X-Requested-With":"XMLHttpRequest"}` {
		t.Fatalf("attribute value = %q", attr.Value)
	}

	got, err := el.RenderBytes(context.Background(), el.DIV(attr))
	if err != nil {
		t.Fatal(err)
	}
	want := `<div hx-headers="{&#34;X-Requested-With&#34;:&#34;XMLHttpRequest&#34;}"></div>`
	if string(got) != want {
		t.Fatalf("rendered HTML = %q, want %q", got, want)
	}
}

func TestOnNameAndOn(t *testing.T) {
	tests := map[string]string{
		"beforeRequest":             "hx-on--before-request",
		"before-request":            "hx-on--before-request",
		"htmx:afterSwap":            "hx-on--after-swap",
		"hx-on:htmx:beforeRequest":  "hx-on--before-request",
		"hx-on::before-request":     "hx-on--before-request",
		"hx-on-htmx-before-request": "hx-on--before-request",
		"hx-on--afterSwap":          "hx-on--after-swap",
		" before-request ":          "hx-on--before-request",
		"click":                     "hx-on-click",
		"hx-on:click":               "hx-on-click",
		"hx-on-click":               "hx-on-click",
	}
	for input, want := range tests {
		if got := OnName(input); got != want {
			t.Errorf("OnName(%q) = %q, want %q", input, got, want)
		}
	}

	attr := On("htmx:afterSwap", "console.log(event)")
	if attr.Key != "hx-on--after-swap" || attr.Value != "console.log(event)" {
		t.Fatalf("On returned %#v", attr)
	}
}

func TestExtensionAndBooleanAttributes(t *testing.T) {
	node := el.DIV(
		Disable(),
		Preserve(),
		HistoryElt("#main"),
		WSSendAttr(),
		SSEConnectAttr("/events"),
	)
	got, err := el.RenderBytes(context.Background(), node)
	if err != nil {
		t.Fatal(err)
	}
	want := `<div hx-disable hx-preserve hx-history-elt="#main" ws-send sse-connect="/events"></div>`
	if string(got) != want {
		t.Fatalf("rendered HTML = %q, want %q", got, want)
	}
}
