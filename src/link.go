package structuredannotations

import "fmt"

type Link struct {
	Text        InlineStyle
	Destination string
}

func (l *Link) Markdown() string {
	return fmt.Sprintf("[%v](%v)", l.Text.Markdown(), l.Destination)
}

func (Link) canInline() {}
