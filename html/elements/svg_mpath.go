package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <mpath> SVG element allows to use the functionality of <animateMotion> to
// animate the <startOffset> attribute of SVG <textPath> elements.
type SVGMPATHElement struct {
	*Element
}

// Create a new SVGMPATHElement element.
// This will create a new element with the tag
// "mpath" during rendering.
func SVG_MPATH(children ...ElementRenderer) *SVGMPATHElement {
	e := NewElement("mpath", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGMPATHElement{Element: e}
}

func (e *SVGMPATHElement) Children(children ...ElementRenderer) *SVGMPATHElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGMPATHElement) IfChildren(condition bool, children ...ElementRenderer) *SVGMPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGMPATHElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGMPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGMPATHElement) Attr(name string, value ...string) *SVGMPATHElement {
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

func (e *SVGMPATHElement) Attrs(attrs ...string) *SVGMPATHElement {
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

func (e *SVGMPATHElement) AttrsMap(attrs map[string]string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGMPATHElement) Text(text string) *SVGMPATHElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGMPATHElement) TextF(format string, args ...any) *SVGMPATHElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGMPATHElement) IfText(condition bool, text string) *SVGMPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGMPATHElement) IfTextF(condition bool, format string, args ...any) *SVGMPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGMPATHElement) Escaped(text string) *SVGMPATHElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGMPATHElement) IfEscaped(condition bool, text string) *SVGMPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGMPATHElement) EscapedF(format string, args ...any) *SVGMPATHElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGMPATHElement) IfEscapedF(condition bool, format string, args ...any) *SVGMPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGMPATHElement) CustomData(key, value string) *SVGMPATHElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGMPATHElement) IfCustomData(condition bool, key, value string) *SVGMPATHElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGMPATHElement) CustomDataF(key, format string, args ...any) *SVGMPATHElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGMPATHElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGMPATHElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGMPATHElement) CustomDataRemove(key string) *SVGMPATHElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// A URI reference to the motion path definition.
func (e *SVGMPATHElement) HREF(s string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("href", s)
	return e
}

func (e *SVGMPATHElement) HREFF(format string, args ...any) *SVGMPATHElement {
	return e.HREF(fmt.Sprintf(format, args...))
}

func (e *SVGMPATHElement) IfHREF(condition bool, s string) *SVGMPATHElement {
	if condition {
		e.HREF(s)
	}
	return e
}

func (e *SVGMPATHElement) IfHREFF(condition bool, format string, args ...any) *SVGMPATHElement {
	if condition {
		e.HREF(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HREF from the element.
func (e *SVGMPATHElement) HREFRemove(s string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("href")
	return e
}

func (e *SVGMPATHElement) HREFRemoveF(format string, args ...any) *SVGMPATHElement {
	return e.HREFRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGMPATHElement) ID(s string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGMPATHElement) IDF(format string, args ...any) *SVGMPATHElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGMPATHElement) IfID(condition bool, s string) *SVGMPATHElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGMPATHElement) IfIDF(condition bool, format string, args ...any) *SVGMPATHElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGMPATHElement) IDRemove(s string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGMPATHElement) IDRemoveF(format string, args ...any) *SVGMPATHElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGMPATHElement) CLASS(s ...string) *SVGMPATHElement {
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

func (e *SVGMPATHElement) IfCLASS(condition bool, s ...string) *SVGMPATHElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGMPATHElement) CLASSRemove(s ...string) *SVGMPATHElement {
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
func (e *SVGMPATHElement) STYLEF(k string, format string, args ...any) *SVGMPATHElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGMPATHElement) IfSTYLE(condition bool, k string, v string) *SVGMPATHElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGMPATHElement) STYLE(k string, v string) *SVGMPATHElement {
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

func (e *SVGMPATHElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGMPATHElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGMPATHElement) STYLEMap(m map[string]string) *SVGMPATHElement {
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
func (e *SVGMPATHElement) STYLEPairs(pairs ...string) *SVGMPATHElement {
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

func (e *SVGMPATHElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGMPATHElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGMPATHElement) STYLERemove(keys ...string) *SVGMPATHElement {
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

func (e *SVGMPATHElement) Z_REQ(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfZ_REQ(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGMPATHElement) Z_REQRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGMPATHElement) Z_TARGET(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfZ_TARGET(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGMPATHElement) Z_TARGETRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGMPATHElement) Z_REQ_SELECTOR(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGMPATHElement) Z_REQ_SELECTORRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGMPATHElement) Z_SWAP(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfZ_SWAP(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGMPATHElement) Z_SWAPRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGMPATHElement) Z_SWAP_PUSH(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGMPATHElement) Z_SWAP_PUSHRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGMPATHElement) Z_TRIGGER(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfZ_TRIGGER(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGMPATHElement) Z_TRIGGERRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGMPATHElement) Z_REQ_METHOD(c SVGMpathZReqMethodChoice) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGMpathZReqMethodChoice string

const (
	// default GET
	SVGMpathZReqMethod_empty SVGMpathZReqMethodChoice = ""
	// GET
	SVGMpathZReqMethod_get SVGMpathZReqMethodChoice = "get"
	// POST
	SVGMpathZReqMethod_post SVGMpathZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGMPATHElement) Z_REQ_METHODRemove(c SVGMpathZReqMethodChoice) *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGMPATHElement) Z_REQ_STRATEGY(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGMPATHElement) Z_REQ_STRATEGYRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGMPATHElement) Z_REQ_HISTORY(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGMPATHElement) Z_REQ_HISTORYRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGMPATHElement) Z_DATA(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfZ_DATA(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGMPATHElement) Z_DATARemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGMPATHElement) Z_JSON(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfZ_JSON(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGMPATHElement) Z_JSONRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGMPATHElement) Z_REQ_BATCH(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGMPATHElement) Z_REQ_BATCHRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGMPATHElement) Z_ACTION(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfZ_ACTION(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGMPATHElement) Z_ACTIONRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGMPATHElement) Z_REQ_BEFORE(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGMPATHElement) Z_REQ_BEFORERemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGMPATHElement) Z_REQ_AFTER(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGMPATHElement) Z_REQ_AFTERRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
