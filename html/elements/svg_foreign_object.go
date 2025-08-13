package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <foreignObject> SVG element allows for inclusion of a foreign XML namespace
// which has its graphical content drawn by a different user agent
// The included foreign graphical content is subject to SVG transformations and
// compositing.
type SVGFOREIGNOBJECTElement struct {
	*Element
}

// Create a new SVGFOREIGNOBJECTElement element.
// This will create a new element with the tag
// "foreignObject" during rendering.
func SVG_FOREIGNOBJECT(children ...ElementRenderer) *SVGFOREIGNOBJECTElement {
	e := NewElement("foreignObject", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFOREIGNOBJECTElement{Element: e}
}

func (e *SVGFOREIGNOBJECTElement) Children(children ...ElementRenderer) *SVGFOREIGNOBJECTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) Attr(name string, value ...string) *SVGFOREIGNOBJECTElement {
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

func (e *SVGFOREIGNOBJECTElement) Attrs(attrs ...string) *SVGFOREIGNOBJECTElement {
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

func (e *SVGFOREIGNOBJECTElement) AttrsMap(attrs map[string]string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) Text(text string) *SVGFOREIGNOBJECTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFOREIGNOBJECTElement) TextF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfText(condition bool, text string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfTextF(condition bool, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) Escaped(text string) *SVGFOREIGNOBJECTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfEscaped(condition bool, text string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) EscapedF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfEscapedF(condition bool, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) CustomData(key, value string) *SVGFOREIGNOBJECTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfCustomData(condition bool, key, value string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) CustomDataF(key, format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) CustomDataRemove(key string) *SVGFOREIGNOBJECTElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The x-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGFOREIGNOBJECTElement) X(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("x", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) XF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.X(fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfX(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.X(s)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfXF(condition bool, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.X(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute X from the element.
func (e *SVGFOREIGNOBJECTElement) XRemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("x")
	return e
}

func (e *SVGFOREIGNOBJECTElement) XRemoveF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.XRemove(fmt.Sprintf(format, args...))
}

// The y-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGFOREIGNOBJECTElement) Y(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("y", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) YF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.Y(fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfY(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Y(s)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfYF(condition bool, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Y(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute Y from the element.
func (e *SVGFOREIGNOBJECTElement) YRemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("y")
	return e
}

func (e *SVGFOREIGNOBJECTElement) YRemoveF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.YRemove(fmt.Sprintf(format, args...))
}

// The width of the rectangular region.
func (e *SVGFOREIGNOBJECTElement) WIDTH(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("width", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) WIDTHF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.WIDTH(fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfWIDTH(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.WIDTH(s)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfWIDTHF(condition bool, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.WIDTH(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute WIDTH from the element.
func (e *SVGFOREIGNOBJECTElement) WIDTHRemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("width")
	return e
}

func (e *SVGFOREIGNOBJECTElement) WIDTHRemoveF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.WIDTHRemove(fmt.Sprintf(format, args...))
}

// The height of the rectangular region.
func (e *SVGFOREIGNOBJECTElement) HEIGHT(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("height", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) HEIGHTF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.HEIGHT(fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfHEIGHT(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.HEIGHT(s)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfHEIGHTF(condition bool, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.HEIGHT(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HEIGHT from the element.
func (e *SVGFOREIGNOBJECTElement) HEIGHTRemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("height")
	return e
}

func (e *SVGFOREIGNOBJECTElement) HEIGHTRemoveF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.HEIGHTRemove(fmt.Sprintf(format, args...))
}

// A space-separated list of required extensions, indicating that the parent SVG
// document must include the specified extensions for this element to be valid.
func (e *SVGFOREIGNOBJECTElement) REQUIRED_EXTENSIONS(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("requiredExtensions", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) REQUIRED_EXTENSIONSF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.REQUIRED_EXTENSIONS(fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfREQUIRED_EXTENSIONS(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.REQUIRED_EXTENSIONS(s)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfREQUIRED_EXTENSIONSF(condition bool, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.REQUIRED_EXTENSIONS(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REQUIRED_EXTENSIONS from the element.
func (e *SVGFOREIGNOBJECTElement) REQUIRED_EXTENSIONSRemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("requiredExtensions")
	return e
}

func (e *SVGFOREIGNOBJECTElement) REQUIRED_EXTENSIONSRemoveF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.REQUIRED_EXTENSIONSRemove(fmt.Sprintf(format, args...))
}

// A space-separated list of required features, indicating that the parent SVG
// document must include support for all of the specified features for this
// element to be valid.
func (e *SVGFOREIGNOBJECTElement) REQUIRED_FEATURES(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("requiredFeatures", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) REQUIRED_FEATURESF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.REQUIRED_FEATURES(fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfREQUIRED_FEATURES(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.REQUIRED_FEATURES(s)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfREQUIRED_FEATURESF(condition bool, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.REQUIRED_FEATURES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REQUIRED_FEATURES from the element.
func (e *SVGFOREIGNOBJECTElement) REQUIRED_FEATURESRemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("requiredFeatures")
	return e
}

func (e *SVGFOREIGNOBJECTElement) REQUIRED_FEATURESRemoveF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.REQUIRED_FEATURESRemove(fmt.Sprintf(format, args...))
}

// A space-separated list of language codes, indicating that the parent SVG
// document must include support for all of the specified languages for this
// element to be valid.
func (e *SVGFOREIGNOBJECTElement) SYSTEM_LANGUAGE(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("systemLanguage", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) SYSTEM_LANGUAGEF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.SYSTEM_LANGUAGE(fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfSYSTEM_LANGUAGE(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.SYSTEM_LANGUAGE(s)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfSYSTEM_LANGUAGEF(condition bool, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.SYSTEM_LANGUAGE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute SYSTEM_LANGUAGE from the element.
func (e *SVGFOREIGNOBJECTElement) SYSTEM_LANGUAGERemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("systemLanguage")
	return e
}

func (e *SVGFOREIGNOBJECTElement) SYSTEM_LANGUAGERemoveF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.SYSTEM_LANGUAGERemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGFOREIGNOBJECTElement) ID(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IDF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfID(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfIDF(condition bool, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFOREIGNOBJECTElement) IDRemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFOREIGNOBJECTElement) IDRemoveF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFOREIGNOBJECTElement) CLASS(s ...string) *SVGFOREIGNOBJECTElement {
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

func (e *SVGFOREIGNOBJECTElement) IfCLASS(condition bool, s ...string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFOREIGNOBJECTElement) CLASSRemove(s ...string) *SVGFOREIGNOBJECTElement {
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
func (e *SVGFOREIGNOBJECTElement) STYLEF(k string, format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfSTYLE(condition bool, k string, v string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) STYLE(k string, v string) *SVGFOREIGNOBJECTElement {
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

func (e *SVGFOREIGNOBJECTElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFOREIGNOBJECTElement) STYLEMap(m map[string]string) *SVGFOREIGNOBJECTElement {
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
func (e *SVGFOREIGNOBJECTElement) STYLEPairs(pairs ...string) *SVGFOREIGNOBJECTElement {
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

func (e *SVGFOREIGNOBJECTElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFOREIGNOBJECTElement) STYLERemove(keys ...string) *SVGFOREIGNOBJECTElement {
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

func (e *SVGFOREIGNOBJECTElement) Z_REQ(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfZ_REQ(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFOREIGNOBJECTElement) Z_REQRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFOREIGNOBJECTElement) Z_TARGET(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfZ_TARGET(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFOREIGNOBJECTElement) Z_TARGETRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFOREIGNOBJECTElement) Z_REQ_SELECTOR(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFOREIGNOBJECTElement) Z_REQ_SELECTORRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFOREIGNOBJECTElement) Z_SWAP(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfZ_SWAP(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFOREIGNOBJECTElement) Z_SWAPRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFOREIGNOBJECTElement) Z_SWAP_PUSH(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFOREIGNOBJECTElement) Z_SWAP_PUSHRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFOREIGNOBJECTElement) Z_TRIGGER(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfZ_TRIGGER(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFOREIGNOBJECTElement) Z_TRIGGERRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFOREIGNOBJECTElement) Z_REQ_METHOD(c SVGForeignObjectZReqMethodChoice) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGForeignObjectZReqMethodChoice string

const (
	// default GET
	SVGForeignObjectZReqMethod_empty SVGForeignObjectZReqMethodChoice = ""
	// GET
	SVGForeignObjectZReqMethod_get SVGForeignObjectZReqMethodChoice = "get"
	// POST
	SVGForeignObjectZReqMethod_post SVGForeignObjectZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFOREIGNOBJECTElement) Z_REQ_METHODRemove(c SVGForeignObjectZReqMethodChoice) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFOREIGNOBJECTElement) Z_REQ_STRATEGY(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFOREIGNOBJECTElement) Z_REQ_STRATEGYRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFOREIGNOBJECTElement) Z_REQ_HISTORY(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFOREIGNOBJECTElement) Z_REQ_HISTORYRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFOREIGNOBJECTElement) Z_DATA(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfZ_DATA(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFOREIGNOBJECTElement) Z_DATARemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFOREIGNOBJECTElement) Z_JSON(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfZ_JSON(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFOREIGNOBJECTElement) Z_JSONRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFOREIGNOBJECTElement) Z_REQ_BATCH(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFOREIGNOBJECTElement) Z_REQ_BATCHRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFOREIGNOBJECTElement) Z_ACTION(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfZ_ACTION(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFOREIGNOBJECTElement) Z_ACTIONRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFOREIGNOBJECTElement) Z_REQ_BEFORE(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFOREIGNOBJECTElement) Z_REQ_BEFORERemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFOREIGNOBJECTElement) Z_REQ_AFTER(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFOREIGNOBJECTElement) Z_REQ_AFTERRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
