package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feGaussianBlur> SVG filter primitive blurs the input image by the amount
// specified in stdDeviation, which defines the bell-curve.
type SVGFEGAUSSIANBLURElement struct {
	*Element
}

// Create a new SVGFEGAUSSIANBLURElement element.
// This will create a new element with the tag
// "feGaussianBlur" during rendering.
func SVG_FEGAUSSIANBLUR(children ...ElementRenderer) *SVGFEGAUSSIANBLURElement {
	e := NewElement("feGaussianBlur", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEGAUSSIANBLURElement{Element: e}
}

func (e *SVGFEGAUSSIANBLURElement) Children(children ...ElementRenderer) *SVGFEGAUSSIANBLURElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) Attr(name string, value ...string) *SVGFEGAUSSIANBLURElement {
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

func (e *SVGFEGAUSSIANBLURElement) Attrs(attrs ...string) *SVGFEGAUSSIANBLURElement {
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

func (e *SVGFEGAUSSIANBLURElement) AttrsMap(attrs map[string]string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) Text(text string) *SVGFEGAUSSIANBLURElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEGAUSSIANBLURElement) TextF(format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEGAUSSIANBLURElement) IfText(condition bool, text string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfTextF(condition bool, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) Escaped(text string) *SVGFEGAUSSIANBLURElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfEscaped(condition bool, text string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) EscapedF(format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEGAUSSIANBLURElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) CustomData(key, value string) *SVGFEGAUSSIANBLURElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfCustomData(condition bool, key, value string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) CustomDataF(key, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEGAUSSIANBLURElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) CustomDataRemove(key string) *SVGFEGAUSSIANBLURElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The input for this filter.
func (e *SVGFEGAUSSIANBLURElement) IN(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) INF(format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.IN(fmt.Sprintf(format, args...))
}

func (e *SVGFEGAUSSIANBLURElement) IfIN(condition bool, s string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.IN(s)
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfINF(condition bool, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.IN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFEGAUSSIANBLURElement) INRemove(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("in")
	return e
}

func (e *SVGFEGAUSSIANBLURElement) INRemoveF(format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.INRemove(fmt.Sprintf(format, args...))
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
func (e *SVGFEGAUSSIANBLURElement) STD_DEVIATION(f float64) *SVGFEGAUSSIANBLURElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("stdDeviation", f)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfSTD_DEVIATION(condition bool, f float64) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.STD_DEVIATION(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFEGAUSSIANBLURElement) ID(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IDF(format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEGAUSSIANBLURElement) IfID(condition bool, s string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfIDF(condition bool, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEGAUSSIANBLURElement) IDRemove(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IDRemoveF(format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEGAUSSIANBLURElement) CLASS(s ...string) *SVGFEGAUSSIANBLURElement {
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

func (e *SVGFEGAUSSIANBLURElement) IfCLASS(condition bool, s ...string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEGAUSSIANBLURElement) CLASSRemove(s ...string) *SVGFEGAUSSIANBLURElement {
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
func (e *SVGFEGAUSSIANBLURElement) STYLEF(k string, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEGAUSSIANBLURElement) IfSTYLE(condition bool, k string, v string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) STYLE(k string, v string) *SVGFEGAUSSIANBLURElement {
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

func (e *SVGFEGAUSSIANBLURElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEGAUSSIANBLURElement) STYLEMap(m map[string]string) *SVGFEGAUSSIANBLURElement {
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
func (e *SVGFEGAUSSIANBLURElement) STYLEPairs(pairs ...string) *SVGFEGAUSSIANBLURElement {
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

func (e *SVGFEGAUSSIANBLURElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEGAUSSIANBLURElement) STYLERemove(keys ...string) *SVGFEGAUSSIANBLURElement {
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

func (e *SVGFEGAUSSIANBLURElement) Z_REQ(expression string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfZ_REQ(condition bool, expression string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEGAUSSIANBLURElement) Z_REQRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEGAUSSIANBLURElement) Z_TARGET(expression string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfZ_TARGET(condition bool, expression string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEGAUSSIANBLURElement) Z_TARGETRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEGAUSSIANBLURElement) Z_REQ_SELECTOR(expression string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEGAUSSIANBLURElement) Z_REQ_SELECTORRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEGAUSSIANBLURElement) Z_SWAP(expression string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfZ_SWAP(condition bool, expression string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEGAUSSIANBLURElement) Z_SWAPRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEGAUSSIANBLURElement) Z_SWAP_PUSH(expression string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEGAUSSIANBLURElement) Z_SWAP_PUSHRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEGAUSSIANBLURElement) Z_TRIGGER(expression string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEGAUSSIANBLURElement) Z_TRIGGERRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEGAUSSIANBLURElement) Z_REQ_METHOD(c SVGFeGaussianBlurZReqMethodChoice) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeGaussianBlurZReqMethodChoice string

const (
	// default GET
	SVGFeGaussianBlurZReqMethod_empty SVGFeGaussianBlurZReqMethodChoice = ""
	// GET
	SVGFeGaussianBlurZReqMethod_get SVGFeGaussianBlurZReqMethodChoice = "get"
	// POST
	SVGFeGaussianBlurZReqMethod_post SVGFeGaussianBlurZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEGAUSSIANBLURElement) Z_REQ_METHODRemove(c SVGFeGaussianBlurZReqMethodChoice) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEGAUSSIANBLURElement) Z_REQ_STRATEGY(expression string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEGAUSSIANBLURElement) Z_REQ_STRATEGYRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEGAUSSIANBLURElement) Z_REQ_HISTORY(expression string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEGAUSSIANBLURElement) Z_REQ_HISTORYRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEGAUSSIANBLURElement) Z_DATA(expression string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfZ_DATA(condition bool, expression string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEGAUSSIANBLURElement) Z_DATARemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEGAUSSIANBLURElement) Z_JSON(expression string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfZ_JSON(condition bool, expression string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEGAUSSIANBLURElement) Z_JSONRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEGAUSSIANBLURElement) Z_REQ_BATCH(expression string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEGAUSSIANBLURElement) Z_REQ_BATCHRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEGAUSSIANBLURElement) Z_ACTION(expression string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfZ_ACTION(condition bool, expression string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEGAUSSIANBLURElement) Z_ACTIONRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEGAUSSIANBLURElement) Z_REQ_BEFORE(expression string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEGAUSSIANBLURElement) Z_REQ_BEFORERemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEGAUSSIANBLURElement) Z_REQ_AFTER(expression string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEGAUSSIANBLURElement) Z_REQ_AFTERRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
