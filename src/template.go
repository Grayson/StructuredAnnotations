package structuredannotations

import (
	"strings"
	"text/template"

	"github.com/google/uuid"
)

type Template struct {
	t      *template.Template
	values map[string]Markdown
}

func NewTemplate(src string) (*Template, error) {
	t, err := template.New(uuid.NewString()).Parse(src)
	if err != nil {
		return nil, err
	}

	temp := Template{
		t:      t,
		values: map[string]Markdown{},
	}
	return &temp, nil
}

func (t *Template) Add(key string, value Markdown) {
	t.values[key] = value
}

func (t *Template) Markdown() string {
	tmp := map[string]string{}

	for k, v := range t.values {
		tmp[k] = v.Markdown()
	}

	var sb strings.Builder
	t.t.Execute(&sb, tmp)
	return sb.String()
}
