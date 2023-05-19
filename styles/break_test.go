package structuredannotations

import "testing"

func TestBreak_Markdown(t *testing.T) {
	tests := []struct {
		name string
		b    *Break
		want string
	}{
		{
			"Test Break",
			&Break{},
			"***",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Break{}
			if got := b.Markdown(); got != tt.want {
				t.Errorf("Break.Markdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
