package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <linearGradient> SVG element lets authors define linear gradients to fill
// or stroke graphical elements.
type SVGLINEARGRADIENTElement struct {
	*Element
}

// Create a new SVGLINEARGRADIENTElement element.
// This will create a new element with the tag
// "linearGradient" during rendering.
func SVG_LINEARGRADIENT(children ...ElementRenderer) *SVGLINEARGRADIENTElement {
	e := NewElement("linearGradient", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGLINEARGRADIENTElement{Element: e}
}

func (e *SVGLINEARGRADIENTElement) Children(children ...ElementRenderer) *SVGLINEARGRADIENTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfChildren(condition bool, children ...ElementRenderer) *SVGLINEARGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGLINEARGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) Attr(name string, value ...string) *SVGLINEARGRADIENTElement {
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

func (e *SVGLINEARGRADIENTElement) Attrs(attrs ...string) *SVGLINEARGRADIENTElement {
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

func (e *SVGLINEARGRADIENTElement) AttrsMap(attrs map[string]string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) Text(text string) *SVGLINEARGRADIENTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGLINEARGRADIENTElement) TextF(format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGLINEARGRADIENTElement) IfText(condition bool, text string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) IfTextF(condition bool, format string, args ...any) *SVGLINEARGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) Escaped(text string) *SVGLINEARGRADIENTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGLINEARGRADIENTElement) IfEscaped(condition bool, text string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) EscapedF(format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGLINEARGRADIENTElement) IfEscapedF(condition bool, format string, args ...any) *SVGLINEARGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) CustomData(key, value string) *SVGLINEARGRADIENTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfCustomData(condition bool, key, value string) *SVGLINEARGRADIENTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) CustomDataF(key, format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGLINEARGRADIENTElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGLINEARGRADIENTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) CustomDataRemove(key string) *SVGLINEARGRADIENTElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The coordinate system for attributes x1, y1, x2 and y2.
func (e *SVGLINEARGRADIENTElement) GRADIENT_UNITS(c SVGLinearGradientGradientUnitsChoice) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("gradientUnits", string(c))
	return e
}

type SVGLinearGradientGradientUnitsChoice string

const (
	// The coordinate system for attributes x1, y1, x2 and y2.
	SVGLinearGradientGradientUnits_userSpaceOnUse SVGLinearGradientGradientUnitsChoice = "userSpaceOnUse"
	// The coordinate system for attributes x1, y1, x2 and y2.
	SVGLinearGradientGradientUnits_objectBoundingBox SVGLinearGradientGradientUnitsChoice = "objectBoundingBox"
)

// Remove the attribute GRADIENT_UNITS from the element.
func (e *SVGLINEARGRADIENTElement) GRADIENT_UNITSRemove(c SVGLinearGradientGradientUnitsChoice) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("gradientUnits")
	return e
}

// The definition of how the gradient is applied, read about <a
// href="https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/gradientTransform">gradientTransform</a>.
func (e *SVGLINEARGRADIENTElement) GRADIENT_TRANSFORM(s string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("gradientTransform", s)
	return e
}

func (e *SVGLINEARGRADIENTElement) GRADIENT_TRANSFORMF(format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.GRADIENT_TRANSFORM(fmt.Sprintf(format, args...))
}

func (e *SVGLINEARGRADIENTElement) IfGRADIENT_TRANSFORM(condition bool, s string) *SVGLINEARGRADIENTElement {
	if condition {
		e.GRADIENT_TRANSFORM(s)
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) IfGRADIENT_TRANSFORMF(condition bool, format string, args ...any) *SVGLINEARGRADIENTElement {
	if condition {
		e.GRADIENT_TRANSFORM(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute GRADIENT_TRANSFORM from the element.
func (e *SVGLINEARGRADIENTElement) GRADIENT_TRANSFORMRemove(s string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("gradientTransform")
	return e
}

func (e *SVGLINEARGRADIENTElement) GRADIENT_TRANSFORMRemoveF(format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.GRADIENT_TRANSFORMRemove(fmt.Sprintf(format, args...))
}

// The method by which to fill a shape.
func (e *SVGLINEARGRADIENTElement) SPREAD_METHOD(c SVGLinearGradientSpreadMethodChoice) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("spreadMethod", string(c))
	return e
}

type SVGLinearGradientSpreadMethodChoice string

const (
	// The method by which to fill a shape.
	SVGLinearGradientSpreadMethod_pad SVGLinearGradientSpreadMethodChoice = "pad"
	// The method by which to fill a shape.
	SVGLinearGradientSpreadMethod_reflect SVGLinearGradientSpreadMethodChoice = "reflect"
	// The method by which to fill a shape.
	SVGLinearGradientSpreadMethod_repeat SVGLinearGradientSpreadMethodChoice = "repeat"
)

// Remove the attribute SPREAD_METHOD from the element.
func (e *SVGLINEARGRADIENTElement) SPREAD_METHODRemove(c SVGLinearGradientSpreadMethodChoice) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("spreadMethod")
	return e
}

// The x-axis coordinate of the start of the gradient.
func (e *SVGLINEARGRADIENTElement) X_1(f float64) *SVGLINEARGRADIENTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("x1", f)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfX_1(condition bool, f float64) *SVGLINEARGRADIENTElement {
	if condition {
		e.X_1(f)
	}
	return e
}

// The y-axis coordinate of the start of the gradient.
func (e *SVGLINEARGRADIENTElement) Y_1(f float64) *SVGLINEARGRADIENTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("y1", f)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfY_1(condition bool, f float64) *SVGLINEARGRADIENTElement {
	if condition {
		e.Y_1(f)
	}
	return e
}

// The x-axis coordinate of the end of the gradient.
func (e *SVGLINEARGRADIENTElement) X_2(f float64) *SVGLINEARGRADIENTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("x2", f)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfX_2(condition bool, f float64) *SVGLINEARGRADIENTElement {
	if condition {
		e.X_2(f)
	}
	return e
}

// The y-axis coordinate of the end of the gradient.
func (e *SVGLINEARGRADIENTElement) Y_2(f float64) *SVGLINEARGRADIENTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("y2", f)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfY_2(condition bool, f float64) *SVGLINEARGRADIENTElement {
	if condition {
		e.Y_2(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGLINEARGRADIENTElement) ID(s string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGLINEARGRADIENTElement) IDF(format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGLINEARGRADIENTElement) IfID(condition bool, s string) *SVGLINEARGRADIENTElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) IfIDF(condition bool, format string, args ...any) *SVGLINEARGRADIENTElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGLINEARGRADIENTElement) IDRemove(s string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGLINEARGRADIENTElement) IDRemoveF(format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGLINEARGRADIENTElement) CLASS(s ...string) *SVGLINEARGRADIENTElement {
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

func (e *SVGLINEARGRADIENTElement) IfCLASS(condition bool, s ...string) *SVGLINEARGRADIENTElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGLINEARGRADIENTElement) CLASSRemove(s ...string) *SVGLINEARGRADIENTElement {
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
func (e *SVGLINEARGRADIENTElement) STYLEF(k string, format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGLINEARGRADIENTElement) IfSTYLE(condition bool, k string, v string) *SVGLINEARGRADIENTElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) STYLE(k string, v string) *SVGLINEARGRADIENTElement {
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

func (e *SVGLINEARGRADIENTElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGLINEARGRADIENTElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGLINEARGRADIENTElement) STYLEMap(m map[string]string) *SVGLINEARGRADIENTElement {
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
func (e *SVGLINEARGRADIENTElement) STYLEPairs(pairs ...string) *SVGLINEARGRADIENTElement {
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

func (e *SVGLINEARGRADIENTElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGLINEARGRADIENTElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGLINEARGRADIENTElement) STYLERemove(keys ...string) *SVGLINEARGRADIENTElement {
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

func (e *SVGLINEARGRADIENTElement) Z_REQ(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfZ_REQ(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGLINEARGRADIENTElement) Z_REQRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGLINEARGRADIENTElement) Z_TARGET(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfZ_TARGET(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGLINEARGRADIENTElement) Z_TARGETRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGLINEARGRADIENTElement) Z_REQ_SELECTOR(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGLINEARGRADIENTElement) Z_REQ_SELECTORRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGLINEARGRADIENTElement) Z_SWAP(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfZ_SWAP(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGLINEARGRADIENTElement) Z_SWAPRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGLINEARGRADIENTElement) Z_SWAP_PUSH(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGLINEARGRADIENTElement) Z_SWAP_PUSHRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGLINEARGRADIENTElement) Z_TRIGGER(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfZ_TRIGGER(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGLINEARGRADIENTElement) Z_TRIGGERRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGLINEARGRADIENTElement) Z_REQ_METHOD(c SVGLinearGradientZReqMethodChoice) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGLinearGradientZReqMethodChoice string

const (
	// default GET
	SVGLinearGradientZReqMethod_empty SVGLinearGradientZReqMethodChoice = ""
	// GET
	SVGLinearGradientZReqMethod_get SVGLinearGradientZReqMethodChoice = "get"
	// POST
	SVGLinearGradientZReqMethod_post SVGLinearGradientZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGLINEARGRADIENTElement) Z_REQ_METHODRemove(c SVGLinearGradientZReqMethodChoice) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGLINEARGRADIENTElement) Z_REQ_STRATEGY(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGLINEARGRADIENTElement) Z_REQ_STRATEGYRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGLINEARGRADIENTElement) Z_REQ_HISTORY(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGLINEARGRADIENTElement) Z_REQ_HISTORYRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGLINEARGRADIENTElement) Z_DATA(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfZ_DATA(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGLINEARGRADIENTElement) Z_DATARemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGLINEARGRADIENTElement) Z_JSON(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfZ_JSON(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGLINEARGRADIENTElement) Z_JSONRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGLINEARGRADIENTElement) Z_REQ_BATCH(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGLINEARGRADIENTElement) Z_REQ_BATCHRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGLINEARGRADIENTElement) Z_ACTION(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfZ_ACTION(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGLINEARGRADIENTElement) Z_ACTIONRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGLINEARGRADIENTElement) Z_REQ_BEFORE(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGLINEARGRADIENTElement) Z_REQ_BEFORERemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGLINEARGRADIENTElement) Z_REQ_AFTER(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGLINEARGRADIENTElement) Z_REQ_AFTERRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
