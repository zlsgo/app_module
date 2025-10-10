package elements

import (
	"fmt"
	"html"
	"io"
	"strings"
	"sync"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zutil"
	"golang.org/x/exp/constraints"
)

var (
	openBracket       = []byte("<")
	closeBracket      = []byte(">")
	spaceCloseBracket = []byte(" >")
	openSlash         = []byte("</")
	equalDblQuote     = []byte("=\"")
	dblQuote          = []byte("\"")
	space             = []byte(" ")

	builderPool = sync.Pool{
		New: func() interface{} {
			b := &strings.Builder{}
			b.Grow(256)
			return b
		},
	}
)

type ElementRenderer interface {
	Render(w io.Writer) error
}

type ElementRendererFunc func() ElementRenderer

type Element struct {
	Tag                  []byte
	IsSelfClosing        bool
	IntAttributes        *zarray.SortMaper[string, int]
	FloatAttributes      *zarray.SortMaper[string, float64]
	StringAttributes     *zarray.SortMaper[string, string]
	DelimitedStrings     *zarray.SortMaper[string, *DelimitedBuilder[string]]
	KVStrings            *zarray.SortMaper[string, *KVBuilder]
	BoolAttributes       *zarray.SortMaper[string, bool]
	CustomDataAttributes *zarray.SortMaper[string, string]
	Descendants          []ElementRenderer
}

func (e *Element) Attr(name string, value ...string) *Element {
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

func (e *Element) Attrs(attrs ...string) *Element {
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

func (e *Element) AttrsMap(attrs map[string]string) *Element {
	if e.StringAttributes == nil {
		e.StringAttributes = zarray.NewSortMap[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *Element) Render(w io.Writer) error {
	w.Write(openBracket)
	w.Write(e.Tag)

	finalKeys := zarray.NewSortMap[string, string]()

	if e.IntAttributes != nil && e.IntAttributes.Len() > 0 {
		e.IntAttributes.ForEach(func(k string, v int) bool {
			finalKeys.Set(k, fmt.Sprint(v))
			return true
		})
	}

	if e.FloatAttributes != nil && e.FloatAttributes.Len() > 0 {
		e.FloatAttributes.ForEach(func(k string, v float64) bool {
			finalKeys.Set(k, fmt.Sprint(v))
			return true
		})
	}

	if e.StringAttributes != nil && e.StringAttributes.Len() > 0 {
		e.StringAttributes.ForEach(func(k string, v string) bool {
			finalKeys.Set(k, v)
			return true
		})
	}

	if e.DelimitedStrings != nil && e.DelimitedStrings.Len() > 0 {
		e.DelimitedStrings.ForEach(func(k string, v *DelimitedBuilder[string]) bool {
			buf := zutil.GetBuff()
			if err := v.Render(buf); err != nil {
				return false
			}
			finalKeys.Set(k, buf.String())
			zutil.PutBuff(buf)
			return true
		})
	}

	if e.KVStrings != nil && e.KVStrings.Len() > 0 {
		e.KVStrings.ForEach(func(k string, v *KVBuilder) bool {
			buf := zutil.GetBuff()
			if err := v.Render(buf); err != nil {
				return false
			}
			finalKeys.Set(k, buf.String())
			zutil.PutBuff(buf)
			return true
		})
	}

	if e.CustomDataAttributes != nil && e.CustomDataAttributes.Len() > 0 {
		e.CustomDataAttributes.ForEach(func(k string, v string) bool {
			// Keys in CustomDataAttributes are already full attribute names
			finalKeys.Set(k, v)
			return true
		})
	}

	if e.BoolAttributes != nil && e.BoolAttributes.Len() > 0 {
		e.BoolAttributes.ForEach(func(k string, v bool) bool {
			if v {
				finalKeys.Set(k, "")
			}
			return true
		})
	}

	if finalKeys.Len() > 0 {
		finalKeys.ForEach(func(k string, v string) bool {
			w.Write(space)
			w.Write([]byte(k))

			if v != "" {
				w.Write(equalDblQuote)
				w.Write([]byte(fmt.Sprint(v)))
				w.Write(dblQuote)
			}
			return true
		})
	}

	if e.IsSelfClosing {
		w.Write(spaceCloseBracket)
		return nil
	}
	w.Write(closeBracket)

	for _, d := range e.Descendants {
		if d == nil {
			continue
		}
		if err := d.Render(w); err != nil {
			return err
		}
	}

	w.Write(openSlash)
	w.Write(e.Tag)
	w.Write(closeBracket)

	return nil
}

type customDataKeyModifier func() string

func customDataKey(key string, modifiers ...customDataKeyModifier) string {
	sb := builderPool.Get().(*strings.Builder)
	defer func() {
		sb.Reset()
		builderPool.Put(sb)
	}()

	sb.WriteString(key)
	for _, m := range modifiers {
		sb.WriteRune('.')
		sb.WriteString(m())
	}
	return sb.String()
}

type DelimitedBuilder[T constraints.Ordered] struct {
	Delimiter string
	Values    *zarray.SortMaper[T, struct{}]
}

func NewDelimitedBuilder[T constraints.Ordered](delimiter string) *DelimitedBuilder[T] {
	return &DelimitedBuilder[T]{
		Delimiter: delimiter,
		Values:    zarray.NewSortMap[T, struct{}](),
	}
}

func (d *DelimitedBuilder[T]) Add(values ...T) *DelimitedBuilder[T] {
	for _, v := range values {
		d.Values.Set(v, struct{}{})
	}
	return d
}

func (d *DelimitedBuilder[T]) Remove(values ...T) *DelimitedBuilder[T] {
	for _, v := range values {
		d.Values.Delete(v)
	}
	return d
}

func (d *DelimitedBuilder[T]) Render(w io.Writer) error {
	count := 0
	total := d.Values.Len()
	d.Values.ForEach(func(k T, v struct{}) bool {
		b := []byte(fmt.Sprint(k))
		if _, err := w.Write(b); err != nil {
			return false
		}

		count++
		if count < total {
			w.Write([]byte(d.Delimiter))
		}
		return true
	})
	return nil
}

type KVBuilder struct {
	KeyPairDelimiter string
	EntryDelimiter   string
	Values           **zarray.SortMaper[string, string]
}

func NewKVBuilder(keyPairDelimiter, entryDelimiter string) *KVBuilder {
	return &KVBuilder{
		KeyPairDelimiter: keyPairDelimiter,
		EntryDelimiter:   entryDelimiter,
		Values:           new(*zarray.SortMaper[string, string]),
	}
}

func (d *KVBuilder) Add(key, value string) *KVBuilder {
	if *d.Values == nil {
		*d.Values = zarray.NewSortMap[string, string]()
	}
	(*d.Values).Set(key, value)
	return d
}

func (d *KVBuilder) Remove(key string) *KVBuilder {
	if *d.Values == nil {
		return d
	}
	(*d.Values).Delete(key)
	return d
}

func (d *KVBuilder) Render(w io.Writer) error {
	count := 0
	total := (*d.Values).Len()
	(*d.Values).ForEach(func(k string, v string) bool {
		w.Write([]byte(k))
		w.Write([]byte(d.KeyPairDelimiter))
		w.Write([]byte(v))
		count++
		if count < total {
			w.Write([]byte(d.EntryDelimiter))
		}
		return true
	})
	return nil
}

type TextContent string

func (tc *TextContent) Render(w io.Writer) error {
	_, err := w.Write([]byte(*tc))
	return err
}

func Text(text string) *TextContent {
	return (*TextContent)(&text)
}

func TextF(format string, args ...interface{}) *TextContent {
	return Text(fmt.Sprintf(format, args...))
}

type EscapedContent string

func (ec *EscapedContent) Render(w io.Writer) error {
	_, err := w.Write([]byte(html.EscapeString(string(*ec))))
	return err
}

func Escaped(text string) *EscapedContent {
	return (*EscapedContent)(&text)
}

func EscapedF(format string, args ...interface{}) *EscapedContent {
	return Escaped(fmt.Sprintf(format, args...))
}

type Grouper struct {
	Children []ElementRenderer
}

func (g *Grouper) Render(w io.Writer) error {
	for _, child := range g.Children {
		if err := child.Render(w); err != nil {
			return fmt.Errorf("failed to build element: %w", err)
		}
	}
	return nil
}

func Group(children ...ElementRenderer) *Grouper {
	return &Grouper{
		Children: children,
	}
}

func If(condition bool, children ...ElementRenderer) ElementRenderer {
	if condition {
		return Group(children...)
	}
	return nil
}

func Tern(condition bool, trueChildren, falseChildren ElementRenderer) ElementRenderer {
	if condition {
		return trueChildren
	}
	return falseChildren
}

func Range[T any](values []T, cb func(T) ElementRenderer) ElementRenderer {
	children := make([]ElementRenderer, 0, len(values))
	for _, value := range values {
		children = append(children, cb(value))
	}
	return Group(children...)
}

func RangeI[T any](values []T, cb func(int, T) ElementRenderer) ElementRenderer {
	children := make([]ElementRenderer, 0, len(values))
	for i, value := range values {
		children = append(children, cb(i, value))
	}
	return Group(children...)
}

func DynGroup(childrenFuncs ...ElementRendererFunc) *Grouper {
	children := make([]ElementRenderer, 0, len(childrenFuncs))
	for _, childFunc := range childrenFuncs {
		child := childFunc()
		if child != nil {
			children = append(children, child)
		}
	}
	return &Grouper{
		Children: children,
	}
}

func DynIf(condition bool, childrenFuncs ...ElementRendererFunc) ElementRenderer {
	if condition {
		children := make([]ElementRenderer, 0, len(childrenFuncs))
		for _, childFunc := range childrenFuncs {
			child := childFunc()
			if child != nil {
				children = append(children, child)
			}
		}
		return Group(children...)
	}
	return nil
}

func DynTern(condition bool, trueChildren, falseChildren ElementRendererFunc) ElementRenderer {
	if condition {
		return trueChildren()
	}
	return falseChildren()
}

func NewElement(tag string, children ...ElementRenderer) *Element {
	return &Element{
		Tag:         []byte(tag),
		Descendants: children,
	}
}

func Error(err error) ElementRenderer {
	return Text(err.Error())
}
