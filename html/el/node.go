package el

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"iter"
	"log"
	"slices"
	"strings"
	"sync"
)

// Item 接口是任意可用作节点或属性项的占位抽象。
type Item interface {
	Item()
}

// Node 接口表示可以在指定上下文中渲染到 io.Writer 的元素，实现需要定义 Render 与 Release 行为。
type Node interface {
	Item
	Render(ChunkWriter)
	Release()
}

// Text 表示渲染器中的纯文本节点，会自动进行 HTML 转义。
type Text string

// Item 方法声明该类型满足 Item 接口约束。
func (Text) Item() {}

// Render 方法将文本内容写入 ChunkWriter，并进行 HTML 转义。
func (t Text) Render(w ChunkWriter) {
	w.Write(StaticChunk(html.EscapeString(string(t))))
}

// Release 方法为占位实现，文本节点无需释放资源。
func (Text) Release() {
	// no-op
}

// Textf 根据格式化模板生成文本节点，便于动态拼接内容。
func Textf(format string, a ...any) Text {
	return Text(fmt.Sprintf(format, a...))
}

// Frag 表示多个节点的集合，可顺序渲染组合内容。
type Frag []Node

// Item 方法声明片段实现了 Item 接口。
func (Frag) Item() {}

// Release 逐个释放片段中的子节点，帮助复用资源。
func (f Frag) Release() {
	for _, node := range f {
		node.Release()
	}
}

// Render 按序渲染片段中的每个子节点。
func (f Frag) Render(w ChunkWriter) {
	for _, node := range f {
		node.Render(w)
	}
}

// Fragment 根据可变参数构造片段节点，便于组合多个 Node。
func Fragment(nodes ...Node) Frag {
	return Frag(nodes)
}

type ProperLifeCycle int

const (
	LifeCycleImmediate = ProperLifeCycle(iota)
	LifeCycleStatic
	LifeCycleDeferred
)

// Proper 接口表示可应用到元素的属性，实现者决定属性如何作用于元素。
type Proper interface {
	Item()
	LifeCycle() ProperLifeCycle
}

// Attribute 表示元素的键值属性对，可直接应用到 Element。
type Attribute struct {
	Key   string
	Value string
}

// Item 方法确保 Attribute 可作为元素构建参数传入。
func (*Attribute) Item() {}

func (*Attribute) LifeCycle() ProperLifeCycle {
	return LifeCycleImmediate
}

// Apply 方法将属性键值写入元素，同时对 class 做追加处理。
func (a *Attribute) Apply(el *Element) {
	if a.Key == "class" {
		el.AddClass(a.Value)
	} else {
		el.SetAttribute(a.Key, a.Value)
	}
}

// DeferredAttribute 表示延迟计算的属性，可在渲染阶段生成值。
type DeferredAttribute struct {
	key string
	fn  func(context.Context) string
}

// Item 方法声明 DeferredAttribute 可作为构建参数。
func (*DeferredAttribute) Item() {}

func (*DeferredAttribute) LifeCycle() ProperLifeCycle {
	return LifeCycleDeferred
}

// Apply 方法在渲染阶段按需生成属性值。
func (a *DeferredAttribute) Apply(ctx context.Context, w io.Writer) error {
	rhs := ""
	value := a.fn(ctx)
	if len(value) > 0 {
		rhs = "=" + "\"" + value + "\""
	}
	_, err := fmt.Fprintf(w, " %s%s", a.key, rhs)
	return err
}

// DeferredAttr 创建延迟计算属性，在渲染阶段动态生成值。
func DeferredAttr(key string, fn func(context.Context) string) *DeferredAttribute {
	return &DeferredAttribute{key: key, fn: fn}
}

// Attr 根据键值创建 Attribute，便于快速设置元素属性。
func Attr(key string, value string) *Attribute {
	return &Attribute{Key: key, Value: value}
}

// Component 表示延迟渲染的组件函数，可在执行阶段生成节点。
type Component func(context.Context) Node

// Item 方法声明 Component 函数可作为元素项。
func (Component) Item() {}

// Release 对组件为占位实现，组件无需额外资源释放。
func (Component) Release() {
	// no-op
}

// Render 将组件转为动态 Chunk，在执行时渲染组件生成的节点。
func (c Component) Render(cw ChunkWriter) {
	cw.Write(DynamicChunk(func(ctx context.Context, w io.Writer) error {
		writer := &teecw{ChunkWriter: cw, fn: func(c Chunk) error {
			if err := RenderChunk(c, ctx, w); err != nil {
				return err
			}
			return nil
		}}
		c(ctx).Render(writer)
		return nil
	}))
}

// Element 表示一个 HTML 元素节点，包含标签名、属性集合、子节点列表及元数据。
type Element struct {
	nodes         []Node
	properties    []Proper
	attributelist []*Attribute
	name          string
	meta          map[string]any
}

// Tag 返回元素的标签名称，并对常规标签进行小写化处理。
func (e *Element) Tag() string {
	if strings.HasPrefix(e.name, "!") {
		return e.name
	}
	return strings.ToLower(e.name)
}

// Set 将自定义元数据写入元素的 meta 存储。
func (e *Element) Set(key string, value any) {
	e.meta[key] = value
}

// Get 从元素的 meta 中读取指定键值。
func (e *Element) Get(key string) any {
	return e.meta[key]
}

// SetAttribute 设置或覆盖元素的属性值。
func (e *Element) SetAttribute(key, value string) {
	for i, attr := range e.attributelist {
		if attr.Key == key {
			e.attributelist[i].Value = value
			return
		}
	}
	e.attributelist = append(e.attributelist, &Attribute{Key: key, Value: value})
}

// AddClass 为元素追加 class 值，自动处理空值和空格。
func (e *Element) AddClass(class string) {
	attr := e.GetAttribute("class")
	if len(attr) == 0 {
		attr = class
	} else {
		attr = attr + " " + class
	}
	e.SetAttribute("class", attr)
}

// GetAttribute 返回指定属性的值，不存在时返回空字符串。
func (e *Element) GetAttribute(key string) string {
	for _, attr := range e.attributelist {
		if attr.Key == key {
			return attr.Value
		}
	}
	return ""
}

// GetAttributes 将当前属性列表复制为键值映射返回。
func (e *Element) GetAttributes() map[string]string {
	attrs := map[string]string{}
	for _, attr := range e.attributelist {
		attrs[attr.Key] = attr.Value
	}
	return attrs
}

// SetAttributes 根据提供的键值映射重建属性列表。
func (e *Element) SetAttributes(list map[string]string) {
	e.attributelist = []*Attribute{}
	for key, value := range list {
		e.attributelist = append(e.attributelist, &Attribute{Key: key, Value: value})
	}
}

// GetNodes 返回元素的所有子节点切片。
func (e *Element) GetNodes() []Node {
	return e.nodes
}

// AppendNode 将新的子节点追加到元素尾部。
func (e *Element) AppendNode(node Node) {
	e.nodes = append(e.nodes, node)
}

// Item 方法声明 Element 同时满足 Item 接口。
func (*Element) Item() {}

// Release 将元素放回对象池，便于内存复用。
func (e *Element) Release() {
	elpool.Put(e)
}

var voidelements = []string{"!DOCTYPE", "area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "source", "track", "wbr"}

// Render 顺序输出标签、属性、子节点，并处理延迟属性与自闭合标签。
func (e *Element) Render(w ChunkWriter) {
	w.Write(StaticChunk(fmt.Appendf(nil, "<%s", e.Tag())))

	if len(e.properties) > 0 {
		for _, prop := range e.properties {
			applier, ok := prop.(interface {
				Apply(context.Context, io.Writer) error
			})
			if !ok {
				w.Write(DynamicChunk(func(ctx context.Context, w io.Writer) error {
					return fmt.Errorf("property with deferred life cycle (%T) is not implementing the applier interface correctly, add a Apply(context.Context, io.Writer) error method", prop)
				}))
			} else {
				w.Write(DynamicChunk(applier.Apply))
			}
		}
	}

	for _, attr := range e.attributelist {
		var rhs string
		if len(attr.Value) != 0 {
			rhs = "=" + "\"" + attr.Value + "\""
		}
		w.Write(StaticChunk(fmt.Appendf(nil, " %s%s", attr.Key, rhs)))
	}

	w.Write(StaticChunk(">"))

	if slices.Contains(voidelements, e.Tag()) {
		return
	}

	for _, node := range e.nodes {
		node.Render(w)
	}

	w.Write(StaticChunk(fmt.Appendf(nil, "</%s>", e.Tag())))
}

// Clone 创建当前元素的浅拷贝，复制属性列表并复用子节点与属性引用。
func (e *Element) Clone() *Element {
	attributelist := make([]*Attribute, len(e.attributelist))
	for i, attr := range e.attributelist {
		attributelist[i] = &Attribute{Key: attr.Key, Value: attr.Value}
	}

	return &Element{
		nodes:         e.nodes,
		properties:    e.properties,
		attributelist: attributelist,
		name:          e.name,
		meta:          e.meta,
	}
}

var propspool = sync.Pool{
	New: func() any {
		return &[]Proper{}
	},
}

var elpool = sync.Pool{
	New: func() any {
		return &Element{
			nodes:         []Node{},
			properties:    []Proper{},
			attributelist: []*Attribute{},
			meta:          make(map[string]any),
		}
	},
}

// El 根据标签名和可变参数构造元素，自动区分节点与属性并按生命周期应用。
func El(name string, items ...Item) *Element {
	el := elpool.New().(*Element)
	el.name = name

	immediate := propspool.New().(*[]Proper)
	defer propspool.Put(immediate)
	static := propspool.New().(*[]Proper)
	defer propspool.Put(static)

	for _, item := range items {
		switch item := item.(type) {
		case Node:
			el.nodes = append(el.nodes, item)
		case Proper:
			cycle := item.LifeCycle()
			switch cycle {
			case LifeCycleImmediate:
				*immediate = append(*immediate, item)
			case LifeCycleStatic:
				*static = append(*static, item)
			case LifeCycleDeferred:
				// TODO: Deferred properties have a serious bug
				el.properties = append(el.properties, item)
			default:
				log.Fatalf("Illegal property lifecycle: %T", item)
			}
		default:
			log.Fatalf("Illegal item type: %T", item)
		}
	}

	for _, prop := range *immediate {
		applier, ok := prop.(interface{ Apply(*Element) })
		if !ok {
			log.Fatalf("Property with immediate life cycle (%T) is not implementing the applier interface correctly, add a Apply(*Element) method", prop)
		}
		applier.Apply(el)
	}

	for _, prop := range *static {
		applier, ok := prop.(interface{ Apply(*Element) })
		if !ok {
			log.Fatalf("Property with static life cycle (%T) is not implementing the applier interface correctly, add a Apply(*Element) method", prop)
		}
		applier.Apply(el)
	}
	return el
}

// VoidEl 构建自闭合元素，适用于无需闭合标签的 HTML 元素。
func VoidEl(name string, items ...Item) *Element {
	return El(name, items...)
}

// Map 将切片中的每个元素映射为 Node，并最终组合成片段返回。
func Map[E any](s []E, fn func(E) Node) Node {
	children := make([]Node, len(s))
	for i, item := range s {
		children[i] = fn(item)
	}
	return Fragment(children...)
}

// MapIdx 在映射节点的同时暴露索引，适合需要位置信息的场景。
func MapIdx[E any](s []E, fn func(E, int) Node) Node {
	children := make([]Node, len(s))
	for i, item := range s {
		children[i] = fn(item, i)
	}
	return Fragment(children...)
}

// Iter 遍历迭代器序列，将每个值映射为 Node 并组成片段返回。
func Iter[T any](s iter.Seq[T], fn func(T) Node) Node {
	children := []Node{}
	for value := range s {
		children = append(children, fn(value))
	}
	return Fragment(children...)
}

// Iter2 遍历键值迭代器，将每个键值映射为 Node 并组合成片段。
func Iter2[K any, V any](s iter.Seq2[K, V], fn func(K, V) Node) Node {
	children := []Node{}
	for key, value := range s {
		children = append(children, fn(key, value))
	}
	return Fragment(children...)
}

// If 根据条件决定返回给定节点或空片段，用于条件渲染。
func If(cond bool, item Item) Item {
	if cond {
		return item
	}
	return Fragment()
}

// IfFn 根据条件执行回调生成节点，便捷实现懒加载逻辑。
func IfFn(cond bool, fn func() Item) Item {
	if cond {
		return fn()
	}
	return Fragment()
}

// SwitchCase 表示 Switch 结构中的单个分支定义。
type SwitchCase[T comparable] struct {
	Value   T
	Fn      func() Item
	Default bool
}

// Switch 遍历所有分支，返回首个匹配值生成的节点，若无匹配则返回空片段。
func Switch[T comparable](expr T, cases ...*SwitchCase[T]) Item {
	var d *SwitchCase[T]
	for _, c := range cases {
		if expr == c.Value {
			return c.Fn()
		} else if c.Default {
			d = c
		}
	}

	if d != nil {
		return d.Fn()
	}
	return Fragment()
}

// Case 将给定值与节点绑定为 SwitchCase，匹配时返回对应节点。
func Case[T comparable](v T, item Item) *SwitchCase[T] {
	return &SwitchCase[T]{
		Value: v,
		Fn: func() Item {
			return item
		},
	}
}

// CaseFn 将值与回调绑定为 SwitchCase，匹配时执行函数生成节点。
func CaseFn[T comparable](v T, fn func() Item) *SwitchCase[T] {
	return &SwitchCase[T]{
		Value: v,
		Fn:    fn,
	}
}

// DefaultEl 创建默认分支，未命中其他分支时返回指定节点。
func DefaultEl[T comparable](item Item) *SwitchCase[T] {
	return &SwitchCase[T]{
		Fn: func() Item {
			return item
		},
		Default: true,
	}
}

// DefaultElFn 创建默认分支并使用回调生成节点。
func DefaultElFn[T comparable](fn func() Item) *SwitchCase[T] {
	return &SwitchCase[T]{
		Fn:      fn,
		Default: true,
	}
}

// JSONNode 表示可序列化为 JSON 的节点，可配置缩进格式化输出。
type JSONNode struct {
	Data   any
	Indent string
}

// Item 方法表明 JSONNode 可作为元素项使用。
func (*JSONNode) Item() {}

// Release 为占位实现，JSONNode 不需要回收资源。
func (*JSONNode) Release() {
	// no-op
}

// WithIndent 设置输出缩进字符串，并返回节点自身以便链式调用。
func (n *JSONNode) WithIndent(indent string) *JSONNode {
	n.Indent = indent
	return n
}

// Render 将数据编码为 JSON 并写入 ChunkWriter，根据需要设置缩进。
func (n *JSONNode) Render(w ChunkWriter) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", n.Indent)
	if err := enc.Encode(n.Data); err != nil {
		log.Fatal(err)
	}
	w.Write(StaticChunk(buf.Bytes()))
}

// JSON 根据任意数据构造 JSONNode，方便在 HTML 中嵌入 JSON。
func JSON(data any) *JSONNode {
	return &JSONNode{Data: data}
}

// Raw 表示原始内容，不进行转义。
type Raw string

// Item 方法声明 RawUnsafe 可直接作为元素项使用。
func (Raw) Item() {}

// Release 为占位实现，原始节点无需释放。
func (Raw) Release() {
	// no-op
}

// Render 直接输出原始内容，调用方需确保内容安全。
func (t Raw) Render(w ChunkWriter) {
	w.Write(StaticChunk(string(t)))
}
