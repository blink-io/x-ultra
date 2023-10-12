package logging

type logging interface {
	Printf(format string, v ...interface{})
}

var _ logging = (Wrap)(nil)

type Wrap func(format string, v ...interface{})

func (w Wrap) Printf(format string, v ...interface{}) {
	w(format, v...)
}
