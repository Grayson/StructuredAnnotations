package structuredannotations

import "testing"

func TestUnorderedList_Markdown(t *testing.T) {
	tests := []struct {
		name  string
		lines []Markdown
		want  string
	}{
		{
			"UnorderedList single item",
			[]Markdown{
				Text("test"),
			},
			"* test",
		},
		{
			"UnorderedList multiple items",
			[]Markdown{
				Text("test"),
				Text("test2"),
			},
			"* test\n* test2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &UnorderedList{
				lines: tt.lines,
			}
			if got := l.Markdown(); got != tt.want {
				t.Errorf("UnorderedList.Markdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
