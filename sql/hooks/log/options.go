package log

import "log"

var (
	DefaultOptions = newDefault()
)

type Fn func(string, ...interface{})

type Options struct {
	Fn Fn
}

func newDefault() *Options {
	return &Options{
		Fn: log.Printf,
	}
}
func setupOptions(o *Options) *Options {
	if o == nil {
		o = &Options{}
	}
	return o
}
