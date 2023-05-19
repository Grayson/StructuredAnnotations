package structuredannotations

type BlockQuote struct {
	text InlineStyle
}

func (q *BlockQuote) Markdown() string {
	return prependToLines("> ", q.text.Markdown())
}
