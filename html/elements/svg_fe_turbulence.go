package elements

import (
	"fmt"
	"github.com/sohaha/zlsgo/zarray"
)

// The <feTurbulence> SVG filter primitive creates an image using the Perlin
// turbulence function
// It allows the synthesis of artificial textures like clouds or marble.
type SVGFETURBULENCEElement struct {
	*Element
}

// Create a new SVGFETURBULENCEElement element.
// This will create a new element with the tag
// "feTurbulence" during rendering.
func SVG_FETURBULENCE(children ...ElementRenderer) *SVGFETURBULENCEElement {
	e := NewElement("feTurbulence", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFETURBULENCEElement{Element: e}
}

func (e *SVGFETURBULENCEElement) Children(children ...ElementRenderer) *SVGFETURBULENCEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFETURBULENCEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFETURBULENCEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFETURBULENCEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFETURBULENCEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFETURBULENCEElement) Attr(name string, value ...string) *SVGFETURBULENCEElement {
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

func (e *SVGFETURBULENCEElement) Attrs(attrs ...string) *SVGFETURBULENCEElement {
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

func (e *SVGFETURBULENCEElement) AttrsMap(attrs map[string]string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFETURBULENCEElement) Text(text string) *SVGFETURBULENCEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFETURBULENCEElement) TextF(format string, args ...any) *SVGFETURBULENCEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFETURBULENCEElement) IfText(condition bool, text string) *SVGFETURBULENCEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFETURBULENCEElement) IfTextF(condition bool, format string, args ...any) *SVGFETURBULENCEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFETURBULENCEElement) Escaped(text string) *SVGFETURBULENCEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFETURBULENCEElement) IfEscaped(condition bool, text string) *SVGFETURBULENCEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFETURBULENCEElement) EscapedF(format string, args ...any) *SVGFETURBULENCEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFETURBULENCEElement) IfEscapedF(condition bool, format string, args ...any) *SVGFETURBULENCEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFETURBULENCEElement) CustomData(key, value string) *SVGFETURBULENCEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = zarray.NewSortMap[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFETURBULENCEElement) IfCustomData(condition bool, key, value string) *SVGFETURBULENCEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFETURBULENCEElement) CustomDataF(key, format string, args ...any) *SVGFETURBULENCEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFETURBULENCEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFETURBULENCEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFETURBULENCEElement) CustomDataRemove(key string) *SVGFETURBULENCEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Delete(key)
	return e
}

// The baseFrequency attribute represent the base frequencies in the X and Y
// directions of the turbulence function.
func (e *SVGFETURBULENCEElement) BASE_FREQUENCY(s string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("baseFrequency", s)
	return e
}

func (e *SVGFETURBULENCEElement) BASE_FREQUENCYF(format string, args ...any) *SVGFETURBULENCEElement {
	return e.BASE_FREQUENCY(fmt.Sprintf(format, args...))
}

func (e *SVGFETURBULENCEElement) IfBASE_FREQUENCY(condition bool, s string) *SVGFETURBULENCEElement {
	if condition {
		e.BASE_FREQUENCY(s)
	}
	return e
}

func (e *SVGFETURBULENCEElement) IfBASE_FREQUENCYF(condition bool, format string, args ...any) *SVGFETURBULENCEElement {
	if condition {
		e.BASE_FREQUENCY(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute BASE_FREQUENCY from the element.
func (e *SVGFETURBULENCEElement) BASE_FREQUENCYRemove(s string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("baseFrequency")
	return e
}

func (e *SVGFETURBULENCEElement) BASE_FREQUENCYRemoveF(format string, args ...any) *SVGFETURBULENCEElement {
	return e.BASE_FREQUENCYRemove(fmt.Sprintf(format, args...))
}

// The numOctaves attribute indicates the number of octaves to be used by the
// noise function.
func (e *SVGFETURBULENCEElement) NUM_OCTAVES(f float64) *SVGFETURBULENCEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("numOctaves", f)
	return e
}

func (e *SVGFETURBULENCEElement) IfNUM_OCTAVES(condition bool, f float64) *SVGFETURBULENCEElement {
	if condition {
		e.NUM_OCTAVES(f)
	}
	return e
}

// The seed attribute indicates which number to use to seed the random number
// generator.
func (e *SVGFETURBULENCEElement) SEED(f float64) *SVGFETURBULENCEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = zarray.NewSortMap[string, float64]()
	}
	e.FloatAttributes.Set("seed", f)
	return e
}

func (e *SVGFETURBULENCEElement) IfSEED(condition bool, f float64) *SVGFETURBULENCEElement {
	if condition {
		e.SEED(f)
	}
	return e
}

// The stitchTiles attribute indicates how the Perlin noise function should be
// tiled
// It is ignored if type is not set to 'turbulence'.
func (e *SVGFETURBULENCEElement) STITCH_TILES(c SVGFeTurbulenceStitchTilesChoice) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("stitchTiles", string(c))
	return e
}

type SVGFeTurbulenceStitchTilesChoice string

const (
	// The <feTurbulence> SVG filter primitive creates an image using the Perlin
	// turbulence function
	// It allows the synthesis of artificial textures like clouds or marble.
	SVGFeTurbulenceStitchTiles_noStitch SVGFeTurbulenceStitchTilesChoice = "noStitch"
	// The <feTurbulence> SVG filter primitive creates an image using the Perlin
	// turbulence function
	// It allows the synthesis of artificial textures like clouds or marble.
	SVGFeTurbulenceStitchTiles_stitch SVGFeTurbulenceStitchTilesChoice = "stitch"
)

// Remove the attribute STITCH_TILES from the element.
func (e *SVGFETURBULENCEElement) STITCH_TILESRemove(c SVGFeTurbulenceStitchTilesChoice) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("stitchTiles")
	return e
}

// The type of turbulence function.
func (e *SVGFETURBULENCEElement) TYPE(c SVGFeTurbulenceTypeChoice) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("type", string(c))
	return e
}

type SVGFeTurbulenceTypeChoice string

const (
	// The type of turbulence function.
	SVGFeTurbulenceType_fractalNoise SVGFeTurbulenceTypeChoice = "fractalNoise"
	// The type of turbulence function.
	SVGFeTurbulenceType_turbulence SVGFeTurbulenceTypeChoice = "turbulence"
)

// Remove the attribute TYPE from the element.
func (e *SVGFETURBULENCEElement) TYPERemove(c SVGFeTurbulenceTypeChoice) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("type")
	return e
}

// Specifies a unique id for an element
func (e *SVGFETURBULENCEElement) ID(s string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFETURBULENCEElement) IDF(format string, args ...any) *SVGFETURBULENCEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFETURBULENCEElement) IfID(condition bool, s string) *SVGFETURBULENCEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFETURBULENCEElement) IfIDF(condition bool, format string, args ...any) *SVGFETURBULENCEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFETURBULENCEElement) IDRemove(s string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("id")
	return e
}

func (e *SVGFETURBULENCEElement) IDRemoveF(format string, args ...any) *SVGFETURBULENCEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFETURBULENCEElement) CLASS(s ...string) *SVGFETURBULENCEElement {
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

func (e *SVGFETURBULENCEElement) IfCLASS(condition bool, s ...string) *SVGFETURBULENCEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFETURBULENCEElement) CLASSRemove(s ...string) *SVGFETURBULENCEElement {
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

// Specifies an inline CSS style for an element
func (e *SVGFETURBULENCEElement) STYLEF(k string, format string, args ...any) *SVGFETURBULENCEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFETURBULENCEElement) IfSTYLE(condition bool, k string, v string) *SVGFETURBULENCEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFETURBULENCEElement) STYLE(k string, v string) *SVGFETURBULENCEElement {
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

func (e *SVGFETURBULENCEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFETURBULENCEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFETURBULENCEElement) STYLEMap(m map[string]string) *SVGFETURBULENCEElement {
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
func (e *SVGFETURBULENCEElement) STYLEPairs(pairs ...string) *SVGFETURBULENCEElement {
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

func (e *SVGFETURBULENCEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFETURBULENCEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFETURBULENCEElement) STYLERemove(keys ...string) *SVGFETURBULENCEElement {
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

// Make a request for an HTML

func (e *SVGFETURBULENCEElement) Z_REQ(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfZ_REQ(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.Z_REQ(expression)
	}
	return e
}

// Remove the attribute Z_REQ from the element.
func (e *SVGFETURBULENCEElement) Z_REQRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req")
	return e
}

// Replace another part of a page with incoming HTML

func (e *SVGFETURBULENCEElement) Z_TARGET(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-target"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfZ_TARGET(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.Z_TARGET(expression)
	}
	return e
}

// Remove the attribute Z_TARGET from the element.
func (e *SVGFETURBULENCEElement) Z_TARGETRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-target")
	return e
}

// Select only a part of a response

func (e *SVGFETURBULENCEElement) Z_REQ_SELECTOR(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-selector"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfZ_REQ_SELECTOR(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.Z_REQ_SELECTOR(expression)
	}
	return e
}

// Remove the attribute Z_REQ_SELECTOR from the element.
func (e *SVGFETURBULENCEElement) Z_REQ_SELECTORRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-selector")
	return e
}

// Select a strategy for HTML replacement

func (e *SVGFETURBULENCEElement) Z_SWAP(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfZ_SWAP(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.Z_SWAP(expression)
	}
	return e
}

// Remove the attribute Z_SWAP from the element.
func (e *SVGFETURBULENCEElement) Z_SWAPRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap")
	return e
}

// Push HTML from server to a client

func (e *SVGFETURBULENCEElement) Z_SWAP_PUSH(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-swap-push"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfZ_SWAP_PUSH(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.Z_SWAP_PUSH(expression)
	}
	return e
}

// Remove the attribute Z_SWAP_PUSH from the element.
func (e *SVGFETURBULENCEElement) Z_SWAP_PUSHRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-swap-push")
	return e
}

// Specify event which triggers the request

func (e *SVGFETURBULENCEElement) Z_TRIGGER(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-trigger"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfZ_TRIGGER(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.Z_TRIGGER(expression)
	}
	return e
}

// Remove the attribute Z_TRIGGER from the element.
func (e *SVGFETURBULENCEElement) Z_TRIGGERRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-trigger")
	return e
}

// Is it GET or POST?
func (e *SVGFETURBULENCEElement) Z_REQ_METHOD(c SVGFeTurbulenceZReqMethodChoice) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	e.StringAttributes.Set("z-req-method", string(c))
	return e
}

type SVGFeTurbulenceZReqMethodChoice string

const (
	// default GET
	SVGFeTurbulenceZReqMethod_empty SVGFeTurbulenceZReqMethodChoice = ""
	// GET
	SVGFeTurbulenceZReqMethod_get SVGFeTurbulenceZReqMethodChoice = "get"
	// POST
	SVGFeTurbulenceZReqMethod_post SVGFeTurbulenceZReqMethodChoice = "post"
)

// Remove the attribute Z_REQ_METHOD from the element.
func (e *SVGFETURBULENCEElement) Z_REQ_METHODRemove(c SVGFeTurbulenceZReqMethodChoice) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-method")
	return e
}

// How to deal with multiple requests being generated

func (e *SVGFETURBULENCEElement) Z_REQ_STRATEGY(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-strategy"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfZ_REQ_STRATEGY(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.Z_REQ_STRATEGY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_STRATEGY from the element.
func (e *SVGFETURBULENCEElement) Z_REQ_STRATEGYRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-strategy")
	return e
}

// Change URL after request

func (e *SVGFETURBULENCEElement) Z_REQ_HISTORY(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-history"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfZ_REQ_HISTORY(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.Z_REQ_HISTORY(expression)
	}
	return e
}

// Remove the attribute Z_REQ_HISTORY from the element.
func (e *SVGFETURBULENCEElement) Z_REQ_HISTORYRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-history")
	return e
}

// Additional data for request

func (e *SVGFETURBULENCEElement) Z_DATA(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-data"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfZ_DATA(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.Z_DATA(expression)
	}
	return e
}

// Remove the attribute Z_DATA from the element.
func (e *SVGFETURBULENCEElement) Z_DATARemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-data")
	return e
}

// As ts-data, but for JSON requests

func (e *SVGFETURBULENCEElement) Z_JSON(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-json"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfZ_JSON(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.Z_JSON(expression)
	}
	return e
}

// Remove the attribute Z_JSON from the element.
func (e *SVGFETURBULENCEElement) Z_JSONRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-json")
	return e
}

// Combine multiple requests into a single one

func (e *SVGFETURBULENCEElement) Z_REQ_BATCH(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-batch"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfZ_REQ_BATCH(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.Z_REQ_BATCH(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BATCH from the element.
func (e *SVGFETURBULENCEElement) Z_REQ_BATCHRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-batch")
	return e
}

// Run actions

func (e *SVGFETURBULENCEElement) Z_ACTION(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-action"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfZ_ACTION(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.Z_ACTION(expression)
	}
	return e
}

// Remove the attribute Z_ACTION from the element.
func (e *SVGFETURBULENCEElement) Z_ACTIONRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-action")
	return e
}

// Actions to run before request

func (e *SVGFETURBULENCEElement) Z_REQ_BEFORE(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-before"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfZ_REQ_BEFORE(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.Z_REQ_BEFORE(expression)
	}
	return e
}

// Remove the attribute Z_REQ_BEFORE from the element.
func (e *SVGFETURBULENCEElement) Z_REQ_BEFORERemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-before")
	return e
}

// Actions to run after request

func (e *SVGFETURBULENCEElement) Z_REQ_AFTER(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}

	key := "z-req-after"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfZ_REQ_AFTER(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.Z_REQ_AFTER(expression)
	}
	return e
}

// Remove the attribute Z_REQ_AFTER from the element.
func (e *SVGFETURBULENCEElement) Z_REQ_AFTERRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Delete("z-req-after")
	return e
}
