package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feComposite> SVG filter primitive performs the combination of two input
// images pixel-wise in image space using one of the Porter-Duff compositing
// operations: over, in, atop, out, xor.
type SVGFECOMPOSITEElement struct {
	*Element
}

// Create a new SVGFECOMPOSITEElement element.
// This will create a new element with the tag
// "feComposite" during rendering.
func SVG_FECOMPOSITE(children ...ElementRenderer) *SVGFECOMPOSITEElement {
	e := NewElement("feComposite", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFECOMPOSITEElement{Element: e}
}

func (e *SVGFECOMPOSITEElement) Children(children ...ElementRenderer) *SVGFECOMPOSITEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFECOMPOSITEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFECOMPOSITEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFECOMPOSITEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFECOMPOSITEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFECOMPOSITEElement) Attr(name string, value ...string) *SVGFECOMPOSITEElement {
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

func (e *SVGFECOMPOSITEElement) Attrs(attrs ...string) *SVGFECOMPOSITEElement {
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

func (e *SVGFECOMPOSITEElement) AttrsMap(attrs map[string]string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFECOMPOSITEElement) Text(text string) *SVGFECOMPOSITEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFECOMPOSITEElement) TextF(format string, args ...any) *SVGFECOMPOSITEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPOSITEElement) IfText(condition bool, text string) *SVGFECOMPOSITEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFECOMPOSITEElement) IfTextF(condition bool, format string, args ...any) *SVGFECOMPOSITEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFECOMPOSITEElement) Escaped(text string) *SVGFECOMPOSITEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFECOMPOSITEElement) IfEscaped(condition bool, text string) *SVGFECOMPOSITEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFECOMPOSITEElement) EscapedF(format string, args ...any) *SVGFECOMPOSITEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPOSITEElement) IfEscapedF(condition bool, format string, args ...any) *SVGFECOMPOSITEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFECOMPOSITEElement) CustomData(key, value string) *SVGFECOMPOSITEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFECOMPOSITEElement) IfCustomData(condition bool, key, value string) *SVGFECOMPOSITEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFECOMPOSITEElement) CustomDataF(key, format string, args ...any) *SVGFECOMPOSITEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPOSITEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFECOMPOSITEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFECOMPOSITEElement) CustomDataRemove(key string) *SVGFECOMPOSITEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// Input for the compositing operation.
func (e *SVGFECOMPOSITEElement) IN(s string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFECOMPOSITEElement) INF(format string, args ...any) *SVGFECOMPOSITEElement {
	return e.IN(fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPOSITEElement) IfIN(condition bool, s string) *SVGFECOMPOSITEElement {
	if condition {
		e.IN(s)
	}
	return e
}

func (e *SVGFECOMPOSITEElement) IfINF(condition bool, format string, args ...any) *SVGFECOMPOSITEElement {
	if condition {
		e.IN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFECOMPOSITEElement) INRemove(s string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("in")
	return e
}

func (e *SVGFECOMPOSITEElement) INRemoveF(format string, args ...any) *SVGFECOMPOSITEElement {
	return e.INRemove(fmt.Sprintf(format, args...))
}

// Second input for the compositing operation.
func (e *SVGFECOMPOSITEElement) IN_2(s string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("in2", s)
	return e
}

func (e *SVGFECOMPOSITEElement) IN_2F(format string, args ...any) *SVGFECOMPOSITEElement {
	return e.IN_2(fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPOSITEElement) IfIN_2(condition bool, s string) *SVGFECOMPOSITEElement {
	if condition {
		e.IN_2(s)
	}
	return e
}

func (e *SVGFECOMPOSITEElement) IfIN_2F(condition bool, format string, args ...any) *SVGFECOMPOSITEElement {
	if condition {
		e.IN_2(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN_2 from the element.
func (e *SVGFECOMPOSITEElement) IN_2Remove(s string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("in2")
	return e
}

func (e *SVGFECOMPOSITEElement) IN_2RemoveF(format string, args ...any) *SVGFECOMPOSITEElement {
	return e.IN_2Remove(fmt.Sprintf(format, args...))
}

// The type of compositing operation.
func (e *SVGFECOMPOSITEElement) OPERATOR(c SVGFeCompositeOperatorChoice) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("operator", string(c))
	return e
}

type SVGFeCompositeOperatorChoice string

const (
	// The source image is composited over the destination image.
	SVGFeCompositeOperator_over SVGFeCompositeOperatorChoice = "over"
	// The part of the source image that lies inside of the destination image is
	// composited over the destination image.
	SVGFeCompositeOperator_in SVGFeCompositeOperatorChoice = "in"
	// The part of the source image that lies outside of the destination image is
	// composited over the destination image.
	SVGFeCompositeOperator_out SVGFeCompositeOperatorChoice = "out"
	// The part of the source image that lies inside of the destination image is
	// composited over the destination image and replaces the destination image.
	SVGFeCompositeOperator_atop SVGFeCompositeOperatorChoice = "atop"
	// The part of the source image that lies outside of the destination image is
	// composited over the destination image.
	SVGFeCompositeOperator_xor SVGFeCompositeOperatorChoice = "xor"
	// A standard arithmetic operator is applied (
	SVGFeCompositeOperator_arithmetic SVGFeCompositeOperatorChoice = "arithmetic"
)

// Remove the attribute OPERATOR from the element.
func (e *SVGFECOMPOSITEElement) OPERATORRemove(c SVGFeCompositeOperatorChoice) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("operator")
	return e
}

// First value to use in the arithmetic operation.
func (e *SVGFECOMPOSITEElement) K_1(f float64) *SVGFECOMPOSITEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("k1", f)
	return e
}

func (e *SVGFECOMPOSITEElement) IfK_1(condition bool, f float64) *SVGFECOMPOSITEElement {
	if condition {
		e.K_1(f)
	}
	return e
}

// Second value to use in the arithmetic operation.
func (e *SVGFECOMPOSITEElement) K_2(f float64) *SVGFECOMPOSITEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("k2", f)
	return e
}

func (e *SVGFECOMPOSITEElement) IfK_2(condition bool, f float64) *SVGFECOMPOSITEElement {
	if condition {
		e.K_2(f)
	}
	return e
}

// Third value to use in the arithmetic operation.
func (e *SVGFECOMPOSITEElement) K_3(f float64) *SVGFECOMPOSITEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("k3", f)
	return e
}

func (e *SVGFECOMPOSITEElement) IfK_3(condition bool, f float64) *SVGFECOMPOSITEElement {
	if condition {
		e.K_3(f)
	}
	return e
}

// Fourth value to use in the arithmetic operation.
func (e *SVGFECOMPOSITEElement) K_4(f float64) *SVGFECOMPOSITEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("k4", f)
	return e
}

func (e *SVGFECOMPOSITEElement) IfK_4(condition bool, f float64) *SVGFECOMPOSITEElement {
	if condition {
		e.K_4(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFECOMPOSITEElement) ID(s string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFECOMPOSITEElement) IDF(format string, args ...any) *SVGFECOMPOSITEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPOSITEElement) IfID(condition bool, s string) *SVGFECOMPOSITEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFECOMPOSITEElement) IfIDF(condition bool, format string, args ...any) *SVGFECOMPOSITEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFECOMPOSITEElement) IDRemove(s string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFECOMPOSITEElement) IDRemoveF(format string, args ...any) *SVGFECOMPOSITEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFECOMPOSITEElement) CLASS(s ...string) *SVGFECOMPOSITEElement {
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

func (e *SVGFECOMPOSITEElement) IfCLASS(condition bool, s ...string) *SVGFECOMPOSITEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFECOMPOSITEElement) CLASSRemove(s ...string) *SVGFECOMPOSITEElement {
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
func (e *SVGFECOMPOSITEElement) STYLEF(k string, format string, args ...any) *SVGFECOMPOSITEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPOSITEElement) IfSTYLE(condition bool, k string, v string) *SVGFECOMPOSITEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFECOMPOSITEElement) STYLE(k string, v string) *SVGFECOMPOSITEElement {
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

func (e *SVGFECOMPOSITEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFECOMPOSITEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFECOMPOSITEElement) STYLEMap(m map[string]string) *SVGFECOMPOSITEElement {
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
func (e *SVGFECOMPOSITEElement) STYLEPairs(pairs ...string) *SVGFECOMPOSITEElement {
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

func (e *SVGFECOMPOSITEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFECOMPOSITEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFECOMPOSITEElement) STYLERemove(keys ...string) *SVGFECOMPOSITEElement {
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

func (e *SVGFECOMPOSITEElement) Z_REQ(expression string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPOSITEElement) IfZ_REQ(condition bool, expression string) *SVGFECOMPOSITEElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFECOMPOSITEElement) Z_REQRemove() *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFECOMPOSITEElement) Z_TARGET(expression string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPOSITEElement) IfZ_TARGET(condition bool, expression string) *SVGFECOMPOSITEElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFECOMPOSITEElement) Z_TARGETRemove() *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFECOMPOSITEElement) Z_REQ_SELECTOR(expression string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPOSITEElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFECOMPOSITEElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFECOMPOSITEElement) Z_REQ_SELECTORRemove() *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFECOMPOSITEElement) Z_SWAP(expression string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPOSITEElement) IfZ_SWAP(condition bool, expression string) *SVGFECOMPOSITEElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFECOMPOSITEElement) Z_SWAPRemove() *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFECOMPOSITEElement) Z_SWAP_PUSH(expression string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPOSITEElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFECOMPOSITEElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFECOMPOSITEElement) Z_SWAP_PUSHRemove() *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFECOMPOSITEElement) Z_TRIGGER(expression string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPOSITEElement) IfZ_TRIGGER(condition bool, expression string) *SVGFECOMPOSITEElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFECOMPOSITEElement) Z_TRIGGERRemove() *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFECOMPOSITEElement) Z_REQ_METHOD(c SVGFeCompositeZReqMethodChoice) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeCompositeZReqMethodChoice string

const (
	// default GET
	SVGFeCompositeZReqMethod_empty SVGFeCompositeZReqMethodChoice = ""
	// GET
	SVGFeCompositeZReqMethod_get SVGFeCompositeZReqMethodChoice = "get"
	// POST
	SVGFeCompositeZReqMethod_post SVGFeCompositeZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFECOMPOSITEElement) Z_REQ_METHODRemove(c SVGFeCompositeZReqMethodChoice) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFECOMPOSITEElement) Z_REQ_STRATEGY(expression string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPOSITEElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFECOMPOSITEElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFECOMPOSITEElement) Z_REQ_STRATEGYRemove() *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFECOMPOSITEElement) Z_REQ_HISTORY(expression string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPOSITEElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFECOMPOSITEElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFECOMPOSITEElement) Z_REQ_HISTORYRemove() *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFECOMPOSITEElement) Z_DATA(expression string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPOSITEElement) IfZ_DATA(condition bool, expression string) *SVGFECOMPOSITEElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFECOMPOSITEElement) Z_DATARemove() *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFECOMPOSITEElement) Z_JSON(expression string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPOSITEElement) IfZ_JSON(condition bool, expression string) *SVGFECOMPOSITEElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFECOMPOSITEElement) Z_JSONRemove() *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFECOMPOSITEElement) Z_REQ_BATCH(expression string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPOSITEElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFECOMPOSITEElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFECOMPOSITEElement) Z_REQ_BATCHRemove() *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFECOMPOSITEElement) Z_ACTION(expression string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPOSITEElement) IfZ_ACTION(condition bool, expression string) *SVGFECOMPOSITEElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFECOMPOSITEElement) Z_ACTIONRemove() *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFECOMPOSITEElement) Z_REQ_BEFORE(expression string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPOSITEElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFECOMPOSITEElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFECOMPOSITEElement) Z_REQ_BEFORERemove() *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFECOMPOSITEElement) Z_REQ_AFTER(expression string) *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPOSITEElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFECOMPOSITEElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFECOMPOSITEElement) Z_REQ_AFTERRemove() *SVGFECOMPOSITEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
