package resty

import (
	"testing"

	"github.com/blink-io/x/testdata"
	"github.com/stretchr/testify/require"
)

func TestHTTP3Client_1(t *testing.T) {
	ctls := testdata.CreateClientTLSConfig()
	h3c := HTTP3Client(ctls)
	require.NotNil(t, h3c)
}
