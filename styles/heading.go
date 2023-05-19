package structuredannotations

import (
	"fmt"
	"strings"
)

type HeadingLevel int

const (
	H1 HeadingLevel = iota
	H2
	H3
	H4
	H5
	H6
)

type Heading struct {
	text  InlineStyle
	level HeadingLevel
}

func (h *Heading) Markdown() string {
	return fmt.Sprintf("%v %v", strings.Repeat("#", int(h.level)+1), h.text.Markdown())
}
