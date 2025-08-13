package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <pattern> SVG element fills a region with a pattern defined by an SVG
// image.
type SVGPATTERNElement struct {
	*Element
}

// Create a new SVGPATTERNElement element.
// This will create a new element with the tag
// "pattern" during rendering.
func SVG_PATTERN(children ...ElementRenderer) *SVGPATTERNElement {
	e := NewElement("pattern", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGPATTERNElement{Element: e}
}

func (e *SVGPATTERNElement) Children(children ...ElementRenderer) *SVGPATTERNElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGPATTERNElement) IfChildren(condition bool, children ...ElementRenderer) *SVGPATTERNElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGPATTERNElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGPATTERNElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGPATTERNElement) Attr(name string, value ...string) *SVGPATTERNElement {
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

func (e *SVGPATTERNElement) Attrs(attrs ...string) *SVGPATTERNElement {
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

func (e *SVGPATTERNElement) AttrsMap(attrs map[string]string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGPATTERNElement) Text(text string) *SVGPATTERNElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGPATTERNElement) TextF(format string, args ...any) *SVGPATTERNElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGPATTERNElement) IfText(condition bool, text string) *SVGPATTERNElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGPATTERNElement) IfTextF(condition bool, format string, args ...any) *SVGPATTERNElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGPATTERNElement) Escaped(text string) *SVGPATTERNElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGPATTERNElement) IfEscaped(condition bool, text string) *SVGPATTERNElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGPATTERNElement) EscapedF(format string, args ...any) *SVGPATTERNElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGPATTERNElement) IfEscapedF(condition bool, format string, args ...any) *SVGPATTERNElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGPATTERNElement) CustomData(key, value string) *SVGPATTERNElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGPATTERNElement) IfCustomData(condition bool, key, value string) *SVGPATTERNElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGPATTERNElement) CustomDataF(key, format string, args ...any) *SVGPATTERNElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGPATTERNElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGPATTERNElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGPATTERNElement) CustomDataRemove(key string) *SVGPATTERNElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The coordinate system for attributes x, y, width and height.
func (e *SVGPATTERNElement) PATTERN_UNITS(c SVGPatternPatternUnitsChoice) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("patternUnits", string(c))
	return e
}

type SVGPatternPatternUnitsChoice string

const (
	// The coordinate system for attributes x, y, width and height.
	SVGPatternPatternUnits_userSpaceOnUse SVGPatternPatternUnitsChoice = "userSpaceOnUse"
	// The coordinate system for attributes x, y, width and height.
	SVGPatternPatternUnits_objectBoundingBox SVGPatternPatternUnitsChoice = "objectBoundingBox"
)

// Remove the attribute PATTERN_UNITS from the element.
func (e *SVGPATTERNElement) PATTERN_UNITSRemove(c SVGPatternPatternUnitsChoice) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("patternUnits")
	return e
}

// The coordinate system for the various length values within the filter.
func (e *SVGPATTERNElement) PATTERN_CONTENT_UNITS(c SVGPatternPatternContentUnitsChoice) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("patternContentUnits", string(c))
	return e
}

type SVGPatternPatternContentUnitsChoice string

const (
	// The coordinate system for the various length values within the filter.
	SVGPatternPatternContentUnits_userSpaceOnUse SVGPatternPatternContentUnitsChoice = "userSpaceOnUse"
	// The coordinate system for the various length values within the filter.
	SVGPatternPatternContentUnits_objectBoundingBox SVGPatternPatternContentUnitsChoice = "objectBoundingBox"
)

// Remove the attribute PATTERN_CONTENT_UNITS from the element.
func (e *SVGPATTERNElement) PATTERN_CONTENT_UNITSRemove(c SVGPatternPatternContentUnitsChoice) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("patternContentUnits")
	return e
}

// The definition of how the pattern is tiled, read about <a
// href="https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/patternTransform">patternTransform</a>.
func (e *SVGPATTERNElement) PATTERN_TRANSFORM(s string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("patternTransform", s)
	return e
}

func (e *SVGPATTERNElement) PATTERN_TRANSFORMF(format string, args ...any) *SVGPATTERNElement {
	return e.PATTERN_TRANSFORM(fmt.Sprintf(format, args...))
}

func (e *SVGPATTERNElement) IfPATTERN_TRANSFORM(condition bool, s string) *SVGPATTERNElement {
	if condition {
		e.PATTERN_TRANSFORM(s)
	}
	return e
}

func (e *SVGPATTERNElement) IfPATTERN_TRANSFORMF(condition bool, format string, args ...any) *SVGPATTERNElement {
	if condition {
		e.PATTERN_TRANSFORM(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute PATTERN_TRANSFORM from the element.
func (e *SVGPATTERNElement) PATTERN_TRANSFORMRemove(s string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("patternTransform")
	return e
}

func (e *SVGPATTERNElement) PATTERN_TRANSFORMRemoveF(format string, args ...any) *SVGPATTERNElement {
	return e.PATTERN_TRANSFORMRemove(fmt.Sprintf(format, args...))
}

// The x-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGPATTERNElement) X(f float64) *SVGPATTERNElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("x", f)
	return e
}

func (e *SVGPATTERNElement) IfX(condition bool, f float64) *SVGPATTERNElement {
	if condition {
		e.X(f)
	}
	return e
}

// The y-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGPATTERNElement) Y(f float64) *SVGPATTERNElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("y", f)
	return e
}

func (e *SVGPATTERNElement) IfY(condition bool, f float64) *SVGPATTERNElement {
	if condition {
		e.Y(f)
	}
	return e
}

// The width of the rectangular region.
func (e *SVGPATTERNElement) WIDTH(f float64) *SVGPATTERNElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("width", f)
	return e
}

func (e *SVGPATTERNElement) IfWIDTH(condition bool, f float64) *SVGPATTERNElement {
	if condition {
		e.WIDTH(f)
	}
	return e
}

// The height of the rectangular region.
func (e *SVGPATTERNElement) HEIGHT(f float64) *SVGPATTERNElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("height", f)
	return e
}

func (e *SVGPATTERNElement) IfHEIGHT(condition bool, f float64) *SVGPATTERNElement {
	if condition {
		e.HEIGHT(f)
	}
	return e
}

// A URI reference to the image to paint.
func (e *SVGPATTERNElement) HREF(s string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("href", s)
	return e
}

func (e *SVGPATTERNElement) HREFF(format string, args ...any) *SVGPATTERNElement {
	return e.HREF(fmt.Sprintf(format, args...))
}

func (e *SVGPATTERNElement) IfHREF(condition bool, s string) *SVGPATTERNElement {
	if condition {
		e.HREF(s)
	}
	return e
}

func (e *SVGPATTERNElement) IfHREFF(condition bool, format string, args ...any) *SVGPATTERNElement {
	if condition {
		e.HREF(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HREF from the element.
func (e *SVGPATTERNElement) HREFRemove(s string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("href")
	return e
}

func (e *SVGPATTERNElement) HREFRemoveF(format string, args ...any) *SVGPATTERNElement {
	return e.HREFRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGPATTERNElement) ID(s string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGPATTERNElement) IDF(format string, args ...any) *SVGPATTERNElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGPATTERNElement) IfID(condition bool, s string) *SVGPATTERNElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGPATTERNElement) IfIDF(condition bool, format string, args ...any) *SVGPATTERNElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGPATTERNElement) IDRemove(s string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGPATTERNElement) IDRemoveF(format string, args ...any) *SVGPATTERNElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGPATTERNElement) CLASS(s ...string) *SVGPATTERNElement {
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

func (e *SVGPATTERNElement) IfCLASS(condition bool, s ...string) *SVGPATTERNElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGPATTERNElement) CLASSRemove(s ...string) *SVGPATTERNElement {
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
func (e *SVGPATTERNElement) STYLEF(k string, format string, args ...any) *SVGPATTERNElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGPATTERNElement) IfSTYLE(condition bool, k string, v string) *SVGPATTERNElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGPATTERNElement) STYLE(k string, v string) *SVGPATTERNElement {
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

func (e *SVGPATTERNElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGPATTERNElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGPATTERNElement) STYLEMap(m map[string]string) *SVGPATTERNElement {
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
func (e *SVGPATTERNElement) STYLEPairs(pairs ...string) *SVGPATTERNElement {
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

func (e *SVGPATTERNElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGPATTERNElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGPATTERNElement) STYLERemove(keys ...string) *SVGPATTERNElement {
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

func (e *SVGPATTERNElement) Z_REQ(expression string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATTERNElement) IfZ_REQ(condition bool, expression string) *SVGPATTERNElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGPATTERNElement) Z_REQRemove() *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGPATTERNElement) Z_TARGET(expression string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATTERNElement) IfZ_TARGET(condition bool, expression string) *SVGPATTERNElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGPATTERNElement) Z_TARGETRemove() *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGPATTERNElement) Z_REQ_SELECTOR(expression string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATTERNElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGPATTERNElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGPATTERNElement) Z_REQ_SELECTORRemove() *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGPATTERNElement) Z_SWAP(expression string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATTERNElement) IfZ_SWAP(condition bool, expression string) *SVGPATTERNElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGPATTERNElement) Z_SWAPRemove() *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGPATTERNElement) Z_SWAP_PUSH(expression string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATTERNElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGPATTERNElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGPATTERNElement) Z_SWAP_PUSHRemove() *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGPATTERNElement) Z_TRIGGER(expression string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATTERNElement) IfZ_TRIGGER(condition bool, expression string) *SVGPATTERNElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGPATTERNElement) Z_TRIGGERRemove() *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGPATTERNElement) Z_REQ_METHOD(c SVGPatternZReqMethodChoice) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGPatternZReqMethodChoice string

const (
	// default GET
	SVGPatternZReqMethod_empty SVGPatternZReqMethodChoice = ""
	// GET
	SVGPatternZReqMethod_get SVGPatternZReqMethodChoice = "get"
	// POST
	SVGPatternZReqMethod_post SVGPatternZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGPATTERNElement) Z_REQ_METHODRemove(c SVGPatternZReqMethodChoice) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGPATTERNElement) Z_REQ_STRATEGY(expression string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATTERNElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGPATTERNElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGPATTERNElement) Z_REQ_STRATEGYRemove() *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGPATTERNElement) Z_REQ_HISTORY(expression string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATTERNElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGPATTERNElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGPATTERNElement) Z_REQ_HISTORYRemove() *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGPATTERNElement) Z_DATA(expression string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATTERNElement) IfZ_DATA(condition bool, expression string) *SVGPATTERNElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGPATTERNElement) Z_DATARemove() *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGPATTERNElement) Z_JSON(expression string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATTERNElement) IfZ_JSON(condition bool, expression string) *SVGPATTERNElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGPATTERNElement) Z_JSONRemove() *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGPATTERNElement) Z_REQ_BATCH(expression string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATTERNElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGPATTERNElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGPATTERNElement) Z_REQ_BATCHRemove() *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGPATTERNElement) Z_ACTION(expression string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATTERNElement) IfZ_ACTION(condition bool, expression string) *SVGPATTERNElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGPATTERNElement) Z_ACTIONRemove() *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGPATTERNElement) Z_REQ_BEFORE(expression string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATTERNElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGPATTERNElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGPATTERNElement) Z_REQ_BEFORERemove() *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGPATTERNElement) Z_REQ_AFTER(expression string) *SVGPATTERNElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPATTERNElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGPATTERNElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGPATTERNElement) Z_REQ_AFTERRemove() *SVGPATTERNElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
