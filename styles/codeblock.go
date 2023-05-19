package structuredannotations

import (
	"fmt"
	"strings"
)

type CodeBlock struct {
	text string
}

func (c *CodeBlock) ConvertableToMarkdown() string {
	return prependToLines("    ", c.text)
}

func prependToLines(prepend string, input string) string {
	lines := strings.FieldsFunc(input, func(r rune) bool { return r == '\n' })
	var builder strings.Builder

	for _, line := range lines {
		fmt.Fprintf(&builder, "%v%v\n", prepend, line)
	}

	return strings.Trim(builder.String(), "\n")
}
