package sql

import (
	"errors"
	"testing"

	"github.com/gocraft/dbr/v2"
	"github.com/stretchr/testify/require"
)

func TestRegister_1(t *testing.T) {
	err1 := dbr.ErrNotFound
	RegisterCommonErrHandler(err1, func(e error) *Error {
		if errors.Is(e, dbr.ErrNotFound) {
			return NewError("dbrErr", "dbrErr", "", e)
		}
		return ErrUnsupported
	})

	var err2 = dbr.ErrNotFound

	var sErr = WrapError(err2)

	var targetErr = NewError("dbrErr", "dbrErr", "", nil)
	require.ErrorIs(t, sErr, targetErr)
}
