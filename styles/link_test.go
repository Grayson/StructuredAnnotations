package structuredannotations

import "testing"

func TestLink_ConvertableToMarkdown(t *testing.T) {
	type fields struct {
		text        string
		destination string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Link test",
			fields{
				"test",
				"dest",
			},
			"[test](dest)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Link{
				text:        tt.fields.text,
				destination: tt.fields.destination,
			}
			if got := l.ConvertableToMarkdown(); got != tt.want {
				t.Errorf("Link.ConvertableToMarkdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
