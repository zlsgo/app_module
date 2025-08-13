package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feSpotLight> SVG filter primitive allows to create a light source placed
// at a point x, y, z.
type SVGFESPOTLIGHTElement struct {
	*Element
}

// Create a new SVGFESPOTLIGHTElement element.
// This will create a new element with the tag
// "feSpotLight" during rendering.
func SVG_FESPOTLIGHT(children ...ElementRenderer) *SVGFESPOTLIGHTElement {
	e := NewElement("feSpotLight", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFESPOTLIGHTElement{Element: e}
}

func (e *SVGFESPOTLIGHTElement) Children(children ...ElementRenderer) *SVGFESPOTLIGHTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFESPOTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFESPOTLIGHTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFESPOTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFESPOTLIGHTElement) Attr(name string, value ...string) *SVGFESPOTLIGHTElement {
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

func (e *SVGFESPOTLIGHTElement) Attrs(attrs ...string) *SVGFESPOTLIGHTElement {
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

func (e *SVGFESPOTLIGHTElement) AttrsMap(attrs map[string]string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFESPOTLIGHTElement) Text(text string) *SVGFESPOTLIGHTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFESPOTLIGHTElement) TextF(format string, args ...any) *SVGFESPOTLIGHTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFESPOTLIGHTElement) IfText(condition bool, text string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFESPOTLIGHTElement) IfTextF(condition bool, format string, args ...any) *SVGFESPOTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFESPOTLIGHTElement) Escaped(text string) *SVGFESPOTLIGHTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFESPOTLIGHTElement) IfEscaped(condition bool, text string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFESPOTLIGHTElement) EscapedF(format string, args ...any) *SVGFESPOTLIGHTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFESPOTLIGHTElement) IfEscapedF(condition bool, format string, args ...any) *SVGFESPOTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFESPOTLIGHTElement) CustomData(key, value string) *SVGFESPOTLIGHTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfCustomData(condition bool, key, value string) *SVGFESPOTLIGHTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFESPOTLIGHTElement) CustomDataF(key, format string, args ...any) *SVGFESPOTLIGHTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFESPOTLIGHTElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFESPOTLIGHTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFESPOTLIGHTElement) CustomDataRemove(key string) *SVGFESPOTLIGHTElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The x attribute indicates the x location of the light source in the coordinate
// system established by attribute 'primitiveUnits' on the <filter> element.
func (e *SVGFESPOTLIGHTElement) X(f float64) *SVGFESPOTLIGHTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("x", f)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfX(condition bool, f float64) *SVGFESPOTLIGHTElement {
	if condition {
		e.X(f)
	}
	return e
}

// The y attribute indicates the y location of the light source in the coordinate
// system established by attribute 'primitiveUnits' on the <filter> element.
func (e *SVGFESPOTLIGHTElement) Y(f float64) *SVGFESPOTLIGHTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("y", f)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfY(condition bool, f float64) *SVGFESPOTLIGHTElement {
	if condition {
		e.Y(f)
	}
	return e
}

// The z attribute indicates the z location of the light source in the coordinate
// system established by attribute 'primitiveUnits' on the <filter> element.
func (e *SVGFESPOTLIGHTElement) Z(f float64) *SVGFESPOTLIGHTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("z", f)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfZ(condition bool, f float64) *SVGFESPOTLIGHTElement {
	if condition {
		e.Z(f)
	}
	return e
}

// The pointsAtX attribute indicates the x location in the coordinate system
// established by attribute 'primitiveUnits' on the <filter> element of the point
// at which the light source is pointing.
func (e *SVGFESPOTLIGHTElement) POINTS_AT_X(f float64) *SVGFESPOTLIGHTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("pointsAtX", f)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfPOINTS_AT_X(condition bool, f float64) *SVGFESPOTLIGHTElement {
	if condition {
		e.POINTS_AT_X(f)
	}
	return e
}

// The pointsAtY attribute indicates the y location in the coordinate system
// established by attribute 'primitiveUnits' on the <filter> element of the point
// at which the light source is pointing.
func (e *SVGFESPOTLIGHTElement) POINTS_AT_Y(f float64) *SVGFESPOTLIGHTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("pointsAtY", f)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfPOINTS_AT_Y(condition bool, f float64) *SVGFESPOTLIGHTElement {
	if condition {
		e.POINTS_AT_Y(f)
	}
	return e
}

// The pointsAtZ attribute indicates the z location in the coordinate system
// established by attribute 'primitiveUnits' on the <filter> element of the point
// at which the light source is pointing.
func (e *SVGFESPOTLIGHTElement) POINTS_AT_Z(f float64) *SVGFESPOTLIGHTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("pointsAtZ", f)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfPOINTS_AT_Z(condition bool, f float64) *SVGFESPOTLIGHTElement {
	if condition {
		e.POINTS_AT_Z(f)
	}
	return e
}

// The specularExponent attribute represents the specular reflection constant.
func (e *SVGFESPOTLIGHTElement) SPECULAR_EXPONENT(f float64) *SVGFESPOTLIGHTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("specularExponent", f)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfSPECULAR_EXPONENT(condition bool, f float64) *SVGFESPOTLIGHTElement {
	if condition {
		e.SPECULAR_EXPONENT(f)
	}
	return e
}

// The limitingConeAngle attribute represents the angle in degrees between the
// spot light axis and the spot light cone.
func (e *SVGFESPOTLIGHTElement) LIMITING_CONE_ANGLE(f float64) *SVGFESPOTLIGHTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("limitingConeAngle", f)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfLIMITING_CONE_ANGLE(condition bool, f float64) *SVGFESPOTLIGHTElement {
	if condition {
		e.LIMITING_CONE_ANGLE(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFESPOTLIGHTElement) ID(s string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFESPOTLIGHTElement) IDF(format string, args ...any) *SVGFESPOTLIGHTElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFESPOTLIGHTElement) IfID(condition bool, s string) *SVGFESPOTLIGHTElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFESPOTLIGHTElement) IfIDF(condition bool, format string, args ...any) *SVGFESPOTLIGHTElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFESPOTLIGHTElement) IDRemove(s string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFESPOTLIGHTElement) IDRemoveF(format string, args ...any) *SVGFESPOTLIGHTElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFESPOTLIGHTElement) CLASS(s ...string) *SVGFESPOTLIGHTElement {
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

func (e *SVGFESPOTLIGHTElement) IfCLASS(condition bool, s ...string) *SVGFESPOTLIGHTElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFESPOTLIGHTElement) CLASSRemove(s ...string) *SVGFESPOTLIGHTElement {
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
func (e *SVGFESPOTLIGHTElement) STYLEF(k string, format string, args ...any) *SVGFESPOTLIGHTElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFESPOTLIGHTElement) IfSTYLE(condition bool, k string, v string) *SVGFESPOTLIGHTElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFESPOTLIGHTElement) STYLE(k string, v string) *SVGFESPOTLIGHTElement {
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

func (e *SVGFESPOTLIGHTElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFESPOTLIGHTElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFESPOTLIGHTElement) STYLEMap(m map[string]string) *SVGFESPOTLIGHTElement {
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
func (e *SVGFESPOTLIGHTElement) STYLEPairs(pairs ...string) *SVGFESPOTLIGHTElement {
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

func (e *SVGFESPOTLIGHTElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFESPOTLIGHTElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFESPOTLIGHTElement) STYLERemove(keys ...string) *SVGFESPOTLIGHTElement {
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

func (e *SVGFESPOTLIGHTElement) Z_REQ(expression string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfZ_REQ(condition bool, expression string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFESPOTLIGHTElement) Z_REQRemove() *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFESPOTLIGHTElement) Z_TARGET(expression string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfZ_TARGET(condition bool, expression string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFESPOTLIGHTElement) Z_TARGETRemove() *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFESPOTLIGHTElement) Z_REQ_SELECTOR(expression string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFESPOTLIGHTElement) Z_REQ_SELECTORRemove() *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFESPOTLIGHTElement) Z_SWAP(expression string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfZ_SWAP(condition bool, expression string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFESPOTLIGHTElement) Z_SWAPRemove() *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFESPOTLIGHTElement) Z_SWAP_PUSH(expression string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFESPOTLIGHTElement) Z_SWAP_PUSHRemove() *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFESPOTLIGHTElement) Z_TRIGGER(expression string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfZ_TRIGGER(condition bool, expression string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFESPOTLIGHTElement) Z_TRIGGERRemove() *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFESPOTLIGHTElement) Z_REQ_METHOD(c SVGFeSpotLightZReqMethodChoice) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeSpotLightZReqMethodChoice string

const (
	// default GET
	SVGFeSpotLightZReqMethod_empty SVGFeSpotLightZReqMethodChoice = ""
	// GET
	SVGFeSpotLightZReqMethod_get SVGFeSpotLightZReqMethodChoice = "get"
	// POST
	SVGFeSpotLightZReqMethod_post SVGFeSpotLightZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFESPOTLIGHTElement) Z_REQ_METHODRemove(c SVGFeSpotLightZReqMethodChoice) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFESPOTLIGHTElement) Z_REQ_STRATEGY(expression string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFESPOTLIGHTElement) Z_REQ_STRATEGYRemove() *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFESPOTLIGHTElement) Z_REQ_HISTORY(expression string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFESPOTLIGHTElement) Z_REQ_HISTORYRemove() *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFESPOTLIGHTElement) Z_DATA(expression string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfZ_DATA(condition bool, expression string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFESPOTLIGHTElement) Z_DATARemove() *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFESPOTLIGHTElement) Z_JSON(expression string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfZ_JSON(condition bool, expression string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFESPOTLIGHTElement) Z_JSONRemove() *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFESPOTLIGHTElement) Z_REQ_BATCH(expression string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFESPOTLIGHTElement) Z_REQ_BATCHRemove() *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFESPOTLIGHTElement) Z_ACTION(expression string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfZ_ACTION(condition bool, expression string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFESPOTLIGHTElement) Z_ACTIONRemove() *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFESPOTLIGHTElement) Z_REQ_BEFORE(expression string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFESPOTLIGHTElement) Z_REQ_BEFORERemove() *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFESPOTLIGHTElement) Z_REQ_AFTER(expression string) *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPOTLIGHTElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFESPOTLIGHTElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFESPOTLIGHTElement) Z_REQ_AFTERRemove() *SVGFESPOTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
