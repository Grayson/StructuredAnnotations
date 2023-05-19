package structuredannotations

import (
	"fmt"
	"strings"
)

type Annotations struct {
	styles []ConvertableToMarkdown
}

func (a *Annotations) Add(styles ...ConvertableToMarkdown) {
	a.styles = append(a.styles, styles...)
}

func (a *Annotations) Markdown() string {
	var sb strings.Builder

	for _, style := range a.styles {
		fmt.Fprint(&sb, style.Markdown())
		fmt.Fprint(&sb, "\n\n")
	}

	return strings.Trim(sb.String(), "\n")
}
