package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feFuncR> SVG filter primitive defines the transfer function for the red
// component of the input graphic of its parent <feComponentTransfer> element.
type SVGFEFUNCRElement struct {
	*Element
}

// Create a new SVGFEFUNCRElement element.
// This will create a new element with the tag
// "feFuncR" during rendering.
func SVG_FEFUNCR(children ...ElementRenderer) *SVGFEFUNCRElement {
	e := NewElement("feFuncR", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEFUNCRElement{Element: e}
}

func (e *SVGFEFUNCRElement) Children(children ...ElementRenderer) *SVGFEFUNCRElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEFUNCRElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEFUNCRElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEFUNCRElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEFUNCRElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEFUNCRElement) Attr(name string, value ...string) *SVGFEFUNCRElement {
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

func (e *SVGFEFUNCRElement) Attrs(attrs ...string) *SVGFEFUNCRElement {
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

func (e *SVGFEFUNCRElement) AttrsMap(attrs map[string]string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEFUNCRElement) Text(text string) *SVGFEFUNCRElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEFUNCRElement) TextF(format string, args ...any) *SVGFEFUNCRElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCRElement) IfText(condition bool, text string) *SVGFEFUNCRElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEFUNCRElement) IfTextF(condition bool, format string, args ...any) *SVGFEFUNCRElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEFUNCRElement) Escaped(text string) *SVGFEFUNCRElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEFUNCRElement) IfEscaped(condition bool, text string) *SVGFEFUNCRElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEFUNCRElement) EscapedF(format string, args ...any) *SVGFEFUNCRElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCRElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEFUNCRElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEFUNCRElement) CustomData(key, value string) *SVGFEFUNCRElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEFUNCRElement) IfCustomData(condition bool, key, value string) *SVGFEFUNCRElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEFUNCRElement) CustomDataF(key, format string, args ...any) *SVGFEFUNCRElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCRElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEFUNCRElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEFUNCRElement) CustomDataRemove(key string) *SVGFEFUNCRElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The type of transfer function.
func (e *SVGFEFUNCRElement) TYPE(c SVGFeFuncRTypeChoice) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("type", string(c))
	return e
}

type SVGFeFuncRTypeChoice string

const (
	// The type of transfer function.
	SVGFeFuncRType_identity SVGFeFuncRTypeChoice = "identity"
	// The type of transfer function.
	SVGFeFuncRType_table SVGFeFuncRTypeChoice = "table"
	// The type of transfer function.
	SVGFeFuncRType_discrete SVGFeFuncRTypeChoice = "discrete"
	// The type of transfer function.
	SVGFeFuncRType_linear SVGFeFuncRTypeChoice = "linear"
	// The type of transfer function.
	SVGFeFuncRType_gamma SVGFeFuncRTypeChoice = "gamma"
)

// Remove the attribute TYPE from the element.
func (e *SVGFEFUNCRElement) TYPERemove(c SVGFeFuncRTypeChoice) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("type")
	return e
}

// Contains the list of <number>s that define the lookup table
// Values must be in the 0-1 range and be equally spaced
// There must be at least two values.
func (e *SVGFEFUNCRElement) TABLE_VALUES(s string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("tableValues", s)
	return e
}

func (e *SVGFEFUNCRElement) TABLE_VALUESF(format string, args ...any) *SVGFEFUNCRElement {
	return e.TABLE_VALUES(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCRElement) IfTABLE_VALUES(condition bool, s string) *SVGFEFUNCRElement {
	if condition {
		e.TABLE_VALUES(s)
	}
	return e
}

func (e *SVGFEFUNCRElement) IfTABLE_VALUESF(condition bool, format string, args ...any) *SVGFEFUNCRElement {
	if condition {
		e.TABLE_VALUES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TABLE_VALUES from the element.
func (e *SVGFEFUNCRElement) TABLE_VALUESRemove(s string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("tableValues")
	return e
}

func (e *SVGFEFUNCRElement) TABLE_VALUESRemoveF(format string, args ...any) *SVGFEFUNCRElement {
	return e.TABLE_VALUESRemove(fmt.Sprintf(format, args...))
}

// The slope attribute indicates the slope of the linear function.
func (e *SVGFEFUNCRElement) SLOPE(f float64) *SVGFEFUNCRElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("slope", f)
	return e
}

func (e *SVGFEFUNCRElement) IfSLOPE(condition bool, f float64) *SVGFEFUNCRElement {
	if condition {
		e.SLOPE(f)
	}
	return e
}

// The intercept attribute indicates the intercept of the linear function.
func (e *SVGFEFUNCRElement) INTERCEPT(f float64) *SVGFEFUNCRElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("intercept", f)
	return e
}

func (e *SVGFEFUNCRElement) IfINTERCEPT(condition bool, f float64) *SVGFEFUNCRElement {
	if condition {
		e.INTERCEPT(f)
	}
	return e
}

// The amplitude attribute indicates the amplitude of the cubic function.
func (e *SVGFEFUNCRElement) AMPLITUDE(f float64) *SVGFEFUNCRElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("amplitude", f)
	return e
}

func (e *SVGFEFUNCRElement) IfAMPLITUDE(condition bool, f float64) *SVGFEFUNCRElement {
	if condition {
		e.AMPLITUDE(f)
	}
	return e
}

// The exponent attribute indicates the exponent of the exponential function.
func (e *SVGFEFUNCRElement) EXPONENT(f float64) *SVGFEFUNCRElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("exponent", f)
	return e
}

func (e *SVGFEFUNCRElement) IfEXPONENT(condition bool, f float64) *SVGFEFUNCRElement {
	if condition {
		e.EXPONENT(f)
	}
	return e
}

// The offset attribute indicates the offset of the function.
func (e *SVGFEFUNCRElement) OFFSET(f float64) *SVGFEFUNCRElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("offset", f)
	return e
}

func (e *SVGFEFUNCRElement) IfOFFSET(condition bool, f float64) *SVGFEFUNCRElement {
	if condition {
		e.OFFSET(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFEFUNCRElement) ID(s string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEFUNCRElement) IDF(format string, args ...any) *SVGFEFUNCRElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCRElement) IfID(condition bool, s string) *SVGFEFUNCRElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEFUNCRElement) IfIDF(condition bool, format string, args ...any) *SVGFEFUNCRElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEFUNCRElement) IDRemove(s string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEFUNCRElement) IDRemoveF(format string, args ...any) *SVGFEFUNCRElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEFUNCRElement) CLASS(s ...string) *SVGFEFUNCRElement {
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

func (e *SVGFEFUNCRElement) IfCLASS(condition bool, s ...string) *SVGFEFUNCRElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEFUNCRElement) CLASSRemove(s ...string) *SVGFEFUNCRElement {
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
func (e *SVGFEFUNCRElement) STYLEF(k string, format string, args ...any) *SVGFEFUNCRElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCRElement) IfSTYLE(condition bool, k string, v string) *SVGFEFUNCRElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEFUNCRElement) STYLE(k string, v string) *SVGFEFUNCRElement {
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

func (e *SVGFEFUNCRElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEFUNCRElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEFUNCRElement) STYLEMap(m map[string]string) *SVGFEFUNCRElement {
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
func (e *SVGFEFUNCRElement) STYLEPairs(pairs ...string) *SVGFEFUNCRElement {
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

func (e *SVGFEFUNCRElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEFUNCRElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEFUNCRElement) STYLERemove(keys ...string) *SVGFEFUNCRElement {
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

func (e *SVGFEFUNCRElement) Z_REQ(expression string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCRElement) IfZ_REQ(condition bool, expression string) *SVGFEFUNCRElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEFUNCRElement) Z_REQRemove() *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEFUNCRElement) Z_TARGET(expression string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCRElement) IfZ_TARGET(condition bool, expression string) *SVGFEFUNCRElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEFUNCRElement) Z_TARGETRemove() *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEFUNCRElement) Z_REQ_SELECTOR(expression string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCRElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEFUNCRElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEFUNCRElement) Z_REQ_SELECTORRemove() *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEFUNCRElement) Z_SWAP(expression string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCRElement) IfZ_SWAP(condition bool, expression string) *SVGFEFUNCRElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEFUNCRElement) Z_SWAPRemove() *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEFUNCRElement) Z_SWAP_PUSH(expression string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCRElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEFUNCRElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEFUNCRElement) Z_SWAP_PUSHRemove() *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEFUNCRElement) Z_TRIGGER(expression string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCRElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEFUNCRElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEFUNCRElement) Z_TRIGGERRemove() *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEFUNCRElement) Z_REQ_METHOD(c SVGFeFuncRZReqMethodChoice) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeFuncRZReqMethodChoice string

const (
	// default GET
	SVGFeFuncRZReqMethod_empty SVGFeFuncRZReqMethodChoice = ""
	// GET
	SVGFeFuncRZReqMethod_get SVGFeFuncRZReqMethodChoice = "get"
	// POST
	SVGFeFuncRZReqMethod_post SVGFeFuncRZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEFUNCRElement) Z_REQ_METHODRemove(c SVGFeFuncRZReqMethodChoice) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEFUNCRElement) Z_REQ_STRATEGY(expression string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCRElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEFUNCRElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEFUNCRElement) Z_REQ_STRATEGYRemove() *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEFUNCRElement) Z_REQ_HISTORY(expression string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCRElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEFUNCRElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEFUNCRElement) Z_REQ_HISTORYRemove() *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEFUNCRElement) Z_DATA(expression string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCRElement) IfZ_DATA(condition bool, expression string) *SVGFEFUNCRElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEFUNCRElement) Z_DATARemove() *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEFUNCRElement) Z_JSON(expression string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCRElement) IfZ_JSON(condition bool, expression string) *SVGFEFUNCRElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEFUNCRElement) Z_JSONRemove() *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEFUNCRElement) Z_REQ_BATCH(expression string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCRElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEFUNCRElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEFUNCRElement) Z_REQ_BATCHRemove() *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEFUNCRElement) Z_ACTION(expression string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCRElement) IfZ_ACTION(condition bool, expression string) *SVGFEFUNCRElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEFUNCRElement) Z_ACTIONRemove() *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEFUNCRElement) Z_REQ_BEFORE(expression string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCRElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEFUNCRElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEFUNCRElement) Z_REQ_BEFORERemove() *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEFUNCRElement) Z_REQ_AFTER(expression string) *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCRElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEFUNCRElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEFUNCRElement) Z_REQ_AFTERRemove() *SVGFEFUNCRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
