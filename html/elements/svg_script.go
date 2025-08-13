package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <script> SVG element includes scripts, which can be used to trigger user
// interface events.
type SVGSCRIPTElement struct {
	*Element
}

// Create a new SVGSCRIPTElement element.
// This will create a new element with the tag
// "script" during rendering.
func SVG_SCRIPT(children ...ElementRenderer) *SVGSCRIPTElement {
	e := NewElement("script", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGSCRIPTElement{Element: e}
}

func (e *SVGSCRIPTElement) Children(children ...ElementRenderer) *SVGSCRIPTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGSCRIPTElement) IfChildren(condition bool, children ...ElementRenderer) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGSCRIPTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGSCRIPTElement) Attr(name string, value ...string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	if len(value) == 0 {
		e.StringAttributes.Set(name, "")
	} else {
		e.StringAttributes.Set(name, value[0])
	}
	return e
}

func (e *SVGSCRIPTElement) Attrs(attrs ...string) *SVGSCRIPTElement {
	if len(attrs)%2 != 0 {
		panic("attrs must be a multiple of 2")
	}
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for i := 0; i < len(attrs); i += 2 {
		k := attrs[i]
		v := attrs[i+1]
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGSCRIPTElement) AttrsMap(attrs map[string]string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGSCRIPTElement) Text(text string) *SVGSCRIPTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGSCRIPTElement) TextF(format string, args ...any) *SVGSCRIPTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfText(condition bool, text string) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGSCRIPTElement) IfTextF(condition bool, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGSCRIPTElement) Escaped(text string) *SVGSCRIPTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGSCRIPTElement) IfEscaped(condition bool, text string) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGSCRIPTElement) EscapedF(format string, args ...any) *SVGSCRIPTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfEscapedF(condition bool, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGSCRIPTElement) CustomData(key, value string) *SVGSCRIPTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGSCRIPTElement) IfCustomData(condition bool, key, value string) *SVGSCRIPTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGSCRIPTElement) CustomDataF(key, format string, args ...any) *SVGSCRIPTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGSCRIPTElement) CustomDataRemove(key string) *SVGSCRIPTElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The scripting language used for the given script element.
func (e *SVGSCRIPTElement) TYPE(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("type", s)
	return e
}

func (e *SVGSCRIPTElement) TYPEF(format string, args ...any) *SVGSCRIPTElement {
	return e.TYPE(fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfTYPE(condition bool, s string) *SVGSCRIPTElement {
	if condition {
		e.TYPE(s)
	}
	return e
}

func (e *SVGSCRIPTElement) IfTYPEF(condition bool, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.TYPE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TYPE from the element.
func (e *SVGSCRIPTElement) TYPERemove(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("type")
	return e
}

func (e *SVGSCRIPTElement) TYPERemoveF(format string, args ...any) *SVGSCRIPTElement {
	return e.TYPERemove(fmt.Sprintf(format, args...))
}

// A Uniform Resource Identifier (URI) reference that specifies the location of
// the script to execute.
func (e *SVGSCRIPTElement) HREF(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("href", s)
	return e
}

func (e *SVGSCRIPTElement) HREFF(format string, args ...any) *SVGSCRIPTElement {
	return e.HREF(fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfHREF(condition bool, s string) *SVGSCRIPTElement {
	if condition {
		e.HREF(s)
	}
	return e
}

func (e *SVGSCRIPTElement) IfHREFF(condition bool, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.HREF(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HREF from the element.
func (e *SVGSCRIPTElement) HREFRemove(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("href")
	return e
}

func (e *SVGSCRIPTElement) HREFRemoveF(format string, args ...any) *SVGSCRIPTElement {
	return e.HREFRemove(fmt.Sprintf(format, args...))
}

// How the element handles crossorigin requests.
func (e *SVGSCRIPTElement) CROSSORIGIN(c SVGScriptCrossoriginChoice) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("crossorigin", string(c))
	return e
}

type SVGScriptCrossoriginChoice string

const (
	// How the element handles crossorigin requests.
	SVGScriptCrossorigin_anonymous SVGScriptCrossoriginChoice = "anonymous"
	// How the element handles crossorigin requests.
	SVGScriptCrossorigin_use_credentials SVGScriptCrossoriginChoice = "use-credentials"
)

// Remove the attribute CROSSORIGIN from the element.
func (e *SVGSCRIPTElement) CROSSORIGINRemove(c SVGScriptCrossoriginChoice) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("crossorigin")
	return e
}

// Specifies a unique id for an element
func (e *SVGSCRIPTElement) ID(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGSCRIPTElement) IDF(format string, args ...any) *SVGSCRIPTElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfID(condition bool, s string) *SVGSCRIPTElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGSCRIPTElement) IfIDF(condition bool, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGSCRIPTElement) IDRemove(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGSCRIPTElement) IDRemoveF(format string, args ...any) *SVGSCRIPTElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGSCRIPTElement) CLASS(s ...string) *SVGSCRIPTElement {
	if e.DelimitedStrings == nil {
		e.DelimitedStrings = zarray.NewSortMap[string, *DelimitedBuilder[string]]()
	}
	ds, ok := e.DelimitedStrings.Get("class")
	if !ok {
		ds = NewDelimitedBuilder[string](" ")
		e.DelimitedStrings.Set("class", ds)
	}
	ds.Add(s...)
	return e
}

func (e *SVGSCRIPTElement) IfCLASS(condition bool, s ...string) *SVGSCRIPTElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGSCRIPTElement) CLASSRemove(s ...string) *SVGSCRIPTElement {
	if e.DelimitedStrings == nil {
		return e
	}
	ds, ok := e.DelimitedStrings.Get("class")
	if !ok {
		return e
	}
	ds.Remove(s...)
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGSCRIPTElement) STYLEF(k string, format string, args ...any) *SVGSCRIPTElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfSTYLE(condition bool, k string, v string) *SVGSCRIPTElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGSCRIPTElement) STYLE(k string, v string) *SVGSCRIPTElement {
	if e.KVStrings == nil {
		e.KVStrings = zarray.NewSortMap[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}
	kv.Add(k, v)
	return e
}

func (e *SVGSCRIPTElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGSCRIPTElement) STYLEMap(m map[string]string) *SVGSCRIPTElement {
	if e.KVStrings == nil {
		e.KVStrings = zarray.NewSortMap[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}
	for k, v := range m {
		kv.Add(k, v)
	}
	return e
}

// Add pairs of attributes to the element.
func (e *SVGSCRIPTElement) STYLEPairs(pairs ...string) *SVGSCRIPTElement {
	if len(pairs)%2 != 0 {
		panic("Must have an even number of pairs")
	}
	if e.KVStrings == nil {
		e.KVStrings = zarray.NewSortMap[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}

	for i := 0; i < len(pairs); i += 2 {
		kv.Add(pairs[i], pairs[i+1])
	}

	return e
}

func (e *SVGSCRIPTElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGSCRIPTElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGSCRIPTElement) STYLERemove(keys ...string) *SVGSCRIPTElement {
	if e.KVStrings == nil {
		return e
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		return e
	}
	for _, k := range keys {
		kv.Remove(k)
	}
	return e
}

// Make a request for an HTML

func (e *SVGSCRIPTElement) Z_REQ(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfZ_REQ(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGSCRIPTElement) Z_REQRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGSCRIPTElement) Z_TARGET(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfZ_TARGET(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGSCRIPTElement) Z_TARGETRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGSCRIPTElement) Z_REQ_SELECTOR(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGSCRIPTElement) Z_REQ_SELECTORRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGSCRIPTElement) Z_SWAP(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfZ_SWAP(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGSCRIPTElement) Z_SWAPRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGSCRIPTElement) Z_SWAP_PUSH(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGSCRIPTElement) Z_SWAP_PUSHRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGSCRIPTElement) Z_TRIGGER(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfZ_TRIGGER(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGSCRIPTElement) Z_TRIGGERRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGSCRIPTElement) Z_REQ_METHOD(c SVGScriptZReqMethodChoice) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGScriptZReqMethodChoice string

const (
	// default GET
	SVGScriptZReqMethod_empty SVGScriptZReqMethodChoice = ""
	// GET
	SVGScriptZReqMethod_get SVGScriptZReqMethodChoice = "get"
	// POST
	SVGScriptZReqMethod_post SVGScriptZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGSCRIPTElement) Z_REQ_METHODRemove(c SVGScriptZReqMethodChoice) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGSCRIPTElement) Z_REQ_STRATEGY(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGSCRIPTElement) Z_REQ_STRATEGYRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGSCRIPTElement) Z_REQ_HISTORY(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGSCRIPTElement) Z_REQ_HISTORYRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGSCRIPTElement) Z_DATA(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfZ_DATA(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGSCRIPTElement) Z_DATARemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGSCRIPTElement) Z_JSON(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfZ_JSON(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGSCRIPTElement) Z_JSONRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGSCRIPTElement) Z_REQ_BATCH(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGSCRIPTElement) Z_REQ_BATCHRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGSCRIPTElement) Z_ACTION(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfZ_ACTION(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGSCRIPTElement) Z_ACTIONRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGSCRIPTElement) Z_REQ_BEFORE(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGSCRIPTElement) Z_REQ_BEFORERemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGSCRIPTElement) Z_REQ_AFTER(expression string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSCRIPTElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGSCRIPTElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGSCRIPTElement) Z_REQ_AFTERRemove() *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
