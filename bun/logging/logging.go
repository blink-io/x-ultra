package logging

type logging interface {
	Printf(format string, v ...interface{})
}

var _ logging = (Fn)(nil)

type Fn func(format string, v ...interface{})

func (f Fn) Printf(format string, v ...interface{}) {
	f(format, v...)
}
