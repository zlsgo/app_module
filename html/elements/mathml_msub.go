package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// This element is used to display a subscript expression.
type MathMLMSUBElement struct {
	*Element
}

// Create a new MathMLMSUBElement element.
// This will create a new element with the tag
// "msub" during rendering.
func MathML_MSUB(children ...ElementRenderer) *MathMLMSUBElement {
	e := NewElement("msub", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMSUBElement{Element: e}
}

func (e *MathMLMSUBElement) Children(children ...ElementRenderer) *MathMLMSUBElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMSUBElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMSUBElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMSUBElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMSUBElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMSUBElement) Attr(name string, value ...string) *MathMLMSUBElement {
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

func (e *MathMLMSUBElement) Attrs(attrs ...string) *MathMLMSUBElement {
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

func (e *MathMLMSUBElement) AttrsMap(attrs map[string]string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMSUBElement) Text(text string) *MathMLMSUBElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMSUBElement) TextF(format string, args ...any) *MathMLMSUBElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBElement) IfText(condition bool, text string) *MathMLMSUBElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMSUBElement) IfTextF(condition bool, format string, args ...any) *MathMLMSUBElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMSUBElement) Escaped(text string) *MathMLMSUBElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMSUBElement) IfEscaped(condition bool, text string) *MathMLMSUBElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMSUBElement) EscapedF(format string, args ...any) *MathMLMSUBElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMSUBElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMSUBElement) CustomData(key, value string) *MathMLMSUBElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMSUBElement) IfCustomData(condition bool, key, value string) *MathMLMSUBElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMSUBElement) CustomDataF(key, format string, args ...any) *MathMLMSUBElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMSUBElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMSUBElement) CustomDataRemove(key string) *MathMLMSUBElement {
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
func (e *MathMLMSUBElement) CLASS(s ...string) *MathMLMSUBElement {
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

func (e *MathMLMSUBElement) IfCLASS(condition bool, s ...string) *MathMLMSUBElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMSUBElement) CLASSRemove(s ...string) *MathMLMSUBElement {
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
func (e *MathMLMSUBElement) DIR(c MathMLMsubDirChoice) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMsubDirChoice string

const (
	// left-to-right
	MathMLMsubDir_ltr MathMLMsubDirChoice = "ltr"
	// right-to-left
	MathMLMsubDir_rtl MathMLMsubDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMSUBElement) DIRRemove(c MathMLMsubDirChoice) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMSUBElement) DISPLAYSTYLE(c MathMLMsubDisplaystyleChoice) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMsubDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMsubDisplaystyle_true MathMLMsubDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMsubDisplaystyle_false MathMLMsubDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMSUBElement) DISPLAYSTYLERemove(c MathMLMsubDisplaystyleChoice) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMSUBElement) ID(s string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMSUBElement) IDF(format string, args ...any) *MathMLMSUBElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBElement) IfID(condition bool, s string) *MathMLMSUBElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *MathMLMSUBElement) IfIDF(condition bool, format string, args ...any) *MathMLMSUBElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMSUBElement) IDRemove(s string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *MathMLMSUBElement) IDRemoveF(format string, args ...any) *MathMLMSUBElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMSUBElement) MATHBACKGROUND(s string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMSUBElement) MATHBACKGROUNDF(format string, args ...any) *MathMLMSUBElement {
	return e.MATHBACKGROUND(fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMSUBElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

func (e *MathMLMSUBElement) IfMATHBACKGROUNDF(condition bool, format string, args ...any) *MathMLMSUBElement {
	if condition {
		e.MATHBACKGROUND(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMSUBElement) MATHBACKGROUNDRemove(s string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathbackground")
	return e
}

func (e *MathMLMSUBElement) MATHBACKGROUNDRemoveF(format string, args ...any) *MathMLMSUBElement {
	return e.MATHBACKGROUNDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMSUBElement) MATHCOLOR(s string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMSUBElement) MATHCOLORF(format string, args ...any) *MathMLMSUBElement {
	return e.MATHCOLOR(fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBElement) IfMATHCOLOR(condition bool, s string) *MathMLMSUBElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

func (e *MathMLMSUBElement) IfMATHCOLORF(condition bool, format string, args ...any) *MathMLMSUBElement {
	if condition {
		e.MATHCOLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMSUBElement) MATHCOLORRemove(s string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathcolor")
	return e
}

func (e *MathMLMSUBElement) MATHCOLORRemoveF(format string, args ...any) *MathMLMSUBElement {
	return e.MATHCOLORRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMSUBElement) MATHSIZE_STR(s string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMSUBElement) MATHSIZE_STRF(format string, args ...any) *MathMLMSUBElement {
	return e.MATHSIZE_STR(fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMSUBElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

func (e *MathMLMSUBElement) IfMATHSIZE_STRF(condition bool, format string, args ...any) *MathMLMSUBElement {
	if condition {
		e.MATHSIZE_STR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMSUBElement) MATHSIZE_STRRemove(s string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathsize")
	return e
}

func (e *MathMLMSUBElement) MATHSIZE_STRRemoveF(format string, args ...any) *MathMLMSUBElement {
	return e.MATHSIZE_STRRemove(fmt.Sprintf(format, args...))
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLMSUBElement) NONCE(s string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMSUBElement) NONCEF(format string, args ...any) *MathMLMSUBElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBElement) IfNONCE(condition bool, s string) *MathMLMSUBElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *MathMLMSUBElement) IfNONCEF(condition bool, format string, args ...any) *MathMLMSUBElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMSUBElement) NONCERemove(s string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("nonce")
	return e
}

func (e *MathMLMSUBElement) NONCERemoveF(format string, args ...any) *MathMLMSUBElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMSUBElement) SCRIPTLEVEL(i int) *MathMLMSUBElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMSUBElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMSUBElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMSUBElement) SCRIPTLEVELRemove(i int) *MathMLMSUBElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMSUBElement) STYLEF(k string, format string, args ...any) *MathMLMSUBElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBElement) IfSTYLE(condition bool, k string, v string) *MathMLMSUBElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMSUBElement) STYLE(k string, v string) *MathMLMSUBElement {
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

func (e *MathMLMSUBElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMSUBElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMSUBElement) STYLEMap(m map[string]string) *MathMLMSUBElement {
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
func (e *MathMLMSUBElement) STYLEPairs(pairs ...string) *MathMLMSUBElement {
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

func (e *MathMLMSUBElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMSUBElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMSUBElement) STYLERemove(keys ...string) *MathMLMSUBElement {
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
func (e *MathMLMSUBElement) TABINDEX(i int) *MathMLMSUBElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMSUBElement) IfTABINDEX(condition bool, i int) *MathMLMSUBElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMSUBElement) TABINDEXRemove(i int) *MathMLMSUBElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("tabindex")
	return e
}

// Make a request for an HTML

func (e *MathMLMSUBElement) Z_REQ(expression string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBElement) IfZ_REQ(condition bool, expression string) *MathMLMSUBElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *MathMLMSUBElement) Z_REQRemove() *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *MathMLMSUBElement) Z_TARGET(expression string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBElement) IfZ_TARGET(condition bool, expression string) *MathMLMSUBElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *MathMLMSUBElement) Z_TARGETRemove() *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *MathMLMSUBElement) Z_REQ_SELECTOR(expression string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBElement) IfZ_REQ_SELECTOR(condition bool, expression string) *MathMLMSUBElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *MathMLMSUBElement) Z_REQ_SELECTORRemove() *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *MathMLMSUBElement) Z_SWAP(expression string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBElement) IfZ_SWAP(condition bool, expression string) *MathMLMSUBElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *MathMLMSUBElement) Z_SWAPRemove() *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *MathMLMSUBElement) Z_SWAP_PUSH(expression string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBElement) IfZ_SWAP_PUSH(condition bool, expression string) *MathMLMSUBElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *MathMLMSUBElement) Z_SWAP_PUSHRemove() *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *MathMLMSUBElement) Z_TRIGGER(expression string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBElement) IfZ_TRIGGER(condition bool, expression string) *MathMLMSUBElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *MathMLMSUBElement) Z_TRIGGERRemove() *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *MathMLMSUBElement) Z_REQ_METHOD(c MathMLMsubZReqMethodChoice) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type MathMLMsubZReqMethodChoice string

const (
	// default GET
	MathMLMsubZReqMethod_empty MathMLMsubZReqMethodChoice = ""
	// GET
	MathMLMsubZReqMethod_get MathMLMsubZReqMethodChoice = "get"
	// POST
	MathMLMsubZReqMethod_post MathMLMsubZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *MathMLMSUBElement) Z_REQ_METHODRemove(c MathMLMsubZReqMethodChoice) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *MathMLMSUBElement) Z_REQ_STRATEGY(expression string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBElement) IfZ_REQ_STRATEGY(condition bool, expression string) *MathMLMSUBElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *MathMLMSUBElement) Z_REQ_STRATEGYRemove() *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *MathMLMSUBElement) Z_REQ_HISTORY(expression string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBElement) IfZ_REQ_HISTORY(condition bool, expression string) *MathMLMSUBElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *MathMLMSUBElement) Z_REQ_HISTORYRemove() *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *MathMLMSUBElement) Z_DATA(expression string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBElement) IfZ_DATA(condition bool, expression string) *MathMLMSUBElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *MathMLMSUBElement) Z_DATARemove() *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *MathMLMSUBElement) Z_JSON(expression string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBElement) IfZ_JSON(condition bool, expression string) *MathMLMSUBElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *MathMLMSUBElement) Z_JSONRemove() *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *MathMLMSUBElement) Z_REQ_BATCH(expression string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBElement) IfZ_REQ_BATCH(condition bool, expression string) *MathMLMSUBElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *MathMLMSUBElement) Z_REQ_BATCHRemove() *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *MathMLMSUBElement) Z_ACTION(expression string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBElement) IfZ_ACTION(condition bool, expression string) *MathMLMSUBElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *MathMLMSUBElement) Z_ACTIONRemove() *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *MathMLMSUBElement) Z_REQ_BEFORE(expression string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBElement) IfZ_REQ_BEFORE(condition bool, expression string) *MathMLMSUBElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *MathMLMSUBElement) Z_REQ_BEFORERemove() *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *MathMLMSUBElement) Z_REQ_AFTER(expression string) *MathMLMSUBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBElement) IfZ_REQ_AFTER(condition bool, expression string) *MathMLMSUBElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *MathMLMSUBElement) Z_REQ_AFTERRemove() *MathMLMSUBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
