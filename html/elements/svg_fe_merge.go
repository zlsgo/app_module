package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feMerge> SVG element allows filter effects to be applied concurrently
// instead of sequentially.
type SVGFEMERGEElement struct {
	*Element
}

// Create a new SVGFEMERGEElement element.
// This will create a new element with the tag
// "feMerge" during rendering.
func SVG_FEMERGE(children ...ElementRenderer) *SVGFEMERGEElement {
	e := NewElement("feMerge", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEMERGEElement{Element: e}
}

func (e *SVGFEMERGEElement) Children(children ...ElementRenderer) *SVGFEMERGEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEMERGEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEMERGEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEMERGEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEMERGEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEMERGEElement) Attr(name string, value ...string) *SVGFEMERGEElement {
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

func (e *SVGFEMERGEElement) Attrs(attrs ...string) *SVGFEMERGEElement {
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

func (e *SVGFEMERGEElement) AttrsMap(attrs map[string]string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEMERGEElement) Text(text string) *SVGFEMERGEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEMERGEElement) TextF(format string, args ...any) *SVGFEMERGEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGEElement) IfText(condition bool, text string) *SVGFEMERGEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEMERGEElement) IfTextF(condition bool, format string, args ...any) *SVGFEMERGEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEMERGEElement) Escaped(text string) *SVGFEMERGEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEMERGEElement) IfEscaped(condition bool, text string) *SVGFEMERGEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEMERGEElement) EscapedF(format string, args ...any) *SVGFEMERGEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGEElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEMERGEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEMERGEElement) CustomData(key, value string) *SVGFEMERGEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEMERGEElement) IfCustomData(condition bool, key, value string) *SVGFEMERGEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEMERGEElement) CustomDataF(key, format string, args ...any) *SVGFEMERGEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEMERGEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEMERGEElement) CustomDataRemove(key string) *SVGFEMERGEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// Specifies a unique id for an element
func (e *SVGFEMERGEElement) ID(s string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEMERGEElement) IDF(format string, args ...any) *SVGFEMERGEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGEElement) IfID(condition bool, s string) *SVGFEMERGEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEMERGEElement) IfIDF(condition bool, format string, args ...any) *SVGFEMERGEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEMERGEElement) IDRemove(s string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEMERGEElement) IDRemoveF(format string, args ...any) *SVGFEMERGEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEMERGEElement) CLASS(s ...string) *SVGFEMERGEElement {
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

func (e *SVGFEMERGEElement) IfCLASS(condition bool, s ...string) *SVGFEMERGEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEMERGEElement) CLASSRemove(s ...string) *SVGFEMERGEElement {
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
func (e *SVGFEMERGEElement) STYLEF(k string, format string, args ...any) *SVGFEMERGEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGEElement) IfSTYLE(condition bool, k string, v string) *SVGFEMERGEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEMERGEElement) STYLE(k string, v string) *SVGFEMERGEElement {
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

func (e *SVGFEMERGEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEMERGEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEMERGEElement) STYLEMap(m map[string]string) *SVGFEMERGEElement {
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
func (e *SVGFEMERGEElement) STYLEPairs(pairs ...string) *SVGFEMERGEElement {
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

func (e *SVGFEMERGEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEMERGEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEMERGEElement) STYLERemove(keys ...string) *SVGFEMERGEElement {
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

func (e *SVGFEMERGEElement) Z_REQ(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfZ_REQ(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEMERGEElement) Z_REQRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEMERGEElement) Z_TARGET(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfZ_TARGET(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEMERGEElement) Z_TARGETRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEMERGEElement) Z_REQ_SELECTOR(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEMERGEElement) Z_REQ_SELECTORRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEMERGEElement) Z_SWAP(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfZ_SWAP(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEMERGEElement) Z_SWAPRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEMERGEElement) Z_SWAP_PUSH(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEMERGEElement) Z_SWAP_PUSHRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEMERGEElement) Z_TRIGGER(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEMERGEElement) Z_TRIGGERRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEMERGEElement) Z_REQ_METHOD(c SVGFeMergeZReqMethodChoice) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeMergeZReqMethodChoice string

const (
	// default GET
	SVGFeMergeZReqMethod_empty SVGFeMergeZReqMethodChoice = ""
	// GET
	SVGFeMergeZReqMethod_get SVGFeMergeZReqMethodChoice = "get"
	// POST
	SVGFeMergeZReqMethod_post SVGFeMergeZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEMERGEElement) Z_REQ_METHODRemove(c SVGFeMergeZReqMethodChoice) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEMERGEElement) Z_REQ_STRATEGY(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEMERGEElement) Z_REQ_STRATEGYRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEMERGEElement) Z_REQ_HISTORY(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEMERGEElement) Z_REQ_HISTORYRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEMERGEElement) Z_DATA(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfZ_DATA(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEMERGEElement) Z_DATARemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEMERGEElement) Z_JSON(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfZ_JSON(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEMERGEElement) Z_JSONRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEMERGEElement) Z_REQ_BATCH(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEMERGEElement) Z_REQ_BATCHRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEMERGEElement) Z_ACTION(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfZ_ACTION(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEMERGEElement) Z_ACTIONRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEMERGEElement) Z_REQ_BEFORE(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEMERGEElement) Z_REQ_BEFORERemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEMERGEElement) Z_REQ_AFTER(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEMERGEElement) Z_REQ_AFTERRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
