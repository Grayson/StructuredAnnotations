package structuredannotations

import "fmt"

type Link struct {
	text        InlineStyle
	destination string
}

func (l *Link) Markdown() string {
	return fmt.Sprintf("[%v](%v)", l.text.Markdown(), l.destination)
}

func (Link) canInline() {}
