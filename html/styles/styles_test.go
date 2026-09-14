package styles

import (
	"strings"
	"testing"

	"github.com/zlsgo/app_module/html/el"
)

func TestPropsInlineAndAttr(t *testing.T) {
	props := Props{BackgroundColor: "blue", Color: "white", FontSize: Pixels(16)}
	if got, want := props.ToInline(), "background-color: blue; color: white; font-size: 16px;"; got != want {
		t.Fatalf("ToInline() = %q, want %q", got, want)
	}
	got, err := el.RenderBytes(nil, el.DIV(props.Attr()))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != `<div style="background-color: blue; color: white; font-size: 16px;"></div>` {
		t.Fatalf("style attr = %q", got)
	}
}

func TestStyleManagerDeterministicAndDeduplicated(t *testing.T) {
	manager := NewStyleManager()
	style := Props{Color: "red", BackgroundColor: "white"}
	first := manager.AddStyle(style)
	second := manager.AddStyle(Props{BackgroundColor: "white", Color: "red"})
	if first != second {
		t.Fatalf("equal style classes differ: %q != %q", first, second)
	}
	css1 := manager.GenerateCSS()
	css2 := manager.GenerateCSS()
	if css1 != css2 || !strings.Contains(css1, "."+first+" { background-color: white; color: red; }") {
		t.Fatalf("CSS is not deterministic or missing rule: %q", css1)
	}
}

func TestStyleManagerZeroValue(t *testing.T) {
	var manager StyleManager
	name := manager.AddStyle(Props{Color: "red"})
	if !strings.Contains(manager.GenerateCSS(), "."+name+" {") {
		t.Fatalf("zero-value manager did not register style: %q", manager.GenerateCSS())
	}
}

func TestStyleManagerElementHelpers(t *testing.T) {
	manager := NewStyleManager()
	class := manager.Class(Props{Color: "red"})
	node := el.DIV(class)
	if got := node.GetAttribute("class"); got == "" {
		t.Fatal("Class helper did not set class")
	}
	composite := manager.CompositeClass(CompositeStyle{Default: Props{Color: "blue"}})
	if got := composite.Value; got == "" {
		t.Fatal("CompositeClass helper returned empty class")
	}
}

func TestStyleManagerCompositeAndAnimation(t *testing.T) {
	manager := NewStyleManager()
	animation := manager.AddAnimation(Keyframes{"to": {Opacity: "0"}, "from": {Opacity: "1"}})
	className := manager.AddCompositeStyle(CompositeStyle{
		Default:        Props{Color: "red"},
		PseudoClasses:  map[string]Props{"hover": {Color: "blue"}},
		PseudoElements: map[string]Props{"before": {Content: `"x"`}},
		MediaQueries:   map[string]Props{"(min-width: 600px)": {Color: "green"}},
	})
	css := manager.GenerateCSS()
	for _, part := range []string{
		"@keyframes " + animation,
		"." + className + " { color: red; }",
		"." + className + ":hover { color: blue; }",
		"." + className + "::before { content: \"x\"; }",
		"@media (min-width: 600px) { ." + className + " { color: green; } }",
	} {
		if !strings.Contains(css, part) {
			t.Errorf("CSS missing %q: %q", part, css)
		}
	}
}

func TestStyleHelpers(t *testing.T) {
	cases := map[string]string{
		Em(1.5): "1.50em", Rem(2): "2.00rem", Pixels(0): "0", Percent(50): "50%",
		RGB(1, 2, 3): "rgb(1,2,3)", URL("a'b"): `url('a\'b')`, Var("--main"): "var(--main)",
	}
	for got, want := range cases {
		if got != want {
			t.Errorf("helper = %q, want %q", got, want)
		}
	}
}
