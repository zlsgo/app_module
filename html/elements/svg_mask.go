package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <mask> SVG element hides portions of SVG elements for user display.
type SVGMASKElement struct {
	*Element
}

// Create a new SVGMASKElement element.
// This will create a new element with the tag
// "mask" during rendering.
func SVG_MASK(children ...ElementRenderer) *SVGMASKElement {
	e := NewElement("mask", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGMASKElement{Element: e}
}

func (e *SVGMASKElement) Children(children ...ElementRenderer) *SVGMASKElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGMASKElement) IfChildren(condition bool, children ...ElementRenderer) *SVGMASKElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGMASKElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGMASKElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGMASKElement) Attr(name string, value ...string) *SVGMASKElement {
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

func (e *SVGMASKElement) Attrs(attrs ...string) *SVGMASKElement {
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

func (e *SVGMASKElement) AttrsMap(attrs map[string]string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGMASKElement) Text(text string) *SVGMASKElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGMASKElement) TextF(format string, args ...any) *SVGMASKElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGMASKElement) IfText(condition bool, text string) *SVGMASKElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGMASKElement) IfTextF(condition bool, format string, args ...any) *SVGMASKElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGMASKElement) Escaped(text string) *SVGMASKElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGMASKElement) IfEscaped(condition bool, text string) *SVGMASKElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGMASKElement) EscapedF(format string, args ...any) *SVGMASKElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGMASKElement) IfEscapedF(condition bool, format string, args ...any) *SVGMASKElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGMASKElement) CustomData(key, value string) *SVGMASKElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGMASKElement) IfCustomData(condition bool, key, value string) *SVGMASKElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGMASKElement) CustomDataF(key, format string, args ...any) *SVGMASKElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGMASKElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGMASKElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGMASKElement) CustomDataRemove(key string) *SVGMASKElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The coordinate system for attributes x, y, width and height.
func (e *SVGMASKElement) MASK_CONTENT_UNITS(c SVGMaskMaskContentUnitsChoice) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("maskContentUnits", string(c))
	return e
}

type SVGMaskMaskContentUnitsChoice string

const (
	// The coordinate system for attributes x, y, width and height.
	SVGMaskMaskContentUnits_userSpaceOnUse SVGMaskMaskContentUnitsChoice = "userSpaceOnUse"
	// The coordinate system for attributes x, y, width and height.
	SVGMaskMaskContentUnits_objectBoundingBox SVGMaskMaskContentUnitsChoice = "objectBoundingBox"
)

// Remove the attribute MASK_CONTENT_UNITS from the element.
func (e *SVGMASKElement) MASK_CONTENT_UNITSRemove(c SVGMaskMaskContentUnitsChoice) *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("maskContentUnits")
	return e
}

// The coordinate system for the various length values within the filter.
func (e *SVGMASKElement) MASK_UNITS(c SVGMaskMaskUnitsChoice) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("maskUnits", string(c))
	return e
}

type SVGMaskMaskUnitsChoice string

const (
	// The coordinate system for the various length values within the filter.
	SVGMaskMaskUnits_userSpaceOnUse SVGMaskMaskUnitsChoice = "userSpaceOnUse"
	// The coordinate system for the various length values within the filter.
	SVGMaskMaskUnits_objectBoundingBox SVGMaskMaskUnitsChoice = "objectBoundingBox"
)

// Remove the attribute MASK_UNITS from the element.
func (e *SVGMASKElement) MASK_UNITSRemove(c SVGMaskMaskUnitsChoice) *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("maskUnits")
	return e
}

// The x-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGMASKElement) X(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("x", s)
	return e
}

func (e *SVGMASKElement) XF(format string, args ...any) *SVGMASKElement {
	return e.X(fmt.Sprintf(format, args...))
}

func (e *SVGMASKElement) IfX(condition bool, s string) *SVGMASKElement {
	if condition {
		e.X(s)
	}
	return e
}

func (e *SVGMASKElement) IfXF(condition bool, format string, args ...any) *SVGMASKElement {
	if condition {
		e.X(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute X from the element.
func (e *SVGMASKElement) XRemove(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("x")
	return e
}

func (e *SVGMASKElement) XRemoveF(format string, args ...any) *SVGMASKElement {
	return e.XRemove(fmt.Sprintf(format, args...))
}

// The y-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGMASKElement) Y(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("y", s)
	return e
}

func (e *SVGMASKElement) YF(format string, args ...any) *SVGMASKElement {
	return e.Y(fmt.Sprintf(format, args...))
}

func (e *SVGMASKElement) IfY(condition bool, s string) *SVGMASKElement {
	if condition {
		e.Y(s)
	}
	return e
}

func (e *SVGMASKElement) IfYF(condition bool, format string, args ...any) *SVGMASKElement {
	if condition {
		e.Y(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute Y from the element.
func (e *SVGMASKElement) YRemove(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("y")
	return e
}

func (e *SVGMASKElement) YRemoveF(format string, args ...any) *SVGMASKElement {
	return e.YRemove(fmt.Sprintf(format, args...))
}

// The width of the rectangular region.
func (e *SVGMASKElement) WIDTH(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("width", s)
	return e
}

func (e *SVGMASKElement) WIDTHF(format string, args ...any) *SVGMASKElement {
	return e.WIDTH(fmt.Sprintf(format, args...))
}

func (e *SVGMASKElement) IfWIDTH(condition bool, s string) *SVGMASKElement {
	if condition {
		e.WIDTH(s)
	}
	return e
}

func (e *SVGMASKElement) IfWIDTHF(condition bool, format string, args ...any) *SVGMASKElement {
	if condition {
		e.WIDTH(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute WIDTH from the element.
func (e *SVGMASKElement) WIDTHRemove(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("width")
	return e
}

func (e *SVGMASKElement) WIDTHRemoveF(format string, args ...any) *SVGMASKElement {
	return e.WIDTHRemove(fmt.Sprintf(format, args...))
}

// The height of the rectangular region.
func (e *SVGMASKElement) HEIGHT(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("height", s)
	return e
}

func (e *SVGMASKElement) HEIGHTF(format string, args ...any) *SVGMASKElement {
	return e.HEIGHT(fmt.Sprintf(format, args...))
}

func (e *SVGMASKElement) IfHEIGHT(condition bool, s string) *SVGMASKElement {
	if condition {
		e.HEIGHT(s)
	}
	return e
}

func (e *SVGMASKElement) IfHEIGHTF(condition bool, format string, args ...any) *SVGMASKElement {
	if condition {
		e.HEIGHT(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HEIGHT from the element.
func (e *SVGMASKElement) HEIGHTRemove(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("height")
	return e
}

func (e *SVGMASKElement) HEIGHTRemoveF(format string, args ...any) *SVGMASKElement {
	return e.HEIGHTRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGMASKElement) ID(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGMASKElement) IDF(format string, args ...any) *SVGMASKElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGMASKElement) IfID(condition bool, s string) *SVGMASKElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGMASKElement) IfIDF(condition bool, format string, args ...any) *SVGMASKElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGMASKElement) IDRemove(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGMASKElement) IDRemoveF(format string, args ...any) *SVGMASKElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGMASKElement) CLASS(s ...string) *SVGMASKElement {
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

func (e *SVGMASKElement) IfCLASS(condition bool, s ...string) *SVGMASKElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGMASKElement) CLASSRemove(s ...string) *SVGMASKElement {
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
func (e *SVGMASKElement) STYLEF(k string, format string, args ...any) *SVGMASKElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGMASKElement) IfSTYLE(condition bool, k string, v string) *SVGMASKElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGMASKElement) STYLE(k string, v string) *SVGMASKElement {
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

func (e *SVGMASKElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGMASKElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGMASKElement) STYLEMap(m map[string]string) *SVGMASKElement {
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
func (e *SVGMASKElement) STYLEPairs(pairs ...string) *SVGMASKElement {
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

func (e *SVGMASKElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGMASKElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGMASKElement) STYLERemove(keys ...string) *SVGMASKElement {
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

func (e *SVGMASKElement) Z_REQ(expression string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMASKElement) IfZ_REQ(condition bool, expression string) *SVGMASKElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGMASKElement) Z_REQRemove() *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGMASKElement) Z_TARGET(expression string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMASKElement) IfZ_TARGET(condition bool, expression string) *SVGMASKElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGMASKElement) Z_TARGETRemove() *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGMASKElement) Z_REQ_SELECTOR(expression string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMASKElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGMASKElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGMASKElement) Z_REQ_SELECTORRemove() *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGMASKElement) Z_SWAP(expression string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMASKElement) IfZ_SWAP(condition bool, expression string) *SVGMASKElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGMASKElement) Z_SWAPRemove() *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGMASKElement) Z_SWAP_PUSH(expression string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMASKElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGMASKElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGMASKElement) Z_SWAP_PUSHRemove() *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGMASKElement) Z_TRIGGER(expression string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMASKElement) IfZ_TRIGGER(condition bool, expression string) *SVGMASKElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGMASKElement) Z_TRIGGERRemove() *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGMASKElement) Z_REQ_METHOD(c SVGMaskZReqMethodChoice) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGMaskZReqMethodChoice string

const (
	// default GET
	SVGMaskZReqMethod_empty SVGMaskZReqMethodChoice = ""
	// GET
	SVGMaskZReqMethod_get SVGMaskZReqMethodChoice = "get"
	// POST
	SVGMaskZReqMethod_post SVGMaskZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGMASKElement) Z_REQ_METHODRemove(c SVGMaskZReqMethodChoice) *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGMASKElement) Z_REQ_STRATEGY(expression string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMASKElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGMASKElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGMASKElement) Z_REQ_STRATEGYRemove() *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGMASKElement) Z_REQ_HISTORY(expression string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMASKElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGMASKElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGMASKElement) Z_REQ_HISTORYRemove() *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGMASKElement) Z_DATA(expression string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMASKElement) IfZ_DATA(condition bool, expression string) *SVGMASKElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGMASKElement) Z_DATARemove() *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGMASKElement) Z_JSON(expression string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMASKElement) IfZ_JSON(condition bool, expression string) *SVGMASKElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGMASKElement) Z_JSONRemove() *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGMASKElement) Z_REQ_BATCH(expression string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMASKElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGMASKElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGMASKElement) Z_REQ_BATCHRemove() *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGMASKElement) Z_ACTION(expression string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMASKElement) IfZ_ACTION(condition bool, expression string) *SVGMASKElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGMASKElement) Z_ACTIONRemove() *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGMASKElement) Z_REQ_BEFORE(expression string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMASKElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGMASKElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGMASKElement) Z_REQ_BEFORERemove() *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGMASKElement) Z_REQ_AFTER(expression string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMASKElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGMASKElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGMASKElement) Z_REQ_AFTERRemove() *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
