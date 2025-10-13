package el

import (
	"io"

	"github.com/sohaha/zlsgo/zutil"
)

// ElementRenderer 元素渲染器接口
// 所有可渲染的HTML元素都必须实现此接口
type ElementRenderer interface {
	Render(w io.Writer) error
}

// ElementRendererFunc 元素渲染器函数类型
type ElementRendererFunc func() ElementRenderer

// ElementConfig HTML元素配置
// 用于描述HTML元素的元数据和验证规则
type ElementConfig struct {
	Tag           string            `json:"tag"`
	SelfClosing   bool              `json:"self_closing"`
	AllowChildren bool              `json:"allow_children"`
	Attributes    []AttributeConfig `json:"attributes"`
	Description   string            `json:"description"`
}

// AttributeConfig 属性配置
// 用于描述HTML属性的验证规则
type AttributeConfig struct {
	Name        string   `json:"name"`
	Type        string   `json:"type"` // string, bool, number
	Required    bool     `json:"required"`
	Default     string   `json:"default"`
	Enum        []string `json:"enum"`
	Description string   `json:"description"`
}

// If 条件渲染
// 当 condition 为 true 时，渲染 children
func If(condition bool, children ...ElementRenderer) ElementRenderer {
	if condition {
		return Group(children...)
	}
	return nil
}

// Tern 三元条件渲染
// condition 为 true 返回 trueChildren，否则返回 falseChildren
func Tern(condition bool, trueChildren, falseChildren ElementRenderer) ElementRenderer {
	if condition {
		return trueChildren
	}
	return falseChildren
}

// Range 遍历切片并生成元素
func Range[T any](values []T, cb func(T) ElementRenderer) ElementRenderer {
	children := make([]ElementRenderer, 0, len(values))
	for _, value := range values {
		children = append(children, cb(value))
	}
	return Group(children...)
}

// RangeI 遍历切片（带索引）并生成元素
func RangeI[T any](values []T, cb func(int, T) ElementRenderer) ElementRenderer {
	children := make([]ElementRenderer, 0, len(values))
	for i, value := range values {
		children = append(children, cb(i, value))
	}
	return Group(children...)
}

// DynGroup 动态生成元素组
// 延迟执行 childrenFuncs 来生成子元素
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

// DynIf 动态条件渲染
// 当 condition 为 true 时，延迟执行 childrenFuncs 来生成子元素
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

// DynTern 动态三元条件渲染
// condition 为 true 执行 trueChildren()，否则执行 falseChildren()
func DynTern(condition bool, trueChildren, falseChildren ElementRendererFunc) ElementRenderer {
	if condition {
		return trueChildren()
	}
	return falseChildren()
}

// Error 从错误创建文本元素
func Error(err error) ElementRenderer {
	return Text(err.Error())
}

// RenderToBytes 将元素渲染为字节数组
func RenderToBytes(element ElementRenderer) ([]byte, error) {
	buf := zutil.GetBuff()
	defer zutil.PutBuff(buf)

	if err := element.Render(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
