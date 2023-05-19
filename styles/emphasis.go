package structuredannotations

import "fmt"

type Emphasis struct {
	text InlineStyle
}

func (e Emphasis) Markdown() string {
	return fmt.Sprintf("*%v*", e.text.Markdown())
}

type StrongEmphasis Emphasis

func (s StrongEmphasis) Markdown() string {
	return fmt.Sprintf("**%v**", s.text.Markdown())
}
