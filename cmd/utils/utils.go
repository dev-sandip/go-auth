package utils

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func MarkdownToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	html := markdown.ToHTML(md, p, nil)
	return html
}
