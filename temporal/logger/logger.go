package logger

// Logger is an interface that can be passed to ClientOptions.Logger.
type Logger interface {
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(string, ...interface{})
}
