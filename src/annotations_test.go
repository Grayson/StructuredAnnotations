package structuredannotations

import (
	"testing"
)

func TestAnnotations_Add(t *testing.T) {
	tests := []struct {
		name          string
		initialStyles []Markdown
		args          []Markdown
	}{
		{
			"Annotation Add one",
			[]Markdown{},
			[]Markdown{
				Text("test"),
			},
		},
		{
			"Annotation Add multiple",
			[]Markdown{},
			[]Markdown{
				Text("test"),
				Text("test"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Annotation{
				elements: tt.initialStyles,
			}
			a.Add(tt.args...)
			if len(a.elements) != len(tt.args) {
				t.Error("Unmatched number of styles!")
			}
		})
	}
}

func TestAnnotations_Markdown(t *testing.T) {
	tests := []struct {
		name   string
		styles []Markdown
		want   string
	}{
		{
			"Annotations Basic Markdown",
			[]Markdown{
				Text("test"),
				CreateParagraph(Text("test2"), Emphasis{Text("test3")}),
			},
			"test\n\ntest2 *test3*",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Annotation{
				elements: tt.styles,
			}
			if got := a.Markdown(); got != tt.want {
				t.Errorf("Annotations.Markdown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAnnotation(t *testing.T) {
	tests := []struct {
		name           string
		configurations []AnnotationConfiguration
		want           *Annotation
	}{
		{
			"New empty annotation",
			nil,
			&Annotation{
				context:  "",
				elements: []Markdown{},
				style:    Default,
			},
		},
		{
			"New annotation with context",
			[]AnnotationConfiguration{WithContext("test")},
			&Annotation{
				context:  "test",
				elements: []Markdown{},
				style:    Default,
			},
		},
		{
			"New annotation with style",
			[]AnnotationConfiguration{WithStyle(Info)},
			&Annotation{
				context:  "",
				elements: []Markdown{},
				style:    Info,
			},
		},
		{
			"New annotation with context and style",
			[]AnnotationConfiguration{
				WithContext("test"),
				WithStyle(Info),
			},
			&Annotation{
				context:  "test",
				elements: []Markdown{},
				style:    Info,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnnotation(tt.configurations...); got.context != tt.want.context || got.style != tt.want.style {
				t.Errorf("NewAnnotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_withUniqueContext(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"Test uniqueness of Annotations",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			l, r := Annotation{}, Annotation{}
			WithUniqueContext()(&l)
			WithUniqueContext()(&r)

			if l.context == r.context {
				t.Errorf("withUniqueContext() failed (%v == %v)", l.context, r.context)
			}
		})
	}
}
