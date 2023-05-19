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

func (Text) canInline() {}

func (t Text) ConvertableToMarkdown() string {
	return string(t)
}
