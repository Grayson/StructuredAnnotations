package structuredannotations

import "testing"

func TestLink_Markdown(t *testing.T) {
	type fields struct {
		text        string
		destination string
	}
	tests := []struct {
		name  string
		style Markdown
		want  string
	}{
		{
			"Link test",
			&Link{
				Text("test"),
				"dest",
			},
			"[test](dest)",
		},
		{
			"Link test",
			&Heading{
				&Link{Text("test"), "dest"},
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
