package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <filter> SVG element defines a custom filter effect by grouping atomic
// filter primitives
// It is never rendered directly
// A filter is referenced by using the filter attribute on the target SVG element
// or via the filter CSS property.
type SVGFILTERElement struct {
	*Element
}

// Create a new SVGFILTERElement element.
// This will create a new element with the tag
// "filter" during rendering.
func SVG_FILTER(children ...ElementRenderer) *SVGFILTERElement {
	e := NewElement("filter", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFILTERElement{Element: e}
}

func (e *SVGFILTERElement) Children(children ...ElementRenderer) *SVGFILTERElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFILTERElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFILTERElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFILTERElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFILTERElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFILTERElement) Attr(name string, value ...string) *SVGFILTERElement {
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

func (e *SVGFILTERElement) Attrs(attrs ...string) *SVGFILTERElement {
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

func (e *SVGFILTERElement) AttrsMap(attrs map[string]string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFILTERElement) Text(text string) *SVGFILTERElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFILTERElement) TextF(format string, args ...any) *SVGFILTERElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFILTERElement) IfText(condition bool, text string) *SVGFILTERElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFILTERElement) IfTextF(condition bool, format string, args ...any) *SVGFILTERElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFILTERElement) Escaped(text string) *SVGFILTERElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFILTERElement) IfEscaped(condition bool, text string) *SVGFILTERElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFILTERElement) EscapedF(format string, args ...any) *SVGFILTERElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFILTERElement) IfEscapedF(condition bool, format string, args ...any) *SVGFILTERElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFILTERElement) CustomData(key, value string) *SVGFILTERElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFILTERElement) IfCustomData(condition bool, key, value string) *SVGFILTERElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFILTERElement) CustomDataF(key, format string, args ...any) *SVGFILTERElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFILTERElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFILTERElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFILTERElement) CustomDataRemove(key string) *SVGFILTERElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The coordinate system for attributes x, y, width and height.
func (e *SVGFILTERElement) FILTER_UNITS(c SVGFilterFilterUnitsChoice) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("filterUnits", string(c))
	return e
}

type SVGFilterFilterUnitsChoice string

const (
	// The coordinate system for attributes x, y, width and height.
	SVGFilterFilterUnits_userSpaceOnUse SVGFilterFilterUnitsChoice = "userSpaceOnUse"
	// The coordinate system for attributes x, y, width and height.
	SVGFilterFilterUnits_objectBoundingBox SVGFilterFilterUnitsChoice = "objectBoundingBox"
)

// Remove the attribute FILTER_UNITS from the element.
func (e *SVGFILTERElement) FILTER_UNITSRemove(c SVGFilterFilterUnitsChoice) *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("filterUnits")
	return e
}

// The coordinate system for the various length values within the filter.
func (e *SVGFILTERElement) PRIMITIVE_UNITS(c SVGFilterPrimitiveUnitsChoice) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("primitiveUnits", string(c))
	return e
}

type SVGFilterPrimitiveUnitsChoice string

const (
	// The coordinate system for the various length values within the filter.
	SVGFilterPrimitiveUnits_userSpaceOnUse SVGFilterPrimitiveUnitsChoice = "userSpaceOnUse"
	// The coordinate system for the various length values within the filter.
	SVGFilterPrimitiveUnits_objectBoundingBox SVGFilterPrimitiveUnitsChoice = "objectBoundingBox"
)

// Remove the attribute PRIMITIVE_UNITS from the element.
func (e *SVGFILTERElement) PRIMITIVE_UNITSRemove(c SVGFilterPrimitiveUnitsChoice) *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("primitiveUnits")
	return e
}

// The x attribute indicates where the left edge of the filter is placed.
func (e *SVGFILTERElement) X(s string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("x", s)
	return e
}

func (e *SVGFILTERElement) XF(format string, args ...any) *SVGFILTERElement {
	return e.X(fmt.Sprintf(format, args...))
}

func (e *SVGFILTERElement) IfX(condition bool, s string) *SVGFILTERElement {
	if condition {
		e.X(s)
	}
	return e
}

func (e *SVGFILTERElement) IfXF(condition bool, format string, args ...any) *SVGFILTERElement {
	if condition {
		e.X(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute X from the element.
func (e *SVGFILTERElement) XRemove(s string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("x")
	return e
}

func (e *SVGFILTERElement) XRemoveF(format string, args ...any) *SVGFILTERElement {
	return e.XRemove(fmt.Sprintf(format, args...))
}

// The y attribute indicates where the top edge of the filter is placed.
func (e *SVGFILTERElement) Y(s string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("y", s)
	return e
}

func (e *SVGFILTERElement) YF(format string, args ...any) *SVGFILTERElement {
	return e.Y(fmt.Sprintf(format, args...))
}

func (e *SVGFILTERElement) IfY(condition bool, s string) *SVGFILTERElement {
	if condition {
		e.Y(s)
	}
	return e
}

func (e *SVGFILTERElement) IfYF(condition bool, format string, args ...any) *SVGFILTERElement {
	if condition {
		e.Y(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute Y from the element.
func (e *SVGFILTERElement) YRemove(s string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("y")
	return e
}

func (e *SVGFILTERElement) YRemoveF(format string, args ...any) *SVGFILTERElement {
	return e.YRemove(fmt.Sprintf(format, args...))
}

// The width attribute indicates the width of the filter primitive box.
func (e *SVGFILTERElement) WIDTH(s string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("width", s)
	return e
}

func (e *SVGFILTERElement) WIDTHF(format string, args ...any) *SVGFILTERElement {
	return e.WIDTH(fmt.Sprintf(format, args...))
}

func (e *SVGFILTERElement) IfWIDTH(condition bool, s string) *SVGFILTERElement {
	if condition {
		e.WIDTH(s)
	}
	return e
}

func (e *SVGFILTERElement) IfWIDTHF(condition bool, format string, args ...any) *SVGFILTERElement {
	if condition {
		e.WIDTH(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute WIDTH from the element.
func (e *SVGFILTERElement) WIDTHRemove(s string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("width")
	return e
}

func (e *SVGFILTERElement) WIDTHRemoveF(format string, args ...any) *SVGFILTERElement {
	return e.WIDTHRemove(fmt.Sprintf(format, args...))
}

// The height attribute indicates the height of the filter primitive box.
func (e *SVGFILTERElement) HEIGHT(s string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("height", s)
	return e
}

func (e *SVGFILTERElement) HEIGHTF(format string, args ...any) *SVGFILTERElement {
	return e.HEIGHT(fmt.Sprintf(format, args...))
}

func (e *SVGFILTERElement) IfHEIGHT(condition bool, s string) *SVGFILTERElement {
	if condition {
		e.HEIGHT(s)
	}
	return e
}

func (e *SVGFILTERElement) IfHEIGHTF(condition bool, format string, args ...any) *SVGFILTERElement {
	if condition {
		e.HEIGHT(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HEIGHT from the element.
func (e *SVGFILTERElement) HEIGHTRemove(s string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("height")
	return e
}

func (e *SVGFILTERElement) HEIGHTRemoveF(format string, args ...any) *SVGFILTERElement {
	return e.HEIGHTRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGFILTERElement) ID(s string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFILTERElement) IDF(format string, args ...any) *SVGFILTERElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFILTERElement) IfID(condition bool, s string) *SVGFILTERElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFILTERElement) IfIDF(condition bool, format string, args ...any) *SVGFILTERElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFILTERElement) IDRemove(s string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFILTERElement) IDRemoveF(format string, args ...any) *SVGFILTERElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFILTERElement) CLASS(s ...string) *SVGFILTERElement {
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

func (e *SVGFILTERElement) IfCLASS(condition bool, s ...string) *SVGFILTERElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFILTERElement) CLASSRemove(s ...string) *SVGFILTERElement {
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
func (e *SVGFILTERElement) STYLEF(k string, format string, args ...any) *SVGFILTERElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFILTERElement) IfSTYLE(condition bool, k string, v string) *SVGFILTERElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFILTERElement) STYLE(k string, v string) *SVGFILTERElement {
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

func (e *SVGFILTERElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFILTERElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFILTERElement) STYLEMap(m map[string]string) *SVGFILTERElement {
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
func (e *SVGFILTERElement) STYLEPairs(pairs ...string) *SVGFILTERElement {
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

func (e *SVGFILTERElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFILTERElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFILTERElement) STYLERemove(keys ...string) *SVGFILTERElement {
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

func (e *SVGFILTERElement) Z_REQ(expression string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFILTERElement) IfZ_REQ(condition bool, expression string) *SVGFILTERElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFILTERElement) Z_REQRemove() *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFILTERElement) Z_TARGET(expression string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFILTERElement) IfZ_TARGET(condition bool, expression string) *SVGFILTERElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFILTERElement) Z_TARGETRemove() *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFILTERElement) Z_REQ_SELECTOR(expression string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFILTERElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFILTERElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFILTERElement) Z_REQ_SELECTORRemove() *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFILTERElement) Z_SWAP(expression string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFILTERElement) IfZ_SWAP(condition bool, expression string) *SVGFILTERElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFILTERElement) Z_SWAPRemove() *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFILTERElement) Z_SWAP_PUSH(expression string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFILTERElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFILTERElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFILTERElement) Z_SWAP_PUSHRemove() *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFILTERElement) Z_TRIGGER(expression string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFILTERElement) IfZ_TRIGGER(condition bool, expression string) *SVGFILTERElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFILTERElement) Z_TRIGGERRemove() *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFILTERElement) Z_REQ_METHOD(c SVGFilterZReqMethodChoice) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFilterZReqMethodChoice string

const (
	// default GET
	SVGFilterZReqMethod_empty SVGFilterZReqMethodChoice = ""
	// GET
	SVGFilterZReqMethod_get SVGFilterZReqMethodChoice = "get"
	// POST
	SVGFilterZReqMethod_post SVGFilterZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFILTERElement) Z_REQ_METHODRemove(c SVGFilterZReqMethodChoice) *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFILTERElement) Z_REQ_STRATEGY(expression string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFILTERElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFILTERElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFILTERElement) Z_REQ_STRATEGYRemove() *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFILTERElement) Z_REQ_HISTORY(expression string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFILTERElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFILTERElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFILTERElement) Z_REQ_HISTORYRemove() *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFILTERElement) Z_DATA(expression string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFILTERElement) IfZ_DATA(condition bool, expression string) *SVGFILTERElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFILTERElement) Z_DATARemove() *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFILTERElement) Z_JSON(expression string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFILTERElement) IfZ_JSON(condition bool, expression string) *SVGFILTERElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFILTERElement) Z_JSONRemove() *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFILTERElement) Z_REQ_BATCH(expression string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFILTERElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFILTERElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFILTERElement) Z_REQ_BATCHRemove() *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFILTERElement) Z_ACTION(expression string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFILTERElement) IfZ_ACTION(condition bool, expression string) *SVGFILTERElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFILTERElement) Z_ACTIONRemove() *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFILTERElement) Z_REQ_BEFORE(expression string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFILTERElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFILTERElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFILTERElement) Z_REQ_BEFORERemove() *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFILTERElement) Z_REQ_AFTER(expression string) *SVGFILTERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFILTERElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFILTERElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFILTERElement) Z_REQ_AFTERRemove() *SVGFILTERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
