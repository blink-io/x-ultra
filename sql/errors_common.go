package sql

var commonErrHandlers = make(map[error]func(error) *Error)

func RegisterCommonErrHandler(e error, f func(error) *Error) {
	commonErrHandlers[e] = f
}

func handleCommonError(e error) (func(error) *Error, bool) {
	fn, ok := commonErrHandlers[e]
	return fn, ok
}
