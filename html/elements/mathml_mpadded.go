package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// This element is used to display an expression with additional spacing.
type MathMLMPADDEDElement struct {
	*Element
}

// Create a new MathMLMPADDEDElement element.
// This will create a new element with the tag
// "mpadded" during rendering.
func MathML_MPADDED(children ...ElementRenderer) *MathMLMPADDEDElement {
	e := NewElement("mpadded", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMPADDEDElement{Element: e}
}

func (e *MathMLMPADDEDElement) Children(children ...ElementRenderer) *MathMLMPADDEDElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMPADDEDElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMPADDEDElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMPADDEDElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMPADDEDElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMPADDEDElement) Attr(name string, value ...string) *MathMLMPADDEDElement {
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

func (e *MathMLMPADDEDElement) Attrs(attrs ...string) *MathMLMPADDEDElement {
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

func (e *MathMLMPADDEDElement) AttrsMap(attrs map[string]string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMPADDEDElement) Text(text string) *MathMLMPADDEDElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMPADDEDElement) TextF(format string, args ...any) *MathMLMPADDEDElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMPADDEDElement) IfText(condition bool, text string) *MathMLMPADDEDElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMPADDEDElement) IfTextF(condition bool, format string, args ...any) *MathMLMPADDEDElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMPADDEDElement) Escaped(text string) *MathMLMPADDEDElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMPADDEDElement) IfEscaped(condition bool, text string) *MathMLMPADDEDElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMPADDEDElement) EscapedF(format string, args ...any) *MathMLMPADDEDElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMPADDEDElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMPADDEDElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMPADDEDElement) CustomData(key, value string) *MathMLMPADDEDElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMPADDEDElement) IfCustomData(condition bool, key, value string) *MathMLMPADDEDElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMPADDEDElement) CustomDataF(key, format string, args ...any) *MathMLMPADDEDElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMPADDEDElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMPADDEDElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMPADDEDElement) CustomDataRemove(key string) *MathMLMPADDEDElement {
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
func (e *MathMLMPADDEDElement) CLASS(s ...string) *MathMLMPADDEDElement {
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

func (e *MathMLMPADDEDElement) IfCLASS(condition bool, s ...string) *MathMLMPADDEDElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMPADDEDElement) CLASSRemove(s ...string) *MathMLMPADDEDElement {
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
func (e *MathMLMPADDEDElement) DIR(c MathMLMpaddedDirChoice) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMpaddedDirChoice string

const (
	// left-to-right
	MathMLMpaddedDir_ltr MathMLMpaddedDirChoice = "ltr"
	// right-to-left
	MathMLMpaddedDir_rtl MathMLMpaddedDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMPADDEDElement) DIRRemove(c MathMLMpaddedDirChoice) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMPADDEDElement) DISPLAYSTYLE(c MathMLMpaddedDisplaystyleChoice) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMpaddedDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMpaddedDisplaystyle_true MathMLMpaddedDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMpaddedDisplaystyle_false MathMLMpaddedDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMPADDEDElement) DISPLAYSTYLERemove(c MathMLMpaddedDisplaystyleChoice) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMPADDEDElement) ID(s string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMPADDEDElement) IDF(format string, args ...any) *MathMLMPADDEDElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *MathMLMPADDEDElement) IfID(condition bool, s string) *MathMLMPADDEDElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *MathMLMPADDEDElement) IfIDF(condition bool, format string, args ...any) *MathMLMPADDEDElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMPADDEDElement) IDRemove(s string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *MathMLMPADDEDElement) IDRemoveF(format string, args ...any) *MathMLMPADDEDElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMPADDEDElement) MATHBACKGROUND(s string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMPADDEDElement) MATHBACKGROUNDF(format string, args ...any) *MathMLMPADDEDElement {
	return e.MATHBACKGROUND(fmt.Sprintf(format, args...))
}

func (e *MathMLMPADDEDElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMPADDEDElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

func (e *MathMLMPADDEDElement) IfMATHBACKGROUNDF(condition bool, format string, args ...any) *MathMLMPADDEDElement {
	if condition {
		e.MATHBACKGROUND(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMPADDEDElement) MATHBACKGROUNDRemove(s string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathbackground")
	return e
}

func (e *MathMLMPADDEDElement) MATHBACKGROUNDRemoveF(format string, args ...any) *MathMLMPADDEDElement {
	return e.MATHBACKGROUNDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMPADDEDElement) MATHCOLOR(s string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMPADDEDElement) MATHCOLORF(format string, args ...any) *MathMLMPADDEDElement {
	return e.MATHCOLOR(fmt.Sprintf(format, args...))
}

func (e *MathMLMPADDEDElement) IfMATHCOLOR(condition bool, s string) *MathMLMPADDEDElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

func (e *MathMLMPADDEDElement) IfMATHCOLORF(condition bool, format string, args ...any) *MathMLMPADDEDElement {
	if condition {
		e.MATHCOLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMPADDEDElement) MATHCOLORRemove(s string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathcolor")
	return e
}

func (e *MathMLMPADDEDElement) MATHCOLORRemoveF(format string, args ...any) *MathMLMPADDEDElement {
	return e.MATHCOLORRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMPADDEDElement) MATHSIZE_STR(s string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMPADDEDElement) MATHSIZE_STRF(format string, args ...any) *MathMLMPADDEDElement {
	return e.MATHSIZE_STR(fmt.Sprintf(format, args...))
}

func (e *MathMLMPADDEDElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMPADDEDElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

func (e *MathMLMPADDEDElement) IfMATHSIZE_STRF(condition bool, format string, args ...any) *MathMLMPADDEDElement {
	if condition {
		e.MATHSIZE_STR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMPADDEDElement) MATHSIZE_STRRemove(s string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathsize")
	return e
}

func (e *MathMLMPADDEDElement) MATHSIZE_STRRemoveF(format string, args ...any) *MathMLMPADDEDElement {
	return e.MATHSIZE_STRRemove(fmt.Sprintf(format, args...))
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLMPADDEDElement) NONCE(s string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMPADDEDElement) NONCEF(format string, args ...any) *MathMLMPADDEDElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *MathMLMPADDEDElement) IfNONCE(condition bool, s string) *MathMLMPADDEDElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *MathMLMPADDEDElement) IfNONCEF(condition bool, format string, args ...any) *MathMLMPADDEDElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMPADDEDElement) NONCERemove(s string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("nonce")
	return e
}

func (e *MathMLMPADDEDElement) NONCERemoveF(format string, args ...any) *MathMLMPADDEDElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMPADDEDElement) SCRIPTLEVEL(i int) *MathMLMPADDEDElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMPADDEDElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMPADDEDElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMPADDEDElement) SCRIPTLEVELRemove(i int) *MathMLMPADDEDElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMPADDEDElement) STYLEF(k string, format string, args ...any) *MathMLMPADDEDElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMPADDEDElement) IfSTYLE(condition bool, k string, v string) *MathMLMPADDEDElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMPADDEDElement) STYLE(k string, v string) *MathMLMPADDEDElement {
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

func (e *MathMLMPADDEDElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMPADDEDElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMPADDEDElement) STYLEMap(m map[string]string) *MathMLMPADDEDElement {
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
func (e *MathMLMPADDEDElement) STYLEPairs(pairs ...string) *MathMLMPADDEDElement {
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

func (e *MathMLMPADDEDElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMPADDEDElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMPADDEDElement) STYLERemove(keys ...string) *MathMLMPADDEDElement {
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
func (e *MathMLMPADDEDElement) TABINDEX(i int) *MathMLMPADDEDElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMPADDEDElement) IfTABINDEX(condition bool, i int) *MathMLMPADDEDElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMPADDEDElement) TABINDEXRemove(i int) *MathMLMPADDEDElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("tabindex")
	return e
}

// Make a request for an HTML

func (e *MathMLMPADDEDElement) Z_REQ(expression string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMPADDEDElement) IfZ_REQ(condition bool, expression string) *MathMLMPADDEDElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *MathMLMPADDEDElement) Z_REQRemove() *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *MathMLMPADDEDElement) Z_TARGET(expression string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMPADDEDElement) IfZ_TARGET(condition bool, expression string) *MathMLMPADDEDElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *MathMLMPADDEDElement) Z_TARGETRemove() *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *MathMLMPADDEDElement) Z_REQ_SELECTOR(expression string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMPADDEDElement) IfZ_REQ_SELECTOR(condition bool, expression string) *MathMLMPADDEDElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *MathMLMPADDEDElement) Z_REQ_SELECTORRemove() *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *MathMLMPADDEDElement) Z_SWAP(expression string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMPADDEDElement) IfZ_SWAP(condition bool, expression string) *MathMLMPADDEDElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *MathMLMPADDEDElement) Z_SWAPRemove() *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *MathMLMPADDEDElement) Z_SWAP_PUSH(expression string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMPADDEDElement) IfZ_SWAP_PUSH(condition bool, expression string) *MathMLMPADDEDElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *MathMLMPADDEDElement) Z_SWAP_PUSHRemove() *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *MathMLMPADDEDElement) Z_TRIGGER(expression string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMPADDEDElement) IfZ_TRIGGER(condition bool, expression string) *MathMLMPADDEDElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *MathMLMPADDEDElement) Z_TRIGGERRemove() *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *MathMLMPADDEDElement) Z_REQ_METHOD(c MathMLMpaddedZReqMethodChoice) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type MathMLMpaddedZReqMethodChoice string

const (
	// default GET
	MathMLMpaddedZReqMethod_empty MathMLMpaddedZReqMethodChoice = ""
	// GET
	MathMLMpaddedZReqMethod_get MathMLMpaddedZReqMethodChoice = "get"
	// POST
	MathMLMpaddedZReqMethod_post MathMLMpaddedZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *MathMLMPADDEDElement) Z_REQ_METHODRemove(c MathMLMpaddedZReqMethodChoice) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *MathMLMPADDEDElement) Z_REQ_STRATEGY(expression string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMPADDEDElement) IfZ_REQ_STRATEGY(condition bool, expression string) *MathMLMPADDEDElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *MathMLMPADDEDElement) Z_REQ_STRATEGYRemove() *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *MathMLMPADDEDElement) Z_REQ_HISTORY(expression string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMPADDEDElement) IfZ_REQ_HISTORY(condition bool, expression string) *MathMLMPADDEDElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *MathMLMPADDEDElement) Z_REQ_HISTORYRemove() *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *MathMLMPADDEDElement) Z_DATA(expression string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMPADDEDElement) IfZ_DATA(condition bool, expression string) *MathMLMPADDEDElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *MathMLMPADDEDElement) Z_DATARemove() *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *MathMLMPADDEDElement) Z_JSON(expression string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMPADDEDElement) IfZ_JSON(condition bool, expression string) *MathMLMPADDEDElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *MathMLMPADDEDElement) Z_JSONRemove() *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *MathMLMPADDEDElement) Z_REQ_BATCH(expression string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMPADDEDElement) IfZ_REQ_BATCH(condition bool, expression string) *MathMLMPADDEDElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *MathMLMPADDEDElement) Z_REQ_BATCHRemove() *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *MathMLMPADDEDElement) Z_ACTION(expression string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMPADDEDElement) IfZ_ACTION(condition bool, expression string) *MathMLMPADDEDElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *MathMLMPADDEDElement) Z_ACTIONRemove() *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *MathMLMPADDEDElement) Z_REQ_BEFORE(expression string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMPADDEDElement) IfZ_REQ_BEFORE(condition bool, expression string) *MathMLMPADDEDElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *MathMLMPADDEDElement) Z_REQ_BEFORERemove() *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *MathMLMPADDEDElement) Z_REQ_AFTER(expression string) *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMPADDEDElement) IfZ_REQ_AFTER(condition bool, expression string) *MathMLMPADDEDElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *MathMLMPADDEDElement) Z_REQ_AFTERRemove() *MathMLMPADDEDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
