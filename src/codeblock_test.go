package structuredannotations

import "testing"

func TestCodeBlock_Markdown(t *testing.T) {
	type fields struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"CodeBlock Test",
			fields{"test"},
			"    test",
		},
		{
			"CodeBlock Multiline Test",
			fields{"test\ntest2"},
			"    test\n    test2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CodeBlock{
				Text: tt.fields.text,
			}
			if got := c.Markdown(); got != tt.want {
				t.Errorf("CodeBlock.Markdown() = |%v|, want |%v|", got, tt.want)
			}
		})
	}
}
