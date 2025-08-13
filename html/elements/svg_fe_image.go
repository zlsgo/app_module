package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feImage> SVG filter primitive fetches image data from an external source
// and provides the pixel data as output (meaning if the external source is an SVG
// image, it is rasterized.)
type SVGFEIMAGEElement struct {
	*Element
}

// Create a new SVGFEIMAGEElement element.
// This will create a new element with the tag
// "feImage" during rendering.
func SVG_FEIMAGE(children ...ElementRenderer) *SVGFEIMAGEElement {
	e := NewElement("feImage", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEIMAGEElement{Element: e}
}

func (e *SVGFEIMAGEElement) Children(children ...ElementRenderer) *SVGFEIMAGEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEIMAGEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEIMAGEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEIMAGEElement) Attr(name string, value ...string) *SVGFEIMAGEElement {
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

func (e *SVGFEIMAGEElement) Attrs(attrs ...string) *SVGFEIMAGEElement {
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

func (e *SVGFEIMAGEElement) AttrsMap(attrs map[string]string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEIMAGEElement) Text(text string) *SVGFEIMAGEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEIMAGEElement) TextF(format string, args ...any) *SVGFEIMAGEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEIMAGEElement) IfText(condition bool, text string) *SVGFEIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEIMAGEElement) IfTextF(condition bool, format string, args ...any) *SVGFEIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEIMAGEElement) Escaped(text string) *SVGFEIMAGEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEIMAGEElement) IfEscaped(condition bool, text string) *SVGFEIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEIMAGEElement) EscapedF(format string, args ...any) *SVGFEIMAGEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEIMAGEElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEIMAGEElement) CustomData(key, value string) *SVGFEIMAGEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEIMAGEElement) IfCustomData(condition bool, key, value string) *SVGFEIMAGEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEIMAGEElement) CustomDataF(key, format string, args ...any) *SVGFEIMAGEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEIMAGEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEIMAGEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEIMAGEElement) CustomDataRemove(key string) *SVGFEIMAGEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// Indicates whether or not to force synchronous behavior.
func (e *SVGFEIMAGEElement) EXTERNAL_RESOURCES_REQUIRED() *SVGFEIMAGEElement {
	e.EXTERNAL_RESOURCES_REQUIREDSet(true)
	return e
}

func (e *SVGFEIMAGEElement) IfEXTERNAL_RESOURCES_REQUIRED(condition bool) *SVGFEIMAGEElement {
	if condition {
		e.EXTERNAL_RESOURCES_REQUIREDSet(true)
	}
	return e
}

// Set the attribute EXTERNAL_RESOURCES_REQUIRED to the value b explicitly.
func (e *SVGFEIMAGEElement) EXTERNAL_RESOURCES_REQUIREDSet(b bool) *SVGFEIMAGEElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = zarray.NewSortMap[string, bool]()
	}
	e.BoolAttributes.Set("externalResourcesRequired", b)
	return e
}

func (e *SVGFEIMAGEElement) IfSetEXTERNAL_RESOURCES_REQUIRED(condition bool, b bool) *SVGFEIMAGEElement {
	if condition {
		e.EXTERNAL_RESOURCES_REQUIREDSet(b)
	}
	return e
}

// Remove the attribute EXTERNAL_RESOURCES_REQUIRED from the element.
func (e *SVGFEIMAGEElement) EXTERNAL_RESOURCES_REQUIREDRemove(b bool) *SVGFEIMAGEElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Delete("externalResourcesRequired")
	return e
}

// Indicates how the fetched image is fitted into the destination rectangle.
func (e *SVGFEIMAGEElement) PRESERVE_ASPECT_RATIO(c SVGFeImagePreserveAspectRatioChoice) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("preserveAspectRatio", string(c))
	return e
}

type SVGFeImagePreserveAspectRatioChoice string

const (
	// Do not force uniform scaling.
	SVGFeImagePreserveAspectRatio_none SVGFeImagePreserveAspectRatioChoice = "none"
	// Scale the image to the smallest size such that both its width and its height
	// can completely fit inside the corresponding dimension of the viewPort.
	SVGFeImagePreserveAspectRatio_xMinYMin SVGFeImagePreserveAspectRatioChoice = "xMinYMin"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGFeImagePreserveAspectRatio_xMidYMin SVGFeImagePreserveAspectRatioChoice = "xMidYMin"
	// Align the image with the corresponding side of the viewPort.
	SVGFeImagePreserveAspectRatio_xMaxYMin SVGFeImagePreserveAspectRatioChoice = "xMaxYMin"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGFeImagePreserveAspectRatio_xMinYMid SVGFeImagePreserveAspectRatioChoice = "xMinYMid"
	// Scale the image to the smallest size such that it can completely fit inside the
	// corresponding dimension of the viewPort.
	SVGFeImagePreserveAspectRatio_xMidYMid SVGFeImagePreserveAspectRatioChoice = "xMidYMid"
	// Align the image with the corresponding side of the viewPort.
	SVGFeImagePreserveAspectRatio_xMaxYMid SVGFeImagePreserveAspectRatioChoice = "xMaxYMid"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGFeImagePreserveAspectRatio_xMinYMax SVGFeImagePreserveAspectRatioChoice = "xMinYMax"
	// Align the image with the corresponding side of the viewPort.
	SVGFeImagePreserveAspectRatio_xMidYMax SVGFeImagePreserveAspectRatioChoice = "xMidYMax"
	// Scale the image to the smallest size such that both its width and its height
	// can completely fit inside the corresponding dimension of the viewPort.
	SVGFeImagePreserveAspectRatio_xMaxYMax SVGFeImagePreserveAspectRatioChoice = "xMaxYMax"
)

// Remove the attribute PRESERVE_ASPECT_RATIO from the element.
func (e *SVGFEIMAGEElement) PRESERVE_ASPECT_RATIORemove(c SVGFeImagePreserveAspectRatioChoice) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("preserveAspectRatio")
	return e
}

// A URI reference to an external resource.
func (e *SVGFEIMAGEElement) HREF(s string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("href", s)
	return e
}

func (e *SVGFEIMAGEElement) HREFF(format string, args ...any) *SVGFEIMAGEElement {
	return e.HREF(fmt.Sprintf(format, args...))
}

func (e *SVGFEIMAGEElement) IfHREF(condition bool, s string) *SVGFEIMAGEElement {
	if condition {
		e.HREF(s)
	}
	return e
}

func (e *SVGFEIMAGEElement) IfHREFF(condition bool, format string, args ...any) *SVGFEIMAGEElement {
	if condition {
		e.HREF(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HREF from the element.
func (e *SVGFEIMAGEElement) HREFRemove(s string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("href")
	return e
}

func (e *SVGFEIMAGEElement) HREFRemoveF(format string, args ...any) *SVGFEIMAGEElement {
	return e.HREFRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGFEIMAGEElement) ID(s string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEIMAGEElement) IDF(format string, args ...any) *SVGFEIMAGEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEIMAGEElement) IfID(condition bool, s string) *SVGFEIMAGEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEIMAGEElement) IfIDF(condition bool, format string, args ...any) *SVGFEIMAGEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEIMAGEElement) IDRemove(s string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEIMAGEElement) IDRemoveF(format string, args ...any) *SVGFEIMAGEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEIMAGEElement) CLASS(s ...string) *SVGFEIMAGEElement {
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

func (e *SVGFEIMAGEElement) IfCLASS(condition bool, s ...string) *SVGFEIMAGEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEIMAGEElement) CLASSRemove(s ...string) *SVGFEIMAGEElement {
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
func (e *SVGFEIMAGEElement) STYLEF(k string, format string, args ...any) *SVGFEIMAGEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEIMAGEElement) IfSTYLE(condition bool, k string, v string) *SVGFEIMAGEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEIMAGEElement) STYLE(k string, v string) *SVGFEIMAGEElement {
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

func (e *SVGFEIMAGEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEIMAGEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEIMAGEElement) STYLEMap(m map[string]string) *SVGFEIMAGEElement {
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
func (e *SVGFEIMAGEElement) STYLEPairs(pairs ...string) *SVGFEIMAGEElement {
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

func (e *SVGFEIMAGEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEIMAGEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEIMAGEElement) STYLERemove(keys ...string) *SVGFEIMAGEElement {
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

func (e *SVGFEIMAGEElement) Z_REQ(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfZ_REQ(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEIMAGEElement) Z_REQRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEIMAGEElement) Z_TARGET(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfZ_TARGET(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEIMAGEElement) Z_TARGETRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEIMAGEElement) Z_REQ_SELECTOR(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEIMAGEElement) Z_REQ_SELECTORRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEIMAGEElement) Z_SWAP(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfZ_SWAP(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEIMAGEElement) Z_SWAPRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEIMAGEElement) Z_SWAP_PUSH(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEIMAGEElement) Z_SWAP_PUSHRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEIMAGEElement) Z_TRIGGER(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEIMAGEElement) Z_TRIGGERRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEIMAGEElement) Z_REQ_METHOD(c SVGFeImageZReqMethodChoice) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeImageZReqMethodChoice string

const (
	// default GET
	SVGFeImageZReqMethod_empty SVGFeImageZReqMethodChoice = ""
	// GET
	SVGFeImageZReqMethod_get SVGFeImageZReqMethodChoice = "get"
	// POST
	SVGFeImageZReqMethod_post SVGFeImageZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEIMAGEElement) Z_REQ_METHODRemove(c SVGFeImageZReqMethodChoice) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEIMAGEElement) Z_REQ_STRATEGY(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEIMAGEElement) Z_REQ_STRATEGYRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEIMAGEElement) Z_REQ_HISTORY(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEIMAGEElement) Z_REQ_HISTORYRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEIMAGEElement) Z_DATA(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfZ_DATA(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEIMAGEElement) Z_DATARemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEIMAGEElement) Z_JSON(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfZ_JSON(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEIMAGEElement) Z_JSONRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEIMAGEElement) Z_REQ_BATCH(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEIMAGEElement) Z_REQ_BATCHRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEIMAGEElement) Z_ACTION(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfZ_ACTION(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEIMAGEElement) Z_ACTIONRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEIMAGEElement) Z_REQ_BEFORE(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEIMAGEElement) Z_REQ_BEFORERemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEIMAGEElement) Z_REQ_AFTER(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEIMAGEElement) Z_REQ_AFTERRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
