package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <radialGradient> SVG element lets authors define radial gradients to fill
// or stroke graphical elements.
type SVGRADIALGRADIENTElement struct {
	*Element
}

// Create a new SVGRADIALGRADIENTElement element.
// This will create a new element with the tag
// "radialGradient" during rendering.
func SVG_RADIALGRADIENT(children ...ElementRenderer) *SVGRADIALGRADIENTElement {
	e := NewElement("radialGradient", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGRADIALGRADIENTElement{Element: e}
}

func (e *SVGRADIALGRADIENTElement) Children(children ...ElementRenderer) *SVGRADIALGRADIENTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfChildren(condition bool, children ...ElementRenderer) *SVGRADIALGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGRADIALGRADIENTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGRADIALGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGRADIALGRADIENTElement) Attr(name string, value ...string) *SVGRADIALGRADIENTElement {
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

func (e *SVGRADIALGRADIENTElement) Attrs(attrs ...string) *SVGRADIALGRADIENTElement {
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

func (e *SVGRADIALGRADIENTElement) AttrsMap(attrs map[string]string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGRADIALGRADIENTElement) Text(text string) *SVGRADIALGRADIENTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGRADIALGRADIENTElement) TextF(format string, args ...any) *SVGRADIALGRADIENTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGRADIALGRADIENTElement) IfText(condition bool, text string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGRADIALGRADIENTElement) IfTextF(condition bool, format string, args ...any) *SVGRADIALGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGRADIALGRADIENTElement) Escaped(text string) *SVGRADIALGRADIENTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGRADIALGRADIENTElement) IfEscaped(condition bool, text string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGRADIALGRADIENTElement) EscapedF(format string, args ...any) *SVGRADIALGRADIENTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGRADIALGRADIENTElement) IfEscapedF(condition bool, format string, args ...any) *SVGRADIALGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGRADIALGRADIENTElement) CustomData(key, value string) *SVGRADIALGRADIENTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfCustomData(condition bool, key, value string) *SVGRADIALGRADIENTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGRADIALGRADIENTElement) CustomDataF(key, format string, args ...any) *SVGRADIALGRADIENTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGRADIALGRADIENTElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGRADIALGRADIENTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGRADIALGRADIENTElement) CustomDataRemove(key string) *SVGRADIALGRADIENTElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The coordinate system for attributes cx, cy and r.
func (e *SVGRADIALGRADIENTElement) GRADIENT_UNITS(c SVGRadialGradientGradientUnitsChoice) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("gradientUnits", string(c))
	return e
}

type SVGRadialGradientGradientUnitsChoice string

const (
	// The coordinate system for attributes cx, cy and r.
	SVGRadialGradientGradientUnits_userSpaceOnUse SVGRadialGradientGradientUnitsChoice = "userSpaceOnUse"
	// The coordinate system for attributes cx, cy and r.
	SVGRadialGradientGradientUnits_objectBoundingBox SVGRadialGradientGradientUnitsChoice = "objectBoundingBox"
)

// Remove the attribute GRADIENT_UNITS from the element.
func (e *SVGRADIALGRADIENTElement) GRADIENT_UNITSRemove(c SVGRadialGradientGradientUnitsChoice) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("gradientUnits")
	return e
}

// The definition of how the gradient is applied, read about <a
// href="https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/gradientTransform">gradientTransform</a>.
func (e *SVGRADIALGRADIENTElement) GRADIENT_TRANSFORM(s string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("gradientTransform", s)
	return e
}

func (e *SVGRADIALGRADIENTElement) GRADIENT_TRANSFORMF(format string, args ...any) *SVGRADIALGRADIENTElement {
	return e.GRADIENT_TRANSFORM(fmt.Sprintf(format, args...))
}

func (e *SVGRADIALGRADIENTElement) IfGRADIENT_TRANSFORM(condition bool, s string) *SVGRADIALGRADIENTElement {
	if condition {
		e.GRADIENT_TRANSFORM(s)
	}
	return e
}

func (e *SVGRADIALGRADIENTElement) IfGRADIENT_TRANSFORMF(condition bool, format string, args ...any) *SVGRADIALGRADIENTElement {
	if condition {
		e.GRADIENT_TRANSFORM(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute GRADIENT_TRANSFORM from the element.
func (e *SVGRADIALGRADIENTElement) GRADIENT_TRANSFORMRemove(s string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("gradientTransform")
	return e
}

func (e *SVGRADIALGRADIENTElement) GRADIENT_TRANSFORMRemoveF(format string, args ...any) *SVGRADIALGRADIENTElement {
	return e.GRADIENT_TRANSFORMRemove(fmt.Sprintf(format, args...))
}

// The x-axis coordinate of the largest (i.e., outermost) circle for the radial
// gradient.
func (e *SVGRADIALGRADIENTElement) CX(f float64) *SVGRADIALGRADIENTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("cx", f)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfCX(condition bool, f float64) *SVGRADIALGRADIENTElement {
	if condition {
		e.CX(f)
	}
	return e
}

// The y-axis coordinate of the largest (i.e., outermost) circle for the radial
// gradient.
func (e *SVGRADIALGRADIENTElement) CY(f float64) *SVGRADIALGRADIENTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("cy", f)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfCY(condition bool, f float64) *SVGRADIALGRADIENTElement {
	if condition {
		e.CY(f)
	}
	return e
}

// The radius of the largest (i.e., outermost) circle for the radial gradient.
func (e *SVGRADIALGRADIENTElement) R(f float64) *SVGRADIALGRADIENTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("r", f)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfR(condition bool, f float64) *SVGRADIALGRADIENTElement {
	if condition {
		e.R(f)
	}
	return e
}

// The x-axis coordinate of the point at which the focal point of the radial
// gradient is placed.
func (e *SVGRADIALGRADIENTElement) FX(f float64) *SVGRADIALGRADIENTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("fx", f)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfFX(condition bool, f float64) *SVGRADIALGRADIENTElement {
	if condition {
		e.FX(f)
	}
	return e
}

// The y-axis coordinate of the point at which the focal point of the radial
// gradient is placed.
func (e *SVGRADIALGRADIENTElement) FY(f float64) *SVGRADIALGRADIENTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("fy", f)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfFY(condition bool, f float64) *SVGRADIALGRADIENTElement {
	if condition {
		e.FY(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGRADIALGRADIENTElement) ID(s string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGRADIALGRADIENTElement) IDF(format string, args ...any) *SVGRADIALGRADIENTElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGRADIALGRADIENTElement) IfID(condition bool, s string) *SVGRADIALGRADIENTElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGRADIALGRADIENTElement) IfIDF(condition bool, format string, args ...any) *SVGRADIALGRADIENTElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGRADIALGRADIENTElement) IDRemove(s string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGRADIALGRADIENTElement) IDRemoveF(format string, args ...any) *SVGRADIALGRADIENTElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGRADIALGRADIENTElement) CLASS(s ...string) *SVGRADIALGRADIENTElement {
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

func (e *SVGRADIALGRADIENTElement) IfCLASS(condition bool, s ...string) *SVGRADIALGRADIENTElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGRADIALGRADIENTElement) CLASSRemove(s ...string) *SVGRADIALGRADIENTElement {
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
func (e *SVGRADIALGRADIENTElement) STYLEF(k string, format string, args ...any) *SVGRADIALGRADIENTElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGRADIALGRADIENTElement) IfSTYLE(condition bool, k string, v string) *SVGRADIALGRADIENTElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGRADIALGRADIENTElement) STYLE(k string, v string) *SVGRADIALGRADIENTElement {
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

func (e *SVGRADIALGRADIENTElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGRADIALGRADIENTElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGRADIALGRADIENTElement) STYLEMap(m map[string]string) *SVGRADIALGRADIENTElement {
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
func (e *SVGRADIALGRADIENTElement) STYLEPairs(pairs ...string) *SVGRADIALGRADIENTElement {
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

func (e *SVGRADIALGRADIENTElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGRADIALGRADIENTElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGRADIALGRADIENTElement) STYLERemove(keys ...string) *SVGRADIALGRADIENTElement {
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

func (e *SVGRADIALGRADIENTElement) Z_REQ(expression string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfZ_REQ(condition bool, expression string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGRADIALGRADIENTElement) Z_REQRemove() *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGRADIALGRADIENTElement) Z_TARGET(expression string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfZ_TARGET(condition bool, expression string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGRADIALGRADIENTElement) Z_TARGETRemove() *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGRADIALGRADIENTElement) Z_REQ_SELECTOR(expression string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGRADIALGRADIENTElement) Z_REQ_SELECTORRemove() *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGRADIALGRADIENTElement) Z_SWAP(expression string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfZ_SWAP(condition bool, expression string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGRADIALGRADIENTElement) Z_SWAPRemove() *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGRADIALGRADIENTElement) Z_SWAP_PUSH(expression string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGRADIALGRADIENTElement) Z_SWAP_PUSHRemove() *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGRADIALGRADIENTElement) Z_TRIGGER(expression string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfZ_TRIGGER(condition bool, expression string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGRADIALGRADIENTElement) Z_TRIGGERRemove() *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGRADIALGRADIENTElement) Z_REQ_METHOD(c SVGRadialGradientZReqMethodChoice) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGRadialGradientZReqMethodChoice string

const (
	// default GET
	SVGRadialGradientZReqMethod_empty SVGRadialGradientZReqMethodChoice = ""
	// GET
	SVGRadialGradientZReqMethod_get SVGRadialGradientZReqMethodChoice = "get"
	// POST
	SVGRadialGradientZReqMethod_post SVGRadialGradientZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGRADIALGRADIENTElement) Z_REQ_METHODRemove(c SVGRadialGradientZReqMethodChoice) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGRADIALGRADIENTElement) Z_REQ_STRATEGY(expression string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGRADIALGRADIENTElement) Z_REQ_STRATEGYRemove() *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGRADIALGRADIENTElement) Z_REQ_HISTORY(expression string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGRADIALGRADIENTElement) Z_REQ_HISTORYRemove() *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGRADIALGRADIENTElement) Z_DATA(expression string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfZ_DATA(condition bool, expression string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGRADIALGRADIENTElement) Z_DATARemove() *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGRADIALGRADIENTElement) Z_JSON(expression string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfZ_JSON(condition bool, expression string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGRADIALGRADIENTElement) Z_JSONRemove() *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGRADIALGRADIENTElement) Z_REQ_BATCH(expression string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGRADIALGRADIENTElement) Z_REQ_BATCHRemove() *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGRADIALGRADIENTElement) Z_ACTION(expression string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfZ_ACTION(condition bool, expression string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGRADIALGRADIENTElement) Z_ACTIONRemove() *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGRADIALGRADIENTElement) Z_REQ_BEFORE(expression string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGRADIALGRADIENTElement) Z_REQ_BEFORERemove() *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGRADIALGRADIENTElement) Z_REQ_AFTER(expression string) *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGRADIALGRADIENTElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGRADIALGRADIENTElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGRADIALGRADIENTElement) Z_REQ_AFTERRemove() *SVGRADIALGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
