package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// This element is used to display a base expression with multiple subscripts and
// superscripts.
type MathMLMMULTISCRIPTSElement struct {
	*Element
}

// Create a new MathMLMMULTISCRIPTSElement element.
// This will create a new element with the tag
// "mmultiscripts" during rendering.
func MathML_MMULTISCRIPTS(children ...ElementRenderer) *MathMLMMULTISCRIPTSElement {
	e := NewElement("mmultiscripts", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMMULTISCRIPTSElement{Element: e}
}

func (e *MathMLMMULTISCRIPTSElement) Children(children ...ElementRenderer) *MathMLMMULTISCRIPTSElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMMULTISCRIPTSElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMMULTISCRIPTSElement) Attr(name string, value ...string) *MathMLMMULTISCRIPTSElement {
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

func (e *MathMLMMULTISCRIPTSElement) Attrs(attrs ...string) *MathMLMMULTISCRIPTSElement {
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

func (e *MathMLMMULTISCRIPTSElement) AttrsMap(attrs map[string]string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMMULTISCRIPTSElement) Text(text string) *MathMLMMULTISCRIPTSElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMMULTISCRIPTSElement) TextF(format string, args ...any) *MathMLMMULTISCRIPTSElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMMULTISCRIPTSElement) IfText(condition bool, text string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfTextF(condition bool, format string, args ...any) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMMULTISCRIPTSElement) Escaped(text string) *MathMLMMULTISCRIPTSElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfEscaped(condition bool, text string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMMULTISCRIPTSElement) EscapedF(format string, args ...any) *MathMLMMULTISCRIPTSElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMMULTISCRIPTSElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMMULTISCRIPTSElement) CustomData(key, value string) *MathMLMMULTISCRIPTSElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfCustomData(condition bool, key, value string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMMULTISCRIPTSElement) CustomDataF(key, format string, args ...any) *MathMLMMULTISCRIPTSElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMMULTISCRIPTSElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMMULTISCRIPTSElement) CustomDataRemove(key string) *MathMLMMULTISCRIPTSElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// Assigns a class name or set of class names to an element
// You may assign the same class name or names to any number of elements
// If you specify multiple class names, they must be separated by whitespace
// characters.
func (e *MathMLMMULTISCRIPTSElement) CLASS(s ...string) *MathMLMMULTISCRIPTSElement {
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

func (e *MathMLMMULTISCRIPTSElement) IfCLASS(condition bool, s ...string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMMULTISCRIPTSElement) CLASSRemove(s ...string) *MathMLMMULTISCRIPTSElement {
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
func (e *MathMLMMULTISCRIPTSElement) DIR(c MathMLMmultiscriptsDirChoice) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMmultiscriptsDirChoice string

const (
	// left-to-right
	MathMLMmultiscriptsDir_ltr MathMLMmultiscriptsDirChoice = "ltr"
	// right-to-left
	MathMLMmultiscriptsDir_rtl MathMLMmultiscriptsDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMMULTISCRIPTSElement) DIRRemove(c MathMLMmultiscriptsDirChoice) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMMULTISCRIPTSElement) DISPLAYSTYLE(c MathMLMmultiscriptsDisplaystyleChoice) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMmultiscriptsDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMmultiscriptsDisplaystyle_true MathMLMmultiscriptsDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMmultiscriptsDisplaystyle_false MathMLMmultiscriptsDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMMULTISCRIPTSElement) DISPLAYSTYLERemove(c MathMLMmultiscriptsDisplaystyleChoice) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMMULTISCRIPTSElement) ID(s string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IDF(format string, args ...any) *MathMLMMULTISCRIPTSElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *MathMLMMULTISCRIPTSElement) IfID(condition bool, s string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfIDF(condition bool, format string, args ...any) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMMULTISCRIPTSElement) IDRemove(s string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IDRemoveF(format string, args ...any) *MathMLMMULTISCRIPTSElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMMULTISCRIPTSElement) MATHBACKGROUND(s string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) MATHBACKGROUNDF(format string, args ...any) *MathMLMMULTISCRIPTSElement {
	return e.MATHBACKGROUND(fmt.Sprintf(format, args...))
}

func (e *MathMLMMULTISCRIPTSElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfMATHBACKGROUNDF(condition bool, format string, args ...any) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.MATHBACKGROUND(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMMULTISCRIPTSElement) MATHBACKGROUNDRemove(s string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathbackground")
	return e
}

func (e *MathMLMMULTISCRIPTSElement) MATHBACKGROUNDRemoveF(format string, args ...any) *MathMLMMULTISCRIPTSElement {
	return e.MATHBACKGROUNDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMMULTISCRIPTSElement) MATHCOLOR(s string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) MATHCOLORF(format string, args ...any) *MathMLMMULTISCRIPTSElement {
	return e.MATHCOLOR(fmt.Sprintf(format, args...))
}

func (e *MathMLMMULTISCRIPTSElement) IfMATHCOLOR(condition bool, s string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfMATHCOLORF(condition bool, format string, args ...any) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.MATHCOLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMMULTISCRIPTSElement) MATHCOLORRemove(s string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathcolor")
	return e
}

func (e *MathMLMMULTISCRIPTSElement) MATHCOLORRemoveF(format string, args ...any) *MathMLMMULTISCRIPTSElement {
	return e.MATHCOLORRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMMULTISCRIPTSElement) MATHSIZE_STR(s string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) MATHSIZE_STRF(format string, args ...any) *MathMLMMULTISCRIPTSElement {
	return e.MATHSIZE_STR(fmt.Sprintf(format, args...))
}

func (e *MathMLMMULTISCRIPTSElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfMATHSIZE_STRF(condition bool, format string, args ...any) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.MATHSIZE_STR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMMULTISCRIPTSElement) MATHSIZE_STRRemove(s string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathsize")
	return e
}

func (e *MathMLMMULTISCRIPTSElement) MATHSIZE_STRRemoveF(format string, args ...any) *MathMLMMULTISCRIPTSElement {
	return e.MATHSIZE_STRRemove(fmt.Sprintf(format, args...))
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLMMULTISCRIPTSElement) NONCE(s string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) NONCEF(format string, args ...any) *MathMLMMULTISCRIPTSElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *MathMLMMULTISCRIPTSElement) IfNONCE(condition bool, s string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfNONCEF(condition bool, format string, args ...any) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMMULTISCRIPTSElement) NONCERemove(s string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("nonce")
	return e
}

func (e *MathMLMMULTISCRIPTSElement) NONCERemoveF(format string, args ...any) *MathMLMMULTISCRIPTSElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMMULTISCRIPTSElement) SCRIPTLEVEL(i int) *MathMLMMULTISCRIPTSElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMMULTISCRIPTSElement) SCRIPTLEVELRemove(i int) *MathMLMMULTISCRIPTSElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMMULTISCRIPTSElement) STYLEF(k string, format string, args ...any) *MathMLMMULTISCRIPTSElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMMULTISCRIPTSElement) IfSTYLE(condition bool, k string, v string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMMULTISCRIPTSElement) STYLE(k string, v string) *MathMLMMULTISCRIPTSElement {
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

func (e *MathMLMMULTISCRIPTSElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMMULTISCRIPTSElement) STYLEMap(m map[string]string) *MathMLMMULTISCRIPTSElement {
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
func (e *MathMLMMULTISCRIPTSElement) STYLEPairs(pairs ...string) *MathMLMMULTISCRIPTSElement {
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

func (e *MathMLMMULTISCRIPTSElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMMULTISCRIPTSElement) STYLERemove(keys ...string) *MathMLMMULTISCRIPTSElement {
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
func (e *MathMLMMULTISCRIPTSElement) TABINDEX(i int) *MathMLMMULTISCRIPTSElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfTABINDEX(condition bool, i int) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMMULTISCRIPTSElement) TABINDEXRemove(i int) *MathMLMMULTISCRIPTSElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("tabindex")
	return e
}

// Make a request for an HTML

func (e *MathMLMMULTISCRIPTSElement) Z_REQ(expression string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfZ_REQ(condition bool, expression string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *MathMLMMULTISCRIPTSElement) Z_REQRemove() *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *MathMLMMULTISCRIPTSElement) Z_TARGET(expression string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfZ_TARGET(condition bool, expression string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *MathMLMMULTISCRIPTSElement) Z_TARGETRemove() *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *MathMLMMULTISCRIPTSElement) Z_REQ_SELECTOR(expression string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfZ_REQ_SELECTOR(condition bool, expression string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *MathMLMMULTISCRIPTSElement) Z_REQ_SELECTORRemove() *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *MathMLMMULTISCRIPTSElement) Z_SWAP(expression string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfZ_SWAP(condition bool, expression string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *MathMLMMULTISCRIPTSElement) Z_SWAPRemove() *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *MathMLMMULTISCRIPTSElement) Z_SWAP_PUSH(expression string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfZ_SWAP_PUSH(condition bool, expression string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *MathMLMMULTISCRIPTSElement) Z_SWAP_PUSHRemove() *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *MathMLMMULTISCRIPTSElement) Z_TRIGGER(expression string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfZ_TRIGGER(condition bool, expression string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *MathMLMMULTISCRIPTSElement) Z_TRIGGERRemove() *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *MathMLMMULTISCRIPTSElement) Z_REQ_METHOD(c MathMLMmultiscriptsZReqMethodChoice) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type MathMLMmultiscriptsZReqMethodChoice string

const (
	// default GET
	MathMLMmultiscriptsZReqMethod_empty MathMLMmultiscriptsZReqMethodChoice = ""
	// GET
	MathMLMmultiscriptsZReqMethod_get MathMLMmultiscriptsZReqMethodChoice = "get"
	// POST
	MathMLMmultiscriptsZReqMethod_post MathMLMmultiscriptsZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *MathMLMMULTISCRIPTSElement) Z_REQ_METHODRemove(c MathMLMmultiscriptsZReqMethodChoice) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *MathMLMMULTISCRIPTSElement) Z_REQ_STRATEGY(expression string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfZ_REQ_STRATEGY(condition bool, expression string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *MathMLMMULTISCRIPTSElement) Z_REQ_STRATEGYRemove() *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *MathMLMMULTISCRIPTSElement) Z_REQ_HISTORY(expression string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfZ_REQ_HISTORY(condition bool, expression string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *MathMLMMULTISCRIPTSElement) Z_REQ_HISTORYRemove() *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *MathMLMMULTISCRIPTSElement) Z_DATA(expression string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfZ_DATA(condition bool, expression string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *MathMLMMULTISCRIPTSElement) Z_DATARemove() *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *MathMLMMULTISCRIPTSElement) Z_JSON(expression string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfZ_JSON(condition bool, expression string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *MathMLMMULTISCRIPTSElement) Z_JSONRemove() *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *MathMLMMULTISCRIPTSElement) Z_REQ_BATCH(expression string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfZ_REQ_BATCH(condition bool, expression string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *MathMLMMULTISCRIPTSElement) Z_REQ_BATCHRemove() *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *MathMLMMULTISCRIPTSElement) Z_ACTION(expression string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfZ_ACTION(condition bool, expression string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *MathMLMMULTISCRIPTSElement) Z_ACTIONRemove() *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *MathMLMMULTISCRIPTSElement) Z_REQ_BEFORE(expression string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfZ_REQ_BEFORE(condition bool, expression string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *MathMLMMULTISCRIPTSElement) Z_REQ_BEFORERemove() *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *MathMLMMULTISCRIPTSElement) Z_REQ_AFTER(expression string) *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMMULTISCRIPTSElement) IfZ_REQ_AFTER(condition bool, expression string) *MathMLMMULTISCRIPTSElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *MathMLMMULTISCRIPTSElement) Z_REQ_AFTERRemove() *MathMLMMULTISCRIPTSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
