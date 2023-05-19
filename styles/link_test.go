package structuredannotations

import "testing"

func TestLink_ConvertableToMarkdown(t *testing.T) {
	type fields struct {
		text        string
		destination string
	}
	tests := []struct {
		name  string
		style ConvertableToMarkdown
		want  string
	}{
		{
			"Link test",
			&Link{
				"test",
				"dest",
			},
			"[test](dest)",
		},
		{
			"Link test",
			&Heading{
				&Link{"test", "dest"},
				H1,
			},
			"# [test](dest)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.style
			if got := l.Markdown(); got != tt.want {
				t.Errorf("Link.Markdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
