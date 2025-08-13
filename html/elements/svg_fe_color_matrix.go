package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feColorMatrix> SVG filter element changes colors based on a transformation
// matrix
// Every pixel's color value (represented by an [R,G,B,A] vector) is matrix
// multiplied to create a new color.
type SVGFECOLORMATRIXElement struct {
	*Element
}

// Create a new SVGFECOLORMATRIXElement element.
// This will create a new element with the tag
// "feColorMatrix" during rendering.
func SVG_FECOLORMATRIX(children ...ElementRenderer) *SVGFECOLORMATRIXElement {
	e := NewElement("feColorMatrix", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFECOLORMATRIXElement{Element: e}
}

func (e *SVGFECOLORMATRIXElement) Children(children ...ElementRenderer) *SVGFECOLORMATRIXElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFECOLORMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFECOLORMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) Attr(name string, value ...string) *SVGFECOLORMATRIXElement {
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

func (e *SVGFECOLORMATRIXElement) Attrs(attrs ...string) *SVGFECOLORMATRIXElement {
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

func (e *SVGFECOLORMATRIXElement) AttrsMap(attrs map[string]string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) Text(text string) *SVGFECOLORMATRIXElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFECOLORMATRIXElement) TextF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFECOLORMATRIXElement) IfText(condition bool, text string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) IfTextF(condition bool, format string, args ...any) *SVGFECOLORMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) Escaped(text string) *SVGFECOLORMATRIXElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFECOLORMATRIXElement) IfEscaped(condition bool, text string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) EscapedF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFECOLORMATRIXElement) IfEscapedF(condition bool, format string, args ...any) *SVGFECOLORMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) CustomData(key, value string) *SVGFECOLORMATRIXElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfCustomData(condition bool, key, value string) *SVGFECOLORMATRIXElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) CustomDataF(key, format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFECOLORMATRIXElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFECOLORMATRIXElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) CustomDataRemove(key string) *SVGFECOLORMATRIXElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The input for this filter.
func (e *SVGFECOLORMATRIXElement) IN(s string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFECOLORMATRIXElement) INF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.IN(fmt.Sprintf(format, args...))
}

func (e *SVGFECOLORMATRIXElement) IfIN(condition bool, s string) *SVGFECOLORMATRIXElement {
	if condition {
		e.IN(s)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) IfINF(condition bool, format string, args ...any) *SVGFECOLORMATRIXElement {
	if condition {
		e.IN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFECOLORMATRIXElement) INRemove(s string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("in")
	return e
}

func (e *SVGFECOLORMATRIXElement) INRemoveF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.INRemove(fmt.Sprintf(format, args...))
}

// The type of matrix operation.
func (e *SVGFECOLORMATRIXElement) TYPE(c SVGFeColorMatrixTypeChoice) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("type", string(c))
	return e
}

type SVGFeColorMatrixTypeChoice string

const (
	// The type of matrix operation.
	SVGFeColorMatrixType_matrix SVGFeColorMatrixTypeChoice = "matrix"
	// The type of matrix operation.
	SVGFeColorMatrixType_saturate SVGFeColorMatrixTypeChoice = "saturate"
	// The type of matrix operation.
	SVGFeColorMatrixType_hueRotate SVGFeColorMatrixTypeChoice = "hueRotate"
	// The type of matrix operation.
	SVGFeColorMatrixType_luminanceToAlpha SVGFeColorMatrixTypeChoice = "luminanceToAlpha"
)

// Remove the attribute TYPE from the element.
func (e *SVGFECOLORMATRIXElement) TYPERemove(c SVGFeColorMatrixTypeChoice) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("type")
	return e
}

// The list of one or more numbers that represent the matrix.
func (e *SVGFECOLORMATRIXElement) VALUES(s string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("values", s)
	return e
}

func (e *SVGFECOLORMATRIXElement) VALUESF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.VALUES(fmt.Sprintf(format, args...))
}

func (e *SVGFECOLORMATRIXElement) IfVALUES(condition bool, s string) *SVGFECOLORMATRIXElement {
	if condition {
		e.VALUES(s)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) IfVALUESF(condition bool, format string, args ...any) *SVGFECOLORMATRIXElement {
	if condition {
		e.VALUES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute VALUES from the element.
func (e *SVGFECOLORMATRIXElement) VALUESRemove(s string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("values")
	return e
}

func (e *SVGFECOLORMATRIXElement) VALUESRemoveF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.VALUESRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGFECOLORMATRIXElement) ID(s string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFECOLORMATRIXElement) IDF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFECOLORMATRIXElement) IfID(condition bool, s string) *SVGFECOLORMATRIXElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) IfIDF(condition bool, format string, args ...any) *SVGFECOLORMATRIXElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFECOLORMATRIXElement) IDRemove(s string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFECOLORMATRIXElement) IDRemoveF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFECOLORMATRIXElement) CLASS(s ...string) *SVGFECOLORMATRIXElement {
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

func (e *SVGFECOLORMATRIXElement) IfCLASS(condition bool, s ...string) *SVGFECOLORMATRIXElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFECOLORMATRIXElement) CLASSRemove(s ...string) *SVGFECOLORMATRIXElement {
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
func (e *SVGFECOLORMATRIXElement) STYLEF(k string, format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFECOLORMATRIXElement) IfSTYLE(condition bool, k string, v string) *SVGFECOLORMATRIXElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) STYLE(k string, v string) *SVGFECOLORMATRIXElement {
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

func (e *SVGFECOLORMATRIXElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFECOLORMATRIXElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFECOLORMATRIXElement) STYLEMap(m map[string]string) *SVGFECOLORMATRIXElement {
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
func (e *SVGFECOLORMATRIXElement) STYLEPairs(pairs ...string) *SVGFECOLORMATRIXElement {
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

func (e *SVGFECOLORMATRIXElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFECOLORMATRIXElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFECOLORMATRIXElement) STYLERemove(keys ...string) *SVGFECOLORMATRIXElement {
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

func (e *SVGFECOLORMATRIXElement) Z_REQ(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfZ_REQ(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFECOLORMATRIXElement) Z_REQRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFECOLORMATRIXElement) Z_TARGET(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfZ_TARGET(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFECOLORMATRIXElement) Z_TARGETRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFECOLORMATRIXElement) Z_REQ_SELECTOR(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFECOLORMATRIXElement) Z_REQ_SELECTORRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFECOLORMATRIXElement) Z_SWAP(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfZ_SWAP(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFECOLORMATRIXElement) Z_SWAPRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFECOLORMATRIXElement) Z_SWAP_PUSH(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFECOLORMATRIXElement) Z_SWAP_PUSHRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFECOLORMATRIXElement) Z_TRIGGER(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfZ_TRIGGER(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFECOLORMATRIXElement) Z_TRIGGERRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFECOLORMATRIXElement) Z_REQ_METHOD(c SVGFeColorMatrixZReqMethodChoice) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeColorMatrixZReqMethodChoice string

const (
	// default GET
	SVGFeColorMatrixZReqMethod_empty SVGFeColorMatrixZReqMethodChoice = ""
	// GET
	SVGFeColorMatrixZReqMethod_get SVGFeColorMatrixZReqMethodChoice = "get"
	// POST
	SVGFeColorMatrixZReqMethod_post SVGFeColorMatrixZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFECOLORMATRIXElement) Z_REQ_METHODRemove(c SVGFeColorMatrixZReqMethodChoice) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFECOLORMATRIXElement) Z_REQ_STRATEGY(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFECOLORMATRIXElement) Z_REQ_STRATEGYRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFECOLORMATRIXElement) Z_REQ_HISTORY(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFECOLORMATRIXElement) Z_REQ_HISTORYRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFECOLORMATRIXElement) Z_DATA(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfZ_DATA(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFECOLORMATRIXElement) Z_DATARemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFECOLORMATRIXElement) Z_JSON(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfZ_JSON(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFECOLORMATRIXElement) Z_JSONRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFECOLORMATRIXElement) Z_REQ_BATCH(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFECOLORMATRIXElement) Z_REQ_BATCHRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFECOLORMATRIXElement) Z_ACTION(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfZ_ACTION(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFECOLORMATRIXElement) Z_ACTIONRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFECOLORMATRIXElement) Z_REQ_BEFORE(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFECOLORMATRIXElement) Z_REQ_BEFORERemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFECOLORMATRIXElement) Z_REQ_AFTER(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFECOLORMATRIXElement) Z_REQ_AFTERRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
