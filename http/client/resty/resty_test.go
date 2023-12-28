package resty

import (
	"testing"

	"github.com/blink-io/x/internal/testdata"
	"github.com/stretchr/testify/require"
)

func TestHTTP3Client_1(t *testing.T) {
	ctls := testdata.GetClientTLSConfig()
	h3c := HTTP3Client(ctls)
	require.NotNil(t, h3c)
}
