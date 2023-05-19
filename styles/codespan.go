package structuredannotations

import "fmt"

type CodeSpan Text

func (s CodeSpan) Markdown() string {
	return fmt.Sprintf("`%v`", s)
}
