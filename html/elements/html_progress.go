package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The HTML <progress> element displays an indicator showing the completion
// progress of a task, typically displayed as a progress bar.
type PROGRESSElement struct {
	*Element
}

// Create a new PROGRESSElement element.
// This will create a new element with the tag
// "progress" during rendering.
func PROGRESS(children ...ElementRenderer) *PROGRESSElement {
	e := NewElement("progress", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &PROGRESSElement{Element: e}
}

func (e *PROGRESSElement) Children(children ...ElementRenderer) *PROGRESSElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *PROGRESSElement) IfChildren(condition bool, children ...ElementRenderer) *PROGRESSElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *PROGRESSElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *PROGRESSElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *PROGRESSElement) Attr(name string, value ...string) *PROGRESSElement {
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

func (e *PROGRESSElement) Attrs(attrs ...string) *PROGRESSElement {
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

func (e *PROGRESSElement) AttrsMap(attrs map[string]string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *PROGRESSElement) Text(text string) *PROGRESSElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *PROGRESSElement) TextF(format string, args ...any) *PROGRESSElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *PROGRESSElement) IfText(condition bool, text string) *PROGRESSElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *PROGRESSElement) IfTextF(condition bool, format string, args ...any) *PROGRESSElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *PROGRESSElement) Escaped(text string) *PROGRESSElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *PROGRESSElement) IfEscaped(condition bool, text string) *PROGRESSElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *PROGRESSElement) EscapedF(format string, args ...any) *PROGRESSElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *PROGRESSElement) IfEscapedF(condition bool, format string, args ...any) *PROGRESSElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *PROGRESSElement) CustomData(key, value string) *PROGRESSElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *PROGRESSElement) IfCustomData(condition bool, key, value string) *PROGRESSElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *PROGRESSElement) CustomDataF(key, format string, args ...any) *PROGRESSElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *PROGRESSElement) IfCustomDataF(condition bool, key, format string, args ...any) *PROGRESSElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *PROGRESSElement) CustomDataRemove(key string) *PROGRESSElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// Upper bound of range.
func (e *PROGRESSElement) MAX(f float64) *PROGRESSElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("max", f)
	return e
}

func (e *PROGRESSElement) IfMAX(condition bool, f float64) *PROGRESSElement {
	if condition {
		e.MAX(f)
	}
	return e
}

// Current value of the element.
func (e *PROGRESSElement) VALUE(f float64) *PROGRESSElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("value", f)
	return e
}

func (e *PROGRESSElement) IfVALUE(condition bool, f float64) *PROGRESSElement {
	if condition {
		e.VALUE(f)
	}
	return e
}

// The accesskey global attribute provides a hint for generating a keyboard
// shortcut for the current element
// The attribute value must consist of a single printable character (which
// includes accented and other characters that can be generated by the keyboard).
func (e *PROGRESSElement) ACCESSKEY(r rune) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("accesskey", string(r))
	return e
}

func (e *PROGRESSElement) IfACCESSKEY(condition bool, r rune) *PROGRESSElement {
	if condition {
		e.ACCESSKEY(r)
	}
	return e
}

// Remove the attribute ACCESSKEY from the element.
func (e *PROGRESSElement) ACCESSKEYRemove() *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("accesskey")
	return e
}

// The autocapitalize global attribute is an enumerated attribute that controls
// whether and how text input is automatically capitalized as it is entered/edited
// by the user
// autocapitalize can be set on <input> and <textarea> elements, and on their
// containing <form> elements
// When autocapitalize is set on a <form> element, it sets the autocapitalize
// behavior for all contained <input>s and <textarea>s, overriding any
// autocapitalize values set on contained elements
// autocapitalize has no effect on the url, email, or password <input> types,
// where autocapitalization is never enabled
// Where autocapitalize is not specified, the adopted default behavior varies
// between browsers
// For example: Chrome and Safari default to on/sentences Firefox defaults to
// off/none.
func (e *PROGRESSElement) AUTOCAPITALIZE(c ProgressAutocapitalizeChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("autocapitalize", string(c))
	return e
}

type ProgressAutocapitalizeChoice string

const (
	// Do not automatically capitalize any text.
	ProgressAutocapitalize_off ProgressAutocapitalizeChoice = "off"
	// Do not automatically capitalize any text.
	ProgressAutocapitalize_none ProgressAutocapitalizeChoice = "none"
	// Automatically capitalize the first character of each sentence.
	ProgressAutocapitalize_sentences ProgressAutocapitalizeChoice = "sentences"
	// Automatically capitalize the first character of each sentence.
	ProgressAutocapitalize_on ProgressAutocapitalizeChoice = "on"
	// Automatically capitalize the first character of each word.
	ProgressAutocapitalize_words ProgressAutocapitalizeChoice = "words"
	// Automatically capitalize all characters.
	ProgressAutocapitalize_characters ProgressAutocapitalizeChoice = "characters"
)

// Remove the attribute AUTOCAPITALIZE from the element.
func (e *PROGRESSElement) AUTOCAPITALIZERemove(c ProgressAutocapitalizeChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("autocapitalize")
	return e
}

// The autofocus global attribute is a Boolean attribute indicating that an
// element should be focused on page load, or when the <dialog> that it is part of
// is displayed.
// 		Accessibility concerns Automatically focusing a form control can confuse
// visually-impaired people using screen-reading technology and people with
// cognitive impairments
// When autofocus is assigned, screen-readers "teleport" their user to the form
// control without warning them beforehand.
// 		Use careful consideration for accessibility when applying the autofocus
// attribute
// Automatically focusing on a control can cause the page to scroll on load
// The focus can also cause dynamic keyboards to display on some touch devices
// While a screen reader will announce the label of the form control receiving
// focus, the screen reader will not announce anything before the label, and the
// sighted user on a small device will equally miss the context created by the
// preceding content.
func (e *PROGRESSElement) AUTOFOCUS() *PROGRESSElement {
	e.AUTOFOCUSSet(true)
	return e
}

func (e *PROGRESSElement) IfAUTOFOCUS(condition bool) *PROGRESSElement {
	if condition {
		e.AUTOFOCUSSet(true)
	}
	return e
}

// Set the attribute AUTOFOCUS to the value b explicitly.
func (e *PROGRESSElement) AUTOFOCUSSet(b bool) *PROGRESSElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = zarray.NewSortMap[string, bool]()
	}
	e.BoolAttributes.Set("autofocus", b)
	return e
}

func (e *PROGRESSElement) IfSetAUTOFOCUS(condition bool, b bool) *PROGRESSElement {
	if condition {
		e.AUTOFOCUSSet(b)
	}
	return e
}

// Remove the attribute AUTOFOCUS from the element.
func (e *PROGRESSElement) AUTOFOCUSRemove(b bool) *PROGRESSElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Delete("autofocus")
	return e
}

// The class global attribute is a space-separated list of the case-sensitive
// classes of the element
// Classes allow CSS and JavaScript to select and access specific elements via the
// class selectors or functions like the DOM method
// document.getElementsByClassName.
func (e *PROGRESSElement) CLASS(s ...string) *PROGRESSElement {
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

func (e *PROGRESSElement) IfCLASS(condition bool, s ...string) *PROGRESSElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *PROGRESSElement) CLASSRemove(s ...string) *PROGRESSElement {
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

// The contenteditable global attribute is an enumerated attribute indicating if
// the element should be editable by the user
// If so, the browser modifies its widget to allow editing.
func (e *PROGRESSElement) CONTENTEDITABLE(c ProgressContenteditableChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("contenteditable", string(c))
	return e
}

type ProgressContenteditableChoice string

const (
	// The element is editable.
	ProgressContenteditable_empty ProgressContenteditableChoice = ""
	// The element is editable.
	ProgressContenteditable_true ProgressContenteditableChoice = "true"
	// The element is not editable.
	ProgressContenteditable_false ProgressContenteditableChoice = "false"
	// which indicates that the element's raw text is editable, but rich text
	// formatting is disabled.
	ProgressContenteditable_plaintext_only ProgressContenteditableChoice = "plaintext-only"
)

// Remove the attribute CONTENTEDITABLE from the element.
func (e *PROGRESSElement) CONTENTEDITABLERemove(c ProgressContenteditableChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("contenteditable")
	return e
}

// The dir global attribute is an enumerated attribute that indicates the
// directionality of the element's text
// Note: This attribute is mandatory for the <bdo> element where it has a
// different semantic meaning
// This attribute is not inherited by the <bdi> element
// If not set, its value is auto
// This attribute can be overridden by the CSS properties direction and
// unicode-bidi, if a CSS page is active and the element supports these properties
// As the directionality of the text is semantically related to its content and
// not to its presentation, it is recommended that web developers use this
// attribute instead of the related CSS properties when possible
// That way, the text will display correctly even on a browser that doesn't
// support CSS or has the CSS deactivated
// The auto value should be used for data with an unknown directionality, like
// data coming from user input, eventually stored in a database.
func (e *PROGRESSElement) DIR(c ProgressDirChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type ProgressDirChoice string

const (
	// which means left to right and is to be used for languages that are written from
	// the left to the right (like English);
	ProgressDir_ltr ProgressDirChoice = "ltr"
	// which means right to left and is to be used for languages that are written from
	// the right to the left (like Arabic);
	ProgressDir_rtl ProgressDirChoice = "rtl"
	// which lets the user agent decide
	// It uses a basic algorithm as it parses the characters inside the element until
	// it finds a character with a strong directionality, then it applies that
	// directionality to the whole element.
	ProgressDir_auto ProgressDirChoice = "auto"
)

// Remove the attribute DIR from the element.
func (e *PROGRESSElement) DIRRemove(c ProgressDirChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("dir")
	return e
}

// The draggable global attribute is an enumerated attribute that indicates
// whether the element can be dragged, either with native browser behavior or the
// HTML Drag and Drop API.
func (e *PROGRESSElement) DRAGGABLE(c ProgressDraggableChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("draggable", string(c))
	return e
}

type ProgressDraggableChoice string

const (
	// The element is draggable.
	ProgressDraggable_true ProgressDraggableChoice = "true"
	// The element is not draggable.
	ProgressDraggable_false ProgressDraggableChoice = "false"
	// drag behavior is the default browser behavior: only text selections, images,
	// and links can be dragged
	// For other elements, the event ondragstart must be set for drag and drop to work
	ProgressDraggable_empty ProgressDraggableChoice = ""
	// drag behavior is the default browser behavior: only text selections, images,
	// and links can be dragged
	// For other elements, the event ondragstart must be set for drag and drop to work
	ProgressDraggable_auto ProgressDraggableChoice = "auto"
)

// Remove the attribute DRAGGABLE from the element.
func (e *PROGRESSElement) DRAGGABLERemove(c ProgressDraggableChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("draggable")
	return e
}

// The enterkeyhint global attribute is an enumerated attribute defining what
// action label (or icon) to present for the enter key on virtual keyboards.
func (e *PROGRESSElement) ENTERKEYHINT(c ProgressEnterkeyhintChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("enterkeyhint", string(c))
	return e
}

type ProgressEnterkeyhintChoice string

const (
	// Typically inserting a new line.
	ProgressEnterkeyhint_enter ProgressEnterkeyhintChoice = "enter"
	// Typically meaning there is nothing more to input and the input method editor
	// (IME) will be closed.
	ProgressEnterkeyhint_done ProgressEnterkeyhintChoice = "done"
	// Typically meaning to take the user to the target of the text they typed.
	ProgressEnterkeyhint_go ProgressEnterkeyhintChoice = "go"
	// Typically meaning to take the user to the next field that will accept text.
	ProgressEnterkeyhint_next ProgressEnterkeyhintChoice = "next"
	// Typically meaning to take the user to the previous field that will accept text.
	ProgressEnterkeyhint_previous ProgressEnterkeyhintChoice = "previous"
	// Typically taking the user to the results of searching for the text they have
	// typed.
	ProgressEnterkeyhint_search ProgressEnterkeyhintChoice = "search"
	// Typically delivering the text to its target.
	ProgressEnterkeyhint_send ProgressEnterkeyhintChoice = "send"
)

// Remove the attribute ENTERKEYHINT from the element.
func (e *PROGRESSElement) ENTERKEYHINTRemove(c ProgressEnterkeyhintChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("enterkeyhint")
	return e
}

// The exportparts global attribute allows you to select and style elements
// existing in nested shadow trees, by exporting their part names
// The shadow tree is an isolated structure where identifiers, classes, and styles
// cannot be reached by selectors or queries belonging to a regular DOM
// To apply a style to an element living in a shadow tree, by CSS rule created
// outside of it, part global attribute has to be used
// It has to be assigned to an element present in Shadow Tree, and its value
// should be some identifier
// Rules present outside of the shadow tree, must use the ::part pseudo-element,
// containing the same identifier as the argument
// The global attribute part makes the element visible on just a single level of
// depth
// When the shadow tree is nested, parts will be visible only to the parent of the
// shadow tree but not to its ancestor
// Exporting parts further down is exactly what exportparts attribute is for
// Attribute exportparts must be placed on a shadow Host, which is the element to
// which the shadow tree is attached
// The value of the attribute should be a comma-separated list of part names
// present in the shadow tree and which should be made available via a DOM outside
// of the current structure.
func (e *PROGRESSElement) EXPORTPARTS(s ...string) *PROGRESSElement {
	if e.DelimitedStrings == nil {
		e.DelimitedStrings = zarray.NewSortMap[string, *DelimitedBuilder[string]]()
	}
	ds, ok := e.DelimitedStrings.Get("exportparts")
	if !ok {
		ds = NewDelimitedBuilder[string](",")
		e.DelimitedStrings.Set("exportparts", ds)
	}
	ds.Add(s...)
	return e
}

func (e *PROGRESSElement) IfEXPORTPARTS(condition bool, s ...string) *PROGRESSElement {
	if condition {
		e.EXPORTPARTS(s...)
	}
	return e
}

// Remove the attribute EXPORTPARTS from the element.
func (e *PROGRESSElement) EXPORTPARTSRemove(s ...string) *PROGRESSElement {
	if e.DelimitedStrings == nil {
		return e
	}
	ds, ok := e.DelimitedStrings.Get("exportparts")
	if !ok {
		return e
	}
	ds.Remove(s...)
	return e
}

// The hidden global attribute is a Boolean attribute indicating that the element
// is not yet, or is no longer, relevant
// For example, it can be used to hide elements of the page that can't be used
// until the login process has been completed
// Note that browsers typically implement hidden until found using
// content-visibility: hidden
// This means that unlike elements in the hidden state, elements in the hidden
// until found state will have generated boxes, meaning that: the element will
// participate in page layout margin, borders, padding, and background for the
// element will be rendered
// Also, the element needs to be affected by layout containment in order to be
// revealed
// This means that if the element in the hidden until found state has a display
// value of none, contents, or inline, then the element will not be revealed by
// find in page or fragment navigation.
func (e *PROGRESSElement) HIDDEN(c ProgressHiddenChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("hidden", string(c))
	return e
}

type ProgressHiddenChoice string

const (
	// set the element to the hidden state
	// Additionally, invalid values set the element to the hidden state.
	ProgressHidden_empty ProgressHiddenChoice = ""
	// set the element to the hidden state
	// Additionally, invalid values set the element to the hidden state.
	ProgressHidden_hidden ProgressHiddenChoice = "hidden"
	// the element is hidden but its content will be accessible to the browser's "find
	// in page" feature or to fragment navigation
	// When these features cause a scroll to an element in a hidden until found
	// subtree, the browser will fire a beforematch event on the hidden element remove
	// the hidden attribute from the element scroll to the element
	//
	ProgressHidden_until_found ProgressHiddenChoice = "until-found"
)

// Remove the attribute HIDDEN from the element.
func (e *PROGRESSElement) HIDDENRemove(c ProgressHiddenChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("hidden")
	return e
}

// The id global attribute defines a unique identifier (ID) which must be unique
// in the whole document
// Its purpose is to identify the element when linking (using a fragment
// identifier), scripting, or styling (with CSS).
func (e *PROGRESSElement) ID(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *PROGRESSElement) IDF(format string, args ...any) *PROGRESSElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *PROGRESSElement) IfID(condition bool, s string) *PROGRESSElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *PROGRESSElement) IfIDF(condition bool, format string, args ...any) *PROGRESSElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *PROGRESSElement) IDRemove(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *PROGRESSElement) IDRemoveF(format string, args ...any) *PROGRESSElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// The inert global attribute is a Boolean attribute indicating that the browser
// will ignore the element
// With the inert attribute, all of the element's flat tree descendants (such as
// modal <dialog>s) that don't otherwise escape inertness are ignored
// The inert attribute also makes the browser ignore input events sent by the
// user, including focus-related events and events from assistive technologies
// Specifically, inert does the following: Prevents the click event from being
// fired when the user clicks on the element
// Prevents the focus event from being raised by preventing the element from
// gaining focus
// Hides the element and its content from assistive technologies by excluding them
// from the accessibility tree.
func (e *PROGRESSElement) INERT() *PROGRESSElement {
	e.INERTSet(true)
	return e
}

func (e *PROGRESSElement) IfINERT(condition bool) *PROGRESSElement {
	if condition {
		e.INERTSet(true)
	}
	return e
}

// Set the attribute INERT to the value b explicitly.
func (e *PROGRESSElement) INERTSet(b bool) *PROGRESSElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = zarray.NewSortMap[string, bool]()
	}
	e.BoolAttributes.Set("inert", b)
	return e
}

func (e *PROGRESSElement) IfSetINERT(condition bool, b bool) *PROGRESSElement {
	if condition {
		e.INERTSet(b)
	}
	return e
}

// Remove the attribute INERT from the element.
func (e *PROGRESSElement) INERTRemove(b bool) *PROGRESSElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Delete("inert")
	return e
}

// The inputmode global attribute is an enumerated attribute that hints at the
// type of data that might be entered by the user while editing the element or its
// contents
// This allows a browser to display an appropriate virtual keyboard
// It is used primarily on <input> elements, but is usable on any element in
// contenteditable mode
// It's important to understand that the inputmode attribute doesn't cause any
// validity requirements to be enforced on input
// To require that input conforms to a particular data type, choose an appropriate
// <input> element type
// For specific guidance on choosing <input> types, see the Values section.
func (e *PROGRESSElement) INPUTMODE(c ProgressInputmodeChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("inputmode", string(c))
	return e
}

type ProgressInputmodeChoice string

const (
	// No virtual keyboard
	// For when the page implements its own keyboard input control.
	ProgressInputmode_none ProgressInputmodeChoice = "none"
	// Standard input keyboard for the user's current locale.
	ProgressInputmode_empty ProgressInputmodeChoice = ""
	// Standard input keyboard for the user's current locale.
	ProgressInputmode_text ProgressInputmodeChoice = "text"
	// Fractional numeric input keyboard containing the digits and decimal separator
	// for the user's locale (typically
	// or ,)
	// Devices may or may not show a minus key (-).
	ProgressInputmode_decimal ProgressInputmodeChoice = "decimal"
	// Numeric input keyboard, but only requires the digits 0–9
	// Devices may or may not show a minus key.
	ProgressInputmode_numeric ProgressInputmodeChoice = "numeric"
	// A telephone keypad input, including the digits 0–9, the asterisk (*), and the
	// pound (#) key
	// Inputs that *require* a telephone number should typically use <input
	// type="tel"> instead.
	ProgressInputmode_tel ProgressInputmodeChoice = "tel"
	// A virtual keyboard optimized for search input
	// For instance, the return/submit key may be labeled "Search", along with
	// possible other optimizations
	// Inputs that require a search query should typically use <input type="search">
	// instead.
	ProgressInputmode_search ProgressInputmodeChoice = "search"
	// A virtual keyboard optimized for entering email addresses
	// Typically includes the @character as well as other optimizations
	// Inputs that require email addresses should typically use <input type="email">
	// instead.
	ProgressInputmode_email ProgressInputmodeChoice = "email"
	// A keypad optimized for entering URLs
	// This may have the / key more prominent, for example
	// Enhanced features could include history access and so on
	// Inputs that require a URL should typically use <input type="url"> instead.
	ProgressInputmode_url ProgressInputmodeChoice = "url"
)

// Remove the attribute INPUTMODE from the element.
func (e *PROGRESSElement) INPUTMODERemove(c ProgressInputmodeChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("inputmode")
	return e
}

// The is global attribute allows you to specify that a standard HTML element
// should behave like a defined custom built-in element (see Using custom elements
// for more details)
// This attribute can only be used if the specified custom element name has been
// successfully defined in the current document, and extends the element type it
// is being applied to.
func (e *PROGRESSElement) IS(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("is", s)
	return e
}

func (e *PROGRESSElement) ISF(format string, args ...any) *PROGRESSElement {
	return e.IS(fmt.Sprintf(format, args...))
}

func (e *PROGRESSElement) IfIS(condition bool, s string) *PROGRESSElement {
	if condition {
		e.IS(s)
	}
	return e
}

func (e *PROGRESSElement) IfISF(condition bool, format string, args ...any) *PROGRESSElement {
	if condition {
		e.IS(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IS from the element.
func (e *PROGRESSElement) ISRemove(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("is")
	return e
}

func (e *PROGRESSElement) ISRemoveF(format string, args ...any) *PROGRESSElement {
	return e.ISRemove(fmt.Sprintf(format, args...))
}

// The itemid global attribute provides microdata in the form of a unique, global
// identifier of an item.
//
// 		An itemid attribute can only be specified for an element that has both
// itemscope and itemtype attributes
// Also, itemid can only be specified on elements that possess an itemscope
// attribute whose corresponding itemtype refers to or defines a vocabulary that
// supports global identifiers
// The exact meaning of an itemtype's global identifier is provided by the
// definition of that identifier within the specified vocabulary
// The vocabulary defines whether several items with the same global identifier
// can coexist and, if so, how items with the same identifier are handled.
func (e *PROGRESSElement) ITEMID(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("itemid", s)
	return e
}

func (e *PROGRESSElement) ITEMIDF(format string, args ...any) *PROGRESSElement {
	return e.ITEMID(fmt.Sprintf(format, args...))
}

func (e *PROGRESSElement) IfITEMID(condition bool, s string) *PROGRESSElement {
	if condition {
		e.ITEMID(s)
	}
	return e
}

func (e *PROGRESSElement) IfITEMIDF(condition bool, format string, args ...any) *PROGRESSElement {
	if condition {
		e.ITEMID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ITEMID from the element.
func (e *PROGRESSElement) ITEMIDRemove(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("itemid")
	return e
}

func (e *PROGRESSElement) ITEMIDRemoveF(format string, args ...any) *PROGRESSElement {
	return e.ITEMIDRemove(fmt.Sprintf(format, args...))
}

// The itemprop global attribute is used to add properties to an item
// Every HTML element can have an itemprop attribute specified, and an itemprop
// consists of a name-value pair
// Each name-value pair is called a property, and a group of one or more
// properties forms an item
// Property values are either a string or a URL and can be associated with a very
// wide range of elements including <audio>, <embed>, <iframe>, <img>, <link>,
// <object>, <source>, <track>, and <video>.
func (e *PROGRESSElement) ITEMPROP(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("itemprop", s)
	return e
}

func (e *PROGRESSElement) ITEMPROPF(format string, args ...any) *PROGRESSElement {
	return e.ITEMPROP(fmt.Sprintf(format, args...))
}

func (e *PROGRESSElement) IfITEMPROP(condition bool, s string) *PROGRESSElement {
	if condition {
		e.ITEMPROP(s)
	}
	return e
}

func (e *PROGRESSElement) IfITEMPROPF(condition bool, format string, args ...any) *PROGRESSElement {
	if condition {
		e.ITEMPROP(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ITEMPROP from the element.
func (e *PROGRESSElement) ITEMPROPRemove(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("itemprop")
	return e
}

func (e *PROGRESSElement) ITEMPROPRemoveF(format string, args ...any) *PROGRESSElement {
	return e.ITEMPROPRemove(fmt.Sprintf(format, args...))
}

// Properties that are not descendants of an element with the itemscope attribute
// can be associated with an item using the global attribute itemref
// itemref provides a list of element IDs (not itemids) elsewhere in the document,
// with additional properties The itemref attribute can only be specified on
// elements that have an itemscope attribute specified.
func (e *PROGRESSElement) ITEMREF(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("itemref", s)
	return e
}

func (e *PROGRESSElement) ITEMREFF(format string, args ...any) *PROGRESSElement {
	return e.ITEMREF(fmt.Sprintf(format, args...))
}

func (e *PROGRESSElement) IfITEMREF(condition bool, s string) *PROGRESSElement {
	if condition {
		e.ITEMREF(s)
	}
	return e
}

func (e *PROGRESSElement) IfITEMREFF(condition bool, format string, args ...any) *PROGRESSElement {
	if condition {
		e.ITEMREF(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ITEMREF from the element.
func (e *PROGRESSElement) ITEMREFRemove(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("itemref")
	return e
}

func (e *PROGRESSElement) ITEMREFRemoveF(format string, args ...any) *PROGRESSElement {
	return e.ITEMREFRemove(fmt.Sprintf(format, args...))
}

// The itemscope global attribute is used to add an item to a microdata DOM tree
// Every HTML element can have an itemscope attribute specified, and an itemscope
// consists of a name-value pair
// Each name-value pair is called a property, and a group of one or more
// properties forms an item
// Property values are either a string or a URL and can be associated with a very
// wide range of elements including <audio>, <embed>, <iframe>, <img>, <link>,
// <object>, <source>, <track>, and <video>.
func (e *PROGRESSElement) ITEMSCOPE() *PROGRESSElement {
	e.ITEMSCOPESet(true)
	return e
}

func (e *PROGRESSElement) IfITEMSCOPE(condition bool) *PROGRESSElement {
	if condition {
		e.ITEMSCOPESet(true)
	}
	return e
}

// Set the attribute ITEMSCOPE to the value b explicitly.
func (e *PROGRESSElement) ITEMSCOPESet(b bool) *PROGRESSElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = zarray.NewSortMap[string, bool]()
	}
	e.BoolAttributes.Set("itemscope", b)
	return e
}

func (e *PROGRESSElement) IfSetITEMSCOPE(condition bool, b bool) *PROGRESSElement {
	if condition {
		e.ITEMSCOPESet(b)
	}
	return e
}

// Remove the attribute ITEMSCOPE from the element.
func (e *PROGRESSElement) ITEMSCOPERemove(b bool) *PROGRESSElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Delete("itemscope")
	return e
}

// The itemtype global attribute is used to add types to an item
// Every HTML element can have an itemtype attribute specified, and an itemtype
// consists of a name-value pair
// Each name-value pair is called a property, and a group of one or more
// properties forms an item
// Property values are either a string or a URL and can be associated with a very
// wide range of elements including <audio>, <embed>, <iframe>, <img>, <link>,
// <object>, <source>, <track>, and <video>.
func (e *PROGRESSElement) ITEMTYPE(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("itemtype", s)
	return e
}

func (e *PROGRESSElement) ITEMTYPEF(format string, args ...any) *PROGRESSElement {
	return e.ITEMTYPE(fmt.Sprintf(format, args...))
}

func (e *PROGRESSElement) IfITEMTYPE(condition bool, s string) *PROGRESSElement {
	if condition {
		e.ITEMTYPE(s)
	}
	return e
}

func (e *PROGRESSElement) IfITEMTYPEF(condition bool, format string, args ...any) *PROGRESSElement {
	if condition {
		e.ITEMTYPE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ITEMTYPE from the element.
func (e *PROGRESSElement) ITEMTYPERemove(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("itemtype")
	return e
}

func (e *PROGRESSElement) ITEMTYPERemoveF(format string, args ...any) *PROGRESSElement {
	return e.ITEMTYPERemove(fmt.Sprintf(format, args...))
}

// The lang global attribute helps define the language of an element: the language
// that non-editable elements are written in or the language that editable
// elements should be written in by the user
// The tag contains one single entry value in the format defines in the Tags for
// Identifying Languages (BCP47) IETF document
// xml:lang has priority over it.
func (e *PROGRESSElement) LANG(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("lang", s)
	return e
}

func (e *PROGRESSElement) LANGF(format string, args ...any) *PROGRESSElement {
	return e.LANG(fmt.Sprintf(format, args...))
}

func (e *PROGRESSElement) IfLANG(condition bool, s string) *PROGRESSElement {
	if condition {
		e.LANG(s)
	}
	return e
}

func (e *PROGRESSElement) IfLANGF(condition bool, format string, args ...any) *PROGRESSElement {
	if condition {
		e.LANG(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute LANG from the element.
func (e *PROGRESSElement) LANGRemove(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("lang")
	return e
}

func (e *PROGRESSElement) LANGRemoveF(format string, args ...any) *PROGRESSElement {
	return e.LANGRemove(fmt.Sprintf(format, args...))
}

// The nonce global attribute is a unique identifier used to declare inline
// scripts and style elements to be used in a specific document
// It is a cryptographic nonce (number used once) that is used by Content Security
// Policy to determine whether or not a given inline script is allowed to execute.
func (e *PROGRESSElement) NONCE(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *PROGRESSElement) NONCEF(format string, args ...any) *PROGRESSElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *PROGRESSElement) IfNONCE(condition bool, s string) *PROGRESSElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *PROGRESSElement) IfNONCEF(condition bool, format string, args ...any) *PROGRESSElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *PROGRESSElement) NONCERemove(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("nonce")
	return e
}

func (e *PROGRESSElement) NONCERemoveF(format string, args ...any) *PROGRESSElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// The part global attribute contains a space-separated list of the part names of
// the element
// Part names allows CSS to select and style specific elements in a shadow tree
// via the ::part pseudo-element.
func (e *PROGRESSElement) PART(s ...string) *PROGRESSElement {
	if e.DelimitedStrings == nil {
		e.DelimitedStrings = zarray.NewSortMap[string, *DelimitedBuilder[string]]()
	}
	ds, ok := e.DelimitedStrings.Get("part")
	if !ok {
		ds = NewDelimitedBuilder[string](" ")
		e.DelimitedStrings.Set("part", ds)
	}
	ds.Add(s...)
	return e
}

func (e *PROGRESSElement) IfPART(condition bool, s ...string) *PROGRESSElement {
	if condition {
		e.PART(s...)
	}
	return e
}

// Remove the attribute PART from the element.
func (e *PROGRESSElement) PARTRemove(s ...string) *PROGRESSElement {
	if e.DelimitedStrings == nil {
		return e
	}
	ds, ok := e.DelimitedStrings.Get("part")
	if !ok {
		return e
	}
	ds.Remove(s...)
	return e
}

// The popover global attribute is used to designate an element as a popover
// element
// Popover elements are hidden via display: none until opened via an
// invoking/control element (i.e
// a <button> or <input type="button"> with a popovertarget attribute) or a
// HTMLElement.showPopover() call
// When open, popover elements will appear above all other elements in the top
// layer, and won't be influenced by parent elements' position or overflow
// styling.
func (e *PROGRESSElement) POPVER(c ProgressPopverChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("popver", string(c))
	return e
}

type ProgressPopverChoice string

const (
	// Popovers that have the auto state can be "light dismissed" by selecting outside
	// the popover area, and generally only allow one popover to be displayed
	// on-screen at a time.
	ProgressPopver_auto ProgressPopverChoice = "auto"
	// Popovers that have the auto state can be "light dismissed" by selecting outside
	// the popover area, and generally only allow one popover to be displayed
	// on-screen at a time.
	ProgressPopver_empty ProgressPopverChoice = ""
	// manual popovers must always be explicitly hidden, but allow for use cases such
	// as nested popovers in menus.
	ProgressPopver_manual ProgressPopverChoice = "manual"
)

// Remove the attribute POPVER from the element.
func (e *PROGRESSElement) POPVERRemove(c ProgressPopverChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("popver")
	return e
}

// The role global attribute is used to define the purpose or state of an element
// to the browser, in order to facilitate assistive technology such as screen
// readers
// It is a simple string value that can be used to describe the role of an
// element.
func (e *PROGRESSElement) ROLE(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("role", s)
	return e
}

func (e *PROGRESSElement) ROLEF(format string, args ...any) *PROGRESSElement {
	return e.ROLE(fmt.Sprintf(format, args...))
}

func (e *PROGRESSElement) IfROLE(condition bool, s string) *PROGRESSElement {
	if condition {
		e.ROLE(s)
	}
	return e
}

func (e *PROGRESSElement) IfROLEF(condition bool, format string, args ...any) *PROGRESSElement {
	if condition {
		e.ROLE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ROLE from the element.
func (e *PROGRESSElement) ROLERemove(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("role")
	return e
}

func (e *PROGRESSElement) ROLERemoveF(format string, args ...any) *PROGRESSElement {
	return e.ROLERemove(fmt.Sprintf(format, args...))
}

// The slot global attribute assigns a slot in a shadow DOM shadow tree to an
// element: An element with a slot attribute is assigned to the slot created by
// the <slot> element whose name attribute's value matches that slot attribute's
// value.
func (e *PROGRESSElement) SLOT(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("slot", s)
	return e
}

func (e *PROGRESSElement) SLOTF(format string, args ...any) *PROGRESSElement {
	return e.SLOT(fmt.Sprintf(format, args...))
}

func (e *PROGRESSElement) IfSLOT(condition bool, s string) *PROGRESSElement {
	if condition {
		e.SLOT(s)
	}
	return e
}

func (e *PROGRESSElement) IfSLOTF(condition bool, format string, args ...any) *PROGRESSElement {
	if condition {
		e.SLOT(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute SLOT from the element.
func (e *PROGRESSElement) SLOTRemove(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("slot")
	return e
}

func (e *PROGRESSElement) SLOTRemoveF(format string, args ...any) *PROGRESSElement {
	return e.SLOTRemove(fmt.Sprintf(format, args...))
}

// The spellcheck global attribute is an enumerated attribute that defines whether
// the element may be checked for spelling errors
// If this attribute is not set, its default value is element-type and
// browser-defined
// This default value may also be inherited, which means that the element content
// will be checked for spelling errors only if its nearest ancestor has a
// spellcheck state of true
// Security and privacy concerns Using spellchecking can have consequences for
// users' security and privacy
// The specification does not regulate how spellchecking is done and the content
// of the element may be sent to a third party for spellchecking results (see
// enhanced spellchecking and "spell-jacking")
// You should consider setting spellcheck to false for elements that can contain
// sensitive information.
func (e *PROGRESSElement) SPELLCHECK(c ProgressSpellcheckChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("spellcheck", string(c))
	return e
}

type ProgressSpellcheckChoice string

const (
	// The element will be checked for spelling errors.
	ProgressSpellcheck_empty ProgressSpellcheckChoice = ""
	// The element will be checked for spelling errors.
	ProgressSpellcheck_true ProgressSpellcheckChoice = "true"
	// The element will not be checked for spelling errors.
	ProgressSpellcheck_false ProgressSpellcheckChoice = "false"
)

// Remove the attribute SPELLCHECK from the element.
func (e *PROGRESSElement) SPELLCHECKRemove(c ProgressSpellcheckChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("spellcheck")
	return e
}

// The style global attribute is used to add styles to an element, such as color,
// font, size, and more
// Styles are written in CSS.
func (e *PROGRESSElement) STYLEF(k string, format string, args ...any) *PROGRESSElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *PROGRESSElement) IfSTYLE(condition bool, k string, v string) *PROGRESSElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *PROGRESSElement) STYLE(k string, v string) *PROGRESSElement {
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

func (e *PROGRESSElement) IfSTYLEF(condition bool, k string, format string, args ...any) *PROGRESSElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *PROGRESSElement) STYLEMap(m map[string]string) *PROGRESSElement {
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
func (e *PROGRESSElement) STYLEPairs(pairs ...string) *PROGRESSElement {
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

func (e *PROGRESSElement) IfSTYLEPairs(condition bool, pairs ...string) *PROGRESSElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *PROGRESSElement) STYLERemove(keys ...string) *PROGRESSElement {
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

// The tabindex global attribute indicates if its element can be focused, and
// if/where it participates in sequential keyboard navigation (usually with the
// Tab key, hence the name)
// It accepts an integer as a value, with different results depending on the
// integer's value: a negative value (usually tabindex="-1") means that the
// element should be focusable, but should not be reachable via sequential
// keyboard navigation; a value of 0 (tabindex="0") means that the element should
// be focusable and reachable via sequential keyboard navigation, but its relative
// order is defined by the platform convention; a positive value means should be
// focusable and reachable via sequential keyboard navigation; its relative order
// is defined by the value of the attribute: the sequential follow the increasing
// number of the tabindex
// If several elements share the same tabindex, their relative order follows their
// relative position in the document.
func (e *PROGRESSElement) TABINDEX(i int) *PROGRESSElement {
	if e.IntAttributes == nil {
		e.IntAttributes = zarray.NewSortMap[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *PROGRESSElement) IfTABINDEX(condition bool, i int) *PROGRESSElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *PROGRESSElement) TABINDEXRemove(i int) *PROGRESSElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Delete("tabindex")
	return e
}

// The title global attribute contains text representing advisory information
// related to the element it belongs to
// Such information can typically, but not necessarily, be presented to the user
// as a tooltip
// The main use of the title attribute is to label <iframe> elements for assistive
// technology
// The title attribute may also be used to label controls in data tables
// The title attribute, when added to <link rel="stylesheet">, creates an
// alternate stylesheet
// When defining an alternative style sheet with <link rel="alternate"> the
// attribute is required and must be set to a non-empty string
// If included on the <abbr> opening tag, the title must be a full expansion of
// the abbreviation or acronym
// Instead of using title, when possible, provide an expansion of the abbreviation
// or acronym in plain text on first use, using the <abbr> to mark up the
// abbreviation
// This enables all users know what name or term the abbreviation or acronym
// shortens while providing a hint to user agents on how to announce the content
// While title can be used to provide a programmatically associated label for an
// <input> element, this is not good practice
// Use a <label> instead.
func (e *PROGRESSElement) TITLE(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("title", s)
	return e
}

func (e *PROGRESSElement) TITLEF(format string, args ...any) *PROGRESSElement {
	return e.TITLE(fmt.Sprintf(format, args...))
}

func (e *PROGRESSElement) IfTITLE(condition bool, s string) *PROGRESSElement {
	if condition {
		e.TITLE(s)
	}
	return e
}

func (e *PROGRESSElement) IfTITLEF(condition bool, format string, args ...any) *PROGRESSElement {
	if condition {
		e.TITLE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TITLE from the element.
func (e *PROGRESSElement) TITLERemove(s string) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("title")
	return e
}

func (e *PROGRESSElement) TITLERemoveF(format string, args ...any) *PROGRESSElement {
	return e.TITLERemove(fmt.Sprintf(format, args...))
}

// The translate global attribute is an enumerated attribute that is used to
// specify whether an element's attribute values and the values of its Text node
// children are to be translated when the page is localized, or whether to leave
// them unchanged.
func (e *PROGRESSElement) TRANSLATE(c ProgressTranslateChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("translate", string(c))
	return e
}

type ProgressTranslateChoice string

const (
	// indicates that the element should be translated when the page is localized.
	ProgressTranslate_empty ProgressTranslateChoice = ""
	// indicates that the element should be translated when the page is localized.
	ProgressTranslate_yes ProgressTranslateChoice = "yes"
	// indicates that the element must not be translated when the page is localized.
	ProgressTranslate_no ProgressTranslateChoice = "no"
)

// Remove the attribute TRANSLATE from the element.
func (e *PROGRESSElement) TRANSLATERemove(c ProgressTranslateChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("translate")
	return e
}

// Make a request for an HTML

func (e *PROGRESSElement) Z_REQ(expression string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *PROGRESSElement) IfZ_REQ(condition bool, expression string) *PROGRESSElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *PROGRESSElement) Z_REQRemove() *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *PROGRESSElement) Z_TARGET(expression string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *PROGRESSElement) IfZ_TARGET(condition bool, expression string) *PROGRESSElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *PROGRESSElement) Z_TARGETRemove() *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *PROGRESSElement) Z_REQ_SELECTOR(expression string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *PROGRESSElement) IfZ_REQ_SELECTOR(condition bool, expression string) *PROGRESSElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *PROGRESSElement) Z_REQ_SELECTORRemove() *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *PROGRESSElement) Z_SWAP(expression string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *PROGRESSElement) IfZ_SWAP(condition bool, expression string) *PROGRESSElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *PROGRESSElement) Z_SWAPRemove() *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *PROGRESSElement) Z_SWAP_PUSH(expression string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *PROGRESSElement) IfZ_SWAP_PUSH(condition bool, expression string) *PROGRESSElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *PROGRESSElement) Z_SWAP_PUSHRemove() *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *PROGRESSElement) Z_TRIGGER(expression string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *PROGRESSElement) IfZ_TRIGGER(condition bool, expression string) *PROGRESSElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *PROGRESSElement) Z_TRIGGERRemove() *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *PROGRESSElement) Z_REQ_METHOD(c ProgressZReqMethodChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type ProgressZReqMethodChoice string

const (
	// default GET
	ProgressZReqMethod_empty ProgressZReqMethodChoice = ""
	// GET
	ProgressZReqMethod_get ProgressZReqMethodChoice = "get"
	// POST
	ProgressZReqMethod_post ProgressZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *PROGRESSElement) Z_REQ_METHODRemove(c ProgressZReqMethodChoice) *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *PROGRESSElement) Z_REQ_STRATEGY(expression string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *PROGRESSElement) IfZ_REQ_STRATEGY(condition bool, expression string) *PROGRESSElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *PROGRESSElement) Z_REQ_STRATEGYRemove() *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *PROGRESSElement) Z_REQ_HISTORY(expression string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *PROGRESSElement) IfZ_REQ_HISTORY(condition bool, expression string) *PROGRESSElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *PROGRESSElement) Z_REQ_HISTORYRemove() *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *PROGRESSElement) Z_DATA(expression string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *PROGRESSElement) IfZ_DATA(condition bool, expression string) *PROGRESSElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *PROGRESSElement) Z_DATARemove() *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *PROGRESSElement) Z_JSON(expression string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *PROGRESSElement) IfZ_JSON(condition bool, expression string) *PROGRESSElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *PROGRESSElement) Z_JSONRemove() *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *PROGRESSElement) Z_REQ_BATCH(expression string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *PROGRESSElement) IfZ_REQ_BATCH(condition bool, expression string) *PROGRESSElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *PROGRESSElement) Z_REQ_BATCHRemove() *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *PROGRESSElement) Z_ACTION(expression string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *PROGRESSElement) IfZ_ACTION(condition bool, expression string) *PROGRESSElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *PROGRESSElement) Z_ACTIONRemove() *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *PROGRESSElement) Z_REQ_BEFORE(expression string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *PROGRESSElement) IfZ_REQ_BEFORE(condition bool, expression string) *PROGRESSElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *PROGRESSElement) Z_REQ_BEFORERemove() *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *PROGRESSElement) Z_REQ_AFTER(expression string) *PROGRESSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *PROGRESSElement) IfZ_REQ_AFTER(condition bool, expression string) *PROGRESSElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *PROGRESSElement) Z_REQ_AFTERRemove() *PROGRESSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
