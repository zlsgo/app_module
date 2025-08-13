package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <view> SVG element is used to define a view into a &lt;svg&gt; element
// It is partially deprecated in SVG 2.0 and should generally not be used.
type SVGVIEWElement struct {
	*Element
}

// Create a new SVGVIEWElement element.
// This will create a new element with the tag
// "view" during rendering.
func SVG_VIEW(children ...ElementRenderer) *SVGVIEWElement {
	e := NewElement("view", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGVIEWElement{Element: e}
}

func (e *SVGVIEWElement) Children(children ...ElementRenderer) *SVGVIEWElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGVIEWElement) IfChildren(condition bool, children ...ElementRenderer) *SVGVIEWElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGVIEWElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGVIEWElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGVIEWElement) Attr(name string, value ...string) *SVGVIEWElement {
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

func (e *SVGVIEWElement) Attrs(attrs ...string) *SVGVIEWElement {
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

func (e *SVGVIEWElement) AttrsMap(attrs map[string]string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGVIEWElement) Text(text string) *SVGVIEWElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGVIEWElement) TextF(format string, args ...any) *SVGVIEWElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGVIEWElement) IfText(condition bool, text string) *SVGVIEWElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGVIEWElement) IfTextF(condition bool, format string, args ...any) *SVGVIEWElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGVIEWElement) Escaped(text string) *SVGVIEWElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGVIEWElement) IfEscaped(condition bool, text string) *SVGVIEWElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGVIEWElement) EscapedF(format string, args ...any) *SVGVIEWElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGVIEWElement) IfEscapedF(condition bool, format string, args ...any) *SVGVIEWElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGVIEWElement) CustomData(key, value string) *SVGVIEWElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGVIEWElement) IfCustomData(condition bool, key, value string) *SVGVIEWElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGVIEWElement) CustomDataF(key, format string, args ...any) *SVGVIEWElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGVIEWElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGVIEWElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGVIEWElement) CustomDataRemove(key string) *SVGVIEWElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The position and size of the viewport (the viewBox) is defined by the viewBox
// attribute.
func (e *SVGVIEWElement) VIEW_BOX(s string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("viewBox", s)
	return e
}

func (e *SVGVIEWElement) VIEW_BOXF(format string, args ...any) *SVGVIEWElement {
	return e.VIEW_BOX(fmt.Sprintf(format, args...))
}

func (e *SVGVIEWElement) IfVIEW_BOX(condition bool, s string) *SVGVIEWElement {
	if condition {
		e.VIEW_BOX(s)
	}
	return e
}

func (e *SVGVIEWElement) IfVIEW_BOXF(condition bool, format string, args ...any) *SVGVIEWElement {
	if condition {
		e.VIEW_BOX(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute VIEW_BOX from the element.
func (e *SVGVIEWElement) VIEW_BOXRemove(s string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("viewBox")
	return e
}

func (e *SVGVIEWElement) VIEW_BOXRemoveF(format string, args ...any) *SVGVIEWElement {
	return e.VIEW_BOXRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGVIEWElement) ID(s string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGVIEWElement) IDF(format string, args ...any) *SVGVIEWElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGVIEWElement) IfID(condition bool, s string) *SVGVIEWElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGVIEWElement) IfIDF(condition bool, format string, args ...any) *SVGVIEWElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGVIEWElement) IDRemove(s string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGVIEWElement) IDRemoveF(format string, args ...any) *SVGVIEWElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGVIEWElement) CLASS(s ...string) *SVGVIEWElement {
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

func (e *SVGVIEWElement) IfCLASS(condition bool, s ...string) *SVGVIEWElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGVIEWElement) CLASSRemove(s ...string) *SVGVIEWElement {
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
func (e *SVGVIEWElement) STYLEF(k string, format string, args ...any) *SVGVIEWElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGVIEWElement) IfSTYLE(condition bool, k string, v string) *SVGVIEWElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGVIEWElement) STYLE(k string, v string) *SVGVIEWElement {
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

func (e *SVGVIEWElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGVIEWElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGVIEWElement) STYLEMap(m map[string]string) *SVGVIEWElement {
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
func (e *SVGVIEWElement) STYLEPairs(pairs ...string) *SVGVIEWElement {
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

func (e *SVGVIEWElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGVIEWElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGVIEWElement) STYLERemove(keys ...string) *SVGVIEWElement {
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

func (e *SVGVIEWElement) Z_REQ(expression string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGVIEWElement) IfZ_REQ(condition bool, expression string) *SVGVIEWElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGVIEWElement) Z_REQRemove() *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGVIEWElement) Z_TARGET(expression string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGVIEWElement) IfZ_TARGET(condition bool, expression string) *SVGVIEWElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGVIEWElement) Z_TARGETRemove() *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGVIEWElement) Z_REQ_SELECTOR(expression string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGVIEWElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGVIEWElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGVIEWElement) Z_REQ_SELECTORRemove() *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGVIEWElement) Z_SWAP(expression string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGVIEWElement) IfZ_SWAP(condition bool, expression string) *SVGVIEWElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGVIEWElement) Z_SWAPRemove() *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGVIEWElement) Z_SWAP_PUSH(expression string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGVIEWElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGVIEWElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGVIEWElement) Z_SWAP_PUSHRemove() *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGVIEWElement) Z_TRIGGER(expression string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGVIEWElement) IfZ_TRIGGER(condition bool, expression string) *SVGVIEWElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGVIEWElement) Z_TRIGGERRemove() *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGVIEWElement) Z_REQ_METHOD(c SVGViewZReqMethodChoice) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGViewZReqMethodChoice string

const (
	// default GET
	SVGViewZReqMethod_empty SVGViewZReqMethodChoice = ""
	// GET
	SVGViewZReqMethod_get SVGViewZReqMethodChoice = "get"
	// POST
	SVGViewZReqMethod_post SVGViewZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGVIEWElement) Z_REQ_METHODRemove(c SVGViewZReqMethodChoice) *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGVIEWElement) Z_REQ_STRATEGY(expression string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGVIEWElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGVIEWElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGVIEWElement) Z_REQ_STRATEGYRemove() *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGVIEWElement) Z_REQ_HISTORY(expression string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGVIEWElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGVIEWElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGVIEWElement) Z_REQ_HISTORYRemove() *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGVIEWElement) Z_DATA(expression string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGVIEWElement) IfZ_DATA(condition bool, expression string) *SVGVIEWElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGVIEWElement) Z_DATARemove() *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGVIEWElement) Z_JSON(expression string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGVIEWElement) IfZ_JSON(condition bool, expression string) *SVGVIEWElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGVIEWElement) Z_JSONRemove() *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGVIEWElement) Z_REQ_BATCH(expression string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGVIEWElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGVIEWElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGVIEWElement) Z_REQ_BATCHRemove() *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGVIEWElement) Z_ACTION(expression string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGVIEWElement) IfZ_ACTION(condition bool, expression string) *SVGVIEWElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGVIEWElement) Z_ACTIONRemove() *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGVIEWElement) Z_REQ_BEFORE(expression string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGVIEWElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGVIEWElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGVIEWElement) Z_REQ_BEFORERemove() *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGVIEWElement) Z_REQ_AFTER(expression string) *SVGVIEWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGVIEWElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGVIEWElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGVIEWElement) Z_REQ_AFTERRemove() *SVGVIEWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
