package gocron

// Logger is the interface that wraps the basic logging methods
// used by gocron. The methods are modeled after the standard
// library slog package. The default logger is a no-op logger.
// To enable logging, use one of the provided New*Logger functions
// or implement your own Logger. The actual level of Log that is logged
// is handled by the implementation.
type Logger interface {
	Debug(msg string, args ...any)
	Error(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
}
