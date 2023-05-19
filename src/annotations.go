package structuredannotations

type Annotations struct {
	styles []ConvertableToMarkdown
}

func (a *Annotations) Add(styles ...ConvertableToMarkdown) {
	a.styles = append(a.styles, styles...)
}
