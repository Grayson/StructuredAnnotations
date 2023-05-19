package structuredannotations

import "testing"

func TestUnorderedList_Markdown(t *testing.T) {
	tests := []struct {
		name  string
		lines []InlineStyle
		want  string
	}{
		{
			"UnorderedList single item",
			[]InlineStyle{
				Text("test"),
			},
			"* test",
		},
		{
			"UnorderedList multiple items",
			[]InlineStyle{
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
