package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <defs> SVG element is used to embed definitions that can be reused inside
// an SVG image.
type SVGDEFSElement struct {
	*Element
}

// Create a new SVGDEFSElement element.
// This will create a new element with the tag
// "defs" during rendering.
func SVG_DEFS(children ...ElementRenderer) *SVGDEFSElement {
	e := NewElement("defs", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGDEFSElement{Element: e}
}

func (e *SVGDEFSElement) Children(children ...ElementRenderer) *SVGDEFSElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGDEFSElement) IfChildren(condition bool, children ...ElementRenderer) *SVGDEFSElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGDEFSElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGDEFSElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGDEFSElement) Attr(name string, value ...string) *SVGDEFSElement {
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

func (e *SVGDEFSElement) Attrs(attrs ...string) *SVGDEFSElement {
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

func (e *SVGDEFSElement) AttrsMap(attrs map[string]string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGDEFSElement) Text(text string) *SVGDEFSElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGDEFSElement) TextF(format string, args ...any) *SVGDEFSElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGDEFSElement) IfText(condition bool, text string) *SVGDEFSElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGDEFSElement) IfTextF(condition bool, format string, args ...any) *SVGDEFSElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGDEFSElement) Escaped(text string) *SVGDEFSElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGDEFSElement) IfEscaped(condition bool, text string) *SVGDEFSElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGDEFSElement) EscapedF(format string, args ...any) *SVGDEFSElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGDEFSElement) IfEscapedF(condition bool, format string, args ...any) *SVGDEFSElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGDEFSElement) CustomData(key, value string) *SVGDEFSElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGDEFSElement) IfCustomData(condition bool, key, value string) *SVGDEFSElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGDEFSElement) CustomDataF(key, format string, args ...any) *SVGDEFSElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGDEFSElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGDEFSElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGDEFSElement) CustomDataRemove(key string) *SVGDEFSElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// Specifies a unique id for an element
func (e *SVGDEFSElement) ID(s string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGDEFSElement) IDF(format string, args ...any) *SVGDEFSElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGDEFSElement) IfID(condition bool, s string) *SVGDEFSElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGDEFSElement) IfIDF(condition bool, format string, args ...any) *SVGDEFSElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGDEFSElement) IDRemove(s string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGDEFSElement) IDRemoveF(format string, args ...any) *SVGDEFSElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGDEFSElement) CLASS(s ...string) *SVGDEFSElement {
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

func (e *SVGDEFSElement) IfCLASS(condition bool, s ...string) *SVGDEFSElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGDEFSElement) CLASSRemove(s ...string) *SVGDEFSElement {
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
func (e *SVGDEFSElement) STYLEF(k string, format string, args ...any) *SVGDEFSElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGDEFSElement) IfSTYLE(condition bool, k string, v string) *SVGDEFSElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGDEFSElement) STYLE(k string, v string) *SVGDEFSElement {
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

func (e *SVGDEFSElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGDEFSElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGDEFSElement) STYLEMap(m map[string]string) *SVGDEFSElement {
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
func (e *SVGDEFSElement) STYLEPairs(pairs ...string) *SVGDEFSElement {
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

func (e *SVGDEFSElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGDEFSElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGDEFSElement) STYLERemove(keys ...string) *SVGDEFSElement {
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

func (e *SVGDEFSElement) Z_REQ(expression string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDEFSElement) IfZ_REQ(condition bool, expression string) *SVGDEFSElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGDEFSElement) Z_REQRemove() *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGDEFSElement) Z_TARGET(expression string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDEFSElement) IfZ_TARGET(condition bool, expression string) *SVGDEFSElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGDEFSElement) Z_TARGETRemove() *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGDEFSElement) Z_REQ_SELECTOR(expression string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDEFSElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGDEFSElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGDEFSElement) Z_REQ_SELECTORRemove() *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGDEFSElement) Z_SWAP(expression string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDEFSElement) IfZ_SWAP(condition bool, expression string) *SVGDEFSElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGDEFSElement) Z_SWAPRemove() *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGDEFSElement) Z_SWAP_PUSH(expression string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDEFSElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGDEFSElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGDEFSElement) Z_SWAP_PUSHRemove() *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGDEFSElement) Z_TRIGGER(expression string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDEFSElement) IfZ_TRIGGER(condition bool, expression string) *SVGDEFSElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGDEFSElement) Z_TRIGGERRemove() *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGDEFSElement) Z_REQ_METHOD(c SVGDefsZReqMethodChoice) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGDefsZReqMethodChoice string

const (
	// default GET
	SVGDefsZReqMethod_empty SVGDefsZReqMethodChoice = ""
	// GET
	SVGDefsZReqMethod_get SVGDefsZReqMethodChoice = "get"
	// POST
	SVGDefsZReqMethod_post SVGDefsZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGDEFSElement) Z_REQ_METHODRemove(c SVGDefsZReqMethodChoice) *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGDEFSElement) Z_REQ_STRATEGY(expression string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDEFSElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGDEFSElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGDEFSElement) Z_REQ_STRATEGYRemove() *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGDEFSElement) Z_REQ_HISTORY(expression string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDEFSElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGDEFSElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGDEFSElement) Z_REQ_HISTORYRemove() *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGDEFSElement) Z_DATA(expression string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDEFSElement) IfZ_DATA(condition bool, expression string) *SVGDEFSElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGDEFSElement) Z_DATARemove() *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGDEFSElement) Z_JSON(expression string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDEFSElement) IfZ_JSON(condition bool, expression string) *SVGDEFSElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGDEFSElement) Z_JSONRemove() *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGDEFSElement) Z_REQ_BATCH(expression string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDEFSElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGDEFSElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGDEFSElement) Z_REQ_BATCHRemove() *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGDEFSElement) Z_ACTION(expression string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDEFSElement) IfZ_ACTION(condition bool, expression string) *SVGDEFSElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGDEFSElement) Z_ACTIONRemove() *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGDEFSElement) Z_REQ_BEFORE(expression string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDEFSElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGDEFSElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGDEFSElement) Z_REQ_BEFORERemove() *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGDEFSElement) Z_REQ_AFTER(expression string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDEFSElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGDEFSElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGDEFSElement) Z_REQ_AFTERRemove() *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
