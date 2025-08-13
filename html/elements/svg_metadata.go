package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <metadata> SVG element allows to add metadata to SVG content
// Metadata is structured information about data
// In XML, metadata can be added to an element using for example attributes.
type SVGMETADATAElement struct {
	*Element
}

// Create a new SVGMETADATAElement element.
// This will create a new element with the tag
// "metadata" during rendering.
func SVG_METADATA(children ...ElementRenderer) *SVGMETADATAElement {
	e := NewElement("metadata", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGMETADATAElement{Element: e}
}

func (e *SVGMETADATAElement) Children(children ...ElementRenderer) *SVGMETADATAElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGMETADATAElement) IfChildren(condition bool, children ...ElementRenderer) *SVGMETADATAElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGMETADATAElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGMETADATAElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGMETADATAElement) Attr(name string, value ...string) *SVGMETADATAElement {
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

func (e *SVGMETADATAElement) Attrs(attrs ...string) *SVGMETADATAElement {
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

func (e *SVGMETADATAElement) AttrsMap(attrs map[string]string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGMETADATAElement) Text(text string) *SVGMETADATAElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGMETADATAElement) TextF(format string, args ...any) *SVGMETADATAElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGMETADATAElement) IfText(condition bool, text string) *SVGMETADATAElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGMETADATAElement) IfTextF(condition bool, format string, args ...any) *SVGMETADATAElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGMETADATAElement) Escaped(text string) *SVGMETADATAElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGMETADATAElement) IfEscaped(condition bool, text string) *SVGMETADATAElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGMETADATAElement) EscapedF(format string, args ...any) *SVGMETADATAElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGMETADATAElement) IfEscapedF(condition bool, format string, args ...any) *SVGMETADATAElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGMETADATAElement) CustomData(key, value string) *SVGMETADATAElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGMETADATAElement) IfCustomData(condition bool, key, value string) *SVGMETADATAElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGMETADATAElement) CustomDataF(key, format string, args ...any) *SVGMETADATAElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGMETADATAElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGMETADATAElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGMETADATAElement) CustomDataRemove(key string) *SVGMETADATAElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// A space-separated list of required extensions, indicating that the parent SVG
// document must include the specified extensions for this element to be valid.
func (e *SVGMETADATAElement) REQUIRED_EXTENSIONS(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("requiredExtensions", s)
	return e
}

func (e *SVGMETADATAElement) REQUIRED_EXTENSIONSF(format string, args ...any) *SVGMETADATAElement {
	return e.REQUIRED_EXTENSIONS(fmt.Sprintf(format, args...))
}

func (e *SVGMETADATAElement) IfREQUIRED_EXTENSIONS(condition bool, s string) *SVGMETADATAElement {
	if condition {
		e.REQUIRED_EXTENSIONS(s)
	}
	return e
}

func (e *SVGMETADATAElement) IfREQUIRED_EXTENSIONSF(condition bool, format string, args ...any) *SVGMETADATAElement {
	if condition {
		e.REQUIRED_EXTENSIONS(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REQUIRED_EXTENSIONS from the element.
func (e *SVGMETADATAElement) REQUIRED_EXTENSIONSRemove(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("requiredExtensions")
	return e
}

func (e *SVGMETADATAElement) REQUIRED_EXTENSIONSRemoveF(format string, args ...any) *SVGMETADATAElement {
	return e.REQUIRED_EXTENSIONSRemove(fmt.Sprintf(format, args...))
}

// A space-separated list of required features, indicating that the parent SVG
// document must include support for all of the specified features for this
// element to be valid.
func (e *SVGMETADATAElement) REQUIRED_FEATURES(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("requiredFeatures", s)
	return e
}

func (e *SVGMETADATAElement) REQUIRED_FEATURESF(format string, args ...any) *SVGMETADATAElement {
	return e.REQUIRED_FEATURES(fmt.Sprintf(format, args...))
}

func (e *SVGMETADATAElement) IfREQUIRED_FEATURES(condition bool, s string) *SVGMETADATAElement {
	if condition {
		e.REQUIRED_FEATURES(s)
	}
	return e
}

func (e *SVGMETADATAElement) IfREQUIRED_FEATURESF(condition bool, format string, args ...any) *SVGMETADATAElement {
	if condition {
		e.REQUIRED_FEATURES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REQUIRED_FEATURES from the element.
func (e *SVGMETADATAElement) REQUIRED_FEATURESRemove(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("requiredFeatures")
	return e
}

func (e *SVGMETADATAElement) REQUIRED_FEATURESRemoveF(format string, args ...any) *SVGMETADATAElement {
	return e.REQUIRED_FEATURESRemove(fmt.Sprintf(format, args...))
}

// A space-separated list of language codes, indicating that the parent SVG
// document must include support for all of the specified languages for this
// element to be valid.
func (e *SVGMETADATAElement) SYSTEM_LANGUAGE(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("systemLanguage", s)
	return e
}

func (e *SVGMETADATAElement) SYSTEM_LANGUAGEF(format string, args ...any) *SVGMETADATAElement {
	return e.SYSTEM_LANGUAGE(fmt.Sprintf(format, args...))
}

func (e *SVGMETADATAElement) IfSYSTEM_LANGUAGE(condition bool, s string) *SVGMETADATAElement {
	if condition {
		e.SYSTEM_LANGUAGE(s)
	}
	return e
}

func (e *SVGMETADATAElement) IfSYSTEM_LANGUAGEF(condition bool, format string, args ...any) *SVGMETADATAElement {
	if condition {
		e.SYSTEM_LANGUAGE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute SYSTEM_LANGUAGE from the element.
func (e *SVGMETADATAElement) SYSTEM_LANGUAGERemove(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("systemLanguage")
	return e
}

func (e *SVGMETADATAElement) SYSTEM_LANGUAGERemoveF(format string, args ...any) *SVGMETADATAElement {
	return e.SYSTEM_LANGUAGERemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGMETADATAElement) ID(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGMETADATAElement) IDF(format string, args ...any) *SVGMETADATAElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGMETADATAElement) IfID(condition bool, s string) *SVGMETADATAElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGMETADATAElement) IfIDF(condition bool, format string, args ...any) *SVGMETADATAElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGMETADATAElement) IDRemove(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGMETADATAElement) IDRemoveF(format string, args ...any) *SVGMETADATAElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGMETADATAElement) CLASS(s ...string) *SVGMETADATAElement {
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

func (e *SVGMETADATAElement) IfCLASS(condition bool, s ...string) *SVGMETADATAElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGMETADATAElement) CLASSRemove(s ...string) *SVGMETADATAElement {
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
func (e *SVGMETADATAElement) STYLEF(k string, format string, args ...any) *SVGMETADATAElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGMETADATAElement) IfSTYLE(condition bool, k string, v string) *SVGMETADATAElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGMETADATAElement) STYLE(k string, v string) *SVGMETADATAElement {
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

func (e *SVGMETADATAElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGMETADATAElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGMETADATAElement) STYLEMap(m map[string]string) *SVGMETADATAElement {
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
func (e *SVGMETADATAElement) STYLEPairs(pairs ...string) *SVGMETADATAElement {
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

func (e *SVGMETADATAElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGMETADATAElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGMETADATAElement) STYLERemove(keys ...string) *SVGMETADATAElement {
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

func (e *SVGMETADATAElement) Z_REQ(expression string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMETADATAElement) IfZ_REQ(condition bool, expression string) *SVGMETADATAElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGMETADATAElement) Z_REQRemove() *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGMETADATAElement) Z_TARGET(expression string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMETADATAElement) IfZ_TARGET(condition bool, expression string) *SVGMETADATAElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGMETADATAElement) Z_TARGETRemove() *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGMETADATAElement) Z_REQ_SELECTOR(expression string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMETADATAElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGMETADATAElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGMETADATAElement) Z_REQ_SELECTORRemove() *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGMETADATAElement) Z_SWAP(expression string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMETADATAElement) IfZ_SWAP(condition bool, expression string) *SVGMETADATAElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGMETADATAElement) Z_SWAPRemove() *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGMETADATAElement) Z_SWAP_PUSH(expression string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMETADATAElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGMETADATAElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGMETADATAElement) Z_SWAP_PUSHRemove() *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGMETADATAElement) Z_TRIGGER(expression string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMETADATAElement) IfZ_TRIGGER(condition bool, expression string) *SVGMETADATAElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGMETADATAElement) Z_TRIGGERRemove() *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGMETADATAElement) Z_REQ_METHOD(c SVGMetadataZReqMethodChoice) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGMetadataZReqMethodChoice string

const (
	// default GET
	SVGMetadataZReqMethod_empty SVGMetadataZReqMethodChoice = ""
	// GET
	SVGMetadataZReqMethod_get SVGMetadataZReqMethodChoice = "get"
	// POST
	SVGMetadataZReqMethod_post SVGMetadataZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGMETADATAElement) Z_REQ_METHODRemove(c SVGMetadataZReqMethodChoice) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGMETADATAElement) Z_REQ_STRATEGY(expression string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMETADATAElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGMETADATAElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGMETADATAElement) Z_REQ_STRATEGYRemove() *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGMETADATAElement) Z_REQ_HISTORY(expression string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMETADATAElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGMETADATAElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGMETADATAElement) Z_REQ_HISTORYRemove() *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGMETADATAElement) Z_DATA(expression string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMETADATAElement) IfZ_DATA(condition bool, expression string) *SVGMETADATAElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGMETADATAElement) Z_DATARemove() *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGMETADATAElement) Z_JSON(expression string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMETADATAElement) IfZ_JSON(condition bool, expression string) *SVGMETADATAElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGMETADATAElement) Z_JSONRemove() *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGMETADATAElement) Z_REQ_BATCH(expression string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMETADATAElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGMETADATAElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGMETADATAElement) Z_REQ_BATCHRemove() *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGMETADATAElement) Z_ACTION(expression string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMETADATAElement) IfZ_ACTION(condition bool, expression string) *SVGMETADATAElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGMETADATAElement) Z_ACTIONRemove() *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGMETADATAElement) Z_REQ_BEFORE(expression string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMETADATAElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGMETADATAElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGMETADATAElement) Z_REQ_BEFORERemove() *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGMETADATAElement) Z_REQ_AFTER(expression string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMETADATAElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGMETADATAElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGMETADATAElement) Z_REQ_AFTERRemove() *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
