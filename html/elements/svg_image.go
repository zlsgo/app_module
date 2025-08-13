package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <image> SVG element includes images inside SVG documents.
type SVGIMAGEElement struct {
	*Element
}

// Create a new SVGIMAGEElement element.
// This will create a new element with the tag
// "image" during rendering.
func SVG_IMAGE(children ...ElementRenderer) *SVGIMAGEElement {
	e := NewElement("image", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGIMAGEElement{Element: e}
}

func (e *SVGIMAGEElement) Children(children ...ElementRenderer) *SVGIMAGEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGIMAGEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGIMAGEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGIMAGEElement) Attr(name string, value ...string) *SVGIMAGEElement {
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

func (e *SVGIMAGEElement) Attrs(attrs ...string) *SVGIMAGEElement {
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

func (e *SVGIMAGEElement) AttrsMap(attrs map[string]string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGIMAGEElement) Text(text string) *SVGIMAGEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGIMAGEElement) TextF(format string, args ...any) *SVGIMAGEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGIMAGEElement) IfText(condition bool, text string) *SVGIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGIMAGEElement) IfTextF(condition bool, format string, args ...any) *SVGIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGIMAGEElement) Escaped(text string) *SVGIMAGEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGIMAGEElement) IfEscaped(condition bool, text string) *SVGIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGIMAGEElement) EscapedF(format string, args ...any) *SVGIMAGEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGIMAGEElement) IfEscapedF(condition bool, format string, args ...any) *SVGIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGIMAGEElement) CustomData(key, value string) *SVGIMAGEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGIMAGEElement) IfCustomData(condition bool, key, value string) *SVGIMAGEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGIMAGEElement) CustomDataF(key, format string, args ...any) *SVGIMAGEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGIMAGEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGIMAGEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGIMAGEElement) CustomDataRemove(key string) *SVGIMAGEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// Indicates how the fetched image is fitted into the destination rectangle.
func (e *SVGIMAGEElement) PRESERVE_ASPECT_RATIO(c SVGImagePreserveAspectRatioChoice) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("preserveAspectRatio", string(c))
	return e
}

type SVGImagePreserveAspectRatioChoice string

const (
	// Do not force uniform scaling.
	SVGImagePreserveAspectRatio_none SVGImagePreserveAspectRatioChoice = "none"
	// Scale the image to the smallest size such that both its width and its height
	// can completely fit inside the corresponding dimension of the viewPort.
	SVGImagePreserveAspectRatio_xMinYMin SVGImagePreserveAspectRatioChoice = "xMinYMin"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGImagePreserveAspectRatio_xMidYMin SVGImagePreserveAspectRatioChoice = "xMidYMin"
	// Align the image with the corresponding side of the viewPort.
	SVGImagePreserveAspectRatio_xMaxYMin SVGImagePreserveAspectRatioChoice = "xMaxYMin"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGImagePreserveAspectRatio_xMinYMid SVGImagePreserveAspectRatioChoice = "xMinYMid"
	// Scale the image to the smallest size such that it can completely fit inside the
	// corresponding dimension of the viewPort.
	SVGImagePreserveAspectRatio_xMidYMid SVGImagePreserveAspectRatioChoice = "xMidYMid"
	// Align the image with the corresponding side of the viewPort.
	SVGImagePreserveAspectRatio_xMaxYMid SVGImagePreserveAspectRatioChoice = "xMaxYMid"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGImagePreserveAspectRatio_xMinYMax SVGImagePreserveAspectRatioChoice = "xMinYMax"
	// Align the image with the corresponding side of the viewPort.
	SVGImagePreserveAspectRatio_xMidYMax SVGImagePreserveAspectRatioChoice = "xMidYMax"
	// Scale the image to the smallest size such that both its width and its height
	// can completely fit inside the corresponding dimension of the viewPort.
	SVGImagePreserveAspectRatio_xMaxYMax SVGImagePreserveAspectRatioChoice = "xMaxYMax"
)

// Remove the attribute PRESERVE_ASPECT_RATIO from the element.
func (e *SVGIMAGEElement) PRESERVE_ASPECT_RATIORemove(c SVGImagePreserveAspectRatioChoice) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("preserveAspectRatio")
	return e
}

// The x-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGIMAGEElement) X(f float64) *SVGIMAGEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("x", f)
	return e
}

func (e *SVGIMAGEElement) IfX(condition bool, f float64) *SVGIMAGEElement {
	if condition {
		e.X(f)
	}
	return e
}

// The y-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGIMAGEElement) Y(f float64) *SVGIMAGEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("y", f)
	return e
}

func (e *SVGIMAGEElement) IfY(condition bool, f float64) *SVGIMAGEElement {
	if condition {
		e.Y(f)
	}
	return e
}

// The width of the rectangular region.
func (e *SVGIMAGEElement) WIDTH(f float64) *SVGIMAGEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("width", f)
	return e
}

func (e *SVGIMAGEElement) IfWIDTH(condition bool, f float64) *SVGIMAGEElement {
	if condition {
		e.WIDTH(f)
	}
	return e
}

// The height of the rectangular region.
func (e *SVGIMAGEElement) HEIGHT(f float64) *SVGIMAGEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("height", f)
	return e
}

func (e *SVGIMAGEElement) IfHEIGHT(condition bool, f float64) *SVGIMAGEElement {
	if condition {
		e.HEIGHT(f)
	}
	return e
}

// A URI reference to the image to embed.
func (e *SVGIMAGEElement) HREF(s string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("href", s)
	return e
}

func (e *SVGIMAGEElement) HREFF(format string, args ...any) *SVGIMAGEElement {
	return e.HREF(fmt.Sprintf(format, args...))
}

func (e *SVGIMAGEElement) IfHREF(condition bool, s string) *SVGIMAGEElement {
	if condition {
		e.HREF(s)
	}
	return e
}

func (e *SVGIMAGEElement) IfHREFF(condition bool, format string, args ...any) *SVGIMAGEElement {
	if condition {
		e.HREF(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HREF from the element.
func (e *SVGIMAGEElement) HREFRemove(s string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("href")
	return e
}

func (e *SVGIMAGEElement) HREFRemoveF(format string, args ...any) *SVGIMAGEElement {
	return e.HREFRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGIMAGEElement) ID(s string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGIMAGEElement) IDF(format string, args ...any) *SVGIMAGEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGIMAGEElement) IfID(condition bool, s string) *SVGIMAGEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGIMAGEElement) IfIDF(condition bool, format string, args ...any) *SVGIMAGEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGIMAGEElement) IDRemove(s string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGIMAGEElement) IDRemoveF(format string, args ...any) *SVGIMAGEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGIMAGEElement) CLASS(s ...string) *SVGIMAGEElement {
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

func (e *SVGIMAGEElement) IfCLASS(condition bool, s ...string) *SVGIMAGEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGIMAGEElement) CLASSRemove(s ...string) *SVGIMAGEElement {
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
func (e *SVGIMAGEElement) STYLEF(k string, format string, args ...any) *SVGIMAGEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGIMAGEElement) IfSTYLE(condition bool, k string, v string) *SVGIMAGEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGIMAGEElement) STYLE(k string, v string) *SVGIMAGEElement {
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

func (e *SVGIMAGEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGIMAGEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGIMAGEElement) STYLEMap(m map[string]string) *SVGIMAGEElement {
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
func (e *SVGIMAGEElement) STYLEPairs(pairs ...string) *SVGIMAGEElement {
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

func (e *SVGIMAGEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGIMAGEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGIMAGEElement) STYLERemove(keys ...string) *SVGIMAGEElement {
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

func (e *SVGIMAGEElement) Z_REQ(expression string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGIMAGEElement) IfZ_REQ(condition bool, expression string) *SVGIMAGEElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGIMAGEElement) Z_REQRemove() *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGIMAGEElement) Z_TARGET(expression string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGIMAGEElement) IfZ_TARGET(condition bool, expression string) *SVGIMAGEElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGIMAGEElement) Z_TARGETRemove() *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGIMAGEElement) Z_REQ_SELECTOR(expression string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGIMAGEElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGIMAGEElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGIMAGEElement) Z_REQ_SELECTORRemove() *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGIMAGEElement) Z_SWAP(expression string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGIMAGEElement) IfZ_SWAP(condition bool, expression string) *SVGIMAGEElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGIMAGEElement) Z_SWAPRemove() *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGIMAGEElement) Z_SWAP_PUSH(expression string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGIMAGEElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGIMAGEElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGIMAGEElement) Z_SWAP_PUSHRemove() *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGIMAGEElement) Z_TRIGGER(expression string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGIMAGEElement) IfZ_TRIGGER(condition bool, expression string) *SVGIMAGEElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGIMAGEElement) Z_TRIGGERRemove() *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGIMAGEElement) Z_REQ_METHOD(c SVGImageZReqMethodChoice) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGImageZReqMethodChoice string

const (
	// default GET
	SVGImageZReqMethod_empty SVGImageZReqMethodChoice = ""
	// GET
	SVGImageZReqMethod_get SVGImageZReqMethodChoice = "get"
	// POST
	SVGImageZReqMethod_post SVGImageZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGIMAGEElement) Z_REQ_METHODRemove(c SVGImageZReqMethodChoice) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGIMAGEElement) Z_REQ_STRATEGY(expression string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGIMAGEElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGIMAGEElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGIMAGEElement) Z_REQ_STRATEGYRemove() *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGIMAGEElement) Z_REQ_HISTORY(expression string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGIMAGEElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGIMAGEElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGIMAGEElement) Z_REQ_HISTORYRemove() *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGIMAGEElement) Z_DATA(expression string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGIMAGEElement) IfZ_DATA(condition bool, expression string) *SVGIMAGEElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGIMAGEElement) Z_DATARemove() *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGIMAGEElement) Z_JSON(expression string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGIMAGEElement) IfZ_JSON(condition bool, expression string) *SVGIMAGEElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGIMAGEElement) Z_JSONRemove() *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGIMAGEElement) Z_REQ_BATCH(expression string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGIMAGEElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGIMAGEElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGIMAGEElement) Z_REQ_BATCHRemove() *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGIMAGEElement) Z_ACTION(expression string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGIMAGEElement) IfZ_ACTION(condition bool, expression string) *SVGIMAGEElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGIMAGEElement) Z_ACTIONRemove() *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGIMAGEElement) Z_REQ_BEFORE(expression string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGIMAGEElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGIMAGEElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGIMAGEElement) Z_REQ_BEFORERemove() *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGIMAGEElement) Z_REQ_AFTER(expression string) *SVGIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGIMAGEElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGIMAGEElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGIMAGEElement) Z_REQ_AFTERRemove() *SVGIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
