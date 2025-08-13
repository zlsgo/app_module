package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <line> SVG element is an SVG basic shape, used to create a line connecting
// two points.
type SVGLINEElement struct {
	*Element
}

// Create a new SVGLINEElement element.
// This will create a new element with the tag
// "line" during rendering.
func SVG_LINE(children ...ElementRenderer) *SVGLINEElement {
	e := NewElement("line", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGLINEElement{Element: e}
}

func (e *SVGLINEElement) Children(children ...ElementRenderer) *SVGLINEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGLINEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGLINEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGLINEElement) Attr(name string, value ...string) *SVGLINEElement {
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

func (e *SVGLINEElement) Attrs(attrs ...string) *SVGLINEElement {
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

func (e *SVGLINEElement) AttrsMap(attrs map[string]string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGLINEElement) Text(text string) *SVGLINEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGLINEElement) TextF(format string, args ...any) *SVGLINEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGLINEElement) IfText(condition bool, text string) *SVGLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGLINEElement) IfTextF(condition bool, format string, args ...any) *SVGLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGLINEElement) Escaped(text string) *SVGLINEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGLINEElement) IfEscaped(condition bool, text string) *SVGLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGLINEElement) EscapedF(format string, args ...any) *SVGLINEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGLINEElement) IfEscapedF(condition bool, format string, args ...any) *SVGLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGLINEElement) CustomData(key, value string) *SVGLINEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGLINEElement) IfCustomData(condition bool, key, value string) *SVGLINEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGLINEElement) CustomDataF(key, format string, args ...any) *SVGLINEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGLINEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGLINEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGLINEElement) CustomDataRemove(key string) *SVGLINEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The x-axis coordinate of the starting point of the line.
func (e *SVGLINEElement) X_1(f float64) *SVGLINEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("x1", f)
	return e
}

func (e *SVGLINEElement) IfX_1(condition bool, f float64) *SVGLINEElement {
	if condition {
		e.X_1(f)
	}
	return e
}

// The y-axis coordinate of the starting point of the line.
func (e *SVGLINEElement) Y_1(f float64) *SVGLINEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("y1", f)
	return e
}

func (e *SVGLINEElement) IfY_1(condition bool, f float64) *SVGLINEElement {
	if condition {
		e.Y_1(f)
	}
	return e
}

// The x-axis coordinate of the ending point of the line.
func (e *SVGLINEElement) X_2(f float64) *SVGLINEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("x2", f)
	return e
}

func (e *SVGLINEElement) IfX_2(condition bool, f float64) *SVGLINEElement {
	if condition {
		e.X_2(f)
	}
	return e
}

// The y-axis coordinate of the ending point of the line.
func (e *SVGLINEElement) Y_2(f float64) *SVGLINEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("y2", f)
	return e
}

func (e *SVGLINEElement) IfY_2(condition bool, f float64) *SVGLINEElement {
	if condition {
		e.Y_2(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGLINEElement) ID(s string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGLINEElement) IDF(format string, args ...any) *SVGLINEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGLINEElement) IfID(condition bool, s string) *SVGLINEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGLINEElement) IfIDF(condition bool, format string, args ...any) *SVGLINEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGLINEElement) IDRemove(s string) *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGLINEElement) IDRemoveF(format string, args ...any) *SVGLINEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGLINEElement) CLASS(s ...string) *SVGLINEElement {
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

func (e *SVGLINEElement) IfCLASS(condition bool, s ...string) *SVGLINEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGLINEElement) CLASSRemove(s ...string) *SVGLINEElement {
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
func (e *SVGLINEElement) STYLEF(k string, format string, args ...any) *SVGLINEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGLINEElement) IfSTYLE(condition bool, k string, v string) *SVGLINEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGLINEElement) STYLE(k string, v string) *SVGLINEElement {
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

func (e *SVGLINEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGLINEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGLINEElement) STYLEMap(m map[string]string) *SVGLINEElement {
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
func (e *SVGLINEElement) STYLEPairs(pairs ...string) *SVGLINEElement {
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

func (e *SVGLINEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGLINEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGLINEElement) STYLERemove(keys ...string) *SVGLINEElement {
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

func (e *SVGLINEElement) Z_REQ(expression string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEElement) IfZ_REQ(condition bool, expression string) *SVGLINEElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGLINEElement) Z_REQRemove() *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGLINEElement) Z_TARGET(expression string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEElement) IfZ_TARGET(condition bool, expression string) *SVGLINEElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGLINEElement) Z_TARGETRemove() *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGLINEElement) Z_REQ_SELECTOR(expression string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGLINEElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGLINEElement) Z_REQ_SELECTORRemove() *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGLINEElement) Z_SWAP(expression string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEElement) IfZ_SWAP(condition bool, expression string) *SVGLINEElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGLINEElement) Z_SWAPRemove() *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGLINEElement) Z_SWAP_PUSH(expression string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGLINEElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGLINEElement) Z_SWAP_PUSHRemove() *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGLINEElement) Z_TRIGGER(expression string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEElement) IfZ_TRIGGER(condition bool, expression string) *SVGLINEElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGLINEElement) Z_TRIGGERRemove() *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGLINEElement) Z_REQ_METHOD(c SVGLineZReqMethodChoice) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGLineZReqMethodChoice string

const (
	// default GET
	SVGLineZReqMethod_empty SVGLineZReqMethodChoice = ""
	// GET
	SVGLineZReqMethod_get SVGLineZReqMethodChoice = "get"
	// POST
	SVGLineZReqMethod_post SVGLineZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGLINEElement) Z_REQ_METHODRemove(c SVGLineZReqMethodChoice) *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGLINEElement) Z_REQ_STRATEGY(expression string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGLINEElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGLINEElement) Z_REQ_STRATEGYRemove() *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGLINEElement) Z_REQ_HISTORY(expression string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGLINEElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGLINEElement) Z_REQ_HISTORYRemove() *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGLINEElement) Z_DATA(expression string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEElement) IfZ_DATA(condition bool, expression string) *SVGLINEElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGLINEElement) Z_DATARemove() *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGLINEElement) Z_JSON(expression string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEElement) IfZ_JSON(condition bool, expression string) *SVGLINEElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGLINEElement) Z_JSONRemove() *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGLINEElement) Z_REQ_BATCH(expression string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGLINEElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGLINEElement) Z_REQ_BATCHRemove() *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGLINEElement) Z_ACTION(expression string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEElement) IfZ_ACTION(condition bool, expression string) *SVGLINEElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGLINEElement) Z_ACTIONRemove() *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGLINEElement) Z_REQ_BEFORE(expression string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGLINEElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGLINEElement) Z_REQ_BEFORERemove() *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGLINEElement) Z_REQ_AFTER(expression string) *SVGLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGLINEElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGLINEElement) Z_REQ_AFTERRemove() *SVGLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
