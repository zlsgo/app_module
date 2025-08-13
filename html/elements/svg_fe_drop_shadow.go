package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feDropShadow> filter primitive creates a drop shadow of the input image
// It is a shorthand filter, and is defined in terms of the <feGaussianBlur> and
// <feOffset> filter primitives.
type SVGFEDROPSHADOWElement struct {
	*Element
}

// Create a new SVGFEDROPSHADOWElement element.
// This will create a new element with the tag
// "feDropShadow" during rendering.
func SVG_FEDROPSHADOW(children ...ElementRenderer) *SVGFEDROPSHADOWElement {
	e := NewElement("feDropShadow", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEDROPSHADOWElement{Element: e}
}

func (e *SVGFEDROPSHADOWElement) Children(children ...ElementRenderer) *SVGFEDROPSHADOWElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEDROPSHADOWElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEDROPSHADOWElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) Attr(name string, value ...string) *SVGFEDROPSHADOWElement {
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

func (e *SVGFEDROPSHADOWElement) Attrs(attrs ...string) *SVGFEDROPSHADOWElement {
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

func (e *SVGFEDROPSHADOWElement) AttrsMap(attrs map[string]string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) Text(text string) *SVGFEDROPSHADOWElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEDROPSHADOWElement) TextF(format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEDROPSHADOWElement) IfText(condition bool, text string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) IfTextF(condition bool, format string, args ...any) *SVGFEDROPSHADOWElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) Escaped(text string) *SVGFEDROPSHADOWElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEDROPSHADOWElement) IfEscaped(condition bool, text string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) EscapedF(format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEDROPSHADOWElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEDROPSHADOWElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) CustomData(key, value string) *SVGFEDROPSHADOWElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfCustomData(condition bool, key, value string) *SVGFEDROPSHADOWElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) CustomDataF(key, format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEDROPSHADOWElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEDROPSHADOWElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) CustomDataRemove(key string) *SVGFEDROPSHADOWElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The amount of offset in the x direction
// If the <length> is 0, the shadow is placed at the same position as the input.
func (e *SVGFEDROPSHADOWElement) DX(f float64) *SVGFEDROPSHADOWElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("dx", f)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfDX(condition bool, f float64) *SVGFEDROPSHADOWElement {
	if condition {
		e.DX(f)
	}
	return e
}

// The amount of offset in the y direction
// If the <length> is 0, the shadow is placed at the same position as the input.
func (e *SVGFEDROPSHADOWElement) DY(f float64) *SVGFEDROPSHADOWElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("dy", f)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfDY(condition bool, f float64) *SVGFEDROPSHADOWElement {
	if condition {
		e.DY(f)
	}
	return e
}

// The standard deviation for the blur operation
// If two <numbers> are provided, the first number represents a standard deviation
// value along the x-axis of the coordinate system established by attribute
// 'primitiveUnits' on the <filter> element
// The second value represents a standard deviation in Y
// If one number is provided, then that value is used for both X and Y
// Negative values are not allowed
// A value of zero disables the effect of the given filter primitive (i.e., the
// result is a transparent black image).
func (e *SVGFEDROPSHADOWElement) STD_DEVIATION(f float64) *SVGFEDROPSHADOWElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("stdDeviation", f)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfSTD_DEVIATION(condition bool, f float64) *SVGFEDROPSHADOWElement {
	if condition {
		e.STD_DEVIATION(f)
	}
	return e
}

// The flood-color attribute indicates what color to use to flood the current
// filter primitive subregion defined through the <feFlood> element
// If attribute 'flood-color' is not specified, then the effect is as if a value
// of black were specified.
func (e *SVGFEDROPSHADOWElement) FLOOD_COLOR(s string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("flood-color", s)
	return e
}

func (e *SVGFEDROPSHADOWElement) FLOOD_COLORF(format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.FLOOD_COLOR(fmt.Sprintf(format, args...))
}

func (e *SVGFEDROPSHADOWElement) IfFLOOD_COLOR(condition bool, s string) *SVGFEDROPSHADOWElement {
	if condition {
		e.FLOOD_COLOR(s)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) IfFLOOD_COLORF(condition bool, format string, args ...any) *SVGFEDROPSHADOWElement {
	if condition {
		e.FLOOD_COLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute FLOOD_COLOR from the element.
func (e *SVGFEDROPSHADOWElement) FLOOD_COLORRemove(s string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("flood-color")
	return e
}

func (e *SVGFEDROPSHADOWElement) FLOOD_COLORRemoveF(format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.FLOOD_COLORRemove(fmt.Sprintf(format, args...))
}

// The flood-opacity attribute indicates the opacity value to use across the
// current filter primitive subregion defined through the <feFlood> element.
func (e *SVGFEDROPSHADOWElement) FLOOD_OPACITY(f float64) *SVGFEDROPSHADOWElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("flood-opacity", f)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfFLOOD_OPACITY(condition bool, f float64) *SVGFEDROPSHADOWElement {
	if condition {
		e.FLOOD_OPACITY(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFEDROPSHADOWElement) ID(s string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEDROPSHADOWElement) IDF(format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEDROPSHADOWElement) IfID(condition bool, s string) *SVGFEDROPSHADOWElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) IfIDF(condition bool, format string, args ...any) *SVGFEDROPSHADOWElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEDROPSHADOWElement) IDRemove(s string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEDROPSHADOWElement) IDRemoveF(format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEDROPSHADOWElement) CLASS(s ...string) *SVGFEDROPSHADOWElement {
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

func (e *SVGFEDROPSHADOWElement) IfCLASS(condition bool, s ...string) *SVGFEDROPSHADOWElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEDROPSHADOWElement) CLASSRemove(s ...string) *SVGFEDROPSHADOWElement {
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
func (e *SVGFEDROPSHADOWElement) STYLEF(k string, format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEDROPSHADOWElement) IfSTYLE(condition bool, k string, v string) *SVGFEDROPSHADOWElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) STYLE(k string, v string) *SVGFEDROPSHADOWElement {
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

func (e *SVGFEDROPSHADOWElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEDROPSHADOWElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEDROPSHADOWElement) STYLEMap(m map[string]string) *SVGFEDROPSHADOWElement {
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
func (e *SVGFEDROPSHADOWElement) STYLEPairs(pairs ...string) *SVGFEDROPSHADOWElement {
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

func (e *SVGFEDROPSHADOWElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEDROPSHADOWElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEDROPSHADOWElement) STYLERemove(keys ...string) *SVGFEDROPSHADOWElement {
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

func (e *SVGFEDROPSHADOWElement) Z_REQ(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfZ_REQ(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEDROPSHADOWElement) Z_REQRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEDROPSHADOWElement) Z_TARGET(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfZ_TARGET(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEDROPSHADOWElement) Z_TARGETRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEDROPSHADOWElement) Z_REQ_SELECTOR(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEDROPSHADOWElement) Z_REQ_SELECTORRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEDROPSHADOWElement) Z_SWAP(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfZ_SWAP(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEDROPSHADOWElement) Z_SWAPRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEDROPSHADOWElement) Z_SWAP_PUSH(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEDROPSHADOWElement) Z_SWAP_PUSHRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEDROPSHADOWElement) Z_TRIGGER(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEDROPSHADOWElement) Z_TRIGGERRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEDROPSHADOWElement) Z_REQ_METHOD(c SVGFeDropShadowZReqMethodChoice) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeDropShadowZReqMethodChoice string

const (
	// default GET
	SVGFeDropShadowZReqMethod_empty SVGFeDropShadowZReqMethodChoice = ""
	// GET
	SVGFeDropShadowZReqMethod_get SVGFeDropShadowZReqMethodChoice = "get"
	// POST
	SVGFeDropShadowZReqMethod_post SVGFeDropShadowZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEDROPSHADOWElement) Z_REQ_METHODRemove(c SVGFeDropShadowZReqMethodChoice) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEDROPSHADOWElement) Z_REQ_STRATEGY(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEDROPSHADOWElement) Z_REQ_STRATEGYRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEDROPSHADOWElement) Z_REQ_HISTORY(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEDROPSHADOWElement) Z_REQ_HISTORYRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEDROPSHADOWElement) Z_DATA(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfZ_DATA(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEDROPSHADOWElement) Z_DATARemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEDROPSHADOWElement) Z_JSON(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfZ_JSON(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEDROPSHADOWElement) Z_JSONRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEDROPSHADOWElement) Z_REQ_BATCH(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEDROPSHADOWElement) Z_REQ_BATCHRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEDROPSHADOWElement) Z_ACTION(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfZ_ACTION(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEDROPSHADOWElement) Z_ACTIONRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEDROPSHADOWElement) Z_REQ_BEFORE(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEDROPSHADOWElement) Z_REQ_BEFORERemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEDROPSHADOWElement) Z_REQ_AFTER(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEDROPSHADOWElement) Z_REQ_AFTERRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
