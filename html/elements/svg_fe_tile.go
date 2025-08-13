package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feTile> SVG filter primitive allows to fill a target rectangle with a
// repeated, tiled pattern of an input image
// The effect is similar to the one of a <pattern> element, but <feTile> can use
// complex (i.e., filter) tree as input, and can be animated.
type SVGFETILEElement struct {
	*Element
}

// Create a new SVGFETILEElement element.
// This will create a new element with the tag
// "feTile" during rendering.
func SVG_FETILE(children ...ElementRenderer) *SVGFETILEElement {
	e := NewElement("feTile", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFETILEElement{Element: e}
}

func (e *SVGFETILEElement) Children(children ...ElementRenderer) *SVGFETILEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFETILEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFETILEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFETILEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFETILEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFETILEElement) Attr(name string, value ...string) *SVGFETILEElement {
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

func (e *SVGFETILEElement) Attrs(attrs ...string) *SVGFETILEElement {
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

func (e *SVGFETILEElement) AttrsMap(attrs map[string]string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFETILEElement) Text(text string) *SVGFETILEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFETILEElement) TextF(format string, args ...any) *SVGFETILEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFETILEElement) IfText(condition bool, text string) *SVGFETILEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFETILEElement) IfTextF(condition bool, format string, args ...any) *SVGFETILEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFETILEElement) Escaped(text string) *SVGFETILEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFETILEElement) IfEscaped(condition bool, text string) *SVGFETILEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFETILEElement) EscapedF(format string, args ...any) *SVGFETILEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFETILEElement) IfEscapedF(condition bool, format string, args ...any) *SVGFETILEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFETILEElement) CustomData(key, value string) *SVGFETILEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFETILEElement) IfCustomData(condition bool, key, value string) *SVGFETILEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFETILEElement) CustomDataF(key, format string, args ...any) *SVGFETILEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFETILEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFETILEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFETILEElement) CustomDataRemove(key string) *SVGFETILEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The input for this filter.
func (e *SVGFETILEElement) IN(s string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFETILEElement) INF(format string, args ...any) *SVGFETILEElement {
	return e.IN(fmt.Sprintf(format, args...))
}

func (e *SVGFETILEElement) IfIN(condition bool, s string) *SVGFETILEElement {
	if condition {
		e.IN(s)
	}
	return e
}

func (e *SVGFETILEElement) IfINF(condition bool, format string, args ...any) *SVGFETILEElement {
	if condition {
		e.IN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFETILEElement) INRemove(s string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("in")
	return e
}

func (e *SVGFETILEElement) INRemoveF(format string, args ...any) *SVGFETILEElement {
	return e.INRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGFETILEElement) ID(s string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFETILEElement) IDF(format string, args ...any) *SVGFETILEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFETILEElement) IfID(condition bool, s string) *SVGFETILEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFETILEElement) IfIDF(condition bool, format string, args ...any) *SVGFETILEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFETILEElement) IDRemove(s string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFETILEElement) IDRemoveF(format string, args ...any) *SVGFETILEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFETILEElement) CLASS(s ...string) *SVGFETILEElement {
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

func (e *SVGFETILEElement) IfCLASS(condition bool, s ...string) *SVGFETILEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFETILEElement) CLASSRemove(s ...string) *SVGFETILEElement {
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
func (e *SVGFETILEElement) STYLEF(k string, format string, args ...any) *SVGFETILEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFETILEElement) IfSTYLE(condition bool, k string, v string) *SVGFETILEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFETILEElement) STYLE(k string, v string) *SVGFETILEElement {
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

func (e *SVGFETILEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFETILEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFETILEElement) STYLEMap(m map[string]string) *SVGFETILEElement {
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
func (e *SVGFETILEElement) STYLEPairs(pairs ...string) *SVGFETILEElement {
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

func (e *SVGFETILEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFETILEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFETILEElement) STYLERemove(keys ...string) *SVGFETILEElement {
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

func (e *SVGFETILEElement) Z_REQ(expression string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETILEElement) IfZ_REQ(condition bool, expression string) *SVGFETILEElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFETILEElement) Z_REQRemove() *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFETILEElement) Z_TARGET(expression string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETILEElement) IfZ_TARGET(condition bool, expression string) *SVGFETILEElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFETILEElement) Z_TARGETRemove() *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFETILEElement) Z_REQ_SELECTOR(expression string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETILEElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFETILEElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFETILEElement) Z_REQ_SELECTORRemove() *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFETILEElement) Z_SWAP(expression string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETILEElement) IfZ_SWAP(condition bool, expression string) *SVGFETILEElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFETILEElement) Z_SWAPRemove() *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFETILEElement) Z_SWAP_PUSH(expression string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETILEElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFETILEElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFETILEElement) Z_SWAP_PUSHRemove() *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFETILEElement) Z_TRIGGER(expression string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETILEElement) IfZ_TRIGGER(condition bool, expression string) *SVGFETILEElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFETILEElement) Z_TRIGGERRemove() *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFETILEElement) Z_REQ_METHOD(c SVGFeTileZReqMethodChoice) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeTileZReqMethodChoice string

const (
	// default GET
	SVGFeTileZReqMethod_empty SVGFeTileZReqMethodChoice = ""
	// GET
	SVGFeTileZReqMethod_get SVGFeTileZReqMethodChoice = "get"
	// POST
	SVGFeTileZReqMethod_post SVGFeTileZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFETILEElement) Z_REQ_METHODRemove(c SVGFeTileZReqMethodChoice) *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFETILEElement) Z_REQ_STRATEGY(expression string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETILEElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFETILEElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFETILEElement) Z_REQ_STRATEGYRemove() *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFETILEElement) Z_REQ_HISTORY(expression string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETILEElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFETILEElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFETILEElement) Z_REQ_HISTORYRemove() *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFETILEElement) Z_DATA(expression string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETILEElement) IfZ_DATA(condition bool, expression string) *SVGFETILEElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFETILEElement) Z_DATARemove() *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFETILEElement) Z_JSON(expression string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETILEElement) IfZ_JSON(condition bool, expression string) *SVGFETILEElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFETILEElement) Z_JSONRemove() *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFETILEElement) Z_REQ_BATCH(expression string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETILEElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFETILEElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFETILEElement) Z_REQ_BATCHRemove() *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFETILEElement) Z_ACTION(expression string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETILEElement) IfZ_ACTION(condition bool, expression string) *SVGFETILEElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFETILEElement) Z_ACTIONRemove() *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFETILEElement) Z_REQ_BEFORE(expression string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETILEElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFETILEElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFETILEElement) Z_REQ_BEFORERemove() *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFETILEElement) Z_REQ_AFTER(expression string) *SVGFETILEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETILEElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFETILEElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFETILEElement) Z_REQ_AFTERRemove() *SVGFETILEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
