package structuredannotations

import "testing"

func TestBlockQuote_Markdown(t *testing.T) {
	type fields struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"BlockQuote single line",
			fields{"test"},
			"> test",
		},
		{
			"BlockQuote multiline",
			fields{"test\ntest2"},
			"> test\n> test2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &BlockQuote{
				Text: Text(tt.fields.text),
			}
			if got := q.Markdown(); got != tt.want {
				t.Errorf("BlockQuote.Markdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
