package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feFlood> SVG filter primitive fills the filter subregion with the color
// and opacity defined by flood-color and flood-opacity.
type SVGFEFLOODElement struct {
	*Element
}

// Create a new SVGFEFLOODElement element.
// This will create a new element with the tag
// "feFlood" during rendering.
func SVG_FEFLOOD(children ...ElementRenderer) *SVGFEFLOODElement {
	e := NewElement("feFlood", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEFLOODElement{Element: e}
}

func (e *SVGFEFLOODElement) Children(children ...ElementRenderer) *SVGFEFLOODElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEFLOODElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEFLOODElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEFLOODElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEFLOODElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEFLOODElement) Attr(name string, value ...string) *SVGFEFLOODElement {
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

func (e *SVGFEFLOODElement) Attrs(attrs ...string) *SVGFEFLOODElement {
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

func (e *SVGFEFLOODElement) AttrsMap(attrs map[string]string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEFLOODElement) Text(text string) *SVGFEFLOODElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEFLOODElement) TextF(format string, args ...any) *SVGFEFLOODElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEFLOODElement) IfText(condition bool, text string) *SVGFEFLOODElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEFLOODElement) IfTextF(condition bool, format string, args ...any) *SVGFEFLOODElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEFLOODElement) Escaped(text string) *SVGFEFLOODElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEFLOODElement) IfEscaped(condition bool, text string) *SVGFEFLOODElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEFLOODElement) EscapedF(format string, args ...any) *SVGFEFLOODElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEFLOODElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEFLOODElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEFLOODElement) CustomData(key, value string) *SVGFEFLOODElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEFLOODElement) IfCustomData(condition bool, key, value string) *SVGFEFLOODElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEFLOODElement) CustomDataF(key, format string, args ...any) *SVGFEFLOODElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEFLOODElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEFLOODElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEFLOODElement) CustomDataRemove(key string) *SVGFEFLOODElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The flood-color attribute indicates what color to use to flood the current
// filter primitive subregion defined through the <feFlood> element
// If attribute 'flood-color' is not specified, then the effect is as if a value
// of black were specified.
func (e *SVGFEFLOODElement) FLOOD_COLOR(s string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("flood-color", s)
	return e
}

func (e *SVGFEFLOODElement) FLOOD_COLORF(format string, args ...any) *SVGFEFLOODElement {
	return e.FLOOD_COLOR(fmt.Sprintf(format, args...))
}

func (e *SVGFEFLOODElement) IfFLOOD_COLOR(condition bool, s string) *SVGFEFLOODElement {
	if condition {
		e.FLOOD_COLOR(s)
	}
	return e
}

func (e *SVGFEFLOODElement) IfFLOOD_COLORF(condition bool, format string, args ...any) *SVGFEFLOODElement {
	if condition {
		e.FLOOD_COLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute FLOOD_COLOR from the element.
func (e *SVGFEFLOODElement) FLOOD_COLORRemove(s string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("flood-color")
	return e
}

func (e *SVGFEFLOODElement) FLOOD_COLORRemoveF(format string, args ...any) *SVGFEFLOODElement {
	return e.FLOOD_COLORRemove(fmt.Sprintf(format, args...))
}

// The flood-opacity attribute indicates the opacity value to use across the
// current filter primitive subregion defined through the <feFlood> element.
func (e *SVGFEFLOODElement) FLOOD_OPACITY(f float64) *SVGFEFLOODElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("flood-opacity", f)
	return e
}

func (e *SVGFEFLOODElement) IfFLOOD_OPACITY(condition bool, f float64) *SVGFEFLOODElement {
	if condition {
		e.FLOOD_OPACITY(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFEFLOODElement) ID(s string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEFLOODElement) IDF(format string, args ...any) *SVGFEFLOODElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEFLOODElement) IfID(condition bool, s string) *SVGFEFLOODElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEFLOODElement) IfIDF(condition bool, format string, args ...any) *SVGFEFLOODElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEFLOODElement) IDRemove(s string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEFLOODElement) IDRemoveF(format string, args ...any) *SVGFEFLOODElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEFLOODElement) CLASS(s ...string) *SVGFEFLOODElement {
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

func (e *SVGFEFLOODElement) IfCLASS(condition bool, s ...string) *SVGFEFLOODElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEFLOODElement) CLASSRemove(s ...string) *SVGFEFLOODElement {
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
func (e *SVGFEFLOODElement) STYLEF(k string, format string, args ...any) *SVGFEFLOODElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEFLOODElement) IfSTYLE(condition bool, k string, v string) *SVGFEFLOODElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEFLOODElement) STYLE(k string, v string) *SVGFEFLOODElement {
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

func (e *SVGFEFLOODElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEFLOODElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEFLOODElement) STYLEMap(m map[string]string) *SVGFEFLOODElement {
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
func (e *SVGFEFLOODElement) STYLEPairs(pairs ...string) *SVGFEFLOODElement {
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

func (e *SVGFEFLOODElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEFLOODElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEFLOODElement) STYLERemove(keys ...string) *SVGFEFLOODElement {
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

func (e *SVGFEFLOODElement) Z_REQ(expression string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFLOODElement) IfZ_REQ(condition bool, expression string) *SVGFEFLOODElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEFLOODElement) Z_REQRemove() *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEFLOODElement) Z_TARGET(expression string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFLOODElement) IfZ_TARGET(condition bool, expression string) *SVGFEFLOODElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEFLOODElement) Z_TARGETRemove() *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEFLOODElement) Z_REQ_SELECTOR(expression string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFLOODElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEFLOODElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEFLOODElement) Z_REQ_SELECTORRemove() *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEFLOODElement) Z_SWAP(expression string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFLOODElement) IfZ_SWAP(condition bool, expression string) *SVGFEFLOODElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEFLOODElement) Z_SWAPRemove() *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEFLOODElement) Z_SWAP_PUSH(expression string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFLOODElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEFLOODElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEFLOODElement) Z_SWAP_PUSHRemove() *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEFLOODElement) Z_TRIGGER(expression string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFLOODElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEFLOODElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEFLOODElement) Z_TRIGGERRemove() *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEFLOODElement) Z_REQ_METHOD(c SVGFeFloodZReqMethodChoice) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeFloodZReqMethodChoice string

const (
	// default GET
	SVGFeFloodZReqMethod_empty SVGFeFloodZReqMethodChoice = ""
	// GET
	SVGFeFloodZReqMethod_get SVGFeFloodZReqMethodChoice = "get"
	// POST
	SVGFeFloodZReqMethod_post SVGFeFloodZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEFLOODElement) Z_REQ_METHODRemove(c SVGFeFloodZReqMethodChoice) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEFLOODElement) Z_REQ_STRATEGY(expression string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFLOODElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEFLOODElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEFLOODElement) Z_REQ_STRATEGYRemove() *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEFLOODElement) Z_REQ_HISTORY(expression string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFLOODElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEFLOODElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEFLOODElement) Z_REQ_HISTORYRemove() *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEFLOODElement) Z_DATA(expression string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFLOODElement) IfZ_DATA(condition bool, expression string) *SVGFEFLOODElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEFLOODElement) Z_DATARemove() *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEFLOODElement) Z_JSON(expression string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFLOODElement) IfZ_JSON(condition bool, expression string) *SVGFEFLOODElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEFLOODElement) Z_JSONRemove() *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEFLOODElement) Z_REQ_BATCH(expression string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFLOODElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEFLOODElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEFLOODElement) Z_REQ_BATCHRemove() *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEFLOODElement) Z_ACTION(expression string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFLOODElement) IfZ_ACTION(condition bool, expression string) *SVGFEFLOODElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEFLOODElement) Z_ACTIONRemove() *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEFLOODElement) Z_REQ_BEFORE(expression string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFLOODElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEFLOODElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEFLOODElement) Z_REQ_BEFORERemove() *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEFLOODElement) Z_REQ_AFTER(expression string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFLOODElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEFLOODElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEFLOODElement) Z_REQ_AFTERRemove() *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
