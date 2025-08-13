package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feComponentTransfer> SVG filter primitive performs color-component-wise
// remapping of data for each pixel
// It allows operations like brightness adjustment, contrast adjustment, color
// balance or thresholding.
type SVGFECOMPONENTTRANSFERElement struct {
	*Element
}

// Create a new SVGFECOMPONENTTRANSFERElement element.
// This will create a new element with the tag
// "feComponentTransfer" during rendering.
func SVG_FECOMPONENTTRANSFER(children ...ElementRenderer) *SVGFECOMPONENTTRANSFERElement {
	e := NewElement("feComponentTransfer", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFECOMPONENTTRANSFERElement{Element: e}
}

func (e *SVGFECOMPONENTTRANSFERElement) Children(children ...ElementRenderer) *SVGFECOMPONENTTRANSFERElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) Attr(name string, value ...string) *SVGFECOMPONENTTRANSFERElement {
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

func (e *SVGFECOMPONENTTRANSFERElement) Attrs(attrs ...string) *SVGFECOMPONENTTRANSFERElement {
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

func (e *SVGFECOMPONENTTRANSFERElement) AttrsMap(attrs map[string]string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) Text(text string) *SVGFECOMPONENTTRANSFERElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) TextF(format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPONENTTRANSFERElement) IfText(condition bool, text string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfTextF(condition bool, format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) Escaped(text string) *SVGFECOMPONENTTRANSFERElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfEscaped(condition bool, text string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) EscapedF(format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPONENTTRANSFERElement) IfEscapedF(condition bool, format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) CustomData(key, value string) *SVGFECOMPONENTTRANSFERElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfCustomData(condition bool, key, value string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) CustomDataF(key, format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPONENTTRANSFERElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) CustomDataRemove(key string) *SVGFECOMPONENTTRANSFERElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The input for this filter.
func (e *SVGFECOMPONENTTRANSFERElement) IN(s string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) INF(format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	return e.IN(fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPONENTTRANSFERElement) IfIN(condition bool, s string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.IN(s)
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfINF(condition bool, format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.IN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFECOMPONENTTRANSFERElement) INRemove(s string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("in")
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) INRemoveF(format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	return e.INRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGFECOMPONENTTRANSFERElement) ID(s string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IDF(format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPONENTTRANSFERElement) IfID(condition bool, s string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfIDF(condition bool, format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFECOMPONENTTRANSFERElement) IDRemove(s string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IDRemoveF(format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFECOMPONENTTRANSFERElement) CLASS(s ...string) *SVGFECOMPONENTTRANSFERElement {
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

func (e *SVGFECOMPONENTTRANSFERElement) IfCLASS(condition bool, s ...string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFECOMPONENTTRANSFERElement) CLASSRemove(s ...string) *SVGFECOMPONENTTRANSFERElement {
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
func (e *SVGFECOMPONENTTRANSFERElement) STYLEF(k string, format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFECOMPONENTTRANSFERElement) IfSTYLE(condition bool, k string, v string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) STYLE(k string, v string) *SVGFECOMPONENTTRANSFERElement {
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

func (e *SVGFECOMPONENTTRANSFERElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFECOMPONENTTRANSFERElement) STYLEMap(m map[string]string) *SVGFECOMPONENTTRANSFERElement {
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
func (e *SVGFECOMPONENTTRANSFERElement) STYLEPairs(pairs ...string) *SVGFECOMPONENTTRANSFERElement {
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

func (e *SVGFECOMPONENTTRANSFERElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFECOMPONENTTRANSFERElement) STYLERemove(keys ...string) *SVGFECOMPONENTTRANSFERElement {
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

func (e *SVGFECOMPONENTTRANSFERElement) Z_REQ(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfZ_REQ(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFECOMPONENTTRANSFERElement) Z_REQRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFECOMPONENTTRANSFERElement) Z_TARGET(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfZ_TARGET(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFECOMPONENTTRANSFERElement) Z_TARGETRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFECOMPONENTTRANSFERElement) Z_REQ_SELECTOR(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFECOMPONENTTRANSFERElement) Z_REQ_SELECTORRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFECOMPONENTTRANSFERElement) Z_SWAP(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfZ_SWAP(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFECOMPONENTTRANSFERElement) Z_SWAPRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFECOMPONENTTRANSFERElement) Z_SWAP_PUSH(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFECOMPONENTTRANSFERElement) Z_SWAP_PUSHRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFECOMPONENTTRANSFERElement) Z_TRIGGER(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfZ_TRIGGER(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFECOMPONENTTRANSFERElement) Z_TRIGGERRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFECOMPONENTTRANSFERElement) Z_REQ_METHOD(c SVGFeComponentTransferZReqMethodChoice) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeComponentTransferZReqMethodChoice string

const (
	// default GET
	SVGFeComponentTransferZReqMethod_empty SVGFeComponentTransferZReqMethodChoice = ""
	// GET
	SVGFeComponentTransferZReqMethod_get SVGFeComponentTransferZReqMethodChoice = "get"
	// POST
	SVGFeComponentTransferZReqMethod_post SVGFeComponentTransferZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFECOMPONENTTRANSFERElement) Z_REQ_METHODRemove(c SVGFeComponentTransferZReqMethodChoice) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFECOMPONENTTRANSFERElement) Z_REQ_STRATEGY(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFECOMPONENTTRANSFERElement) Z_REQ_STRATEGYRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFECOMPONENTTRANSFERElement) Z_REQ_HISTORY(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFECOMPONENTTRANSFERElement) Z_REQ_HISTORYRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFECOMPONENTTRANSFERElement) Z_DATA(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfZ_DATA(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFECOMPONENTTRANSFERElement) Z_DATARemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFECOMPONENTTRANSFERElement) Z_JSON(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfZ_JSON(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFECOMPONENTTRANSFERElement) Z_JSONRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFECOMPONENTTRANSFERElement) Z_REQ_BATCH(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFECOMPONENTTRANSFERElement) Z_REQ_BATCHRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFECOMPONENTTRANSFERElement) Z_ACTION(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfZ_ACTION(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFECOMPONENTTRANSFERElement) Z_ACTIONRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFECOMPONENTTRANSFERElement) Z_REQ_BEFORE(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFECOMPONENTTRANSFERElement) Z_REQ_BEFORERemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFECOMPONENTTRANSFERElement) Z_REQ_AFTER(expression string) *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOMPONENTTRANSFERElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFECOMPONENTTRANSFERElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFECOMPONENTTRANSFERElement) Z_REQ_AFTERRemove() *SVGFECOMPONENTTRANSFERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
