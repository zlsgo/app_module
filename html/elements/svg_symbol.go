package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <symbol> SVG element is used to define graphical template objects which can
// be instantiated by a <use> element
// The use of symbol elements for graphics that are used multiple times in the
// same document adds structure and semantics
// Documents that are rich in structure may be rendered graphically, as speech, or
// as Braille, and thus promote accessibility
// note that a symbol element itself is not rendered
// Only instances of a symbol element (i.e., a reference to a symbol by a <use>
// element) are rendered
// To render a 'stand-alone' graphic that has been defined using a symbol, a
// reference to the symbol is referenced using a <use> element.
type SVGSYMBOLElement struct {
	*Element
}

// Create a new SVGSYMBOLElement element.
// This will create a new element with the tag
// "symbol" during rendering.
func SVG_SYMBOL(children ...ElementRenderer) *SVGSYMBOLElement {
	e := NewElement("symbol", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGSYMBOLElement{Element: e}
}

func (e *SVGSYMBOLElement) Children(children ...ElementRenderer) *SVGSYMBOLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGSYMBOLElement) IfChildren(condition bool, children ...ElementRenderer) *SVGSYMBOLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGSYMBOLElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGSYMBOLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGSYMBOLElement) Attr(name string, value ...string) *SVGSYMBOLElement {
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

func (e *SVGSYMBOLElement) Attrs(attrs ...string) *SVGSYMBOLElement {
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

func (e *SVGSYMBOLElement) AttrsMap(attrs map[string]string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGSYMBOLElement) Text(text string) *SVGSYMBOLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGSYMBOLElement) TextF(format string, args ...any) *SVGSYMBOLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGSYMBOLElement) IfText(condition bool, text string) *SVGSYMBOLElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGSYMBOLElement) IfTextF(condition bool, format string, args ...any) *SVGSYMBOLElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGSYMBOLElement) Escaped(text string) *SVGSYMBOLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGSYMBOLElement) IfEscaped(condition bool, text string) *SVGSYMBOLElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGSYMBOLElement) EscapedF(format string, args ...any) *SVGSYMBOLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGSYMBOLElement) IfEscapedF(condition bool, format string, args ...any) *SVGSYMBOLElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGSYMBOLElement) CustomData(key, value string) *SVGSYMBOLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGSYMBOLElement) IfCustomData(condition bool, key, value string) *SVGSYMBOLElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGSYMBOLElement) CustomDataF(key, format string, args ...any) *SVGSYMBOLElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGSYMBOLElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGSYMBOLElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGSYMBOLElement) CustomDataRemove(key string) *SVGSYMBOLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// Indicates how the fetched image is fitted into the destination rectangle.
func (e *SVGSYMBOLElement) PRESERVE_ASPECT_RATIO(c SVGSymbolPreserveAspectRatioChoice) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("preserveAspectRatio", string(c))
	return e
}

type SVGSymbolPreserveAspectRatioChoice string

const (
	// Do not force uniform scaling.
	SVGSymbolPreserveAspectRatio_none SVGSymbolPreserveAspectRatioChoice = "none"
	// Scale the image to the smallest size such that both its width and its height
	// can completely fit inside the corresponding dimension of the viewPort.
	SVGSymbolPreserveAspectRatio_xMinYMin SVGSymbolPreserveAspectRatioChoice = "xMinYMin"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGSymbolPreserveAspectRatio_xMidYMin SVGSymbolPreserveAspectRatioChoice = "xMidYMin"
	// Align the image with the corresponding side of the viewPort.
	SVGSymbolPreserveAspectRatio_xMaxYMin SVGSymbolPreserveAspectRatioChoice = "xMaxYMin"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGSymbolPreserveAspectRatio_xMinYMid SVGSymbolPreserveAspectRatioChoice = "xMinYMid"
	// Scale the image to the smallest size such that it can completely fit inside the
	// corresponding dimension of the viewPort.
	SVGSymbolPreserveAspectRatio_xMidYMid SVGSymbolPreserveAspectRatioChoice = "xMidYMid"
	// Align the image with the corresponding side of the viewPort.
	SVGSymbolPreserveAspectRatio_xMaxYMid SVGSymbolPreserveAspectRatioChoice = "xMaxYMid"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGSymbolPreserveAspectRatio_xMinYMax SVGSymbolPreserveAspectRatioChoice = "xMinYMax"
	// Align the image with the corresponding side of the viewPort.
	SVGSymbolPreserveAspectRatio_xMidYMax SVGSymbolPreserveAspectRatioChoice = "xMidYMax"
	// Scale the image to the smallest size such that both its width and its height
	// can completely fit inside the corresponding dimension of the viewPort.
	SVGSymbolPreserveAspectRatio_xMaxYMax SVGSymbolPreserveAspectRatioChoice = "xMaxYMax"
)

// Remove the attribute PRESERVE_ASPECT_RATIO from the element.
func (e *SVGSYMBOLElement) PRESERVE_ASPECT_RATIORemove(c SVGSymbolPreserveAspectRatioChoice) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("preserveAspectRatio")
	return e
}

// Specifies a unique id for an element
func (e *SVGSYMBOLElement) ID(s string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGSYMBOLElement) IDF(format string, args ...any) *SVGSYMBOLElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGSYMBOLElement) IfID(condition bool, s string) *SVGSYMBOLElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGSYMBOLElement) IfIDF(condition bool, format string, args ...any) *SVGSYMBOLElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGSYMBOLElement) IDRemove(s string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGSYMBOLElement) IDRemoveF(format string, args ...any) *SVGSYMBOLElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGSYMBOLElement) CLASS(s ...string) *SVGSYMBOLElement {
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

func (e *SVGSYMBOLElement) IfCLASS(condition bool, s ...string) *SVGSYMBOLElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGSYMBOLElement) CLASSRemove(s ...string) *SVGSYMBOLElement {
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
func (e *SVGSYMBOLElement) STYLEF(k string, format string, args ...any) *SVGSYMBOLElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGSYMBOLElement) IfSTYLE(condition bool, k string, v string) *SVGSYMBOLElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGSYMBOLElement) STYLE(k string, v string) *SVGSYMBOLElement {
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

func (e *SVGSYMBOLElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGSYMBOLElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGSYMBOLElement) STYLEMap(m map[string]string) *SVGSYMBOLElement {
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
func (e *SVGSYMBOLElement) STYLEPairs(pairs ...string) *SVGSYMBOLElement {
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

func (e *SVGSYMBOLElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGSYMBOLElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGSYMBOLElement) STYLERemove(keys ...string) *SVGSYMBOLElement {
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

func (e *SVGSYMBOLElement) Z_REQ(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfZ_REQ(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGSYMBOLElement) Z_REQRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGSYMBOLElement) Z_TARGET(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfZ_TARGET(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGSYMBOLElement) Z_TARGETRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGSYMBOLElement) Z_REQ_SELECTOR(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGSYMBOLElement) Z_REQ_SELECTORRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGSYMBOLElement) Z_SWAP(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfZ_SWAP(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGSYMBOLElement) Z_SWAPRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGSYMBOLElement) Z_SWAP_PUSH(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGSYMBOLElement) Z_SWAP_PUSHRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGSYMBOLElement) Z_TRIGGER(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfZ_TRIGGER(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGSYMBOLElement) Z_TRIGGERRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGSYMBOLElement) Z_REQ_METHOD(c SVGSymbolZReqMethodChoice) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGSymbolZReqMethodChoice string

const (
	// default GET
	SVGSymbolZReqMethod_empty SVGSymbolZReqMethodChoice = ""
	// GET
	SVGSymbolZReqMethod_get SVGSymbolZReqMethodChoice = "get"
	// POST
	SVGSymbolZReqMethod_post SVGSymbolZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGSYMBOLElement) Z_REQ_METHODRemove(c SVGSymbolZReqMethodChoice) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGSYMBOLElement) Z_REQ_STRATEGY(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGSYMBOLElement) Z_REQ_STRATEGYRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGSYMBOLElement) Z_REQ_HISTORY(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGSYMBOLElement) Z_REQ_HISTORYRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGSYMBOLElement) Z_DATA(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfZ_DATA(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGSYMBOLElement) Z_DATARemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGSYMBOLElement) Z_JSON(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfZ_JSON(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGSYMBOLElement) Z_JSONRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGSYMBOLElement) Z_REQ_BATCH(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGSYMBOLElement) Z_REQ_BATCHRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGSYMBOLElement) Z_ACTION(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfZ_ACTION(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGSYMBOLElement) Z_ACTIONRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGSYMBOLElement) Z_REQ_BEFORE(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGSYMBOLElement) Z_REQ_BEFORERemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGSYMBOLElement) Z_REQ_AFTER(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGSYMBOLElement) Z_REQ_AFTERRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
