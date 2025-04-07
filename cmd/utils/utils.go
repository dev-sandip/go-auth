package utils

import (
	"bytes"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	htmlrenderer "github.com/yuin/goldmark/renderer/html"
)

func MarkdownToHTML(md []byte) ([]byte, error) {

	markdown := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Table,
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
				highlighting.WithFormatOptions(
					html.WithLineNumbers(true),
				),
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			htmlrenderer.WithHardWraps(),
			htmlrenderer.WithXHTML(),
			htmlrenderer.WithUnsafe(),
		),
	)

	var buf bytes.Buffer
	if err := markdown.Convert(md, &buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
