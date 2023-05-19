package structuredannotations

import (
	"fmt"
	"strings"
)

type Paragraph struct {
	text []InlineStyle
}

func CreateParagraph(text ...InlineStyle) *Paragraph {
	return &Paragraph{
		text,
	}
}

func (p *Paragraph) Add(style InlineStyle) {
	p.text = append(p.text, style)
}

func (p *Paragraph) Markdown() string {
	var sb strings.Builder
	for _, style := range p.text {
		fmt.Fprint(&sb, style.Markdown())
		fmt.Fprint(&sb, " ")
	}
	return strings.Trim(sb.String(), " ")
}
