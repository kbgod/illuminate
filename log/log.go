package log

type Logger interface {
	Error(err error, message string, fields map[string]any)
	Info(message string, fields map[string]any)
	Debug(message string, fields map[string]any)
}
