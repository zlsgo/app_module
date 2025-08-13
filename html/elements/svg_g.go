package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <g> SVG element is a container used to group other SVG elements.
type SVGGElement struct {
	*Element
}

// Create a new SVGGElement element.
// This will create a new element with the tag
// "g" during rendering.
func SVG_G(children ...ElementRenderer) *SVGGElement {
	e := NewElement("g", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGGElement{Element: e}
}

func (e *SVGGElement) Children(children ...ElementRenderer) *SVGGElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGGElement) IfChildren(condition bool, children ...ElementRenderer) *SVGGElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGGElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGGElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGGElement) Attr(name string, value ...string) *SVGGElement {
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

func (e *SVGGElement) Attrs(attrs ...string) *SVGGElement {
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

func (e *SVGGElement) AttrsMap(attrs map[string]string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGGElement) Text(text string) *SVGGElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGGElement) TextF(format string, args ...any) *SVGGElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGGElement) IfText(condition bool, text string) *SVGGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGGElement) IfTextF(condition bool, format string, args ...any) *SVGGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGGElement) Escaped(text string) *SVGGElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGGElement) IfEscaped(condition bool, text string) *SVGGElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGGElement) EscapedF(format string, args ...any) *SVGGElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGGElement) IfEscapedF(condition bool, format string, args ...any) *SVGGElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGGElement) CustomData(key, value string) *SVGGElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGGElement) IfCustomData(condition bool, key, value string) *SVGGElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGGElement) CustomDataF(key, format string, args ...any) *SVGGElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGGElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGGElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGGElement) CustomDataRemove(key string) *SVGGElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// A space-separated list of required extensions, indicating that the parent SVG
// document must include the specified extensions for this element to be valid.
func (e *SVGGElement) REQUIRED_EXTENSIONS(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("requiredExtensions", s)
	return e
}

func (e *SVGGElement) REQUIRED_EXTENSIONSF(format string, args ...any) *SVGGElement {
	return e.REQUIRED_EXTENSIONS(fmt.Sprintf(format, args...))
}

func (e *SVGGElement) IfREQUIRED_EXTENSIONS(condition bool, s string) *SVGGElement {
	if condition {
		e.REQUIRED_EXTENSIONS(s)
	}
	return e
}

func (e *SVGGElement) IfREQUIRED_EXTENSIONSF(condition bool, format string, args ...any) *SVGGElement {
	if condition {
		e.REQUIRED_EXTENSIONS(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REQUIRED_EXTENSIONS from the element.
func (e *SVGGElement) REQUIRED_EXTENSIONSRemove(s string) *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("requiredExtensions")
	return e
}

func (e *SVGGElement) REQUIRED_EXTENSIONSRemoveF(format string, args ...any) *SVGGElement {
	return e.REQUIRED_EXTENSIONSRemove(fmt.Sprintf(format, args...))
}

// A space-separated list of required features, indicating that the parent SVG
// document must include support for all of the specified features for this
// element to be valid.
func (e *SVGGElement) REQUIRED_FEATURES(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("requiredFeatures", s)
	return e
}

func (e *SVGGElement) REQUIRED_FEATURESF(format string, args ...any) *SVGGElement {
	return e.REQUIRED_FEATURES(fmt.Sprintf(format, args...))
}

func (e *SVGGElement) IfREQUIRED_FEATURES(condition bool, s string) *SVGGElement {
	if condition {
		e.REQUIRED_FEATURES(s)
	}
	return e
}

func (e *SVGGElement) IfREQUIRED_FEATURESF(condition bool, format string, args ...any) *SVGGElement {
	if condition {
		e.REQUIRED_FEATURES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REQUIRED_FEATURES from the element.
func (e *SVGGElement) REQUIRED_FEATURESRemove(s string) *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("requiredFeatures")
	return e
}

func (e *SVGGElement) REQUIRED_FEATURESRemoveF(format string, args ...any) *SVGGElement {
	return e.REQUIRED_FEATURESRemove(fmt.Sprintf(format, args...))
}

// A space-separated list of language codes, indicating that the parent SVG
// document must include support for all of the specified languages for this
// element to be valid.
func (e *SVGGElement) SYSTEM_LANGUAGE(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("systemLanguage", s)
	return e
}

func (e *SVGGElement) SYSTEM_LANGUAGEF(format string, args ...any) *SVGGElement {
	return e.SYSTEM_LANGUAGE(fmt.Sprintf(format, args...))
}

func (e *SVGGElement) IfSYSTEM_LANGUAGE(condition bool, s string) *SVGGElement {
	if condition {
		e.SYSTEM_LANGUAGE(s)
	}
	return e
}

func (e *SVGGElement) IfSYSTEM_LANGUAGEF(condition bool, format string, args ...any) *SVGGElement {
	if condition {
		e.SYSTEM_LANGUAGE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute SYSTEM_LANGUAGE from the element.
func (e *SVGGElement) SYSTEM_LANGUAGERemove(s string) *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("systemLanguage")
	return e
}

func (e *SVGGElement) SYSTEM_LANGUAGERemoveF(format string, args ...any) *SVGGElement {
	return e.SYSTEM_LANGUAGERemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGGElement) ID(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGGElement) IDF(format string, args ...any) *SVGGElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGGElement) IfID(condition bool, s string) *SVGGElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGGElement) IfIDF(condition bool, format string, args ...any) *SVGGElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGGElement) IDRemove(s string) *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGGElement) IDRemoveF(format string, args ...any) *SVGGElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGGElement) CLASS(s ...string) *SVGGElement {
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

func (e *SVGGElement) IfCLASS(condition bool, s ...string) *SVGGElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGGElement) CLASSRemove(s ...string) *SVGGElement {
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
func (e *SVGGElement) STYLEF(k string, format string, args ...any) *SVGGElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGGElement) IfSTYLE(condition bool, k string, v string) *SVGGElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGGElement) STYLE(k string, v string) *SVGGElement {
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

func (e *SVGGElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGGElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGGElement) STYLEMap(m map[string]string) *SVGGElement {
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
func (e *SVGGElement) STYLEPairs(pairs ...string) *SVGGElement {
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

func (e *SVGGElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGGElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGGElement) STYLERemove(keys ...string) *SVGGElement {
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

func (e *SVGGElement) Z_REQ(expression string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGGElement) IfZ_REQ(condition bool, expression string) *SVGGElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGGElement) Z_REQRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGGElement) Z_TARGET(expression string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGGElement) IfZ_TARGET(condition bool, expression string) *SVGGElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGGElement) Z_TARGETRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGGElement) Z_REQ_SELECTOR(expression string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGGElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGGElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGGElement) Z_REQ_SELECTORRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGGElement) Z_SWAP(expression string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGGElement) IfZ_SWAP(condition bool, expression string) *SVGGElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGGElement) Z_SWAPRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGGElement) Z_SWAP_PUSH(expression string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGGElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGGElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGGElement) Z_SWAP_PUSHRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGGElement) Z_TRIGGER(expression string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGGElement) IfZ_TRIGGER(condition bool, expression string) *SVGGElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGGElement) Z_TRIGGERRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGGElement) Z_REQ_METHOD(c SVGGZReqMethodChoice) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGGZReqMethodChoice string

const (
	// default GET
	SVGGZReqMethod_empty SVGGZReqMethodChoice = ""
	// GET
	SVGGZReqMethod_get SVGGZReqMethodChoice = "get"
	// POST
	SVGGZReqMethod_post SVGGZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGGElement) Z_REQ_METHODRemove(c SVGGZReqMethodChoice) *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGGElement) Z_REQ_STRATEGY(expression string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGGElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGGElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGGElement) Z_REQ_STRATEGYRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGGElement) Z_REQ_HISTORY(expression string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGGElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGGElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGGElement) Z_REQ_HISTORYRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGGElement) Z_DATA(expression string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGGElement) IfZ_DATA(condition bool, expression string) *SVGGElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGGElement) Z_DATARemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGGElement) Z_JSON(expression string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGGElement) IfZ_JSON(condition bool, expression string) *SVGGElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGGElement) Z_JSONRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGGElement) Z_REQ_BATCH(expression string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGGElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGGElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGGElement) Z_REQ_BATCHRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGGElement) Z_ACTION(expression string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGGElement) IfZ_ACTION(condition bool, expression string) *SVGGElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGGElement) Z_ACTIONRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGGElement) Z_REQ_BEFORE(expression string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGGElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGGElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGGElement) Z_REQ_BEFORERemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGGElement) Z_REQ_AFTER(expression string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGGElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGGElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGGElement) Z_REQ_AFTERRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
