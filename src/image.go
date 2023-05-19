package structuredannotations

import "fmt"

type Image struct {
	Source      string
	Title       string
	Description string
}

func (i *Image) Markdown() string {
	return fmt.Sprintf("![%v](%v %v)", i.Source, i.Title, i.Description)
}
