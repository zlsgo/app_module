package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// This element is the root element of a MathML expression
// It is used to identify the document as a MathML document, and to specify the
// namespaces used in the document.
type MathMLMATHElement struct {
	*Element
}

// Create a new MathMLMATHElement element.
// This will create a new element with the tag
// "math" during rendering.
func MathML_MATH(children ...ElementRenderer) *MathMLMATHElement {
	e := NewElement("math", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMATHElement{Element: e}
}

func (e *MathMLMATHElement) Children(children ...ElementRenderer) *MathMLMATHElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMATHElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMATHElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMATHElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMATHElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMATHElement) Attr(name string, value ...string) *MathMLMATHElement {
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

func (e *MathMLMATHElement) Attrs(attrs ...string) *MathMLMATHElement {
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

func (e *MathMLMATHElement) AttrsMap(attrs map[string]string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMATHElement) Text(text string) *MathMLMATHElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMATHElement) TextF(format string, args ...any) *MathMLMATHElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMATHElement) IfText(condition bool, text string) *MathMLMATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMATHElement) IfTextF(condition bool, format string, args ...any) *MathMLMATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMATHElement) Escaped(text string) *MathMLMATHElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMATHElement) IfEscaped(condition bool, text string) *MathMLMATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMATHElement) EscapedF(format string, args ...any) *MathMLMATHElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMATHElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMATHElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMATHElement) CustomData(key, value string) *MathMLMATHElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMATHElement) IfCustomData(condition bool, key, value string) *MathMLMATHElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMATHElement) CustomDataF(key, format string, args ...any) *MathMLMATHElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMATHElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMATHElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMATHElement) CustomDataRemove(key string) *MathMLMATHElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// This attribute specifies the default namespace for elements and attributes in
// the document
// Possible values are http://www.w3.org/1998/Math/MathML and
// http://www.w3.org/1999/xhtml.
func (e *MathMLMATHElement) XMLNS(c MathMLMathXmlnsChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("xmlns", string(c))
	return e
}

type MathMLMathXmlnsChoice string

const (
	MathMLMathXmlns_http___www_w3_org_1998_Math_MathML MathMLMathXmlnsChoice = "http://www.w3.org/1998/Math/MathML"

	MathMLMathXmlns_http___www_w3_org_1999_xhtml MathMLMathXmlnsChoice = "http://www.w3.org/1999/xhtml"
)

// Remove the attribute XMLNS from the element.
func (e *MathMLMATHElement) XMLNSRemove(c MathMLMathXmlnsChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("xmlns")
	return e
}

// This attribute specifies the namespace for elements and attributes in the
// document whose names start with the letter m
// Possible values are http://www.w3.org/1998/Math/MathML and
// http://www.w3.org/1999/xhtml.
func (e *MathMLMATHElement) XMLNS_M(c MathMLMathXmlnsmChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("xmlns:m", string(c))
	return e
}

type MathMLMathXmlnsmChoice string

const (
	MathMLMathXmlnsm_http___www_w3_org_1998_Math_MathML MathMLMathXmlnsmChoice = "http://www.w3.org/1998/Math/MathML"

	MathMLMathXmlnsm_http___www_w3_org_1999_xhtml MathMLMathXmlnsmChoice = "http://www.w3.org/1999/xhtml"
)

// Remove the attribute XMLNS_M from the element.
func (e *MathMLMATHElement) XMLNS_MRemove(c MathMLMathXmlnsmChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("xmlns:m")
	return e
}

// This attribute specifies the namespace for elements and attributes in the
// document whose names start with the letters xlink
// Possible values are http://www.w3.org/1999/xlink and
// http://www.w3.org/1999/xhtml.
func (e *MathMLMATHElement) XMLNS_XLINK(c MathMLMathXmlnsxlinkChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("xmlns:xlink", string(c))
	return e
}

type MathMLMathXmlnsxlinkChoice string

const (
	MathMLMathXmlnsxlink_http___www_w3_org_1999_xlink MathMLMathXmlnsxlinkChoice = "http://www.w3.org/1999/xlink"

	MathMLMathXmlnsxlink_http___www_w3_org_1999_xhtml MathMLMathXmlnsxlinkChoice = "http://www.w3.org/1999/xhtml"
)

// Remove the attribute XMLNS_XLINK from the element.
func (e *MathMLMATHElement) XMLNS_XLINKRemove(c MathMLMathXmlnsxlinkChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("xmlns:xlink")
	return e
}

// This attribute specifies the namespace for elements and attributes in the
// document whose names start with the letters xml
// Possible values are http://www.w3.org/XML/1998/namespace and
// http://www.w3.org/1999/xhtml.
func (e *MathMLMATHElement) XMLNS_XML(c MathMLMathXmlnsxmlChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("xmlns:xml", string(c))
	return e
}

type MathMLMathXmlnsxmlChoice string

const (
	MathMLMathXmlnsxml_http___www_w3_org_XML_1998_namespace MathMLMathXmlnsxmlChoice = "http://www.w3.org/XML/1998/namespace"

	MathMLMathXmlnsxml_http___www_w3_org_1999_xhtml MathMLMathXmlnsxmlChoice = "http://www.w3.org/1999/xhtml"
)

// Remove the attribute XMLNS_XML from the element.
func (e *MathMLMATHElement) XMLNS_XMLRemove(c MathMLMathXmlnsxmlChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("xmlns:xml")
	return e
}

// Assigns a class name or set of class names to an element
// You may assign the same class name or names to any number of elements
// If you specify multiple class names, they must be separated by whitespace
// characters.
func (e *MathMLMATHElement) CLASS(s ...string) *MathMLMATHElement {
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

func (e *MathMLMATHElement) IfCLASS(condition bool, s ...string) *MathMLMATHElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMATHElement) CLASSRemove(s ...string) *MathMLMATHElement {
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

// This attribute specifies the text directionality of the element, merely
// indicating what direction the text flows when surrounded by text with inherent
// directionality (such as Arabic or Hebrew)
// Possible values are ltr (left-to-right) and rtl (right-to-left).
func (e *MathMLMATHElement) DIR(c MathMLMathDirChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMathDirChoice string

const (
	// left-to-right
	MathMLMathDir_ltr MathMLMathDirChoice = "ltr"
	// right-to-left
	MathMLMathDir_rtl MathMLMathDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMATHElement) DIRRemove(c MathMLMathDirChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMATHElement) DISPLAYSTYLE(c MathMLMathDisplaystyleChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMathDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMathDisplaystyle_true MathMLMathDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMathDisplaystyle_false MathMLMathDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMATHElement) DISPLAYSTYLERemove(c MathMLMathDisplaystyleChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMATHElement) ID(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMATHElement) IDF(format string, args ...any) *MathMLMATHElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *MathMLMATHElement) IfID(condition bool, s string) *MathMLMATHElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *MathMLMATHElement) IfIDF(condition bool, format string, args ...any) *MathMLMATHElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMATHElement) IDRemove(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *MathMLMATHElement) IDRemoveF(format string, args ...any) *MathMLMATHElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMATHElement) MATHBACKGROUND(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMATHElement) MATHBACKGROUNDF(format string, args ...any) *MathMLMATHElement {
	return e.MATHBACKGROUND(fmt.Sprintf(format, args...))
}

func (e *MathMLMATHElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMATHElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

func (e *MathMLMATHElement) IfMATHBACKGROUNDF(condition bool, format string, args ...any) *MathMLMATHElement {
	if condition {
		e.MATHBACKGROUND(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMATHElement) MATHBACKGROUNDRemove(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathbackground")
	return e
}

func (e *MathMLMATHElement) MATHBACKGROUNDRemoveF(format string, args ...any) *MathMLMATHElement {
	return e.MATHBACKGROUNDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMATHElement) MATHCOLOR(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMATHElement) MATHCOLORF(format string, args ...any) *MathMLMATHElement {
	return e.MATHCOLOR(fmt.Sprintf(format, args...))
}

func (e *MathMLMATHElement) IfMATHCOLOR(condition bool, s string) *MathMLMATHElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

func (e *MathMLMATHElement) IfMATHCOLORF(condition bool, format string, args ...any) *MathMLMATHElement {
	if condition {
		e.MATHCOLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMATHElement) MATHCOLORRemove(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathcolor")
	return e
}

func (e *MathMLMATHElement) MATHCOLORRemoveF(format string, args ...any) *MathMLMATHElement {
	return e.MATHCOLORRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMATHElement) MATHSIZE_STR(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMATHElement) MATHSIZE_STRF(format string, args ...any) *MathMLMATHElement {
	return e.MATHSIZE_STR(fmt.Sprintf(format, args...))
}

func (e *MathMLMATHElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMATHElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

func (e *MathMLMATHElement) IfMATHSIZE_STRF(condition bool, format string, args ...any) *MathMLMATHElement {
	if condition {
		e.MATHSIZE_STR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMATHElement) MATHSIZE_STRRemove(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathsize")
	return e
}

func (e *MathMLMATHElement) MATHSIZE_STRRemoveF(format string, args ...any) *MathMLMATHElement {
	return e.MATHSIZE_STRRemove(fmt.Sprintf(format, args...))
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLMATHElement) NONCE(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMATHElement) NONCEF(format string, args ...any) *MathMLMATHElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *MathMLMATHElement) IfNONCE(condition bool, s string) *MathMLMATHElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *MathMLMATHElement) IfNONCEF(condition bool, format string, args ...any) *MathMLMATHElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMATHElement) NONCERemove(s string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("nonce")
	return e
}

func (e *MathMLMATHElement) NONCERemoveF(format string, args ...any) *MathMLMATHElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMATHElement) SCRIPTLEVEL(i int) *MathMLMATHElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMATHElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMATHElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMATHElement) SCRIPTLEVELRemove(i int) *MathMLMATHElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMATHElement) STYLEF(k string, format string, args ...any) *MathMLMATHElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMATHElement) IfSTYLE(condition bool, k string, v string) *MathMLMATHElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMATHElement) STYLE(k string, v string) *MathMLMATHElement {
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

func (e *MathMLMATHElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMATHElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMATHElement) STYLEMap(m map[string]string) *MathMLMATHElement {
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
func (e *MathMLMATHElement) STYLEPairs(pairs ...string) *MathMLMATHElement {
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

func (e *MathMLMATHElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMATHElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMATHElement) STYLERemove(keys ...string) *MathMLMATHElement {
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

// This attribute specifies the position of the current element in the tabbing
// order for the current document
// This value must be a number between 0 and 32767
// User agents should ignore leading zeros.
func (e *MathMLMATHElement) TABINDEX(i int) *MathMLMATHElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMATHElement) IfTABINDEX(condition bool, i int) *MathMLMATHElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMATHElement) TABINDEXRemove(i int) *MathMLMATHElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("tabindex")
	return e
}

// Make a request for an HTML

func (e *MathMLMATHElement) Z_REQ(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfZ_REQ(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *MathMLMATHElement) Z_REQRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *MathMLMATHElement) Z_TARGET(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfZ_TARGET(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *MathMLMATHElement) Z_TARGETRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *MathMLMATHElement) Z_REQ_SELECTOR(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfZ_REQ_SELECTOR(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *MathMLMATHElement) Z_REQ_SELECTORRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *MathMLMATHElement) Z_SWAP(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfZ_SWAP(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *MathMLMATHElement) Z_SWAPRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *MathMLMATHElement) Z_SWAP_PUSH(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfZ_SWAP_PUSH(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *MathMLMATHElement) Z_SWAP_PUSHRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *MathMLMATHElement) Z_TRIGGER(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfZ_TRIGGER(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *MathMLMATHElement) Z_TRIGGERRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *MathMLMATHElement) Z_REQ_METHOD(c MathMLMathZReqMethodChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type MathMLMathZReqMethodChoice string

const (
	// default GET
	MathMLMathZReqMethod_empty MathMLMathZReqMethodChoice = ""
	// GET
	MathMLMathZReqMethod_get MathMLMathZReqMethodChoice = "get"
	// POST
	MathMLMathZReqMethod_post MathMLMathZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *MathMLMATHElement) Z_REQ_METHODRemove(c MathMLMathZReqMethodChoice) *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *MathMLMATHElement) Z_REQ_STRATEGY(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfZ_REQ_STRATEGY(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *MathMLMATHElement) Z_REQ_STRATEGYRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *MathMLMATHElement) Z_REQ_HISTORY(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfZ_REQ_HISTORY(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *MathMLMATHElement) Z_REQ_HISTORYRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *MathMLMATHElement) Z_DATA(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfZ_DATA(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *MathMLMATHElement) Z_DATARemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *MathMLMATHElement) Z_JSON(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfZ_JSON(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *MathMLMATHElement) Z_JSONRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *MathMLMATHElement) Z_REQ_BATCH(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfZ_REQ_BATCH(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *MathMLMATHElement) Z_REQ_BATCHRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *MathMLMATHElement) Z_ACTION(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfZ_ACTION(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *MathMLMATHElement) Z_ACTIONRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *MathMLMATHElement) Z_REQ_BEFORE(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfZ_REQ_BEFORE(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *MathMLMATHElement) Z_REQ_BEFORERemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *MathMLMATHElement) Z_REQ_AFTER(expression string) *MathMLMATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMATHElement) IfZ_REQ_AFTER(condition bool, expression string) *MathMLMATHElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *MathMLMATHElement) Z_REQ_AFTERRemove() *MathMLMATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
