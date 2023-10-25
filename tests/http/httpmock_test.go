package http

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestHTTP_Mock_1(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "",
		httpmock.NewStringResponder(200, "abc"))

	httpmock.GetCallCountInfo()
}

func TestHTTP_Server_1(t *testing.T) {
	rr := chi.NewRouter()

	err := http.ListenAndServe("localhost:8181", rr)
	require.NoError(t, err)
}
