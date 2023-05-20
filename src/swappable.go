package structuredannotations

type Swappable struct {
	inner Markdown
}

func CreateSwappable(initial Markdown) *Swappable {
	return &Swappable{
		initial,
	}
}

func (s *Swappable) Swap(other Markdown) {
	s.inner = other
}

func (s *Swappable) Delete() {
	s.inner = nil
}

func (s *Swappable) Markdown() string {
	if s.inner == nil {
		return ""
	}

	return s.inner.Markdown()
}
