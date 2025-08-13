package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <animateTransform> SVG element animates a transformation attribute on a
// target element, thereby allowing animations to control translation, scaling,
// rotation and/or skewing.
type SVGANIMATETRANSFORMElement struct {
	*Element
}

// Create a new SVGANIMATETRANSFORMElement element.
// This will create a new element with the tag
// "animateTransform" during rendering.
func SVG_ANIMATETRANSFORM(children ...ElementRenderer) *SVGANIMATETRANSFORMElement {
	e := NewElement("animateTransform", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGANIMATETRANSFORMElement{Element: e}
}

func (e *SVGANIMATETRANSFORMElement) Children(children ...ElementRenderer) *SVGANIMATETRANSFORMElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfChildren(condition bool, children ...ElementRenderer) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) Attr(name string, value ...string) *SVGANIMATETRANSFORMElement {
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

func (e *SVGANIMATETRANSFORMElement) Attrs(attrs ...string) *SVGANIMATETRANSFORMElement {
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

func (e *SVGANIMATETRANSFORMElement) AttrsMap(attrs map[string]string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) Text(text string) *SVGANIMATETRANSFORMElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGANIMATETRANSFORMElement) TextF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfText(condition bool, text string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfTextF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) Escaped(text string) *SVGANIMATETRANSFORMElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfEscaped(condition bool, text string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) EscapedF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfEscapedF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) CustomData(key, value string) *SVGANIMATETRANSFORMElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfCustomData(condition bool, key, value string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) CustomDataF(key, format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) CustomDataRemove(key string) *SVGANIMATETRANSFORMElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// Controls whether or not the animation is cumulative.
func (e *SVGANIMATETRANSFORMElement) ACCUMULATE(c SVGAnimateTransformAccumulateChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("accumulate", string(c))
	return e
}

type SVGAnimateTransformAccumulateChoice string

const (
	// The animation is not cumulative
	// Each iteration starts over from the beginning.
	SVGAnimateTransformAccumulate_none SVGAnimateTransformAccumulateChoice = "none"
	// The animation is cumulative
	// Each iteration the animation picks up where it left off in the previous
	// iteration.
	SVGAnimateTransformAccumulate_sum SVGAnimateTransformAccumulateChoice = "sum"
)

// Remove the attribute ACCUMULATE from the element.
func (e *SVGANIMATETRANSFORMElement) ACCUMULATERemove(c SVGAnimateTransformAccumulateChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("accumulate")
	return e
}

// Controls whether or not the animation is additive.
func (e *SVGANIMATETRANSFORMElement) ADDITIVE(c SVGAnimateTransformAdditiveChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("additive", string(c))
	return e
}

type SVGAnimateTransformAdditiveChoice string

const (
	// The animation is not additive
	// The animation replaces the underlying value.
	SVGAnimateTransformAdditive_replace SVGAnimateTransformAdditiveChoice = "replace"
	// The animation is additive
	// The animation adds to the underlying value.
	SVGAnimateTransformAdditive_sum SVGAnimateTransformAdditiveChoice = "sum"
)

// Remove the attribute ADDITIVE from the element.
func (e *SVGANIMATETRANSFORMElement) ADDITIVERemove(c SVGAnimateTransformAdditiveChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("additive")
	return e
}

// The name of the attribute to animate.
func (e *SVGANIMATETRANSFORMElement) ATTRIBUTE_NAME(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("attributeName", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) ATTRIBUTE_NAMEF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.ATTRIBUTE_NAME(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfATTRIBUTE_NAME(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.ATTRIBUTE_NAME(s)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfATTRIBUTE_NAMEF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.ATTRIBUTE_NAME(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ATTRIBUTE_NAME from the element.
func (e *SVGANIMATETRANSFORMElement) ATTRIBUTE_NAMERemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("attributeName")
	return e
}

func (e *SVGANIMATETRANSFORMElement) ATTRIBUTE_NAMERemoveF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.ATTRIBUTE_NAMERemove(fmt.Sprintf(format, args...))
}

// The namespace of the attribute to animate.
func (e *SVGANIMATETRANSFORMElement) ATTRIBUTE_TYPE(c SVGAnimateTransformAttributeTypeChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("attributeType", string(c))
	return e
}

type SVGAnimateTransformAttributeTypeChoice string

const (
	// If the attribute is a presentation attribute, the animation will use the target
	// element's corresponding baseVal
	// If the attribute is not a presentation attribute, the animation will use the
	// target element's corresponding animVal.
	SVGAnimateTransformAttributeType_auto SVGAnimateTransformAttributeTypeChoice = "auto"
	// The animation will use the CSS namespace.
	SVGAnimateTransformAttributeType_CSS SVGAnimateTransformAttributeTypeChoice = "CSS"
	// The animation will use the XML namespace.
	SVGAnimateTransformAttributeType_XML SVGAnimateTransformAttributeTypeChoice = "XML"
	// The animation will use the XML ID namespace.
	SVGAnimateTransformAttributeType_XMLID SVGAnimateTransformAttributeTypeChoice = "XMLID"
	// The animation will use the XML LANG namespace.
	SVGAnimateTransformAttributeType_XMLLANG SVGAnimateTransformAttributeTypeChoice = "XMLLANG"
	// The animation will use the XML SPACE namespace.
	SVGAnimateTransformAttributeType_XMLSPACE SVGAnimateTransformAttributeTypeChoice = "XMLSPACE"
)

// Remove the attribute ATTRIBUTE_TYPE from the element.
func (e *SVGANIMATETRANSFORMElement) ATTRIBUTE_TYPERemove(c SVGAnimateTransformAttributeTypeChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("attributeType")
	return e
}

// Defines when the animation should begin.
func (e *SVGANIMATETRANSFORMElement) BEGIN(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("begin", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) BEGINF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.BEGIN(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfBEGIN(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.BEGIN(s)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfBEGINF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.BEGIN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute BEGIN from the element.
func (e *SVGANIMATETRANSFORMElement) BEGINRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("begin")
	return e
}

func (e *SVGANIMATETRANSFORMElement) BEGINRemoveF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.BEGINRemove(fmt.Sprintf(format, args...))
}

// Defines a relative offset value for the animation.
func (e *SVGANIMATETRANSFORMElement) BY(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("by", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) BYF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.BY(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfBY(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.BY(s)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfBYF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.BY(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute BY from the element.
func (e *SVGANIMATETRANSFORMElement) BYRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("by")
	return e
}

func (e *SVGANIMATETRANSFORMElement) BYRemoveF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.BYRemove(fmt.Sprintf(format, args...))
}

// Defines the pacing of the animation.
func (e *SVGANIMATETRANSFORMElement) CALC_MODE(c SVGAnimateTransformCalcModeChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("calcMode", string(c))
	return e
}

type SVGAnimateTransformCalcModeChoice string

const (
	// The animation is not paced
	// Each iteration of the animation is displayed as fast as possible.
	SVGAnimateTransformCalcMode_discrete SVGAnimateTransformCalcModeChoice = "discrete"
	// The animation is paced such that it takes the same amount of time to go from
	// the start value to the end value throughout the animation.
	SVGAnimateTransformCalcMode_linear SVGAnimateTransformCalcModeChoice = "linear"
	// The animation is paced according to a cubic function.
	SVGAnimateTransformCalcMode_paced SVGAnimateTransformCalcModeChoice = "paced"
	// The animation is paced according to a cubic function, but with easing at both
	// the start and end.
	SVGAnimateTransformCalcMode_spline SVGAnimateTransformCalcModeChoice = "spline"
)

// Remove the attribute CALC_MODE from the element.
func (e *SVGANIMATETRANSFORMElement) CALC_MODERemove(c SVGAnimateTransformCalcModeChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("calcMode")
	return e
}

// Defines the duration of the animation.
func (e *SVGANIMATETRANSFORMElement) DUR(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("dur", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) DURF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.DUR(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfDUR(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.DUR(s)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfDURF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.DUR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute DUR from the element.
func (e *SVGANIMATETRANSFORMElement) DURRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("dur")
	return e
}

func (e *SVGANIMATETRANSFORMElement) DURRemoveF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.DURRemove(fmt.Sprintf(format, args...))
}

// Defines when the animation should end.
func (e *SVGANIMATETRANSFORMElement) END(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("end", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) ENDF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.END(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfEND(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.END(s)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfENDF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.END(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute END from the element.
func (e *SVGANIMATETRANSFORMElement) ENDRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("end")
	return e
}

func (e *SVGANIMATETRANSFORMElement) ENDRemoveF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.ENDRemove(fmt.Sprintf(format, args...))
}

// Defines the fill behavior for the animation.
func (e *SVGANIMATETRANSFORMElement) FILL(c SVGAnimateTransformFillChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("fill", string(c))
	return e
}

type SVGAnimateTransformFillChoice string

const (
	// The animation will hold the attribute value when the animation ends.
	SVGAnimateTransformFill_freeze SVGAnimateTransformFillChoice = "freeze"
	// The animation will remove the attribute value when the animation ends.
	SVGAnimateTransformFill_remove SVGAnimateTransformFillChoice = "remove"
)

// Remove the attribute FILL from the element.
func (e *SVGANIMATETRANSFORMElement) FILLRemove(c SVGAnimateTransformFillChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("fill")
	return e
}

// Defines the initial value of the attribute.
func (e *SVGANIMATETRANSFORMElement) FROM(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("from", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) FROMF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.FROM(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfFROM(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.FROM(s)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfFROMF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.FROM(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute FROM from the element.
func (e *SVGANIMATETRANSFORMElement) FROMRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("from")
	return e
}

func (e *SVGANIMATETRANSFORMElement) FROMRemoveF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.FROMRemove(fmt.Sprintf(format, args...))
}

// Defines the values for a cubic Bézier function that controls interval pacing.
func (e *SVGANIMATETRANSFORMElement) KEY_SPLINES(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("keySplines", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) KEY_SPLINESF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.KEY_SPLINES(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfKEY_SPLINES(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.KEY_SPLINES(s)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfKEY_SPLINESF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.KEY_SPLINES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute KEY_SPLINES from the element.
func (e *SVGANIMATETRANSFORMElement) KEY_SPLINESRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("keySplines")
	return e
}

func (e *SVGANIMATETRANSFORMElement) KEY_SPLINESRemoveF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.KEY_SPLINESRemove(fmt.Sprintf(format, args...))
}

// Defines when the animation should take place in terms of time fractions.
func (e *SVGANIMATETRANSFORMElement) KEY_TIMES(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("keyTimes", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) KEY_TIMESF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.KEY_TIMES(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfKEY_TIMES(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.KEY_TIMES(s)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfKEY_TIMESF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.KEY_TIMES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute KEY_TIMES from the element.
func (e *SVGANIMATETRANSFORMElement) KEY_TIMESRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("keyTimes")
	return e
}

func (e *SVGANIMATETRANSFORMElement) KEY_TIMESRemoveF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.KEY_TIMESRemove(fmt.Sprintf(format, args...))
}

// Defines the maximum value allowed for the attribute.
func (e *SVGANIMATETRANSFORMElement) MAX(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("max", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) MAXF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.MAX(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfMAX(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.MAX(s)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfMAXF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.MAX(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MAX from the element.
func (e *SVGANIMATETRANSFORMElement) MAXRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("max")
	return e
}

func (e *SVGANIMATETRANSFORMElement) MAXRemoveF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.MAXRemove(fmt.Sprintf(format, args...))
}

// Defines the minimum value allowed for the attribute.
func (e *SVGANIMATETRANSFORMElement) MIN(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("min", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) MINF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.MIN(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfMIN(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.MIN(s)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfMINF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.MIN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MIN from the element.
func (e *SVGANIMATETRANSFORMElement) MINRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("min")
	return e
}

func (e *SVGANIMATETRANSFORMElement) MINRemoveF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.MINRemove(fmt.Sprintf(format, args...))
}

// Defines the number of times the animation should repeat.
func (e *SVGANIMATETRANSFORMElement) REPEAT_COUNT(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("repeatCount", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) REPEAT_COUNTF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.REPEAT_COUNT(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfREPEAT_COUNT(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.REPEAT_COUNT(s)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfREPEAT_COUNTF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.REPEAT_COUNT(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REPEAT_COUNT from the element.
func (e *SVGANIMATETRANSFORMElement) REPEAT_COUNTRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("repeatCount")
	return e
}

func (e *SVGANIMATETRANSFORMElement) REPEAT_COUNTRemoveF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.REPEAT_COUNTRemove(fmt.Sprintf(format, args...))
}

// Defines the duration for repeating an animation.
func (e *SVGANIMATETRANSFORMElement) REPEAT_DUR(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("repeatDur", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) REPEAT_DURF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.REPEAT_DUR(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfREPEAT_DUR(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.REPEAT_DUR(s)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfREPEAT_DURF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.REPEAT_DUR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REPEAT_DUR from the element.
func (e *SVGANIMATETRANSFORMElement) REPEAT_DURRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("repeatDur")
	return e
}

func (e *SVGANIMATETRANSFORMElement) REPEAT_DURRemoveF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.REPEAT_DURRemove(fmt.Sprintf(format, args...))
}

// Defines if an animation should restart after it completes.
func (e *SVGANIMATETRANSFORMElement) RESTART(c SVGAnimateTransformRestartChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("restart", string(c))
	return e
}

type SVGAnimateTransformRestartChoice string

const (
	// The animation will restart indefinitely.
	SVGAnimateTransformRestart_always SVGAnimateTransformRestartChoice = "always"
	// The animation will not restart after it completes.
	SVGAnimateTransformRestart_never SVGAnimateTransformRestartChoice = "never"
	// The animation will restart after it completes if the animation is not currently
	// active.
	SVGAnimateTransformRestart_whenNotActive SVGAnimateTransformRestartChoice = "whenNotActive"
)

// Remove the attribute RESTART from the element.
func (e *SVGANIMATETRANSFORMElement) RESTARTRemove(c SVGAnimateTransformRestartChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("restart")
	return e
}

// Defines the ending value of the attribute.
func (e *SVGANIMATETRANSFORMElement) TO(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("to", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) TOF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.TO(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfTO(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.TO(s)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfTOF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.TO(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TO from the element.
func (e *SVGANIMATETRANSFORMElement) TORemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("to")
	return e
}

func (e *SVGANIMATETRANSFORMElement) TORemoveF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.TORemove(fmt.Sprintf(format, args...))
}

// Defines which transform to use.
func (e *SVGANIMATETRANSFORMElement) TYPE(c SVGAnimateTransformTypeChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("type", string(c))
	return e
}

type SVGAnimateTransformTypeChoice string

const (
	// The animation will use the rotate transform.
	SVGAnimateTransformType_rotate SVGAnimateTransformTypeChoice = "rotate"
	// The animation will use the scale transform.
	SVGAnimateTransformType_scale SVGAnimateTransformTypeChoice = "scale"
	// The animation will use the translate transform.
	SVGAnimateTransformType_translate SVGAnimateTransformTypeChoice = "translate"
)

// Remove the attribute TYPE from the element.
func (e *SVGANIMATETRANSFORMElement) TYPERemove(c SVGAnimateTransformTypeChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("type")
	return e
}

// Defines a list of discrete values to interpolate.
func (e *SVGANIMATETRANSFORMElement) VALUES(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("values", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) VALUESF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.VALUES(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfVALUES(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.VALUES(s)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfVALUESF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.VALUES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute VALUES from the element.
func (e *SVGANIMATETRANSFORMElement) VALUESRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("values")
	return e
}

func (e *SVGANIMATETRANSFORMElement) VALUESRemoveF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.VALUESRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGANIMATETRANSFORMElement) ID(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IDF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfID(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfIDF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGANIMATETRANSFORMElement) IDRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGANIMATETRANSFORMElement) IDRemoveF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGANIMATETRANSFORMElement) CLASS(s ...string) *SVGANIMATETRANSFORMElement {
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

func (e *SVGANIMATETRANSFORMElement) IfCLASS(condition bool, s ...string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGANIMATETRANSFORMElement) CLASSRemove(s ...string) *SVGANIMATETRANSFORMElement {
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
func (e *SVGANIMATETRANSFORMElement) STYLEF(k string, format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfSTYLE(condition bool, k string, v string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) STYLE(k string, v string) *SVGANIMATETRANSFORMElement {
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

func (e *SVGANIMATETRANSFORMElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGANIMATETRANSFORMElement) STYLEMap(m map[string]string) *SVGANIMATETRANSFORMElement {
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
func (e *SVGANIMATETRANSFORMElement) STYLEPairs(pairs ...string) *SVGANIMATETRANSFORMElement {
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

func (e *SVGANIMATETRANSFORMElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGANIMATETRANSFORMElement) STYLERemove(keys ...string) *SVGANIMATETRANSFORMElement {
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

func (e *SVGANIMATETRANSFORMElement) Z_REQ(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfZ_REQ(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGANIMATETRANSFORMElement) Z_REQRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGANIMATETRANSFORMElement) Z_TARGET(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfZ_TARGET(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGANIMATETRANSFORMElement) Z_TARGETRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGANIMATETRANSFORMElement) Z_REQ_SELECTOR(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGANIMATETRANSFORMElement) Z_REQ_SELECTORRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGANIMATETRANSFORMElement) Z_SWAP(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfZ_SWAP(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGANIMATETRANSFORMElement) Z_SWAPRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGANIMATETRANSFORMElement) Z_SWAP_PUSH(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGANIMATETRANSFORMElement) Z_SWAP_PUSHRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGANIMATETRANSFORMElement) Z_TRIGGER(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfZ_TRIGGER(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGANIMATETRANSFORMElement) Z_TRIGGERRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGANIMATETRANSFORMElement) Z_REQ_METHOD(c SVGAnimateTransformZReqMethodChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGAnimateTransformZReqMethodChoice string

const (
	// default GET
	SVGAnimateTransformZReqMethod_empty SVGAnimateTransformZReqMethodChoice = ""
	// GET
	SVGAnimateTransformZReqMethod_get SVGAnimateTransformZReqMethodChoice = "get"
	// POST
	SVGAnimateTransformZReqMethod_post SVGAnimateTransformZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGANIMATETRANSFORMElement) Z_REQ_METHODRemove(c SVGAnimateTransformZReqMethodChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGANIMATETRANSFORMElement) Z_REQ_STRATEGY(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGANIMATETRANSFORMElement) Z_REQ_STRATEGYRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGANIMATETRANSFORMElement) Z_REQ_HISTORY(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGANIMATETRANSFORMElement) Z_REQ_HISTORYRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGANIMATETRANSFORMElement) Z_DATA(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfZ_DATA(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGANIMATETRANSFORMElement) Z_DATARemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGANIMATETRANSFORMElement) Z_JSON(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfZ_JSON(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGANIMATETRANSFORMElement) Z_JSONRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGANIMATETRANSFORMElement) Z_REQ_BATCH(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGANIMATETRANSFORMElement) Z_REQ_BATCHRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGANIMATETRANSFORMElement) Z_ACTION(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfZ_ACTION(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGANIMATETRANSFORMElement) Z_ACTIONRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGANIMATETRANSFORMElement) Z_REQ_BEFORE(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGANIMATETRANSFORMElement) Z_REQ_BEFORERemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGANIMATETRANSFORMElement) Z_REQ_AFTER(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGANIMATETRANSFORMElement) Z_REQ_AFTERRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
