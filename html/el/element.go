package el

import (
	"fmt"
	"html"
	"io"
	"strings"
	"sync"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/ztype"
)

// Element 统一的HTML元素实现
type Element struct {
	tag         string
	selfClosing bool

	// 快速属性访问字段
	fastID    string          // id属性
	fastClass *DelimitedValue // class属性
	fastStyle *KVValue        // style属性

	attributes *zarray.SortMaper[string, AttributeValue]
	children   []ElementRenderer
	config     *ElementConfig
}

func (d *DelimitedValue) AppendTo(sb *strings.Builder) {
	if d.Values.Len() == 0 {
		return
	}
	sb.WriteString(`="`)
	first := true
	d.Values.ForEach(func(k string, _ struct{}) bool {
		if !first {
			sb.WriteString(d.Delimiter)
		}
		first = false
		sb.WriteString(html.EscapeString(k))
		return true
	})
	sb.WriteByte('"')
}

// AttributeValue 属性值接口
// 支持多种类型的属性值，实现策略模式
type AttributeValue interface {
	Render(w io.Writer) error
	String() string
}

type AttributeAppender interface {
	AppendTo(sb *strings.Builder)
}

// StringValue 字符串属性值
type StringValue string

func (s StringValue) Render(w io.Writer) error {
	if s == "" {
		return nil // 空值属性只渲染属性名
	}
	_, err := fmt.Fprintf(w, `="%s"`, html.EscapeString(string(s)))
	return err
}

func (s StringValue) String() string {
	return string(s)
}

func (s StringValue) AppendTo(sb *strings.Builder) {
	if s == "" {
		return
	}
	sb.WriteString(`="`)
	sb.WriteString(html.EscapeString(string(s)))
	sb.WriteByte('"')
}

// BoolValue 布尔属性值
type BoolValue bool

func (b BoolValue) Render(w io.Writer) error {
	// 布尔属性只渲染属性名，不渲染值
	return nil
}

func (b BoolValue) String() string {
	if b {
		return "true"
	}
	return "false"
}

// DelimitedValue 分隔符分隔的值（用于 class, style 等）
type DelimitedValue struct {
	Values    *zarray.SortMaper[string, struct{}]
	Delimiter string
}

func NewDelimitedValue(delimiter string) *DelimitedValue {
	return &DelimitedValue{
		Values:    zarray.NewSortMap[string, struct{}](),
		Delimiter: delimiter,
	}
}

func (d *DelimitedValue) Add(values ...string) *DelimitedValue {
	for _, v := range values {
		d.Values.Set(v, struct{}{})
	}
	return d
}

func (d *DelimitedValue) Remove(values ...string) *DelimitedValue {
	for _, v := range values {
		d.Values.Delete(v)
	}
	return d
}

func (d *DelimitedValue) Render(w io.Writer) error {
	if d.Values.Len() == 0 {
		return nil
	}

	var sb strings.Builder
	sb.Grow(d.Values.Len() * 16)
	first := true
	d.Values.ForEach(func(k string, _ struct{}) bool {
		if !first {
			sb.WriteString(d.Delimiter)
		}
		first = false
		sb.WriteString(k)
		return true
	})

	escaped := html.EscapeString(sb.String())
	_, err := fmt.Fprintf(w, `="%s"`, escaped)
	return err
}

func (d *DelimitedValue) String() string {
	parts := make([]string, 0, d.Values.Len())
	d.Values.ForEach(func(k string, _ struct{}) bool {
		parts = append(parts, k)
		return true
	})
	return strings.Join(parts, d.Delimiter)
}

// KVValue 键值对属性值（用于 style 等）
type KVValue struct {
	Pairs         *zarray.SortMaper[string, string]
	PairDelimiter string // 键值对之间的分隔符，如 "; "
	KeyDelimiter  string // 键和值之间的分隔符，如 ": "
}

func NewKVValue(keyDelimiter, pairDelimiter string) *KVValue {
	return &KVValue{
		Pairs:         zarray.NewSortMap[string, string](),
		KeyDelimiter:  keyDelimiter,
		PairDelimiter: pairDelimiter,
	}
}

func (kv *KVValue) Set(key, value string) *KVValue {
	kv.Pairs.Set(key, value)
	return kv
}

func (kv *KVValue) Delete(key string) *KVValue {
	kv.Pairs.Delete(key)
	return kv
}

func (kv *KVValue) Render(w io.Writer) error {
	if kv.Pairs.Len() == 0 {
		return nil
	}

	var sb strings.Builder
	sb.Grow(kv.Pairs.Len() * 32)
	first := true
	kv.Pairs.ForEach(func(k, v string) bool {
		if !first {
			sb.WriteString(kv.PairDelimiter)
		}
		first = false
		sb.WriteString(k)
		sb.WriteString(kv.KeyDelimiter)
		sb.WriteString(v)
		return true
	})

	escaped := html.EscapeString(sb.String())
	_, err := fmt.Fprintf(w, `="%s"`, escaped)
	return err
}

func (kv *KVValue) String() string {
	parts := make([]string, 0, kv.Pairs.Len())
	kv.Pairs.ForEach(func(k, v string) bool {
		parts = append(parts, k+kv.KeyDelimiter+v)
		return true
	})
	return strings.Join(parts, kv.PairDelimiter)
}

func (kv *KVValue) AppendTo(sb *strings.Builder) {
	if kv.Pairs.Len() == 0 {
		return
	}

	sb.WriteString(`="`)
	first := true
	kv.Pairs.ForEach(func(k, v string) bool {
		if !first {
			sb.WriteString(kv.PairDelimiter)
		}
		first = false
		sb.WriteString(html.EscapeString(k))
		sb.WriteString(kv.KeyDelimiter)
		sb.WriteString(html.EscapeString(v))
		return true
	})
	sb.WriteByte('"')
}

// NewElement 创建新元素
func NewElement(tag string, children ...ElementRenderer) *Element {
	return &Element{
		tag:      tag,
		children: children,
		// attributes 延迟初始化，优先使用快速访问字段
	}
}

// NewElementWithConfig 使用配置创建元素
func NewElementWithConfig(config *ElementConfig, children ...ElementRenderer) *Element {
	element := &Element{
		tag:         config.Tag,
		selfClosing: config.SelfClosing,
		children:    make([]ElementRenderer, 0),
		config:      config,
		// attributes 延迟初始化，优先使用快速访问字段
	}

	if config.AllowChildren {
		element.children = append(element.children, children...)
	}

	return element
}

// 分层内存池
var (
	smallBufPool = sync.Pool{
		New: func() interface{} {
			sb := &strings.Builder{}
			sb.Grow(256) // 小元素：基础标签+少量属性
			return sb
		},
	}
	mediumBufPool = sync.Pool{
		New: func() interface{} {
			sb := &strings.Builder{}
			sb.Grow(1024) // 中等元素：多属性或嵌套
			return sb
		},
	}
	largeBufPool = sync.Pool{
		New: func() interface{} {
			sb := &strings.Builder{}
			sb.Grow(4096) // 复杂元素：深嵌套或大量属性
			return sb
		},
	}
)

// 智能缓冲区选择器
func (e *Element) selectBufPool() *sync.Pool {
	complexity := e.estimateComplexity()
	switch {
	case complexity <= 256:
		return &smallBufPool
	case complexity <= 1024:
		return &mediumBufPool
	default:
		return &largeBufPool
	}
}

// 估算元素复杂度（字节）
func (e *Element) estimateComplexity() int {
	size := len(e.tag)*2 + 5

	// 快速属性复杂度
	if e.fastID != "" {
		size += len(e.fastID) + 8
	}
	if e.fastClass != nil && e.fastClass.Values.Len() > 0 {
		size += e.fastClass.Values.Len() * 16
	}
	if e.fastStyle != nil && e.fastStyle.Pairs.Len() > 0 {
		size += e.fastStyle.Pairs.Len() * 24
	}

	// 其他属性复杂度
	if e.attributes != nil {
		size += e.attributes.Len() * 32
	}

	// 子元素复杂度（递归估算）
	for _, child := range e.children {
		if elem, ok := child.(*Element); ok {
			size += elem.estimateComplexity()
		} else {
			size += 64
		}
	}

	return size
}

// Render 渲染HTML到 io.Writer
func (e *Element) Render(w io.Writer) error {
	if len(e.children) == 0 {
		if !e.selfClosing && e.fastID == "" && e.fastClass == nil && e.fastStyle == nil &&
			(e.attributes == nil || e.attributes.Len() == 0) {
			_, err := w.Write([]byte("<" + e.tag + "></" + e.tag + ">"))
			return err
		}
	}

	pool := e.selectBufPool()
	sb := pool.Get().(*strings.Builder)
	sb.Reset()
	defer pool.Put(sb)

	e.renderToBuilder(sb)
	_, err := w.Write([]byte(sb.String()))
	return err
}

func (e *Element) renderToBuilder(sb *strings.Builder) {
	sb.WriteByte('<')
	sb.WriteString(e.tag)

	e.renderFastAttributes(sb)

	if e.attributes != nil && e.attributes.Len() > 0 {
		e.attributes.ForEach(func(name string, value AttributeValue) bool {
			sb.WriteByte(' ')
			sb.WriteString(name)
			if app, ok := value.(AttributeAppender); ok {
				app.AppendTo(sb)
			} else if sv, ok := value.(StringValue); ok {
				if sv != "" {
					sb.WriteString(`="`)
					sb.WriteString(html.EscapeString(string(sv)))
					sb.WriteByte('"')
				}
			} else if _, ok := value.(BoolValue); !ok {
				sb.WriteString(`="`)
				sb.WriteString(ztype.ToString(value))
				sb.WriteByte('"')
			}
			return true
		})
	}

	// 自闭合标签
	if e.selfClosing {
		sb.WriteString(" />")
		return
	}

	// 闭合开始标签
	sb.WriteByte('>')

	for _, child := range e.children {
		if child == nil {
			continue
		}
		if elem, ok := child.(*Element); ok {
			elem.renderToBuilder(sb)
		} else {
			child.Render(&builderWriter{sb})
		}
	}

	// 结束标签
	sb.WriteString("</")
	sb.WriteString(e.tag)
	sb.WriteByte('>')
}

// builderWriter 将strings.Builder适配为io.Writer接口
type builderWriter struct {
	sb *strings.Builder
}

func (bw *builderWriter) Write(p []byte) (int, error) {
	return bw.sb.Write(p)
}

// renderFastAttributes 渲染常用属性
func (e *Element) renderFastAttributes(sb *strings.Builder) {
	// ID属性
	if e.fastID != "" {
		sb.WriteString(` id="`)
		sb.WriteString(html.EscapeString(e.fastID))
		sb.WriteByte('"')
	}

	// Class属性
	if e.fastClass != nil && e.fastClass.Values.Len() > 0 {
		sb.WriteString(` class`)
		e.fastClass.AppendTo(sb)
	}

	// Style属性
	if e.fastStyle != nil && e.fastStyle.Pairs.Len() > 0 {
		sb.WriteString(` style`)
		e.fastStyle.AppendTo(sb)
	}
}

// Attr 设置字符串属性
func (e *Element) Attr(name string, value ...string) *Element {
	val := ""
	if len(value) > 0 {
		val = value[0]
	}

	switch name {
	case "id":
		e.fastID = val
		return e
	}

	if e.attributes == nil {
		e.attributes = zarray.NewSortMap[string, AttributeValue]()
	}
	if val == "" {
		e.attributes.Set(name, StringValue(""))
	} else {
		e.attributes.Set(name, StringValue(val))
	}
	return e
}

// Attrs 批量设置属性
func (e *Element) Attrs(attrs ...string) *Element {
	for i := 0; i < len(attrs); i += 2 {
		k := attrs[i]
		v := ""
		if i+1 < len(attrs) {
			v = attrs[i+1]
		}
		e.Attr(k, v)
	}
	return e
}

// AttrsMap 从 map 设置属性
func (e *Element) AttrsMap(attrs map[string]string) *Element {
	for k, v := range attrs {
		e.Attr(k, v)
	}
	return e
}

// BoolAttr 设置布尔属性
func (e *Element) BoolAttr(name string, value bool) *Element {
	if e.attributes == nil {
		e.attributes = zarray.NewSortMap[string, AttributeValue]()
	}
	if value {
		e.attributes.Set(name, BoolValue(true))
	} else {
		e.attributes.Delete(name)
	}
	return e
}

// Class 添加 CSS 类 - 使用快速访问字段
func (e *Element) Class(classes ...string) *Element {
	if e.fastClass == nil {
		e.fastClass = NewDelimitedValue(" ")
	}
	e.fastClass.Add(classes...)
	return e
}

// RemoveClass 移除 CSS 类 - 使用快速访问字段
func (e *Element) RemoveClass(classes ...string) *Element {
	if e.fastClass != nil {
		e.fastClass.Remove(classes...)
	}
	return e
}

// Style 设置 CSS 样式 - 使用快速访问字段
func (e *Element) Style(key, value string) *Element {
	if e.fastStyle == nil {
		e.fastStyle = NewKVValue(": ", "; ")
	}
	e.fastStyle.Set(key, value)
	return e
}

// Children 添加子元素
func (e *Element) Children(children ...ElementRenderer) *Element {
	e.children = append(e.children, children...)
	return e
}

// Text 添加文本内容
func (e *Element) Text(text string) *Element {
	e.children = append(e.children, Text(text))
	return e
}

// TextF 添加格式化文本
func (e *Element) TextF(format string, args ...interface{}) *Element {
	return e.Text(fmt.Sprintf(format, args...))
}

// Escaped 添加转义文本
func (e *Element) Escaped(text string) *Element {
	e.children = append(e.children, Escaped(text))
	return e
}

// EscapedF 添加格式化转义文本
func (e *Element) EscapedF(format string, args ...interface{}) *Element {
	return e.Escaped(fmt.Sprintf(format, args...))
}

// IfText 条件添加文本
func (e *Element) IfText(condition bool, text string) *Element {
	if condition {
		return e.Text(text)
	}
	return e
}

// IfEscaped 条件添加转义文本
func (e *Element) IfEscaped(condition bool, text string) *Element {
	if condition {
		return e.Escaped(text)
	}
	return e
}
