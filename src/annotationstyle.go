package structuredannotations

type AnnotationStyle int

const (
	Default AnnotationStyle = iota
	Error
	Warning
	Info
)

func (a AnnotationStyle) String() string {
	switch a {
	case Default:
		return "default"
	case Error:
		return "error"
	case Warning:
		return "warning"
	case Info:
		return "info"
	}
	panic("Unexpected AnnotationStyle")
}
