package el

import (
	"context"
	"fmt"
	"io"
	"sync"
)

type Chunk interface {
	// chunk 是内部标记方法，用于确保只有受支持的类型实现 Chunk 接口。
	chunk()
}

// StaticChunk 表示静态内容片段，会直接写入输出缓冲。
type StaticChunk []byte

// DynamicChunk 表示动态内容片段，在渲染阶段执行回调生成输出。
type DynamicChunk func(context.Context, io.Writer) error

func (StaticChunk) chunk()  {}
func (DynamicChunk) chunk() {}

// RenderChunk 根据 Chunk 的具体类型将内容渲染到 io.Writer 中。
// 若为 DynamicChunk，则执行回调；若为 StaticChunk，则直接写入原始字节。
func RenderChunk(chunk Chunk, ctx context.Context, w io.Writer) error {
	switch chunk := chunk.(type) {
	case DynamicChunk:
		return chunk(ensureContext(ctx), w)
	case StaticChunk:
		_, err := w.Write(chunk)
		return err
	default:
		return fmt.Errorf("invalid chunk type %t", chunk)
	}
}

type ChunkWriter interface {
	// Write 写入一个或多个 Chunk。
	Write(...Chunk)
	// Chunks 返回内部缓存的全部 Chunk 列表。
	Chunks() []Chunk
}

var chunkpool = sync.Pool{
	New: func() any {
		return &[]Chunk{}
	},
}

type cw struct {
	buf *[]Chunk
}

// Write 将新的 Chunk 追加进缓冲区。
func (cw *cw) Write(chunks ...Chunk) {
	if len(chunks) == 0 {
		return
	}
	if cw.buf == nil {
		buf := chunkpool.Get().(*[]Chunk)
		*buf = (*buf)[:0]
		cw.buf = buf
	}
	*cw.buf = append(*cw.buf, chunks...)
}

// Chunks 返回当前缓冲内容，并在返回后自动释放内部缓冲。
func (cw *cw) Chunks() []Chunk {
	if cw.buf == nil {
		return nil
	}

	buf := *cw.buf
	chunks := append([]Chunk(nil), buf...)
	*cw.buf = buf[:0]
	chunkpool.Put(cw.buf)
	cw.buf = nil
	return chunks
}

// Release 将缓冲区放回对象池，避免重复分配。
func (cw *cw) Release() {
	if cw.buf == nil {
		return
	}

	*cw.buf = (*cw.buf)[:0]
	chunkpool.Put(cw.buf)
	cw.buf = nil
}

// NewChunkWriter 创建一个新的 ChunkWriter，并复用对象池中的缓冲。
func NewChunkWriter() ChunkWriter {
	return &cw{}
}

type teecw struct {
	ChunkWriter
	fn func(Chunk) error
}

// Write 先调用回调处理每个 Chunk，再写入基础 ChunkWriter。
func (cw *teecw) Write(chunks ...Chunk) {
	for _, chunk := range chunks {
		cw.fn(chunk)
	}
	cw.ChunkWriter.Write(chunks...)
}
