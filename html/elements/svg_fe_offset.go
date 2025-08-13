package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feOffset> SVG filter primitive allows to offset the input image
// The amount of offset can be controlled by attributes dx and dy.
type SVGFEOFFSETElement struct {
	*Element
}

// Create a new SVGFEOFFSETElement element.
// This will create a new element with the tag
// "feOffset" during rendering.
func SVG_FEOFFSET(children ...ElementRenderer) *SVGFEOFFSETElement {
	e := NewElement("feOffset", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEOFFSETElement{Element: e}
}

func (e *SVGFEOFFSETElement) Children(children ...ElementRenderer) *SVGFEOFFSETElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEOFFSETElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEOFFSETElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEOFFSETElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEOFFSETElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEOFFSETElement) Attr(name string, value ...string) *SVGFEOFFSETElement {
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

func (e *SVGFEOFFSETElement) Attrs(attrs ...string) *SVGFEOFFSETElement {
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

func (e *SVGFEOFFSETElement) AttrsMap(attrs map[string]string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEOFFSETElement) Text(text string) *SVGFEOFFSETElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEOFFSETElement) TextF(format string, args ...any) *SVGFEOFFSETElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEOFFSETElement) IfText(condition bool, text string) *SVGFEOFFSETElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEOFFSETElement) IfTextF(condition bool, format string, args ...any) *SVGFEOFFSETElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEOFFSETElement) Escaped(text string) *SVGFEOFFSETElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEOFFSETElement) IfEscaped(condition bool, text string) *SVGFEOFFSETElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEOFFSETElement) EscapedF(format string, args ...any) *SVGFEOFFSETElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEOFFSETElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEOFFSETElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEOFFSETElement) CustomData(key, value string) *SVGFEOFFSETElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEOFFSETElement) IfCustomData(condition bool, key, value string) *SVGFEOFFSETElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEOFFSETElement) CustomDataF(key, format string, args ...any) *SVGFEOFFSETElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEOFFSETElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEOFFSETElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEOFFSETElement) CustomDataRemove(key string) *SVGFEOFFSETElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The dx attribute indicates a shift along the x-axis on the kernel matrix.
func (e *SVGFEOFFSETElement) DX(f float64) *SVGFEOFFSETElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("dx", f)
	return e
}

func (e *SVGFEOFFSETElement) IfDX(condition bool, f float64) *SVGFEOFFSETElement {
	if condition {
		e.DX(f)
	}
	return e
}

// The dy attribute indicates a shift along the y-axis on the kernel matrix.
func (e *SVGFEOFFSETElement) DY(f float64) *SVGFEOFFSETElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("dy", f)
	return e
}

func (e *SVGFEOFFSETElement) IfDY(condition bool, f float64) *SVGFEOFFSETElement {
	if condition {
		e.DY(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFEOFFSETElement) ID(s string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEOFFSETElement) IDF(format string, args ...any) *SVGFEOFFSETElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEOFFSETElement) IfID(condition bool, s string) *SVGFEOFFSETElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEOFFSETElement) IfIDF(condition bool, format string, args ...any) *SVGFEOFFSETElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEOFFSETElement) IDRemove(s string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEOFFSETElement) IDRemoveF(format string, args ...any) *SVGFEOFFSETElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEOFFSETElement) CLASS(s ...string) *SVGFEOFFSETElement {
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

func (e *SVGFEOFFSETElement) IfCLASS(condition bool, s ...string) *SVGFEOFFSETElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEOFFSETElement) CLASSRemove(s ...string) *SVGFEOFFSETElement {
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
func (e *SVGFEOFFSETElement) STYLEF(k string, format string, args ...any) *SVGFEOFFSETElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEOFFSETElement) IfSTYLE(condition bool, k string, v string) *SVGFEOFFSETElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEOFFSETElement) STYLE(k string, v string) *SVGFEOFFSETElement {
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

func (e *SVGFEOFFSETElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEOFFSETElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEOFFSETElement) STYLEMap(m map[string]string) *SVGFEOFFSETElement {
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
func (e *SVGFEOFFSETElement) STYLEPairs(pairs ...string) *SVGFEOFFSETElement {
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

func (e *SVGFEOFFSETElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEOFFSETElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEOFFSETElement) STYLERemove(keys ...string) *SVGFEOFFSETElement {
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

func (e *SVGFEOFFSETElement) Z_REQ(expression string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEOFFSETElement) IfZ_REQ(condition bool, expression string) *SVGFEOFFSETElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEOFFSETElement) Z_REQRemove() *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEOFFSETElement) Z_TARGET(expression string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEOFFSETElement) IfZ_TARGET(condition bool, expression string) *SVGFEOFFSETElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEOFFSETElement) Z_TARGETRemove() *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEOFFSETElement) Z_REQ_SELECTOR(expression string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEOFFSETElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEOFFSETElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEOFFSETElement) Z_REQ_SELECTORRemove() *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEOFFSETElement) Z_SWAP(expression string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEOFFSETElement) IfZ_SWAP(condition bool, expression string) *SVGFEOFFSETElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEOFFSETElement) Z_SWAPRemove() *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEOFFSETElement) Z_SWAP_PUSH(expression string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEOFFSETElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEOFFSETElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEOFFSETElement) Z_SWAP_PUSHRemove() *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEOFFSETElement) Z_TRIGGER(expression string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEOFFSETElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEOFFSETElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEOFFSETElement) Z_TRIGGERRemove() *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEOFFSETElement) Z_REQ_METHOD(c SVGFeOffsetZReqMethodChoice) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeOffsetZReqMethodChoice string

const (
	// default GET
	SVGFeOffsetZReqMethod_empty SVGFeOffsetZReqMethodChoice = ""
	// GET
	SVGFeOffsetZReqMethod_get SVGFeOffsetZReqMethodChoice = "get"
	// POST
	SVGFeOffsetZReqMethod_post SVGFeOffsetZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEOFFSETElement) Z_REQ_METHODRemove(c SVGFeOffsetZReqMethodChoice) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEOFFSETElement) Z_REQ_STRATEGY(expression string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEOFFSETElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEOFFSETElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEOFFSETElement) Z_REQ_STRATEGYRemove() *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEOFFSETElement) Z_REQ_HISTORY(expression string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEOFFSETElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEOFFSETElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEOFFSETElement) Z_REQ_HISTORYRemove() *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEOFFSETElement) Z_DATA(expression string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEOFFSETElement) IfZ_DATA(condition bool, expression string) *SVGFEOFFSETElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEOFFSETElement) Z_DATARemove() *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEOFFSETElement) Z_JSON(expression string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEOFFSETElement) IfZ_JSON(condition bool, expression string) *SVGFEOFFSETElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEOFFSETElement) Z_JSONRemove() *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEOFFSETElement) Z_REQ_BATCH(expression string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEOFFSETElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEOFFSETElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEOFFSETElement) Z_REQ_BATCHRemove() *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEOFFSETElement) Z_ACTION(expression string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEOFFSETElement) IfZ_ACTION(condition bool, expression string) *SVGFEOFFSETElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEOFFSETElement) Z_ACTIONRemove() *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEOFFSETElement) Z_REQ_BEFORE(expression string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEOFFSETElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEOFFSETElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEOFFSETElement) Z_REQ_BEFORERemove() *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEOFFSETElement) Z_REQ_AFTER(expression string) *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEOFFSETElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEOFFSETElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEOFFSETElement) Z_REQ_AFTERRemove() *SVGFEOFFSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
