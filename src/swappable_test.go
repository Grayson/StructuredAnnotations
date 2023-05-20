package structuredannotations

import (
	"testing"
)

func TestSwappable_Markdown(t *testing.T) {
	tests := []struct {
		name string
		orig Markdown
		swap Markdown
		want string
	}{
		{
			"Simple Swap",
			Text("test"),
			Text("other"),
			"other",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Swappable{
				inner: tt.orig,
			}
			s.Swap(tt.swap)
			if got := s.Markdown(); got != tt.want {
				t.Errorf("Swappable.ToMarkdown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwappable_Delete(t *testing.T) {
	tests := []struct {
		name string
		orig Markdown
		want string
	}{
		{
			"Simple Delete",
			Text("test"),
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Swappable{
				inner: tt.orig,
			}
			s.Delete()
			if got := s.Markdown(); got != tt.want {
				t.Errorf("Swappable.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
