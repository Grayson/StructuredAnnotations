package structuredannotations

import "testing"

func TestParagraph_Markdown(t *testing.T) {
	tests := []struct {
		name string
		text []InlineStyle
		want string
	}{
		{
			"Paragraph test",
			[]InlineStyle{Text("test"), Text("test2")},
			"test test2",
		},
		{
			"Paragraph style test",
			[]InlineStyle{Text("test"), Emphasis{Text("test2")}, Text("test3")},
			"test *test2* test3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Paragraph{
				text: tt.text,
			}
			if got := p.Markdown(); got != tt.want {
				t.Errorf("Paragraph.Markdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
