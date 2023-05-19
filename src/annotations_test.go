package structuredannotations

import "testing"

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
			a := &Annotations{
				styles: tt.initialStyles,
			}
			a.Add(tt.args...)
			if len(a.styles) != len(tt.args) {
				t.Error("Unmatched number of styles!")
			}
		})
	}
}
