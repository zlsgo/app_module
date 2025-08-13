package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feBlend> SVG filter primitive composes two objects together ruled by a
// certain blending mode
// This is similar to what is known from image editing software when blending two
// layers
// The mode is defined by the mode attribute.
type SVGFEBLENDElement struct {
	*Element
}

// Create a new SVGFEBLENDElement element.
// This will create a new element with the tag
// "feBlend" during rendering.
func SVG_FEBLEND(children ...ElementRenderer) *SVGFEBLENDElement {
	e := NewElement("feBlend", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEBLENDElement{Element: e}
}

func (e *SVGFEBLENDElement) Children(children ...ElementRenderer) *SVGFEBLENDElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEBLENDElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEBLENDElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEBLENDElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEBLENDElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEBLENDElement) Attr(name string, value ...string) *SVGFEBLENDElement {
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

func (e *SVGFEBLENDElement) Attrs(attrs ...string) *SVGFEBLENDElement {
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

func (e *SVGFEBLENDElement) AttrsMap(attrs map[string]string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEBLENDElement) Text(text string) *SVGFEBLENDElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEBLENDElement) TextF(format string, args ...any) *SVGFEBLENDElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEBLENDElement) IfText(condition bool, text string) *SVGFEBLENDElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEBLENDElement) IfTextF(condition bool, format string, args ...any) *SVGFEBLENDElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEBLENDElement) Escaped(text string) *SVGFEBLENDElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEBLENDElement) IfEscaped(condition bool, text string) *SVGFEBLENDElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEBLENDElement) EscapedF(format string, args ...any) *SVGFEBLENDElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEBLENDElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEBLENDElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEBLENDElement) CustomData(key, value string) *SVGFEBLENDElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEBLENDElement) IfCustomData(condition bool, key, value string) *SVGFEBLENDElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEBLENDElement) CustomDataF(key, format string, args ...any) *SVGFEBLENDElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEBLENDElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEBLENDElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEBLENDElement) CustomDataRemove(key string) *SVGFEBLENDElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// Input for the blending.
func (e *SVGFEBLENDElement) IN(s string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFEBLENDElement) INF(format string, args ...any) *SVGFEBLENDElement {
	return e.IN(fmt.Sprintf(format, args...))
}

func (e *SVGFEBLENDElement) IfIN(condition bool, s string) *SVGFEBLENDElement {
	if condition {
		e.IN(s)
	}
	return e
}

func (e *SVGFEBLENDElement) IfINF(condition bool, format string, args ...any) *SVGFEBLENDElement {
	if condition {
		e.IN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFEBLENDElement) INRemove(s string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("in")
	return e
}

func (e *SVGFEBLENDElement) INRemoveF(format string, args ...any) *SVGFEBLENDElement {
	return e.INRemove(fmt.Sprintf(format, args...))
}

// Second input for the blending.
func (e *SVGFEBLENDElement) IN_2(s string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("in2", s)
	return e
}

func (e *SVGFEBLENDElement) IN_2F(format string, args ...any) *SVGFEBLENDElement {
	return e.IN_2(fmt.Sprintf(format, args...))
}

func (e *SVGFEBLENDElement) IfIN_2(condition bool, s string) *SVGFEBLENDElement {
	if condition {
		e.IN_2(s)
	}
	return e
}

func (e *SVGFEBLENDElement) IfIN_2F(condition bool, format string, args ...any) *SVGFEBLENDElement {
	if condition {
		e.IN_2(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN_2 from the element.
func (e *SVGFEBLENDElement) IN_2Remove(s string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("in2")
	return e
}

func (e *SVGFEBLENDElement) IN_2RemoveF(format string, args ...any) *SVGFEBLENDElement {
	return e.IN_2Remove(fmt.Sprintf(format, args...))
}

// The mode used to blend the two inputs together.
func (e *SVGFEBLENDElement) MODE(c SVGFeBlendModeChoice) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mode", string(c))
	return e
}

type SVGFeBlendModeChoice string

const (
	// The input image is placed over the backdrop image, then the parts of the input
	// image that are outside the backdrop are discarded.
	SVGFeBlendMode_normal SVGFeBlendModeChoice = "normal"
	// The input image is multiplied by the backdrop image.
	SVGFeBlendMode_multiply SVGFeBlendModeChoice = "multiply"
	// Multiplies the complements of the backdrop and input image color values, then
	// complements the result.
	SVGFeBlendMode_screen SVGFeBlendModeChoice = "screen"
	// Selects the darker of the backdrop and input image pixels.
	SVGFeBlendMode_darken SVGFeBlendModeChoice = "darken"
	// Selects the lighter of the backdrop and input image pixels.
	SVGFeBlendMode_lighten SVGFeBlendModeChoice = "lighten"
)

// Remove the attribute MODE from the element.
func (e *SVGFEBLENDElement) MODERemove(c SVGFeBlendModeChoice) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mode")
	return e
}

// Specifies a unique id for an element
func (e *SVGFEBLENDElement) ID(s string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEBLENDElement) IDF(format string, args ...any) *SVGFEBLENDElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEBLENDElement) IfID(condition bool, s string) *SVGFEBLENDElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEBLENDElement) IfIDF(condition bool, format string, args ...any) *SVGFEBLENDElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEBLENDElement) IDRemove(s string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEBLENDElement) IDRemoveF(format string, args ...any) *SVGFEBLENDElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEBLENDElement) CLASS(s ...string) *SVGFEBLENDElement {
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

func (e *SVGFEBLENDElement) IfCLASS(condition bool, s ...string) *SVGFEBLENDElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEBLENDElement) CLASSRemove(s ...string) *SVGFEBLENDElement {
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
func (e *SVGFEBLENDElement) STYLEF(k string, format string, args ...any) *SVGFEBLENDElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEBLENDElement) IfSTYLE(condition bool, k string, v string) *SVGFEBLENDElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEBLENDElement) STYLE(k string, v string) *SVGFEBLENDElement {
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

func (e *SVGFEBLENDElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEBLENDElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEBLENDElement) STYLEMap(m map[string]string) *SVGFEBLENDElement {
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
func (e *SVGFEBLENDElement) STYLEPairs(pairs ...string) *SVGFEBLENDElement {
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

func (e *SVGFEBLENDElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEBLENDElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEBLENDElement) STYLERemove(keys ...string) *SVGFEBLENDElement {
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

func (e *SVGFEBLENDElement) Z_REQ(expression string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEBLENDElement) IfZ_REQ(condition bool, expression string) *SVGFEBLENDElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEBLENDElement) Z_REQRemove() *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEBLENDElement) Z_TARGET(expression string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEBLENDElement) IfZ_TARGET(condition bool, expression string) *SVGFEBLENDElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEBLENDElement) Z_TARGETRemove() *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEBLENDElement) Z_REQ_SELECTOR(expression string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEBLENDElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEBLENDElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEBLENDElement) Z_REQ_SELECTORRemove() *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEBLENDElement) Z_SWAP(expression string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEBLENDElement) IfZ_SWAP(condition bool, expression string) *SVGFEBLENDElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEBLENDElement) Z_SWAPRemove() *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEBLENDElement) Z_SWAP_PUSH(expression string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEBLENDElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEBLENDElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEBLENDElement) Z_SWAP_PUSHRemove() *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEBLENDElement) Z_TRIGGER(expression string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEBLENDElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEBLENDElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEBLENDElement) Z_TRIGGERRemove() *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEBLENDElement) Z_REQ_METHOD(c SVGFeBlendZReqMethodChoice) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeBlendZReqMethodChoice string

const (
	// default GET
	SVGFeBlendZReqMethod_empty SVGFeBlendZReqMethodChoice = ""
	// GET
	SVGFeBlendZReqMethod_get SVGFeBlendZReqMethodChoice = "get"
	// POST
	SVGFeBlendZReqMethod_post SVGFeBlendZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEBLENDElement) Z_REQ_METHODRemove(c SVGFeBlendZReqMethodChoice) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEBLENDElement) Z_REQ_STRATEGY(expression string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEBLENDElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEBLENDElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEBLENDElement) Z_REQ_STRATEGYRemove() *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEBLENDElement) Z_REQ_HISTORY(expression string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEBLENDElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEBLENDElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEBLENDElement) Z_REQ_HISTORYRemove() *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEBLENDElement) Z_DATA(expression string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEBLENDElement) IfZ_DATA(condition bool, expression string) *SVGFEBLENDElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEBLENDElement) Z_DATARemove() *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEBLENDElement) Z_JSON(expression string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEBLENDElement) IfZ_JSON(condition bool, expression string) *SVGFEBLENDElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEBLENDElement) Z_JSONRemove() *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEBLENDElement) Z_REQ_BATCH(expression string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEBLENDElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEBLENDElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEBLENDElement) Z_REQ_BATCHRemove() *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEBLENDElement) Z_ACTION(expression string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEBLENDElement) IfZ_ACTION(condition bool, expression string) *SVGFEBLENDElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEBLENDElement) Z_ACTIONRemove() *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEBLENDElement) Z_REQ_BEFORE(expression string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEBLENDElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEBLENDElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEBLENDElement) Z_REQ_BEFORERemove() *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEBLENDElement) Z_REQ_AFTER(expression string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEBLENDElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEBLENDElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEBLENDElement) Z_REQ_AFTERRemove() *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
