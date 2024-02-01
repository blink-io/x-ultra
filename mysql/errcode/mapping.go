package errcode

// ErrorNum is a number error code.
type ErrorNum int

// Name returns a more human friendly rendering of the error code, namely the
// "condition name".
func (ec ErrorNum) Name() string {
	return errorNumNames[ec]
}

var errorNumNames = map[ErrorNum]string{}
