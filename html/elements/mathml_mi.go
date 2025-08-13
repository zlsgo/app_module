package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// This element is used to display a single identifier or a single operator.
type MathMLMIElement struct {
	*Element
}

// Create a new MathMLMIElement element.
// This will create a new element with the tag
// "mi" during rendering.
func MathML_MI(children ...ElementRenderer) *MathMLMIElement {
	e := NewElement("mi", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMIElement{Element: e}
}

func (e *MathMLMIElement) Children(children ...ElementRenderer) *MathMLMIElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMIElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMIElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMIElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMIElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMIElement) Attr(name string, value ...string) *MathMLMIElement {
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

func (e *MathMLMIElement) Attrs(attrs ...string) *MathMLMIElement {
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

func (e *MathMLMIElement) AttrsMap(attrs map[string]string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMIElement) Text(text string) *MathMLMIElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMIElement) TextF(format string, args ...any) *MathMLMIElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMIElement) IfText(condition bool, text string) *MathMLMIElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMIElement) IfTextF(condition bool, format string, args ...any) *MathMLMIElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMIElement) Escaped(text string) *MathMLMIElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMIElement) IfEscaped(condition bool, text string) *MathMLMIElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMIElement) EscapedF(format string, args ...any) *MathMLMIElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMIElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMIElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMIElement) CustomData(key, value string) *MathMLMIElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMIElement) IfCustomData(condition bool, key, value string) *MathMLMIElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMIElement) CustomDataF(key, format string, args ...any) *MathMLMIElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMIElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMIElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMIElement) CustomDataRemove(key string) *MathMLMIElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// This attribute specifies the variant form of the character
// Possible values are normal, bold, italic, bold-italic, double-struck,
// bold-fraktur, script, bold-script, fraktur, sans-serif, bold-sans-serif,
// sans-serif-italic, sans-serif-bold-italic, monospace, initial, and tailed.
func (e *MathMLMIElement) MATHVARIANT(c MathMLMiMathvariantChoice) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathvariant", string(c))
	return e
}

type MathMLMiMathvariantChoice string

const (
	MathMLMiMathvariant_normal MathMLMiMathvariantChoice = "normal"

	MathMLMiMathvariant_bold MathMLMiMathvariantChoice = "bold"

	MathMLMiMathvariant_italic MathMLMiMathvariantChoice = "italic"

	MathMLMiMathvariant_bold_italic MathMLMiMathvariantChoice = "bold-italic"

	MathMLMiMathvariant_double_struck MathMLMiMathvariantChoice = "double-struck"

	MathMLMiMathvariant_bold_fraktur MathMLMiMathvariantChoice = "bold-fraktur"

	MathMLMiMathvariant_script MathMLMiMathvariantChoice = "script"

	MathMLMiMathvariant_bold_script MathMLMiMathvariantChoice = "bold-script"

	MathMLMiMathvariant_fraktur MathMLMiMathvariantChoice = "fraktur"

	MathMLMiMathvariant_sans_serif MathMLMiMathvariantChoice = "sans-serif"

	MathMLMiMathvariant_bold_sans_serif MathMLMiMathvariantChoice = "bold-sans-serif"

	MathMLMiMathvariant_sans_serif_italic MathMLMiMathvariantChoice = "sans-serif-italic"

	MathMLMiMathvariant_sans_serif_bold_italic MathMLMiMathvariantChoice = "sans-serif-bold-italic"

	MathMLMiMathvariant_monospace MathMLMiMathvariantChoice = "monospace"

	MathMLMiMathvariant_initial MathMLMiMathvariantChoice = "initial"

	MathMLMiMathvariant_tailed MathMLMiMathvariantChoice = "tailed"
)

// Remove the attribute MATHVARIANT from the element.
func (e *MathMLMIElement) MATHVARIANTRemove(c MathMLMiMathvariantChoice) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathvariant")
	return e
}

// Assigns a class name or set of class names to an element
// You may assign the same class name or names to any number of elements
// If you specify multiple class names, they must be separated by whitespace
// characters.
func (e *MathMLMIElement) CLASS(s ...string) *MathMLMIElement {
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

func (e *MathMLMIElement) IfCLASS(condition bool, s ...string) *MathMLMIElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMIElement) CLASSRemove(s ...string) *MathMLMIElement {
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
func (e *MathMLMIElement) DIR(c MathMLMiDirChoice) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMiDirChoice string

const (
	// left-to-right
	MathMLMiDir_ltr MathMLMiDirChoice = "ltr"
	// right-to-left
	MathMLMiDir_rtl MathMLMiDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMIElement) DIRRemove(c MathMLMiDirChoice) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMIElement) DISPLAYSTYLE(c MathMLMiDisplaystyleChoice) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMiDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMiDisplaystyle_true MathMLMiDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMiDisplaystyle_false MathMLMiDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMIElement) DISPLAYSTYLERemove(c MathMLMiDisplaystyleChoice) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMIElement) ID(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMIElement) IDF(format string, args ...any) *MathMLMIElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *MathMLMIElement) IfID(condition bool, s string) *MathMLMIElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *MathMLMIElement) IfIDF(condition bool, format string, args ...any) *MathMLMIElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMIElement) IDRemove(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *MathMLMIElement) IDRemoveF(format string, args ...any) *MathMLMIElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMIElement) MATHBACKGROUND(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMIElement) MATHBACKGROUNDF(format string, args ...any) *MathMLMIElement {
	return e.MATHBACKGROUND(fmt.Sprintf(format, args...))
}

func (e *MathMLMIElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMIElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

func (e *MathMLMIElement) IfMATHBACKGROUNDF(condition bool, format string, args ...any) *MathMLMIElement {
	if condition {
		e.MATHBACKGROUND(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMIElement) MATHBACKGROUNDRemove(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathbackground")
	return e
}

func (e *MathMLMIElement) MATHBACKGROUNDRemoveF(format string, args ...any) *MathMLMIElement {
	return e.MATHBACKGROUNDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMIElement) MATHCOLOR(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMIElement) MATHCOLORF(format string, args ...any) *MathMLMIElement {
	return e.MATHCOLOR(fmt.Sprintf(format, args...))
}

func (e *MathMLMIElement) IfMATHCOLOR(condition bool, s string) *MathMLMIElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

func (e *MathMLMIElement) IfMATHCOLORF(condition bool, format string, args ...any) *MathMLMIElement {
	if condition {
		e.MATHCOLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMIElement) MATHCOLORRemove(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathcolor")
	return e
}

func (e *MathMLMIElement) MATHCOLORRemoveF(format string, args ...any) *MathMLMIElement {
	return e.MATHCOLORRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMIElement) MATHSIZE_STR(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMIElement) MATHSIZE_STRF(format string, args ...any) *MathMLMIElement {
	return e.MATHSIZE_STR(fmt.Sprintf(format, args...))
}

func (e *MathMLMIElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMIElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

func (e *MathMLMIElement) IfMATHSIZE_STRF(condition bool, format string, args ...any) *MathMLMIElement {
	if condition {
		e.MATHSIZE_STR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMIElement) MATHSIZE_STRRemove(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathsize")
	return e
}

func (e *MathMLMIElement) MATHSIZE_STRRemoveF(format string, args ...any) *MathMLMIElement {
	return e.MATHSIZE_STRRemove(fmt.Sprintf(format, args...))
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLMIElement) NONCE(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMIElement) NONCEF(format string, args ...any) *MathMLMIElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *MathMLMIElement) IfNONCE(condition bool, s string) *MathMLMIElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *MathMLMIElement) IfNONCEF(condition bool, format string, args ...any) *MathMLMIElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMIElement) NONCERemove(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("nonce")
	return e
}

func (e *MathMLMIElement) NONCERemoveF(format string, args ...any) *MathMLMIElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMIElement) SCRIPTLEVEL(i int) *MathMLMIElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMIElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMIElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMIElement) SCRIPTLEVELRemove(i int) *MathMLMIElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMIElement) STYLEF(k string, format string, args ...any) *MathMLMIElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMIElement) IfSTYLE(condition bool, k string, v string) *MathMLMIElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMIElement) STYLE(k string, v string) *MathMLMIElement {
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

func (e *MathMLMIElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMIElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMIElement) STYLEMap(m map[string]string) *MathMLMIElement {
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
func (e *MathMLMIElement) STYLEPairs(pairs ...string) *MathMLMIElement {
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

func (e *MathMLMIElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMIElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMIElement) STYLERemove(keys ...string) *MathMLMIElement {
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
func (e *MathMLMIElement) TABINDEX(i int) *MathMLMIElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMIElement) IfTABINDEX(condition bool, i int) *MathMLMIElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMIElement) TABINDEXRemove(i int) *MathMLMIElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("tabindex")
	return e
}

// Make a request for an HTML

func (e *MathMLMIElement) Z_REQ(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfZ_REQ(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *MathMLMIElement) Z_REQRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *MathMLMIElement) Z_TARGET(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfZ_TARGET(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *MathMLMIElement) Z_TARGETRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *MathMLMIElement) Z_REQ_SELECTOR(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfZ_REQ_SELECTOR(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *MathMLMIElement) Z_REQ_SELECTORRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *MathMLMIElement) Z_SWAP(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfZ_SWAP(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *MathMLMIElement) Z_SWAPRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *MathMLMIElement) Z_SWAP_PUSH(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfZ_SWAP_PUSH(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *MathMLMIElement) Z_SWAP_PUSHRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *MathMLMIElement) Z_TRIGGER(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfZ_TRIGGER(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *MathMLMIElement) Z_TRIGGERRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *MathMLMIElement) Z_REQ_METHOD(c MathMLMiZReqMethodChoice) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type MathMLMiZReqMethodChoice string

const (
	// default GET
	MathMLMiZReqMethod_empty MathMLMiZReqMethodChoice = ""
	// GET
	MathMLMiZReqMethod_get MathMLMiZReqMethodChoice = "get"
	// POST
	MathMLMiZReqMethod_post MathMLMiZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *MathMLMIElement) Z_REQ_METHODRemove(c MathMLMiZReqMethodChoice) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *MathMLMIElement) Z_REQ_STRATEGY(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfZ_REQ_STRATEGY(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *MathMLMIElement) Z_REQ_STRATEGYRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *MathMLMIElement) Z_REQ_HISTORY(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfZ_REQ_HISTORY(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *MathMLMIElement) Z_REQ_HISTORYRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *MathMLMIElement) Z_DATA(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfZ_DATA(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *MathMLMIElement) Z_DATARemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *MathMLMIElement) Z_JSON(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfZ_JSON(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *MathMLMIElement) Z_JSONRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *MathMLMIElement) Z_REQ_BATCH(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfZ_REQ_BATCH(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *MathMLMIElement) Z_REQ_BATCHRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *MathMLMIElement) Z_ACTION(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfZ_ACTION(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *MathMLMIElement) Z_ACTIONRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *MathMLMIElement) Z_REQ_BEFORE(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfZ_REQ_BEFORE(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *MathMLMIElement) Z_REQ_BEFORERemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *MathMLMIElement) Z_REQ_AFTER(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfZ_REQ_AFTER(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *MathMLMIElement) Z_REQ_AFTERRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
