package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <use> SVG element includes a reference to a <symbol> element and attempts
// to display the referenced content
// The reference is drawn exactly as it was defined
// It can be reused as often as needed and can be programmatically manipulated.
type SVGUSEElement struct {
	*Element
}

// Create a new SVGUSEElement element.
// This will create a new element with the tag
// "use" during rendering.
func SVG_USE(children ...ElementRenderer) *SVGUSEElement {
	e := NewElement("use", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGUSEElement{Element: e}
}

func (e *SVGUSEElement) Children(children ...ElementRenderer) *SVGUSEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGUSEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGUSEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGUSEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGUSEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGUSEElement) Attr(name string, value ...string) *SVGUSEElement {
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

func (e *SVGUSEElement) Attrs(attrs ...string) *SVGUSEElement {
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

func (e *SVGUSEElement) AttrsMap(attrs map[string]string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGUSEElement) Text(text string) *SVGUSEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGUSEElement) TextF(format string, args ...any) *SVGUSEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGUSEElement) IfText(condition bool, text string) *SVGUSEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGUSEElement) IfTextF(condition bool, format string, args ...any) *SVGUSEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGUSEElement) Escaped(text string) *SVGUSEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGUSEElement) IfEscaped(condition bool, text string) *SVGUSEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGUSEElement) EscapedF(format string, args ...any) *SVGUSEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGUSEElement) IfEscapedF(condition bool, format string, args ...any) *SVGUSEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGUSEElement) CustomData(key, value string) *SVGUSEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGUSEElement) IfCustomData(condition bool, key, value string) *SVGUSEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGUSEElement) CustomDataF(key, format string, args ...any) *SVGUSEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGUSEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGUSEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGUSEElement) CustomDataRemove(key string) *SVGUSEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// A URI reference to the symbol to use.
func (e *SVGUSEElement) HREF(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("href", s)
	return e
}

func (e *SVGUSEElement) HREFF(format string, args ...any) *SVGUSEElement {
	return e.HREF(fmt.Sprintf(format, args...))
}

func (e *SVGUSEElement) IfHREF(condition bool, s string) *SVGUSEElement {
	if condition {
		e.HREF(s)
	}
	return e
}

func (e *SVGUSEElement) IfHREFF(condition bool, format string, args ...any) *SVGUSEElement {
	if condition {
		e.HREF(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HREF from the element.
func (e *SVGUSEElement) HREFRemove(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("href")
	return e
}

func (e *SVGUSEElement) HREFRemoveF(format string, args ...any) *SVGUSEElement {
	return e.HREFRemove(fmt.Sprintf(format, args...))
}

// The x-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGUSEElement) X(f float64) *SVGUSEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("x", f)
	return e
}

func (e *SVGUSEElement) IfX(condition bool, f float64) *SVGUSEElement {
	if condition {
		e.X(f)
	}
	return e
}

// The y-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGUSEElement) Y(f float64) *SVGUSEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("y", f)
	return e
}

func (e *SVGUSEElement) IfY(condition bool, f float64) *SVGUSEElement {
	if condition {
		e.Y(f)
	}
	return e
}

// The width of the rectangular region.
func (e *SVGUSEElement) WIDTH(f float64) *SVGUSEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("width", f)
	return e
}

func (e *SVGUSEElement) IfWIDTH(condition bool, f float64) *SVGUSEElement {
	if condition {
		e.WIDTH(f)
	}
	return e
}

// The height of the rectangular region.
func (e *SVGUSEElement) HEIGHT(f float64) *SVGUSEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("height", f)
	return e
}

func (e *SVGUSEElement) IfHEIGHT(condition bool, f float64) *SVGUSEElement {
	if condition {
		e.HEIGHT(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGUSEElement) ID(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGUSEElement) IDF(format string, args ...any) *SVGUSEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGUSEElement) IfID(condition bool, s string) *SVGUSEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGUSEElement) IfIDF(condition bool, format string, args ...any) *SVGUSEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGUSEElement) IDRemove(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGUSEElement) IDRemoveF(format string, args ...any) *SVGUSEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGUSEElement) CLASS(s ...string) *SVGUSEElement {
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

func (e *SVGUSEElement) IfCLASS(condition bool, s ...string) *SVGUSEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGUSEElement) CLASSRemove(s ...string) *SVGUSEElement {
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
func (e *SVGUSEElement) STYLEF(k string, format string, args ...any) *SVGUSEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGUSEElement) IfSTYLE(condition bool, k string, v string) *SVGUSEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGUSEElement) STYLE(k string, v string) *SVGUSEElement {
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

func (e *SVGUSEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGUSEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGUSEElement) STYLEMap(m map[string]string) *SVGUSEElement {
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
func (e *SVGUSEElement) STYLEPairs(pairs ...string) *SVGUSEElement {
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

func (e *SVGUSEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGUSEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGUSEElement) STYLERemove(keys ...string) *SVGUSEElement {
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

func (e *SVGUSEElement) Z_REQ(expression string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGUSEElement) IfZ_REQ(condition bool, expression string) *SVGUSEElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGUSEElement) Z_REQRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGUSEElement) Z_TARGET(expression string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGUSEElement) IfZ_TARGET(condition bool, expression string) *SVGUSEElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGUSEElement) Z_TARGETRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGUSEElement) Z_REQ_SELECTOR(expression string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGUSEElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGUSEElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGUSEElement) Z_REQ_SELECTORRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGUSEElement) Z_SWAP(expression string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGUSEElement) IfZ_SWAP(condition bool, expression string) *SVGUSEElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGUSEElement) Z_SWAPRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGUSEElement) Z_SWAP_PUSH(expression string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGUSEElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGUSEElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGUSEElement) Z_SWAP_PUSHRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGUSEElement) Z_TRIGGER(expression string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGUSEElement) IfZ_TRIGGER(condition bool, expression string) *SVGUSEElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGUSEElement) Z_TRIGGERRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGUSEElement) Z_REQ_METHOD(c SVGUseZReqMethodChoice) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGUseZReqMethodChoice string

const (
	// default GET
	SVGUseZReqMethod_empty SVGUseZReqMethodChoice = ""
	// GET
	SVGUseZReqMethod_get SVGUseZReqMethodChoice = "get"
	// POST
	SVGUseZReqMethod_post SVGUseZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGUSEElement) Z_REQ_METHODRemove(c SVGUseZReqMethodChoice) *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGUSEElement) Z_REQ_STRATEGY(expression string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGUSEElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGUSEElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGUSEElement) Z_REQ_STRATEGYRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGUSEElement) Z_REQ_HISTORY(expression string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGUSEElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGUSEElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGUSEElement) Z_REQ_HISTORYRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGUSEElement) Z_DATA(expression string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGUSEElement) IfZ_DATA(condition bool, expression string) *SVGUSEElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGUSEElement) Z_DATARemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGUSEElement) Z_JSON(expression string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGUSEElement) IfZ_JSON(condition bool, expression string) *SVGUSEElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGUSEElement) Z_JSONRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGUSEElement) Z_REQ_BATCH(expression string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGUSEElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGUSEElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGUSEElement) Z_REQ_BATCHRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGUSEElement) Z_ACTION(expression string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGUSEElement) IfZ_ACTION(condition bool, expression string) *SVGUSEElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGUSEElement) Z_ACTIONRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGUSEElement) Z_REQ_BEFORE(expression string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGUSEElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGUSEElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGUSEElement) Z_REQ_BEFORERemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGUSEElement) Z_REQ_AFTER(expression string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGUSEElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGUSEElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGUSEElement) Z_REQ_AFTERRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
