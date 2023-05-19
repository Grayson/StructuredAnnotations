package structuredannotations

import (
	"testing"
)

func TestAnnotations_Add(t *testing.T) {
	tests := []struct {
		name          string
		initialStyles []ConvertableToMarkdown
		args          []ConvertableToMarkdown
	}{
		{
			"Annotation Add one",
			[]ConvertableToMarkdown{},
			[]ConvertableToMarkdown{
				Text("test"),
			},
		},
		{
			"Annotation Add multiple",
			[]ConvertableToMarkdown{},
			[]ConvertableToMarkdown{
				Text("test"),
				Text("test"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Annotation{
				styles: tt.initialStyles,
			}
			a.Add(tt.args...)
			if len(a.styles) != len(tt.args) {
				t.Error("Unmatched number of styles!")
			}
		})
	}
}

func TestAnnotations_Markdown(t *testing.T) {
	tests := []struct {
		name   string
		styles []ConvertableToMarkdown
		want   string
	}{
		{
			"Annotations Basic Markdown",
			[]ConvertableToMarkdown{
				Text("test"),
				CreateParagraph(Text("test2"), Emphasis{Text("test3")}),
			},
			"test\n\ntest2 *test3*",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Annotation{
				styles: tt.styles,
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
				context: "",
				styles:  []ConvertableToMarkdown{},
				style:   Default,
			},
		},
		{
			"New annotation with context",
			[]AnnotationConfiguration{withContext("test")},
			&Annotation{
				context: "test",
				styles:  []ConvertableToMarkdown{},
				style:   Default,
			},
		},
		{
			"New annotation with style",
			[]AnnotationConfiguration{withStyle(Info)},
			&Annotation{
				context: "",
				styles:  []ConvertableToMarkdown{},
				style:   Info,
			},
		},
		{
			"New annotation with context and style",
			[]AnnotationConfiguration{
				withContext("test"),
				withStyle(Info),
			},
			&Annotation{
				context: "test",
				styles:  []ConvertableToMarkdown{},
				style:   Info,
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
