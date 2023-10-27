package timing

type Option func(*hook)

func Logf(logf func(string, ...any)) Option {
	return func(h *hook) {
		h.logf = logf
	}
}
