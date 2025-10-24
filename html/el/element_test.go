package el

import (
	"testing"

	"github.com/sohaha/zlsgo"
)

func TestElementSetGet(t *testing.T) {
	tt := zlsgo.NewTest(t)

	el := DIV()
	el.Set("custom", "value")
	tt.Equal("value", el.Get("custom"))

	el.Set("custom", "new")
	tt.Equal("new", el.Get("custom"))
}

func TestElementSetAttribute(t *testing.T) {
	tt := zlsgo.NewTest(t)

	t.Run("set new attribute", func(t *testing.T) {
		el := DIV()
		el.SetAttribute("title", "test")
		tt.Equal("test", el.GetAttribute("title"))
	})

	t.Run("overwrite attribute", func(t *testing.T) {
		el := DIV(Attr("id", "old"))
		tt.Equal("old", el.GetAttribute("id"))
		el.SetAttribute("id", "new")
		tt.Equal("new", el.GetAttribute("id"))
	})
}

func TestElementAddClass(t *testing.T) {
	tt := zlsgo.NewTest(t)

	t.Run("add to empty", func(t *testing.T) {
		el := DIV()
		el.AddClass("foo")
		tt.Equal("foo", el.GetAttribute("class"))
	})

	t.Run("add to existing", func(t *testing.T) {
		el := DIV(Class("bar"))
		el.AddClass("baz")
		tt.Equal("bar baz", el.GetAttribute("class"))
	})
}

func TestElementGetAttributes(t *testing.T) {
	tt := zlsgo.NewTest(t)

	el := DIV(ID("app"), Class("container"))
	attrs := el.GetAttributes()

	tt.Equal(2, len(attrs))
	tt.Equal("app", attrs["id"])
	tt.Equal("container", attrs["class"])
}

func TestElementSetAttributes(t *testing.T) {
	tt := zlsgo.NewTest(t)

	el := DIV(ID("old"))
	el.SetAttributes(map[string]string{
		"id":    "new",
		"class": "active",
	})

	tt.Equal("new", el.GetAttribute("id"))
	tt.Equal("active", el.GetAttribute("class"))
}

func TestElementAppendNode(t *testing.T) {
	tt := zlsgo.NewTest(t)

	el := DIV()
	tt.Equal(0, len(el.GetNodes()))

	el.AppendNode(Text("hello"))
	tt.Equal(1, len(el.GetNodes()))

	el.AppendNode(SPAN(Text("world")))
	tt.Equal(2, len(el.GetNodes()))

	ChunkTest{
		Node:     el,
		Rendered: "<div>hello<span>world</span></div>",
	}.Assert(tt)
}

func TestElementClone(t *testing.T) {
	tt := zlsgo.NewTest(t)

	original := DIV(
		ID("original"),
		Class("foo"),
		Text("content"),
	)
	original.Set("meta", "data")

	clone := original.Clone()
	clone.SetAttribute("id", "cloned")
	clone.AddClass("bar")

	tt.Equal("original", original.GetAttribute("id"))
	tt.Equal("foo", original.GetAttribute("class"))

	tt.Equal("cloned", clone.GetAttribute("id"))
	tt.Equal("foo bar", clone.GetAttribute("class"))

	tt.Equal(len(original.GetNodes()), len(clone.GetNodes()))
}

func TestElementTag(t *testing.T) {
	tt := zlsgo.NewTest(t)

	t.Run("normal tag", func(t *testing.T) {
		el := El("DiV")
		tt.Equal("div", el.Tag())
	})

	t.Run("doctype", func(t *testing.T) {
		el := VoidEl("!DOCTYPE", Attr("html", ""))
		tt.Equal("!DOCTYPE", el.Tag())
	})
}

func TestVoidElement(t *testing.T) {
	tt := zlsgo.NewTest(t)

	tests := []struct {
		node     Node
		rendered string
	}{
		{BR(), "<br>"},
		{HR(), "<hr>"},
		{IMG(Alt("test")), `<img alt="test">`},
		{INPUT(Type("text")), `<input type="text">`},
		{LINK(Href("style.css")), `<link href="style.css">`},
		{META(Charset("utf-8")), `<meta charset="utf-8">`},
	}

	for _, test := range tests {
		ChunkTest{
			Node:     test.node,
			Rendered: test.rendered,
		}.Assert(tt)
	}
}

func TestCommonElements(t *testing.T) {
	tt := zlsgo.NewTest(t)

	tests := []struct {
		node     Node
		rendered string
	}{
		{A(Href("#"), Text("link")), `<a href="#">link</a>`},
		{ABBR(Title("title"), Text("abbr")), `<abbr title="title">abbr</abbr>`},
		{ADDRESS(Text("addr")), "<address>addr</address>"},
		{ARTICLE(Text("art")), "<article>art</article>"},
		{ASIDE(Text("aside")), "<aside>aside</aside>"},
		{B(Text("bold")), "<b>bold</b>"},
		{BUTTON(Text("click")), "<button>click</button>"},
		{CODE(Text("code")), "<code>code</code>"},
		{EM(Text("emphasis")), "<em>emphasis</em>"},
		{FOOTER(Text("foot")), "<footer>foot</footer>"},
		{FORM(Action("/submit")), `<form action="/submit"></form>`},
		{H1(Text("h1")), "<h1>h1</h1>"},
		{H2(Text("h2")), "<h2>h2</h2>"},
		{H3(Text("h3")), "<h3>h3</h3>"},
		{H4(Text("h4")), "<h4>h4</h4>"},
		{H5(Text("h5")), "<h5>h5</h5>"},
		{H6(Text("h6")), "<h6>h6</h6>"},
		{HEADER(Text("head")), "<header>head</header>"},
		{I(Text("italic")), "<i>italic</i>"},
		{LABEL(For("input")), `<label for="input"></label>`},
		{LI(Text("item")), "<li>item</li>"},
		{MAIN(Text("main")), "<main>main</main>"},
		{MARK(Text("mark")), "<mark>mark</mark>"},
		{NAV(Text("nav")), "<nav>nav</nav>"},
		{OL(LI(Text("1"))), "<ol><li>1</li></ol>"},
		{OPTION(Value("v"), Text("opt")), `<option value="v">opt</option>`},
		{PRE(Text("pre")), "<pre>pre</pre>"},
		{Q(Text("quote")), "<q>quote</q>"},
		{S(Text("strike")), "<s>strike</s>"},
		{SECTION(Text("sec")), "<section>sec</section>"},
		{SELECT(OPTION(Text("opt"))), "<select><option>opt</option></select>"},
		{SMALL(Text("small")), "<small>small</small>"},
		{STRONG(Text("strong")), "<strong>strong</strong>"},
		{SUB(Text("sub")), "<sub>sub</sub>"},
		{SUP(Text("sup")), "<sup>sup</sup>"},
		{TABLE(TR(TD(Text("cell")))), "<table><tr><td>cell</td></tr></table>"},
		{TBODY(TR(TD(Text("b")))), "<tbody><tr><td>b</td></tr></tbody>"},
		{TEXTAREA(Text("text")), "<textarea>text</textarea>"},
		{TFOOT(TR(TD(Text("f")))), "<tfoot><tr><td>f</td></tr></tfoot>"},
		{TH(Text("h")), "<th>h</th>"},
		{THEAD(TR(TH(Text("h")))), "<thead><tr><th>h</th></tr></thead>"},
		{U(Text("under")), "<u>under</u>"},
		{UL(LI(Text("item"))), "<ul><li>item</li></ul>"},
	}

	for _, test := range tests {
		ChunkTest{
			Node:     test.node,
			Rendered: test.rendered,
		}.Assert(tt)
	}
}

func TestCommonAttributes(t *testing.T) {
	tt := zlsgo.NewTest(t)

	tests := []struct {
		node     Node
		rendered string
	}{
		{DIV(Accept("image/*")), `<div accept="image/*"></div>`},
		{DIV(AcceptCharset("utf-8")), `<div accept-charset="utf-8"></div>`},
		{DIV(AccessKey("a")), `<div accesskey="a"></div>`},
		{DIV(Action("/path")), `<div action="/path"></div>`},
		{DIV(Align("center")), `<div align="center"></div>`},
		{IMG(Alt("image")), `<img alt="image">`},
		{DIV(Aria("label", "test")), `<div aria-label="test"></div>`},
		{DIV(Async("true")), `<div async="true"></div>`},
		{DIV(Autocomplete("off")), `<div autocomplete="off"></div>`},
		{DIV(Autofocus("true")), `<div autofocus="true"></div>`},
		{DIV(Autoplay("true")), `<div autoplay="true"></div>`},
		{DIV(Charset("utf-8")), `<div charset="utf-8"></div>`},
		{DIV(Color("red")), `<div color="red"></div>`},
		{DIV(Cols("10")), `<div cols="10"></div>`},
		{DIV(Colspan("2")), `<div colspan="2"></div>`},
		{DIV(Content("text")), `<div content="text"></div>`},
		{DIV(Controls("true")), `<div controls="true"></div>`},
		{DIV(Coords("0,0")), `<div coords="0,0"></div>`},
		{DIV(Crossorigin("anonymous")), `<div crossorigin="anonymous"></div>`},
		{DIV(Datetime("2023-01-01")), `<div datetime="2023-01-01"></div>`},
		{DIV(Default("true")), `<div default="true"></div>`},
		{DIV(Dirname("dir")), `<div dirname="dir"></div>`},
		{DIV(Disabled), `<div disabled></div>`},
		{DIV(Download("file")), `<div download="file"></div>`},
		{DIV(Draggable("true")), `<div draggable="true"></div>`},
		{DIV(Enctype("multipart/form-data")), `<div enctype="multipart/form-data"></div>`},
		{DIV(EnterKeyHint("go")), `<div enterkeyhint="go"></div>`},
		{DIV(FormAction("/action")), `<div formaction="/action"></div>`},
		{DIV(Headers("h1 h2")), `<div headers="h1 h2"></div>`},
		{DIV(Height("100")), `<div height="100"></div>`},
		{DIV(Hidden("true")), `<div hidden="true"></div>`},
		{DIV(High("90")), `<div high="90"></div>`},
		{DIV(Href("/page")), `<div href="/page"></div>`},
		{DIV(HrefLang("en")), `<div hreflang="en"></div>`},
		{DIV(HttpEquiv("refresh")), `<div http-equiv="refresh"></div>`},
		{DIV(Inert("true")), `<div inert="true"></div>`},
		{DIV(InputMode("numeric")), `<div inputmode="numeric"></div>`},
		{DIV(IsMap("true")), `<div ismap="true"></div>`},
		{DIV(Kind("captions")), `<div kind="captions"></div>`},
		{DIV(Label("text")), `<div label="text"></div>`},
		{DIV(Src("/img.png")), `<div src="/img.png"></div>`},
		{DIV(Role("button")), `<div role="button"></div>`},
		{DIV(Lang("en")), `<div lang="en"></div>`},
		{DIV(List("list1")), `<div list="list1"></div>`},
		{DIV(Loop("true")), `<div loop="true"></div>`},
		{DIV(Low("10")), `<div low="10"></div>`},
		{DIV(Max("100")), `<div max="100"></div>`},
		{DIV(MaxLength("255")), `<div maxlength="255"></div>`},
		{DIV(Media("screen")), `<div media="screen"></div>`},
		{DIV(Method("post")), `<div method="post"></div>`},
		{DIV(Min("0")), `<div min="0"></div>`},
		{DIV(MinLength("5")), `<div minlength="5"></div>`},
		{DIV(Multiple("true")), `<div multiple="true"></div>`},
		{DIV(Muted("true")), `<div muted="true"></div>`},
		{DIV(Name("field")), `<div name="field"></div>`},
		{DIV(Placeholder("hint")), `<div placeholder="hint"></div>`},
		{DIV(Tabindex("1")), `<div tabindex="1"></div>`},
		{DIV(Type("text")), `<div type="text"></div>`},
		{DIV(Rel("stylesheet")), `<div rel="stylesheet"></div>`},
		{DIV(Width("200")), `<div width="200"></div>`},
		{DIV(Value("test")), `<div value="test"></div>`},
		{DIV(ShadowRootMode("open")), `<div shadowrootmode="open"></div>`},
		{DIV(Slot("content")), `<div slot="content"></div>`},
		{DIV(Property("og:title"), Content("Title")), `<div property="og:title" content="Title"></div>`},
	}

	for _, test := range tests {
		ChunkTest{
			Node:     test.node,
			Rendered: test.rendered,
		}.Assert(tt)
	}
}

func TestSpecialAttributes(t *testing.T) {
	tt := zlsgo.NewTest(t)

	t.Run("ContentEditable", func(t *testing.T) {
		node := DIV(ContentEditable)
		ChunkTest{
			Node:     node,
			Rendered: `<div contenteditable></div>`,
		}.Assert(tt)
	})

	t.Run("Checked", func(t *testing.T) {
		node := INPUT(Type("checkbox"), Checked)
		ChunkTest{
			Node:     node,
			Rendered: `<input type="checkbox" checked>`,
		}.Assert(tt)
	})

	t.Run("Defer", func(t *testing.T) {
		node := SCRIPT(Defer)
		ChunkTest{
			Node:     node,
			Rendered: `<script defer></script>`,
		}.Assert(tt)
	})

	t.Run("As", func(t *testing.T) {
		node := LINK(As("style"))
		ChunkTest{
			Node:     node,
			Rendered: `<link as="style">`,
		}.Assert(tt)
	})

	t.Run("FormAttr", func(t *testing.T) {
		node := INPUT(Form("form1"))
		ChunkTest{
			Node:     node,
			Rendered: `<input form="form1">`,
		}.Assert(tt)
	})

	t.Run("For", func(t *testing.T) {
		node := LABEL(For("input1"))
		ChunkTest{
			Node:     node,
			Rendered: `<label for="input1"></label>`,
		}.Assert(tt)
	})
}
