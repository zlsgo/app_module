package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feDisplacementMap> SVG filter primitive uses the pixel values from the
// image from in2 to spatially displace the image from in.
type SVGFEDISPLACEMENTMAPElement struct {
	*Element
}

// Create a new SVGFEDISPLACEMENTMAPElement element.
// This will create a new element with the tag
// "feDisplacementMap" during rendering.
func SVG_FEDISPLACEMENTMAP(children ...ElementRenderer) *SVGFEDISPLACEMENTMAPElement {
	e := NewElement("feDisplacementMap", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEDISPLACEMENTMAPElement{Element: e}
}

func (e *SVGFEDISPLACEMENTMAPElement) Children(children ...ElementRenderer) *SVGFEDISPLACEMENTMAPElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) Attr(name string, value ...string) *SVGFEDISPLACEMENTMAPElement {
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

func (e *SVGFEDISPLACEMENTMAPElement) Attrs(attrs ...string) *SVGFEDISPLACEMENTMAPElement {
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

func (e *SVGFEDISPLACEMENTMAPElement) AttrsMap(attrs map[string]string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) Text(text string) *SVGFEDISPLACEMENTMAPElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) TextF(format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEDISPLACEMENTMAPElement) IfText(condition bool, text string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfTextF(condition bool, format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) Escaped(text string) *SVGFEDISPLACEMENTMAPElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfEscaped(condition bool, text string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) EscapedF(format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEDISPLACEMENTMAPElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) CustomData(key, value string) *SVGFEDISPLACEMENTMAPElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfCustomData(condition bool, key, value string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) CustomDataF(key, format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEDISPLACEMENTMAPElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) CustomDataRemove(key string) *SVGFEDISPLACEMENTMAPElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The input for this filter.
func (e *SVGFEDISPLACEMENTMAPElement) IN(s string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) INF(format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	return e.IN(fmt.Sprintf(format, args...))
}

func (e *SVGFEDISPLACEMENTMAPElement) IfIN(condition bool, s string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.IN(s)
	}
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfINF(condition bool, format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.IN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFEDISPLACEMENTMAPElement) INRemove(s string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("in")
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) INRemoveF(format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	return e.INRemove(fmt.Sprintf(format, args...))
}

// The displacement map
// This attribute can take on the same values as the 'in' attribute.
func (e *SVGFEDISPLACEMENTMAPElement) IN_2(s string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("in2", s)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IN_2F(format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	return e.IN_2(fmt.Sprintf(format, args...))
}

func (e *SVGFEDISPLACEMENTMAPElement) IfIN_2(condition bool, s string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.IN_2(s)
	}
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfIN_2F(condition bool, format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.IN_2(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN_2 from the element.
func (e *SVGFEDISPLACEMENTMAPElement) IN_2Remove(s string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("in2")
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IN_2RemoveF(format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	return e.IN_2Remove(fmt.Sprintf(format, args...))
}

// The scale attribute defines the maximum value for the in2 displacement
// A value of 0 disables the effect of the displacement map.
func (e *SVGFEDISPLACEMENTMAPElement) SCALE(f float64) *SVGFEDISPLACEMENTMAPElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("scale", f)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfSCALE(condition bool, f float64) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.SCALE(f)
	}
	return e
}

// The xChannelSelector attribute indicates which color channel from in2 to use to
// displace the pixels in in the horizontal direction.
func (e *SVGFEDISPLACEMENTMAPElement) X_CHANNEL_SELECTOR(c SVGFeDisplacementMapXChannelSelectorChoice) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("xChannelSelector", string(c))
	return e
}

type SVGFeDisplacementMapXChannelSelectorChoice string

const (
	// The red channel of in2 is used to displace the x coordinate of each pixel.
	SVGFeDisplacementMapXChannelSelector_R SVGFeDisplacementMapXChannelSelectorChoice = "R"
	// The green channel of in2 is used to displace the x coordinate of each pixel.
	SVGFeDisplacementMapXChannelSelector_G SVGFeDisplacementMapXChannelSelectorChoice = "G"
	// The blue channel of in2 is used to displace the x coordinate of each pixel.
	SVGFeDisplacementMapXChannelSelector_B SVGFeDisplacementMapXChannelSelectorChoice = "B"
	// The alpha channel of in2 is used to displace the x coordinate of each pixel.
	SVGFeDisplacementMapXChannelSelector_A SVGFeDisplacementMapXChannelSelectorChoice = "A"
)

// Remove the attribute X_CHANNEL_SELECTOR from the element.
func (e *SVGFEDISPLACEMENTMAPElement) X_CHANNEL_SELECTORRemove(c SVGFeDisplacementMapXChannelSelectorChoice) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("xChannelSelector")
	return e
}

// The yChannelSelector attribute indicates which color channel from in2 to use to
// displace the pixels in in the vertical direction.
func (e *SVGFEDISPLACEMENTMAPElement) Y_CHANNEL_SELECTOR(c SVGFeDisplacementMapYChannelSelectorChoice) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("yChannelSelector", string(c))
	return e
}

type SVGFeDisplacementMapYChannelSelectorChoice string

const (
	// The red channel of in2 is used to displace the y coordinate of each pixel.
	SVGFeDisplacementMapYChannelSelector_R SVGFeDisplacementMapYChannelSelectorChoice = "R"
	// The green channel of in2 is used to displace the y coordinate of each pixel.
	SVGFeDisplacementMapYChannelSelector_G SVGFeDisplacementMapYChannelSelectorChoice = "G"
	// The blue channel of in2 is used to displace the y coordinate of each pixel.
	SVGFeDisplacementMapYChannelSelector_B SVGFeDisplacementMapYChannelSelectorChoice = "B"
	// The alpha channel of in2 is used to displace the y coordinate of each pixel.
	SVGFeDisplacementMapYChannelSelector_A SVGFeDisplacementMapYChannelSelectorChoice = "A"
)

// Remove the attribute Y_CHANNEL_SELECTOR from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Y_CHANNEL_SELECTORRemove(c SVGFeDisplacementMapYChannelSelectorChoice) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("yChannelSelector")
	return e
}

// Specifies a unique id for an element
func (e *SVGFEDISPLACEMENTMAPElement) ID(s string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IDF(format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEDISPLACEMENTMAPElement) IfID(condition bool, s string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfIDF(condition bool, format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEDISPLACEMENTMAPElement) IDRemove(s string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IDRemoveF(format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEDISPLACEMENTMAPElement) CLASS(s ...string) *SVGFEDISPLACEMENTMAPElement {
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

func (e *SVGFEDISPLACEMENTMAPElement) IfCLASS(condition bool, s ...string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEDISPLACEMENTMAPElement) CLASSRemove(s ...string) *SVGFEDISPLACEMENTMAPElement {
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
func (e *SVGFEDISPLACEMENTMAPElement) STYLEF(k string, format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEDISPLACEMENTMAPElement) IfSTYLE(condition bool, k string, v string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) STYLE(k string, v string) *SVGFEDISPLACEMENTMAPElement {
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

func (e *SVGFEDISPLACEMENTMAPElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEDISPLACEMENTMAPElement) STYLEMap(m map[string]string) *SVGFEDISPLACEMENTMAPElement {
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
func (e *SVGFEDISPLACEMENTMAPElement) STYLEPairs(pairs ...string) *SVGFEDISPLACEMENTMAPElement {
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

func (e *SVGFEDISPLACEMENTMAPElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEDISPLACEMENTMAPElement) STYLERemove(keys ...string) *SVGFEDISPLACEMENTMAPElement {
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

func (e *SVGFEDISPLACEMENTMAPElement) Z_REQ(expression string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfZ_REQ(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Z_REQRemove() *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFEDISPLACEMENTMAPElement) Z_TARGET(expression string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfZ_TARGET(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Z_TARGETRemove() *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFEDISPLACEMENTMAPElement) Z_REQ_SELECTOR(expression string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Z_REQ_SELECTORRemove() *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFEDISPLACEMENTMAPElement) Z_SWAP(expression string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfZ_SWAP(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Z_SWAPRemove() *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFEDISPLACEMENTMAPElement) Z_SWAP_PUSH(expression string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Z_SWAP_PUSHRemove() *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFEDISPLACEMENTMAPElement) Z_TRIGGER(expression string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfZ_TRIGGER(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Z_TRIGGERRemove() *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFEDISPLACEMENTMAPElement) Z_REQ_METHOD(c SVGFeDisplacementMapZReqMethodChoice) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeDisplacementMapZReqMethodChoice string

const (
	// default GET
	SVGFeDisplacementMapZReqMethod_empty SVGFeDisplacementMapZReqMethodChoice = ""
	// GET
	SVGFeDisplacementMapZReqMethod_get SVGFeDisplacementMapZReqMethodChoice = "get"
	// POST
	SVGFeDisplacementMapZReqMethod_post SVGFeDisplacementMapZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Z_REQ_METHODRemove(c SVGFeDisplacementMapZReqMethodChoice) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFEDISPLACEMENTMAPElement) Z_REQ_STRATEGY(expression string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Z_REQ_STRATEGYRemove() *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFEDISPLACEMENTMAPElement) Z_REQ_HISTORY(expression string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Z_REQ_HISTORYRemove() *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFEDISPLACEMENTMAPElement) Z_DATA(expression string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfZ_DATA(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Z_DATARemove() *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFEDISPLACEMENTMAPElement) Z_JSON(expression string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfZ_JSON(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Z_JSONRemove() *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFEDISPLACEMENTMAPElement) Z_REQ_BATCH(expression string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Z_REQ_BATCHRemove() *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFEDISPLACEMENTMAPElement) Z_ACTION(expression string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfZ_ACTION(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Z_ACTIONRemove() *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFEDISPLACEMENTMAPElement) Z_REQ_BEFORE(expression string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Z_REQ_BEFORERemove() *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFEDISPLACEMENTMAPElement) Z_REQ_AFTER(expression string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFEDISPLACEMENTMAPElement) Z_REQ_AFTERRemove() *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
