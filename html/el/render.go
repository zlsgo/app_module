package el

import (
	"context"
	"fmt"
	"io"

	"github.com/sohaha/zlsgo/zutil"
)

// RenderNode 渲染单个 Node，并将输出写入 io.Writer。
// 该方法内部会处理 ChunkWriter 以及节点的释放，便于直接输出节点内容。
func RenderNode(ctx context.Context, w io.Writer, node Node) error {
	if node == nil || w == nil {
		return nil
	}

	cw := NewChunkWriter()
	node.Render(cw)
	chunks := cw.Chunks()
	defer node.Release()

	ctx = ensureContext(ctx)
	for _, chunk := range chunks {
		if err := RenderChunk(chunk, ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// Render 依次渲染多个 Item，常用于直接输出顶层节点或片段。
// 仅 Node 类型会被渲染，其他 Item 会返回错误。
func Render(ctx context.Context, w io.Writer, items ...Item) error {
	if len(items) == 0 || w == nil {
		return nil
	}

	nodes := make([]Node, 0, len(items))
	for _, item := range items {
		switch item := item.(type) {
		case nil:
			continue
		case Node:
			nodes = append(nodes, item)
		default:
			return fmt.Errorf("html: unsupported item type %T passed to Render; wrap it in an element before rendering", item)
		}
	}

	if len(nodes) == 0 {
		return nil
	}

	if len(nodes) == 1 {
		return RenderNode(ctx, w, nodes[0])
	}

	return RenderNode(ctx, w, Fragment(nodes...))
}

// RenderBytes 渲染并返回字节结果，便于快速获取输出内容。
func RenderBytes(ctx context.Context, items ...Item) ([]byte, error) {
	buf := zutil.GetBuff()
	defer zutil.PutBuff(buf)

	if err := Render(ctx, buf, items...); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}

// MustRenderBytes 渲染并返回字节结果，并在出现错误时 panic，适合初始化时使用。
func MustRenderBytes(ctx context.Context, items ...Item) []byte {
	result, err := RenderBytes(ctx, items...)
	if err != nil {
		panic(err)
	}
	return result
}

func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return ctx
}
