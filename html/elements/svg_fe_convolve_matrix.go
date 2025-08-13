package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feConvolveMatrix> SVG filter primitive applies a matrix convolution filter
// effect.
type SVGFECONVOLVEMATRIXElement struct {
	*Element
}

// Create a new SVGFECONVOLVEMATRIXElement element.
// This will create a new element with the tag
// "feConvolveMatrix" during rendering.
func SVG_FECONVOLVEMATRIX(children ...ElementRenderer) *SVGFECONVOLVEMATRIXElement {
	e := NewElement("feConvolveMatrix", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFECONVOLVEMATRIXElement{Element: e}
}

func (e *SVGFECONVOLVEMATRIXElement) Children(children ...ElementRenderer) *SVGFECONVOLVEMATRIXElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) Attr(name string, value ...string) *SVGFECONVOLVEMATRIXElement {
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

func (e *SVGFECONVOLVEMATRIXElement) Attrs(attrs ...string) *SVGFECONVOLVEMATRIXElement {
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

func (e *SVGFECONVOLVEMATRIXElement) AttrsMap(attrs map[string]string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) Text(text string) *SVGFECONVOLVEMATRIXElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) TextF(format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFECONVOLVEMATRIXElement) IfText(condition bool, text string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfTextF(condition bool, format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) Escaped(text string) *SVGFECONVOLVEMATRIXElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfEscaped(condition bool, text string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) EscapedF(format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFECONVOLVEMATRIXElement) IfEscapedF(condition bool, format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) CustomData(key, value string) *SVGFECONVOLVEMATRIXElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfCustomData(condition bool, key, value string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) CustomDataF(key, format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFECONVOLVEMATRIXElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) CustomDataRemove(key string) *SVGFECONVOLVEMATRIXElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The input for this filter.
func (e *SVGFECONVOLVEMATRIXElement) IN(s string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) INF(format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	return e.IN(fmt.Sprintf(format, args...))
}

func (e *SVGFECONVOLVEMATRIXElement) IfIN(condition bool, s string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.IN(s)
	}
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfINF(condition bool, format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.IN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFECONVOLVEMATRIXElement) INRemove(s string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("in")
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) INRemoveF(format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	return e.INRemove(fmt.Sprintf(format, args...))
}

// The number of cells in each dimension for 'kernelMatrix'
func (e *SVGFECONVOLVEMATRIXElement) ORDER(s string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("order", s)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) ORDERF(format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	return e.ORDER(fmt.Sprintf(format, args...))
}

func (e *SVGFECONVOLVEMATRIXElement) IfORDER(condition bool, s string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.ORDER(s)
	}
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfORDERF(condition bool, format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.ORDER(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ORDER from the element.
func (e *SVGFECONVOLVEMATRIXElement) ORDERRemove(s string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("order")
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) ORDERRemoveF(format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	return e.ORDERRemove(fmt.Sprintf(format, args...))
}

// A list of numbers that make up the kernel matrix for the convolution.
func (e *SVGFECONVOLVEMATRIXElement) KERNEL_MATRIX(s string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("kernelMatrix", s)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) KERNEL_MATRIXF(format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	return e.KERNEL_MATRIX(fmt.Sprintf(format, args...))
}

func (e *SVGFECONVOLVEMATRIXElement) IfKERNEL_MATRIX(condition bool, s string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.KERNEL_MATRIX(s)
	}
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfKERNEL_MATRIXF(condition bool, format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.KERNEL_MATRIX(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute KERNEL_MATRIX from the element.
func (e *SVGFECONVOLVEMATRIXElement) KERNEL_MATRIXRemove(s string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("kernelMatrix")
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) KERNEL_MATRIXRemoveF(format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	return e.KERNEL_MATRIXRemove(fmt.Sprintf(format, args...))
}

// The divisor attribute specifies the value by which to divide the result of
// applying the convolution operator.
func (e *SVGFECONVOLVEMATRIXElement) DIVISOR(f float64) *SVGFECONVOLVEMATRIXElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("divisor", f)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfDIVISOR(condition bool, f float64) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.DIVISOR(f)
	}
	return e
}

// The bias attribute shifts the range of the filter
// After applying the matrix operation, this bias value is added to each
// component.
func (e *SVGFECONVOLVEMATRIXElement) BIAS(f float64) *SVGFECONVOLVEMATRIXElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("bias", f)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfBIAS(condition bool, f float64) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.BIAS(f)
	}
	return e
}

// The targetX attribute determines the positioning in X of the convolution matrix
// relative to a given target pixel in the input image.
func (e *SVGFECONVOLVEMATRIXElement) TARGET_X(f float64) *SVGFECONVOLVEMATRIXElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("targetX", f)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfTARGET_X(condition bool, f float64) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.TARGET_X(f)
	}
	return e
}

// The targetY attribute determines the positioning in Y of the convolution matrix
// relative to a given target pixel in the input image.
func (e *SVGFECONVOLVEMATRIXElement) TARGET_Y(f float64) *SVGFECONVOLVEMATRIXElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("targetY", f)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfTARGET_Y(condition bool, f float64) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.TARGET_Y(f)
	}
	return e
}

// The edgeMode attribute determines how to extend the input image as necessary
// with color values so that the matrix operations can be applied when the kernel
// is positioned at or near the edge of the input image.
func (e *SVGFECONVOLVEMATRIXElement) EDGE_MODE(c SVGFeConvolveMatrixEdgeModeChoice) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("edgeMode", string(c))
	return e
}

type SVGFeConvolveMatrixEdgeModeChoice string

const (
	// The input image is extended along each of its borders as necessary by
	// duplicating the color values at the given edge of the input image.
	SVGFeConvolveMatrixEdgeMode_duplicate SVGFeConvolveMatrixEdgeModeChoice = "duplicate"
	// The input image is extended by taking the component values from the opposite
	// edge of the image.
	SVGFeConvolveMatrixEdgeMode_wrap SVGFeConvolveMatrixEdgeModeChoice = "wrap"
	// Any values outside the input image are assumed to be transparent black.
	SVGFeConvolveMatrixEdgeMode_none SVGFeConvolveMatrixEdgeModeChoice = "none"
)

// Remove the attribute EDGE_MODE from the element.
func (e *SVGFECONVOLVEMATRIXElement) EDGE_MODERemove(c SVGFeConvolveMatrixEdgeModeChoice) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("edgeMode")
	return e
}

// The kernelUnitLength attribute defines the intended distance in current filter
// units (i.e., units as determined by the value of attribute 'primitiveUnits')
// for dx and dy in the surface normal calculation formulas.
func (e *SVGFECONVOLVEMATRIXElement) KERNEL_UNIT_LENGTH(s string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("kernelUnitLength", s)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) KERNEL_UNIT_LENGTHF(format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	return e.KERNEL_UNIT_LENGTH(fmt.Sprintf(format, args...))
}

func (e *SVGFECONVOLVEMATRIXElement) IfKERNEL_UNIT_LENGTH(condition bool, s string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.KERNEL_UNIT_LENGTH(s)
	}
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfKERNEL_UNIT_LENGTHF(condition bool, format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.KERNEL_UNIT_LENGTH(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute KERNEL_UNIT_LENGTH from the element.
func (e *SVGFECONVOLVEMATRIXElement) KERNEL_UNIT_LENGTHRemove(s string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("kernelUnitLength")
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) KERNEL_UNIT_LENGTHRemoveF(format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	return e.KERNEL_UNIT_LENGTHRemove(fmt.Sprintf(format, args...))
}

// The preserveAlpha attribute indicates how the convolution will handle the alpha
// channel of the input image.
func (e *SVGFECONVOLVEMATRIXElement) PRESERVE_ALPHA() *SVGFECONVOLVEMATRIXElement {
	e.PRESERVE_ALPHASet(true)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfPRESERVE_ALPHA(condition bool) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.PRESERVE_ALPHASet(true)
	}
	return e
}

// Set the attribute PRESERVE_ALPHA to the value b explicitly.
func (e *SVGFECONVOLVEMATRIXElement) PRESERVE_ALPHASet(b bool) *SVGFECONVOLVEMATRIXElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = zarray.NewSortMap[string, bool]()
	}
	e.BoolAttributes.Set("preserveAlpha", b)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfSetPRESERVE_ALPHA(condition bool, b bool) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.PRESERVE_ALPHASet(b)
	}
	return e
}

// Remove the attribute PRESERVE_ALPHA from the element.
func (e *SVGFECONVOLVEMATRIXElement) PRESERVE_ALPHARemove(b bool) *SVGFECONVOLVEMATRIXElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Delete("preserveAlpha")
	return e
}

// Specifies a unique id for an element
func (e *SVGFECONVOLVEMATRIXElement) ID(s string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IDF(format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFECONVOLVEMATRIXElement) IfID(condition bool, s string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfIDF(condition bool, format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFECONVOLVEMATRIXElement) IDRemove(s string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IDRemoveF(format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFECONVOLVEMATRIXElement) CLASS(s ...string) *SVGFECONVOLVEMATRIXElement {
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

func (e *SVGFECONVOLVEMATRIXElement) IfCLASS(condition bool, s ...string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFECONVOLVEMATRIXElement) CLASSRemove(s ...string) *SVGFECONVOLVEMATRIXElement {
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
func (e *SVGFECONVOLVEMATRIXElement) STYLEF(k string, format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFECONVOLVEMATRIXElement) IfSTYLE(condition bool, k string, v string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) STYLE(k string, v string) *SVGFECONVOLVEMATRIXElement {
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

func (e *SVGFECONVOLVEMATRIXElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFECONVOLVEMATRIXElement) STYLEMap(m map[string]string) *SVGFECONVOLVEMATRIXElement {
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
func (e *SVGFECONVOLVEMATRIXElement) STYLEPairs(pairs ...string) *SVGFECONVOLVEMATRIXElement {
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

func (e *SVGFECONVOLVEMATRIXElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFECONVOLVEMATRIXElement) STYLERemove(keys ...string) *SVGFECONVOLVEMATRIXElement {
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

func (e *SVGFECONVOLVEMATRIXElement) Z_REQ(expression string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfZ_REQ(condition bool, expression string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFECONVOLVEMATRIXElement) Z_REQRemove() *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFECONVOLVEMATRIXElement) Z_TARGET(expression string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfZ_TARGET(condition bool, expression string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFECONVOLVEMATRIXElement) Z_TARGETRemove() *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFECONVOLVEMATRIXElement) Z_REQ_SELECTOR(expression string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFECONVOLVEMATRIXElement) Z_REQ_SELECTORRemove() *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFECONVOLVEMATRIXElement) Z_SWAP(expression string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfZ_SWAP(condition bool, expression string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFECONVOLVEMATRIXElement) Z_SWAPRemove() *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFECONVOLVEMATRIXElement) Z_SWAP_PUSH(expression string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFECONVOLVEMATRIXElement) Z_SWAP_PUSHRemove() *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFECONVOLVEMATRIXElement) Z_TRIGGER(expression string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfZ_TRIGGER(condition bool, expression string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFECONVOLVEMATRIXElement) Z_TRIGGERRemove() *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFECONVOLVEMATRIXElement) Z_REQ_METHOD(c SVGFeConvolveMatrixZReqMethodChoice) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeConvolveMatrixZReqMethodChoice string

const (
	// default GET
	SVGFeConvolveMatrixZReqMethod_empty SVGFeConvolveMatrixZReqMethodChoice = ""
	// GET
	SVGFeConvolveMatrixZReqMethod_get SVGFeConvolveMatrixZReqMethodChoice = "get"
	// POST
	SVGFeConvolveMatrixZReqMethod_post SVGFeConvolveMatrixZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFECONVOLVEMATRIXElement) Z_REQ_METHODRemove(c SVGFeConvolveMatrixZReqMethodChoice) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFECONVOLVEMATRIXElement) Z_REQ_STRATEGY(expression string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFECONVOLVEMATRIXElement) Z_REQ_STRATEGYRemove() *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFECONVOLVEMATRIXElement) Z_REQ_HISTORY(expression string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFECONVOLVEMATRIXElement) Z_REQ_HISTORYRemove() *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFECONVOLVEMATRIXElement) Z_DATA(expression string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfZ_DATA(condition bool, expression string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFECONVOLVEMATRIXElement) Z_DATARemove() *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFECONVOLVEMATRIXElement) Z_JSON(expression string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfZ_JSON(condition bool, expression string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFECONVOLVEMATRIXElement) Z_JSONRemove() *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFECONVOLVEMATRIXElement) Z_REQ_BATCH(expression string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFECONVOLVEMATRIXElement) Z_REQ_BATCHRemove() *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFECONVOLVEMATRIXElement) Z_ACTION(expression string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfZ_ACTION(condition bool, expression string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFECONVOLVEMATRIXElement) Z_ACTIONRemove() *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFECONVOLVEMATRIXElement) Z_REQ_BEFORE(expression string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFECONVOLVEMATRIXElement) Z_REQ_BEFORERemove() *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFECONVOLVEMATRIXElement) Z_REQ_AFTER(expression string) *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECONVOLVEMATRIXElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFECONVOLVEMATRIXElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFECONVOLVEMATRIXElement) Z_REQ_AFTERRemove() *SVGFECONVOLVEMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
