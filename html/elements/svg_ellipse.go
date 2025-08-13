package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <ellipse> SVG element is an SVG basic shape, used to create ellipses based
// on a center coordinate, and both their x and y radius.
type SVGELLIPSEElement struct {
	*Element
}

// Create a new SVGELLIPSEElement element.
// This will create a new element with the tag
// "ellipse" during rendering.
func SVG_ELLIPSE(children ...ElementRenderer) *SVGELLIPSEElement {
	e := NewElement("ellipse", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGELLIPSEElement{Element: e}
}

func (e *SVGELLIPSEElement) Children(children ...ElementRenderer) *SVGELLIPSEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGELLIPSEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGELLIPSEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGELLIPSEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGELLIPSEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGELLIPSEElement) Attr(name string, value ...string) *SVGELLIPSEElement {
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

func (e *SVGELLIPSEElement) Attrs(attrs ...string) *SVGELLIPSEElement {
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

func (e *SVGELLIPSEElement) AttrsMap(attrs map[string]string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGELLIPSEElement) Text(text string) *SVGELLIPSEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGELLIPSEElement) TextF(format string, args ...any) *SVGELLIPSEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGELLIPSEElement) IfText(condition bool, text string) *SVGELLIPSEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGELLIPSEElement) IfTextF(condition bool, format string, args ...any) *SVGELLIPSEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGELLIPSEElement) Escaped(text string) *SVGELLIPSEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGELLIPSEElement) IfEscaped(condition bool, text string) *SVGELLIPSEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGELLIPSEElement) EscapedF(format string, args ...any) *SVGELLIPSEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGELLIPSEElement) IfEscapedF(condition bool, format string, args ...any) *SVGELLIPSEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGELLIPSEElement) CustomData(key, value string) *SVGELLIPSEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGELLIPSEElement) IfCustomData(condition bool, key, value string) *SVGELLIPSEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGELLIPSEElement) CustomDataF(key, format string, args ...any) *SVGELLIPSEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGELLIPSEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGELLIPSEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGELLIPSEElement) CustomDataRemove(key string) *SVGELLIPSEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The x-axis coordinate of the center of the ellipse.
func (e *SVGELLIPSEElement) CX(f float64) *SVGELLIPSEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("cx", f)
	return e
}

func (e *SVGELLIPSEElement) IfCX(condition bool, f float64) *SVGELLIPSEElement {
	if condition {
		e.CX(f)
	}
	return e
}

// The y-axis coordinate of the center of the ellipse.
func (e *SVGELLIPSEElement) CY(f float64) *SVGELLIPSEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("cy", f)
	return e
}

func (e *SVGELLIPSEElement) IfCY(condition bool, f float64) *SVGELLIPSEElement {
	if condition {
		e.CY(f)
	}
	return e
}

// The x-axis radius of the ellipse.
func (e *SVGELLIPSEElement) RX(f float64) *SVGELLIPSEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("rx", f)
	return e
}

func (e *SVGELLIPSEElement) IfRX(condition bool, f float64) *SVGELLIPSEElement {
	if condition {
		e.RX(f)
	}
	return e
}

// The y-axis radius of the ellipse.
func (e *SVGELLIPSEElement) RY(f float64) *SVGELLIPSEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("ry", f)
	return e
}

func (e *SVGELLIPSEElement) IfRY(condition bool, f float64) *SVGELLIPSEElement {
	if condition {
		e.RY(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGELLIPSEElement) ID(s string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGELLIPSEElement) IDF(format string, args ...any) *SVGELLIPSEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGELLIPSEElement) IfID(condition bool, s string) *SVGELLIPSEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGELLIPSEElement) IfIDF(condition bool, format string, args ...any) *SVGELLIPSEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGELLIPSEElement) IDRemove(s string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGELLIPSEElement) IDRemoveF(format string, args ...any) *SVGELLIPSEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGELLIPSEElement) CLASS(s ...string) *SVGELLIPSEElement {
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

func (e *SVGELLIPSEElement) IfCLASS(condition bool, s ...string) *SVGELLIPSEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGELLIPSEElement) CLASSRemove(s ...string) *SVGELLIPSEElement {
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
func (e *SVGELLIPSEElement) STYLEF(k string, format string, args ...any) *SVGELLIPSEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGELLIPSEElement) IfSTYLE(condition bool, k string, v string) *SVGELLIPSEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGELLIPSEElement) STYLE(k string, v string) *SVGELLIPSEElement {
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

func (e *SVGELLIPSEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGELLIPSEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGELLIPSEElement) STYLEMap(m map[string]string) *SVGELLIPSEElement {
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
func (e *SVGELLIPSEElement) STYLEPairs(pairs ...string) *SVGELLIPSEElement {
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

func (e *SVGELLIPSEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGELLIPSEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGELLIPSEElement) STYLERemove(keys ...string) *SVGELLIPSEElement {
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

func (e *SVGELLIPSEElement) Z_REQ(expression string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGELLIPSEElement) IfZ_REQ(condition bool, expression string) *SVGELLIPSEElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGELLIPSEElement) Z_REQRemove() *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGELLIPSEElement) Z_TARGET(expression string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGELLIPSEElement) IfZ_TARGET(condition bool, expression string) *SVGELLIPSEElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGELLIPSEElement) Z_TARGETRemove() *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGELLIPSEElement) Z_REQ_SELECTOR(expression string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGELLIPSEElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGELLIPSEElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGELLIPSEElement) Z_REQ_SELECTORRemove() *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGELLIPSEElement) Z_SWAP(expression string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGELLIPSEElement) IfZ_SWAP(condition bool, expression string) *SVGELLIPSEElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGELLIPSEElement) Z_SWAPRemove() *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGELLIPSEElement) Z_SWAP_PUSH(expression string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGELLIPSEElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGELLIPSEElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGELLIPSEElement) Z_SWAP_PUSHRemove() *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGELLIPSEElement) Z_TRIGGER(expression string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGELLIPSEElement) IfZ_TRIGGER(condition bool, expression string) *SVGELLIPSEElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGELLIPSEElement) Z_TRIGGERRemove() *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGELLIPSEElement) Z_REQ_METHOD(c SVGEllipseZReqMethodChoice) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGEllipseZReqMethodChoice string

const (
	// default GET
	SVGEllipseZReqMethod_empty SVGEllipseZReqMethodChoice = ""
	// GET
	SVGEllipseZReqMethod_get SVGEllipseZReqMethodChoice = "get"
	// POST
	SVGEllipseZReqMethod_post SVGEllipseZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGELLIPSEElement) Z_REQ_METHODRemove(c SVGEllipseZReqMethodChoice) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGELLIPSEElement) Z_REQ_STRATEGY(expression string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGELLIPSEElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGELLIPSEElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGELLIPSEElement) Z_REQ_STRATEGYRemove() *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGELLIPSEElement) Z_REQ_HISTORY(expression string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGELLIPSEElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGELLIPSEElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGELLIPSEElement) Z_REQ_HISTORYRemove() *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGELLIPSEElement) Z_DATA(expression string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGELLIPSEElement) IfZ_DATA(condition bool, expression string) *SVGELLIPSEElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGELLIPSEElement) Z_DATARemove() *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGELLIPSEElement) Z_JSON(expression string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGELLIPSEElement) IfZ_JSON(condition bool, expression string) *SVGELLIPSEElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGELLIPSEElement) Z_JSONRemove() *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGELLIPSEElement) Z_REQ_BATCH(expression string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGELLIPSEElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGELLIPSEElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGELLIPSEElement) Z_REQ_BATCHRemove() *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGELLIPSEElement) Z_ACTION(expression string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGELLIPSEElement) IfZ_ACTION(condition bool, expression string) *SVGELLIPSEElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGELLIPSEElement) Z_ACTIONRemove() *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGELLIPSEElement) Z_REQ_BEFORE(expression string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGELLIPSEElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGELLIPSEElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGELLIPSEElement) Z_REQ_BEFORERemove() *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGELLIPSEElement) Z_REQ_AFTER(expression string) *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGELLIPSEElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGELLIPSEElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGELLIPSEElement) Z_REQ_AFTERRemove() *SVGELLIPSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
