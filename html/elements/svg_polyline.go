package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <polyline> SVG element is an SVG basic shape, used to create a series of
// straight lines connecting several points
// Typically a polyline is used to create open shapes as the last point doesn't
// have to be connected to the first point
// For closed shapes see the <polygon> element.
type SVGPOLYLINEElement struct {
	*Element
}

// Create a new SVGPOLYLINEElement element.
// This will create a new element with the tag
// "polyline" during rendering.
func SVG_POLYLINE(children ...ElementRenderer) *SVGPOLYLINEElement {
	e := NewElement("polyline", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGPOLYLINEElement{Element: e}
}

func (e *SVGPOLYLINEElement) Children(children ...ElementRenderer) *SVGPOLYLINEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGPOLYLINEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGPOLYLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGPOLYLINEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGPOLYLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGPOLYLINEElement) Attr(name string, value ...string) *SVGPOLYLINEElement {
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

func (e *SVGPOLYLINEElement) Attrs(attrs ...string) *SVGPOLYLINEElement {
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

func (e *SVGPOLYLINEElement) AttrsMap(attrs map[string]string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGPOLYLINEElement) Text(text string) *SVGPOLYLINEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGPOLYLINEElement) TextF(format string, args ...any) *SVGPOLYLINEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGPOLYLINEElement) IfText(condition bool, text string) *SVGPOLYLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGPOLYLINEElement) IfTextF(condition bool, format string, args ...any) *SVGPOLYLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGPOLYLINEElement) Escaped(text string) *SVGPOLYLINEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGPOLYLINEElement) IfEscaped(condition bool, text string) *SVGPOLYLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGPOLYLINEElement) EscapedF(format string, args ...any) *SVGPOLYLINEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGPOLYLINEElement) IfEscapedF(condition bool, format string, args ...any) *SVGPOLYLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGPOLYLINEElement) CustomData(key, value string) *SVGPOLYLINEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGPOLYLINEElement) IfCustomData(condition bool, key, value string) *SVGPOLYLINEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGPOLYLINEElement) CustomDataF(key, format string, args ...any) *SVGPOLYLINEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGPOLYLINEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGPOLYLINEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGPOLYLINEElement) CustomDataRemove(key string) *SVGPOLYLINEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// A list of points, each of which is a coordinate pair.
func (e *SVGPOLYLINEElement) POINTS(s string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("points", s)
	return e
}

func (e *SVGPOLYLINEElement) POINTSF(format string, args ...any) *SVGPOLYLINEElement {
	return e.POINTS(fmt.Sprintf(format, args...))
}

func (e *SVGPOLYLINEElement) IfPOINTS(condition bool, s string) *SVGPOLYLINEElement {
	if condition {
		e.POINTS(s)
	}
	return e
}

func (e *SVGPOLYLINEElement) IfPOINTSF(condition bool, format string, args ...any) *SVGPOLYLINEElement {
	if condition {
		e.POINTS(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute POINTS from the element.
func (e *SVGPOLYLINEElement) POINTSRemove(s string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("points")
	return e
}

func (e *SVGPOLYLINEElement) POINTSRemoveF(format string, args ...any) *SVGPOLYLINEElement {
	return e.POINTSRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGPOLYLINEElement) ID(s string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGPOLYLINEElement) IDF(format string, args ...any) *SVGPOLYLINEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGPOLYLINEElement) IfID(condition bool, s string) *SVGPOLYLINEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGPOLYLINEElement) IfIDF(condition bool, format string, args ...any) *SVGPOLYLINEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGPOLYLINEElement) IDRemove(s string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGPOLYLINEElement) IDRemoveF(format string, args ...any) *SVGPOLYLINEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGPOLYLINEElement) CLASS(s ...string) *SVGPOLYLINEElement {
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

func (e *SVGPOLYLINEElement) IfCLASS(condition bool, s ...string) *SVGPOLYLINEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGPOLYLINEElement) CLASSRemove(s ...string) *SVGPOLYLINEElement {
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
func (e *SVGPOLYLINEElement) STYLEF(k string, format string, args ...any) *SVGPOLYLINEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGPOLYLINEElement) IfSTYLE(condition bool, k string, v string) *SVGPOLYLINEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGPOLYLINEElement) STYLE(k string, v string) *SVGPOLYLINEElement {
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

func (e *SVGPOLYLINEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGPOLYLINEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGPOLYLINEElement) STYLEMap(m map[string]string) *SVGPOLYLINEElement {
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
func (e *SVGPOLYLINEElement) STYLEPairs(pairs ...string) *SVGPOLYLINEElement {
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

func (e *SVGPOLYLINEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGPOLYLINEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGPOLYLINEElement) STYLERemove(keys ...string) *SVGPOLYLINEElement {
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

func (e *SVGPOLYLINEElement) Z_REQ(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfZ_REQ(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGPOLYLINEElement) Z_REQRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGPOLYLINEElement) Z_TARGET(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfZ_TARGET(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGPOLYLINEElement) Z_TARGETRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGPOLYLINEElement) Z_REQ_SELECTOR(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGPOLYLINEElement) Z_REQ_SELECTORRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGPOLYLINEElement) Z_SWAP(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfZ_SWAP(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGPOLYLINEElement) Z_SWAPRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGPOLYLINEElement) Z_SWAP_PUSH(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGPOLYLINEElement) Z_SWAP_PUSHRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGPOLYLINEElement) Z_TRIGGER(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfZ_TRIGGER(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGPOLYLINEElement) Z_TRIGGERRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGPOLYLINEElement) Z_REQ_METHOD(c SVGPolylineZReqMethodChoice) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGPolylineZReqMethodChoice string

const (
	// default GET
	SVGPolylineZReqMethod_empty SVGPolylineZReqMethodChoice = ""
	// GET
	SVGPolylineZReqMethod_get SVGPolylineZReqMethodChoice = "get"
	// POST
	SVGPolylineZReqMethod_post SVGPolylineZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGPOLYLINEElement) Z_REQ_METHODRemove(c SVGPolylineZReqMethodChoice) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGPOLYLINEElement) Z_REQ_STRATEGY(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGPOLYLINEElement) Z_REQ_STRATEGYRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGPOLYLINEElement) Z_REQ_HISTORY(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGPOLYLINEElement) Z_REQ_HISTORYRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGPOLYLINEElement) Z_DATA(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfZ_DATA(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGPOLYLINEElement) Z_DATARemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGPOLYLINEElement) Z_JSON(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfZ_JSON(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGPOLYLINEElement) Z_JSONRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGPOLYLINEElement) Z_REQ_BATCH(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGPOLYLINEElement) Z_REQ_BATCHRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGPOLYLINEElement) Z_ACTION(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfZ_ACTION(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGPOLYLINEElement) Z_ACTIONRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGPOLYLINEElement) Z_REQ_BEFORE(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGPOLYLINEElement) Z_REQ_BEFORERemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGPOLYLINEElement) Z_REQ_AFTER(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGPOLYLINEElement) Z_REQ_AFTERRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
