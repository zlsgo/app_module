package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <style> SVG element allows style sheets to be embedded directly within SVG
// content
// SVG's style element has the same attributes as the corresponding element in
// HTML (see HTML's <style> element).
type SVGSTYLEElement struct {
	*Element
}

// Create a new SVGSTYLEElement element.
// This will create a new element with the tag
// "style" during rendering.
func SVG_STYLE(children ...ElementRenderer) *SVGSTYLEElement {
	e := NewElement("style", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGSTYLEElement{Element: e}
}

func (e *SVGSTYLEElement) Children(children ...ElementRenderer) *SVGSTYLEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGSTYLEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGSTYLEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGSTYLEElement) Attr(name string, value ...string) *SVGSTYLEElement {
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

func (e *SVGSTYLEElement) Attrs(attrs ...string) *SVGSTYLEElement {
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

func (e *SVGSTYLEElement) AttrsMap(attrs map[string]string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGSTYLEElement) Text(text string) *SVGSTYLEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGSTYLEElement) TextF(format string, args ...any) *SVGSTYLEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfText(condition bool, text string) *SVGSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGSTYLEElement) IfTextF(condition bool, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGSTYLEElement) Escaped(text string) *SVGSTYLEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGSTYLEElement) IfEscaped(condition bool, text string) *SVGSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGSTYLEElement) EscapedF(format string, args ...any) *SVGSTYLEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfEscapedF(condition bool, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGSTYLEElement) CustomData(key, value string) *SVGSTYLEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGSTYLEElement) IfCustomData(condition bool, key, value string) *SVGSTYLEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGSTYLEElement) CustomDataF(key, format string, args ...any) *SVGSTYLEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGSTYLEElement) CustomDataRemove(key string) *SVGSTYLEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The style sheet language.
func (e *SVGSTYLEElement) TYPE(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("type", s)
	return e
}

func (e *SVGSTYLEElement) TYPEF(format string, args ...any) *SVGSTYLEElement {
	return e.TYPE(fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfTYPE(condition bool, s string) *SVGSTYLEElement {
	if condition {
		e.TYPE(s)
	}
	return e
}

func (e *SVGSTYLEElement) IfTYPEF(condition bool, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.TYPE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TYPE from the element.
func (e *SVGSTYLEElement) TYPERemove(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("type")
	return e
}

func (e *SVGSTYLEElement) TYPERemoveF(format string, args ...any) *SVGSTYLEElement {
	return e.TYPERemove(fmt.Sprintf(format, args...))
}

// The intended destination medium for style information.
func (e *SVGSTYLEElement) MEDIA(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("media", s)
	return e
}

func (e *SVGSTYLEElement) MEDIAF(format string, args ...any) *SVGSTYLEElement {
	return e.MEDIA(fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfMEDIA(condition bool, s string) *SVGSTYLEElement {
	if condition {
		e.MEDIA(s)
	}
	return e
}

func (e *SVGSTYLEElement) IfMEDIAF(condition bool, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.MEDIA(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MEDIA from the element.
func (e *SVGSTYLEElement) MEDIARemove(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("media")
	return e
}

func (e *SVGSTYLEElement) MEDIARemoveF(format string, args ...any) *SVGSTYLEElement {
	return e.MEDIARemove(fmt.Sprintf(format, args...))
}

// The advisory title.
func (e *SVGSTYLEElement) TITLE(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("title", s)
	return e
}

func (e *SVGSTYLEElement) TITLEF(format string, args ...any) *SVGSTYLEElement {
	return e.TITLE(fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfTITLE(condition bool, s string) *SVGSTYLEElement {
	if condition {
		e.TITLE(s)
	}
	return e
}

func (e *SVGSTYLEElement) IfTITLEF(condition bool, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.TITLE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TITLE from the element.
func (e *SVGSTYLEElement) TITLERemove(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("title")
	return e
}

func (e *SVGSTYLEElement) TITLERemoveF(format string, args ...any) *SVGSTYLEElement {
	return e.TITLERemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGSTYLEElement) ID(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGSTYLEElement) IDF(format string, args ...any) *SVGSTYLEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfID(condition bool, s string) *SVGSTYLEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGSTYLEElement) IfIDF(condition bool, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGSTYLEElement) IDRemove(s string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGSTYLEElement) IDRemoveF(format string, args ...any) *SVGSTYLEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGSTYLEElement) CLASS(s ...string) *SVGSTYLEElement {
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

func (e *SVGSTYLEElement) IfCLASS(condition bool, s ...string) *SVGSTYLEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGSTYLEElement) CLASSRemove(s ...string) *SVGSTYLEElement {
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
func (e *SVGSTYLEElement) STYLEF(k string, format string, args ...any) *SVGSTYLEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGSTYLEElement) IfSTYLE(condition bool, k string, v string) *SVGSTYLEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGSTYLEElement) STYLE(k string, v string) *SVGSTYLEElement {
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

func (e *SVGSTYLEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGSTYLEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGSTYLEElement) STYLEMap(m map[string]string) *SVGSTYLEElement {
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
func (e *SVGSTYLEElement) STYLEPairs(pairs ...string) *SVGSTYLEElement {
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

func (e *SVGSTYLEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGSTYLEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGSTYLEElement) STYLERemove(keys ...string) *SVGSTYLEElement {
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

func (e *SVGSTYLEElement) Z_REQ(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfZ_REQ(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGSTYLEElement) Z_REQRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGSTYLEElement) Z_TARGET(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfZ_TARGET(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGSTYLEElement) Z_TARGETRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGSTYLEElement) Z_REQ_SELECTOR(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGSTYLEElement) Z_REQ_SELECTORRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGSTYLEElement) Z_SWAP(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfZ_SWAP(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGSTYLEElement) Z_SWAPRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGSTYLEElement) Z_SWAP_PUSH(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGSTYLEElement) Z_SWAP_PUSHRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGSTYLEElement) Z_TRIGGER(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfZ_TRIGGER(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGSTYLEElement) Z_TRIGGERRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGSTYLEElement) Z_REQ_METHOD(c SVGStyleZReqMethodChoice) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGStyleZReqMethodChoice string

const (
	// default GET
	SVGStyleZReqMethod_empty SVGStyleZReqMethodChoice = ""
	// GET
	SVGStyleZReqMethod_get SVGStyleZReqMethodChoice = "get"
	// POST
	SVGStyleZReqMethod_post SVGStyleZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGSTYLEElement) Z_REQ_METHODRemove(c SVGStyleZReqMethodChoice) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGSTYLEElement) Z_REQ_STRATEGY(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGSTYLEElement) Z_REQ_STRATEGYRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGSTYLEElement) Z_REQ_HISTORY(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGSTYLEElement) Z_REQ_HISTORYRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGSTYLEElement) Z_DATA(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfZ_DATA(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGSTYLEElement) Z_DATARemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGSTYLEElement) Z_JSON(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfZ_JSON(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGSTYLEElement) Z_JSONRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGSTYLEElement) Z_REQ_BATCH(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGSTYLEElement) Z_REQ_BATCHRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGSTYLEElement) Z_ACTION(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfZ_ACTION(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGSTYLEElement) Z_ACTIONRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGSTYLEElement) Z_REQ_BEFORE(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGSTYLEElement) Z_REQ_BEFORERemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGSTYLEElement) Z_REQ_AFTER(expression string) *SVGSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSTYLEElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGSTYLEElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGSTYLEElement) Z_REQ_AFTERRemove() *SVGSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
