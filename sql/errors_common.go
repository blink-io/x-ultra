package sql

var commonErrHandlers = make(map[error]func(error) *StateError)

func RegisterCommonErrHandler(e error, f func(error) *StateError) {
	commonErrHandlers[e] = f
}
