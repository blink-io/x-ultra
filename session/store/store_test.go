package store

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNilStruct(t *testing.T) {
	ns := NilStruct

	require.Nilf(t, ns, "NilStruct is really nil")
}
