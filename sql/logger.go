package sql

type PrintfLogger func(format string, args ...any)

func (l PrintfLogger) Printf(format string, args ...any) {
	l(format, args)
}
