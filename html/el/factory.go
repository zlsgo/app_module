package el

import (
	"strings"
)

// HTML元素配置注册表
var elementConfigs = map[string]*ElementConfig{
	// 文档结构元素
	"html": {
		Tag:           "html",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "The root element of an HTML document",
		Attributes: []AttributeConfig{
			{Name: "lang", Type: "string", Description: "Language of the document"},
			{Name: "dir", Type: "string", Enum: []string{"ltr", "rtl", "auto"}, Description: "Text direction"},
		},
	},
	"head": {
		Tag:           "head",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Container for document metadata",
	},
	"body": {
		Tag:           "body",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Container for document content",
	},
	"title": {
		Tag:           "title",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Document title",
	},

	// 元数据元素
	"meta": {
		Tag:         "meta",
		SelfClosing: true,
		Description: "Document metadata",
		Attributes: []AttributeConfig{
			{Name: "charset", Type: "string", Description: "Character encoding"},
			{Name: "name", Type: "string", Description: "Metadata name"},
			{Name: "content", Type: "string", Description: "Metadata content"},
			{Name: "http-equiv", Type: "string", Description: "HTTP header name"},
		},
	},
	"link": {
		Tag:         "link",
		SelfClosing: true,
		Description: "External resource link",
		Attributes: []AttributeConfig{
			{Name: "rel", Type: "string", Required: true, Description: "Relationship"},
			{Name: "href", Type: "string", Description: "Resource URL"},
			{Name: "type", Type: "string", Description: "MIME type"},
			{Name: "media", Type: "string", Description: "Media query"},
		},
	},
	"script": {
		Tag:           "script",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "JavaScript code or reference",
		Attributes: []AttributeConfig{
			{Name: "src", Type: "string", Description: "Script URL"},
			{Name: "type", Type: "string", Default: "text/javascript", Description: "Script type"},
			{Name: "async", Type: "bool", Description: "Asynchronous loading"},
			{Name: "defer", Type: "bool", Description: "Deferred execution"},
		},
	},
	"style": {
		Tag:           "style",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "CSS styles",
		Attributes: []AttributeConfig{
			{Name: "type", Type: "string", Default: "text/css", Description: "Style type"},
			{Name: "media", Type: "string", Description: "Media query"},
		},
	},

	// 区块元素
	"div": {
		Tag:           "div",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Generic container element",
	},
	"span": {
		Tag:           "span",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Generic inline element",
	},
	"p": {
		Tag:           "p",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Paragraph",
	},
	"br": {
		Tag:         "br",
		SelfClosing: true,
		Description: "Line break",
	},
	"hr": {
		Tag:         "hr",
		SelfClosing: true,
		Description: "Horizontal rule",
	},

	// 标题元素
	"h1": {
		Tag:           "h1",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Level 1 heading",
	},
	"h2": {
		Tag:           "h2",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Level 2 heading",
	},
	"h3": {
		Tag:           "h3",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Level 3 heading",
	},
	"h4": {
		Tag:           "h4",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Level 4 heading",
	},
	"h5": {
		Tag:           "h5",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Level 5 heading",
	},
	"h6": {
		Tag:           "h6",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Level 6 heading",
	},

	// 表单元素
	"form": {
		Tag:           "form",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Form container",
		Attributes: []AttributeConfig{
			{Name: "action", Type: "string", Description: "Form submission URL"},
			{Name: "method", Type: "string", Enum: []string{"get", "post", "put", "delete", "patch"}, Default: "get", Description: "HTTP method"},
			{Name: "enctype", Type: "string", Enum: []string{"application/x-www-form-urlencoded", "multipart/form-data", "text/plain"}, Description: "Encoding type"},
			{Name: "target", Type: "string", Enum: []string{"_self", "_blank", "_parent", "_top"}, Description: "Target window"},
			{Name: "novalidate", Type: "bool", Description: "Disable validation"},
		},
	},
	"input": {
		Tag:         "input",
		SelfClosing: true,
		Description: "Input field",
		Attributes: []AttributeConfig{
			{Name: "type", Type: "string", Enum: []string{"text", "password", "email", "number", "tel", "url", "search", "color", "date", "datetime-local", "month", "time", "week", "file", "hidden", "checkbox", "radio", "submit", "reset", "button"}, Default: "text", Description: "Input type"},
			{Name: "name", Type: "string", Description: "Field name"},
			{Name: "value", Type: "string", Description: "Field value"},
			{Name: "placeholder", Type: "string", Description: "Placeholder text"},
			{Name: "required", Type: "bool", Description: "Required field"},
			{Name: "disabled", Type: "bool", Description: "Disabled field"},
			{Name: "readonly", Type: "bool", Description: "Read-only field"},
			{Name: "checked", Type: "bool", Description: "Checked state (checkbox/radio)"},
			{Name: "multiple", Type: "bool", Description: "Multiple selection"},
			{Name: "min", Type: "string", Description: "Minimum value"},
			{Name: "max", Type: "string", Description: "Maximum value"},
			{Name: "step", Type: "string", Description: "Step value"},
			{Name: "pattern", Type: "string", Description: "Validation pattern"},
			{Name: "maxlength", Type: "number", Description: "Maximum length"},
			{Name: "minlength", Type: "number", Description: "Minimum length"},
		},
	},
	"textarea": {
		Tag:           "textarea",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Multi-line text input",
		Attributes: []AttributeConfig{
			{Name: "name", Type: "string", Description: "Field name"},
			{Name: "rows", Type: "number", Description: "Number of rows"},
			{Name: "cols", Type: "number", Description: "Number of columns"},
			{Name: "placeholder", Type: "string", Description: "Placeholder text"},
			{Name: "required", Type: "bool", Description: "Required field"},
			{Name: "disabled", Type: "bool", Description: "Disabled field"},
			{Name: "readonly", Type: "bool", Description: "Read-only field"},
			{Name: "maxlength", Type: "number", Description: "Maximum length"},
			{Name: "minlength", Type: "number", Description: "Minimum length"},
		},
	},
	"select": {
		Tag:           "select",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Selection dropdown",
		Attributes: []AttributeConfig{
			{Name: "name", Type: "string", Description: "Field name"},
			{Name: "required", Type: "bool", Description: "Required field"},
			{Name: "disabled", Type: "bool", Description: "Disabled field"},
			{Name: "multiple", Type: "bool", Description: "Multiple selection"},
			{Name: "size", Type: "number", Description: "Visible options"},
		},
	},
	"option": {
		Tag:           "option",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Selection option",
		Attributes: []AttributeConfig{
			{Name: "value", Type: "string", Description: "Option value"},
			{Name: "selected", Type: "bool", Description: "Selected option"},
			{Name: "disabled", Type: "bool", Description: "Disabled option"},
		},
	},
	"button": {
		Tag:           "button",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Button element",
		Attributes: []AttributeConfig{
			{Name: "type", Type: "string", Enum: []string{"button", "submit", "reset"}, Default: "button", Description: "Button type"},
			{Name: "disabled", Type: "bool", Description: "Disabled button"},
			{Name: "name", Type: "string", Description: "Button name"},
			{Name: "value", Type: "string", Description: "Button value"},
		},
	},
	"label": {
		Tag:           "label",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Form field label",
		Attributes: []AttributeConfig{
			{Name: "for", Type: "string", Description: "Associated field ID"},
		},
	},

	// 表格元素
	"table": {
		Tag:           "table",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Table container",
	},
	"thead": {
		Tag:           "thead",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Table header group",
	},
	"tbody": {
		Tag:           "tbody",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Table body group",
	},
	"tfoot": {
		Tag:           "tfoot",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Table footer group",
	},
	"tr": {
		Tag:           "tr",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Table row",
	},
	"th": {
		Tag:           "th",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Table header cell",
		Attributes: []AttributeConfig{
			{Name: "scope", Type: "string", Enum: []string{"row", "col", "rowgroup", "colgroup"}, Description: "Header scope"},
			{Name: "colspan", Type: "number", Description: "Column span"},
			{Name: "rowspan", Type: "number", Description: "Row span"},
		},
	},
	"td": {
		Tag:           "td",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Table data cell",
		Attributes: []AttributeConfig{
			{Name: "colspan", Type: "number", Description: "Column span"},
			{Name: "rowspan", Type: "number", Description: "Row span"},
		},
	},

	// 列表元素
	"ul": {
		Tag:           "ul",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Unordered list",
	},
	"ol": {
		Tag:           "ol",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Ordered list",
		Attributes: []AttributeConfig{
			{Name: "start", Type: "number", Description: "Starting number"},
			{Name: "type", Type: "string", Enum: []string{"1", "A", "a", "I", "i"}, Description: "Numbering type"},
		},
	},
	"li": {
		Tag:           "li",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "List item",
		Attributes: []AttributeConfig{
			{Name: "value", Type: "number", Description: "Item value (for ordered lists)"},
		},
	},

	// 链接和媒体元素
	"a": {
		Tag:           "a",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Hyperlink",
		Attributes: []AttributeConfig{
			{Name: "href", Type: "string", Description: "Link URL"},
			{Name: "target", Type: "string", Enum: []string{"_self", "_blank", "_parent", "_top"}, Description: "Target window"},
			{Name: "rel", Type: "string", Description: "Link relationship"},
			{Name: "download", Type: "string", Description: "Download filename"},
		},
	},
	"img": {
		Tag:         "img",
		SelfClosing: true,
		Description: "Image",
		Attributes: []AttributeConfig{
			{Name: "src", Type: "string", Required: true, Description: "Image URL"},
			{Name: "alt", Type: "string", Required: true, Description: "Alternative text"},
			{Name: "width", Type: "number", Description: "Image width"},
			{Name: "height", Type: "number", Description: "Image height"},
			{Name: "loading", Type: "string", Enum: []string{"eager", "lazy"}, Description: "Loading behavior"},
		},
	},

	// 语义化元素
	"header": {
		Tag:           "header",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Header section",
	},
	"footer": {
		Tag:           "footer",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Footer section",
	},
	"nav": {
		Tag:           "nav",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Navigation section",
	},
	"main": {
		Tag:           "main",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Main content section",
	},
	"section": {
		Tag:           "section",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Document section",
	},
	"article": {
		Tag:           "article",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Article content",
	},
	"aside": {
		Tag:           "aside",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Sidebar content",
	},

	// 文本语义元素
	"strong": {
		Tag:           "strong",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Strong emphasis",
	},
	"em": {
		Tag:           "em",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Emphasized text",
	},
	"b": {
		Tag:           "b",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Bold text",
	},
	"i": {
		Tag:           "i",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Italic text",
	},
	"u": {
		Tag:           "u",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Underlined text",
	},
	"small": {
		Tag:           "small",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Small text",
	},
	"code": {
		Tag:           "code",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Inline code",
	},
	"pre": {
		Tag:           "pre",
		SelfClosing:   false,
		AllowChildren: true,
		Description:   "Preformatted text",
	},
}

// Create 创建元素（统一接口）
func Create(tagName string, children ...ElementRenderer) *Element {
	tagName = strings.ToLower(tagName)
	config, ok := elementConfigs[tagName]
	if !ok {
		// 未知元素，创建基本配置
		config = &ElementConfig{
			Tag:           tagName,
			SelfClosing:   false,
			AllowChildren: true,
			Description:   "Custom element: " + tagName,
		}
	}

	return NewElementWithConfig(config, children...)
}

func HTML(children ...ElementRenderer) *Element {
	return Create("html", children...)
}

func HEAD(children ...ElementRenderer) *Element {
	return Create("head", children...)
}

func TITLE(children ...ElementRenderer) *Element {
	return Create("title", children...)
}

func META(children ...ElementRenderer) *Element {
	return Create("meta", children...)
}

func BODY(children ...ElementRenderer) *Element {
	return Create("body", children...)
}

func DIV(children ...ElementRenderer) *Element {
	return Create("div", children...)
}

func SCRIPT(children ...ElementRenderer) *Element {
	return Create("script", children...)
}

func LABEL(children ...ElementRenderer) *Element {
	return Create("label", children...)
}

func LINK(children ...ElementRenderer) *Element {
	return Create("link", children...)
}

func SPAN(children ...ElementRenderer) *Element {
	return Create("span", children...)
}

func P(children ...ElementRenderer) *Element {
	return Create("p", children...)
}

func H1(children ...ElementRenderer) *Element {
	return Create("h1", children...)
}

func H2(children ...ElementRenderer) *Element {
	return Create("h2", children...)
}

func H3(children ...ElementRenderer) *Element {
	return Create("h3", children...)
}

func A(children ...ElementRenderer) *Element {
	return Create("a", children...)
}

func BUTTON(children ...ElementRenderer) *Element {
	return Create("button", children...)
}

func INPUT() *Element {
	return Create("input")
}

func TEXTAREA(children ...ElementRenderer) *Element {
	return Create("textarea", children...)
}

func SELECT(children ...ElementRenderer) *Element {
	return Create("select", children...)
}

func OPTION(children ...ElementRenderer) *Element {
	return Create("option", children...)
}

func FORM(children ...ElementRenderer) *Element {
	return Create("form", children...)
}

func TABLE(children ...ElementRenderer) *Element {
	return Create("table", children...)
}

func TR(children ...ElementRenderer) *Element {
	return Create("tr", children...)
}

func TD(children ...ElementRenderer) *Element {
	return Create("td", children...)
}

func TH(children ...ElementRenderer) *Element {
	return Create("th", children...)
}

func UL(children ...ElementRenderer) *Element {
	return Create("ul", children...)
}

func OL(children ...ElementRenderer) *Element {
	return Create("ol", children...)
}

func LI(children ...ElementRenderer) *Element {
	return Create("li", children...)
}

func IMG() *Element {
	return Create("img")
}

func BR() *Element {
	return Create("br")
}

func HR() *Element {
	return Create("hr")
}

// RegisterElement 注册自定义元素配置
func RegisterElement(tagName string, config *ElementConfig) {
	elementConfigs[strings.ToLower(tagName)] = config
}

// GetElementConfig 获取元素配置
func GetElementConfig(tagName string) (*ElementConfig, bool) {
	config, ok := elementConfigs[strings.ToLower(tagName)]
	return config, ok
}

// ListSupportedElements 列出支持的元素
func ListSupportedElements() []string {
	elements := make([]string, 0, len(elementConfigs))
	for tag := range elementConfigs {
		elements = append(elements, tag)
	}
	return elements
}
