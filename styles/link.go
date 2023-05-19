package structuredannotations

import "fmt"

type Link struct {
	text        string
	destination string
}

func (l *Link) ConvertableToMarkdown() string {
	return fmt.Sprintf("[%v](%v)", l.text, l.destination)
}
