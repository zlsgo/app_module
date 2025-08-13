package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// This element is used to display a fraction.
type MathMLMFRACElement struct {
	*Element
}

// Create a new MathMLMFRACElement element.
// This will create a new element with the tag
// "mfrac" during rendering.
func MathML_MFRAC(children ...ElementRenderer) *MathMLMFRACElement {
	e := NewElement("mfrac", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMFRACElement{Element: e}
}

func (e *MathMLMFRACElement) Children(children ...ElementRenderer) *MathMLMFRACElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMFRACElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMFRACElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMFRACElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMFRACElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMFRACElement) Attr(name string, value ...string) *MathMLMFRACElement {
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

func (e *MathMLMFRACElement) Attrs(attrs ...string) *MathMLMFRACElement {
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

func (e *MathMLMFRACElement) AttrsMap(attrs map[string]string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMFRACElement) Text(text string) *MathMLMFRACElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMFRACElement) TextF(format string, args ...any) *MathMLMFRACElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfText(condition bool, text string) *MathMLMFRACElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMFRACElement) IfTextF(condition bool, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMFRACElement) Escaped(text string) *MathMLMFRACElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMFRACElement) IfEscaped(condition bool, text string) *MathMLMFRACElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMFRACElement) EscapedF(format string, args ...any) *MathMLMFRACElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMFRACElement) CustomData(key, value string) *MathMLMFRACElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMFRACElement) IfCustomData(condition bool, key, value string) *MathMLMFRACElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMFRACElement) CustomDataF(key, format string, args ...any) *MathMLMFRACElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMFRACElement) CustomDataRemove(key string) *MathMLMFRACElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// This attribute specifies whether the fraction line is to be drawn straight or
// to beveled
// Possible values are true and false.
func (e *MathMLMFRACElement) BEVELLED(c MathMLMfracBevelledChoice) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("bevelled", string(c))
	return e
}

type MathMLMfracBevelledChoice string

const (
	MathMLMfracBevelled_true MathMLMfracBevelledChoice = "true"

	MathMLMfracBevelled_false MathMLMfracBevelledChoice = "false"
)

// Remove the attribute BEVELLED from the element.
func (e *MathMLMFRACElement) BEVELLEDRemove(c MathMLMfracBevelledChoice) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("bevelled")
	return e
}

// Assigns a class name or set of class names to an element
// You may assign the same class name or names to any number of elements
// If you specify multiple class names, they must be separated by whitespace
// characters.
func (e *MathMLMFRACElement) CLASS(s ...string) *MathMLMFRACElement {
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

func (e *MathMLMFRACElement) IfCLASS(condition bool, s ...string) *MathMLMFRACElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMFRACElement) CLASSRemove(s ...string) *MathMLMFRACElement {
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
func (e *MathMLMFRACElement) DIR(c MathMLMfracDirChoice) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMfracDirChoice string

const (
	// left-to-right
	MathMLMfracDir_ltr MathMLMfracDirChoice = "ltr"
	// right-to-left
	MathMLMfracDir_rtl MathMLMfracDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMFRACElement) DIRRemove(c MathMLMfracDirChoice) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMFRACElement) DISPLAYSTYLE(c MathMLMfracDisplaystyleChoice) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMfracDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMfracDisplaystyle_true MathMLMfracDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMfracDisplaystyle_false MathMLMfracDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMFRACElement) DISPLAYSTYLERemove(c MathMLMfracDisplaystyleChoice) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMFRACElement) ID(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMFRACElement) IDF(format string, args ...any) *MathMLMFRACElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfID(condition bool, s string) *MathMLMFRACElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *MathMLMFRACElement) IfIDF(condition bool, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMFRACElement) IDRemove(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *MathMLMFRACElement) IDRemoveF(format string, args ...any) *MathMLMFRACElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMFRACElement) MATHBACKGROUND(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMFRACElement) MATHBACKGROUNDF(format string, args ...any) *MathMLMFRACElement {
	return e.MATHBACKGROUND(fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMFRACElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

func (e *MathMLMFRACElement) IfMATHBACKGROUNDF(condition bool, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.MATHBACKGROUND(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMFRACElement) MATHBACKGROUNDRemove(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathbackground")
	return e
}

func (e *MathMLMFRACElement) MATHBACKGROUNDRemoveF(format string, args ...any) *MathMLMFRACElement {
	return e.MATHBACKGROUNDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMFRACElement) MATHCOLOR(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMFRACElement) MATHCOLORF(format string, args ...any) *MathMLMFRACElement {
	return e.MATHCOLOR(fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfMATHCOLOR(condition bool, s string) *MathMLMFRACElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

func (e *MathMLMFRACElement) IfMATHCOLORF(condition bool, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.MATHCOLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMFRACElement) MATHCOLORRemove(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathcolor")
	return e
}

func (e *MathMLMFRACElement) MATHCOLORRemoveF(format string, args ...any) *MathMLMFRACElement {
	return e.MATHCOLORRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMFRACElement) MATHSIZE_STR(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMFRACElement) MATHSIZE_STRF(format string, args ...any) *MathMLMFRACElement {
	return e.MATHSIZE_STR(fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMFRACElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

func (e *MathMLMFRACElement) IfMATHSIZE_STRF(condition bool, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.MATHSIZE_STR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMFRACElement) MATHSIZE_STRRemove(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathsize")
	return e
}

func (e *MathMLMFRACElement) MATHSIZE_STRRemoveF(format string, args ...any) *MathMLMFRACElement {
	return e.MATHSIZE_STRRemove(fmt.Sprintf(format, args...))
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLMFRACElement) NONCE(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMFRACElement) NONCEF(format string, args ...any) *MathMLMFRACElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfNONCE(condition bool, s string) *MathMLMFRACElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *MathMLMFRACElement) IfNONCEF(condition bool, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMFRACElement) NONCERemove(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("nonce")
	return e
}

func (e *MathMLMFRACElement) NONCERemoveF(format string, args ...any) *MathMLMFRACElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMFRACElement) SCRIPTLEVEL(i int) *MathMLMFRACElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMFRACElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMFRACElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMFRACElement) SCRIPTLEVELRemove(i int) *MathMLMFRACElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMFRACElement) STYLEF(k string, format string, args ...any) *MathMLMFRACElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfSTYLE(condition bool, k string, v string) *MathMLMFRACElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMFRACElement) STYLE(k string, v string) *MathMLMFRACElement {
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

func (e *MathMLMFRACElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMFRACElement) STYLEMap(m map[string]string) *MathMLMFRACElement {
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
func (e *MathMLMFRACElement) STYLEPairs(pairs ...string) *MathMLMFRACElement {
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

func (e *MathMLMFRACElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMFRACElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMFRACElement) STYLERemove(keys ...string) *MathMLMFRACElement {
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
func (e *MathMLMFRACElement) TABINDEX(i int) *MathMLMFRACElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMFRACElement) IfTABINDEX(condition bool, i int) *MathMLMFRACElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMFRACElement) TABINDEXRemove(i int) *MathMLMFRACElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("tabindex")
	return e
}

// Make a request for an HTML

func (e *MathMLMFRACElement) Z_REQ(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfZ_REQ(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *MathMLMFRACElement) Z_REQRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *MathMLMFRACElement) Z_TARGET(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfZ_TARGET(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *MathMLMFRACElement) Z_TARGETRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *MathMLMFRACElement) Z_REQ_SELECTOR(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfZ_REQ_SELECTOR(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *MathMLMFRACElement) Z_REQ_SELECTORRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *MathMLMFRACElement) Z_SWAP(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfZ_SWAP(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *MathMLMFRACElement) Z_SWAPRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *MathMLMFRACElement) Z_SWAP_PUSH(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfZ_SWAP_PUSH(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *MathMLMFRACElement) Z_SWAP_PUSHRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *MathMLMFRACElement) Z_TRIGGER(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfZ_TRIGGER(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *MathMLMFRACElement) Z_TRIGGERRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *MathMLMFRACElement) Z_REQ_METHOD(c MathMLMfracZReqMethodChoice) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type MathMLMfracZReqMethodChoice string

const (
	// default GET
	MathMLMfracZReqMethod_empty MathMLMfracZReqMethodChoice = ""
	// GET
	MathMLMfracZReqMethod_get MathMLMfracZReqMethodChoice = "get"
	// POST
	MathMLMfracZReqMethod_post MathMLMfracZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *MathMLMFRACElement) Z_REQ_METHODRemove(c MathMLMfracZReqMethodChoice) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *MathMLMFRACElement) Z_REQ_STRATEGY(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfZ_REQ_STRATEGY(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *MathMLMFRACElement) Z_REQ_STRATEGYRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *MathMLMFRACElement) Z_REQ_HISTORY(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfZ_REQ_HISTORY(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *MathMLMFRACElement) Z_REQ_HISTORYRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *MathMLMFRACElement) Z_DATA(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfZ_DATA(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *MathMLMFRACElement) Z_DATARemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *MathMLMFRACElement) Z_JSON(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfZ_JSON(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *MathMLMFRACElement) Z_JSONRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *MathMLMFRACElement) Z_REQ_BATCH(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfZ_REQ_BATCH(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *MathMLMFRACElement) Z_REQ_BATCHRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *MathMLMFRACElement) Z_ACTION(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfZ_ACTION(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *MathMLMFRACElement) Z_ACTIONRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *MathMLMFRACElement) Z_REQ_BEFORE(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfZ_REQ_BEFORE(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *MathMLMFRACElement) Z_REQ_BEFORERemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *MathMLMFRACElement) Z_REQ_AFTER(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfZ_REQ_AFTER(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *MathMLMFRACElement) Z_REQ_AFTERRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
