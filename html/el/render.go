package el

import (
	"fmt"
	"html"
	"io"
)

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
