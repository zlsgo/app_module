package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <desc> SVG element provides a description container for SVG content.
type SVGDESCElement struct {
	*Element
}

// Create a new SVGDESCElement element.
// This will create a new element with the tag
// "desc" during rendering.
func SVG_DESC(children ...ElementRenderer) *SVGDESCElement {
	e := NewElement("desc", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGDESCElement{Element: e}
}

func (e *SVGDESCElement) Children(children ...ElementRenderer) *SVGDESCElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGDESCElement) IfChildren(condition bool, children ...ElementRenderer) *SVGDESCElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGDESCElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGDESCElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGDESCElement) Attr(name string, value ...string) *SVGDESCElement {
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

func (e *SVGDESCElement) Attrs(attrs ...string) *SVGDESCElement {
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

func (e *SVGDESCElement) AttrsMap(attrs map[string]string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGDESCElement) Text(text string) *SVGDESCElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGDESCElement) TextF(format string, args ...any) *SVGDESCElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGDESCElement) IfText(condition bool, text string) *SVGDESCElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGDESCElement) IfTextF(condition bool, format string, args ...any) *SVGDESCElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGDESCElement) Escaped(text string) *SVGDESCElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGDESCElement) IfEscaped(condition bool, text string) *SVGDESCElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGDESCElement) EscapedF(format string, args ...any) *SVGDESCElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGDESCElement) IfEscapedF(condition bool, format string, args ...any) *SVGDESCElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGDESCElement) CustomData(key, value string) *SVGDESCElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGDESCElement) IfCustomData(condition bool, key, value string) *SVGDESCElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGDESCElement) CustomDataF(key, format string, args ...any) *SVGDESCElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGDESCElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGDESCElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGDESCElement) CustomDataRemove(key string) *SVGDESCElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// Specifies a unique id for an element
func (e *SVGDESCElement) ID(s string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGDESCElement) IDF(format string, args ...any) *SVGDESCElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGDESCElement) IfID(condition bool, s string) *SVGDESCElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGDESCElement) IfIDF(condition bool, format string, args ...any) *SVGDESCElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGDESCElement) IDRemove(s string) *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGDESCElement) IDRemoveF(format string, args ...any) *SVGDESCElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGDESCElement) CLASS(s ...string) *SVGDESCElement {
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

func (e *SVGDESCElement) IfCLASS(condition bool, s ...string) *SVGDESCElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGDESCElement) CLASSRemove(s ...string) *SVGDESCElement {
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
func (e *SVGDESCElement) STYLEF(k string, format string, args ...any) *SVGDESCElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGDESCElement) IfSTYLE(condition bool, k string, v string) *SVGDESCElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGDESCElement) STYLE(k string, v string) *SVGDESCElement {
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

func (e *SVGDESCElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGDESCElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGDESCElement) STYLEMap(m map[string]string) *SVGDESCElement {
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
func (e *SVGDESCElement) STYLEPairs(pairs ...string) *SVGDESCElement {
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

func (e *SVGDESCElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGDESCElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGDESCElement) STYLERemove(keys ...string) *SVGDESCElement {
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

func (e *SVGDESCElement) Z_REQ(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfZ_REQ(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGDESCElement) Z_REQRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGDESCElement) Z_TARGET(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfZ_TARGET(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGDESCElement) Z_TARGETRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGDESCElement) Z_REQ_SELECTOR(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGDESCElement) Z_REQ_SELECTORRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGDESCElement) Z_SWAP(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfZ_SWAP(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGDESCElement) Z_SWAPRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGDESCElement) Z_SWAP_PUSH(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGDESCElement) Z_SWAP_PUSHRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGDESCElement) Z_TRIGGER(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfZ_TRIGGER(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGDESCElement) Z_TRIGGERRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGDESCElement) Z_REQ_METHOD(c SVGDescZReqMethodChoice) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGDescZReqMethodChoice string

const (
	// default GET
	SVGDescZReqMethod_empty SVGDescZReqMethodChoice = ""
	// GET
	SVGDescZReqMethod_get SVGDescZReqMethodChoice = "get"
	// POST
	SVGDescZReqMethod_post SVGDescZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGDESCElement) Z_REQ_METHODRemove(c SVGDescZReqMethodChoice) *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGDESCElement) Z_REQ_STRATEGY(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGDESCElement) Z_REQ_STRATEGYRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGDESCElement) Z_REQ_HISTORY(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGDESCElement) Z_REQ_HISTORYRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGDESCElement) Z_DATA(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfZ_DATA(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGDESCElement) Z_DATARemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGDESCElement) Z_JSON(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfZ_JSON(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGDESCElement) Z_JSONRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGDESCElement) Z_REQ_BATCH(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGDESCElement) Z_REQ_BATCHRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGDESCElement) Z_ACTION(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfZ_ACTION(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGDESCElement) Z_ACTIONRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGDESCElement) Z_REQ_BEFORE(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGDESCElement) Z_REQ_BEFORERemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGDESCElement) Z_REQ_AFTER(expression string) *SVGDESCElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGDESCElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGDESCElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGDESCElement) Z_REQ_AFTERRemove() *SVGDESCElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
