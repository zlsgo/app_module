package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <marker> SVG element defines the graphics that is to be used for drawing
// arrowheads or polymarkers on a given <path>, <line>, <polyline> or <polygon>
// element.
type SVGMARKERElement struct {
	*Element
}

// Create a new SVGMARKERElement element.
// This will create a new element with the tag
// "marker" during rendering.
func SVG_MARKER(children ...ElementRenderer) *SVGMARKERElement {
	e := NewElement("marker", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGMARKERElement{Element: e}
}

func (e *SVGMARKERElement) Children(children ...ElementRenderer) *SVGMARKERElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGMARKERElement) IfChildren(condition bool, children ...ElementRenderer) *SVGMARKERElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGMARKERElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGMARKERElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGMARKERElement) Attr(name string, value ...string) *SVGMARKERElement {
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

func (e *SVGMARKERElement) Attrs(attrs ...string) *SVGMARKERElement {
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

func (e *SVGMARKERElement) AttrsMap(attrs map[string]string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGMARKERElement) Text(text string) *SVGMARKERElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGMARKERElement) TextF(format string, args ...any) *SVGMARKERElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGMARKERElement) IfText(condition bool, text string) *SVGMARKERElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGMARKERElement) IfTextF(condition bool, format string, args ...any) *SVGMARKERElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGMARKERElement) Escaped(text string) *SVGMARKERElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGMARKERElement) IfEscaped(condition bool, text string) *SVGMARKERElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGMARKERElement) EscapedF(format string, args ...any) *SVGMARKERElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGMARKERElement) IfEscapedF(condition bool, format string, args ...any) *SVGMARKERElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGMARKERElement) CustomData(key, value string) *SVGMARKERElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGMARKERElement) IfCustomData(condition bool, key, value string) *SVGMARKERElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGMARKERElement) CustomDataF(key, format string, args ...any) *SVGMARKERElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGMARKERElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGMARKERElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGMARKERElement) CustomDataRemove(key string) *SVGMARKERElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The x-axis coordinate of the reference point which is to be aligned exactly at
// the marker position.
func (e *SVGMARKERElement) REF_X(f float64) *SVGMARKERElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("refX", f)
	return e
}

func (e *SVGMARKERElement) IfREF_X(condition bool, f float64) *SVGMARKERElement {
	if condition {
		e.REF_X(f)
	}
	return e
}

// The y-axis coordinate of the reference point which is to be aligned exactly at
// the marker position.
func (e *SVGMARKERElement) REF_Y(f float64) *SVGMARKERElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("refY", f)
	return e
}

func (e *SVGMARKERElement) IfREF_Y(condition bool, f float64) *SVGMARKERElement {
	if condition {
		e.REF_Y(f)
	}
	return e
}

// The width of the marker viewport.
func (e *SVGMARKERElement) MARKER_WIDTH(f float64) *SVGMARKERElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("markerWidth", f)
	return e
}

func (e *SVGMARKERElement) IfMARKER_WIDTH(condition bool, f float64) *SVGMARKERElement {
	if condition {
		e.MARKER_WIDTH(f)
	}
	return e
}

// The height of the marker viewport.
func (e *SVGMARKERElement) MARKER_HEIGHT(f float64) *SVGMARKERElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("markerHeight", f)
	return e
}

func (e *SVGMARKERElement) IfMARKER_HEIGHT(condition bool, f float64) *SVGMARKERElement {
	if condition {
		e.MARKER_HEIGHT(f)
	}
	return e
}

// The orientation of the marker relative to the shape it is attached to.
func (e *SVGMARKERElement) ORIENT(c SVGMarkerOrientChoice) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("orient", string(c))
	return e
}

type SVGMarkerOrientChoice string

const (
	// The orientation of the marker relative to the shape it is attached to.
	SVGMarkerOrient_auto SVGMarkerOrientChoice = "auto"
	// The orientation of the marker relative to the shape it is attached to.
	SVGMarkerOrient_auto_start_reverse SVGMarkerOrientChoice = "auto-start-reverse"
	// The orientation of the marker relative to the shape it is attached to.
	SVGMarkerOrient_angle SVGMarkerOrientChoice = "angle"
)

// Remove the attribute ORIENT from the element.
func (e *SVGMARKERElement) ORIENTRemove(c SVGMarkerOrientChoice) *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("orient")
	return e
}

// The coordinate system for the various length values within the marker.
func (e *SVGMARKERElement) MARKER_UNITS(c SVGMarkerMarkerUnitsChoice) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("markerUnits", string(c))
	return e
}

type SVGMarkerMarkerUnitsChoice string

const (
	// The coordinate system for the various length values within the marker.
	SVGMarkerMarkerUnits_userSpaceOnUse SVGMarkerMarkerUnitsChoice = "userSpaceOnUse"
	// The coordinate system for the various length values within the marker.
	SVGMarkerMarkerUnits_strokeWidth SVGMarkerMarkerUnitsChoice = "strokeWidth"
)

// Remove the attribute MARKER_UNITS from the element.
func (e *SVGMARKERElement) MARKER_UNITSRemove(c SVGMarkerMarkerUnitsChoice) *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("markerUnits")
	return e
}

// The position and size of the marker viewport (the bounds of the marker).
func (e *SVGMARKERElement) VIEW_BOX(s string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("viewBox", s)
	return e
}

func (e *SVGMARKERElement) VIEW_BOXF(format string, args ...any) *SVGMARKERElement {
	return e.VIEW_BOX(fmt.Sprintf(format, args...))
}

func (e *SVGMARKERElement) IfVIEW_BOX(condition bool, s string) *SVGMARKERElement {
	if condition {
		e.VIEW_BOX(s)
	}
	return e
}

func (e *SVGMARKERElement) IfVIEW_BOXF(condition bool, format string, args ...any) *SVGMARKERElement {
	if condition {
		e.VIEW_BOX(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute VIEW_BOX from the element.
func (e *SVGMARKERElement) VIEW_BOXRemove(s string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("viewBox")
	return e
}

func (e *SVGMARKERElement) VIEW_BOXRemoveF(format string, args ...any) *SVGMARKERElement {
	return e.VIEW_BOXRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGMARKERElement) ID(s string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGMARKERElement) IDF(format string, args ...any) *SVGMARKERElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGMARKERElement) IfID(condition bool, s string) *SVGMARKERElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGMARKERElement) IfIDF(condition bool, format string, args ...any) *SVGMARKERElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGMARKERElement) IDRemove(s string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGMARKERElement) IDRemoveF(format string, args ...any) *SVGMARKERElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGMARKERElement) CLASS(s ...string) *SVGMARKERElement {
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

func (e *SVGMARKERElement) IfCLASS(condition bool, s ...string) *SVGMARKERElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGMARKERElement) CLASSRemove(s ...string) *SVGMARKERElement {
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
func (e *SVGMARKERElement) STYLEF(k string, format string, args ...any) *SVGMARKERElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGMARKERElement) IfSTYLE(condition bool, k string, v string) *SVGMARKERElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGMARKERElement) STYLE(k string, v string) *SVGMARKERElement {
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

func (e *SVGMARKERElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGMARKERElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGMARKERElement) STYLEMap(m map[string]string) *SVGMARKERElement {
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
func (e *SVGMARKERElement) STYLEPairs(pairs ...string) *SVGMARKERElement {
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

func (e *SVGMARKERElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGMARKERElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGMARKERElement) STYLERemove(keys ...string) *SVGMARKERElement {
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

func (e *SVGMARKERElement) Z_REQ(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfZ_REQ(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGMARKERElement) Z_REQRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGMARKERElement) Z_TARGET(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfZ_TARGET(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGMARKERElement) Z_TARGETRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGMARKERElement) Z_REQ_SELECTOR(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGMARKERElement) Z_REQ_SELECTORRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGMARKERElement) Z_SWAP(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfZ_SWAP(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGMARKERElement) Z_SWAPRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGMARKERElement) Z_SWAP_PUSH(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGMARKERElement) Z_SWAP_PUSHRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGMARKERElement) Z_TRIGGER(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfZ_TRIGGER(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGMARKERElement) Z_TRIGGERRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGMARKERElement) Z_REQ_METHOD(c SVGMarkerZReqMethodChoice) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGMarkerZReqMethodChoice string

const (
	// default GET
	SVGMarkerZReqMethod_empty SVGMarkerZReqMethodChoice = ""
	// GET
	SVGMarkerZReqMethod_get SVGMarkerZReqMethodChoice = "get"
	// POST
	SVGMarkerZReqMethod_post SVGMarkerZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGMARKERElement) Z_REQ_METHODRemove(c SVGMarkerZReqMethodChoice) *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGMARKERElement) Z_REQ_STRATEGY(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGMARKERElement) Z_REQ_STRATEGYRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGMARKERElement) Z_REQ_HISTORY(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGMARKERElement) Z_REQ_HISTORYRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGMARKERElement) Z_DATA(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfZ_DATA(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGMARKERElement) Z_DATARemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGMARKERElement) Z_JSON(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfZ_JSON(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGMARKERElement) Z_JSONRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGMARKERElement) Z_REQ_BATCH(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGMARKERElement) Z_REQ_BATCHRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGMARKERElement) Z_ACTION(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfZ_ACTION(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGMARKERElement) Z_ACTIONRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGMARKERElement) Z_REQ_BEFORE(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGMARKERElement) Z_REQ_BEFORERemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGMARKERElement) Z_REQ_AFTER(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGMARKERElement) Z_REQ_AFTERRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
