package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feSpecularLighting> SVG filter primitive lights a source graphic using the
// alpha channel as a bump map
// The resulting image is an RGBA image based on the light color
// The lighting calculation follows the standard specular component of the Phong
// lighting model
// The resulting image depends on the light color, light position and surface
// geometry of the input bump map.
type SVGFESPECULARLIGHTINGElement struct {
	*Element
}

// Create a new SVGFESPECULARLIGHTINGElement element.
// This will create a new element with the tag
// "feSpecularLighting" during rendering.
func SVG_FESPECULARLIGHTING(children ...ElementRenderer) *SVGFESPECULARLIGHTINGElement {
	e := NewElement("feSpecularLighting", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFESPECULARLIGHTINGElement{Element: e}
}

func (e *SVGFESPECULARLIGHTINGElement) Children(children ...ElementRenderer) *SVGFESPECULARLIGHTINGElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) Attr(name string, value ...string) *SVGFESPECULARLIGHTINGElement {
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

func (e *SVGFESPECULARLIGHTINGElement) Attrs(attrs ...string) *SVGFESPECULARLIGHTINGElement {
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

func (e *SVGFESPECULARLIGHTINGElement) AttrsMap(attrs map[string]string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) Text(text string) *SVGFESPECULARLIGHTINGElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) TextF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFESPECULARLIGHTINGElement) IfText(condition bool, text string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfTextF(condition bool, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) Escaped(text string) *SVGFESPECULARLIGHTINGElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfEscaped(condition bool, text string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) EscapedF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFESPECULARLIGHTINGElement) IfEscapedF(condition bool, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) CustomData(key, value string) *SVGFESPECULARLIGHTINGElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfCustomData(condition bool, key, value string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) CustomDataF(key, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFESPECULARLIGHTINGElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) CustomDataRemove(key string) *SVGFESPECULARLIGHTINGElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The input for this filter.
func (e *SVGFESPECULARLIGHTINGElement) IN(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) INF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.IN(fmt.Sprintf(format, args...))
}

func (e *SVGFESPECULARLIGHTINGElement) IfIN(condition bool, s string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.IN(s)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfINF(condition bool, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.IN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFESPECULARLIGHTINGElement) INRemove(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("in")
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) INRemoveF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.INRemove(fmt.Sprintf(format, args...))
}

// The 'surfaceScale' attribute indicates the height of the surface when the alpha
// channel is 1.0.
func (e *SVGFESPECULARLIGHTINGElement) SURFACE_SCALE(f float64) *SVGFESPECULARLIGHTINGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("surfaceScale", f)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfSURFACE_SCALE(condition bool, f float64) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.SURFACE_SCALE(f)
	}
	return e
}

// The specularConstant attribute represents the diffuse reflection constant.
func (e *SVGFESPECULARLIGHTINGElement) SPECULAR_CONSTANT(f float64) *SVGFESPECULARLIGHTINGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("specularConstant", f)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfSPECULAR_CONSTANT(condition bool, f float64) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.SPECULAR_CONSTANT(f)
	}
	return e
}

// The specularExponent attribute represents the specular reflection constant.
func (e *SVGFESPECULARLIGHTINGElement) SPECULAR_EXPONENT(f float64) *SVGFESPECULARLIGHTINGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("specularExponent", f)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfSPECULAR_EXPONENT(condition bool, f float64) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.SPECULAR_EXPONENT(f)
	}
	return e
}

// The kernelUnitLength attribute defines the intended distance in current filter
// units (i.e., units as determined by the value of attribute 'primitiveUnits')
// for dx and dy in the surface normal calculation formulas.
func (e *SVGFESPECULARLIGHTINGElement) KERNEL_UNIT_LENGTH(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("kernelUnitLength", s)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) KERNEL_UNIT_LENGTHF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.KERNEL_UNIT_LENGTH(fmt.Sprintf(format, args...))
}

func (e *SVGFESPECULARLIGHTINGElement) IfKERNEL_UNIT_LENGTH(condition bool, s string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.KERNEL_UNIT_LENGTH(s)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfKERNEL_UNIT_LENGTHF(condition bool, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.KERNEL_UNIT_LENGTH(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute KERNEL_UNIT_LENGTH from the element.
func (e *SVGFESPECULARLIGHTINGElement) KERNEL_UNIT_LENGTHRemove(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("kernelUnitLength")
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) KERNEL_UNIT_LENGTHRemoveF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.KERNEL_UNIT_LENGTHRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGFESPECULARLIGHTINGElement) ID(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IDF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFESPECULARLIGHTINGElement) IfID(condition bool, s string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfIDF(condition bool, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFESPECULARLIGHTINGElement) IDRemove(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IDRemoveF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFESPECULARLIGHTINGElement) CLASS(s ...string) *SVGFESPECULARLIGHTINGElement {
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

func (e *SVGFESPECULARLIGHTINGElement) IfCLASS(condition bool, s ...string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFESPECULARLIGHTINGElement) CLASSRemove(s ...string) *SVGFESPECULARLIGHTINGElement {
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
func (e *SVGFESPECULARLIGHTINGElement) STYLEF(k string, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFESPECULARLIGHTINGElement) IfSTYLE(condition bool, k string, v string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) STYLE(k string, v string) *SVGFESPECULARLIGHTINGElement {
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

func (e *SVGFESPECULARLIGHTINGElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFESPECULARLIGHTINGElement) STYLEMap(m map[string]string) *SVGFESPECULARLIGHTINGElement {
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
func (e *SVGFESPECULARLIGHTINGElement) STYLEPairs(pairs ...string) *SVGFESPECULARLIGHTINGElement {
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

func (e *SVGFESPECULARLIGHTINGElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFESPECULARLIGHTINGElement) STYLERemove(keys ...string) *SVGFESPECULARLIGHTINGElement {
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

func (e *SVGFESPECULARLIGHTINGElement) Z_REQ(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfZ_REQ(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFESPECULARLIGHTINGElement) Z_REQRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFESPECULARLIGHTINGElement) Z_TARGET(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfZ_TARGET(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFESPECULARLIGHTINGElement) Z_TARGETRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFESPECULARLIGHTINGElement) Z_REQ_SELECTOR(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFESPECULARLIGHTINGElement) Z_REQ_SELECTORRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFESPECULARLIGHTINGElement) Z_SWAP(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfZ_SWAP(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFESPECULARLIGHTINGElement) Z_SWAPRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFESPECULARLIGHTINGElement) Z_SWAP_PUSH(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFESPECULARLIGHTINGElement) Z_SWAP_PUSHRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFESPECULARLIGHTINGElement) Z_TRIGGER(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfZ_TRIGGER(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFESPECULARLIGHTINGElement) Z_TRIGGERRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFESPECULARLIGHTINGElement) Z_REQ_METHOD(c SVGFeSpecularLightingZReqMethodChoice) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeSpecularLightingZReqMethodChoice string

const (
	// default GET
	SVGFeSpecularLightingZReqMethod_empty SVGFeSpecularLightingZReqMethodChoice = ""
	// GET
	SVGFeSpecularLightingZReqMethod_get SVGFeSpecularLightingZReqMethodChoice = "get"
	// POST
	SVGFeSpecularLightingZReqMethod_post SVGFeSpecularLightingZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFESPECULARLIGHTINGElement) Z_REQ_METHODRemove(c SVGFeSpecularLightingZReqMethodChoice) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFESPECULARLIGHTINGElement) Z_REQ_STRATEGY(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFESPECULARLIGHTINGElement) Z_REQ_STRATEGYRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFESPECULARLIGHTINGElement) Z_REQ_HISTORY(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFESPECULARLIGHTINGElement) Z_REQ_HISTORYRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFESPECULARLIGHTINGElement) Z_DATA(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfZ_DATA(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFESPECULARLIGHTINGElement) Z_DATARemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFESPECULARLIGHTINGElement) Z_JSON(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfZ_JSON(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFESPECULARLIGHTINGElement) Z_JSONRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFESPECULARLIGHTINGElement) Z_REQ_BATCH(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFESPECULARLIGHTINGElement) Z_REQ_BATCHRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFESPECULARLIGHTINGElement) Z_ACTION(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfZ_ACTION(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFESPECULARLIGHTINGElement) Z_ACTIONRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFESPECULARLIGHTINGElement) Z_REQ_BEFORE(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFESPECULARLIGHTINGElement) Z_REQ_BEFORERemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFESPECULARLIGHTINGElement) Z_REQ_AFTER(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFESPECULARLIGHTINGElement) Z_REQ_AFTERRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
