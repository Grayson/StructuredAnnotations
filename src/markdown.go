package structuredannotations

type Markdown interface {
	Markdown() string
}

type Inlineable interface {
	canInline()
}

type InlineStyle interface {
	Markdown
	Inlineable
}

type Text string

func (t Text) Markdown() string {
	return string(t)
}

func (t Text) canInline() {}
