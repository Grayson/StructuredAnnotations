package structuredannotations

import "testing"

func TestOrderedList_Markdown(t *testing.T) {
	tests := []struct {
		name  string
		lines []InlineStyle
		want  string
	}{
		{
			"OrderedList single item",
			[]InlineStyle{Text("test")},
			"1. test",
		},
		{
			"OrderedList multiple items",
			[]InlineStyle{Text("test"), Text("test2")},
			"1. test\n2. test2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &OrderedList{
				lines: tt.lines,
			}
			if got := l.Markdown(); got != tt.want {
				t.Errorf("OrderedList.Markdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
