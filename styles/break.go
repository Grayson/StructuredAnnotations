package structuredannotations

type Break struct{}

func (b *Break) Markdown() string {
	return "***"
}
