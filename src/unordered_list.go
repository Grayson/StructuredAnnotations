package structuredannotations

import (
	"fmt"
	"strings"
)

type UnorderedList struct {
	lines []InlineStyle
}

func CreateUnorderedList(lines ...InlineStyle) *UnorderedList {
	return &UnorderedList{lines}
}

func (l *UnorderedList) Add(line InlineStyle) {
	l.lines = append(l.lines, line)
}

func (l *UnorderedList) Markdown() string {
	var sb strings.Builder

	for _, line := range l.lines {
		fmt.Fprintf(&sb, "* %v\n", line.Markdown())
	}

	return strings.Trim(sb.String(), "\n")
}
