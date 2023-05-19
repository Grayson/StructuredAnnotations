package structuredannotations

type BlockQuote struct {
	Text InlineStyle
}

func (q *BlockQuote) Markdown() string {
	return prependToLines("> ", q.Text.Markdown())
}
