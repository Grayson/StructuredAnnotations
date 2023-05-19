package structuredannotations

import "testing"

func TestCodeSpan_Markdown(t *testing.T) {
	tests := []struct {
		name string
		s    CodeSpan
		want string
	}{
		{
			"CodeSpan test",
			CodeSpan("test"),
			"`test`",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Markdown(); got != tt.want {
				t.Errorf("CodeSpan.Markdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
