package structuredannotations

import "fmt"

type Link struct {
	text        string
	destination string
}

func (l *Link) Markdown() string {
	return fmt.Sprintf("[%v](%v)", l.text, l.destination)
}

func (Link) canInline() {}
