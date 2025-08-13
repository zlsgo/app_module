package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feMergeNode> SVG element allows a series of filter primitives to be
// connected together graphically
// Incoming nodes are blended into the background via the defined compositing
// operator.
type SVGFEMERGENODEElement struct {
	*Element
}

// Create a new SVGFEMERGENODEElement element.
// This will create a new element with the tag
// "feMergeNode" during rendering.
func SVG_FEMERGENODE(children ...ElementRenderer) *SVGFEMERGENODEElement {
	e := NewElement("feMergeNode", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEMERGENODEElement{Element: e}
}

func (e *SVGFEMERGENODEElement) Children(children ...ElementRenderer) *SVGFEMERGENODEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEMERGENODEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEMERGENODEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEMERGENODEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEMERGENODEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEMERGENODEElement) Attr(name string, value ...string) *SVGFEMERGENODEElement {
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

func (e *SVGFEMERGENODEElement) Attrs(attrs ...string) *SVGFEMERGENODEElement {
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

func (e *SVGFEMERGENODEElement) AttrsMap(attrs map[string]string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEMERGENODEElement) Text(text string) *SVGFEMERGENODEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEMERGENODEElement) TextF(format string, args ...any) *SVGFEMERGENODEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGENODEElement) IfText(condition bool, text string) *SVGFEMERGENODEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEMERGENODEElement) IfTextF(condition bool, format string, args ...any) *SVGFEMERGENODEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEMERGENODEElement) Escaped(text string) *SVGFEMERGENODEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEMERGENODEElement) IfEscaped(condition bool, text string) *SVGFEMERGENODEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEMERGENODEElement) EscapedF(format string, args ...any) *SVGFEMERGENODEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGENODEElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEMERGENODEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEMERGENODEElement) CustomData(key, value string) *SVGFEMERGENODEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEMERGENODEElement) IfCustomData(condition bool, key, value string) *SVGFEMERGENODEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEMERGENODEElement) CustomDataF(key, format string, args ...any) *SVGFEMERGENODEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGENODEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEMERGENODEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEMERGENODEElement) CustomDataRemove(key string) *SVGFEMERGENODEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The identifier for the input SVGAnimatedString attribute on the given
// 'feMergeNode' element.
func (e *SVGFEMERGENODEElement) IN(s string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFEMERGENODEElement) INF(format string, args ...any) *SVGFEMERGENODEElement {
	return e.IN(fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGENODEElement) IfIN(condition bool, s string) *SVGFEMERGENODEElement {
	if condition {
		e.IN(s)
	}
	return e
}

func (e *SVGFEMERGENODEElement) IfINF(condition bool, format string, args ...any) *SVGFEMERGENODEElement {
	if condition {
		e.IN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFEMERGENODEElement) INRemove(s string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("in")
	return e
}

func (e *SVGFEMERGENODEElement) INRemoveF(format string, args ...any) *SVGFEMERGENODEElement {
	return e.INRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGFEMERGENODEElement) ID(s string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEMERGENODEElement) IDF(format string, args ...any) *SVGFEMERGENODEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGENODEElement) IfID(condition bool, s string) *SVGFEMERGENODEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEMERGENODEElement) IfIDF(condition bool, format string, args ...any) *SVGFEMERGENODEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEMERGENODEElement) IDRemove(s string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEMERGENODEElement) IDRemoveF(format string, args ...any) *SVGFEMERGENODEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEMERGENODEElement) CLASS(s ...string) *SVGFEMERGENODEElement {
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

func (e *SVGFEMERGENODEElement) IfCLASS(condition bool, s ...string) *SVGFEMERGENODEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEMERGENODEElement) CLASSRemove(s ...string) *SVGFEMERGENODEElement {
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
func (e *SVGFEMERGENODEElement) STYLEF(k string, format string, args ...any) *SVGFEMERGENODEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGENODEElement) IfSTYLE(condition bool, k string, v string) *SVGFEMERGENODEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEMERGENODEElement) STYLE(k string, v string) *SVGFEMERGENODEElement {
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

func (e *SVGFEMERGENODEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEMERGENODEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEMERGENODEElement) STYLEMap(m map[string]string) *SVGFEMERGENODEElement {
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
func (e *SVGFEMERGENODEElement) STYLEPairs(pairs ...string) *SVGFEMERGENODEElement {
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

func (e *SVGFEMERGENODEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEMERGENODEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEMERGENODEElement) STYLERemove(keys ...string) *SVGFEMERGENODEElement {
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

func (e *SVGFEMERGENODEElement) Z_REQ(expression string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGENODEElement) IfZ_REQ(condition bool, expression string) *SVGFEMERGENODEElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEMERGENODEElement) Z_REQRemove() *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEMERGENODEElement) Z_TARGET(expression string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGENODEElement) IfZ_TARGET(condition bool, expression string) *SVGFEMERGENODEElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEMERGENODEElement) Z_TARGETRemove() *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEMERGENODEElement) Z_REQ_SELECTOR(expression string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGENODEElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEMERGENODEElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEMERGENODEElement) Z_REQ_SELECTORRemove() *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEMERGENODEElement) Z_SWAP(expression string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGENODEElement) IfZ_SWAP(condition bool, expression string) *SVGFEMERGENODEElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEMERGENODEElement) Z_SWAPRemove() *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEMERGENODEElement) Z_SWAP_PUSH(expression string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGENODEElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEMERGENODEElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEMERGENODEElement) Z_SWAP_PUSHRemove() *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEMERGENODEElement) Z_TRIGGER(expression string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGENODEElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEMERGENODEElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEMERGENODEElement) Z_TRIGGERRemove() *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEMERGENODEElement) Z_REQ_METHOD(c SVGFeMergeNodeZReqMethodChoice) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeMergeNodeZReqMethodChoice string

const (
	// default GET
	SVGFeMergeNodeZReqMethod_empty SVGFeMergeNodeZReqMethodChoice = ""
	// GET
	SVGFeMergeNodeZReqMethod_get SVGFeMergeNodeZReqMethodChoice = "get"
	// POST
	SVGFeMergeNodeZReqMethod_post SVGFeMergeNodeZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEMERGENODEElement) Z_REQ_METHODRemove(c SVGFeMergeNodeZReqMethodChoice) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEMERGENODEElement) Z_REQ_STRATEGY(expression string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGENODEElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEMERGENODEElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEMERGENODEElement) Z_REQ_STRATEGYRemove() *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEMERGENODEElement) Z_REQ_HISTORY(expression string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGENODEElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEMERGENODEElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEMERGENODEElement) Z_REQ_HISTORYRemove() *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEMERGENODEElement) Z_DATA(expression string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGENODEElement) IfZ_DATA(condition bool, expression string) *SVGFEMERGENODEElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEMERGENODEElement) Z_DATARemove() *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEMERGENODEElement) Z_JSON(expression string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGENODEElement) IfZ_JSON(condition bool, expression string) *SVGFEMERGENODEElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEMERGENODEElement) Z_JSONRemove() *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEMERGENODEElement) Z_REQ_BATCH(expression string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGENODEElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEMERGENODEElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEMERGENODEElement) Z_REQ_BATCHRemove() *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEMERGENODEElement) Z_ACTION(expression string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGENODEElement) IfZ_ACTION(condition bool, expression string) *SVGFEMERGENODEElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEMERGENODEElement) Z_ACTIONRemove() *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEMERGENODEElement) Z_REQ_BEFORE(expression string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGENODEElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEMERGENODEElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEMERGENODEElement) Z_REQ_BEFORERemove() *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEMERGENODEElement) Z_REQ_AFTER(expression string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGENODEElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEMERGENODEElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEMERGENODEElement) Z_REQ_AFTERRemove() *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
