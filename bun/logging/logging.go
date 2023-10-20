package logging

type logging interface {
	Printf(format string, v ...interface{})
}

var _ logging = (Func)(nil)

type Func func(format string, v ...interface{})

func (f Func) Printf(format string, v ...interface{}) {
	f(format, v...)
}
