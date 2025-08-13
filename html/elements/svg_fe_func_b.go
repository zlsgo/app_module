package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feFuncB> SVG filter primitive defines the transfer function for the blue
// component of the input graphic of its parent <feComponentTransfer> element.
type SVGFEFUNCBElement struct {
	*Element
}

// Create a new SVGFEFUNCBElement element.
// This will create a new element with the tag
// "feFuncB" during rendering.
func SVG_FEFUNCB(children ...ElementRenderer) *SVGFEFUNCBElement {
	e := NewElement("feFuncB", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEFUNCBElement{Element: e}
}

func (e *SVGFEFUNCBElement) Children(children ...ElementRenderer) *SVGFEFUNCBElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEFUNCBElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEFUNCBElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEFUNCBElement) Attr(name string, value ...string) *SVGFEFUNCBElement {
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

func (e *SVGFEFUNCBElement) Attrs(attrs ...string) *SVGFEFUNCBElement {
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

func (e *SVGFEFUNCBElement) AttrsMap(attrs map[string]string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEFUNCBElement) Text(text string) *SVGFEFUNCBElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEFUNCBElement) TextF(format string, args ...any) *SVGFEFUNCBElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfText(condition bool, text string) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEFUNCBElement) IfTextF(condition bool, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEFUNCBElement) Escaped(text string) *SVGFEFUNCBElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEFUNCBElement) IfEscaped(condition bool, text string) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEFUNCBElement) EscapedF(format string, args ...any) *SVGFEFUNCBElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEFUNCBElement) CustomData(key, value string) *SVGFEFUNCBElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEFUNCBElement) IfCustomData(condition bool, key, value string) *SVGFEFUNCBElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEFUNCBElement) CustomDataF(key, format string, args ...any) *SVGFEFUNCBElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEFUNCBElement) CustomDataRemove(key string) *SVGFEFUNCBElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The type of transfer function.
func (e *SVGFEFUNCBElement) TYPE(c SVGFeFuncBTypeChoice) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("type", string(c))
	return e
}

type SVGFeFuncBTypeChoice string

const (
	// The type of transfer function.
	SVGFeFuncBType_identity SVGFeFuncBTypeChoice = "identity"
	// The type of transfer function.
	SVGFeFuncBType_table SVGFeFuncBTypeChoice = "table"
	// The type of transfer function.
	SVGFeFuncBType_discrete SVGFeFuncBTypeChoice = "discrete"
	// The type of transfer function.
	SVGFeFuncBType_linear SVGFeFuncBTypeChoice = "linear"
	// The type of transfer function.
	SVGFeFuncBType_gamma SVGFeFuncBTypeChoice = "gamma"
)

// Remove the attribute TYPE from the element.
func (e *SVGFEFUNCBElement) TYPERemove(c SVGFeFuncBTypeChoice) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("type")
	return e
}

// Contains the list of <number>s that define the lookup table
// Values must be in the 0-1 range and be equally spaced
// There must be at least two values.
func (e *SVGFEFUNCBElement) TABLE_VALUES(s string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("tableValues", s)
	return e
}

func (e *SVGFEFUNCBElement) TABLE_VALUESF(format string, args ...any) *SVGFEFUNCBElement {
	return e.TABLE_VALUES(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfTABLE_VALUES(condition bool, s string) *SVGFEFUNCBElement {
	if condition {
		e.TABLE_VALUES(s)
	}
	return e
}

func (e *SVGFEFUNCBElement) IfTABLE_VALUESF(condition bool, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.TABLE_VALUES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TABLE_VALUES from the element.
func (e *SVGFEFUNCBElement) TABLE_VALUESRemove(s string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("tableValues")
	return e
}

func (e *SVGFEFUNCBElement) TABLE_VALUESRemoveF(format string, args ...any) *SVGFEFUNCBElement {
	return e.TABLE_VALUESRemove(fmt.Sprintf(format, args...))
}

// The slope attribute indicates the slope of the linear function.
func (e *SVGFEFUNCBElement) SLOPE(f float64) *SVGFEFUNCBElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("slope", f)
	return e
}

func (e *SVGFEFUNCBElement) IfSLOPE(condition bool, f float64) *SVGFEFUNCBElement {
	if condition {
		e.SLOPE(f)
	}
	return e
}

// The intercept attribute indicates the intercept of the linear function.
func (e *SVGFEFUNCBElement) INTERCEPT(f float64) *SVGFEFUNCBElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("intercept", f)
	return e
}

func (e *SVGFEFUNCBElement) IfINTERCEPT(condition bool, f float64) *SVGFEFUNCBElement {
	if condition {
		e.INTERCEPT(f)
	}
	return e
}

// The amplitude attribute indicates the amplitude of the cubic function.
func (e *SVGFEFUNCBElement) AMPLITUDE(f float64) *SVGFEFUNCBElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("amplitude", f)
	return e
}

func (e *SVGFEFUNCBElement) IfAMPLITUDE(condition bool, f float64) *SVGFEFUNCBElement {
	if condition {
		e.AMPLITUDE(f)
	}
	return e
}

// The exponent attribute indicates the exponent of the exponential function.
func (e *SVGFEFUNCBElement) EXPONENT(f float64) *SVGFEFUNCBElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("exponent", f)
	return e
}

func (e *SVGFEFUNCBElement) IfEXPONENT(condition bool, f float64) *SVGFEFUNCBElement {
	if condition {
		e.EXPONENT(f)
	}
	return e
}

// The offset attribute indicates the offset of the function.
func (e *SVGFEFUNCBElement) OFFSET(f float64) *SVGFEFUNCBElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("offset", f)
	return e
}

func (e *SVGFEFUNCBElement) IfOFFSET(condition bool, f float64) *SVGFEFUNCBElement {
	if condition {
		e.OFFSET(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFEFUNCBElement) ID(s string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEFUNCBElement) IDF(format string, args ...any) *SVGFEFUNCBElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfID(condition bool, s string) *SVGFEFUNCBElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEFUNCBElement) IfIDF(condition bool, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEFUNCBElement) IDRemove(s string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEFUNCBElement) IDRemoveF(format string, args ...any) *SVGFEFUNCBElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEFUNCBElement) CLASS(s ...string) *SVGFEFUNCBElement {
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

func (e *SVGFEFUNCBElement) IfCLASS(condition bool, s ...string) *SVGFEFUNCBElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEFUNCBElement) CLASSRemove(s ...string) *SVGFEFUNCBElement {
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
func (e *SVGFEFUNCBElement) STYLEF(k string, format string, args ...any) *SVGFEFUNCBElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfSTYLE(condition bool, k string, v string) *SVGFEFUNCBElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEFUNCBElement) STYLE(k string, v string) *SVGFEFUNCBElement {
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

func (e *SVGFEFUNCBElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEFUNCBElement) STYLEMap(m map[string]string) *SVGFEFUNCBElement {
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
func (e *SVGFEFUNCBElement) STYLEPairs(pairs ...string) *SVGFEFUNCBElement {
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

func (e *SVGFEFUNCBElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEFUNCBElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEFUNCBElement) STYLERemove(keys ...string) *SVGFEFUNCBElement {
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

func (e *SVGFEFUNCBElement) Z_REQ(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfZ_REQ(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEFUNCBElement) Z_REQRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEFUNCBElement) Z_TARGET(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfZ_TARGET(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEFUNCBElement) Z_TARGETRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEFUNCBElement) Z_REQ_SELECTOR(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEFUNCBElement) Z_REQ_SELECTORRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEFUNCBElement) Z_SWAP(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfZ_SWAP(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEFUNCBElement) Z_SWAPRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEFUNCBElement) Z_SWAP_PUSH(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEFUNCBElement) Z_SWAP_PUSHRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEFUNCBElement) Z_TRIGGER(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEFUNCBElement) Z_TRIGGERRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEFUNCBElement) Z_REQ_METHOD(c SVGFeFuncBZReqMethodChoice) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeFuncBZReqMethodChoice string

const (
	// default GET
	SVGFeFuncBZReqMethod_empty SVGFeFuncBZReqMethodChoice = ""
	// GET
	SVGFeFuncBZReqMethod_get SVGFeFuncBZReqMethodChoice = "get"
	// POST
	SVGFeFuncBZReqMethod_post SVGFeFuncBZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEFUNCBElement) Z_REQ_METHODRemove(c SVGFeFuncBZReqMethodChoice) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEFUNCBElement) Z_REQ_STRATEGY(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEFUNCBElement) Z_REQ_STRATEGYRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEFUNCBElement) Z_REQ_HISTORY(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEFUNCBElement) Z_REQ_HISTORYRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEFUNCBElement) Z_DATA(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfZ_DATA(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEFUNCBElement) Z_DATARemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEFUNCBElement) Z_JSON(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfZ_JSON(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEFUNCBElement) Z_JSONRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEFUNCBElement) Z_REQ_BATCH(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEFUNCBElement) Z_REQ_BATCHRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEFUNCBElement) Z_ACTION(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfZ_ACTION(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEFUNCBElement) Z_ACTIONRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEFUNCBElement) Z_REQ_BEFORE(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEFUNCBElement) Z_REQ_BEFORERemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEFUNCBElement) Z_REQ_AFTER(expression string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCBElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEFUNCBElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEFUNCBElement) Z_REQ_AFTERRemove() *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
