package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// This element is used to include comments or annotations within a MathML
// expression
// It can be used to provide additional information about the expression, or to
// include comments for the author of the expression.
type MathMLANNOTATION_XMLElement struct {
	*Element
}

// Create a new MathMLANNOTATION_XMLElement element.
// This will create a new element with the tag
// "annotation-xml" during rendering.
func MathML_ANNOTATION_XML(children ...ElementRenderer) *MathMLANNOTATION_XMLElement {
	e := NewElement("annotation-xml", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLANNOTATION_XMLElement{Element: e}
}

func (e *MathMLANNOTATION_XMLElement) Children(children ...ElementRenderer) *MathMLANNOTATION_XMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) Attr(name string, value ...string) *MathMLANNOTATION_XMLElement {
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

func (e *MathMLANNOTATION_XMLElement) Attrs(attrs ...string) *MathMLANNOTATION_XMLElement {
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

func (e *MathMLANNOTATION_XMLElement) AttrsMap(attrs map[string]string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) Text(text string) *MathMLANNOTATION_XMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLANNOTATION_XMLElement) TextF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATION_XMLElement) IfText(condition bool, text string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfTextF(condition bool, format string, args ...any) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) Escaped(text string) *MathMLANNOTATION_XMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfEscaped(condition bool, text string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) EscapedF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATION_XMLElement) IfEscapedF(condition bool, format string, args ...any) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) CustomData(key, value string) *MathMLANNOTATION_XMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfCustomData(condition bool, key, value string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) CustomDataF(key, format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATION_XMLElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLANNOTATION_XMLElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) CustomDataRemove(key string) *MathMLANNOTATION_XMLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// This attribute specifies the encoding used for the text content of the element
// Possible values are text/plain, text/html, and application/x-tex.
func (e *MathMLANNOTATION_XMLElement) ENCODING(c MathMLAnnotationXmlEncodingChoice) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("encoding", string(c))
	return e
}

type MathMLAnnotationXmlEncodingChoice string

const (
	MathMLAnnotationXmlEncoding_text_plain MathMLAnnotationXmlEncodingChoice = "text/plain"

	MathMLAnnotationXmlEncoding_text_html MathMLAnnotationXmlEncodingChoice = "text/html"

	MathMLAnnotationXmlEncoding_application_x_tex MathMLAnnotationXmlEncodingChoice = "application/x-tex"
)

// Remove the attribute ENCODING from the element.
func (e *MathMLANNOTATION_XMLElement) ENCODINGRemove(c MathMLAnnotationXmlEncodingChoice) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("encoding")
	return e
}

// This attribute specifies the name of the annotation.
func (e *MathMLANNOTATION_XMLElement) NAME(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("name", s)
	return e
}

func (e *MathMLANNOTATION_XMLElement) NAMEF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.NAME(fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATION_XMLElement) IfNAME(condition bool, s string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.NAME(s)
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfNAMEF(condition bool, format string, args ...any) *MathMLANNOTATION_XMLElement {
	if condition {
		e.NAME(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NAME from the element.
func (e *MathMLANNOTATION_XMLElement) NAMERemove(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("name")
	return e
}

func (e *MathMLANNOTATION_XMLElement) NAMERemoveF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.NAMERemove(fmt.Sprintf(format, args...))
}

// Assigns a class name or set of class names to an element
// You may assign the same class name or names to any number of elements
// If you specify multiple class names, they must be separated by whitespace
// characters.
func (e *MathMLANNOTATION_XMLElement) CLASS(s ...string) *MathMLANNOTATION_XMLElement {
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

func (e *MathMLANNOTATION_XMLElement) IfCLASS(condition bool, s ...string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLANNOTATION_XMLElement) CLASSRemove(s ...string) *MathMLANNOTATION_XMLElement {
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
func (e *MathMLANNOTATION_XMLElement) DIR(c MathMLAnnotationXmlDirChoice) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLAnnotationXmlDirChoice string

const (
	// left-to-right
	MathMLAnnotationXmlDir_ltr MathMLAnnotationXmlDirChoice = "ltr"
	// right-to-left
	MathMLAnnotationXmlDir_rtl MathMLAnnotationXmlDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLANNOTATION_XMLElement) DIRRemove(c MathMLAnnotationXmlDirChoice) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLANNOTATION_XMLElement) DISPLAYSTYLE(c MathMLAnnotationXmlDisplaystyleChoice) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLAnnotationXmlDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLAnnotationXmlDisplaystyle_true MathMLAnnotationXmlDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLAnnotationXmlDisplaystyle_false MathMLAnnotationXmlDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLANNOTATION_XMLElement) DISPLAYSTYLERemove(c MathMLAnnotationXmlDisplaystyleChoice) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLANNOTATION_XMLElement) ID(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IDF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATION_XMLElement) IfID(condition bool, s string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfIDF(condition bool, format string, args ...any) *MathMLANNOTATION_XMLElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLANNOTATION_XMLElement) IDRemove(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *MathMLANNOTATION_XMLElement) IDRemoveF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLANNOTATION_XMLElement) MATHBACKGROUND(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLANNOTATION_XMLElement) MATHBACKGROUNDF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.MATHBACKGROUND(fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATION_XMLElement) IfMATHBACKGROUND(condition bool, s string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfMATHBACKGROUNDF(condition bool, format string, args ...any) *MathMLANNOTATION_XMLElement {
	if condition {
		e.MATHBACKGROUND(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLANNOTATION_XMLElement) MATHBACKGROUNDRemove(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathbackground")
	return e
}

func (e *MathMLANNOTATION_XMLElement) MATHBACKGROUNDRemoveF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.MATHBACKGROUNDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLANNOTATION_XMLElement) MATHCOLOR(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLANNOTATION_XMLElement) MATHCOLORF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.MATHCOLOR(fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATION_XMLElement) IfMATHCOLOR(condition bool, s string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfMATHCOLORF(condition bool, format string, args ...any) *MathMLANNOTATION_XMLElement {
	if condition {
		e.MATHCOLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLANNOTATION_XMLElement) MATHCOLORRemove(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathcolor")
	return e
}

func (e *MathMLANNOTATION_XMLElement) MATHCOLORRemoveF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.MATHCOLORRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLANNOTATION_XMLElement) MATHSIZE_STR(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLANNOTATION_XMLElement) MATHSIZE_STRF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.MATHSIZE_STR(fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATION_XMLElement) IfMATHSIZE_STR(condition bool, s string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfMATHSIZE_STRF(condition bool, format string, args ...any) *MathMLANNOTATION_XMLElement {
	if condition {
		e.MATHSIZE_STR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLANNOTATION_XMLElement) MATHSIZE_STRRemove(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathsize")
	return e
}

func (e *MathMLANNOTATION_XMLElement) MATHSIZE_STRRemoveF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.MATHSIZE_STRRemove(fmt.Sprintf(format, args...))
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLANNOTATION_XMLElement) NONCE(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLANNOTATION_XMLElement) NONCEF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATION_XMLElement) IfNONCE(condition bool, s string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfNONCEF(condition bool, format string, args ...any) *MathMLANNOTATION_XMLElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLANNOTATION_XMLElement) NONCERemove(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("nonce")
	return e
}

func (e *MathMLANNOTATION_XMLElement) NONCERemoveF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLANNOTATION_XMLElement) SCRIPTLEVEL(i int) *MathMLANNOTATION_XMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLANNOTATION_XMLElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLANNOTATION_XMLElement) SCRIPTLEVELRemove(i int) *MathMLANNOTATION_XMLElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLANNOTATION_XMLElement) STYLEF(k string, format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATION_XMLElement) IfSTYLE(condition bool, k string, v string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) STYLE(k string, v string) *MathMLANNOTATION_XMLElement {
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

func (e *MathMLANNOTATION_XMLElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLANNOTATION_XMLElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLANNOTATION_XMLElement) STYLEMap(m map[string]string) *MathMLANNOTATION_XMLElement {
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
func (e *MathMLANNOTATION_XMLElement) STYLEPairs(pairs ...string) *MathMLANNOTATION_XMLElement {
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

func (e *MathMLANNOTATION_XMLElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLANNOTATION_XMLElement) STYLERemove(keys ...string) *MathMLANNOTATION_XMLElement {
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
func (e *MathMLANNOTATION_XMLElement) TABINDEX(i int) *MathMLANNOTATION_XMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfTABINDEX(condition bool, i int) *MathMLANNOTATION_XMLElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLANNOTATION_XMLElement) TABINDEXRemove(i int) *MathMLANNOTATION_XMLElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("tabindex")
	return e
}

// Make a request for an HTML

func (e *MathMLANNOTATION_XMLElement) Z_REQ(expression string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfZ_REQ(condition bool, expression string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *MathMLANNOTATION_XMLElement) Z_REQRemove() *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *MathMLANNOTATION_XMLElement) Z_TARGET(expression string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfZ_TARGET(condition bool, expression string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *MathMLANNOTATION_XMLElement) Z_TARGETRemove() *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *MathMLANNOTATION_XMLElement) Z_REQ_SELECTOR(expression string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfZ_REQ_SELECTOR(condition bool, expression string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *MathMLANNOTATION_XMLElement) Z_REQ_SELECTORRemove() *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *MathMLANNOTATION_XMLElement) Z_SWAP(expression string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfZ_SWAP(condition bool, expression string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *MathMLANNOTATION_XMLElement) Z_SWAPRemove() *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *MathMLANNOTATION_XMLElement) Z_SWAP_PUSH(expression string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfZ_SWAP_PUSH(condition bool, expression string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *MathMLANNOTATION_XMLElement) Z_SWAP_PUSHRemove() *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *MathMLANNOTATION_XMLElement) Z_TRIGGER(expression string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfZ_TRIGGER(condition bool, expression string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *MathMLANNOTATION_XMLElement) Z_TRIGGERRemove() *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *MathMLANNOTATION_XMLElement) Z_REQ_METHOD(c MathMLAnnotationXmlZReqMethodChoice) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type MathMLAnnotationXmlZReqMethodChoice string

const (
	// default GET
	MathMLAnnotationXmlZReqMethod_empty MathMLAnnotationXmlZReqMethodChoice = ""
	// GET
	MathMLAnnotationXmlZReqMethod_get MathMLAnnotationXmlZReqMethodChoice = "get"
	// POST
	MathMLAnnotationXmlZReqMethod_post MathMLAnnotationXmlZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *MathMLANNOTATION_XMLElement) Z_REQ_METHODRemove(c MathMLAnnotationXmlZReqMethodChoice) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *MathMLANNOTATION_XMLElement) Z_REQ_STRATEGY(expression string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfZ_REQ_STRATEGY(condition bool, expression string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *MathMLANNOTATION_XMLElement) Z_REQ_STRATEGYRemove() *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *MathMLANNOTATION_XMLElement) Z_REQ_HISTORY(expression string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfZ_REQ_HISTORY(condition bool, expression string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *MathMLANNOTATION_XMLElement) Z_REQ_HISTORYRemove() *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *MathMLANNOTATION_XMLElement) Z_DATA(expression string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfZ_DATA(condition bool, expression string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *MathMLANNOTATION_XMLElement) Z_DATARemove() *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *MathMLANNOTATION_XMLElement) Z_JSON(expression string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfZ_JSON(condition bool, expression string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *MathMLANNOTATION_XMLElement) Z_JSONRemove() *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *MathMLANNOTATION_XMLElement) Z_REQ_BATCH(expression string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfZ_REQ_BATCH(condition bool, expression string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *MathMLANNOTATION_XMLElement) Z_REQ_BATCHRemove() *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *MathMLANNOTATION_XMLElement) Z_ACTION(expression string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfZ_ACTION(condition bool, expression string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *MathMLANNOTATION_XMLElement) Z_ACTIONRemove() *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *MathMLANNOTATION_XMLElement) Z_REQ_BEFORE(expression string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfZ_REQ_BEFORE(condition bool, expression string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *MathMLANNOTATION_XMLElement) Z_REQ_BEFORERemove() *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *MathMLANNOTATION_XMLElement) Z_REQ_AFTER(expression string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfZ_REQ_AFTER(condition bool, expression string) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *MathMLANNOTATION_XMLElement) Z_REQ_AFTERRemove() *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
