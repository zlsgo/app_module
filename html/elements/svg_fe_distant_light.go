package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feDistantLight> SVG filter primitive defines a distant light source that
// can be used within a lighting filter primitive: <feDiffuseLighting> or
// <feSpecularLighting>
type SVGFEDISTANTLIGHTElement struct {
	*Element
}

// Create a new SVGFEDISTANTLIGHTElement element.
// This will create a new element with the tag
// "feDistantLight" during rendering.
func SVG_FEDISTANTLIGHT(children ...ElementRenderer) *SVGFEDISTANTLIGHTElement {
	e := NewElement("feDistantLight", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEDISTANTLIGHTElement{Element: e}
}

func (e *SVGFEDISTANTLIGHTElement) Children(children ...ElementRenderer) *SVGFEDISTANTLIGHTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) Attr(name string, value ...string) *SVGFEDISTANTLIGHTElement {
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

func (e *SVGFEDISTANTLIGHTElement) Attrs(attrs ...string) *SVGFEDISTANTLIGHTElement {
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

func (e *SVGFEDISTANTLIGHTElement) AttrsMap(attrs map[string]string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) Text(text string) *SVGFEDISTANTLIGHTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEDISTANTLIGHTElement) TextF(format string, args ...any) *SVGFEDISTANTLIGHTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEDISTANTLIGHTElement) IfText(condition bool, text string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfTextF(condition bool, format string, args ...any) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) Escaped(text string) *SVGFEDISTANTLIGHTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfEscaped(condition bool, text string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) EscapedF(format string, args ...any) *SVGFEDISTANTLIGHTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEDISTANTLIGHTElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) CustomData(key, value string) *SVGFEDISTANTLIGHTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfCustomData(condition bool, key, value string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) CustomDataF(key, format string, args ...any) *SVGFEDISTANTLIGHTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEDISTANTLIGHTElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) CustomDataRemove(key string) *SVGFEDISTANTLIGHTElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The azimuth attribute represent the direction vector of the light source in the
// XY plane (clockwise), in degrees from the x axis.
func (e *SVGFEDISTANTLIGHTElement) AZIMUTH(f float64) *SVGFEDISTANTLIGHTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("azimuth", f)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfAZIMUTH(condition bool, f float64) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.AZIMUTH(f)
	}
	return e
}

// The elevation attribute represent the direction vector of the light source
// perpendicular to the XY plane, in degrees from the XY plane towards the z axis
// (clockwise).
func (e *SVGFEDISTANTLIGHTElement) ELEVATION(f float64) *SVGFEDISTANTLIGHTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("elevation", f)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfELEVATION(condition bool, f float64) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.ELEVATION(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFEDISTANTLIGHTElement) ID(s string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IDF(format string, args ...any) *SVGFEDISTANTLIGHTElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEDISTANTLIGHTElement) IfID(condition bool, s string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfIDF(condition bool, format string, args ...any) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEDISTANTLIGHTElement) IDRemove(s string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IDRemoveF(format string, args ...any) *SVGFEDISTANTLIGHTElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEDISTANTLIGHTElement) CLASS(s ...string) *SVGFEDISTANTLIGHTElement {
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

func (e *SVGFEDISTANTLIGHTElement) IfCLASS(condition bool, s ...string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEDISTANTLIGHTElement) CLASSRemove(s ...string) *SVGFEDISTANTLIGHTElement {
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
func (e *SVGFEDISTANTLIGHTElement) STYLEF(k string, format string, args ...any) *SVGFEDISTANTLIGHTElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEDISTANTLIGHTElement) IfSTYLE(condition bool, k string, v string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) STYLE(k string, v string) *SVGFEDISTANTLIGHTElement {
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

func (e *SVGFEDISTANTLIGHTElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEDISTANTLIGHTElement) STYLEMap(m map[string]string) *SVGFEDISTANTLIGHTElement {
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
func (e *SVGFEDISTANTLIGHTElement) STYLEPairs(pairs ...string) *SVGFEDISTANTLIGHTElement {
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

func (e *SVGFEDISTANTLIGHTElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEDISTANTLIGHTElement) STYLERemove(keys ...string) *SVGFEDISTANTLIGHTElement {
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

func (e *SVGFEDISTANTLIGHTElement) Z_REQ(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfZ_REQ(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEDISTANTLIGHTElement) Z_REQRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEDISTANTLIGHTElement) Z_TARGET(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfZ_TARGET(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEDISTANTLIGHTElement) Z_TARGETRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEDISTANTLIGHTElement) Z_REQ_SELECTOR(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEDISTANTLIGHTElement) Z_REQ_SELECTORRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEDISTANTLIGHTElement) Z_SWAP(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfZ_SWAP(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEDISTANTLIGHTElement) Z_SWAPRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEDISTANTLIGHTElement) Z_SWAP_PUSH(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEDISTANTLIGHTElement) Z_SWAP_PUSHRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEDISTANTLIGHTElement) Z_TRIGGER(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEDISTANTLIGHTElement) Z_TRIGGERRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEDISTANTLIGHTElement) Z_REQ_METHOD(c SVGFeDistantLightZReqMethodChoice) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeDistantLightZReqMethodChoice string

const (
	// default GET
	SVGFeDistantLightZReqMethod_empty SVGFeDistantLightZReqMethodChoice = ""
	// GET
	SVGFeDistantLightZReqMethod_get SVGFeDistantLightZReqMethodChoice = "get"
	// POST
	SVGFeDistantLightZReqMethod_post SVGFeDistantLightZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEDISTANTLIGHTElement) Z_REQ_METHODRemove(c SVGFeDistantLightZReqMethodChoice) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEDISTANTLIGHTElement) Z_REQ_STRATEGY(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEDISTANTLIGHTElement) Z_REQ_STRATEGYRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEDISTANTLIGHTElement) Z_REQ_HISTORY(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEDISTANTLIGHTElement) Z_REQ_HISTORYRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEDISTANTLIGHTElement) Z_DATA(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfZ_DATA(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEDISTANTLIGHTElement) Z_DATARemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEDISTANTLIGHTElement) Z_JSON(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfZ_JSON(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEDISTANTLIGHTElement) Z_JSONRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEDISTANTLIGHTElement) Z_REQ_BATCH(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEDISTANTLIGHTElement) Z_REQ_BATCHRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEDISTANTLIGHTElement) Z_ACTION(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfZ_ACTION(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEDISTANTLIGHTElement) Z_ACTIONRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEDISTANTLIGHTElement) Z_REQ_BEFORE(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEDISTANTLIGHTElement) Z_REQ_BEFORERemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEDISTANTLIGHTElement) Z_REQ_AFTER(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEDISTANTLIGHTElement) Z_REQ_AFTERRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
