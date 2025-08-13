package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <tspan> SVG element lets authors explicitly specify the location of a glyph
// along the given path via the attributes.
type SVGTSPANElement struct {
	*Element
}

// Create a new SVGTSPANElement element.
// This will create a new element with the tag
// "tspan" during rendering.
func SVG_TSPAN(children ...ElementRenderer) *SVGTSPANElement {
	e := NewElement("tspan", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGTSPANElement{Element: e}
}

func (e *SVGTSPANElement) Children(children ...ElementRenderer) *SVGTSPANElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGTSPANElement) IfChildren(condition bool, children ...ElementRenderer) *SVGTSPANElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGTSPANElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGTSPANElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGTSPANElement) Attr(name string, value ...string) *SVGTSPANElement {
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

func (e *SVGTSPANElement) Attrs(attrs ...string) *SVGTSPANElement {
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

func (e *SVGTSPANElement) AttrsMap(attrs map[string]string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGTSPANElement) Text(text string) *SVGTSPANElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGTSPANElement) TextF(format string, args ...any) *SVGTSPANElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGTSPANElement) IfText(condition bool, text string) *SVGTSPANElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGTSPANElement) IfTextF(condition bool, format string, args ...any) *SVGTSPANElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGTSPANElement) Escaped(text string) *SVGTSPANElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGTSPANElement) IfEscaped(condition bool, text string) *SVGTSPANElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGTSPANElement) EscapedF(format string, args ...any) *SVGTSPANElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGTSPANElement) IfEscapedF(condition bool, format string, args ...any) *SVGTSPANElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGTSPANElement) CustomData(key, value string) *SVGTSPANElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGTSPANElement) IfCustomData(condition bool, key, value string) *SVGTSPANElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGTSPANElement) CustomDataF(key, format string, args ...any) *SVGTSPANElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGTSPANElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGTSPANElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGTSPANElement) CustomDataRemove(key string) *SVGTSPANElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The x-axis coordinate of the current text position.
func (e *SVGTSPANElement) X(f float64) *SVGTSPANElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("x", f)
	return e
}

func (e *SVGTSPANElement) IfX(condition bool, f float64) *SVGTSPANElement {
	if condition {
		e.X(f)
	}
	return e
}

// The y-axis coordinate of the current text position.
func (e *SVGTSPANElement) Y(f float64) *SVGTSPANElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("y", f)
	return e
}

func (e *SVGTSPANElement) IfY(condition bool, f float64) *SVGTSPANElement {
	if condition {
		e.Y(f)
	}
	return e
}

// The x-axis coordinate of the current text position.
func (e *SVGTSPANElement) DX(f float64) *SVGTSPANElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("dx", f)
	return e
}

func (e *SVGTSPANElement) IfDX(condition bool, f float64) *SVGTSPANElement {
	if condition {
		e.DX(f)
	}
	return e
}

// The y-axis coordinate of the current text position.
func (e *SVGTSPANElement) DY(f float64) *SVGTSPANElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("dy", f)
	return e
}

func (e *SVGTSPANElement) IfDY(condition bool, f float64) *SVGTSPANElement {
	if condition {
		e.DY(f)
	}
	return e
}

// The rotation angle about the current text position.
func (e *SVGTSPANElement) ROTATE(f float64) *SVGTSPANElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("rotate", f)
	return e
}

func (e *SVGTSPANElement) IfROTATE(condition bool, f float64) *SVGTSPANElement {
	if condition {
		e.ROTATE(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGTSPANElement) ID(s string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGTSPANElement) IDF(format string, args ...any) *SVGTSPANElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGTSPANElement) IfID(condition bool, s string) *SVGTSPANElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGTSPANElement) IfIDF(condition bool, format string, args ...any) *SVGTSPANElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGTSPANElement) IDRemove(s string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGTSPANElement) IDRemoveF(format string, args ...any) *SVGTSPANElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGTSPANElement) CLASS(s ...string) *SVGTSPANElement {
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

func (e *SVGTSPANElement) IfCLASS(condition bool, s ...string) *SVGTSPANElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGTSPANElement) CLASSRemove(s ...string) *SVGTSPANElement {
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
func (e *SVGTSPANElement) STYLEF(k string, format string, args ...any) *SVGTSPANElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGTSPANElement) IfSTYLE(condition bool, k string, v string) *SVGTSPANElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGTSPANElement) STYLE(k string, v string) *SVGTSPANElement {
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

func (e *SVGTSPANElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGTSPANElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGTSPANElement) STYLEMap(m map[string]string) *SVGTSPANElement {
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
func (e *SVGTSPANElement) STYLEPairs(pairs ...string) *SVGTSPANElement {
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

func (e *SVGTSPANElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGTSPANElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGTSPANElement) STYLERemove(keys ...string) *SVGTSPANElement {
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

func (e *SVGTSPANElement) Z_REQ(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfZ_REQ(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGTSPANElement) Z_REQRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGTSPANElement) Z_TARGET(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfZ_TARGET(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGTSPANElement) Z_TARGETRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGTSPANElement) Z_REQ_SELECTOR(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGTSPANElement) Z_REQ_SELECTORRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGTSPANElement) Z_SWAP(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfZ_SWAP(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGTSPANElement) Z_SWAPRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGTSPANElement) Z_SWAP_PUSH(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGTSPANElement) Z_SWAP_PUSHRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGTSPANElement) Z_TRIGGER(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfZ_TRIGGER(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGTSPANElement) Z_TRIGGERRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGTSPANElement) Z_REQ_METHOD(c SVGTspanZReqMethodChoice) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGTspanZReqMethodChoice string

const (
	// default GET
	SVGTspanZReqMethod_empty SVGTspanZReqMethodChoice = ""
	// GET
	SVGTspanZReqMethod_get SVGTspanZReqMethodChoice = "get"
	// POST
	SVGTspanZReqMethod_post SVGTspanZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGTSPANElement) Z_REQ_METHODRemove(c SVGTspanZReqMethodChoice) *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGTSPANElement) Z_REQ_STRATEGY(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGTSPANElement) Z_REQ_STRATEGYRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGTSPANElement) Z_REQ_HISTORY(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGTSPANElement) Z_REQ_HISTORYRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGTSPANElement) Z_DATA(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfZ_DATA(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGTSPANElement) Z_DATARemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGTSPANElement) Z_JSON(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfZ_JSON(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGTSPANElement) Z_JSONRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGTSPANElement) Z_REQ_BATCH(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGTSPANElement) Z_REQ_BATCHRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGTSPANElement) Z_ACTION(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfZ_ACTION(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGTSPANElement) Z_ACTIONRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGTSPANElement) Z_REQ_BEFORE(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGTSPANElement) Z_REQ_BEFORERemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGTSPANElement) Z_REQ_AFTER(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGTSPANElement) Z_REQ_AFTERRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
