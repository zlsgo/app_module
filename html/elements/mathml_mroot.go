package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// This element is used to display an expression with a radical.
type MathMLMROOTElement struct {
	*Element
}

// Create a new MathMLMROOTElement element.
// This will create a new element with the tag
// "mroot" during rendering.
func MathML_MROOT(children ...ElementRenderer) *MathMLMROOTElement {
	e := NewElement("mroot", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMROOTElement{Element: e}
}

func (e *MathMLMROOTElement) Children(children ...ElementRenderer) *MathMLMROOTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMROOTElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMROOTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMROOTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMROOTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMROOTElement) Attr(name string, value ...string) *MathMLMROOTElement {
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

func (e *MathMLMROOTElement) Attrs(attrs ...string) *MathMLMROOTElement {
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

func (e *MathMLMROOTElement) AttrsMap(attrs map[string]string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMROOTElement) Text(text string) *MathMLMROOTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMROOTElement) TextF(format string, args ...any) *MathMLMROOTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMROOTElement) IfText(condition bool, text string) *MathMLMROOTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMROOTElement) IfTextF(condition bool, format string, args ...any) *MathMLMROOTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMROOTElement) Escaped(text string) *MathMLMROOTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMROOTElement) IfEscaped(condition bool, text string) *MathMLMROOTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMROOTElement) EscapedF(format string, args ...any) *MathMLMROOTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMROOTElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMROOTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMROOTElement) CustomData(key, value string) *MathMLMROOTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMROOTElement) IfCustomData(condition bool, key, value string) *MathMLMROOTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMROOTElement) CustomDataF(key, format string, args ...any) *MathMLMROOTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMROOTElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMROOTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMROOTElement) CustomDataRemove(key string) *MathMLMROOTElement {
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
func (e *MathMLMROOTElement) CLASS(s ...string) *MathMLMROOTElement {
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

func (e *MathMLMROOTElement) IfCLASS(condition bool, s ...string) *MathMLMROOTElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMROOTElement) CLASSRemove(s ...string) *MathMLMROOTElement {
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
func (e *MathMLMROOTElement) DIR(c MathMLMrootDirChoice) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMrootDirChoice string

const (
	// left-to-right
	MathMLMrootDir_ltr MathMLMrootDirChoice = "ltr"
	// right-to-left
	MathMLMrootDir_rtl MathMLMrootDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMROOTElement) DIRRemove(c MathMLMrootDirChoice) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMROOTElement) DISPLAYSTYLE(c MathMLMrootDisplaystyleChoice) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMrootDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMrootDisplaystyle_true MathMLMrootDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMrootDisplaystyle_false MathMLMrootDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMROOTElement) DISPLAYSTYLERemove(c MathMLMrootDisplaystyleChoice) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMROOTElement) ID(s string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMROOTElement) IDF(format string, args ...any) *MathMLMROOTElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *MathMLMROOTElement) IfID(condition bool, s string) *MathMLMROOTElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *MathMLMROOTElement) IfIDF(condition bool, format string, args ...any) *MathMLMROOTElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMROOTElement) IDRemove(s string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *MathMLMROOTElement) IDRemoveF(format string, args ...any) *MathMLMROOTElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMROOTElement) MATHBACKGROUND(s string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMROOTElement) MATHBACKGROUNDF(format string, args ...any) *MathMLMROOTElement {
	return e.MATHBACKGROUND(fmt.Sprintf(format, args...))
}

func (e *MathMLMROOTElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMROOTElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

func (e *MathMLMROOTElement) IfMATHBACKGROUNDF(condition bool, format string, args ...any) *MathMLMROOTElement {
	if condition {
		e.MATHBACKGROUND(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMROOTElement) MATHBACKGROUNDRemove(s string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathbackground")
	return e
}

func (e *MathMLMROOTElement) MATHBACKGROUNDRemoveF(format string, args ...any) *MathMLMROOTElement {
	return e.MATHBACKGROUNDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMROOTElement) MATHCOLOR(s string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMROOTElement) MATHCOLORF(format string, args ...any) *MathMLMROOTElement {
	return e.MATHCOLOR(fmt.Sprintf(format, args...))
}

func (e *MathMLMROOTElement) IfMATHCOLOR(condition bool, s string) *MathMLMROOTElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

func (e *MathMLMROOTElement) IfMATHCOLORF(condition bool, format string, args ...any) *MathMLMROOTElement {
	if condition {
		e.MATHCOLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMROOTElement) MATHCOLORRemove(s string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathcolor")
	return e
}

func (e *MathMLMROOTElement) MATHCOLORRemoveF(format string, args ...any) *MathMLMROOTElement {
	return e.MATHCOLORRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMROOTElement) MATHSIZE_STR(s string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMROOTElement) MATHSIZE_STRF(format string, args ...any) *MathMLMROOTElement {
	return e.MATHSIZE_STR(fmt.Sprintf(format, args...))
}

func (e *MathMLMROOTElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMROOTElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

func (e *MathMLMROOTElement) IfMATHSIZE_STRF(condition bool, format string, args ...any) *MathMLMROOTElement {
	if condition {
		e.MATHSIZE_STR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMROOTElement) MATHSIZE_STRRemove(s string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathsize")
	return e
}

func (e *MathMLMROOTElement) MATHSIZE_STRRemoveF(format string, args ...any) *MathMLMROOTElement {
	return e.MATHSIZE_STRRemove(fmt.Sprintf(format, args...))
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLMROOTElement) NONCE(s string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMROOTElement) NONCEF(format string, args ...any) *MathMLMROOTElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *MathMLMROOTElement) IfNONCE(condition bool, s string) *MathMLMROOTElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *MathMLMROOTElement) IfNONCEF(condition bool, format string, args ...any) *MathMLMROOTElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMROOTElement) NONCERemove(s string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("nonce")
	return e
}

func (e *MathMLMROOTElement) NONCERemoveF(format string, args ...any) *MathMLMROOTElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMROOTElement) SCRIPTLEVEL(i int) *MathMLMROOTElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMROOTElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMROOTElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMROOTElement) SCRIPTLEVELRemove(i int) *MathMLMROOTElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMROOTElement) STYLEF(k string, format string, args ...any) *MathMLMROOTElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMROOTElement) IfSTYLE(condition bool, k string, v string) *MathMLMROOTElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMROOTElement) STYLE(k string, v string) *MathMLMROOTElement {
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

func (e *MathMLMROOTElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMROOTElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMROOTElement) STYLEMap(m map[string]string) *MathMLMROOTElement {
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
func (e *MathMLMROOTElement) STYLEPairs(pairs ...string) *MathMLMROOTElement {
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

func (e *MathMLMROOTElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMROOTElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMROOTElement) STYLERemove(keys ...string) *MathMLMROOTElement {
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
func (e *MathMLMROOTElement) TABINDEX(i int) *MathMLMROOTElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMROOTElement) IfTABINDEX(condition bool, i int) *MathMLMROOTElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMROOTElement) TABINDEXRemove(i int) *MathMLMROOTElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("tabindex")
	return e
}

// Make a request for an HTML

func (e *MathMLMROOTElement) Z_REQ(expression string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROOTElement) IfZ_REQ(condition bool, expression string) *MathMLMROOTElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *MathMLMROOTElement) Z_REQRemove() *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *MathMLMROOTElement) Z_TARGET(expression string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROOTElement) IfZ_TARGET(condition bool, expression string) *MathMLMROOTElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *MathMLMROOTElement) Z_TARGETRemove() *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *MathMLMROOTElement) Z_REQ_SELECTOR(expression string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROOTElement) IfZ_REQ_SELECTOR(condition bool, expression string) *MathMLMROOTElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *MathMLMROOTElement) Z_REQ_SELECTORRemove() *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *MathMLMROOTElement) Z_SWAP(expression string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROOTElement) IfZ_SWAP(condition bool, expression string) *MathMLMROOTElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *MathMLMROOTElement) Z_SWAPRemove() *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *MathMLMROOTElement) Z_SWAP_PUSH(expression string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROOTElement) IfZ_SWAP_PUSH(condition bool, expression string) *MathMLMROOTElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *MathMLMROOTElement) Z_SWAP_PUSHRemove() *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *MathMLMROOTElement) Z_TRIGGER(expression string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROOTElement) IfZ_TRIGGER(condition bool, expression string) *MathMLMROOTElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *MathMLMROOTElement) Z_TRIGGERRemove() *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *MathMLMROOTElement) Z_REQ_METHOD(c MathMLMrootZReqMethodChoice) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type MathMLMrootZReqMethodChoice string

const (
	// default GET
	MathMLMrootZReqMethod_empty MathMLMrootZReqMethodChoice = ""
	// GET
	MathMLMrootZReqMethod_get MathMLMrootZReqMethodChoice = "get"
	// POST
	MathMLMrootZReqMethod_post MathMLMrootZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *MathMLMROOTElement) Z_REQ_METHODRemove(c MathMLMrootZReqMethodChoice) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *MathMLMROOTElement) Z_REQ_STRATEGY(expression string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROOTElement) IfZ_REQ_STRATEGY(condition bool, expression string) *MathMLMROOTElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *MathMLMROOTElement) Z_REQ_STRATEGYRemove() *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *MathMLMROOTElement) Z_REQ_HISTORY(expression string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROOTElement) IfZ_REQ_HISTORY(condition bool, expression string) *MathMLMROOTElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *MathMLMROOTElement) Z_REQ_HISTORYRemove() *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *MathMLMROOTElement) Z_DATA(expression string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROOTElement) IfZ_DATA(condition bool, expression string) *MathMLMROOTElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *MathMLMROOTElement) Z_DATARemove() *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *MathMLMROOTElement) Z_JSON(expression string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROOTElement) IfZ_JSON(condition bool, expression string) *MathMLMROOTElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *MathMLMROOTElement) Z_JSONRemove() *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *MathMLMROOTElement) Z_REQ_BATCH(expression string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROOTElement) IfZ_REQ_BATCH(condition bool, expression string) *MathMLMROOTElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *MathMLMROOTElement) Z_REQ_BATCHRemove() *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *MathMLMROOTElement) Z_ACTION(expression string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROOTElement) IfZ_ACTION(condition bool, expression string) *MathMLMROOTElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *MathMLMROOTElement) Z_ACTIONRemove() *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *MathMLMROOTElement) Z_REQ_BEFORE(expression string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROOTElement) IfZ_REQ_BEFORE(condition bool, expression string) *MathMLMROOTElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *MathMLMROOTElement) Z_REQ_BEFORERemove() *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *MathMLMROOTElement) Z_REQ_AFTER(expression string) *MathMLMROOTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMROOTElement) IfZ_REQ_AFTER(condition bool, expression string) *MathMLMROOTElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *MathMLMROOTElement) Z_REQ_AFTERRemove() *MathMLMROOTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
