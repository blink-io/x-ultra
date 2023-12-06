package logging

type Option func(*hook)

func Logf(logf func(format string, args ...any)) Option {
	return func(h *hook) {
		h.logf = logf
	}
}
