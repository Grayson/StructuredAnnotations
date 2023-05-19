package structuredannotations

type ConvertableToMarkdown interface {
	Markdown() string
}

type Inlineable interface {
	canInline()
}

type InlineStyle interface {
	ConvertableToMarkdown
	Inlineable
}

type Text string

func (t Text) Markdown() string {
	return string(t)
}

func (t Text) canInline() {}
