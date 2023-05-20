package structuredannotations

import (
	"fmt"
	"strings"
)

type OrderedList struct {
	lines []Markdown
}

func CreateOrderedList(lines ...Markdown) *OrderedList {
	return &OrderedList{lines}
}

func (l *OrderedList) Add(line Markdown) {
	l.lines = append(l.lines, line)
}

func (l *OrderedList) Markdown() string {
	var sb strings.Builder

	for idx, line := range l.lines {
		fmt.Fprintf(&sb, "%v. %v\n", idx+1, line.Markdown())
	}

	return strings.Trim(sb.String(), "\n")
}
