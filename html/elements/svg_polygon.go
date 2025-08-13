package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <polygon> SVG element is an SVG basic shape, used to create a vector-based
// polygonal shape.
type SVGPOLYGONElement struct {
	*Element
}

// Create a new SVGPOLYGONElement element.
// This will create a new element with the tag
// "polygon" during rendering.
func SVG_POLYGON(children ...ElementRenderer) *SVGPOLYGONElement {
	e := NewElement("polygon", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGPOLYGONElement{Element: e}
}

func (e *SVGPOLYGONElement) Children(children ...ElementRenderer) *SVGPOLYGONElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGPOLYGONElement) IfChildren(condition bool, children ...ElementRenderer) *SVGPOLYGONElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGPOLYGONElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGPOLYGONElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGPOLYGONElement) Attr(name string, value ...string) *SVGPOLYGONElement {
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

func (e *SVGPOLYGONElement) Attrs(attrs ...string) *SVGPOLYGONElement {
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

func (e *SVGPOLYGONElement) AttrsMap(attrs map[string]string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGPOLYGONElement) Text(text string) *SVGPOLYGONElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGPOLYGONElement) TextF(format string, args ...any) *SVGPOLYGONElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGPOLYGONElement) IfText(condition bool, text string) *SVGPOLYGONElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGPOLYGONElement) IfTextF(condition bool, format string, args ...any) *SVGPOLYGONElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGPOLYGONElement) Escaped(text string) *SVGPOLYGONElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGPOLYGONElement) IfEscaped(condition bool, text string) *SVGPOLYGONElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGPOLYGONElement) EscapedF(format string, args ...any) *SVGPOLYGONElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGPOLYGONElement) IfEscapedF(condition bool, format string, args ...any) *SVGPOLYGONElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGPOLYGONElement) CustomData(key, value string) *SVGPOLYGONElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGPOLYGONElement) IfCustomData(condition bool, key, value string) *SVGPOLYGONElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGPOLYGONElement) CustomDataF(key, format string, args ...any) *SVGPOLYGONElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGPOLYGONElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGPOLYGONElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGPOLYGONElement) CustomDataRemove(key string) *SVGPOLYGONElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// A list of points, each of which is a coordinate pair.
func (e *SVGPOLYGONElement) POINTS(s string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("points", s)
	return e
}

func (e *SVGPOLYGONElement) POINTSF(format string, args ...any) *SVGPOLYGONElement {
	return e.POINTS(fmt.Sprintf(format, args...))
}

func (e *SVGPOLYGONElement) IfPOINTS(condition bool, s string) *SVGPOLYGONElement {
	if condition {
		e.POINTS(s)
	}
	return e
}

func (e *SVGPOLYGONElement) IfPOINTSF(condition bool, format string, args ...any) *SVGPOLYGONElement {
	if condition {
		e.POINTS(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute POINTS from the element.
func (e *SVGPOLYGONElement) POINTSRemove(s string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("points")
	return e
}

func (e *SVGPOLYGONElement) POINTSRemoveF(format string, args ...any) *SVGPOLYGONElement {
	return e.POINTSRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGPOLYGONElement) ID(s string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGPOLYGONElement) IDF(format string, args ...any) *SVGPOLYGONElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGPOLYGONElement) IfID(condition bool, s string) *SVGPOLYGONElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGPOLYGONElement) IfIDF(condition bool, format string, args ...any) *SVGPOLYGONElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGPOLYGONElement) IDRemove(s string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGPOLYGONElement) IDRemoveF(format string, args ...any) *SVGPOLYGONElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGPOLYGONElement) CLASS(s ...string) *SVGPOLYGONElement {
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

func (e *SVGPOLYGONElement) IfCLASS(condition bool, s ...string) *SVGPOLYGONElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGPOLYGONElement) CLASSRemove(s ...string) *SVGPOLYGONElement {
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
func (e *SVGPOLYGONElement) STYLEF(k string, format string, args ...any) *SVGPOLYGONElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGPOLYGONElement) IfSTYLE(condition bool, k string, v string) *SVGPOLYGONElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGPOLYGONElement) STYLE(k string, v string) *SVGPOLYGONElement {
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

func (e *SVGPOLYGONElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGPOLYGONElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGPOLYGONElement) STYLEMap(m map[string]string) *SVGPOLYGONElement {
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
func (e *SVGPOLYGONElement) STYLEPairs(pairs ...string) *SVGPOLYGONElement {
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

func (e *SVGPOLYGONElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGPOLYGONElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGPOLYGONElement) STYLERemove(keys ...string) *SVGPOLYGONElement {
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

func (e *SVGPOLYGONElement) Z_REQ(expression string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYGONElement) IfZ_REQ(condition bool, expression string) *SVGPOLYGONElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGPOLYGONElement) Z_REQRemove() *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGPOLYGONElement) Z_TARGET(expression string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYGONElement) IfZ_TARGET(condition bool, expression string) *SVGPOLYGONElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGPOLYGONElement) Z_TARGETRemove() *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGPOLYGONElement) Z_REQ_SELECTOR(expression string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYGONElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGPOLYGONElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGPOLYGONElement) Z_REQ_SELECTORRemove() *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGPOLYGONElement) Z_SWAP(expression string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYGONElement) IfZ_SWAP(condition bool, expression string) *SVGPOLYGONElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGPOLYGONElement) Z_SWAPRemove() *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGPOLYGONElement) Z_SWAP_PUSH(expression string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYGONElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGPOLYGONElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGPOLYGONElement) Z_SWAP_PUSHRemove() *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGPOLYGONElement) Z_TRIGGER(expression string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYGONElement) IfZ_TRIGGER(condition bool, expression string) *SVGPOLYGONElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGPOLYGONElement) Z_TRIGGERRemove() *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGPOLYGONElement) Z_REQ_METHOD(c SVGPolygonZReqMethodChoice) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGPolygonZReqMethodChoice string

const (
	// default GET
	SVGPolygonZReqMethod_empty SVGPolygonZReqMethodChoice = ""
	// GET
	SVGPolygonZReqMethod_get SVGPolygonZReqMethodChoice = "get"
	// POST
	SVGPolygonZReqMethod_post SVGPolygonZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGPOLYGONElement) Z_REQ_METHODRemove(c SVGPolygonZReqMethodChoice) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGPOLYGONElement) Z_REQ_STRATEGY(expression string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYGONElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGPOLYGONElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGPOLYGONElement) Z_REQ_STRATEGYRemove() *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGPOLYGONElement) Z_REQ_HISTORY(expression string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYGONElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGPOLYGONElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGPOLYGONElement) Z_REQ_HISTORYRemove() *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGPOLYGONElement) Z_DATA(expression string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYGONElement) IfZ_DATA(condition bool, expression string) *SVGPOLYGONElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGPOLYGONElement) Z_DATARemove() *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGPOLYGONElement) Z_JSON(expression string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYGONElement) IfZ_JSON(condition bool, expression string) *SVGPOLYGONElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGPOLYGONElement) Z_JSONRemove() *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGPOLYGONElement) Z_REQ_BATCH(expression string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYGONElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGPOLYGONElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGPOLYGONElement) Z_REQ_BATCHRemove() *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGPOLYGONElement) Z_ACTION(expression string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYGONElement) IfZ_ACTION(condition bool, expression string) *SVGPOLYGONElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGPOLYGONElement) Z_ACTIONRemove() *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGPOLYGONElement) Z_REQ_BEFORE(expression string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYGONElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGPOLYGONElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGPOLYGONElement) Z_REQ_BEFORERemove() *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGPOLYGONElement) Z_REQ_AFTER(expression string) *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYGONElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGPOLYGONElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGPOLYGONElement) Z_REQ_AFTERRemove() *SVGPOLYGONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
