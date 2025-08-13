package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <switch> SVG element evaluates the requiredFeatures, requiredExtensions and
// systemLanguage attributes on its direct child elements in order, and then
// processes and renders the first child for which these attributes evaluate to
// true
// All others will be bypassed and therefore not rendered
// If the child element is a container element such as a <g>, then the entire
// subtree is either processed/rendered or bypassed/not rendered.
type SVGSWITCHElement struct {
	*Element
}

// Create a new SVGSWITCHElement element.
// This will create a new element with the tag
// "switch" during rendering.
func SVG_SWITCH(children ...ElementRenderer) *SVGSWITCHElement {
	e := NewElement("switch", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGSWITCHElement{Element: e}
}

func (e *SVGSWITCHElement) Children(children ...ElementRenderer) *SVGSWITCHElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGSWITCHElement) IfChildren(condition bool, children ...ElementRenderer) *SVGSWITCHElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGSWITCHElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGSWITCHElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGSWITCHElement) Attr(name string, value ...string) *SVGSWITCHElement {
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

func (e *SVGSWITCHElement) Attrs(attrs ...string) *SVGSWITCHElement {
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

func (e *SVGSWITCHElement) AttrsMap(attrs map[string]string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGSWITCHElement) Text(text string) *SVGSWITCHElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGSWITCHElement) TextF(format string, args ...any) *SVGSWITCHElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfText(condition bool, text string) *SVGSWITCHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGSWITCHElement) IfTextF(condition bool, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGSWITCHElement) Escaped(text string) *SVGSWITCHElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGSWITCHElement) IfEscaped(condition bool, text string) *SVGSWITCHElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGSWITCHElement) EscapedF(format string, args ...any) *SVGSWITCHElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfEscapedF(condition bool, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGSWITCHElement) CustomData(key, value string) *SVGSWITCHElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGSWITCHElement) IfCustomData(condition bool, key, value string) *SVGSWITCHElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGSWITCHElement) CustomDataF(key, format string, args ...any) *SVGSWITCHElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGSWITCHElement) CustomDataRemove(key string) *SVGSWITCHElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// A space-separated list of required features, indicating that the parent SVG
// document must include support for all of the specified features for this
// element to be valid.
func (e *SVGSWITCHElement) REQUIRED_FEATURES(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("requiredFeatures", s)
	return e
}

func (e *SVGSWITCHElement) REQUIRED_FEATURESF(format string, args ...any) *SVGSWITCHElement {
	return e.REQUIRED_FEATURES(fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfREQUIRED_FEATURES(condition bool, s string) *SVGSWITCHElement {
	if condition {
		e.REQUIRED_FEATURES(s)
	}
	return e
}

func (e *SVGSWITCHElement) IfREQUIRED_FEATURESF(condition bool, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.REQUIRED_FEATURES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REQUIRED_FEATURES from the element.
func (e *SVGSWITCHElement) REQUIRED_FEATURESRemove(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("requiredFeatures")
	return e
}

func (e *SVGSWITCHElement) REQUIRED_FEATURESRemoveF(format string, args ...any) *SVGSWITCHElement {
	return e.REQUIRED_FEATURESRemove(fmt.Sprintf(format, args...))
}

// A space-separated list of required extensions, indicating that the parent SVG
// document must include the specified extensions for this element to be valid.
func (e *SVGSWITCHElement) REQUIRED_EXTENSIONS(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("requiredExtensions", s)
	return e
}

func (e *SVGSWITCHElement) REQUIRED_EXTENSIONSF(format string, args ...any) *SVGSWITCHElement {
	return e.REQUIRED_EXTENSIONS(fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfREQUIRED_EXTENSIONS(condition bool, s string) *SVGSWITCHElement {
	if condition {
		e.REQUIRED_EXTENSIONS(s)
	}
	return e
}

func (e *SVGSWITCHElement) IfREQUIRED_EXTENSIONSF(condition bool, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.REQUIRED_EXTENSIONS(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REQUIRED_EXTENSIONS from the element.
func (e *SVGSWITCHElement) REQUIRED_EXTENSIONSRemove(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("requiredExtensions")
	return e
}

func (e *SVGSWITCHElement) REQUIRED_EXTENSIONSRemoveF(format string, args ...any) *SVGSWITCHElement {
	return e.REQUIRED_EXTENSIONSRemove(fmt.Sprintf(format, args...))
}

// A space-separated list of language codes, indicating that the parent SVG
// document must include support for all of the specified languages for this
// element to be valid.
func (e *SVGSWITCHElement) SYSTEM_LANGUAGE(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("systemLanguage", s)
	return e
}

func (e *SVGSWITCHElement) SYSTEM_LANGUAGEF(format string, args ...any) *SVGSWITCHElement {
	return e.SYSTEM_LANGUAGE(fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfSYSTEM_LANGUAGE(condition bool, s string) *SVGSWITCHElement {
	if condition {
		e.SYSTEM_LANGUAGE(s)
	}
	return e
}

func (e *SVGSWITCHElement) IfSYSTEM_LANGUAGEF(condition bool, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.SYSTEM_LANGUAGE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute SYSTEM_LANGUAGE from the element.
func (e *SVGSWITCHElement) SYSTEM_LANGUAGERemove(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("systemLanguage")
	return e
}

func (e *SVGSWITCHElement) SYSTEM_LANGUAGERemoveF(format string, args ...any) *SVGSWITCHElement {
	return e.SYSTEM_LANGUAGERemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGSWITCHElement) ID(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGSWITCHElement) IDF(format string, args ...any) *SVGSWITCHElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfID(condition bool, s string) *SVGSWITCHElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGSWITCHElement) IfIDF(condition bool, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGSWITCHElement) IDRemove(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGSWITCHElement) IDRemoveF(format string, args ...any) *SVGSWITCHElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGSWITCHElement) CLASS(s ...string) *SVGSWITCHElement {
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

func (e *SVGSWITCHElement) IfCLASS(condition bool, s ...string) *SVGSWITCHElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGSWITCHElement) CLASSRemove(s ...string) *SVGSWITCHElement {
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
func (e *SVGSWITCHElement) STYLEF(k string, format string, args ...any) *SVGSWITCHElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfSTYLE(condition bool, k string, v string) *SVGSWITCHElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGSWITCHElement) STYLE(k string, v string) *SVGSWITCHElement {
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

func (e *SVGSWITCHElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGSWITCHElement) STYLEMap(m map[string]string) *SVGSWITCHElement {
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
func (e *SVGSWITCHElement) STYLEPairs(pairs ...string) *SVGSWITCHElement {
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

func (e *SVGSWITCHElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGSWITCHElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGSWITCHElement) STYLERemove(keys ...string) *SVGSWITCHElement {
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

func (e *SVGSWITCHElement) Z_REQ(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfZ_REQ(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGSWITCHElement) Z_REQRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGSWITCHElement) Z_TARGET(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfZ_TARGET(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGSWITCHElement) Z_TARGETRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGSWITCHElement) Z_REQ_SELECTOR(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGSWITCHElement) Z_REQ_SELECTORRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGSWITCHElement) Z_SWAP(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfZ_SWAP(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGSWITCHElement) Z_SWAPRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGSWITCHElement) Z_SWAP_PUSH(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGSWITCHElement) Z_SWAP_PUSHRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGSWITCHElement) Z_TRIGGER(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfZ_TRIGGER(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGSWITCHElement) Z_TRIGGERRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGSWITCHElement) Z_REQ_METHOD(c SVGSwitchZReqMethodChoice) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGSwitchZReqMethodChoice string

const (
	// default GET
	SVGSwitchZReqMethod_empty SVGSwitchZReqMethodChoice = ""
	// GET
	SVGSwitchZReqMethod_get SVGSwitchZReqMethodChoice = "get"
	// POST
	SVGSwitchZReqMethod_post SVGSwitchZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGSWITCHElement) Z_REQ_METHODRemove(c SVGSwitchZReqMethodChoice) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGSWITCHElement) Z_REQ_STRATEGY(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGSWITCHElement) Z_REQ_STRATEGYRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGSWITCHElement) Z_REQ_HISTORY(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGSWITCHElement) Z_REQ_HISTORYRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGSWITCHElement) Z_DATA(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfZ_DATA(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGSWITCHElement) Z_DATARemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGSWITCHElement) Z_JSON(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfZ_JSON(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGSWITCHElement) Z_JSONRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGSWITCHElement) Z_REQ_BATCH(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGSWITCHElement) Z_REQ_BATCHRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGSWITCHElement) Z_ACTION(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfZ_ACTION(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGSWITCHElement) Z_ACTIONRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGSWITCHElement) Z_REQ_BEFORE(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGSWITCHElement) Z_REQ_BEFORERemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGSWITCHElement) Z_REQ_AFTER(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGSWITCHElement) Z_REQ_AFTERRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
