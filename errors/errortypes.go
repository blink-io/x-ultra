package errors

// ConstError is a prototype for a certain type of error
type ConstError string

// ConstError implements error
func (e ConstError) Error() string {
	return string(e)
}
