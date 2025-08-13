package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// This element is used to display text.
type MathMLMTEXTElement struct {
	*Element
}

// Create a new MathMLMTEXTElement element.
// This will create a new element with the tag
// "mtext" during rendering.
func MathML_MTEXT(children ...ElementRenderer) *MathMLMTEXTElement {
	e := NewElement("mtext", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMTEXTElement{Element: e}
}

func (e *MathMLMTEXTElement) Children(children ...ElementRenderer) *MathMLMTEXTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMTEXTElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMTEXTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMTEXTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMTEXTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMTEXTElement) Attr(name string, value ...string) *MathMLMTEXTElement {
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

func (e *MathMLMTEXTElement) Attrs(attrs ...string) *MathMLMTEXTElement {
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

func (e *MathMLMTEXTElement) AttrsMap(attrs map[string]string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMTEXTElement) Text(text string) *MathMLMTEXTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMTEXTElement) TextF(format string, args ...any) *MathMLMTEXTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMTEXTElement) IfText(condition bool, text string) *MathMLMTEXTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMTEXTElement) IfTextF(condition bool, format string, args ...any) *MathMLMTEXTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMTEXTElement) Escaped(text string) *MathMLMTEXTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMTEXTElement) IfEscaped(condition bool, text string) *MathMLMTEXTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMTEXTElement) EscapedF(format string, args ...any) *MathMLMTEXTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMTEXTElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMTEXTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMTEXTElement) CustomData(key, value string) *MathMLMTEXTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMTEXTElement) IfCustomData(condition bool, key, value string) *MathMLMTEXTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMTEXTElement) CustomDataF(key, format string, args ...any) *MathMLMTEXTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMTEXTElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMTEXTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMTEXTElement) CustomDataRemove(key string) *MathMLMTEXTElement {
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
func (e *MathMLMTEXTElement) CLASS(s ...string) *MathMLMTEXTElement {
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

func (e *MathMLMTEXTElement) IfCLASS(condition bool, s ...string) *MathMLMTEXTElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMTEXTElement) CLASSRemove(s ...string) *MathMLMTEXTElement {
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
func (e *MathMLMTEXTElement) DIR(c MathMLMtextDirChoice) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMtextDirChoice string

const (
	// left-to-right
	MathMLMtextDir_ltr MathMLMtextDirChoice = "ltr"
	// right-to-left
	MathMLMtextDir_rtl MathMLMtextDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMTEXTElement) DIRRemove(c MathMLMtextDirChoice) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMTEXTElement) DISPLAYSTYLE(c MathMLMtextDisplaystyleChoice) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMtextDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMtextDisplaystyle_true MathMLMtextDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMtextDisplaystyle_false MathMLMtextDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMTEXTElement) DISPLAYSTYLERemove(c MathMLMtextDisplaystyleChoice) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMTEXTElement) ID(s string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMTEXTElement) IDF(format string, args ...any) *MathMLMTEXTElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *MathMLMTEXTElement) IfID(condition bool, s string) *MathMLMTEXTElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *MathMLMTEXTElement) IfIDF(condition bool, format string, args ...any) *MathMLMTEXTElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMTEXTElement) IDRemove(s string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *MathMLMTEXTElement) IDRemoveF(format string, args ...any) *MathMLMTEXTElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMTEXTElement) MATHBACKGROUND(s string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMTEXTElement) MATHBACKGROUNDF(format string, args ...any) *MathMLMTEXTElement {
	return e.MATHBACKGROUND(fmt.Sprintf(format, args...))
}

func (e *MathMLMTEXTElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMTEXTElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

func (e *MathMLMTEXTElement) IfMATHBACKGROUNDF(condition bool, format string, args ...any) *MathMLMTEXTElement {
	if condition {
		e.MATHBACKGROUND(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMTEXTElement) MATHBACKGROUNDRemove(s string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathbackground")
	return e
}

func (e *MathMLMTEXTElement) MATHBACKGROUNDRemoveF(format string, args ...any) *MathMLMTEXTElement {
	return e.MATHBACKGROUNDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMTEXTElement) MATHCOLOR(s string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMTEXTElement) MATHCOLORF(format string, args ...any) *MathMLMTEXTElement {
	return e.MATHCOLOR(fmt.Sprintf(format, args...))
}

func (e *MathMLMTEXTElement) IfMATHCOLOR(condition bool, s string) *MathMLMTEXTElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

func (e *MathMLMTEXTElement) IfMATHCOLORF(condition bool, format string, args ...any) *MathMLMTEXTElement {
	if condition {
		e.MATHCOLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMTEXTElement) MATHCOLORRemove(s string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathcolor")
	return e
}

func (e *MathMLMTEXTElement) MATHCOLORRemoveF(format string, args ...any) *MathMLMTEXTElement {
	return e.MATHCOLORRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMTEXTElement) MATHSIZE_STR(s string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMTEXTElement) MATHSIZE_STRF(format string, args ...any) *MathMLMTEXTElement {
	return e.MATHSIZE_STR(fmt.Sprintf(format, args...))
}

func (e *MathMLMTEXTElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMTEXTElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

func (e *MathMLMTEXTElement) IfMATHSIZE_STRF(condition bool, format string, args ...any) *MathMLMTEXTElement {
	if condition {
		e.MATHSIZE_STR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMTEXTElement) MATHSIZE_STRRemove(s string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("mathsize")
	return e
}

func (e *MathMLMTEXTElement) MATHSIZE_STRRemoveF(format string, args ...any) *MathMLMTEXTElement {
	return e.MATHSIZE_STRRemove(fmt.Sprintf(format, args...))
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLMTEXTElement) NONCE(s string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMTEXTElement) NONCEF(format string, args ...any) *MathMLMTEXTElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *MathMLMTEXTElement) IfNONCE(condition bool, s string) *MathMLMTEXTElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *MathMLMTEXTElement) IfNONCEF(condition bool, format string, args ...any) *MathMLMTEXTElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMTEXTElement) NONCERemove(s string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("nonce")
	return e
}

func (e *MathMLMTEXTElement) NONCERemoveF(format string, args ...any) *MathMLMTEXTElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMTEXTElement) SCRIPTLEVEL(i int) *MathMLMTEXTElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMTEXTElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMTEXTElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMTEXTElement) SCRIPTLEVELRemove(i int) *MathMLMTEXTElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMTEXTElement) STYLEF(k string, format string, args ...any) *MathMLMTEXTElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMTEXTElement) IfSTYLE(condition bool, k string, v string) *MathMLMTEXTElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMTEXTElement) STYLE(k string, v string) *MathMLMTEXTElement {
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

func (e *MathMLMTEXTElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMTEXTElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMTEXTElement) STYLEMap(m map[string]string) *MathMLMTEXTElement {
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
func (e *MathMLMTEXTElement) STYLEPairs(pairs ...string) *MathMLMTEXTElement {
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

func (e *MathMLMTEXTElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMTEXTElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMTEXTElement) STYLERemove(keys ...string) *MathMLMTEXTElement {
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
func (e *MathMLMTEXTElement) TABINDEX(i int) *MathMLMTEXTElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMTEXTElement) IfTABINDEX(condition bool, i int) *MathMLMTEXTElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMTEXTElement) TABINDEXRemove(i int) *MathMLMTEXTElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("tabindex")
	return e
}

// Make a request for an HTML

func (e *MathMLMTEXTElement) Z_REQ(expression string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTEXTElement) IfZ_REQ(condition bool, expression string) *MathMLMTEXTElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *MathMLMTEXTElement) Z_REQRemove() *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *MathMLMTEXTElement) Z_TARGET(expression string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTEXTElement) IfZ_TARGET(condition bool, expression string) *MathMLMTEXTElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *MathMLMTEXTElement) Z_TARGETRemove() *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *MathMLMTEXTElement) Z_REQ_SELECTOR(expression string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTEXTElement) IfZ_REQ_SELECTOR(condition bool, expression string) *MathMLMTEXTElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *MathMLMTEXTElement) Z_REQ_SELECTORRemove() *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *MathMLMTEXTElement) Z_SWAP(expression string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTEXTElement) IfZ_SWAP(condition bool, expression string) *MathMLMTEXTElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *MathMLMTEXTElement) Z_SWAPRemove() *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *MathMLMTEXTElement) Z_SWAP_PUSH(expression string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTEXTElement) IfZ_SWAP_PUSH(condition bool, expression string) *MathMLMTEXTElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *MathMLMTEXTElement) Z_SWAP_PUSHRemove() *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *MathMLMTEXTElement) Z_TRIGGER(expression string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTEXTElement) IfZ_TRIGGER(condition bool, expression string) *MathMLMTEXTElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *MathMLMTEXTElement) Z_TRIGGERRemove() *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *MathMLMTEXTElement) Z_REQ_METHOD(c MathMLMtextZReqMethodChoice) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type MathMLMtextZReqMethodChoice string

const (
	// default GET
	MathMLMtextZReqMethod_empty MathMLMtextZReqMethodChoice = ""
	// GET
	MathMLMtextZReqMethod_get MathMLMtextZReqMethodChoice = "get"
	// POST
	MathMLMtextZReqMethod_post MathMLMtextZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *MathMLMTEXTElement) Z_REQ_METHODRemove(c MathMLMtextZReqMethodChoice) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *MathMLMTEXTElement) Z_REQ_STRATEGY(expression string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTEXTElement) IfZ_REQ_STRATEGY(condition bool, expression string) *MathMLMTEXTElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *MathMLMTEXTElement) Z_REQ_STRATEGYRemove() *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *MathMLMTEXTElement) Z_REQ_HISTORY(expression string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTEXTElement) IfZ_REQ_HISTORY(condition bool, expression string) *MathMLMTEXTElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *MathMLMTEXTElement) Z_REQ_HISTORYRemove() *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *MathMLMTEXTElement) Z_DATA(expression string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTEXTElement) IfZ_DATA(condition bool, expression string) *MathMLMTEXTElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *MathMLMTEXTElement) Z_DATARemove() *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *MathMLMTEXTElement) Z_JSON(expression string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTEXTElement) IfZ_JSON(condition bool, expression string) *MathMLMTEXTElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *MathMLMTEXTElement) Z_JSONRemove() *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *MathMLMTEXTElement) Z_REQ_BATCH(expression string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTEXTElement) IfZ_REQ_BATCH(condition bool, expression string) *MathMLMTEXTElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *MathMLMTEXTElement) Z_REQ_BATCHRemove() *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *MathMLMTEXTElement) Z_ACTION(expression string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTEXTElement) IfZ_ACTION(condition bool, expression string) *MathMLMTEXTElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *MathMLMTEXTElement) Z_ACTIONRemove() *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *MathMLMTEXTElement) Z_REQ_BEFORE(expression string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTEXTElement) IfZ_REQ_BEFORE(condition bool, expression string) *MathMLMTEXTElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *MathMLMTEXTElement) Z_REQ_BEFORERemove() *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *MathMLMTEXTElement) Z_REQ_AFTER(expression string) *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTEXTElement) IfZ_REQ_AFTER(condition bool, expression string) *MathMLMTEXTElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *MathMLMTEXTElement) Z_REQ_AFTERRemove() *MathMLMTEXTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
