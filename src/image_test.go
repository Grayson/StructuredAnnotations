package structuredannotations

import "testing"

func TestImage_Markdown(t *testing.T) {
	type fields struct {
		source      string
		title       string
		description string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Image simple test",
			fields{
				"src",
				"title",
				"desc",
			},
			"![src](title desc)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Image{
				Source:      tt.fields.source,
				Title:       tt.fields.title,
				Description: tt.fields.description,
			}
			if got := i.Markdown(); got != tt.want {
				t.Errorf("Image.Markdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
