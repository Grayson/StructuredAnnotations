package structuredannotations

import (
	"fmt"
	"strings"
)

type CodeBlock struct {
	Text string
}

func (c *CodeBlock) Markdown() string {
	return prependToLines("    ", c.Text)
}

func prependToLines(prepend string, input string) string {
	lines := strings.FieldsFunc(input, func(r rune) bool { return r == '\n' })
	var builder strings.Builder

	for _, line := range lines {
		fmt.Fprintf(&builder, "%v%v\n", prepend, line)
	}

	return strings.Trim(builder.String(), "\n")
}
