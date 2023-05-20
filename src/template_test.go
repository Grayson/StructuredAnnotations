package structuredannotations

import (
	"testing"
)

func TestTemplate_Markdown(t *testing.T) {
	simple, _ := NewTemplate("# {{.key}}")

	tests := []struct {
		name     string
		template *Template
		values   map[string]Markdown
		want     string
	}{
		{
			"Template simple test",
			simple,
			map[string]Markdown{"key": Text("test")},
			"# test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.values {
				tt.template.Add(k, v)
			}
			if got := tt.template.Markdown(); got != tt.want {
				t.Errorf("Template.Markdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
