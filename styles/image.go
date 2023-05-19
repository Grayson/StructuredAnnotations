package structuredannotations

import "fmt"

type Image struct {
	source      string
	title       string
	description string
}

func (i *Image) Markdown() string {
	return fmt.Sprintf("![%v](%v %v)", i.source, i.title, i.description)
}
