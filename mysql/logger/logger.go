package logger

// Logger is used to log critical error messages.
type Logger interface {
	Print(v ...any)
}
