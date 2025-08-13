package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <rect> SVG element is a basic SVG shape that draws rectangles, defined by
// their position, width, and height
// The shape is created by connecting a line from one point to the other three
// points.
type SVGRECTElement struct {
	*Element
}

// Create a new SVGRECTElement element.
// This will create a new element with the tag
// "rect" during rendering.
func SVG_RECT(children ...ElementRenderer) *SVGRECTElement {
	e := NewElement("rect", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGRECTElement{Element: e}
}

func (e *SVGRECTElement) Children(children ...ElementRenderer) *SVGRECTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGRECTElement) IfChildren(condition bool, children ...ElementRenderer) *SVGRECTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGRECTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGRECTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGRECTElement) Attr(name string, value ...string) *SVGRECTElement {
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

func (e *SVGRECTElement) Attrs(attrs ...string) *SVGRECTElement {
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

func (e *SVGRECTElement) AttrsMap(attrs map[string]string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGRECTElement) Text(text string) *SVGRECTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGRECTElement) TextF(format string, args ...any) *SVGRECTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGRECTElement) IfText(condition bool, text string) *SVGRECTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGRECTElement) IfTextF(condition bool, format string, args ...any) *SVGRECTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGRECTElement) Escaped(text string) *SVGRECTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGRECTElement) IfEscaped(condition bool, text string) *SVGRECTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGRECTElement) EscapedF(format string, args ...any) *SVGRECTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGRECTElement) IfEscapedF(condition bool, format string, args ...any) *SVGRECTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGRECTElement) CustomData(key, value string) *SVGRECTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGRECTElement) IfCustomData(condition bool, key, value string) *SVGRECTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGRECTElement) CustomDataF(key, format string, args ...any) *SVGRECTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGRECTElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGRECTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGRECTElement) CustomDataRemove(key string) *SVGRECTElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The x-axis coordinate of the side of the rectangle which has the smaller x-axis
// value.
func (e *SVGRECTElement) X(f float64) *SVGRECTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("x", f)
	return e
}

func (e *SVGRECTElement) IfX(condition bool, f float64) *SVGRECTElement {
	if condition {
		e.X(f)
	}
	return e
}

// The y-axis coordinate of the side of the rectangle which has the smaller y-axis
// value.
func (e *SVGRECTElement) Y(f float64) *SVGRECTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("y", f)
	return e
}

func (e *SVGRECTElement) IfY(condition bool, f float64) *SVGRECTElement {
	if condition {
		e.Y(f)
	}
	return e
}

// The width of the rectangle.
func (e *SVGRECTElement) WIDTH(f float64) *SVGRECTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("width", f)
	return e
}

func (e *SVGRECTElement) IfWIDTH(condition bool, f float64) *SVGRECTElement {
	if condition {
		e.WIDTH(f)
	}
	return e
}

// The height of the rectangle.
func (e *SVGRECTElement) HEIGHT(f float64) *SVGRECTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("height", f)
	return e
}

func (e *SVGRECTElement) IfHEIGHT(condition bool, f float64) *SVGRECTElement {
	if condition {
		e.HEIGHT(f)
	}
	return e
}

// The x-axis radius of the ellipse used to round off the corners of the
// rectangle.
func (e *SVGRECTElement) RX(f float64) *SVGRECTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("rx", f)
	return e
}

func (e *SVGRECTElement) IfRX(condition bool, f float64) *SVGRECTElement {
	if condition {
		e.RX(f)
	}
	return e
}

// The y-axis radius of the ellipse used to round off the corners of the
// rectangle.
func (e *SVGRECTElement) RY(f float64) *SVGRECTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("ry", f)
	return e
}

func (e *SVGRECTElement) IfRY(condition bool, f float64) *SVGRECTElement {
	if condition {
		e.RY(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGRECTElement) ID(s string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGRECTElement) IDF(format string, args ...any) *SVGRECTElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGRECTElement) IfID(condition bool, s string) *SVGRECTElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGRECTElement) IfIDF(condition bool, format string, args ...any) *SVGRECTElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGRECTElement) IDRemove(s string) *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGRECTElement) IDRemoveF(format string, args ...any) *SVGRECTElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGRECTElement) CLASS(s ...string) *SVGRECTElement {
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

func (e *SVGRECTElement) IfCLASS(condition bool, s ...string) *SVGRECTElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGRECTElement) CLASSRemove(s ...string) *SVGRECTElement {
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
func (e *SVGRECTElement) STYLEF(k string, format string, args ...any) *SVGRECTElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGRECTElement) IfSTYLE(condition bool, k string, v string) *SVGRECTElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGRECTElement) STYLE(k string, v string) *SVGRECTElement {
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

func (e *SVGRECTElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGRECTElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGRECTElement) STYLEMap(m map[string]string) *SVGRECTElement {
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
func (e *SVGRECTElement) STYLEPairs(pairs ...string) *SVGRECTElement {
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

func (e *SVGRECTElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGRECTElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGRECTElement) STYLERemove(keys ...string) *SVGRECTElement {
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

func (e *SVGRECTElement) Z_REQ(expression string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRECTElement) IfZ_REQ(condition bool, expression string) *SVGRECTElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGRECTElement) Z_REQRemove() *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGRECTElement) Z_TARGET(expression string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRECTElement) IfZ_TARGET(condition bool, expression string) *SVGRECTElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGRECTElement) Z_TARGETRemove() *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGRECTElement) Z_REQ_SELECTOR(expression string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRECTElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGRECTElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGRECTElement) Z_REQ_SELECTORRemove() *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGRECTElement) Z_SWAP(expression string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRECTElement) IfZ_SWAP(condition bool, expression string) *SVGRECTElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGRECTElement) Z_SWAPRemove() *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGRECTElement) Z_SWAP_PUSH(expression string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRECTElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGRECTElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGRECTElement) Z_SWAP_PUSHRemove() *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGRECTElement) Z_TRIGGER(expression string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRECTElement) IfZ_TRIGGER(condition bool, expression string) *SVGRECTElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGRECTElement) Z_TRIGGERRemove() *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGRECTElement) Z_REQ_METHOD(c SVGRectZReqMethodChoice) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGRectZReqMethodChoice string

const (
	// default GET
	SVGRectZReqMethod_empty SVGRectZReqMethodChoice = ""
	// GET
	SVGRectZReqMethod_get SVGRectZReqMethodChoice = "get"
	// POST
	SVGRectZReqMethod_post SVGRectZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGRECTElement) Z_REQ_METHODRemove(c SVGRectZReqMethodChoice) *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGRECTElement) Z_REQ_STRATEGY(expression string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRECTElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGRECTElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGRECTElement) Z_REQ_STRATEGYRemove() *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGRECTElement) Z_REQ_HISTORY(expression string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRECTElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGRECTElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGRECTElement) Z_REQ_HISTORYRemove() *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGRECTElement) Z_DATA(expression string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRECTElement) IfZ_DATA(condition bool, expression string) *SVGRECTElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGRECTElement) Z_DATARemove() *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGRECTElement) Z_JSON(expression string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRECTElement) IfZ_JSON(condition bool, expression string) *SVGRECTElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGRECTElement) Z_JSONRemove() *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGRECTElement) Z_REQ_BATCH(expression string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRECTElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGRECTElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGRECTElement) Z_REQ_BATCHRemove() *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGRECTElement) Z_ACTION(expression string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRECTElement) IfZ_ACTION(condition bool, expression string) *SVGRECTElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGRECTElement) Z_ACTIONRemove() *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGRECTElement) Z_REQ_BEFORE(expression string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRECTElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGRECTElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGRECTElement) Z_REQ_BEFORERemove() *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGRECTElement) Z_REQ_AFTER(expression string) *SVGRECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRECTElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGRECTElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGRECTElement) Z_REQ_AFTERRemove() *SVGRECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
