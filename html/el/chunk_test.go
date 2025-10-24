package el

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"

	"github.com/sohaha/zlsgo"
)

func TestChunkWriter(t *testing.T) {
	tt := zlsgo.NewTest(t)

	t.Run("basic write and retrieve", func(t *testing.T) {
		cw := NewChunkWriter()
		cw.Write(StaticChunk("hello"))
		cw.Write(StaticChunk(" world"))

		chunks := cw.Chunks()
		tt.Equal(2, len(chunks))
	})

	t.Run("empty write", func(t *testing.T) {
		cw := NewChunkWriter()
		cw.Write()
		chunks := cw.Chunks()
		tt.Equal(0, len(chunks))
	})

	t.Run("Release method", func(t *testing.T) {
		cw := NewChunkWriter().(*cw)
		cw.Write(StaticChunk("test"))
		cw.Release()
		tt.Equal(true, cw.buf == nil)
	})

	t.Run("Release empty writer", func(t *testing.T) {
		cw := NewChunkWriter().(*cw)
		cw.Release()
		tt.Equal(true, cw.buf == nil)
	})

	t.Run("multiple Chunks calls", func(t *testing.T) {
		cw := NewChunkWriter()
		cw.Write(StaticChunk("first"))
		chunks1 := cw.Chunks()
		tt.Equal(1, len(chunks1))

		chunks2 := cw.Chunks()
		tt.Equal(0, len(chunks2))
	})
}

func TestRenderChunk(t *testing.T) {
	tt := zlsgo.NewTest(t)

	t.Run("static chunk", func(t *testing.T) {
		buf := new(bytes.Buffer)
		chunk := StaticChunk("hello")
		err := RenderChunk(chunk, context.Background(), buf)
		tt.NoError(err)
		tt.Equal("hello", buf.String())
	})

	t.Run("dynamic chunk", func(t *testing.T) {
		buf := new(bytes.Buffer)
		chunk := DynamicChunk(func(ctx context.Context, w io.Writer) error {
			_, err := w.Write([]byte("dynamic"))
			return err
		})
		err := RenderChunk(chunk, context.Background(), buf)
		tt.NoError(err)
		tt.Equal("dynamic", buf.String())
	})

	t.Run("dynamic chunk with error", func(t *testing.T) {
		buf := new(bytes.Buffer)
		chunk := DynamicChunk(func(ctx context.Context, w io.Writer) error {
			return errors.New("test error")
		})
		err := RenderChunk(chunk, context.Background(), buf)
		tt.Equal(true, err != nil)
	})

	t.Run("nil context", func(t *testing.T) {
		buf := new(bytes.Buffer)
		chunk := StaticChunk("test")
		err := RenderChunk(chunk, nil, buf)
		tt.NoError(err)
		tt.Equal("test", buf.String())
	})
}

func TestTeecw(t *testing.T) {
	tt := zlsgo.NewTest(t)

	t.Run("tee writer", func(t *testing.T) {
		var collected []Chunk
		baseCw := NewChunkWriter()
		teeCw := &teecw{
			ChunkWriter: baseCw,
			fn: func(c Chunk) error {
				collected = append(collected, c)
				return nil
			},
		}

		teeCw.Write(StaticChunk("a"), StaticChunk("b"))
		tt.Equal(2, len(collected))

		chunks := baseCw.Chunks()
		tt.Equal(2, len(chunks))
	})
}
