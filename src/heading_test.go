package structuredannotations

import "testing"

func TestHeading_Markdown(t *testing.T) {
	type fields struct {
		text  string
		level HeadingLevel
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"H1 Test",
			fields{
				"test",
				H1,
			},
			"# test",
		},
		{
			"H6 Test",
			fields{
				"test",
				H6,
			},
			"###### test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heading{
				Text:  Text(tt.fields.text),
				Level: tt.fields.level,
			}
			if got := h.Markdown(); got != tt.want {
				t.Errorf("Heading.Markdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
