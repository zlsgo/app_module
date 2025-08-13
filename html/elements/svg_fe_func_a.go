package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feFuncA> SVG filter primitive defines the transfer function for the alpha
// component of the input graphic of its parent <feComponentTransfer> element.
type SVGFEFUNCAElement struct {
	*Element
}

// Create a new SVGFEFUNCAElement element.
// This will create a new element with the tag
// "feFuncA" during rendering.
func SVG_FEFUNCA(children ...ElementRenderer) *SVGFEFUNCAElement {
	e := NewElement("feFuncA", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEFUNCAElement{Element: e}
}

func (e *SVGFEFUNCAElement) Children(children ...ElementRenderer) *SVGFEFUNCAElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEFUNCAElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEFUNCAElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEFUNCAElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEFUNCAElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEFUNCAElement) Attr(name string, value ...string) *SVGFEFUNCAElement {
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

func (e *SVGFEFUNCAElement) Attrs(attrs ...string) *SVGFEFUNCAElement {
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

func (e *SVGFEFUNCAElement) AttrsMap(attrs map[string]string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEFUNCAElement) Text(text string) *SVGFEFUNCAElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEFUNCAElement) TextF(format string, args ...any) *SVGFEFUNCAElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCAElement) IfText(condition bool, text string) *SVGFEFUNCAElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEFUNCAElement) IfTextF(condition bool, format string, args ...any) *SVGFEFUNCAElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEFUNCAElement) Escaped(text string) *SVGFEFUNCAElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEFUNCAElement) IfEscaped(condition bool, text string) *SVGFEFUNCAElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEFUNCAElement) EscapedF(format string, args ...any) *SVGFEFUNCAElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCAElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEFUNCAElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEFUNCAElement) CustomData(key, value string) *SVGFEFUNCAElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEFUNCAElement) IfCustomData(condition bool, key, value string) *SVGFEFUNCAElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEFUNCAElement) CustomDataF(key, format string, args ...any) *SVGFEFUNCAElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCAElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEFUNCAElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEFUNCAElement) CustomDataRemove(key string) *SVGFEFUNCAElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The type of transfer function.
func (e *SVGFEFUNCAElement) TYPE(c SVGFeFuncATypeChoice) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("type", string(c))
	return e
}

type SVGFeFuncATypeChoice string

const (
	// The type of transfer function.
	SVGFeFuncAType_identity SVGFeFuncATypeChoice = "identity"
	// The type of transfer function.
	SVGFeFuncAType_table SVGFeFuncATypeChoice = "table"
	// The type of transfer function.
	SVGFeFuncAType_discrete SVGFeFuncATypeChoice = "discrete"
	// The type of transfer function.
	SVGFeFuncAType_linear SVGFeFuncATypeChoice = "linear"
	// The type of transfer function.
	SVGFeFuncAType_gamma SVGFeFuncATypeChoice = "gamma"
)

// Remove the attribute TYPE from the element.
func (e *SVGFEFUNCAElement) TYPERemove(c SVGFeFuncATypeChoice) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("type")
	return e
}

// Contains the list of <number>s that define the lookup table
// Values must be in the 0-1 range and be equally spaced
// There must be at least two values.
func (e *SVGFEFUNCAElement) TABLE_VALUES(s string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("tableValues", s)
	return e
}

func (e *SVGFEFUNCAElement) TABLE_VALUESF(format string, args ...any) *SVGFEFUNCAElement {
	return e.TABLE_VALUES(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCAElement) IfTABLE_VALUES(condition bool, s string) *SVGFEFUNCAElement {
	if condition {
		e.TABLE_VALUES(s)
	}
	return e
}

func (e *SVGFEFUNCAElement) IfTABLE_VALUESF(condition bool, format string, args ...any) *SVGFEFUNCAElement {
	if condition {
		e.TABLE_VALUES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TABLE_VALUES from the element.
func (e *SVGFEFUNCAElement) TABLE_VALUESRemove(s string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("tableValues")
	return e
}

func (e *SVGFEFUNCAElement) TABLE_VALUESRemoveF(format string, args ...any) *SVGFEFUNCAElement {
	return e.TABLE_VALUESRemove(fmt.Sprintf(format, args...))
}

// The slope attribute indicates the slope of the linear function.
func (e *SVGFEFUNCAElement) SLOPE(f float64) *SVGFEFUNCAElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("slope", f)
	return e
}

func (e *SVGFEFUNCAElement) IfSLOPE(condition bool, f float64) *SVGFEFUNCAElement {
	if condition {
		e.SLOPE(f)
	}
	return e
}

// The intercept attribute indicates the intercept of the linear function.
func (e *SVGFEFUNCAElement) INTERCEPT(f float64) *SVGFEFUNCAElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("intercept", f)
	return e
}

func (e *SVGFEFUNCAElement) IfINTERCEPT(condition bool, f float64) *SVGFEFUNCAElement {
	if condition {
		e.INTERCEPT(f)
	}
	return e
}

// The amplitude attribute indicates the amplitude of the cubic function.
func (e *SVGFEFUNCAElement) AMPLITUDE(f float64) *SVGFEFUNCAElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("amplitude", f)
	return e
}

func (e *SVGFEFUNCAElement) IfAMPLITUDE(condition bool, f float64) *SVGFEFUNCAElement {
	if condition {
		e.AMPLITUDE(f)
	}
	return e
}

// The exponent attribute indicates the exponent of the exponential function.
func (e *SVGFEFUNCAElement) EXPONENT(f float64) *SVGFEFUNCAElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("exponent", f)
	return e
}

func (e *SVGFEFUNCAElement) IfEXPONENT(condition bool, f float64) *SVGFEFUNCAElement {
	if condition {
		e.EXPONENT(f)
	}
	return e
}

// The offset attribute indicates the offset of the function.
func (e *SVGFEFUNCAElement) OFFSET(f float64) *SVGFEFUNCAElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("offset", f)
	return e
}

func (e *SVGFEFUNCAElement) IfOFFSET(condition bool, f float64) *SVGFEFUNCAElement {
	if condition {
		e.OFFSET(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFEFUNCAElement) ID(s string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEFUNCAElement) IDF(format string, args ...any) *SVGFEFUNCAElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCAElement) IfID(condition bool, s string) *SVGFEFUNCAElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEFUNCAElement) IfIDF(condition bool, format string, args ...any) *SVGFEFUNCAElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEFUNCAElement) IDRemove(s string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEFUNCAElement) IDRemoveF(format string, args ...any) *SVGFEFUNCAElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEFUNCAElement) CLASS(s ...string) *SVGFEFUNCAElement {
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

func (e *SVGFEFUNCAElement) IfCLASS(condition bool, s ...string) *SVGFEFUNCAElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEFUNCAElement) CLASSRemove(s ...string) *SVGFEFUNCAElement {
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
func (e *SVGFEFUNCAElement) STYLEF(k string, format string, args ...any) *SVGFEFUNCAElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCAElement) IfSTYLE(condition bool, k string, v string) *SVGFEFUNCAElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEFUNCAElement) STYLE(k string, v string) *SVGFEFUNCAElement {
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

func (e *SVGFEFUNCAElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEFUNCAElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEFUNCAElement) STYLEMap(m map[string]string) *SVGFEFUNCAElement {
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
func (e *SVGFEFUNCAElement) STYLEPairs(pairs ...string) *SVGFEFUNCAElement {
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

func (e *SVGFEFUNCAElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEFUNCAElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEFUNCAElement) STYLERemove(keys ...string) *SVGFEFUNCAElement {
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

func (e *SVGFEFUNCAElement) Z_REQ(expression string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCAElement) IfZ_REQ(condition bool, expression string) *SVGFEFUNCAElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEFUNCAElement) Z_REQRemove() *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEFUNCAElement) Z_TARGET(expression string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCAElement) IfZ_TARGET(condition bool, expression string) *SVGFEFUNCAElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEFUNCAElement) Z_TARGETRemove() *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEFUNCAElement) Z_REQ_SELECTOR(expression string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCAElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEFUNCAElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEFUNCAElement) Z_REQ_SELECTORRemove() *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEFUNCAElement) Z_SWAP(expression string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCAElement) IfZ_SWAP(condition bool, expression string) *SVGFEFUNCAElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEFUNCAElement) Z_SWAPRemove() *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEFUNCAElement) Z_SWAP_PUSH(expression string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCAElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEFUNCAElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEFUNCAElement) Z_SWAP_PUSHRemove() *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEFUNCAElement) Z_TRIGGER(expression string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCAElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEFUNCAElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEFUNCAElement) Z_TRIGGERRemove() *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEFUNCAElement) Z_REQ_METHOD(c SVGFeFuncAZReqMethodChoice) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeFuncAZReqMethodChoice string

const (
	// default GET
	SVGFeFuncAZReqMethod_empty SVGFeFuncAZReqMethodChoice = ""
	// GET
	SVGFeFuncAZReqMethod_get SVGFeFuncAZReqMethodChoice = "get"
	// POST
	SVGFeFuncAZReqMethod_post SVGFeFuncAZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEFUNCAElement) Z_REQ_METHODRemove(c SVGFeFuncAZReqMethodChoice) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEFUNCAElement) Z_REQ_STRATEGY(expression string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCAElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEFUNCAElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEFUNCAElement) Z_REQ_STRATEGYRemove() *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEFUNCAElement) Z_REQ_HISTORY(expression string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCAElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEFUNCAElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEFUNCAElement) Z_REQ_HISTORYRemove() *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEFUNCAElement) Z_DATA(expression string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCAElement) IfZ_DATA(condition bool, expression string) *SVGFEFUNCAElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEFUNCAElement) Z_DATARemove() *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEFUNCAElement) Z_JSON(expression string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCAElement) IfZ_JSON(condition bool, expression string) *SVGFEFUNCAElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEFUNCAElement) Z_JSONRemove() *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEFUNCAElement) Z_REQ_BATCH(expression string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCAElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEFUNCAElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEFUNCAElement) Z_REQ_BATCHRemove() *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEFUNCAElement) Z_ACTION(expression string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCAElement) IfZ_ACTION(condition bool, expression string) *SVGFEFUNCAElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEFUNCAElement) Z_ACTIONRemove() *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEFUNCAElement) Z_REQ_BEFORE(expression string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCAElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEFUNCAElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEFUNCAElement) Z_REQ_BEFORERemove() *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEFUNCAElement) Z_REQ_AFTER(expression string) *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCAElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEFUNCAElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEFUNCAElement) Z_REQ_AFTERRemove() *SVGFEFUNCAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
