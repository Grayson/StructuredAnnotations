package structuredannotations

import (
	"fmt"
	"strings"
)

type CodeBlock struct {
	text string
}

func (c *CodeBlock) ConvertableToMarkdown() string {
	lines := strings.FieldsFunc(c.text, func(r rune) bool { return r == '\n' })
	var builder strings.Builder

	for _, line := range lines {
		fmt.Fprintf(&builder, "    %v\n", line)
	}

	return strings.Trim(builder.String(), "\n")
}
