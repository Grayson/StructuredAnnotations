package structuredannotations

import (
	"fmt"
	"strings"

	"github.com/bitrise-io/bitrise-plugins-annotations/service"
	"github.com/google/uuid"
)

type Annotation struct {
	elements []Markdown
	context  string
	style    AnnotationStyle
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

func WithContext(ctx string) AnnotationConfiguration {
	return func(a *Annotation) {
		a.context = ctx
	}
}

func withUniqueContext() AnnotationConfiguration {
	return func(a *Annotation) {
		a.context = uuid.New().String()
	}
}

func WithStyle(style AnnotationStyle) AnnotationConfiguration {
	return func(a *Annotation) {
		a.style = style
	}
}

// Annotation methods

func (a *Annotation) Add(styles ...Markdown) {
	a.elements = append(a.elements, styles...)
}

func (a *Annotation) Send() error {
	annotation := service.Annotation{
		Context:  a.context,
		Markdown: a.Markdown(),
		Style:    a.style.String(),
	}
	return service.Annotate(annotation)
}

func (a *Annotation) Markdown() string {
	var sb strings.Builder

	for _, style := range a.elements {
		fmt.Fprint(&sb, style.Markdown())
		fmt.Fprint(&sb, "\n\n")
	}

	return strings.Trim(sb.String(), "\n")
}
