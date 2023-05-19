package structuredannotations

import "testing"

func TestEmphasis_Markdown(t *testing.T) {
	tests := []struct {
		name string
		text InlineStyle
		want string
	}{
		{
			"Emphasis simple",
			Text("test"),
			"*test*",
		},
		{
			"Emphasis with additional style",
			&Link{Text("test"), "dest"},
			"*[test](dest)*",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Emphasis{
				text: tt.text,
			}
			if got := e.Markdown(); got != tt.want {
				t.Errorf("Emphasis.Markdown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrongEmphasis_Markdown(t *testing.T) {
	tests := []struct {
		name string
		text InlineStyle
		want string
	}{
		{
			"StrongEmphasis simple",
			Text("test"),
			"**test**",
		},
		{
			"StrongEmphasis with additional style",
			&Link{Text("test"), "dest"},
			"**[test](dest)**",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StrongEmphasis{
				text: tt.text,
			}
			if got := s.Markdown(); got != tt.want {
				t.Errorf("StrongEmphasis.Markdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
