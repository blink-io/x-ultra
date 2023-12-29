package http

import (
	"net/http"
	"testing"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/stretchr/testify/require"
)

func TestHandler_1(t *testing.T) {
	hdlr := StdHandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
	hdl := hdlr.Handle
	require.NotNil(t, hdlr)
	require.NotNil(t, hdl)
}

func TestCompat(t *testing.T) {
	var _ RouteRegistrar = (*khttp.Server)(nil)
}
