package structuredannotations

type BlockQuote struct {
	text string
}

func (q *BlockQuote) ConvertableToMarkdown() string {
	return prependToLines("> ", q.text)
}
