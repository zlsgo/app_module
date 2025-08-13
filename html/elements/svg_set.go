package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <set> SVG element provides a simple means of just setting the value of an
// attribute for a specified duration
// It supports all attribute types, including those that cannot reasonably be
// interpolated, such as string and boolean values
// The <set> element is non-additive
// The additive and accumulate attributes are not allowed, and will be ignored if
// specified.
type SVGSETElement struct {
	*Element
}

// Create a new SVGSETElement element.
// This will create a new element with the tag
// "set" during rendering.
func SVG_SET(children ...ElementRenderer) *SVGSETElement {
	e := NewElement("set", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGSETElement{Element: e}
}

func (e *SVGSETElement) Children(children ...ElementRenderer) *SVGSETElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGSETElement) IfChildren(condition bool, children ...ElementRenderer) *SVGSETElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGSETElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGSETElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGSETElement) Attr(name string, value ...string) *SVGSETElement {
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

func (e *SVGSETElement) Attrs(attrs ...string) *SVGSETElement {
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

func (e *SVGSETElement) AttrsMap(attrs map[string]string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGSETElement) Text(text string) *SVGSETElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGSETElement) TextF(format string, args ...any) *SVGSETElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGSETElement) IfText(condition bool, text string) *SVGSETElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGSETElement) IfTextF(condition bool, format string, args ...any) *SVGSETElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGSETElement) Escaped(text string) *SVGSETElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGSETElement) IfEscaped(condition bool, text string) *SVGSETElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGSETElement) EscapedF(format string, args ...any) *SVGSETElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGSETElement) IfEscapedF(condition bool, format string, args ...any) *SVGSETElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGSETElement) CustomData(key, value string) *SVGSETElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGSETElement) IfCustomData(condition bool, key, value string) *SVGSETElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGSETElement) CustomDataF(key, format string, args ...any) *SVGSETElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGSETElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGSETElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGSETElement) CustomDataRemove(key string) *SVGSETElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The target attribute value to assign on end.
func (e *SVGSETElement) TO(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("to", s)
	return e
}

func (e *SVGSETElement) TOF(format string, args ...any) *SVGSETElement {
	return e.TO(fmt.Sprintf(format, args...))
}

func (e *SVGSETElement) IfTO(condition bool, s string) *SVGSETElement {
	if condition {
		e.TO(s)
	}
	return e
}

func (e *SVGSETElement) IfTOF(condition bool, format string, args ...any) *SVGSETElement {
	if condition {
		e.TO(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TO from the element.
func (e *SVGSETElement) TORemove(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("to")
	return e
}

func (e *SVGSETElement) TORemoveF(format string, args ...any) *SVGSETElement {
	return e.TORemove(fmt.Sprintf(format, args...))
}

// The name of the attribute to assign.
func (e *SVGSETElement) ATTRIBUTE_NAME(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("attributeName", s)
	return e
}

func (e *SVGSETElement) ATTRIBUTE_NAMEF(format string, args ...any) *SVGSETElement {
	return e.ATTRIBUTE_NAME(fmt.Sprintf(format, args...))
}

func (e *SVGSETElement) IfATTRIBUTE_NAME(condition bool, s string) *SVGSETElement {
	if condition {
		e.ATTRIBUTE_NAME(s)
	}
	return e
}

func (e *SVGSETElement) IfATTRIBUTE_NAMEF(condition bool, format string, args ...any) *SVGSETElement {
	if condition {
		e.ATTRIBUTE_NAME(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ATTRIBUTE_NAME from the element.
func (e *SVGSETElement) ATTRIBUTE_NAMERemove(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("attributeName")
	return e
}

func (e *SVGSETElement) ATTRIBUTE_NAMERemoveF(format string, args ...any) *SVGSETElement {
	return e.ATTRIBUTE_NAMERemove(fmt.Sprintf(format, args...))
}

// The namespace in which the target attribute and its associated values are
// defined.
func (e *SVGSETElement) ATTRIBUTE_TYPE(c SVGSetAttributeTypeChoice) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("attributeType", string(c))
	return e
}

type SVGSetAttributeTypeChoice string

const (
	// The namespace in which the target attribute and its associated values are
	// defined.
	SVGSetAttributeType_auto SVGSetAttributeTypeChoice = "auto"
	// The namespace in which the target attribute and its associated values are
	// defined.
	SVGSetAttributeType_CSS SVGSetAttributeTypeChoice = "CSS"
	// The namespace in which the target attribute and its associated values are
	// defined.
	SVGSetAttributeType_XML SVGSetAttributeTypeChoice = "XML"
	// The namespace in which the target attribute and its associated values are
	// defined.
	SVGSetAttributeType_XMLNS SVGSetAttributeTypeChoice = "XMLNS"
	// The namespace in which the target attribute and its associated values are
	// defined.
	SVGSetAttributeType_empty SVGSetAttributeTypeChoice = "empty"
)

// Remove the attribute ATTRIBUTE_TYPE from the element.
func (e *SVGSETElement) ATTRIBUTE_TYPERemove(c SVGSetAttributeTypeChoice) *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("attributeType")
	return e
}

// The begin time for the element.
func (e *SVGSETElement) BEGIN(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("begin", s)
	return e
}

func (e *SVGSETElement) BEGINF(format string, args ...any) *SVGSETElement {
	return e.BEGIN(fmt.Sprintf(format, args...))
}

func (e *SVGSETElement) IfBEGIN(condition bool, s string) *SVGSETElement {
	if condition {
		e.BEGIN(s)
	}
	return e
}

func (e *SVGSETElement) IfBEGINF(condition bool, format string, args ...any) *SVGSETElement {
	if condition {
		e.BEGIN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute BEGIN from the element.
func (e *SVGSETElement) BEGINRemove(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("begin")
	return e
}

func (e *SVGSETElement) BEGINRemoveF(format string, args ...any) *SVGSETElement {
	return e.BEGINRemove(fmt.Sprintf(format, args...))
}

// The simple duration for the element.
func (e *SVGSETElement) DUR(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("dur", s)
	return e
}

func (e *SVGSETElement) DURF(format string, args ...any) *SVGSETElement {
	return e.DUR(fmt.Sprintf(format, args...))
}

func (e *SVGSETElement) IfDUR(condition bool, s string) *SVGSETElement {
	if condition {
		e.DUR(s)
	}
	return e
}

func (e *SVGSETElement) IfDURF(condition bool, format string, args ...any) *SVGSETElement {
	if condition {
		e.DUR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute DUR from the element.
func (e *SVGSETElement) DURRemove(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("dur")
	return e
}

func (e *SVGSETElement) DURRemoveF(format string, args ...any) *SVGSETElement {
	return e.DURRemove(fmt.Sprintf(format, args...))
}

// The end for the element.
func (e *SVGSETElement) END(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("end", s)
	return e
}

func (e *SVGSETElement) ENDF(format string, args ...any) *SVGSETElement {
	return e.END(fmt.Sprintf(format, args...))
}

func (e *SVGSETElement) IfEND(condition bool, s string) *SVGSETElement {
	if condition {
		e.END(s)
	}
	return e
}

func (e *SVGSETElement) IfENDF(condition bool, format string, args ...any) *SVGSETElement {
	if condition {
		e.END(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute END from the element.
func (e *SVGSETElement) ENDRemove(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("end")
	return e
}

func (e *SVGSETElement) ENDRemoveF(format string, args ...any) *SVGSETElement {
	return e.ENDRemove(fmt.Sprintf(format, args...))
}

// The minimum value allowed for the attribute.
func (e *SVGSETElement) MIN(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("min", s)
	return e
}

func (e *SVGSETElement) MINF(format string, args ...any) *SVGSETElement {
	return e.MIN(fmt.Sprintf(format, args...))
}

func (e *SVGSETElement) IfMIN(condition bool, s string) *SVGSETElement {
	if condition {
		e.MIN(s)
	}
	return e
}

func (e *SVGSETElement) IfMINF(condition bool, format string, args ...any) *SVGSETElement {
	if condition {
		e.MIN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MIN from the element.
func (e *SVGSETElement) MINRemove(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("min")
	return e
}

func (e *SVGSETElement) MINRemoveF(format string, args ...any) *SVGSETElement {
	return e.MINRemove(fmt.Sprintf(format, args...))
}

// The maximum value allowed for the attribute.
func (e *SVGSETElement) MAX(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("max", s)
	return e
}

func (e *SVGSETElement) MAXF(format string, args ...any) *SVGSETElement {
	return e.MAX(fmt.Sprintf(format, args...))
}

func (e *SVGSETElement) IfMAX(condition bool, s string) *SVGSETElement {
	if condition {
		e.MAX(s)
	}
	return e
}

func (e *SVGSETElement) IfMAXF(condition bool, format string, args ...any) *SVGSETElement {
	if condition {
		e.MAX(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MAX from the element.
func (e *SVGSETElement) MAXRemove(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("max")
	return e
}

func (e *SVGSETElement) MAXRemoveF(format string, args ...any) *SVGSETElement {
	return e.MAXRemove(fmt.Sprintf(format, args...))
}

// Defines how the element is restarted.
func (e *SVGSETElement) RESTART(c SVGSetRestartChoice) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("restart", string(c))
	return e
}

type SVGSetRestartChoice string

const (
	// Defines how the element is restarted.
	SVGSetRestart_always SVGSetRestartChoice = "always"
	// Defines how the element is restarted.
	SVGSetRestart_whenNotActive SVGSetRestartChoice = "whenNotActive"
	// Defines how the element is restarted.
	SVGSetRestart_never SVGSetRestartChoice = "never"
)

// Remove the attribute RESTART from the element.
func (e *SVGSETElement) RESTARTRemove(c SVGSetRestartChoice) *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("restart")
	return e
}

// Defines the number of times the element is repeated.
func (e *SVGSETElement) REPEAT_COUNT(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("repeatCount", s)
	return e
}

func (e *SVGSETElement) REPEAT_COUNTF(format string, args ...any) *SVGSETElement {
	return e.REPEAT_COUNT(fmt.Sprintf(format, args...))
}

func (e *SVGSETElement) IfREPEAT_COUNT(condition bool, s string) *SVGSETElement {
	if condition {
		e.REPEAT_COUNT(s)
	}
	return e
}

func (e *SVGSETElement) IfREPEAT_COUNTF(condition bool, format string, args ...any) *SVGSETElement {
	if condition {
		e.REPEAT_COUNT(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REPEAT_COUNT from the element.
func (e *SVGSETElement) REPEAT_COUNTRemove(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("repeatCount")
	return e
}

func (e *SVGSETElement) REPEAT_COUNTRemoveF(format string, args ...any) *SVGSETElement {
	return e.REPEAT_COUNTRemove(fmt.Sprintf(format, args...))
}

// Defines the duration for the element to repeat.
func (e *SVGSETElement) REPEAT_DUR(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("repeatDur", s)
	return e
}

func (e *SVGSETElement) REPEAT_DURF(format string, args ...any) *SVGSETElement {
	return e.REPEAT_DUR(fmt.Sprintf(format, args...))
}

func (e *SVGSETElement) IfREPEAT_DUR(condition bool, s string) *SVGSETElement {
	if condition {
		e.REPEAT_DUR(s)
	}
	return e
}

func (e *SVGSETElement) IfREPEAT_DURF(condition bool, format string, args ...any) *SVGSETElement {
	if condition {
		e.REPEAT_DUR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REPEAT_DUR from the element.
func (e *SVGSETElement) REPEAT_DURRemove(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("repeatDur")
	return e
}

func (e *SVGSETElement) REPEAT_DURRemoveF(format string, args ...any) *SVGSETElement {
	return e.REPEAT_DURRemove(fmt.Sprintf(format, args...))
}

// Defines the value the animation will have before the begin event.
func (e *SVGSETElement) FILL(c SVGSetFillChoice) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("fill", string(c))
	return e
}

type SVGSetFillChoice string

const (
	// Defines the value the animation will have before the begin event.
	SVGSetFill_remove SVGSetFillChoice = "remove"
	// Defines the value the animation will have before the begin event.
	SVGSetFill_freeze SVGSetFillChoice = "freeze"
)

// Remove the attribute FILL from the element.
func (e *SVGSETElement) FILLRemove(c SVGSetFillChoice) *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("fill")
	return e
}

// Specifies a unique id for an element
func (e *SVGSETElement) ID(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGSETElement) IDF(format string, args ...any) *SVGSETElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGSETElement) IfID(condition bool, s string) *SVGSETElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGSETElement) IfIDF(condition bool, format string, args ...any) *SVGSETElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGSETElement) IDRemove(s string) *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGSETElement) IDRemoveF(format string, args ...any) *SVGSETElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGSETElement) CLASS(s ...string) *SVGSETElement {
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

func (e *SVGSETElement) IfCLASS(condition bool, s ...string) *SVGSETElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGSETElement) CLASSRemove(s ...string) *SVGSETElement {
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
func (e *SVGSETElement) STYLEF(k string, format string, args ...any) *SVGSETElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGSETElement) IfSTYLE(condition bool, k string, v string) *SVGSETElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGSETElement) STYLE(k string, v string) *SVGSETElement {
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

func (e *SVGSETElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGSETElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGSETElement) STYLEMap(m map[string]string) *SVGSETElement {
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
func (e *SVGSETElement) STYLEPairs(pairs ...string) *SVGSETElement {
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

func (e *SVGSETElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGSETElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGSETElement) STYLERemove(keys ...string) *SVGSETElement {
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

func (e *SVGSETElement) Z_REQ(expression string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSETElement) IfZ_REQ(condition bool, expression string) *SVGSETElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGSETElement) Z_REQRemove() *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGSETElement) Z_TARGET(expression string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSETElement) IfZ_TARGET(condition bool, expression string) *SVGSETElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGSETElement) Z_TARGETRemove() *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGSETElement) Z_REQ_SELECTOR(expression string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSETElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGSETElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGSETElement) Z_REQ_SELECTORRemove() *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGSETElement) Z_SWAP(expression string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSETElement) IfZ_SWAP(condition bool, expression string) *SVGSETElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGSETElement) Z_SWAPRemove() *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGSETElement) Z_SWAP_PUSH(expression string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSETElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGSETElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGSETElement) Z_SWAP_PUSHRemove() *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGSETElement) Z_TRIGGER(expression string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSETElement) IfZ_TRIGGER(condition bool, expression string) *SVGSETElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGSETElement) Z_TRIGGERRemove() *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGSETElement) Z_REQ_METHOD(c SVGSetZReqMethodChoice) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGSetZReqMethodChoice string

const (
	// default GET
	SVGSetZReqMethod_empty SVGSetZReqMethodChoice = ""
	// GET
	SVGSetZReqMethod_get SVGSetZReqMethodChoice = "get"
	// POST
	SVGSetZReqMethod_post SVGSetZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGSETElement) Z_REQ_METHODRemove(c SVGSetZReqMethodChoice) *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGSETElement) Z_REQ_STRATEGY(expression string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSETElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGSETElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGSETElement) Z_REQ_STRATEGYRemove() *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGSETElement) Z_REQ_HISTORY(expression string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSETElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGSETElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGSETElement) Z_REQ_HISTORYRemove() *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGSETElement) Z_DATA(expression string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSETElement) IfZ_DATA(condition bool, expression string) *SVGSETElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGSETElement) Z_DATARemove() *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGSETElement) Z_JSON(expression string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSETElement) IfZ_JSON(condition bool, expression string) *SVGSETElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGSETElement) Z_JSONRemove() *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGSETElement) Z_REQ_BATCH(expression string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSETElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGSETElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGSETElement) Z_REQ_BATCHRemove() *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGSETElement) Z_ACTION(expression string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSETElement) IfZ_ACTION(condition bool, expression string) *SVGSETElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGSETElement) Z_ACTIONRemove() *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGSETElement) Z_REQ_BEFORE(expression string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSETElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGSETElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGSETElement) Z_REQ_BEFORERemove() *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGSETElement) Z_REQ_AFTER(expression string) *SVGSETElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSETElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGSETElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGSETElement) Z_REQ_AFTERRemove() *SVGSETElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
