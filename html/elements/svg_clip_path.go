package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <clipPath> SVG element defines a clipping path
// A clipping path is used/referenced using the clip-path property.
type SVGCLIPPATHElement struct {
	*Element
}

// Create a new SVGCLIPPATHElement element.
// This will create a new element with the tag
// "clipPath" during rendering.
func SVG_CLIPPATH(children ...ElementRenderer) *SVGCLIPPATHElement {
	e := NewElement("clipPath", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGCLIPPATHElement{Element: e}
}

func (e *SVGCLIPPATHElement) Children(children ...ElementRenderer) *SVGCLIPPATHElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGCLIPPATHElement) IfChildren(condition bool, children ...ElementRenderer) *SVGCLIPPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGCLIPPATHElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGCLIPPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGCLIPPATHElement) Attr(name string, value ...string) *SVGCLIPPATHElement {
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

func (e *SVGCLIPPATHElement) Attrs(attrs ...string) *SVGCLIPPATHElement {
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

func (e *SVGCLIPPATHElement) AttrsMap(attrs map[string]string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGCLIPPATHElement) Text(text string) *SVGCLIPPATHElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGCLIPPATHElement) TextF(format string, args ...any) *SVGCLIPPATHElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGCLIPPATHElement) IfText(condition bool, text string) *SVGCLIPPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGCLIPPATHElement) IfTextF(condition bool, format string, args ...any) *SVGCLIPPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGCLIPPATHElement) Escaped(text string) *SVGCLIPPATHElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGCLIPPATHElement) IfEscaped(condition bool, text string) *SVGCLIPPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGCLIPPATHElement) EscapedF(format string, args ...any) *SVGCLIPPATHElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGCLIPPATHElement) IfEscapedF(condition bool, format string, args ...any) *SVGCLIPPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGCLIPPATHElement) CustomData(key, value string) *SVGCLIPPATHElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGCLIPPATHElement) IfCustomData(condition bool, key, value string) *SVGCLIPPATHElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGCLIPPATHElement) CustomDataF(key, format string, args ...any) *SVGCLIPPATHElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGCLIPPATHElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGCLIPPATHElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGCLIPPATHElement) CustomDataRemove(key string) *SVGCLIPPATHElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The coordinate system for the contents of the <clipPath> element.
func (e *SVGCLIPPATHElement) CLIP_PATH_UNITS(c SVGClipPathClipPathUnitsChoice) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("clipPathUnits", string(c))
	return e
}

type SVGClipPathClipPathUnitsChoice string

const (
	// The contents of the <clipPath> element represent values in the current user
	// coordinate system.
	SVGClipPathClipPathUnits_userSpaceOnUse SVGClipPathClipPathUnitsChoice = "userSpaceOnUse"
	// The contents of the <clipPath> element represent values in the coordinate
	// system that results from taking the current user coordinate system in place at
	// the time when the <clipPath> element is referenced (i.e., the user coordinate
	// system for the element referencing the <clipPath> element via a clip-path
	// property).
	SVGClipPathClipPathUnits_objectBoundingBox SVGClipPathClipPathUnitsChoice = "objectBoundingBox"
)

// Remove the attribute CLIP_PATH_UNITS from the element.
func (e *SVGCLIPPATHElement) CLIP_PATH_UNITSRemove(c SVGClipPathClipPathUnitsChoice) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("clipPathUnits")
	return e
}

// Specifies a unique id for an element
func (e *SVGCLIPPATHElement) ID(s string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGCLIPPATHElement) IDF(format string, args ...any) *SVGCLIPPATHElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGCLIPPATHElement) IfID(condition bool, s string) *SVGCLIPPATHElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGCLIPPATHElement) IfIDF(condition bool, format string, args ...any) *SVGCLIPPATHElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGCLIPPATHElement) IDRemove(s string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGCLIPPATHElement) IDRemoveF(format string, args ...any) *SVGCLIPPATHElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGCLIPPATHElement) CLASS(s ...string) *SVGCLIPPATHElement {
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

func (e *SVGCLIPPATHElement) IfCLASS(condition bool, s ...string) *SVGCLIPPATHElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGCLIPPATHElement) CLASSRemove(s ...string) *SVGCLIPPATHElement {
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
func (e *SVGCLIPPATHElement) STYLEF(k string, format string, args ...any) *SVGCLIPPATHElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGCLIPPATHElement) IfSTYLE(condition bool, k string, v string) *SVGCLIPPATHElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGCLIPPATHElement) STYLE(k string, v string) *SVGCLIPPATHElement {
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

func (e *SVGCLIPPATHElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGCLIPPATHElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGCLIPPATHElement) STYLEMap(m map[string]string) *SVGCLIPPATHElement {
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
func (e *SVGCLIPPATHElement) STYLEPairs(pairs ...string) *SVGCLIPPATHElement {
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

func (e *SVGCLIPPATHElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGCLIPPATHElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGCLIPPATHElement) STYLERemove(keys ...string) *SVGCLIPPATHElement {
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

func (e *SVGCLIPPATHElement) Z_REQ(expression string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCLIPPATHElement) IfZ_REQ(condition bool, expression string) *SVGCLIPPATHElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGCLIPPATHElement) Z_REQRemove() *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGCLIPPATHElement) Z_TARGET(expression string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCLIPPATHElement) IfZ_TARGET(condition bool, expression string) *SVGCLIPPATHElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGCLIPPATHElement) Z_TARGETRemove() *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGCLIPPATHElement) Z_REQ_SELECTOR(expression string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCLIPPATHElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGCLIPPATHElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGCLIPPATHElement) Z_REQ_SELECTORRemove() *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGCLIPPATHElement) Z_SWAP(expression string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCLIPPATHElement) IfZ_SWAP(condition bool, expression string) *SVGCLIPPATHElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGCLIPPATHElement) Z_SWAPRemove() *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGCLIPPATHElement) Z_SWAP_PUSH(expression string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCLIPPATHElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGCLIPPATHElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGCLIPPATHElement) Z_SWAP_PUSHRemove() *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGCLIPPATHElement) Z_TRIGGER(expression string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCLIPPATHElement) IfZ_TRIGGER(condition bool, expression string) *SVGCLIPPATHElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGCLIPPATHElement) Z_TRIGGERRemove() *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGCLIPPATHElement) Z_REQ_METHOD(c SVGClipPathZReqMethodChoice) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGClipPathZReqMethodChoice string

const (
	// default GET
	SVGClipPathZReqMethod_empty SVGClipPathZReqMethodChoice = ""
	// GET
	SVGClipPathZReqMethod_get SVGClipPathZReqMethodChoice = "get"
	// POST
	SVGClipPathZReqMethod_post SVGClipPathZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGCLIPPATHElement) Z_REQ_METHODRemove(c SVGClipPathZReqMethodChoice) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGCLIPPATHElement) Z_REQ_STRATEGY(expression string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCLIPPATHElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGCLIPPATHElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGCLIPPATHElement) Z_REQ_STRATEGYRemove() *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGCLIPPATHElement) Z_REQ_HISTORY(expression string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCLIPPATHElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGCLIPPATHElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGCLIPPATHElement) Z_REQ_HISTORYRemove() *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGCLIPPATHElement) Z_DATA(expression string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCLIPPATHElement) IfZ_DATA(condition bool, expression string) *SVGCLIPPATHElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGCLIPPATHElement) Z_DATARemove() *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGCLIPPATHElement) Z_JSON(expression string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCLIPPATHElement) IfZ_JSON(condition bool, expression string) *SVGCLIPPATHElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGCLIPPATHElement) Z_JSONRemove() *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGCLIPPATHElement) Z_REQ_BATCH(expression string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCLIPPATHElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGCLIPPATHElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGCLIPPATHElement) Z_REQ_BATCHRemove() *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGCLIPPATHElement) Z_ACTION(expression string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCLIPPATHElement) IfZ_ACTION(condition bool, expression string) *SVGCLIPPATHElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGCLIPPATHElement) Z_ACTIONRemove() *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGCLIPPATHElement) Z_REQ_BEFORE(expression string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCLIPPATHElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGCLIPPATHElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGCLIPPATHElement) Z_REQ_BEFORERemove() *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGCLIPPATHElement) Z_REQ_AFTER(expression string) *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGCLIPPATHElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGCLIPPATHElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGCLIPPATHElement) Z_REQ_AFTERRemove() *SVGCLIPPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
