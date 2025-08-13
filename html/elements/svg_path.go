package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <path> SVG element is the generic element to define a shape
// All the basic shapes can be created with a path element.
type SVGPATHElement struct {
	*Element
}

// Create a new SVGPATHElement element.
// This will create a new element with the tag
// "path" during rendering.
func SVG_PATH(children ...ElementRenderer) *SVGPATHElement {
	e := NewElement("path", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGPATHElement{Element: e}
}

func (e *SVGPATHElement) Children(children ...ElementRenderer) *SVGPATHElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGPATHElement) IfChildren(condition bool, children ...ElementRenderer) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGPATHElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGPATHElement) Attr(name string, value ...string) *SVGPATHElement {
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

func (e *SVGPATHElement) Attrs(attrs ...string) *SVGPATHElement {
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

func (e *SVGPATHElement) AttrsMap(attrs map[string]string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGPATHElement) Text(text string) *SVGPATHElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGPATHElement) TextF(format string, args ...any) *SVGPATHElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfText(condition bool, text string) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGPATHElement) IfTextF(condition bool, format string, args ...any) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGPATHElement) Escaped(text string) *SVGPATHElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGPATHElement) IfEscaped(condition bool, text string) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGPATHElement) EscapedF(format string, args ...any) *SVGPATHElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfEscapedF(condition bool, format string, args ...any) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGPATHElement) CustomData(key, value string) *SVGPATHElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGPATHElement) IfCustomData(condition bool, key, value string) *SVGPATHElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGPATHElement) CustomDataF(key, format string, args ...any) *SVGPATHElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGPATHElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGPATHElement) CustomDataRemove(key string) *SVGPATHElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The definition of the outline of a shape.
func (e *SVGPATHElement) D(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("d", s)
	return e
}

func (e *SVGPATHElement) DF(format string, args ...any) *SVGPATHElement {
	return e.D(fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfD(condition bool, s string) *SVGPATHElement {
	if condition {
		e.D(s)
	}
	return e
}

func (e *SVGPATHElement) IfDF(condition bool, format string, args ...any) *SVGPATHElement {
	if condition {
		e.D(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute D from the element.
func (e *SVGPATHElement) DRemove(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("d")
	return e
}

func (e *SVGPATHElement) DRemoveF(format string, args ...any) *SVGPATHElement {
	return e.DRemove(fmt.Sprintf(format, args...))
}

// The <path> SVG element is the generic element to define a shape
// All the basic shapes can be created with a path element.
func (e *SVGPATHElement) FILL(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("fill", s)
	return e
}

func (e *SVGPATHElement) FILLF(format string, args ...any) *SVGPATHElement {
	return e.FILL(fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfFILL(condition bool, s string) *SVGPATHElement {
	if condition {
		e.FILL(s)
	}
	return e
}

func (e *SVGPATHElement) IfFILLF(condition bool, format string, args ...any) *SVGPATHElement {
	if condition {
		e.FILL(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute FILL from the element.
func (e *SVGPATHElement) FILLRemove(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("fill")
	return e
}

func (e *SVGPATHElement) FILLRemoveF(format string, args ...any) *SVGPATHElement {
	return e.FILLRemove(fmt.Sprintf(format, args...))
}

// The <path> SVG element is the generic element to define a shape
// All the basic shapes can be created with a path element.
func (e *SVGPATHElement) FILL_OPACITY(f float64) *SVGPATHElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("fill-opacity", f)
	return e
}

func (e *SVGPATHElement) IfFILL_OPACITY(condition bool, f float64) *SVGPATHElement {
	if condition {
		e.FILL_OPACITY(f)
	}
	return e
}

// The total length for the path, in user units.
func (e *SVGPATHElement) PATH_LENGTH(f float64) *SVGPATHElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("pathLength", f)
	return e
}

func (e *SVGPATHElement) IfPATH_LENGTH(condition bool, f float64) *SVGPATHElement {
	if condition {
		e.PATH_LENGTH(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGPATHElement) ID(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGPATHElement) IDF(format string, args ...any) *SVGPATHElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfID(condition bool, s string) *SVGPATHElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGPATHElement) IfIDF(condition bool, format string, args ...any) *SVGPATHElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGPATHElement) IDRemove(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGPATHElement) IDRemoveF(format string, args ...any) *SVGPATHElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGPATHElement) CLASS(s ...string) *SVGPATHElement {
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

func (e *SVGPATHElement) IfCLASS(condition bool, s ...string) *SVGPATHElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGPATHElement) CLASSRemove(s ...string) *SVGPATHElement {
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
func (e *SVGPATHElement) STYLEF(k string, format string, args ...any) *SVGPATHElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfSTYLE(condition bool, k string, v string) *SVGPATHElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGPATHElement) STYLE(k string, v string) *SVGPATHElement {
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

func (e *SVGPATHElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGPATHElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGPATHElement) STYLEMap(m map[string]string) *SVGPATHElement {
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
func (e *SVGPATHElement) STYLEPairs(pairs ...string) *SVGPATHElement {
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

func (e *SVGPATHElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGPATHElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGPATHElement) STYLERemove(keys ...string) *SVGPATHElement {
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

func (e *SVGPATHElement) Z_REQ(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfZ_REQ(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGPATHElement) Z_REQRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGPATHElement) Z_TARGET(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfZ_TARGET(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGPATHElement) Z_TARGETRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGPATHElement) Z_REQ_SELECTOR(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGPATHElement) Z_REQ_SELECTORRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGPATHElement) Z_SWAP(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfZ_SWAP(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGPATHElement) Z_SWAPRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGPATHElement) Z_SWAP_PUSH(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGPATHElement) Z_SWAP_PUSHRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGPATHElement) Z_TRIGGER(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfZ_TRIGGER(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGPATHElement) Z_TRIGGERRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGPATHElement) Z_REQ_METHOD(c SVGPathZReqMethodChoice) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGPathZReqMethodChoice string

const (
	// default GET
	SVGPathZReqMethod_empty SVGPathZReqMethodChoice = ""
	// GET
	SVGPathZReqMethod_get SVGPathZReqMethodChoice = "get"
	// POST
	SVGPathZReqMethod_post SVGPathZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGPATHElement) Z_REQ_METHODRemove(c SVGPathZReqMethodChoice) *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGPATHElement) Z_REQ_STRATEGY(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGPATHElement) Z_REQ_STRATEGYRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGPATHElement) Z_REQ_HISTORY(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGPATHElement) Z_REQ_HISTORYRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGPATHElement) Z_DATA(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfZ_DATA(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGPATHElement) Z_DATARemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGPATHElement) Z_JSON(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfZ_JSON(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGPATHElement) Z_JSONRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGPATHElement) Z_REQ_BATCH(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGPATHElement) Z_REQ_BATCHRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGPATHElement) Z_ACTION(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfZ_ACTION(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGPATHElement) Z_ACTIONRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGPATHElement) Z_REQ_BEFORE(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGPATHElement) Z_REQ_BEFORERemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGPATHElement) Z_REQ_AFTER(expression string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATHElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGPATHElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGPATHElement) Z_REQ_AFTERRemove() *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
