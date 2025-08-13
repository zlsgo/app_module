package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <svg> element is a container that defines a new coordinate system and
// viewport
// It is used as the outermost element of SVG documents, but it can also be used
// to embed a SVG fragment inside an SVG or HTML document.
type SVGSVGElement struct {
	*Element
}

// Create a new SVGSVGElement element.
// This will create a new element with the tag
// "svg" during rendering.
func SVG_SVG(children ...ElementRenderer) *SVGSVGElement {
	e := NewElement("svg", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGSVGElement{Element: e}
}

func (e *SVGSVGElement) Children(children ...ElementRenderer) *SVGSVGElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGSVGElement) IfChildren(condition bool, children ...ElementRenderer) *SVGSVGElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGSVGElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGSVGElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGSVGElement) Attr(name string, value ...string) *SVGSVGElement {
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

func (e *SVGSVGElement) Attrs(attrs ...string) *SVGSVGElement {
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

func (e *SVGSVGElement) AttrsMap(attrs map[string]string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGSVGElement) Text(text string) *SVGSVGElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGSVGElement) TextF(format string, args ...any) *SVGSVGElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGSVGElement) IfText(condition bool, text string) *SVGSVGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGSVGElement) IfTextF(condition bool, format string, args ...any) *SVGSVGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGSVGElement) Escaped(text string) *SVGSVGElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGSVGElement) IfEscaped(condition bool, text string) *SVGSVGElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGSVGElement) EscapedF(format string, args ...any) *SVGSVGElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGSVGElement) IfEscapedF(condition bool, format string, args ...any) *SVGSVGElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGSVGElement) CustomData(key, value string) *SVGSVGElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGSVGElement) IfCustomData(condition bool, key, value string) *SVGSVGElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGSVGElement) CustomDataF(key, format string, args ...any) *SVGSVGElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGSVGElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGSVGElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGSVGElement) CustomDataRemove(key string) *SVGSVGElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The x-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGSVGElement) X(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("x", s)
	return e
}

func (e *SVGSVGElement) XF(format string, args ...any) *SVGSVGElement {
	return e.X(fmt.Sprintf(format, args...))
}

func (e *SVGSVGElement) IfX(condition bool, s string) *SVGSVGElement {
	if condition {
		e.X(s)
	}
	return e
}

func (e *SVGSVGElement) IfXF(condition bool, format string, args ...any) *SVGSVGElement {
	if condition {
		e.X(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute X from the element.
func (e *SVGSVGElement) XRemove(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("x")
	return e
}

func (e *SVGSVGElement) XRemoveF(format string, args ...any) *SVGSVGElement {
	return e.XRemove(fmt.Sprintf(format, args...))
}

// The y-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGSVGElement) Y(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("y", s)
	return e
}

func (e *SVGSVGElement) YF(format string, args ...any) *SVGSVGElement {
	return e.Y(fmt.Sprintf(format, args...))
}

func (e *SVGSVGElement) IfY(condition bool, s string) *SVGSVGElement {
	if condition {
		e.Y(s)
	}
	return e
}

func (e *SVGSVGElement) IfYF(condition bool, format string, args ...any) *SVGSVGElement {
	if condition {
		e.Y(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute Y from the element.
func (e *SVGSVGElement) YRemove(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("y")
	return e
}

func (e *SVGSVGElement) YRemoveF(format string, args ...any) *SVGSVGElement {
	return e.YRemove(fmt.Sprintf(format, args...))
}

// The width of the rectangular region.
func (e *SVGSVGElement) WIDTH(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("width", s)
	return e
}

func (e *SVGSVGElement) WIDTHF(format string, args ...any) *SVGSVGElement {
	return e.WIDTH(fmt.Sprintf(format, args...))
}

func (e *SVGSVGElement) IfWIDTH(condition bool, s string) *SVGSVGElement {
	if condition {
		e.WIDTH(s)
	}
	return e
}

func (e *SVGSVGElement) IfWIDTHF(condition bool, format string, args ...any) *SVGSVGElement {
	if condition {
		e.WIDTH(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute WIDTH from the element.
func (e *SVGSVGElement) WIDTHRemove(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("width")
	return e
}

func (e *SVGSVGElement) WIDTHRemoveF(format string, args ...any) *SVGSVGElement {
	return e.WIDTHRemove(fmt.Sprintf(format, args...))
}

// The height of the rectangular region.
func (e *SVGSVGElement) HEIGHT(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("height", s)
	return e
}

func (e *SVGSVGElement) HEIGHTF(format string, args ...any) *SVGSVGElement {
	return e.HEIGHT(fmt.Sprintf(format, args...))
}

func (e *SVGSVGElement) IfHEIGHT(condition bool, s string) *SVGSVGElement {
	if condition {
		e.HEIGHT(s)
	}
	return e
}

func (e *SVGSVGElement) IfHEIGHTF(condition bool, format string, args ...any) *SVGSVGElement {
	if condition {
		e.HEIGHT(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HEIGHT from the element.
func (e *SVGSVGElement) HEIGHTRemove(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("height")
	return e
}

func (e *SVGSVGElement) HEIGHTRemoveF(format string, args ...any) *SVGSVGElement {
	return e.HEIGHTRemove(fmt.Sprintf(format, args...))
}

// The position and size of the viewport (the viewBox) is defined by the viewBox
// attribute.
func (e *SVGSVGElement) VIEW_BOX(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("viewBox", s)
	return e
}

func (e *SVGSVGElement) VIEW_BOXF(format string, args ...any) *SVGSVGElement {
	return e.VIEW_BOX(fmt.Sprintf(format, args...))
}

func (e *SVGSVGElement) IfVIEW_BOX(condition bool, s string) *SVGSVGElement {
	if condition {
		e.VIEW_BOX(s)
	}
	return e
}

func (e *SVGSVGElement) IfVIEW_BOXF(condition bool, format string, args ...any) *SVGSVGElement {
	if condition {
		e.VIEW_BOX(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute VIEW_BOX from the element.
func (e *SVGSVGElement) VIEW_BOXRemove(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("viewBox")
	return e
}

func (e *SVGSVGElement) VIEW_BOXRemoveF(format string, args ...any) *SVGSVGElement {
	return e.VIEW_BOXRemove(fmt.Sprintf(format, args...))
}

// Indicates how the viewport is fitted to the rectangle of the given width and
// height.
func (e *SVGSVGElement) PRESERVE_ASPECT_RATIO(c SVGSvgPreserveAspectRatioChoice) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("preserveAspectRatio", string(c))
	return e
}

type SVGSvgPreserveAspectRatioChoice string

const (
	// Do not force uniform scaling.
	SVGSvgPreserveAspectRatio_none SVGSvgPreserveAspectRatioChoice = "none"
	// Scale the image to the smallest size such that both its width and its height
	// can completely fit inside the corresponding dimension of the viewPort.
	SVGSvgPreserveAspectRatio_xMinYMin SVGSvgPreserveAspectRatioChoice = "xMinYMin"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGSvgPreserveAspectRatio_xMidYMin SVGSvgPreserveAspectRatioChoice = "xMidYMin"
	// Align the image with the corresponding side of the viewPort.
	SVGSvgPreserveAspectRatio_xMaxYMin SVGSvgPreserveAspectRatioChoice = "xMaxYMin"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGSvgPreserveAspectRatio_xMinYMid SVGSvgPreserveAspectRatioChoice = "xMinYMid"
	// Scale the image to the smallest size such that it can completely fit inside the
	// corresponding dimension of the viewPort.
	SVGSvgPreserveAspectRatio_xMidYMid SVGSvgPreserveAspectRatioChoice = "xMidYMid"
	// Align the image with the corresponding side of the viewPort.
	SVGSvgPreserveAspectRatio_xMaxYMid SVGSvgPreserveAspectRatioChoice = "xMaxYMid"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGSvgPreserveAspectRatio_xMinYMax SVGSvgPreserveAspectRatioChoice = "xMinYMax"
	// Align the image with the corresponding side of the viewPort.
	SVGSvgPreserveAspectRatio_xMidYMax SVGSvgPreserveAspectRatioChoice = "xMidYMax"
	// Scale the image to the smallest size such that both its width and its height
	// can completely fit inside the corresponding dimension of the viewPort.
	SVGSvgPreserveAspectRatio_xMaxYMax SVGSvgPreserveAspectRatioChoice = "xMaxYMax"
)

// Remove the attribute PRESERVE_ASPECT_RATIO from the element.
func (e *SVGSVGElement) PRESERVE_ASPECT_RATIORemove(c SVGSvgPreserveAspectRatioChoice) *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("preserveAspectRatio")
	return e
}

// Specifies a unique id for an element
func (e *SVGSVGElement) ID(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGSVGElement) IDF(format string, args ...any) *SVGSVGElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGSVGElement) IfID(condition bool, s string) *SVGSVGElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGSVGElement) IfIDF(condition bool, format string, args ...any) *SVGSVGElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGSVGElement) IDRemove(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGSVGElement) IDRemoveF(format string, args ...any) *SVGSVGElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGSVGElement) CLASS(s ...string) *SVGSVGElement {
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

func (e *SVGSVGElement) IfCLASS(condition bool, s ...string) *SVGSVGElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGSVGElement) CLASSRemove(s ...string) *SVGSVGElement {
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
func (e *SVGSVGElement) STYLEF(k string, format string, args ...any) *SVGSVGElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGSVGElement) IfSTYLE(condition bool, k string, v string) *SVGSVGElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGSVGElement) STYLE(k string, v string) *SVGSVGElement {
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

func (e *SVGSVGElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGSVGElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGSVGElement) STYLEMap(m map[string]string) *SVGSVGElement {
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
func (e *SVGSVGElement) STYLEPairs(pairs ...string) *SVGSVGElement {
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

func (e *SVGSVGElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGSVGElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGSVGElement) STYLERemove(keys ...string) *SVGSVGElement {
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

func (e *SVGSVGElement) Z_REQ(expression string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSVGElement) IfZ_REQ(condition bool, expression string) *SVGSVGElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGSVGElement) Z_REQRemove() *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGSVGElement) Z_TARGET(expression string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSVGElement) IfZ_TARGET(condition bool, expression string) *SVGSVGElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGSVGElement) Z_TARGETRemove() *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGSVGElement) Z_REQ_SELECTOR(expression string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSVGElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGSVGElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGSVGElement) Z_REQ_SELECTORRemove() *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGSVGElement) Z_SWAP(expression string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSVGElement) IfZ_SWAP(condition bool, expression string) *SVGSVGElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGSVGElement) Z_SWAPRemove() *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGSVGElement) Z_SWAP_PUSH(expression string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSVGElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGSVGElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGSVGElement) Z_SWAP_PUSHRemove() *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGSVGElement) Z_TRIGGER(expression string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSVGElement) IfZ_TRIGGER(condition bool, expression string) *SVGSVGElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGSVGElement) Z_TRIGGERRemove() *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGSVGElement) Z_REQ_METHOD(c SVGSvgZReqMethodChoice) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGSvgZReqMethodChoice string

const (
	// default GET
	SVGSvgZReqMethod_empty SVGSvgZReqMethodChoice = ""
	// GET
	SVGSvgZReqMethod_get SVGSvgZReqMethodChoice = "get"
	// POST
	SVGSvgZReqMethod_post SVGSvgZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGSVGElement) Z_REQ_METHODRemove(c SVGSvgZReqMethodChoice) *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGSVGElement) Z_REQ_STRATEGY(expression string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSVGElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGSVGElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGSVGElement) Z_REQ_STRATEGYRemove() *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGSVGElement) Z_REQ_HISTORY(expression string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSVGElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGSVGElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGSVGElement) Z_REQ_HISTORYRemove() *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGSVGElement) Z_DATA(expression string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSVGElement) IfZ_DATA(condition bool, expression string) *SVGSVGElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGSVGElement) Z_DATARemove() *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGSVGElement) Z_JSON(expression string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSVGElement) IfZ_JSON(condition bool, expression string) *SVGSVGElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGSVGElement) Z_JSONRemove() *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGSVGElement) Z_REQ_BATCH(expression string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSVGElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGSVGElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGSVGElement) Z_REQ_BATCHRemove() *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGSVGElement) Z_ACTION(expression string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSVGElement) IfZ_ACTION(condition bool, expression string) *SVGSVGElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGSVGElement) Z_ACTIONRemove() *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGSVGElement) Z_REQ_BEFORE(expression string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSVGElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGSVGElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGSVGElement) Z_REQ_BEFORERemove() *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGSVGElement) Z_REQ_AFTER(expression string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSVGElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGSVGElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGSVGElement) Z_REQ_AFTERRemove() *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
