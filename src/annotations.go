package structuredannotations

import (
	"fmt"
	"strings"
)

type Annotation struct {
	styles  []ConvertableToMarkdown
	context string
	style   AnnotationStyle
}

type AnnotationConfiguration func(*Annotation)

func NewAnnotation(configurations ...AnnotationConfiguration) *Annotation {
	annotation := Annotation{}

	for _, conf := range configurations {
		conf(&annotation)
	}

	return &annotation
}

// Annotation configurations

func withContext(ctx string) AnnotationConfiguration {
	return func(a *Annotation) {
		a.context = ctx
	}
}

func withStyle(style AnnotationStyle) AnnotationConfiguration {
	return func(a *Annotation) {
		a.style = style
	}
}

func (a *Annotations) Add(styles ...ConvertableToMarkdown) {
// Annotation methods

func (a *Annotation) Add(styles ...ConvertableToMarkdown) {
	a.styles = append(a.styles, styles...)
}

func (a *Annotation) Markdown() string {
	var sb strings.Builder

	for _, style := range a.styles {
		fmt.Fprint(&sb, style.Markdown())
		fmt.Fprint(&sb, "\n\n")
	}

	return strings.Trim(sb.String(), "\n")
}
