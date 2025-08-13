package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feMorphology> SVG filter primitive is used to erode or dilate the input
// image
// It's usefulness lies especially in fattening or thinning effects.
type SVGFEMORPHOLOGYElement struct {
	*Element
}

// Create a new SVGFEMORPHOLOGYElement element.
// This will create a new element with the tag
// "feMorphology" during rendering.
func SVG_FEMORPHOLOGY(children ...ElementRenderer) *SVGFEMORPHOLOGYElement {
	e := NewElement("feMorphology", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEMORPHOLOGYElement{Element: e}
}

func (e *SVGFEMORPHOLOGYElement) Children(children ...ElementRenderer) *SVGFEMORPHOLOGYElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEMORPHOLOGYElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEMORPHOLOGYElement) Attr(name string, value ...string) *SVGFEMORPHOLOGYElement {
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

func (e *SVGFEMORPHOLOGYElement) Attrs(attrs ...string) *SVGFEMORPHOLOGYElement {
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

func (e *SVGFEMORPHOLOGYElement) AttrsMap(attrs map[string]string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEMORPHOLOGYElement) Text(text string) *SVGFEMORPHOLOGYElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEMORPHOLOGYElement) TextF(format string, args ...any) *SVGFEMORPHOLOGYElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEMORPHOLOGYElement) IfText(condition bool, text string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfTextF(condition bool, format string, args ...any) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEMORPHOLOGYElement) Escaped(text string) *SVGFEMORPHOLOGYElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfEscaped(condition bool, text string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEMORPHOLOGYElement) EscapedF(format string, args ...any) *SVGFEMORPHOLOGYElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEMORPHOLOGYElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEMORPHOLOGYElement) CustomData(key, value string) *SVGFEMORPHOLOGYElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfCustomData(condition bool, key, value string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEMORPHOLOGYElement) CustomDataF(key, format string, args ...any) *SVGFEMORPHOLOGYElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEMORPHOLOGYElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEMORPHOLOGYElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEMORPHOLOGYElement) CustomDataRemove(key string) *SVGFEMORPHOLOGYElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The input for this filter.
func (e *SVGFEMORPHOLOGYElement) IN(s string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFEMORPHOLOGYElement) INF(format string, args ...any) *SVGFEMORPHOLOGYElement {
	return e.IN(fmt.Sprintf(format, args...))
}

func (e *SVGFEMORPHOLOGYElement) IfIN(condition bool, s string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.IN(s)
	}
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfINF(condition bool, format string, args ...any) *SVGFEMORPHOLOGYElement {
	if condition {
		e.IN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFEMORPHOLOGYElement) INRemove(s string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("in")
	return e
}

func (e *SVGFEMORPHOLOGYElement) INRemoveF(format string, args ...any) *SVGFEMORPHOLOGYElement {
	return e.INRemove(fmt.Sprintf(format, args...))
}

// The operator attribute defines what type of operation is performed.
func (e *SVGFEMORPHOLOGYElement) OPERATOR(c SVGFeMorphologyOperatorChoice) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("operator", string(c))
	return e
}

type SVGFeMorphologyOperatorChoice string

const (
	// The operator attribute defines what type of operation is performed.
	SVGFeMorphologyOperator_erode SVGFeMorphologyOperatorChoice = "erode"
	// The operator attribute defines what type of operation is performed.
	SVGFeMorphologyOperator_dilate SVGFeMorphologyOperatorChoice = "dilate"
)

// Remove the attribute OPERATOR from the element.
func (e *SVGFEMORPHOLOGYElement) OPERATORRemove(c SVGFeMorphologyOperatorChoice) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("operator")
	return e
}

// The radius attribute indicates the size of the matrix.
func (e *SVGFEMORPHOLOGYElement) RADIUS(f float64) *SVGFEMORPHOLOGYElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("radius", f)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfRADIUS(condition bool, f float64) *SVGFEMORPHOLOGYElement {
	if condition {
		e.RADIUS(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFEMORPHOLOGYElement) ID(s string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IDF(format string, args ...any) *SVGFEMORPHOLOGYElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEMORPHOLOGYElement) IfID(condition bool, s string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfIDF(condition bool, format string, args ...any) *SVGFEMORPHOLOGYElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEMORPHOLOGYElement) IDRemove(s string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEMORPHOLOGYElement) IDRemoveF(format string, args ...any) *SVGFEMORPHOLOGYElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEMORPHOLOGYElement) CLASS(s ...string) *SVGFEMORPHOLOGYElement {
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

func (e *SVGFEMORPHOLOGYElement) IfCLASS(condition bool, s ...string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEMORPHOLOGYElement) CLASSRemove(s ...string) *SVGFEMORPHOLOGYElement {
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
func (e *SVGFEMORPHOLOGYElement) STYLEF(k string, format string, args ...any) *SVGFEMORPHOLOGYElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEMORPHOLOGYElement) IfSTYLE(condition bool, k string, v string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEMORPHOLOGYElement) STYLE(k string, v string) *SVGFEMORPHOLOGYElement {
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

func (e *SVGFEMORPHOLOGYElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEMORPHOLOGYElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEMORPHOLOGYElement) STYLEMap(m map[string]string) *SVGFEMORPHOLOGYElement {
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
func (e *SVGFEMORPHOLOGYElement) STYLEPairs(pairs ...string) *SVGFEMORPHOLOGYElement {
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

func (e *SVGFEMORPHOLOGYElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEMORPHOLOGYElement) STYLERemove(keys ...string) *SVGFEMORPHOLOGYElement {
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

func (e *SVGFEMORPHOLOGYElement) Z_REQ(expression string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfZ_REQ(condition bool, expression string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEMORPHOLOGYElement) Z_REQRemove() *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEMORPHOLOGYElement) Z_TARGET(expression string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfZ_TARGET(condition bool, expression string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEMORPHOLOGYElement) Z_TARGETRemove() *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEMORPHOLOGYElement) Z_REQ_SELECTOR(expression string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEMORPHOLOGYElement) Z_REQ_SELECTORRemove() *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEMORPHOLOGYElement) Z_SWAP(expression string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfZ_SWAP(condition bool, expression string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEMORPHOLOGYElement) Z_SWAPRemove() *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEMORPHOLOGYElement) Z_SWAP_PUSH(expression string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEMORPHOLOGYElement) Z_SWAP_PUSHRemove() *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEMORPHOLOGYElement) Z_TRIGGER(expression string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEMORPHOLOGYElement) Z_TRIGGERRemove() *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEMORPHOLOGYElement) Z_REQ_METHOD(c SVGFeMorphologyZReqMethodChoice) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeMorphologyZReqMethodChoice string

const (
	// default GET
	SVGFeMorphologyZReqMethod_empty SVGFeMorphologyZReqMethodChoice = ""
	// GET
	SVGFeMorphologyZReqMethod_get SVGFeMorphologyZReqMethodChoice = "get"
	// POST
	SVGFeMorphologyZReqMethod_post SVGFeMorphologyZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEMORPHOLOGYElement) Z_REQ_METHODRemove(c SVGFeMorphologyZReqMethodChoice) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEMORPHOLOGYElement) Z_REQ_STRATEGY(expression string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEMORPHOLOGYElement) Z_REQ_STRATEGYRemove() *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEMORPHOLOGYElement) Z_REQ_HISTORY(expression string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEMORPHOLOGYElement) Z_REQ_HISTORYRemove() *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEMORPHOLOGYElement) Z_DATA(expression string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfZ_DATA(condition bool, expression string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEMORPHOLOGYElement) Z_DATARemove() *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEMORPHOLOGYElement) Z_JSON(expression string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfZ_JSON(condition bool, expression string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEMORPHOLOGYElement) Z_JSONRemove() *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEMORPHOLOGYElement) Z_REQ_BATCH(expression string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEMORPHOLOGYElement) Z_REQ_BATCHRemove() *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEMORPHOLOGYElement) Z_ACTION(expression string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfZ_ACTION(condition bool, expression string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEMORPHOLOGYElement) Z_ACTIONRemove() *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEMORPHOLOGYElement) Z_REQ_BEFORE(expression string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEMORPHOLOGYElement) Z_REQ_BEFORERemove() *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEMORPHOLOGYElement) Z_REQ_AFTER(expression string) *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMORPHOLOGYElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEMORPHOLOGYElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEMORPHOLOGYElement) Z_REQ_AFTERRemove() *SVGFEMORPHOLOGYElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
