package i18n

type (
	Logger func(format string, args ...any)
)

func SetLogger(l Logger) {
	if log != nil {
		log = l
	}
}
