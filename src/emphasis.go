package structuredannotations

import "fmt"

type Emphasis struct {
	Text InlineStyle
}

func (Emphasis) canInline() {}

func (e Emphasis) Markdown() string {
	return fmt.Sprintf("*%v*", e.Text.Markdown())
}

type StrongEmphasis Emphasis

func (StrongEmphasis) canInline() {}

func (s StrongEmphasis) Markdown() string {
	return fmt.Sprintf("**%v**", s.Text.Markdown())
}
