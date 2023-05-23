package structuredannotations

type AnnotationStyle int

const (
	Info AnnotationStyle = iota
	Error
	Warning
	Success
)

func (a AnnotationStyle) String() string {
	switch a {
	case Info:
		return "info"
	case Error:
		return "error"
	case Warning:
		return "warning"
	case Success:
		return "success"
	}
	panic("Unexpected AnnotationStyle")
}
