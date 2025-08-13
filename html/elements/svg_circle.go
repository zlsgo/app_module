package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <circle> SVG element is an SVG basic shape, used to create circles based on
// a center point and a radius.
type SVGCIRCLEElement struct {
	*Element
}

// Create a new SVGCIRCLEElement element.
// This will create a new element with the tag
// "circle" during rendering.
func SVG_CIRCLE(children ...ElementRenderer) *SVGCIRCLEElement {
	e := NewElement("circle", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGCIRCLEElement{Element: e}
}

func (e *SVGCIRCLEElement) Children(children ...ElementRenderer) *SVGCIRCLEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGCIRCLEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGCIRCLEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGCIRCLEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGCIRCLEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGCIRCLEElement) Attr(name string, value ...string) *SVGCIRCLEElement {
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

func (e *SVGCIRCLEElement) Attrs(attrs ...string) *SVGCIRCLEElement {
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

func (e *SVGCIRCLEElement) AttrsMap(attrs map[string]string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGCIRCLEElement) Text(text string) *SVGCIRCLEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGCIRCLEElement) TextF(format string, args ...any) *SVGCIRCLEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGCIRCLEElement) IfText(condition bool, text string) *SVGCIRCLEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGCIRCLEElement) IfTextF(condition bool, format string, args ...any) *SVGCIRCLEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGCIRCLEElement) Escaped(text string) *SVGCIRCLEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGCIRCLEElement) IfEscaped(condition bool, text string) *SVGCIRCLEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGCIRCLEElement) EscapedF(format string, args ...any) *SVGCIRCLEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGCIRCLEElement) IfEscapedF(condition bool, format string, args ...any) *SVGCIRCLEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGCIRCLEElement) CustomData(key, value string) *SVGCIRCLEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGCIRCLEElement) IfCustomData(condition bool, key, value string) *SVGCIRCLEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGCIRCLEElement) CustomDataF(key, format string, args ...any) *SVGCIRCLEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGCIRCLEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGCIRCLEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGCIRCLEElement) CustomDataRemove(key string) *SVGCIRCLEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The x-axis coordinate of the center of the circle.
func (e *SVGCIRCLEElement) CX(f float64) *SVGCIRCLEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("cx", f)
	return e
}

func (e *SVGCIRCLEElement) IfCX(condition bool, f float64) *SVGCIRCLEElement {
	if condition {
		e.CX(f)
	}
	return e
}

// The y-axis coordinate of the center of the circle.
func (e *SVGCIRCLEElement) CY(f float64) *SVGCIRCLEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("cy", f)
	return e
}

func (e *SVGCIRCLEElement) IfCY(condition bool, f float64) *SVGCIRCLEElement {
	if condition {
		e.CY(f)
	}
	return e
}

// The radius of the circle.
func (e *SVGCIRCLEElement) R(f float64) *SVGCIRCLEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("r", f)
	return e
}

func (e *SVGCIRCLEElement) IfR(condition bool, f float64) *SVGCIRCLEElement {
	if condition {
		e.R(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGCIRCLEElement) ID(s string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGCIRCLEElement) IDF(format string, args ...any) *SVGCIRCLEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGCIRCLEElement) IfID(condition bool, s string) *SVGCIRCLEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGCIRCLEElement) IfIDF(condition bool, format string, args ...any) *SVGCIRCLEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGCIRCLEElement) IDRemove(s string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGCIRCLEElement) IDRemoveF(format string, args ...any) *SVGCIRCLEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGCIRCLEElement) CLASS(s ...string) *SVGCIRCLEElement {
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

func (e *SVGCIRCLEElement) IfCLASS(condition bool, s ...string) *SVGCIRCLEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGCIRCLEElement) CLASSRemove(s ...string) *SVGCIRCLEElement {
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
func (e *SVGCIRCLEElement) STYLEF(k string, format string, args ...any) *SVGCIRCLEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGCIRCLEElement) IfSTYLE(condition bool, k string, v string) *SVGCIRCLEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGCIRCLEElement) STYLE(k string, v string) *SVGCIRCLEElement {
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

func (e *SVGCIRCLEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGCIRCLEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGCIRCLEElement) STYLEMap(m map[string]string) *SVGCIRCLEElement {
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
func (e *SVGCIRCLEElement) STYLEPairs(pairs ...string) *SVGCIRCLEElement {
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

func (e *SVGCIRCLEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGCIRCLEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGCIRCLEElement) STYLERemove(keys ...string) *SVGCIRCLEElement {
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

func (e *SVGCIRCLEElement) Z_REQ(expression string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCIRCLEElement) IfZ_REQ(condition bool, expression string) *SVGCIRCLEElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGCIRCLEElement) Z_REQRemove() *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGCIRCLEElement) Z_TARGET(expression string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCIRCLEElement) IfZ_TARGET(condition bool, expression string) *SVGCIRCLEElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGCIRCLEElement) Z_TARGETRemove() *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGCIRCLEElement) Z_REQ_SELECTOR(expression string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCIRCLEElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGCIRCLEElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGCIRCLEElement) Z_REQ_SELECTORRemove() *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGCIRCLEElement) Z_SWAP(expression string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCIRCLEElement) IfZ_SWAP(condition bool, expression string) *SVGCIRCLEElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGCIRCLEElement) Z_SWAPRemove() *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGCIRCLEElement) Z_SWAP_PUSH(expression string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCIRCLEElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGCIRCLEElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGCIRCLEElement) Z_SWAP_PUSHRemove() *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGCIRCLEElement) Z_TRIGGER(expression string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCIRCLEElement) IfZ_TRIGGER(condition bool, expression string) *SVGCIRCLEElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGCIRCLEElement) Z_TRIGGERRemove() *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGCIRCLEElement) Z_REQ_METHOD(c SVGCircleZReqMethodChoice) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGCircleZReqMethodChoice string

const (
	// default GET
	SVGCircleZReqMethod_empty SVGCircleZReqMethodChoice = ""
	// GET
	SVGCircleZReqMethod_get SVGCircleZReqMethodChoice = "get"
	// POST
	SVGCircleZReqMethod_post SVGCircleZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGCIRCLEElement) Z_REQ_METHODRemove(c SVGCircleZReqMethodChoice) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGCIRCLEElement) Z_REQ_STRATEGY(expression string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCIRCLEElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGCIRCLEElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGCIRCLEElement) Z_REQ_STRATEGYRemove() *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGCIRCLEElement) Z_REQ_HISTORY(expression string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCIRCLEElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGCIRCLEElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGCIRCLEElement) Z_REQ_HISTORYRemove() *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGCIRCLEElement) Z_DATA(expression string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCIRCLEElement) IfZ_DATA(condition bool, expression string) *SVGCIRCLEElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGCIRCLEElement) Z_DATARemove() *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGCIRCLEElement) Z_JSON(expression string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCIRCLEElement) IfZ_JSON(condition bool, expression string) *SVGCIRCLEElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGCIRCLEElement) Z_JSONRemove() *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGCIRCLEElement) Z_REQ_BATCH(expression string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCIRCLEElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGCIRCLEElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGCIRCLEElement) Z_REQ_BATCHRemove() *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGCIRCLEElement) Z_ACTION(expression string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCIRCLEElement) IfZ_ACTION(condition bool, expression string) *SVGCIRCLEElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGCIRCLEElement) Z_ACTIONRemove() *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGCIRCLEElement) Z_REQ_BEFORE(expression string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCIRCLEElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGCIRCLEElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGCIRCLEElement) Z_REQ_BEFORERemove() *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGCIRCLEElement) Z_REQ_AFTER(expression string) *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCIRCLEElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGCIRCLEElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGCIRCLEElement) Z_REQ_AFTERRemove() *SVGCIRCLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
