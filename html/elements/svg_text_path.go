package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <textPath> SVG element defines a set of glyphs that exactly fit along a
// curve.
type SVGTEXTPATHElement struct {
	*Element
}

// Create a new SVGTEXTPATHElement element.
// This will create a new element with the tag
// "textPath" during rendering.
func SVG_TEXTPATH(children ...ElementRenderer) *SVGTEXTPATHElement {
	e := NewElement("textPath", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGTEXTPATHElement{Element: e}
}

func (e *SVGTEXTPATHElement) Children(children ...ElementRenderer) *SVGTEXTPATHElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGTEXTPATHElement) IfChildren(condition bool, children ...ElementRenderer) *SVGTEXTPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGTEXTPATHElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGTEXTPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGTEXTPATHElement) Attr(name string, value ...string) *SVGTEXTPATHElement {
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

func (e *SVGTEXTPATHElement) Attrs(attrs ...string) *SVGTEXTPATHElement {
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

func (e *SVGTEXTPATHElement) AttrsMap(attrs map[string]string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGTEXTPATHElement) Text(text string) *SVGTEXTPATHElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGTEXTPATHElement) TextF(format string, args ...any) *SVGTEXTPATHElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGTEXTPATHElement) IfText(condition bool, text string) *SVGTEXTPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGTEXTPATHElement) IfTextF(condition bool, format string, args ...any) *SVGTEXTPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGTEXTPATHElement) Escaped(text string) *SVGTEXTPATHElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGTEXTPATHElement) IfEscaped(condition bool, text string) *SVGTEXTPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGTEXTPATHElement) EscapedF(format string, args ...any) *SVGTEXTPATHElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGTEXTPATHElement) IfEscapedF(condition bool, format string, args ...any) *SVGTEXTPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGTEXTPATHElement) CustomData(key, value string) *SVGTEXTPATHElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGTEXTPATHElement) IfCustomData(condition bool, key, value string) *SVGTEXTPATHElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGTEXTPATHElement) CustomDataF(key, format string, args ...any) *SVGTEXTPATHElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGTEXTPATHElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGTEXTPATHElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGTEXTPATHElement) CustomDataRemove(key string) *SVGTEXTPATHElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// A URI reference to the path to render along.
func (e *SVGTEXTPATHElement) HREF(s string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("href", s)
	return e
}

func (e *SVGTEXTPATHElement) HREFF(format string, args ...any) *SVGTEXTPATHElement {
	return e.HREF(fmt.Sprintf(format, args...))
}

func (e *SVGTEXTPATHElement) IfHREF(condition bool, s string) *SVGTEXTPATHElement {
	if condition {
		e.HREF(s)
	}
	return e
}

func (e *SVGTEXTPATHElement) IfHREFF(condition bool, format string, args ...any) *SVGTEXTPATHElement {
	if condition {
		e.HREF(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HREF from the element.
func (e *SVGTEXTPATHElement) HREFRemove(s string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("href")
	return e
}

func (e *SVGTEXTPATHElement) HREFRemoveF(format string, args ...any) *SVGTEXTPATHElement {
	return e.HREFRemove(fmt.Sprintf(format, args...))
}

// Indicates an offset from the start of the path, where the first character is
// rendered.
func (e *SVGTEXTPATHElement) START_OFFSET(s string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("startOffset", s)
	return e
}

func (e *SVGTEXTPATHElement) START_OFFSETF(format string, args ...any) *SVGTEXTPATHElement {
	return e.START_OFFSET(fmt.Sprintf(format, args...))
}

func (e *SVGTEXTPATHElement) IfSTART_OFFSET(condition bool, s string) *SVGTEXTPATHElement {
	if condition {
		e.START_OFFSET(s)
	}
	return e
}

func (e *SVGTEXTPATHElement) IfSTART_OFFSETF(condition bool, format string, args ...any) *SVGTEXTPATHElement {
	if condition {
		e.START_OFFSET(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute START_OFFSET from the element.
func (e *SVGTEXTPATHElement) START_OFFSETRemove(s string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("startOffset")
	return e
}

func (e *SVGTEXTPATHElement) START_OFFSETRemoveF(format string, args ...any) *SVGTEXTPATHElement {
	return e.START_OFFSETRemove(fmt.Sprintf(format, args...))
}

// Indicates the method by which text should be rendered along the path.
func (e *SVGTEXTPATHElement) METHOD(c SVGTextPathMethodChoice) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("method", string(c))
	return e
}

type SVGTextPathMethodChoice string

const (
	// Indicates the method by which text should be rendered along the path.
	SVGTextPathMethod_align SVGTextPathMethodChoice = "align"
	// Indicates the method by which text should be rendered along the path.
	SVGTextPathMethod_stretch SVGTextPathMethodChoice = "stretch"
)

// Remove the attribute METHOD from the element.
func (e *SVGTEXTPATHElement) METHODRemove(c SVGTextPathMethodChoice) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("method")
	return e
}

// Indicates the spacing behavior between characters.
func (e *SVGTEXTPATHElement) SPACING(c SVGTextPathSpacingChoice) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("spacing", string(c))
	return e
}

type SVGTextPathSpacingChoice string

const (
	// Indicates the spacing behavior between characters.
	SVGTextPathSpacing_auto SVGTextPathSpacingChoice = "auto"
	// Indicates the spacing behavior between characters.
	SVGTextPathSpacing_exact SVGTextPathSpacingChoice = "exact"
)

// Remove the attribute SPACING from the element.
func (e *SVGTEXTPATHElement) SPACINGRemove(c SVGTextPathSpacingChoice) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("spacing")
	return e
}

// Specifies a unique id for an element
func (e *SVGTEXTPATHElement) ID(s string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGTEXTPATHElement) IDF(format string, args ...any) *SVGTEXTPATHElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGTEXTPATHElement) IfID(condition bool, s string) *SVGTEXTPATHElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGTEXTPATHElement) IfIDF(condition bool, format string, args ...any) *SVGTEXTPATHElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGTEXTPATHElement) IDRemove(s string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGTEXTPATHElement) IDRemoveF(format string, args ...any) *SVGTEXTPATHElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGTEXTPATHElement) CLASS(s ...string) *SVGTEXTPATHElement {
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

func (e *SVGTEXTPATHElement) IfCLASS(condition bool, s ...string) *SVGTEXTPATHElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGTEXTPATHElement) CLASSRemove(s ...string) *SVGTEXTPATHElement {
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
func (e *SVGTEXTPATHElement) STYLEF(k string, format string, args ...any) *SVGTEXTPATHElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGTEXTPATHElement) IfSTYLE(condition bool, k string, v string) *SVGTEXTPATHElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGTEXTPATHElement) STYLE(k string, v string) *SVGTEXTPATHElement {
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

func (e *SVGTEXTPATHElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGTEXTPATHElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGTEXTPATHElement) STYLEMap(m map[string]string) *SVGTEXTPATHElement {
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
func (e *SVGTEXTPATHElement) STYLEPairs(pairs ...string) *SVGTEXTPATHElement {
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

func (e *SVGTEXTPATHElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGTEXTPATHElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGTEXTPATHElement) STYLERemove(keys ...string) *SVGTEXTPATHElement {
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

func (e *SVGTEXTPATHElement) Z_REQ(expression string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTPATHElement) IfZ_REQ(condition bool, expression string) *SVGTEXTPATHElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGTEXTPATHElement) Z_REQRemove() *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGTEXTPATHElement) Z_TARGET(expression string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTPATHElement) IfZ_TARGET(condition bool, expression string) *SVGTEXTPATHElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGTEXTPATHElement) Z_TARGETRemove() *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGTEXTPATHElement) Z_REQ_SELECTOR(expression string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTPATHElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGTEXTPATHElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGTEXTPATHElement) Z_REQ_SELECTORRemove() *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGTEXTPATHElement) Z_SWAP(expression string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTPATHElement) IfZ_SWAP(condition bool, expression string) *SVGTEXTPATHElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGTEXTPATHElement) Z_SWAPRemove() *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGTEXTPATHElement) Z_SWAP_PUSH(expression string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTPATHElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGTEXTPATHElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGTEXTPATHElement) Z_SWAP_PUSHRemove() *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGTEXTPATHElement) Z_TRIGGER(expression string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTPATHElement) IfZ_TRIGGER(condition bool, expression string) *SVGTEXTPATHElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGTEXTPATHElement) Z_TRIGGERRemove() *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGTEXTPATHElement) Z_REQ_METHOD(c SVGTextPathZReqMethodChoice) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGTextPathZReqMethodChoice string

const (
	// default GET
	SVGTextPathZReqMethod_empty SVGTextPathZReqMethodChoice = ""
	// GET
	SVGTextPathZReqMethod_get SVGTextPathZReqMethodChoice = "get"
	// POST
	SVGTextPathZReqMethod_post SVGTextPathZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGTEXTPATHElement) Z_REQ_METHODRemove(c SVGTextPathZReqMethodChoice) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGTEXTPATHElement) Z_REQ_STRATEGY(expression string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTPATHElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGTEXTPATHElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGTEXTPATHElement) Z_REQ_STRATEGYRemove() *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGTEXTPATHElement) Z_REQ_HISTORY(expression string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTPATHElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGTEXTPATHElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGTEXTPATHElement) Z_REQ_HISTORYRemove() *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGTEXTPATHElement) Z_DATA(expression string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTPATHElement) IfZ_DATA(condition bool, expression string) *SVGTEXTPATHElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGTEXTPATHElement) Z_DATARemove() *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGTEXTPATHElement) Z_JSON(expression string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTPATHElement) IfZ_JSON(condition bool, expression string) *SVGTEXTPATHElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGTEXTPATHElement) Z_JSONRemove() *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGTEXTPATHElement) Z_REQ_BATCH(expression string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTPATHElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGTEXTPATHElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGTEXTPATHElement) Z_REQ_BATCHRemove() *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGTEXTPATHElement) Z_ACTION(expression string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTPATHElement) IfZ_ACTION(condition bool, expression string) *SVGTEXTPATHElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGTEXTPATHElement) Z_ACTIONRemove() *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGTEXTPATHElement) Z_REQ_BEFORE(expression string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTPATHElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGTEXTPATHElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGTEXTPATHElement) Z_REQ_BEFORERemove() *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGTEXTPATHElement) Z_REQ_AFTER(expression string) *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTEXTPATHElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGTEXTPATHElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGTEXTPATHElement) Z_REQ_AFTERRemove() *SVGTEXTPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
