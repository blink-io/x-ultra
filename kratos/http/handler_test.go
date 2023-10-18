package http

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_1(t *testing.T) {
	hdlr := StdHandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
	hdl := hdlr.Handle
	require.NotNil(t, hdlr)
	require.NotNil(t, hdl)
}
