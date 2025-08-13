package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feFuncG> SVG filter primitive defines the transfer function for the green
// component of the input graphic of its parent <feComponentTransfer> element.
type SVGFEFUNCGElement struct {
	*Element
}

// Create a new SVGFEFUNCGElement element.
// This will create a new element with the tag
// "feFuncG" during rendering.
func SVG_FEFUNCG(children ...ElementRenderer) *SVGFEFUNCGElement {
	e := NewElement("feFuncG", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEFUNCGElement{Element: e}
}

func (e *SVGFEFUNCGElement) Children(children ...ElementRenderer) *SVGFEFUNCGElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEFUNCGElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEFUNCGElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEFUNCGElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEFUNCGElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEFUNCGElement) Attr(name string, value ...string) *SVGFEFUNCGElement {
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

func (e *SVGFEFUNCGElement) Attrs(attrs ...string) *SVGFEFUNCGElement {
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

func (e *SVGFEFUNCGElement) AttrsMap(attrs map[string]string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEFUNCGElement) Text(text string) *SVGFEFUNCGElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEFUNCGElement) TextF(format string, args ...any) *SVGFEFUNCGElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCGElement) IfText(condition bool, text string) *SVGFEFUNCGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEFUNCGElement) IfTextF(condition bool, format string, args ...any) *SVGFEFUNCGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEFUNCGElement) Escaped(text string) *SVGFEFUNCGElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEFUNCGElement) IfEscaped(condition bool, text string) *SVGFEFUNCGElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEFUNCGElement) EscapedF(format string, args ...any) *SVGFEFUNCGElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCGElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEFUNCGElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEFUNCGElement) CustomData(key, value string) *SVGFEFUNCGElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEFUNCGElement) IfCustomData(condition bool, key, value string) *SVGFEFUNCGElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEFUNCGElement) CustomDataF(key, format string, args ...any) *SVGFEFUNCGElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCGElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEFUNCGElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEFUNCGElement) CustomDataRemove(key string) *SVGFEFUNCGElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The type of transfer function.
func (e *SVGFEFUNCGElement) TYPE(c SVGFeFuncGTypeChoice) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("type", string(c))
	return e
}

type SVGFeFuncGTypeChoice string

const (
	// The type of transfer function.
	SVGFeFuncGType_identity SVGFeFuncGTypeChoice = "identity"
	// The type of transfer function.
	SVGFeFuncGType_table SVGFeFuncGTypeChoice = "table"
	// The type of transfer function.
	SVGFeFuncGType_discrete SVGFeFuncGTypeChoice = "discrete"
	// The type of transfer function.
	SVGFeFuncGType_linear SVGFeFuncGTypeChoice = "linear"
	// The type of transfer function.
	SVGFeFuncGType_gamma SVGFeFuncGTypeChoice = "gamma"
)

// Remove the attribute TYPE from the element.
func (e *SVGFEFUNCGElement) TYPERemove(c SVGFeFuncGTypeChoice) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("type")
	return e
}

// Contains the list of <number>s that define the lookup table
// Values must be in the 0-1 range and be equally spaced
// There must be at least two values.
func (e *SVGFEFUNCGElement) TABLE_VALUES(s string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("tableValues", s)
	return e
}

func (e *SVGFEFUNCGElement) TABLE_VALUESF(format string, args ...any) *SVGFEFUNCGElement {
	return e.TABLE_VALUES(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCGElement) IfTABLE_VALUES(condition bool, s string) *SVGFEFUNCGElement {
	if condition {
		e.TABLE_VALUES(s)
	}
	return e
}

func (e *SVGFEFUNCGElement) IfTABLE_VALUESF(condition bool, format string, args ...any) *SVGFEFUNCGElement {
	if condition {
		e.TABLE_VALUES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TABLE_VALUES from the element.
func (e *SVGFEFUNCGElement) TABLE_VALUESRemove(s string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("tableValues")
	return e
}

func (e *SVGFEFUNCGElement) TABLE_VALUESRemoveF(format string, args ...any) *SVGFEFUNCGElement {
	return e.TABLE_VALUESRemove(fmt.Sprintf(format, args...))
}

// The slope attribute indicates the slope of the linear function.
func (e *SVGFEFUNCGElement) SLOPE(f float64) *SVGFEFUNCGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("slope", f)
	return e
}

func (e *SVGFEFUNCGElement) IfSLOPE(condition bool, f float64) *SVGFEFUNCGElement {
	if condition {
		e.SLOPE(f)
	}
	return e
}

// The intercept attribute indicates the intercept of the linear function.
func (e *SVGFEFUNCGElement) INTERCEPT(f float64) *SVGFEFUNCGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("intercept", f)
	return e
}

func (e *SVGFEFUNCGElement) IfINTERCEPT(condition bool, f float64) *SVGFEFUNCGElement {
	if condition {
		e.INTERCEPT(f)
	}
	return e
}

// The amplitude attribute indicates the amplitude of the cubic function.
func (e *SVGFEFUNCGElement) AMPLITUDE(f float64) *SVGFEFUNCGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("amplitude", f)
	return e
}

func (e *SVGFEFUNCGElement) IfAMPLITUDE(condition bool, f float64) *SVGFEFUNCGElement {
	if condition {
		e.AMPLITUDE(f)
	}
	return e
}

// The exponent attribute indicates the exponent of the exponential function.
func (e *SVGFEFUNCGElement) EXPONENT(f float64) *SVGFEFUNCGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("exponent", f)
	return e
}

func (e *SVGFEFUNCGElement) IfEXPONENT(condition bool, f float64) *SVGFEFUNCGElement {
	if condition {
		e.EXPONENT(f)
	}
	return e
}

// The offset attribute indicates the offset of the function.
func (e *SVGFEFUNCGElement) OFFSET(f float64) *SVGFEFUNCGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("offset", f)
	return e
}

func (e *SVGFEFUNCGElement) IfOFFSET(condition bool, f float64) *SVGFEFUNCGElement {
	if condition {
		e.OFFSET(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFEFUNCGElement) ID(s string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEFUNCGElement) IDF(format string, args ...any) *SVGFEFUNCGElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCGElement) IfID(condition bool, s string) *SVGFEFUNCGElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEFUNCGElement) IfIDF(condition bool, format string, args ...any) *SVGFEFUNCGElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEFUNCGElement) IDRemove(s string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEFUNCGElement) IDRemoveF(format string, args ...any) *SVGFEFUNCGElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEFUNCGElement) CLASS(s ...string) *SVGFEFUNCGElement {
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

func (e *SVGFEFUNCGElement) IfCLASS(condition bool, s ...string) *SVGFEFUNCGElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEFUNCGElement) CLASSRemove(s ...string) *SVGFEFUNCGElement {
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
func (e *SVGFEFUNCGElement) STYLEF(k string, format string, args ...any) *SVGFEFUNCGElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCGElement) IfSTYLE(condition bool, k string, v string) *SVGFEFUNCGElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEFUNCGElement) STYLE(k string, v string) *SVGFEFUNCGElement {
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

func (e *SVGFEFUNCGElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEFUNCGElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEFUNCGElement) STYLEMap(m map[string]string) *SVGFEFUNCGElement {
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
func (e *SVGFEFUNCGElement) STYLEPairs(pairs ...string) *SVGFEFUNCGElement {
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

func (e *SVGFEFUNCGElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEFUNCGElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEFUNCGElement) STYLERemove(keys ...string) *SVGFEFUNCGElement {
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

func (e *SVGFEFUNCGElement) Z_REQ(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfZ_REQ(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEFUNCGElement) Z_REQRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEFUNCGElement) Z_TARGET(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfZ_TARGET(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEFUNCGElement) Z_TARGETRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEFUNCGElement) Z_REQ_SELECTOR(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEFUNCGElement) Z_REQ_SELECTORRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEFUNCGElement) Z_SWAP(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfZ_SWAP(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEFUNCGElement) Z_SWAPRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEFUNCGElement) Z_SWAP_PUSH(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEFUNCGElement) Z_SWAP_PUSHRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEFUNCGElement) Z_TRIGGER(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEFUNCGElement) Z_TRIGGERRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEFUNCGElement) Z_REQ_METHOD(c SVGFeFuncGZReqMethodChoice) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeFuncGZReqMethodChoice string

const (
	// default GET
	SVGFeFuncGZReqMethod_empty SVGFeFuncGZReqMethodChoice = ""
	// GET
	SVGFeFuncGZReqMethod_get SVGFeFuncGZReqMethodChoice = "get"
	// POST
	SVGFeFuncGZReqMethod_post SVGFeFuncGZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEFUNCGElement) Z_REQ_METHODRemove(c SVGFeFuncGZReqMethodChoice) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEFUNCGElement) Z_REQ_STRATEGY(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEFUNCGElement) Z_REQ_STRATEGYRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEFUNCGElement) Z_REQ_HISTORY(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEFUNCGElement) Z_REQ_HISTORYRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEFUNCGElement) Z_DATA(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfZ_DATA(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEFUNCGElement) Z_DATARemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEFUNCGElement) Z_JSON(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfZ_JSON(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEFUNCGElement) Z_JSONRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEFUNCGElement) Z_REQ_BATCH(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEFUNCGElement) Z_REQ_BATCHRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEFUNCGElement) Z_ACTION(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfZ_ACTION(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEFUNCGElement) Z_ACTIONRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEFUNCGElement) Z_REQ_BEFORE(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEFUNCGElement) Z_REQ_BEFORERemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEFUNCGElement) Z_REQ_AFTER(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEFUNCGElement) Z_REQ_AFTERRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
