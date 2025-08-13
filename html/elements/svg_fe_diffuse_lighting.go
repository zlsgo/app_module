package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feDiffuseLighting> SVG filter primitive lights an image using the alpha
// channel as a bump map
// The resulting image, which is an RGBA opaque image, depends on the light color,
// light position and surface geometry of the input bump map.
type SVGFEDIFFUSELIGHTINGElement struct {
	*Element
}

// Create a new SVGFEDIFFUSELIGHTINGElement element.
// This will create a new element with the tag
// "feDiffuseLighting" during rendering.
func SVG_FEDIFFUSELIGHTING(children ...ElementRenderer) *SVGFEDIFFUSELIGHTINGElement {
	e := NewElement("feDiffuseLighting", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEDIFFUSELIGHTINGElement{Element: e}
}

func (e *SVGFEDIFFUSELIGHTINGElement) Children(children ...ElementRenderer) *SVGFEDIFFUSELIGHTINGElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) Attr(name string, value ...string) *SVGFEDIFFUSELIGHTINGElement {
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

func (e *SVGFEDIFFUSELIGHTINGElement) Attrs(attrs ...string) *SVGFEDIFFUSELIGHTINGElement {
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

func (e *SVGFEDIFFUSELIGHTINGElement) AttrsMap(attrs map[string]string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) Text(text string) *SVGFEDIFFUSELIGHTINGElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) TextF(format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfText(condition bool, text string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfTextF(condition bool, format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) Escaped(text string) *SVGFEDIFFUSELIGHTINGElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfEscaped(condition bool, text string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) EscapedF(format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) CustomData(key, value string) *SVGFEDIFFUSELIGHTINGElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfCustomData(condition bool, key, value string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) CustomDataF(key, format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) CustomDataRemove(key string) *SVGFEDIFFUSELIGHTINGElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The input for this filter.
func (e *SVGFEDIFFUSELIGHTINGElement) IN(s string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) INF(format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	return e.IN(fmt.Sprintf(format, args...))
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfIN(condition bool, s string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.IN(s)
	}
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfINF(condition bool, format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.IN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) INRemove(s string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("in")
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) INRemoveF(format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	return e.INRemove(fmt.Sprintf(format, args...))
}

// The 'surfaceScale' attribute indicates the height of the surface when the alpha
// channel is 1.0.
func (e *SVGFEDIFFUSELIGHTINGElement) SURFACE_SCALE(f float64) *SVGFEDIFFUSELIGHTINGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("surfaceScale", f)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfSURFACE_SCALE(condition bool, f float64) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.SURFACE_SCALE(f)
	}
	return e
}

// The diffuseConstant attribute represents the proportion of the light that is
// reflected by the surface.
func (e *SVGFEDIFFUSELIGHTINGElement) DIFFUSE_CONSTANT(f float64) *SVGFEDIFFUSELIGHTINGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("diffuseConstant", f)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfDIFFUSE_CONSTANT(condition bool, f float64) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.DIFFUSE_CONSTANT(f)
	}
	return e
}

// The kernelUnitLength attribute defines the intended distance in current filter
// units (i.e., units as determined by the value of attribute 'primitiveUnits')
// for dx and dy in the surface normal calculation formulas.
func (e *SVGFEDIFFUSELIGHTINGElement) KERNEL_UNIT_LENGTH(s string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("kernelUnitLength", s)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) KERNEL_UNIT_LENGTHF(format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	return e.KERNEL_UNIT_LENGTH(fmt.Sprintf(format, args...))
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfKERNEL_UNIT_LENGTH(condition bool, s string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.KERNEL_UNIT_LENGTH(s)
	}
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfKERNEL_UNIT_LENGTHF(condition bool, format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.KERNEL_UNIT_LENGTH(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute KERNEL_UNIT_LENGTH from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) KERNEL_UNIT_LENGTHRemove(s string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("kernelUnitLength")
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) KERNEL_UNIT_LENGTHRemoveF(format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	return e.KERNEL_UNIT_LENGTHRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGFEDIFFUSELIGHTINGElement) ID(s string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IDF(format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfID(condition bool, s string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfIDF(condition bool, format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) IDRemove(s string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IDRemoveF(format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEDIFFUSELIGHTINGElement) CLASS(s ...string) *SVGFEDIFFUSELIGHTINGElement {
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

func (e *SVGFEDIFFUSELIGHTINGElement) IfCLASS(condition bool, s ...string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) CLASSRemove(s ...string) *SVGFEDIFFUSELIGHTINGElement {
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
func (e *SVGFEDIFFUSELIGHTINGElement) STYLEF(k string, format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfSTYLE(condition bool, k string, v string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) STYLE(k string, v string) *SVGFEDIFFUSELIGHTINGElement {
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

func (e *SVGFEDIFFUSELIGHTINGElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEDIFFUSELIGHTINGElement) STYLEMap(m map[string]string) *SVGFEDIFFUSELIGHTINGElement {
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
func (e *SVGFEDIFFUSELIGHTINGElement) STYLEPairs(pairs ...string) *SVGFEDIFFUSELIGHTINGElement {
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

func (e *SVGFEDIFFUSELIGHTINGElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) STYLERemove(keys ...string) *SVGFEDIFFUSELIGHTINGElement {
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

func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQ(expression string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfZ_REQ(condition bool, expression string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQRemove() *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEDIFFUSELIGHTINGElement) Z_TARGET(expression string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfZ_TARGET(condition bool, expression string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) Z_TARGETRemove() *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQ_SELECTOR(expression string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQ_SELECTORRemove() *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEDIFFUSELIGHTINGElement) Z_SWAP(expression string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfZ_SWAP(condition bool, expression string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) Z_SWAPRemove() *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEDIFFUSELIGHTINGElement) Z_SWAP_PUSH(expression string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) Z_SWAP_PUSHRemove() *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEDIFFUSELIGHTINGElement) Z_TRIGGER(expression string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) Z_TRIGGERRemove() *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQ_METHOD(c SVGFeDiffuseLightingZReqMethodChoice) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeDiffuseLightingZReqMethodChoice string

const (
	// default GET
	SVGFeDiffuseLightingZReqMethod_empty SVGFeDiffuseLightingZReqMethodChoice = ""
	// GET
	SVGFeDiffuseLightingZReqMethod_get SVGFeDiffuseLightingZReqMethodChoice = "get"
	// POST
	SVGFeDiffuseLightingZReqMethod_post SVGFeDiffuseLightingZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQ_METHODRemove(c SVGFeDiffuseLightingZReqMethodChoice) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQ_STRATEGY(expression string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQ_STRATEGYRemove() *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQ_HISTORY(expression string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQ_HISTORYRemove() *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEDIFFUSELIGHTINGElement) Z_DATA(expression string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfZ_DATA(condition bool, expression string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) Z_DATARemove() *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEDIFFUSELIGHTINGElement) Z_JSON(expression string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfZ_JSON(condition bool, expression string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) Z_JSONRemove() *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQ_BATCH(expression string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQ_BATCHRemove() *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEDIFFUSELIGHTINGElement) Z_ACTION(expression string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfZ_ACTION(condition bool, expression string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) Z_ACTIONRemove() *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQ_BEFORE(expression string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQ_BEFORERemove() *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQ_AFTER(expression string) *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDIFFUSELIGHTINGElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEDIFFUSELIGHTINGElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEDIFFUSELIGHTINGElement) Z_REQ_AFTERRemove() *SVGFEDIFFUSELIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
