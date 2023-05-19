package structuredannotations

type BlockQuote struct {
	text InlineStyle
}

func (q *BlockQuote) ConvertableToMarkdown() string {
	return prependToLines("> ", q.text.Markdown())
}
