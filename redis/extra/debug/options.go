package debug

type Option func(*hook)

func Log(l func(string, ...any)) Option {
	return func(h *hook) {
		h.log = l
	}
}
